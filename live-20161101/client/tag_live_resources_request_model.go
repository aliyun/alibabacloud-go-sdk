// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagLiveResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *TagLiveResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagLiveResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagLiveResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagLiveResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagLiveResourcesRequestTag) *TagLiveResourcesRequest
	GetTag() []*TagLiveResourcesRequestTag
}

type TagLiveResourcesRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resources. Set the value to **DOMAIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagLiveResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagLiveResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagLiveResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagLiveResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagLiveResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagLiveResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagLiveResourcesRequest) GetTag() []*TagLiveResourcesRequestTag {
	return s.Tag
}

func (s *TagLiveResourcesRequest) SetOwnerId(v int64) *TagLiveResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagLiveResourcesRequest) SetRegionId(v string) *TagLiveResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagLiveResourcesRequest) SetResourceId(v []*string) *TagLiveResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagLiveResourcesRequest) SetResourceType(v string) *TagLiveResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagLiveResourcesRequest) SetTag(v []*TagLiveResourcesRequestTag) *TagLiveResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagLiveResourcesRequest) Validate() error {
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

type TagLiveResourcesRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagLiveResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagLiveResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagLiveResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagLiveResourcesRequestTag) SetKey(v string) *TagLiveResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagLiveResourcesRequestTag) SetValue(v string) *TagLiveResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagLiveResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
