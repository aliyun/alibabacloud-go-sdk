// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsInstanceIds(v string) *ListNodesRequest
	GetEcsInstanceIds() *string
	SetEcsInstanceName(v string) *ListNodesRequest
	GetEcsInstanceName() *string
	SetPage(v int32) *ListNodesRequest
	GetPage() *int32
	SetSize(v int32) *ListNodesRequest
	GetSize() *int32
	SetTags(v string) *ListNodesRequest
	GetTags() *string
}

type ListNodesRequest struct {
	// The IDs of the ECS instances.
	//
	// example:
	//
	// i-bp1ei8ysh7orb6eq****
	EcsInstanceIds *string `json:"ecsInstanceIds,omitempty" xml:"ecsInstanceIds,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// test
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The tags of the ECS instance. You must configure tagKey and tagValue.
	//
	// example:
	//
	// [{"tagKey":"abc","tagValue":"xyz"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetEcsInstanceIds() *string {
	return s.EcsInstanceIds
}

func (s *ListNodesRequest) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *ListNodesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListNodesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListNodesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListNodesRequest) SetEcsInstanceIds(v string) *ListNodesRequest {
	s.EcsInstanceIds = &v
	return s
}

func (s *ListNodesRequest) SetEcsInstanceName(v string) *ListNodesRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *ListNodesRequest) SetPage(v int32) *ListNodesRequest {
	s.Page = &v
	return s
}

func (s *ListNodesRequest) SetSize(v int32) *ListNodesRequest {
	s.Size = &v
	return s
}

func (s *ListNodesRequest) SetTags(v string) *ListNodesRequest {
	s.Tags = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
