// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsInstanceIds(v string) *ListEcsInstancesRequest
	GetEcsInstanceIds() *string
	SetEcsInstanceName(v string) *ListEcsInstancesRequest
	GetEcsInstanceName() *string
	SetPage(v int32) *ListEcsInstancesRequest
	GetPage() *int32
	SetSize(v int32) *ListEcsInstancesRequest
	GetSize() *int32
	SetTags(v string) *ListEcsInstancesRequest
	GetTags() *string
	SetVpcId(v string) *ListEcsInstancesRequest
	GetVpcId() *string
}

type ListEcsInstancesRequest struct {
	// test
	//
	// example:
	//
	// ["i-bp13y63575oypr9d****","i-bp1gyhphjaj73jsr****"]
	EcsInstanceIds *string `json:"ecsInstanceIds,omitempty" xml:"ecsInstanceIds,omitempty"`
	// [{ "tagKey":"a","tagValue":"b"}]
	//
	// example:
	//
	// test
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	// 10
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// ["i-bp13y63575oypr9d\\*\\*\\*\\*","i-bp1gyhphjaj73jsr\\*\\*\\*\\*"]
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// vpc-bp16k1dvzxtmagcva\\*\\*\\*\\*
	//
	// example:
	//
	// [{ "tagKey":"a","tagValue":"b"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEcsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesRequest) GetEcsInstanceIds() *string {
	return s.EcsInstanceIds
}

func (s *ListEcsInstancesRequest) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *ListEcsInstancesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListEcsInstancesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListEcsInstancesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListEcsInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEcsInstancesRequest) SetEcsInstanceIds(v string) *ListEcsInstancesRequest {
	s.EcsInstanceIds = &v
	return s
}

func (s *ListEcsInstancesRequest) SetEcsInstanceName(v string) *ListEcsInstancesRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *ListEcsInstancesRequest) SetPage(v int32) *ListEcsInstancesRequest {
	s.Page = &v
	return s
}

func (s *ListEcsInstancesRequest) SetSize(v int32) *ListEcsInstancesRequest {
	s.Size = &v
	return s
}

func (s *ListEcsInstancesRequest) SetTags(v string) *ListEcsInstancesRequest {
	s.Tags = &v
	return s
}

func (s *ListEcsInstancesRequest) SetVpcId(v string) *ListEcsInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *ListEcsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
