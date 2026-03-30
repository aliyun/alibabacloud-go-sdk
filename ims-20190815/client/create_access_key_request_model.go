// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *CreateAccessKeyRequest
	GetUserPrincipalName() *string
}

type CreateAccessKeyRequest struct {
	// The logon name of the RAM user.
	//
	// If this parameter is empty, an AccessKey pair is created for the current user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateAccessKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateAccessKeyRequest) SetUserPrincipalName(v string) *CreateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateAccessKeyRequest) Validate() error {
	return dara.Validate(s)
}
