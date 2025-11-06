// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMdsCubeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMdsCubeResourceRequest
	GetAppId() *string
	SetMockDataUrl(v string) *UpdateMdsCubeResourceRequest
	GetMockDataUrl() *string
	SetOnexFlag(v bool) *UpdateMdsCubeResourceRequest
	GetOnexFlag() *bool
	SetTemplateResourceId(v int64) *UpdateMdsCubeResourceRequest
	GetTemplateResourceId() *int64
	SetTenantId(v string) *UpdateMdsCubeResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *UpdateMdsCubeResourceRequest
	GetWorkspaceId() *string
}

type UpdateMdsCubeResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// url
	MockDataUrl *string `json:"MockDataUrl,omitempty" xml:"MockDataUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateResourceId *int64 `json:"TemplateResourceId,omitempty" xml:"TemplateResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMdsCubeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMdsCubeResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateMdsCubeResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMdsCubeResourceRequest) GetMockDataUrl() *string {
	return s.MockDataUrl
}

func (s *UpdateMdsCubeResourceRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *UpdateMdsCubeResourceRequest) GetTemplateResourceId() *int64 {
	return s.TemplateResourceId
}

func (s *UpdateMdsCubeResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMdsCubeResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMdsCubeResourceRequest) SetAppId(v string) *UpdateMdsCubeResourceRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) SetMockDataUrl(v string) *UpdateMdsCubeResourceRequest {
	s.MockDataUrl = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) SetOnexFlag(v bool) *UpdateMdsCubeResourceRequest {
	s.OnexFlag = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) SetTemplateResourceId(v int64) *UpdateMdsCubeResourceRequest {
	s.TemplateResourceId = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) SetTenantId(v string) *UpdateMdsCubeResourceRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) SetWorkspaceId(v string) *UpdateMdsCubeResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMdsCubeResourceRequest) Validate() error {
	return dara.Validate(s)
}
