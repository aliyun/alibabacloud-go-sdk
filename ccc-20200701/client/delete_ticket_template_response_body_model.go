// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTicketTemplateResponseBody
	GetCode() *string
	SetData(v interface{}) *DeleteTicketTemplateResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *DeleteTicketTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTicketTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteTicketTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteTicketTemplateResponseBody
	GetRequestId() *string
}

type DeleteTicketTemplateResponseBody struct {
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
	// 01B12EE4-6AF2-4730-8B78-EC15F4E5C025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTicketTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTicketTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTicketTemplateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteTicketTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTicketTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTicketTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteTicketTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTicketTemplateResponseBody) SetCode(v string) *DeleteTicketTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTicketTemplateResponseBody) SetData(v interface{}) *DeleteTicketTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTicketTemplateResponseBody) SetHttpStatusCode(v int32) *DeleteTicketTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTicketTemplateResponseBody) SetMessage(v string) *DeleteTicketTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTicketTemplateResponseBody) SetParams(v []*string) *DeleteTicketTemplateResponseBody {
	s.Params = v
	return s
}

func (s *DeleteTicketTemplateResponseBody) SetRequestId(v string) *DeleteTicketTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTicketTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
