// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyChatappTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyChatappTemplateResponseBody
	GetCode() *string
	SetData(v *ModifyChatappTemplateResponseBodyData) *ModifyChatappTemplateResponseBody
	GetData() *ModifyChatappTemplateResponseBodyData
	SetMessage(v string) *ModifyChatappTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyChatappTemplateResponseBody
	GetRequestId() *string
}

type ModifyChatappTemplateResponseBody struct {
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
	Data *ModifyChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// NONE
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyChatappTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyChatappTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyChatappTemplateResponseBody) GetData() *ModifyChatappTemplateResponseBodyData {
	return s.Data
}

func (s *ModifyChatappTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyChatappTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *ModifyChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetCode(v string) *ModifyChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetData(v *ModifyChatappTemplateResponseBodyData) *ModifyChatappTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetMessage(v string) *ModifyChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) SetRequestId(v string) *ModifyChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyChatappTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyChatappTemplateResponseBodyData struct {
	// The code of the message template.
	//
	// example:
	//
	// 8472929283883
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyChatappTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifyChatappTemplateResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyChatappTemplateResponseBodyData) SetTemplateCode(v string) *ModifyChatappTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplateResponseBodyData) SetTemplateName(v string) *ModifyChatappTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ModifyChatappTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
