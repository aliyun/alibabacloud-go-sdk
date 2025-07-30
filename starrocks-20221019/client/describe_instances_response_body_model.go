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
	Data []*DescribeInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyData struct {
	// example:
	//
	// acl-bp1xc6b9vs013jjtp****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// example:
	//
	// onECS
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// 1733292921000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// true
	EnableAutoMinorVersionUpgrade *bool `json:"EnableAutoMinorVersionUpgrade,omitempty" xml:"EnableAutoMinorVersionUpgrade,omitempty"`
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// example:
	//
	// true
	EnabledAuditLoader *bool `json:"EnabledAuditLoader,omitempty" xml:"EnabledAuditLoader,omitempty"`
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// example:
	//
	// 4889001600000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// c-d4be777ff5e8cXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// sr_test_1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// true
	IsolateLeader *bool `json:"IsolateLeader,omitempty" xml:"IsolateLeader,omitempty"`
	// example:
	//
	// rewqfds****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// example:
	//
	// 00:00-06:00
	MaintainablePeriod *string `json:"MaintainablePeriod,omitempty" xml:"MaintainablePeriod,omitempty"`
	// example:
	//
	// 3.2.11-1.79-1.6.5
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// cms
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// oss://sr-c-****
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// example:
	//
	// official
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// postPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmytyuofb****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// shared_nothing
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// example:
	//
	// 3645445
	RunningTime *int64 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// example:
	//
	// true
	SecurityGroupManaged *bool `json:"SecurityGroupManaged,omitempty" xml:"SecurityGroupManaged,omitempty"`
	// example:
	//
	// sg-8vbaz2etr66a62b9****
	SgId      *string                                       `json:"SgId,omitempty" xml:"SgId,omitempty"`
	Tags      []*DescribeInstancesResponseBodyDataTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitches []*DescribeInstancesResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// example:
	//
	// 3.2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// VPC ID。
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

func (s *DescribeInstancesResponseBodyData) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeInstancesResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeInstancesResponseBodyData) GetEnableAutoMinorVersionUpgrade() *bool {
	return s.EnableAutoMinorVersionUpgrade
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

func (s *DescribeInstancesResponseBodyData) SetArchitecture(v string) *DescribeInstancesResponseBodyData {
	s.Architecture = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetBeginTime(v int64) *DescribeInstancesResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyData) SetEnableAutoMinorVersionUpgrade(v bool) *DescribeInstancesResponseBodyData {
	s.EnableAutoMinorVersionUpgrade = &v
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
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyDataTags struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// example:
	//
	// vsw-bp1p0mldwx5av55v0xXXX
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
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
