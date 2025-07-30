// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest
	GetDBInstanceId() *string
	SetNetType(v string) *AllocateInstancePublicConnectionRequest
	GetNetType() *string
	SetRegionId(v string) *AllocateInstancePublicConnectionRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateInstancePublicConnectionRequest struct {
	// The prefix of the endpoint. Set the parameter to the prefix of the value of **ConnectionString**.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-h033cn****-pub-i3
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the endpoint to be applied for. Set the value to Public.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateInstancePublicConnectionRequest) GetNetType() *string {
	return s.NetType
}

func (s *AllocateInstancePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetNetType(v string) *AllocateInstancePublicConnectionRequest {
	s.NetType = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetRegionId(v string) *AllocateInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
