// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcceleratorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerators(v []*ListAcceleratorsResponseBodyAccelerators) *ListAcceleratorsResponseBody
	GetAccelerators() []*ListAcceleratorsResponseBodyAccelerators
	SetPageNumber(v int32) *ListAcceleratorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAcceleratorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAcceleratorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAcceleratorsResponseBody
	GetTotalCount() *int32
}

type ListAcceleratorsResponseBody struct {
	// The information about the GA instances.
	Accelerators []*ListAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcceleratorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBody) GetAccelerators() []*ListAcceleratorsResponseBodyAccelerators {
	return s.Accelerators
}

func (s *ListAcceleratorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAcceleratorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAcceleratorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAcceleratorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAcceleratorsResponseBody) SetAccelerators(v []*ListAcceleratorsResponseBodyAccelerators) *ListAcceleratorsResponseBody {
	s.Accelerators = v
	return s
}

func (s *ListAcceleratorsResponseBody) SetPageNumber(v int32) *ListAcceleratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetPageSize(v int32) *ListAcceleratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetRequestId(v string) *ListAcceleratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetTotalCount(v int32) *ListAcceleratorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAcceleratorsResponseBody) Validate() error {
	if s.Accelerators != nil {
		for _, item := range s.Accelerators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAcceleratorsResponseBodyAccelerators struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth of the GA instance. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth metering method. Valid values:
	//
	// 	- **BandwidthPackage**: billed based on bandwidth plans.
	//
	// 	- **CDT**: billed based on data transfer.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The information about the basic bandwidth plan that is associated with the GA instance.
	BasicBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The ID of the Cloud Enterprise Network (CEN) instance that is associated with the GA instance.
	//
	// example:
	//
	// cen-hjfufhjfuwff****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The timestamp that indicates when the GA instance was created.
	//
	// example:
	//
	// 1650643200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of cross-border acceleration. This parameter is returned for GA instances whose bandwidth metering method is pay-by-data-transfer.
	//
	// 	- **bpgPro**: BGP (Multi-ISP) Pro lines.
	//
	// 	- **private**: cross-border Express Connect circuit.
	//
	// example:
	//
	// bpgPro
	CrossBorderMode *string `json:"CrossBorderMode,omitempty" xml:"CrossBorderMode,omitempty"`
	// Indicates whether cross-border acceleration is enabled for the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The information about the cross-border acceleration bandwidth plan that is associated with the GA instance.
	//
	// This array is returned only for GA instances that are created on the international site (alibabacloud.com).
	CrossDomainBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	DdosConfigList              []*ListAcceleratorsResponseBodyAcceleratorsDdosConfigList            `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro or Anti-DDOS Premium instance that is associated with the GA instance.
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
	// The CNAME that is assigned to the GA instance.
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
	IpSetConfig *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	// The name of the GA instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The ID of the region where GA instance is deployed. Only **cn-hangzhou*	- may be returned.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekztkx4zwc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The CNAME that is used to associate the GA instance with an Anti-DDoS Pro instance or an Anti-DDOS Premium instance.
	//
	// example:
	//
	// ga-bp1f609c76zg6zuna****-1.aliyunga0047.com
	SecondDnsName *string `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter takes effect only if the value of **ServiceManaged*	- is **true**.
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
	// > 	- This parameter takes effect only if the value of **ServiceManaged*	- is **true**.
	//
	// > 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The specification of the GA instance. Valid values:
	//
	// 	- **1**: Small Ⅰ.
	//
	// 	- **2**: Small Ⅱ.
	//
	// 	- **3**: Small Ⅲ.
	//
	// 	- **5**: Medium Ⅰ.
	//
	// 	- **8**: Medium Ⅱ.
	//
	// 	- **10**: Medium Ⅲ.
	//
	// 	- **20**: Large Ⅰ.
	//
	// 	- **30**: Large Ⅱ.
	//
	// 	- **40**: Large Ⅲ.
	//
	// 	- **50**: Large IV.
	//
	// 	- **60**: Large V.
	//
	// 	- **70**: Large VI.
	//
	// 	- **80**: Large VII.
	//
	// 	- **90**: Large VIII.
	//
	// 	- **100**: Super Large Ⅰ.
	//
	// 	- **200**: Super Large Ⅱ.
	//
	// >  The Large Ⅲ specification and higher specifications are available only to accounts that are added to the whitelist. To use these specifications, contact your Alibaba Cloud account manager.
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
	// The tags that are added to the resource.
	Tags []*ListAcceleratorsResponseBodyAcceleratorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// An invalid parameter.
	//
	// example:
	//
	// Invalid parameter
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Indicates whether the GA instance can be upgraded. Valid values:
	//
	// 	- **notUpgradable**: The GA instance does not need to be upgraded.
	//
	// 	- **upgradable**: The GA instance can be upgraded to the latest version.
	//
	// 	- **upgradeFailed**: The GA instance failed to be upgraded.
	//
	// example:
	//
	// notUpgradable
	UpgradableStatus *string `json:"UpgradableStatus,omitempty" xml:"UpgradableStatus,omitempty"`
}

func (s ListAcceleratorsResponseBodyAccelerators) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAccelerators) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetBasicBandwidthPackage() *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	return s.BasicBandwidthPackage
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetCenId() *string {
	return s.CenId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetCrossBorderMode() *string {
	return s.CrossBorderMode
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetCrossBorderStatus() *bool {
	return s.CrossBorderStatus
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetCrossDomainBandwidthPackage() *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	return s.CrossDomainBandwidthPackage
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetDdosConfigList() []*ListAcceleratorsResponseBodyAcceleratorsDdosConfigList {
	return s.DdosConfigList
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetDdosId() *string {
	return s.DdosId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetDescription() *string {
	return s.Description
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetDnsName() *string {
	return s.DnsName
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetIpSetConfig() *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig {
	return s.IpSetConfig
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetName() *string {
	return s.Name
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetSecondDnsName() *string {
	return s.SecondDnsName
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetServiceManagedInfos() []*ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetSpec() *string {
	return s.Spec
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetState() *string {
	return s.State
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetTags() []*ListAcceleratorsResponseBodyAcceleratorsTags {
	return s.Tags
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetType() *string {
	return s.Type
}

func (s *ListAcceleratorsResponseBodyAccelerators) GetUpgradableStatus() *string {
	return s.UpgradableStatus
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetAcceleratorId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAccelerators {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBandwidthBillingType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.BandwidthBillingType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBasicBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.BasicBandwidthPackage = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCenId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.CenId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCreateTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.CreateTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCrossBorderMode(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.CrossBorderMode = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCrossBorderStatus(v bool) *ListAcceleratorsResponseBodyAccelerators {
	s.CrossBorderStatus = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCrossDomainBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDdosConfigList(v []*ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) *ListAcceleratorsResponseBodyAccelerators {
	s.DdosConfigList = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDdosId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DdosId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDescription(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Description = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDnsName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DnsName = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetExpiredTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.ExpiredTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetInstanceChargeType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetIpSetConfig(v *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) *ListAcceleratorsResponseBodyAccelerators {
	s.IpSetConfig = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Name = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetRegionId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetResourceGroupId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetSecondDnsName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.SecondDnsName = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetServiceId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.ServiceId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetServiceManaged(v bool) *ListAcceleratorsResponseBodyAccelerators {
	s.ServiceManaged = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetServiceManagedInfos(v []*ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) *ListAcceleratorsResponseBodyAccelerators {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetSpec(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Spec = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetState(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.State = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetTags(v []*ListAcceleratorsResponseBodyAcceleratorsTags) *ListAcceleratorsResponseBodyAccelerators {
	s.Tags = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Type = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetUpgradableStatus(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.UpgradableStatus = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) Validate() error {
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

type ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage struct {
	// The bandwidth value of the basic bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth type that is provided by the basic bandwidth plan. Valid values:
	//
	// 	- **Basic**
	//
	// 	- **Enhanced**
	//
	// 	- **Advanced**
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

func (s ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidthType(v string) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetInstanceId(v string) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage struct {
	// The bandwidth value of the cross-border acceleration bandwidth plan. Unit: Mbit/s.
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

func (s ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetInstanceId(v string) *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type ListAcceleratorsResponseBodyAcceleratorsDdosConfigList struct {
	DdosId       *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) GetDdosId() *string {
	return s.DdosId
}

func (s *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) SetDdosId(v string) *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList {
	s.DdosId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) SetDdosRegionId(v string) *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList {
	s.DdosRegionId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsDdosConfigList) Validate() error {
	return dara.Validate(s)
}

type ListAcceleratorsResponseBodyAcceleratorsIpSetConfig struct {
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

func (s ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) GetAccessMode() *string {
	return s.AccessMode
}

func (s *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) SetAccessMode(v string) *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig {
	s.AccessMode = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) Validate() error {
	return dara.Validate(s)
}

type ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos struct {
	// The name of the action that is performed on the managed instance. Valid values:
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
	// 	- **Listener**: listener.
	//
	// 	- **IpSet**: acceleration region.
	//
	// 	- **EndpointGroup**: endpoint group.
	//
	// 	- **ForwardingRule**: forwarding rule.
	//
	// 	- **Endpoint**: endpoint.
	//
	// 	- **EndpointGroupDestination**: protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter takes effect only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) SetAction(v string) *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) SetChildType(v string) *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) SetIsManaged(v bool) *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type ListAcceleratorsResponseBodyAcceleratorsTags struct {
	// The key of the tag.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tast-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsTags) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) GetKey() *string {
	return s.Key
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) GetValue() *string {
	return s.Value
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) SetKey(v string) *ListAcceleratorsResponseBodyAcceleratorsTags {
	s.Key = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) SetValue(v string) *ListAcceleratorsResponseBodyAcceleratorsTags {
	s.Value = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) Validate() error {
	return dara.Validate(s)
}
