// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetTemplateID(v string) *UpdateTemplateResponseBody
	GetTemplateID() *string
}

type UpdateTemplateResponseBody struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetTemplateID() *string {
	return s.TemplateID
}

func (s *UpdateTemplateResponseBody) SetCode(v string) *UpdateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetMessage(v string) *UpdateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateID(v string) *UpdateTemplateResponseBody {
	s.TemplateID = &v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
