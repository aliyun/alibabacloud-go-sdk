// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListRolesShrinkRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListRolesShrinkRequest
	GetMaxItems() *int32
	SetTagShrink(v string) *ListRolesShrinkRequest
	GetTagShrink() *string
}

type ListRolesShrinkRequest struct {
	// The marker used to retrieve the remaining results when the response is truncated.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of entries to return. If the number of entries exceeds this value, the response is truncated and the `IsTruncated` parameter is set to `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRolesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRolesShrinkRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListRolesShrinkRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListRolesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListRolesShrinkRequest) SetMarker(v string) *ListRolesShrinkRequest {
	s.Marker = &v
	return s
}

func (s *ListRolesShrinkRequest) SetMaxItems(v int32) *ListRolesShrinkRequest {
	s.MaxItems = &v
	return s
}

func (s *ListRolesShrinkRequest) SetTagShrink(v string) *ListRolesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListRolesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
