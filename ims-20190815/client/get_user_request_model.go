// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *GetUserRequest
	GetUserAccessKeyId() *string
	SetUserId(v string) *GetUserRequest
	GetUserId() *string
	SetUserPrincipalName(v string) *GetUserRequest
	GetUserPrincipalName() *string
}

type GetUserRequest struct {
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetUserRequest) SetUserAccessKeyId(v string) *GetUserRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) SetUserPrincipalName(v string) *GetUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
