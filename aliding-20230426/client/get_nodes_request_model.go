// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v []*string) *GetNodesRequest
	GetNodeIds() []*string
	SetOption(v *GetNodesRequestOption) *GetNodesRequest
	GetOption() *GetNodesRequestOption
	SetTenantContext(v *GetNodesRequestTenantContext) *GetNodesRequest
	GetTenantContext() *GetNodesRequestTenantContext
}

type GetNodesRequest struct {
	// This parameter is required.
	NodeIds       []*string                     `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	Option        *GetNodesRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	TenantContext *GetNodesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodesRequest) GoString() string {
	return s.String()
}

func (s *GetNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *GetNodesRequest) GetOption() *GetNodesRequestOption {
	return s.Option
}

func (s *GetNodesRequest) GetTenantContext() *GetNodesRequestTenantContext {
	return s.TenantContext
}

func (s *GetNodesRequest) SetNodeIds(v []*string) *GetNodesRequest {
	s.NodeIds = v
	return s
}

func (s *GetNodesRequest) SetOption(v *GetNodesRequestOption) *GetNodesRequest {
	s.Option = v
	return s
}

func (s *GetNodesRequest) SetTenantContext(v *GetNodesRequestTenantContext) *GetNodesRequest {
	s.TenantContext = v
	return s
}

func (s *GetNodesRequest) Validate() error {
	return dara.Validate(s)
}

type GetNodesRequestOption struct {
	// example:
	//
	// false
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
	// example:
	//
	// false
	WithStatisticalInfo *bool `json:"WithStatisticalInfo,omitempty" xml:"WithStatisticalInfo,omitempty"`
}

func (s GetNodesRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetNodesRequestOption) GoString() string {
	return s.String()
}

func (s *GetNodesRequestOption) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetNodesRequestOption) GetWithStatisticalInfo() *bool {
	return s.WithStatisticalInfo
}

func (s *GetNodesRequestOption) SetWithPermissionRole(v bool) *GetNodesRequestOption {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodesRequestOption) SetWithStatisticalInfo(v bool) *GetNodesRequestOption {
	s.WithStatisticalInfo = &v
	return s
}

func (s *GetNodesRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetNodesRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetNodesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetNodesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetNodesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNodesRequestTenantContext) SetTenantId(v string) *GetNodesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetNodesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
