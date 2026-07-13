// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPhoneNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCloudPhoneNodeRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateCloudPhoneNodeRequest
	GetAutoRenew() *bool
	SetBandwidthPackageId(v string) *CreateCloudPhoneNodeRequest
	GetBandwidthPackageId() *string
	SetBandwidthPackageType(v string) *CreateCloudPhoneNodeRequest
	GetBandwidthPackageType() *string
	SetBizRegionId(v string) *CreateCloudPhoneNodeRequest
	GetBizRegionId() *string
	SetChannelCookie(v string) *CreateCloudPhoneNodeRequest
	GetChannelCookie() *string
	SetChargeType(v string) *CreateCloudPhoneNodeRequest
	GetChargeType() *string
	SetCount(v string) *CreateCloudPhoneNodeRequest
	GetCount() *string
	SetDisplayConfig(v *CreateCloudPhoneNodeRequestDisplayConfig) *CreateCloudPhoneNodeRequest
	GetDisplayConfig() *CreateCloudPhoneNodeRequestDisplayConfig
	SetDownBandwidthLimit(v int32) *CreateCloudPhoneNodeRequest
	GetDownBandwidthLimit() *int32
	SetImageId(v string) *CreateCloudPhoneNodeRequest
	GetImageId() *string
	SetInstanceType(v string) *CreateCloudPhoneNodeRequest
	GetInstanceType() *string
	SetIsSingleImgDisk(v bool) *CreateCloudPhoneNodeRequest
	GetIsSingleImgDisk() *bool
	SetNetworkId(v string) *CreateCloudPhoneNodeRequest
	GetNetworkId() *string
	SetNetworkInfo(v *CreateCloudPhoneNodeRequestNetworkInfo) *CreateCloudPhoneNodeRequest
	GetNetworkInfo() *CreateCloudPhoneNodeRequestNetworkInfo
	SetNetworkType(v string) *CreateCloudPhoneNodeRequest
	GetNetworkType() *string
	SetNodeName(v string) *CreateCloudPhoneNodeRequest
	GetNodeName() *string
	SetPaidCallBackUrl(v string) *CreateCloudPhoneNodeRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *CreateCloudPhoneNodeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCloudPhoneNodeRequest
	GetPeriodUnit() *string
	SetPhoneCount(v int32) *CreateCloudPhoneNodeRequest
	GetPhoneCount() *int32
	SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeRequest
	GetPhoneDataVolume() *int32
	SetPromotionId(v string) *CreateCloudPhoneNodeRequest
	GetPromotionId() *string
	SetResolutionHeight(v int32) *CreateCloudPhoneNodeRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *CreateCloudPhoneNodeRequest
	GetResolutionWidth() *int32
	SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeRequest
	GetServerShareDataVolume() *int32
	SetServerType(v string) *CreateCloudPhoneNodeRequest
	GetServerType() *string
	SetStreamMode(v int32) *CreateCloudPhoneNodeRequest
	GetStreamMode() *int32
	SetSwapSize(v int32) *CreateCloudPhoneNodeRequest
	GetSwapSize() *int32
	SetTag(v []*CreateCloudPhoneNodeRequestTag) *CreateCloudPhoneNodeRequest
	GetTag() []*CreateCloudPhoneNodeRequestTag
	SetUpBandwidthLimit(v int32) *CreateCloudPhoneNodeRequest
	GetUpBandwidthLimit() *int32
	SetUseTemplate(v string) *CreateCloudPhoneNodeRequest
	GetUseTemplate() *string
	SetVSwitchId(v string) *CreateCloudPhoneNodeRequest
	GetVSwitchId() *string
}

type CreateCloudPhoneNodeRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// cbwp-uf6g3hgg*******8s3lxiob3
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The bandwidth type.
	//
	// example:
	//
	// cbwp_ecd
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId   *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billing type. Only subscription is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of cloud phone matrices to purchase.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The display settings.
	DisplayConfig *CreateCloudPhoneNodeRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	// The downstream bandwidth throttling. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	DownBandwidthLimit *int32 `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type.
	//
	// > To purchase more instance types, [contact pre-sales support](https://smartservice.console.aliyun.com/service/pre-sales-chat?spm=5176.6d6ecb63.0.0.729adda2VqVQx7).
	//
	// example:
	//
	// ac.max
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsSingleImgDisk *bool   `json:"IsSingleImgDisk,omitempty" xml:"IsSingleImgDisk,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The network mapping information of the instance.
	NetworkInfo *CreateCloudPhoneNodeRequestNetworkInfo `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty" type:"Struct"`
	// The network type of the instance.
	//
	// example:
	//
	// network_pro_ecd
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The name of the cloud phone matrix.
	//
	// example:
	//
	// node_name
	NodeName        *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The subscription duration. The unit is specified by PeriodUnit.
	//
	// - If PeriodUnit is set to **year**, the value can only be 1.
	//
	// - If PeriodUnit is set to **month**, valid values are 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of cloud phone instances to create in a single matrix.
	//
	// example:
	//
	// 25
	PhoneCount *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	// The size of the independent device storage. Unit: GiB.
	//
	// example:
	//
	// 10
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	PromotionId     *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The size of the shared device storage. Unit: GiB.
	//
	// > The minimum value of the shared device storage must be greater than the number of instances in the matrix multiplied by 10 GiB.
	//
	// example:
	//
	// 200
	ServerShareDataVolume *int32 `json:"ServerShareDataVolume,omitempty" xml:"ServerShareDataVolume,omitempty"`
	// The specifications of the cloud phone matrix.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpm.gx7.10xlarge
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The streaming mode for instances in the cloud phone matrix. If this parameter is not specified, the default value is preemptive mode.
	//
	// example:
	//
	// 1
	StreamMode *int32 `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	SwapSize   *int32 `json:"SwapSize,omitempty" xml:"SwapSize,omitempty"`
	// The tags of the resource.
	Tag []*CreateCloudPhoneNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The upstream bandwidth throttling. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	UpBandwidthLimit *int32 `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
	// Specifies whether to use a template during creation. Set this parameter to `Random` to use a random template from the template list. Alternatively, specify a template ID to use that template.
	//
	// example:
	//
	// Random
	UseTemplate *string `json:"UseTemplate,omitempty" xml:"UseTemplate,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeekryyc1q3sm72l****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateCloudPhoneNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCloudPhoneNodeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateCloudPhoneNodeRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateCloudPhoneNodeRequest) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *CreateCloudPhoneNodeRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateCloudPhoneNodeRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateCloudPhoneNodeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateCloudPhoneNodeRequest) GetCount() *string {
	return s.Count
}

func (s *CreateCloudPhoneNodeRequest) GetDisplayConfig() *CreateCloudPhoneNodeRequestDisplayConfig {
	return s.DisplayConfig
}

func (s *CreateCloudPhoneNodeRequest) GetDownBandwidthLimit() *int32 {
	return s.DownBandwidthLimit
}

func (s *CreateCloudPhoneNodeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateCloudPhoneNodeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCloudPhoneNodeRequest) GetIsSingleImgDisk() *bool {
	return s.IsSingleImgDisk
}

func (s *CreateCloudPhoneNodeRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateCloudPhoneNodeRequest) GetNetworkInfo() *CreateCloudPhoneNodeRequestNetworkInfo {
	return s.NetworkInfo
}

func (s *CreateCloudPhoneNodeRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateCloudPhoneNodeRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateCloudPhoneNodeRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *CreateCloudPhoneNodeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateCloudPhoneNodeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCloudPhoneNodeRequest) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *CreateCloudPhoneNodeRequest) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *CreateCloudPhoneNodeRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateCloudPhoneNodeRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *CreateCloudPhoneNodeRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *CreateCloudPhoneNodeRequest) GetServerShareDataVolume() *int32 {
	return s.ServerShareDataVolume
}

func (s *CreateCloudPhoneNodeRequest) GetServerType() *string {
	return s.ServerType
}

func (s *CreateCloudPhoneNodeRequest) GetStreamMode() *int32 {
	return s.StreamMode
}

func (s *CreateCloudPhoneNodeRequest) GetSwapSize() *int32 {
	return s.SwapSize
}

func (s *CreateCloudPhoneNodeRequest) GetTag() []*CreateCloudPhoneNodeRequestTag {
	return s.Tag
}

func (s *CreateCloudPhoneNodeRequest) GetUpBandwidthLimit() *int32 {
	return s.UpBandwidthLimit
}

func (s *CreateCloudPhoneNodeRequest) GetUseTemplate() *string {
	return s.UseTemplate
}

func (s *CreateCloudPhoneNodeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateCloudPhoneNodeRequest) SetAutoPay(v bool) *CreateCloudPhoneNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetAutoRenew(v bool) *CreateCloudPhoneNodeRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetBandwidthPackageId(v string) *CreateCloudPhoneNodeRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetBandwidthPackageType(v string) *CreateCloudPhoneNodeRequest {
	s.BandwidthPackageType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetBizRegionId(v string) *CreateCloudPhoneNodeRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetChannelCookie(v string) *CreateCloudPhoneNodeRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetChargeType(v string) *CreateCloudPhoneNodeRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetCount(v string) *CreateCloudPhoneNodeRequest {
	s.Count = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetDisplayConfig(v *CreateCloudPhoneNodeRequestDisplayConfig) *CreateCloudPhoneNodeRequest {
	s.DisplayConfig = v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetDownBandwidthLimit(v int32) *CreateCloudPhoneNodeRequest {
	s.DownBandwidthLimit = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetImageId(v string) *CreateCloudPhoneNodeRequest {
	s.ImageId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetInstanceType(v string) *CreateCloudPhoneNodeRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetIsSingleImgDisk(v bool) *CreateCloudPhoneNodeRequest {
	s.IsSingleImgDisk = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNetworkId(v string) *CreateCloudPhoneNodeRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNetworkInfo(v *CreateCloudPhoneNodeRequestNetworkInfo) *CreateCloudPhoneNodeRequest {
	s.NetworkInfo = v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNetworkType(v string) *CreateCloudPhoneNodeRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetNodeName(v string) *CreateCloudPhoneNodeRequest {
	s.NodeName = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPaidCallBackUrl(v string) *CreateCloudPhoneNodeRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPeriod(v int32) *CreateCloudPhoneNodeRequest {
	s.Period = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPeriodUnit(v string) *CreateCloudPhoneNodeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPhoneCount(v int32) *CreateCloudPhoneNodeRequest {
	s.PhoneCount = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetPromotionId(v string) *CreateCloudPhoneNodeRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetResolutionHeight(v int32) *CreateCloudPhoneNodeRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetResolutionWidth(v int32) *CreateCloudPhoneNodeRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeRequest {
	s.ServerShareDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetServerType(v string) *CreateCloudPhoneNodeRequest {
	s.ServerType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetStreamMode(v int32) *CreateCloudPhoneNodeRequest {
	s.StreamMode = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetSwapSize(v int32) *CreateCloudPhoneNodeRequest {
	s.SwapSize = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetTag(v []*CreateCloudPhoneNodeRequestTag) *CreateCloudPhoneNodeRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetUpBandwidthLimit(v int32) *CreateCloudPhoneNodeRequest {
	s.UpBandwidthLimit = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetUseTemplate(v string) *CreateCloudPhoneNodeRequest {
	s.UseTemplate = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) SetVSwitchId(v string) *CreateCloudPhoneNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateCloudPhoneNodeRequest) Validate() error {
	if s.DisplayConfig != nil {
		if err := s.DisplayConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInfo != nil {
		if err := s.NetworkInfo.Validate(); err != nil {
			return err
		}
	}
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

type CreateCloudPhoneNodeRequestDisplayConfig struct {
	// The DPI. Valid values: 72 to 600.
	//
	// example:
	//
	// 240
	Dpi *int32 `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// > This parameter is not yet available for public use.
	//
	// example:
	//
	// null
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
}

func (s CreateCloudPhoneNodeRequestDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) GetDpi() *int32 {
	return s.Dpi
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) GetFps() *int32 {
	return s.Fps
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) GetLockResolution() *string {
	return s.LockResolution
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetDpi(v int32) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetFps(v int32) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.Fps = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) SetLockResolution(v string) *CreateCloudPhoneNodeRequestDisplayConfig {
	s.LockResolution = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestDisplayConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCloudPhoneNodeRequestNetworkInfo struct {
	// The name of the bandwidth plan.
	//
	// example:
	//
	// inst-bandwidth-pkg-1
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" xml:"BandwidthPackageName,omitempty"`
	// The private CIDR block.
	//
	// example:
	//
	// 10.10.13.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The billable methods of the bandwidth plan. Valid values:
	//
	// <props="china">
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// - **PayBy95**: pay-by-95th-percentile. IPv6 Internet bandwidth does not support pay-by-95th-percentile billing by default. To use this billing method, contact your account manager.
	//
	//
	// <props="intl">
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ratio of IP addresses to instances.
	//
	// example:
	//
	// 128
	IpRatio *int32 `json:"IpRatio,omitempty" xml:"IpRatio,omitempty"`
	// The line type.
	//
	// example:
	//
	// ChinaTelecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The maximum bandwidth of the bandwidth plan. Unit: Mbit/s. The total bandwidth of all instances that use this bandwidth plan is subject to this limit.
	//
	// example:
	//
	// 200
	LimitedBandwidth *int32  `json:"LimitedBandwidth,omitempty" xml:"LimitedBandwidth,omitempty"`
	PaidCallbackUrl  *string `json:"PaidCallbackUrl,omitempty" xml:"PaidCallbackUrl,omitempty"`
	// The billing type.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The visibility scope.
	//
	// example:
	//
	// CPS
	VisibleType *string `json:"VisibleType,omitempty" xml:"VisibleType,omitempty"`
}

func (s CreateCloudPhoneNodeRequestNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetBandwidthPackageName() *string {
	return s.BandwidthPackageName
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetIpRatio() *int32 {
	return s.IpRatio
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetIsp() *string {
	return s.Isp
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetLimitedBandwidth() *int32 {
	return s.LimitedBandwidth
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetPaidCallbackUrl() *string {
	return s.PaidCallbackUrl
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetPayType() *string {
	return s.PayType
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) GetVisibleType() *string {
	return s.VisibleType
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetBandwidthPackageName(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.BandwidthPackageName = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetCidrBlock(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.CidrBlock = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetInternetChargeType(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.InternetChargeType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetIpRatio(v int32) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.IpRatio = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetIsp(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.Isp = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetLimitedBandwidth(v int32) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.LimitedBandwidth = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetPaidCallbackUrl(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.PaidCallbackUrl = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetPayType(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.PayType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) SetVisibleType(v string) *CreateCloudPhoneNodeRequestNetworkInfo {
	s.VisibleType = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestNetworkInfo) Validate() error {
	return dara.Validate(s)
}

type CreateCloudPhoneNodeRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// keyname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// valuename
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudPhoneNodeRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCloudPhoneNodeRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCloudPhoneNodeRequestTag) SetKey(v string) *CreateCloudPhoneNodeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestTag) SetValue(v string) *CreateCloudPhoneNodeRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCloudPhoneNodeRequestTag) Validate() error {
	return dara.Validate(s)
}
