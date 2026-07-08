// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListDocsRequest
	GetCategoryId() *string
	SetDocName(v string) *ListDocsRequest
	GetDocName() *string
	SetDocType(v string) *ListDocsRequest
	GetDocType() *string
	SetMaxResults(v int32) *ListDocsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocsRequest
	GetNextToken() *string
	SetSkip(v int32) *ListDocsRequest
	GetSkip() *int32
	SetStatuses(v []*int32) *ListDocsRequest
	GetStatuses() []*int32
	SetWorkspaceId(v string) *ListDocsRequest
	GetWorkspaceId() *string
}

type ListDocsRequest struct {
	// Document folder.
	//
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Document name.
	//
	// example:
	//
	// 企业2022年報
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// Document type.
	//
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Maximum number of results returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next page token.
	//
	// example:
	//
	// 52a33dc83779f63641e16f5146cd7125
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Skip n records for paging.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// Document status list.
	Statuses []*int32 `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// Alibaba Cloud Model Studio workspace ID. For more information, see [How to Use Workspaces](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDocsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocsRequest) GoString() string {
	return s.String()
}

func (s *ListDocsRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListDocsRequest) GetDocName() *string {
	return s.DocName
}

func (s *ListDocsRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDocsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListDocsRequest) GetStatuses() []*int32 {
	return s.Statuses
}

func (s *ListDocsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDocsRequest) SetCategoryId(v string) *ListDocsRequest {
	s.CategoryId = &v
	return s
}

func (s *ListDocsRequest) SetDocName(v string) *ListDocsRequest {
	s.DocName = &v
	return s
}

func (s *ListDocsRequest) SetDocType(v string) *ListDocsRequest {
	s.DocType = &v
	return s
}

func (s *ListDocsRequest) SetMaxResults(v int32) *ListDocsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocsRequest) SetNextToken(v string) *ListDocsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocsRequest) SetSkip(v int32) *ListDocsRequest {
	s.Skip = &v
	return s
}

func (s *ListDocsRequest) SetStatuses(v []*int32) *ListDocsRequest {
	s.Statuses = v
	return s
}

func (s *ListDocsRequest) SetWorkspaceId(v string) *ListDocsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDocsRequest) Validate() error {
	return dara.Validate(s)
}
