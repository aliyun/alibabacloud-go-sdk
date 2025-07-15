// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyPhoneNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyPhoneNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPhoneNumberResponseBody
	GetRequestId() *string
}

type ModifyPhoneNumberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyPhoneNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPhoneNumberResponseBody) SetCode(v string) *ModifyPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetHttpStatusCode(v int32) *ModifyPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetMessage(v string) *ModifyPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetRequestId(v string) *ModifyPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
