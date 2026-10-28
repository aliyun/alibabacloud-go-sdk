// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePxfsPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *AllocatePxfsPublicConnectionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *AllocatePxfsPublicConnectionRequest
	GetRegionId() *string
}

type AllocatePxfsPublicConnectionRequest struct {
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

func (s AllocatePxfsPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePxfsPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocatePxfsPublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocatePxfsPublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocatePxfsPublicConnectionRequest) SetDBInstanceName(v string) *AllocatePxfsPublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocatePxfsPublicConnectionRequest) SetRegionId(v string) *AllocatePxfsPublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocatePxfsPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
