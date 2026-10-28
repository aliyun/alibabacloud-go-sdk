// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertHistoryId(v string) *ListAlertHistoriesRequest
	GetAlertHistoryId() *string
	SetAlertRuleId(v string) *ListAlertHistoriesRequest
	GetAlertRuleId() *string
	SetBizSource(v string) *ListAlertHistoriesRequest
	GetBizSource() *string
	SetCmsNamespace(v string) *ListAlertHistoriesRequest
	GetCmsNamespace() *string
	SetCmsProductCategory(v string) *ListAlertHistoriesRequest
	GetCmsProductCategory() *string
	SetDisplayNameKeyWord(v string) *ListAlertHistoriesRequest
	GetDisplayNameKeyWord() *string
	SetInstanceKey(v string) *ListAlertHistoriesRequest
	GetInstanceKey() *string
	SetLabelFilter(v *AlertRuleLabelFilter) *ListAlertHistoriesRequest
	GetLabelFilter() *AlertRuleLabelFilter
	SetLatestLevel(v string) *ListAlertHistoriesRequest
	GetLatestLevel() *string
	SetMaxLevel(v string) *ListAlertHistoriesRequest
	GetMaxLevel() *string
	SetMaxResults(v int32) *ListAlertHistoriesRequest
	GetMaxResults() *int32
	SetMigrationStatus(v int32) *ListAlertHistoriesRequest
	GetMigrationStatus() *int32
	SetNextToken(v string) *ListAlertHistoriesRequest
	GetNextToken() *string
	SetNotificationFilter(v *AlertRuleNotificationFilter) *ListAlertHistoriesRequest
	GetNotificationFilter() *AlertRuleNotificationFilter
	SetObserveRelationType(v string) *ListAlertHistoriesRequest
	GetObserveRelationType() *string
	SetObserveResourceDomain(v string) *ListAlertHistoriesRequest
	GetObserveResourceDomain() *string
	SetObserveResourceGlobalScope(v bool) *ListAlertHistoriesRequest
	GetObserveResourceGlobalScope() *bool
	SetObserveResourceInstanceId(v string) *ListAlertHistoriesRequest
	GetObserveResourceInstanceId() *string
	SetObserveResourceType(v string) *ListAlertHistoriesRequest
	GetObserveResourceType() *string
	SetPageNumber(v int32) *ListAlertHistoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAlertHistoriesRequest
	GetPageSize() *int32
	SetPartitionKey(v string) *ListAlertHistoriesRequest
	GetPartitionKey() *string
	SetStartTimeFrom(v int64) *ListAlertHistoriesRequest
	GetStartTimeFrom() *int64
	SetStartTimeTo(v int64) *ListAlertHistoriesRequest
	GetStartTimeTo() *int64
	SetStatus(v string) *ListAlertHistoriesRequest
	GetStatus() *string
	SetWorkspace(v string) *ListAlertHistoriesRequest
	GetWorkspace() *string
}

type ListAlertHistoriesRequest struct {
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
	LabelFilter *AlertRuleLabelFilter `json:"labelFilter,omitempty" xml:"labelFilter,omitempty"`
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
	NotificationFilter *AlertRuleNotificationFilter `json:"notificationFilter,omitempty" xml:"notificationFilter,omitempty"`
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

func (s ListAlertHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertHistoriesRequest) GetAlertHistoryId() *string {
	return s.AlertHistoryId
}

func (s *ListAlertHistoriesRequest) GetAlertRuleId() *string {
	return s.AlertRuleId
}

func (s *ListAlertHistoriesRequest) GetBizSource() *string {
	return s.BizSource
}

func (s *ListAlertHistoriesRequest) GetCmsNamespace() *string {
	return s.CmsNamespace
}

func (s *ListAlertHistoriesRequest) GetCmsProductCategory() *string {
	return s.CmsProductCategory
}

func (s *ListAlertHistoriesRequest) GetDisplayNameKeyWord() *string {
	return s.DisplayNameKeyWord
}

func (s *ListAlertHistoriesRequest) GetInstanceKey() *string {
	return s.InstanceKey
}

func (s *ListAlertHistoriesRequest) GetLabelFilter() *AlertRuleLabelFilter {
	return s.LabelFilter
}

func (s *ListAlertHistoriesRequest) GetLatestLevel() *string {
	return s.LatestLevel
}

func (s *ListAlertHistoriesRequest) GetMaxLevel() *string {
	return s.MaxLevel
}

func (s *ListAlertHistoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertHistoriesRequest) GetMigrationStatus() *int32 {
	return s.MigrationStatus
}

func (s *ListAlertHistoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertHistoriesRequest) GetNotificationFilter() *AlertRuleNotificationFilter {
	return s.NotificationFilter
}

func (s *ListAlertHistoriesRequest) GetObserveRelationType() *string {
	return s.ObserveRelationType
}

func (s *ListAlertHistoriesRequest) GetObserveResourceDomain() *string {
	return s.ObserveResourceDomain
}

func (s *ListAlertHistoriesRequest) GetObserveResourceGlobalScope() *bool {
	return s.ObserveResourceGlobalScope
}

func (s *ListAlertHistoriesRequest) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *ListAlertHistoriesRequest) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *ListAlertHistoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertHistoriesRequest) GetPartitionKey() *string {
	return s.PartitionKey
}

