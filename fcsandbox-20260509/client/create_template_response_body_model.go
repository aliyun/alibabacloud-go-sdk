// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *CreateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
	SetTemplateID(v string) *CreateTemplateResponseBody
	GetTemplateID() *string
}

type CreateTemplateResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The unique identifier of the template.
	//
	// example:
	//
	// tpl-9f3a2b7c8d1e4f5a6b0c7d8e9f1a2b3c
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) GetTemplateID() *string {
	return s.TemplateID
}

func (s *CreateTemplateResponseBody) SetCode(v string) *CreateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateResponseBody) SetMessage(v string) *CreateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateID(v string) *CreateTemplateResponseBody {
	s.TemplateID = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
