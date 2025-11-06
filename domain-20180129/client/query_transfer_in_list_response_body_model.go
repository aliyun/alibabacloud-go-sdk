// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryTransferInListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryTransferInListResponseBodyData) *QueryTransferInListResponseBody
	GetData() *QueryTransferInListResponseBodyData
	SetNextPage(v bool) *QueryTransferInListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryTransferInListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryTransferInListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryTransferInListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryTransferInListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryTransferInListResponseBody
	GetTotalPageNum() *int32
}

type QueryTransferInListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTransferInListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 40
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 2
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTransferInListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryTransferInListResponseBody) GetData() *QueryTransferInListResponseBodyData {
	return s.Data
}

func (s *QueryTransferInListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryTransferInListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTransferInListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryTransferInListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTransferInListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryTransferInListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryTransferInListResponseBody) SetCurrentPageNum(v int32) *QueryTransferInListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetData(v *QueryTransferInListResponseBodyData) *QueryTransferInListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTransferInListResponseBody) SetNextPage(v bool) *QueryTransferInListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetPageSize(v int32) *QueryTransferInListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetPrePage(v bool) *QueryTransferInListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetRequestId(v string) *QueryTransferInListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetTotalItemNum(v int32) *QueryTransferInListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetTotalPageNum(v int32) *QueryTransferInListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTransferInListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTransferInListResponseBodyData struct {
	TransferInInfo []*QueryTransferInListResponseBodyDataTransferInInfo `json:"TransferInInfo,omitempty" xml:"TransferInInfo,omitempty" type:"Repeated"`
}

func (s QueryTransferInListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBodyData) GetTransferInInfo() []*QueryTransferInListResponseBodyDataTransferInInfo {
	return s.TransferInInfo
}

func (s *QueryTransferInListResponseBodyData) SetTransferInInfo(v []*QueryTransferInListResponseBodyDataTransferInInfo) *QueryTransferInListResponseBodyData {
	s.TransferInInfo = v
	return s
}

func (s *QueryTransferInListResponseBodyData) Validate() error {
	if s.TransferInInfo != nil {
		for _, item := range s.TransferInInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTransferInListResponseBodyDataTransferInInfo struct {
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
	// FAIL
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

func (s QueryTransferInListResponseBodyDataTransferInInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInListResponseBodyDataTransferInInfo) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetEmail() *string {
	return s.Email
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetModificationDate() *string {
	return s.ModificationDate
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetModificationDateLong() *int64 {
	return s.ModificationDateLong
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetNeedMailCheck() *bool {
	return s.NeedMailCheck
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetProgressBarType() *int32 {
	return s.ProgressBarType
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetResultDate() *string {
	return s.ResultDate
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetResultDateLong() *int64 {
	return s.ResultDateLong
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetSimpleTransferInStatus() *string {
	return s.SimpleTransferInStatus
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetSubmissionDate() *string {
	return s.SubmissionDate
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetSubmissionDateLong() *int64 {
	return s.SubmissionDateLong
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetTransferAuthorizationCodeSubmissionDate() *string {
	return s.TransferAuthorizationCodeSubmissionDate
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetTransferAuthorizationCodeSubmissionDateLong() *int64 {
	return s.TransferAuthorizationCodeSubmissionDateLong
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetUserId() *string {
	return s.UserId
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) GetWhoisMailStatus() *bool {
	return s.WhoisMailStatus
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetDomainName(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetEmail(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.Email = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetExpirationDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetExpirationDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetInstanceId(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetModificationDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ModificationDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetModificationDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ModificationDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetNeedMailCheck(v bool) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.NeedMailCheck = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetProgressBarType(v int32) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ProgressBarType = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultCode(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultMsg(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSimpleTransferInStatus(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetStatus(v int32) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.Status = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSubmissionDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SubmissionDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSubmissionDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SubmissionDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetTransferAuthorizationCodeSubmissionDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.TransferAuthorizationCodeSubmissionDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetTransferAuthorizationCodeSubmissionDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.TransferAuthorizationCodeSubmissionDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetUserId(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.UserId = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetWhoisMailStatus(v bool) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.WhoisMailStatus = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) Validate() error {
	return dara.Validate(s)
}
