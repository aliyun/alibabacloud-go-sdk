// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocType(v string) *CreateWorkspaceDocShrinkRequest
	GetDocType() *string
	SetName(v string) *CreateWorkspaceDocShrinkRequest
	GetName() *string
	SetParentNodeId(v string) *CreateWorkspaceDocShrinkRequest
	GetParentNodeId() *string
	SetTemplateId(v string) *CreateWorkspaceDocShrinkRequest
	GetTemplateId() *string
	SetTemplateType(v string) *CreateWorkspaceDocShrinkRequest
	GetTemplateType() *string
	SetTenantContextShrink(v string) *CreateWorkspaceDocShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *CreateWorkspaceDocShrinkRequest
	GetWorkspaceId() *string
}

type CreateWorkspaceDocShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DOC
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试文档
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// YRBGv0Ye
	ParentNodeId *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	// example:
	//
	// 123243
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// team_template
	TemplateType        *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *CreateWorkspaceDocShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceDocShrinkRequest) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *CreateWorkspaceDocShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkspaceDocShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateWorkspaceDocShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateWorkspaceDocShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceDocShrinkRequest) SetDocType(v string) *CreateWorkspaceDocShrinkRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetName(v string) *CreateWorkspaceDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetParentNodeId(v string) *CreateWorkspaceDocShrinkRequest {
	s.ParentNodeId = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTemplateId(v string) *CreateWorkspaceDocShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTemplateType(v string) *CreateWorkspaceDocShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTenantContextShrink(v string) *CreateWorkspaceDocShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetWorkspaceId(v string) *CreateWorkspaceDocShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
