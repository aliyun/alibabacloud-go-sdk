// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeTaskRequest
	GetResourceGroupId() *string
	SetTaskId(v string) *DescribeTaskRequest
	GetTaskId() *string
	SetToken(v string) *DescribeTaskRequest
	GetToken() *string
}

type DescribeTaskRequest struct {
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvywqfey5njq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The access token.
	//
	// example:
	//
	// 01W3ZZOQ
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeTaskRequest) SetResourceGroupId(v string) *DescribeTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskRequest) SetToken(v string) *DescribeTaskRequest {
	s.Token = &v
	return s
}

func (s *DescribeTaskRequest) Validate() error {
	return dara.Validate(s)
}
