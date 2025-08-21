// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListUsersInRecycleBinRequest
	GetFilter() *string
	SetMarker(v string) *ListUsersInRecycleBinRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListUsersInRecycleBinRequest
	GetMaxItems() *int32
}

type ListUsersInRecycleBinRequest struct {
	// The filter condition that is used to query information about a specified RAM user in the recycle bin.
	//
	// You must specify this parameter in the `UserPrincipalName eq <username>@<AccountAlias>.onaliyun.com` format.
	//
	// example:
	//
	// UserPrincipalName eq test@example.onaliyun.com
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListUsersInRecycleBinRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersInRecycleBinRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListUsersInRecycleBinRequest) SetFilter(v string) *ListUsersInRecycleBinRequest {
	s.Filter = &v
	return s
}

func (s *ListUsersInRecycleBinRequest) SetMarker(v string) *ListUsersInRecycleBinRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersInRecycleBinRequest) SetMaxItems(v int32) *ListUsersInRecycleBinRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUsersInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
