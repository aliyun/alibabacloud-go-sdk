// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryClusterInfoResponseBodyData) *QueryClusterInfoResponseBody
	GetData() *QueryClusterInfoResponseBodyData
	SetErrorCode(v string) *QueryClusterInfoResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryClusterInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryClusterInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryClusterInfoResponseBody
	GetSuccess() *bool
}

type QueryClusterInfoResponseBody struct {
	// The details of the data.
	Data *QueryClusterInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 415088B3-A7BE-56F6-9CD9-C42DE895CD41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoResponseBody) GetData() *QueryClusterInfoResponseBodyData {
	return s.Data
}

func (s *QueryClusterInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryClusterInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryClusterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryClusterInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryClusterInfoResponseBody) SetData(v *QueryClusterInfoResponseBodyData) *QueryClusterInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterInfoResponseBody) SetErrorCode(v string) *QueryClusterInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterInfoResponseBody) SetMessage(v string) *QueryClusterInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterInfoResponseBody) SetRequestId(v string) *QueryClusterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterInfoResponseBody) SetSuccess(v bool) *QueryClusterInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryClusterInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryClusterInfoResponseBodyData struct {
	// The public IP address whitelist.
	//
	// example:
	//
	// ["127.0.0.0/32"]
	AclEntryList *string `json:"AclEntryList,omitempty" xml:"AclEntryList,omitempty"`
	// The ID of the instance in the public IP address whitelist.
	//
	// example:
	//
	// acl-bp144q24cgqvzckmxxxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// 2.1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CanUpdate  *bool   `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	// The billing method, such as subscription or pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The alias of the instance.
	//
	// example:
	//
	// Development environment
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	// The full name of the instance.
	//
	// example:
	//
	// mse-74355150-xxxxxxx
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The engine specifications.
	//
	// example:
	//
	// MSE_SC_2_4_60_c
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	// The type of the instance. Valid values: ZooKeeper, Nacos-Ans, and Eureka.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The version of the order.
	//
	// example:
	//
	// NACOS_2_0_0
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// null
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// null
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-12-15 10:02:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// null
	DiskCapacity *int64 `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// null
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the instance that is associated with the Elastic IP Address (EIP).
	//
	// example:
	//
	// eip-bp1uujshd5funmyy8rcl9
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The time when the subscription instance expires.
	//
	// example:
	//
	// 2021-08-01 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The zones to which the current cluster can be distributed.
	ExpectZones []*string `json:"ExpectZones,omitempty" xml:"ExpectZones,omitempty" type:"Repeated"`
	// The status of the instance.
	//
	// example:
	//
	// INIT_SUCCESS
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The time that is required to initialize the instance. Unit: milliseconds.
	//
	// example:
	//
	// 53353
	InitCostTime *int64 `json:"InitCostTime,omitempty" xml:"InitCostTime,omitempty"`
	// The initial status of the instance.
	//
	// example:
	//
	// INIT_SUCCESS
	InitStatus *string `json:"InitStatus,omitempty" xml:"InitStatus,omitempty"`
	// The number of instance nodes.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reserved structure.
	InstanceModels []*QueryClusterInfoResponseBodyDataInstanceModels `json:"InstanceModels,omitempty" xml:"InstanceModels,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// mse-xxxxxx-p.nacos-ans.mse.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// The instance ports that are accessible over the Internet.
	//
	// example:
	//
	// 8848,6443,9848,8761
	InternetPort *string `json:"InternetPort,omitempty" xml:"InternetPort,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// mse-xxxxx-nacos-ans.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The instance ports that are accessible over an internal network.
	//
	// example:
	//
	// 8848,6443,9848,8761
	IntranetPort *string `json:"IntranetPort,omitempty" xml:"IntranetPort,omitempty"`
	// The O\\&M time window.
	MaintenancePeriod *QueryClusterInfoResponseBodyDataMaintenancePeriod `json:"MaintenancePeriod,omitempty" xml:"MaintenancePeriod,omitempty" type:"Struct"`
	// A deprecated parameter.
	//
	// example:
	//
	// null
	MemoryCapacity *int64 `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The network connection type of the instance.
	//
	// example:
	//
	// privatenet
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The billing method.
	//
	// example:
	//
	// Pay-as-you-go
	PayInfo *string `json:"PayInfo,omitempty" xml:"PayInfo,omitempty"`
	// The public bandwidth. Unit: Mbit/s.\\
	//
	// Valid values: 0 to 5000. The value 0 indicates no access to the Internet.
	//
	// example:
	//
	// 1
	PubNetworkFlow *string `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the elastic network interface (ENI) is connected.
	//
	// example:
	//
	// sg-uf6hgwe067prhg68agfa
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The type of the security group to which the ENI is connected.
	//
	// example:
	//
	// enterprise
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// The tag.
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1egfakxxxxx
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionLifecycle *string `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// The ID of the VPC where the instance resides.
	//
	// example:
	//
	// vpc-bp1v5nbauzh8xxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryClusterInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoResponseBodyData) GetAclEntryList() *string {
	return s.AclEntryList
}