func (s *ListAlertHistoriesRequest) GetStartTimeFrom() *int64 {
	return s.StartTimeFrom
}

func (s *ListAlertHistoriesRequest) GetStartTimeTo() *int64 {
	return s.StartTimeTo
}

func (s *ListAlertHistoriesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAlertHistoriesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertHistoriesRequest) SetAlertHistoryId(v string) *ListAlertHistoriesRequest {
	s.AlertHistoryId = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetAlertRuleId(v string) *ListAlertHistoriesRequest {
	s.AlertRuleId = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetBizSource(v string) *ListAlertHistoriesRequest {
	s.BizSource = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetCmsNamespace(v string) *ListAlertHistoriesRequest {
	s.CmsNamespace = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetCmsProductCategory(v string) *ListAlertHistoriesRequest {
	s.CmsProductCategory = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetDisplayNameKeyWord(v string) *ListAlertHistoriesRequest {
	s.DisplayNameKeyWord = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetInstanceKey(v string) *ListAlertHistoriesRequest {
	s.InstanceKey = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetLabelFilter(v *AlertRuleLabelFilter) *ListAlertHistoriesRequest {
	s.LabelFilter = v
	return s
}

func (s *ListAlertHistoriesRequest) SetLatestLevel(v string) *ListAlertHistoriesRequest {
	s.LatestLevel = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetMaxLevel(v string) *ListAlertHistoriesRequest {
	s.MaxLevel = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetMaxResults(v int32) *ListAlertHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetMigrationStatus(v int32) *ListAlertHistoriesRequest {
	s.MigrationStatus = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetNextToken(v string) *ListAlertHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetNotificationFilter(v *AlertRuleNotificationFilter) *ListAlertHistoriesRequest {
	s.NotificationFilter = v
	return s
}

func (s *ListAlertHistoriesRequest) SetObserveRelationType(v string) *ListAlertHistoriesRequest {
	s.ObserveRelationType = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetObserveResourceDomain(v string) *ListAlertHistoriesRequest {
	s.ObserveResourceDomain = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetObserveResourceGlobalScope(v bool) *ListAlertHistoriesRequest {
	s.ObserveResourceGlobalScope = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetObserveResourceInstanceId(v string) *ListAlertHistoriesRequest {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetObserveResourceType(v string) *ListAlertHistoriesRequest {
	s.ObserveResourceType = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetPageNumber(v int32) *ListAlertHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetPageSize(v int32) *ListAlertHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetPartitionKey(v string) *ListAlertHistoriesRequest {
	s.PartitionKey = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetStartTimeFrom(v int64) *ListAlertHistoriesRequest {
	s.StartTimeFrom = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetStartTimeTo(v int64) *ListAlertHistoriesRequest {
	s.StartTimeTo = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetStatus(v string) *ListAlertHistoriesRequest {
	s.Status = &v
	return s
}

func (s *ListAlertHistoriesRequest) SetWorkspace(v string) *ListAlertHistoriesRequest {
	s.Workspace = &v
	return s
}

func (s *ListAlertHistoriesRequest) Validate() error {
	if s.LabelFilter != nil {
		if err := s.LabelFilter.Validate(); err != nil {
			return err
		}
	}
	if s.NotificationFilter != nil {
		if err := s.NotificationFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}
