// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListQuotasResponseBody
	GetNextToken() *string
	SetData(v *ListQuotasResponseBodyData) *ListQuotasResponseBody
	GetData() *ListQuotasResponseBodyData
	SetMarker(v string) *ListQuotasResponseBody
	GetMarker() *string
	SetMaxItem(v int64) *ListQuotasResponseBody
	GetMaxItem() *int64
	SetQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoList) *ListQuotasResponseBody
	GetQuotaInfoList() []*ListQuotasResponseBodyQuotaInfoList
	SetRequestId(v string) *ListQuotasResponseBody
	GetRequestId() *string
}

type ListQuotasResponseBody struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2w6Olxc+cMPjUtUMo/CvPe4IK7f7kIQFrIZjyc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The returned data.
	Data *ListQuotasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc12e6f16677875480593081d2956
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotasResponseBody) GetData() *ListQuotasResponseBodyData {
	return s.Data
}

func (s *ListQuotasResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListQuotasResponseBody) GetMaxItem() *int64 {
	return s.MaxItem
}

func (s *ListQuotasResponseBody) GetQuotaInfoList() []*ListQuotasResponseBodyQuotaInfoList {
	return s.QuotaInfoList
}

func (s *ListQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotasResponseBody) SetNextToken(v string) *ListQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBody) SetData(v *ListQuotasResponseBodyData) *ListQuotasResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasResponseBody) SetMarker(v string) *ListQuotasResponseBody {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBody) SetMaxItem(v int64) *ListQuotasResponseBody {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBody) SetQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoList) *ListQuotasResponseBody {
	s.QuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// "abcde"
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotasResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListQuotasResponseBodyData) GetMaxItem() *int64 {
	return s.MaxItem
}

func (s *ListQuotasResponseBodyData) GetQuotaInfoList() []*ListQuotasResponseBodyDataQuotaInfoList {
	return s.QuotaInfoList
}

func (s *ListQuotasResponseBodyData) SetNextToken(v string) *ListQuotasResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMarker(v string) *ListQuotasResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMaxItem(v int64) *ListQuotasResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoList) *ListQuotasResponseBodyData {
	s.QuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyDataQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
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
	// 280747109771520
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

func (s ListQuotasResponseBodyDataQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetTags() []*ListQuotasResponseBodyDataQuotaInfoListTags {
	return s.Tags
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetBillingPolicy() *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetSaleTag() *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetScheduleInfo() *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetSubQuotaInfoList() []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTags(v []*ListQuotasResponseBodyDataQuotaInfoListTags) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Department
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// a12351qHDP6YEQMt
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagValue = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListBillingPolicy struct {
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

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListScheduleInfo struct {
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

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	Parameter *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
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
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
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
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version of the algorithm image.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetBillingPolicy() *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetParameter() *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetSaleTag() *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetScheduleInfo() *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParameter(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy struct {
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

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter struct {
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

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetEnablePriority(v bool) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetForceReservedMin(v bool) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetMaxCU(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetMinCU(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetSchedulerType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo struct {
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

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	SaleTag *ListQuotasResponseBodyQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	//
	// example:
	//
	// ON
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
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
	// 280747109771520
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	//
	// example:
	//
	// FUXI_ONLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetTags() []*ListQuotasResponseBodyQuotaInfoListTags {
	return s.Tags
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetBillingPolicy() *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetParameter() map[string]interface{} {
	return s.Parameter
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetSaleTag() *ListQuotasResponseBodyQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetScheduleInfo() *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetSubQuotaInfoList() []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListQuotasResponseBodyQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTags(v []*ListQuotasResponseBodyQuotaInfoListTags) *ListQuotasResponseBodyQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Department
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListTags) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagValue = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListBillingPolicy struct {
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

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "project"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListScheduleInfo struct {
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

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
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
	Parameter *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
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
	SaleTag *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
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
	// 280747109771520
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

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetBillingPolicy() *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	return s.BillingPolicy
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetId() *string {
	return s.Id
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetParameter() *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetParentId() *string {
	return s.ParentId
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetSaleTag() *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	return s.SaleTag
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetScheduleInfo() *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetTag() *string {
	return s.Tag
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GetVersion() *string {
	return s.Version
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParameter(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy struct {
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

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GetOdpsSpecCode() *string {
	return s.OdpsSpecCode
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GetOrderId() *string {
	return s.OrderId
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter struct {
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

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetEnablePriority(v bool) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetForceReservedMin(v bool) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetMaxCU(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetMinCU(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetSchedulerType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	//
	// example:
	//
	// "quota"
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo struct {
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

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetCurrPlan() *string {
	return s.CurrPlan
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetCurrTime() *string {
	return s.CurrTime
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetNextPlan() *string {
	return s.NextPlan
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetNextTime() *string {
	return s.NextTime
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetOncePlan() *string {
	return s.OncePlan
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetOnceTime() *string {
	return s.OnceTime
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GetTimezone() *string {
	return s.Timezone
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
