// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListDisasterRecoveryPlansResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListDisasterRecoveryPlansResponseBody
	GetCode() *string
	SetData(v *ListDisasterRecoveryPlansResponseBodyData) *ListDisasterRecoveryPlansResponseBody
	GetData() *ListDisasterRecoveryPlansResponseBodyData
	SetDynamicCode(v string) *ListDisasterRecoveryPlansResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListDisasterRecoveryPlansResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListDisasterRecoveryPlansResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDisasterRecoveryPlansResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDisasterRecoveryPlansResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDisasterRecoveryPlansResponseBody
	GetSuccess() *bool
}

type ListDisasterRecoveryPlansResponseBody struct {
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
	Data *ListDisasterRecoveryPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The returned dynamic error message.
	//
	// example:
	//
	// InstanceId
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
	// 855EF8E6-9C1D-5DE2-9E84-924E13Exxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the information about the service was queried.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDisasterRecoveryPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListDisasterRecoveryPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDisasterRecoveryPlansResponseBody) GetData() *ListDisasterRecoveryPlansResponseBodyData {
	return s.Data
}

func (s *ListDisasterRecoveryPlansResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListDisasterRecoveryPlansResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListDisasterRecoveryPlansResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDisasterRecoveryPlansResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDisasterRecoveryPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDisasterRecoveryPlansResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDisasterRecoveryPlansResponseBody) SetAccessDeniedDetail(v string) *ListDisasterRecoveryPlansResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetCode(v string) *ListDisasterRecoveryPlansResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetData(v *ListDisasterRecoveryPlansResponseBodyData) *ListDisasterRecoveryPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetDynamicCode(v string) *ListDisasterRecoveryPlansResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetDynamicMessage(v string) *ListDisasterRecoveryPlansResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetHttpStatusCode(v int32) *ListDisasterRecoveryPlansResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetMessage(v string) *ListDisasterRecoveryPlansResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetRequestId(v string) *ListDisasterRecoveryPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) SetSuccess(v bool) *ListDisasterRecoveryPlansResponseBody {
	s.Success = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryPlansResponseBodyData struct {
	// The Global Replicator tasks.
	List []*ListDisasterRecoveryPlansResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scroll ID of the request. The ID is automatically generated by the system. The result can be paginated only if this parameter is included in the pagination request.
	//
	// example:
	//
	// B13D0B07-F24B-4790-88D8-D47A38063D00
	ScrollId *string `json:"scrollId,omitempty" xml:"scrollId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 28
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDisasterRecoveryPlansResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponseBodyData) GetList() []*ListDisasterRecoveryPlansResponseBodyDataList {
	return s.List
}

func (s *ListDisasterRecoveryPlansResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryPlansResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDisasterRecoveryPlansResponseBodyData) GetScrollId() *string {
	return s.ScrollId
}

func (s *ListDisasterRecoveryPlansResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDisasterRecoveryPlansResponseBodyData) SetList(v []*ListDisasterRecoveryPlansResponseBodyDataList) *ListDisasterRecoveryPlansResponseBodyData {
	s.List = v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyData) SetPageNumber(v int64) *ListDisasterRecoveryPlansResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyData) SetPageSize(v int64) *ListDisasterRecoveryPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyData) SetScrollId(v string) *ListDisasterRecoveryPlansResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyData) SetTotalCount(v int64) *ListDisasterRecoveryPlansResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryPlansResponseBodyDataList struct {
	// Indicates whether automatic consumer progress synchronization is enabled.
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
	Instances []*ListDisasterRecoveryPlansResponseBodyDataListInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The description of the Global Replicator task.
	//
	// example:
	//
	// xxx
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

func (s ListDisasterRecoveryPlansResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetAutoSyncCheckpoint() *bool {
	return s.AutoSyncCheckpoint
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetExtInfo() map[string]*string {
	return s.ExtInfo
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetInstances() []*ListDisasterRecoveryPlansResponseBodyDataListInstances {
	return s.Instances
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetPlanId() *int64 {
	return s.PlanId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetPlanName() *string {
	return s.PlanName
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetPlanType() *string {
	return s.PlanType
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetSyncCheckpointEnabled() *bool {
	return s.SyncCheckpointEnabled
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetAutoSyncCheckpoint(v bool) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.AutoSyncCheckpoint = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetCreateTime(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetExtInfo(v map[string]*string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.ExtInfo = v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetInstances(v []*ListDisasterRecoveryPlansResponseBodyDataListInstances) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.Instances = v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetPlanDesc(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.PlanDesc = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetPlanId(v int64) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetPlanName(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.PlanName = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetPlanStatus(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.PlanStatus = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetPlanType(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.PlanType = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetSyncCheckpointEnabled(v bool) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.SyncCheckpointEnabled = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) SetUpdateTime(v string) *ListDisasterRecoveryPlansResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryPlansResponseBodyDataListInstances struct {
	// The authentication type.
	//
	// example:
	//
	// NO_AUTH
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// GID_ui_xxx
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
	// rmq-cn-ot93rbxxx
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
	MessageProperty *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty `json:"messageProperty,omitempty" xml:"messageProperty,omitempty" type:"Struct"`
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
	// cn-hangzhou
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
	// vpc-bp13docqysrgxtbxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListDisasterRecoveryPlansResponseBodyDataListInstances) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponseBodyDataListInstances) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetAuthType() *string {
	return s.AuthType
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetMessageProperty() *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty {
	return s.MessageProperty
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetPassword() *string {
	return s.Password
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetUsername() *string {
	return s.Username
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetAuthType(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.AuthType = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetConsumerGroupId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetEndpointUrl(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.EndpointUrl = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetInstanceId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetInstanceRole(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.InstanceRole = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetInstanceType(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.InstanceType = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetMessageProperty(v *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.MessageProperty = v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetNetworkType(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.NetworkType = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetPassword(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.Password = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetRegionId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.RegionId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetSecurityGroupId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetUsername(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.Username = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetVSwitchId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.VSwitchId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) SetVpcId(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstances {
	s.VpcId = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstances) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty struct {
	// The attribute key.
	//
	// example:
	//
	// aaa
	PropertyKey *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// bbb
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) SetPropertyKey(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty {
	s.PropertyKey = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) SetPropertyValue(v string) *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty {
	s.PropertyValue = &v
	return s
}

func (s *ListDisasterRecoveryPlansResponseBodyDataListInstancesMessageProperty) Validate() error {
	return dara.Validate(s)
}
