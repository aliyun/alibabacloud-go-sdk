// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionKey(v string) *ModifyInstanceTDERequest
	GetEncryptionKey() *string
	SetEncryptionName(v string) *ModifyInstanceTDERequest
	GetEncryptionName() *string
	SetInstanceId(v string) *ModifyInstanceTDERequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceTDERequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceTDERequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *ModifyInstanceTDERequest
	GetRoleArn() *string
	SetSecurityToken(v string) *ModifyInstanceTDERequest
	GetSecurityToken() *string
	SetTDEStatus(v string) *ModifyInstanceTDERequest
	GetTDEStatus() *string
}

type ModifyInstanceTDERequest struct {
	// The ID of the custom key. You can call the [DescribeEncryptionKeyList](https://help.aliyun.com/document_detail/473860.html) operation to query the key ID.
	//
	// >
	//
	// 	- If you do not specify this parameter, [Key Management Service (KMS)](https://help.aliyun.com/document_detail/28935.html) automatically generates a key.
	//
	// 	- To create a custom key, you can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation of the KMS API.
	//
	// example:
	//
	// ad463061-992d-4195-8a94-ed63********
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption algorithm. Default value: AES-CTR-256.
	//
	// > This parameter is available only if the **TDEStatus*	- parameter is set to **Enabled**.
	//
	// example:
	//
	// AES-CTR-256
	EncryptionName *string `json:"EncryptionName,omitempty" xml:"EncryptionName,omitempty"`
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
	// The Alibaba Cloud Resource Name (ARN) of the Resource Access Management (RAM) role that you want to attach to your Tair (Redis OSS-compatible) instance. The ARN must be in the format of `acs:ram::$accountID:role/$roleName`. After the role is attached, your Tair (Redis OSS-compatible) instance can use KMS.
	//
	// >
	//
	// 	- `$accountID`: the ID of the Alibaba Cloud account. To view the account ID, log on to the Alibaba Cloud console, move the pointer over your profile picture in the upper-right corner of the page, and then click **Security Settings**.
	//
	// 	- `$roleName`: the name of the RAM role. Replace $roleName with **AliyunRdsInstanceEncryptionDefaultRole**.
	//
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to enable TDE. Set the value to **Enabled**.
	//
	// > TDE cannot be disabled after it is enabled. Before you enable it, evaluate whether this feature affects your business. For more information, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s ModifyInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTDERequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyInstanceTDERequest) GetEncryptionName() *string {
	return s.EncryptionName
}

func (s *ModifyInstanceTDERequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceTDERequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ModifyInstanceTDERequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceTDERequest) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *ModifyInstanceTDERequest) SetEncryptionKey(v string) *ModifyInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetEncryptionName(v string) *ModifyInstanceTDERequest {
	s.EncryptionName = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetInstanceId(v string) *ModifyInstanceTDERequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetOwnerAccount(v string) *ModifyInstanceTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetOwnerId(v int64) *ModifyInstanceTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetResourceOwnerAccount(v string) *ModifyInstanceTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetResourceOwnerId(v int64) *ModifyInstanceTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetRoleArn(v string) *ModifyInstanceTDERequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetSecurityToken(v string) *ModifyInstanceTDERequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceTDERequest) SetTDEStatus(v string) *ModifyInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
