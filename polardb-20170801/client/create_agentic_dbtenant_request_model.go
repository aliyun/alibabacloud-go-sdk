// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAgenticDBTenantRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateAgenticDBTenantRequest
	GetDescription() *string
	SetRegionId(v string) *CreateAgenticDBTenantRequest
	GetRegionId() *string
	SetTenantName(v string) *CreateAgenticDBTenantRequest
	GetTenantName() *string
}

type CreateAgenticDBTenantRequest struct {
	// The AgenticDB cluster ID, which starts with pagc-. You can obtain this ID after creating a cluster by calling CreateDBCluster with AgenticDbType set to dedicated. Standard PolarDB cluster IDs start with pc- and are not supported by this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The tenant description. The description can be up to 256 characters in length.
	//
	// example:
	//
	// Sample tenant
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreateAgenticDBTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBTenantRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticDBTenantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgenticDBTenantRequest) GetTenantName() *string {
	return s.TenantName
}

func (s *CreateAgenticDBTenantRequest) SetDBClusterId(v string) *CreateAgenticDBTenantRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBTenantRequest) SetDescription(v string) *CreateAgenticDBTenantRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticDBTenantRequest) SetRegionId(v string) *CreateAgenticDBTenantRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgenticDBTenantRequest) SetTenantName(v string) *CreateAgenticDBTenantRequest {
	s.TenantName = &v
	return s
}

func (s *CreateAgenticDBTenantRequest) Validate() error {
	return dara.Validate(s)
}
