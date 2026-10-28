// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertHistoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertHistoryId(v string) *ListAlertHistoriesShrinkRequest
	GetAlertHistoryId() *string
	SetAlertRuleId(v string) *ListAlertHistoriesShrinkRequest
	GetAlertRuleId() *string
	SetBizSource(v string) *ListAlertHistoriesShrinkRequest
	GetBizSource() *string
	SetCmsNamespace(v string) *ListAlertHistoriesShrinkRequest
	GetCmsNamespace() *string
	SetCmsProductCategory(v string) *ListAlertHistoriesShrinkRequest
	GetCmsProductCategory() *string
	SetDisplayNameKeyWord(v string) *ListAlertHistoriesShrinkRequest
	GetDisplayNameKeyWord() *string
	SetInstanceKey(v string) *ListAlertHistoriesShrinkRequest
	GetInstanceKey() *string
	SetLabelFilterShrink(v string) *ListAlertHistoriesShrinkRequest
	GetLabelFilterShrink() *string
	SetLatestLevel(v string) *ListAlertHistoriesShrinkRequest
	GetLatestLevel() *string
	SetMaxLevel(v string) *ListAlertHistoriesShrinkRequest
	GetMaxLevel() *string
	SetMaxResults(v int32) *ListAlertHistoriesShrinkRequest
	GetMaxResults() *int32
	SetMigrationStatus(v int32) *ListAlertHistoriesShrinkRequest
	GetMigrationStatus() *int32
	SetNextToken(v string) *ListAlertHistoriesShrinkRequest
	GetNextToken() *string
	SetNotificationFilterShrink(v string) *ListAlertHistoriesShrinkRequest
	GetNotificationFilterShrink() *string
	SetObserveRelationType(v string) *ListAlertHistoriesShrinkRequest
	GetObserveRelationType() *string
	SetObserveResourceDomain(v string) *ListAlertHistoriesShrinkRequest
	GetObserveResourceDomain() *string
	SetObserveResourceGlobalScope(v bool) *ListAlertHistoriesShrinkRequest
	GetObserveResourceGlobalScope() *bool
	SetObserveResourceInstanceId(v string) *ListAlertHistoriesShrinkRequest
	GetObserveResourceInstanceId() *string
	SetObserveResourceType(v string) *ListAlertHistoriesShrinkRequest
	GetObserveResourceType() *string
	SetPageNumber(v int32) *ListAlertHistoriesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertHistoriesShrinkRequest
	GetPageSize() *int32
	SetPartitionKey(v string) *ListAlertHistoriesShrinkRequest
	GetPartitionKey() *string
	SetStartTimeFrom(v int64) *ListAlertHistoriesShrinkRequest
	GetStartTimeFrom() *int64
	SetStartTimeTo(v int64) *ListAlertHistoriesShrinkRequest
	GetStartTimeTo() *int64
	SetStatus(v string) *ListAlertHistoriesShrinkRequest
	GetStatus() *string
	SetWorkspace(v string) *ListAlertHistoriesShrinkRequest
	GetWorkspace() *string
}

