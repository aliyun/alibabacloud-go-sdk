// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalIngresses(v []*DescribeRenderingInstanceResponseBodyAdditionalIngresses) *DescribeRenderingInstanceResponseBody
	GetAdditionalIngresses() []*DescribeRenderingInstanceResponseBodyAdditionalIngresses
	SetConfigInfo(v *DescribeRenderingInstanceResponseBodyConfigInfo) *DescribeRenderingInstanceResponseBody
	GetConfigInfo() *DescribeRenderingInstanceResponseBodyConfigInfo
	SetCreationTime(v string) *DescribeRenderingInstanceResponseBody
	GetCreationTime() *string
	SetEgressIp(v string) *DescribeRenderingInstanceResponseBody
	GetEgressIp() *string
	SetExpiredTime(v string) *DescribeRenderingInstanceResponseBody
	GetExpiredTime() *string
	SetHostname(v string) *DescribeRenderingInstanceResponseBody
	GetHostname() *string
	SetInstanceChargeType(v string) *DescribeRenderingInstanceResponseBody
	GetInstanceChargeType() *string
	SetInstanceName(v string) *DescribeRenderingInstanceResponseBody
	GetInstanceName() *string
	SetInternalIp(v string) *DescribeRenderingInstanceResponseBody
	GetInternalIp() *string
	SetIsp(v string) *DescribeRenderingInstanceResponseBody
	GetIsp() *string
	SetPortMappings(v []*DescribeRenderingInstanceResponseBodyPortMappings) *DescribeRenderingInstanceResponseBody
	GetPortMappings() []*DescribeRenderingInstanceResponseBodyPortMappings
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceResponseBody
	GetRenderingInstanceId() *string
	SetRenderingSpec(v string) *DescribeRenderingInstanceResponseBody
	GetRenderingSpec() *string
	SetRenderingStatus(v *DescribeRenderingInstanceResponseBodyRenderingStatus) *DescribeRenderingInstanceResponseBody
	GetRenderingStatus() *DescribeRenderingInstanceResponseBodyRenderingStatus
	SetRequestId(v string) *DescribeRenderingInstanceResponseBody
	GetRequestId() *string
	SetResourceAttributes(v *DescribeRenderingInstanceResponseBodyResourceAttributes) *DescribeRenderingInstanceResponseBody
	GetResourceAttributes() *DescribeRenderingInstanceResponseBodyResourceAttributes
	SetResourceStatus(v *DescribeRenderingInstanceResponseBodyResourceStatus) *DescribeRenderingInstanceResponseBody
	GetResourceStatus() *DescribeRenderingInstanceResponseBodyResourceStatus
	SetStorageSize(v int32) *DescribeRenderingInstanceResponseBody
	GetStorageSize() *int32
	SetSystemInfo(v *DescribeRenderingInstanceResponseBodySystemInfo) *DescribeRenderingInstanceResponseBody
	GetSystemInfo() *DescribeRenderingInstanceResponseBodySystemInfo
}

type DescribeRenderingInstanceResponseBody struct {
	// The additional optional ingress network information.
	AdditionalIngresses []*DescribeRenderingInstanceResponseBodyAdditionalIngresses `json:"AdditionalIngresses,omitempty" xml:"AdditionalIngresses,omitempty" type:"Repeated"`
	// The instance configuration.
	ConfigInfo *DescribeRenderingInstanceResponseBodyConfigInfo `json:"ConfigInfo,omitempty" xml:"ConfigInfo,omitempty" type:"Struct"`
	// The instance creation time. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2024-05-07T02:27:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The egress IP address.
	//
	// example:
	//
	// 1.1.8.8
	EgressIp *string `json:"EgressIp,omitempty" xml:"EgressIp,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2029-03-28T16:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The domain name or access IP address of the cloud application service instance.
	//
	// example:
	//
	// cn-xxx.ecr.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The instance billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// TestInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 10.1.17.32
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The ISP code. Valid values:
	//
	// - cmcc
	//
	// - unicom
	//
	// - telecom
	//
	// example:
	//
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The port mappings.
	PortMappings []*DescribeRenderingInstanceResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The instance type of the cloud application service instance.
	//
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	// The instance operation status.
	RenderingStatus *DescribeRenderingInstanceResponseBodyRenderingStatus `json:"RenderingStatus,omitempty" xml:"RenderingStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource attributes of the cloud application service instance.
	ResourceAttributes *DescribeRenderingInstanceResponseBodyResourceAttributes `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Struct"`
	// The compute resource status information.
	ResourceStatus *DescribeRenderingInstanceResponseBodyResourceStatus `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty" type:"Struct"`
	// The storage capacity of the cloud application service instance.
	//
	// example:
	//
	// 20
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The system information, including the resolution.
	SystemInfo *DescribeRenderingInstanceResponseBodySystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Struct"`
}

