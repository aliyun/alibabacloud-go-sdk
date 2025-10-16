// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstanceId(v string) *DescribeFirewallTaskRequest
	GetChildInstanceId() *string
	SetLang(v string) *DescribeFirewallTaskRequest
	GetLang() *string
	SetTaskId(v string) *DescribeFirewallTaskRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeFirewallTaskRequest
	GetTaskType() *string
}

type DescribeFirewallTaskRequest struct {
	// example:
	//
	// vfw-tr-cd6000c588214403****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 199431783
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// VPC
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeFirewallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTaskRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeFirewallTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFirewallTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeFirewallTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeFirewallTaskRequest) SetChildInstanceId(v string) *DescribeFirewallTaskRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeFirewallTaskRequest) SetLang(v string) *DescribeFirewallTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFirewallTaskRequest) SetTaskId(v string) *DescribeFirewallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeFirewallTaskRequest) SetTaskType(v string) *DescribeFirewallTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeFirewallTaskRequest) Validate() error {
	return dara.Validate(s)
}
