// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageId(v string) *CreateAppInstanceGroupRequest
	GetAppCenterImageId() *string
	SetAppInstanceGroupName(v string) *CreateAppInstanceGroupRequest
	GetAppInstanceGroupName() *string
	SetAppPackageType(v string) *CreateAppInstanceGroupRequest
	GetAppPackageType() *string
	SetAppPolicyId(v string) *CreateAppInstanceGroupRequest
	GetAppPolicyId() *string
	SetAuthMode(v string) *CreateAppInstanceGroupRequest
	GetAuthMode() *string
	SetAutoPay(v bool) *CreateAppInstanceGroupRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateAppInstanceGroupRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateAppInstanceGroupRequest
	GetBizRegionId() *string
	SetChargeResourceMode(v string) *CreateAppInstanceGroupRequest
	GetChargeResourceMode() *string
	SetChargeType(v string) *CreateAppInstanceGroupRequest
	GetChargeType() *string
	SetClusterId(v string) *CreateAppInstanceGroupRequest
	GetClusterId() *string
	SetNetwork(v *CreateAppInstanceGroupRequestNetwork) *CreateAppInstanceGroupRequest
	GetNetwork() *CreateAppInstanceGroupRequestNetwork
	SetNodePool(v *CreateAppInstanceGroupRequestNodePool) *CreateAppInstanceGroupRequest
	GetNodePool() *CreateAppInstanceGroupRequestNodePool
	SetPeriod(v int32) *CreateAppInstanceGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAppInstanceGroupRequest
	GetPeriodUnit() *string
	SetPreOpenAppId(v string) *CreateAppInstanceGroupRequest
	GetPreOpenAppId() *string
	SetProductType(v string) *CreateAppInstanceGroupRequest
	GetProductType() *string
	SetPromotionId(v string) *CreateAppInstanceGroupRequest
	GetPromotionId() *string
	SetRuntimePolicy(v *CreateAppInstanceGroupRequestRuntimePolicy) *CreateAppInstanceGroupRequest
	GetRuntimePolicy() *CreateAppInstanceGroupRequestRuntimePolicy
	SetSecurityPolicy(v *CreateAppInstanceGroupRequestSecurityPolicy) *CreateAppInstanceGroupRequest
	GetSecurityPolicy() *CreateAppInstanceGroupRequestSecurityPolicy
	SetSessionTimeout(v int32) *CreateAppInstanceGroupRequest
	GetSessionTimeout() *int32
	SetStoragePolicy(v *CreateAppInstanceGroupRequestStoragePolicy) *CreateAppInstanceGroupRequest
	GetStoragePolicy() *CreateAppInstanceGroupRequestStoragePolicy
	SetSubPayType(v string) *CreateAppInstanceGroupRequest
	GetSubPayType() *string
	SetUserDefinePolicy(v *CreateAppInstanceGroupRequestUserDefinePolicy) *CreateAppInstanceGroupRequest
	GetUserDefinePolicy() *CreateAppInstanceGroupRequestUserDefinePolicy
	SetUserGroupIds(v []*string) *CreateAppInstanceGroupRequest
	GetUserGroupIds() []*string
	SetUserInfo(v *CreateAppInstanceGroupRequestUserInfo) *CreateAppInstanceGroupRequest
	GetUserInfo() *CreateAppInstanceGroupRequestUserInfo
	SetUsers(v []*string) *CreateAppInstanceGroupRequest
	GetUsers() []*string
	SetVideoPolicy(v *CreateAppInstanceGroupRequestVideoPolicy) *CreateAppInstanceGroupRequest
	GetVideoPolicy() *CreateAppInstanceGroupRequestVideoPolicy
}

