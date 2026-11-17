// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomDomainsResponseBody
	GetCode() *string
	SetCustomDomains(v []*CustomDomainResp) *ListCustomDomainsResponseBody
	GetCustomDomains() []*CustomDomainResp
	SetHasMore(v bool) *ListCustomDomainsResponseBody
	GetHasMore() *bool
	SetMaxResults(v int32) *ListCustomDomainsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCustomDomainsResponseBody
	GetMessage() *string
	SetNextOffset(v int32) *ListCustomDomainsResponseBody
	GetNextOffset() *int32
	SetNextToken(v string) *ListCustomDomainsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCustomDomainsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCustomDomainsResponseBody
	GetTotal() *int64
}

type ListCustomDomainsResponseBody struct {
	// The response code. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The custom domain names.
	//
	// example:
	//
	// []
	CustomDomains []*CustomDomainResp `json:"customDomains" xml:"customDomains" type:"Repeated"`
	// Indicates whether a next page is available.
	//
	// example:
	//
	// False
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The maximum number of records to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The offset for the next page.
	//
	// example:
	//
	// 100
	NextOffset *int32 `json:"nextOffset,omitempty" xml:"nextOffset,omitempty"`
	// The pagination token for the next page. This parameter is empty if no more results are available.
	//
	// example:
	//
	// cae5f900-8b1d-4c0e-9c2a-1a2b3c4d5e6f
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BCFAE0A-9FA9-5F72-8E8B-724632BC19A9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 65
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCustomDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomDomainsResponseBody) GetCustomDomains() []*CustomDomainResp {
	return s.CustomDomains
}

func (s *ListCustomDomainsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListCustomDomainsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomDomainsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomDomainsResponseBody) GetNextOffset() *int32 {
	return s.NextOffset
}

func (s *ListCustomDomainsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomDomainsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCustomDomainsResponseBody) SetCode(v string) *ListCustomDomainsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetCustomDomains(v []*CustomDomainResp) *ListCustomDomainsResponseBody {
	s.CustomDomains = v
	return s
}

func (s *ListCustomDomainsResponseBody) SetHasMore(v bool) *ListCustomDomainsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetMaxResults(v int32) *ListCustomDomainsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetMessage(v string) *ListCustomDomainsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetNextOffset(v int32) *ListCustomDomainsResponseBody {
	s.NextOffset = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetNextToken(v string) *ListCustomDomainsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetRequestId(v string) *ListCustomDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomDomainsResponseBody) SetTotal(v int64) *ListCustomDomainsResponseBody {
	s.Total = &v
	return s
}

func (s *ListCustomDomainsResponseBody) Validate() error {
	if s.CustomDomains != nil {
		for _, item := range s.CustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
