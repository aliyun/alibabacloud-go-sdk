// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateUserStatusRequest
	GetDirectoryId() *string
	SetNewStatus(v string) *UpdateUserStatusRequest
	GetNewStatus() *string
	SetUserId(v string) *UpdateUserStatusRequest
	GetUserId() *string
}

type UpdateUserStatusRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new status of the user. Valid values:
	//
	// - Enabled: The logon of the user is enabled.
	//
	// - Disabled: The logon of the user is disabled.
	//
	// example:
	//
	// Disabled
	NewStatus *string `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserStatusRequest) GetNewStatus() *string {
	return s.NewStatus
}

func (s *UpdateUserStatusRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserStatusRequest) SetDirectoryId(v string) *UpdateUserStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserStatusRequest) SetNewStatus(v string) *UpdateUserStatusRequest {
	s.NewStatus = &v
	return s
}

func (s *UpdateUserStatusRequest) SetUserId(v string) *UpdateUserStatusRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserStatusRequest) Validate() error {
	return dara.Validate(s)
}
