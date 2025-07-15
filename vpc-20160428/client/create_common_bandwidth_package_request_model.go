// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommonBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *CreateCommonBandwidthPackageRequest
	GetBandwidth() *int32
	SetClientToken(v string) *CreateCommonBandwidthPackageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCommonBandwidthPackageRequest
	GetDescription() *string
	SetISP(v string) *CreateCommonBandwidthPackageRequest
	GetISP() *string
	SetInternetChargeType(v string) *CreateCommonBandwidthPackageRequest
	GetInternetChargeType() *string
	SetName(v string) *CreateCommonBandwidthPackageRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCommonBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCommonBandwidthPackageRequest
	GetOwnerId() *int64
	SetRatio(v int32) *CreateCommonBandwidthPackageRequest
	GetRatio() *int32
	SetRegionId(v string) *CreateCommonBandwidthPackageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCommonBandwidthPackageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCommonBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCommonBandwidthPackageRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionTypes(v []*string) *CreateCommonBandwidthPackageRequest
	GetSecurityProtectionTypes() []*string
	SetTag(v []*CreateCommonBandwidthPackageRequestTag) *CreateCommonBandwidthPackageRequest
	GetTag() []*CreateCommonBandwidthPackageRequestTag
	SetZone(v string) *CreateCommonBandwidthPackageRequest
	GetZone() *string
}

type CreateCommonBandwidthPackageRequest struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// Valid values: **1*	- to **1000**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Internet Shared Bandwidth instance.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default) All regions support BGP (Multi-ISP).
	//
	// 	- **BGP_PRO*	- BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	//
	// 	- **ChinaTelecom**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaTelecom_L2**
	//
	// 	- **ChinaUnicom_L2**
	//
	// 	- **ChinaMobile_L2**
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The billing method of the Internet Shared Bandwidth instance. Set the value to **PayByTraffic**, which specifies the pay-by-data-transfer billing method.
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The name of the Internet Shared Bandwidth instance.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test123
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The percentage of the minimum bandwidth commitment. Set the parameter to **20**.
	//
	// > This parameter is available only on the Alibaba Cloud China site.
	//
	// example:
	//
	// 20
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// AntiDDoS_Enhanced
	SecurityProtectionTypes []*string                                 `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	Tag                     []*CreateCommonBandwidthPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// example:
	//
	// ap-southeast-1-lzdvn-cb
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateCommonBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommonBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateCommonBandwidthPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateCommonBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCommonBandwidthPackageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCommonBandwidthPackageRequest) GetISP() *string {
	return s.ISP
}

func (s *CreateCommonBandwidthPackageRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateCommonBandwidthPackageRequest) GetName() *string {
	return s.Name
}

func (s *CreateCommonBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCommonBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCommonBandwidthPackageRequest) GetRatio() *int32 {
	return s.Ratio
}

func (s *CreateCommonBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCommonBandwidthPackageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCommonBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCommonBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCommonBandwidthPackageRequest) GetSecurityProtectionTypes() []*string {
	return s.SecurityProtectionTypes
}

func (s *CreateCommonBandwidthPackageRequest) GetTag() []*CreateCommonBandwidthPackageRequestTag {
	return s.Tag
}

func (s *CreateCommonBandwidthPackageRequest) GetZone() *string {
	return s.Zone
}

func (s *CreateCommonBandwidthPackageRequest) SetBandwidth(v int32) *CreateCommonBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetClientToken(v string) *CreateCommonBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetDescription(v string) *CreateCommonBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetISP(v string) *CreateCommonBandwidthPackageRequest {
	s.ISP = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetInternetChargeType(v string) *CreateCommonBandwidthPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetName(v string) *CreateCommonBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetOwnerAccount(v string) *CreateCommonBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetOwnerId(v int64) *CreateCommonBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetRatio(v int32) *CreateCommonBandwidthPackageRequest {
	s.Ratio = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetRegionId(v string) *CreateCommonBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetResourceGroupId(v string) *CreateCommonBandwidthPackageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetResourceOwnerAccount(v string) *CreateCommonBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetResourceOwnerId(v int64) *CreateCommonBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetSecurityProtectionTypes(v []*string) *CreateCommonBandwidthPackageRequest {
	s.SecurityProtectionTypes = v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetTag(v []*CreateCommonBandwidthPackageRequestTag) *CreateCommonBandwidthPackageRequest {
	s.Tag = v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) SetZone(v string) *CreateCommonBandwidthPackageRequest {
	s.Zone = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCommonBandwidthPackageRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCommonBandwidthPackageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCommonBandwidthPackageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCommonBandwidthPackageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCommonBandwidthPackageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCommonBandwidthPackageRequestTag) SetKey(v string) *CreateCommonBandwidthPackageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequestTag) SetValue(v string) *CreateCommonBandwidthPackageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCommonBandwidthPackageRequestTag) Validate() error {
	return dara.Validate(s)
}
