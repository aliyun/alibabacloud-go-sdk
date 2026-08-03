// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionInsId(v string) *DescribeInspectionTaskReportRequest
	GetInspectionInsId() *string
	SetInstanceId(v string) *DescribeInspectionTaskReportRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DescribeInspectionTaskReportRequest
	GetSecurityToken() *string
	SetTaskId(v string) *DescribeInspectionTaskReportRequest
	GetTaskId() *string
}

type DescribeInspectionTaskReportRequest struct {
	// example:
	//
	// r-bp19f4f6994813xxx
	InspectionInsId *string `json:"InspectionInsId,omitempty" xml:"InspectionInsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ta-bp19f4f6994813xxx
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tit-dca42f85c73644e0ab5c80ef64121axxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTaskReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskReportRequest) GetInspectionInsId() *string {
	return s.InspectionInsId
}

func (s *DescribeInspectionTaskReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInspectionTaskReportRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInspectionTaskReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTaskReportRequest) SetInspectionInsId(v string) *DescribeInspectionTaskReportRequest {
	s.InspectionInsId = &v
	return s
}

func (s *DescribeInspectionTaskReportRequest) SetInstanceId(v string) *DescribeInspectionTaskReportRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInspectionTaskReportRequest) SetSecurityToken(v string) *DescribeInspectionTaskReportRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInspectionTaskReportRequest) SetTaskId(v string) *DescribeInspectionTaskReportRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTaskReportRequest) Validate() error {
	return dara.Validate(s)
}
