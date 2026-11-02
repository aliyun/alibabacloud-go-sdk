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
	// If the response is truncated, you can use the Marker parameter to retrieve the remaining items.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of entries to return. If the response is truncated because the maximum number of entries is reached, the IsTruncated response parameter is set to true.
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
