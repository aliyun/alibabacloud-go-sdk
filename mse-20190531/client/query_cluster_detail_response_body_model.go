// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryClusterDetailResponseBodyData) *QueryClusterDetailResponseBody
	GetData() *QueryClusterDetailResponseBodyData
	SetErrorCode(v string) *QueryClusterDetailResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryClusterDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryClusterDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryClusterDetailResponseBody
	GetSuccess() *bool
}

type QueryClusterDetailResponseBody struct {
	// The details of the data.
	Data *QueryClusterDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9515ACA4-E94D-440D-989E-C379FCED****
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

func (s QueryClusterDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBody) GetData() *QueryClusterDetailResponseBodyData {
	return s.Data
}

func (s *QueryClusterDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryClusterDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryClusterDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryClusterDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryClusterDetailResponseBody) SetData(v *QueryClusterDetailResponseBodyData) *QueryClusterDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterDetailResponseBody) SetErrorCode(v string) *QueryClusterDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetMessage(v string) *QueryClusterDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetRequestId(v string) *QueryClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetSuccess(v bool) *QueryClusterDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryClusterDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryClusterDetailResponseBodyData struct {
	// The whitelist.
	//
	// example:
	//
	// []
	AclEntryList *string `json:"AclEntryList,omitempty" xml:"AclEntryList,omitempty"`
	// The ID of the whitelist.
	//
	// example:
	//
	// acl-bp17020kiqvzutwwj****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.2.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The billing method, such as subscription or pay-as-you-go.
	//
	// example:
	//
	// Pay-as-you-go
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The alias of the instance.
	//
	// example:
	//
	// mse-7413****
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// mse-bc1a29b0-160230875****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The engine specifications.
	//
	// example:
	//
	// MSE_SC_1_2_200_c
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	// The type of the instance. Valid values: ZooKeeper, Nacos-Ans, and Eureka.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// 1.2.1
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// The network connection type. Valid values:
	//
	// 	- slb
	//
	// 	- eni
	//
	// example:
	//
	// slb
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-07-31 11:36:08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The capacity of the disk. Unit: GB.
	//
	// example:
	//
	// 60
	DiskCapacity *int64 `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	// The type of the disk.
	//
	// example:
	//
	// alicloud-disk-ssd-multi-zone
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The health status of the instance.
	//
	// example:
	//
	// RESTART_SUCCESS
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The amount of time taken to create the instance. Unit: milliseconds.
	//
	// example:
	//
	// 98408
	InitCostTime *int64 `json:"InitCostTime,omitempty" xml:"InitCostTime,omitempty"`
	// The creation status of the instance.
	//
	// example:
	//
	// RESTART_SUCCESS
	InitStatus *string `json:"InitStatus,omitempty" xml:"InitStatus,omitempty"`
	// The number of instance nodes.
	//
	// example:
	//
	// 3
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of instance nodes.
	InstanceModels []*QueryClusterDetailResponseBodyDataInstanceModels `json:"InstanceModels,omitempty" xml:"InstanceModels,omitempty" type:"Repeated"`
	// The public IP address of the instance.
	//
	// example:
	//
	// 47.98.XX.XX
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The public endpoint of the instance.
	//
	// example:
	//
	// mse-7413****-p.eureka.mse.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// The private port number.
	//
	// example:
	//
	// 8761
	InternetPort *string `json:"InternetPort,omitempty" xml:"InternetPort,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// mse-7413****-eureka.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The private port number.
	//
	// example:
	//
	// 8761
	IntranetPort *string `json:"IntranetPort,omitempty" xml:"IntranetPort,omitempty"`
	// The size of the memory. Unit: GB.
	//
	// example:
	//
	// 2
	MemoryCapacity *int64 `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
	// The edition of Microservices Engine (MSE).
	//
	// example:
	//
	// mse_basic
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- `privatenet`: VPC
	//
	// 	- `pubnet`: Internet
	//
	// example:
	//
	// privatenet
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The version number of the original order.
	//
	// example:
	//
	// 1.2.0
	OrderClusterVersion *string `json:"OrderClusterVersion,omitempty" xml:"OrderClusterVersion,omitempty"`
	// The billing method, such as subscription or pay-as-you-go.
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
	// 3
	PubNetworkFlow *string `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2dhgysj*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are attached to the instance.
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-xxx-xxxx
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VersionLifecycle *string `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1hcg467ekqsv0zr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryClusterDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyData) GetAclEntryList() *string {
	return s.AclEntryList
}

func (s *QueryClusterDetailResponseBodyData) GetAclId() *string {
	return s.AclId
}

func (s *QueryClusterDetailResponseBodyData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *QueryClusterDetailResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryClusterDetailResponseBodyData) GetClusterAliasName() *string {
	return s.ClusterAliasName
}

func (s *QueryClusterDetailResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryClusterDetailResponseBodyData) GetClusterSpecification() *string {
	return s.ClusterSpecification
}

func (s *QueryClusterDetailResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *QueryClusterDetailResponseBodyData) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *QueryClusterDetailResponseBodyData) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *QueryClusterDetailResponseBodyData) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryClusterDetailResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryClusterDetailResponseBodyData) GetDiskCapacity() *int64 {
	return s.DiskCapacity
}

func (s *QueryClusterDetailResponseBodyData) GetDiskType() *string {
	return s.DiskType
}

func (s *QueryClusterDetailResponseBodyData) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *QueryClusterDetailResponseBodyData) GetInitCostTime() *int64 {
	return s.InitCostTime
}

func (s *QueryClusterDetailResponseBodyData) GetInitStatus() *string {
	return s.InitStatus
}

func (s *QueryClusterDetailResponseBodyData) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *QueryClusterDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryClusterDetailResponseBodyData) GetInstanceModels() []*QueryClusterDetailResponseBodyDataInstanceModels {
	return s.InstanceModels
}

func (s *QueryClusterDetailResponseBodyData) GetInternetAddress() *string {
	return s.InternetAddress
}

func (s *QueryClusterDetailResponseBodyData) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *QueryClusterDetailResponseBodyData) GetInternetPort() *string {
	return s.InternetPort
}

func (s *QueryClusterDetailResponseBodyData) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *QueryClusterDetailResponseBodyData) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *QueryClusterDetailResponseBodyData) GetIntranetPort() *string {
	return s.IntranetPort
}

func (s *QueryClusterDetailResponseBodyData) GetMemoryCapacity() *int64 {
	return s.MemoryCapacity
}

func (s *QueryClusterDetailResponseBodyData) GetMseVersion() *string {
	return s.MseVersion
}

func (s *QueryClusterDetailResponseBodyData) GetNetType() *string {
	return s.NetType
}

func (s *QueryClusterDetailResponseBodyData) GetOrderClusterVersion() *string {
	return s.OrderClusterVersion
}

func (s *QueryClusterDetailResponseBodyData) GetPayInfo() *string {
	return s.PayInfo
}

func (s *QueryClusterDetailResponseBodyData) GetPubNetworkFlow() *string {
	return s.PubNetworkFlow
}

func (s *QueryClusterDetailResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryClusterDetailResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryClusterDetailResponseBodyData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *QueryClusterDetailResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *QueryClusterDetailResponseBodyData) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *QueryClusterDetailResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryClusterDetailResponseBodyData) SetAclEntryList(v string) *QueryClusterDetailResponseBodyData {
	s.AclEntryList = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetAclId(v string) *QueryClusterDetailResponseBodyData {
	s.AclId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetAppVersion(v string) *QueryClusterDetailResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetChargeType(v string) *QueryClusterDetailResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterAliasName(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterAliasName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterName(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterSpecification(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterSpecification = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterType(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterVersion(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetConnectionType(v string) *QueryClusterDetailResponseBodyData {
	s.ConnectionType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetCpu(v int32) *QueryClusterDetailResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetCreateTime(v string) *QueryClusterDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetDiskCapacity(v int64) *QueryClusterDetailResponseBodyData {
	s.DiskCapacity = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetDiskType(v string) *QueryClusterDetailResponseBodyData {
	s.DiskType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetHealthStatus(v string) *QueryClusterDetailResponseBodyData {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInitCostTime(v int64) *QueryClusterDetailResponseBodyData {
	s.InitCostTime = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInitStatus(v string) *QueryClusterDetailResponseBodyData {
	s.InitStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceCount(v int32) *QueryClusterDetailResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceId(v string) *QueryClusterDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceModels(v []*QueryClusterDetailResponseBodyDataInstanceModels) *QueryClusterDetailResponseBodyData {
	s.InstanceModels = v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetAddress(v string) *QueryClusterDetailResponseBodyData {
	s.InternetAddress = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetDomain(v string) *QueryClusterDetailResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetPort(v string) *QueryClusterDetailResponseBodyData {
	s.InternetPort = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetAddress(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetAddress = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetDomain(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetPort(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetPort = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetMemoryCapacity(v int64) *QueryClusterDetailResponseBodyData {
	s.MemoryCapacity = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetMseVersion(v string) *QueryClusterDetailResponseBodyData {
	s.MseVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetNetType(v string) *QueryClusterDetailResponseBodyData {
	s.NetType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetOrderClusterVersion(v string) *QueryClusterDetailResponseBodyData {
	s.OrderClusterVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetPayInfo(v string) *QueryClusterDetailResponseBodyData {
	s.PayInfo = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetPubNetworkFlow(v string) *QueryClusterDetailResponseBodyData {
	s.PubNetworkFlow = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetRegionId(v string) *QueryClusterDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetResourceGroupId(v string) *QueryClusterDetailResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetTags(v map[string]interface{}) *QueryClusterDetailResponseBodyData {
	s.Tags = v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetVSwitchId(v string) *QueryClusterDetailResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetVersionLifecycle(v string) *QueryClusterDetailResponseBodyData {
	s.VersionLifecycle = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetVpcId(v string) *QueryClusterDetailResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) Validate() error {
	if s.InstanceModels != nil {
		for _, item := range s.InstanceModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryClusterDetailResponseBodyDataInstanceModels struct {
	// The timestamp when the instance was created.
	//
	// example:
	//
	// 1578575377732
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The health status of the instance.
	//
	// example:
	//
	// Running
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.98.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.12.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// mse-7413****-159616656****-reg-center-0-0
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The role.
	//
	// example:
	//
	// Peer
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The single-thread IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SingleTunnelVip *string `json:"SingleTunnelVip,omitempty" xml:"SingleTunnelVip,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shanghai-f
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s QueryClusterDetailResponseBodyDataInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDetailResponseBodyDataInstanceModels) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetInternetIp() *string {
	return s.InternetIp
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetIp() *string {
	return s.Ip
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetPodName() *string {
	return s.PodName
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetRole() *string {
	return s.Role
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetSingleTunnelVip() *string {
	return s.SingleTunnelVip
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) GetZone() *string {
	return s.Zone
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetCreationTimestamp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.CreationTimestamp = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetHealthStatus(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetInternetIp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.InternetIp = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetIp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.Ip = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetPodName(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.PodName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetRole(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.Role = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetSingleTunnelVip(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.SingleTunnelVip = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetZone(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.Zone = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) Validate() error {
	return dara.Validate(s)
}