type CreateAppInstanceGroupRequest struct {
	// The image ID of the application. To obtain the image ID, log on to the [App Streaming console](https://appstreaming.console.aliyun.com/). In the left-side navigation pane, choose **Maintenance*	- > **Custom Images*	- or Maintenance > **System Images**.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// Package type.
	//
	// example:
	//
	// browser.package.5.250.appstreaming.general.basic
	AppPackageType *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// pg-0clfzcy0adpcf****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The authentication mode of the delivery group.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: manual payment. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: manual payment. This is the default value.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The sales mode.
	//
	// Valid value:
	//
	// 	- Node: by resource
	//
	// This parameter is required.
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Cluster ID.
	//
	// example:
	//
	// cls-d39iq73l5c0a8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The network settings.
	//
	// >  If you want to use this parameter, submit a ticket.
	Network *CreateAppInstanceGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The node pool object.
	NodePool *CreateAppInstanceGroupRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// The subscription duration of resources. This parameter is required if you set `ChargeType` to `PrePaid`. The unit of this parameter is specified by `PeriodUnit`.
	//
	// 	- Valid value if you set `PeriodUnit` to `Week`:
	//
	//     	- 1
	//
	// 	- Valid values if you set `PeriodUnit` to `Month`:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	//     	- 6
	//
	// 	- Valid values if you set `PeriodUnit` to `Year`:
	//
	//     	- 1
	//
	//     	- 2
	//
	//     	- 3
	//
	// >  If you set `ChargeType` to `PostPaid`, set this parameter to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter is available if you set `ChargeType` to `PrePaid`.
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify an invalid value combination for Period and PeriodUnit, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// >  If you set `ChargeType` to `PostPaid`, set this parameter to `Month`.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Week
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the pre-open application.
	//
	// example:
	//
	// cag-b2ron*******
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The promotion ID. You can call the [GetResourcePrice](https://help.aliyun.com/document_detail/428503.html) operation to obtain the ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The runtime policy.
	RuntimePolicy *CreateAppInstanceGroupRequestRuntimePolicy `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty" type:"Struct"`
	// The security policy.
	SecurityPolicy *CreateAppInstanceGroupRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The period of time during which the application can be recycled. The recycling period is the period of time between the time when the end user disconnects from the application and the time when processes exit the application. If you do not want to recycle the application, set this parameter to `-1`. Valid values:-1 and 3 to 300. The value must be an integer. Default value: `15`. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicy *CreateAppInstanceGroupRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	// Payment method subtype.
	//
	// example:
	//
	// postPaid
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The custom policy.
	UserDefinePolicy *CreateAppInstanceGroupRequestUserDefinePolicy `json:"UserDefinePolicy,omitempty" xml:"UserDefinePolicy,omitempty" type:"Struct"`
	// List of authorized user group IDs.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The information about the user that you want to add to the assigned user list of the delivery group. This parameter is required if you configure `Users`.
	UserInfo *CreateAppInstanceGroupRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// The users that you want to add to the assigned user list of the delivery group.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	// Display policy.
	VideoPolicy *CreateAppInstanceGroupRequestVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
}

func (s CreateAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequest) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *CreateAppInstanceGroupRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *CreateAppInstanceGroupRequest) GetAppPackageType() *string {
	return s.AppPackageType
}

func (s *CreateAppInstanceGroupRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *CreateAppInstanceGroupRequest) GetAuthMode() *string {
	return s.AuthMode
}

func (s *CreateAppInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAppInstanceGroupRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAppInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAppInstanceGroupRequest) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *CreateAppInstanceGroupRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAppInstanceGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAppInstanceGroupRequest) GetNetwork() *CreateAppInstanceGroupRequestNetwork {
	return s.Network
}

func (s *CreateAppInstanceGroupRequest) GetNodePool() *CreateAppInstanceGroupRequestNodePool {
	return s.NodePool
}

func (s *CreateAppInstanceGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAppInstanceGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAppInstanceGroupRequest) GetPreOpenAppId() *string {
	return s.PreOpenAppId
}

func (s *CreateAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateAppInstanceGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAppInstanceGroupRequest) GetRuntimePolicy() *CreateAppInstanceGroupRequestRuntimePolicy {
	return s.RuntimePolicy
}

func (s *CreateAppInstanceGroupRequest) GetSecurityPolicy() *CreateAppInstanceGroupRequestSecurityPolicy {
	return s.SecurityPolicy
}

func (s *CreateAppInstanceGroupRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *CreateAppInstanceGroupRequest) GetStoragePolicy() *CreateAppInstanceGroupRequestStoragePolicy {
	return s.StoragePolicy
}

func (s *CreateAppInstanceGroupRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *CreateAppInstanceGroupRequest) GetUserDefinePolicy() *CreateAppInstanceGroupRequestUserDefinePolicy {
	return s.UserDefinePolicy
}

func (s *CreateAppInstanceGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateAppInstanceGroupRequest) GetUserInfo() *CreateAppInstanceGroupRequestUserInfo {
	return s.UserInfo
}

