// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	roa "github.com/alibabacloud-go/tea-roa/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateContactGroupForAlertRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s UpdateContactGroupForAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactGroupForAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupForAlertRequest) SetHeaders(v map[string]*string) *UpdateContactGroupForAlertRequest {
	s.Headers = v
	return s
}

type UpdateContactGroupForAlertResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateContactGroupForAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactGroupForAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupForAlertResponse) SetHeaders(v map[string]*string) *UpdateContactGroupForAlertResponse {
	s.Headers = v
	return s
}

type StopAlertRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s StopAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAlertRequest) GoString() string {
	return s.String()
}

func (s *StopAlertRequest) SetHeaders(v map[string]*string) *StopAlertRequest {
	s.Headers = v
	return s
}

type StopAlertResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s StopAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAlertResponse) GoString() string {
	return s.String()
}

func (s *StopAlertResponse) SetHeaders(v map[string]*string) *StopAlertResponse {
	s.Headers = v
	return s
}

type StartAlertRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s StartAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) SetHeaders(v map[string]*string) *StartAlertRequest {
	s.Headers = v
	return s
}

type StartAlertResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s StartAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAlertResponse) GoString() string {
	return s.String()
}

func (s *StartAlertResponse) SetHeaders(v map[string]*string) *StartAlertResponse {
	s.Headers = v
	return s
}

type DeleteAlertContactRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) SetHeaders(v map[string]*string) *DeleteAlertContactRequest {
	s.Headers = v
	return s
}

type DeleteAlertContactResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponse) SetHeaders(v map[string]*string) *DeleteAlertContactResponse {
	s.Headers = v
	return s
}

type DeleteAlertContactGroupRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) SetHeaders(v map[string]*string) *DeleteAlertContactGroupRequest {
	s.Headers = v
	return s
}

type DeleteAlertContactGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponse) SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse {
	s.Headers = v
	return s
}

type DescribeEventsQuery struct {
	ClusterId  *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	PageSize   *int64  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PageNumber *int64  `json:"page_number,omitempty" xml:"page_number,omitempty"`
}

func (s DescribeEventsQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsQuery) GoString() string {
	return s.String()
}

func (s *DescribeEventsQuery) SetClusterId(v string) *DescribeEventsQuery {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsQuery) SetType(v string) *DescribeEventsQuery {
	s.Type = &v
	return s
}

func (s *DescribeEventsQuery) SetPageSize(v int64) *DescribeEventsQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsQuery) SetPageNumber(v int64) *DescribeEventsQuery {
	s.PageNumber = &v
	return s
}

type DescribeEventsRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeEventsQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetHeaders(v map[string]*string) *DescribeEventsRequest {
	s.Headers = v
	return s
}

func (s *DescribeEventsRequest) SetQuery(v *DescribeEventsQuery) *DescribeEventsRequest {
	s.Query = v
	return s
}

type DescribeEventsResponseBody struct {
	Events   []*DescribeEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" require:"true" type:"Repeated"`
	PageInfo *DescribeEventsResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetEvents(v []*DescribeEventsResponseBodyEvents) *DescribeEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageInfo(v *DescribeEventsResponseBodyPageInfo) *DescribeEventsResponseBody {
	s.PageInfo = v
	return s
}

type DescribeEventsResponseBodyEvents struct {
	ClusterId *string                               `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	EventId   *string                               `json:"event_id,omitempty" xml:"event_id,omitempty" require:"true"`
	Subject   *string                               `json:"subject,omitempty" xml:"subject,omitempty" require:"true"`
	Time      *string                               `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	Source    *string                               `json:"source,omitempty" xml:"source,omitempty" require:"true"`
	Type      *string                               `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Data      *DescribeEventsResponseBodyEventsData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEvents) SetClusterId(v string) *DescribeEventsResponseBodyEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetEventId(v string) *DescribeEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetSubject(v string) *DescribeEventsResponseBodyEvents {
	s.Subject = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetTime(v string) *DescribeEventsResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetSource(v string) *DescribeEventsResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetType(v string) *DescribeEventsResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *DescribeEventsResponseBodyEvents) SetData(v *DescribeEventsResponseBodyEventsData) *DescribeEventsResponseBodyEvents {
	s.Data = v
	return s
}

type DescribeEventsResponseBodyEventsData struct {
	Reason  *string `json:"reason,omitempty" xml:"reason,omitempty" require:"true"`
	Level   *string `json:"level,omitempty" xml:"level,omitempty" require:"true"`
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DescribeEventsResponseBodyEventsData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyEventsData) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventsData) SetReason(v string) *DescribeEventsResponseBodyEventsData {
	s.Reason = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsData) SetLevel(v string) *DescribeEventsResponseBodyEventsData {
	s.Level = &v
	return s
}

func (s *DescribeEventsResponseBodyEventsData) SetMessage(v string) *DescribeEventsResponseBodyEventsData {
	s.Message = &v
	return s
}

type DescribeEventsResponseBodyPageInfo struct {
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty" require:"true"`
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	PageSize   *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s DescribeEventsResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyPageInfo) SetPageNumber(v int64) *DescribeEventsResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBodyPageInfo) SetTotalCount(v int64) *DescribeEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBodyPageInfo) SetPageSize(v int64) *DescribeEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

type DescribeEventsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type MigrateClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s MigrateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateClusterRequest) GoString() string {
	return s.String()
}

func (s *MigrateClusterRequest) SetHeaders(v map[string]*string) *MigrateClusterRequest {
	s.Headers = v
	return s
}

type MigrateClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s MigrateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateClusterResponse) GoString() string {
	return s.String()
}

func (s *MigrateClusterResponse) SetHeaders(v map[string]*string) *MigrateClusterResponse {
	s.Headers = v
	return s
}

type OpenAckServiceQuery struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OpenAckServiceQuery) String() string {
	return tea.Prettify(s)
}

func (s OpenAckServiceQuery) GoString() string {
	return s.String()
}

func (s *OpenAckServiceQuery) SetType(v string) *OpenAckServiceQuery {
	s.Type = &v
	return s
}

type OpenAckServiceRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *OpenAckServiceQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s OpenAckServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenAckServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAckServiceRequest) SetHeaders(v map[string]*string) *OpenAckServiceRequest {
	s.Headers = v
	return s
}

func (s *OpenAckServiceRequest) SetQuery(v *OpenAckServiceQuery) *OpenAckServiceRequest {
	s.Query = v
	return s
}

type OpenAckServiceResponseBody struct {
	OrderId   *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
}

func (s OpenAckServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenAckServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAckServiceResponseBody) SetOrderId(v string) *OpenAckServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAckServiceResponseBody) SetRequestId(v string) *OpenAckServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenAckServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenAckServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenAckServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenAckServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAckServiceResponse) SetHeaders(v map[string]*string) *OpenAckServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAckServiceResponse) SetBody(v *OpenAckServiceResponseBody) *OpenAckServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetHeaders(v map[string]*string) *TagResourcesRequest {
	s.Headers = v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

type UntagResourcesQuery struct {
	RegionId     *string   `json:"region_id,omitempty" xml:"region_id,omitempty" require:"true"`
	ResourceType *string   `json:"resource_type,omitempty" xml:"resource_type,omitempty" require:"true"`
	ResourceIds  []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" require:"true" type:"Repeated"`
	TagKeys      []*string `json:"tag_keys,omitempty" xml:"tag_keys,omitempty" require:"true" type:"Repeated"`
}

func (s UntagResourcesQuery) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesQuery) GoString() string {
	return s.String()
}

func (s *UntagResourcesQuery) SetRegionId(v string) *UntagResourcesQuery {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesQuery) SetResourceType(v string) *UntagResourcesQuery {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesQuery) SetResourceIds(v []*string) *UntagResourcesQuery {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesQuery) SetTagKeys(v []*string) *UntagResourcesQuery {
	s.TagKeys = v
	return s
}

type UntagResourcesRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UntagResourcesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetHeaders(v map[string]*string) *UntagResourcesRequest {
	s.Headers = v
	return s
}

