// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodTagResourcesRequest
	GetOwnerId() *int64
	SetResourceId(v []*string) *DescribeVodTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *DescribeVodTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*DescribeVodTagResourcesRequestTag) *DescribeVodTagResourcesRequest
	GetTag() []*DescribeVodTagResourcesRequestTag
}

type DescribeVodTagResourcesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string                              `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*DescribeVodTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeVodTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeVodTagResourcesRequest) GetTag() []*DescribeVodTagResourcesRequestTag {
	return s.Tag
}

func (s *DescribeVodTagResourcesRequest) SetOwnerId(v int64) *DescribeVodTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetResourceId(v []*string) *DescribeVodTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetResourceType(v string) *DescribeVodTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetTag(v []*DescribeVodTagResourcesRequestTag) *DescribeVodTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *DescribeVodTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVodTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVodTagResourcesRequestTag) SetKey(v string) *DescribeVodTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVodTagResourcesRequestTag) SetValue(v string) *DescribeVodTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVodTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
