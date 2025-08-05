// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillingPolicy(v *GetQuotaResponseBodyBillingPolicy) *GetQuotaResponseBody
	GetBillingPolicy() *GetQuotaResponseBodyBillingPolicy
	SetCluster(v string) *GetQuotaResponseBody
	GetCluster() *string
	SetCreateTime(v int64) *GetQuotaResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *GetQuotaResponseBody
	GetCreatorId() *string
	SetData(v *GetQuotaResponseBodyData) *GetQuotaResponseBody
	GetData() *GetQuotaResponseBodyData
	SetId(v string) *GetQuotaResponseBody
	GetId() *string
	SetName(v string) *GetQuotaResponseBody
	GetName() *string
	SetNickName(v string) *GetQuotaResponseBody
	GetNickName() *string
	SetParameter(v map[string]interface{}) *GetQuotaResponseBody
	GetParameter() map[string]interface{}
	SetParentId(v string) *GetQuotaResponseBody
	GetParentId() *string
	SetRegionId(v string) *GetQuotaResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetQuotaResponseBody
	GetRequestId() *string
	SetSaleTag(v *GetQuotaResponseBodySaleTag) *GetQuotaResponseBody
	GetSaleTag() *GetQuotaResponseBodySaleTag
	SetScheduleInfo(v *GetQuotaResponseBodyScheduleInfo) *GetQuotaResponseBody
	GetScheduleInfo() *GetQuotaResponseBodyScheduleInfo
	SetStatus(v string) *GetQuotaResponseBody
	GetStatus() *string
	SetSubQuotaInfoList(v []*GetQuotaResponseBodySubQuotaInfoList) *GetQuotaResponseBody
	GetSubQuotaInfoList() []*GetQuotaResponseBodySubQuotaInfoList
	SetTag(v string) *GetQuotaResponseBody
	GetTag() *string
	SetTenantId(v string) *GetQuotaResponseBody
	GetTenantId() *string
	SetType(v string) *GetQuotaResponseBody
	GetType() *string
	SetVersion(v string) *GetQuotaResponseBody
	GetVersion() *string
}

type GetQuotaResponseBody struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The returned data.
	Data *GetQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The quota ID.
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
	// The ID of the parent resource.
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
	// The request ID.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodySaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information about the level-2 quota.
	SubQuotaInfoList []*GetQuotaResponseBodySubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
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

func (s GetQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) GetBillingPolicy() *GetQuotaResponseBodyBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaResponseBody) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaResponseBody) GetData() *GetQuotaResponseBodyData {
	return s.Data
}

func (s *GetQuotaResponseBody) GetId() *string {
	return s.Id
}

func (s *GetQuotaResponseBody) GetName() *string {
	return s.Name
}

func (s *GetQuotaResponseBody) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaResponseBody) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetQuotaResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaResponseBody) GetSaleTag() *GetQuotaResponseBodySaleTag {
	return s.SaleTag
}

func (s *GetQuotaResponseBody) GetScheduleInfo() *GetQuotaResponseBodyScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaResponseBody) GetSubQuotaInfoList() []*GetQuotaResponseBodySubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *GetQuotaResponseBody) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaResponseBody) GetType() *string {
	return s.Type
}

func (s *GetQuotaResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaResponseBody) SetBillingPolicy(v *GetQuotaResponseBodyBillingPolicy) *GetQuotaResponseBody {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBody) SetCluster(v string) *GetQuotaResponseBody {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreateTime(v int64) *GetQuotaResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreatorId(v string) *GetQuotaResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBody) SetData(v *GetQuotaResponseBodyData) *GetQuotaResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaResponseBody) SetId(v string) *GetQuotaResponseBody {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBody) SetName(v string) *GetQuotaResponseBody {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBody) SetNickName(v string) *GetQuotaResponseBody {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBody) SetParameter(v map[string]interface{}) *GetQuotaResponseBody {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBody) SetParentId(v string) *GetQuotaResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRegionId(v string) *GetQuotaResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetSaleTag(v *GetQuotaResponseBodySaleTag) *GetQuotaResponseBody {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBody) SetScheduleInfo(v *GetQuotaResponseBodyScheduleInfo) *GetQuotaResponseBody {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBody) SetStatus(v string) *GetQuotaResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBody) SetSubQuotaInfoList(v []*GetQuotaResponseBodySubQuotaInfoList) *GetQuotaResponseBody {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBody) SetTag(v string) *GetQuotaResponseBody {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBody) SetTenantId(v string) *GetQuotaResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBody) SetType(v string) *GetQuotaResponseBody {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBody) SetVersion(v string) *GetQuotaResponseBody {
	s.Version = &v
	return s
}

func (s *GetQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyBillingPolicy struct {
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
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaResponseBodyBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaResponseBodyBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaResponseBodyBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyData struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyDataBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	// 2022-09-06T02:14:44Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	//
	// example:
	//
	// 672863518
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
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
	// The ID of the parent resource.
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
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyDataScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information about the level-2 quota.
	SubQuotaInfoList []*GetQuotaResponseBodyDataSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
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

func (s GetQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyData) GetBillingPolicy() *GetQuotaResponseBodyDataBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaResponseBodyData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetQuotaResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetQuotaResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaResponseBodyData) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *GetQuotaResponseBodyData) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaResponseBodyData) GetSaleTag() *GetQuotaResponseBodyDataSaleTag {
	return s.SaleTag
}

