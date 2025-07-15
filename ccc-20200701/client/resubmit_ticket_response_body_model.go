// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResubmitTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResubmitTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *ResubmitTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *ResubmitTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResubmitTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *ResubmitTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *ResubmitTicketResponseBody
	GetRequestId() *string
}

type ResubmitTicketResponseBody struct {
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

func (s ResubmitTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResubmitTicketResponseBody) GoString() string {
	return s.String()
}

func (s *ResubmitTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResubmitTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ResubmitTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResubmitTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResubmitTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ResubmitTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResubmitTicketResponseBody) SetCode(v string) *ResubmitTicketResponseBody {
	s.Code = &v
	return s
}

func (s *ResubmitTicketResponseBody) SetData(v interface{}) *ResubmitTicketResponseBody {
	s.Data = v
	return s
}

func (s *ResubmitTicketResponseBody) SetHttpStatusCode(v int32) *ResubmitTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResubmitTicketResponseBody) SetMessage(v string) *ResubmitTicketResponseBody {
	s.Message = &v
	return s
}

func (s *ResubmitTicketResponseBody) SetParams(v []*string) *ResubmitTicketResponseBody {
	s.Params = v
	return s
}

func (s *ResubmitTicketResponseBody) SetRequestId(v string) *ResubmitTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResubmitTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
