// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayResponseBody
	GetCode() *int32
	SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody
	GetData() *GetGatewayResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayResponseBody
	GetSuccess() *bool
}

type GetGatewayResponseBody struct {
	// The status code returned. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the gateway.
	Data *GetGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// 9C96CDF8-9E6C-5AB6-B83C-8F87A10117E6
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

func (s GetGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayResponseBody) GetData() *GetGatewayResponseBodyData {
	return s.Data
}

func (s *GetGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayResponseBody) SetCode(v int32) *GetGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetHttpStatusCode(v int32) *GetGatewayResponseBody {
	s.HttpStatusCode = &v
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

func (s *GetGatewayResponseBody) SetSuccess(v bool) *GetGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResponseBodyData struct {
	// The billing method, such as subscription or pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Indicates whether auto scale-out is enabled.
	//
	// example:
	//
	// true
	Elastic *bool `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// The auto scale-out policy.
	ElasticPolicy *GetGatewayResponseBodyDataElasticPolicy `json:"ElasticPolicy,omitempty" xml:"ElasticPolicy,omitempty" type:"Struct"`
	// The number of replicas that are automatically scaled out.
	//
	// example:
	//
	// 1
	ElasticReplica *int32 `json:"ElasticReplica,omitempty" xml:"ElasticReplica,omitempty"`
	// The type of auto scale-out. Valid value:
	//
	// 	- CronHPA: scale-out by time
	//
	// example:
	//
	// CronHPA
	ElasticType *string `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	// The time when the gateway expires.
	//
	// example:
	//
	// 2021-08-01 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the gateway was created. The time is displayed in GMT. The time is the local time of the region in which the gateway resides.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the gateway was last modified.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The log configuration.
	LogConfigDetails *GetGatewayResponseBodyDataLogConfigDetails `json:"LogConfigDetails,omitempty" xml:"LogConfigDetails,omitempty" type:"Struct"`
	// The tag of the resource.
	//
	// example:
	//
	// {"TagKey":"TagValue"}
	MseTag *string `json:"MseTag,omitempty" xml:"MseTag,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// DEFAULT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the user who created the gateway.
	//
	// example:
	//
	// 1231254
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of gateway replicas.
	//
	// example:
	//
	// 2
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm34x43l*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp11ufzqn6mmb8dtzz82
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The specifications of the gateway.
	//
	// example:
	//
	// MSE_GTW_16_32_200_c
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the gateway. Valid values:
	//
	// 	- 0: The gateway is being created.
	//
	// 	- 1: The gateway fails to be created.
	//
	// 	- 2: The gateway is running.
	//
	// 	- 3: The gateway is changing.
	//
	// 	- 4: The gateway is scaling in.
	//
	// 	- 6: The gateway is scaling out.
	//
	// 	- 8: The gateway is being deleted.
	//
	// 	- 10: The gateway is restarting.
	//
	// 	- 11: The gateway is being rebuilt.
	//
	// 	- 12: The gateway is updating.
	//
	// 	- 13: The gateway fails to be updated.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// Restarting
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The total number of replicas, including the number of replicas that are automatically scaled out.
	//
	// example:
	//
	// 3
	TotalReplica *int32 `json:"TotalReplica,omitempty" xml:"TotalReplica,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1328cm01m6uel42b5zb
	Vpc *string `json:"Vpc,omitempty" xml:"Vpc,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp18zeqxx6mpuq843z4n5
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	// The ID of the secondary vSwitch.
	//
	// example:
	//
	// vpc-bp1j6ivhav15ve7q54kq2
	Vswitch2 *string `json:"Vswitch2,omitempty" xml:"Vswitch2,omitempty"`
	// The details of Tracing Analysis.
	XtraceDetails *GetGatewayResponseBodyDataXtraceDetails `json:"XtraceDetails,omitempty" xml:"XtraceDetails,omitempty" type:"Struct"`
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

func (s *GetGatewayResponseBodyData) GetElastic() *bool {
	return s.Elastic
}

func (s *GetGatewayResponseBodyData) GetElasticPolicy() *GetGatewayResponseBodyDataElasticPolicy {
	return s.ElasticPolicy
}

func (s *GetGatewayResponseBodyData) GetElasticReplica() *int32 {
	return s.ElasticReplica
}

func (s *GetGatewayResponseBodyData) GetElasticType() *string {
	return s.ElasticType
}

func (s *GetGatewayResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *GetGatewayResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetGatewayResponseBodyData) GetLogConfigDetails() *GetGatewayResponseBodyDataLogConfigDetails {
	return s.LogConfigDetails
}

func (s *GetGatewayResponseBodyData) GetMseTag() *string {
	return s.MseTag
}

func (s *GetGatewayResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayResponseBodyData) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *GetGatewayResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetGatewayResponseBodyData) GetReplica() *int32 {
	return s.Replica
}

func (s *GetGatewayResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetGatewayResponseBodyData) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *GetGatewayResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *GetGatewayResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetGatewayResponseBodyData) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetGatewayResponseBodyData) GetTotalReplica() *int32 {
	return s.TotalReplica
}

func (s *GetGatewayResponseBodyData) GetVpc() *string {
	return s.Vpc
}

func (s *GetGatewayResponseBodyData) GetVswitch() *string {
	return s.Vswitch
}

func (s *GetGatewayResponseBodyData) GetVswitch2() *string {
	return s.Vswitch2
}

func (s *GetGatewayResponseBodyData) GetXtraceDetails() *GetGatewayResponseBodyDataXtraceDetails {
	return s.XtraceDetails
}

func (s *GetGatewayResponseBodyData) SetChargeType(v string) *GetGatewayResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetElastic(v bool) *GetGatewayResponseBodyData {
	s.Elastic = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetElasticPolicy(v *GetGatewayResponseBodyDataElasticPolicy) *GetGatewayResponseBodyData {
	s.ElasticPolicy = v
	return s
}

func (s *GetGatewayResponseBodyData) SetElasticReplica(v int32) *GetGatewayResponseBodyData {
	s.ElasticReplica = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetElasticType(v string) *GetGatewayResponseBodyData {
	s.ElasticType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEndDate(v string) *GetGatewayResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGmtCreate(v string) *GetGatewayResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGmtModified(v string) *GetGatewayResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetId(v int64) *GetGatewayResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetInstanceId(v string) *GetGatewayResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLogConfigDetails(v *GetGatewayResponseBodyDataLogConfigDetails) *GetGatewayResponseBodyData {
	s.LogConfigDetails = v
	return s
}

func (s *GetGatewayResponseBodyData) SetMseTag(v string) *GetGatewayResponseBodyData {
	s.MseTag = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetPrimaryUser(v string) *GetGatewayResponseBodyData {
	s.PrimaryUser = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetRegion(v string) *GetGatewayResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetReplica(v int32) *GetGatewayResponseBodyData {
	s.Replica = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetResourceGroupId(v string) *GetGatewayResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSecurityGroup(v string) *GetGatewayResponseBodyData {
	s.SecurityGroup = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSpec(v string) *GetGatewayResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatus(v int32) *GetGatewayResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatusDesc(v string) *GetGatewayResponseBodyData {
	s.StatusDesc = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetTotalReplica(v int32) *GetGatewayResponseBodyData {
	s.TotalReplica = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVpc(v string) *GetGatewayResponseBodyData {
	s.Vpc = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVswitch(v string) *GetGatewayResponseBodyData {
	s.Vswitch = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVswitch2(v string) *GetGatewayResponseBodyData {
	s.Vswitch2 = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetXtraceDetails(v *GetGatewayResponseBodyDataXtraceDetails) *GetGatewayResponseBodyData {
	s.XtraceDetails = v
	return s
}

func (s *GetGatewayResponseBodyData) Validate() error {
	if s.ElasticPolicy != nil {
		if err := s.ElasticPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfigDetails != nil {
		if err := s.LogConfigDetails.Validate(); err != nil {
			return err
		}
	}
	if s.XtraceDetails != nil {
		if err := s.XtraceDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResponseBodyDataElasticPolicy struct {
	// The type of auto scale-out. Valid value:
	//
	// 	- CronHPA: scale-out by time
	//
	// example:
	//
	// CronHPA
	ElasticType *string `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	// The maximum number of instances that are automatically scaled out. This parameter is used for horizontal scale-out.
	//
	// example:
	//
	// 5
	MaxReplica *int32 `json:"MaxReplica,omitempty" xml:"MaxReplica,omitempty"`
	// The policy of scale-out by time.
	TimePolicyList []*GetGatewayResponseBodyDataElasticPolicyTimePolicyList `json:"TimePolicyList,omitempty" xml:"TimePolicyList,omitempty" type:"Repeated"`
}

func (s GetGatewayResponseBodyDataElasticPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataElasticPolicy) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataElasticPolicy) GetElasticType() *string {
	return s.ElasticType
}

func (s *GetGatewayResponseBodyDataElasticPolicy) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *GetGatewayResponseBodyDataElasticPolicy) GetTimePolicyList() []*GetGatewayResponseBodyDataElasticPolicyTimePolicyList {
	return s.TimePolicyList
}

