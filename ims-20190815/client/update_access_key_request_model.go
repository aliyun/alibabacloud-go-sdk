// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateAccessKeyRequest
	GetStatus() *string
	SetUserAccessKeyId(v string) *UpdateAccessKeyRequest
	GetUserAccessKeyId() *string
	SetUserPrincipalName(v string) *UpdateAccessKeyRequest
	GetUserPrincipalName() *string
}

type UpdateAccessKeyRequest struct {
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// This parameter is required.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The AccessKey ID of the AccessKey pair for which you want to modify the status.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is empty, the status of the AccessKey pair for the current user is modified.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAccessKeyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *UpdateAccessKeyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UpdateAccessKeyRequest) SetStatus(v string) *UpdateAccessKeyRequest {
	s.Status = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserAccessKeyId(v string) *UpdateAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserPrincipalName(v string) *UpdateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
