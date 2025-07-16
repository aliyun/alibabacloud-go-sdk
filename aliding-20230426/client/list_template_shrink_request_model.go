// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateShrinkRequest
	GetNextToken() *string
	SetTemplateType(v string) *ListTemplateShrinkRequest
	GetTemplateType() *string
	SetTenantContextShrink(v string) *ListTemplateShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *ListTemplateShrinkRequest
	GetWorkspaceId() *string
}

type ListTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// zzz
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// public_template
	TemplateType        *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// workspaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplateShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListTemplateShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplateShrinkRequest) SetMaxResults(v int32) *ListTemplateShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetNextToken(v string) *ListTemplateShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetTemplateType(v string) *ListTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetTenantContextShrink(v string) *ListTemplateShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetWorkspaceId(v string) *ListTemplateShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
