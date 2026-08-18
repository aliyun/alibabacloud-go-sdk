// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfuseSecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfuseSecurityIpsRequest
	GetRegionId() *string
}

type DescribePxfuseSecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePxfuseSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfuseSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfuseSecurityIpsRequest) SetDBInstanceName(v string) *DescribePxfuseSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfuseSecurityIpsRequest) SetRegionId(v string) *DescribePxfuseSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfuseSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
