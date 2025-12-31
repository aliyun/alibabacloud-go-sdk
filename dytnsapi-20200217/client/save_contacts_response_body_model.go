// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveContactsResponseBody
	GetCode() *string
	SetData(v string) *SaveContactsResponseBody
	GetData() *string
	SetMessage(v string) *SaveContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveContactsResponseBody
	GetRequestId() *string
}

type SaveContactsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 29E058D7-4B28-55EE-BE3B-61D5AE488A9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveContactsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveContactsResponseBody) GetData() *string {
	return s.Data
}

func (s *SaveContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveContactsResponseBody) SetCode(v string) *SaveContactsResponseBody {
	s.Code = &v
	return s
}

func (s *SaveContactsResponseBody) SetData(v string) *SaveContactsResponseBody {
	s.Data = &v
	return s
}

func (s *SaveContactsResponseBody) SetMessage(v string) *SaveContactsResponseBody {
	s.Message = &v
	return s
}

func (s *SaveContactsResponseBody) SetRequestId(v string) *SaveContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
