// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagVodResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *TagVodResourcesRequest
	GetOwnerId() *int64
	SetResourceId(v []*string) *TagVodResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagVodResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagVodResourcesRequestTag) *TagVodResourcesRequest
	GetTag() []*TagVodResourcesRequestTag
}

type TagVodResourcesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagVodResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagVodResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagVodResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagVodResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagVodResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagVodResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagVodResourcesRequest) GetTag() []*TagVodResourcesRequestTag {
	return s.Tag
}

func (s *TagVodResourcesRequest) SetOwnerId(v int64) *TagVodResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagVodResourcesRequest) SetResourceId(v []*string) *TagVodResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagVodResourcesRequest) SetResourceType(v string) *TagVodResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagVodResourcesRequest) SetTag(v []*TagVodResourcesRequestTag) *TagVodResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagVodResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type TagVodResourcesRequestTag struct {
	// This parameter is required.
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagVodResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagVodResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagVodResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagVodResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagVodResourcesRequestTag) SetKey(v string) *TagVodResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagVodResourcesRequestTag) SetValue(v string) *TagVodResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagVodResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
