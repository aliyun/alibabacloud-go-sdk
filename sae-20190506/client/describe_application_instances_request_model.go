// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationInstancesRequest
	GetAppId() *string
	SetCurrentPage(v int32) *DescribeApplicationInstancesRequest
	GetCurrentPage() *int32
	SetGroupId(v string) *DescribeApplicationInstancesRequest
	GetGroupId() *string
	SetInstanceId(v string) *DescribeApplicationInstancesRequest
	GetInstanceId() *string
	SetPageSize(v int32) *DescribeApplicationInstancesRequest
	GetPageSize() *int32
	SetPipelineId(v string) *DescribeApplicationInstancesRequest
	GetPipelineId() *string
	SetReverse(v bool) *DescribeApplicationInstancesRequest
	GetReverse() *bool
}

type DescribeApplicationInstancesRequest struct {
	// d700e680-aa4d-4ec1-afc2-6566b5ff\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 1
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// b2a8a925-477a-4ed7-b825-d5e22500\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// demo-faaca66c-5959-45cc-b3bf-d26093b2e9c0******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 85750d48-6cfc-4dbf-8ea0-840397******
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// true
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s DescribeApplicationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationInstancesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApplicationInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApplicationInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationInstancesRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribeApplicationInstancesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *DescribeApplicationInstancesRequest) SetAppId(v string) *DescribeApplicationInstancesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetCurrentPage(v int32) *DescribeApplicationInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetGroupId(v string) *DescribeApplicationInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetInstanceId(v string) *DescribeApplicationInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetPageSize(v int32) *DescribeApplicationInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetPipelineId(v string) *DescribeApplicationInstancesRequest {
	s.PipelineId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetReverse(v bool) *DescribeApplicationInstancesRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) Validate() error {
	return dara.Validate(s)
}
