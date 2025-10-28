// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesInput interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *TagResourcesInput
	GetResourceId() []*string
	SetResourceType(v string) *TagResourcesInput
	GetResourceType() *string
	SetTag(v []*Tag) *TagResourcesInput
	GetTag() []*Tag
}

type TagResourcesInput struct {
	// This parameter is required.
	ResourceId []*string `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	// example:
	//
	// FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*Tag `json:"Tag" xml:"Tag" type:"Repeated"`
}

func (s TagResourcesInput) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesInput) GoString() string {
	return s.String()
}

func (s *TagResourcesInput) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesInput) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesInput) GetTag() []*Tag {
	return s.Tag
}

func (s *TagResourcesInput) SetResourceId(v []*string) *TagResourcesInput {
	s.ResourceId = v
	return s
}

func (s *TagResourcesInput) SetResourceType(v string) *TagResourcesInput {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesInput) SetTag(v []*Tag) *TagResourcesInput {
	s.Tag = v
	return s
}

func (s *TagResourcesInput) Validate() error {
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
