// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterTDERequest
	GetDBClusterId() *string
	SetEnableAutomaticRotation(v string) *ModifyDBClusterTDERequest
	GetEnableAutomaticRotation() *string
	SetEncryptNewTables(v string) *ModifyDBClusterTDERequest
	GetEncryptNewTables() *string
	SetEncryptionKey(v string) *ModifyDBClusterTDERequest
	GetEncryptionKey() *string
	SetOwnerAccount(v string) *ModifyDBClusterTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterTDERequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterTDERequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *ModifyDBClusterTDERequest
	GetRoleArn() *string
	SetTDEStatus(v string) *ModifyDBClusterTDERequest
	GetTDEStatus() *string
}

type ModifyDBClusterTDERequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to allow the TDE key of the cluster to be automatically rotated within the next maintenance window after a lapse of the rotation period when a change in the KMS key version is detected. This parameter is supported only for custom keys. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is supported only for a PolarDB for PostgreSQL or PolarDB for PostgreSQL (Compatible with Oracle) cluster.
	//
	// example:
	//
	// false
	EnableAutomaticRotation *string `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// Specifies whether to enable automatic encryption for new tables. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// >  This parameter takes effect only for a PolarDB for MySQL cluster.
	//
	// example:
	//
	// ON
	EncryptNewTables *string `json:"EncryptNewTables,omitempty" xml:"EncryptNewTables,omitempty"`
	// The ID of the custom key.
	//
	// example:
	//
	// 749c1df7-****-****-****-*********
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role. A RAM role is a virtual identity that you can create within your Alibaba Cloud account. For more information, see [RAM role overview](https://help.aliyun.com/document_detail/93689.html).
	//
	// example:
	//
	// acs:ram::1406926*****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// Modifies the TDE status. Set the value to **Enable**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s ModifyDBClusterTDERequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDERequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterTDERequest) GetEnableAutomaticRotation() *string {
	return s.EnableAutomaticRotation
}

func (s *ModifyDBClusterTDERequest) GetEncryptNewTables() *string {
	return s.EncryptNewTables
}

func (s *ModifyDBClusterTDERequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyDBClusterTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterTDERequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ModifyDBClusterTDERequest) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *ModifyDBClusterTDERequest) SetDBClusterId(v string) *ModifyDBClusterTDERequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetEnableAutomaticRotation(v string) *ModifyDBClusterTDERequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetEncryptNewTables(v string) *ModifyDBClusterTDERequest {
	s.EncryptNewTables = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetEncryptionKey(v string) *ModifyDBClusterTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetOwnerAccount(v string) *ModifyDBClusterTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetOwnerId(v int64) *ModifyDBClusterTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetResourceOwnerAccount(v string) *ModifyDBClusterTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetResourceOwnerId(v int64) *ModifyDBClusterTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetRoleArn(v string) *ModifyDBClusterTDERequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetTDEStatus(v string) *ModifyDBClusterTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyDBClusterTDERequest) Validate() error {
	return dara.Validate(s)
}
