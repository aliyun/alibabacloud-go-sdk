// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDeviceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCode(v string) *GenerateDeviceCodeResponseBody
	GetDeviceCode() *string
	SetExpiresAt(v int64) *GenerateDeviceCodeResponseBody
	GetExpiresAt() *int64
	SetExpiresIn(v int64) *GenerateDeviceCodeResponseBody
	GetExpiresIn() *int64
	SetInterval(v int64) *GenerateDeviceCodeResponseBody
	GetInterval() *int64
	SetUserCode(v string) *GenerateDeviceCodeResponseBody
	GetUserCode() *string
	SetVerificationUri(v string) *GenerateDeviceCodeResponseBody
	GetVerificationUri() *string
	SetVerificationUriComplete(v string) *GenerateDeviceCodeResponseBody
	GetVerificationUriComplete() *string
}

type GenerateDeviceCodeResponseBody struct {
	// The device code.
	//
	// example:
	//
	// xxxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// The time when the token expires. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1653288641
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// The remaining validity period of the device code. Unit: seconds.
	//
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The timeout period of the request token. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// The user authorization code.
	//
	// example:
	//
	// xxxxx
	UserCode *string `json:"user_code,omitempty" xml:"user_code,omitempty"`
	// The verification URI.
	//
	// example:
	//
	// https://example.com/authorize/device
	VerificationUri *string `json:"verification_uri,omitempty" xml:"verification_uri,omitempty"`
	// The complete verification URI.
	//
	// example:
	//
	// https://example.com/authorize/device?user_code=
	//
	// xxxx
	VerificationUriComplete *string `json:"verification_uri_complete,omitempty" xml:"verification_uri_complete,omitempty"`
}

func (s GenerateDeviceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDeviceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponseBody) GetDeviceCode() *string {
	return s.DeviceCode
}

func (s *GenerateDeviceCodeResponseBody) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *GenerateDeviceCodeResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GenerateDeviceCodeResponseBody) GetInterval() *int64 {
	return s.Interval
}

func (s *GenerateDeviceCodeResponseBody) GetUserCode() *string {
	return s.UserCode
}

func (s *GenerateDeviceCodeResponseBody) GetVerificationUri() *string {
	return s.VerificationUri
}

func (s *GenerateDeviceCodeResponseBody) GetVerificationUriComplete() *string {
	return s.VerificationUriComplete
}

func (s *GenerateDeviceCodeResponseBody) SetDeviceCode(v string) *GenerateDeviceCodeResponseBody {
	s.DeviceCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresAt(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresIn(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetInterval(v int64) *GenerateDeviceCodeResponseBody {
	s.Interval = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetUserCode(v string) *GenerateDeviceCodeResponseBody {
	s.UserCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUri(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUri = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUriComplete(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUriComplete = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
