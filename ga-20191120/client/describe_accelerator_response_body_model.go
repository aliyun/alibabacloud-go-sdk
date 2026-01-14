// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeAcceleratorResponseBody
	GetAcceleratorId() *string
	SetBandwidth(v int32) *DescribeAcceleratorResponseBody
	GetBandwidth() *int32
	SetBandwidthBillingType(v string) *DescribeAcceleratorResponseBody
	GetBandwidthBillingType() *string
	SetBasicBandwidthPackage(v *DescribeAcceleratorResponseBodyBasicBandwidthPackage) *DescribeAcceleratorResponseBody
	GetBasicBandwidthPackage() *DescribeAcceleratorResponseBodyBasicBandwidthPackage
	SetCenId(v string) *DescribeAcceleratorResponseBody
	GetCenId() *string
	SetCreateTime(v int64) *DescribeAcceleratorResponseBody
	GetCreateTime() *int64
	SetCrossBorderMode(v string) *DescribeAcceleratorResponseBody
	GetCrossBorderMode() *string
	SetCrossBorderStatus(v bool) *DescribeAcceleratorResponseBody
	GetCrossBorderStatus() *bool
	SetCrossDomainBandwidthPackage(v *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) *DescribeAcceleratorResponseBody
	GetCrossDomainBandwidthPackage() *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage
	SetCrossPrivateState(v string) *DescribeAcceleratorResponseBody
	GetCrossPrivateState() *string
	SetDdosConfigList(v []*DescribeAcceleratorResponseBodyDdosConfigList) *DescribeAcceleratorResponseBody
	GetDdosConfigList() []*DescribeAcceleratorResponseBodyDdosConfigList
	SetDdosId(v string) *DescribeAcceleratorResponseBody
	GetDdosId() *string
	SetDescription(v string) *DescribeAcceleratorResponseBody
	GetDescription() *string
	SetDnsName(v string) *DescribeAcceleratorResponseBody
	GetDnsName() *string
	SetExpiredTime(v int64) *DescribeAcceleratorResponseBody
	GetExpiredTime() *int64
	SetInstanceChargeType(v string) *DescribeAcceleratorResponseBody
	GetInstanceChargeType() *string
	SetIpSetConfig(v *DescribeAcceleratorResponseBodyIpSetConfig) *DescribeAcceleratorResponseBody
	GetIpSetConfig() *DescribeAcceleratorResponseBodyIpSetConfig
	SetName(v string) *DescribeAcceleratorResponseBody
	GetName() *string
	SetRegionId(v string) *DescribeAcceleratorResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAcceleratorResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeAcceleratorResponseBody
	GetResourceGroupId() *string
	SetSecondDnsName(v string) *DescribeAcceleratorResponseBody
	GetSecondDnsName() *string
	SetServiceId(v string) *DescribeAcceleratorResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeAcceleratorResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeAcceleratorResponseBodyServiceManagedInfos) *DescribeAcceleratorResponseBody
	GetServiceManagedInfos() []*DescribeAcceleratorResponseBodyServiceManagedInfos
	SetSpec(v string) *DescribeAcceleratorResponseBody
	GetSpec() *string
	SetState(v string) *DescribeAcceleratorResponseBody
	GetState() *string
	SetTags(v []*DescribeAcceleratorResponseBodyTags) *DescribeAcceleratorResponseBody
	GetTags() []*DescribeAcceleratorResponseBodyTags
	SetUpgradableStatus(v string) *DescribeAcceleratorResponseBody
	GetUpgradableStatus() *string
}

type DescribeAcceleratorResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth metering method. Valid values:
	//
	// 	- **BandwidthPackage:*	- billed based on bandwidth plans.
	//
	// 	- **CDT**: billed based on data transfer.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The details about the basic bandwidth plan that is associated with the GA instance.
	BasicBandwidthPackage *DescribeAcceleratorResponseBodyBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The ID of the Cloud Enterprise Network (CEN) instance with which the GA instance is associated.
	//
	// example:
	//
	// cen-hjkduu767hc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The timestamp that indicates when the GA instance is created.
	//
	// example:
	//
	// 1650643200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of cross-border acceleration. This parameter is returned for GA instances whose bandwidth metering method is pay-by-data-transfer (CDT).
	//
	// Only **bpgPro*	- may be returned, which indicates BGP (Multi-ISP) Pro lines.
	//
	// example:
	//
	// bpgPro
	CrossBorderMode *string `json:"CrossBorderMode,omitempty" xml:"CrossBorderMode,omitempty"`
	// Indicates whether cross-border acceleration is enabled.
	//
	// - **true**: yes
	//
	// - **false**: no
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The details about the cross-border acceleration bandwidth plan that is associated with the GA instance.
	//
	// This array is returned only for GA instances that are created on the international site (alibabacloud.com).
	CrossDomainBandwidthPackage *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	// Indicates whether cross-border acceleration is enabled.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	CrossPrivateState *string                                          `json:"CrossPrivateState,omitempty" xml:"CrossPrivateState,omitempty"`
	DdosConfigList    []*DescribeAcceleratorResponseBodyDdosConfigList `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro/Premium instance that is associated with the GA instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The description of the GA instance.
	//
	// example:
	//
	// Accelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The canonical name (CNAME) that is assigned to the GA instance.
	//
	// example:
	//
	// ga-bp15u1i2hmtbk8c3i****.aliyunga0019.com
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// The timestamp that indicates when the GA instance expires.
	//
	// example:
	//
	// 1653235200
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the GA instance.
	//
	// example:
	//
	// PREPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The configurations of the acceleration area.
	IpSetConfig *DescribeAcceleratorResponseBodyIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	// The name of the GA instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The region ID of the GA instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw2vwdbujqbq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The CNAME that is used to integrate the GA instance with the Anti-DDoS service.
	//
	// example:
	//
	// ga-bp1f609c76zg6zuna****-1.aliyunga0047.com
	SecondDnsName *string `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	// The ID of the service that manages the GA instance.
	//
	// >  This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the GA instance is managed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that users can perform on the managed instance.
	//
	// >
	//
	// 	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*DescribeAcceleratorResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The specification of the GA instance. Valid values:
	//
	// 	- **1**: Small Ⅰ
	//
	// 	- **2**: Small Ⅱ
	//
	// 	- **3**: Small Ⅲ
	//
	// 	- **5**: Medium Ⅰ
	//
	// 	- **8**: Medium Ⅱ
	//
	// 	- **10**: Medium Ⅲ
	//
	// 	- **20**: Large Ⅰ
	//
	// 	- **30**: Large Ⅱ
	//
	// 	- **40**: Large Ⅲ
	//
	// 	- **50**: Large Ⅳ
	//
	// 	- **60**: Large Ⅴ
	//
	// 	- **70**: Large Ⅵ
	//
	// 	- **80**: Large VⅡ
	//
	// 	- **90**: Large VⅢ
	//
	// 	- **100**: Super Large Ⅰ
	//
	// 	- **200**: Super Large Ⅱ
	//
	// >  The Large Ⅲ specification and higher specifications are available only to users that are added to the whitelist. To use these specifications, contact your Alibaba Cloud account manager.
	//
	// Different specifications provide different capabilities. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/153127.html).
	//
	// example:
	//
	// 1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the GA instance. Valid values:
	//
	// 	- **init**: The GA instance is being initialized.
	//
	// 	- **active**: The GA instance is available.
	//
	// 	- **configuring**: The GA instance is being configured.
	//
	// 	- **binding**: The GA instance is being associated.
	//
	// 	- **unbinding**: The GA instance is being disassociated.
	//
	// 	- **deleting**: The GA instance is being deleted.
	//
	// 	- **finacialLocked**: The GA instance is locked due to overdue payments.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the GA instance.
	Tags []*DescribeAcceleratorResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether the GA instance can be upgraded. Valid values:
	//
	// 	- **notUpgradable:*	- The GA instance does not need to be upgraded.
	//
	// 	- **upgradable:*	- The GA instance can be upgraded to the latest version.
	//
	// 	- **upgradeFailed:*	- The GA instance failed to be upgraded.
	//
	// example:
	//
	// notUpgradable
	UpgradableStatus *string `json:"UpgradableStatus,omitempty" xml:"UpgradableStatus,omitempty"`
}

func (s DescribeAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeAcceleratorResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeAcceleratorResponseBody) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *DescribeAcceleratorResponseBody) GetBasicBandwidthPackage() *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	return s.BasicBandwidthPackage
}

