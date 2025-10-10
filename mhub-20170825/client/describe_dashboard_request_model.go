// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeDashboardRequest
	GetAppKey() *string
	SetAppType(v int32) *DescribeDashboardRequest
	GetAppType() *int32
	SetAppVersion(v string) *DescribeDashboardRequest
	GetAppVersion() *string
	SetEndTime(v int64) *DescribeDashboardRequest
	GetEndTime() *int64
	SetKeyword(v string) *DescribeDashboardRequest
	GetKeyword() *string
	SetProxyAction(v string) *DescribeDashboardRequest
	GetProxyAction() *string
	SetServiceName(v string) *DescribeDashboardRequest
	GetServiceName() *string
	SetStartTime(v int64) *DescribeDashboardRequest
	GetStartTime() *int64
}

type DescribeDashboardRequest struct {
	// example:
	//
	// 29201799
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 4.8
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1681985390256
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 4.8
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// queryAppVersions
	ProxyAction *string `json:"ProxyAction,omitempty" xml:"ProxyAction,omitempty"`
	// example:
	//
	// mqc
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// 1681369984564
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeDashboardRequest) GetAppType() *int32 {
	return s.AppType
}

func (s *DescribeDashboardRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeDashboardRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDashboardRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDashboardRequest) GetProxyAction() *string {
	return s.ProxyAction
}

func (s *DescribeDashboardRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeDashboardRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDashboardRequest) SetAppKey(v string) *DescribeDashboardRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeDashboardRequest) SetAppType(v int32) *DescribeDashboardRequest {
	s.AppType = &v
	return s
}

func (s *DescribeDashboardRequest) SetAppVersion(v string) *DescribeDashboardRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeDashboardRequest) SetEndTime(v int64) *DescribeDashboardRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDashboardRequest) SetKeyword(v string) *DescribeDashboardRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDashboardRequest) SetProxyAction(v string) *DescribeDashboardRequest {
	s.ProxyAction = &v
	return s
}

func (s *DescribeDashboardRequest) SetServiceName(v string) *DescribeDashboardRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeDashboardRequest) SetStartTime(v int64) *DescribeDashboardRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDashboardRequest) Validate() error {
	return dara.Validate(s)
}
