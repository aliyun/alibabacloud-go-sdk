// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSecurityDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationSecurityDomainsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationSecurityDomainsResponseBody
	GetNextToken() *string
	SetNormalizationSecurityDomains(v []*ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) *ListNormalizationSecurityDomainsResponseBody
	GetNormalizationSecurityDomains() []*ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains
	SetRequestId(v string) *ListNormalizationSecurityDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationSecurityDomainsResponseBody
	GetTotalCount() *int32
}

type ListNormalizationSecurityDomainsResponseBody struct {
	// The maximum number of entries returned in this query.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of security domains.
	NormalizationSecurityDomains []*ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains `json:"NormalizationSecurityDomains,omitempty" xml:"NormalizationSecurityDomains,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNormalizationSecurityDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSecurityDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationSecurityDomainsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationSecurityDomainsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationSecurityDomainsResponseBody) GetNormalizationSecurityDomains() []*ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains {
	return s.NormalizationSecurityDomains
}

func (s *ListNormalizationSecurityDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationSecurityDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationSecurityDomainsResponseBody) SetMaxResults(v int32) *ListNormalizationSecurityDomainsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBody) SetNextToken(v string) *ListNormalizationSecurityDomainsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBody) SetNormalizationSecurityDomains(v []*ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) *ListNormalizationSecurityDomainsResponseBody {
	s.NormalizationSecurityDomains = v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBody) SetRequestId(v string) *ListNormalizationSecurityDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBody) SetTotalCount(v int32) *ListNormalizationSecurityDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBody) Validate() error {
	if s.NormalizationSecurityDomains != nil {
		for _, item := range s.NormalizationSecurityDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains struct {
	// The security domain ID.
	//
	// example:
	//
	// OTHER
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The security domain name.
	//
	// example:
	//
	// net
	NormalizationSecurityDomainName *string `json:"NormalizationSecurityDomainName,omitempty" xml:"NormalizationSecurityDomainName,omitempty"`
}

func (s ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) GoString() string {
	return s.String()
}

func (s *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) GetNormalizationSecurityDomainName() *string {
	return s.NormalizationSecurityDomainName
}

func (s *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) SetNormalizationSecurityDomainId(v string) *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) SetNormalizationSecurityDomainName(v string) *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains {
	s.NormalizationSecurityDomainName = &v
	return s
}

func (s *ListNormalizationSecurityDomainsResponseBodyNormalizationSecurityDomains) Validate() error {
	return dara.Validate(s)
}
