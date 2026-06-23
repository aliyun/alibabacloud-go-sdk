// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpgradeMinorVersionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *UpgradeMinorVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeMinorVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest
	GetResourceOwnerId() *int64
	SetUpgradeImmediately(v bool) *UpgradeMinorVersionRequest
	GetUpgradeImmediately() *bool
	SetUpgradeTime(v string) *UpgradeMinorVersionRequest
	GetUpgradeTime() *string
	SetUpgradeVersion(v string) *UpgradeMinorVersionRequest
	GetUpgradeVersion() *string
}

type UpgradeMinorVersionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to upgrade the minor engine version immediately. Valid values:
	//
	// - **true**: The minor engine version is upgraded immediately.
	//
	// - **false**: The minor engine version is upgraded at a specified time or within the maintenance window.
	//
	// > The **UpgradeTime*	- parameter is required if you want to upgrade the minor engine version at a specified time.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	UpgradeImmediately *bool `json:"UpgradeImmediately,omitempty" xml:"UpgradeImmediately,omitempty"`
	// The time to perform the upgrade. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > If you do not specify this parameter, the ApsaraDB for ClickHouse cluster is upgraded within the maintenance window.
	//
	// example:
	//
	// 2022-08-07T16:38Z
	UpgradeTime *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
	// The minor engine version to which you want to upgrade.
	//
	// > If you do not specify this parameter, the cluster is upgraded to the latest minor engine version.
	//
	// example:
	//
	// 1.37.0
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpgradeMinorVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeMinorVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeMinorVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeMinorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeMinorVersionRequest) GetUpgradeImmediately() *bool {
	return s.UpgradeImmediately
}

func (s *UpgradeMinorVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeMinorVersionRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *UpgradeMinorVersionRequest) SetDBClusterId(v string) *UpgradeMinorVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeImmediately(v bool) *UpgradeMinorVersionRequest {
	s.UpgradeImmediately = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeTime(v string) *UpgradeMinorVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeVersion(v string) *UpgradeMinorVersionRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *UpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
