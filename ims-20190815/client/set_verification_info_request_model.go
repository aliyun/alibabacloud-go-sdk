// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVerificationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *SetVerificationInfoRequest
	GetEmail() *string
	SetMobilePhone(v string) *SetVerificationInfoRequest
	GetMobilePhone() *string
	SetUserPrincipalName(v string) *SetVerificationInfoRequest
	GetUserPrincipalName() *string
	SetVerifyType(v string) *SetVerificationInfoRequest
	GetVerifyType() *string
}

type SetVerificationInfoRequest struct {
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	VerifyType        *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SetVerificationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoRequest) GetEmail() *string {
	return s.Email
}

func (s *SetVerificationInfoRequest) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *SetVerificationInfoRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *SetVerificationInfoRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *SetVerificationInfoRequest) SetEmail(v string) *SetVerificationInfoRequest {
	s.Email = &v
	return s
}

func (s *SetVerificationInfoRequest) SetMobilePhone(v string) *SetVerificationInfoRequest {
	s.MobilePhone = &v
	return s
}

func (s *SetVerificationInfoRequest) SetUserPrincipalName(v string) *SetVerificationInfoRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *SetVerificationInfoRequest) SetVerifyType(v string) *SetVerificationInfoRequest {
	s.VerifyType = &v
	return s
}

func (s *SetVerificationInfoRequest) Validate() error {
	return dara.Validate(s)
}
