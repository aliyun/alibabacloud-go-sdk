// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditTermsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAuditTermsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuditTermsRequest
	GetNextToken() *string
	SetTermsName(v string) *ListAuditTermsRequest
	GetTermsName() *string
	SetWorkspaceId(v string) *ListAuditTermsRequest
	GetWorkspaceId() *string
}

type ListAuditTermsRequest struct {
	// Maximum number of records to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next page.
	//
	// example:
	//
	// XXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Dictionary name. Optional. Defaults to Default if not specified.
	TermsName *string `json:"TermsName,omitempty" xml:"TermsName,omitempty"`
	// Unique identifier of the Alibaba Cloud Model Studio workspace. Get your [Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAuditTermsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuditTermsRequest) GoString() string {
	return s.String()
}

func (s *ListAuditTermsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuditTermsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuditTermsRequest) GetTermsName() *string {
	return s.TermsName
}

func (s *ListAuditTermsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAuditTermsRequest) SetMaxResults(v int32) *ListAuditTermsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuditTermsRequest) SetNextToken(v string) *ListAuditTermsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuditTermsRequest) SetTermsName(v string) *ListAuditTermsRequest {
	s.TermsName = &v
	return s
}

func (s *ListAuditTermsRequest) SetWorkspaceId(v string) *ListAuditTermsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAuditTermsRequest) Validate() error {
	return dara.Validate(s)
}
