// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteServiceAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteServiceAccountRequest
	GetRegionId() *string
	SetServiceAccountType(v string) *DeleteServiceAccountRequest
	GetServiceAccountType() *string
}

type DeleteServiceAccountRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// METADATA_READONLY
	ServiceAccountType *string `json:"ServiceAccountType,omitempty" xml:"ServiceAccountType,omitempty"`
}

func (s DeleteServiceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteServiceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceAccountRequest) GetServiceAccountType() *string {
	return s.ServiceAccountType
}

func (s *DeleteServiceAccountRequest) SetDBInstanceName(v string) *DeleteServiceAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteServiceAccountRequest) SetRegionId(v string) *DeleteServiceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceAccountRequest) SetServiceAccountType(v string) *DeleteServiceAccountRequest {
	s.ServiceAccountType = &v
	return s
}

func (s *DeleteServiceAccountRequest) Validate() error {
	return dara.Validate(s)
}
