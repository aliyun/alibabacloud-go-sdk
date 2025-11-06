// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsCubeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMdsCubeTemplateRequest
	GetAppId() *string
	SetTemplateId(v string) *DeleteMdsCubeTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *DeleteMdsCubeTemplateRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMdsCubeTemplateRequest
	GetWorkspaceId() *string
}

type DeleteMdsCubeTemplateRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMdsCubeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsCubeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteMdsCubeTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMdsCubeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteMdsCubeTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMdsCubeTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMdsCubeTemplateRequest) SetAppId(v string) *DeleteMdsCubeTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMdsCubeTemplateRequest) SetTemplateId(v string) *DeleteMdsCubeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteMdsCubeTemplateRequest) SetTenantId(v string) *DeleteMdsCubeTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMdsCubeTemplateRequest) SetWorkspaceId(v string) *DeleteMdsCubeTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMdsCubeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
