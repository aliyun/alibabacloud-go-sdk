// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInspectReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListClusterInspectReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListClusterInspectReportsRequest
	GetNextToken() *string
}

type ListClusterInspectReportsRequest struct {
	// The maximum number of returned results. Maximum value: 50.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// AK8uQQrxgFKsI3OiS5TbhUQ9R3kPme4I3
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListClusterInspectReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInspectReportsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterInspectReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClusterInspectReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterInspectReportsRequest) SetMaxResults(v int32) *ListClusterInspectReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClusterInspectReportsRequest) SetNextToken(v string) *ListClusterInspectReportsRequest {
	s.NextToken = &v
	return s
}

func (s *ListClusterInspectReportsRequest) Validate() error {
	return dara.Validate(s)
}
