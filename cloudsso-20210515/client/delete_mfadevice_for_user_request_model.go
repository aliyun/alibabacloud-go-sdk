// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMFADeviceForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DeleteMFADeviceForUserRequest
	GetDirectoryId() *string
	SetMFADeviceId(v string) *DeleteMFADeviceForUserRequest
	GetMFADeviceId() *string
	SetMfaType(v string) *DeleteMFADeviceForUserRequest
	GetMfaType() *string
	SetUserId(v string) *DeleteMFADeviceForUserRequest
	GetUserId() *string
}

type DeleteMFADeviceForUserRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The MFA device ID.
	//
	// You can call [ListMFADevicesForUser](https://help.aliyun.com/document_detail/333531.html) to query the MFA device ID.
	//
	// example:
	//
	// mfa-00ujhet8pycljj7j****
	MFADeviceId *string `json:"MFADeviceId,omitempty" xml:"MFADeviceId,omitempty"`
	MfaType     *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteMFADeviceForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMFADeviceForUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteMFADeviceForUserRequest) GetMFADeviceId() *string {
	return s.MFADeviceId
}

func (s *DeleteMFADeviceForUserRequest) GetMfaType() *string {
	return s.MfaType
}

func (s *DeleteMFADeviceForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteMFADeviceForUserRequest) SetDirectoryId(v string) *DeleteMFADeviceForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) SetMFADeviceId(v string) *DeleteMFADeviceForUserRequest {
	s.MFADeviceId = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) SetMfaType(v string) *DeleteMFADeviceForUserRequest {
	s.MfaType = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) SetUserId(v string) *DeleteMFADeviceForUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMFADeviceForUserRequest) Validate() error {
	return dara.Validate(s)
}
