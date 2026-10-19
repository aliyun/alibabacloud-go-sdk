// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInspectionTasksRequest
	GetInstanceId() *string
	SetPageNum(v string) *DescribeInspectionTasksRequest
	GetPageNum() *string
	SetPageSize(v string) *DescribeInspectionTasksRequest
	GetPageSize() *string
	SetSecurityToken(v string) *DescribeInspectionTasksRequest
	GetSecurityToken() *string
	SetType(v string) *DescribeInspectionTasksRequest
	GetType() *string
}

type DescribeInspectionTasksRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task type.
	//
	// - 1: manual task
	//
	// - 2: scheduled task
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInspectionTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInspectionTasksRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *DescribeInspectionTasksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInspectionTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInspectionTasksRequest) GetType() *string {
	return s.Type
}

func (s *DescribeInspectionTasksRequest) SetInstanceId(v string) *DescribeInspectionTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInspectionTasksRequest) SetPageNum(v string) *DescribeInspectionTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionTasksRequest) SetPageSize(v string) *DescribeInspectionTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionTasksRequest) SetSecurityToken(v string) *DescribeInspectionTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInspectionTasksRequest) SetType(v string) *DescribeInspectionTasksRequest {
	s.Type = &v
	return s
}

func (s *DescribeInspectionTasksRequest) Validate() error {
	return dara.Validate(s)
}
