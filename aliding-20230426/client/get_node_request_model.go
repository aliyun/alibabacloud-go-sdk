// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *GetNodeRequest
	GetNodeId() *string
	SetTenantContext(v *GetNodeRequestTenantContext) *GetNodeRequest
	GetTenantContext() *GetNodeRequestTenantContext
	SetWithPermissionRole(v bool) *GetNodeRequest
	GetWithPermissionRole() *bool
	SetWithStatisticalInfo(v bool) *GetNodeRequest
	GetWithStatisticalInfo() *bool
}

type GetNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a9E05BDRVQ9K600yf1NplNDxV63zgkYA
	NodeId        *string                      `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *GetNodeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// true
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
	// example:
	//
	// true
	WithStatisticalInfo *bool `json:"WithStatisticalInfo,omitempty" xml:"WithStatisticalInfo,omitempty"`
}

func (s GetNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeRequest) GetTenantContext() *GetNodeRequestTenantContext {
	return s.TenantContext
}

func (s *GetNodeRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetNodeRequest) GetWithStatisticalInfo() *bool {
	return s.WithStatisticalInfo
}

func (s *GetNodeRequest) SetNodeId(v string) *GetNodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeRequest) SetTenantContext(v *GetNodeRequestTenantContext) *GetNodeRequest {
	s.TenantContext = v
	return s
}

func (s *GetNodeRequest) SetWithPermissionRole(v bool) *GetNodeRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodeRequest) SetWithStatisticalInfo(v bool) *GetNodeRequest {
	s.WithStatisticalInfo = &v
	return s
}

func (s *GetNodeRequest) Validate() error {
	return dara.Validate(s)
}

type GetNodeRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetNodeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetNodeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNodeRequestTenantContext) SetTenantId(v string) *GetNodeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetNodeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
