// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageId(v string) *CreateAppInstanceGroupShrinkRequest
	GetAppCenterImageId() *string
	SetAppInstanceGroupName(v string) *CreateAppInstanceGroupShrinkRequest
	GetAppInstanceGroupName() *string
	SetAppPackageType(v string) *CreateAppInstanceGroupShrinkRequest
	GetAppPackageType() *string
	SetAppPolicyId(v string) *CreateAppInstanceGroupShrinkRequest
	GetAppPolicyId() *string
	SetAuthMode(v string) *CreateAppInstanceGroupShrinkRequest
	GetAuthMode() *string
	SetAutoPay(v bool) *CreateAppInstanceGroupShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateAppInstanceGroupShrinkRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateAppInstanceGroupShrinkRequest
	GetBizRegionId() *string
	SetChargeResourceMode(v string) *CreateAppInstanceGroupShrinkRequest
	GetChargeResourceMode() *string
	SetChargeType(v string) *CreateAppInstanceGroupShrinkRequest
	GetChargeType() *string
	SetClusterId(v string) *CreateAppInstanceGroupShrinkRequest
	GetClusterId() *string
	SetNetworkShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetNetworkShrink() *string
	SetNodePoolShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetNodePoolShrink() *string
	SetPeriod(v int32) *CreateAppInstanceGroupShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAppInstanceGroupShrinkRequest
	GetPeriodUnit() *string
	SetPreOpenAppId(v string) *CreateAppInstanceGroupShrinkRequest
	GetPreOpenAppId() *string
	SetProductType(v string) *CreateAppInstanceGroupShrinkRequest
	GetProductType() *string
	SetPromotionId(v string) *CreateAppInstanceGroupShrinkRequest
	GetPromotionId() *string
	SetRuntimePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetRuntimePolicyShrink() *string
	SetSecurityPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetSecurityPolicyShrink() *string
	SetSessionTimeout(v int32) *CreateAppInstanceGroupShrinkRequest
	GetSessionTimeout() *int32
	SetStoragePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetStoragePolicyShrink() *string
	SetSubPayType(v string) *CreateAppInstanceGroupShrinkRequest
	GetSubPayType() *string
	SetUserDefinePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetUserDefinePolicyShrink() *string
	SetUserGroupIds(v []*string) *CreateAppInstanceGroupShrinkRequest
	GetUserGroupIds() []*string
	SetUserInfoShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetUserInfoShrink() *string
	SetUsers(v []*string) *CreateAppInstanceGroupShrinkRequest
	GetUsers() []*string
	SetVideoPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest
	GetVideoPolicyShrink() *string
}

type CreateAppInstanceGroupShrinkRequest struct {
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
	NetworkShrink *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The node pool object.
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
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
	RuntimePolicyShrink *string `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty"`
	// The security policy.
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// The period of time during which the application can be recycled. The recycling period is the period of time between the time when the end user disconnects from the application and the time when processes exit the application. If you do not want to recycle the application, set this parameter to `-1`. Valid values:-1 and 3 to 300. The value must be an integer. Default value: `15`. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicyShrink *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	// Payment method subtype.
	//
	// example:
	//
	// postPaid
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The custom policy.
	UserDefinePolicyShrink *string `json:"UserDefinePolicy,omitempty" xml:"UserDefinePolicy,omitempty"`
	// List of authorized user group IDs.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The information about the user that you want to add to the assigned user list of the delivery group. This parameter is required if you configure `Users`.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	// The users that you want to add to the assigned user list of the delivery group.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	// Display policy.
	VideoPolicyShrink *string `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty"`
}

func (s CreateAppInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAppPackageType() *string {
	return s.AppPackageType
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAuthMode() *string {
	return s.AuthMode
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAppInstanceGroupShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAppInstanceGroupShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *CreateAppInstanceGroupShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAppInstanceGroupShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetNetworkShrink() *string {
	return s.NetworkShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetNodePoolShrink() *string {
	return s.NodePoolShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAppInstanceGroupShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAppInstanceGroupShrinkRequest) GetPreOpenAppId() *string {
	return s.PreOpenAppId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateAppInstanceGroupShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAppInstanceGroupShrinkRequest) GetRuntimePolicyShrink() *string {
	return s.RuntimePolicyShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetSecurityPolicyShrink() *string {
	return s.SecurityPolicyShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *CreateAppInstanceGroupShrinkRequest) GetStoragePolicyShrink() *string {
	return s.StoragePolicyShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *CreateAppInstanceGroupShrinkRequest) GetUserDefinePolicyShrink() *string {
	return s.UserDefinePolicyShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateAppInstanceGroupShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) GetUsers() []*string {
	return s.Users
}

func (s *CreateAppInstanceGroupShrinkRequest) GetVideoPolicyShrink() *string {
	return s.VideoPolicyShrink
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppPackageType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppPolicyId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppPolicyId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAuthMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AuthMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetClusterId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNetworkShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNodePoolShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetProductType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetRuntimePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.RuntimePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSecurityPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetStoragePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSubPayType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserDefinePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserDefinePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserGroupIds(v []*string) *CreateAppInstanceGroupShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserInfoShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUsers(v []*string) *CreateAppInstanceGroupShrinkRequest {
	s.Users = v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetVideoPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.VideoPolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
