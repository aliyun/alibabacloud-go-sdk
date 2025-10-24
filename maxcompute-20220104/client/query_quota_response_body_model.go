// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryQuotaResponseBodyData) *QueryQuotaResponseBody
	GetData() *QueryQuotaResponseBodyData
	SetErrorCode(v string) *QueryQuotaResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryQuotaResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *QueryQuotaResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *QueryQuotaResponseBody
	GetRequestId() *string
}

type QueryQuotaResponseBody struct {
	// The data returned.
	Data *QueryQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Exception information
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1eeed16675342848904412e08dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBody) GetData() *QueryQuotaResponseBodyData {
	return s.Data
}

func (s *QueryQuotaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryQuotaResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryQuotaResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryQuotaResponseBody) SetData(v *QueryQuotaResponseBodyData) *QueryQuotaResponseBody {
	s.Data = v
	return s
}

func (s *QueryQuotaResponseBody) SetErrorCode(v string) *QueryQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryQuotaResponseBody) SetErrorMsg(v string) *QueryQuotaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryQuotaResponseBody) SetHttpCode(v int32) *QueryQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryQuotaResponseBody) SetRequestId(v string) *QueryQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryQuotaResponseBodyData struct {
	// The order information.
	BillingPolicy *QueryQuotaResponseBodyDataBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the Managed Service for Prometheus cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 1714356241163
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the quota plan.
	//
	// example:
	//
	// 1248953767546358
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The group name.
	//
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 2523
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The quota name.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The quota alias.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The quota description.
	//
	// example:
	//
	// {"minCU":10,
	//
	// "adhocCU":0,
	//
	// "maxCU":10,
	//
	// "schedulerType":"Fair",
	//
	// }
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The parent resource ID.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifiers of MaxCompute quota objects. The identifiers are the same as those in the Alibaba Cloud sales bill. This parameter is used for tags.
	SaleTag *QueryQuotaResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *QueryQuotaResponseBodyDataScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The level-2 quotas.
	SubQuotaInfoList []*QueryQuotaResponseBodyDataSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The resource tag of a quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter corresponds to the resourceSystemType field.
	//
	// example:
	//
	// FUXI_OFFLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyData) GetBillingPolicy() *QueryQuotaResponseBodyDataBillingPolicy {
	return s.BillingPolicy
}

func (s *QueryQuotaResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *QueryQuotaResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryQuotaResponseBodyData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryQuotaResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryQuotaResponseBodyData) GetId() *string {
	return s.Id
}

func (s *QueryQuotaResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryQuotaResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *QueryQuotaResponseBodyData) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *QueryQuotaResponseBodyData) GetParentId() *string {
	return s.ParentId
}

func (s *QueryQuotaResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryQuotaResponseBodyData) GetSaleTag() *QueryQuotaResponseBodyDataSaleTag {
	return s.SaleTag
}

func (s *QueryQuotaResponseBodyData) GetScheduleInfo() *QueryQuotaResponseBodyDataScheduleInfo {
	return s.ScheduleInfo
}

func (s *QueryQuotaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryQuotaResponseBodyData) GetSubQuotaInfoList() []*QueryQuotaResponseBodyDataSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *QueryQuotaResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *QueryQuotaResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryQuotaResponseBodyData) GetType() *string {
	return s.Type
}

