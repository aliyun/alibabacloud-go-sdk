// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobHistorysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListPipelineJobHistorysRequest
	GetCategory() *string
	SetIdentifier(v string) *ListPipelineJobHistorysRequest
	GetIdentifier() *string
	SetMaxResults(v int64) *ListPipelineJobHistorysRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPipelineJobHistorysRequest
	GetNextToken() *string
}

type ListPipelineJobHistorysRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DEPLOY
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10_ssasasa
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xsaxsa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPipelineJobHistorysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysRequest) GetCategory() *string {
	return s.Category
}

func (s *ListPipelineJobHistorysRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListPipelineJobHistorysRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPipelineJobHistorysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineJobHistorysRequest) SetCategory(v string) *ListPipelineJobHistorysRequest {
	s.Category = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetIdentifier(v string) *ListPipelineJobHistorysRequest {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetMaxResults(v int64) *ListPipelineJobHistorysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetNextToken(v string) *ListPipelineJobHistorysRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) Validate() error {
	return dara.Validate(s)
}
