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
	// The ID of the Global Accelerator instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth of the standard Global Accelerator instance. Unit: Mbps.
	//
	// > This parameter is valid only when the access mode of the acceleration area is Anycast.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the bandwidth. Valid values:
	//
	// - **BandwidthPackage**: pay-by-bandwidth-plan.
	//
	// - **CDT**: pay-by-data-transfer.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The details of the basic bandwidth plan that is associated with the Global Accelerator instance.
	BasicBandwidthPackage *DescribeAcceleratorResponseBodyBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The ID of the Cloud Enterprise Network (CEN) instance that is associated with the Global Accelerator instance.
	//
	// example:
	//
	// cen-hjkduu767hc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The timestamp that indicates when the Global Accelerator instance was created.
	//
	// example:
	//
	// 1650643200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of cross-border acceleration. This parameter is returned for pay-by-data-transfer instances.
	//
	// **bpgPro**: premium bandwidth for cross-border acceleration.
	//
	// example:
	//
	// bpgPro
	CrossBorderMode *string `json:"CrossBorderMode,omitempty" xml:"CrossBorderMode,omitempty"`
	// Indicates whether the cross-border line feature is enabled for the Global Accelerator instance. Valid values:
	//
	// - **true**: The cross-border line feature is enabled. You can use Global Accelerator to accelerate data transmission across borders.
	//
	// - **false**: The cross-border line feature is disabled. You cannot use Global Accelerator to accelerate data transmission across borders.
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The details of the cross-region bandwidth plan that is associated with the Global Accelerator instance.
	//
	// This parameter is returned only by the Alibaba Cloud International Website (www\\.alibabacloud.com).
	CrossDomainBandwidthPackage *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	// Indicates whether cross-border bandwidth is enabled.
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	CrossPrivateState *string `json:"CrossPrivateState,omitempty" xml:"CrossPrivateState,omitempty"`
	// The list of Anti-DDoS instances that are associated with the Global Accelerator instance.
	DdosConfigList []*DescribeAcceleratorResponseBodyDdosConfigList `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the Anti-DDoS instance that is associated with the Global Accelerator instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The description of the Global Accelerator instance.
	//
	// example:
	//
	// Accelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The canonical name (CNAME) that is assigned to the Global Accelerator instance.
	//
	// example:
	//
	// ga-bp15u1i2hmtbk8c3i****.aliyunga0019.com
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// The timestamp that indicates when the Global Accelerator instance expires.
	//
	// example:
	//
	// 1653235200
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the Global Accelerator instance.
	//
	// example:
	//
	// PREPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The configuration of the acceleration area.
	IpSetConfig *DescribeAcceleratorResponseBodyIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	// The name of the Global Accelerator instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The region where the Global Accelerator instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
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
	// The CNAME of the Anti-DDoS instance that is associated with the Global Accelerator instance.
	//
	// example:
	//
	// ga-bp1f609c76zg6zuna****-1.aliyunga0047.com
	SecondDnsName *string `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	// The ID of the service that manages the instance.
	//
	// > This parameter is valid only when **ServiceManaged*	- is set to **True**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the instance is a managed instance. Valid values:
	//
	// - **true**: The instance is a managed instance.
	//
	// - **false**: The instance is not a managed instance.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that you can perform on the managed instance.
	//
	// > - This parameter is valid only when **ServiceManaged*	- is set to **True**.
	//
	// >
	//
	// > - When the instance is managed, you cannot perform some operations on the instance.
	ServiceManagedInfos []*DescribeAcceleratorResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The specification of the Global Accelerator instance. Valid values:
	//
	// - **1**: Small I.
	//
	// - **2**: Small II.
	//
	// - **3**: Small III.
	//
	// - **5**: Medium I.
	//
	// - **8**: Medium II.
	//
	// - **10**: Medium III.
	//
	// - **20**: Large I.
	//
	// - **30**: Large II.
	//
	// - **40**: Large III.
	//
	// - **50**: Large IV.
	//
	// - **60**: Large V.
	//
	// - **70**: Large VI.
	//
	// - **80**: Large VII.
	//
	// - **90**: Large VIII.
	//
	// - **100**: Super Large I.
	//
	// - **200**: Super Large II.
	//
	// > The Large III and higher specifications are available only to users on the whitelist. To use these specifications, contact your account manager.
	//
	// The definitions of different specifications vary. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/153127.html).
	//
	// example:
	//
	// 1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the Global Accelerator instance. Valid values:
	//
	// - **init**: The instance is being initialized.
	//
	// - **active**: The instance is available.
	//
	// - **configuring**: The instance is being configured.
	//
	// - **binding**: The instance is being associated.
	//
	// - **unbinding**: The instance is being disassociated.
	//
	// - **deleting**: The instance is being deleted.
	//
	// - **finacialLocked**: The instance is financially locked.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the resource.
	Tags []*DescribeAcceleratorResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The upgrade status of the Global Accelerator instance. Valid values:
	//
	// - **notUpgradable**: The instance does not need to be upgraded.
	//
	// - **upgradable**: The instance can be upgraded.
	//
	// - **upgradeFailed**: The instance failed to be upgraded.
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
	// The bandwidth of the basic bandwidth plan. Unit: Mbps.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of the bandwidth of the basic bandwidth plan. Valid values:
	//
	// - **Basic**: standard acceleration bandwidth.
	//
	// - **Enhanced**: enhanced acceleration bandwidth.
	//
	// - **Advanced**: premium acceleration bandwidth.
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
	// The bandwidth of the cross-region bandwidth plan. Unit: Mbps.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the cross-region bandwidth plan.
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
	// The ID of the Anti-DDoS instance that is associated with the Global Accelerator instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The region where the Anti-DDoS instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// ap-southeast-1
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
	// - **UserDefine**: custom. You can select acceleration areas and regions based on your business needs. Global Accelerator provides a separate elastic IP address (EIP) for each acceleration region.
	//
	// - **Anycast**: automatic. You do not need to configure an acceleration area. Global Accelerator provides an Anycast EIP for multiple regions. Users can connect to the nearest access point of the Alibaba Cloud network using the Anycast EIP.
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
	// The name of the action on the managed instance. Valid values:
	//
	// - **Create**: creates an instance.
	//
	// - **Update**: updates the current instance.
	//
	// - **Delete**: deletes the current instance.
	//
	// - **Associate**: associates the instance with other resources.
	//
	// - **UserUnmanaged**: unmanages the instance.
	//
	// - **CreateChild**: creates a child resource in the instance.
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// - **Listener**: listener.
	//
	// - **IpSet**: acceleration region.
	//
	// - **EndpointGroup**: endpoint group.
	//
	// - **ForwardingRule**: forwarding rule.
	//
	// - **Endpoint**: endpoint.
	//
	// - **EndpointGroupDestination**: protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// - **EndpointPolicy**: access policy of an endpoint associated with a custom routing listener.
	//
	// > This parameter is valid only when **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified action is managed. Valid values:
	//
	// - **true**: The action is managed. You cannot perform the specified action on the managed instance.
	//
	// - **false**: The action is not managed. You can perform the specified action on the managed instance.
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
	// The tag key.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
