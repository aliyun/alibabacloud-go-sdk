// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v []*string) *TagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest
	GetTags() []*TagResourcesRequestTags
}

type TagResourcesRequest struct {
	// The resource IDs.
	//
	// This parameter is required.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type. Set this parameter to INSTANCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag information.
	//
	// This parameter is required.
	Tags []*TagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() []*TagResourcesRequestTags {
	return s.Tags
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TagResourcesRequestTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// IT
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