func (s *DescribeAcceleratorResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *DescribeAcceleratorResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeAcceleratorResponseBody) GetCrossBorderMode() *string {
	return s.CrossBorderMode
}

func (s *DescribeAcceleratorResponseBody) GetCrossBorderStatus() *bool {
	return s.CrossBorderStatus
}

func (s *DescribeAcceleratorResponseBody) GetCrossDomainBandwidthPackage() *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage {
	return s.CrossDomainBandwidthPackage
}

func (s *DescribeAcceleratorResponseBody) GetCrossPrivateState() *string {
	return s.CrossPrivateState
}

func (s *DescribeAcceleratorResponseBody) GetDdosConfigList() []*DescribeAcceleratorResponseBodyDdosConfigList {
	return s.DdosConfigList
}

func (s *DescribeAcceleratorResponseBody) GetDdosId() *string {
	return s.DdosId
}

func (s *DescribeAcceleratorResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAcceleratorResponseBody) GetDnsName() *string {
	return s.DnsName
}

func (s *DescribeAcceleratorResponseBody) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *DescribeAcceleratorResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAcceleratorResponseBody) GetIpSetConfig() *DescribeAcceleratorResponseBodyIpSetConfig {
	return s.IpSetConfig
}

func (s *DescribeAcceleratorResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeAcceleratorResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAcceleratorResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAcceleratorResponseBody) GetSecondDnsName() *string {
	return s.SecondDnsName
}

func (s *DescribeAcceleratorResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeAcceleratorResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeAcceleratorResponseBody) GetServiceManagedInfos() []*DescribeAcceleratorResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeAcceleratorResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *DescribeAcceleratorResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeAcceleratorResponseBody) GetTags() []*DescribeAcceleratorResponseBodyTags {
	return s.Tags
}

func (s *DescribeAcceleratorResponseBody) GetUpgradableStatus() *string {
	return s.UpgradableStatus
}

