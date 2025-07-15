// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTicketTaskResponseBody
	GetCode() *string
	SetData(v interface{}) *AddTicketTaskResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *AddTicketTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddTicketTaskResponseBody
	GetMessage() *string
	SetParams(v []*string) *AddTicketTaskResponseBody
	GetParams() []*string
	SetRequestId(v string) *AddTicketTaskResponseBody
	GetRequestId() *string
}

type AddTicketTaskResponseBody struct {
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
	// B06B3244-1B44-481B-90C4-F2F92E59D6B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTicketTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTicketTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddTicketTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTicketTaskResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AddTicketTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddTicketTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTicketTaskResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AddTicketTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTicketTaskResponseBody) SetCode(v string) *AddTicketTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AddTicketTaskResponseBody) SetData(v interface{}) *AddTicketTaskResponseBody {
	s.Data = v
	return s
}

func (s *AddTicketTaskResponseBody) SetHttpStatusCode(v int32) *AddTicketTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTicketTaskResponseBody) SetMessage(v string) *AddTicketTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AddTicketTaskResponseBody) SetParams(v []*string) *AddTicketTaskResponseBody {
	s.Params = v
	return s
}

func (s *AddTicketTaskResponseBody) SetRequestId(v string) *AddTicketTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTicketTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
