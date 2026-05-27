// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSupabaseProjectTagsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListSupabaseProjectTagsRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListSupabaseProjectTagsRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListSupabaseProjectTagsRequest
	GetResourceType() *string
	SetTag(v []*ListSupabaseProjectTagsRequestTag) *ListSupabaseProjectTagsRequest
	GetTag() []*ListSupabaseProjectTagsRequestTag
}

type ListSupabaseProjectTagsRequest struct {
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// example:
	//
	// instance
	ResourceType *string                              `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListSupabaseProjectTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSupabaseProjectTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectTagsRequest) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseProjectTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupabaseProjectTagsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListSupabaseProjectTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSupabaseProjectTagsRequest) GetTag() []*ListSupabaseProjectTagsRequestTag {
	return s.Tag
}

func (s *ListSupabaseProjectTagsRequest) SetNextToken(v string) *ListSupabaseProjectTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseProjectTagsRequest) SetRegionId(v string) *ListSupabaseProjectTagsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupabaseProjectTagsRequest) SetResourceId(v []*string) *ListSupabaseProjectTagsRequest {
	s.ResourceId = v
	return s
}

func (s *ListSupabaseProjectTagsRequest) SetResourceType(v string) *ListSupabaseProjectTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSupabaseProjectTagsRequest) SetTag(v []*ListSupabaseProjectTagsRequestTag) *ListSupabaseProjectTagsRequest {
	s.Tag = v
	return s
}

func (s *ListSupabaseProjectTagsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupabaseProjectTagsRequestTag struct {
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSupabaseProjectTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectTagsRequestTag) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListSupabaseProjectTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListSupabaseProjectTagsRequestTag) SetKey(v string) *ListSupabaseProjectTagsRequestTag {
	s.Key = &v
	return s
}

func (s *ListSupabaseProjectTagsRequestTag) SetValue(v string) *ListSupabaseProjectTagsRequestTag {
	s.Value = &v
	return s
}

func (s *ListSupabaseProjectTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
