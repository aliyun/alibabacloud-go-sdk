// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationCode1(v string) *BindMFADeviceRequest
	GetAuthenticationCode1() *string
	SetAuthenticationCode2(v string) *BindMFADeviceRequest
	GetAuthenticationCode2() *string
	SetSerialNumber(v string) *BindMFADeviceRequest
	GetSerialNumber() *string
	SetUserPrincipalName(v string) *BindMFADeviceRequest
	GetUserPrincipalName() *string
}

type BindMFADeviceRequest struct {
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
	SerialNumber        *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// This parameter is required.
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s BindMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *BindMFADeviceRequest) GetAuthenticationCode1() *string {
	return s.AuthenticationCode1
}

func (s *BindMFADeviceRequest) GetAuthenticationCode2() *string {
	return s.AuthenticationCode2
}

func (s *BindMFADeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *BindMFADeviceRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *BindMFADeviceRequest) SetAuthenticationCode1(v string) *BindMFADeviceRequest {
	s.AuthenticationCode1 = &v
	return s
}

func (s *BindMFADeviceRequest) SetAuthenticationCode2(v string) *BindMFADeviceRequest {
	s.AuthenticationCode2 = &v
	return s
}

func (s *BindMFADeviceRequest) SetSerialNumber(v string) *BindMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindMFADeviceRequest) SetUserPrincipalName(v string) *BindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *BindMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
