// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcompany(v string) *QueryDomainListRequest
	GetCcompany() *string
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
	// The name of the domain name registrant.
	//
	// example:
	//
	// Guangzhou Jinye Renewable Resources Recycling Co., Ltd
	Ccompany *string `json:"Ccompany,omitempty" xml:"Ccompany,omitempty"`
	// The ID of the domain name group.
	//
	// example:
	//
	// 123456
	DomainGroupId *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// The domain name. You can search for the domain name in the domain name list.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query domain names based on expiration dates. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed from January 1, 1970, 00:00:00 UTC to the time you perform the query. Only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	EndExpirationDate *int64 `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	// The end of the time range to query domain names based on registration dates. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	EndRegistrationDate *int64 `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order of the information based on which the domain names are sorted, such as the registration date and expiration date. Valid values:
	//
	// 	- **ASC**: ascending order
	//
	// 	- **DESC**: descending order
	//
	// >  If this parameter is not specified, the default value **DESC*	- is used.
	//
	// example:
	//
	// ASC
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The field that you use to sort the domain names. Valid values:
	//
	// 	- **RegistrationDate**: registration date
	//
	// 	- **ExpirationDate**: expiration date
	//
	// >  If this parameter is not specified, the domain names are sorted by the time when they were added to the database.
	//
	// example:
	//
	// RegistrationDate
	OrderKeyType *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **New gTLD**: new generic top-level domain names
	//
	// 	- **gTLD**: generic top-level domain names
	//
	// 	- **ccTLD**: country code top-level domain names
	//
	// example:
	//
	// New gTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// The category of the domain names that you want to query. Valid values:
	//
	// 	- **1**: the domain names that need to be renewed
	//
	// 	- **2**: the domain names that need to be redeemed
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
	// The beginning of the time range to query domain names based on expiration dates. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed from January 1, 1970, 00:00:00 UTC to the time you perform the query. Only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	StartExpirationDate *int64 `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	// The beginning of the time range to query domain names based on registration dates. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Only queries by day are supported.
	//
	// example:
	//
	// 1522080000000
	StartRegistrationDate *int64 `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	// The tags to add to the resource.
	Tag []*QueryDomainListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The IP address of the client. Set the value to **127.0.0.1**.
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

func (s *QueryDomainListRequest) GetCcompany() *string {
	return s.Ccompany
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

func (s *QueryDomainListRequest) SetCcompany(v string) *QueryDomainListRequest {
	s.Ccompany = &v
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
	// The key of the tag to add to the resource.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource.
	//
	// example:
	//
	// testValue
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
