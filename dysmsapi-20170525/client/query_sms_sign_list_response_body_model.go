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
	// The HTTP status code. Valid values:
	//
	//
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
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
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of signatures to return on each page. Default value: **10**. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned results.
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
	if s.SmsSignList != nil {
		for _, item := range s.SmsSignList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySmsSignListResponseBodySmsSignList struct {
	// The APP-ICP filing entity ID.
	//
	// example:
	//
	// 1000001***123
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// The audit status of the signature. Valid values:
	//
	// - **AUDIT_STATE_INIT**: under review.
	//
	// - **AUDIT_STATE_PASS**: approved.
	//
	// - **AUDIT_STATE_NOT_PASS**: rejected. You can view the rejection reason in the Reason response parameter.
	//
	// - **AUDIT_STATE_CANCEL**: review canceled.
	//
	// example:
	//
	// AUDIT_STATE_NOT_PASS
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The ID of the letter of authorization.
	//
	// example:
	//
	// 1000********1234
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// The scenario type of the signature. Valid values:
	//
	// - Verification code.
	//
	// - General-purpose.
	//
	// example:
	//
	// 验证码类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The time when the SMS signature was created. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-06-04 11:42:17
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The order ID.
	//
	// This parameter is used by auditors when querying the audit. You must provide this order ID if you need to expedite the audit.
	//
	// example:
	//
	// 2005098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The audit remarks.
	//
	// - If the audit status is **approved*	- or **under review**, the Reason parameter is displayed as "No audit remarks".
	//
	// - If the audit status is **rejected**, the Reason parameter displays the specific reason for the rejection.
	Reason *QuerySmsSignListResponseBodySmsSignListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The signature name.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The trademark entity ID.
	//
	// example:
	//
	// 1000009081***
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
	// The audit status of the letter of authorization. Valid values:
	//
	// - true: approved.
	//
	// - false: not approved (includes all statuses other than approved).
	//
	// example:
	//
	// true
	AuthorizationLetterAuditPass *bool `json:"authorizationLetterAuditPass,omitempty" xml:"authorizationLetterAuditPass,omitempty"`
}

func (s QuerySmsSignListResponseBodySmsSignList) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListResponseBodySmsSignList) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
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

func (s *QuerySmsSignListResponseBodySmsSignList) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *QuerySmsSignListResponseBodySmsSignList) GetAuthorizationLetterAuditPass() *bool {
	return s.AuthorizationLetterAuditPass
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAppIcpRecordId(v int64) *QuerySmsSignListResponseBodySmsSignList {
	s.AppIcpRecordId = &v
	return s
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

func (s *QuerySmsSignListResponseBodySmsSignList) SetTrademarkId(v int64) *QuerySmsSignListResponseBodySmsSignList {
	s.TrademarkId = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAuthorizationLetterAuditPass(v bool) *QuerySmsSignListResponseBodySmsSignList {
	s.AuthorizationLetterAuditPass = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) Validate() error {
	if s.Reason != nil {
		if err := s.Reason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySmsSignListResponseBodySmsSignListReason struct {
	// The time when the signature was rejected. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-06-04 13:35:10
	RejectDate *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	// The reason for the rejection.
	//
	// example:
	//
	// 文件不能证明信息真实性，请重新上传。
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The remarks for the rejection.
	//
	// example:
	//
	// 文件不能证明信息真实性，请重新上传。
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
