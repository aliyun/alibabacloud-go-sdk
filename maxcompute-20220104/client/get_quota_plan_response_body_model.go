// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetQuotaPlanResponseBodyData) *GetQuotaPlanResponseBody
	GetData() *GetQuotaPlanResponseBodyData
	SetRequestId(v string) *GetQuotaPlanResponseBody
	GetRequestId() *string
}

type GetQuotaPlanResponseBody struct {
	// The returned data.
	Data *GetQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBody) GetData() *GetQuotaPlanResponseBodyData {
	return s.Data
}

func (s *GetQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaPlanResponseBody) SetData(v *GetQuotaPlanResponseBodyData) *GetQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaPlanResponseBody) SetRequestId(v string) *GetQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaPlanResponseBodyData struct {
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
	Quota *GetQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetQuotaPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQuotaPlanResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetQuotaPlanResponseBodyData) GetQuota() *GetQuotaPlanResponseBodyDataQuota {
	return s.Quota
}

func (s *GetQuotaPlanResponseBodyData) SetCreateTime(v string) *GetQuotaPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetName(v string) *GetQuotaPlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetQuota(v *GetQuotaPlanResponseBodyDataQuota) *GetQuotaPlanResponseBodyData {
	s.Quota = v
	return s
}

func (s *GetQuotaPlanResponseBodyData) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaPlanResponseBodyDataQuota struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
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

func (s GetQuotaPlanResponseBodyDataQuota) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetBillingPolicy() *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetId() *string {
	return s.Id
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetName() *string {
	return s.Name
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetScheduleInfo() *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetSubQuotaInfoList() []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetType() *string {
	return s.Type
}

func (s *GetQuotaPlanResponseBodyDataQuota) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) *GetQuotaPlanResponseBodyDataQuota {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) *GetQuotaPlanResponseBodyDataQuota {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) *GetQuotaPlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTag(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetType(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) Validate() error {
	if s.BillingPolicy != nil {
		if err := s.BillingPolicy.Validate(); err != nil {
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

type GetQuotaPlanResponseBodyDataQuotaBillingPolicy struct {
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

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaPlanResponseBodyDataQuotaScheduleInfo struct {
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

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	//
	// example:
	//
	// AT-120N
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the user who created the quota plan.
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
	// The alias of the level-2 quota.
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
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
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

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetBillingPolicy() *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetScheduleInfo() *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTag(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) Validate() error {
	if s.BillingPolicy != nil {
		if err := s.BillingPolicy.Validate(); err != nil {
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

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy struct {
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

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo struct {
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

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
