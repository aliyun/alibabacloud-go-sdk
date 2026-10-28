// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfsInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribePxfsInfoRequest
	GetRegionId() *string
}

type DescribePxfsInfoRequest struct {
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

func (s DescribePxfsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsInfoRequest) SetDBInstanceName(v string) *DescribePxfsInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsInfoRequest) SetRegionId(v string) *DescribePxfsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsInfoRequest) Validate() error {
	return dara.Validate(s)
}
