// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPolicySetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPolicySetsRequest
	GetNextToken() *string
}

type ListPolicySetsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPolicySetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicySetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPolicySetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicySetsRequest) SetMaxResults(v int32) *ListPolicySetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPolicySetsRequest) SetNextToken(v string) *ListPolicySetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPolicySetsRequest) Validate() error {
	return dara.Validate(s)
}
