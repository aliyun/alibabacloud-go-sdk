// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarDBXInstanceNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddDNSpec(v string) *UpdatePolarDBXInstanceNodeRequest
	GetAddDNSpec() *string
	SetCNNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest
	GetCNNodeCount() *int32
	SetClientToken(v string) *UpdatePolarDBXInstanceNodeRequest
	GetClientToken() *string
	SetDBInstanceName(v string) *UpdatePolarDBXInstanceNodeRequest
	GetDBInstanceName() *string
	SetDNNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest
	GetDNNodeCount() *int32
	SetDbInstanceNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest
	GetDbInstanceNodeCount() *int32
	SetDeleteDNIds(v string) *UpdatePolarDBXInstanceNodeRequest
	GetDeleteDNIds() *string
	SetRegionId(v string) *UpdatePolarDBXInstanceNodeRequest
	GetRegionId() *string
	SetStoragePoolName(v string) *UpdatePolarDBXInstanceNodeRequest
	GetStoragePoolName() *string
}

type UpdatePolarDBXInstanceNodeRequest struct {
	AddDNSpec *string `json:"AddDNSpec,omitempty" xml:"AddDNSpec,omitempty"`
	// example:
	//
	// 2
	CNNodeCount *int32 `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 2
	DNNodeCount *int32 `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// example:
	//
	// 3
	DbInstanceNodeCount *int32  `json:"DbInstanceNodeCount,omitempty" xml:"DbInstanceNodeCount,omitempty"`
	DeleteDNIds         *string `json:"DeleteDNIds,omitempty" xml:"DeleteDNIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoragePoolName *string `json:"StoragePoolName,omitempty" xml:"StoragePoolName,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetAddDNSpec() *string {
	return s.AddDNSpec
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetCNNodeCount() *int32 {
	return s.CNNodeCount
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetDNNodeCount() *int32 {
	return s.DNNodeCount
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetDbInstanceNodeCount() *int32 {
	return s.DbInstanceNodeCount
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetDeleteDNIds() *string {
	return s.DeleteDNIds
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePolarDBXInstanceNodeRequest) GetStoragePoolName() *string {
	return s.StoragePoolName
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetAddDNSpec(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.AddDNSpec = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetCNNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest {
	s.CNNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetClientToken(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDBInstanceName(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDNNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest {
	s.DNNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDbInstanceNodeCount(v int32) *UpdatePolarDBXInstanceNodeRequest {
	s.DbInstanceNodeCount = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetDeleteDNIds(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.DeleteDNIds = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetRegionId(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) SetStoragePoolName(v string) *UpdatePolarDBXInstanceNodeRequest {
	s.StoragePoolName = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeRequest) Validate() error {
	return dara.Validate(s)
}
