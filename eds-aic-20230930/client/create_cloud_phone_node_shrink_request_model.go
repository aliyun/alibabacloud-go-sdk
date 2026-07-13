// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPhoneNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCloudPhoneNodeShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateCloudPhoneNodeShrinkRequest
	GetAutoRenew() *bool
	SetBandwidthPackageId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetBandwidthPackageId() *string
	SetBandwidthPackageType(v string) *CreateCloudPhoneNodeShrinkRequest
	GetBandwidthPackageType() *string
	SetBizRegionId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetBizRegionId() *string
	SetChannelCookie(v string) *CreateCloudPhoneNodeShrinkRequest
	GetChannelCookie() *string
	SetChargeType(v string) *CreateCloudPhoneNodeShrinkRequest
	GetChargeType() *string
	SetCount(v string) *CreateCloudPhoneNodeShrinkRequest
	GetCount() *string
	SetDisplayConfigShrink(v string) *CreateCloudPhoneNodeShrinkRequest
	GetDisplayConfigShrink() *string
	SetDownBandwidthLimit(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetDownBandwidthLimit() *int32
	SetImageId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetImageId() *string
	SetInstanceType(v string) *CreateCloudPhoneNodeShrinkRequest
	GetInstanceType() *string
	SetIsSingleImgDisk(v bool) *CreateCloudPhoneNodeShrinkRequest
	GetIsSingleImgDisk() *bool
	SetNetworkId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetNetworkId() *string
	SetNetworkInfoShrink(v string) *CreateCloudPhoneNodeShrinkRequest
	GetNetworkInfoShrink() *string
	SetNetworkType(v string) *CreateCloudPhoneNodeShrinkRequest
	GetNetworkType() *string
	SetNodeName(v string) *CreateCloudPhoneNodeShrinkRequest
	GetNodeName() *string
	SetPaidCallBackUrl(v string) *CreateCloudPhoneNodeShrinkRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCloudPhoneNodeShrinkRequest
	GetPeriodUnit() *string
	SetPhoneCount(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetPhoneCount() *int32
	SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetPhoneDataVolume() *int32
	SetPromotionId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetPromotionId() *string
	SetResolutionHeight(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetResolutionWidth() *int32
	SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetServerShareDataVolume() *int32
	SetServerType(v string) *CreateCloudPhoneNodeShrinkRequest
	GetServerType() *string
	SetStreamMode(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetStreamMode() *int32
	SetSwapSize(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetSwapSize() *int32
	SetTag(v []*CreateCloudPhoneNodeShrinkRequestTag) *CreateCloudPhoneNodeShrinkRequest
	GetTag() []*CreateCloudPhoneNodeShrinkRequestTag
	SetUpBandwidthLimit(v int32) *CreateCloudPhoneNodeShrinkRequest
	GetUpBandwidthLimit() *int32
	SetUseTemplate(v string) *CreateCloudPhoneNodeShrinkRequest
	GetUseTemplate() *string
	SetVSwitchId(v string) *CreateCloudPhoneNodeShrinkRequest
	GetVSwitchId() *string
}

type CreateCloudPhoneNodeShrinkRequest struct {
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
	DisplayConfigShrink *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
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
	NetworkInfoShrink *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
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
	Tag []*CreateCloudPhoneNodeShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateCloudPhoneNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetCount() *string {
	return s.Count
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetDisplayConfigShrink() *string {
	return s.DisplayConfigShrink
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetDownBandwidthLimit() *int32 {
	return s.DownBandwidthLimit
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetIsSingleImgDisk() *bool {
	return s.IsSingleImgDisk
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetNetworkInfoShrink() *string {
	return s.NetworkInfoShrink
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetServerShareDataVolume() *int32 {
	return s.ServerShareDataVolume
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetServerType() *string {
	return s.ServerType
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetStreamMode() *int32 {
	return s.StreamMode
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetSwapSize() *int32 {
	return s.SwapSize
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetTag() []*CreateCloudPhoneNodeShrinkRequestTag {
	return s.Tag
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetUpBandwidthLimit() *int32 {
	return s.UpBandwidthLimit
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetUseTemplate() *string {
	return s.UseTemplate
}

func (s *CreateCloudPhoneNodeShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetAutoPay(v bool) *CreateCloudPhoneNodeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetAutoRenew(v bool) *CreateCloudPhoneNodeShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetBandwidthPackageId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetBandwidthPackageType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.BandwidthPackageType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetBizRegionId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetChannelCookie(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetChargeType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetCount(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.Count = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetDisplayConfigShrink(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetDownBandwidthLimit(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.DownBandwidthLimit = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetImageId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetInstanceType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetIsSingleImgDisk(v bool) *CreateCloudPhoneNodeShrinkRequest {
	s.IsSingleImgDisk = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNetworkId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNetworkInfoShrink(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NetworkInfoShrink = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNetworkType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetNodeName(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.NodeName = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPaidCallBackUrl(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPeriod(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPeriodUnit(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPhoneCount(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.PhoneCount = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetPromotionId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetResolutionHeight(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetResolutionWidth(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetServerShareDataVolume(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.ServerShareDataVolume = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetServerType(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.ServerType = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetStreamMode(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.StreamMode = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetSwapSize(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.SwapSize = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetTag(v []*CreateCloudPhoneNodeShrinkRequestTag) *CreateCloudPhoneNodeShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetUpBandwidthLimit(v int32) *CreateCloudPhoneNodeShrinkRequest {
	s.UpBandwidthLimit = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetUseTemplate(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.UseTemplate = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) SetVSwitchId(v string) *CreateCloudPhoneNodeShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequest) Validate() error {
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

type CreateCloudPhoneNodeShrinkRequestTag struct {
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

func (s CreateCloudPhoneNodeShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) SetKey(v string) *CreateCloudPhoneNodeShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) SetValue(v string) *CreateCloudPhoneNodeShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCloudPhoneNodeShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
