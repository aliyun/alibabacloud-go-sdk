// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *DeleteAccessKeyRequest
	GetUserAccessKeyId() *string
	SetUserPrincipalName(v string) *DeleteAccessKeyRequest
	GetUserPrincipalName() *string
}

type DeleteAccessKeyRequest struct {
	// The ID of the AccessKey pair that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is empty, the AccessKey pair of the current user is deleted.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *DeleteAccessKeyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DeleteAccessKeyRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyRequest) SetUserPrincipalName(v string) *DeleteAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeleteAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
