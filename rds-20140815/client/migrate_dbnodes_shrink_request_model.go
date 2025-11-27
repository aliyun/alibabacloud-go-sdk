// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *MigrateDBNodesShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *MigrateDBNodesShrinkRequest
	GetDBInstanceId() *string
	SetDBNodeShrink(v string) *MigrateDBNodesShrinkRequest
	GetDBNodeShrink() *string
	SetEffectiveTime(v string) *MigrateDBNodesShrinkRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *MigrateDBNodesShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateDBNodesShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateDBNodesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateDBNodesShrinkRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *MigrateDBNodesShrinkRequest
	GetSwitchTime() *string
	SetVSwitchId(v string) *MigrateDBNodesShrinkRequest
	GetVSwitchId() *string
}

type MigrateDBNodesShrinkRequest struct {
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the nodes.
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	// The time when you want the system to start the migration. Valid value:
	//
	// 	- **Immediately**: The system immediately starts the migration. This is the default value.
	//
	// 	- **MaintainTime**: The system starts the migration during the specified maintenance window.
	//
	// 	- **Specified**: The system starts the migration at the specified point in time.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies the time when the modification is performed. We recommend that you apply the specification during off-peak hours. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s MigrateDBNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBNodesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *MigrateDBNodesShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateDBNodesShrinkRequest) GetDBNodeShrink() *string {
	return s.DBNodeShrink
}

func (s *MigrateDBNodesShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateDBNodesShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateDBNodesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateDBNodesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateDBNodesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateDBNodesShrinkRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *MigrateDBNodesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *MigrateDBNodesShrinkRequest) SetClientToken(v string) *MigrateDBNodesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetDBInstanceId(v string) *MigrateDBNodesShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetDBNodeShrink(v string) *MigrateDBNodesShrinkRequest {
	s.DBNodeShrink = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetEffectiveTime(v string) *MigrateDBNodesShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetOwnerAccount(v string) *MigrateDBNodesShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetOwnerId(v int64) *MigrateDBNodesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetResourceOwnerAccount(v string) *MigrateDBNodesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetResourceOwnerId(v int64) *MigrateDBNodesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetSwitchTime(v string) *MigrateDBNodesShrinkRequest {
	s.SwitchTime = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) SetVSwitchId(v string) *MigrateDBNodesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateDBNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
