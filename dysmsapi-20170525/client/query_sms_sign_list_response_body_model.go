// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySmsSignListResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *QuerySmsSignListResponseBody
	GetCurrentPage() *int32
	SetMessage(v string) *QuerySmsSignListResponseBody
	GetMessage() *string
	SetPageSize(v int32) *QuerySmsSignListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QuerySmsSignListResponseBody
	GetRequestId() *string
	SetSmsSignList(v []*QuerySmsSignListResponseBodySmsSignList) *QuerySmsSignListResponseBody
	GetSmsSignList() []*QuerySmsSignListResponseBodySmsSignList
	SetTotalCount(v int64) *QuerySmsSignListResponseBody
	GetTotalCount() *int64
}

type QuerySmsSignListResponseBody struct {
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
	// The number of signatures per page. Valid values: **1 to 50**.
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
	// The queried message signatures.
	SmsSignList []*QuerySmsSignListResponseBodySmsSignList `json:"SmsSignList,omitempty" xml:"SmsSignList,omitempty" type:"Repeated"`
	// The total number of signatures.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySmsSignListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsSignListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QuerySmsSignListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsSignListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySmsSignListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsSignListResponseBody) GetSmsSignList() []*QuerySmsSignListResponseBodySmsSignList {
	return s.SmsSignList
}

func (s *QuerySmsSignListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QuerySmsSignListResponseBody) SetCode(v string) *QuerySmsSignListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetCurrentPage(v int32) *QuerySmsSignListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetMessage(v string) *QuerySmsSignListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetPageSize(v int32) *QuerySmsSignListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetRequestId(v string) *QuerySmsSignListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetSmsSignList(v []*QuerySmsSignListResponseBodySmsSignList) *QuerySmsSignListResponseBody {
	s.SmsSignList = v
	return s
}

func (s *QuerySmsSignListResponseBody) SetTotalCount(v int64) *QuerySmsSignListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySmsSignListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySmsSignListResponseBodySmsSignList struct {
	// The approval status of the signature. Valid values:
	//
	// 	- **AUDIT_STATE_INIT**: The signature is pending approval.
	//
	// 	- **AUDIT_STATE_PASS**: The signature is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The signature is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **AUDIT_STATE_CANCEL**: The approval is canceled.
	//
	// example:
	//
	// AUDIT_STATE_NOT_PASS
	AuditStatus           *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationLetterId *int64  `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// The type of the signature scenario. The return value ends with "type". Valid values:
	//
	// 	- Verification code type
	//
	// 	- General-purpose type
	//
	// example:
	//
	// Verification code type
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The time when the signature was created. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 236****5
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the signature is rejected is returned.
	Reason *QuerySmsSignListResponseBodySmsSignListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The name of the signature.
	//
	// example:
	//
	// Aliyun
	SignName                     *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	AuthorizationLetterAuditPass *bool   `json:"authorizationLetterAuditPass,omitempty" xml:"authorizationLetterAuditPass,omitempty"`
}

func (s QuerySmsSignListResponseBodySmsSignList) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListResponseBodySmsSignList) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetOrderId() *string {
	return s.OrderId
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetReason() *QuerySmsSignListResponseBodySmsSignListReason {
	return s.Reason
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetAuthorizationLetterAuditPass() *bool {
	return s.AuthorizationLetterAuditPass
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAuditStatus(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.AuditStatus = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAuthorizationLetterId(v int64) *QuerySmsSignListResponseBodySmsSignList {
	s.AuthorizationLetterId = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetBusinessType(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.BusinessType = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetCreateDate(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetOrderId(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.OrderId = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetReason(v *QuerySmsSignListResponseBodySmsSignListReason) *QuerySmsSignListResponseBodySmsSignList {
	s.Reason = v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetSignName(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.SignName = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAuthorizationLetterAuditPass(v bool) *QuerySmsSignListResponseBodySmsSignList {
	s.AuthorizationLetterAuditPass = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) Validate() error {
	return dara.Validate(s)
}

type QuerySmsSignListResponseBodySmsSignListReason struct {
	// The time when the signature was rejected. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-08 19:02:13
	RejectDate *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	// The reason why the signature was rejected.
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

func (s QuerySmsSignListResponseBodySmsSignListReason) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListResponseBodySmsSignListReason) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) GetRejectDate() *string {
	return s.RejectDate
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) GetRejectInfo() *string {
	return s.RejectInfo
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) GetRejectSubInfo() *string {
	return s.RejectSubInfo
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectDate(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectDate = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectInfo(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectInfo = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectSubInfo(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectSubInfo = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) Validate() error {
	return dara.Validate(s)
}
