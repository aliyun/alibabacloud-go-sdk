// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReleaseMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *QueryReleaseMetricRequest
	GetChangeOrderId() *string
	SetCreateTime(v int64) *QueryReleaseMetricRequest
	GetCreateTime() *int64
	SetMetricType(v string) *QueryReleaseMetricRequest
	GetMetricType() *string
	SetPid(v string) *QueryReleaseMetricRequest
	GetPid() *string
	SetProxyUserId(v string) *QueryReleaseMetricRequest
	GetProxyUserId() *string
	SetReleaseEndTime(v int64) *QueryReleaseMetricRequest
	GetReleaseEndTime() *int64
	SetReleaseStartTime(v int64) *QueryReleaseMetricRequest
	GetReleaseStartTime() *int64
	SetService(v string) *QueryReleaseMetricRequest
	GetService() *string
}

type QueryReleaseMetricRequest struct {
	// The ID of the change order.
	//
	// This parameter is required.
	//
	// example:
	//
	// a341a2f2-ed07-4257-aae9-dfb1be******
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The time when the change order was created.
	//
	// example:
	//
	// 1634005438000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the metric that you want to query.
	//
	// example:
	//
	// SystemContrast
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The ID of the Enterprise Distributed Application Service (EDAS) or Kubernetes application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8b46d03f-5947-449d-90fd-3a96c2******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is not in use.
	//
	// example:
	//
	// null
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// The end time of the version release.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1632798718632
	ReleaseEndTime *int64 `json:"ReleaseEndTime,omitempty" xml:"ReleaseEndTime,omitempty"`
	// The start time of the version release.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1632798686692
	ReleaseStartTime *int64 `json:"ReleaseStartTime,omitempty" xml:"ReleaseStartTime,omitempty"`
	// The service that you want to query.
	//
	// example:
	//
	// clothservice
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s QueryReleaseMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReleaseMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *QueryReleaseMetricRequest) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryReleaseMetricRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *QueryReleaseMetricRequest) GetPid() *string {
	return s.Pid
}

func (s *QueryReleaseMetricRequest) GetProxyUserId() *string {
	return s.ProxyUserId
}

func (s *QueryReleaseMetricRequest) GetReleaseEndTime() *int64 {
	return s.ReleaseEndTime
}

func (s *QueryReleaseMetricRequest) GetReleaseStartTime() *int64 {
	return s.ReleaseStartTime
}

func (s *QueryReleaseMetricRequest) GetService() *string {
	return s.Service
}

func (s *QueryReleaseMetricRequest) SetChangeOrderId(v string) *QueryReleaseMetricRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetCreateTime(v int64) *QueryReleaseMetricRequest {
	s.CreateTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetMetricType(v string) *QueryReleaseMetricRequest {
	s.MetricType = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetPid(v string) *QueryReleaseMetricRequest {
	s.Pid = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetProxyUserId(v string) *QueryReleaseMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetReleaseEndTime(v int64) *QueryReleaseMetricRequest {
	s.ReleaseEndTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetReleaseStartTime(v int64) *QueryReleaseMetricRequest {
	s.ReleaseStartTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetService(v string) *QueryReleaseMetricRequest {
	s.Service = &v
	return s
}

func (s *QueryReleaseMetricRequest) Validate() error {
	return dara.Validate(s)
}
