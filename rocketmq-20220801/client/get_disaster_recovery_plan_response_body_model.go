// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisasterRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetDisasterRecoveryPlanResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetDisasterRecoveryPlanResponseBody
	GetCode() *string
	SetData(v *GetDisasterRecoveryPlanResponseBodyData) *GetDisasterRecoveryPlanResponseBody
	GetData() *GetDisasterRecoveryPlanResponseBodyData
	SetDynamicCode(v string) *GetDisasterRecoveryPlanResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetDisasterRecoveryPlanResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetDisasterRecoveryPlanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDisasterRecoveryPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDisasterRecoveryPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDisasterRecoveryPlanResponseBody
	GetSuccess() *bool
}

type GetDisasterRecoveryPlanResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetDisasterRecoveryPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// xxx
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-611C63E0xxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDisasterRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryPlanResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetDisasterRecoveryPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDisasterRecoveryPlanResponseBody) GetData() *GetDisasterRecoveryPlanResponseBodyData {
	return s.Data
}

func (s *GetDisasterRecoveryPlanResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetDisasterRecoveryPlanResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetDisasterRecoveryPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDisasterRecoveryPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDisasterRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDisasterRecoveryPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDisasterRecoveryPlanResponseBody) SetAccessDeniedDetail(v string) *GetDisasterRecoveryPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetCode(v string) *GetDisasterRecoveryPlanResponseBody {
	s.Code = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetData(v *GetDisasterRecoveryPlanResponseBodyData) *GetDisasterRecoveryPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetDynamicCode(v string) *GetDisasterRecoveryPlanResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetDynamicMessage(v string) *GetDisasterRecoveryPlanResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetHttpStatusCode(v int32) *GetDisasterRecoveryPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetMessage(v string) *GetDisasterRecoveryPlanResponseBody {
	s.Message = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetRequestId(v string) *GetDisasterRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) SetSuccess(v bool) *GetDisasterRecoveryPlanResponseBody {
	s.Success = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDisasterRecoveryPlanResponseBodyData struct {
	// Indicates whether automatic consumer progress synchronization is enabled.
	//
	// >  This parameter takes effect only when `syncCheckpointEnabled` is set to true.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoSyncCheckpoint *bool `json:"autoSyncCheckpoint,omitempty" xml:"autoSyncCheckpoint,omitempty"`
	// The time when the query task was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The extended information.
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// The instances involved in the Global Replicator task.
	Instances []*GetDisasterRecoveryPlanResponseBodyDataInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The description of the Global Replicator task.
	//
	// example:
	//
	// xxxx
	PlanDesc *string `json:"planDesc,omitempty" xml:"planDesc,omitempty"`
	// The ID of the Global Replicator task.
	//
	// example:
	//
	// 1300000016
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// The name of the Global Replicator task.
	//
	// example:
	//
	// xxx
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The status of the Global Replicator task. Valid values:
	//
	// 	- CREATED
	//
	// 	- RUNNING
	//
	// 	- DELETED
	//
	// example:
	//
	// RUNNING
	PlanStatus *string `json:"planStatus,omitempty" xml:"planStatus,omitempty"`
	// The type of the Global Replicator task. Valid values:
	//
	// 	- ACTIVE_PASSIVE: one-way backup
	//
	// 	- ACTIVE_ACTIVE: two-way backup
	//
	// example:
	//
	// ACTIVE_PASSIVE
	PlanType *string `json:"planType,omitempty" xml:"planType,omitempty"`
	// Indicates whether consumer progress synchronization is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SyncCheckpointEnabled *bool `json:"syncCheckpointEnabled,omitempty" xml:"syncCheckpointEnabled,omitempty"`
	// The time when the query task was last modified.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDisasterRecoveryPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetAutoSyncCheckpoint() *bool {
	return s.AutoSyncCheckpoint
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetExtInfo() map[string]*string {
	return s.ExtInfo
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetInstances() []*GetDisasterRecoveryPlanResponseBodyDataInstances {
	return s.Instances
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetPlanName() *string {
	return s.PlanName
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetPlanType() *string {
	return s.PlanType
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetSyncCheckpointEnabled() *bool {
	return s.SyncCheckpointEnabled
}

func (s *GetDisasterRecoveryPlanResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetAutoSyncCheckpoint(v bool) *GetDisasterRecoveryPlanResponseBodyData {
	s.AutoSyncCheckpoint = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetCreateTime(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetExtInfo(v map[string]*string) *GetDisasterRecoveryPlanResponseBodyData {
	s.ExtInfo = v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetInstances(v []*GetDisasterRecoveryPlanResponseBodyDataInstances) *GetDisasterRecoveryPlanResponseBodyData {
	s.Instances = v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetPlanDesc(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.PlanDesc = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetPlanId(v int64) *GetDisasterRecoveryPlanResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetPlanName(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.PlanName = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetPlanStatus(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.PlanStatus = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetPlanType(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.PlanType = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetSyncCheckpointEnabled(v bool) *GetDisasterRecoveryPlanResponseBodyData {
	s.SyncCheckpointEnabled = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) SetUpdateTime(v string) *GetDisasterRecoveryPlanResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDisasterRecoveryPlanResponseBodyDataInstances struct {
	// The authentication type.
	//
	// example:
	//
	// ACL_AUTH
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The consumer Group ID.
	//
	// example:
	//
	// GID_DS_XXX_YYY
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// xxx
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-gpz3qtcdxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance role.
	//
	// example:
	//
	// ACTIVE
	InstanceRole *string `json:"instanceRole,omitempty" xml:"instanceRole,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The message attribute.
	MessageProperty *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty `json:"messageProperty,omitempty" xml:"messageProperty,omitempty" type:"Struct"`
	// The network type.
	//
	// example:
	//
	// TCP_INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The password used for authentication.
	//
	// example:
	//
	// xxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp17hpmgz9******
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The username used for authentication.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-wz9qt50xhtj9krb******
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetDisasterRecoveryPlanResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryPlanResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetAuthType() *string {
	return s.AuthType
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetMessageProperty() *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty {
	return s.MessageProperty
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetPassword() *string {
	return s.Password
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetUsername() *string {
	return s.Username
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetAuthType(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.AuthType = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetConsumerGroupId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetEndpointUrl(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.EndpointUrl = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetInstanceId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetInstanceRole(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.InstanceRole = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetInstanceType(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetMessageProperty(v *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.MessageProperty = v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetNetworkType(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.NetworkType = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetPassword(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.Password = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetRegionId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.RegionId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetSecurityGroupId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetUsername(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.Username = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetVSwitchId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.VSwitchId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) SetVpcId(v string) *GetDisasterRecoveryPlanResponseBodyDataInstances {
	s.VpcId = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}

type GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty struct {
	// The attribute key.
	//
	// example:
	//
	// xxx
	PropertyKey *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// xxx
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) SetPropertyKey(v string) *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty {
	s.PropertyKey = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) SetPropertyValue(v string) *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty {
	s.PropertyValue = &v
	return s
}

func (s *GetDisasterRecoveryPlanResponseBodyDataInstancesMessageProperty) Validate() error {
	return dara.Validate(s)
}
