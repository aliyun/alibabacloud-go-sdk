// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateChatappTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateChatappTemplateResponseBody
	GetCode() *string
	SetData(v *CreateChatappTemplateResponseBodyData) *CreateChatappTemplateResponseBody
	GetData() *CreateChatappTemplateResponseBodyData
	SetMessage(v string) *CreateChatappTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatappTemplateResponseBody
	GetRequestId() *string
}

type CreateChatappTemplateResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {"templateCode": "****4b5c79c9432497a075bdfca36bf5"，"templateName": "hello_whatsapp"}
	Data *CreateChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatappTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateChatappTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatappTemplateResponseBody) GetData() *CreateChatappTemplateResponseBodyData {
	return s.Data
}

func (s *CreateChatappTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatappTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *CreateChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetCode(v string) *CreateChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetData(v *CreateChatappTemplateResponseBodyData) *CreateChatappTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetMessage(v string) *CreateChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetRequestId(v string) *CreateChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateChatappTemplateResponseBodyData struct {
	// The code of the message template.
	//
	// example:
	//
	// SMS_232907****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateChatappTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateChatappTemplateResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateCode(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateName(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *CreateChatappTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
