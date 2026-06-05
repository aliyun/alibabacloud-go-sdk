// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPublishHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppPublishHistoryRequest
	GetBizId() *string
	SetKeyword(v string) *ListAppPublishHistoryRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAppPublishHistoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppPublishHistoryRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListAppPublishHistoryRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppPublishHistoryRequest
	GetPageSize() *int32
	SetSort(v string) *ListAppPublishHistoryRequest
	GetSort() *string
	SetStatus(v string) *ListAppPublishHistoryRequest
	GetStatus() *string
	SetWebsiteDomain(v string) *ListAppPublishHistoryRequest
	GetWebsiteDomain() *string
}

type ListAppPublishHistoryRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EndTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// www.aliyun.com
	WebsiteDomain *string `json:"WebsiteDomain,omitempty" xml:"WebsiteDomain,omitempty"`
}

func (s ListAppPublishHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppPublishHistoryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAppPublishHistoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPublishHistoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPublishHistoryRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppPublishHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppPublishHistoryRequest) GetSort() *string {
	return s.Sort
}

func (s *ListAppPublishHistoryRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAppPublishHistoryRequest) GetWebsiteDomain() *string {
	return s.WebsiteDomain
}

func (s *ListAppPublishHistoryRequest) SetBizId(v string) *ListAppPublishHistoryRequest {
	s.BizId = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetKeyword(v string) *ListAppPublishHistoryRequest {
	s.Keyword = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetMaxResults(v int32) *ListAppPublishHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetNextToken(v string) *ListAppPublishHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetPageNum(v int32) *ListAppPublishHistoryRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetPageSize(v int32) *ListAppPublishHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetSort(v string) *ListAppPublishHistoryRequest {
	s.Sort = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetStatus(v string) *ListAppPublishHistoryRequest {
	s.Status = &v
	return s
}

func (s *ListAppPublishHistoryRequest) SetWebsiteDomain(v string) *ListAppPublishHistoryRequest {
	s.WebsiteDomain = &v
	return s
}

func (s *ListAppPublishHistoryRequest) Validate() error {
	return dara.Validate(s)
}
