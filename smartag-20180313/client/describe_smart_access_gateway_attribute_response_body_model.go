// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateBandwidth(v int64) *DescribeSmartAccessGatewayAttributeResponseBody
	GetAccelerateBandwidth() *int64
	SetAccessPointId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetAccessPointId() *string
	SetAclIds(v *DescribeSmartAccessGatewayAttributeResponseBodyAclIds) *DescribeSmartAccessGatewayAttributeResponseBody
	GetAclIds() *DescribeSmartAccessGatewayAttributeResponseBodyAclIds
	SetApplicationBandwidthPackageBussinessStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetApplicationBandwidthPackageBussinessStatus() *string
	SetApplicationBandwidthPackageId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetApplicationBandwidthPackageId() *string
	SetApplicationBandwidthPackageName(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetApplicationBandwidthPackageName() *string
	SetApplicationBandwidthPackageOperationLocks(v *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) *DescribeSmartAccessGatewayAttributeResponseBody
	GetApplicationBandwidthPackageOperationLocks() *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks
	SetAssociatedCcnId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetAssociatedCcnId() *string
	SetAssociatedCcnName(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetAssociatedCcnName() *string
	SetBackupBoxControllerIp(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetBackupBoxControllerIp() *string
	SetBoxControllerIp(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetBoxControllerIp() *string
	SetCidrBlock(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetCidrBlock() *string
	SetCity(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetCity() *string
	SetCreateTime(v int64) *DescribeSmartAccessGatewayAttributeResponseBody
	GetCreateTime() *int64
	SetDataPlan(v int64) *DescribeSmartAccessGatewayAttributeResponseBody
	GetDataPlan() *int64
	SetDescription(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetDescription() *string
	SetDevices(v *DescribeSmartAccessGatewayAttributeResponseBodyDevices) *DescribeSmartAccessGatewayAttributeResponseBody
	GetDevices() *DescribeSmartAccessGatewayAttributeResponseBodyDevices
	SetEnableOptimization(v bool) *DescribeSmartAccessGatewayAttributeResponseBody
	GetEnableOptimization() *bool
	SetEnableSoftwareConnectionAudit(v bool) *DescribeSmartAccessGatewayAttributeResponseBody
	GetEnableSoftwareConnectionAudit() *bool
	SetEndTime(v int64) *DescribeSmartAccessGatewayAttributeResponseBody
	GetEndTime() *int64
	SetFlowLogIds(v *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) *DescribeSmartAccessGatewayAttributeResponseBody
	GetFlowLogIds() *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds
	SetIRIds(v *DescribeSmartAccessGatewayAttributeResponseBodyIRIds) *DescribeSmartAccessGatewayAttributeResponseBody
	GetIRIds() *DescribeSmartAccessGatewayAttributeResponseBodyIRIds
	SetInstanceType(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetInstanceType() *string
	SetIpsecStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetIpsecStatus() *string
	SetLinks(v *DescribeSmartAccessGatewayAttributeResponseBodyLinks) *DescribeSmartAccessGatewayAttributeResponseBody
	GetLinks() *DescribeSmartAccessGatewayAttributeResponseBodyLinks
	SetMaxBandwidth(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetMaxBandwidth() *string
	SetName(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetName() *string
	SetOptimizationType(v bool) *DescribeSmartAccessGatewayAttributeResponseBody
	GetOptimizationType() *bool
	SetPosition(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetPosition() *string
	SetQosIds(v *DescribeSmartAccessGatewayAttributeResponseBodyQosIds) *DescribeSmartAccessGatewayAttributeResponseBody
	GetQosIds() *DescribeSmartAccessGatewayAttributeResponseBodyQosIds
	SetRequestId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetRequestId() *string
	SetResellerInstanceId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetResellerInstanceId() *string
	SetResellerUid(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetResellerUid() *string
	SetResourceGroupId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetRoutingStrategy(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetRoutingStrategy() *string
	SetSecurityLockThreshold(v int32) *DescribeSmartAccessGatewayAttributeResponseBody
	GetSecurityLockThreshold() *int32
	SetSerialNumber(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetSerialNumber() *string
	SetSmartAGId(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetSmartAGId() *string
	SetStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetStatus() *string
	SetTrafficMasterSn(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetTrafficMasterSn() *string
	SetUpBandwidth4G(v int32) *DescribeSmartAccessGatewayAttributeResponseBody
	GetUpBandwidth4G() *int32
	SetUpBandwidthWan(v int32) *DescribeSmartAccessGatewayAttributeResponseBody
	GetUpBandwidthWan() *int32
	SetUserCount(v int32) *DescribeSmartAccessGatewayAttributeResponseBody
	GetUserCount() *int32
	SetVpnStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody
	GetVpnStatus() *string
}

type DescribeSmartAccessGatewayAttributeResponseBody struct {
	// The maximum bandwidth value for application acceleration. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	AccelerateBandwidth *int64 `json:"AccelerateBandwidth,omitempty" xml:"AccelerateBandwidth,omitempty"`
	// The ID of the access point for the SAG instance.
	//
	// example:
	//
	// 238
	AccessPointId *string                                                `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AclIds        *DescribeSmartAccessGatewayAttributeResponseBodyAclIds `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Struct"`
	// The status of the bandwidth plan for application acceleration. Valid value:
	//
	// 	- **Abnormal**: abnormal
	//
	// 	- **Normal**: normal
	//
	// example:
	//
	// Normal
	ApplicationBandwidthPackageBussinessStatus *string `json:"ApplicationBandwidthPackageBussinessStatus,omitempty" xml:"ApplicationBandwidthPackageBussinessStatus,omitempty"`
	// The ID of the bandwidth plan for application acceleration that is associated with the SAG instance.
	//
	// example:
	//
	// abwp-7963l7iqnquyj3****
	ApplicationBandwidthPackageId *string `json:"ApplicationBandwidthPackageId,omitempty" xml:"ApplicationBandwidthPackageId,omitempty"`
	// The name of the bandwidth plan for application acceleration that is associated with the SAG instance.
	//
	// example:
	//
	// testname
	ApplicationBandwidthPackageName *string `json:"ApplicationBandwidthPackageName,omitempty" xml:"ApplicationBandwidthPackageName,omitempty"`
	// Indicates whether the bandwidth plan is locked.
	ApplicationBandwidthPackageOperationLocks *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks `json:"ApplicationBandwidthPackageOperationLocks,omitempty" xml:"ApplicationBandwidthPackageOperationLocks,omitempty" type:"Struct"`
	// The ID of the Cloud Connect Network (CCN) instance with which the SAG instance is associated.
	//
	// example:
	//
	// ccn-iz26o9zye6lhoo****
	AssociatedCcnId *string `json:"AssociatedCcnId,omitempty" xml:"AssociatedCcnId,omitempty"`
	// The ID of the Cloud Connect Network (CCN) instance with which the SAG instance is associated.
	//
	// example:
	//
	// testname
	AssociatedCcnName *string `json:"AssociatedCcnName,omitempty" xml:"AssociatedCcnName,omitempty"`
	// The public IP address of the standby SAG device.
	//
	// example:
	//
	// 112.XX.XX.27
	BackupBoxControllerIp *string `json:"BackupBoxControllerIp,omitempty" xml:"BackupBoxControllerIp,omitempty"`
	// The public IP address of the active SAG device.
	//
	// example:
	//
	// 112.XX.XX.25
	BoxControllerIp *string `json:"BoxControllerIp,omitempty" xml:"BoxControllerIp,omitempty"`
	// The private CIDR block of the destination network with which the on-premises network or client needs to communicate.
	//
	// example:
	//
	// 10.0.9.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The ID of the city where the SAG device is deployed.
	//
	// example:
	//
	// cn-shanghai
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The timestamp when the SAG instance was created.
	//
	// example:
	//
	// 1622617250000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data transfer plan of the SAG instance. Unit: GB.
	//
	// >  Each client account has a data transfer plan free of charge for 5 GB each month.
	//
	// example:
	//
	// 5
	DataPlan *int64 `json:"DataPlan,omitempty" xml:"DataPlan,omitempty"`
	// The description of the SAG instance.
	//
	// example:
	//
	// testdesc
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Devices     *DescribeSmartAccessGatewayAttributeResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Struct"`
	// Indicates whether the transmission optimization feature is enabled.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	EnableOptimization *bool `json:"EnableOptimization,omitempty" xml:"EnableOptimization,omitempty"`
	// Indicates whether the audit log for connections to the SAG app instance is enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	EnableSoftwareConnectionAudit *bool `json:"EnableSoftwareConnectionAudit,omitempty" xml:"EnableSoftwareConnectionAudit,omitempty"`
	// The timestamp when the SAG instance expires.
	//
	// example:
	//
	// 1628265600000
	EndTime    *int64                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowLogIds *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds `json:"FlowLogIds,omitempty" xml:"FlowLogIds,omitempty" type:"Struct"`
	IRIds      *DescribeSmartAccessGatewayAttributeResponseBodyIRIds      `json:"IRIds,omitempty" xml:"IRIds,omitempty" type:"Struct"`
	// The type of the SAG instance. Valid values:
	//
	// 	- **sag-1000**: indicates an SAG CPE instance and the instance is associated with an SAG-1000 device.
	//
	// 	- **sag-10wm**: indicates an SAG CPE instance and the instance is associated with an SAG-100WM device.
	//
	// 	- **sag-software**: indicates an SAG app instance.
	//
	// 	- **sag-vcpe**: an SAG vCPE instance.
	//
	// example:
	//
	// sag-vcpe
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The status of the IPsec-VPN connection. Valid values:
	//
	// 	- **up**: normal
	//
	// 	- **down**: abnormal
	//
	// example:
	//
	// up
	IpsecStatus *string                                               `json:"IpsecStatus,omitempty" xml:"IpsecStatus,omitempty"`
	Links       *DescribeSmartAccessGatewayAttributeResponseBodyLinks `json:"Links,omitempty" xml:"Links,omitempty" type:"Struct"`
	// The maximum bandwidth value of the SAG instance. Unit: Mbit/s.
	//
	// example:
	//
	// 50 M
	MaxBandwidth *string `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The name of the SAG instance.
	//
	// example:
	//
	// testname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The transmission optimization type of the SAG instance. If transmission optimization is enabled, the default value is **fec**.
	//
	// example:
	//
	// fec
	OptimizationType *bool `json:"OptimizationType,omitempty" xml:"OptimizationType,omitempty"`
	// The location of the SAG instance.
	Position *string                                                `json:"Position,omitempty" xml:"Position,omitempty"`
	QosIds   *DescribeSmartAccessGatewayAttributeResponseBodyQosIds `json:"QosIds,omitempty" xml:"QosIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F02D092B-A0B7-4BA1-BCA7-014B953C5DC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the instance provided by the third-party reseller.
	//
	// example:
	//
	// sag-v0fkpk4akfz5******
	ResellerInstanceId *string `json:"ResellerInstanceId,omitempty" xml:"ResellerInstanceId,omitempty"`
	// The ID of the third-party reseller.
	//
	// example:
	//
	// 1210123456123456
	ResellerUid *string `json:"ResellerUid,omitempty" xml:"ResellerUid,omitempty"`
	// The ID of the resource group to which the SAG instance belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The method that the SAG instance uses to advertise routes to Alibaba Cloud.
	//
	// 	- **static**: static routing
	//
	// 	- **dynamic**: dynamic routing
	//
	// example:
	//
	// static
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// The time threshold. If the SAG device remains disconnected for the specified period of time, the SAG device is locked.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 3600
	SecurityLockThreshold *int32 `json:"SecurityLockThreshold,omitempty" xml:"SecurityLockThreshold,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sage6gsdllbidl****,sage6nniq3d****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-6z21oj0vjjrx6s****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The status of the SAG instance. Valid values:
	//
	// 	- **Ordered**: The order is to be shipped.
	//
	// 	- **Delivered**: The SAG instance is shipped.
	//
	// 	- **Received**: The SAG instance is activated.
	//
	// 	- **Unconfirmed**: The SAG instance is to be confirmed.
	//
	// 	- **Active**: The SAG instance is available.
	//
	// 	- **Offline**: The SAG instance is disconnected.
	//
	// 	- **Arrearage**: The SAG device is locked due to overdue payments.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The serial number of the active SAG device.
	//
	// example:
	//
	// sage6nniq3d****
	TrafficMasterSn *string `json:"TrafficMasterSn,omitempty" xml:"TrafficMasterSn,omitempty"`
	// The maximum upstream bandwidth of 4G network connections established by the SAG device. Unit: Mbit/s.
	//
	// example:
	//
	// 3
	UpBandwidth4G *int32 `json:"UpBandwidth4G,omitempty" xml:"UpBandwidth4G,omitempty"`
	// The maximum upstream bandwidth of network connections established on the WAN port of the SAG device. Unit: Mbit/s.
	//
	// example:
	//
	// 4
	UpBandwidthWan *int32 `json:"UpBandwidthWan,omitempty" xml:"UpBandwidthWan,omitempty"`
	// The number of client accounts on the SAG instance.
	//
	// example:
	//
	// 3
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The status of the VPN connection. Valid values:
	//
	// 	- **up**: normal
	//
	// 	- **down**: abnormal
	//
	// example:
	//
	// down
	VpnStatus *string `json:"VpnStatus,omitempty" xml:"VpnStatus,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetAccelerateBandwidth() *int64 {
	return s.AccelerateBandwidth
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetAclIds() *DescribeSmartAccessGatewayAttributeResponseBodyAclIds {
	return s.AclIds
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetApplicationBandwidthPackageBussinessStatus() *string {
	return s.ApplicationBandwidthPackageBussinessStatus
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetApplicationBandwidthPackageId() *string {
	return s.ApplicationBandwidthPackageId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetApplicationBandwidthPackageName() *string {
	return s.ApplicationBandwidthPackageName
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetApplicationBandwidthPackageOperationLocks() *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks {
	return s.ApplicationBandwidthPackageOperationLocks
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetAssociatedCcnId() *string {
	return s.AssociatedCcnId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetAssociatedCcnName() *string {
	return s.AssociatedCcnName
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetBackupBoxControllerIp() *string {
	return s.BackupBoxControllerIp
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetBoxControllerIp() *string {
	return s.BoxControllerIp
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetCity() *string {
	return s.City
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetDataPlan() *int64 {
	return s.DataPlan
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetDevices() *DescribeSmartAccessGatewayAttributeResponseBodyDevices {
	return s.Devices
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetEnableOptimization() *bool {
	return s.EnableOptimization
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetEnableSoftwareConnectionAudit() *bool {
	return s.EnableSoftwareConnectionAudit
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetFlowLogIds() *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds {
	return s.FlowLogIds
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetIRIds() *DescribeSmartAccessGatewayAttributeResponseBodyIRIds {
	return s.IRIds
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetIpsecStatus() *string {
	return s.IpsecStatus
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetLinks() *DescribeSmartAccessGatewayAttributeResponseBodyLinks {
	return s.Links
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetMaxBandwidth() *string {
	return s.MaxBandwidth
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetOptimizationType() *bool {
	return s.OptimizationType
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetPosition() *string {
	return s.Position
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetQosIds() *DescribeSmartAccessGatewayAttributeResponseBodyQosIds {
	return s.QosIds
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetResellerInstanceId() *string {
	return s.ResellerInstanceId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetResellerUid() *string {
	return s.ResellerUid
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetSecurityLockThreshold() *int32 {
	return s.SecurityLockThreshold
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetTrafficMasterSn() *string {
	return s.TrafficMasterSn
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetUpBandwidth4G() *int32 {
	return s.UpBandwidth4G
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetUpBandwidthWan() *int32 {
	return s.UpBandwidthWan
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) GetVpnStatus() *string {
	return s.VpnStatus
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetAccelerateBandwidth(v int64) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.AccelerateBandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetAccessPointId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.AccessPointId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetAclIds(v *DescribeSmartAccessGatewayAttributeResponseBodyAclIds) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.AclIds = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetApplicationBandwidthPackageBussinessStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ApplicationBandwidthPackageBussinessStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetApplicationBandwidthPackageId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ApplicationBandwidthPackageId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetApplicationBandwidthPackageName(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ApplicationBandwidthPackageName = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetApplicationBandwidthPackageOperationLocks(v *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ApplicationBandwidthPackageOperationLocks = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetAssociatedCcnId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.AssociatedCcnId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetAssociatedCcnName(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.AssociatedCcnName = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetBackupBoxControllerIp(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.BackupBoxControllerIp = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetBoxControllerIp(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.BoxControllerIp = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetCidrBlock(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetCity(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.City = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetCreateTime(v int64) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetDataPlan(v int64) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.DataPlan = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetDescription(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetDevices(v *DescribeSmartAccessGatewayAttributeResponseBodyDevices) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetEnableOptimization(v bool) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.EnableOptimization = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetEnableSoftwareConnectionAudit(v bool) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.EnableSoftwareConnectionAudit = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetEndTime(v int64) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetFlowLogIds(v *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.FlowLogIds = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetIRIds(v *DescribeSmartAccessGatewayAttributeResponseBodyIRIds) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.IRIds = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetInstanceType(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetIpsecStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.IpsecStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetLinks(v *DescribeSmartAccessGatewayAttributeResponseBodyLinks) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Links = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetMaxBandwidth(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.MaxBandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetName(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetOptimizationType(v bool) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.OptimizationType = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetPosition(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Position = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetQosIds(v *DescribeSmartAccessGatewayAttributeResponseBodyQosIds) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.QosIds = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetRequestId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetResellerInstanceId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ResellerInstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetResellerUid(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ResellerUid = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetResourceGroupId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetRoutingStrategy(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.RoutingStrategy = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetSecurityLockThreshold(v int32) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.SecurityLockThreshold = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetSerialNumber(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetSmartAGId(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetTrafficMasterSn(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.TrafficMasterSn = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetUpBandwidth4G(v int32) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.UpBandwidth4G = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetUpBandwidthWan(v int32) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.UpBandwidthWan = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetUserCount(v int32) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.UserCount = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) SetVpnStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBody {
	s.VpnStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBody) Validate() error {
	if s.AclIds != nil {
		if err := s.AclIds.Validate(); err != nil {
			return err
		}
	}
	if s.ApplicationBandwidthPackageOperationLocks != nil {
		if err := s.ApplicationBandwidthPackageOperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.Devices != nil {
		if err := s.Devices.Validate(); err != nil {
			return err
		}
	}
	if s.FlowLogIds != nil {
		if err := s.FlowLogIds.Validate(); err != nil {
			return err
		}
	}
	if s.IRIds != nil {
		if err := s.IRIds.Validate(); err != nil {
			return err
		}
	}
	if s.Links != nil {
		if err := s.Links.Validate(); err != nil {
			return err
		}
	}
	if s.QosIds != nil {
		if err := s.QosIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewayAttributeResponseBodyAclIds struct {
	AclId []*string `json:"AclId,omitempty" xml:"AclId,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyAclIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyAclIds) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyAclIds) GetAclId() []*string {
	return s.AclId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyAclIds) SetAclId(v []*string) *DescribeSmartAccessGatewayAttributeResponseBodyAclIds {
	s.AclId = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyAclIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks struct {
	// The reason why the instance was locked.
	//
	// example:
	//
	// Message
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The lock mode of the instance. The value is set to **FinancialLocked**.
	//
	// example:
	//
	// FinancialLocked
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) GetLockType() *string {
	return s.LockType
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) SetLockReason(v string) *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks {
	s.LockReason = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) SetLockType(v string) *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks {
	s.LockType = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyApplicationBandwidthPackageOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyDevices struct {
	Device []*DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevices) GetDevice() []*DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	return s.Device
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevices) SetDevice(v []*DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) *DescribeSmartAccessGatewayAttributeResponseBodyDevices {
	s.Device = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevices) Validate() error {
	if s.Device != nil {
		for _, item := range s.Device {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice struct {
	DistributeSkStatus    *int32  `json:"DistributeSkStatus,omitempty" xml:"DistributeSkStatus,omitempty"`
	DpiSignatureDbVersion *string `json:"DpiSignatureDbVersion,omitempty" xml:"DpiSignatureDbVersion,omitempty"`
	HaState               *string `json:"HaState,omitempty" xml:"HaState,omitempty"`
	HcState               *string `json:"HcState,omitempty" xml:"HcState,omitempty"`
	MonitorVersion        *string `json:"MonitorVersion,omitempty" xml:"MonitorVersion,omitempty"`
	SecretKey             *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	SerialNumber          *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SoftwareVersion       *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetDistributeSkStatus() *int32 {
	return s.DistributeSkStatus
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetDpiSignatureDbVersion() *string {
	return s.DpiSignatureDbVersion
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetHaState() *string {
	return s.HaState
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetHcState() *string {
	return s.HcState
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetMonitorVersion() *string {
	return s.MonitorVersion
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetDistributeSkStatus(v int32) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.DistributeSkStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetDpiSignatureDbVersion(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.DpiSignatureDbVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetHaState(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.HaState = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetHcState(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.HcState = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetMonitorVersion(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.MonitorVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetSecretKey(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.SecretKey = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetSerialNumber(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) SetSoftwareVersion(v string) *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice {
	s.SoftwareVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyDevicesDevice) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds struct {
	FlowLogId []*string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) GetFlowLogId() []*string {
	return s.FlowLogId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) SetFlowLogId(v []*string) *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds {
	s.FlowLogId = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyFlowLogIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyIRIds struct {
	IRId []*string `json:"IRId,omitempty" xml:"IRId,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyIRIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyIRIds) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyIRIds) GetIRId() []*string {
	return s.IRId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyIRIds) SetIRId(v []*string) *DescribeSmartAccessGatewayAttributeResponseBodyIRIds {
	s.IRId = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyIRIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyLinks struct {
	Link []*DescribeSmartAccessGatewayAttributeResponseBodyLinksLink `json:"Link,omitempty" xml:"Link,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyLinks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyLinks) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinks) GetLink() []*DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	return s.Link
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinks) SetLink(v []*DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) *DescribeSmartAccessGatewayAttributeResponseBodyLinks {
	s.Link = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinks) Validate() error {
	if s.Link != nil {
		for _, item := range s.Link {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewayAttributeResponseBodyLinksLink struct {
	Bandwidth              *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CommodityType          *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	EndTime                *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HealthCheckTargetIp    *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RelateInstanceId       *string `json:"RelateInstanceId,omitempty" xml:"RelateInstanceId,omitempty"`
	RelateInstanceRegionId *string `json:"RelateInstanceRegionId,omitempty" xml:"RelateInstanceRegionId,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetCommodityType() *string {
	return s.CommodityType
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetRelateInstanceId() *string {
	return s.RelateInstanceId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetRelateInstanceRegionId() *string {
	return s.RelateInstanceRegionId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetStatus() *string {
	return s.Status
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) GetType() *string {
	return s.Type
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetBandwidth(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.Bandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetCommodityType(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.CommodityType = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetEndTime(v int64) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.EndTime = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetHealthCheckTargetIp(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetInstanceId(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.InstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetRelateInstanceId(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.RelateInstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetRelateInstanceRegionId(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.RelateInstanceRegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetStatus(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.Status = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) SetType(v string) *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink {
	s.Type = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyLinksLink) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartAccessGatewayAttributeResponseBodyQosIds struct {
	QosId []*string `json:"QosId,omitempty" xml:"QosId,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyQosIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeResponseBodyQosIds) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyQosIds) GetQosId() []*string {
	return s.QosId
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyQosIds) SetQosId(v []*string) *DescribeSmartAccessGatewayAttributeResponseBodyQosIds {
	s.QosId = v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeResponseBodyQosIds) Validate() error {
	return dara.Validate(s)
}
