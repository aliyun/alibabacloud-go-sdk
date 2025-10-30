// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAxgGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *OperateAxgGroupRequest
	GetGroupId() *int64
	SetNumbers(v string) *OperateAxgGroupRequest
	GetNumbers() *string
	SetOperateType(v string) *OperateAxgGroupRequest
	GetOperateType() *string
	SetOwnerId(v int64) *OperateAxgGroupRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *OperateAxgGroupRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *OperateAxgGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OperateAxgGroupRequest
	GetResourceOwnerId() *int64
}

type OperateAxgGroupRequest struct {
	// The group ID in the AXG binding.
	//
	// You can view the group ID by using either of the following methods:
	//
	// 	- On the **Number Pool Management*	- page in the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account), filter AXG private numbers and click **Number Group G Management*	- to view the group IDs of number groups G.****
	//
	// 	- If the [CreateAxgGroup](https://help.aliyun.com/document_detail/110250.html) operation is called to create number group G, the value of **GroupId*	- in the response of the CreateAxgGroup operation is specified as the value of the request parameter **GroupId*	- of the OperateAxgGroup operation.
	//
	// >  Number group G must have one or more phone numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The phone numbers that you add to number group G. Separate the phone numbers with commas (,). You can add up to 200 phone numbers at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****,1380000****
	Numbers *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	// The type of the operation on number group G. Valid values:
	//
	// 	- **addNumbers**: adds phone numbers to number group G.
	//
	// 	- **deleteNumbers**: deletes phone numbers from number group G.
	//
	// 	- **overwriteNumbers**: replaces all phone numbers in number group G. All the original phone numbers are deleted from number group G, and new phone numbers are added to number group G.
	//
	// >
	//
	// 	- When you replace all phone numbers in number group G, the quantity of original phone numbers in number group G cannot exceed 200.
	//
	// 	- You can add up to 200 phone numbers to number group G at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// addNumbers
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC123456
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OperateAxgGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *OperateAxgGroupRequest) GetNumbers() *string {
	return s.Numbers
}

func (s *OperateAxgGroupRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateAxgGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OperateAxgGroupRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *OperateAxgGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OperateAxgGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OperateAxgGroupRequest) SetGroupId(v int64) *OperateAxgGroupRequest {
	s.GroupId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetNumbers(v string) *OperateAxgGroupRequest {
	s.Numbers = &v
	return s
}

func (s *OperateAxgGroupRequest) SetOperateType(v string) *OperateAxgGroupRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAxgGroupRequest) SetOwnerId(v int64) *OperateAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateAxgGroupRequest) SetPoolKey(v string) *OperateAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerAccount(v string) *OperateAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateAxgGroupRequest) SetResourceOwnerId(v int64) *OperateAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateAxgGroupRequest) Validate() error {
	return dara.Validate(s)
}
