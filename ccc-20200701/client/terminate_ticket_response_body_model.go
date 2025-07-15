// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TerminateTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *TerminateTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *TerminateTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TerminateTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *TerminateTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *TerminateTicketResponseBody
	GetRequestId() *string
}

type TerminateTicketResponseBody struct {
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
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *TerminateTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *TerminateTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TerminateTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TerminateTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *TerminateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateTicketResponseBody) SetCode(v string) *TerminateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *TerminateTicketResponseBody) SetData(v interface{}) *TerminateTicketResponseBody {
	s.Data = v
	return s
}

func (s *TerminateTicketResponseBody) SetHttpStatusCode(v int32) *TerminateTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TerminateTicketResponseBody) SetMessage(v string) *TerminateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *TerminateTicketResponseBody) SetParams(v []*string) *TerminateTicketResponseBody {
	s.Params = v
	return s
}

func (s *TerminateTicketResponseBody) SetRequestId(v string) *TerminateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
