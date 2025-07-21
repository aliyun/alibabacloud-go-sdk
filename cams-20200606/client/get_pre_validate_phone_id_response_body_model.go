// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreValidatePhoneIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPreValidatePhoneIdResponseBody
	GetCode() *string
	SetData(v *GetPreValidatePhoneIdResponseBodyData) *GetPreValidatePhoneIdResponseBody
	GetData() *GetPreValidatePhoneIdResponseBodyData
	SetMessage(v string) *GetPreValidatePhoneIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPreValidatePhoneIdResponseBody
	GetRequestId() *string
}

type GetPreValidatePhoneIdResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPreValidatePhoneIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPreValidatePhoneIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPreValidatePhoneIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPreValidatePhoneIdResponseBody) GetData() *GetPreValidatePhoneIdResponseBodyData {
	return s.Data
}

func (s *GetPreValidatePhoneIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPreValidatePhoneIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPreValidatePhoneIdResponseBody) SetCode(v string) *GetPreValidatePhoneIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetData(v *GetPreValidatePhoneIdResponseBodyData) *GetPreValidatePhoneIdResponseBody {
	s.Data = v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetMessage(v string) *GetPreValidatePhoneIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) SetRequestId(v string) *GetPreValidatePhoneIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPreValidatePhoneIdResponseBodyData struct {
	// The phone number.
	//
	// example:
	//
	// 929833
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The ID of the phone number.
	//
	// example:
	//
	// 8613800000000
	PhoneNumberId *string `json:"PhoneNumberId,omitempty" xml:"PhoneNumberId,omitempty"`
}

func (s GetPreValidatePhoneIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPreValidatePhoneIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPreValidatePhoneIdResponseBodyData) GetPhoneNumberId() *string {
	return s.PhoneNumberId
}

func (s *GetPreValidatePhoneIdResponseBodyData) SetPhoneNumber(v string) *GetPreValidatePhoneIdResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBodyData) SetPhoneNumberId(v string) *GetPreValidatePhoneIdResponseBodyData {
	s.PhoneNumberId = &v
	return s
}

func (s *GetPreValidatePhoneIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