func (s *QueryQuotaResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *QueryQuotaResponseBodyData) SetBillingPolicy(v *QueryQuotaResponseBodyDataBillingPolicy) *QueryQuotaResponseBodyData {
	s.BillingPolicy = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCluster(v string) *QueryQuotaResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCreateTime(v int64) *QueryQuotaResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetCreatorId(v string) *QueryQuotaResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetGroupName(v string) *QueryQuotaResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetId(v string) *QueryQuotaResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetName(v string) *QueryQuotaResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetNickName(v string) *QueryQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetParameter(v map[string]interface{}) *QueryQuotaResponseBodyData {
	s.Parameter = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetParentId(v string) *QueryQuotaResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetRegionId(v string) *QueryQuotaResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetSaleTag(v *QueryQuotaResponseBodyDataSaleTag) *QueryQuotaResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetScheduleInfo(v *QueryQuotaResponseBodyDataScheduleInfo) *QueryQuotaResponseBodyData {
	s.ScheduleInfo = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetStatus(v string) *QueryQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetSubQuotaInfoList(v []*QueryQuotaResponseBodyDataSubQuotaInfoList) *QueryQuotaResponseBodyData {
	s.SubQuotaInfoList = v
	return s
}

func (s *QueryQuotaResponseBodyData) SetTag(v string) *QueryQuotaResponseBodyData {
	s.Tag = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetTenantId(v string) *QueryQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetType(v string) *QueryQuotaResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryQuotaResponseBodyData) SetVersion(v string) *QueryQuotaResponseBodyData {
	s.Version = &v
	return s
}

func (s *QueryQuotaResponseBodyData) Validate() error {
	if s.BillingPolicy != nil {
		if err := s.BillingPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.SaleTag != nil {
		if err := s.SaleTag.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SubQuotaInfoList != nil {
		for _, item := range s.SubQuotaInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryQuotaResponseBodyDataBillingPolicy struct {
	// The billing method. Valid values:
	//
	// 	- subscription: the subscription quota.
	//
	// 	- payasyougo: the pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// In MaxCompute, instanceId and orderId are considered the same.
	//
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The order specifications.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s QueryQuotaResponseBodyDataBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataBillingPolicy) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetBillingMethod(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetInstanceId(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.InstanceId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetOdpsSpecCode(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) SetOrderId(v string) *QueryQuotaResponseBodyDataBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type QueryQuotaResponseBodyDataSaleTag struct {
	// The identifier of a MaxCompute quota object. This identifier exists in the Alibaba Cloud sales bill. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The object type. Valid values: quota and project.
	//
	// example:
	//
	// project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s QueryQuotaResponseBodyDataSaleTag) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *QueryQuotaResponseBodyDataSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryQuotaResponseBodyDataSaleTag) SetResourceIds(v []*string) *QueryQuotaResponseBodyDataSaleTag {
	s.ResourceIds = v
	return s
}

func (s *QueryQuotaResponseBodyDataSaleTag) SetResourceType(v string) *QueryQuotaResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSaleTag) Validate() error {
	return dara.Validate(s)
}

type QueryQuotaResponseBodyDataScheduleInfo struct {
	// The current quota plan that has taken effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the plan specified by currPlan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the plan specified by nextPlan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan specified by this parameter is triggered and the plan is different from the plan specified by currPlan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan specified by oncePlan is scheduled.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The operator name.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryQuotaResponseBodyDataScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataScheduleInfo) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetCurrPlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetCurrTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetNextPlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetNextTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOncePlan(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOnceTime(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetOperatorName(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) SetTimezone(v string) *QueryQuotaResponseBodyDataScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *QueryQuotaResponseBodyDataScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type QueryQuotaResponseBodyDataSubQuotaInfoList struct {
	// The order information.
	BillingPolicy *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 1688653978768
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the quota plan.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The role group name.
	//
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The ID of the level-2 quota.
	//
	// example:
	//
	// 1000048
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The quota description.
	//
	// example:
	//
	// {\\"maxCU\\": 10, \\"minCU\\": 10, \\"adhocCU\\": 0, \\"schedulerType\\": \\"Fifo\\"}
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The parent resource ID.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifiers of MaxCompute quota objects. The identifiers are the same as those in the Alibaba Cloud sales bill. This parameter is used for tags.
	SaleTag *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The resource tag of a quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter corresponds to the resourceSystemType field.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetBillingPolicy() *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetSaleTag() *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetScheduleInfo() *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetBillingPolicy(v *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCluster(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCreateTime(v int64) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetCreatorId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetGroupName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.GroupName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetNickName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetParameter(v map[string]interface{}) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetParentId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetRegionId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetSaleTag(v *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetScheduleInfo(v *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetStatus(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetTag(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetTenantId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetType(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) SetVersion(v string) *QueryQuotaResponseBodyDataSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoList) Validate() error {
	if s.BillingPolicy != nil {
		if err := s.BillingPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.SaleTag != nil {
		if err := s.SaleTag.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleInfo != nil {
		if err := s.ScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy struct {
	// The billing method. Valid values:
	//
	// 	- subscription: the subscription quota.
	//
	// 	- payasyougo: the pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// In MaxCompute, instanceId and orderId are considered the same.
	//
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The order specifications.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 880c0d0d-5d79-4217-b683-8e8bdd2a2523
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetInstanceId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.InstanceId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOrderId(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag struct {
	// The identifier of a MaxCompute quota object. This identifier exists in the Alibaba Cloud sales bill. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The object type. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceType(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo struct {
	// The current quota plan that has taken effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the plan specified by currPlan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the plan specified by nextPlan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan specified by this parameter is triggered and the plan is different from the plan specified by currPlan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan specified by oncePlan is scheduled.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The operator name.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetTimezone(v string) *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *QueryQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
