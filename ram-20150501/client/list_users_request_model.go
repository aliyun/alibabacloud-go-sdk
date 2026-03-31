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
}

type ListUsersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of MaxItems, the value of `IsTruncatedg` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
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

func (s *ListUsersRequest) SetMarker(v string) *ListUsersRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersRequest) SetMaxItems(v int32) *ListUsersRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
