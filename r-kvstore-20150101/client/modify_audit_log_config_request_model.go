// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbAudit(v bool) *ModifyAuditLogConfigRequest
	GetDbAudit() *bool
	SetInstanceId(v string) *ModifyAuditLogConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyAuditLogConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAuditLogConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest
	GetResourceOwnerId() *int64
	SetRetention(v int32) *ModifyAuditLogConfigRequest
	GetRetention() *int32
	SetSecurityToken(v string) *ModifyAuditLogConfigRequest
	GetSecurityToken() *string
}

type ModifyAuditLogConfigRequest struct {
	// Specifies whether to enable the audit log feature. Default value: true. Valid values:
	//
	// 	- **true**: enables the audit log feature.
	//
	// 	- **false**: disables the audit log feature.
	//
	// > If the instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html) or [read/write splitting architecture](https://help.aliyun.com/document_detail/62870.html), the audit log feature is enabled or disabled for both the data nodes and proxy nodes. You cannot separately enable the audit log feature for the data nodes or proxy nodes.
	//
	// example:
	//
	// true
	DbAudit *bool `json:"DbAudit,omitempty" xml:"DbAudit,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of audit logs. Valid values: **1*	- to **365**. Unit: days.
	//
	// >
	//
	// 	- This parameter is required only when the **DbAudit*	- parameter is set to **true**.
	//
	// 	- The value of this parameter takes effect for all instances in the current region.
	//
	// example:
	//
	// 10
	Retention     *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) GetDbAudit() *bool {
	return s.DbAudit
}

func (s *ModifyAuditLogConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAuditLogConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAuditLogConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAuditLogConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAuditLogConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAuditLogConfigRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyAuditLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAuditLogConfigRequest) SetDbAudit(v bool) *ModifyAuditLogConfigRequest {
	s.DbAudit = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetInstanceId(v string) *ModifyAuditLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRetention(v int32) *ModifyAuditLogConfigRequest {
	s.Retention = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetSecurityToken(v string) *ModifyAuditLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
