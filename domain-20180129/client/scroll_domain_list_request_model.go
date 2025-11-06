// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScrollDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupId(v int64) *ScrollDomainListRequest
	GetDomainGroupId() *int64
	SetDomainStatus(v int32) *ScrollDomainListRequest
	GetDomainStatus() *int32
	SetEndExpirationDate(v int64) *ScrollDomainListRequest
	GetEndExpirationDate() *int64
	SetEndLength(v int32) *ScrollDomainListRequest
	GetEndLength() *int32
	SetEndRegistrationDate(v int64) *ScrollDomainListRequest
	GetEndRegistrationDate() *int64
	SetExcluded(v string) *ScrollDomainListRequest
	GetExcluded() *string
	SetExcludedPrefix(v bool) *ScrollDomainListRequest
	GetExcludedPrefix() *bool
	SetExcludedSuffix(v bool) *ScrollDomainListRequest
	GetExcludedSuffix() *bool
	SetForm(v int32) *ScrollDomainListRequest
	GetForm() *int32
	SetKeyWord(v string) *ScrollDomainListRequest
	GetKeyWord() *string
	SetKeyWordPrefix(v bool) *ScrollDomainListRequest
	GetKeyWordPrefix() *bool
	SetKeyWordSuffix(v bool) *ScrollDomainListRequest
	GetKeyWordSuffix() *bool
	SetLang(v string) *ScrollDomainListRequest
	GetLang() *string
	SetPageSize(v int32) *ScrollDomainListRequest
	GetPageSize() *int32
	SetProductDomainType(v string) *ScrollDomainListRequest
	GetProductDomainType() *string
	SetResourceGroupId(v string) *ScrollDomainListRequest
	GetResourceGroupId() *string
	SetScrollId(v string) *ScrollDomainListRequest
	GetScrollId() *string
	SetStartExpirationDate(v int64) *ScrollDomainListRequest
	GetStartExpirationDate() *int64
	SetStartLength(v int32) *ScrollDomainListRequest
	GetStartLength() *int32
	SetStartRegistrationDate(v int64) *ScrollDomainListRequest
	GetStartRegistrationDate() *int64
	SetSuffixs(v string) *ScrollDomainListRequest
	GetSuffixs() *string
	SetTradeType(v int32) *ScrollDomainListRequest
	GetTradeType() *int32
	SetUserClientIp(v string) *ScrollDomainListRequest
	GetUserClientIp() *string
}

type ScrollDomainListRequest struct {
	// The ID of the domain name group. You can call the [QueryDomainGroupList](https://help.aliyun.com/document_detail/69362.html) operation to obtain the ID of the domain name group.
	//
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **0**: All.
	//
	// 	- **1**: The domain name needs to be renewed.
	//
	// 	- **2**: The domain name needs to be redeemed.
	//
	// 	- **3**: The domain name is normal.
	//
	// 	- **4**: The domain name is being transferred from Alibaba Cloud.
	//
	// 	- **5**: The information about the domain name registrant is being modified.
	//
	// 	- **6**: Real-name verification is not performed on the domain name.
	//
	// 	- **7**: Real-name verification for the domain name fails. Real-name reverification is required.
	//
	// 	- **8**: The domain name is being reviewed.
	//
	// example:
	//
	// 0
	DomainStatus *int32 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The end of the time range to query domain names based on expiration dates. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1541520000000
	EndExpirationDate *int64 `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	// The end of domain name length to query.
	//
	// example:
	//
	// 3
	EndLength *int32 `json:"EndLength,omitempty" xml:"EndLength,omitempty"`
	// The end of the time range to query domain names based on registration dates. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1541520000000
	EndRegistrationDate *int64 `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	// The keyword that is used to exclude domain names.
	//
	// example:
	//
	// test
	Excluded *string `json:"Excluded,omitempty" xml:"Excluded,omitempty"`
	// Specifies whether to exclude the prefix keyword.
	//
	// example:
	//
	// false
	ExcludedPrefix *bool `json:"ExcludedPrefix,omitempty" xml:"ExcludedPrefix,omitempty"`
	// Specifies whether to exclude the suffix keyword.
	//
	// example:
	//
	// true
	ExcludedSuffix *bool `json:"ExcludedSuffix,omitempty" xml:"ExcludedSuffix,omitempty"`
	// The composition of the domain name.
	//
	// example:
	//
	// 1
	Form *int32 `json:"Form,omitempty" xml:"Form,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// Specifies whether the keyword is the prefix.
	//
	// example:
	//
	// true
	KeyWordPrefix *bool `json:"KeyWordPrefix,omitempty" xml:"KeyWordPrefix,omitempty"`
	// Specifies whether the keyword is the suffix.
	//
	// example:
	//
	// false
	KeyWordSuffix *bool `json:"KeyWordSuffix,omitempty" xml:"KeyWordSuffix,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **New gTLD**
	//
	// 	- **gTLD**
	//
	// 	- **ccTLD**
	//
	// 	- **other**
	//
	// example:
	//
	// gTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw6bpc6n7zai
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scroll ID. This parameter is a technical parameter.
	//
	// example:
	//
	// test
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	// The beginning of the time range to query domain names based on expiration dates. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1541520000000
	StartExpirationDate *int64 `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	// The start of the domain name length to query.
	//
	// example:
	//
	// 0
	StartLength *int32 `json:"StartLength,omitempty" xml:"StartLength,omitempty"`
	// The beginning of the time range to query domain names based on registration dates. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1541520000000
	StartRegistrationDate *int64 `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	// The suffixes of domain names to be queried. Separate multiple suffixes with commas (,).
	//
	// example:
	//
	// com
	Suffixs *string `json:"Suffixs,omitempty" xml:"Suffixs,omitempty"`
	// The publishing status of the domain name. Valid values:
	//
	// 	- **2**: The domain name is published at a fixed price.
	//
	// 	- **3**: The domain name is published with the price negotiable.
	//
	// 	- **4**: The domain name is published for bidding.
	//
	// 	- **6**: The domain name is published with price push.
	//
	// 	- **-1**: The domain name is not published.
	//
	// example:
	//
	// -1
	TradeType *int32 `json:"TradeType,omitempty" xml:"TradeType,omitempty"`
	// The IP address of the client. Set the value to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ScrollDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListRequest) GoString() string {
	return s.String()
}

