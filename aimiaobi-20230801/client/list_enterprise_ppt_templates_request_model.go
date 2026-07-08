// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterprisePptTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEnterprisePptTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEnterprisePptTemplatesRequest
	GetNextToken() *string
	SetSkip(v int32) *ListEnterprisePptTemplatesRequest
	GetSkip() *int32
	SetWorkspaceId(v string) *ListEnterprisePptTemplatesRequest
	GetWorkspaceId() *string
}

type ListEnterprisePptTemplatesRequest struct {
	// The maximum number of results to return per page. Note: This parameter is not yet in effect.
	//
	// example:
	//
	// null
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. Note: This parameter is not yet in effect.
	//
	// example:
	//
	// XXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to skip (the \\"offset\\"). Note: This parameter is not yet in effect.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListEnterprisePptTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisePptTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListEnterprisePptTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEnterprisePptTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnterprisePptTemplatesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListEnterprisePptTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListEnterprisePptTemplatesRequest) SetMaxResults(v int32) *ListEnterprisePptTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEnterprisePptTemplatesRequest) SetNextToken(v string) *ListEnterprisePptTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEnterprisePptTemplatesRequest) SetSkip(v int32) *ListEnterprisePptTemplatesRequest {
	s.Skip = &v
	return s
}

func (s *ListEnterprisePptTemplatesRequest) SetWorkspaceId(v string) *ListEnterprisePptTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListEnterprisePptTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
