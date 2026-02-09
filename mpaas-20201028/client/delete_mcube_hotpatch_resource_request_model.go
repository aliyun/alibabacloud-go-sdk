// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeHotpatchResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DeleteMcubeHotpatchResourceRequest
	GetAppCode() *string
	SetAppId(v string) *DeleteMcubeHotpatchResourceRequest
	GetAppId() *string
	SetId(v int64) *DeleteMcubeHotpatchResourceRequest
	GetId() *int64
	SetTenantId(v string) *DeleteMcubeHotpatchResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcubeHotpatchResourceRequest
	GetWorkspaceId() *string
}

type DeleteMcubeHotpatchResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 321594
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
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

func (s DeleteMcubeHotpatchResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeHotpatchResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcubeHotpatchResourceRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteMcubeHotpatchResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcubeHotpatchResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteMcubeHotpatchResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcubeHotpatchResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcubeHotpatchResourceRequest) SetAppCode(v string) *DeleteMcubeHotpatchResourceRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceRequest) SetAppId(v string) *DeleteMcubeHotpatchResourceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceRequest) SetId(v int64) *DeleteMcubeHotpatchResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceRequest) SetTenantId(v string) *DeleteMcubeHotpatchResourceRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceRequest) SetWorkspaceId(v string) *DeleteMcubeHotpatchResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceRequest) Validate() error {
	return dara.Validate(s)
}