func (s *ScrollDomainListRequest) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *ScrollDomainListRequest) GetDomainStatus() *int32 {
	return s.DomainStatus
}

func (s *ScrollDomainListRequest) GetEndExpirationDate() *int64 {
	return s.EndExpirationDate
}

func (s *ScrollDomainListRequest) GetEndLength() *int32 {
	return s.EndLength
}

func (s *ScrollDomainListRequest) GetEndRegistrationDate() *int64 {
	return s.EndRegistrationDate
}

func (s *ScrollDomainListRequest) GetExcluded() *string {
	return s.Excluded
}

func (s *ScrollDomainListRequest) GetExcludedPrefix() *bool {
	return s.ExcludedPrefix
}

func (s *ScrollDomainListRequest) GetExcludedSuffix() *bool {
	return s.ExcludedSuffix
}

func (s *ScrollDomainListRequest) GetForm() *int32 {
	return s.Form
}

func (s *ScrollDomainListRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ScrollDomainListRequest) GetKeyWordPrefix() *bool {
	return s.KeyWordPrefix
}

func (s *ScrollDomainListRequest) GetKeyWordSuffix() *bool {
	return s.KeyWordSuffix
}

func (s *ScrollDomainListRequest) GetLang() *string {
	return s.Lang
}

func (s *ScrollDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ScrollDomainListRequest) GetProductDomainType() *string {
	return s.ProductDomainType
}

func (s *ScrollDomainListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ScrollDomainListRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *ScrollDomainListRequest) GetStartExpirationDate() *int64 {
	return s.StartExpirationDate
}

func (s *ScrollDomainListRequest) GetStartLength() *int32 {
	return s.StartLength
}

func (s *ScrollDomainListRequest) GetStartRegistrationDate() *int64 {
	return s.StartRegistrationDate
}

func (s *ScrollDomainListRequest) GetSuffixs() *string {
	return s.Suffixs
}

func (s *ScrollDomainListRequest) GetTradeType() *int32 {
	return s.TradeType
}

func (s *ScrollDomainListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ScrollDomainListRequest) SetDomainGroupId(v int64) *ScrollDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *ScrollDomainListRequest) SetDomainStatus(v int32) *ScrollDomainListRequest {
	s.DomainStatus = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndExpirationDate(v int64) *ScrollDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndLength(v int32) *ScrollDomainListRequest {
	s.EndLength = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndRegistrationDate(v int64) *ScrollDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcluded(v string) *ScrollDomainListRequest {
	s.Excluded = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcludedPrefix(v bool) *ScrollDomainListRequest {
	s.ExcludedPrefix = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcludedSuffix(v bool) *ScrollDomainListRequest {
	s.ExcludedSuffix = &v
	return s
}

func (s *ScrollDomainListRequest) SetForm(v int32) *ScrollDomainListRequest {
	s.Form = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWord(v string) *ScrollDomainListRequest {
	s.KeyWord = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWordPrefix(v bool) *ScrollDomainListRequest {
	s.KeyWordPrefix = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWordSuffix(v bool) *ScrollDomainListRequest {
	s.KeyWordSuffix = &v
	return s
}

func (s *ScrollDomainListRequest) SetLang(v string) *ScrollDomainListRequest {
	s.Lang = &v
	return s
}

func (s *ScrollDomainListRequest) SetPageSize(v int32) *ScrollDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *ScrollDomainListRequest) SetProductDomainType(v string) *ScrollDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *ScrollDomainListRequest) SetResourceGroupId(v string) *ScrollDomainListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ScrollDomainListRequest) SetScrollId(v string) *ScrollDomainListRequest {
	s.ScrollId = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartExpirationDate(v int64) *ScrollDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartLength(v int32) *ScrollDomainListRequest {
	s.StartLength = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartRegistrationDate(v int64) *ScrollDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetSuffixs(v string) *ScrollDomainListRequest {
	s.Suffixs = &v
	return s
}

func (s *ScrollDomainListRequest) SetTradeType(v int32) *ScrollDomainListRequest {
	s.TradeType = &v
	return s
}

func (s *ScrollDomainListRequest) SetUserClientIp(v string) *ScrollDomainListRequest {
	s.UserClientIp = &v
	return s
}

func (s *ScrollDomainListRequest) Validate() error {
	return dara.Validate(s)
}
