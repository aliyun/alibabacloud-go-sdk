// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TransferTicketTaskResponseBody
	GetCode() *string
	SetData(v interface{}) *TransferTicketTaskResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *TransferTicketTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TransferTicketTaskResponseBody
	GetMessage() *string
	SetParams(v []*string) *TransferTicketTaskResponseBody
	GetParams() []*string
	SetRequestId(v string) *TransferTicketTaskResponseBody
	GetRequestId() *string
}

type TransferTicketTaskResponseBody struct {
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

func (s TransferTicketTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TransferTicketTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferTicketTaskResponseBody) GetData() interface{} {
	return s.Data
}

func (s *TransferTicketTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TransferTicketTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TransferTicketTaskResponseBody) GetParams() []*string {
	return s.Params
}

func (s *TransferTicketTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferTicketTaskResponseBody) SetCode(v string) *TransferTicketTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TransferTicketTaskResponseBody) SetData(v interface{}) *TransferTicketTaskResponseBody {
	s.Data = v
	return s
}

func (s *TransferTicketTaskResponseBody) SetHttpStatusCode(v int32) *TransferTicketTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferTicketTaskResponseBody) SetMessage(v string) *TransferTicketTaskResponseBody {
	s.Message = &v
	return s
}

func (s *TransferTicketTaskResponseBody) SetParams(v []*string) *TransferTicketTaskResponseBody {
	s.Params = v
	return s
}

func (s *TransferTicketTaskResponseBody) SetRequestId(v string) *TransferTicketTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferTicketTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
