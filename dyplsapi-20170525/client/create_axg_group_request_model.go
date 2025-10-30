// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAxgGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateAxgGroupRequest
	GetName() *string
	SetNumbers(v string) *CreateAxgGroupRequest
	GetNumbers() *string
	SetOwnerId(v int64) *CreateAxgGroupRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *CreateAxgGroupRequest
	GetPoolKey() *string
	SetRemark(v string) *CreateAxgGroupRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateAxgGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAxgGroupRequest
	GetResourceOwnerId() *int64
}

type CreateAxgGroupRequest struct {
	// The name of number group G. If the name of number group G is not specified, the ID of number group G is used as the name of number group G.
	//
	// >  The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// Aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone numbers that you add to number group G. Separate multiple phone numbers with commas (,). A maximum of 200 phone numbers can be added to number group G.
	//
	// example:
	//
	// 1390000****,1380000****
	Numbers *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The remarks of number group G. The value must be 0 to 256 characters in length.
	//
	// example:
	//
	// test
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAxgGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateAxgGroupRequest) GetNumbers() *string {
	return s.Numbers
}

func (s *CreateAxgGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAxgGroupRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *CreateAxgGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateAxgGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAxgGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAxgGroupRequest) SetName(v string) *CreateAxgGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateAxgGroupRequest) SetNumbers(v string) *CreateAxgGroupRequest {
	s.Numbers = &v
	return s
}

func (s *CreateAxgGroupRequest) SetOwnerId(v int64) *CreateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAxgGroupRequest) SetPoolKey(v string) *CreateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateAxgGroupRequest) SetRemark(v string) *CreateAxgGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerAccount(v string) *CreateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAxgGroupRequest) SetResourceOwnerId(v int64) *CreateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAxgGroupRequest) Validate() error {
	return dara.Validate(s)
}
