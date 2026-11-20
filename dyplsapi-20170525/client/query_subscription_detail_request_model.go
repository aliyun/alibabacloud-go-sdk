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
	// The privacy number in the binding relationship, which is the X number.
	//
	// Logon to the privacy protection console and go to the [Number Management](https://dyplsnext.console.aliyun.com/number/list) page to obtain the X number. Alternatively, check the **SecretNo*	- in the response parameters when you invoke number binding API operations such as [BindAxb](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-bindaxb).
	//
	// > When specifying the X number, make sure it is associated with the **SubsId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1111
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
	// The key of the number pool.
	//
	// Log on to the privacy protection console and go to the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool?spm=a2c4g.11186623.0.0.7e6449bbtiHP8m) page to obtain the **Number Pool Key**.
	//
	// > This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC10000012345****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Valid values:
	//
	// - **AXB_170**
	//
	// - **AXN_170**
	//
	// - **AXN_95**
	//
	// - **AXN_EXTENSION_REUSE**
	//
	// > - This parameter is applicable to former Alibaba Dayu users. Alibaba Cloud users can ignore this parameter.
	//
	// > - This parameter is required when **PoolKey*	- is left empty.
	//
	// example:
	//
	// AXB_170
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding relationship ID.
	//
	// Logon to the privacy protection console, go to the [Number Management](https://dyplsnext.console.aliyun.com/number/list) page, filter the numbers, and click **Details*	- to view the binding relationship ID. Alternatively, check the **SubsId*	- in the response parameters when you invoke number binding API operations such as [BindAxb](https://www.alibabacloud.com/help/en/pnp/developer-reference/api-dyplsapi-2017-05-25-bindaxb).
	//
	// > When specifying the binding relationship ID, make sure it is associated with the **PhoneNoX*	- parameter.
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