func (s DescribeRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBody) GetAdditionalIngresses() []*DescribeRenderingInstanceResponseBodyAdditionalIngresses {
	return s.AdditionalIngresses
}

func (s *DescribeRenderingInstanceResponseBody) GetConfigInfo() *DescribeRenderingInstanceResponseBodyConfigInfo {
	return s.ConfigInfo
}

func (s *DescribeRenderingInstanceResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRenderingInstanceResponseBody) GetEgressIp() *string {
	return s.EgressIp
}

func (s *DescribeRenderingInstanceResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeRenderingInstanceResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeRenderingInstanceResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRenderingInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRenderingInstanceResponseBody) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeRenderingInstanceResponseBody) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRenderingInstanceResponseBody) GetPortMappings() []*DescribeRenderingInstanceResponseBodyPortMappings {
	return s.PortMappings
}

func (s *DescribeRenderingInstanceResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceResponseBody) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *DescribeRenderingInstanceResponseBody) GetRenderingStatus() *DescribeRenderingInstanceResponseBodyRenderingStatus {
	return s.RenderingStatus
}

func (s *DescribeRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenderingInstanceResponseBody) GetResourceAttributes() *DescribeRenderingInstanceResponseBodyResourceAttributes {
	return s.ResourceAttributes
}

func (s *DescribeRenderingInstanceResponseBody) GetResourceStatus() *DescribeRenderingInstanceResponseBodyResourceStatus {
	return s.ResourceStatus
}

func (s *DescribeRenderingInstanceResponseBody) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeRenderingInstanceResponseBody) GetSystemInfo() *DescribeRenderingInstanceResponseBodySystemInfo {
	return s.SystemInfo
}

