// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionScheduleReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInspectionScheduleReportsRequest
	GetInstanceId() *string
	SetPageNum(v int64) *DescribeInspectionScheduleReportsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeInspectionScheduleReportsRequest
	GetPageSize() *int64
	SetScheduleId(v string) *DescribeInspectionScheduleReportsRequest
	GetScheduleId() *string
	SetSecurityToken(v string) *DescribeInspectionScheduleReportsRequest
	GetSecurityToken() *string
}

type DescribeInspectionScheduleReportsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
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

func (s DescribeInspectionScheduleReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionScheduleReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionScheduleReportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInspectionScheduleReportsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeInspectionScheduleReportsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionScheduleReportsRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DescribeInspectionScheduleReportsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInspectionScheduleReportsRequest) SetInstanceId(v string) *DescribeInspectionScheduleReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInspectionScheduleReportsRequest) SetPageNum(v int64) *DescribeInspectionScheduleReportsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionScheduleReportsRequest) SetPageSize(v int64) *DescribeInspectionScheduleReportsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionScheduleReportsRequest) SetScheduleId(v string) *DescribeInspectionScheduleReportsRequest {
	s.ScheduleId = &v
	return s
}

func (s *DescribeInspectionScheduleReportsRequest) SetSecurityToken(v string) *DescribeInspectionScheduleReportsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInspectionScheduleReportsRequest) Validate() error {
	return dara.Validate(s)
}
