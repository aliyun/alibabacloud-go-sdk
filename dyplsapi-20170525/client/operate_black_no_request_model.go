// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBlackNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackNo(v string) *OperateBlackNoRequest
	GetBlackNo() *string
	SetOperateType(v string) *OperateBlackNoRequest
	GetOperateType() *string
	SetOwnerId(v int64) *OperateBlackNoRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *OperateBlackNoRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *OperateBlackNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OperateBlackNoRequest
	GetResourceOwnerId() *int64
	SetTips(v string) *OperateBlackNoRequest
	GetTips() *string
}

type OperateBlackNoRequest struct {
	// The phone number to be added to or deleted from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 150****0000
	BlackNo *string `json:"BlackNo,omitempty" xml:"BlackNo,omitempty"`
	// The type of the operation on the phone number. Valid values:
	//
	// 	- **AddBlack**: adds the phone number to the blacklist.
	//
	// 	- **DeleteBlack**: deletes the phone number from the blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// AddBlack
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description.
	//
	// example:
	//
	// abcdef
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s OperateBlackNoRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateBlackNoRequest) GoString() string {
	return s.String()
}

func (s *OperateBlackNoRequest) GetBlackNo() *string {
	return s.BlackNo
}

func (s *OperateBlackNoRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateBlackNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OperateBlackNoRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *OperateBlackNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OperateBlackNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OperateBlackNoRequest) GetTips() *string {
	return s.Tips
}

func (s *OperateBlackNoRequest) SetBlackNo(v string) *OperateBlackNoRequest {
	s.BlackNo = &v
	return s
}

func (s *OperateBlackNoRequest) SetOperateType(v string) *OperateBlackNoRequest {
	s.OperateType = &v
	return s
}

func (s *OperateBlackNoRequest) SetOwnerId(v int64) *OperateBlackNoRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetPoolKey(v string) *OperateBlackNoRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerAccount(v string) *OperateBlackNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateBlackNoRequest) SetResourceOwnerId(v int64) *OperateBlackNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateBlackNoRequest) SetTips(v string) *OperateBlackNoRequest {
	s.Tips = &v
	return s
}

func (s *OperateBlackNoRequest) Validate() error {
	return dara.Validate(s)
}
