// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgradeMinorVersion(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetAutoUpgradeMinorVersion() *string
	SetClientToken(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceAutoUpgradeMinorVersionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceAutoUpgradeMinorVersionRequest struct {
	// The method that is used to update the minor engine version of the instance. Valid values:
	//
	// 	- **Auto:*	- automatic update.
	//
	// 	- **Manual**: manual update. ApsaraDB RDS automatically updates the current minor engine version of the instance only when the current minor engine version is phased out.
	//
	// This parameter is required.
	//
	// example:
	//
	// Auto
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
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
	// rm-uf6wjk5xxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetAutoUpgradeMinorVersion(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetClientToken(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetDBInstanceId(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetOwnerId(v int64) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceAutoUpgradeMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
