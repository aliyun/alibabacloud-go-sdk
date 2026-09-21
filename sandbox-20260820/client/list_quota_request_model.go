// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQuotaRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaRequest
	GetNextToken() *string
}

type ListQuotaRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// Trigger_fasdatalake_deductionfee
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaRequest) SetMaxResults(v int32) *ListQuotaRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaRequest) SetNextToken(v string) *ListQuotaRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaRequest) Validate() error {
	return dara.Validate(s)
}
