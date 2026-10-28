// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfsSecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfsSecurityIpsRequest
	GetRegionId() *string
}

type DescribePxfsSecurityIpsRequest struct {
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
}

func (s DescribePxfsSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsSecurityIpsRequest) SetDBInstanceName(v string) *DescribePxfsSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsSecurityIpsRequest) SetRegionId(v string) *DescribePxfsSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
