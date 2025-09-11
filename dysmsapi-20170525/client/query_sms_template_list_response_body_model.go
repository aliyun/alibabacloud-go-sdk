// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySmsTemplateListResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *QuerySmsTemplateListResponseBody
	GetCurrentPage() *int32
	SetMessage(v string) *QuerySmsTemplateListResponseBody
	GetMessage() *string
	SetPageSize(v int32) *QuerySmsTemplateListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QuerySmsTemplateListResponseBody
	GetRequestId() *string
	SetSmsTemplateList(v []*QuerySmsTemplateListResponseBodySmsTemplateList) *QuerySmsTemplateListResponseBody
	GetSmsTemplateList() []*QuerySmsTemplateListResponseBodySmsTemplateList
	SetTotalCount(v int64) *QuerySmsTemplateListResponseBody
	GetTotalCount() *int64
}

type QuerySmsTemplateListResponseBody struct {
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
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of templates per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried message templates.
	SmsTemplateList []*QuerySmsTemplateListResponseBodySmsTemplateList `json:"SmsTemplateList,omitempty" xml:"SmsTemplateList,omitempty" type:"Repeated"`
	// The total number of templates.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySmsTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsTemplateListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QuerySmsTemplateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsTemplateListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySmsTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsTemplateListResponseBody) GetSmsTemplateList() []*QuerySmsTemplateListResponseBodySmsTemplateList {
	return s.SmsTemplateList
}

func (s *QuerySmsTemplateListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QuerySmsTemplateListResponseBody) SetCode(v string) *QuerySmsTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetCurrentPage(v int32) *QuerySmsTemplateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetMessage(v string) *QuerySmsTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetPageSize(v int32) *QuerySmsTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetRequestId(v string) *QuerySmsTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetSmsTemplateList(v []*QuerySmsTemplateListResponseBodySmsTemplateList) *QuerySmsTemplateListResponseBody {
	s.SmsTemplateList = v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetTotalCount(v int64) *QuerySmsTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySmsTemplateListResponseBodySmsTemplateList struct {
	// The approval status of the message template. Valid values:
	//
	// 	- **AUDIT_STATE_INIT**: The message template is pending approval.
	//
	// 	- **AUDIT_STATE_PASS**: The message template is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The message template is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **AUDIT_STATE_CANCEL*	- or **AUDIT_SATE_CANCEL**: The approval is canceled.
	//
	// example:
	//
	// AUDIT_STATE_PASS
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The time when the message template was created. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-06-04 11:42:17
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 2361****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The type of the message template. We recommend that you specify this parameter. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// 	- **7**: digital message
	//
	// > The template type is the same as the value of the TemplateType parameter in the AddSmsTemplate and ModifySmsTemplate operations.
	//
	// example:
	//
	// 0
	OuterTemplateType *int32 `json:"OuterTemplateType,omitempty" xml:"OuterTemplateType,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the message template is rejected is returned.
	Reason        *QuerySmsTemplateListResponseBodySmsTemplateListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	SignatureName *string                                                `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// example:
	//
	// SMS_1525***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the message template.
	//
	// example:
	//
	// 123456789
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template. Valid values:
	//
	// 	- **0**: notification message
	//
	// 	- **1**: promotional message
	//
	// 	- **2**: verification code
	//
	// 	- **6**: message sent to countries or regions outside the Chinese mainland
	//
	// 	- **7**: digital message
	//
	// example:
	//
	// 7
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySmsTemplateListResponseBodySmsTemplateList) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateListResponseBodySmsTemplateList) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetOrderId() *string {
	return s.OrderId
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetOuterTemplateType() *int32 {
	return s.OuterTemplateType
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetReason() *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	return s.Reason
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetSignatureName() *string {
	return s.SignatureName
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetAuditStatus(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.AuditStatus = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetCreateDate(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetOrderId(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.OrderId = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetOuterTemplateType(v int32) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.OuterTemplateType = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetReason(v *QuerySmsTemplateListResponseBodySmsTemplateListReason) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.Reason = v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetSignatureName(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.SignatureName = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateCode(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateCode = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateContent(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateContent = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateName(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateName = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateType(v int32) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateType = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) Validate() error {
	return dara.Validate(s)
}

type QuerySmsTemplateListResponseBodySmsTemplateListReason struct {
	// The time when the message template was rejected. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-06-04 16:01:17
	RejectDate *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	// The reason why the message template was rejected.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The remarks about the rejection.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	RejectSubInfo *string `json:"RejectSubInfo,omitempty" xml:"RejectSubInfo,omitempty"`
}

func (s QuerySmsTemplateListResponseBodySmsTemplateListReason) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateListResponseBodySmsTemplateListReason) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) GetRejectDate() *string {
	return s.RejectDate
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) GetRejectInfo() *string {
	return s.RejectInfo
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) GetRejectSubInfo() *string {
	return s.RejectSubInfo
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectDate(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectDate = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectInfo(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectInfo = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectSubInfo(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectSubInfo = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) Validate() error {
	return dara.Validate(s)
}
