// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListLogstashRequest
	GetDescription() *string
	SetInstanceId(v string) *ListLogstashRequest
	GetInstanceId() *string
	SetPage(v int32) *ListLogstashRequest
	GetPage() *int32
	SetResourceGroupId(v string) *ListLogstashRequest
	GetResourceGroupId() *string
	SetSize(v int32) *ListLogstashRequest
	GetSize() *int32
	SetTags(v string) *ListLogstashRequest
	GetTags() *string
	SetVersion(v string) *ListLogstashRequest
	GetVersion() *string
}

type ListLogstashRequest struct {
	// rg-acfm2h5vbzd\\*\\*\\*\\*
	//
	// example:
	//
	// ls-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// [{"tagKey":"key1","tagValue":"value1"}]
	//
	// example:
	//
	// ls-cn-n6w1o5jq****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// ls-cn-n6w1o5jq\\*\\*\\*\\*
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// Details of the request header.
	//
	// example:
	//
	// rg-acfm2h5vbzd****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 5.5.3_with_X-Pack
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// [{"tagKey":"key1","tagValue":"value1"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5.5.3_with_X-Pack
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashRequest) GetDescription() *string {
	return s.Description
}

func (s *ListLogstashRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLogstashRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListLogstashRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLogstashRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListLogstashRequest) GetTags() *string {
	return s.Tags
}

func (s *ListLogstashRequest) GetVersion() *string {
	return s.Version
}

func (s *ListLogstashRequest) SetDescription(v string) *ListLogstashRequest {
	s.Description = &v
	return s
}

func (s *ListLogstashRequest) SetInstanceId(v string) *ListLogstashRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashRequest) SetPage(v int32) *ListLogstashRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashRequest) SetResourceGroupId(v string) *ListLogstashRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLogstashRequest) SetSize(v int32) *ListLogstashRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashRequest) SetTags(v string) *ListLogstashRequest {
	s.Tags = &v
	return s
}

func (s *ListLogstashRequest) SetVersion(v string) *ListLogstashRequest {
	s.Version = &v
	return s
}

func (s *ListLogstashRequest) Validate() error {
	return dara.Validate(s)
}
