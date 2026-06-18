// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateStoragePoolRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateStoragePoolRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStoragePoolRequest
	GetResourceGroupId() *string
	SetStoragePoolDNList(v string) *CreateStoragePoolRequest
	GetStoragePoolDNList() *string
	SetStoragePoolName(v string) *CreateStoragePoolRequest
	GetStoragePoolName() *string
}

type CreateStoragePoolRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-bjxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of storage pool DNs.
	//
	// example:
	//
	// node6
	StoragePoolDNList *string `json:"StoragePoolDNList,omitempty" xml:"StoragePoolDNList,omitempty"`
	// The name of the storage pool.
	//
	// example:
	//
	// test
	StoragePoolName *string `json:"StoragePoolName,omitempty" xml:"StoragePoolName,omitempty"`
}

func (s CreateStoragePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateStoragePoolRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateStoragePoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStoragePoolRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStoragePoolRequest) GetStoragePoolDNList() *string {
	return s.StoragePoolDNList
}

func (s *CreateStoragePoolRequest) GetStoragePoolName() *string {
	return s.StoragePoolName
}

func (s *CreateStoragePoolRequest) SetDBInstanceName(v string) *CreateStoragePoolRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateStoragePoolRequest) SetRegionId(v string) *CreateStoragePoolRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStoragePoolRequest) SetResourceGroupId(v string) *CreateStoragePoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStoragePoolRequest) SetStoragePoolDNList(v string) *CreateStoragePoolRequest {
	s.StoragePoolDNList = &v
	return s
}

func (s *CreateStoragePoolRequest) SetStoragePoolName(v string) *CreateStoragePoolRequest {
	s.StoragePoolName = &v
	return s
}

func (s *CreateStoragePoolRequest) Validate() error {
	return dara.Validate(s)
}
