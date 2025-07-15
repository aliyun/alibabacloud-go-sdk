// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditContentErrorTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAuditContentErrorTypesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuditContentErrorTypesRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListAuditContentErrorTypesRequest
	GetWorkspaceId() *string
}

type ListAuditContentErrorTypesRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAuditContentErrorTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuditContentErrorTypesRequest) GoString() string {
	return s.String()
}

func (s *ListAuditContentErrorTypesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuditContentErrorTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuditContentErrorTypesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAuditContentErrorTypesRequest) SetMaxResults(v int32) *ListAuditContentErrorTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuditContentErrorTypesRequest) SetNextToken(v string) *ListAuditContentErrorTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuditContentErrorTypesRequest) SetWorkspaceId(v string) *ListAuditContentErrorTypesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAuditContentErrorTypesRequest) Validate() error {
	return dara.Validate(s)
}