func (s *DescribeAcceleratorResponseBody) SetAcceleratorId(v string) *DescribeAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBandwidth(v int32) *DescribeAcceleratorResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBandwidthBillingType(v string) *DescribeAcceleratorResponseBody {
	s.BandwidthBillingType = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBasicBandwidthPackage(v *DescribeAcceleratorResponseBodyBasicBandwidthPackage) *DescribeAcceleratorResponseBody {
	s.BasicBandwidthPackage = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCenId(v string) *DescribeAcceleratorResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCreateTime(v int64) *DescribeAcceleratorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCrossBorderMode(v string) *DescribeAcceleratorResponseBody {
	s.CrossBorderMode = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCrossBorderStatus(v bool) *DescribeAcceleratorResponseBody {
	s.CrossBorderStatus = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCrossDomainBandwidthPackage(v *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) *DescribeAcceleratorResponseBody {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCrossPrivateState(v string) *DescribeAcceleratorResponseBody {
	s.CrossPrivateState = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDdosConfigList(v []*DescribeAcceleratorResponseBodyDdosConfigList) *DescribeAcceleratorResponseBody {
	s.DdosConfigList = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDdosId(v string) *DescribeAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDescription(v string) *DescribeAcceleratorResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDnsName(v string) *DescribeAcceleratorResponseBody {
	s.DnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetExpiredTime(v int64) *DescribeAcceleratorResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetInstanceChargeType(v string) *DescribeAcceleratorResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetIpSetConfig(v *DescribeAcceleratorResponseBodyIpSetConfig) *DescribeAcceleratorResponseBody {
	s.IpSetConfig = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetName(v string) *DescribeAcceleratorResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRegionId(v string) *DescribeAcceleratorResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRequestId(v string) *DescribeAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetResourceGroupId(v string) *DescribeAcceleratorResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSecondDnsName(v string) *DescribeAcceleratorResponseBody {
	s.SecondDnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetServiceId(v string) *DescribeAcceleratorResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetServiceManaged(v bool) *DescribeAcceleratorResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetServiceManagedInfos(v []*DescribeAcceleratorResponseBodyServiceManagedInfos) *DescribeAcceleratorResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSpec(v string) *DescribeAcceleratorResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetState(v string) *DescribeAcceleratorResponseBody {
	s.State = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetTags(v []*DescribeAcceleratorResponseBodyTags) *DescribeAcceleratorResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetUpgradableStatus(v string) *DescribeAcceleratorResponseBody {
	s.UpgradableStatus = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) Validate() error {
	if s.BasicBandwidthPackage != nil {
		if err := s.BasicBandwidthPackage.Validate(); err != nil {
			return err
		}
	}
	if s.CrossDomainBandwidthPackage != nil {
		if err := s.CrossDomainBandwidthPackage.Validate(); err != nil {
			return err
		}
	}
	if s.DdosConfigList != nil {
		for _, item := range s.DdosConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpSetConfig != nil {
		if err := s.IpSetConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceManagedInfos != nil {
		for _, item := range s.ServiceManagedInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAcceleratorResponseBodyBasicBandwidthPackage struct {
	// The bandwidth value of the basic bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of the bandwidth that is provided by the basic bandwidth plan. Valid values:
	//
	// 	- **Basic**: basic
	//
	// 	- **Enhanced**: enhanced
	//
	// 	- **Advanced**: premium
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the basic bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1d8xk8bg139j0fw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAcceleratorResponseBodyBasicBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidth(v int32) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidthType(v string) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetInstanceId(v string) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage struct {
	// The bandwidth that is provided by the cross-border acceleration bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the cross-border acceleration bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1d8xk8bg139j0fw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) SetBandwidth(v int32) *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) SetInstanceId(v string) *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type DescribeAcceleratorResponseBodyDdosConfigList struct {
	DdosId       *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeAcceleratorResponseBodyDdosConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyDdosConfigList) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyDdosConfigList) GetDdosId() *string {
	return s.DdosId
}

func (s *DescribeAcceleratorResponseBodyDdosConfigList) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeAcceleratorResponseBodyDdosConfigList) SetDdosId(v string) *DescribeAcceleratorResponseBodyDdosConfigList {
	s.DdosId = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyDdosConfigList) SetDdosRegionId(v string) *DescribeAcceleratorResponseBodyDdosConfigList {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyDdosConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeAcceleratorResponseBodyIpSetConfig struct {
	// The access mode of the acceleration area. Valid values:
	//
	// 	- **UserDefine**: custom nearby access mode. You can select acceleration areas and regions based on your business requirements. GA allocates a separate elastic IP address (EIP) to each acceleration region.
	//
	// 	- **Anycast**: automatic nearby access mode. You do not need to specify an acceleration area. GA allocates an Anycast EIP to multiple regions across the globe. Users can connect to the nearest access point of the Alibaba Cloud global transmission network by sending requests to the Anycast EIP.
	//
	// example:
	//
	// UserDefine
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s DescribeAcceleratorResponseBodyIpSetConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyIpSetConfig) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyIpSetConfig) GetAccessMode() *string {
	return s.AccessMode
}

func (s *DescribeAcceleratorResponseBodyIpSetConfig) SetAccessMode(v string) *DescribeAcceleratorResponseBodyIpSetConfig {
	s.AccessMode = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyIpSetConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAcceleratorResponseBodyServiceManagedInfos struct {
	// The name of the action performed on the managed instance. Valid values:
	//
	// 	- **Create**
	//
	// 	- **Update**
	//
	// 	- **Delete**
	//
	// 	- **Associate**
	//
	// 	- **UserUnmanaged**
	//
	// 	- **CreateChild**
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// 	- **Listener**: a listener.
	//
	// 	- **IpSet**: an acceleration region.
	//
	// 	- **EndpointGroup**: an endpoint group.
	//
	// 	- **ForwardingRule**: a forwarding rule.
	//
	// 	- **Endpoint**: an endpoint.
	//
	// 	- **EndpointGroupDestination**: a protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: a traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter is returned only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and you cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and you can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeAcceleratorResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) SetAction(v string) *DescribeAcceleratorResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeAcceleratorResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeAcceleratorResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeAcceleratorResponseBodyTags struct {
	// The key of tag N that is added to the GA instance.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the GA instance.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAcceleratorResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAcceleratorResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAcceleratorResponseBodyTags) SetKey(v string) *DescribeAcceleratorResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyTags) SetValue(v string) *DescribeAcceleratorResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