func (s *UntagResourcesRequest) SetQuery(v *UntagResourcesQuery) *UntagResourcesRequest {
	s.Query = v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

type DescribeClusterAddonsUpgradeStatusQuery struct {
	ComponentIds []*string `json:"componentIds,omitempty" xml:"componentIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClusterAddonsUpgradeStatusQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusQuery) SetComponentIds(v []*string) *DescribeClusterAddonsUpgradeStatusQuery {
	s.ComponentIds = v
	return s
}

type DescribeClusterAddonsUpgradeStatusRequest struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClusterAddonsUpgradeStatusQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DescribeClusterAddonsUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) SetHeaders(v map[string]*string) *DescribeClusterAddonsUpgradeStatusRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusRequest) SetQuery(v *DescribeClusterAddonsUpgradeStatusQuery) *DescribeClusterAddonsUpgradeStatusRequest {
	s.Query = v
	return s
}

type DescribeClusterAddonsUpgradeStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterAddonsUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsUpgradeStatusResponse {
	s.Headers = v
	return s
}

type ModifyClusterConfigurationBody struct {
	Body *ModifyClusterConfigurationBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ModifyClusterConfigurationBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationBody) SetBody(v *ModifyClusterConfigurationBodyBody) *ModifyClusterConfigurationBody {
	s.Body = v
	return s
}

type ModifyClusterConfigurationBodyBody struct {
	CustomizeConfig []*ModifyClusterConfigurationBodyBodyCustomizeConfig `json:"customize_config,omitempty" xml:"customize_config,omitempty" type:"Repeated"`
}

func (s ModifyClusterConfigurationBodyBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationBodyBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationBodyBody) SetCustomizeConfig(v []*ModifyClusterConfigurationBodyBodyCustomizeConfig) *ModifyClusterConfigurationBodyBody {
	s.CustomizeConfig = v
	return s
}

type ModifyClusterConfigurationBodyBodyCustomizeConfig struct {
	Name    *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	Configs []*ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
}

func (s ModifyClusterConfigurationBodyBodyCustomizeConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationBodyBodyCustomizeConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationBodyBodyCustomizeConfig) SetName(v string) *ModifyClusterConfigurationBodyBodyCustomizeConfig {
	s.Name = &v
	return s
}

func (s *ModifyClusterConfigurationBodyBodyCustomizeConfig) SetConfigs(v []*ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs) *ModifyClusterConfigurationBodyBodyCustomizeConfig {
	s.Configs = v
	return s
}

type ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs) SetKey(v string) *ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs {
	s.Key = &v
	return s
}

func (s *ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs) SetValue(v string) *ModifyClusterConfigurationBodyBodyCustomizeConfigConfigs {
	s.Value = &v
	return s
}

type ModifyClusterConfigurationRequest struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *ModifyClusterConfigurationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationRequest) SetHeaders(v map[string]*string) *ModifyClusterConfigurationRequest {
	s.Headers = v
	return s
}

func (s *ModifyClusterConfigurationRequest) SetBody(v *ModifyClusterConfigurationBody) *ModifyClusterConfigurationRequest {
	s.Body = v
	return s
}

type ModifyClusterConfigurationResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyClusterConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationResponse) SetHeaders(v map[string]*string) *ModifyClusterConfigurationResponse {
	s.Headers = v
	return s
}

type DeleteKubernetesTriggerRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteKubernetesTriggerRequest) SetHeaders(v map[string]*string) *DeleteKubernetesTriggerRequest {
	s.Headers = v
	return s
}

type DeleteKubernetesTriggerResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteKubernetesTriggerResponse) SetHeaders(v map[string]*string) *DeleteKubernetesTriggerResponse {
	s.Headers = v
	return s
}

type CreateKubernetesTriggerBody struct {
	Body *CreateKubernetesTriggerBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateKubernetesTriggerBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerBody) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerBody) SetBody(v *CreateKubernetesTriggerBodyBody) *CreateKubernetesTriggerBody {
	s.Body = v
	return s
}

type CreateKubernetesTriggerBodyBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Action    *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateKubernetesTriggerBodyBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerBodyBody) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerBodyBody) SetClusterId(v string) *CreateKubernetesTriggerBodyBody {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerBodyBody) SetProjectId(v string) *CreateKubernetesTriggerBodyBody {
	s.ProjectId = &v
	return s
}

func (s *CreateKubernetesTriggerBodyBody) SetAction(v string) *CreateKubernetesTriggerBodyBody {
	s.Action = &v
	return s
}

func (s *CreateKubernetesTriggerBodyBody) SetType(v string) *CreateKubernetesTriggerBodyBody {
	s.Type = &v
	return s
}

type CreateKubernetesTriggerRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *CreateKubernetesTriggerBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerRequest) SetHeaders(v map[string]*string) *CreateKubernetesTriggerRequest {
	s.Headers = v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetBody(v *CreateKubernetesTriggerBody) *CreateKubernetesTriggerRequest {
	s.Body = v
	return s
}

type CreateKubernetesTriggerResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Action    *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	Id        *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s CreateKubernetesTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponseBody) SetClusterId(v string) *CreateKubernetesTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetProjectId(v string) *CreateKubernetesTriggerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetAction(v string) *CreateKubernetesTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetId(v string) *CreateKubernetesTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetType(v string) *CreateKubernetesTriggerResponseBody {
	s.Type = &v
	return s
}

type CreateKubernetesTriggerResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKubernetesTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponse) SetHeaders(v map[string]*string) *CreateKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateKubernetesTriggerResponse) SetBody(v *CreateKubernetesTriggerResponseBody) *CreateKubernetesTriggerResponse {
	s.Body = v
	return s
}

type GetKubernetesTriggerQuery struct {
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty" require:"true"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Action    *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s GetKubernetesTriggerQuery) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerQuery) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerQuery) SetNamespace(v string) *GetKubernetesTriggerQuery {
	s.Namespace = &v
	return s
}

func (s *GetKubernetesTriggerQuery) SetType(v string) *GetKubernetesTriggerQuery {
	s.Type = &v
	return s
}

func (s *GetKubernetesTriggerQuery) SetName(v string) *GetKubernetesTriggerQuery {
	s.Name = &v
	return s
}

func (s *GetKubernetesTriggerQuery) SetAction(v string) *GetKubernetesTriggerQuery {
	s.Action = &v
	return s
}

type GetKubernetesTriggerRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetKubernetesTriggerQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetKubernetesTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerRequest) SetHeaders(v map[string]*string) *GetKubernetesTriggerRequest {
	s.Headers = v
	return s
}

func (s *GetKubernetesTriggerRequest) SetQuery(v *GetKubernetesTriggerQuery) *GetKubernetesTriggerRequest {
	s.Query = v
	return s
}

type GetKubernetesTriggerResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetKubernetesTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerResponse) SetHeaders(v map[string]*string) *GetKubernetesTriggerResponse {
	s.Headers = v
	return s
}

type ListTagResourcesQuery struct {
	ResourceType *string   `json:"resource_type,omitempty" xml:"resource_type,omitempty" require:"true"`
	RegionId     *string   `json:"region_id,omitempty" xml:"region_id,omitempty" require:"true"`
	NextToken    *string   `json:"next_token,omitempty" xml:"next_token,omitempty"`
	ResourceIds  []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesQuery) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesQuery) GoString() string {
	return s.String()
}

func (s *ListTagResourcesQuery) SetResourceType(v string) *ListTagResourcesQuery {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesQuery) SetRegionId(v string) *ListTagResourcesQuery {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesQuery) SetNextToken(v string) *ListTagResourcesQuery {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesQuery) SetResourceIds(v []*string) *ListTagResourcesQuery {
	s.ResourceIds = v
	return s
}

type ListTagResourcesRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListTagResourcesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetHeaders(v map[string]*string) *ListTagResourcesRequest {
	s.Headers = v
	return s
}

func (s *ListTagResourcesRequest) SetQuery(v *ListTagResourcesQuery) *ListTagResourcesRequest {
	s.Query = v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"next_token,omitempty" xml:"next_token,omitempty" require:"true"`
	RequestId    *string                                   `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"tag_resources,omitempty" xml:"tag_resources,omitempty" require:"true" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"tag_resource,omitempty" xml:"tag_resource,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	TagValue     *string `json:"tag_value,omitempty" xml:"tag_value,omitempty" require:"true"`
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty" require:"true"`
	ResourceId   *string `json:"resource_id,omitempty" xml:"resource_id,omitempty" require:"true"`
	TagKey       *string `json:"tag_key,omitempty" xml:"tag_key,omitempty" require:"true"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ResumeComponentUpgradeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ResumeComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeRequest) SetHeaders(v map[string]*string) *ResumeComponentUpgradeRequest {
	s.Headers = v
	return s
}

type ResumeComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ResumeComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeResponse) SetHeaders(v map[string]*string) *ResumeComponentUpgradeResponse {
	s.Headers = v
	return s
}

type PauseComponentUpgradeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s PauseComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeRequest) SetHeaders(v map[string]*string) *PauseComponentUpgradeRequest {
	s.Headers = v
	return s
}

type PauseComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PauseComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeResponse) SetHeaders(v map[string]*string) *PauseComponentUpgradeResponse {
	s.Headers = v
	return s
}

type CancelComponentUpgradeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CancelComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeRequest) SetHeaders(v map[string]*string) *CancelComponentUpgradeRequest {
	s.Headers = v
	return s
}

type CancelComponentUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeResponse) SetHeaders(v map[string]*string) *CancelComponentUpgradeResponse {
	s.Headers = v
	return s
}

type DeleteClusterNodepoolRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteClusterNodepoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodepoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodepoolRequest) SetHeaders(v map[string]*string) *DeleteClusterNodepoolRequest {
	s.Headers = v
	return s
}

type DeleteClusterNodepoolResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterNodepoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodepoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodepoolResponse) SetHeaders(v map[string]*string) *DeleteClusterNodepoolResponse {
	s.Headers = v
	return s
}

type ScaleClusterNodePoolBody struct {
	Body *ScaleClusterNodePoolBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ScaleClusterNodePoolBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterNodePoolBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolBody) SetBody(v *ScaleClusterNodePoolBodyBody) *ScaleClusterNodePoolBody {
	s.Body = v
	return s
}

type ScaleClusterNodePoolBodyBody struct {
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s ScaleClusterNodePoolBodyBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterNodePoolBodyBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolBodyBody) SetCount(v int64) *ScaleClusterNodePoolBodyBody {
	s.Count = &v
	return s
}

type ScaleClusterNodePoolRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *ScaleClusterNodePoolBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleClusterNodePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolRequest) SetHeaders(v map[string]*string) *ScaleClusterNodePoolRequest {
	s.Headers = v
	return s
}

func (s *ScaleClusterNodePoolRequest) SetBody(v *ScaleClusterNodePoolBody) *ScaleClusterNodePoolRequest {
	s.Body = v
	return s
}

type ScaleClusterNodePoolResponseBody struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s ScaleClusterNodePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolResponseBody) SetTaskId(v string) *ScaleClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

type ScaleClusterNodePoolResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleClusterNodePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolResponse) SetHeaders(v map[string]*string) *ScaleClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterNodePoolResponse) SetBody(v *ScaleClusterNodePoolResponseBody) *ScaleClusterNodePoolResponse {
	s.Body = v
	return s
}

type DescribeClusterNodePoolDetailRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterNodePoolDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodePoolDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailRequest) SetHeaders(v map[string]*string) *DescribeClusterNodePoolDetailRequest {
	s.Headers = v
	return s
}

type DescribeClusterNodePoolDetailResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNodePoolDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterNodePoolDetailResponse {
	s.Headers = v
	return s
}

type DescribeClusterNodePoolsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterNodePoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodePoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsRequest) SetHeaders(v map[string]*string) *DescribeClusterNodePoolsRequest {
	s.Headers = v
	return s
}

type DescribeClusterNodePoolsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNodePoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodePoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponse) SetHeaders(v map[string]*string) *DescribeClusterNodePoolsResponse {
	s.Headers = v
	return s
}

type CreateClusterNodePoolRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CreateClusterNodePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequest) SetHeaders(v map[string]*string) *CreateClusterNodePoolRequest {
	s.Headers = v
	return s
}

type CreateClusterNodePoolResponseBody struct {
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty" require:"true"`
}

func (s CreateClusterNodePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBody) SetNodepoolId(v string) *CreateClusterNodePoolResponseBody {
	s.NodepoolId = &v
	return s
}

type CreateClusterNodePoolResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterNodePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponse) SetHeaders(v map[string]*string) *CreateClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterNodePoolResponse) SetBody(v *CreateClusterNodePoolResponseBody) *CreateClusterNodePoolResponse {
	s.Body = v
	return s
}

type ModifyClusterNodePoolRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ModifyClusterNodePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequest) SetHeaders(v map[string]*string) *ModifyClusterNodePoolRequest {
	s.Headers = v
	return s
}

type ModifyClusterNodePoolResponseBody struct {
	TaskId     *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty" require:"true"`
}

func (s ModifyClusterNodePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolResponseBody) SetTaskId(v string) *ModifyClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyClusterNodePoolResponseBody) SetNodepoolId(v string) *ModifyClusterNodePoolResponseBody {
	s.NodepoolId = &v
	return s
}

type ModifyClusterNodePoolResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterNodePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolResponse) SetHeaders(v map[string]*string) *ModifyClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNodePoolResponse) SetBody(v *ModifyClusterNodePoolResponseBody) *ModifyClusterNodePoolResponse {
	s.Body = v
	return s
}

type CancelWorkflowBody struct {
	Body *CancelWorkflowBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CancelWorkflowBody) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowBody) GoString() string {
	return s.String()
}

func (s *CancelWorkflowBody) SetBody(v *CancelWorkflowBodyBody) *CancelWorkflowBody {
	s.Body = v
	return s
}

type CancelWorkflowBodyBody struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s CancelWorkflowBodyBody) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowBodyBody) GoString() string {
	return s.String()
}

func (s *CancelWorkflowBodyBody) SetAction(v string) *CancelWorkflowBodyBody {
	s.Action = &v
	return s
}

type CancelWorkflowRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *CancelWorkflowBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CancelWorkflowRequest) SetHeaders(v map[string]*string) *CancelWorkflowRequest {
	s.Headers = v
	return s
}

func (s *CancelWorkflowRequest) SetBody(v *CancelWorkflowBody) *CancelWorkflowRequest {
	s.Body = v
	return s
}

type CancelWorkflowResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CancelWorkflowResponse) SetHeaders(v map[string]*string) *CancelWorkflowResponse {
	s.Headers = v
	return s
}

type DescirbeWorkflowRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescirbeWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescirbeWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DescirbeWorkflowRequest) SetHeaders(v map[string]*string) *DescirbeWorkflowRequest {
	s.Headers = v
	return s
}

type DescirbeWorkflowResponseBody struct {
	Duration       *string `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
	JobNamespace   *string `json:"job_namespace,omitempty" xml:"job_namespace,omitempty" require:"true"`
	TotalReads     *string `json:"total_reads,omitempty" xml:"total_reads,omitempty" require:"true"`
	JobName        *string `json:"job_name,omitempty" xml:"job_name,omitempty" require:"true"`
	CreateTime     *string `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	InputDataSize  *string `json:"input_data_size,omitempty" xml:"input_data_size,omitempty" require:"true"`
	UserInputData  *string `json:"user_input_data,omitempty" xml:"user_input_data,omitempty" require:"true"`
	TotalBases     *string `json:"total_bases,omitempty" xml:"total_bases,omitempty" require:"true"`
	OutputDataSize *string `json:"output_data_size,omitempty" xml:"output_data_size,omitempty" require:"true"`
	FinishTime     *string `json:"finish_time,omitempty" xml:"finish_time,omitempty" require:"true"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s DescirbeWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescirbeWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DescirbeWorkflowResponseBody) SetDuration(v string) *DescirbeWorkflowResponseBody {
	s.Duration = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetJobNamespace(v string) *DescirbeWorkflowResponseBody {
	s.JobNamespace = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetTotalReads(v string) *DescirbeWorkflowResponseBody {
	s.TotalReads = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetJobName(v string) *DescirbeWorkflowResponseBody {
	s.JobName = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetCreateTime(v string) *DescirbeWorkflowResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetInputDataSize(v string) *DescirbeWorkflowResponseBody {
	s.InputDataSize = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetUserInputData(v string) *DescirbeWorkflowResponseBody {
	s.UserInputData = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetTotalBases(v string) *DescirbeWorkflowResponseBody {
	s.TotalBases = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetOutputDataSize(v string) *DescirbeWorkflowResponseBody {
	s.OutputDataSize = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetFinishTime(v string) *DescirbeWorkflowResponseBody {
	s.FinishTime = &v
	return s
}

func (s *DescirbeWorkflowResponseBody) SetStatus(v string) *DescirbeWorkflowResponseBody {
	s.Status = &v
	return s
}

type DescirbeWorkflowResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescirbeWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescirbeWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescirbeWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DescirbeWorkflowResponse) SetHeaders(v map[string]*string) *DescirbeWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DescirbeWorkflowResponse) SetBody(v *DescirbeWorkflowResponseBody) *DescirbeWorkflowResponse {
	s.Body = v
	return s
}

type RemoveWorkflowRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s RemoveWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkflowRequest) GoString() string {
	return s.String()
}

func (s *RemoveWorkflowRequest) SetHeaders(v map[string]*string) *RemoveWorkflowRequest {
	s.Headers = v
	return s
}

type RemoveWorkflowResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkflowResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkflowResponse) SetHeaders(v map[string]*string) *RemoveWorkflowResponse {
	s.Headers = v
	return s
}

type DescribeWorkflowsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeWorkflowsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsRequest) SetHeaders(v map[string]*string) *DescribeWorkflowsRequest {
	s.Headers = v
	return s
}

type DescribeWorkflowsResponseBody struct {
	Jobs []*DescribeWorkflowsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeWorkflowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsResponseBody) SetJobs(v []*DescribeWorkflowsResponseBodyJobs) *DescribeWorkflowsResponseBody {
	s.Jobs = v
	return s
}

type DescribeWorkflowsResponseBodyJobs struct {
	ClusterId  *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	CreateTime *string `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	JobName    *string `json:"job_name,omitempty" xml:"job_name,omitempty" require:"true"`
}

func (s DescribeWorkflowsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsResponseBodyJobs) SetClusterId(v string) *DescribeWorkflowsResponseBodyJobs {
	s.ClusterId = &v
	return s
}

func (s *DescribeWorkflowsResponseBodyJobs) SetCreateTime(v string) *DescribeWorkflowsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *DescribeWorkflowsResponseBodyJobs) SetJobName(v string) *DescribeWorkflowsResponseBodyJobs {
	s.JobName = &v
	return s
}

type DescribeWorkflowsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWorkflowsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsResponse) SetHeaders(v map[string]*string) *DescribeWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkflowsResponse) SetBody(v *DescribeWorkflowsResponseBody) *DescribeWorkflowsResponse {
	s.Body = v
	return s
}

type StartWorkflowBody struct {
	Body *StartWorkflowBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s StartWorkflowBody) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowBody) GoString() string {
	return s.String()
}

func (s *StartWorkflowBody) SetBody(v *StartWorkflowBodyBody) *StartWorkflowBody {
	s.Body = v
	return s
}

type StartWorkflowBodyBody struct {
	WgsVcfOutFilename          *string `json:"wgs_vcf_out_filename,omitempty" xml:"wgs_vcf_out_filename,omitempty"`
	WgsFastqPath               *string `json:"wgs_fastq_path,omitempty" xml:"wgs_fastq_path,omitempty"`
	MappingBucketName          *string `json:"mapping_bucket_name,omitempty" xml:"mapping_bucket_name,omitempty"`
	MappingReferencePath       *string `json:"mapping_reference_path,omitempty" xml:"mapping_reference_path,omitempty"`
	WorkflowType               *string `json:"workflow_type,omitempty" xml:"workflow_type,omitempty" require:"true"`
	WgsBucketName              *string `json:"wgs_bucket_name,omitempty" xml:"wgs_bucket_name,omitempty"`
	MappingOssRegion           *string `json:"mapping_oss_region,omitempty" xml:"mapping_oss_region,omitempty"`
	WgsReferencePath           *string `json:"wgs_reference_path,omitempty" xml:"wgs_reference_path,omitempty"`
	WgsFastqFirstFilename      *string `json:"wgs_fastq_first_filename,omitempty" xml:"wgs_fastq_first_filename,omitempty"`
	MappingFastqFirstFilename  *string `json:"mapping_fastq_first_filename,omitempty" xml:"mapping_fastq_first_filename,omitempty"`
	WgsOssRegion               *string `json:"wgs_oss_region,omitempty" xml:"wgs_oss_region,omitempty"`
	MappingBamOutFilename      *string `json:"mapping_bam_out_filename,omitempty" xml:"mapping_bam_out_filename,omitempty"`
	MappingFastqPath           *string `json:"mapping_fastq_path,omitempty" xml:"mapping_fastq_path,omitempty"`
	Service                    *string `json:"service,omitempty" xml:"service,omitempty"`
	WgsVcfOutPath              *string `json:"wgs_vcf_out_path,omitempty" xml:"wgs_vcf_out_path,omitempty"`
	MappingIsMarkDup           *string `json:"mapping_is_mark_dup,omitempty" xml:"mapping_is_mark_dup,omitempty"`
	MappingBamOutPath          *string `json:"mapping_bam_out_path,omitempty" xml:"mapping_bam_out_path,omitempty"`
	MappingFastqSecondFilename *string `json:"mapping_fastq_second_filename,omitempty" xml:"mapping_fastq_second_filename,omitempty"`
	WgsFastqSecondFilename     *string `json:"wgs_fastq_second_filename,omitempty" xml:"wgs_fastq_second_filename,omitempty"`
}

func (s StartWorkflowBodyBody) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowBodyBody) GoString() string {
	return s.String()
}

