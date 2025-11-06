// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAdvancedDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryAdvancedDomainListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryAdvancedDomainListResponseBodyData) *QueryAdvancedDomainListResponseBody
	GetData() *QueryAdvancedDomainListResponseBodyData
	SetNextPage(v bool) *QueryAdvancedDomainListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryAdvancedDomainListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryAdvancedDomainListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryAdvancedDomainListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryAdvancedDomainListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryAdvancedDomainListResponseBody
	GetTotalPageNum() *int32
}

type QueryAdvancedDomainListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryAdvancedDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// D200000-C0B9-4CD3-B92A-9B44A000000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 549
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 275
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryAdvancedDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryAdvancedDomainListResponseBody) GetData() *QueryAdvancedDomainListResponseBodyData {
	return s.Data
}

func (s *QueryAdvancedDomainListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryAdvancedDomainListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAdvancedDomainListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryAdvancedDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAdvancedDomainListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryAdvancedDomainListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryAdvancedDomainListResponseBody) SetCurrentPageNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetData(v *QueryAdvancedDomainListResponseBodyData) *QueryAdvancedDomainListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetNextPage(v bool) *QueryAdvancedDomainListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetPageSize(v int32) *QueryAdvancedDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetPrePage(v bool) *QueryAdvancedDomainListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetRequestId(v string) *QueryAdvancedDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetTotalItemNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetTotalPageNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAdvancedDomainListResponseBodyData struct {
	Domain []*QueryAdvancedDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s QueryAdvancedDomainListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyData) GetDomain() []*QueryAdvancedDomainListResponseBodyDataDomain {
	return s.Domain
}

func (s *QueryAdvancedDomainListResponseBodyData) SetDomain(v []*QueryAdvancedDomainListResponseBodyDataDomain) *QueryAdvancedDomainListResponseBodyData {
	s.Domain = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyData) Validate() error {
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

type QueryAdvancedDomainListResponseBodyDataDomain struct {
	DnsList *QueryAdvancedDomainListResponseBodyDataDomainDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	// example:
	//
	// SUCCEED
	DomainAuditStatus *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	// example:
	//
	// -1
	DomainGroupId   *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 5
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// gTLD
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 356
	ExpirationCurrDateDiff *int32 `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	// example:
	//
	// 2019-04-09 17:07:03
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// example:
	//
	// 1554800823000
	ExpirationDateLong *int64 `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	// example:
	//
	// 1
	ExpirationDateStatus *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	// example:
	//
	// S20182000000000
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	Premium *bool `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// example:
	//
	// 2a
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// Tom
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// example:
	//
	// 2018-04-09 17:07:03
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// example:
	//
	// 1523264823000
	RegistrationDateLong *int64  `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// rg-aek2yyciz557g3q
	ResourceGroupId *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             *QueryAdvancedDomainListResponseBodyDataDomainTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// example:
	//
	// Tom
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDnsList() *QueryAdvancedDomainListResponseBodyDataDomainDnsList {
	return s.DnsList
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainAuditStatus() *string {
	return s.DomainAuditStatus
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainGroupId() *string {
	return s.DomainGroupId
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetEmail() *string {
	return s.Email
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetExpirationCurrDateDiff() *int32 {
	return s.ExpirationCurrDateDiff
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetExpirationDateStatus() *string {
	return s.ExpirationDateStatus
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetPremium() *bool {
	return s.Premium
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetProductId() *string {
	return s.ProductId
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetRegistrationDateLong() *int64 {
	return s.RegistrationDateLong
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetRemark() *string {
	return s.Remark
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetTag() *QueryAdvancedDomainListResponseBodyDataDomainTag {
	return s.Tag
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDnsList(v *QueryAdvancedDomainListResponseBodyDataDomainDnsList) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DnsList = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainName(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainType(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetEmail(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Email = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDate(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetInstanceId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetPremium(v bool) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetProductId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrantOrganization(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrantType(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRemark(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetResourceGroupId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetTag(v *QueryAdvancedDomainListResponseBodyDataDomainTag) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Tag = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetZhRegistrantOrganization(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) Validate() error {
	if s.DnsList != nil {
		if err := s.DnsList.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAdvancedDomainListResponseBodyDataDomainDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomainDnsList) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomainDnsList) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainDnsList) GetDns() []*string {
	return s.Dns
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainDnsList) SetDns(v []*string) *QueryAdvancedDomainListResponseBodyDataDomainDnsList {
	s.Dns = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainDnsList) Validate() error {
	return dara.Validate(s)
}

type QueryAdvancedDomainListResponseBodyDataDomainTag struct {
	Tag []*QueryAdvancedDomainListResponseBodyDataDomainTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomainTag) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomainTag) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTag) GetTag() []*QueryAdvancedDomainListResponseBodyDataDomainTagTag {
	return s.Tag
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTag) SetTag(v []*QueryAdvancedDomainListResponseBodyDataDomainTagTag) *QueryAdvancedDomainListResponseBodyDataDomainTag {
	s.Tag = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTag) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAdvancedDomainListResponseBodyDataDomainTagTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomainTagTag) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomainTagTag) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTagTag) GetKey() *string {
	return s.Key
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTagTag) GetValue() *string {
	return s.Value
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTagTag) SetKey(v string) *QueryAdvancedDomainListResponseBodyDataDomainTagTag {
	s.Key = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTagTag) SetValue(v string) *QueryAdvancedDomainListResponseBodyDataDomainTagTag {
	s.Value = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainTagTag) Validate() error {
	return dara.Validate(s)
}
