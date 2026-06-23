// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListConditionalAccessPoliciesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListConditionalAccessPoliciesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListConditionalAccessPoliciesRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListConditionalAccessPoliciesRequest
	GetPreviousToken() *string
}

type ListConditionalAccessPoliciesRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Number of rows per page in paginated queries.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next page query token.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Previous page query token.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListConditionalAccessPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListConditionalAccessPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConditionalAccessPoliciesRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListConditionalAccessPoliciesRequest) SetInstanceId(v string) *ListConditionalAccessPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesRequest) SetMaxResults(v int64) *ListConditionalAccessPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConditionalAccessPoliciesRequest) SetNextToken(v string) *ListConditionalAccessPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListConditionalAccessPoliciesRequest) SetPreviousToken(v string) *ListConditionalAccessPoliciesRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListConditionalAccessPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
