// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePxfsPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ReleasePxfsPublicConnectionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ReleasePxfsPublicConnectionRequest
	GetRegionId() *string
}

type ReleasePxfsPublicConnectionRequest struct {
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
}

func (s ReleasePxfsPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePxfsPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleasePxfsPublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleasePxfsPublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleasePxfsPublicConnectionRequest) SetDBInstanceName(v string) *ReleasePxfsPublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleasePxfsPublicConnectionRequest) SetRegionId(v string) *ReleasePxfsPublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleasePxfsPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
