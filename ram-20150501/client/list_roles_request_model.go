// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListRolesRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListRolesRequest
	GetMaxItems() *int32
	SetTag(v []*ListRolesRequestTag) *ListRolesRequest
	GetTag() []*ListRolesRequestTag
}

type ListRolesRequest struct {
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
	// The tags.
	Tag []*ListRolesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListRolesRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListRolesRequest) GetTag() []*ListRolesRequestTag {
	return s.Tag
}

func (s *ListRolesRequest) SetMarker(v string) *ListRolesRequest {
	s.Marker = &v
	return s
}

func (s *ListRolesRequest) SetMaxItems(v int32) *ListRolesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListRolesRequest) SetTag(v []*ListRolesRequestTag) *ListRolesRequest {
	s.Tag = v
	return s
}

func (s *ListRolesRequest) Validate() error {
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

type ListRolesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRolesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequestTag) GoString() string {
	return s.String()
}

func (s *ListRolesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListRolesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListRolesRequestTag) SetKey(v string) *ListRolesRequestTag {
	s.Key = &v
	return s
}

func (s *ListRolesRequestTag) SetValue(v string) *ListRolesRequestTag {
	s.Value = &v
	return s
}

func (s *ListRolesRequestTag) Validate() error {
	return dara.Validate(s)
}
