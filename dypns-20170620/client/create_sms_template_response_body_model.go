// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSmsTemplateResponseBody
	GetCode() *string
	SetData(v string) *CreateSmsTemplateResponseBody
	GetData() *string
	SetMessage(v string) *CreateSmsTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSmsTemplateResponseBody
	GetRequestId() *string
}

type CreateSmsTemplateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsTemplateResponseBody) SetCode(v string) *CreateSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetData(v string) *CreateSmsTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetMessage(v string) *CreateSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetRequestId(v string) *CreateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
