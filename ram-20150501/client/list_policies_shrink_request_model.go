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
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The type of the policies. Valid values: `System` and `Custom`. If you do not specify the parameter, all policies are returned.``
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
