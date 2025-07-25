// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeDomainLogsRequest
	GetGroupId() *string
	SetKeyWord(v string) *DescribeDomainLogsRequest
	GetKeyWord() *string
	SetLang(v string) *DescribeDomainLogsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDomainLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainLogsRequest
	GetPageSize() *int64
	SetStartDate(v string) *DescribeDomainLogsRequest
	GetStartDate() *string
	SetType(v string) *DescribeDomainLogsRequest
	GetType() *string
	SetEndDate(v string) *DescribeDomainLogsRequest
	GetEndDate() *string
}

type DescribeDomainLogsRequest struct {
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The keyword for the query in "%KeyWord%" mode. The keyword is not case-sensitive.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number to return. The page number starts from **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time for the query. Format: **YYYY-MM-DD**
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The type of object of which you want to query operation logs. Valid values:
	//
	// 	- domain: domain name
	//
	// 	- slavedns: secondary Domain Name System (DNS)
	//
	// example:
	//
	// domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The end time for the query. Format: **YYYY-MM-DD**
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
}

func (s DescribeDomainLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainLogsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeDomainLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainLogsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDomainLogsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDomainLogsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDomainLogsRequest) SetGroupId(v string) *DescribeDomainLogsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetKeyWord(v string) *DescribeDomainLogsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetLang(v string) *DescribeDomainLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetPageNumber(v int64) *DescribeDomainLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetPageSize(v int64) *DescribeDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetStartDate(v string) *DescribeDomainLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetType(v string) *DescribeDomainLogsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDomainLogsRequest) SetEndDate(v string) *DescribeDomainLogsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainLogsRequest) Validate() error {
	return dara.Validate(s)
}
