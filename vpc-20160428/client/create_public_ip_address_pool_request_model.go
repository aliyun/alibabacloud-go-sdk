// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublicIpAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreatePublicIpAddressPoolRequest
	GetBizType() *string
	SetClientToken(v string) *CreatePublicIpAddressPoolRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePublicIpAddressPoolRequest
	GetDescription() *string
	SetDryRun(v bool) *CreatePublicIpAddressPoolRequest
	GetDryRun() *bool
	SetIsp(v string) *CreatePublicIpAddressPoolRequest
	GetIsp() *string
	SetName(v string) *CreatePublicIpAddressPoolRequest
	GetName() *string
	SetOwnerAccount(v string) *CreatePublicIpAddressPoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePublicIpAddressPoolRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreatePublicIpAddressPoolRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePublicIpAddressPoolRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreatePublicIpAddressPoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePublicIpAddressPoolRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionTypes(v []*string) *CreatePublicIpAddressPoolRequest
	GetSecurityProtectionTypes() []*string
	SetTag(v []*CreatePublicIpAddressPoolRequestTag) *CreatePublicIpAddressPoolRequest
	GetTag() []*CreatePublicIpAddressPoolRequestTag
	SetZones(v []*string) *CreatePublicIpAddressPoolRequest
	GetZones() []*string
}

type CreatePublicIpAddressPoolRequest struct {
	// The service type of the IP address pool. Valid values:
	//
	// 	- **CloudBox*	- Only cloud box users can select this type.
	//
	// 	- **Default**: This is the default value.
	//
	// example:
	//
	// Default
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a value, and you must make sure that each request has a unique token value. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IP address pool.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to precheck only this request. Valid values:
	//
	// 	- **true**: prechecks the request without creating an IP address pool. The system checks the required parameters, request format, and service limits. If the request fails to pass the precheck, an error code is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. This is the default value. If the request passes the precheck, a 2xx HTTP status code is returned and the IP address pool is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default)
	//
	// 	- **BGP_PRO**
	//
	// For more information about BGP (Multi-ISP) lines and BGP (Multi-ISP) Pro lines, see the "Line types" section in the [What is EIP?](https://help.aliyun.com/document_detail/32321.html) topic.
	//
	// 	- If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	//
	//     	- **ChinaTelecom**
	//
	//     	- **ChinaUnicom**
	//
	//     	- **ChinaMobile**
	//
	//     	- **ChinaTelecom_L2**
	//
	//     	- **ChinaUnicom_L2**
	//
	//     	- **ChinaMobile_L2**
	//
	// 	- If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the IP address pool.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to create the IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IP address pool belongs.
	//
	// example:
	//
	// rg-acfmxazb4pcdvf****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The editions of Anti-DDoS.
	//
	// - If you do not specify this parameter, Anti-DDoS Origin Basic is used.
	//
	// - If you set the parameter to AntiDDoS_Enhanced, Anti-DDoS Pro/Premium is used.
	SecurityProtectionTypes []*string `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	// The tag of the resource.
	Tag []*CreatePublicIpAddressPoolRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone of the IP address pool. If you set **BizType*	- to **CloudBox**, this parameter is required.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s CreatePublicIpAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicIpAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *CreatePublicIpAddressPoolRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreatePublicIpAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePublicIpAddressPoolRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePublicIpAddressPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreatePublicIpAddressPoolRequest) GetIsp() *string {
	return s.Isp
}

func (s *CreatePublicIpAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *CreatePublicIpAddressPoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePublicIpAddressPoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePublicIpAddressPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePublicIpAddressPoolRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePublicIpAddressPoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePublicIpAddressPoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePublicIpAddressPoolRequest) GetSecurityProtectionTypes() []*string {
	return s.SecurityProtectionTypes
}

func (s *CreatePublicIpAddressPoolRequest) GetTag() []*CreatePublicIpAddressPoolRequestTag {
	return s.Tag
}

func (s *CreatePublicIpAddressPoolRequest) GetZones() []*string {
	return s.Zones
}

func (s *CreatePublicIpAddressPoolRequest) SetBizType(v string) *CreatePublicIpAddressPoolRequest {
	s.BizType = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetClientToken(v string) *CreatePublicIpAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetDescription(v string) *CreatePublicIpAddressPoolRequest {
	s.Description = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetDryRun(v bool) *CreatePublicIpAddressPoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetIsp(v string) *CreatePublicIpAddressPoolRequest {
	s.Isp = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetName(v string) *CreatePublicIpAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetOwnerAccount(v string) *CreatePublicIpAddressPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetOwnerId(v int64) *CreatePublicIpAddressPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetRegionId(v string) *CreatePublicIpAddressPoolRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetResourceGroupId(v string) *CreatePublicIpAddressPoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetResourceOwnerAccount(v string) *CreatePublicIpAddressPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetResourceOwnerId(v int64) *CreatePublicIpAddressPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetSecurityProtectionTypes(v []*string) *CreatePublicIpAddressPoolRequest {
	s.SecurityProtectionTypes = v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetTag(v []*CreatePublicIpAddressPoolRequestTag) *CreatePublicIpAddressPoolRequest {
	s.Tag = v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) SetZones(v []*string) *CreatePublicIpAddressPoolRequest {
	s.Zones = v
	return s
}

func (s *CreatePublicIpAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePublicIpAddressPoolRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePublicIpAddressPoolRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicIpAddressPoolRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePublicIpAddressPoolRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePublicIpAddressPoolRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePublicIpAddressPoolRequestTag) SetKey(v string) *CreatePublicIpAddressPoolRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequestTag) SetValue(v string) *CreatePublicIpAddressPoolRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePublicIpAddressPoolRequestTag) Validate() error {
	return dara.Validate(s)
}
