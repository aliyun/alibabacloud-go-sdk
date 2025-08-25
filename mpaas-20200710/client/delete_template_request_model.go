// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteTemplateRequest
	GetAppId() *string
	SetTemplateId(v string) *DeleteTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *DeleteTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *DeleteTemplateRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteTemplateRequest
	GetWorkspaceId() *string
}

type DeleteTemplateRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// COEYM44ER0465E8G
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// NPHTGKNR
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DeleteTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteTemplateRequest) SetAppId(v string) *DeleteTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateId(v string) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteTemplateRequest) SetTenantId(v string) *DeleteTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteTemplateRequest) SetWorkspaceId(v string) *DeleteTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
