// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesShrinkRequest
	GetNextToken() *string
	SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *ListTagResourcesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *ListTagResourcesShrinkRequest
	GetTagShrink() *string
}

type ListTagResourcesShrinkRequest struct {
	// The token for the next query.
	//
	// example:
	//
	// caeb235b-xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs. You must specify at least one of ResourceId and Tag. If both are empty, the API returns the InvalidParameter.BothEmpty (400) error.
	//
	// example:
	//
	// ["gw-xxx","gw-yyy"]
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Note: Although the documentation specifies a default value of Gateway, you must explicitly pass this parameter in actual API calls. Otherwise, the API returns the InvalidParameter.UnsupportedTagResourceType (400) error. Valid values: Gateway.
	//
	// example:
	//
	// Gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags. You must specify at least one of ResourceId and Tag. If both are empty, the API returns the InvalidParameter.BothEmpty (400) error.
	//
	// example:
	//
	// [{"key":"env","value":"prod"}]
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
}

func (s *ListTagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
