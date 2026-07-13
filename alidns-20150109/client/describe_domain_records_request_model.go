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
	// The sorting direction. Valid values: DESC, ASC. Default value: DESC.
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name.<props="china"> For more information, see [DescribeDomains](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0).
	//
	// <props="intl">For more information, see [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain group.
	//
	// - If you do not specify GroupId, the query is performed on all domain names.
	//
	// - If you set GroupId to 0, an empty result is returned.
	//
	// - If you set GroupId to -1, the query is performed on the default group. The default group includes domain names that are not assigned to a group.
	//
	// - If you set GroupId to -2, the query is performed on all domain names.
	//
	// - If you specify another value for GroupId, the query is performed on the specified group.
	//
	//   <props="china">For more information, see [DescribeDomainGroups](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomaingroups?spm=a2c4g.11186623.help-menu-search-29697.d_0).
	//
	//   <props="intl">For more information, see [DescribeDomainGroups](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaingroups?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// example:
	//
	// 2****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The DNS resolution line. Default value: **default**.
	//
	// For more information, see:
	//
	// <props="china">[Enumeration of DNS resolution lines](https://help.aliyun.com/document_detail/29807.html).
	//
	// <props="intl">
	//
	// [Enumeration of DNS resolution lines](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm)
	//
	// example:
	//
	// cn_mobile_anhui
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The sorting method. Records are sorted in descending order based on the time they were added.
	//
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. The value starts from **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Maximum value: **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword for the host record. The search uses a fuzzy match pattern and is not case-sensitive.
	//
	// example:
	//
	// www
	RRKeyWord *string `json:"RRKeyWord,omitempty" xml:"RRKeyWord,omitempty"`
	// The search mode. Valid values: LIKE, EXACT, ADVANCED, and **COMBINATION**.
	//
	// - If you set this parameter to LIKE or EXACT, use the KeyWord parameter. The RRKeyWord, TypeKeyWord, ValueKeyWord, Type, Line, and Status parameters are invalid.
	//
	// - If you set this parameter to ADVANCED, use the RRKeyWord, TypeKeyWord, ValueKeyWord, Type, Line, and Status parameters. The RRKeyWord and ValueKeyWord parameters support fuzzy matching.
	//
	// - If you set this parameter to COMBINATION, use the RRKeyWord, TypeKeyWord, ValueKeyWord, Type, Line, and Status parameters. All these parameters support only exact matching.
	//
	// - If you do not specify this parameter:
	//
	//   - If you specify the keyWord parameter, the search mode is set to LIKE.
	//
	//   - If you do not specify the keyWord parameter, RRKeyWord and ValueKeyWord support fuzzy matching, and TypeKeyWord, Type, Line, and Status support exact matching.
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The status of the DNS record. Valid values:
	//
	// Enable: The DNS record is enabled.
	//
	// Disable: The DNS record is paused.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the DNS record. For more information, see:
	//
	// <props="china">[DNS record types](https://help.aliyun.com/document_detail/29805.html)
	//
	// <props="intl">[DNS record types](https://www.alibabacloud.com/help/zh/doc-detail/29805.htm)
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The keyword for the record type. The search uses an exact match and is not case-sensitive.
	//
	// example:
	//
	// MX
	TypeKeyWord *string `json:"TypeKeyWord,omitempty" xml:"TypeKeyWord,omitempty"`
	// The keyword for the record value. The search uses a fuzzy match pattern and is not case-sensitive.
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
