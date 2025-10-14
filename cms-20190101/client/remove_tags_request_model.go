// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *RemoveTagsRequest
	GetGroupIds() []*string
	SetRegionId(v string) *RemoveTagsRequest
	GetRegionId() *string
	SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest
	GetTag() []*RemoveTagsRequestTag
}

type RemoveTagsRequest struct {
	// The IDs of the application groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// This parameter is required.
	Tag []*RemoveTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s RemoveTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *RemoveTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveTagsRequest) GetTag() []*RemoveTagsRequestTag {
	return s.Tag
}

func (s *RemoveTagsRequest) SetGroupIds(v []*string) *RemoveTagsRequest {
	s.GroupIds = v
	return s
}

func (s *RemoveTagsRequest) SetRegionId(v string) *RemoveTagsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveTagsRequest) SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest {
	s.Tag = v
	return s
}

func (s *RemoveTagsRequest) Validate() error {
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

type RemoveTagsRequestTag struct {
	// The tag key.
	//
	// > The tag key (`Tag.N.Key`) and tag value (`Tag.N.Value`) must be specified at the same time.
	//
	// This parameter is required.
	//
	// example:
	//
	// Key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > The tag key (`Tag.N.Key`) and tag value (`Tag.N.Value`) must be specified at the same time.
	//
	// This parameter is required.
	//
	// example:
	//
	// Value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RemoveTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *RemoveTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *RemoveTagsRequestTag) SetKey(v string) *RemoveTagsRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsRequestTag) SetValue(v string) *RemoveTagsRequestTag {
	s.Value = &v
	return s
}

func (s *RemoveTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
