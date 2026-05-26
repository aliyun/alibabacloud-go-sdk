// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTokenVaultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTokenVaultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTokenVaultsRequest
	GetNextToken() *string
}

type ListTokenVaultsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTokenVaultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTokenVaultsRequest) GoString() string {
	return s.String()
}

func (s *ListTokenVaultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTokenVaultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTokenVaultsRequest) SetMaxResults(v int32) *ListTokenVaultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTokenVaultsRequest) SetNextToken(v string) *ListTokenVaultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTokenVaultsRequest) Validate() error {
	return dara.Validate(s)
}
