// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesSystemTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesSystemTagsRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesSystemTagsRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagResourcesSystemTagsRequest
	GetResourceType() *string
	SetScope(v string) *TagResourcesSystemTagsRequest
	GetScope() *string
	SetTag(v []*TagResourcesSystemTagsRequestTag) *TagResourcesSystemTagsRequest
	GetTag() []*TagResourcesSystemTagsRequestTag
	SetTagOwnerUid(v int64) *TagResourcesSystemTagsRequest
	GetTagOwnerUid() *int64
}

type TagResourcesSystemTagsRequest struct {
	RegionId     *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                           `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope        *string                             `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Tag          []*TagResourcesSystemTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TagOwnerUid  *int64                              `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s TagResourcesSystemTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesSystemTagsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesSystemTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesSystemTagsRequest) GetScope() *string {
	return s.Scope
}

func (s *TagResourcesSystemTagsRequest) GetTag() []*TagResourcesSystemTagsRequestTag {
	return s.Tag
}

func (s *TagResourcesSystemTagsRequest) GetTagOwnerUid() *int64 {
	return s.TagOwnerUid
}

func (s *TagResourcesSystemTagsRequest) SetRegionId(v string) *TagResourcesSystemTagsRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetResourceId(v []*string) *TagResourcesSystemTagsRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetResourceType(v string) *TagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetScope(v string) *TagResourcesSystemTagsRequest {
	s.Scope = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTag(v []*TagResourcesSystemTagsRequestTag) *TagResourcesSystemTagsRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *TagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) Validate() error {
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

type TagResourcesSystemTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesSystemTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesSystemTagsRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesSystemTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesSystemTagsRequestTag) SetKey(v string) *TagResourcesSystemTagsRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesSystemTagsRequestTag) SetValue(v string) *TagResourcesSystemTagsRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesSystemTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
