// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthYears(v int32) *CreateAuthKeyRequest
	GetAuthYears() *int32
	SetBizType(v string) *CreateAuthKeyRequest
	GetBizType() *string
	SetTest(v bool) *CreateAuthKeyRequest
	GetTest() *bool
	SetUserDeviceId(v string) *CreateAuthKeyRequest
	GetUserDeviceId() *string
}

type CreateAuthKeyRequest struct {
	// When the Test flag is false or empty, AuthYears is required, in years, with a range of [1,100]. A value of 100 indicates permanent authorization.
	//
	// example:
	//
	// 1
	AuthYears *int32 `json:"AuthYears,omitempty" xml:"AuthYears,omitempty"`
	// Business type. No more than 64 characters. Can be used to note specific business, such as different face usage scenarios of the access party or the customer identifier to be delivered. It is recommended to pass this parameter.
	//
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Test flag. If true, it indicates using test authorization with a default duration of 30 days; if false, the authorization duration will be based on AuthYears.
	//
	// example:
	//
	// false
	Test *bool `json:"Test,omitempty" xml:"Test,omitempty"`
	// User device ID. No more than 64 characters. Can be used to identify a specific device, and it is suggested to use the physical number of the device. It is recommended to pass this parameter.
	//
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s CreateAuthKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyRequest) GetAuthYears() *int32 {
	return s.AuthYears
}

func (s *CreateAuthKeyRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateAuthKeyRequest) GetTest() *bool {
	return s.Test
}

func (s *CreateAuthKeyRequest) GetUserDeviceId() *string {
	return s.UserDeviceId
}

func (s *CreateAuthKeyRequest) SetAuthYears(v int32) *CreateAuthKeyRequest {
	s.AuthYears = &v
	return s
}

func (s *CreateAuthKeyRequest) SetBizType(v string) *CreateAuthKeyRequest {
	s.BizType = &v
	return s
}

func (s *CreateAuthKeyRequest) SetTest(v bool) *CreateAuthKeyRequest {
	s.Test = &v
	return s
}

func (s *CreateAuthKeyRequest) SetUserDeviceId(v string) *CreateAuthKeyRequest {
	s.UserDeviceId = &v
	return s
}

func (s *CreateAuthKeyRequest) Validate() error {
	return dara.Validate(s)
}
