// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeDomainRecordsRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeDomainRecordsRequest
	GetDomainName() *string
	SetGroupId(v int64) *DescribeDomainRecordsRequest
	GetGroupId() *int64
	SetKeyWord(v string) *DescribeDomainRecordsRequest
	GetKeyWord() *string
	SetLang(v string) *DescribeDomainRecordsRequest
	GetLang() *string
	SetLine(v string) *DescribeDomainRecordsRequest
	GetLine() *string
	SetOrderBy(v string) *DescribeDomainRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *DescribeDomainRecordsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainRecordsRequest
	GetPageSize() *int64
	SetRRKeyWord(v string) *DescribeDomainRecordsRequest
	GetRRKeyWord() *string
	SetSearchMode(v string) *DescribeDomainRecordsRequest
	GetSearchMode() *string
	SetStatus(v string) *DescribeDomainRecordsRequest
	GetStatus() *string
	SetType(v string) *DescribeDomainRecordsRequest
	GetType() *string
	SetTypeKeyWord(v string) *DescribeDomainRecordsRequest
	GetTypeKeyWord() *string
	SetValueKeyWord(v string) *DescribeDomainRecordsRequest
	GetValueKeyWord() *string
}

type DescribeDomainRecordsRequest struct {
	// The order in which you want to sort the returned DNS records. Valid values: DESC and ASC. Default value: DESC.
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group.
	//
	// 	- If you do not specify GroupId, all domain names are queried.
	//
	// 	- If you set GroupId to 0, no value is returned.
	//
	// 	- If you set GroupId to 1, the domain names in the default group are queried.
	//
	// 	- If you set GroupId to -2, all domain names are queried.
	//
	// 	- You can also specify GroupId based on the actual group ID.
	//
	// You can call the [DescribeDomainGroups ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaingroups?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line. Default value: **default**.
	//
	// For more information, see
	//
	// [DNS lines](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm).
	//
	// example:
	//
	// cn_mobile_anhui
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The method that is used to sort the returned DNS records. By default, the DNS records are sorted in reverse chronological order based on the time when they were added.
	//
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The hostname keyword based on which the system queries DNS records. The system queries DNS records based on the value of this parameter in fuzzy match mode. The value is not case-sensitive.
	//
	// example:
	//
	// www
	RRKeyWord *string `json:"RRKeyWord,omitempty" xml:"RRKeyWord,omitempty"`
	// The search mode. Valid values: **LIKE, EXACT, and ADVANCED**.
	//
	// 	- If you set SearchMode to LIKE or EXACT, specify KeyWord. In this case, RRKeyWord, TypeKeyWord, ValueKeyWord, Type, Line, and Status are invalid.
	//
	// 	- If you set SearchMode to ADVANCED, specify RRKeyWord, TypeKeyWord, ValueKeyWord, Type, Line, and Status.
	//
	// 	- If you do not specify SearchMode, the system determines the search mode based on the following rules:
	//
	//     	- If KeyWord is specified, the system uses the LIKE mode.
	//
	//     	- If KeyWord is not specified, the system queries DNS records based on values of RRKeyWord and ValueKeyWord in fuzzy match mode, and based on the values of TypeKeyWord, Type, Line, and Status in exact match mode.
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The status of the DNS records to query. Valid values: **Enable and Disable**.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the DNS records to query. For more information, see
	//
	// [DNS record types](https://www.alibabacloud.com/help/zh/doc-detail/29805.htm).
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type keyword based on which the system queries DNS records. The system queries DNS records based on the value of this parameter in exact match mode. The value is not case-sensitive.
	//
	// example:
	//
	// MX
	TypeKeyWord *string `json:"TypeKeyWord,omitempty" xml:"TypeKeyWord,omitempty"`
	// The record value keyword based on which the system queries DNS records. The system queries DNS records based on the value of this parameter in fuzzy match mode. The value is not case-sensitive.
	//
	// example:
	//
	// com
	ValueKeyWord *string `json:"ValueKeyWord,omitempty" xml:"ValueKeyWord,omitempty"`
}

func (s DescribeDomainRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeDomainRecordsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRecordsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeDomainRecordsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeDomainRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainRecordsRequest) GetLine() *string {
	return s.Line
}

func (s *DescribeDomainRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDomainRecordsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainRecordsRequest) GetRRKeyWord() *string {
	return s.RRKeyWord
}

func (s *DescribeDomainRecordsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeDomainRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainRecordsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDomainRecordsRequest) GetTypeKeyWord() *string {
	return s.TypeKeyWord
}

func (s *DescribeDomainRecordsRequest) GetValueKeyWord() *string {
	return s.ValueKeyWord
}

func (s *DescribeDomainRecordsRequest) SetDirection(v string) *DescribeDomainRecordsRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetDomainName(v string) *DescribeDomainRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetGroupId(v int64) *DescribeDomainRecordsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetKeyWord(v string) *DescribeDomainRecordsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetLang(v string) *DescribeDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetLine(v string) *DescribeDomainRecordsRequest {
	s.Line = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetOrderBy(v string) *DescribeDomainRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetPageNumber(v int64) *DescribeDomainRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetPageSize(v int64) *DescribeDomainRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetRRKeyWord(v string) *DescribeDomainRecordsRequest {
	s.RRKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetSearchMode(v string) *DescribeDomainRecordsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetStatus(v string) *DescribeDomainRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetType(v string) *DescribeDomainRecordsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetTypeKeyWord(v string) *DescribeDomainRecordsRequest {
	s.TypeKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) SetValueKeyWord(v string) *DescribeDomainRecordsRequest {
	s.ValueKeyWord = &v
	return s
}

func (s *DescribeDomainRecordsRequest) Validate() error {
	return dara.Validate(s)
}
