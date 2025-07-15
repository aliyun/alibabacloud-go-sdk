// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemovePhoneNumbersResponseBody
	GetCode() *string
	SetData(v []*string) *RemovePhoneNumbersResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *RemovePhoneNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemovePhoneNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *RemovePhoneNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *RemovePhoneNumbersResponseBody
	GetRequestId() *string
}

type RemovePhoneNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePhoneNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemovePhoneNumbersResponseBody) GetData() []*string {
	return s.Data
}

func (s *RemovePhoneNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemovePhoneNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemovePhoneNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RemovePhoneNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePhoneNumbersResponseBody) SetCode(v string) *RemovePhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetData(v []*string) *RemovePhoneNumbersResponseBody {
	s.Data = v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetHttpStatusCode(v int32) *RemovePhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetMessage(v string) *RemovePhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetParams(v []*string) *RemovePhoneNumbersResponseBody {
	s.Params = v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetRequestId(v string) *RemovePhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
