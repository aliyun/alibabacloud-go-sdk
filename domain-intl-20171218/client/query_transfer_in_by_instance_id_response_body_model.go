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
	DomainName                                  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                                       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationDate                              *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong                          *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	InstanceId                                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModificationDate                            *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	ModificationDateLong                        *int64  `json:"ModificationDateLong,omitempty" xml:"ModificationDateLong,omitempty"`
	NeedMailCheck                               *bool   `json:"NeedMailCheck,omitempty" xml:"NeedMailCheck,omitempty"`
	ProgressBarType                             *int32  `json:"ProgressBarType,omitempty" xml:"ProgressBarType,omitempty"`
	RequestId                                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                                  *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultDate                                  *string `json:"ResultDate,omitempty" xml:"ResultDate,omitempty"`
	ResultDateLong                              *int64  `json:"ResultDateLong,omitempty" xml:"ResultDateLong,omitempty"`
	ResultMsg                                   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	SimpleTransferInStatus                      *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	Status                                      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmissionDate                              *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
	SubmissionDateLong                          *int64  `json:"SubmissionDateLong,omitempty" xml:"SubmissionDateLong,omitempty"`
	TransferAuthorizationCodeSubmissionDate     *string `json:"TransferAuthorizationCodeSubmissionDate,omitempty" xml:"TransferAuthorizationCodeSubmissionDate,omitempty"`
	TransferAuthorizationCodeSubmissionDateLong *int64  `json:"TransferAuthorizationCodeSubmissionDateLong,omitempty" xml:"TransferAuthorizationCodeSubmissionDateLong,omitempty"`
	UserId                                      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WhoisMailStatus                             *bool   `json:"WhoisMailStatus,omitempty" xml:"WhoisMailStatus,omitempty"`
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