func (s *QueryClusterInfoResponseBodyData) GetAclId() *string {
	return s.AclId
}

func (s *QueryClusterInfoResponseBodyData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *QueryClusterInfoResponseBodyData) GetCanUpdate() *bool {
	return s.CanUpdate
}

func (s *QueryClusterInfoResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryClusterInfoResponseBodyData) GetClusterAliasName() *string {
	return s.ClusterAliasName
}

func (s *QueryClusterInfoResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryClusterInfoResponseBodyData) GetClusterSpecification() *string {
	return s.ClusterSpecification
}

func (s *QueryClusterInfoResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *QueryClusterInfoResponseBodyData) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *QueryClusterInfoResponseBodyData) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *QueryClusterInfoResponseBodyData) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryClusterInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryClusterInfoResponseBodyData) GetDiskCapacity() *int64 {
	return s.DiskCapacity
}

func (s *QueryClusterInfoResponseBodyData) GetDiskType() *string {
	return s.DiskType
}

func (s *QueryClusterInfoResponseBodyData) GetEipInstanceId() *string {
	return s.EipInstanceId
}

func (s *QueryClusterInfoResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryClusterInfoResponseBodyData) GetExpectZones() []*string {
	return s.ExpectZones
}

func (s *QueryClusterInfoResponseBodyData) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *QueryClusterInfoResponseBodyData) GetInitCostTime() *int64 {
	return s.InitCostTime
}

func (s *QueryClusterInfoResponseBodyData) GetInitStatus() *string {
	return s.InitStatus
}

func (s *QueryClusterInfoResponseBodyData) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *QueryClusterInfoResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryClusterInfoResponseBodyData) GetInstanceModels() []*QueryClusterInfoResponseBodyDataInstanceModels {
	return s.InstanceModels
}

func (s *QueryClusterInfoResponseBodyData) GetInternetAddress() *string {
	return s.InternetAddress
}

func (s *QueryClusterInfoResponseBodyData) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *QueryClusterInfoResponseBodyData) GetInternetPort() *string {
	return s.InternetPort
}

func (s *QueryClusterInfoResponseBodyData) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *QueryClusterInfoResponseBodyData) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *QueryClusterInfoResponseBodyData) GetIntranetPort() *string {
	return s.IntranetPort
}

func (s *QueryClusterInfoResponseBodyData) GetMaintenancePeriod() *QueryClusterInfoResponseBodyDataMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *QueryClusterInfoResponseBodyData) GetMemoryCapacity() *int64 {
	return s.MemoryCapacity
}

func (s *QueryClusterInfoResponseBodyData) GetMseVersion() *string {
	return s.MseVersion
}

func (s *QueryClusterInfoResponseBodyData) GetNetType() *string {
	return s.NetType
}

func (s *QueryClusterInfoResponseBodyData) GetPayInfo() *string {
	return s.PayInfo
}

func (s *QueryClusterInfoResponseBodyData) GetPubNetworkFlow() *string {
	return s.PubNetworkFlow
}

