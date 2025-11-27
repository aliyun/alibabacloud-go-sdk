// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *StopDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *StopDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopDBInstanceRequest
	GetResourceOwnerId() *int64
}

type StopDBInstanceRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1q915x50h******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StopDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopDBInstanceRequest) SetDBInstanceId(v string) *StopDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StopDBInstanceRequest) SetOwnerId(v int64) *StopDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDBInstanceRequest) SetRegionId(v string) *StopDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopDBInstanceRequest) SetResourceOwnerAccount(v string) *StopDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopDBInstanceRequest) SetResourceOwnerId(v int64) *StopDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
