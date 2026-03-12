// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectStatuses(v string) *DescribeConfigHistoryRequest
	GetEffectStatuses() *string
	SetEndTime(v int64) *DescribeConfigHistoryRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeConfigHistoryRequest
	GetInstanceId() *string
	SetNeedTotal(v bool) *DescribeConfigHistoryRequest
	GetNeedTotal() *bool
	SetPageNumber(v int32) *DescribeConfigHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConfigHistoryRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeConfigHistoryRequest
	GetStartTime() *int64
}

type DescribeConfigHistoryRequest struct {
	// example:
	//
	// effective
	EffectStatuses *string `json:"EffectStatuses,omitempty" xml:"EffectStatuses,omitempty"`
	// example:
	//
	// 1742178604000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	NeedTotal *bool `json:"NeedTotal,omitempty" xml:"NeedTotal,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1742178604000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeConfigHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryRequest) GetEffectStatuses() *string {
	return s.EffectStatuses
}

func (s *DescribeConfigHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeConfigHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeConfigHistoryRequest) GetNeedTotal() *bool {
	return s.NeedTotal
}

func (s *DescribeConfigHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConfigHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConfigHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeConfigHistoryRequest) SetEffectStatuses(v string) *DescribeConfigHistoryRequest {
	s.EffectStatuses = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetEndTime(v int64) *DescribeConfigHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetInstanceId(v string) *DescribeConfigHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetNeedTotal(v bool) *DescribeConfigHistoryRequest {
	s.NeedTotal = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetPageNumber(v int32) *DescribeConfigHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetPageSize(v int32) *DescribeConfigHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetStartTime(v int64) *DescribeConfigHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeConfigHistoryRequest) Validate() error {
	return dara.Validate(s)
}
