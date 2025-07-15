// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FinishTicketTaskResponseBody
	GetCode() *string
	SetData(v interface{}) *FinishTicketTaskResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *FinishTicketTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FinishTicketTaskResponseBody
	GetMessage() *string
	SetParams(v []*string) *FinishTicketTaskResponseBody
	GetParams() []*string
	SetRequestId(v string) *FinishTicketTaskResponseBody
	GetRequestId() *string
}

type FinishTicketTaskResponseBody struct {
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

func (s FinishTicketTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FinishTicketTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FinishTicketTaskResponseBody) GetData() interface{} {
	return s.Data
}

func (s *FinishTicketTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FinishTicketTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FinishTicketTaskResponseBody) GetParams() []*string {
	return s.Params
}

func (s *FinishTicketTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishTicketTaskResponseBody) SetCode(v string) *FinishTicketTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FinishTicketTaskResponseBody) SetData(v interface{}) *FinishTicketTaskResponseBody {
	s.Data = v
	return s
}

func (s *FinishTicketTaskResponseBody) SetHttpStatusCode(v int32) *FinishTicketTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FinishTicketTaskResponseBody) SetMessage(v string) *FinishTicketTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FinishTicketTaskResponseBody) SetParams(v []*string) *FinishTicketTaskResponseBody {
	s.Params = v
	return s
}

func (s *FinishTicketTaskResponseBody) SetRequestId(v string) *FinishTicketTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishTicketTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
