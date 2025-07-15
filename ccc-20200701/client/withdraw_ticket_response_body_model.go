// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WithdrawTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *WithdrawTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *WithdrawTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *WithdrawTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *WithdrawTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *WithdrawTicketResponseBody
	GetRequestId() *string
}

type WithdrawTicketResponseBody struct {
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
	// 2778FA12-EDD6-42AA-9B15-AF855072E5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawTicketResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *WithdrawTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *WithdrawTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WithdrawTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WithdrawTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *WithdrawTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawTicketResponseBody) SetCode(v string) *WithdrawTicketResponseBody {
	s.Code = &v
	return s
}

func (s *WithdrawTicketResponseBody) SetData(v interface{}) *WithdrawTicketResponseBody {
	s.Data = v
	return s
}

func (s *WithdrawTicketResponseBody) SetHttpStatusCode(v int32) *WithdrawTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WithdrawTicketResponseBody) SetMessage(v string) *WithdrawTicketResponseBody {
	s.Message = &v
	return s
}

func (s *WithdrawTicketResponseBody) SetParams(v []*string) *WithdrawTicketResponseBody {
	s.Params = v
	return s
}

func (s *WithdrawTicketResponseBody) SetRequestId(v string) *WithdrawTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