func (s *DescribeRenderingInstanceResponseBody) SetAdditionalIngresses(v []*DescribeRenderingInstanceResponseBodyAdditionalIngresses) *DescribeRenderingInstanceResponseBody {
	s.AdditionalIngresses = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetConfigInfo(v *DescribeRenderingInstanceResponseBodyConfigInfo) *DescribeRenderingInstanceResponseBody {
	s.ConfigInfo = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetCreationTime(v string) *DescribeRenderingInstanceResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetEgressIp(v string) *DescribeRenderingInstanceResponseBody {
	s.EgressIp = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetExpiredTime(v string) *DescribeRenderingInstanceResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetHostname(v string) *DescribeRenderingInstanceResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetInstanceChargeType(v string) *DescribeRenderingInstanceResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetInstanceName(v string) *DescribeRenderingInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetInternalIp(v string) *DescribeRenderingInstanceResponseBody {
	s.InternalIp = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetIsp(v string) *DescribeRenderingInstanceResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetPortMappings(v []*DescribeRenderingInstanceResponseBodyPortMappings) *DescribeRenderingInstanceResponseBody {
	s.PortMappings = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetRenderingInstanceId(v string) *DescribeRenderingInstanceResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetRenderingSpec(v string) *DescribeRenderingInstanceResponseBody {
	s.RenderingSpec = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetRenderingStatus(v *DescribeRenderingInstanceResponseBodyRenderingStatus) *DescribeRenderingInstanceResponseBody {
	s.RenderingStatus = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetRequestId(v string) *DescribeRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetResourceAttributes(v *DescribeRenderingInstanceResponseBodyResourceAttributes) *DescribeRenderingInstanceResponseBody {
	s.ResourceAttributes = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetResourceStatus(v *DescribeRenderingInstanceResponseBodyResourceStatus) *DescribeRenderingInstanceResponseBody {
	s.ResourceStatus = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetStorageSize(v int32) *DescribeRenderingInstanceResponseBody {
	s.StorageSize = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetSystemInfo(v *DescribeRenderingInstanceResponseBodySystemInfo) *DescribeRenderingInstanceResponseBody {
	s.SystemInfo = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) Validate() error {
	if s.AdditionalIngresses != nil {
		for _, item := range s.AdditionalIngresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConfigInfo != nil {
		if err := s.ConfigInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PortMappings != nil {
		for _, item := range s.PortMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RenderingStatus != nil {
		if err := s.RenderingStatus.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceAttributes != nil {
		if err := s.ResourceAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceStatus != nil {
		if err := s.ResourceStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SystemInfo != nil {
		if err := s.SystemInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenderingInstanceResponseBodyAdditionalIngresses struct {
	// The domain name or IP address of the cloud application service instance.
	//
	// example:
	//
	// 101.66.165.213
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The ISP code. Valid values:
	//
	// - cmcc
	//
	// - unicom
	//
	// - telecom
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The port mappings.
	PortMappings []*DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
}

func (s DescribeRenderingInstanceResponseBodyAdditionalIngresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyAdditionalIngresses) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) GetPortMappings() []*DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings {
	return s.PortMappings
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) SetHostname(v string) *DescribeRenderingInstanceResponseBodyAdditionalIngresses {
	s.Hostname = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) SetIsp(v string) *DescribeRenderingInstanceResponseBodyAdditionalIngresses {
	s.Isp = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) SetPortMappings(v []*DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) *DescribeRenderingInstanceResponseBodyAdditionalIngresses {
	s.PortMappings = v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngresses) Validate() error {
	if s.PortMappings != nil {
		for _, item := range s.PortMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings struct {
	// The external port or port range, such as 22. If a port range is specified, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// example:
	//
	// 12500/12519
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The internal port or port range. The internal ports correspond to the external ports one by one. If a port range is specified, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// example:
	//
	// 11120/11139
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) SetExternalPort(v string) *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings {
	s.ExternalPort = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) SetInternalPort(v string) *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings {
	s.InternalPort = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyConfigInfo struct {
	// The list of configured physical device simulation modules.
	Configuration []*DescribeRenderingInstanceResponseBodyConfigInfoConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// The outbound and inbound bandwidth throttling. Unit: Mbit/s.
	NetworkConfig *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
}

func (s DescribeRenderingInstanceResponseBodyConfigInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyConfigInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfo) GetConfiguration() []*DescribeRenderingInstanceResponseBodyConfigInfoConfiguration {
	return s.Configuration
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfo) GetNetworkConfig() *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig {
	return s.NetworkConfig
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfo) SetConfiguration(v []*DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) *DescribeRenderingInstanceResponseBodyConfigInfo {
	s.Configuration = v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfo) SetNetworkConfig(v *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) *DescribeRenderingInstanceResponseBodyConfigInfo {
	s.NetworkConfig = v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfo) Validate() error {
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenderingInstanceResponseBodyConfigInfoConfiguration struct {
	// The attributes.
	Attributes []*DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the physical device simulation module. Valid values:
	//
	// - ctl: control module
	//
	// - prop: property module
	//
	// - location: location module
	//
	// - battery: battery module
	//
	// - network: network module
	//
	// - bluetooth: Bluetooth module
	//
	// - sim: SIM card module
	//
	// - display: display module
	//
	// - system: basic module
	//
	// example:
	//
	// location
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) GetAttributes() []*DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes {
	return s.Attributes
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) SetAttributes(v []*DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration {
	s.Attributes = v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) SetModuleName(v string) *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration {
	s.ModuleName = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfiguration) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes struct {
	// The attribute name.
	//
	// example:
	//
	// lon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// 100
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) GetName() *string {
	return s.Name
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) GetValue() interface{} {
	return s.Value
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) SetName(v string) *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes {
	s.Name = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) SetValue(v interface{}) *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes {
	s.Value = v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig struct {
	// The bandwidth application status. Valid values:
	//
	// - waiting: taking effect
	//
	// - success: taken effect
	//
	// - failed: failed to take effect
	//
	// example:
	//
	// success
	BandwidthStatus *string `json:"BandwidthStatus,omitempty" xml:"BandwidthStatus,omitempty"`
	// The maximum outbound bandwidth limit. Unit: Mbit/s. A value of 0 indicates that no bandwidth limit is configured.
	//
	// example:
	//
	// 100
	MaxEgressBandwidth *int32 `json:"MaxEgressBandwidth,omitempty" xml:"MaxEgressBandwidth,omitempty"`
	// The maximum inbound bandwidth limit. Unit: Mbit/s. A value of 0 indicates that no bandwidth limit is configured.
	//
	// example:
	//
	// 100
	MaxIngressBandwidth *int32 `json:"MaxIngressBandwidth,omitempty" xml:"MaxIngressBandwidth,omitempty"`
	// The update time. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-08-17T09:54:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) GetBandwidthStatus() *string {
	return s.BandwidthStatus
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) GetMaxEgressBandwidth() *int32 {
	return s.MaxEgressBandwidth
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) GetMaxIngressBandwidth() *int32 {
	return s.MaxIngressBandwidth
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) SetBandwidthStatus(v string) *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig {
	s.BandwidthStatus = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) SetMaxEgressBandwidth(v int32) *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig {
	s.MaxEgressBandwidth = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) SetMaxIngressBandwidth(v int32) *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig {
	s.MaxIngressBandwidth = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) SetUpdateTime(v string) *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig {
	s.UpdateTime = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyPortMappings struct {
	// The external port or port range, such as 22. If a port range is specified, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// example:
	//
	// 10013/10020
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The internal port or port range. The internal ports correspond to the external ports one by one. If a port range is specified, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// example:
	//
	// 49008/49015
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyPortMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyPortMappings) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeRenderingInstanceResponseBodyPortMappings) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeRenderingInstanceResponseBodyPortMappings) SetExternalPort(v string) *DescribeRenderingInstanceResponseBodyPortMappings {
	s.ExternalPort = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyPortMappings) SetInternalPort(v string) *DescribeRenderingInstanceResponseBodyPortMappings {
	s.InternalPort = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyPortMappings) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyRenderingStatus struct {
	// The status description.
	//
	// example:
	//
	// Working
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the latest instance operation.
	//
	// example:
	//
	// MigrateLocalData
	LatestAction *string `json:"LatestAction,omitempty" xml:"LatestAction,omitempty"`
	// The instance operation status. Valid values:
	//
	// - Preparing: initializing
	//
	// - Rebooting: restarting
	//
	// - Resetting: resetting
	//
	// - Working: working. This is a final state.
	//
	// - Failure: operation failed. This is a final state.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyRenderingStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyRenderingStatus) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) GetLatestAction() *string {
	return s.LatestAction
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) SetDescription(v string) *DescribeRenderingInstanceResponseBodyRenderingStatus {
	s.Description = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) SetLatestAction(v string) *DescribeRenderingInstanceResponseBodyRenderingStatus {
	s.LatestAction = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) SetStatus(v string) *DescribeRenderingInstanceResponseBodyRenderingStatus {
	s.Status = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyRenderingStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyResourceAttributes struct {
	// The streaming service. Valid values:
	//
	// - ON: enabled
	//
	// - OFF: disabled
	//
	// example:
	//
	// ON
	EdgeMediaService *string `json:"EdgeMediaService,omitempty" xml:"EdgeMediaService,omitempty"`
	// The inbound access configuration. Valid values:
	//
	// - ON: enabled. The resource can be accessed from external networks.
	//
	// - OFF: disabled.
	//
	// example:
	//
	// ON
	InAccess *string `json:"InAccess,omitempty" xml:"InAccess,omitempty"`
	// The outbound access configuration. Valid values:
	//
	// - ON: enabled. The resource can access the Internet.
	//
	// - OFF: disabled.
	//
	// example:
	//
	// ON
	OutAccess *string `json:"OutAccess,omitempty" xml:"OutAccess,omitempty"`
	// The resource zone. Valid values:
	//
	// - Private
	//
	// - Public
	//
	// example:
	//
	// Public
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyResourceAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyResourceAttributes) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) GetEdgeMediaService() *string {
	return s.EdgeMediaService
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) GetInAccess() *string {
	return s.InAccess
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) GetOutAccess() *string {
	return s.OutAccess
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) GetZone() *string {
	return s.Zone
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) SetEdgeMediaService(v string) *DescribeRenderingInstanceResponseBodyResourceAttributes {
	s.EdgeMediaService = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) SetInAccess(v string) *DescribeRenderingInstanceResponseBodyResourceAttributes {
	s.InAccess = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) SetOutAccess(v string) *DescribeRenderingInstanceResponseBodyResourceAttributes {
	s.OutAccess = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) SetZone(v string) *DescribeRenderingInstanceResponseBodyResourceAttributes {
	s.Zone = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyResourceAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyResourceStatus struct {
	// The running state of the compute resource. Valid values:
	//
	// 1. running: The edge instance status is normal.
	//
	// 2. operating: The edge instance is under maintenance.
	//
	// 3. error: An edge anomaly is detected.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodyResourceStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodyResourceStatus) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodyResourceStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeRenderingInstanceResponseBodyResourceStatus) SetStatus(v string) *DescribeRenderingInstanceResponseBodyResourceStatus {
	s.Status = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodyResourceStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodySystemInfo struct {
	// The refresh rate of the instance, in Hz.
	//
	// example:
	//
	// 60
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The resolution. Valid values:
	//
	// - 1920*864
	//
	// - 1080*1920
	//
	// - 1920*1080
	//
	// - 720*1280
	//
	// - 2400*1080
	//
	// - 1080*2400
	//
	// - 1280*720
	//
	// - 864*1920
	//
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s DescribeRenderingInstanceResponseBodySystemInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponseBodySystemInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponseBodySystemInfo) GetFrequency() *int32 {
	return s.Frequency
}

func (s *DescribeRenderingInstanceResponseBodySystemInfo) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeRenderingInstanceResponseBodySystemInfo) SetFrequency(v int32) *DescribeRenderingInstanceResponseBodySystemInfo {
	s.Frequency = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodySystemInfo) SetResolution(v string) *DescribeRenderingInstanceResponseBodySystemInfo {
	s.Resolution = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBodySystemInfo) Validate() error {
	return dara.Validate(s)
}
