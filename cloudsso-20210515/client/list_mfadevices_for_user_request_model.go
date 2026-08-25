// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMFADevicesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListMFADevicesForUserRequest
	GetDirectoryId() *string
	SetUserId(v string) *ListMFADevicesForUserRequest
	GetUserId() *string
}

type ListMFADevicesForUserRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMFADevicesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMFADevicesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListMFADevicesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMFADevicesForUserRequest) SetDirectoryId(v string) *ListMFADevicesForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListMFADevicesForUserRequest) SetUserId(v string) *ListMFADevicesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListMFADevicesForUserRequest) Validate() error {
	return dara.Validate(s)
}
