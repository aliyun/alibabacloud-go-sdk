// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchRollbackTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeHotpatchRollbackTaskRequest
	GetAppId() *string
	SetId(v int64) *CreateMcubeHotpatchRollbackTaskRequest
	GetId() *int64
	SetProductId(v string) *CreateMcubeHotpatchRollbackTaskRequest
	GetProductId() *string
	SetProductVersion(v string) *CreateMcubeHotpatchRollbackTaskRequest
	GetProductVersion() *string
	SetTenantId(v string) *CreateMcubeHotpatchRollbackTaskRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMcubeHotpatchRollbackTaskRequest
	GetWorkspaceId() *string
}

type CreateMcubeHotpatchRollbackTaskRequest struct {
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
	// 1653905
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419_Android-default
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
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

func (s CreateMcubeHotpatchRollbackTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchRollbackTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetAppId(v string) *CreateMcubeHotpatchRollbackTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetId(v int64) *CreateMcubeHotpatchRollbackTaskRequest {
	s.Id = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetProductId(v string) *CreateMcubeHotpatchRollbackTaskRequest {
	s.ProductId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetProductVersion(v string) *CreateMcubeHotpatchRollbackTaskRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetTenantId(v string) *CreateMcubeHotpatchRollbackTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) SetWorkspaceId(v string) *CreateMcubeHotpatchRollbackTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskRequest) Validate() error {
	return dara.Validate(s)
}
