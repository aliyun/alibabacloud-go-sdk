// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeInstancesResponseBodyData) *DescribeInstancesResponseBody
	GetData() []*DescribeInstancesResponseBodyData
	SetErrCode(v string) *DescribeInstancesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInstancesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInstancesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstancesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeInstancesResponseBody
	GetTotal() *int32
}

type DescribeInstancesResponseBody struct {
	// The query results.
	Data []*DescribeInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetData() []*DescribeInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeInstancesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInstancesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstancesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstancesResponseBody) SetData(v []*DescribeInstancesResponseBodyData) *DescribeInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstancesResponseBody) SetErrCode(v string) *DescribeInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetErrMessage(v string) *DescribeInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetHttpStatusCode(v int32) *DescribeInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetSuccess(v bool) *DescribeInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotal(v int32) *DescribeInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyData struct {
	// The ID of the network access control list (ACL).
	//
	// example:
	//
	// acl-bp1xc6b9vs013jjtp****
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AiFunctionInstanceId *string `json:"AiFunctionInstanceId,omitempty" xml:"AiFunctionInstanceId,omitempty"`
	// The instance architecture. Valid values:
	//
	// - onEci: deployed on ECI.
	//
	// - onECS: deployed on ECS.
	//
	// - onBareMetal: deployed on a bare metal resource pool.
	//
	// example:
	//
	// onECS
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 1733292921000
	BeginTime        *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EnableAiFunction *bool  `json:"EnableAiFunction,omitempty" xml:"EnableAiFunction,omitempty"`
	// Indicates whether automatic minor version upgrades are enabled.
	//
	// example:
	//
	// true
	EnableAutoMinorVersionUpgrade *bool `json:"EnableAutoMinorVersionUpgrade,omitempty" xml:"EnableAutoMinorVersionUpgrade,omitempty"`
	EnableMultiAz                 *bool `json:"EnableMultiAz,omitempty" xml:"EnableMultiAz,omitempty"`
	// Indicates whether SSL is enabled.
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Indicates whether the audit plugin is enabled.
	//
	// example:
	//
	// true
	EnabledAuditLoader *bool `json:"EnabledAuditLoader,omitempty" xml:"EnabledAuditLoader,omitempty"`
	// Indicates whether encryption is enabled.
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 4889001600000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// c-d4be777ff5e8cXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// sr_test_1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status. Valid values:
	//
	// - not_init: The instance is not initialized.
	//
	// - unpaid: The instance is pending payment.
	//
	// - paid: The payment is made.
	//
	// - creating: The instance is being created.
	//
	// - creating_failed: The instance failed to be created.
	//
	// - created: The instance is created.
	//
	// - running: The instance is running.
	//
	// - updating: The instance is being upgraded.
	//
	// - agent_creating: The agent is being created.
	//
	// - agent_scaling_up: The agent is being upgraded.
	//
	// - modifying_config: The configurations are being updated.
	//
	// - scaling_out: The instance is being scaled out.
	//
	// - restarting: The instance is restarting.
	//
	// - scaling_in: The instance is being scaled in.
	//
	// - scaling_up: The instance is being upgraded.
	//
	// - scaling_down: The instance is being downgraded.
	//
	// - upgrading: The instance is being upgraded.
	//
	// - enable_public_network: The public endpoint is being enabled.
	//
	// - disable_public_network: The public endpoint is being disabled.
	//
	// - convert_from_trial_to_official: The instance edition is being changed.
	//
	// - migration_cluster_to_serverless: The cluster is being migrated.
	//
	// - modifying_timezone: The time zone is being modified.
	//
	// - switch_az: The primary and secondary zones are being switched.
	//
	// - enabling: The instance is being resumed.
	//
	// - disable: The instance is unavailable.
	//
	// - actively_disabled: The instance is unavailable.
	//
	// - deleting: The instance is being deleted.
	//
	// - deleting_failed: The instance failed to be deleted.
	//
	// - deleted_with_error: The instance is deleted due to a creation failure.
	//
	// - deleted: The instance is deleted.
	//
	// example:
	//
	// running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether read/write splitting is enabled. If this parameter is set to true, the leader FE node processes write requests, and the other FE nodes process read requests.
	//
	// example:
	//
	// true
	IsolateLeader *bool `json:"IsolateLeader,omitempty" xml:"IsolateLeader,omitempty"`
	// The ID of the KMS key.
	//
	// example:
	//
	// rewqfds****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maintenance window of the instance. Valid values:
	//
	// - 00:00-06:00
	//
	// - 06:00-07:00
	//
	// - 07:00-08:00
	//
	// - 08:00-09:00
	//
	// - 09:00-10:00
	//
	// - 10:00-11:00
	//
	// - 11:00-12:00
	//
	// - 12:00-13:00
	//
	// - 13:00-14:00
	//
	// - 14:00-15:00
	//
	// - 15:00-16:00
	//
	// - 16:00-17:00
	//
	// - 17:00-18:00
	//
	// - 18:00-19:00
	//
	// - 19:00-20:00
	//
	// - 20:00-21:00
	//
	// - 21:00-22:00
	//
	// - 22:00-23:00
	//
	// - 23:00-24:00
	//
	// example:
	//
	// 00:00-06:00
	MaintainablePeriod *string `json:"MaintainablePeriod,omitempty" xml:"MaintainablePeriod,omitempty"`
	// The minor version number.
	//
	// example:
	//
	// 3.2.11-1.79-1.6.5
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The type of the monitoring service.
	//
	// example:
	//
	// cms
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// The OSS URL.
	//
	// example:
	//
	// oss://sr-c-****
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The instance edition. Valid values:
	//
	// - trial: Trial Edition.
	//
	// - official: Standard Edition.
	//
	// example:
	//
	// official
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The billing method:
	//
	// - prePaid: subscription.
	//
	// - postPaid: pay-as-you-go.
	//
	// example:
	//
	// postPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmytyuofb****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The running mode of the cluster:
	//
	// - shared_nothing: all-in-one.
	//
	// - shared_data: storage-compute separation.
	//
	// - lakehouse: data lake analytics.
	//
	// example:
	//
	// shared_nothing
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The duration for which the cluster has been running. Unit: seconds.
	//
	// example:
	//
	// 3645445
	RunningTime *int64 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// Indicates whether the security group is a managed security group.
	//
	// example:
	//
	// true
	SecurityGroupManaged *bool `json:"SecurityGroupManaged,omitempty" xml:"SecurityGroupManaged,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-8vbaz2etr66a62b9****
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// The tags attached to the instance.
	Tags []*DescribeInstancesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of vSwitches.
	VSwitches []*DescribeInstancesResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The cluster version.
	//
	// example:
	//
	// 3.2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1rbvag1cafkj4prwXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyData) GetAclId() *string {
	return s.AclId
}

func (s *DescribeInstancesResponseBodyData) GetAiFunctionInstanceId() *string {
	return s.AiFunctionInstanceId
}

func (s *DescribeInstancesResponseBodyData) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeInstancesResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeInstancesResponseBodyData) GetEnableAiFunction() *bool {
	return s.EnableAiFunction
}

func (s *DescribeInstancesResponseBodyData) GetEnableAutoMinorVersionUpgrade() *bool {
	return s.EnableAutoMinorVersionUpgrade
}

func (s *DescribeInstancesResponseBodyData) GetEnableMultiAz() *bool {
	return s.EnableMultiAz
}

func (s *DescribeInstancesResponseBodyData) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *DescribeInstancesResponseBodyData) GetEnabledAuditLoader() *bool {
	return s.EnabledAuditLoader
}

func (s *DescribeInstancesResponseBodyData) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeInstancesResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyData) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesResponseBodyData) GetIsolateLeader() *bool {
	return s.IsolateLeader
}

func (s *DescribeInstancesResponseBodyData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *DescribeInstancesResponseBodyData) GetMaintainablePeriod() *string {
	return s.MaintainablePeriod
}

func (s *DescribeInstancesResponseBodyData) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeInstancesResponseBodyData) GetMonitorType() *string {
	return s.MonitorType
}

func (s *DescribeInstancesResponseBodyData) GetOssLocation() *string {
	return s.OssLocation
}

func (s *DescribeInstancesResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeInstancesResponseBodyData) GetPayType() *string {
	return s.PayType
}

func (s *DescribeInstancesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyData) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeInstancesResponseBodyData) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *DescribeInstancesResponseBodyData) GetSecurityGroupManaged() *bool {
	return s.SecurityGroupManaged
}

func (s *DescribeInstancesResponseBodyData) GetSgId() *string {
	return s.SgId
}

func (s *DescribeInstancesResponseBodyData) GetTags() []*DescribeInstancesResponseBodyDataTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyData) GetVSwitches() []*DescribeInstancesResponseBodyDataVSwitches {
	return s.VSwitches
}

func (s *DescribeInstancesResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeInstancesResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyData) SetAclId(v string) *DescribeInstancesResponseBodyData {
	s.AclId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetAiFunctionInstanceId(v string) *DescribeInstancesResponseBodyData {
	s.AiFunctionInstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetArchitecture(v string) *DescribeInstancesResponseBodyData {
	s.Architecture = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetBeginTime(v int64) *DescribeInstancesResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnableAiFunction(v bool) *DescribeInstancesResponseBodyData {
	s.EnableAiFunction = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnableAutoMinorVersionUpgrade(v bool) *DescribeInstancesResponseBodyData {
	s.EnableAutoMinorVersionUpgrade = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnableMultiAz(v bool) *DescribeInstancesResponseBodyData {
	s.EnableMultiAz = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnableSSL(v bool) *DescribeInstancesResponseBodyData {
	s.EnableSSL = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnabledAuditLoader(v bool) *DescribeInstancesResponseBodyData {
	s.EnabledAuditLoader = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEncrypted(v bool) *DescribeInstancesResponseBodyData {
	s.Encrypted = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetExpireTime(v int64) *DescribeInstancesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetInstanceId(v string) *DescribeInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetInstanceName(v string) *DescribeInstancesResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetInstanceStatus(v string) *DescribeInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetIsolateLeader(v bool) *DescribeInstancesResponseBodyData {
	s.IsolateLeader = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetKmsKeyId(v string) *DescribeInstancesResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetMaintainablePeriod(v string) *DescribeInstancesResponseBodyData {
	s.MaintainablePeriod = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetMinorVersion(v string) *DescribeInstancesResponseBodyData {
	s.MinorVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetMonitorType(v string) *DescribeInstancesResponseBodyData {
	s.MonitorType = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetOssLocation(v string) *DescribeInstancesResponseBodyData {
	s.OssLocation = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetPackageType(v string) *DescribeInstancesResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetPayType(v string) *DescribeInstancesResponseBodyData {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetRegionId(v string) *DescribeInstancesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetResourceGroupId(v string) *DescribeInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetRunMode(v string) *DescribeInstancesResponseBodyData {
	s.RunMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetRunningTime(v int64) *DescribeInstancesResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetSecurityGroupManaged(v bool) *DescribeInstancesResponseBodyData {
	s.SecurityGroupManaged = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetSgId(v string) *DescribeInstancesResponseBodyData {
	s.SgId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetTags(v []*DescribeInstancesResponseBodyDataTags) *DescribeInstancesResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetVSwitches(v []*DescribeInstancesResponseBodyDataVSwitches) *DescribeInstancesResponseBodyData {
	s.VSwitches = v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetVersion(v string) *DescribeInstancesResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetVpcId(v string) *DescribeInstancesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyDataTags) SetKey(v string) *DescribeInstancesResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyDataTags) SetValue(v string) *DescribeInstancesResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyDataVSwitches struct {
	// Indicates whether the vSwitch is the primary vSwitch.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1p0mldwx5av55v0xXXX
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyDataVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyDataVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyDataVSwitches) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeInstancesResponseBodyDataVSwitches) GetVswId() *string {
	return s.VswId
}

func (s *DescribeInstancesResponseBodyDataVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyDataVSwitches) SetPrimary(v bool) *DescribeInstancesResponseBodyDataVSwitches {
	s.Primary = &v
	return s
}

func (s *DescribeInstancesResponseBodyDataVSwitches) SetVswId(v string) *DescribeInstancesResponseBodyDataVSwitches {
	s.VswId = &v
	return s
}

func (s *DescribeInstancesResponseBodyDataVSwitches) SetZoneId(v string) *DescribeInstancesResponseBodyDataVSwitches {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyDataVSwitches) Validate() error {
	return dara.Validate(s)
}
