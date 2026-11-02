// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListPoliciesShrinkRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListPoliciesShrinkRequest
	GetMaxItems() *int32
	SetPolicyType(v string) *ListPoliciesShrinkRequest
	GetPolicyType() *string
	SetTagShrink(v string) *ListPoliciesShrinkRequest
	GetTagShrink() *string
}

type ListPoliciesShrinkRequest struct {
	// The marker. If the response is truncated, you can use `Marker` to obtain the content that starts from the position after the truncation point.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If the response is truncated because it reaches the `MaxItems` limit, the `IsTruncated` response parameter equals `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The type of the access policy. Valid values: `System` and `Custom`. If this parameter is not specified, all access policies are listed.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListPoliciesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesShrinkRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListPoliciesShrinkRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListPoliciesShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListPoliciesShrinkRequest) SetMarker(v string) *ListPoliciesShrinkRequest {
	s.Marker = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetMaxItems(v int32) *ListPoliciesShrinkRequest {
	s.MaxItems = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetPolicyType(v string) *ListPoliciesShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesShrinkRequest) SetTagShrink(v string) *ListPoliciesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListPoliciesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
