// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetTemplateRequest
	GetAppId() *string
	SetTemplateId(v string) *GetTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *GetTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *GetTemplateRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetTemplateRequest
	GetWorkspaceId() *string
}

type GetTemplateRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 146552
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// CGAKLRCS
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTemplateRequest) SetAppId(v string) *GetTemplateRequest {
	s.AppId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateName(v string) *GetTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateRequest) SetTenantId(v string) *GetTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *GetTemplateRequest) SetWorkspaceId(v string) *GetTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
