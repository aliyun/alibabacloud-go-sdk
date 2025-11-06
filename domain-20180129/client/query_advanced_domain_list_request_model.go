// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAdvancedDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupId(v int64) *QueryAdvancedDomainListRequest
	GetDomainGroupId() *int64
	SetDomainNameSort(v bool) *QueryAdvancedDomainListRequest
	GetDomainNameSort() *bool
	SetDomainStatus(v int32) *QueryAdvancedDomainListRequest
	GetDomainStatus() *int32
	SetEndExpirationDate(v int64) *QueryAdvancedDomainListRequest
	GetEndExpirationDate() *int64
	SetEndLength(v int32) *QueryAdvancedDomainListRequest
	GetEndLength() *int32
	SetEndRegistrationDate(v int64) *QueryAdvancedDomainListRequest
	GetEndRegistrationDate() *int64
	SetExcluded(v string) *QueryAdvancedDomainListRequest
	GetExcluded() *string
	SetExcludedPrefix(v bool) *QueryAdvancedDomainListRequest
	GetExcludedPrefix() *bool
	SetExcludedSuffix(v bool) *QueryAdvancedDomainListRequest
	GetExcludedSuffix() *bool
	SetExpirationDateSort(v bool) *QueryAdvancedDomainListRequest
	GetExpirationDateSort() *bool
	SetForm(v int32) *QueryAdvancedDomainListRequest
	GetForm() *int32
	SetIsPremiumDomain(v bool) *QueryAdvancedDomainListRequest
	GetIsPremiumDomain() *bool
	SetKeyWord(v string) *QueryAdvancedDomainListRequest
	GetKeyWord() *string
	SetKeyWordPrefix(v bool) *QueryAdvancedDomainListRequest
	GetKeyWordPrefix() *bool
	SetKeyWordSuffix(v bool) *QueryAdvancedDomainListRequest
	GetKeyWordSuffix() *bool
	SetLang(v string) *QueryAdvancedDomainListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryAdvancedDomainListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryAdvancedDomainListRequest
	GetPageSize() *int32
	SetProductDomainType(v string) *QueryAdvancedDomainListRequest
	GetProductDomainType() *string
	SetProductDomainTypeSort(v bool) *QueryAdvancedDomainListRequest
	GetProductDomainTypeSort() *bool
	SetRegistrationDateSort(v bool) *QueryAdvancedDomainListRequest
	GetRegistrationDateSort() *bool
	SetResourceGroupId(v string) *QueryAdvancedDomainListRequest
	GetResourceGroupId() *string
	SetStartExpirationDate(v int64) *QueryAdvancedDomainListRequest
	GetStartExpirationDate() *int64
	SetStartLength(v int32) *QueryAdvancedDomainListRequest
	GetStartLength() *int32
	SetStartRegistrationDate(v int64) *QueryAdvancedDomainListRequest
	GetStartRegistrationDate() *int64
	SetSuffixs(v string) *QueryAdvancedDomainListRequest
	GetSuffixs() *string
	SetTag(v []*QueryAdvancedDomainListRequestTag) *QueryAdvancedDomainListRequest
	GetTag() []*QueryAdvancedDomainListRequestTag
	SetTradeType(v int32) *QueryAdvancedDomainListRequest
	GetTradeType() *int32
	SetUserClientIp(v string) *QueryAdvancedDomainListRequest
	GetUserClientIp() *string
}

