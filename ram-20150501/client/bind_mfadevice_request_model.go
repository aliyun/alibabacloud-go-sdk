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
	SetUserName(v string) *BindMFADeviceRequest
	GetUserName() *string
}

type BindMFADeviceRequest struct {
	// The first authentication code.
	//
	// example:
	//
	// 11****
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	// The second authentication code.
	//
	// example:
	//
	// 33****
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *BindMFADeviceRequest) GetUserName() *string {
	return s.UserName
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

func (s *BindMFADeviceRequest) SetUserName(v string) *BindMFADeviceRequest {
	s.UserName = &v
	return s
}

func (s *BindMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