func (s *QueryClusterInfoResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryClusterInfoResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *QueryClusterInfoResponseBodyData) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *QueryClusterInfoResponseBodyData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *QueryClusterInfoResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *QueryClusterInfoResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *QueryClusterInfoResponseBodyData) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *QueryClusterInfoResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryClusterInfoResponseBodyData) SetAclEntryList(v string) *QueryClusterInfoResponseBodyData {
	s.AclEntryList = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetAclId(v string) *QueryClusterInfoResponseBodyData {
	s.AclId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetAppVersion(v string) *QueryClusterInfoResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetCanUpdate(v bool) *QueryClusterInfoResponseBodyData {
	s.CanUpdate = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetChargeType(v string) *QueryClusterInfoResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetClusterAliasName(v string) *QueryClusterInfoResponseBodyData {
	s.ClusterAliasName = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetClusterName(v string) *QueryClusterInfoResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetClusterSpecification(v string) *QueryClusterInfoResponseBodyData {
	s.ClusterSpecification = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetClusterType(v string) *QueryClusterInfoResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetClusterVersion(v string) *QueryClusterInfoResponseBodyData {
	s.ClusterVersion = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetConnectionType(v string) *QueryClusterInfoResponseBodyData {
	s.ConnectionType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetCpu(v int32) *QueryClusterInfoResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetCreateTime(v string) *QueryClusterInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetDiskCapacity(v int64) *QueryClusterInfoResponseBodyData {
	s.DiskCapacity = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetDiskType(v string) *QueryClusterInfoResponseBodyData {
	s.DiskType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetEipInstanceId(v string) *QueryClusterInfoResponseBodyData {
	s.EipInstanceId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetEndDate(v string) *QueryClusterInfoResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetExpectZones(v []*string) *QueryClusterInfoResponseBodyData {
	s.ExpectZones = v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetHealthStatus(v string) *QueryClusterInfoResponseBodyData {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInitCostTime(v int64) *QueryClusterInfoResponseBodyData {
	s.InitCostTime = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInitStatus(v string) *QueryClusterInfoResponseBodyData {
	s.InitStatus = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInstanceCount(v int32) *QueryClusterInfoResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInstanceId(v string) *QueryClusterInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInstanceModels(v []*QueryClusterInfoResponseBodyDataInstanceModels) *QueryClusterInfoResponseBodyData {
	s.InstanceModels = v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInternetAddress(v string) *QueryClusterInfoResponseBodyData {
	s.InternetAddress = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInternetDomain(v string) *QueryClusterInfoResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetInternetPort(v string) *QueryClusterInfoResponseBodyData {
	s.InternetPort = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetIntranetAddress(v string) *QueryClusterInfoResponseBodyData {
	s.IntranetAddress = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetIntranetDomain(v string) *QueryClusterInfoResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetIntranetPort(v string) *QueryClusterInfoResponseBodyData {
	s.IntranetPort = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetMaintenancePeriod(v *QueryClusterInfoResponseBodyDataMaintenancePeriod) *QueryClusterInfoResponseBodyData {
	s.MaintenancePeriod = v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetMemoryCapacity(v int64) *QueryClusterInfoResponseBodyData {
	s.MemoryCapacity = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetMseVersion(v string) *QueryClusterInfoResponseBodyData {
	s.MseVersion = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetNetType(v string) *QueryClusterInfoResponseBodyData {
	s.NetType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetPayInfo(v string) *QueryClusterInfoResponseBodyData {
	s.PayInfo = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetPubNetworkFlow(v string) *QueryClusterInfoResponseBodyData {
	s.PubNetworkFlow = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetRegionId(v string) *QueryClusterInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetSecurityGroupId(v string) *QueryClusterInfoResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetSecurityGroupType(v string) *QueryClusterInfoResponseBodyData {
	s.SecurityGroupType = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetTags(v map[string]interface{}) *QueryClusterInfoResponseBodyData {
	s.Tags = v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetVSwitchId(v string) *QueryClusterInfoResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetVersionCode(v string) *QueryClusterInfoResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetVersionLifecycle(v string) *QueryClusterInfoResponseBodyData {
	s.VersionLifecycle = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) SetVpcId(v string) *QueryClusterInfoResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *QueryClusterInfoResponseBodyData) Validate() error {
	if s.InstanceModels != nil {
		for _, item := range s.InstanceModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MaintenancePeriod != nil {
		if err := s.MaintenancePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryClusterInfoResponseBodyDataInstanceModels struct {
	// A reserved parameter.
	//
	// example:
	//
	// null
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	SingleTunnelVip *string `json:"SingleTunnelVip,omitempty" xml:"SingleTunnelVip,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s QueryClusterInfoResponseBodyDataInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoResponseBodyDataInstanceModels) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetInternetIp() *string {
	return s.InternetIp
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetIp() *string {
	return s.Ip
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetPodName() *string {
	return s.PodName
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetRole() *string {
	return s.Role
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetSingleTunnelVip() *string {
	return s.SingleTunnelVip
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) GetZone() *string {
	return s.Zone
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetCreationTimestamp(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.CreationTimestamp = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetHealthStatus(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetInternetIp(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.InternetIp = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetIp(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.Ip = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetPodName(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.PodName = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetRole(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.Role = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetSingleTunnelVip(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.SingleTunnelVip = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) SetZone(v string) *QueryClusterInfoResponseBodyDataInstanceModels {
	s.Zone = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataInstanceModels) Validate() error {
	return dara.Validate(s)
}

type QueryClusterInfoResponseBodyDataMaintenancePeriod struct {
	// The start time of the O\\&M time window.
	//
	// example:
	//
	// 02:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end time of the O\\&M time window.
	//
	// example:
	//
	// 06:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryClusterInfoResponseBodyDataMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterInfoResponseBodyDataMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *QueryClusterInfoResponseBodyDataMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryClusterInfoResponseBodyDataMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryClusterInfoResponseBodyDataMaintenancePeriod) SetEndTime(v string) *QueryClusterInfoResponseBodyDataMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataMaintenancePeriod) SetStartTime(v string) *QueryClusterInfoResponseBodyDataMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *QueryClusterInfoResponseBodyDataMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}
