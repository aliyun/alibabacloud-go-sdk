// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHADiagnoseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyHADiagnoseConfigRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyHADiagnoseConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHADiagnoseConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHADiagnoseConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHADiagnoseConfigRequest
	GetResourceOwnerId() *int64
	SetTcpConnectionType(v string) *ModifyHADiagnoseConfigRequest
	GetTcpConnectionType() *string
}

type ModifyHADiagnoseConfigRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
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
	// The availability check method of the instance. Valid values:
	//
	// 	- **SHORT**: Alibaba Cloud uses short-lived connections to check the availability of the instance.
	//
	// 	- **LONG**: Alibaba Cloud uses persistent connections to check the availability of the instance.
	//
	// example:
	//
	// SHORT
	TcpConnectionType *string `json:"TcpConnectionType,omitempty" xml:"TcpConnectionType,omitempty"`
}

func (s ModifyHADiagnoseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHADiagnoseConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyHADiagnoseConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyHADiagnoseConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHADiagnoseConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHADiagnoseConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHADiagnoseConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHADiagnoseConfigRequest) GetTcpConnectionType() *string {
	return s.TcpConnectionType
}

func (s *ModifyHADiagnoseConfigRequest) SetDBInstanceId(v string) *ModifyHADiagnoseConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) SetOwnerId(v int64) *ModifyHADiagnoseConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) SetRegionId(v string) *ModifyHADiagnoseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) SetResourceOwnerAccount(v string) *ModifyHADiagnoseConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) SetResourceOwnerId(v int64) *ModifyHADiagnoseConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) SetTcpConnectionType(v string) *ModifyHADiagnoseConfigRequest {
	s.TcpConnectionType = &v
	return s
}

func (s *ModifyHADiagnoseConfigRequest) Validate() error {
	return dara.Validate(s)
}
