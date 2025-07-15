// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *UpdateTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *UpdateTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateTicketResponseBody
	GetRequestId() *string
}

type UpdateTicketResponseBody struct {
	// example:
	//
	// OK
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTicketResponseBody) SetCode(v string) *UpdateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTicketResponseBody) SetData(v interface{}) *UpdateTicketResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTicketResponseBody) SetHttpStatusCode(v int32) *UpdateTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTicketResponseBody) SetMessage(v string) *UpdateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTicketResponseBody) SetParams(v []*string) *UpdateTicketResponseBody {
	s.Params = v
	return s
}

func (s *UpdateTicketResponseBody) SetRequestId(v string) *UpdateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
