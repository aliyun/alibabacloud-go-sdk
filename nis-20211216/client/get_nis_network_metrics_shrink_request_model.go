// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkMetricsShrinkRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkMetricsShrinkRequest
	GetBeginTime() *int64
	SetDimensionsShrink(v string) *GetNisNetworkMetricsShrinkRequest
	GetDimensionsShrink() *string
	SetEndTime(v int64) *GetNisNetworkMetricsShrinkRequest
	GetEndTime() *int64
	SetMetricName(v string) *GetNisNetworkMetricsShrinkRequest
	GetMetricName() *string
	SetRegionNo(v string) *GetNisNetworkMetricsShrinkRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkMetricsShrinkRequest
	GetResourceType() *string
	SetScanBy(v string) *GetNisNetworkMetricsShrinkRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisNetworkMetricsShrinkRequest
	GetStepMinutes() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkMetricsShrinkRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkMetricsShrinkRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIPV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// TimestampAscending
	ScanBy      *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	StepMinutes *int32  `json:"StepMinutes,omitempty" xml:"StepMinutes,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsShrinkRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkMetricsShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkMetricsShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *GetNisNetworkMetricsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkMetricsShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisNetworkMetricsShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkMetricsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkMetricsShrinkRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisNetworkMetricsShrinkRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisNetworkMetricsShrinkRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkMetricsShrinkRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetBeginTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetDimensionsShrink(v string) *GetNisNetworkMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetEndTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetMetricName(v string) *GetNisNetworkMetricsShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetRegionNo(v string) *GetNisNetworkMetricsShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetResourceType(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetScanBy(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetStepMinutes(v int32) *GetNisNetworkMetricsShrinkRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsShrinkRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
