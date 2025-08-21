// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVerificationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *GetVerificationInfoRequest
	GetUserPrincipalName() *string
}

type GetVerificationInfoRequest struct {
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetVerificationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetVerificationInfoRequest) SetUserPrincipalName(v string) *GetVerificationInfoRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetVerificationInfoRequest) Validate() error {
	return dara.Validate(s)
}
