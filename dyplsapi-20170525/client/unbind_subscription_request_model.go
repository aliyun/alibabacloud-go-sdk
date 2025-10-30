// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *UnbindSubscriptionRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *UnbindSubscriptionRequest
	GetPoolKey() *string
	SetProductType(v string) *UnbindSubscriptionRequest
	GetProductType() *string
	SetResourceOwnerAccount(v string) *UnbindSubscriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindSubscriptionRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *UnbindSubscriptionRequest
	GetSecretNo() *string
	SetSubsId(v string) *UnbindSubscriptionRequest
	GetSubsId() *string
}

type UnbindSubscriptionRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Fixed value: **AXB_170**.
	//
	// >
	//
	// 	- This parameter is applicable to the original key accounts of Alibaba Cloud. This parameter can be ignored for Alibaba Cloud users.
	//
	// 	- This parameter is required when **PoolKey*	- is left empty.
	//
	// example:
	//
	// AXB_170
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private number, that is, phone number X specified in an API operation for a phone number binding such as [BindAXG](https://help.aliyun.com/document_detail/110249.html) or automatically assigned after such an operation is called.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The binding ID.
	//
	// Log on to the Phone Number Protection console, choose **Number and Number Pool*	- > **Number Management**. On the Number Management page, select the desired record and click Details to view the binding ID. Alternatively, you can view the value of the **SubsId*	- parameter returned by an API operation for a phone number binding such as BindAxb. The value of this parameter indicates a binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1************2
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s UnbindSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindSubscriptionRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *UnbindSubscriptionRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UnbindSubscriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindSubscriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindSubscriptionRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *UnbindSubscriptionRequest) GetSubsId() *string {
	return s.SubsId
}

func (s *UnbindSubscriptionRequest) SetOwnerId(v int64) *UnbindSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetPoolKey(v string) *UnbindSubscriptionRequest {
	s.PoolKey = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetProductType(v string) *UnbindSubscriptionRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerAccount(v string) *UnbindSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetResourceOwnerId(v int64) *UnbindSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSecretNo(v string) *UnbindSubscriptionRequest {
	s.SecretNo = &v
	return s
}

func (s *UnbindSubscriptionRequest) SetSubsId(v string) *UnbindSubscriptionRequest {
	s.SubsId = &v
	return s
}

func (s *UnbindSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
