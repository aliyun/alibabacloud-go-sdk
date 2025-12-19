// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkloadIdentitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkloadIdentitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkloadIdentitiesRequest
	GetNextToken() *string
}

type ListWorkloadIdentitiesRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lksg167PctRcRw0nyoPjdX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListWorkloadIdentitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkloadIdentitiesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkloadIdentitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkloadIdentitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkloadIdentitiesRequest) SetMaxResults(v int32) *ListWorkloadIdentitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkloadIdentitiesRequest) SetNextToken(v string) *ListWorkloadIdentitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkloadIdentitiesRequest) Validate() error {
	return dara.Validate(s)
}
