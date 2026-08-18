// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfuseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreatePxfuseRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreatePxfuseRequest
	GetRegionId() *string
}

type CreatePxfuseRequest struct {
	// The instance ID.
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

func (s CreatePxfuseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfuseRequest) GoString() string {
	return s.String()
}

func (s *CreatePxfuseRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfuseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePxfuseRequest) SetDBInstanceName(v string) *CreatePxfuseRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfuseRequest) SetRegionId(v string) *CreatePxfuseRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePxfuseRequest) Validate() error {
	return dara.Validate(s)
}
