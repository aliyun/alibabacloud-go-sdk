// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAgenticDBTenantRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteAgenticDBTenantRequest
	GetRegionId() *string
	SetTenantId(v string) *DeleteAgenticDBTenantRequest
	GetTenantId() *string
}

type DeleteAgenticDBTenantRequest struct {
	// The ID of the AgenticDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAgenticDBTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBTenantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBTenantRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAgenticDBTenantRequest) SetDBClusterId(v string) *DeleteAgenticDBTenantRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBTenantRequest) SetRegionId(v string) *DeleteAgenticDBTenantRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBTenantRequest) SetTenantId(v string) *DeleteAgenticDBTenantRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAgenticDBTenantRequest) Validate() error {
	return dara.Validate(s)
}