func (s *StartWorkflowBodyBody) SetWgsVcfOutFilename(v string) *StartWorkflowBodyBody {
	s.WgsVcfOutFilename = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsFastqPath(v string) *StartWorkflowBodyBody {
	s.WgsFastqPath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingBucketName(v string) *StartWorkflowBodyBody {
	s.MappingBucketName = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingReferencePath(v string) *StartWorkflowBodyBody {
	s.MappingReferencePath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWorkflowType(v string) *StartWorkflowBodyBody {
	s.WorkflowType = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsBucketName(v string) *StartWorkflowBodyBody {
	s.WgsBucketName = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingOssRegion(v string) *StartWorkflowBodyBody {
	s.MappingOssRegion = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsReferencePath(v string) *StartWorkflowBodyBody {
	s.WgsReferencePath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsFastqFirstFilename(v string) *StartWorkflowBodyBody {
	s.WgsFastqFirstFilename = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingFastqFirstFilename(v string) *StartWorkflowBodyBody {
	s.MappingFastqFirstFilename = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsOssRegion(v string) *StartWorkflowBodyBody {
	s.WgsOssRegion = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingBamOutFilename(v string) *StartWorkflowBodyBody {
	s.MappingBamOutFilename = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingFastqPath(v string) *StartWorkflowBodyBody {
	s.MappingFastqPath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetService(v string) *StartWorkflowBodyBody {
	s.Service = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsVcfOutPath(v string) *StartWorkflowBodyBody {
	s.WgsVcfOutPath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingIsMarkDup(v string) *StartWorkflowBodyBody {
	s.MappingIsMarkDup = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingBamOutPath(v string) *StartWorkflowBodyBody {
	s.MappingBamOutPath = &v
	return s
}

func (s *StartWorkflowBodyBody) SetMappingFastqSecondFilename(v string) *StartWorkflowBodyBody {
	s.MappingFastqSecondFilename = &v
	return s
}

func (s *StartWorkflowBodyBody) SetWgsFastqSecondFilename(v string) *StartWorkflowBodyBody {
	s.WgsFastqSecondFilename = &v
	return s
}

type StartWorkflowRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *StartWorkflowBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowRequest) GoString() string {
	return s.String()
}

func (s *StartWorkflowRequest) SetHeaders(v map[string]*string) *StartWorkflowRequest {
	s.Headers = v
	return s
}

func (s *StartWorkflowRequest) SetBody(v *StartWorkflowBody) *StartWorkflowRequest {
	s.Body = v
	return s
}

type StartWorkflowResponseBody struct {
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty" require:"true"`
}

func (s StartWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *StartWorkflowResponseBody) SetJobName(v string) *StartWorkflowResponseBody {
	s.JobName = &v
	return s
}

type StartWorkflowResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowResponse) GoString() string {
	return s.String()
}

func (s *StartWorkflowResponse) SetHeaders(v map[string]*string) *StartWorkflowResponse {
	s.Headers = v
	return s
}

func (s *StartWorkflowResponse) SetBody(v *StartWorkflowResponseBody) *StartWorkflowResponse {
	s.Body = v
	return s
}

type DescribeUserPermissionRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionRequest) SetHeaders(v map[string]*string) *DescribeUserPermissionRequest {
	s.Headers = v
	return s
}

type DescribeUserPermissionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionResponse) SetHeaders(v map[string]*string) *DescribeUserPermissionResponse {
	s.Headers = v
	return s
}

type GrantPermissionsBody struct {
	Body []*GrantPermissionsBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GrantPermissionsBody) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionsBody) SetBody(v []*GrantPermissionsBodyBody) *GrantPermissionsBody {
	s.Body = v
	return s
}

type GrantPermissionsBodyBody struct {
	RoleName  *string `json:"role_name,omitempty" xml:"role_name,omitempty" require:"true"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	IsRamRole *bool   `json:"is_ram_role,omitempty" xml:"is_ram_role,omitempty"`
	IsCustom  *bool   `json:"is_custom,omitempty" xml:"is_custom,omitempty"`
	RoleType  *string `json:"role_type,omitempty" xml:"role_type,omitempty" require:"true"`
	Cluster   *string `json:"cluster,omitempty" xml:"cluster,omitempty" require:"true"`
}

func (s GrantPermissionsBodyBody) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsBodyBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionsBodyBody) SetRoleName(v string) *GrantPermissionsBodyBody {
	s.RoleName = &v
	return s
}

func (s *GrantPermissionsBodyBody) SetNamespace(v string) *GrantPermissionsBodyBody {
	s.Namespace = &v
	return s
}

func (s *GrantPermissionsBodyBody) SetIsRamRole(v bool) *GrantPermissionsBodyBody {
	s.IsRamRole = &v
	return s
}

func (s *GrantPermissionsBodyBody) SetIsCustom(v bool) *GrantPermissionsBodyBody {
	s.IsCustom = &v
	return s
}

func (s *GrantPermissionsBodyBody) SetRoleType(v string) *GrantPermissionsBodyBody {
	s.RoleType = &v
	return s
}

func (s *GrantPermissionsBodyBody) SetCluster(v string) *GrantPermissionsBodyBody {
	s.Cluster = &v
	return s
}

type GrantPermissionsRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *GrantPermissionsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantPermissionsRequest) SetHeaders(v map[string]*string) *GrantPermissionsRequest {
	s.Headers = v
	return s
}

func (s *GrantPermissionsRequest) SetBody(v *GrantPermissionsBody) *GrantPermissionsRequest {
	s.Body = v
	return s
}

type GrantPermissionsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GrantPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantPermissionsResponse) SetHeaders(v map[string]*string) *GrantPermissionsResponse {
	s.Headers = v
	return s
}

type UnInstallClusterAddonsBody struct {
	Addons []*UnInstallClusterAddonsBodyAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
}

func (s UnInstallClusterAddonsBody) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsBody) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsBody) SetAddons(v []*UnInstallClusterAddonsBodyAddons) *UnInstallClusterAddonsBody {
	s.Addons = v
	return s
}

type UnInstallClusterAddonsBodyAddons struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UnInstallClusterAddonsBodyAddons) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsBodyAddons) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsBodyAddons) SetName(v string) *UnInstallClusterAddonsBodyAddons {
	s.Name = &v
	return s
}

type UnInstallClusterAddonsRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UnInstallClusterAddonsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnInstallClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequest) SetHeaders(v map[string]*string) *UnInstallClusterAddonsRequest {
	s.Headers = v
	return s
}

func (s *UnInstallClusterAddonsRequest) SetBody(v *UnInstallClusterAddonsBody) *UnInstallClusterAddonsRequest {
	s.Body = v
	return s
}

type UnInstallClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UnInstallClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponse) SetHeaders(v map[string]*string) *UnInstallClusterAddonsResponse {
	s.Headers = v
	return s
}

type DescribeAddonsQuery struct {
	Region      *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
}

func (s DescribeAddonsQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsQuery) GoString() string {
	return s.String()
}

func (s *DescribeAddonsQuery) SetRegion(v string) *DescribeAddonsQuery {
	s.Region = &v
	return s
}

func (s *DescribeAddonsQuery) SetClusterType(v string) *DescribeAddonsQuery {
	s.ClusterType = &v
	return s
}

type DescribeAddonsRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeAddonsQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DescribeAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonsRequest) SetHeaders(v map[string]*string) *DescribeAddonsRequest {
	s.Headers = v
	return s
}

func (s *DescribeAddonsRequest) SetQuery(v *DescribeAddonsQuery) *DescribeAddonsRequest {
	s.Query = v
	return s
}

type DescribeAddonsResponseBody struct {
	StandardComponents map[string]interface{}                       `json:"StandardComponents,omitempty" xml:"StandardComponents,omitempty" require:"true"`
	ComponentGroups    []*DescribeAddonsResponseBodyComponentGroups `json:"ComponentGroups,omitempty" xml:"ComponentGroups,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBody) SetStandardComponents(v map[string]interface{}) *DescribeAddonsResponseBody {
	s.StandardComponents = v
	return s
}

func (s *DescribeAddonsResponseBody) SetComponentGroups(v []*DescribeAddonsResponseBodyComponentGroups) *DescribeAddonsResponseBody {
	s.ComponentGroups = v
	return s
}

type DescribeAddonsResponseBodyComponentGroups struct {
	GroupName *string                                           `json:"group_name,omitempty" xml:"group_name,omitempty" require:"true"`
	Items     []*DescribeAddonsResponseBodyComponentGroupsItems `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAddonsResponseBodyComponentGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroups) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetGroupName(v string) *DescribeAddonsResponseBodyComponentGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetItems(v []*DescribeAddonsResponseBodyComponentGroupsItems) *DescribeAddonsResponseBodyComponentGroups {
	s.Items = v
	return s
}

type DescribeAddonsResponseBodyComponentGroupsItems struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetName(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Name = &v
	return s
}

type DescribeAddonsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponse) SetHeaders(v map[string]*string) *DescribeAddonsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonsResponse) SetBody(v *DescribeAddonsResponseBody) *DescribeAddonsResponse {
	s.Body = v
	return s
}

type UpdateK8sClusterUserConfigExpireRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s UpdateK8sClusterUserConfigExpireRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireRequest) SetHeaders(v map[string]*string) *UpdateK8sClusterUserConfigExpireRequest {
	s.Headers = v
	return s
}

type UpdateK8sClusterUserConfigExpireResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateK8sClusterUserConfigExpireResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireResponse) SetHeaders(v map[string]*string) *UpdateK8sClusterUserConfigExpireResponse {
	s.Headers = v
	return s
}

type CancelClusterUpgradeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CancelClusterUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelClusterUpgradeRequest) GoString() string {
	return s.String()
}

func (s *CancelClusterUpgradeRequest) SetHeaders(v map[string]*string) *CancelClusterUpgradeRequest {
	s.Headers = v
	return s
}

type CancelClusterUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelClusterUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelClusterUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelClusterUpgradeResponse) SetHeaders(v map[string]*string) *CancelClusterUpgradeResponse {
	s.Headers = v
	return s
}

type DescribeClustersV1Query struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	PageSize    *int64  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PageNumber  *int64  `json:"page_number,omitempty" xml:"page_number,omitempty"`
}

func (s DescribeClustersV1Query) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1Query) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Query) SetName(v string) *DescribeClustersV1Query {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1Query) SetClusterType(v string) *DescribeClustersV1Query {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1Query) SetPageSize(v int64) *DescribeClustersV1Query {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersV1Query) SetPageNumber(v int64) *DescribeClustersV1Query {
	s.PageNumber = &v
	return s
}

type DescribeClustersV1Request struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClustersV1Query `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClustersV1Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1Request) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Request) SetHeaders(v map[string]*string) *DescribeClustersV1Request {
	s.Headers = v
	return s
}

func (s *DescribeClustersV1Request) SetQuery(v *DescribeClustersV1Query) *DescribeClustersV1Request {
	s.Query = v
	return s
}

type DescribeClustersV1Response struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClustersV1Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersV1Response) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Response) SetHeaders(v map[string]*string) *DescribeClustersV1Response {
	s.Headers = v
	return s
}

type DescribeUserQuotaRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaRequest) SetHeaders(v map[string]*string) *DescribeUserQuotaRequest {
	s.Headers = v
	return s
}

type DescribeUserQuotaResponseBody struct {
	ClusterNodepoolQuota *int64 `json:"cluster_nodepool_quota,omitempty" xml:"cluster_nodepool_quota,omitempty" require:"true"`
	AmkClusterQuota      *int64 `json:"amk_cluster_quota,omitempty" xml:"amk_cluster_quota,omitempty" require:"true"`
	ClusterQuota         *int64 `json:"cluster_quota,omitempty" xml:"cluster_quota,omitempty" require:"true"`
	NodeQuota            *int64 `json:"node_quota,omitempty" xml:"node_quota,omitempty" require:"true"`
	AskClusterQuota      *int64 `json:"ask_cluster_quota,omitempty" xml:"ask_cluster_quota,omitempty" require:"true"`
}

func (s DescribeUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBody) SetClusterNodepoolQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterNodepoolQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetAmkClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AmkClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetNodeQuota(v int64) *DescribeUserQuotaResponseBody {
	s.NodeQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetAskClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AskClusterQuota = &v
	return s
}

type DescribeUserQuotaResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserQuotaResponse) SetBody(v *DescribeUserQuotaResponseBody) *DescribeUserQuotaResponse {
	s.Body = v
	return s
}

type DescribeClusterV2UserKubeconfigQuery struct {
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigQuery) SetPrivateIpAddress(v bool) *DescribeClusterV2UserKubeconfigQuery {
	s.PrivateIpAddress = &v
	return s
}

type DescribeClusterV2UserKubeconfigRequest struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClusterV2UserKubeconfigQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetHeaders(v map[string]*string) *DescribeClusterV2UserKubeconfigRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetQuery(v *DescribeClusterV2UserKubeconfigQuery) *DescribeClusterV2UserKubeconfigRequest {
	s.Query = v
	return s
}

type DescribeClusterV2UserKubeconfigResponseBody struct {
	Config *string `json:"config,omitempty" xml:"config,omitempty" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterV2UserKubeconfigResponseBody {
	s.Config = &v
	return s
}

type DescribeClusterV2UserKubeconfigResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterV2UserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterV2UserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetBody(v *DescribeClusterV2UserKubeconfigResponseBody) *DescribeClusterV2UserKubeconfigResponse {
	s.Body = v
	return s
}

type RemoveClusterNodesBody struct {
	Body *RemoveClusterNodesBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s RemoveClusterNodesBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesBody) SetBody(v *RemoveClusterNodesBodyBody) *RemoveClusterNodesBody {
	s.Body = v
	return s
}

type RemoveClusterNodesBodyBody struct {
	ReleaseNode *bool     `json:"release_node,omitempty" xml:"release_node,omitempty"`
	DrainNode   *bool     `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	Nodes       []*string `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s RemoveClusterNodesBodyBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesBodyBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesBodyBody) SetReleaseNode(v bool) *RemoveClusterNodesBodyBody {
	s.ReleaseNode = &v
	return s
}

func (s *RemoveClusterNodesBodyBody) SetDrainNode(v bool) *RemoveClusterNodesBodyBody {
	s.DrainNode = &v
	return s
}

func (s *RemoveClusterNodesBodyBody) SetNodes(v []*string) *RemoveClusterNodesBodyBody {
	s.Nodes = v
	return s
}

type RemoveClusterNodesRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *RemoveClusterNodesBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesRequest) SetHeaders(v map[string]*string) *RemoveClusterNodesRequest {
	s.Headers = v
	return s
}

func (s *RemoveClusterNodesRequest) SetBody(v *RemoveClusterNodesBody) *RemoveClusterNodesRequest {
	s.Body = v
	return s
}

type RemoveClusterNodesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RemoveClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesResponse) SetHeaders(v map[string]*string) *RemoveClusterNodesResponse {
	s.Headers = v
	return s
}

type UpgradeClusterBody struct {
	Body *UpgradeClusterBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpgradeClusterBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterBody) SetBody(v *UpgradeClusterBodyBody) *UpgradeClusterBody {
	s.Body = v
	return s
}

type UpgradeClusterBodyBody struct {
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
	NextVersion   *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
}

func (s UpgradeClusterBodyBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterBodyBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterBodyBody) SetComponentName(v string) *UpgradeClusterBodyBody {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterBodyBody) SetVersion(v string) *UpgradeClusterBodyBody {
	s.Version = &v
	return s
}

func (s *UpgradeClusterBodyBody) SetNextVersion(v string) *UpgradeClusterBodyBody {
	s.NextVersion = &v
	return s
}

type UpgradeClusterRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UpgradeClusterBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequest) SetHeaders(v map[string]*string) *UpgradeClusterRequest {
	s.Headers = v
	return s
}

func (s *UpgradeClusterRequest) SetBody(v *UpgradeClusterBody) *UpgradeClusterRequest {
	s.Body = v
	return s
}

type UpgradeClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponse) SetHeaders(v map[string]*string) *UpgradeClusterResponse {
	s.Headers = v
	return s
}

type PauseClusterUpgradeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s PauseClusterUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseClusterUpgradeRequest) GoString() string {
	return s.String()
}

func (s *PauseClusterUpgradeRequest) SetHeaders(v map[string]*string) *PauseClusterUpgradeRequest {
	s.Headers = v
	return s
}

type PauseClusterUpgradeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PauseClusterUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseClusterUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseClusterUpgradeResponse) SetHeaders(v map[string]*string) *PauseClusterUpgradeResponse {
	s.Headers = v
	return s
}

type ResumeUpgradeClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ResumeUpgradeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeUpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *ResumeUpgradeClusterRequest) SetHeaders(v map[string]*string) *ResumeUpgradeClusterRequest {
	s.Headers = v
	return s
}

type ResumeUpgradeClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ResumeUpgradeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeUpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *ResumeUpgradeClusterResponse) SetHeaders(v map[string]*string) *ResumeUpgradeClusterResponse {
	s.Headers = v
	return s
}

type GetUpgradeStatusRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusRequest) SetHeaders(v map[string]*string) *GetUpgradeStatusRequest {
	s.Headers = v
	return s
}

type GetUpgradeStatusResponseBody struct {
	ErrorMessage     *string                                  `json:"error_message,omitempty" xml:"error_message,omitempty" require:"true"`
	PrecheckReportId *string                                  `json:"precheck_report_id,omitempty" xml:"precheck_report_id,omitempty" require:"true"`
	UpgradeStep      *string                                  `json:"upgrade_step,omitempty" xml:"upgrade_step,omitempty" require:"true"`
	Status           *string                                  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	UpgradeTask      *GetUpgradeStatusResponseBodyUpgradeTask `json:"upgrade_task,omitempty" xml:"upgrade_task,omitempty" require:"true" type:"Struct"`
}

func (s GetUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBody) SetErrorMessage(v string) *GetUpgradeStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetPrecheckReportId(v string) *GetUpgradeStatusResponseBody {
	s.PrecheckReportId = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetUpgradeStep(v string) *GetUpgradeStatusResponseBody {
	s.UpgradeStep = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetStatus(v string) *GetUpgradeStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetUpgradeTask(v *GetUpgradeStatusResponseBodyUpgradeTask) *GetUpgradeStatusResponseBody {
	s.UpgradeTask = v
	return s
}

type GetUpgradeStatusResponseBodyUpgradeTask struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s GetUpgradeStatusResponseBodyUpgradeTask) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusResponseBodyUpgradeTask) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) SetMessage(v string) *GetUpgradeStatusResponseBodyUpgradeTask {
	s.Message = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) SetStatus(v string) *GetUpgradeStatusResponseBodyUpgradeTask {
	s.Status = &v
	return s
}

type GetUpgradeStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponse) SetHeaders(v map[string]*string) *GetUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUpgradeStatusResponse) SetBody(v *GetUpgradeStatusResponseBody) *GetUpgradeStatusResponse {
	s.Body = v
	return s
}

type ModifyClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ModifyClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequest) SetHeaders(v map[string]*string) *ModifyClusterRequest {
	s.Headers = v
	return s
}

type ModifyClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
}

func (s ModifyClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponseBody) SetClusterId(v string) *ModifyClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetTaskId(v string) *ModifyClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyClusterResponseBody) SetRequestId(v string) *ModifyClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponse) SetHeaders(v map[string]*string) *ModifyClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterResponse) SetBody(v *ModifyClusterResponseBody) *ModifyClusterResponse {
	s.Body = v
	return s
}

type InstallClusterAddonsBody struct {
	Body []*InstallClusterAddonsBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s InstallClusterAddonsBody) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsBody) SetBody(v []*InstallClusterAddonsBodyBody) *InstallClusterAddonsBody {
	s.Body = v
	return s
}

type InstallClusterAddonsBodyBody struct {
	Config  *string `json:"config,omitempty" xml:"config,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s InstallClusterAddonsBodyBody) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsBodyBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsBodyBody) SetConfig(v string) *InstallClusterAddonsBodyBody {
	s.Config = &v
	return s
}

func (s *InstallClusterAddonsBodyBody) SetName(v string) *InstallClusterAddonsBodyBody {
	s.Name = &v
	return s
}

func (s *InstallClusterAddonsBodyBody) SetVersion(v string) *InstallClusterAddonsBodyBody {
	s.Version = &v
	return s
}

type InstallClusterAddonsRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *InstallClusterAddonsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequest) SetHeaders(v map[string]*string) *InstallClusterAddonsRequest {
	s.Headers = v
	return s
}

func (s *InstallClusterAddonsRequest) SetBody(v *InstallClusterAddonsBody) *InstallClusterAddonsRequest {
	s.Body = v
	return s
}

type InstallClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s InstallClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsResponse) SetHeaders(v map[string]*string) *InstallClusterAddonsResponse {
	s.Headers = v
	return s
}

type ModifyClusterTagsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ModifyClusterTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsRequest) SetHeaders(v map[string]*string) *ModifyClusterTagsRequest {
	s.Headers = v
	return s
}

type ModifyClusterTagsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyClusterTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsResponse) SetHeaders(v map[string]*string) *ModifyClusterTagsResponse {
	s.Headers = v
	return s
}

type DescribeClusterNamespacesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNamespacesRequest) SetHeaders(v map[string]*string) *DescribeClusterNamespacesRequest {
	s.Headers = v
	return s
}

type DescribeClusterNamespacesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeClusterNamespacesResponse {
	s.Headers = v
	return s
}

type DescribeExternalAgentQuery struct {
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeExternalAgentQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentQuery) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentQuery) SetPrivateIpAddress(v string) *DescribeExternalAgentQuery {
	s.PrivateIpAddress = &v
	return s
}

type DescribeExternalAgentRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeExternalAgentQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeExternalAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentRequest) SetHeaders(v map[string]*string) *DescribeExternalAgentRequest {
	s.Headers = v
	return s
}

func (s *DescribeExternalAgentRequest) SetQuery(v *DescribeExternalAgentQuery) *DescribeExternalAgentRequest {
	s.Query = v
	return s
}

type DescribeExternalAgentResponseBody struct {
	Config *string `json:"config,omitempty" xml:"config,omitempty" require:"true"`
}

func (s DescribeExternalAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponseBody) SetConfig(v string) *DescribeExternalAgentResponseBody {
	s.Config = &v
	return s
}

type DescribeExternalAgentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExternalAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponse) SetHeaders(v map[string]*string) *DescribeExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *DescribeExternalAgentResponse) SetBody(v *DescribeExternalAgentResponseBody) *DescribeExternalAgentResponse {
	s.Body = v
	return s
}

type DescribeClusterAttachScriptsBody struct {
	Body *DescribeClusterAttachScriptsBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s DescribeClusterAttachScriptsBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsBody) SetBody(v *DescribeClusterAttachScriptsBodyBody) *DescribeClusterAttachScriptsBody {
	s.Body = v
	return s
}

