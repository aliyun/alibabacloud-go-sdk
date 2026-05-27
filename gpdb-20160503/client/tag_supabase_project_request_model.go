// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagSupabaseProjectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagSupabaseProjectRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagSupabaseProjectRequest
	GetResourceType() *string
	SetTag(v []*TagSupabaseProjectRequestTag) *TagSupabaseProjectRequest
	GetTag() []*TagSupabaseProjectRequestTag
}

type TagSupabaseProjectRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// example:
	//
	// instance
	ResourceType *string                         `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagSupabaseProjectRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s TagSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *TagSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagSupabaseProjectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagSupabaseProjectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagSupabaseProjectRequest) GetTag() []*TagSupabaseProjectRequestTag {
	return s.Tag
}

func (s *TagSupabaseProjectRequest) SetRegionId(v string) *TagSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *TagSupabaseProjectRequest) SetResourceId(v []*string) *TagSupabaseProjectRequest {
	s.ResourceId = v
	return s
}

func (s *TagSupabaseProjectRequest) SetResourceType(v string) *TagSupabaseProjectRequest {
	s.ResourceType = &v
	return s
}

func (s *TagSupabaseProjectRequest) SetTag(v []*TagSupabaseProjectRequestTag) *TagSupabaseProjectRequest {
	s.Tag = v
	return s
}

func (s *TagSupabaseProjectRequest) Validate() error {
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

type TagSupabaseProjectRequestTag struct {
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagSupabaseProjectRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagSupabaseProjectRequestTag) GoString() string {
	return s.String()
}

func (s *TagSupabaseProjectRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagSupabaseProjectRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagSupabaseProjectRequestTag) SetKey(v string) *TagSupabaseProjectRequestTag {
	s.Key = &v
	return s
}

func (s *TagSupabaseProjectRequestTag) SetValue(v string) *TagSupabaseProjectRequestTag {
	s.Value = &v
	return s
}

func (s *TagSupabaseProjectRequestTag) Validate() error {
	return dara.Validate(s)
}
