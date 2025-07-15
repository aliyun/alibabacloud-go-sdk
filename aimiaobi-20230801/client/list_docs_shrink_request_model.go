// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListDocsShrinkRequest
	GetCategoryId() *string
	SetDocName(v string) *ListDocsShrinkRequest
	GetDocName() *string
	SetDocType(v string) *ListDocsShrinkRequest
	GetDocType() *string
	SetMaxResults(v int32) *ListDocsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocsShrinkRequest
	GetNextToken() *string
	SetSkip(v int32) *ListDocsShrinkRequest
	GetSkip() *int32
	SetStatusesShrink(v string) *ListDocsShrinkRequest
	GetStatusesShrink() *string
	SetWorkspaceId(v string) *ListDocsShrinkRequest
	GetWorkspaceId() *string
}

type ListDocsShrinkRequest struct {
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DocName    *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 52a33dc83779f63641e16f5146cd7125
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Skip           *int32  `json:"Skip,omitempty" xml:"Skip,omitempty"`
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDocsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDocsShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListDocsShrinkRequest) GetDocName() *string {
	return s.DocName
}

func (s *ListDocsShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDocsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocsShrinkRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListDocsShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListDocsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDocsShrinkRequest) SetCategoryId(v string) *ListDocsShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *ListDocsShrinkRequest) SetDocName(v string) *ListDocsShrinkRequest {
	s.DocName = &v
	return s
}

func (s *ListDocsShrinkRequest) SetDocType(v string) *ListDocsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListDocsShrinkRequest) SetMaxResults(v int32) *ListDocsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocsShrinkRequest) SetNextToken(v string) *ListDocsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocsShrinkRequest) SetSkip(v int32) *ListDocsShrinkRequest {
	s.Skip = &v
	return s
}

func (s *ListDocsShrinkRequest) SetStatusesShrink(v string) *ListDocsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListDocsShrinkRequest) SetWorkspaceId(v string) *ListDocsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDocsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
