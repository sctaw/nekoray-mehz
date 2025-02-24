package speedtest

import (
	"context"
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"time"
)

type Result struct {
	SpeedMB    float64
	Duration   time.Duration
	ServerName string
	ServerAddr string
	Latency    time.Duration
}

func GetDownloadSpeed(ctx context.Context, proxyAddress string) (*Result, error) {
	s := speedtest.New(speedtest.WithUserConfig(&speedtest.UserConfig{Proxy: proxyAddress}))
	user, err := s.FetchUserInfo()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Running speedtest for user with IP: %s and ISP: %s", user.IP, user.Isp)

	servers, err := s.FetchServerListContext(ctx)
	if err != nil {
		return nil, err
	}
	server, err := servers.FindServer(nil)
	if err != nil {
		return nil, err
	}
	_ = server[0].PingTestContext(ctx, nil)
	_ = server[0].DownloadTestContext(ctx)
	return &Result{
		SpeedMB:    server[0].DLSpeed.Mbps(),
		Duration:   server[0].TestDuration.Download.Abs(),
		ServerName: server[0].Name,
		ServerAddr: server[0].Host,
		Latency:    server[0].Latency,
	}, nil
}

func GetUploadSpeed(ctx context.Context, proxyAddress string) (*Result, error) {
	s := speedtest.New(speedtest.WithUserConfig(&speedtest.UserConfig{Proxy: proxyAddress}))
	user, err := s.FetchUserInfo()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Running speedtest for user with IP: %s and ISP: %s", user.IP, user.Isp)

	servers, err := s.FetchServerListContext(ctx)
	if err != nil {
		return nil, err
	}
	server, err := servers.FindServer(nil)
	if err != nil {
		return nil, err
	}
	_ = server[0].PingTestContext(ctx, nil)
	_ = server[0].UploadTestContext(ctx)
	return &Result{
		SpeedMB:    server[0].ULSpeed.Mbps(),
		Duration:   server[0].TestDuration.Upload.Abs(),
		ServerName: server[0].Name,
		ServerAddr: server[0].Host,
		Latency:    server[0].Latency,
	}, nil
}