func (s *GetQuotaResponseBodyData) GetScheduleInfo() *GetQuotaResponseBodyDataScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaResponseBodyData) GetSubQuotaInfoList() []*GetQuotaResponseBodyDataSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *GetQuotaResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetQuotaResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaResponseBodyData) SetBillingPolicy(v *GetQuotaResponseBodyDataBillingPolicy) *GetQuotaResponseBodyData {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyData) SetCluster(v string) *GetQuotaResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreateTime(v int64) *GetQuotaResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreatorId(v string) *GetQuotaResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetId(v string) *GetQuotaResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetName(v string) *GetQuotaResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetNickName(v string) *GetQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetParameter(v map[string]interface{}) *GetQuotaResponseBodyData {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyData) SetParentId(v string) *GetQuotaResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetRegionId(v string) *GetQuotaResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSaleTag(v *GetQuotaResponseBodyDataSaleTag) *GetQuotaResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyData) SetScheduleInfo(v *GetQuotaResponseBodyDataScheduleInfo) *GetQuotaResponseBodyData {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyData) SetStatus(v string) *GetQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSubQuotaInfoList(v []*GetQuotaResponseBodyDataSubQuotaInfoList) *GetQuotaResponseBodyData {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBodyData) SetTag(v string) *GetQuotaResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetTenantId(v string) *GetQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetType(v string) *GetQuotaResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetVersion(v string) *GetQuotaResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataBillingPolicy struct {
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
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaResponseBodyDataBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaResponseBodyDataBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSaleTag) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *GetQuotaResponseBodyDataSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBodyDataSaleTag) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataScheduleInfo struct {
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
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
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
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaResponseBodyDataScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSubQuotaInfoList struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	Parameter *GetQuotaResponseBodyDataSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The ID of the parent resource.
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
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
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
	// The tenant ID.
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

func (s GetQuotaResponseBodyDataSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetBillingPolicy() *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetParameter() *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetSaleTag() *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetScheduleInfo() *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParameter(v *GetQuotaResponseBodyDataSubQuotaInfoListParameter) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetType(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy struct {
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
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSubQuotaInfoListParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	EnablePriority    *bool  `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	ForceReservedMin  *bool  `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// This parameter is required.
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// This parameter is required.
	MinCU            *int64  `json:"minCU,omitempty" xml:"minCU,omitempty"`
	SchedulerType    *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	SingleJobCULimit *int64  `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetEnablePriority(v bool) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetForceReservedMin(v bool) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetMaxCU(v int64) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetMinCU(v int64) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetSchedulerType(v string) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *GetQuotaResponseBodyDataSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo struct {
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
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
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
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySaleTag) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *GetQuotaResponseBodySaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetQuotaResponseBodySaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySaleTag) SetResourceType(v string) *GetQuotaResponseBodySaleTag {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBodySaleTag) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodyScheduleInfo struct {
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
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
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
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodyScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaResponseBodyScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaResponseBodyScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaResponseBodyScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaResponseBodyScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaResponseBodyScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaResponseBodyScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaResponseBodyScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySubQuotaInfoList struct {
	// The information about the order.
	BillingPolicy *GetQuotaResponseBodySubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	// The alias of the level-2 quota.
	//
	// example:
	//
	// subquotaA
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
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
	Parameter *GetQuotaResponseBodySubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The ID of the parent resource.
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
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *GetQuotaResponseBodySubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information about the scheduling plan.
	ScheduleInfo *GetQuotaResponseBodySubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
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
	// The tenant ID.
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

func (s GetQuotaResponseBodySubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetBillingPolicy() *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetParameter() *GetQuotaResponseBodySubQuotaInfoListParameter {
	return s.Parameter
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetSaleTag() *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetScheduleInfo() *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *GetQuotaResponseBodySubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) *GetQuotaResponseBodySubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParameter(v *GetQuotaResponseBodySubQuotaInfoListParameter) *GetQuotaResponseBodySubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodySubQuotaInfoListSaleTag) *GetQuotaResponseBodySubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) *GetQuotaResponseBodySubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetType(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySubQuotaInfoListBillingPolicy struct {
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
	// The order ID.
	//
	// example:
	//
	// 45245678
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySubQuotaInfoListParameter struct {
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	EnablePriority    *bool  `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	ForceReservedMin  *bool  `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// This parameter is required.
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// This parameter is required.
	MinCU            *int64  `json:"minCU,omitempty" xml:"minCU,omitempty"`
	SchedulerType    *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	SingleJobCULimit *int64  `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetElasticReservedCU(v int64) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetEnablePriority(v bool) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetForceReservedMin(v bool) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetMaxCU(v int64) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetMinCU(v int64) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetSchedulerType(v string) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *GetQuotaResponseBodySubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type GetQuotaResponseBodySubQuotaInfoListScheduleInfo struct {
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
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
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
	// The time zone of the project.
	//
	// example:
	//
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
