// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListUsersRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListUsersRequest
	GetMaxItems() *int32
	SetStatus(v string) *ListUsersRequest
	GetStatus() *string
	SetTag(v []*ListUsersRequestTag) *ListUsersRequest
	GetTag() []*ListUsersRequestTag
}

type ListUsersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 1000. Default value: 1000.
	//
	// example:
	//
	// 1000
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The status of the RAM users that you want to query. Valid values: active, freeze, and active,freeze. If you leave the parameter empty, the value active is used by default. If you specify a value for the Tag parameter, users in both states are queried.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags. A maximum number of 20 tags are supported.
	Tag []*ListUsersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUsersRequest) GetTag() []*ListUsersRequestTag {
	return s.Tag
}

func (s *ListUsersRequest) SetMarker(v string) *ListUsersRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersRequest) SetMaxItems(v int32) *ListUsersRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) SetTag(v []*ListUsersRequestTag) *ListUsersRequest {
	s.Tag = v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}

type ListUsersRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUsersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequestTag) GoString() string {
	return s.String()
}

func (s *ListUsersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListUsersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListUsersRequestTag) SetKey(v string) *ListUsersRequestTag {
	s.Key = &v
	return s
}

func (s *ListUsersRequestTag) SetValue(v string) *ListUsersRequestTag {
	s.Value = &v
	return s
}

func (s *ListUsersRequestTag) Validate() error {
	return dara.Validate(s)
}
