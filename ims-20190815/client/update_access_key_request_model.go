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
	// This parameter is required.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
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