type ListAlertHistoriesShrinkRequest struct {
	// The unique identifier (transId) of an alert history record.
	//
	// example:
	//
	// lme6hk5qih5o769m379jjuo5i1
	AlertHistoryId *string `json:"alertHistoryId,omitempty" xml:"alertHistoryId,omitempty"`
	// Filters the history records by a specific alert rule. Specify the unique identifier (ruleUuid) of the alert rule.
	//
	// example:
	//
	// mig_prom_v2_11308451
	AlertRuleId *string `json:"alertRuleId,omitempty" xml:"alertRuleId,omitempty"`
	// The business source of the alert history. For example, `managed_service_for_prometheus` indicates the managed type of the Prometheus integration center. The default value is empty. Leave this parameter empty if you are unsure about the value.
	//
	// example:
	//
	// managed_service_for_prometheus
	BizSource *string `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	// The namespace of CloudMonitor. Examples: acs_rds_dashboard and acs_ecs_dashboard.
	//
	// example:
	//
	// acs_rds_dashboard
	CmsNamespace *string `json:"cmsNamespace,omitempty" xml:"cmsNamespace,omitempty"`
	// The product category of CloudMonitor.
	//
	// example:
	//
	// rds
	CmsProductCategory *string `json:"cmsProductCategory,omitempty" xml:"cmsProductCategory,omitempty"`
	// The keyword used for fuzzy matching and filtering based on the display name of the alert rule.
	//
	// example:
	//
	// keyword
	DisplayNameKeyWord *string `json:"displayNameKeyWord,omitempty" xml:"displayNameKeyWord,omitempty"`
	// The unique key (instanceKey) used to identify the alert object instance.
	//
	// example:
	//
	// entity/apm/apm.service/6dc9e46fda6d012537b48c3250ff806b
	InstanceKey *string `json:"instanceKey,omitempty" xml:"instanceKey,omitempty"`
	// The filter condition for rule labels. The value is a JSON-formatted string that supports filtering by label key-value pairs.
	LabelFilterShrink *string `json:"labelFilter,omitempty" xml:"labelFilter,omitempty"`
	// The level of the most recent alert during the alert duration. Valid values: CRITICAL, WARN, and INFO.
	//
	// example:
	//
	// CRITICAL
	LatestLevel *string `json:"latestLevel,omitempty" xml:"latestLevel,omitempty"`
	// The highest alert level reached during the alert duration. Valid values: CRITICAL, WARN, and INFO.
	//
	// example:
	//
	// CRITICAL
	MaxLevel *string `json:"maxLevel,omitempty" xml:"maxLevel,omitempty"`
	// The maximum number of results to return on each page. This parameter is not supported. Use pageSize and pageNumber for pagination.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The migration status of the alert history.
	//
	// example:
	//
	// 0
	MigrationStatus *int32 `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// The pagination token used for the next query. This parameter is not supported. Use pageSize and pageNumber for pagination.
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters alert rules that contain specific notification recipients. Only rules that contain at least one of the notification recipients specified in notifyTargets are returned in the result list.
	NotificationFilterShrink *string `json:"notificationFilter,omitempty" xml:"notificationFilter,omitempty"`
	// The relation type of the observed object. Examples: GROUP, INSTANCE, and TAG.
	//
	// example:
	//
	// INSTANCE
	ObserveRelationType *string `json:"observeRelationType,omitempty" xml:"observeRelationType,omitempty"`
	// The dimension column (domain) of the observed object.
	//
	// example:
	//
	// acs
	ObserveResourceDomain *string `json:"observeResourceDomain,omitempty" xml:"observeResourceDomain,omitempty"`
	// Specifies whether the observed object is in the global scope.
	//
	// example:
	//
	// false
	ObserveResourceGlobalScope *bool `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	// Filters the alert history by resource instance ID.
	//
	// example:
	//
	// i-bp1234567890abcdef
	ObserveResourceInstanceId *string `json:"observeResourceInstanceId,omitempty" xml:"observeResourceInstanceId,omitempty"`
	// The type of the observed object. Examples: PROMETHEUS, APM, RUM, and cms_basic.
	//
	// example:
	//
	// PROMETHEUS
	ObserveResourceType *string `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of alert history records to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The key used for data partitioning. Typically, you do not need to specify this parameter.
	//
	// example:
	//
	// cn-hangzhou
	PartitionKey *string `json:"partitionKey,omitempty" xml:"partitionKey,omitempty"`
	// The start timestamp of the time range for the alert start time. The value is a UNIX timestamp in milliseconds. If not specified, the default value is the time 30 days ago.
	//
	// example:
	//
	// 1739866071668
	StartTimeFrom *int64 `json:"startTimeFrom,omitempty" xml:"startTimeFrom,omitempty"`
	// The end timestamp of the time range for the alert start time. The value is a UNIX timestamp in milliseconds. If not specified, the default value is the current time.
	//
	// example:
	//
	// 1739866071668
	StartTimeTo *int64 `json:"startTimeTo,omitempty" xml:"startTimeTo,omitempty"`
	// The current status of the alert history. Valid values: OK and ALARM.
	//
	// example:
	//
	// ALARM
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The identifier of the workspace to which the alert history belongs. If not specified, the default workspace is used.
	//
	// example:
	//
	// default-cms-<userId>-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertHistoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertHistoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertHistoriesShrinkRequest) GetAlertHistoryId() *string {
	return s.AlertHistoryId
}

func (s *ListAlertHistoriesShrinkRequest) GetAlertRuleId() *string {
	return s.AlertRuleId
}

func (s *ListAlertHistoriesShrinkRequest) GetBizSource() *string {
	return s.BizSource
}

func (s *ListAlertHistoriesShrinkRequest) GetCmsNamespace() *string {
	return s.CmsNamespace
}

func (s *ListAlertHistoriesShrinkRequest) GetCmsProductCategory() *string {
	return s.CmsProductCategory
}

