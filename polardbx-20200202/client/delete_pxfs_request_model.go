// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeletePxfsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeletePxfsRequest
	GetRegionId() *string
}

type DeletePxfsRequest struct {
	// The ID of the PolarDB-X instance.
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

func (s DeletePxfsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsRequest) GoString() string {
	return s.String()
}

func (s *DeletePxfsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeletePxfsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePxfsRequest) SetDBInstanceName(v string) *DeletePxfsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeletePxfsRequest) SetRegionId(v string) *DeletePxfsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePxfsRequest) Validate() error {
	return dara.Validate(s)
}
