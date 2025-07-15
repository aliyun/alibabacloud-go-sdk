// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RejectTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *RejectTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *RejectTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RejectTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *RejectTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *RejectTicketResponseBody
	GetRequestId() *string
}

type RejectTicketResponseBody struct {
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
	// 678F7002-CA01-4ABF-A112-585AFBDF3A3B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RejectTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *RejectTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *RejectTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RejectTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RejectTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RejectTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectTicketResponseBody) SetCode(v string) *RejectTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RejectTicketResponseBody) SetData(v interface{}) *RejectTicketResponseBody {
	s.Data = v
	return s
}

func (s *RejectTicketResponseBody) SetHttpStatusCode(v int32) *RejectTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RejectTicketResponseBody) SetMessage(v string) *RejectTicketResponseBody {
	s.Message = &v
	return s
}

func (s *RejectTicketResponseBody) SetParams(v []*string) *RejectTicketResponseBody {
	s.Params = v
	return s
}

func (s *RejectTicketResponseBody) SetRequestId(v string) *RejectTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
