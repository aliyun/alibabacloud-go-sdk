// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySmsTemplateResponseBody
	GetCode() *string
	SetCreateDate(v string) *QuerySmsTemplateResponseBody
	GetCreateDate() *string
	SetMessage(v string) *QuerySmsTemplateResponseBody
	GetMessage() *string
	SetReason(v string) *QuerySmsTemplateResponseBody
	GetReason() *string
	SetRequestId(v string) *QuerySmsTemplateResponseBody
	GetRequestId() *string
	SetTemplateCode(v string) *QuerySmsTemplateResponseBody
	GetTemplateCode() *string
	SetTemplateContent(v string) *QuerySmsTemplateResponseBody
	GetTemplateContent() *string
	SetTemplateName(v string) *QuerySmsTemplateResponseBody
	GetTemplateName() *string
	SetTemplateStatus(v int32) *QuerySmsTemplateResponseBody
	GetTemplateStatus() *int32
	SetTemplateType(v int32) *QuerySmsTemplateResponseBody
	GetTemplateType() *int32
}

type QuerySmsTemplateResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the message template was created.
	//
	// example:
	//
	// 2019-06-04 11:42:17
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the message template is rejected is returned.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A974B78-02BF-4C79-ADF3-90CFBA1B55B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_16703****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the message template.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The approval status of the message template. Valid values:
	//
	// 	- **0**: The message template is pending approval.
	//
	// 	- **1**: The message template is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The message template is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **10**: The approval is canceled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsTemplateResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QuerySmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsTemplateResponseBody) GetReason() *string {
	return s.Reason
}

func (s *QuerySmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QuerySmsTemplateResponseBody) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *QuerySmsTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QuerySmsTemplateResponseBody) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *QuerySmsTemplateResponseBody) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *QuerySmsTemplateResponseBody) SetCode(v string) *QuerySmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetCreateDate(v string) *QuerySmsTemplateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetMessage(v string) *QuerySmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetReason(v string) *QuerySmsTemplateResponseBody {
	s.Reason = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetRequestId(v string) *QuerySmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateCode(v string) *QuerySmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateContent(v string) *QuerySmsTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateName(v string) *QuerySmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateStatus(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateType(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
