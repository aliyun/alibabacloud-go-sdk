// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryDomainListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryDomainListResponseBodyData) *QueryDomainListResponseBody
	GetData() *QueryDomainListResponseBodyData
	SetNextPage(v bool) *QueryDomainListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryDomainListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryDomainListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryDomainListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryDomainListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryDomainListResponseBody
	GetTotalPageNum() *int32
}

type QueryDomainListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// True
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 2
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryDomainListResponseBody) GetData() *QueryDomainListResponseBodyData {
	return s.Data
}

func (s *QueryDomainListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryDomainListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDomainListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryDomainListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryDomainListResponseBody) SetCurrentPageNum(v int32) *QueryDomainListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryDomainListResponseBody) SetData(v *QueryDomainListResponseBodyData) *QueryDomainListResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainListResponseBody) SetNextPage(v bool) *QueryDomainListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryDomainListResponseBody) SetPageSize(v int32) *QueryDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDomainListResponseBody) SetPrePage(v bool) *QueryDomainListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryDomainListResponseBody) SetRequestId(v string) *QueryDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainListResponseBody) SetTotalItemNum(v int32) *QueryDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryDomainListResponseBody) SetTotalPageNum(v int32) *QueryDomainListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryDomainListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainListResponseBodyData struct {
	Domain []*QueryDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s QueryDomainListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyData) GetDomain() []*QueryDomainListResponseBodyDataDomain {
	return s.Domain
}

func (s *QueryDomainListResponseBodyData) SetDomain(v []*QueryDomainListResponseBodyDataDomain) *QueryDomainListResponseBodyData {
	s.Domain = v
	return s
}

func (s *QueryDomainListResponseBodyData) Validate() error {
	if s.Domain != nil {
		for _, item := range s.Domain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainListResponseBodyDataDomain struct {
	DomainAuditStatus      *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus           *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType             *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	ExpirationCurrDateDiff *int32  `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate         *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong     *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus   *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                *bool   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	ProductId              *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegistrantType         *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrationDate       *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong   *int64  `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                 *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryDomainListResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainAuditStatus() *string {
	return s.DomainAuditStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *QueryDomainListResponseBodyDataDomain) GetExpirationCurrDateDiff() *int32 {
	return s.ExpirationCurrDateDiff
}

func (s *QueryDomainListResponseBodyDataDomain) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainListResponseBodyDataDomain) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryDomainListResponseBodyDataDomain) GetExpirationDateStatus() *string {
	return s.ExpirationDateStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainListResponseBodyDataDomain) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainListResponseBodyDataDomain) GetProductId() *string {
	return s.ProductId
}

func (s *QueryDomainListResponseBodyDataDomain) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryDomainListResponseBodyDataDomain) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *QueryDomainListResponseBodyDataDomain) GetRegistrationDateLong() *int64 {
	return s.RegistrationDateLong
}

func (s *QueryDomainListResponseBodyDataDomain) GetRemark() *string {
	return s.Remark
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainName(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainType(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetInstanceId(v string) *QueryDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetPremium(v bool) *QueryDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetProductId(v string) *QueryDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrantType(v string) *QueryDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRemark(v string) *QueryDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) Validate() error {
	return dara.Validate(s)
}
