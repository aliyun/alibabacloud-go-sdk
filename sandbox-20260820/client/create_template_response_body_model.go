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
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
