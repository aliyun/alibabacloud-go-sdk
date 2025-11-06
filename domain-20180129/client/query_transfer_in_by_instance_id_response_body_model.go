// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryTransferInByInstanceIdResponseBody
	GetDomainName() *string
	SetEmail(v string) *QueryTransferInByInstanceIdResponseBody
	GetEmail() *string
	SetExpirationDate(v string) *QueryTransferInByInstanceIdResponseBody
	GetExpirationDate() *string
	SetExpirationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody
	GetExpirationDateLong() *int64
	SetInstanceId(v string) *QueryTransferInByInstanceIdResponseBody
	GetInstanceId() *string
	SetModificationDate(v string) *QueryTransferInByInstanceIdResponseBody
	GetModificationDate() *string
	SetModificationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody
	GetModificationDateLong() *int64
	SetNeedMailCheck(v bool) *QueryTransferInByInstanceIdResponseBody
	GetNeedMailCheck() *bool
	SetProgressBarType(v int32) *QueryTransferInByInstanceIdResponseBody
	GetProgressBarType() *int32
	SetRequestId(v string) *QueryTransferInByInstanceIdResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryTransferInByInstanceIdResponseBody
	GetResultCode() *string
	SetResultDate(v string) *QueryTransferInByInstanceIdResponseBody
	GetResultDate() *string
	SetResultDateLong(v int64) *QueryTransferInByInstanceIdResponseBody
	GetResultDateLong() *int64
	SetResultMsg(v string) *QueryTransferInByInstanceIdResponseBody
	GetResultMsg() *string
	SetSimpleTransferInStatus(v string) *QueryTransferInByInstanceIdResponseBody
	GetSimpleTransferInStatus() *string
	SetStatus(v int32) *QueryTransferInByInstanceIdResponseBody
	GetStatus() *int32
	SetSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody
	GetSubmissionDate() *string
	SetSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody
	GetSubmissionDateLong() *int64
	SetTransferAuthorizationCodeSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody
	GetTransferAuthorizationCodeSubmissionDate() *string
	SetTransferAuthorizationCodeSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody
	GetTransferAuthorizationCodeSubmissionDateLong() *int64
	SetUserId(v string) *QueryTransferInByInstanceIdResponseBody
	GetUserId() *string
	SetWhoisMailStatus(v bool) *QueryTransferInByInstanceIdResponseBody
	GetWhoisMailStatus() *bool
}

type QueryTransferInByInstanceIdResponseBody struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// example:
	//
	// 1514428524669
	ExpirationDateLong *int64 `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	// example:
	//
	// S20181T0WLI85212
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	ModificationDate *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	// example:
	//
	// 1514428524669
	ModificationDateLong *int64 `json:"ModificationDateLong,omitempty" xml:"ModificationDateLong,omitempty"`
	// example:
	//
	// true
	NeedMailCheck *bool `json:"NeedMailCheck,omitempty" xml:"NeedMailCheck,omitempty"`
	// example:
	//
	// 0
	ProgressBarType *int32 `json:"ProgressBarType,omitempty" xml:"ProgressBarType,omitempty"`
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// clientCancelled
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	ResultDate *string `json:"ResultDate,omitempty" xml:"ResultDate,omitempty"`
	// example:
	//
	// 1514428524669
	ResultDateLong *int64  `json:"ResultDateLong,omitempty" xml:"ResultDateLong,omitempty"`
	ResultMsg      *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// SUCCESS
	SimpleTransferInStatus *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	// example:
	//
	// 11
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	SubmissionDate *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
	// example:
	//
	// 1514428524669
	SubmissionDateLong *int64 `json:"SubmissionDateLong,omitempty" xml:"SubmissionDateLong,omitempty"`
	// example:
	//
	// 2018-03-28 00:41:42
	TransferAuthorizationCodeSubmissionDate *string `json:"TransferAuthorizationCodeSubmissionDate,omitempty" xml:"TransferAuthorizationCodeSubmissionDate,omitempty"`
	// example:
	//
	// 1514428524669
	TransferAuthorizationCodeSubmissionDateLong *int64 `json:"TransferAuthorizationCodeSubmissionDateLong,omitempty" xml:"TransferAuthorizationCodeSubmissionDateLong,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// true
	WhoisMailStatus *bool `json:"WhoisMailStatus,omitempty" xml:"WhoisMailStatus,omitempty"`
}

func (s QueryTransferInByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTransferInByInstanceIdResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryTransferInByInstanceIdResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryTransferInByInstanceIdResponseBody) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryTransferInByInstanceIdResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTransferInByInstanceIdResponseBody) GetModificationDate() *string {
	return s.ModificationDate
}

func (s *QueryTransferInByInstanceIdResponseBody) GetModificationDateLong() *int64 {
	return s.ModificationDateLong
}

func (s *QueryTransferInByInstanceIdResponseBody) GetNeedMailCheck() *bool {
	return s.NeedMailCheck
}

func (s *QueryTransferInByInstanceIdResponseBody) GetProgressBarType() *int32 {
	return s.ProgressBarType
}

func (s *QueryTransferInByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTransferInByInstanceIdResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryTransferInByInstanceIdResponseBody) GetResultDate() *string {
	return s.ResultDate
}

func (s *QueryTransferInByInstanceIdResponseBody) GetResultDateLong() *int64 {
	return s.ResultDateLong
}

func (s *QueryTransferInByInstanceIdResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryTransferInByInstanceIdResponseBody) GetSimpleTransferInStatus() *string {
	return s.SimpleTransferInStatus
}

func (s *QueryTransferInByInstanceIdResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTransferInByInstanceIdResponseBody) GetSubmissionDate() *string {
	return s.SubmissionDate
}

func (s *QueryTransferInByInstanceIdResponseBody) GetSubmissionDateLong() *int64 {
	return s.SubmissionDateLong
}

func (s *QueryTransferInByInstanceIdResponseBody) GetTransferAuthorizationCodeSubmissionDate() *string {
	return s.TransferAuthorizationCodeSubmissionDate
}

func (s *QueryTransferInByInstanceIdResponseBody) GetTransferAuthorizationCodeSubmissionDateLong() *int64 {
	return s.TransferAuthorizationCodeSubmissionDateLong
}

func (s *QueryTransferInByInstanceIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryTransferInByInstanceIdResponseBody) GetWhoisMailStatus() *bool {
	return s.WhoisMailStatus
}

func (s *QueryTransferInByInstanceIdResponseBody) SetDomainName(v string) *QueryTransferInByInstanceIdResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetEmail(v string) *QueryTransferInByInstanceIdResponseBody {
	s.Email = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetExpirationDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetExpirationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetInstanceId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetModificationDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetModificationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ModificationDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetNeedMailCheck(v bool) *QueryTransferInByInstanceIdResponseBody {
	s.NeedMailCheck = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetProgressBarType(v int32) *QueryTransferInByInstanceIdResponseBody {
	s.ProgressBarType = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetRequestId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultCode(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ResultDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultMsg(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSimpleTransferInStatus(v string) *QueryTransferInByInstanceIdResponseBody {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetStatus(v int32) *QueryTransferInByInstanceIdResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.SubmissionDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.SubmissionDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetTransferAuthorizationCodeSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.TransferAuthorizationCodeSubmissionDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetTransferAuthorizationCodeSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.TransferAuthorizationCodeSubmissionDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetUserId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetWhoisMailStatus(v bool) *QueryTransferInByInstanceIdResponseBody {
	s.WhoisMailStatus = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) Validate() error {
	return dara.Validate(s)
}