func (s *ListAlertHistoriesShrinkRequest) GetDisplayNameKeyWord() *string {
	return s.DisplayNameKeyWord
}

func (s *ListAlertHistoriesShrinkRequest) GetInstanceKey() *string {
	return s.InstanceKey
}

func (s *ListAlertHistoriesShrinkRequest) GetLabelFilterShrink() *string {
	return s.LabelFilterShrink
}

func (s *ListAlertHistoriesShrinkRequest) GetLatestLevel() *string {
	return s.LatestLevel
}

func (s *ListAlertHistoriesShrinkRequest) GetMaxLevel() *string {
	return s.MaxLevel
}

func (s *ListAlertHistoriesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertHistoriesShrinkRequest) GetMigrationStatus() *int32 {
	return s.MigrationStatus
}

func (s *ListAlertHistoriesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertHistoriesShrinkRequest) GetNotificationFilterShrink() *string {
	return s.NotificationFilterShrink
}

func (s *ListAlertHistoriesShrinkRequest) GetObserveRelationType() *string {
	return s.ObserveRelationType
}

func (s *ListAlertHistoriesShrinkRequest) GetObserveResourceDomain() *string {
	return s.ObserveResourceDomain
}

func (s *ListAlertHistoriesShrinkRequest) GetObserveResourceGlobalScope() *bool {
	return s.ObserveResourceGlobalScope
}

func (s *ListAlertHistoriesShrinkRequest) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *ListAlertHistoriesShrinkRequest) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *ListAlertHistoriesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertHistoriesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertHistoriesShrinkRequest) GetPartitionKey() *string {
	return s.PartitionKey
}

func (s *ListAlertHistoriesShrinkRequest) GetStartTimeFrom() *int64 {
	return s.StartTimeFrom
}

func (s *ListAlertHistoriesShrinkRequest) GetStartTimeTo() *int64 {
	return s.StartTimeTo
}

func (s *ListAlertHistoriesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAlertHistoriesShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertHistoriesShrinkRequest) SetAlertHistoryId(v string) *ListAlertHistoriesShrinkRequest {
	s.AlertHistoryId = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetAlertRuleId(v string) *ListAlertHistoriesShrinkRequest {
	s.AlertRuleId = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetBizSource(v string) *ListAlertHistoriesShrinkRequest {
	s.BizSource = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetCmsNamespace(v string) *ListAlertHistoriesShrinkRequest {
	s.CmsNamespace = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetCmsProductCategory(v string) *ListAlertHistoriesShrinkRequest {
	s.CmsProductCategory = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetDisplayNameKeyWord(v string) *ListAlertHistoriesShrinkRequest {
	s.DisplayNameKeyWord = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetInstanceKey(v string) *ListAlertHistoriesShrinkRequest {
	s.InstanceKey = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetLabelFilterShrink(v string) *ListAlertHistoriesShrinkRequest {
	s.LabelFilterShrink = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetLatestLevel(v string) *ListAlertHistoriesShrinkRequest {
	s.LatestLevel = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetMaxLevel(v string) *ListAlertHistoriesShrinkRequest {
	s.MaxLevel = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetMaxResults(v int32) *ListAlertHistoriesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetMigrationStatus(v int32) *ListAlertHistoriesShrinkRequest {
	s.MigrationStatus = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetNextToken(v string) *ListAlertHistoriesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetNotificationFilterShrink(v string) *ListAlertHistoriesShrinkRequest {
	s.NotificationFilterShrink = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetObserveRelationType(v string) *ListAlertHistoriesShrinkRequest {
	s.ObserveRelationType = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetObserveResourceDomain(v string) *ListAlertHistoriesShrinkRequest {
	s.ObserveResourceDomain = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetObserveResourceGlobalScope(v bool) *ListAlertHistoriesShrinkRequest {
	s.ObserveResourceGlobalScope = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetObserveResourceInstanceId(v string) *ListAlertHistoriesShrinkRequest {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetObserveResourceType(v string) *ListAlertHistoriesShrinkRequest {
	s.ObserveResourceType = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetPageNumber(v int32) *ListAlertHistoriesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetPageSize(v int32) *ListAlertHistoriesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetPartitionKey(v string) *ListAlertHistoriesShrinkRequest {
	s.PartitionKey = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetStartTimeFrom(v int64) *ListAlertHistoriesShrinkRequest {
	s.StartTimeFrom = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetStartTimeTo(v int64) *ListAlertHistoriesShrinkRequest {
	s.StartTimeTo = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetStatus(v string) *ListAlertHistoriesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) SetWorkspace(v string) *ListAlertHistoriesShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListAlertHistoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
