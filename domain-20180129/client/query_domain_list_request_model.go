// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewEnabled(v bool) *QueryDomainListRequest
	GetAutoRenewEnabled() *bool
	SetCcompany(v string) *QueryDomainListRequest
	GetCcompany() *string
	SetDns(v string) *QueryDomainListRequest
	GetDns() *string
	SetDomainGroupId(v string) *QueryDomainListRequest
	GetDomainGroupId() *string
	SetDomainName(v string) *QueryDomainListRequest
	GetDomainName() *string
	SetEndExpirationDate(v int64) *QueryDomainListRequest
	GetEndExpirationDate() *int64
	SetEndRegistrationDate(v int64) *QueryDomainListRequest
	GetEndRegistrationDate() *int64
	SetLang(v string) *QueryDomainListRequest
	GetLang() *string
	SetOrderByType(v string) *QueryDomainListRequest
	GetOrderByType() *string
	SetOrderKeyType(v string) *QueryDomainListRequest
	GetOrderKeyType() *string
	SetPageNum(v int32) *QueryDomainListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryDomainListRequest
	GetPageSize() *int32
	SetProductDomainType(v string) *QueryDomainListRequest
	GetProductDomainType() *string
	SetQueryType(v string) *QueryDomainListRequest
	GetQueryType() *string
	SetRegistrar(v string) *QueryDomainListRequest
	GetRegistrar() *string
	SetResourceGroupId(v string) *QueryDomainListRequest
	GetResourceGroupId() *string
	SetStartExpirationDate(v int64) *QueryDomainListRequest
	GetStartExpirationDate() *int64
	SetStartRegistrationDate(v int64) *QueryDomainListRequest
	GetStartRegistrationDate() *int64
	SetTag(v []*QueryDomainListRequestTag) *QueryDomainListRequest
	GetTag() []*QueryDomainListRequestTag
	SetUserClientIp(v string) *QueryDomainListRequest
	GetUserClientIp() *string
}

type QueryDomainListRequest struct {
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// The name of the domain owner.
	//
	// example:
	//
	// 广州金烨再生资源回收有限公司
	Ccompany *string `json:"Ccompany,omitempty" xml:"Ccompany,omitempty"`
	Dns      *string `json:"Dns,omitempty" xml:"Dns,omitempty"`
	// <props="china">The ID of the domain group. You can obtain this ID by calling the [QueryDomainGroupList](https://help.aliyun.com/document_detail/69362.html) operation.
	//
	// <props="intl">The ID of the domain group.
	//
	// example:
	//
	// 123456
	DomainGroupId *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// The domain name to query.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the expiration date range. The value is a Unix timestamp in milliseconds. Currently, only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	EndExpirationDate *int64 `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	// The end of the registration date range. The value is a Unix timestamp in milliseconds. Currently, only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	EndRegistrationDate *int64 `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	// The language for API error messages. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// The default value is **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sort order for the results. Valid values:
	//
	// - **ASC**: Ascending.
	//
	// - **DESC**: Descending.
	//
	// > The default value is **DESC**.
	//
	// example:
	//
	// ASC
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The field to use for sorting. Valid values:
	//
	// - **RegistrationDate**: Sorts by registration date.
	//
	// - **ExpirationDate**: Sorts by expiration date.
	//
	// > By default, the results are sorted by the time they were added to the system.
	//
	// example:
	//
	// RegistrationDate
	OrderKeyType *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	// The page number for the paginated results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The domain type. Valid values:
	//
	// - **New gTLD**: new generic top-level domain.
	//
	// - **gTLD**: generic top-level domain.
	//
	// - **ccTLD**: country-code top-level domain.
	//
	// example:
	//
	// New gTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// The type of list to return. Valid values:
	//
	// - **1**: Domain names that require urgent renewal.
	//
	// - **2**: Domain names that require urgent redemption.
	//
	// example:
	//
	// 1
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	Registrar *string `json:"Registrar,omitempty" xml:"Registrar,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2indvyxgpfti
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start of the expiration date range. The value is a Unix timestamp in milliseconds. Currently, only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	StartExpirationDate *int64 `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	// The start of the registration date range. The value is a Unix timestamp in milliseconds. Currently, only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	StartRegistrationDate *int64 `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	// A list of tags.
	Tag []*QueryDomainListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The user\\"s client IP address. You can set this parameter to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainListRequest) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *QueryDomainListRequest) GetCcompany() *string {
	return s.Ccompany
}

