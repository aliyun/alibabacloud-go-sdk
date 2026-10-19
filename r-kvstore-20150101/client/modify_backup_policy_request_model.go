// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest
	GetBackupRetentionPeriod() *int32
	SetEnableBackupLog(v int32) *ModifyBackupPolicyRequest
	GetEnableBackupLog() *int32
	SetInstanceId(v string) *ModifyBackupPolicyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyBackupPolicyRequest
	GetSecurityToken() *string
}

type ModifyBackupPolicyRequest struct {
	// The number of days for which data backups are retained. Default value: 7. Valid values: 7 to 730.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Specifies whether to enable the data flashback feature. Valid values:
	//
	// 	- **1**: Enabled. Make sure that append-only file (AOF) persistence is enabled in the parameter settings of the instance (`appendonly` is set to `yes`). The data flashback feature is available only after AOF persistence is enabled.
	//
	// 	- **0**: Disabled. This is the default value.
	//
	// > This parameter is supported only by Tair (Enterprise Edition) memory-optimized and persistent memory-optimized instances. For more information, see [Data flashback](https://help.aliyun.com/document_detail/148479.html).
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// > You can specify multiple values. Separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time. Specify the time in the <i>HH:mm</i>Z-<i>HH:mm</i>Z format (UTC).
	//
	// > The time must be on the hour, and the interval must be 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07:00Z-08:00Z
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetEnableBackupLog() *int32 {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v int32) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetInstanceId(v string) *ModifyBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetSecurityToken(v string) *ModifyBackupPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
