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
	SetHostname(v string) *DescribeRenderingInstanceResponseBody
	GetHostname() *string
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
	SetStorageSize(v int32) *DescribeRenderingInstanceResponseBody
	GetStorageSize() *int32
	SetSystemInfo(v *DescribeRenderingInstanceResponseBodySystemInfo) *DescribeRenderingInstanceResponseBody
	GetSystemInfo() *DescribeRenderingInstanceResponseBodySystemInfo
}

type DescribeRenderingInstanceResponseBody struct {
	AdditionalIngresses []*DescribeRenderingInstanceResponseBodyAdditionalIngresses `json:"AdditionalIngresses,omitempty" xml:"AdditionalIngresses,omitempty" type:"Repeated"`
	ConfigInfo          *DescribeRenderingInstanceResponseBodyConfigInfo            `json:"ConfigInfo,omitempty" xml:"ConfigInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-07T02:27:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EgressIp     *string `json:"EgressIp,omitempty" xml:"EgressIp,omitempty"`
	// example:
	//
	// cn-xxx.ecr.aliyuncs.com
	Hostname     *string                                              `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InternalIp   *string                                              `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	Isp          *string                                              `json:"Isp,omitempty" xml:"Isp,omitempty"`
	PortMappings []*DescribeRenderingInstanceResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string                                               `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	RenderingSpec       *string                                               `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	RenderingStatus     *DescribeRenderingInstanceResponseBodyRenderingStatus `json:"RenderingStatus,omitempty" xml:"RenderingStatus,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageSize *int32                                           `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SystemInfo  *DescribeRenderingInstanceResponseBodySystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Struct"`
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

func (s *DescribeRenderingInstanceResponseBody) GetHostname() *string {
	return s.Hostname
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

func (s *DescribeRenderingInstanceResponseBody) SetHostname(v string) *DescribeRenderingInstanceResponseBody {
	s.Hostname = &v
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

func (s *DescribeRenderingInstanceResponseBody) SetStorageSize(v int32) *DescribeRenderingInstanceResponseBody {
	s.StorageSize = &v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) SetSystemInfo(v *DescribeRenderingInstanceResponseBodySystemInfo) *DescribeRenderingInstanceResponseBody {
	s.SystemInfo = v
	return s
}

func (s *DescribeRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyAdditionalIngresses struct {
	Hostname     *string                                                                 `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	Isp          *string                                                                 `json:"Isp,omitempty" xml:"Isp,omitempty"`
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
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyAdditionalIngressesPortMappings struct {
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
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
	Configuration []*DescribeRenderingInstanceResponseBodyConfigInfoConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	NetworkConfig *DescribeRenderingInstanceResponseBodyConfigInfoNetworkConfig   `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyConfigInfoConfiguration struct {
	Attributes []*DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type DescribeRenderingInstanceResponseBodyConfigInfoConfigurationAttributes struct {
	// example:
	//
	// lon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// success
	BandwidthStatus *string `json:"BandwidthStatus,omitempty" xml:"BandwidthStatus,omitempty"`
	// example:
	//
	// 100
	MaxEgressBandwidth *int32 `json:"MaxEgressBandwidth,omitempty" xml:"MaxEgressBandwidth,omitempty"`
	// example:
	//
	// 100
	MaxIngressBandwidth *int32 `json:"MaxIngressBandwidth,omitempty" xml:"MaxIngressBandwidth,omitempty"`
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
	// example:
	//
	// 10013/10020
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MigrateLocalData
	LatestAction *string `json:"LatestAction,omitempty" xml:"LatestAction,omitempty"`
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

type DescribeRenderingInstanceResponseBodySystemInfo struct {
	// example:
	//
	// 60
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
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
