// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetInstanceId(v string) *ModifyAccountDescriptionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyAccountDescriptionRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *ModifyAccountDescriptionRequest
	GetSourceBiz() *string
}

type ModifyAccountDescriptionRequest struct {
	// The account description to set.
	//
	// 	- The description must start with a Chinese character or an English letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain Chinese characters, English letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testescription
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name. You can call [DescribeAccounts](https://help.aliyun.com/document_detail/473816.html) to obtain the account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The instance ID.
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
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is used only for internal maintenance. You do not need to specify this parameter.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAccountDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountDescriptionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAccountDescriptionRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetSecurityToken(v string) *ModifyAccountDescriptionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetSourceBiz(v string) *ModifyAccountDescriptionRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
