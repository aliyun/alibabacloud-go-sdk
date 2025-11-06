// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScrollDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScrollDomainListResponseBodyData) *ScrollDomainListResponseBody
	GetData() *ScrollDomainListResponseBodyData
	SetPageSize(v int32) *ScrollDomainListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ScrollDomainListResponseBody
	GetRequestId() *string
	SetScrollId(v string) *ScrollDomainListResponseBody
	GetScrollId() *string
	SetTotalItemNum(v int32) *ScrollDomainListResponseBody
	GetTotalItemNum() *int32
}

type ScrollDomainListResponseBody struct {
	// The domain names.
	Data *ScrollDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 722AB7F5-61F0-408C-A012-4784AFD34083
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scroll ID.
	//
	// example:
	//
	// test
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	// The number of remaining domain names to be queried.
	//
	// example:
	//
	// 200
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
}

func (s ScrollDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBody) GetData() *ScrollDomainListResponseBodyData {
	return s.Data
}

func (s *ScrollDomainListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ScrollDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScrollDomainListResponseBody) GetScrollId() *string {
	return s.ScrollId
}

func (s *ScrollDomainListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ScrollDomainListResponseBody) SetData(v *ScrollDomainListResponseBodyData) *ScrollDomainListResponseBody {
	s.Data = v
	return s
}

