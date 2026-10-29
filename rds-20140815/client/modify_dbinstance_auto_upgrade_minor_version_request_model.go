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
	// The minor version upgrade method for the instance. Valid values:
	//
	// 	- **Auto**: Automatically upgrades the minor version.
	//
	// 	- **Manual**: Does not automatically upgrade the minor version. The minor version is forcibly upgraded only when the current version is retired.
	//
	// This parameter is required.
	//
	// example:
	//
	// Auto
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can be up to 64 characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
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
