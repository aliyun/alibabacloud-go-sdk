// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppDomainRedirectRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppDomainRedirectRecordsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListAppDomainRedirectRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppDomainRedirectRecordsRequest
	GetNextToken() *string
}

type ListAppDomainRedirectRecordsRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// yPkgBbQln0gmUnIZSGizGw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAppDomainRedirectRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppDomainRedirectRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAppDomainRedirectRecordsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppDomainRedirectRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppDomainRedirectRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppDomainRedirectRecordsRequest) SetBizId(v string) *ListAppDomainRedirectRecordsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppDomainRedirectRecordsRequest) SetMaxResults(v int32) *ListAppDomainRedirectRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppDomainRedirectRecordsRequest) SetNextToken(v string) *ListAppDomainRedirectRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppDomainRedirectRecordsRequest) Validate() error {
	return dara.Validate(s)
}
