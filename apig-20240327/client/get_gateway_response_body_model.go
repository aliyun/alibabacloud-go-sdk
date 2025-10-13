// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayResponseBody
	GetCode() *string
	SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody
	GetData() *GetGatewayResponseBodyData
	SetMessage(v string) *GetGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayResponseBody
	GetRequestId() *string
}

type GetGatewayResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetGatewayResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F138FFC-6E2B-56C1-9BAB-A67462E339D1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayResponseBody) GetData() *GetGatewayResponseBodyData {
	return s.Data
}

func (s *GetGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayResponseBody) SetCode(v string) *GetGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetMessage(v string) *GetGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResponseBody) SetRequestId(v string) *GetGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyData struct {
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The creation source of the instance. Valid values:
	//
	// 	- Console
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// The creation timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The list of environments associated with the instance.
	Environments []*GetGatewayResponseBodyDataEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The time when the instance expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// example:
	//
	// Serverless
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gw-cq2vundlhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// the gateway type, which is categorized into the following two types:
	//
	// - API: indicates an API gateway
	//
	// - AI: Indicates an AI gateway
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	Isp         *string `json:"isp,omitempty" xml:"isp,omitempty"`
	// The ingress addresses of the instance.
	LoadBalancers     []*GetGatewayResponseBodyDataLoadBalancers   `json:"loadBalancers,omitempty" xml:"loadBalancers,omitempty" type:"Repeated"`
	MaintenancePeriod *GetGatewayResponseBodyDataMaintenancePeriod `json:"maintenancePeriod,omitempty" xml:"maintenancePeriod,omitempty" type:"Struct"`
	// The instance name.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node quantity of the instance.
	//
	// example:
	//
	// 2
	Replicas *string `json:"replicas,omitempty" xml:"replicas,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2s3cvc4jzfxi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The security group of the instance.
	SecurityGroup *GetGatewayResponseBodyDataSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" type:"Struct"`
	// The instance specification. Valid values:
	//
	// 	- apigw.small.x1
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The instance state. Valid values:
	//
	// 	- Running: The instance is running.
	//
	// 	- Creating: The instance is being created.
	//
	// 	- CreateFailed: The instance failed to be created.
	//
	// 	- Upgrading: The instance is being upgraded.
	//
	// 	- UpgradeFailed: The instance failed to be upgraded.
	//
	// 	- Restarting: The instance is being restarted.
	//
	// 	- RestartFailed: The instance failed to be restarted.
	//
	// 	- Deleting: The instance is being released.
	//
	// 	- DeleteFailed: The instance failed to be released.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The resource tags.
	Tags []*GetGatewayResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The destination version of the instance. If the value is inconsistent with the version value, you can upgrade the instance.
	//
	// example:
	//
	// 2.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// The last update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
	// The vSwitch associated with the instance.
	VSwitch *GetGatewayResponseBodyDataVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The instance version.
	//
	// example:
	//
	// 2.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The VPC associated with the instance.
	Vpc *GetGatewayResponseBodyDataVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
	// The list of zones associated with the instance.
	Zones []*GetGatewayResponseBodyDataZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetGatewayResponseBodyData) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *GetGatewayResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetGatewayResponseBodyData) GetEnvironments() []*GetGatewayResponseBodyDataEnvironments {
	return s.Environments
}

func (s *GetGatewayResponseBodyData) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *GetGatewayResponseBodyData) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *GetGatewayResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetGatewayResponseBodyData) GetIsp() *string {
	return s.Isp
}

func (s *GetGatewayResponseBodyData) GetLoadBalancers() []*GetGatewayResponseBodyDataLoadBalancers {
	return s.LoadBalancers
}

func (s *GetGatewayResponseBodyData) GetMaintenancePeriod() *GetGatewayResponseBodyDataMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *GetGatewayResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyData) GetReplicas() *string {
	return s.Replicas
}

func (s *GetGatewayResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetGatewayResponseBodyData) GetSecurityGroup() *GetGatewayResponseBodyDataSecurityGroup {
	return s.SecurityGroup
}

func (s *GetGatewayResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *GetGatewayResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayResponseBodyData) GetTags() []*GetGatewayResponseBodyDataTags {
	return s.Tags
}

func (s *GetGatewayResponseBodyData) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *GetGatewayResponseBodyData) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetGatewayResponseBodyData) GetVSwitch() *GetGatewayResponseBodyDataVSwitch {
	return s.VSwitch
}