type DescribeClusterAttachScriptsBodyBody struct {
	KeepInstanceName *bool     `json:"keep_instance_name,omitempty" xml:"keep_instance_name,omitempty"`
	Options          *string   `json:"options,omitempty" xml:"options,omitempty"`
	Arch             *string   `json:"arch,omitempty" xml:"arch,omitempty"`
	FormatDisk       *bool     `json:"format_disk,omitempty" xml:"format_disk,omitempty"`
	NodepoolId       *string   `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	RdsInstances     []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttachScriptsBodyBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsBodyBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsBodyBody) SetKeepInstanceName(v bool) *DescribeClusterAttachScriptsBodyBody {
	s.KeepInstanceName = &v
	return s
}

func (s *DescribeClusterAttachScriptsBodyBody) SetOptions(v string) *DescribeClusterAttachScriptsBodyBody {
	s.Options = &v
	return s
}

func (s *DescribeClusterAttachScriptsBodyBody) SetArch(v string) *DescribeClusterAttachScriptsBodyBody {
	s.Arch = &v
	return s
}

func (s *DescribeClusterAttachScriptsBodyBody) SetFormatDisk(v bool) *DescribeClusterAttachScriptsBodyBody {
	s.FormatDisk = &v
	return s
}

func (s *DescribeClusterAttachScriptsBodyBody) SetNodepoolId(v string) *DescribeClusterAttachScriptsBodyBody {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterAttachScriptsBodyBody) SetRdsInstances(v []*string) *DescribeClusterAttachScriptsBodyBody {
	s.RdsInstances = v
	return s
}

type DescribeClusterAttachScriptsRequest struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *DescribeClusterAttachScriptsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAttachScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsRequest) SetHeaders(v map[string]*string) *DescribeClusterAttachScriptsRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterAttachScriptsRequest) SetBody(v *DescribeClusterAttachScriptsBody) *DescribeClusterAttachScriptsRequest {
	s.Body = v
	return s
}

type DescribeClusterAttachScriptsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterAttachScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAttachScriptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsResponse) SetHeaders(v map[string]*string) *DescribeClusterAttachScriptsResponse {
	s.Headers = v
	return s
}

type ScaleOutClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ScaleOutClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequest) SetHeaders(v map[string]*string) *ScaleOutClusterRequest {
	s.Headers = v
	return s
}

type ScaleOutClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
}

func (s ScaleOutClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponseBody) SetClusterId(v string) *ScaleOutClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetTaskId(v string) *ScaleOutClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetRequestId(v string) *ScaleOutClusterResponseBody {
	s.RequestId = &v
	return s
}

type ScaleOutClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleOutClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleOutClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponse) SetHeaders(v map[string]*string) *ScaleOutClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutClusterResponse) SetBody(v *ScaleOutClusterResponseBody) *ScaleOutClusterResponse {
	s.Body = v
	return s
}

type DescribeClusterResourcesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesRequest) SetHeaders(v map[string]*string) *DescribeClusterResourcesRequest {
	s.Headers = v
	return s
}

type DescribeClusterResourcesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponse) SetHeaders(v map[string]*string) *DescribeClusterResourcesResponse {
	s.Headers = v
	return s
}

type DescribeKubernetesVersionMetadataQuery struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	ClusterType       *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	Profile           *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s DescribeKubernetesVersionMetadataQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataQuery) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataQuery) SetRegion(v string) *DescribeKubernetesVersionMetadataQuery {
	s.Region = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataQuery) SetClusterType(v string) *DescribeKubernetesVersionMetadataQuery {
	s.ClusterType = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataQuery) SetKubernetesVersion(v string) *DescribeKubernetesVersionMetadataQuery {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataQuery) SetProfile(v string) *DescribeKubernetesVersionMetadataQuery {
	s.Profile = &v
	return s
}

type DescribeKubernetesVersionMetadataRequest struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeKubernetesVersionMetadataQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DescribeKubernetesVersionMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataRequest) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataRequest) SetHeaders(v map[string]*string) *DescribeKubernetesVersionMetadataRequest {
	s.Headers = v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetQuery(v *DescribeKubernetesVersionMetadataQuery) *DescribeKubernetesVersionMetadataRequest {
	s.Query = v
	return s
}

type DescribeKubernetesVersionMetadataResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeKubernetesVersionMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataResponse) SetHeaders(v map[string]*string) *DescribeKubernetesVersionMetadataResponse {
	s.Headers = v
	return s
}

type UpgradeClusterAddonsBody struct {
	Body []*UpgradeClusterAddonsBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpgradeClusterAddonsBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsBody) SetBody(v []*UpgradeClusterAddonsBodyBody) *UpgradeClusterAddonsBody {
	s.Body = v
	return s
}

type UpgradeClusterAddonsBodyBody struct {
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty" require:"true"`
	NextVersion   *string `json:"next_version,omitempty" xml:"next_version,omitempty" require:"true"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s UpgradeClusterAddonsBodyBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsBodyBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsBodyBody) SetComponentName(v string) *UpgradeClusterAddonsBodyBody {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterAddonsBodyBody) SetNextVersion(v string) *UpgradeClusterAddonsBodyBody {
	s.NextVersion = &v
	return s
}

func (s *UpgradeClusterAddonsBodyBody) SetVersion(v string) *UpgradeClusterAddonsBodyBody {
	s.Version = &v
	return s
}

type UpgradeClusterAddonsRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UpgradeClusterAddonsBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequest) SetHeaders(v map[string]*string) *UpgradeClusterAddonsRequest {
	s.Headers = v
	return s
}

func (s *UpgradeClusterAddonsRequest) SetBody(v *UpgradeClusterAddonsBody) *UpgradeClusterAddonsRequest {
	s.Body = v
	return s
}

type UpgradeClusterAddonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpgradeClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsResponse) SetHeaders(v map[string]*string) *UpgradeClusterAddonsResponse {
	s.Headers = v
	return s
}

type DescribeClusterAddonsVersionRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterAddonsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionRequest) SetHeaders(v map[string]*string) *DescribeClusterAddonsVersionRequest {
	s.Headers = v
	return s
}

type DescribeClusterAddonsVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterAddonsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsVersionResponse {
	s.Headers = v
	return s
}

type DescribeClusterAddonUpgradeStatusRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterAddonUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusRequest) SetHeaders(v map[string]*string) *DescribeClusterAddonUpgradeStatusRequest {
	s.Headers = v
	return s
}

type DescribeClusterAddonUpgradeStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterAddonUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonUpgradeStatusResponse {
	s.Headers = v
	return s
}

type DeleteClusterNodesBody struct {
	Body *DeleteClusterNodesBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s DeleteClusterNodesBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesBody) SetBody(v *DeleteClusterNodesBodyBody) *DeleteClusterNodesBody {
	s.Body = v
	return s
}

type DeleteClusterNodesBodyBody struct {
	ReleaseNode *bool     `json:"release_node,omitempty" xml:"release_node,omitempty"`
	DrainNode   *bool     `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	Nodes       []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s DeleteClusterNodesBodyBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesBodyBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesBodyBody) SetReleaseNode(v bool) *DeleteClusterNodesBodyBody {
	s.ReleaseNode = &v
	return s
}

func (s *DeleteClusterNodesBodyBody) SetDrainNode(v bool) *DeleteClusterNodesBodyBody {
	s.DrainNode = &v
	return s
}

func (s *DeleteClusterNodesBodyBody) SetNodes(v []*string) *DeleteClusterNodesBodyBody {
	s.Nodes = v
	return s
}

type DeleteClusterNodesRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *DeleteClusterNodesBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesRequest) SetHeaders(v map[string]*string) *DeleteClusterNodesRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodesRequest) SetBody(v *DeleteClusterNodesBody) *DeleteClusterNodesRequest {
	s.Body = v
	return s
}

type DeleteClusterNodesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponse) SetHeaders(v map[string]*string) *DeleteClusterNodesResponse {
	s.Headers = v
	return s
}

type UpdateTemplateBody struct {
	Body *UpdateTemplateBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateTemplateBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateBody) SetBody(v *UpdateTemplateBodyBody) *UpdateTemplateBody {
	s.Body = v
	return s
}

type UpdateTemplateBodyBody struct {
	Template     *string `json:"template,omitempty" xml:"template,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	Tags         *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s UpdateTemplateBodyBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateBodyBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateBodyBody) SetTemplate(v string) *UpdateTemplateBodyBody {
	s.Template = &v
	return s
}

func (s *UpdateTemplateBodyBody) SetName(v string) *UpdateTemplateBodyBody {
	s.Name = &v
	return s
}

func (s *UpdateTemplateBodyBody) SetDescription(v string) *UpdateTemplateBodyBody {
	s.Description = &v
	return s
}

func (s *UpdateTemplateBodyBody) SetTemplateType(v string) *UpdateTemplateBodyBody {
	s.TemplateType = &v
	return s
}

func (s *UpdateTemplateBodyBody) SetTags(v string) *UpdateTemplateBodyBody {
	s.Tags = &v
	return s
}

type UpdateTemplateRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UpdateTemplateBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetHeaders(v map[string]*string) *UpdateTemplateRequest {
	s.Headers = v
	return s
}

func (s *UpdateTemplateRequest) SetBody(v *UpdateTemplateBody) *UpdateTemplateRequest {
	s.Body = v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

type DeleteTemplateRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetHeaders(v map[string]*string) *DeleteTemplateRequest {
	s.Headers = v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

type DescribeClusterUserKubeconfigQuery struct {
	PrivateIpAddress         *bool  `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	TemporaryDurationMinutes *int64 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
}

func (s DescribeClusterUserKubeconfigQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigQuery) SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigQuery {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeClusterUserKubeconfigQuery) SetTemporaryDurationMinutes(v int64) *DescribeClusterUserKubeconfigQuery {
	s.TemporaryDurationMinutes = &v
	return s
}

type DescribeClusterUserKubeconfigRequest struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClusterUserKubeconfigQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClusterUserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigRequest) SetHeaders(v map[string]*string) *DescribeClusterUserKubeconfigRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) SetQuery(v *DescribeClusterUserKubeconfigQuery) *DescribeClusterUserKubeconfigRequest {
	s.Query = v
	return s
}

type DescribeClusterUserKubeconfigResponseBody struct {
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty" require:"true"`
	Config     *string `json:"config,omitempty" xml:"config,omitempty" require:"true"`
}

func (s DescribeClusterUserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetExpiration(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Expiration = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Config = &v
	return s
}

type DescribeClusterUserKubeconfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterUserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterUserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterUserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) SetBody(v *DescribeClusterUserKubeconfigResponseBody) *DescribeClusterUserKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeClusterNodesQuery struct {
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	NodepoolId  *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	PageSize    *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PageNumber  *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s DescribeClusterNodesQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesQuery) SetInstanceIds(v string) *DescribeClusterNodesQuery {
	s.InstanceIds = &v
	return s
}

func (s *DescribeClusterNodesQuery) SetNodepoolId(v string) *DescribeClusterNodesQuery {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesQuery) SetState(v string) *DescribeClusterNodesQuery {
	s.State = &v
	return s
}

func (s *DescribeClusterNodesQuery) SetPageSize(v string) *DescribeClusterNodesQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesQuery) SetPageNumber(v string) *DescribeClusterNodesQuery {
	s.PageNumber = &v
	return s
}

type DescribeClusterNodesRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClusterNodesQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesRequest) SetHeaders(v map[string]*string) *DescribeClusterNodesRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodesRequest) SetQuery(v *DescribeClusterNodesQuery) *DescribeClusterNodesRequest {
	s.Query = v
	return s
}

type DescribeClusterNodesResponseBody struct {
	Nodes []*DescribeClusterNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
	Page  *DescribeClusterNodesResponseBodyPage    `json:"page,omitempty" xml:"page,omitempty" require:"true" type:"Struct"`
}

func (s DescribeClusterNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBody) SetNodes(v []*DescribeClusterNodesResponseBodyNodes) *DescribeClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeClusterNodesResponseBody) SetPage(v *DescribeClusterNodesResponseBodyPage) *DescribeClusterNodesResponseBody {
	s.Page = v
	return s
}

