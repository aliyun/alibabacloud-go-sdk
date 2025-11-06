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
	// 0
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The domain names.
	Data *QueryDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// false
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the current page is preceded by a page.
	//
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B7AB5469-5E38-4AA9-A920-C65B7A9C8E6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
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
	// The name of the domain name registrant.
	//
	// example:
	//
	// Guangzhou Forest Advertising Decoration Co., LTD
	Ccompany *string `json:"Ccompany,omitempty" xml:"Ccompany,omitempty"`
	// domain transfer status. value:
	//
	// - 0: domain status normal.
	//
	// - 1: domain is pending change holder.
	//
	// - 2: change holder failed.
	//
	// example:
	//
	// 0
	ChgholderStatus *string `json:"ChgholderStatus,omitempty" xml:"ChgholderStatus,omitempty"`
	// The state of real-name verification for the domain name. Valid values:
	//
	// 	- **FAILED**: Real-name verification for the domain name fails.
	//
	// 	- **SUCCEED**: Real-name verification for the domain name is successful.
	//
	// 	- **NONAUDIT**: Real-name verification for the domain name is not performed.
	//
	// 	- **AUDITING**: Real-name verification for the domain name is in progress.
	//
	// example:
	//
	// FAILED
	DomainAuditStatus *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 123456
	DomainGroupId *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// The name of the domain name group.
	//
	// example:
	//
	// test group
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The state of the domain name. Valid values:
	//
	// 	- **1**: The domain name needs to be renewed.
	//
	// 	- **2**: The domain name needs to be redeemed.
	//
	// 	- **3**: The domain name is normal.
	//
	// example:
	//
	// 3
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **New gTLD**
	//
	// 	- **gTLD**
	//
	// 	- **ccTLD**
	//
	// example:
	//
	// gTLD
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The number of days from the expiration date of the domain name to the current date.
	//
	// example:
	//
	// -30
	ExpirationCurrDateDiff *int32 `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	// The time when the domain name expires.
	//
	// example:
	//
	// 2017-11-02 04:00:45
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The validity period of the domain name. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1522080000000
	ExpirationDateLong *int64 `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	// Indicates whether the domain name expires. Valid values:
	//
	// 	- **1**: The domain name does not expire.
	//
	// 	- **2**: The domain name expires.
	//
	// example:
	//
	// 1
	ExpirationDateStatus *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ST20151102120031118
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the domain name is a premium domain name.
	//
	// example:
	//
	// true
	Premium *bool `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 2a
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The registration type of the domain name. Valid values:
	//
	// 	- **1**: individual
	//
	// 	- **2**: enterprise
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	Registrar      *string `json:"Registrar,omitempty" xml:"Registrar,omitempty"`
	// The time when the domain name was registered.
	//
	// example:
	//
	// 2017-11-02 04:00:45
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// Indicates how long the domain name has been registered. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1522080000000
	RegistrationDateLong *int64 `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	// The remarks of the domain name.
	//
	// example:
	//
	// test remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the domain name belongs.
	//
	// example:
	//
	// rg-aek2yyciz557g3q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the resource.
	Tag *QueryDomainListResponseBodyDataDomainTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
}

func (s QueryDomainListResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomain) GetCcompany() *string {
	return s.Ccompany
}

func (s *QueryDomainListResponseBodyDataDomain) GetChgholderStatus() *string {
	return s.ChgholderStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainAuditStatus() *string {
	return s.DomainAuditStatus
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainGroupId() *string {
	return s.DomainGroupId
}

func (s *QueryDomainListResponseBodyDataDomain) GetDomainGroupName() *string {
	return s.DomainGroupName
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

func (s *QueryDomainListResponseBodyDataDomain) GetRegistrar() *string {
	return s.Registrar
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

func (s *QueryDomainListResponseBodyDataDomain) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryDomainListResponseBodyDataDomain) GetTag() *QueryDomainListResponseBodyDataDomainTag {
	return s.Tag
}

func (s *QueryDomainListResponseBodyDataDomain) SetCcompany(v string) *QueryDomainListResponseBodyDataDomain {
	s.Ccompany = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetChgholderStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.ChgholderStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
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

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrar(v string) *QueryDomainListResponseBodyDataDomain {
	s.Registrar = &v
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

func (s *QueryDomainListResponseBodyDataDomain) SetResourceGroupId(v string) *QueryDomainListResponseBodyDataDomain {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetTag(v *QueryDomainListResponseBodyDataDomainTag) *QueryDomainListResponseBodyDataDomain {
	s.Tag = v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) Validate() error {
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainListResponseBodyDataDomainTag struct {
	Tag []*QueryDomainListResponseBodyDataDomainTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryDomainListResponseBodyDataDomainTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomainTag) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomainTag) GetTag() []*QueryDomainListResponseBodyDataDomainTagTag {
	return s.Tag
}

func (s *QueryDomainListResponseBodyDataDomainTag) SetTag(v []*QueryDomainListResponseBodyDataDomainTagTag) *QueryDomainListResponseBodyDataDomainTag {
	s.Tag = v
	return s
}

func (s *QueryDomainListResponseBodyDataDomainTag) Validate() error {
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

type QueryDomainListResponseBodyDataDomainTagTag struct {
	// The key of the tag added to the resource.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag added to the resource.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDomainListResponseBodyDataDomainTagTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomainTagTag) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomainTagTag) GetKey() *string {
	return s.Key
}

func (s *QueryDomainListResponseBodyDataDomainTagTag) GetValue() *string {
	return s.Value
}

func (s *QueryDomainListResponseBodyDataDomainTagTag) SetKey(v string) *QueryDomainListResponseBodyDataDomainTagTag {
	s.Key = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomainTagTag) SetValue(v string) *QueryDomainListResponseBodyDataDomainTagTag {
	s.Value = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomainTagTag) Validate() error {
	return dara.Validate(s)
}
