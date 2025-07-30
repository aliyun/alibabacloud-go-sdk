// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricsEnabled(v string) *GetRayDashboardResponseBody
	GetMetricsEnabled() *string
	SetUrl(v string) *GetRayDashboardResponseBody
	GetUrl() *string
}

type GetRayDashboardResponseBody struct {
	// Indicates whether the dashboard has been integrated with CloudMonitor and supports ray metrics
	//
	// example:
	//
	// true
	MetricsEnabled *string `json:"metricsEnabled,omitempty" xml:"metricsEnabled,omitempty"`
	// The Ray Dashboard URL
	//
	// example:
	//
	// https://pre-pai-dlc-proxy-cn-hangzhou.aliyun.com/ray/dashboard/dlc1k7426goc7bvy
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRayDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRayDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayDashboardResponseBody) GetMetricsEnabled() *string {
	return s.MetricsEnabled
}

func (s *GetRayDashboardResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetRayDashboardResponseBody) SetMetricsEnabled(v string) *GetRayDashboardResponseBody {
	s.MetricsEnabled = &v
	return s
}

func (s *GetRayDashboardResponseBody) SetUrl(v string) *GetRayDashboardResponseBody {
	s.Url = &v
	return s
}

func (s *GetRayDashboardResponseBody) Validate() error {
	return dara.Validate(s)
}
