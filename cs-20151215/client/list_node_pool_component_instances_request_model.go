// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodePoolComponentInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentInstancesRequest
	GetNextToken() *string
}

type ListNodePoolComponentInstancesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
}

func (s ListNodePoolComponentInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentInstancesRequest) SetMaxResults(v int32) *ListNodePoolComponentInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentInstancesRequest) SetNextToken(v string) *ListNodePoolComponentInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentInstancesRequest) Validate() error {
	return dara.Validate(s)
}
