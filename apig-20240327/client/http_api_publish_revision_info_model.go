// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiPublishRevisionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBackendScene(v string) *HttpApiPublishRevisionInfo
	GetBackendScene() *string
	SetBackendType(v string) *HttpApiPublishRevisionInfo
	GetBackendType() *string
	SetCloudProductConfig(v *HttpApiPublishRevisionInfoCloudProductConfig) *HttpApiPublishRevisionInfo
	GetCloudProductConfig() *HttpApiPublishRevisionInfoCloudProductConfig
	SetCreateTimestamp(v int64) *HttpApiPublishRevisionInfo
	GetCreateTimestamp() *int64
	SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo
	GetCustomDomains() []*HttpApiDomainInfo
	SetDnsConfigs(v []*HttpApiPublishRevisionInfoDnsConfigs) *HttpApiPublishRevisionInfo
	GetDnsConfigs() []*HttpApiPublishRevisionInfoDnsConfigs
	SetEnvironmentInfo(v *HttpApiPublishRevisionInfoEnvironmentInfo) *HttpApiPublishRevisionInfo
	GetEnvironmentInfo() *HttpApiPublishRevisionInfoEnvironmentInfo
	SetIsCurrentVersion(v bool) *HttpApiPublishRevisionInfo
	GetIsCurrentVersion() *bool
	SetOperations(v []*HttpApiOperationInfo) *HttpApiPublishRevisionInfo
	GetOperations() []*HttpApiOperationInfo
	SetRevisionId(v string) *HttpApiPublishRevisionInfo
	GetRevisionId() *string
	SetServiceConfigs(v []*HttpApiPublishRevisionInfoServiceConfigs) *HttpApiPublishRevisionInfo
	GetServiceConfigs() []*HttpApiPublishRevisionInfoServiceConfigs
	SetSubDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo
	GetSubDomains() []*HttpApiDomainInfo
	SetVipConfigs(v []*HttpApiPublishRevisionInfoVipConfigs) *HttpApiPublishRevisionInfo
	GetVipConfigs() []*HttpApiPublishRevisionInfoVipConfigs
}

