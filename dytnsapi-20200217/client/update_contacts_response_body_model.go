// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContactsResponseBody
	GetCode() *string
	SetData(v bool) *UpdateContactsResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContactsResponseBody
	GetRequestId() *string
}

type UpdateContactsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContactsResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactsResponseBody) SetCode(v string) *UpdateContactsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContactsResponseBody) SetData(v bool) *UpdateContactsResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateContactsResponseBody) SetMessage(v string) *UpdateContactsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContactsResponseBody) SetRequestId(v string) *UpdateContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