type DescribeClusterNodesResponseBodyNodes struct {
	ErrorMessage       *string   `json:"error_message,omitempty" xml:"error_message,omitempty" require:"true"`
	CreationTime       *string   `json:"creation_time,omitempty" xml:"creation_time,omitempty" require:"true"`
	NodeStatus         *string   `json:"node_status,omitempty" xml:"node_status,omitempty" require:"true"`
	InstanceName       *string   `json:"instance_name,omitempty" xml:"instance_name,omitempty" require:"true"`
	IsAliyunNode       *bool     `json:"is_aliyun_node,omitempty" xml:"is_aliyun_node,omitempty" require:"true"`
	NodeName           *string   `json:"node_name,omitempty" xml:"node_name,omitempty" require:"true"`
	SpotStrategy       *string   `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty" require:"true"`
	ExpiredTime        *string   `json:"expired_time,omitempty" xml:"expired_time,omitempty" require:"true"`
	Source             *string   `json:"source,omitempty" xml:"source,omitempty" require:"true"`
	InstanceTypeFamily *string   `json:"instance_type_family,omitempty" xml:"instance_type_family,omitempty" require:"true"`
	InstanceId         *string   `json:"instance_id,omitempty" xml:"instance_id,omitempty" require:"true"`
	InstanceChargeType *string   `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty" require:"true"`
	InstanceRole       *string   `json:"instance_role,omitempty" xml:"instance_role,omitempty" require:"true"`
	State              *string   `json:"state,omitempty" xml:"state,omitempty" require:"true"`
	InstanceStatus     *string   `json:"instance_status,omitempty" xml:"instance_status,omitempty" require:"true"`
	ImageId            *string   `json:"image_id,omitempty" xml:"image_id,omitempty" require:"true"`
	NodepoolId         *string   `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty" require:"true"`
	InstanceType       *string   `json:"instance_type,omitempty" xml:"instance_type,omitempty" require:"true"`
	HostName           *string   `json:"host_name,omitempty" xml:"host_name,omitempty" require:"true"`
	IpAddress          []*string `json:"ip_address,omitempty" xml:"ip_address,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClusterNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyNodes) SetErrorMessage(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetCreationTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIsAliyunNode(v bool) *DescribeClusterNodesResponseBodyNodes {
	s.IsAliyunNode = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetSpotStrategy(v string) *DescribeClusterNodesResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetExpiredTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetSource(v string) *DescribeClusterNodesResponseBodyNodes {
	s.Source = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceTypeFamily(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceChargeType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceRole(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceRole = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetState(v string) *DescribeClusterNodesResponseBodyNodes {
	s.State = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetImageId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodepoolId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetHostName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIpAddress(v []*string) *DescribeClusterNodesResponseBodyNodes {
	s.IpAddress = v
	return s
}

type DescribeClusterNodesResponseBodyPage struct {
	PageNumber *int `json:"page_number,omitempty" xml:"page_number,omitempty" require:"true"`
	TotalCount *int `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	PageSize   *int `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s DescribeClusterNodesResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageNumber(v int) *DescribeClusterNodesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetTotalCount(v int) *DescribeClusterNodesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageSize(v int) *DescribeClusterNodesResponseBodyPage {
	s.PageSize = &v
	return s
}

type DescribeClusterNodesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponse) SetHeaders(v map[string]*string) *DescribeClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodesResponse) SetBody(v *DescribeClusterNodesResponseBody) *DescribeClusterNodesResponse {
	s.Body = v
	return s
}

type DescribeClusterLogsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsRequest) SetHeaders(v map[string]*string) *DescribeClusterLogsRequest {
	s.Headers = v
	return s
}

type DescribeClusterLogsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponse) SetHeaders(v map[string]*string) *DescribeClusterLogsResponse {
	s.Headers = v
	return s
}

type DescribeTaskInfoRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoRequest) SetHeaders(v map[string]*string) *DescribeTaskInfoRequest {
	s.Headers = v
	return s
}

type DescribeTaskInfoResponseBody struct {
	ClusterId  *string                                   `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	Created    *string                                   `json:"created,omitempty" xml:"created,omitempty" require:"true"`
	TaskId     *string                                   `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	State      *string                                   `json:"state,omitempty" xml:"state,omitempty" require:"true"`
	TaskType   *string                                   `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	Updated    *string                                   `json:"updated,omitempty" xml:"updated,omitempty" require:"true"`
	TaskResult []*DescribeTaskInfoResponseBodyTaskResult `json:"task_result,omitempty" xml:"task_result,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBody) SetClusterId(v string) *DescribeTaskInfoResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetCreated(v string) *DescribeTaskInfoResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskId(v string) *DescribeTaskInfoResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetState(v string) *DescribeTaskInfoResponseBody {
	s.State = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskType(v string) *DescribeTaskInfoResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetUpdated(v string) *DescribeTaskInfoResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskResult(v []*DescribeTaskInfoResponseBodyTaskResult) *DescribeTaskInfoResponseBody {
	s.TaskResult = v
	return s
}

type DescribeTaskInfoResponseBodyTaskResult struct {
	Data   *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s DescribeTaskInfoResponseBodyTaskResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBodyTaskResult) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBodyTaskResult) SetData(v string) *DescribeTaskInfoResponseBodyTaskResult {
	s.Data = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskResult) SetStatus(v string) *DescribeTaskInfoResponseBodyTaskResult {
	s.Status = &v
	return s
}

type DescribeTaskInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInfoResponse) SetBody(v *DescribeTaskInfoResponseBody) *DescribeTaskInfoResponse {
	s.Body = v
	return s
}

type AttachInstancesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetHeaders(v map[string]*string) *AttachInstancesRequest {
	s.Headers = v
	return s
}

type AttachInstancesResponseBody struct {
	TaskId *string                            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	List   []*AttachInstancesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetTaskId(v string) *AttachInstancesResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetList(v []*AttachInstancesResponseBodyList) *AttachInstancesResponseBody {
	s.List = v
	return s
}

type AttachInstancesResponseBodyList struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
	Code       *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AttachInstancesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBodyList) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBodyList) SetInstanceId(v string) *AttachInstancesResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetCode(v string) *AttachInstancesResponseBodyList {
	s.Code = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetMessage(v string) *AttachInstancesResponseBodyList {
	s.Message = &v
	return s
}

type AttachInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]*string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type DescribeTemplatesQuery struct {
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty" require:"true"`
	PageNum      *int64  `json:"page_num,omitempty" xml:"page_num,omitempty"`
	PageSize     *int64  `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s DescribeTemplatesQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesQuery) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesQuery) SetTemplateType(v string) *DescribeTemplatesQuery {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesQuery) SetPageNum(v int64) *DescribeTemplatesQuery {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesQuery) SetPageSize(v int64) *DescribeTemplatesQuery {
	s.PageSize = &v
	return s
}

type DescribeTemplatesRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeTemplatesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DescribeTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) SetHeaders(v map[string]*string) *DescribeTemplatesRequest {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesRequest) SetQuery(v *DescribeTemplatesQuery) *DescribeTemplatesRequest {
	s.Query = v
	return s
}

type DescribeTemplatesResponseBody struct {
	Templates []*DescribeTemplatesResponseBodyTemplates `json:"templates,omitempty" xml:"templates,omitempty" require:"true" type:"Repeated"`
	PageInfo  *DescribeTemplatesResponseBodyPageInfo    `json:"page_info,omitempty" xml:"page_info,omitempty" require:"true" type:"Struct"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) SetTemplates(v []*DescribeTemplatesResponseBodyTemplates) *DescribeTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeTemplatesResponseBody) SetPageInfo(v *DescribeTemplatesResponseBodyPageInfo) *DescribeTemplatesResponseBody {
	s.PageInfo = v
	return s
}

type DescribeTemplatesResponseBodyTemplates struct {
	Template           *string `json:"template,omitempty" xml:"template,omitempty" require:"true"`
	TemplateWithHistId *string `json:"template_with_hist_id,omitempty" xml:"template_with_hist_id,omitempty" require:"true"`
	Created            *string `json:"created,omitempty" xml:"created,omitempty" require:"true"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	TemplateType       *string `json:"template_type,omitempty" xml:"template_type,omitempty" require:"true"`
	Id                 *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	Acl                *string `json:"acl,omitempty" xml:"acl,omitempty" require:"true"`
	Updated            *string `json:"updated,omitempty" xml:"updated,omitempty" require:"true"`
	Tags               *string `json:"tags,omitempty" xml:"tags,omitempty" require:"true"`
}

func (s DescribeTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplate(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Template = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplateWithHistId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.TemplateWithHistId = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetCreated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Created = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetName(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetId(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetAcl(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Acl = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetUpdated(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Updated = &v
	return s
}

func (s *DescribeTemplatesResponseBodyTemplates) SetTags(v string) *DescribeTemplatesResponseBodyTemplates {
	s.Tags = &v
	return s
}

type DescribeTemplatesResponseBodyPageInfo struct {
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty" require:"true"`
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	PageSize   *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s DescribeTemplatesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageNumber(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetTotalCount(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeTemplatesResponseBodyPageInfo) SetPageSize(v int64) *DescribeTemplatesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

type DescribeTemplatesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesResponse) SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse {
	s.Body = v
	return s
}

type DescribeTemplateAttributeQuery struct {
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s DescribeTemplateAttributeQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAttributeQuery) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeQuery) SetTemplateType(v string) *DescribeTemplateAttributeQuery {
	s.TemplateType = &v
	return s
}

type DescribeTemplateAttributeRequest struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeTemplateAttributeQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeRequest) SetHeaders(v map[string]*string) *DescribeTemplateAttributeRequest {
	s.Headers = v
	return s
}

func (s *DescribeTemplateAttributeRequest) SetQuery(v *DescribeTemplateAttributeQuery) *DescribeTemplateAttributeRequest {
	s.Query = v
	return s
}

type DescribeTemplateAttributeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeResponse) SetHeaders(v map[string]*string) *DescribeTemplateAttributeResponse {
	s.Headers = v
	return s
}

type CreateTemplateBody struct {
	Body *CreateTemplateBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateTemplateBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateBody) SetBody(v *CreateTemplateBodyBody) *CreateTemplateBody {
	s.Body = v
	return s
}

type CreateTemplateBodyBody struct {
	Template     *string `json:"template,omitempty" xml:"template,omitempty" require:"true"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	Tags         *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s CreateTemplateBodyBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateBodyBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateBodyBody) SetTemplate(v string) *CreateTemplateBodyBody {
	s.Template = &v
	return s
}

func (s *CreateTemplateBodyBody) SetName(v string) *CreateTemplateBodyBody {
	s.Name = &v
	return s
}

func (s *CreateTemplateBodyBody) SetDescription(v string) *CreateTemplateBodyBody {
	s.Description = &v
	return s
}

func (s *CreateTemplateBodyBody) SetTemplateType(v string) *CreateTemplateBodyBody {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateBodyBody) SetTags(v string) *CreateTemplateBodyBody {
	s.Tags = &v
	return s
}

type CreateTemplateRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *CreateTemplateBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetHeaders(v map[string]*string) *CreateTemplateRequest {
	s.Headers = v
	return s
}

func (s *CreateTemplateRequest) SetBody(v *CreateTemplateBody) *CreateTemplateRequest {
	s.Body = v
	return s
}

type CreateTemplateResponseBody struct {
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetHeaders(v map[string]*string) *CreateClusterRequest {
	s.Headers = v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type ScaleClusterBody struct {
	Body *ScaleClusterBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ScaleClusterBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterBody) SetBody(v *ScaleClusterBodyBody) *ScaleClusterBody {
	s.Body = v
	return s
}

type ScaleClusterBodyBody struct {
	WorkerDataDisk           *bool                                  `json:"worker_data_disk,omitempty" xml:"worker_data_disk,omitempty"`
	KeyPair                  *string                                `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	WorkerSystemDiskCategory *string                                `json:"worker_system_disk_category,omitempty" xml:"worker_system_disk_category,omitempty"`
	Count                    *int64                                 `json:"count,omitempty" xml:"count,omitempty"`
	CloudMonitorFlags        *bool                                  `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	WorkerPeriodUnit         *string                                `json:"worker_period_unit,omitempty" xml:"worker_period_unit,omitempty"`
	WorkerAutoRenew          *bool                                  `json:"worker_auto_renew,omitempty" xml:"worker_auto_renew,omitempty"`
	WorkerAutoRenewPeriod    *int64                                 `json:"worker_auto_renew_period,omitempty" xml:"worker_auto_renew_period,omitempty"`
	WorkerPeriod             *int64                                 `json:"worker_period,omitempty" xml:"worker_period,omitempty"`
	LoginPassword            *string                                `json:"login_password,omitempty" xml:"login_password,omitempty"`
	WorkerSystemDiskSize     *int64                                 `json:"worker_system_disk_size,omitempty" xml:"worker_system_disk_size,omitempty"`
	CpuPolicy                *string                                `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	DisableRollback          *bool                                  `json:"disable_rollback,omitempty" xml:"disable_rollback,omitempty"`
	WorkerInstanceChargeType *string                                `json:"worker_instance_charge_type,omitempty" xml:"worker_instance_charge_type,omitempty"`
	Tags                     []*ScaleClusterBodyBodyTags            `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Taints                   []*ScaleClusterBodyBodyTaints          `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	WorkerDataDisks          []*ScaleClusterBodyBodyWorkerDataDisks `json:"worker_data_disks,omitempty" xml:"worker_data_disks,omitempty" type:"Repeated"`
	VswitchIds               []*string                              `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	WorkerInstanceTypes      []*string                              `json:"worker_instance_types,omitempty" xml:"worker_instance_types,omitempty" type:"Repeated"`
}

func (s ScaleClusterBodyBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterBodyBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterBodyBody) SetWorkerDataDisk(v bool) *ScaleClusterBodyBody {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleClusterBodyBody) SetKeyPair(v string) *ScaleClusterBodyBody {
	s.KeyPair = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerSystemDiskCategory(v string) *ScaleClusterBodyBody {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleClusterBodyBody) SetCount(v int64) *ScaleClusterBodyBody {
	s.Count = &v
	return s
}

func (s *ScaleClusterBodyBody) SetCloudMonitorFlags(v bool) *ScaleClusterBodyBody {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerPeriodUnit(v string) *ScaleClusterBodyBody {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerAutoRenew(v bool) *ScaleClusterBodyBody {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerAutoRenewPeriod(v int64) *ScaleClusterBodyBody {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerPeriod(v int64) *ScaleClusterBodyBody {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleClusterBodyBody) SetLoginPassword(v string) *ScaleClusterBodyBody {
	s.LoginPassword = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerSystemDiskSize(v int64) *ScaleClusterBodyBody {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleClusterBodyBody) SetCpuPolicy(v string) *ScaleClusterBodyBody {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleClusterBodyBody) SetDisableRollback(v bool) *ScaleClusterBodyBody {
	s.DisableRollback = &v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerInstanceChargeType(v string) *ScaleClusterBodyBody {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleClusterBodyBody) SetTags(v []*ScaleClusterBodyBodyTags) *ScaleClusterBodyBody {
	s.Tags = v
	return s
}

func (s *ScaleClusterBodyBody) SetTaints(v []*ScaleClusterBodyBodyTaints) *ScaleClusterBodyBody {
	s.Taints = v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerDataDisks(v []*ScaleClusterBodyBodyWorkerDataDisks) *ScaleClusterBodyBody {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleClusterBodyBody) SetVswitchIds(v []*string) *ScaleClusterBodyBody {
	s.VswitchIds = v
	return s
}

func (s *ScaleClusterBodyBody) SetWorkerInstanceTypes(v []*string) *ScaleClusterBodyBody {
	s.WorkerInstanceTypes = v
	return s
}

type ScaleClusterBodyBodyTags struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ScaleClusterBodyBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterBodyBodyTags) GoString() string {
	return s.String()
}

func (s *ScaleClusterBodyBodyTags) SetKey(v string) *ScaleClusterBodyBodyTags {
	s.Key = &v
	return s
}

type ScaleClusterBodyBodyTaints struct {
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
}

func (s ScaleClusterBodyBodyTaints) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterBodyBodyTaints) GoString() string {
	return s.String()
}

func (s *ScaleClusterBodyBodyTaints) SetKey(v string) *ScaleClusterBodyBodyTaints {
	s.Key = &v
	return s
}

func (s *ScaleClusterBodyBodyTaints) SetValue(v string) *ScaleClusterBodyBodyTaints {
	s.Value = &v
	return s
}

func (s *ScaleClusterBodyBodyTaints) SetEffect(v string) *ScaleClusterBodyBodyTaints {
	s.Effect = &v
	return s
}

type ScaleClusterBodyBodyWorkerDataDisks struct {
	Category  *string `json:"category,omitempty" xml:"category,omitempty"`
	Size      *string `json:"size,omitempty" xml:"size,omitempty"`
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
}

func (s ScaleClusterBodyBodyWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterBodyBodyWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleClusterBodyBodyWorkerDataDisks) SetCategory(v string) *ScaleClusterBodyBodyWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleClusterBodyBodyWorkerDataDisks) SetSize(v string) *ScaleClusterBodyBodyWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *ScaleClusterBodyBodyWorkerDataDisks) SetEncrypted(v string) *ScaleClusterBodyBodyWorkerDataDisks {
	s.Encrypted = &v
	return s
}

type ScaleClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *ScaleClusterBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequest) SetHeaders(v map[string]*string) *ScaleClusterRequest {
	s.Headers = v
	return s
}

func (s *ScaleClusterRequest) SetBody(v *ScaleClusterBody) *ScaleClusterRequest {
	s.Body = v
	return s
}

type ScaleClusterResponseBody struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty" require:"true"`
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
}

