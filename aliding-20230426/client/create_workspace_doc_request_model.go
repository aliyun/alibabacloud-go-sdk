// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocType(v string) *CreateWorkspaceDocRequest
	GetDocType() *string
	SetName(v string) *CreateWorkspaceDocRequest
	GetName() *string
	SetParentNodeId(v string) *CreateWorkspaceDocRequest
	GetParentNodeId() *string
	SetTemplateId(v string) *CreateWorkspaceDocRequest
	GetTemplateId() *string
	SetTemplateType(v string) *CreateWorkspaceDocRequest
	GetTemplateType() *string
	SetTenantContext(v *CreateWorkspaceDocRequestTenantContext) *CreateWorkspaceDocRequest
	GetTenantContext() *CreateWorkspaceDocRequestTenantContext
	SetWorkspaceId(v string) *CreateWorkspaceDocRequest
	GetWorkspaceId() *string
}

type CreateWorkspaceDocRequest struct {
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
	TemplateType  *string                                 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContext *CreateWorkspaceDocRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequest) GetDocType() *string {
	return s.DocType
}

func (s *CreateWorkspaceDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceDocRequest) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *CreateWorkspaceDocRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkspaceDocRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateWorkspaceDocRequest) GetTenantContext() *CreateWorkspaceDocRequestTenantContext {
	return s.TenantContext
}

func (s *CreateWorkspaceDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceDocRequest) SetDocType(v string) *CreateWorkspaceDocRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetName(v string) *CreateWorkspaceDocRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetParentNodeId(v string) *CreateWorkspaceDocRequest {
	s.ParentNodeId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateId(v string) *CreateWorkspaceDocRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateType(v string) *CreateWorkspaceDocRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTenantContext(v *CreateWorkspaceDocRequestTenantContext) *CreateWorkspaceDocRequest {
	s.TenantContext = v
	return s
}

func (s *CreateWorkspaceDocRequest) SetWorkspaceId(v string) *CreateWorkspaceDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceDocRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkspaceDocRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkspaceDocRequestTenantContext) SetTenantId(v string) *CreateWorkspaceDocRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateWorkspaceDocRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
