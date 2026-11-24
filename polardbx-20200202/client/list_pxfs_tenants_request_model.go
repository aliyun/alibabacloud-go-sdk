// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPxfsTenantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ListPxfsTenantsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ListPxfsTenantsRequest
	GetRegionId() *string
}

type ListPxfsTenantsRequest struct {
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

func (s ListPxfsTenantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsRequest) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ListPxfsTenantsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPxfsTenantsRequest) SetDBInstanceName(v string) *ListPxfsTenantsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ListPxfsTenantsRequest) SetRegionId(v string) *ListPxfsTenantsRequest {
	s.RegionId = &v
	return s
}

func (s *ListPxfsTenantsRequest) Validate() error {
	return dara.Validate(s)
}