func (s *GetGatewayResponseBodyDataElasticPolicy) SetElasticType(v string) *GetGatewayResponseBodyDataElasticPolicy {
	s.ElasticType = &v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicy) SetMaxReplica(v int32) *GetGatewayResponseBodyDataElasticPolicy {
	s.MaxReplica = &v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicy) SetTimePolicyList(v []*GetGatewayResponseBodyDataElasticPolicyTimePolicyList) *GetGatewayResponseBodyDataElasticPolicy {
	s.TimePolicyList = v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicy) Validate() error {
	if s.TimePolicyList != nil {
		for _, item := range s.TimePolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayResponseBodyDataElasticPolicyTimePolicyList struct {
	// The number of expected replicas.
	//
	// example:
	//
	// 2
	DesiredReplica *int32 `json:"DesiredReplica,omitempty" xml:"DesiredReplica,omitempty"`
	// The end time of auto scale-out.
	//
	// example:
	//
	// 18:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of auto scale-out.
	//
	// example:
	//
	// 16:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetGatewayResponseBodyDataElasticPolicyTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataElasticPolicyTimePolicyList) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) GetDesiredReplica() *int32 {
	return s.DesiredReplica
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) SetDesiredReplica(v int32) *GetGatewayResponseBodyDataElasticPolicyTimePolicyList {
	s.DesiredReplica = &v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) SetEndTime(v string) *GetGatewayResponseBodyDataElasticPolicyTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) SetStartTime(v string) *GetGatewayResponseBodyDataElasticPolicyTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *GetGatewayResponseBodyDataElasticPolicyTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataLogConfigDetails struct {
	// Indicates whether Log Service is activated.
	//
	// example:
	//
	// true
	LogEnabled *bool `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// mse_access_log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// k8s-log-c173117256e934a96b7fefdf2ef8a8057
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetGatewayResponseBodyDataLogConfigDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataLogConfigDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) GetLogEnabled() *bool {
	return s.LogEnabled
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetLogEnabled(v bool) *GetGatewayResponseBodyDataLogConfigDetails {
	s.LogEnabled = &v
	return s
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetLogStoreName(v string) *GetGatewayResponseBodyDataLogConfigDetails {
	s.LogStoreName = &v
	return s
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetProjectName(v string) *GetGatewayResponseBodyDataLogConfigDetails {
	s.ProjectName = &v
	return s
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) Validate() error {
	return dara.Validate(s)
}

type GetGatewayResponseBodyDataXtraceDetails struct {
	// The sampling rate of Tracing Analysis.
	//
	// example:
	//
	// 10
	Sample *int32 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// Indicates whether sampling by using Tracing Analysis is enabled.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"TraceOn,omitempty" xml:"TraceOn,omitempty"`
}

func (s GetGatewayResponseBodyDataXtraceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResponseBodyDataXtraceDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataXtraceDetails) GetSample() *int32 {
	return s.Sample
}

func (s *GetGatewayResponseBodyDataXtraceDetails) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *GetGatewayResponseBodyDataXtraceDetails) SetSample(v int32) *GetGatewayResponseBodyDataXtraceDetails {
	s.Sample = &v
	return s
}

func (s *GetGatewayResponseBodyDataXtraceDetails) SetTraceOn(v bool) *GetGatewayResponseBodyDataXtraceDetails {
	s.TraceOn = &v
	return s
}

func (s *GetGatewayResponseBodyDataXtraceDetails) Validate() error {
	return dara.Validate(s)
}
