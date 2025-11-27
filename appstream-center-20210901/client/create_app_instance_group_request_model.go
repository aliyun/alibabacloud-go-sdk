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
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId     *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// example:
	//
	// browser.package.5.250.appstreaming.general.basic
	AppPackageType *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
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
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// cls-d39iq73l5c0a8****
	ClusterId *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Network   *CreateAppInstanceGroupRequestNetwork  `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	NodePool  *CreateAppInstanceGroupRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// cag-b2ron*******
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The runtime policy.
	RuntimePolicy  *CreateAppInstanceGroupRequestRuntimePolicy  `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty" type:"Struct"`
	SecurityPolicy *CreateAppInstanceGroupRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	SessionTimeout *int32                                      `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicy  *CreateAppInstanceGroupRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	// example:
	//
	// postPaid
	SubPayType       *string                                        `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	UserDefinePolicy *CreateAppInstanceGroupRequestUserDefinePolicy `json:"UserDefinePolicy,omitempty" xml:"UserDefinePolicy,omitempty" type:"Struct"`
	// if can be null:
	// true
	UserGroupIds []*string                                 `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	UserInfo     *CreateAppInstanceGroupRequestUserInfo    `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	Users        []*string                                 `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	VideoPolicy  *CreateAppInstanceGroupRequestVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
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
	DomainRules []*CreateAppInstanceGroupRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	IpExpireMinutes *int32 `json:"IpExpireMinutes,omitempty" xml:"IpExpireMinutes,omitempty"`
	// example:
	//
	// cn-hongkong+dir-842567****
	OfficeSiteId *string                                       `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Routes       []*CreateAppInstanceGroupRequestNetworkRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	// example:
	//
	// Shared
	StrategyType *string   `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	VSwitchIds   []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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
	// example:
	//
	// 139.196.XX.XX/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
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
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// example:
	//
	// 10
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// example:
	//
	// appstreaming.general.4c8g
	NodeInstanceType    *string                                                     `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	RecurrenceSchedules []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
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
	// example:
	//
	// weekly
	RecurrenceType   *string                                                                 `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
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
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// Specifies whether only one app can be opened in a session.
	//
	// 	- After you enable this feature, the system assigns a session to each app if you open multiple apps in a delivery group. This consumes a larger number of sessions.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// DYNAMIC
	PersistentAppInstanceScheduleMode *string `json:"PersistentAppInstanceScheduleMode,omitempty" xml:"PersistentAppInstanceScheduleMode,omitempty"`
	// Specifies whether to enable pre-open for sessions.
	//
	// 	- Default value: true
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// 	- wyid. In this case, you must set sessionPreOpen to false.
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
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
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
	StorageTypeList []*string                                              `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	UserProfile     *CreateAppInstanceGroupRequestStoragePolicyUserProfile `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
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
	// example:
	//
	// ID20250101
	RemoteStoragePath *string `json:"RemoteStoragePath,omitempty" xml:"RemoteStoragePath,omitempty"`
	// example:
	//
	// NAS
	RemoteStorageType *string `json:"RemoteStorageType,omitempty" xml:"RemoteStorageType,omitempty"`
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
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// example:
	//
	// 1920
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	// example:
	//
	// video
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// example:
	//
	// false
	TerminalResolutionAdaptive *bool `json:"TerminalResolutionAdaptive,omitempty" xml:"TerminalResolutionAdaptive,omitempty"`
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
