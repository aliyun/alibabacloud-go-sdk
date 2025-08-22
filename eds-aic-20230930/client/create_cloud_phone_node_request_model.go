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
	SetNetworkId(v string) *CreateCloudPhoneNodeRequest
	GetNetworkId() *string
	SetNetworkInfo(v *CreateCloudPhoneNodeRequestNetworkInfo) *CreateCloudPhoneNodeRequest
	GetNetworkInfo() *CreateCloudPhoneNodeRequestNetworkInfo
	SetNetworkType(v string) *CreateCloudPhoneNodeRequest
	GetNetworkType() *string
	SetNodeName(v string) *CreateCloudPhoneNodeRequest
	GetNodeName() *string
	SetPeriod(v int32) *CreateCloudPhoneNodeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateCloudPhoneNodeRequest
	GetPeriodUnit() *string
	SetPhoneCount(v int32) *CreateCloudPhoneNodeRequest
	GetPhoneCount() *int32
	SetPhoneDataVolume(v int32) *CreateCloudPhoneNodeRequest
	GetPhoneDataVolume() *int32
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
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- False (default): You must manually complete the payment in the Alibaba Cloud Expenses and Costs console.
	//
	// 	- true: enables the auto-payment feature.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-renewal feature. In this case, the system automatically renews instances upon expiration.
	//
	// 	- false (default): disables the auto-renewal feature. In this case, you need to manually renew instances upon expiration.
	//
	// example:
	//
	// true
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method. Only the subscription billing method is supported.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of cloud phone matrixes you want to purchase.
	//
	// example:
	//
	// 1
	Count              *string                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	DisplayConfig      *CreateCloudPhoneNodeRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	DownBandwidthLimit *int32                                    `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance specification.
	//
	// Valid values:
	//
	// 	- ac.max: By default, this specification allows up to 25 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// 	- ac.plus: By default, this specification allows up to 40 instances. You can adjust this number by using PhoneCount (Value range: 4 to 40).
	//
	// example:
	//
	// ac.max
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId   *string                                 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkInfo *CreateCloudPhoneNodeRequestNetworkInfo `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty" type:"Struct"`
	NetworkType *string                                 `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The name of the cloud phone matrix.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The subscription duration. The unit is specified by `PeriodUnit`. Valid values:
	//
	// 	- When `PeriodUnit` is set to **year**: 1.
	//
	// 	- When `PeriodUnit` is set to **month**: 1, 2, 3, and 6.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of instances per cloud phone matrix.
	//
	// example:
	//
	// 25
	PhoneCount      *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	// The resolution height. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width. Unit: pixel.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The shared storage size Unit: GiB.
	//
	// example:
	//
	// 200
	ServerShareDataVolume *int32 `json:"ServerShareDataVolume,omitempty" xml:"ServerShareDataVolume,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// This parameter is required.
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	StreamMode *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	// The resource tags.
	Tag              []*CreateCloudPhoneNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UpBandwidthLimit *int32                            `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
	UseTemplate      *string                           `json:"UseTemplate,omitempty" xml:"UseTemplate,omitempty"`
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
	return dara.Validate(s)
}

type CreateCloudPhoneNodeRequestDisplayConfig struct {
	Dpi            *int32  `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	Fps            *int32  `json:"Fps,omitempty" xml:"Fps,omitempty"`
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
	BandwidthPackageName *string `json:"BandwidthPackageName,omitempty" xml:"BandwidthPackageName,omitempty"`
	CidrBlock            *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	InternetChargeType   *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpRatio              *int32  `json:"IpRatio,omitempty" xml:"IpRatio,omitempty"`
	Isp                  *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	LimitedBandwidth     *int32  `json:"LimitedBandwidth,omitempty" xml:"LimitedBandwidth,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VisibleType          *string `json:"VisibleType,omitempty" xml:"VisibleType,omitempty"`
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
