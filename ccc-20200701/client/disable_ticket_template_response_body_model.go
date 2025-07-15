// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTicketTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableTicketTemplateResponseBody
	GetCode() *string
	SetData(v interface{}) *DisableTicketTemplateResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *DisableTicketTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableTicketTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *DisableTicketTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *DisableTicketTemplateResponseBody
	GetRequestId() *string
}

type DisableTicketTemplateResponseBody struct {
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
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableTicketTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableTicketTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DisableTicketTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableTicketTemplateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DisableTicketTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableTicketTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableTicketTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DisableTicketTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableTicketTemplateResponseBody) SetCode(v string) *DisableTicketTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DisableTicketTemplateResponseBody) SetData(v interface{}) *DisableTicketTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DisableTicketTemplateResponseBody) SetHttpStatusCode(v int32) *DisableTicketTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableTicketTemplateResponseBody) SetMessage(v string) *DisableTicketTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DisableTicketTemplateResponseBody) SetParams(v []*string) *DisableTicketTemplateResponseBody {
	s.Params = v
	return s
}

func (s *DisableTicketTemplateResponseBody) SetRequestId(v string) *DisableTicketTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableTicketTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