func (s *CreateAppInstanceGroupRequest) GetUsers() []*string {
	return s.Users
}

func (s *CreateAppInstanceGroupRequest) GetVideoPolicy() *CreateAppInstanceGroupRequestVideoPolicy {
	return s.VideoPolicy
}

func (s *CreateAppInstanceGroupRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppPackageType(v string) *CreateAppInstanceGroupRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppPolicyId(v string) *CreateAppInstanceGroupRequest {
	s.AppPolicyId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAuthMode(v string) *CreateAppInstanceGroupRequest {
	s.AuthMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoPay(v bool) *CreateAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetBizRegionId(v string) *CreateAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeType(v string) *CreateAppInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetClusterId(v string) *CreateAppInstanceGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNetwork(v *CreateAppInstanceGroupRequestNetwork) *CreateAppInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNodePool(v *CreateAppInstanceGroupRequestNodePool) *CreateAppInstanceGroupRequest {
	s.NodePool = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriod(v int32) *CreateAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetProductType(v string) *CreateAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPromotionId(v string) *CreateAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetRuntimePolicy(v *CreateAppInstanceGroupRequestRuntimePolicy) *CreateAppInstanceGroupRequest {
	s.RuntimePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSecurityPolicy(v *CreateAppInstanceGroupRequestSecurityPolicy) *CreateAppInstanceGroupRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetStoragePolicy(v *CreateAppInstanceGroupRequestStoragePolicy) *CreateAppInstanceGroupRequest {
	s.StoragePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSubPayType(v string) *CreateAppInstanceGroupRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserDefinePolicy(v *CreateAppInstanceGroupRequestUserDefinePolicy) *CreateAppInstanceGroupRequest {
	s.UserDefinePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserGroupIds(v []*string) *CreateAppInstanceGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserInfo(v *CreateAppInstanceGroupRequestUserInfo) *CreateAppInstanceGroupRequest {
	s.UserInfo = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUsers(v []*string) *CreateAppInstanceGroupRequest {
	s.Users = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetVideoPolicy(v *CreateAppInstanceGroupRequestVideoPolicy) *CreateAppInstanceGroupRequest {
	s.VideoPolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.NodePool != nil {
		if err := s.NodePool.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimePolicy != nil {
		if err := s.RuntimePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityPolicy != nil {
		if err := s.SecurityPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.StoragePolicy != nil {
		if err := s.StoragePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.UserDefinePolicy != nil {
		if err := s.UserDefinePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VideoPolicy != nil {
		if err := s.VideoPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppInstanceGroupRequestNetwork struct {
	// The domain name rules.
	DomainRules []*CreateAppInstanceGroupRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
	// The validity period of the public IP address. If the specified value is exceeded, the IP address is updated at next logon. Minimum value: 60. Unit: minutes.
	//
	// example:
	//
	// 60
	IpExpireMinutes *int32 `json:"IpExpireMinutes,omitempty" xml:"IpExpireMinutes,omitempty"`
	// Office Network ID.
	//
	// example:
	//
	// cn-hongkong+dir-842567****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The route settings. This parameter is available only if you set `StrategyType` to `Mixed`.
	Routes []*CreateAppInstanceGroupRequestNetworkRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	// The type of the network policy.
	//
	// Valid values:
	//
	// 	- Mixed: the hybrid mode. In this mode, a device is deployed in one virtual private cloud (VPC). Two NICs are provided and an independent public IP address is configured for the device.
	//
	// 	- Shared: the shared mode. In this mode, a single NIC is provided for a device and the network is accessed by using NAT Gateway.
	//
	// example:
	//
	// Shared
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// List of virtual switch IDs.
	//
	// - Valid only for custom office networks.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetwork) GetDomainRules() []*CreateAppInstanceGroupRequestNetworkDomainRules {
	return s.DomainRules
}

func (s *CreateAppInstanceGroupRequestNetwork) GetIpExpireMinutes() *int32 {
	return s.IpExpireMinutes
}

func (s *CreateAppInstanceGroupRequestNetwork) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateAppInstanceGroupRequestNetwork) GetRoutes() []*CreateAppInstanceGroupRequestNetworkRoutes {
	return s.Routes
}

func (s *CreateAppInstanceGroupRequestNetwork) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CreateAppInstanceGroupRequestNetwork) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateAppInstanceGroupRequestNetwork) SetDomainRules(v []*CreateAppInstanceGroupRequestNetworkDomainRules) *CreateAppInstanceGroupRequestNetwork {
	s.DomainRules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetIpExpireMinutes(v int32) *CreateAppInstanceGroupRequestNetwork {
	s.IpExpireMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetOfficeSiteId(v string) *CreateAppInstanceGroupRequestNetwork {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetRoutes(v []*CreateAppInstanceGroupRequestNetworkRoutes) *CreateAppInstanceGroupRequestNetwork {
	s.Routes = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetStrategyType(v string) *CreateAppInstanceGroupRequestNetwork {
	s.StrategyType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetVSwitchIds(v []*string) *CreateAppInstanceGroupRequestNetwork {
	s.VSwitchIds = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) Validate() error {
	if s.DomainRules != nil {
		for _, item := range s.DomainRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppInstanceGroupRequestNetworkDomainRules struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The policy used for the domain name.
	//
	// Valid values:
	//
	// 	- allow
	//
	// 	- block
	//
	// example:
	//
	// block
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) GetDomain() *string {
	return s.Domain
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) GetPolicy() *string {
	return s.Policy
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetDomain(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetPolicy(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestNetworkRoutes struct {
	// The destination. The value is a CIDR block.
	//
	// example:
	//
	// 139.196.XX.XX/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The network egress mode.
	//
	// Valid value:
	//
	// 	- Shared: accesses the network by using NAT Gateway.
	//
	// example:
	//
	// Shared
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) GetDestination() *string {
	return s.Destination
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) GetMode() *string {
	return s.Mode
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetDestination(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Destination = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetMode(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Mode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestNodePool struct {
	// Maximum number of idle sessions. When this value is specified, auto-scaling is triggered only if the session utilization exceeds `ScalingUsageThreshold` and the current number of idle sessions in the delivery group is less than `MaxIdleAppInstanceAmount`. Otherwise, it is considered that sufficient idle sessions are available, and no auto-scaling will occur. This parameter allows flexible control over elastic scaling behavior and helps reduce usage costs.
	//
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// The maximum number of resources that can be created for scale-out. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_USAGE`.
	//
	// example:
	//
	// 10
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// The number of resources that you want to purchase. Valid values: 1 to 100.
	//
	// >
	//
	// 	- This parameter is required if the resources are subscription resources.
	//
	// 	- If the resources are pay-as-you-go resources, this parameter is required only if you set `StrategyType` to `NODE_FIXED` or `NODE_SCALING_BY_USAGE`.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The maximum number of sessions to which a resource can connect at the same time. If a resource connects to a large number of sessions at the same time, the user experience can be compromised. The value range varies based on the resource type. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The ID of the resource type that you want to purchase. You can call the [ListNodeInstanceType](https://help.aliyun.com/document_detail/428502.html) operation to obtain the ID.
	//
	// Valid values:
	//
	// 	- appstreaming.vgpu.8c16g.4g: WUYING - Graphics_8 vCPUs, 16 GiB Memory, 4 GiB GPU Memory
	//
	// 	- appstreaming.general.8c16g: WUYING - General_8 vCPUs, 16 GiB Memory
	//
	// 	- appstreaming.general.4c8g: WUYING - General_4 vCPUs, 8 GiB Memory
	//
	// 	- appstreaming.vgpu.14c93g.12g: WUYING - Graphics_14 vCPUs, 93 GiB Memory, 12 GiB GPU Memory.
	//
	// 	- appstreaming.vgpu.8c31g.16g: WUYING - Graphics_8 vCPUs, 31 GiB Memory, 16 GiB GPU Memory
	//
	// example:
	//
	// appstreaming.general.4c8g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The schedules of the scaling policy. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_SCHEDULE`.
	RecurrenceSchedules []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The maximum retention period of a resource to which no session is connected. If no session is connected to a resource, the resource is automatically scaled in after the specified retention period elapses. Valid values: 5 to 120. Default value: 5. Unit: minutes. If one of the following situations occurs, the resource is not scaled in.
	//
	// 	- If automatic scale-out is triggered after the resource is scaled in, the scale-in is not executed. This prevents repeated scale-in and scale-out.
	//
	// 	- If automatic scale-out is triggered due to an increase in the number of sessions during the specified period of time, the resource is not scaled in and the countdown restarts.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of resources that are created each time resources are scaled out. Valid values: 1 to 10. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_USAGE`.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper limit of session usage. If the session usage exceeds the specified upper limit, auto scaling is automatically triggered. The session usage is calculated by using the following formula: `Session usage = Number of current sessions/(Total number of resources × Number of concurrent sessions) × 100%`. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_USAGE`. Valid values: 0 to 100. Default value: 85.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The expiration date of the scaling policy. Format: yyyy-MM-dd. The interval between the expiration date and the effective date must be from 7 days to 1 year. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_SCHEDULE`.
	//
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// The effective date of the scaling policy. Format: yyyy-MM-dd. The date must be the same as or later than the current date. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_SCHEDULE`.
	//
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// The scaling policy of resources.
	//
	// >
	//
	// 	- `NODE_FIXED`: fixed number of resources. This value is applicable to pay-as-you-go resources and subscription resources.
	//
	// 	- `NODE_SCALING_BY_USAGE`: auto scaling. This value is applicable to pay-as-you-go resources and subscription resources.
	//
	// 	- `NODE_SCALING_BY_SCHEDULE`: scheduled scaling. This value is applicable only to pay-as-you-go resources.
	//
	// Valid values:
	//
	// 	- NODE_FIXED: fixed number of resources
	//
	// 	- NODE_SCALING_BY_SCHEDULE: scheduled scaling
	//
	// 	- NODE_SCALING_BY_USAGE: auto scaling
	//
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// Specifies whether to enable the warmup policy for resources. This parameter is required if you set `StrategyType` to `NODE_SCALING_BY_SCHEDULE`.
	//
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePool) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePool) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePool) GetMaxIdleAppInstanceAmount() *int32 {
	return s.MaxIdleAppInstanceAmount
}

func (s *CreateAppInstanceGroupRequestNodePool) GetMaxScalingAmount() *int32 {
	return s.MaxScalingAmount
}

func (s *CreateAppInstanceGroupRequestNodePool) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *CreateAppInstanceGroupRequestNodePool) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *CreateAppInstanceGroupRequestNodePool) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *CreateAppInstanceGroupRequestNodePool) GetRecurrenceSchedules() []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	return s.RecurrenceSchedules
}

func (s *CreateAppInstanceGroupRequestNodePool) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *CreateAppInstanceGroupRequestNodePool) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *CreateAppInstanceGroupRequestNodePool) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *CreateAppInstanceGroupRequestNodePool) GetStrategyDisableDate() *string {
	return s.StrategyDisableDate
}

func (s *CreateAppInstanceGroupRequestNodePool) GetStrategyEnableDate() *string {
	return s.StrategyEnableDate
}

func (s *CreateAppInstanceGroupRequestNodePool) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CreateAppInstanceGroupRequestNodePool) GetWarmUp() *bool {
	return s.WarmUp
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxIdleAppInstanceAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxScalingAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeCapacity(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeInstanceType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetRecurrenceSchedules(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) *CreateAppInstanceGroupRequestNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingDownAfterIdleMinutes(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingStep(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingStep = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingUsageThreshold(v string) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyDisableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyEnableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetWarmUp(v bool) *CreateAppInstanceGroupRequestNodePool {
	s.WarmUp = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) Validate() error {
	if s.RecurrenceSchedules != nil {
		for _, item := range s.RecurrenceSchedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules struct {
	// The schedule type of the scaling policy. This parameter must be configured together with `RecurrenceValues`.``
	//
	// Valid value:
	//
	// 	- Weekly: The scaling policy is executed on specific days each week.
	//
	// example:
	//
	// weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The days of each week on which the scaling policy is executed.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The time periods during which the scaling policy can be executed. The time periods must meet the following requirements:
	//
	// 	- Up to three time periods can be added.
	//
	// 	- Time periods cannot be overlapped.
	//
	// 	- The interval between two consecutive time periods must be greater than or equal to 5 minutes.
	//
	// 	- Each time period must be greater than or equal to 15 minutes.
	//
	// 	- The total length of the time periods that you specify cannot be greater than a day.
	TimerPeriods []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GetTimerPeriods() []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	return s.TimerPeriods
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetTimerPeriods(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) Validate() error {
	if s.TimerPeriods != nil {
		for _, item := range s.TimerPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods struct {
	// The number of resources.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The end time of the time period. Format: HH:mm.
	//
	// example:
	//
	// 15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the time period. Format: HH:mm.
	//
	// example:
	//
	// 12:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestRuntimePolicy struct {
	// Specifies whether to enable the debugging mode. If you want to call the `GetDebugAppInstance` and `CreateImageFromAppInstanceGroup` operations, you must set this parameter to `ON`.
	//
	// Valid values:
	//
	// 	- OFF
	//
	// 	- ON
	//
	// example:
	//
	// OFF
	DebugMode *string `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	// Only one application is allowed to be opened within a single session.
	//
	// - When enabled, launching multiple applications from the delivery group will allocate a separate session for each application, resulting in higher session consumption.
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// Persistent session scheduling mode.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// DYNAMIC
	PersistentAppInstanceScheduleMode *string `json:"PersistentAppInstanceScheduleMode,omitempty" xml:"PersistentAppInstanceScheduleMode,omitempty"`
	// Session pre-launch toggle.
	//
	// - If not specified, the default value is true.
	//
	// example:
	//
	// false
	SessionPreOpen *string `json:"SessionPreOpen,omitempty" xml:"SessionPreOpen,omitempty"`
	// The session type.
	//
	// Valid values:
	//
	// 	- CONSOLE: console session
	//
	// 	- NORMAL: Remote Desktop Protocol (RDP)-based O\\&M session
	//
	// example:
	//
	// NORMAL
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The generation mode of the session users. Valid value:
	//
	// - wyid. In this case, you must set sessionPreOpen to false.
	//
	// example:
	//
	// wyid
	SessionUserGenerationMode *string `json:"SessionUserGenerationMode,omitempty" xml:"SessionUserGenerationMode,omitempty"`
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetDebugMode() *string {
	return s.DebugMode
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetPerSessionPerApp() *bool {
	return s.PerSessionPerApp
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetPersistentAppInstanceScheduleMode() *string {
	return s.PersistentAppInstanceScheduleMode
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetSessionPreOpen() *string {
	return s.SessionPreOpen
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetSessionType() *string {
	return s.SessionType
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) GetSessionUserGenerationMode() *string {
	return s.SessionUserGenerationMode
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetDebugMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.DebugMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetPerSessionPerApp(v bool) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.PerSessionPerApp = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetPersistentAppInstanceScheduleMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.PersistentAppInstanceScheduleMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionPreOpen(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionPreOpen = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionType(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionUserGenerationMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionUserGenerationMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestSecurityPolicy struct {
	// Specifies whether to reset after unbinding from a delivery group.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	// Specifies whether to skip user permission verification.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) GetResetAfterUnbind() *bool {
	return s.ResetAfterUnbind
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) GetSkipUserAuthCheck() *bool {
	return s.SkipUserAuthCheck
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetResetAfterUnbind(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestStoragePolicy struct {
	// The storage types.
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	// User data roaming configuration.
	UserProfile *CreateAppInstanceGroupRequestStoragePolicyUserProfile `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
}

func (s CreateAppInstanceGroupRequestStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) GetStorageTypeList() []*string {
	return s.StorageTypeList
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) GetUserProfile() *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	return s.UserProfile
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) SetStorageTypeList(v []*string) *CreateAppInstanceGroupRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) SetUserProfile(v *CreateAppInstanceGroupRequestStoragePolicyUserProfile) *CreateAppInstanceGroupRequestStoragePolicy {
	s.UserProfile = v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) Validate() error {
	if s.UserProfile != nil {
		if err := s.UserProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppInstanceGroupRequestStoragePolicyUserProfile struct {
	// Remote storage path for user data roaming.
	//
	// - If left empty, the default value is the delivery group ID.
	//
	// - For cross-delivery-group (within the same VPC) user data roaming, the same value must be configured for all participating delivery groups.
	//
	// example:
	//
	// ID20250101
	RemoteStoragePath *string `json:"RemoteStoragePath,omitempty" xml:"RemoteStoragePath,omitempty"`
	// Remote storage type used for user data roaming.
	//
	// example:
	//
	// NAS
	RemoteStorageType *string `json:"RemoteStorageType,omitempty" xml:"RemoteStorageType,omitempty"`
	// User data roaming toggle.
	//
	// example:
	//
	// false
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s CreateAppInstanceGroupRequestStoragePolicyUserProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) GetRemoteStoragePath() *string {
	return s.RemoteStoragePath
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) GetRemoteStorageType() *string {
	return s.RemoteStorageType
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) GetUserProfileSwitch() *bool {
	return s.UserProfileSwitch
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetRemoteStoragePath(v string) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.RemoteStoragePath = &v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetRemoteStorageType(v string) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.RemoteStorageType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestUserDefinePolicy struct {
	// The content of the custom policy. The content must meet the specifications of image versions. To use this parameter, submit a ticket to apply to enable the whitelist feature.
	//
	// example:
	//
	// [{"target":"agent","config":{"abc":"xxx"}}]
	CustomConfig *string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserDefinePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserDefinePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserDefinePolicy) GetCustomConfig() *string {
	return s.CustomConfig
}

func (s *CreateAppInstanceGroupRequestUserDefinePolicy) SetCustomConfig(v string) *CreateAppInstanceGroupRequestUserDefinePolicy {
	s.CustomConfig = &v
	return s
}

func (s *CreateAppInstanceGroupRequestUserDefinePolicy) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestUserInfo struct {
	// The account type of the user.
	//
	// Valid value:
	//
	// 	- Simple: convenience account
	//
	// example:
	//
	// Simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserInfo) GetType() *string {
	return s.Type
}

func (s *CreateAppInstanceGroupRequestUserInfo) SetType(v string) *CreateAppInstanceGroupRequestUserInfo {
	s.Type = &v
	return s
}

func (s *CreateAppInstanceGroupRequestUserInfo) Validate() error {
	return dara.Validate(s)
}

type CreateAppInstanceGroupRequestVideoPolicy struct {
	// Frame rate (FPS).
	//
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// Resolution height, in pixels.
	//
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// Resolution width, in pixels.
	//
	// example:
	//
	// 1920
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	// Streaming mode. Combined with the Webrtc parameter, it indicates the protocol type.
	//
	// - When Webrtc=true and StreamingMode=video, it indicates a WebRTC stream.
	//
	// - When Webrtc=false and StreamingMode=video, it indicates a video stream.
	//
	// - When Webrtc=false and StreamingMode=mix, it indicates a mixed stream.
	//
	// example:
	//
	// video
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// Whether to use adaptive resolution.
	//
	// - true: The session resolution follows changes in the terminal\\"s display area. In this case, SessionResolutionWidth and SessionResolutionHeight represent the maximum values for resolution adjustment.
	//
	// - false: The session resolution does not follow changes in the terminal\\"s display area. In this case, the resolution is fixed to the values of SessionResolutionWidth and SessionResolutionHeight.
	//
	// example:
	//
	// false
	TerminalResolutionAdaptive *bool `json:"TerminalResolutionAdaptive,omitempty" xml:"TerminalResolutionAdaptive,omitempty"`
	// Whether to enable WebRTC. Combined with the StreamingMode parameter, it indicates the protocol type.
	//
	// - When Webrtc=true and StreamingMode=video, it indicates a WebRTC stream.
	//
	// - When Webrtc=false and StreamingMode=video, it indicates a video stream.
	//
	// - When Webrtc=false and StreamingMode=mix, it indicates a mixed stream.
	//
	// example:
	//
	// true
	Webrtc *bool `json:"Webrtc,omitempty" xml:"Webrtc,omitempty"`
}

func (s CreateAppInstanceGroupRequestVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupRequestVideoPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetSessionResolutionHeight() *int32 {
	return s.SessionResolutionHeight
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetSessionResolutionWidth() *int32 {
	return s.SessionResolutionWidth
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetTerminalResolutionAdaptive() *bool {
	return s.TerminalResolutionAdaptive
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) GetWebrtc() *bool {
	return s.Webrtc
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetFrameRate(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetSessionResolutionHeight(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.SessionResolutionHeight = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetSessionResolutionWidth(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.SessionResolutionWidth = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetStreamingMode(v string) *CreateAppInstanceGroupRequestVideoPolicy {
	s.StreamingMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetTerminalResolutionAdaptive(v bool) *CreateAppInstanceGroupRequestVideoPolicy {
	s.TerminalResolutionAdaptive = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetWebrtc(v bool) *CreateAppInstanceGroupRequestVideoPolicy {
	s.Webrtc = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) Validate() error {
	return dara.Validate(s)
}
