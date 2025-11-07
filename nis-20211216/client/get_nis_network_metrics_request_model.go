// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkMetricsRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkMetricsRequest
	GetBeginTime() *int64
	SetDimensions(v []*GetNisNetworkMetricsRequestDimensions) *GetNisNetworkMetricsRequest
	GetDimensions() []*GetNisNetworkMetricsRequestDimensions
	SetEndTime(v int64) *GetNisNetworkMetricsRequest
	GetEndTime() *int64
	SetMetricName(v string) *GetNisNetworkMetricsRequest
	GetMetricName() *string
	SetRegionNo(v string) *GetNisNetworkMetricsRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkMetricsRequest
	GetResourceType() *string
	SetScanBy(v string) *GetNisNetworkMetricsRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisNetworkMetricsRequest
	GetStepMinutes() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkMetricsRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkMetricsRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	Dimensions []*GetNisNetworkMetricsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
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

func (s GetNisNetworkMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkMetricsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkMetricsRequest) GetDimensions() []*GetNisNetworkMetricsRequestDimensions {
	return s.Dimensions
}

func (s *GetNisNetworkMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkMetricsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisNetworkMetricsRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkMetricsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkMetricsRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisNetworkMetricsRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisNetworkMetricsRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkMetricsRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetBeginTime(v int64) *GetNisNetworkMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetDimensions(v []*GetNisNetworkMetricsRequestDimensions) *GetNisNetworkMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetEndTime(v int64) *GetNisNetworkMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetMetricName(v string) *GetNisNetworkMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetRegionNo(v string) *GetNisNetworkMetricsRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetResourceType(v string) *GetNisNetworkMetricsRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetScanBy(v string) *GetNisNetworkMetricsRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetStepMinutes(v int32) *GetNisNetworkMetricsRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisNetworkMetricsRequestDimensions struct {
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// eip-sample*
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkMetricsRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsRequestDimensions) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequestDimensions) GetName() *string {
	return s.Name
}

func (s *GetNisNetworkMetricsRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *GetNisNetworkMetricsRequestDimensions) SetName(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Name = &v
	return s
}

func (s *GetNisNetworkMetricsRequestDimensions) SetValue(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Value = &v
	return s
}

func (s *GetNisNetworkMetricsRequestDimensions) Validate() error {
	return dara.Validate(s)
}
