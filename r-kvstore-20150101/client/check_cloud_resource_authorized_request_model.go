// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckCloudResourceAuthorizedRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest
	GetRoleArn() *string
	SetSecurityToken(v string) *CheckCloudResourceAuthorizedRequest
	GetSecurityToken() *string
}

type CheckCloudResourceAuthorizedRequest struct {
	// The instance ID. You can call [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) to query the instance ID.
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
	// The global resource descriptor Alibaba Cloud Resource Name (ARN) of the role to be authorized. After the role is granted authorization, the related KMS can be used. Format: `acs:ram::$accountID:role/$roleName`.
	//
	// > 	- `$accountID`: the ID of the Alibaba Cloud account. Log on to the Alibaba Cloud Management Console, move the mouse pointer over the profile picture in the upper-right corner, and then click **Security Settings*	- to view the account ID.
	//
	// > 	- `$roleName`: the name of the RAM role. Set this value to **AliyunRdsInstanceEncryptionDefaultRole**.
	//
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckCloudResourceAuthorizedRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckCloudResourceAuthorizedRequest) SetInstanceId(v string) *CheckCloudResourceAuthorizedRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRoleArn(v string) *CheckCloudResourceAuthorizedRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetSecurityToken(v string) *CheckCloudResourceAuthorizedRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
