// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTicketResponseBody
	GetCode() *string
	SetData(v string) *CreateTicketResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTicketResponseBody
	GetRequestId() *string
}

type CreateTicketResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 3d26b90a-c5d2-4b09-8219-60cda1******
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A450574A-337F-43E2-BC59-9C6594C994C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTicketResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicketResponseBody) SetCode(v string) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v string) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetHttpStatusCode(v int32) *CreateTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
