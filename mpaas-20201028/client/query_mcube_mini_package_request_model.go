// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcubeMiniPackageRequest
	GetAppId() *string
	SetH5Id(v string) *QueryMcubeMiniPackageRequest
	GetH5Id() *string
	SetId(v string) *QueryMcubeMiniPackageRequest
	GetId() *string
	SetTenantId(v string) *QueryMcubeMiniPackageRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcubeMiniPackageRequest
	GetWorkspaceId() *string
}

type QueryMcubeMiniPackageRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcubeMiniPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniPackageRequest) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniPackageRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcubeMiniPackageRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *QueryMcubeMiniPackageRequest) GetId() *string {
	return s.Id
}

func (s *QueryMcubeMiniPackageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcubeMiniPackageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcubeMiniPackageRequest) SetAppId(v string) *QueryMcubeMiniPackageRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcubeMiniPackageRequest) SetH5Id(v string) *QueryMcubeMiniPackageRequest {
	s.H5Id = &v
	return s
}

func (s *QueryMcubeMiniPackageRequest) SetId(v string) *QueryMcubeMiniPackageRequest {
	s.Id = &v
	return s
}

func (s *QueryMcubeMiniPackageRequest) SetTenantId(v string) *QueryMcubeMiniPackageRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcubeMiniPackageRequest) SetWorkspaceId(v string) *QueryMcubeMiniPackageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcubeMiniPackageRequest) Validate() error {
	return dara.Validate(s)
}
