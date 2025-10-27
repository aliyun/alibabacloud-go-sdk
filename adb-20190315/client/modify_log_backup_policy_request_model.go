// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyLogBackupPolicyRequest
	GetDBClusterId() *string
	SetEnableBackupLog(v string) *ModifyLogBackupPolicyRequest
	GetEnableBackupLog() *string
	SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest
	GetLogBackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLogBackupPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyLogBackupPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyLogBackupPolicyRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which to retain log backup files. Valid values: 7 to 730.
	//
	// >  If you do not specify this parameter, the default value 7 is used.
	//
	// example:
	//
	// 30
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm4f7oger****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyLogBackupPolicyRequest) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *ModifyLogBackupPolicyRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyLogBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLogBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLogBackupPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyLogBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLogBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyLogBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceGroupId(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
