// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *GetUserMFAInfoRequest
	GetUserPrincipalName() *string
}

type GetUserMFAInfoRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetUserMFAInfoRequest) SetUserPrincipalName(v string) *GetUserMFAInfoRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserMFAInfoRequest) Validate() error {
	return dara.Validate(s)
}
