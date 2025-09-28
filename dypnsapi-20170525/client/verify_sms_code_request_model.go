// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySmsCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumber(v string) *VerifySmsCodeRequest
	GetPhoneNumber() *string
	SetSmsCode(v string) *VerifySmsCodeRequest
	GetSmsCode() *string
	SetSmsToken(v string) *VerifySmsCodeRequest
	GetSmsToken() *string
}

type VerifySmsCodeRequest struct {
	// The phone number, which is used to receive SMS verification codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1321111****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The SMS verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	SmsCode *string `json:"SmsCode,omitempty" xml:"SmsCode,omitempty"`
	// The text message verification code. After you successfully call the corresponding API operation to send the SMS verification code, the end users receive the SMS verification code. SmsToken is returned by the SDK for SMS verification for you to verify the text message verification code. For an Android client, sendVerifyCode is called to send the verification code. For an iOS client, sendVerifyCodeWithTimeout is called to send the verification code. For more information, see [Overview](https://help.aliyun.com/document_detail/400434.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// sddsbvdbvjd****
	SmsToken *string `json:"SmsToken,omitempty" xml:"SmsToken,omitempty"`
}

func (s VerifySmsCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifySmsCodeRequest) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *VerifySmsCodeRequest) GetSmsCode() *string {
	return s.SmsCode
}

func (s *VerifySmsCodeRequest) GetSmsToken() *string {
	return s.SmsToken
}

func (s *VerifySmsCodeRequest) SetPhoneNumber(v string) *VerifySmsCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsCode(v string) *VerifySmsCodeRequest {
	s.SmsCode = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsToken(v string) *VerifySmsCodeRequest {
	s.SmsToken = &v
	return s
}

func (s *VerifySmsCodeRequest) Validate() error {
	return dara.Validate(s)
}
