// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMdsCubeTemplateRequest
	GetAppId() *string
	SetTemplateDesc(v string) *CreateMdsCubeTemplateRequest
	GetTemplateDesc() *string
	SetTemplateId(v string) *CreateMdsCubeTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *CreateMdsCubeTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *CreateMdsCubeTemplateRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMdsCubeTemplateRequest
	GetWorkspaceId() *string
}

type CreateMdsCubeTemplateRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test_card
	TemplateDesc *string `json:"TemplateDesc,omitempty" xml:"TemplateDesc,omitempty"`
	// example:
	//
	// test12435
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// test_card
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMdsCubeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMdsCubeTemplateRequest) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *CreateMdsCubeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateMdsCubeTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateMdsCubeTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMdsCubeTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMdsCubeTemplateRequest) SetAppId(v string) *CreateMdsCubeTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) SetTemplateDesc(v string) *CreateMdsCubeTemplateRequest {
	s.TemplateDesc = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) SetTemplateId(v string) *CreateMdsCubeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) SetTemplateName(v string) *CreateMdsCubeTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) SetTenantId(v string) *CreateMdsCubeTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) SetWorkspaceId(v string) *CreateMdsCubeTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMdsCubeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