func (s ScaleClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponseBody) SetClusterId(v string) *ScaleClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetTaskId(v string) *ScaleClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleClusterResponseBody) SetRequestId(v string) *ScaleClusterResponseBody {
	s.RequestId = &v
	return s
}

type ScaleClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponse) SetHeaders(v map[string]*string) *ScaleClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterResponse) SetBody(v *ScaleClusterResponseBody) *ScaleClusterResponse {
	s.Body = v
	return s
}

type DescribeClustersQuery struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
}

func (s DescribeClustersQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersQuery) GoString() string {
	return s.String()
}

func (s *DescribeClustersQuery) SetName(v string) *DescribeClustersQuery {
	s.Name = &v
	return s
}

func (s *DescribeClustersQuery) SetClusterType(v string) *DescribeClustersQuery {
	s.ClusterType = &v
	return s
}

type DescribeClustersRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClustersQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetHeaders(v map[string]*string) *DescribeClustersRequest {
	s.Headers = v
	return s
}

func (s *DescribeClustersRequest) SetQuery(v *DescribeClustersQuery) *DescribeClustersRequest {
	s.Query = v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

type DescribeClusterDetailRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) SetHeaders(v map[string]*string) *DescribeClusterDetailRequest {
	s.Headers = v
	return s
}

type DescribeClusterDetailResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

type DeleteClusterQuery struct {
	RetainAllResources *bool     `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	KeepSlb            *bool     `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	RetainResources    []*string `json:"retain_resources,omitempty" xml:"retain_resources,omitempty" type:"Repeated"`
}

func (s DeleteClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterQuery) GoString() string {
	return s.String()
}

func (s *DeleteClusterQuery) SetRetainAllResources(v bool) *DeleteClusterQuery {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteClusterQuery) SetKeepSlb(v bool) *DeleteClusterQuery {
	s.KeepSlb = &v
	return s
}

func (s *DeleteClusterQuery) SetRetainResources(v []*string) *DeleteClusterQuery {
	s.RetainResources = v
	return s
}

type DeleteClusterRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteClusterQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetHeaders(v map[string]*string) *DeleteClusterRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterRequest) SetQuery(v *DeleteClusterQuery) *DeleteClusterRequest {
	s.Query = v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

type Client struct {
	roa.Client
}

func NewClient(config *roa.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *roa.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cs.aliyuncs.com"),
		"cn-fujian":                   tea.String("cs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cs-vpc.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cs-vpc.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cs-vpc.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cs.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("cs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.EndpointHost, _err = client.GetEndpoint(tea.String("cs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.EndpointHost)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) UpdateContactGroupForAlertWithOptions(clusterId *string, request *UpdateContactGroupForAlertRequest, runtime *util.RuntimeOptions) (_result *UpdateContactGroupForAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateContactGroupForAlertResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateContactGroupForAlert"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/alert/"+tea.StringValue(clusterId)+"/alert_rule/contact_groups"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContactGroupForAlert(clusterId *string, request *UpdateContactGroupForAlertRequest) (_result *UpdateContactGroupForAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContactGroupForAlertResponse{}
	_body, _err := client.UpdateContactGroupForAlertWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAlertWithOptions(clusterId *string, request *StopAlertRequest, runtime *util.RuntimeOptions) (_result *StopAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopAlertResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StopAlert"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/alert/"+tea.StringValue(clusterId)+"/alert_rule/stop"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAlert(clusterId *string, request *StopAlertRequest) (_result *StopAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAlertResponse{}
	_body, _err := client.StopAlertWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAlertWithOptions(clusterId *string, request *StartAlertRequest, runtime *util.RuntimeOptions) (_result *StartAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartAlertResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StartAlert"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/alert/"+tea.StringValue(clusterId)+"/alert_rule/start"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAlert(clusterId *string, request *StartAlertRequest) (_result *StartAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAlertResponse{}
	_body, _err := client.StartAlertWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactWithOptions(request *DeleteAlertContactRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteAlertContact"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/alert/contacts"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContact(request *DeleteAlertContactRequest) (_result *DeleteAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DeleteAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactGroupWithOptions(request *DeleteAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteAlertContactGroup"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/alert/contact_groups"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContactGroup(request *DeleteAlertContactGroupRequest) (_result *DeleteAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DeleteAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeEvents"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/events"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateClusterWithOptions(cluster_id *string, request *MigrateClusterRequest, runtime *util.RuntimeOptions) (_result *MigrateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &MigrateClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("MigrateCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(cluster_id)+"/migrate"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateCluster(cluster_id *string, request *MigrateClusterRequest) (_result *MigrateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateClusterResponse{}
	_body, _err := client.MigrateClusterWithOptions(cluster_id, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenAckServiceWithOptions(request *OpenAckServiceRequest, runtime *util.RuntimeOptions) (_result *OpenAckServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &OpenAckServiceResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("OpenAckService"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/service/open"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenAckService(request *OpenAckServiceRequest) (_result *OpenAckServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenAckServiceResponse{}
	_body, _err := client.OpenAckServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("TagResources"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/tags"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UntagResources"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/tags"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonsUpgradeStatusWithOptions(clusterId *string, request *DescribeClusterAddonsUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterAddonsUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/upgradestatus"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonsUpgradeStatus(clusterId *string, request *DescribeClusterAddonsUpgradeStatusRequest) (_result *DescribeClusterAddonsUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAddonsUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonsUpgradeStatusWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterConfigurationWithOptions(clusterId *string, request *ModifyClusterConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyClusterConfigurationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ModifyClusterConfiguration"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/configuration"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterConfiguration(clusterId *string, request *ModifyClusterConfigurationRequest) (_result *ModifyClusterConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterConfigurationResponse{}
	_body, _err := client.ModifyClusterConfigurationWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKubernetesTriggerWithOptions(id *string, request *DeleteKubernetesTriggerRequest, runtime *util.RuntimeOptions) (_result *DeleteKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/triggers/revoke/"+tea.StringValue(id)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKubernetesTrigger(id *string, request *DeleteKubernetesTriggerRequest) (_result *DeleteKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKubernetesTriggerResponse{}
	_body, _err := client.DeleteKubernetesTriggerWithOptions(id, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKubernetesTriggerWithOptions(request *CreateKubernetesTriggerRequest, runtime *util.RuntimeOptions) (_result *CreateKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/triggers"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKubernetesTrigger(request *CreateKubernetesTriggerRequest) (_result *CreateKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKubernetesTriggerResponse{}
	_body, _err := client.CreateKubernetesTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetKubernetesTriggerWithOptions(clusterId *string, request *GetKubernetesTriggerRequest, runtime *util.RuntimeOptions) (_result *GetKubernetesTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetKubernetesTrigger"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/triggers/"+tea.StringValue(clusterId)), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetKubernetesTrigger(clusterId *string, request *GetKubernetesTriggerRequest) (_result *GetKubernetesTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKubernetesTriggerResponse{}
	_body, _err := client.GetKubernetesTriggerWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListTagResources"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/tags"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeComponentUpgradeWithOptions(clusterid *string, componentid *string, request *ResumeComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *ResumeComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ResumeComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterid)+"/components/"+tea.StringValue(componentid)+"/resume"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeComponentUpgrade(clusterid *string, componentid *string, request *ResumeComponentUpgradeRequest) (_result *ResumeComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.ResumeComponentUpgradeWithOptions(clusterid, componentid, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseComponentUpgradeWithOptions(clusterid *string, componentid *string, request *PauseComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *PauseComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("PauseComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterid)+"/components/"+tea.StringValue(componentid)+"/pause"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseComponentUpgrade(clusterid *string, componentid *string, request *PauseComponentUpgradeRequest) (_result *PauseComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.PauseComponentUpgradeWithOptions(clusterid, componentid, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelComponentUpgradeWithOptions(clusterId *string, componentId *string, request *CancelComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *CancelComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CancelComponentUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/"+tea.StringValue(componentId)+"/cancel"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelComponentUpgrade(clusterId *string, componentId *string, request *CancelComponentUpgradeRequest) (_result *CancelComponentUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.CancelComponentUpgradeWithOptions(clusterId, componentId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterNodepoolWithOptions(clusterId *string, nodepoolId *string, request *DeleteClusterNodepoolRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterNodepoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterNodepoolResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteClusterNodepool"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools/"+tea.StringValue(nodepoolId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterNodepool(clusterId *string, nodepoolId *string, request *DeleteClusterNodepoolRequest) (_result *DeleteClusterNodepoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterNodepoolResponse{}
	_body, _err := client.DeleteClusterNodepoolWithOptions(clusterId, nodepoolId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleClusterNodePoolWithOptions(clusterId *string, nodepoolId *string, request *ScaleClusterNodePoolRequest, runtime *util.RuntimeOptions) (_result *ScaleClusterNodePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleClusterNodePoolResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleClusterNodePool"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools/"+tea.StringValue(nodepoolId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleClusterNodePool(clusterId *string, nodepoolId *string, request *ScaleClusterNodePoolRequest) (_result *ScaleClusterNodePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleClusterNodePoolResponse{}
	_body, _err := client.ScaleClusterNodePoolWithOptions(clusterId, nodepoolId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodePoolDetailWithOptions(clusterId *string, nodepoolId *string, request *DescribeClusterNodePoolDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodePoolDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodePoolDetailResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodePoolDetail"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools/"+tea.StringValue(nodepoolId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodePoolDetail(clusterId *string, nodepoolId *string, request *DescribeClusterNodePoolDetailRequest) (_result *DescribeClusterNodePoolDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodePoolDetailResponse{}
	_body, _err := client.DescribeClusterNodePoolDetailWithOptions(clusterId, nodepoolId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodePoolsWithOptions(clusterId *string, request *DescribeClusterNodePoolsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodePoolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodePoolsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodePools"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodePools(clusterId *string, request *DescribeClusterNodePoolsRequest) (_result *DescribeClusterNodePoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodePoolsResponse{}
	_body, _err := client.DescribeClusterNodePoolsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterNodePoolWithOptions(clusterId *string, request *CreateClusterNodePoolRequest, runtime *util.RuntimeOptions) (_result *CreateClusterNodePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateClusterNodePoolResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateClusterNodePool"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterNodePool(clusterId *string, request *CreateClusterNodePoolRequest) (_result *CreateClusterNodePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterNodePoolResponse{}
	_body, _err := client.CreateClusterNodePoolWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterNodePoolWithOptions(clusterId *string, nodepoolId *string, request *ModifyClusterNodePoolRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterNodePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyClusterNodePoolResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ModifyClusterNodePool"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodepools/"+tea.StringValue(nodepoolId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterNodePool(clusterId *string, nodepoolId *string, request *ModifyClusterNodePoolRequest) (_result *ModifyClusterNodePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterNodePoolResponse{}
	_body, _err := client.ModifyClusterNodePoolWithOptions(clusterId, nodepoolId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelWorkflowWithOptions(workflowName *string, request *CancelWorkflowRequest, runtime *util.RuntimeOptions) (_result *CancelWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelWorkflowResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CancelWorkflow"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/gs/workflow/"+tea.StringValue(workflowName)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelWorkflow(workflowName *string, request *CancelWorkflowRequest) (_result *CancelWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelWorkflowResponse{}
	_body, _err := client.CancelWorkflowWithOptions(workflowName, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescirbeWorkflowWithOptions(workflowName *string, request *DescirbeWorkflowRequest, runtime *util.RuntimeOptions) (_result *DescirbeWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescirbeWorkflowResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescirbeWorkflow"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/gs/workflow/"+tea.StringValue(workflowName)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescirbeWorkflow(workflowName *string, request *DescirbeWorkflowRequest) (_result *DescirbeWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescirbeWorkflowResponse{}
	_body, _err := client.DescirbeWorkflowWithOptions(workflowName, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveWorkflowWithOptions(workflowName *string, request *RemoveWorkflowRequest, runtime *util.RuntimeOptions) (_result *RemoveWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveWorkflowResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RemoveWorkflow"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/gs/workflow/"+tea.StringValue(workflowName)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveWorkflow(workflowName *string, request *RemoveWorkflowRequest) (_result *RemoveWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveWorkflowResponse{}
	_body, _err := client.RemoveWorkflowWithOptions(workflowName, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWorkflowsWithOptions(request *DescribeWorkflowsRequest, runtime *util.RuntimeOptions) (_result *DescribeWorkflowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeWorkflowsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeWorkflows"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/gs/workflows"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWorkflows(request *DescribeWorkflowsRequest) (_result *DescribeWorkflowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWorkflowsResponse{}
	_body, _err := client.DescribeWorkflowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartWorkflowWithOptions(request *StartWorkflowRequest, runtime *util.RuntimeOptions) (_result *StartWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartWorkflowResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StartWorkflow"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/gs/workflow"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartWorkflow(request *StartWorkflowRequest) (_result *StartWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartWorkflowResponse{}
	_body, _err := client.StartWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserPermissionWithOptions(uid *string, request *DescribeUserPermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserPermissionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeUserPermission"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/permissions/users/"+tea.StringValue(uid)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserPermission(uid *string, request *DescribeUserPermissionRequest) (_result *DescribeUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserPermissionResponse{}
	_body, _err := client.DescribeUserPermissionWithOptions(uid, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantPermissionsWithOptions(uid *string, request *GrantPermissionsRequest, runtime *util.RuntimeOptions) (_result *GrantPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GrantPermissionsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GrantPermissions"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/permissions/users/"+tea.StringValue(uid)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantPermissions(uid *string, request *GrantPermissionsRequest) (_result *GrantPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantPermissionsResponse{}
	_body, _err := client.GrantPermissionsWithOptions(uid, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnInstallClusterAddonsWithOptions(clusterId *string, request *UnInstallClusterAddonsRequest, runtime *util.RuntimeOptions) (_result *UnInstallClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UnInstallClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/uninstall"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnInstallClusterAddons(clusterId *string, request *UnInstallClusterAddonsRequest) (_result *UnInstallClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.UnInstallClusterAddonsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAddonsWithOptions(request *DescribeAddonsRequest, runtime *util.RuntimeOptions) (_result *DescribeAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/components/metadata"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAddons(request *DescribeAddonsRequest) (_result *DescribeAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DescribeAddonsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateK8sClusterUserConfigExpireWithOptions(clusterId *string, request *UpdateK8sClusterUserConfigExpireRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateK8sClusterUserConfigExpireResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateK8sClusterUserConfigExpire"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(clusterId)+"/user_config/expire"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sClusterUserConfigExpire(clusterId *string, request *UpdateK8sClusterUserConfigExpireRequest) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateK8sClusterUserConfigExpireResponse{}
	_body, _err := client.UpdateK8sClusterUserConfigExpireWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelClusterUpgradeWithOptions(clusterId *string, request *CancelClusterUpgradeRequest, runtime *util.RuntimeOptions) (_result *CancelClusterUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CancelClusterUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/upgrade/cancel"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelClusterUpgrade(clusterId *string, request *CancelClusterUpgradeRequest) (_result *CancelClusterUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelClusterUpgradeResponse{}
	_body, _err := client.CancelClusterUpgradeWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersV1WithOptions(request *DescribeClustersV1Request, runtime *util.RuntimeOptions) (_result *DescribeClustersV1Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClustersV1Response{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClustersV1"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/clusters"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClustersV1(request *DescribeClustersV1Request) (_result *DescribeClustersV1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersV1Response{}
	_body, _err := client.DescribeClustersV1WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserQuotaWithOptions(request *DescribeUserQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeUserQuota"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/quota"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserQuota(request *DescribeUserQuotaRequest) (_result *DescribeUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.DescribeUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterV2UserKubeconfigWithOptions(clusterId *string, request *DescribeClusterV2UserKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterV2UserKubeconfig"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v2/k8s/"+tea.StringValue(clusterId)+"/user_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterV2UserKubeconfig(clusterId *string, request *DescribeClusterV2UserKubeconfigRequest) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DescribeClusterV2UserKubeconfigWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveClusterNodesWithOptions(clusterId *string, request *RemoveClusterNodesRequest, runtime *util.RuntimeOptions) (_result *RemoveClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RemoveClusterNodes"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/nodes/remove"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveClusterNodes(clusterId *string, request *RemoveClusterNodesRequest) (_result *RemoveClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClusterNodesResponse{}
	_body, _err := client.RemoveClusterNodesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterWithOptions(clusterId *string, request *UpgradeClusterRequest, runtime *util.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpgradeCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/upgrade"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeCluster(clusterId *string, request *UpgradeClusterRequest) (_result *UpgradeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.UpgradeClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseClusterUpgradeWithOptions(clusterId *string, request *PauseClusterUpgradeRequest, runtime *util.RuntimeOptions) (_result *PauseClusterUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PauseClusterUpgradeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("PauseClusterUpgrade"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/upgrade/pause"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseClusterUpgrade(clusterId *string, request *PauseClusterUpgradeRequest) (_result *PauseClusterUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseClusterUpgradeResponse{}
	_body, _err := client.PauseClusterUpgradeWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeUpgradeClusterWithOptions(clusterId *string, request *ResumeUpgradeClusterRequest, runtime *util.RuntimeOptions) (_result *ResumeUpgradeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResumeUpgradeClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ResumeUpgradeCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/upgrade/resume"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeUpgradeCluster(clusterId *string, request *ResumeUpgradeClusterRequest) (_result *ResumeUpgradeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeUpgradeClusterResponse{}
	_body, _err := client.ResumeUpgradeClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUpgradeStatusWithOptions(clusterId *string, request *GetUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *GetUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)+"/upgrade/status"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUpgradeStatus(clusterId *string, request *GetUpgradeStatusRequest) (_result *GetUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.GetUpgradeStatusWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterWithOptions(clusterId *string, request *ModifyClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ModifyCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCluster(clusterId *string, request *ModifyClusterRequest) (_result *ModifyClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterResponse{}
	_body, _err := client.ModifyClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallClusterAddonsWithOptions(clusterId *string, request *InstallClusterAddonsRequest, runtime *util.RuntimeOptions) (_result *InstallClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InstallClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/install"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallClusterAddons(clusterId *string, request *InstallClusterAddonsRequest) (_result *InstallClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallClusterAddonsResponse{}
	_body, _err := client.InstallClusterAddonsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterTagsWithOptions(clusterId *string, request *ModifyClusterTagsRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ModifyClusterTags"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/tags"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterTags(clusterId *string, request *ModifyClusterTagsRequest) (_result *ModifyClusterTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.ModifyClusterTagsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNamespacesWithOptions(clusterId *string, request *DescribeClusterNamespacesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNamespacesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNamespaces"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(clusterId)+"/namespaces"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNamespaces(clusterId *string, request *DescribeClusterNamespacesRequest) (_result *DescribeClusterNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNamespacesResponse{}
	_body, _err := client.DescribeClusterNamespacesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExternalAgentWithOptions(clusterId *string, request *DescribeExternalAgentRequest, runtime *util.RuntimeOptions) (_result *DescribeExternalAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeExternalAgent"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(clusterId)+"/external/agent/deployment"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExternalAgent(clusterId *string, request *DescribeExternalAgentRequest) (_result *DescribeExternalAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExternalAgentResponse{}
	_body, _err := client.DescribeExternalAgentWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAttachScriptsWithOptions(clusterId *string, request *DescribeClusterAttachScriptsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterAttachScripts"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/attachscript"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAttachScripts(clusterId *string, request *DescribeClusterAttachScriptsRequest) (_result *DescribeClusterAttachScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAttachScriptsResponse{}
	_body, _err := client.DescribeClusterAttachScriptsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleOutClusterWithOptions(clusterId *string, request *ScaleOutClusterRequest, runtime *util.RuntimeOptions) (_result *ScaleOutClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleOutCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v2/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleOutCluster(clusterId *string, request *ScaleOutClusterRequest) (_result *ScaleOutClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.ScaleOutClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterResourcesWithOptions(clusterId *string, request *DescribeClusterResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterResources"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/resources"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterResources(clusterId *string, request *DescribeClusterResourcesRequest) (_result *DescribeClusterResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResourcesResponse{}
	_body, _err := client.DescribeClusterResourcesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKubernetesVersionMetadataWithOptions(request *DescribeKubernetesVersionMetadataRequest, runtime *util.RuntimeOptions) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKubernetesVersionMetadataResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeKubernetesVersionMetadata"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/metadata/versions"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKubernetesVersionMetadata(request *DescribeKubernetesVersionMetadataRequest) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKubernetesVersionMetadataResponse{}
	_body, _err := client.DescribeKubernetesVersionMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterAddonsWithOptions(clusterId *string, request *UpgradeClusterAddonsRequest, runtime *util.RuntimeOptions) (_result *UpgradeClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpgradeClusterAddons"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/upgrade"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClusterAddons(clusterId *string, request *UpgradeClusterAddonsRequest) (_result *UpgradeClusterAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClusterAddonsResponse{}
	_body, _err := client.UpgradeClusterAddonsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonsVersionWithOptions(clusterId *string, request *DescribeClusterAddonsVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterAddonsVersion"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/version"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonsVersion(clusterId *string, request *DescribeClusterAddonsVersionRequest) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DescribeClusterAddonsVersionWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAddonUpgradeStatusWithOptions(clusterId *string, componentId *string, request *DescribeClusterAddonUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterAddonUpgradeStatus"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/components/"+tea.StringValue(componentId)+"/upgradestatus"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonUpgradeStatus(clusterId *string, componentId *string, request *DescribeClusterAddonUpgradeStatusRequest) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DescribeClusterAddonUpgradeStatusWithOptions(clusterId, componentId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterNodesWithOptions(clusterId *string, request *DeleteClusterNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterNodesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteClusterNodes"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodes"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterNodes(clusterId *string, request *DeleteClusterNodesRequest) (_result *DeleteClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterNodesResponse{}
	_body, _err := client.DeleteClusterNodesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(templateId *string, request *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateTemplate"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/templates/"+tea.StringValue(templateId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(templateId *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(templateId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(templateId *string, request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteTemplate"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/templates/"+tea.StringValue(templateId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(templateId *string, request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(templateId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterUserKubeconfigWithOptions(clusterId *string, request *DescribeClusterUserKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterUserKubeconfig"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/"+tea.StringValue(clusterId)+"/user_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterUserKubeconfig(clusterId *string, request *DescribeClusterUserKubeconfigRequest) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DescribeClusterUserKubeconfigWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodesWithOptions(clusterId *string, request *DescribeClusterNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodes"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodes"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodes(clusterId *string, request *DescribeClusterNodesRequest) (_result *DescribeClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DescribeClusterNodesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterLogsWithOptions(clusterId *string, request *DescribeClusterLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterLogs"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/logs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterLogs(clusterId *string, request *DescribeClusterLogsRequest) (_result *DescribeClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DescribeClusterLogsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskInfoWithOptions(task_id *string, request *DescribeTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTaskInfo"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/tasks/"+tea.StringValue(task_id)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTaskInfo(task_id *string, request *DescribeTaskInfoRequest) (_result *DescribeTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DescribeTaskInfoWithOptions(task_id, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachInstancesWithOptions(clusterId *string, request *AttachInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AttachInstances"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/attach"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachInstances(clusterId *string, request *AttachInstancesRequest) (_result *AttachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachInstancesResponse{}
	_body, _err := client.AttachInstancesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTemplates"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/templates"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateAttributeWithOptions(templateId *string, request *DescribeTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTemplateAttribute"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/templates/"+tea.StringValue(templateId)), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateAttribute(templateId *string, request *DescribeTemplateAttributeRequest) (_result *DescribeTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.DescribeTemplateAttributeWithOptions(templateId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateTemplate"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/templates"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleClusterWithOptions(clusterId *string, request *ScaleClusterRequest, runtime *util.RuntimeOptions) (_result *ScaleClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleCluster(clusterId *string, request *ScaleClusterRequest) (_result *ScaleClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleClusterResponse{}
	_body, _err := client.ScaleClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusters"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDetailWithOptions(clusterId *string, request *DescribeClusterDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterDetail"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDetail(clusterId *string, request *DescribeClusterDetailRequest) (_result *DescribeClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DescribeClusterDetailWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(clusterId *string, request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteCluster"), tea.String("2015-12-15"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(clusterId *string, request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
