// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatappTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatappTemplateResponseBody
	GetCode() *string
	SetListTemplate(v []*ListChatappTemplateResponseBodyListTemplate) *ListChatappTemplateResponseBody
	GetListTemplate() []*ListChatappTemplateResponseBodyListTemplate
	SetMessage(v string) *ListChatappTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatappTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatappTemplateResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListChatappTemplateResponseBody
	GetTotal() *int32
}

type ListChatappTemplateResponseBody struct {
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
	// The message templates.
	ListTemplate []*ListChatappTemplateResponseBodyListTemplate `json:"ListTemplate,omitempty" xml:"ListTemplate,omitempty" type:"Repeated"`
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
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChatappTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatappTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatappTemplateResponseBody) GetListTemplate() []*ListChatappTemplateResponseBodyListTemplate {
	return s.ListTemplate
}

func (s *ListChatappTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatappTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatappTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatappTemplateResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListChatappTemplateResponseBody) SetAccessDeniedDetail(v string) *ListChatappTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetCode(v string) *ListChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetListTemplate(v []*ListChatappTemplateResponseBodyListTemplate) *ListChatappTemplateResponseBody {
	s.ListTemplate = v
	return s
}

func (s *ListChatappTemplateResponseBody) SetMessage(v string) *ListChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetRequestId(v string) *ListChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetSuccess(v bool) *ListChatappTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetTotal(v int32) *ListChatappTemplateResponseBody {
	s.Total = &v
	return s
}

func (s *ListChatappTemplateResponseBody) Validate() error {
	if s.ListTemplate != nil {
		for _, item := range s.ListTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatappTemplateResponseBodyListTemplate struct {
	// The review state of the template. Valid values:
	//
	// 	- **pass**: The template is approved.
	//
	// 	- **fail**: The template is rejected.
	//
	// 	- **auditing**: The template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the WhatsApp message template. Valid values:
	//
	// 	- **UTILITY**
	//
	// 	- **MARKETING**
	//
	// 	- **AUTHENTICATION**
	//
	// The category of the Viber template. Valid values:
	//
	// 	- **text**: template that contains only text
	//
	// 	- **image**: template that contains only images
	//
	// 	- **text_image_button**: template that contains text, images, and buttons
	//
	// 	- **text_button**: template that contains text and buttons
	//
	// 	- **document**: template that contains only documents
	//
	// 	- **video**: template that contains only videos
	//
	// 	- **text_video**: template that contains text and videos
	//
	// 	- **text_video_button**: template that contains text, videos, and buttons
	//
	// 	- **text_image**: template that contains text and images
	//
	// example:
	//
	// TRANSACTIONAL
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time when the template was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1711006633000
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The reason why the template was rejected.
	//
	// example:
	//
	// None
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// hello_whatsapp
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template. Valid values: WHATSAPP and VIBER.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListChatappTemplateResponseBodyListTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateResponseBodyListTemplate) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetCategory() *string {
	return s.Category
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetLanguage() *string {
	return s.Language
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetReason() *string {
	return s.Reason
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListChatappTemplateResponseBodyListTemplate) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetAuditStatus(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetCategory(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Category = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetLanguage(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetLastUpdateTime(v int64) *ListChatappTemplateResponseBodyListTemplate {
	s.LastUpdateTime = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetReason(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Reason = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateCode(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateName(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateName = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateType(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateType = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) Validate() error {
	return dara.Validate(s)
}