type HttpApiPublishRevisionInfo struct {
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// example:
	//
	// Service
	BackendType        *string                                       `json:"backendType,omitempty" xml:"backendType,omitempty"`
	CloudProductConfig *HttpApiPublishRevisionInfoCloudProductConfig `json:"cloudProductConfig,omitempty" xml:"cloudProductConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1718807057927
	CreateTimestamp *int64                                  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	CustomDomains   []*HttpApiDomainInfo                    `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	DnsConfigs      []*HttpApiPublishRevisionInfoDnsConfigs `json:"dnsConfigs,omitempty" xml:"dnsConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentInfo *HttpApiPublishRevisionInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsCurrentVersion *bool                   `json:"isCurrentVersion,omitempty" xml:"isCurrentVersion,omitempty"`
	Operations       []*HttpApiOperationInfo `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// example:
	//
	// apr-xxx
	RevisionId     *string                                     `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
	ServiceConfigs []*HttpApiPublishRevisionInfoServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx.com
	SubDomains []*HttpApiDomainInfo                    `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
	VipConfigs []*HttpApiPublishRevisionInfoVipConfigs `json:"vipConfigs,omitempty" xml:"vipConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfo) GetBackendScene() *string {
	return s.BackendScene
}

func (s *HttpApiPublishRevisionInfo) GetBackendType() *string {
	return s.BackendType
}

func (s *HttpApiPublishRevisionInfo) GetCloudProductConfig() *HttpApiPublishRevisionInfoCloudProductConfig {
	return s.CloudProductConfig
}

func (s *HttpApiPublishRevisionInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *HttpApiPublishRevisionInfo) GetCustomDomains() []*HttpApiDomainInfo {
	return s.CustomDomains
}

func (s *HttpApiPublishRevisionInfo) GetDnsConfigs() []*HttpApiPublishRevisionInfoDnsConfigs {
	return s.DnsConfigs
}

func (s *HttpApiPublishRevisionInfo) GetEnvironmentInfo() *HttpApiPublishRevisionInfoEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *HttpApiPublishRevisionInfo) GetIsCurrentVersion() *bool {
	return s.IsCurrentVersion
}

func (s *HttpApiPublishRevisionInfo) GetOperations() []*HttpApiOperationInfo {
	return s.Operations
}

func (s *HttpApiPublishRevisionInfo) GetRevisionId() *string {
	return s.RevisionId
}

func (s *HttpApiPublishRevisionInfo) GetServiceConfigs() []*HttpApiPublishRevisionInfoServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiPublishRevisionInfo) GetSubDomains() []*HttpApiDomainInfo {
	return s.SubDomains
}

func (s *HttpApiPublishRevisionInfo) GetVipConfigs() []*HttpApiPublishRevisionInfoVipConfigs {
	return s.VipConfigs
}

func (s *HttpApiPublishRevisionInfo) SetBackendScene(v string) *HttpApiPublishRevisionInfo {
	s.BackendScene = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetBackendType(v string) *HttpApiPublishRevisionInfo {
	s.BackendType = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCloudProductConfig(v *HttpApiPublishRevisionInfoCloudProductConfig) *HttpApiPublishRevisionInfo {
	s.CloudProductConfig = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCreateTimestamp(v int64) *HttpApiPublishRevisionInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.CustomDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetDnsConfigs(v []*HttpApiPublishRevisionInfoDnsConfigs) *HttpApiPublishRevisionInfo {
	s.DnsConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetEnvironmentInfo(v *HttpApiPublishRevisionInfoEnvironmentInfo) *HttpApiPublishRevisionInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetIsCurrentVersion(v bool) *HttpApiPublishRevisionInfo {
	s.IsCurrentVersion = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetOperations(v []*HttpApiOperationInfo) *HttpApiPublishRevisionInfo {
	s.Operations = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetRevisionId(v string) *HttpApiPublishRevisionInfo {
	s.RevisionId = &v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetServiceConfigs(v []*HttpApiPublishRevisionInfoServiceConfigs) *HttpApiPublishRevisionInfo {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetSubDomains(v []*HttpApiDomainInfo) *HttpApiPublishRevisionInfo {
	s.SubDomains = v
	return s
}

func (s *HttpApiPublishRevisionInfo) SetVipConfigs(v []*HttpApiPublishRevisionInfoVipConfigs) *HttpApiPublishRevisionInfo {
	s.VipConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoCloudProductConfig struct {
	// example:
	//
	// FC
	CloudProductType        *string                                                                `json:"cloudProductType,omitempty" xml:"cloudProductType,omitempty"`
	ContainerServiceConfigs []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs `json:"containerServiceConfigs,omitempty" xml:"containerServiceConfigs,omitempty" type:"Repeated"`
	FunctionConfigs         []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs         `json:"functionConfigs,omitempty" xml:"functionConfigs,omitempty" type:"Repeated"`
	MseNacosConfigs         []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs         `json:"mseNacosConfigs,omitempty" xml:"mseNacosConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetCloudProductType() *string {
	return s.CloudProductType
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetContainerServiceConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	return s.ContainerServiceConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetFunctionConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	return s.FunctionConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) GetMseNacosConfigs() []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	return s.MseNacosConfigs
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetCloudProductType(v string) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.CloudProductType = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetContainerServiceConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.ContainerServiceConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetFunctionConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.FunctionConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) SetMseNacosConfigs(v []*HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) *HttpApiPublishRevisionInfoCloudProductConfig {
	s.MseNacosConfigs = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 100
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetNamespace() *string {
	return s.Namespace
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) GetWeight() *string {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) SetWeight(v string) *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigContainerServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// demo-function
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetQualifier() *string {
	return s.Qualifier
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetQualifier(v string) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Qualifier = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigFunctionConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string                        `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// spring-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetGroupName() *string {
	return s.GroupName
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetNamespace() *string {
	return s.Namespace
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetGroupName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.GroupName = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetName(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetNamespace(v string) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Namespace = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoCloudProductConfigMseNacosConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoDnsConfigs struct {
	DnsList []*string                      `json:"dnsList,omitempty" xml:"dnsList,omitempty" type:"Repeated"`
	Match   *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoDnsConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoDnsConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetDnsList() []*string {
	return s.DnsList
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetDnsList(v []*string) *HttpApiPublishRevisionInfoDnsConfigs {
	s.DnsList = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoDnsConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoDnsConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoEnvironmentInfo struct {
	// example:
	//
	// 测试
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// env-xxx
	EnvironmentId *string                                               `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayInfo   *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetAlias() *string {
	return s.Alias
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetGatewayInfo() *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	return s.GatewayInfo
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetAlias(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Alias = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetEnvironmentId(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetGatewayInfo(v *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfo {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo struct {
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// 实例1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetGatewayId(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) SetName(v string) *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo {
	s.Name = &v
	return s
}

func (s *HttpApiPublishRevisionInfoEnvironmentInfoGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoServiceConfigs struct {
	// example:
	//
	// gs-xxx
	GatewayServiceId *string                        `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	Match            *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetGatewayServiceId(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetPort(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetProtocol(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetVersion(v string) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPublishRevisionInfoVipConfigs struct {
	Endpoints []*string                      `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiPublishRevisionInfoVipConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPublishRevisionInfoVipConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiPublishRevisionInfoVipConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetEndpoints(v []*string) *HttpApiPublishRevisionInfoVipConfigs {
	s.Endpoints = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiPublishRevisionInfoVipConfigs {
	s.Match = v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) SetWeight(v int32) *HttpApiPublishRevisionInfoVipConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiPublishRevisionInfoVipConfigs) Validate() error {
	return dara.Validate(s)
}