type QueryAdvancedDomainListRequest struct {
	// example:
	//
	// -1
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// example:
	//
	// false
	DomainNameSort *bool `json:"DomainNameSort,omitempty" xml:"DomainNameSort,omitempty"`
	// example:
	//
	// 1
	DomainStatus *int32 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// 1522080000000
	EndExpirationDate *int64 `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	// example:
	//
	// 5
	EndLength *int32 `json:"EndLength,omitempty" xml:"EndLength,omitempty"`
	// example:
	//
	// 1522080000000
	EndRegistrationDate *int64 `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	// example:
	//
	// test
	Excluded *string `json:"Excluded,omitempty" xml:"Excluded,omitempty"`
	// example:
	//
	// false
	ExcludedPrefix *bool `json:"ExcludedPrefix,omitempty" xml:"ExcludedPrefix,omitempty"`
	// example:
	//
	// false
	ExcludedSuffix *bool `json:"ExcludedSuffix,omitempty" xml:"ExcludedSuffix,omitempty"`
	// example:
	//
	// false
	ExpirationDateSort *bool `json:"ExpirationDateSort,omitempty" xml:"ExpirationDateSort,omitempty"`
	// example:
	//
	// 1
	Form            *int32 `json:"Form,omitempty" xml:"Form,omitempty"`
	IsPremiumDomain *bool  `json:"IsPremiumDomain,omitempty" xml:"IsPremiumDomain,omitempty"`
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// example:
	//
	// false
	KeyWordPrefix *bool `json:"KeyWordPrefix,omitempty" xml:"KeyWordPrefix,omitempty"`
	// example:
	//
	// true
	KeyWordSuffix *bool `json:"KeyWordSuffix,omitempty" xml:"KeyWordSuffix,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// gTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// example:
	//
	// false
	ProductDomainTypeSort *bool `json:"ProductDomainTypeSort,omitempty" xml:"ProductDomainTypeSort,omitempty"`
	// example:
	//
	// false
	RegistrationDateSort *bool `json:"RegistrationDateSort,omitempty" xml:"RegistrationDateSort,omitempty"`
	// example:
	//
	// rg-acfmw6bpc6n7zai
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1522080000000
	StartExpirationDate *int64 `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	// example:
	//
	// 5
	StartLength *int32 `json:"StartLength,omitempty" xml:"StartLength,omitempty"`
	// example:
	//
	// 1522080000000
	StartRegistrationDate *int64 `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	// example:
	//
	// com.cn
	Suffixs *string                              `json:"Suffixs,omitempty" xml:"Suffixs,omitempty"`
	Tag     []*QueryAdvancedDomainListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// -1
	TradeType *int32 `json:"TradeType,omitempty" xml:"TradeType,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryAdvancedDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListRequest) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *QueryAdvancedDomainListRequest) GetDomainNameSort() *bool {
	return s.DomainNameSort
}

func (s *QueryAdvancedDomainListRequest) GetDomainStatus() *int32 {
	return s.DomainStatus
}

func (s *QueryAdvancedDomainListRequest) GetEndExpirationDate() *int64 {
	return s.EndExpirationDate
}

func (s *QueryAdvancedDomainListRequest) GetEndLength() *int32 {
	return s.EndLength
}

func (s *QueryAdvancedDomainListRequest) GetEndRegistrationDate() *int64 {
	return s.EndRegistrationDate
}

func (s *QueryAdvancedDomainListRequest) GetExcluded() *string {
	return s.Excluded
}

func (s *QueryAdvancedDomainListRequest) GetExcludedPrefix() *bool {
	return s.ExcludedPrefix
}

func (s *QueryAdvancedDomainListRequest) GetExcludedSuffix() *bool {
	return s.ExcludedSuffix
}

func (s *QueryAdvancedDomainListRequest) GetExpirationDateSort() *bool {
	return s.ExpirationDateSort
}

func (s *QueryAdvancedDomainListRequest) GetForm() *int32 {
	return s.Form
}

func (s *QueryAdvancedDomainListRequest) GetIsPremiumDomain() *bool {
	return s.IsPremiumDomain
}

func (s *QueryAdvancedDomainListRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryAdvancedDomainListRequest) GetKeyWordPrefix() *bool {
	return s.KeyWordPrefix
}

func (s *QueryAdvancedDomainListRequest) GetKeyWordSuffix() *bool {
	return s.KeyWordSuffix
}

func (s *QueryAdvancedDomainListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryAdvancedDomainListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAdvancedDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAdvancedDomainListRequest) GetProductDomainType() *string {
	return s.ProductDomainType
}

func (s *QueryAdvancedDomainListRequest) GetProductDomainTypeSort() *bool {
	return s.ProductDomainTypeSort
}

func (s *QueryAdvancedDomainListRequest) GetRegistrationDateSort() *bool {
	return s.RegistrationDateSort
}

