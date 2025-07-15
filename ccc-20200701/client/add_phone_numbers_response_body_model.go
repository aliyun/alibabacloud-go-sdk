// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddPhoneNumbersResponseBody
	GetCode() *string
	SetData(v []*string) *AddPhoneNumbersResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *AddPhoneNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddPhoneNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPhoneNumbersResponseBody
	GetRequestId() *string
}

type AddPhoneNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s AddPhoneNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddPhoneNumbersResponseBody) GetData() []*string {
	return s.Data
}

func (s *AddPhoneNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddPhoneNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPhoneNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPhoneNumbersResponseBody) SetCode(v string) *AddPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetData(v []*string) *AddPhoneNumbersResponseBody {
	s.Data = v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetHttpStatusCode(v int32) *AddPhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetMessage(v string) *AddPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetRequestId(v string) *AddPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
