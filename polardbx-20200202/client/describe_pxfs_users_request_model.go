// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfsUsersRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfsUsersRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribePxfsUsersRequest
	GetTenantId() *string
}

type DescribePxfsUsersRequest struct {
	// The instance ID of the PolarDB-X instance.
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

func (s DescribePxfsUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribePxfsUsersRequest) SetDBInstanceName(v string) *DescribePxfsUsersRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsUsersRequest) SetRegionId(v string) *DescribePxfsUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsUsersRequest) SetTenantId(v string) *DescribePxfsUsersRequest {
	s.TenantId = &v
	return s
}

func (s *DescribePxfsUsersRequest) Validate() error {
	return dara.Validate(s)
}
