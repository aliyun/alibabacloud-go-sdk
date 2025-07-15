// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTicketResponseBody
	GetCode() *string
	SetData(v interface{}) *DeleteTicketResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *DeleteTicketResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTicketResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteTicketResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteTicketResponseBody
	GetRequestId() *string
}

type DeleteTicketResponseBody struct {
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
	// 2263B273-AC1B-44EB-BA98-87F2322C6780
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTicketResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteTicketResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTicketResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTicketResponseBody) SetCode(v string) *DeleteTicketResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTicketResponseBody) SetData(v interface{}) *DeleteTicketResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTicketResponseBody) SetHttpStatusCode(v int32) *DeleteTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTicketResponseBody) SetMessage(v string) *DeleteTicketResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTicketResponseBody) SetParams(v []*string) *DeleteTicketResponseBody {
	s.Params = v
	return s
}

func (s *DeleteTicketResponseBody) SetRequestId(v string) *DeleteTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
