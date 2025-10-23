// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMailAddressByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *QueryMailAddressByParamResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryMailAddressByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMailAddressByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryMailAddressByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryMailAddressByParamResponseBodyData) *QueryMailAddressByParamResponseBody
	GetData() *QueryMailAddressByParamResponseBodyData
}

type QueryMailAddressByParamResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of sending addresses
	Data *QueryMailAddressByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryMailAddressByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMailAddressByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryMailAddressByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMailAddressByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMailAddressByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryMailAddressByParamResponseBody) GetData() *QueryMailAddressByParamResponseBodyData {
	return s.Data
}

func (s *QueryMailAddressByParamResponseBody) SetPageNumber(v int32) *QueryMailAddressByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetPageSize(v int32) *QueryMailAddressByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetRequestId(v string) *QueryMailAddressByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetTotalCount(v int32) *QueryMailAddressByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetData(v *QueryMailAddressByParamResponseBodyData) *QueryMailAddressByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryMailAddressByParamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMailAddressByParamResponseBodyData struct {
	MailAddress []*QueryMailAddressByParamResponseBodyDataMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Repeated"`
}

func (s QueryMailAddressByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMailAddressByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBodyData) GetMailAddress() []*QueryMailAddressByParamResponseBodyDataMailAddress {
	return s.MailAddress
}

func (s *QueryMailAddressByParamResponseBodyData) SetMailAddress(v []*QueryMailAddressByParamResponseBodyDataMailAddress) *QueryMailAddressByParamResponseBodyData {
	s.MailAddress = v
	return s
}

func (s *QueryMailAddressByParamResponseBodyData) Validate() error {
	if s.MailAddress != nil {
		for _, item := range s.MailAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMailAddressByParamResponseBodyDataMailAddress struct {
	// Sending address
	//
	// example:
	//
	// sender@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Account status, frozen: 1, normal: 0.
	//
	// example:
	//
	// 0
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	ConfigSetId   *string `json:"ConfigSetId,omitempty" xml:"ConfigSetId,omitempty"`
	ConfigSetName *string `json:"ConfigSetName,omitempty" xml:"ConfigSetName,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Daily quota limit
	//
	// example:
	//
	// 10000
	DailyCount *string `json:"DailyCount,omitempty" xml:"DailyCount,omitempty"`
	// Daily quota
	//
	// example:
	//
	// 100
	DailyReqCount *string `json:"DailyReqCount,omitempty" xml:"DailyReqCount,omitempty"`
	// Domain status, 0 indicates normal, 1 indicates abnormal.
	//
	// example:
	//
	// 0
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// Sending address ID
	//
	// example:
	//
	// 12122
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	// Monthly quota limit
	//
	// example:
	//
	// 300000
	MonthCount *string `json:"MonthCount,omitempty" xml:"MonthCount,omitempty"`
	// Monthly quota
	//
	// example:
	//
	// 20000
	MonthReqCount *string `json:"MonthReqCount,omitempty" xml:"MonthReqCount,omitempty"`
	// Reply address
	//
	// example:
	//
	// test@example.com
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// Reply address status
	//
	// example:
	//
	// 0
	ReplyStatus *string `json:"ReplyStatus,omitempty" xml:"ReplyStatus,omitempty"`
	// Sending address type. Values:
	//
	// - batch: bulk email
	//
	// - trigger: triggered email
	//
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s QueryMailAddressByParamResponseBodyDataMailAddress) String() string {
	return dara.Prettify(s)
}

func (s QueryMailAddressByParamResponseBodyDataMailAddress) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetConfigSetId() *string {
	return s.ConfigSetId
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetConfigSetName() *string {
	return s.ConfigSetName
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetDailyCount() *string {
	return s.DailyCount
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetDailyReqCount() *string {
	return s.DailyReqCount
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetMailAddressId() *string {
	return s.MailAddressId
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetMonthCount() *string {
	return s.MonthCount
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetMonthReqCount() *string {
	return s.MonthReqCount
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetReplyStatus() *string {
	return s.ReplyStatus
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) GetSendtype() *string {
	return s.Sendtype
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetAccountName(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.AccountName = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetAccountStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.AccountStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetConfigSetId(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ConfigSetId = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetConfigSetName(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ConfigSetName = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetCreateTime(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.CreateTime = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDailyCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DailyCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDailyReqCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DailyReqCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDomainStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DomainStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMailAddressId(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MailAddressId = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMonthCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MonthCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMonthReqCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MonthReqCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetReplyAddress(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ReplyAddress = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetReplyStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ReplyStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetSendtype(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.Sendtype = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) Validate() error {
	return dara.Validate(s)
}
