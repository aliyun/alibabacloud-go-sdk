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
	// The details of the Global Accelerator instances.
	Accelerators []*ListAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
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
	// The ID of the Global Accelerator instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth of the Global Accelerator instance. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method for the bandwidth.
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
	BasicBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The Cloud Enterprise Network (CEN) instance that is bound to the Global Accelerator instance.
	//
	// example:
	//
	// cen-hjfufhjfuwff****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The timestamp that indicates when the Global Accelerator instance was created.
	//
	// example:
	//
	// 1650643200
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of cross-border acceleration for the instance that uses the pay-by-data-transfer billing method. Valid values:
	//
	// - **bpgPro**: premium bandwidth for cross-border acceleration.
	//
	// - **private**: <props="china">China Unicom leased line for cross-border acceleration.<props="intl">Leased line for cross-domain acceleration.
	//
	// example:
	//
	// bpgPro
	CrossBorderMode *string `json:"CrossBorderMode,omitempty" xml:"CrossBorderMode,omitempty"`
	// Indicates whether cross-border data transfer is enabled for the Global Accelerator instance. Valid values:
	//
	// - **true**: Cross-border data transfer is enabled, which can accelerate data transfer across borders.
	//
	// - **false**: Cross-border data transfer is disabled, which cannot accelerate data transfer across borders.
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The details of the cross-domain acceleration bandwidth plan that is associated with the Global Accelerator instance.
	//
	// This array is returned only by the Alibaba Cloud International Website (www\\.alibabacloud.com).
	CrossDomainBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	// The list of Anti-DDoS Pro/Premium configurations.
	DdosConfigList []*ListAcceleratorsResponseBodyAcceleratorsDdosConfigList `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro/Premium instance that is associated with the Global Accelerator instance.
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
	// The canonical name (CNAME) assigned to the Global Accelerator instance.
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
	IpSetConfig *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	// The name of the Global Accelerator instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// The region ID of the Global Accelerator instance. The value is set to **cn-hangzhou**.
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
	// The CNAME that is assigned to the Global Accelerator instance after it is associated with an Anti-DDoS Pro/Premium instance.
	//
	// example:
	//
	// ga-bp1f609c76zg6zuna****-1.aliyunga0047.com
	SecondDnsName *string `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	// The ID of the service that manages the instance.
	//
	// > This parameter is returned only when **ServiceManaged*	- is set to **True**.
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
	// > - This parameter is returned only when **ServiceManaged*	- is set to **True**.
	//
	// >
	//
	// > - When the instance is in a managed state, you have limited permissions to perform operations on the instance.
	ServiceManagedInfos []*ListAcceleratorsResponseBodyAcceleratorsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The instance type of the Global Accelerator instance. Valid values:
	//
	// - **1**: Small I
	//
	// - **2**: Small II
	//
	// - **3**: Small III
	//
	// - **5**: Medium I
	//
	// - **8**: Medium II
	//
	// - **10**: Medium III
	//
	// - **20**: Large I
	//
	// - **30**: Large II
	//
	// - **40**: Large III
	//
	// - **50**: Large IV
	//
	// - **60**: Large V
	//
	// - **70**: Large VI
	//
	// - **80**: Large VII
	//
	// - **90**: Large VIII
	//
	// - **100**: Extra Large I
	//
	// - **200**: Extra Large II
	//
	// > Currently, the Large III and higher instance types are available only to users on the whitelist. To use these instance types, contact your account manager.
	//
	// Different instance types have different definitions. For more information, see [Instance types](https://help.aliyun.com/document_detail/153127.html).
	//
	// example:
	//
	// 1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The state of the Global Accelerator instance. Valid values:
	//
	// - **init**: The instance is being initialized.
	//
	// - **active**: The instance is active.
	//
	// - **configuring**: The instance is being configured.
	//
	// - **binding**: The instance is being bound.
	//
	// - **unbinding**: The instance is being unbound.
	//
	// - **deleting**: The instance is being deleted.
	//
	// - **finacialLocked**: The instance is locked due to an overdue payment.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the resource.
	Tags []*ListAcceleratorsResponseBodyAcceleratorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is invalid.
	//
	// example:
	//
	// None
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The upgrade status of the Global Accelerator instance. Valid values:
	//
	// - **notUpgradable**: The instance does not need to be upgraded.
	//
	// - **upgradable**: The instance can be upgraded to the latest version.
	//
	// - **upgradeFailed**: The instance failed to be upgraded.
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
	// The bandwidth of the basic bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of the bandwidth. Valid values:
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
	// The bandwidth of the cross-domain acceleration bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the cross-domain acceleration bandwidth plan.
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
	// The ID of the Anti-DDoS Pro/Premium instance.
	//
	// example:
	//
	// ddoscoo-cn-a8w4ekcb**
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The region ID of the Anti-DDoS Pro/Premium instance.
	//
	// example:
	//
	// cn-hangzhou
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
	// - **UserDefine**: custom nearby access mode. You can select acceleration areas and regions based on your business needs. Global Accelerator provides a separate elastic IP address (EIP) for each acceleration region.
	//
	// - **Anycast**: automatic nearby access mode. You do not need to configure an acceleration area. Global Accelerator provides an Anycast EIP for multiple regions across the globe. Users can connect to the nearest access point of the Alibaba Cloud network using the Anycast EIP.
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
	// The name of the action on the managed instance. Valid values:
	//
	// - **Create**: create an instance.
	//
	// - **Update**: update the current instance.
	//
	// - **Delete**: delete the current instance.
	//
	// - **Associate**: associate the instance with other resources.
	//
	// - **UserUnmanaged**: disassociate the instance from the service.
	//
	// - **CreateChild**: create a child resource in the instance.
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
	// - **EndpointPolicy**: traffic policy of an endpoint associated with a custom routing listener.
	//
	// > This parameter is returned only when **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified action is managed. Valid values:
	//
	// - **true**: The specified action is managed, and you cannot perform the specified action on the managed instance.
	//
	// - **false**: The specified action is not managed, and you can perform the specified action on the managed instance.
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
	// The tag key.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
