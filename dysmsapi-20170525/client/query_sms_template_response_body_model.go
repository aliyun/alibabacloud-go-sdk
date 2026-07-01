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
	// The status code of the request.
	//
	// - OK indicates that the request was successful.
	//
	// - For a list of other error codes, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2024-06-03 10:02:34
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The review notes for the template.
	//
	// - If the review status is **Approved*	- or **Reviewing**, the message "No review remarks" is returned.
	//
	// - If the review status is **Rejected**, the reason for the rejection is returned.
	//
	// example:
	//
	// 无审批备注
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A974B78-02BF-4C79-ADF3-90CFBA1B55B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template code.
	//
	// example:
	//
	// SMS_1525****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template content.
	//
	// example:
	//
	// 亲爱的会员！阿里云短信服务祝您新年快乐！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The template name.
	//
	// example:
	//
	// 通知短信
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The review status of the template. Valid values:
	//
	// - **0**: Reviewing.
	//
	// - **1**: Approved.
	//
	// - **2**: Rejected. The reason for the rejection is returned in the response. For more information, see [Suggestions for handling a failed review](https://help.aliyun.com/document_detail/65990.html). You can then call the [ModifySmsTemplate](https://help.aliyun.com/document_detail/419287.html) API or modify the template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
	//
	// - **10**: Canceled.
	//
	// example:
	//
	// 0
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The message type. Valid values:
	//
	// - **0**: Verification code.
	//
	// - **1**: Message notification.
	//
	// - **2**: Promotional message.
	//
	// - **3**: International message.
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