func (s *ScrollDomainListResponseBody) SetPageSize(v int32) *ScrollDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetRequestId(v string) *ScrollDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetScrollId(v string) *ScrollDomainListResponseBody {
	s.ScrollId = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetTotalItemNum(v int32) *ScrollDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ScrollDomainListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScrollDomainListResponseBodyData struct {
	Domain []*ScrollDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s ScrollDomainListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyData) GetDomain() []*ScrollDomainListResponseBodyDataDomain {
	return s.Domain
}

func (s *ScrollDomainListResponseBodyData) SetDomain(v []*ScrollDomainListResponseBodyDataDomain) *ScrollDomainListResponseBodyData {
	s.Domain = v
	return s
}

func (s *ScrollDomainListResponseBodyData) Validate() error {
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

type ScrollDomainListResponseBodyDataDomain struct {
	// The Domain Name System (DNS) servers of the domain name.
	DnsList *ScrollDomainListResponseBodyDataDomainDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	// The status of real-name verification for the domain name. Valid values:
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
	// NONAUDIT
	DomainAuditStatus *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 1234
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
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **1**: The domain name needs to be renewed.
	//
	// 	- **2**: The domain name needs to be redeemed.
	//
	// 	- **3**: The domain name is normal.
	//
	// 	- **4**: The domain name is being transferred out.
	//
	// 	- **5**: The information about the domain name registrant is being modified.
	//
	// 	- **6**: Real-name verification is not performed on the domain name.
	//
	// 	- **7**: Real-name verification for the domain name fails.
	//
	// 	- **8**: The real-name verification is being reviewed.
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
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The number of days from the expiration date of the domain name to the current date.
	//
	// example:
	//
	// 10
	ExpirationCurrDateDiff *int32 `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	// The time when the domain name expires.
	//
	// example:
	//
	// 2019-02-15 17:30:35
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The time when the domain name expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1550223035000
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
	// The instance ID of the domain name.
	//
	// example:
	//
	// S1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the domain name is a premium domain name.
	//
	// example:
	//
	// false
	Premium *bool `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 2a
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The registrant of the domain name.
	//
	// example:
	//
	// alibaba cloud
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// The registration type of the domain name. Valid values:
	//
	// 	- **1**: individual.
	//
	// 	- **2**: enterprise.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// The time when the domain name was registered.
	//
	// example:
	//
	// 2017-02-15 00:00:00
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// The time when the domain name was registered. This value is a UNIX timestamp that indicates the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1487088000000
	RegistrationDateLong *int64 `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	// The remarks on the domain name.
	//
	// example:
	//
	// test domain
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2yyciz557g3q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tag.
	Tag *ScrollDomainListResponseBodyDataDomainTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// The Chinese name of the domain name registrant.
	//
	// example:
	//
	// 阿里云
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s ScrollDomainListResponseBodyDataDomain) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDnsList() *ScrollDomainListResponseBodyDataDomainDnsList {
	return s.DnsList
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainAuditStatus() *string {
	return s.DomainAuditStatus
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainGroupId() *string {
	return s.DomainGroupId
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *ScrollDomainListResponseBodyDataDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *ScrollDomainListResponseBodyDataDomain) GetEmail() *string {
	return s.Email
}

func (s *ScrollDomainListResponseBodyDataDomain) GetExpirationCurrDateDiff() *int32 {
	return s.ExpirationCurrDateDiff
}

func (s *ScrollDomainListResponseBodyDataDomain) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *ScrollDomainListResponseBodyDataDomain) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *ScrollDomainListResponseBodyDataDomain) GetExpirationDateStatus() *string {
	return s.ExpirationDateStatus
}

func (s *ScrollDomainListResponseBodyDataDomain) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ScrollDomainListResponseBodyDataDomain) GetPremium() *bool {
	return s.Premium
}

func (s *ScrollDomainListResponseBodyDataDomain) GetProductId() *string {
	return s.ProductId
}

func (s *ScrollDomainListResponseBodyDataDomain) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *ScrollDomainListResponseBodyDataDomain) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *ScrollDomainListResponseBodyDataDomain) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *ScrollDomainListResponseBodyDataDomain) GetRegistrationDateLong() *int64 {
	return s.RegistrationDateLong
}

func (s *ScrollDomainListResponseBodyDataDomain) GetRemark() *string {
	return s.Remark
}

func (s *ScrollDomainListResponseBodyDataDomain) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ScrollDomainListResponseBodyDataDomain) GetTag() *ScrollDomainListResponseBodyDataDomainTag {
	return s.Tag
}

func (s *ScrollDomainListResponseBodyDataDomain) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDnsList(v *ScrollDomainListResponseBodyDataDomainDnsList) *ScrollDomainListResponseBodyDataDomain {
	s.DnsList = v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainName(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainType(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetEmail(v string) *ScrollDomainListResponseBodyDataDomain {
	s.Email = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDate(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetInstanceId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetPremium(v bool) *ScrollDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetProductId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrantOrganization(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrantOrganization = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrantType(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRemark(v string) *ScrollDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetResourceGroupId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ResourceGroupId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetTag(v *ScrollDomainListResponseBodyDataDomainTag) *ScrollDomainListResponseBodyDataDomain {
	s.Tag = v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetZhRegistrantOrganization(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) Validate() error {
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

type ScrollDomainListResponseBodyDataDomainDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s ScrollDomainListResponseBodyDataDomainDnsList) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomainDnsList) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomainDnsList) GetDns() []*string {
	return s.Dns
}

func (s *ScrollDomainListResponseBodyDataDomainDnsList) SetDns(v []*string) *ScrollDomainListResponseBodyDataDomainDnsList {
	s.Dns = v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomainDnsList) Validate() error {
	return dara.Validate(s)
}

type ScrollDomainListResponseBodyDataDomainTag struct {
	Tag []*ScrollDomainListResponseBodyDataDomainTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ScrollDomainListResponseBodyDataDomainTag) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomainTag) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomainTag) GetTag() []*ScrollDomainListResponseBodyDataDomainTagTag {
	return s.Tag
}

func (s *ScrollDomainListResponseBodyDataDomainTag) SetTag(v []*ScrollDomainListResponseBodyDataDomainTagTag) *ScrollDomainListResponseBodyDataDomainTag {
	s.Tag = v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomainTag) Validate() error {
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

type ScrollDomainListResponseBodyDataDomainTagTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ScrollDomainListResponseBodyDataDomainTagTag) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomainTagTag) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomainTagTag) GetKey() *string {
	return s.Key
}

func (s *ScrollDomainListResponseBodyDataDomainTagTag) GetValue() *string {
	return s.Value
}

func (s *ScrollDomainListResponseBodyDataDomainTagTag) SetKey(v string) *ScrollDomainListResponseBodyDataDomainTagTag {
	s.Key = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomainTagTag) SetValue(v string) *ScrollDomainListResponseBodyDataDomainTagTag {
	s.Value = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomainTagTag) Validate() error {
	return dara.Validate(s)
}
