// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQuotasPlansResponseBodyData) *ListQuotasPlansResponseBody
	GetData() *ListQuotasPlansResponseBodyData
	SetRequestId(v string) *ListQuotasPlansResponseBody
	GetRequestId() *string
}

type ListQuotasPlansResponseBody struct {
	// The returned data.
	Data *ListQuotasPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0bd16661643917136451ebf55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBody) GetData() *ListQuotasPlansResponseBodyData {
	return s.Data
}

func (s *ListQuotasPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotasPlansResponseBody) SetData(v *ListQuotasPlansResponseBodyData) *ListQuotasPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasPlansResponseBody) SetRequestId(v string) *ListQuotasPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasPlansResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyData struct {
	// The list of quota plans.
	PlanList []*ListQuotasPlansResponseBodyDataPlanList `json:"planList,omitempty" xml:"planList,omitempty" type:"Repeated"`
}

func (s ListQuotasPlansResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyData) GetPlanList() []*ListQuotasPlansResponseBodyDataPlanList {
	return s.PlanList
}

func (s *ListQuotasPlansResponseBodyData) SetPlanList(v []*ListQuotasPlansResponseBodyDataPlanList) *ListQuotasPlansResponseBodyData {
	s.PlanList = v
	return s
}

func (s *ListQuotasPlansResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanList struct {
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-05-16T06:07:45Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *ListQuotasPlansResponseBodyDataPlanListQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s ListQuotasPlansResponseBodyDataPlanList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQuotasPlansResponseBodyDataPlanList) GetName() *string {
	return s.Name
}

func (s *ListQuotasPlansResponseBodyDataPlanList) GetQuota() *ListQuotasPlansResponseBodyDataPlanListQuota {
	return s.Quota
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetCreateTime(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetQuota(v *ListQuotasPlansResponseBodyDataPlanListQuota) *ListQuotasPlansResponseBodyDataPlanList {
	s.Quota = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuota struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// 0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	//
	// example:
	//
	// quota_a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	//
	// example:
	//
	// quota_nickname
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// null
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
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

func (s ListQuotasPlansResponseBodyDataPlanListQuota) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuota) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetBillingPolicy() *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetId() *string {
	return s.Id
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetName() *string {
	return s.Name
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetScheduleInfo() *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetSubQuotaInfoList() []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetType() *string {
	return s.Type
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetSubQuotaInfoList(v []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Version = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
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
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	//
	// example:
	//
	// 0
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	//
	// example:
	//
	// abc
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
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

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetBillingPolicy() *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetScheduleInfo() *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// 	- subscription: a subscription quota.
	//
	// 	- payasyougo: a pay-as-you-go quota.
	//
	// example:
	//
	// subscription
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	//
	// example:
	//
	// OdpsStandard
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	//
	// example:
	//
	// planA
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	//
	// example:
	//
	// 0800
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	//
	// example:
	//
	// planB
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	//
	// example:
	//
	// 1700
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	//
	// example:
	//
	// planC
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	//
	// example:
	//
	// 1500
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// userA
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
