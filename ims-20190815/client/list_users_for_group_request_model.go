// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListUsersForGroupRequest
	GetGroupName() *string
	SetMarker(v string) *ListUsersForGroupRequest
	GetMarker() *string
	SetMaxItems(v int32) *ListUsersForGroupRequest
	GetMaxItems() *int32
}

type ListUsersForGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Marker    *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems  *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListUsersForGroupRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersForGroupRequest) GetMaxItems() *int32 {
	return s.MaxItems
}

func (s *ListUsersForGroupRequest) SetGroupName(v string) *ListUsersForGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMarker(v string) *ListUsersForGroupRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMaxItems(v int32) *ListUsersForGroupRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUsersForGroupRequest) Validate() error {
	return dara.Validate(s)
}
