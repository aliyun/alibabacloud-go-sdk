// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppPackageType(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetAppPackageType() *string
	SetAuthNotificationEnabled(v bool) *CreateBrowserInstanceGroupShrinkRequest
	GetAuthNotificationEnabled() *bool
	SetAutoPay(v bool) *CreateBrowserInstanceGroupShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateBrowserInstanceGroupShrinkRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetBizRegionId() *string
	SetBrowserConfigShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetBrowserConfigShrink() *string
	SetChargeResourceMode(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetChargeResourceMode() *string
	SetChargeType(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetChargeType() *string
	SetCloudBrowserName(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetCloudBrowserName() *string
	SetImageId(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetImageId() *string
	SetInstanceType(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetInstanceType() *string
	SetMaxAmount(v int32) *CreateBrowserInstanceGroupShrinkRequest
	GetMaxAmount() *int32
	SetNetworkShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetNetworkShrink() *string
	SetNodePoolShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetNodePoolShrink() *string
	SetOsType(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetOsType() *string
	SetPeriod(v int32) *CreateBrowserInstanceGroupShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetPeriodUnit() *string
	SetPolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetPolicyShrink() *string
	SetPromotionId(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetPromotionId() *string
	SetSecurityPolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetSecurityPolicyShrink() *string
	SetStoragePolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetStoragePolicyShrink() *string
	SetSubPayType(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetSubPayType() *string
	SetTagShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetTagShrink() *string
	SetTimersShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetTimersShrink() *string
	SetUserGroupIds(v []*string) *CreateBrowserInstanceGroupShrinkRequest
	GetUserGroupIds() []*string
	SetUserInfoShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetUserInfoShrink() *string
	SetUsersShrink(v string) *CreateBrowserInstanceGroupShrinkRequest
	GetUsersShrink() *string
}

type CreateBrowserInstanceGroupShrinkRequest struct {
	// The plan identifier.
	//
	// Do not specify this parameter.
	//
	// example:
	//
	// -
	AppPackageType *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
	// Specifies whether to send authorization and deauthorization notification emails.
	//
	// - `true`: Sends the notification.
	//
	// - `false`: Does not send the notification.
	//
	// example:
	//
	// true
	AuthNotificationEnabled *bool `json:"AuthNotificationEnabled,omitempty" xml:"AuthNotificationEnabled,omitempty"`
	// The automatic payment parameter.
	//
	// Do not specify this parameter.
	//
	// example:
	//
	// -
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The auto-renewal parameter.
	//
	// Do not specify this parameter.
	//
	// example:
	//
	// -
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business region ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The browser configuration.
	BrowserConfigShrink *string `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty"`
	// The resource billing mode.
	//
	// **For MAU scenarios:*	- Set this parameter to `AppInstance` to bill by instance resource.
	//
	// example:
	//
	// AppInstance
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing type.
	//
	// **For MAU scenarios:*	- Set this parameter to `PostPaid`, which indicates pay-as-you-go billing.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The name of the cloud browser group. This parameter cannot be empty. The name is used to distinguish different browser groups in business management scenarios.
	//
	// This parameter is required.
	//
	// example:
	//
	// BusinessOfficeBrowser
	CloudBrowserName *string `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	// The image identifier used by the cloud browser. The image must be compatible with the operating system.
	//
	// If this parameter is omitted, the default image available for the account is used. If no default image is available, the creation may fail.
	//
	// **Usage condition:*	- When `CookiesSync` is enabled, explicitly specify an image that supports cookie synchronization.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type identifier. Select an instance type that matches the target region, operating system, and inventory conditions.
	//
	// If this parameter is omitted, the default instance type is used.
	//
	// example:
	//
	// appstreaming.general.basic
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The capacity configuration for the MAU billing scenario.
	//
	// example:
	//
	// 5
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The office network and website access restriction configurations. The selected office network must belong to the current account and be located in the region specified by `BizRegionId`.
	NetworkShrink *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The node pool configuration.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	// The operating system type. This parameter is required.
	//
	// Only `Windows` is supported. Other operating systems are not supported.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The number of subscription periods.
	//
	// Do not specify this parameter.
	//
	// example:
	//
	// -
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period.
	//
	// Do not specify this parameter.
	//
	// example:
	//
	// -
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The clipboard, video, watermark, session, and client access policy configurations.
	PolicyShrink *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The promotion ID. Specifies the promotional campaign to apply to the order.
	//
	// Whether the promotion is applicable depends on the campaign rules. Do not specify this parameter if no promotional campaign is used.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The connection security policy for the browser group.
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// The user data storage configuration for the browser group.
	StoragePolicyShrink *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	// The billing subtype.
	//
	// **Set this parameter to `mau` explicitly, which indicates billing by monthly active users.*	- Omitting this field does not enable MAU billing.
	//
	// example:
	//
	// mau
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// Not supported. You do not need to specify this parameter.
	//
	// example:
	//
	// -
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Not supported. You do not need to specify this parameter.
	//
	// example:
	//
	// -
	TimersShrink *string `json:"Timers,omitempty" xml:"Timers,omitempty"`
	// The list of authorized user group identifiers. A maximum of 10 items are supported. The user groups must belong to the current account and match the workspace network account type.
	//
	// **Limit:*	- Cannot be specified together with a non-empty `Users`.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The authorized user account information. The value must match the user and workspace network type.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	// The list of authorized users. A maximum of 200 users can be specified. Users must be created in advance and must match the account type.
	//
	// **Restriction:*	- This parameter cannot be specified together with a non-empty `UserGroupIds`.
	UsersShrink *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s CreateBrowserInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetAppPackageType() *string {
	return s.AppPackageType
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetAuthNotificationEnabled() *bool {
	return s.AuthNotificationEnabled
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetBrowserConfigShrink() *string {
	return s.BrowserConfigShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetCloudBrowserName() *string {
	return s.CloudBrowserName
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetNetworkShrink() *string {
	return s.NetworkShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetNodePoolShrink() *string {
	return s.NodePoolShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetSecurityPolicyShrink() *string {
	return s.SecurityPolicyShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetStoragePolicyShrink() *string {
	return s.StoragePolicyShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetTimersShrink() *string {
	return s.TimersShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) GetUsersShrink() *string {
	return s.UsersShrink
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetAppPackageType(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetAuthNotificationEnabled(v bool) *CreateBrowserInstanceGroupShrinkRequest {
	s.AuthNotificationEnabled = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateBrowserInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateBrowserInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetBrowserConfigShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.BrowserConfigShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetChargeResourceMode(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetChargeType(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetCloudBrowserName(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.CloudBrowserName = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetImageId(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetInstanceType(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetMaxAmount(v int32) *CreateBrowserInstanceGroupShrinkRequest {
	s.MaxAmount = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetNetworkShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetNodePoolShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetOsType(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.OsType = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateBrowserInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetPolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetSecurityPolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetStoragePolicyShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetSubPayType(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetTagShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetTimersShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.TimersShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetUserGroupIds(v []*string) *CreateBrowserInstanceGroupShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetUserInfoShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) SetUsersShrink(v string) *CreateBrowserInstanceGroupShrinkRequest {
	s.UsersShrink = &v
	return s
}

func (s *CreateBrowserInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
