// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateRequest
	GetNextToken() *string
	SetTemplateType(v string) *ListTemplateRequest
	GetTemplateType() *string
	SetTenantContext(v *ListTemplateRequestTenantContext) *ListTemplateRequest
	GetTenantContext() *ListTemplateRequestTenantContext
	SetWorkspaceId(v string) *ListTemplateRequest
	GetWorkspaceId() *string
}

type ListTemplateRequest struct {
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
	TemplateType  *string                           `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContext *ListTemplateRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// workspaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplateRequest) GetTenantContext() *ListTemplateRequestTenantContext {
	return s.TenantContext
}

func (s *ListTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplateRequest) SetMaxResults(v int32) *ListTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateRequest) SetNextToken(v string) *ListTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateRequest) SetTemplateType(v string) *ListTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplateRequest) SetTenantContext(v *ListTemplateRequestTenantContext) *ListTemplateRequest {
	s.TenantContext = v
	return s
}

func (s *ListTemplateRequest) SetWorkspaceId(v string) *ListTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplateRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTemplateRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTemplateRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListTemplateRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTemplateRequestTenantContext) SetTenantId(v string) *ListTemplateRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListTemplateRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