func (s *GetGatewayResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetGatewayResponseBodyData) GetVpc() *GetGatewayResponseBodyDataVpc {
	return s.Vpc
}

func (s *GetGatewayResponseBodyData) GetZones() []*GetGatewayResponseBodyDataZones {
	return s.Zones
}

func (s *GetGatewayResponseBodyData) SetChargeType(v string) *GetGatewayResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateFrom(v string) *GetGatewayResponseBodyData {
	s.CreateFrom = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCreateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEnvironments(v []*GetGatewayResponseBodyDataEnvironments) *GetGatewayResponseBodyData {
	s.Environments = v
	return s
}

func (s *GetGatewayResponseBodyData) SetExpireTimestamp(v int64) *GetGatewayResponseBodyData {
	s.ExpireTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayEdition(v string) *GetGatewayResponseBodyData {
	s.GatewayEdition = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayId(v string) *GetGatewayResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayType(v string) *GetGatewayResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetIsp(v string) *GetGatewayResponseBodyData {
	s.Isp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLoadBalancers(v []*GetGatewayResponseBodyDataLoadBalancers) *GetGatewayResponseBodyData {
	s.LoadBalancers = v
	return s
}

func (s *GetGatewayResponseBodyData) SetMaintenancePeriod(v *GetGatewayResponseBodyDataMaintenancePeriod) *GetGatewayResponseBodyData {
	s.MaintenancePeriod = v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetReplicas(v string) *GetGatewayResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetResourceGroupId(v string) *GetGatewayResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSecurityGroup(v *GetGatewayResponseBodyDataSecurityGroup) *GetGatewayResponseBodyData {
	s.SecurityGroup = v
	return s
}

func (s *GetGatewayResponseBodyData) SetSpec(v string) *GetGatewayResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatus(v string) *GetGatewayResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetTags(v []*GetGatewayResponseBodyDataTags) *GetGatewayResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetGatewayResponseBodyData) SetTargetVersion(v string) *GetGatewayResponseBodyData {
	s.TargetVersion = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetUpdateTimestamp(v int64) *GetGatewayResponseBodyData {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVSwitch(v *GetGatewayResponseBodyDataVSwitch) *GetGatewayResponseBodyData {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyData) SetVersion(v string) *GetGatewayResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVpc(v *GetGatewayResponseBodyDataVpc) *GetGatewayResponseBodyData {
	s.Vpc = v
	return s
}

func (s *GetGatewayResponseBodyData) SetZones(v []*GetGatewayResponseBodyDataZones) *GetGatewayResponseBodyData {
	s.Zones = v
	return s
}

func (s *GetGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataEnvironments struct {
	// The environment alias.
	//
	// example:
	//
	// Default environment
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cp9uhudlht***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// default-gw-cp9ugg5***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetGatewayResponseBodyDataEnvironments) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataEnvironments) GetAlias() *string {
	return s.Alias
}

func (s *GetGatewayResponseBodyDataEnvironments) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetGatewayResponseBodyDataEnvironments) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataEnvironments) SetAlias(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Alias = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetEnvironmentId(v string) *GetGatewayResponseBodyDataEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) SetName(v string) *GetGatewayResponseBodyDataEnvironments {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataEnvironments) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataLoadBalancers struct {
	// The load balancer IP address.
	//
	// example:
	//
	// nlb-xoh3pghr***.cn-hangzhou.nlb.aliyuncs.com
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The IP version of the address. Valid values:
	//
	// 	- ipv4
	//
	// 	- ipv6
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"addressIpVersion,omitempty" xml:"addressIpVersion,omitempty"`
	// The load balancer address type. Valid values:
	//
	// 	- Internet
	//
	// 	- Intranet
	//
	// example:
	//
	// Internet
	AddressType *string `json:"addressType,omitempty" xml:"addressType,omitempty"`
	// Indicates whether the address is the default ingress address of the instance.
	//
	// example:
	//
	// true
	GatewayDefault *bool `json:"gatewayDefault,omitempty" xml:"gatewayDefault,omitempty"`
	// The load balancer ID.
	//
	// example:
	//
	// nlb-xoh3pghru7c***
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The mode in which the load balancer is provided. Valid values:
	//
	// 	- Managed: Cloud-native API Gateway manages and provides the load balancer.
	//
	// example:
	//
	// Managed
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The list of listened ports.
	Ports []*GetGatewayResponseBodyDataLoadBalancersPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	// The load balancer status. Valid values:
	//
	// 	- Ready: The load balancer is available.
	//
	// 	- NotCreate: The load balancer is not associated with the instance.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The load balancer type. Valid values:
	//
	// 	- NLB: Network Load Balancer
	//
	// 	- CLB: Classic Load Balancer
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancers) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddress() *string {
	return s.Address
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetAddressType() *string {
	return s.AddressType
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetGatewayDefault() *bool {
	return s.GatewayDefault
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetMode() *string {
	return s.Mode
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetPorts() []*GetGatewayResponseBodyDataLoadBalancersPorts {
	return s.Ports
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayResponseBodyDataLoadBalancers) GetType() *string {
	return s.Type
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddress(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Address = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressIpVersion(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetAddressType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetGatewayDefault(v bool) *GetGatewayResponseBodyDataLoadBalancers {
	s.GatewayDefault = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetLoadBalancerId(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetMode(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Mode = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetPorts(v []*GetGatewayResponseBodyDataLoadBalancersPorts) *GetGatewayResponseBodyDataLoadBalancers {
	s.Ports = v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetStatus(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) SetType(v string) *GetGatewayResponseBodyDataLoadBalancers {
	s.Type = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancers) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataLoadBalancersPorts struct {
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataLoadBalancersPorts) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) GetPort() *int32 {
	return s.Port
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetPort(v int32) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Port = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) SetProtocol(v string) *GetGatewayResponseBodyDataLoadBalancersPorts {
	s.Protocol = &v
	return s
}

func (s *GetGatewayResponseBodyDataLoadBalancersPorts) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataMaintenancePeriod struct {
	// example:
	//
	// 06:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 02:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetGatewayResponseBodyDataMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) SetEndTime(v string) *GetGatewayResponseBodyDataMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) SetStartTime(v string) *GetGatewayResponseBodyDataMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataSecurityGroup struct {
	// The security group name.
	//
	// example:
	//
	// APIG-sg-gw-cq7ke5ll***
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp16tafq9***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s GetGatewayResponseBodyDataSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataSecurityGroup) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataSecurityGroup) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetName(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataSecurityGroup) SetSecurityGroupId(v string) *GetGatewayResponseBodyDataSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyDataSecurityGroup) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetGatewayResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetGatewayResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetGatewayResponseBodyDataTags) SetKey(v string) *GetGatewayResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetGatewayResponseBodyDataTags) SetValue(v string) *GetGatewayResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetGatewayResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataVSwitch struct {
	// The vSwitch name.
	//
	// example:
	//
	// HangzhouVPCvSwitch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataVSwitch) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVSwitch) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetGatewayResponseBodyDataVSwitch) SetName(v string) *GetGatewayResponseBodyDataVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *GetGatewayResponseBodyDataVSwitch) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataVpc struct {
	// The VPC name.
	//
	// example:
	//
	// HangzhouVPC
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1llj52lvj6xc***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetGatewayResponseBodyDataVpc) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataVpc) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataVpc) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetGatewayResponseBodyDataVpc) SetName(v string) *GetGatewayResponseBodyDataVpc {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataVpc) SetVpcId(v string) *GetGatewayResponseBodyDataVpc {
	s.VpcId = &v
	return s
}

func (s *GetGatewayResponseBodyDataVpc) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataZones struct {
	// The zone name.
	//
	// example:
	//
	// HangzhouZoneE
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The vSwitch information.
	VSwitch *GetGatewayResponseBodyDataZonesVSwitch `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetGatewayResponseBodyDataZones) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataZones) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZones) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataZones) GetVSwitch() *GetGatewayResponseBodyDataZonesVSwitch {
	return s.VSwitch
}

func (s *GetGatewayResponseBodyDataZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetGatewayResponseBodyDataZones) SetName(v string) *GetGatewayResponseBodyDataZones {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetVSwitch(v *GetGatewayResponseBodyDataZonesVSwitch) *GetGatewayResponseBodyDataZones {
	s.VSwitch = v
	return s
}

func (s *GetGatewayResponseBodyDataZones) SetZoneId(v string) *GetGatewayResponseBodyDataZones {
	s.ZoneId = &v
	return s
}

func (s *GetGatewayResponseBodyDataZones) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataZonesVSwitch struct {
	// The vSwitch name.
	//
	// example:
	//
	// HangzhouVPCvSwitch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1c7ggkj***
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s GetGatewayResponseBodyDataZonesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataZonesVSwitch) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetName(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) SetVSwitchId(v string) *GetGatewayResponseBodyDataZonesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *GetGatewayResponseBodyDataZonesVSwitch) Validate() error {
	return dara.Validate(s)
}
