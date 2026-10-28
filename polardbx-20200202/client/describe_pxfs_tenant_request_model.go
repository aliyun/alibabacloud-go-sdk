// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfsTenantRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfsTenantRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribePxfsTenantRequest
	GetTenantId() *string
}

type DescribePxfsTenantRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribePxfsTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsTenantRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfsTenantRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsTenantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsTenantRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribePxfsTenantRequest) SetDBInstanceName(v string) *DescribePxfsTenantRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsTenantRequest) SetRegionId(v string) *DescribePxfsTenantRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsTenantRequest) SetTenantId(v string) *DescribePxfsTenantRequest {
	s.TenantId = &v
	return s
}

func (s *DescribePxfsTenantRequest) Validate() error {
	return dara.Validate(s)
}
