// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfuseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeletePxfuseRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeletePxfuseRequest
	GetRegionId() *string
}

type DeletePxfuseRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePxfuseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfuseRequest) GoString() string {
	return s.String()
}

func (s *DeletePxfuseRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeletePxfuseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePxfuseRequest) SetDBInstanceName(v string) *DeletePxfuseRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeletePxfuseRequest) SetRegionId(v string) *DeletePxfuseRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePxfuseRequest) Validate() error {
	return dara.Validate(s)
}
