// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubscriptionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySubscriptionDetailRequest
	GetOwnerId() *int64
	SetPhoneNoX(v string) *QuerySubscriptionDetailRequest
	GetPhoneNoX() *string
	SetPoolKey(v string) *QuerySubscriptionDetailRequest
	GetPoolKey() *string
	SetProductType(v string) *QuerySubscriptionDetailRequest
	GetProductType() *string
	SetResourceOwnerAccount(v string) *QuerySubscriptionDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySubscriptionDetailRequest
	GetResourceOwnerId() *int64
	SetSubsId(v string) *QuerySubscriptionDetailRequest
	GetSubsId() *string
}

type QuerySubscriptionDetailRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private number in the binding, that is, phone number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13900001234
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Valid values:
	//
	// 	- **AXB_170**
	//
	// 	- **AXN_170**
	//
	// 	- **AXN_95**
	//
	// 	- **AXN_EXTENSION_REUSE**
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
	// The binding ID.
	//
	// Log on to the Phone Number Protection console, choose **Number and Number Pool*	- > **Number Management**. On the Number Management page, select the desired record and click Details to view the binding ID. Alternatively, you can view the value of the **SubsId*	- parameter returned by an API operation for a phone number binding such as [BindAxb](https://help.aliyun.com/document_detail/110248.html). The value of this parameter indicates a binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000076879****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubscriptionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySubscriptionDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySubscriptionDetailRequest) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *QuerySubscriptionDetailRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *QuerySubscriptionDetailRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySubscriptionDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySubscriptionDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySubscriptionDetailRequest) GetSubsId() *string {
	return s.SubsId
}

func (s *QuerySubscriptionDetailRequest) SetOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPhoneNoX(v string) *QuerySubscriptionDetailRequest {
	s.PhoneNoX = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetPoolKey(v string) *QuerySubscriptionDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetProductType(v string) *QuerySubscriptionDetailRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerAccount(v string) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetResourceOwnerId(v int64) *QuerySubscriptionDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) SetSubsId(v string) *QuerySubscriptionDetailRequest {
	s.SubsId = &v
	return s
}

func (s *QuerySubscriptionDetailRequest) Validate() error {
	return dara.Validate(s)
}
