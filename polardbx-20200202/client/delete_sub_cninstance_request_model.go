// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCNInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteSubCNInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteSubCNInstanceRequest
	GetRegionId() *string
}

type DeleteSubCNInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSubCNInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCNInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubCNInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteSubCNInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSubCNInstanceRequest) SetDBInstanceName(v string) *DeleteSubCNInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteSubCNInstanceRequest) SetRegionId(v string) *DeleteSubCNInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubCNInstanceRequest) Validate() error {
	return dara.Validate(s)
}
