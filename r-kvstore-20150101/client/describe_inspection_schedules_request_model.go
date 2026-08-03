// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v int64) *DescribeInspectionSchedulesRequest
	GetEnabled() *int64
	SetInstanceId(v string) *DescribeInspectionSchedulesRequest
	GetInstanceId() *string
	SetPageNum(v int64) *DescribeInspectionSchedulesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeInspectionSchedulesRequest
	GetPageSize() *int64
	SetScheduleId(v string) *DescribeInspectionSchedulesRequest
	GetScheduleId() *string
	SetSecurityToken(v string) *DescribeInspectionSchedulesRequest
	GetSecurityToken() *string
}

type DescribeInspectionSchedulesRequest struct {
	// example:
	//
	// 1
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// scheduleId-202604141xxxx
	ScheduleId    *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeInspectionSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionSchedulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionSchedulesRequest) GetEnabled() *int64 {
	return s.Enabled
}

func (s *DescribeInspectionSchedulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInspectionSchedulesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeInspectionSchedulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionSchedulesRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DescribeInspectionSchedulesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInspectionSchedulesRequest) SetEnabled(v int64) *DescribeInspectionSchedulesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) SetInstanceId(v string) *DescribeInspectionSchedulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) SetPageNum(v int64) *DescribeInspectionSchedulesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) SetPageSize(v int64) *DescribeInspectionSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) SetScheduleId(v string) *DescribeInspectionSchedulesRequest {
	s.ScheduleId = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) SetSecurityToken(v string) *DescribeInspectionSchedulesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInspectionSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
