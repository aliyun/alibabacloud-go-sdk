// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceEngineVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeDBInstanceEngineVersionRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *UpgradeDBInstanceEngineVersionRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest
	GetResourceOwnerId() *int64
}

type UpgradeDBInstanceEngineVersionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective time. Valid values:
	//
	// 	- **Immediate**: This is the default value.
	//
	// 	- **MaintainTime**: The effective time is within the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The major engine version that the new instance runs. Valid values:
	//
	// 	- **8.0**
	//
	// 	- **5.7**
	//
	// 	- **5.6**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetClientToken(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEffectiveTime(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EngineVersion = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) Validate() error {
	return dara.Validate(s)
}
