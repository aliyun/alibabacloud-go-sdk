// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateServiceAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateServiceAccountRequest
	GetRegionId() *string
	SetServiceAccountType(v string) *CreateServiceAccountRequest
	GetServiceAccountType() *string
}

type CreateServiceAccountRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service account.
	//
	// This parameter is required.
	//
	// example:
	//
	// METADATA_READONLY
	ServiceAccountType *string `json:"ServiceAccountType,omitempty" xml:"ServiceAccountType,omitempty"`
}

func (s CreateServiceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateServiceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceAccountRequest) GetServiceAccountType() *string {
	return s.ServiceAccountType
}

func (s *CreateServiceAccountRequest) SetDBInstanceName(v string) *CreateServiceAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateServiceAccountRequest) SetRegionId(v string) *CreateServiceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceAccountRequest) SetServiceAccountType(v string) *CreateServiceAccountRequest {
	s.ServiceAccountType = &v
	return s
}

func (s *CreateServiceAccountRequest) Validate() error {
	return dara.Validate(s)
}