func (s *QueryAdvancedDomainListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryAdvancedDomainListRequest) GetStartExpirationDate() *int64 {
	return s.StartExpirationDate
}

func (s *QueryAdvancedDomainListRequest) GetStartLength() *int32 {
	return s.StartLength
}

func (s *QueryAdvancedDomainListRequest) GetStartRegistrationDate() *int64 {
	return s.StartRegistrationDate
}

func (s *QueryAdvancedDomainListRequest) GetSuffixs() *string {
	return s.Suffixs
}

func (s *QueryAdvancedDomainListRequest) GetTag() []*QueryAdvancedDomainListRequestTag {
	return s.Tag
}

func (s *QueryAdvancedDomainListRequest) GetTradeType() *int32 {
	return s.TradeType
}

func (s *QueryAdvancedDomainListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryAdvancedDomainListRequest) SetDomainGroupId(v int64) *QueryAdvancedDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetDomainNameSort(v bool) *QueryAdvancedDomainListRequest {
	s.DomainNameSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetDomainStatus(v int32) *QueryAdvancedDomainListRequest {
	s.DomainStatus = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndExpirationDate(v int64) *QueryAdvancedDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndLength(v int32) *QueryAdvancedDomainListRequest {
	s.EndLength = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndRegistrationDate(v int64) *QueryAdvancedDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcluded(v string) *QueryAdvancedDomainListRequest {
	s.Excluded = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcludedPrefix(v bool) *QueryAdvancedDomainListRequest {
	s.ExcludedPrefix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcludedSuffix(v bool) *QueryAdvancedDomainListRequest {
	s.ExcludedSuffix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExpirationDateSort(v bool) *QueryAdvancedDomainListRequest {
	s.ExpirationDateSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetForm(v int32) *QueryAdvancedDomainListRequest {
	s.Form = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetIsPremiumDomain(v bool) *QueryAdvancedDomainListRequest {
	s.IsPremiumDomain = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWord(v string) *QueryAdvancedDomainListRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWordPrefix(v bool) *QueryAdvancedDomainListRequest {
	s.KeyWordPrefix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWordSuffix(v bool) *QueryAdvancedDomainListRequest {
	s.KeyWordSuffix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetLang(v string) *QueryAdvancedDomainListRequest {
	s.Lang = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetPageNum(v int32) *QueryAdvancedDomainListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetPageSize(v int32) *QueryAdvancedDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetProductDomainType(v string) *QueryAdvancedDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetProductDomainTypeSort(v bool) *QueryAdvancedDomainListRequest {
	s.ProductDomainTypeSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetRegistrationDateSort(v bool) *QueryAdvancedDomainListRequest {
	s.RegistrationDateSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetResourceGroupId(v string) *QueryAdvancedDomainListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartExpirationDate(v int64) *QueryAdvancedDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartLength(v int32) *QueryAdvancedDomainListRequest {
	s.StartLength = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartRegistrationDate(v int64) *QueryAdvancedDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetSuffixs(v string) *QueryAdvancedDomainListRequest {
	s.Suffixs = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetTag(v []*QueryAdvancedDomainListRequestTag) *QueryAdvancedDomainListRequest {
	s.Tag = v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetTradeType(v int32) *QueryAdvancedDomainListRequest {
	s.TradeType = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetUserClientIp(v string) *QueryAdvancedDomainListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) Validate() error {
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

type QueryAdvancedDomainListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryAdvancedDomainListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListRequestTag) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListRequestTag) GetKey() *string {
	return s.Key
}

func (s *QueryAdvancedDomainListRequestTag) GetValue() *string {
	return s.Value
}

func (s *QueryAdvancedDomainListRequestTag) SetKey(v string) *QueryAdvancedDomainListRequestTag {
	s.Key = &v
	return s
}

func (s *QueryAdvancedDomainListRequestTag) SetValue(v string) *QueryAdvancedDomainListRequestTag {
	s.Value = &v
	return s
}

func (s *QueryAdvancedDomainListRequestTag) Validate() error {
	return dara.Validate(s)
}
