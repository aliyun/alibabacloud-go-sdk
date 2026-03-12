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
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the current page follows another page.
	//
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 5
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
	DeadDate          *string `json:"DeadDate,omitempty" xml:"DeadDate,omitempty"`
	DeadDateLong      *int64  `json:"DeadDateLong,omitempty" xml:"DeadDateLong,omitempty"`
	DeadDateStatus    *string `json:"DeadDateStatus,omitempty" xml:"DeadDateStatus,omitempty"`
	DomainAuditStatus *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRegType     *string `json:"DomainRegType,omitempty" xml:"DomainRegType,omitempty"`
	DomainStatus      *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType        *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Premium           *bool   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	ProductId         *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegDate           *string `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
	RegDateLong       *int64  `json:"RegDateLong,omitempty" xml:"RegDateLong,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SaleId            *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
}

func (s QueryDomainListResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomain) GetDeadDate() *string {
	return s.DeadDate
}

func (s *QueryDomainListResponseBodyDataDomain) GetDeadDateLong() *int64 {
	return s.DeadDateLong
}

func (s *QueryDomainListResponseBodyDataDomain) GetDeadDateStatus() *string {
	return s.DeadDateStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainAuditStatus() *string {
	return s.DomainAuditStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainRegType() *string {
	return s.DomainRegType
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *QueryDomainListResponseBodyDataDomain) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryDomainListResponseBodyDataDomain) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainListResponseBodyDataDomain) GetProductId() *string {
	return s.ProductId
}

func (s *QueryDomainListResponseBodyDataDomain) GetRegDate() *string {
	return s.RegDate
}

func (s *QueryDomainListResponseBodyDataDomain) GetRegDateLong() *int64 {
	return s.RegDateLong
}

func (s *QueryDomainListResponseBodyDataDomain) GetRemark() *string {
	return s.Remark
}

func (s *QueryDomainListResponseBodyDataDomain) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainListResponseBodyDataDomain) SetDeadDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.DeadDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDeadDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.DeadDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDeadDateStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DeadDateStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainName(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainRegType(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainRegType = &v
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

func (s *QueryDomainListResponseBodyDataDomain) SetGroupId(v string) *QueryDomainListResponseBodyDataDomain {
	s.GroupId = &v
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

func (s *QueryDomainListResponseBodyDataDomain) SetRegDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.RegDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.RegDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRemark(v string) *QueryDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetSaleId(v string) *QueryDomainListResponseBodyDataDomain {
	s.SaleId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) Validate() error {
	return dara.Validate(s)
}
