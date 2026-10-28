// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAlertHistoriesResponseBodyData) *ListAlertHistoriesResponseBody
	GetData() []*ListAlertHistoriesResponseBodyData
	SetMaxResults(v int32) *ListAlertHistoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertHistoriesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAlertHistoriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertHistoriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAlertHistoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlertHistoriesResponseBody
	GetTotalCount() *int32
}

type ListAlertHistoriesResponseBody struct {
	// The list of alert histories that match the query conditions.
	//
	// example:
	//
	// [{"alertHistoryId":"ah_123456","alertRuleId":"rule_789","status":"ALARM","startTime":1704067200000}]
	Data []*ListAlertHistoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of alert history records that match the query conditions.
	//
	// example:
	//
	// 81
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAlertHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertHistoriesResponseBody) GetData() []*ListAlertHistoriesResponseBodyData {
	return s.Data
}

func (s *ListAlertHistoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertHistoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertHistoriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertHistoriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertHistoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertHistoriesResponseBody) SetData(v []*ListAlertHistoriesResponseBodyData) *ListAlertHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetMaxResults(v int32) *ListAlertHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetNextToken(v string) *ListAlertHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetPageNumber(v int32) *ListAlertHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetPageSize(v int32) *ListAlertHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetRequestId(v string) *ListAlertHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) SetTotalCount(v int32) *ListAlertHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlertHistoriesResponseBody) Validate() error {
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

type ListAlertHistoriesResponseBodyData struct {
	// The unique identifier of the alert history.
	//
	// example:
	//
	// u1oe4i7c197uucfmvscum4u797
	AlertHistoryId *string `json:"alertHistoryId,omitempty" xml:"alertHistoryId,omitempty"`
	// The unique identifier of the alert rule that generated the alert history.
	//
	// example:
	//
	// 51d0da46-bb65-4a14-8d65-de73f41e495f
	AlertRuleId *string `json:"alertRuleId,omitempty" xml:"alertRuleId,omitempty"`
	// The snapshot of the annotations of the alert rule. The annotations are stored as key-value pairs and copied from the rule when the alert is triggered.
	//
	// example:
	//
	// {"_cms_alert_explore_url":"https://sls.console.aliyun.com/..."}
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// The business source.
	//
	// example:
	//
	// managed_service_for_prometheus
	BizSource *string `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	// The cumulative number of alerts (alert events) during the alert duration.
	//
	// example:
	//
	// 200
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The end time (recovery time) of the alert. This field has a value only when the alert history enters the OK state. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1740535517000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The creation time of the alert history record. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1751608363000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time of the alert history record. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1751608363000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The detailed information of the alert object. The information is stored as key-value pairs.
	//
	// example:
	//
	// {"datasourceInfo":{"type":"SLS_MULTI_DS"}}
	InstanceDetail map[string]*string `json:"instanceDetail,omitempty" xml:"instanceDetail,omitempty"`
	// The unique key (instanceKey) used to identify the alert object instance.
	//
	// example:
	//
	// entity/apm/apm.service/6dc9e46fda6d012537b48c3250ff806b
	InstanceKey *string `json:"instanceKey,omitempty" xml:"instanceKey,omitempty"`
	// The snapshot of the labels of the alert rule. The labels are stored as key-value pairs and copied from the rule when the alert is triggered.
	//
	// example:
	//
	// {"_cms_rule_id":"42c0f464-a3b0-4c4b-b426-6242289244f8","tenant_id":"mygt"}
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The level of the most recent alert during the alert duration, such as CRITICAL, WARN, or INFO.
	//
	// example:
	//
	// CRITICAL
	LatestLevel *string `json:"latestLevel,omitempty" xml:"latestLevel,omitempty"`
	// The highest alert level reached during the alert duration.
	//
	// example:
	//
	// CRITICAL
	MaxLevel *string `json:"maxLevel,omitempty" xml:"maxLevel,omitempty"`
	// The content of the latest alert message.
	//
	// example:
	//
	// The number of zombie child processes of the process {{$tags.comm}} running on the instance {{$tags.instanceId}} is greater than 128
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the observed object is in the global scope. The value is copied from the rule when the alert history is created.
	//
	// example:
	//
	// false
	ObserveResourceGlobalScope *bool `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	// The serialized string of the observed object list. The value is copied from the rule when the alert history is created.
	//
	// example:
	//
	// [{"resourceId":"i-bp1234567890abcdef","resourceType":"ACS::ECS::Instance"}]
	ObserveResourceList *string `json:"observeResourceList,omitempty" xml:"observeResourceList,omitempty"`
	// The type of the observed object, such as PROMETHEUS, APM, RUM, or cms_basic. The value is copied from the rule when the alert history is created.
	//
	// example:
	//
	// PROMETHEUS
	ObserveResourceType *string `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	// The ID of the region to which the alert history belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The display name of the corresponding rule.
	//
	// example:
	//
	// myTestRule
	RuleDisplayName *string `json:"ruleDisplayName,omitempty" xml:"ruleDisplayName,omitempty"`
	// The notification sending configuration information corresponding to the alert history.
	//
	// example:
	//
	// {"notification":{"contacts":["ops-team"],"silenceTime":3600}}
	Send *AlertRuleSend `json:"send,omitempty" xml:"send,omitempty"`
	// The start time of the alert history. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1751508233000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The current alert status of the rule. Valid values: OK (recovered) and ALARM (alerting).
	//
	// example:
	//
	// ALARM
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The identifier of the workspace to which the alert history belongs.
	//
	// example:
	//
	// default-cms-<userId>-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertHistoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlertHistoriesResponseBodyData) GetAlertHistoryId() *string {
	return s.AlertHistoryId
}

func (s *ListAlertHistoriesResponseBodyData) GetAlertRuleId() *string {
	return s.AlertRuleId
}

func (s *ListAlertHistoriesResponseBodyData) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *ListAlertHistoriesResponseBodyData) GetBizSource() *string {
	return s.BizSource
}

func (s *ListAlertHistoriesResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *ListAlertHistoriesResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAlertHistoriesResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListAlertHistoriesResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListAlertHistoriesResponseBodyData) GetInstanceDetail() map[string]*string {
	return s.InstanceDetail
}

func (s *ListAlertHistoriesResponseBodyData) GetInstanceKey() *string {
	return s.InstanceKey
}

func (s *ListAlertHistoriesResponseBodyData) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListAlertHistoriesResponseBodyData) GetLatestLevel() *string {
	return s.LatestLevel
}

func (s *ListAlertHistoriesResponseBodyData) GetMaxLevel() *string {
	return s.MaxLevel
}

func (s *ListAlertHistoriesResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListAlertHistoriesResponseBodyData) GetObserveResourceGlobalScope() *bool {
	return s.ObserveResourceGlobalScope
}

func (s *ListAlertHistoriesResponseBodyData) GetObserveResourceList() *string {
	return s.ObserveResourceList
}

func (s *ListAlertHistoriesResponseBodyData) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *ListAlertHistoriesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertHistoriesResponseBodyData) GetRuleDisplayName() *string {
	return s.RuleDisplayName
}

func (s *ListAlertHistoriesResponseBodyData) GetSend() *AlertRuleSend {
	return s.Send
}

func (s *ListAlertHistoriesResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAlertHistoriesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAlertHistoriesResponseBodyData) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertHistoriesResponseBodyData) SetAlertHistoryId(v string) *ListAlertHistoriesResponseBodyData {
	s.AlertHistoryId = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetAlertRuleId(v string) *ListAlertHistoriesResponseBodyData {
	s.AlertRuleId = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetAnnotations(v map[string]*string) *ListAlertHistoriesResponseBodyData {
	s.Annotations = v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetBizSource(v string) *ListAlertHistoriesResponseBodyData {
	s.BizSource = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetCount(v int64) *ListAlertHistoriesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetEndTime(v int64) *ListAlertHistoriesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetGmtCreate(v int64) *ListAlertHistoriesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetGmtModified(v int64) *ListAlertHistoriesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetInstanceDetail(v map[string]*string) *ListAlertHistoriesResponseBodyData {
	s.InstanceDetail = v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetInstanceKey(v string) *ListAlertHistoriesResponseBodyData {
	s.InstanceKey = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetLabels(v map[string]*string) *ListAlertHistoriesResponseBodyData {
	s.Labels = v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetLatestLevel(v string) *ListAlertHistoriesResponseBodyData {
	s.LatestLevel = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetMaxLevel(v string) *ListAlertHistoriesResponseBodyData {
	s.MaxLevel = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetMessage(v string) *ListAlertHistoriesResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetObserveResourceGlobalScope(v bool) *ListAlertHistoriesResponseBodyData {
	s.ObserveResourceGlobalScope = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetObserveResourceList(v string) *ListAlertHistoriesResponseBodyData {
	s.ObserveResourceList = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetObserveResourceType(v string) *ListAlertHistoriesResponseBodyData {
	s.ObserveResourceType = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetRegionId(v string) *ListAlertHistoriesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetRuleDisplayName(v string) *ListAlertHistoriesResponseBodyData {
	s.RuleDisplayName = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetSend(v *AlertRuleSend) *ListAlertHistoriesResponseBodyData {
	s.Send = v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetStartTime(v int64) *ListAlertHistoriesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetStatus(v string) *ListAlertHistoriesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) SetWorkspace(v string) *ListAlertHistoriesResponseBodyData {
	s.Workspace = &v
	return s
}

func (s *ListAlertHistoriesResponseBodyData) Validate() error {
	if s.Send != nil {
		if err := s.Send.Validate(); err != nil {
			return err
		}
	}
	return nil
}