func (s *QueryDomainListRequest) GetDns() *string {
	return s.Dns
}

func (s *QueryDomainListRequest) GetDomainGroupId() *string {
	return s.DomainGroupId
}

func (s *QueryDomainListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainListRequest) GetEndExpirationDate() *int64 {
	return s.EndExpirationDate
}

func (s *QueryDomainListRequest) GetEndRegistrationDate() *int64 {
	return s.EndRegistrationDate
}

func (s *QueryDomainListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainListRequest) GetOrderByType() *string {
	return s.OrderByType
}

func (s *QueryDomainListRequest) GetOrderKeyType() *string {
	return s.OrderKeyType
}

func (s *QueryDomainListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDomainListRequest) GetProductDomainType() *string {
	return s.ProductDomainType
}

func (s *QueryDomainListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryDomainListRequest) GetRegistrar() *string {
	return s.Registrar
}

func (s *QueryDomainListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryDomainListRequest) GetStartExpirationDate() *int64 {
	return s.StartExpirationDate
}

func (s *QueryDomainListRequest) GetStartRegistrationDate() *int64 {
	return s.StartRegistrationDate
}

func (s *QueryDomainListRequest) GetTag() []*QueryDomainListRequestTag {
	return s.Tag
}

func (s *QueryDomainListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainListRequest) SetAutoRenewEnabled(v bool) *QueryDomainListRequest {
	s.AutoRenewEnabled = &v
	return s
}

func (s *QueryDomainListRequest) SetCcompany(v string) *QueryDomainListRequest {
	s.Ccompany = &v
	return s
}

func (s *QueryDomainListRequest) SetDns(v string) *QueryDomainListRequest {
	s.Dns = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainGroupId(v string) *QueryDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainName(v string) *QueryDomainListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListRequest) SetEndExpirationDate(v int64) *QueryDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetEndRegistrationDate(v int64) *QueryDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetLang(v string) *QueryDomainListRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderByType(v string) *QueryDomainListRequest {
	s.OrderByType = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderKeyType(v string) *QueryDomainListRequest {
	s.OrderKeyType = &v
	return s
}

func (s *QueryDomainListRequest) SetPageNum(v int32) *QueryDomainListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDomainListRequest) SetPageSize(v int32) *QueryDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainListRequest) SetProductDomainType(v string) *QueryDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *QueryDomainListRequest) SetQueryType(v string) *QueryDomainListRequest {
	s.QueryType = &v
	return s
}

func (s *QueryDomainListRequest) SetRegistrar(v string) *QueryDomainListRequest {
	s.Registrar = &v
	return s
}

func (s *QueryDomainListRequest) SetResourceGroupId(v string) *QueryDomainListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryDomainListRequest) SetStartExpirationDate(v int64) *QueryDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetStartRegistrationDate(v int64) *QueryDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetTag(v []*QueryDomainListRequestTag) *QueryDomainListRequest {
	s.Tag = v
	return s
}

func (s *QueryDomainListRequest) SetUserClientIp(v string) *QueryDomainListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainListRequest) Validate() error {
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

type QueryDomainListRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// 备注
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 标签1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDomainListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListRequestTag) GoString() string {
	return s.String()
}

func (s *QueryDomainListRequestTag) GetKey() *string {
	return s.Key
}

func (s *QueryDomainListRequestTag) GetValue() *string {
	return s.Value
}

func (s *QueryDomainListRequestTag) SetKey(v string) *QueryDomainListRequestTag {
	s.Key = &v
	return s
}

func (s *QueryDomainListRequestTag) SetValue(v string) *QueryDomainListRequestTag {
	s.Value = &v
	return s
}

func (s *QueryDomainListRequestTag) Validate() error {
	return dara.Validate(s)
}
