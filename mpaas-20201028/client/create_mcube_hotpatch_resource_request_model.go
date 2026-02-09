// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeHotpatchResourceRequest
	GetAppId() *string
	SetFileUrl(v string) *CreateMcubeHotpatchResourceRequest
	GetFileUrl() *string
	SetFixDesc(v string) *CreateMcubeHotpatchResourceRequest
	GetFixDesc() *string
	SetOnexFlag(v bool) *CreateMcubeHotpatchResourceRequest
	GetOnexFlag() *bool
	SetPlatform(v string) *CreateMcubeHotpatchResourceRequest
	GetPlatform() *string
	SetProductVersion(v string) *CreateMcubeHotpatchResourceRequest
	GetProductVersion() *string
	SetTenantId(v string) *CreateMcubeHotpatchResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcubeHotpatchResourceRequest
	GetWorkspaceId() *string
}

type CreateMcubeHotpatchResourceRequest struct {
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
	// https://mcube-prod.oss-cn-hangzhou.aliyuncs.com/cubecard/tempFileForOnex/ONEXE99ED22171502/preProd/TPHWQYXG/8a6177ce-d7c3-434e-8c62-676a54db9667/main.zip
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FixDesc *string `json:"FixDesc,omitempty" xml:"FixDesc,omitempty"`
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
	// iOS,Android
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// ZXCXMAHQ
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeHotpatchResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeHotpatchResourceRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateMcubeHotpatchResourceRequest) GetFixDesc() *string {
	return s.FixDesc
}

func (s *CreateMcubeHotpatchResourceRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *CreateMcubeHotpatchResourceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcubeHotpatchResourceRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *CreateMcubeHotpatchResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeHotpatchResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeHotpatchResourceRequest) SetAppId(v string) *CreateMcubeHotpatchResourceRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetFileUrl(v string) *CreateMcubeHotpatchResourceRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetFixDesc(v string) *CreateMcubeHotpatchResourceRequest {
	s.FixDesc = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetOnexFlag(v bool) *CreateMcubeHotpatchResourceRequest {
	s.OnexFlag = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetPlatform(v string) *CreateMcubeHotpatchResourceRequest {
	s.Platform = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetProductVersion(v string) *CreateMcubeHotpatchResourceRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetTenantId(v string) *CreateMcubeHotpatchResourceRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) SetWorkspaceId(v string) *CreateMcubeHotpatchResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeHotpatchResourceRequest) Validate() error {
	return dara.Validate(s)
}
