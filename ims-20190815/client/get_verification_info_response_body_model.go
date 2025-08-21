// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVerificationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVerificationInfoResponseBody
	GetRequestId() *string
	SetSecurityEmailDevice(v *GetVerificationInfoResponseBodySecurityEmailDevice) *GetVerificationInfoResponseBody
	GetSecurityEmailDevice() *GetVerificationInfoResponseBodySecurityEmailDevice
	SetSecurityPhoneDevice(v *GetVerificationInfoResponseBodySecurityPhoneDevice) *GetVerificationInfoResponseBody
	GetSecurityPhoneDevice() *GetVerificationInfoResponseBodySecurityPhoneDevice
}

type GetVerificationInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the email.
	SecurityEmailDevice *GetVerificationInfoResponseBodySecurityEmailDevice `json:"SecurityEmailDevice,omitempty" xml:"SecurityEmailDevice,omitempty" type:"Struct"`
	// The information about the mobile phone.
	SecurityPhoneDevice *GetVerificationInfoResponseBodySecurityPhoneDevice `json:"SecurityPhoneDevice,omitempty" xml:"SecurityPhoneDevice,omitempty" type:"Struct"`
}

func (s GetVerificationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVerificationInfoResponseBody) GetSecurityEmailDevice() *GetVerificationInfoResponseBodySecurityEmailDevice {
	return s.SecurityEmailDevice
}

func (s *GetVerificationInfoResponseBody) GetSecurityPhoneDevice() *GetVerificationInfoResponseBodySecurityPhoneDevice {
	return s.SecurityPhoneDevice
}

func (s *GetVerificationInfoResponseBody) SetRequestId(v string) *GetVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVerificationInfoResponseBody) SetSecurityEmailDevice(v *GetVerificationInfoResponseBodySecurityEmailDevice) *GetVerificationInfoResponseBody {
	s.SecurityEmailDevice = v
	return s
}

func (s *GetVerificationInfoResponseBody) SetSecurityPhoneDevice(v *GetVerificationInfoResponseBodySecurityPhoneDevice) *GetVerificationInfoResponseBody {
	s.SecurityPhoneDevice = v
	return s
}

func (s *GetVerificationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVerificationInfoResponseBodySecurityEmailDevice struct {
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The status of the email. Valid values:
	//
	// 	- active: The email is activated.
	//
	// 	- pending: The email is pending for activation.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVerificationInfoResponseBodySecurityEmailDevice) String() string {
	return dara.Prettify(s)
}

func (s GetVerificationInfoResponseBodySecurityEmailDevice) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) GetEmail() *string {
	return s.Email
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) GetStatus() *string {
	return s.Status
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) SetEmail(v string) *GetVerificationInfoResponseBodySecurityEmailDevice {
	s.Email = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) SetStatus(v string) *GetVerificationInfoResponseBodySecurityEmailDevice {
	s.Status = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) Validate() error {
	return dara.Validate(s)
}

type GetVerificationInfoResponseBodySecurityPhoneDevice struct {
	// The international dialing code.
	//
	// example:
	//
	// 86
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 13900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The status of the mobile phone. Valid values:
	//
	// 	- active: The mobile phone is activated.
	//
	// 	- pending: The mobile phone is pending for activation.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVerificationInfoResponseBodySecurityPhoneDevice) String() string {
	return dara.Prettify(s)
}

func (s GetVerificationInfoResponseBodySecurityPhoneDevice) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) GetAreaCode() *string {
	return s.AreaCode
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) GetStatus() *string {
	return s.Status
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetAreaCode(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.AreaCode = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetPhoneNumber(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.PhoneNumber = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetStatus(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.Status = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) Validate() error {
	return dara.Validate(s)
}
