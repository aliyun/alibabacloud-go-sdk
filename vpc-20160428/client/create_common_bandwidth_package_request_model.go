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
	// The peak bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// <props="intl"><ph>Default value range: **1*	- to **1000**. Default value: **1**.</ph>
	//
	// <props="china">
	//
	// - If **InternetChargeType*	- is set to **PayByBandwidth**, which indicates that the billable method of the Internet Shared Bandwidth instance is pay-by-bandwidth, the default value range of **Bandwidth*	- is **2*	- to **20000**.
	//
	// - If **InternetChargeType*	- is set to **PayBy95**, which indicates that the billable method of the Internet Shared Bandwidth instance is pay-by-enhanced-95th-percentile, the default value range of **Bandwidth*	- is **200*	- to **20000**.
	//
	// - If **InternetChargeType*	- is set to **PayByDominantTraffic**, which indicates that the billable method of the Internet Shared Bandwidth instance is pay-by-dominant-traffic, the default value range of **Bandwidth*	- is **1*	- to **2000**.
	//
	//  Default value: **1000**.
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
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
	// - **BGP*	- (default): BGP (multi-ISP) lines. All regions support BGP (multi-ISP) lines.
	//
	// - **BGP_PRO**: BGP (multi-ISP) premium lines. Currently, only the Hong Kong (China), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions support BGP (multi-ISP) premium Internet Shared Bandwidth instances.
	//
	// If you are a single-ISP bandwidth whitelist user, you can also select the following types:
	//
	// - **ChinaTelecom**: China Telecom
	//
	// - **ChinaUnicom**: China Unicom
	//
	// - **ChinaMobile**: China Mobile
	//
	// - **ChinaTelecom_L2**: China Telecom L2
	//
	// - **ChinaUnicom_L2**: China Unicom L2
	//
	// - **ChinaMobile_L2**: China Mobile L2
	//
	// If you are a Finance Cloud user in the China (Hangzhou) region, this parameter is required. Set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The billable method of the Internet Shared Bandwidth instance. Valid values:
	//
	// <props="intl">**PayByTraffic*	- (pay-by-data-transfer).
	//
	// <props="china">
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// - **PayBy95**: pay-by-enhanced-95th-percentile.
	//
	// - **PayByDominantTraffic**: pay-by-dominant-traffic.
	//
	// .
	//
	// example:
	//
	// 中国站示例值：PayByBandwidth，国际站示例值：PayByTraffic
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
	// The minimum bandwidth commitment percentage of the Internet Shared Bandwidth instance. Set the value to **20**.
	//
	//  <props="china"><ph>This parameter is required when **InternetChargeType*	- is set to **PayBy95**.</ph>
	//
	// >This parameter is supported only on the China site.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security protection level.
	//
	// - If you do not set this parameter, Anti-DDoS Origin Basic is used by default.
	//
	// - If you set this parameter to **AntiDDoS_Enhanced**, Anti-DDoS Origin Enhanced is used.
	//
	// <props="china"><ph>You can set this parameter when **InternetChargeType*	- is set to **PayBy95**.</ph>
	//
	// You can specify up to 10 security protection levels.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// AntiDDoS_Enhanced
	SecurityProtectionTypes []*string `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	// The list of tags for the Internet Shared Bandwidth instance.
	Tag []*CreateCommonBandwidthPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone of the Internet Shared Bandwidth instance.
	//
	// This parameter is required when you create an Internet Shared Bandwidth instance for a CloudBox.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCommonBandwidthPackageRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
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
