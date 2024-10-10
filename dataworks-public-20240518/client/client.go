// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbolishDeploymentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1606087c-9ac4-43f0-83a8-0b5ced21XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AbolishDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentRequest) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentRequest) SetId(v string) *AbolishDeploymentRequest {
	s.Id = &v
	return s
}

func (s *AbolishDeploymentRequest) SetProjectId(v string) *AbolishDeploymentRequest {
	s.ProjectId = &v
	return s
}

type AbolishDeploymentResponseBody struct {
	// example:
	//
	// 55D786C9-DD57-524D-884C-C5239278XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbolishDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponseBody) SetRequestId(v string) *AbolishDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishDeploymentResponseBody) SetSuccess(v bool) *AbolishDeploymentResponseBody {
	s.Success = &v
	return s
}

type AbolishDeploymentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponse) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponse) SetHeaders(v map[string]*string) *AbolishDeploymentResponse {
	s.Headers = v
	return s
}

func (s *AbolishDeploymentResponse) SetStatusCode(v int32) *AbolishDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishDeploymentResponse) SetBody(v *AbolishDeploymentResponseBody) *AbolishDeploymentResponse {
	s.Body = v
	return s
}

type CloneDataSourceRequest struct {
	// example:
	//
	// demo_holo_datasource
	CloneDataSourceName *string `json:"CloneDataSourceName,omitempty" xml:"CloneDataSourceName,omitempty"`
	// example:
	//
	// 16036
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CloneDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CloneDataSourceRequest) SetCloneDataSourceName(v string) *CloneDataSourceRequest {
	s.CloneDataSourceName = &v
	return s
}

func (s *CloneDataSourceRequest) SetId(v int64) *CloneDataSourceRequest {
	s.Id = &v
	return s
}

type CloneDataSourceResponseBody struct {
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// FCD583B9-346B-5E75-82C1-4A7C192C48DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponseBody) SetId(v int64) *CloneDataSourceResponseBody {
	s.Id = &v
	return s
}

func (s *CloneDataSourceResponseBody) SetRequestId(v string) *CloneDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CloneDataSourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponse) SetHeaders(v map[string]*string) *CloneDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CloneDataSourceResponse) SetStatusCode(v int32) *CloneDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneDataSourceResponse) SetBody(v *CloneDataSourceResponseBody) *CloneDataSourceResponse {
	s.Body = v
	return s
}

type CreateDIJobRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationDataSourceSettings []*CreateDIJobRequestDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName     *string                        `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettings *CreateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	ResourceSettings *CreateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// This parameter is required.
	SourceDataSourceSettings []*CreateDIJobRequestSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// This parameter is required.
	TableMappings       []*CreateDIJobRequestTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules []*CreateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequest) SetDescription(v string) *CreateDIJobRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceSettings(v []*CreateDIJobRequestDestinationDataSourceSettings) *CreateDIJobRequest {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceType(v string) *CreateDIJobRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetJobName(v string) *CreateDIJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobRequest) SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *CreateDIJobRequest) SetMigrationType(v string) *CreateDIJobRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobRequest) SetProjectId(v int64) *CreateDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobRequest) SetResourceSettings(v *CreateDIJobRequestResourceSettings) *CreateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceSettings(v []*CreateDIJobRequestSourceDataSourceSettings) *CreateDIJobRequest {
	s.SourceDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceType(v string) *CreateDIJobRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetTableMappings(v []*CreateDIJobRequestTableMappings) *CreateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *CreateDIJobRequest) SetTransformationRules(v []*CreateDIJobRequestTransformationRules) *CreateDIJobRequest {
	s.TransformationRules = v
	return s
}

type CreateDIJobRequestDestinationDataSourceSettings struct {
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s CreateDIJobRequestDestinationDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

type CreateDIJobRequestJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*CreateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *CreateDIJobRequestJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*CreateDIJobRequestJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*CreateDIJobRequestJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestJobSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettings) SetChannelSettings(v string) *CreateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*CreateDIJobRequestJobSettingsColumnDataTypeSettings) *CreateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetCycleScheduleSettings(v *CreateDIJobRequestJobSettingsCycleScheduleSettings) *CreateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*CreateDIJobRequestJobSettingsDdlHandlingSettings) *CreateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetRuntimeSettings(v []*CreateDIJobRequestJobSettingsRuntimeSettings) *CreateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

type CreateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type CreateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// Full
	CycleMigrationType *string `json:"CycleMigrationType,omitempty" xml:"CycleMigrationType,omitempty"`
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type CreateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type CreateDIJobRequestJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type CreateDIJobRequestResourceSettings struct {
	OfflineResourceSettings  *CreateDIJobRequestResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *CreateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *CreateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *CreateDIJobRequestResourceSettingsOfflineResourceSettings) *CreateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) *CreateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *CreateDIJobRequestResourceSettingsScheduleResourceSettings) *CreateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type CreateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestSourceDataSourceSettings struct {
	// example:
	//
	// mysql_datasource_1
	DataSourceName       *string                                                         `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceProperties *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestSourceDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceProperties(v *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

type CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties struct {
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

type CreateDIJobRequestTableMappings struct {
	SourceObjectSelectionRules []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*CreateDIJobRequestTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestTableMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules) *CreateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *CreateDIJobRequestTableMappings) SetTransformationRules(v []*CreateDIJobRequestTableMappingsTransformationRules) *CreateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

type CreateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type CreateDIJobRequestTableMappingsTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s CreateDIJobRequestTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type CreateDIJobRequestTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s CreateDIJobRequestTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleExpression(v string) *CreateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleName(v string) *CreateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

type CreateDIJobShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationDataSourceSettingsShrink *string `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName           *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// This parameter is required.
	SourceDataSourceSettingsShrink *string `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// This parameter is required.
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s CreateDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobShrinkRequest) SetDescription(v string) *CreateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobName(v string) *CreateDIJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetMigrationType(v string) *CreateDIJobShrinkRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetProjectId(v int64) *CreateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

type CreateDIJobResponseBody struct {
	// example:
	//
	// 11792
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// 4F6AB6B3-41FB-5EBB-AFB2-0C98D49DA2BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponseBody) SetDIJobId(v int64) *CreateDIJobResponseBody {
	s.DIJobId = &v
	return s
}

func (s *CreateDIJobResponseBody) SetRequestId(v string) *CreateDIJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponse) SetHeaders(v map[string]*string) *CreateDIJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDIJobResponse) SetStatusCode(v int32) *CreateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDIJobResponse) SetBody(v *CreateDIJobResponseBody) *CreateDIJobResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// this is a holo datasource
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_datasource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetConnectionProperties(v string) *CreateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *CreateDataSourceRequest) SetConnectionPropertiesMode(v string) *CreateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetProjectId(v int64) *CreateDataSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

type CreateDataSourceResponseBody struct {
	// example:
	//
	// 22130
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// B62EC203-B39E-5DC1-B5B8-EB3C61707009
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetId(v int64) *CreateDataSourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateDataSourceSharedRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 144544
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 1107550004253538
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 106560
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s CreateDataSourceSharedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleRequest) SetDataSourceId(v int64) *CreateDataSourceSharedRuleRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetEnvType(v string) *CreateDataSourceSharedRuleRequest {
	s.EnvType = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetSharedUser(v string) *CreateDataSourceSharedRuleRequest {
	s.SharedUser = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetTargetProjectId(v int64) *CreateDataSourceSharedRuleRequest {
	s.TargetProjectId = &v
	return s
}

type CreateDataSourceSharedRuleResponseBody struct {
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 46F594E6-84AB-5FA5-8144-6F3D149961E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceSharedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponseBody) SetId(v int64) *CreateDataSourceSharedRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponseBody) SetRequestId(v string) *CreateDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceSharedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *CreateDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetStatusCode(v int32) *CreateDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetBody(v *CreateDataSourceSharedRuleResponseBody) *CreateDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

type CreateDeploymentRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// 项目Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentRequest) SetDescription(v string) *CreateDeploymentRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentRequest) SetObjectIds(v []*string) *CreateDeploymentRequest {
	s.ObjectIds = v
	return s
}

func (s *CreateDeploymentRequest) SetProjectId(v string) *CreateDeploymentRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentRequest) SetType(v string) *CreateDeploymentRequest {
	s.Type = &v
	return s
}

type CreateDeploymentShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// 项目Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentShrinkRequest) SetDescription(v string) *CreateDeploymentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetObjectIdsShrink(v string) *CreateDeploymentShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetProjectId(v string) *CreateDeploymentShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetType(v string) *CreateDeploymentShrinkRequest {
	s.Type = &v
	return s
}

type CreateDeploymentResponseBody struct {
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponseBody) SetId(v string) *CreateDeploymentResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetRequestId(v string) *CreateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponse) SetHeaders(v map[string]*string) *CreateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentResponse) SetStatusCode(v int32) *CreateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentResponse) SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetProjectId(v string) *CreateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFunctionRequest) SetSpec(v string) *CreateFunctionRequest {
	s.Spec = &v
	return s
}

type CreateFunctionResponseBody struct {
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// AE49C88D-5BEE-5ADD-8B8C-C4BBC0D7XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetId(v string) *CreateFunctionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

type CreateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateNodeRequest struct {
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) SetContainerId(v string) *CreateNodeRequest {
	s.ContainerId = &v
	return s
}

func (s *CreateNodeRequest) SetProjectId(v string) *CreateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeRequest) SetScene(v string) *CreateNodeRequest {
	s.Scene = &v
	return s
}

func (s *CreateNodeRequest) SetSpec(v string) *CreateNodeRequest {
	s.Spec = &v
	return s
}

type CreateNodeResponseBody struct {
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) SetId(v string) *CreateNodeResponseBody {
	s.Id = &v
	return s
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

type CreateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeResponse) SetHeaders(v map[string]*string) *CreateNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeResponse) SetStatusCode(v int32) *CreateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeResponse) SetBody(v *CreateNodeResponseBody) *CreateNodeResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string                                   `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*CreateProjectRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetAliyunResourceGroupId(v string) *CreateProjectRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectRequest) SetAliyunResourceTags(v []*CreateProjectRequestAliyunResourceTags) *CreateProjectRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectRequest) SetDevRoleDisabled(v bool) *CreateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectRequest) SetDisplayName(v string) *CreateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetPaiTaskEnabled(v bool) *CreateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

type CreateProjectRequestAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProjectRequestAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestAliyunResourceTags) SetKey(v string) *CreateProjectRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *CreateProjectRequestAliyunResourceTags) SetValue(v string) *CreateProjectRequestAliyunResourceTags {
	s.Value = &v
	return s
}

type CreateProjectShrinkRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId    *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceGroupId(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceTagsShrink(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevRoleDisabled(v bool) *CreateProjectShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDisplayName(v string) *CreateProjectShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetName(v string) *CreateProjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetPaiTaskEnabled(v bool) *CreateProjectShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

type CreateProjectResponseBody struct {
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v int64) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// 资源文件的项目id
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetProjectId(v string) *CreateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceRequest) SetSpec(v string) *CreateResourceRequest {
	s.Spec = &v
	return s
}

type CreateResourceResponseBody struct {
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// A5B97987-66EA-5563-9599-A2752292XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetId(v string) *CreateResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type CreateWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionRequest) SetProjectId(v string) *CreateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowDefinitionRequest) SetSpec(v string) *CreateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type CreateWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0EF298E5-0940-5AC7-9CB0-65025070XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponseBody) SetId(v string) *CreateWorkflowDefinitionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWorkflowDefinitionResponseBody) SetRequestId(v string) *CreateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

type CreateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *CreateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetStatusCode(v int32) *CreateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetBody(v *CreateWorkflowDefinitionResponseBody) *CreateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type DeleteDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11126
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
}

func (s DeleteDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIJobRequest) SetDIJobId(v int64) *DeleteDIJobRequest {
	s.DIJobId = &v
	return s
}

type DeleteDIJobResponseBody struct {
	// example:
	//
	// D33D4A51-5845-579A-B4BA-FAADD0F83D53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponseBody) SetRequestId(v string) *DeleteDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDIJobResponseBody) SetSuccess(v bool) *DeleteDIJobResponseBody {
	s.Success = &v
	return s
}

type DeleteDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponse) SetHeaders(v map[string]*string) *DeleteDIJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDIJobResponse) SetStatusCode(v int32) *DeleteDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDIJobResponse) SetBody(v *DeleteDIJobResponseBody) *DeleteDIJobResponse {
	s.Body = v
	return s
}

type DeleteDataSourceRequest struct {
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) SetId(v int64) *DeleteDataSourceRequest {
	s.Id = &v
	return s
}

type DeleteDataSourceResponseBody struct {
	// example:
	//
	// B56432E0-2112-5C97-88D0-AA0AE5C75C74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetSuccess(v bool) *DeleteDataSourceResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteDataSourceSharedRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 22127
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataSourceSharedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleRequest) SetId(v int64) *DeleteDataSourceSharedRuleRequest {
	s.Id = &v
	return s
}

type DeleteDataSourceSharedRuleResponseBody struct {
	// example:
	//
	// 64B-587A-8CED-969E1973887FXXX-TT
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetRequestId(v string) *DeleteDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetSuccess(v bool) *DeleteDataSourceSharedRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *DeleteDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetStatusCode(v int32) *DeleteDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetBody(v *DeleteDataSourceSharedRuleResponseBody) *DeleteDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetId(v string) *DeleteFunctionRequest {
	s.Id = &v
	return s
}

func (s *DeleteFunctionRequest) SetProjectId(v string) *DeleteFunctionRequest {
	s.ProjectId = &v
	return s
}

type DeleteFunctionResponseBody struct {
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetSuccess(v bool) *DeleteFunctionResponseBody {
	s.Success = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
	s.Body = v
	return s
}

type DeleteNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) SetId(v string) *DeleteNodeRequest {
	s.Id = &v
	return s
}

func (s *DeleteNodeRequest) SetProjectId(v string) *DeleteNodeRequest {
	s.ProjectId = &v
	return s
}

type DeleteNodeResponseBody struct {
	// example:
	//
	// A1E54497-5122-505E-91C6-BAC14980XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetSuccess(v bool) *DeleteNodeResponseBody {
	s.Success = &v
	return s
}

type DeleteNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponse) SetHeaders(v map[string]*string) *DeleteNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeResponse) SetStatusCode(v int32) *DeleteNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeResponse) SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetId(v int64) *DeleteProjectRequest {
	s.Id = &v
	return s
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetId(v string) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v string) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

type DeleteResourceResponseBody struct {
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetSuccess(v bool) *DeleteResourceResponseBody {
	s.Success = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type DeleteWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionRequest) SetId(v string) *DeleteWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DeleteWorkflowDefinitionRequest) SetProjectId(v string) *DeleteWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type DeleteWorkflowDefinitionResponseBody struct {
	// example:
	//
	// B17730C0-D959-548A-AE23-E754177CXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponseBody) SetRequestId(v string) *DeleteWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponseBody) SetSuccess(v bool) *DeleteWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *DeleteWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetStatusCode(v int32) *DeleteWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetBody(v *DeleteWorkflowDefinitionResponseBody) *DeleteWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type ExecDeploymentStageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecDeploymentStageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageRequest) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageRequest) SetCode(v string) *ExecDeploymentStageRequest {
	s.Code = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetId(v string) *ExecDeploymentStageRequest {
	s.Id = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetProjectId(v string) *ExecDeploymentStageRequest {
	s.ProjectId = &v
	return s
}

type ExecDeploymentStageResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecDeploymentStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponseBody) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponseBody) SetRequestId(v string) *ExecDeploymentStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecDeploymentStageResponseBody) SetSuccess(v bool) *ExecDeploymentStageResponseBody {
	s.Success = &v
	return s
}

type ExecDeploymentStageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecDeploymentStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDeploymentStageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponse) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponse) SetHeaders(v map[string]*string) *ExecDeploymentStageResponse {
	s.Headers = v
	return s
}

func (s *ExecDeploymentStageResponse) SetStatusCode(v int32) *ExecDeploymentStageResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecDeploymentStageResponse) SetBody(v *ExecDeploymentStageResponseBody) *ExecDeploymentStageResponse {
	s.Body = v
	return s
}

type GetDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// true
	WithDetails *bool `json:"WithDetails,omitempty" xml:"WithDetails,omitempty"`
}

func (s GetDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobRequest) SetDIJobId(v string) *GetDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobRequest) SetWithDetails(v bool) *GetDIJobRequest {
	s.WithDetails = &v
	return s
}

type GetDIJobResponseBody struct {
	PagingInfo *GetDIJobResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBody) SetPagingInfo(v *GetDIJobResponseBodyPagingInfo) *GetDIJobResponseBody {
	s.PagingInfo = v
	return s
}

func (s *GetDIJobResponseBody) SetRequestId(v string) *GetDIJobResponseBody {
	s.RequestId = &v
	return s
}

type GetDIJobResponseBodyPagingInfo struct {
	// example:
	//
	// 32601
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// description
	Description                   *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDataSourceSettings []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// example:
	//
	// imp_ods_dms_det_dealer_info_df
	JobName     *string                                    `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettings *GetDIJobResponseBodyPagingInfoJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 98330
	ProjectId                *int64                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceSettings         *GetDIJobResponseBodyPagingInfoResourceSettings           `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	SourceDataSourceSettings []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// example:
	//
	// Mysql
	SourceDataSourceType *string                                              `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	TableMappings        []*GetDIJobResponseBodyPagingInfoTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules  []*GetDIJobResponseBodyPagingInfoTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfo) SetDIJobId(v string) *GetDIJobResponseBodyPagingInfo {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDescription(v string) *GetDIJobResponseBodyPagingInfo {
	s.Description = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobName(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobSettings(v *GetDIJobResponseBodyPagingInfoJobSettings) *GetDIJobResponseBodyPagingInfo {
	s.JobSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetMigrationType(v string) *GetDIJobResponseBodyPagingInfo {
	s.MigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetProjectId(v int64) *GetDIJobResponseBodyPagingInfo {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.ResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTableMappings(v []*GetDIJobResponseBodyPagingInfoTableMappings) *GetDIJobResponseBodyPagingInfo {
	s.TableMappings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTransformationRules) *GetDIJobResponseBodyPagingInfo {
	s.TransformationRules = v
	return s
}

type GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings struct {
	// example:
	//
	// dw_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                            `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetChannelSettings(v string) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetColumnDataTypeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetCycleScheduleSettings(v *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetDdlHandlingSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetRuntimeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.RuntimeSettings = v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// Full
	CycleMigrationType *string `json:"CycleMigrationType,omitempty" xml:"CycleMigrationType,omitempty"`
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// CreateTable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetAction(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetName(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetValue(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettings struct {
	OfflineResourceSettings  *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetOfflineResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetRealtimeResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetScheduleResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_7708_1667792816832
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1718359176885
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettings struct {
	// example:
	//
	// dw_mysql
	DataSourceName       *string                                                                     `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceProperties *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceProperties(v *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties struct {
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappings struct {
	SourceObjectSelectionRules []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetSourceObjectSelectionRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.TransformationRules = v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetAction(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpression(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules struct {
	// example:
	//
	// AddColumn
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleExpression(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleTargetType = &v
	return s
}

type GetDIJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobResponse) SetHeaders(v map[string]*string) *GetDIJobResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobResponse) SetStatusCode(v int32) *GetDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobResponse) SetBody(v *GetDIJobResponseBody) *GetDIJobResponse {
	s.Body = v
	return s
}

type GetDIJobLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// 10
	FailoverId *int64 `json:"FailoverId,omitempty" xml:"FailoverId,omitempty"`
	// example:
	//
	// 6153616438
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDIJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobLogRequest) SetDIJobId(v int64) *GetDIJobLogRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobLogRequest) SetFailoverId(v int64) *GetDIJobLogRequest {
	s.FailoverId = &v
	return s
}

func (s *GetDIJobLogRequest) SetInstanceId(v int64) *GetDIJobLogRequest {
	s.InstanceId = &v
	return s
}

type GetDIJobLogResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// >>>>>>>> stdout:n++++++++++++++++++executing sql: create database if not exists jindo_test location \\"oss://pangbei-hdfs/tmp/hive\\" n++n
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponseBody) SetLog(v string) *GetDIJobLogResponseBody {
	s.Log = &v
	return s
}

func (s *GetDIJobLogResponseBody) SetRequestId(v string) *GetDIJobLogResponseBody {
	s.RequestId = &v
	return s
}

type GetDIJobLogResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponse) SetHeaders(v map[string]*string) *GetDIJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobLogResponse) SetStatusCode(v int32) *GetDIJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobLogResponse) SetBody(v *GetDIJobLogResponseBody) *GetDIJobLogResponse {
	s.Body = v
	return s
}

type GetDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceRequest) SetId(v int64) *GetDataSourceRequest {
	s.Id = &v
	return s
}

type GetDataSourceResponseBody struct {
	DataSource *GetDataSourceResponseBodyDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetDataSource(v *GetDataSourceResponseBodyDataSource) *GetDataSourceResponseBody {
	s.DataSource = v
	return s
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type GetDataSourceResponseBodyDataSource struct {
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// 1698286929333
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 16738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1698286929333
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1107550004253538
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 52660
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1107550004253538:cn-beijing:holo:hgprecn-cn-x0r3oun4k001:testdb
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionProperties(v interface{}) *GetDataSourceResponseBodyDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionPropertiesMode(v string) *GetDataSourceResponseBodyDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateUser(v string) *GetDataSourceResponseBodyDataSource {
	s.CreateUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetDescription(v string) *GetDataSourceResponseBodyDataSource {
	s.Description = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetId(v int64) *GetDataSourceResponseBodyDataSource {
	s.Id = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.ModifyTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyUser(v string) *GetDataSourceResponseBodyDataSource {
	s.ModifyUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetName(v string) *GetDataSourceResponseBodyDataSource {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetProjectId(v int64) *GetDataSourceResponseBodyDataSource {
	s.ProjectId = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetQualifiedName(v string) *GetDataSourceResponseBodyDataSource {
	s.QualifiedName = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetType(v string) *GetDataSourceResponseBodyDataSource {
	s.Type = &v
	return s
}

type GetDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponse) SetHeaders(v map[string]*string) *GetDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceResponse) SetStatusCode(v int32) *GetDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

type GetDeploymentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentRequest) SetId(v string) *GetDeploymentRequest {
	s.Id = &v
	return s
}

func (s *GetDeploymentRequest) SetProjectId(v string) *GetDeploymentRequest {
	s.ProjectId = &v
	return s
}

type GetDeploymentResponseBody struct {
	Pipeline *GetDeploymentResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// example:
	//
	// 08468352-032C-5262-AEDC-68C9FA05XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBody) SetPipeline(v *GetDeploymentResponseBodyPipeline) *GetDeploymentResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetDeploymentResponseBody) SetRequestId(v string) *GetDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type GetDeploymentResponseBodyPipeline struct {
	// 发布包创建时间戳
	//
	// example:
	//
	// 1724984066000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 创建人
	//
	// example:
	//
	// 137946317766XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 发布流程Id
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1724984066000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 56160
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 步骤详情
	Stages []*GetDeploymentResponseBodyPipelineStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// 发布流程状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeploymentResponseBodyPipeline) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipeline) SetCreateTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetCreator(v string) *GetDeploymentResponseBodyPipeline {
	s.Creator = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetId(v string) *GetDeploymentResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetMessage(v string) *GetDeploymentResponseBodyPipeline {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetModifyTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.ModifyTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetProjectId(v string) *GetDeploymentResponseBodyPipeline {
	s.ProjectId = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStages(v []*GetDeploymentResponseBodyPipelineStages) *GetDeploymentResponseBodyPipeline {
	s.Stages = v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStatus(v string) *GetDeploymentResponseBodyPipeline {
	s.Status = &v
	return s
}

type GetDeploymentResponseBodyPipelineStages struct {
	// 阶段代号
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 阶段描述
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail      map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// 阶段信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 阶段名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 阶段状态
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 步骤
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// 阶段类型
	//
	// example:
	//
	// BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDeploymentResponseBodyPipelineStages) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipelineStages) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipelineStages) SetCode(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Code = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDescription(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Description = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDetail(v map[string]interface{}) *GetDeploymentResponseBodyPipelineStages {
	s.Detail = v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetMessage(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetName(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Name = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStatus(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Status = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStep(v int32) *GetDeploymentResponseBodyPipelineStages {
	s.Step = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetType(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Type = &v
	return s
}

type GetDeploymentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponse) SetHeaders(v map[string]*string) *GetDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentResponse) SetStatusCode(v int32) *GetDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentResponse) SetBody(v *GetDeploymentResponseBody) *GetDeploymentResponse {
	s.Body = v
	return s
}

type GetFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetId(v string) *GetFunctionRequest {
	s.Id = &v
	return s
}

func (s *GetFunctionRequest) SetProjectId(v string) *GetFunctionRequest {
	s.ProjectId = &v
	return s
}

type GetFunctionResponseBody struct {
	Function *GetFunctionResponseBodyFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// example:
	//
	// 6CF95929-6D12-5A88-8CC3-4B2F4C2EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) SetFunction(v *GetFunctionResponseBodyFunction) *GetFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetFunctionResponseBody) SetRequestId(v string) *GetFunctionResponseBody {
	s.RequestId = &v
	return s
}

type GetFunctionResponseBodyFunction struct {
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1724506661000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetFunctionResponseBodyFunction) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBodyFunction) SetCreateTime(v int64) *GetFunctionResponseBodyFunction {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetId(v string) *GetFunctionResponseBodyFunction {
	s.Id = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetModifyTime(v int64) *GetFunctionResponseBodyFunction {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetName(v string) *GetFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetOwner(v string) *GetFunctionResponseBodyFunction {
	s.Owner = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetProjectId(v string) *GetFunctionResponseBodyFunction {
	s.ProjectId = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetSpec(v string) *GetFunctionResponseBodyFunction {
	s.Spec = &v
	return s
}

type GetFunctionResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) SetId(v string) *GetNodeRequest {
	s.Id = &v
	return s
}

func (s *GetNodeRequest) SetProjectId(v string) *GetNodeRequest {
	s.ProjectId = &v
	return s
}

type GetNodeResponseBody struct {
	Node *GetNodeResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeResponseBodyNode struct {
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1700539206000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 196596664824XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetNodeResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v int64) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetId(v string) *GetNodeResponseBodyNode {
	s.Id = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifyTime(v int64) *GetNodeResponseBodyNode {
	s.ModifyTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetOwner(v string) *GetNodeResponseBodyNode {
	s.Owner = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetProjectId(v string) *GetNodeResponseBodyNode {
	s.ProjectId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSpec(v string) *GetNodeResponseBodyNode {
	s.Spec = &v
	return s
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetId(v int64) *GetProjectRequest {
	s.Id = &v
	return s
}

type GetProjectResponseBody struct {
	Project *GetProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponseBodyProject struct {
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string                                            `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*GetProjectResponseBodyProjectAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	Description           *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 28477242
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 207947397706614299
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceGroupId(v string) *GetProjectResponseBodyProject {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceTags(v []*GetProjectResponseBodyProjectAliyunResourceTags) *GetProjectResponseBodyProject {
	s.AliyunResourceTags = v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevEnvironmentEnabled(v bool) *GetProjectResponseBodyProject {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevRoleDisabled(v bool) *GetProjectResponseBodyProject {
	s.DevRoleDisabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDisplayName(v string) *GetProjectResponseBodyProject {
	s.DisplayName = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetId(v int64) *GetProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetName(v string) *GetProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetOwner(v string) *GetProjectResponseBodyProject {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetPaiTaskEnabled(v bool) *GetProjectResponseBodyProject {
	s.PaiTaskEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetStatus(v string) *GetProjectResponseBodyProject {
	s.Status = &v
	return s
}

type GetProjectResponseBodyProjectAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProjectResponseBodyProjectAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyProjectAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetKey(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetValue(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Value = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetId(v string) *GetResourceRequest {
	s.Id = &v
	return s
}

func (s *GetResourceRequest) SetProjectId(v string) *GetResourceRequest {
	s.ProjectId = &v
	return s
}

type GetResourceResponseBody struct {
	// example:
	//
	// E871F6C0-2EFF-5790-A00D-C57543EEXXXX
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *GetResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody {
	s.Resource = v
	return s
}

type GetResourceResponseBodyResource struct {
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1700539206000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetResourceResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResource) SetCreateTime(v int64) *GetResourceResponseBodyResource {
	s.CreateTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetId(v string) *GetResourceResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetModifyTime(v int64) *GetResourceResponseBodyResource {
	s.ModifyTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetName(v string) *GetResourceResponseBodyResource {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetOwner(v string) *GetResourceResponseBodyResource {
	s.Owner = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetProjectId(v string) *GetResourceResponseBodyResource {
	s.ProjectId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetSpec(v string) *GetResourceResponseBodyResource {
	s.Spec = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionRequest) SetId(v string) *GetWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) SetProjectId(v string) *GetWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type GetWorkflowDefinitionResponseBody struct {
	// example:
	//
	// F2BDD628-8A21-5BD1-B930-1A2D5989XXXX
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowDefinition *GetWorkflowDefinitionResponseBodyWorkflowDefinition `json:"WorkflowDefinition,omitempty" xml:"WorkflowDefinition,omitempty" type:"Struct"`
}

func (s GetWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBody) SetRequestId(v string) *GetWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBody) SetWorkflowDefinition(v *GetWorkflowDefinitionResponseBodyWorkflowDefinition) *GetWorkflowDefinitionResponseBody {
	s.WorkflowDefinition = v
	return s
}

type GetWorkflowDefinitionResponseBodyWorkflowDefinition struct {
	// example:
	//
	// 1708481905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1708481905000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetCreateTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetModifyTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetName(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Name = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetOwner(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Owner = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetProjectId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetSpec(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Spec = &v
	return s
}

type GetWorkflowDefinitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *GetWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetStatusCode(v int32) *GetWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetBody(v *GetWorkflowDefinitionResponseBody) *GetWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type ListDIJobEventsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1717971005
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Alarm
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1716971005
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsRequest) SetDIJobId(v string) *ListDIJobEventsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEndTime(v int64) *ListDIJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEventType(v string) *ListDIJobEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageNumber(v int64) *ListDIJobEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageSize(v int64) *ListDIJobEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsRequest) SetStartTime(v int64) *ListDIJobEventsRequest {
	s.StartTime = &v
	return s
}

type ListDIJobEventsResponseBody struct {
	PagingInfo *ListDIJobEventsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 645F6D68-9C29-5961-80B1-BDD4B794C22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBody) SetPagingInfo(v *ListDIJobEventsResponseBodyPagingInfo) *ListDIJobEventsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobEventsResponseBody) SetRequestId(v string) *ListDIJobEventsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobEventsResponseBodyPagingInfo struct {
	DIJobEvent []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent `json:"DIJobEvent,omitempty" xml:"DIJobEvent,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2524
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetDIJobEvent(v []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent) *ListDIJobEventsResponseBodyPagingInfo {
	s.DIJobEvent = v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIJobEventsResponseBodyPagingInfoDIJobEvent struct {
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// Phone
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// 1663573162
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Detail     *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// alter table table2 ***
	DstSql *string `json:"DstSql,omitempty" xml:"DstSql,omitempty"`
	// example:
	//
	// table2
	DstTable *string `json:"DstTable,omitempty" xml:"DstTable,omitempty"`
	// example:
	//
	// 2024-05-29 15:11:31,377 [main] INFO com.*.**.di.core.metrics.:21 []  {****}
	//
	// 2024-05-29 15:11:31,384 [main] INFO *.aliyun.*.di.*.*.metrics.*:27 [] - Open MarioDiReporter
	//
	// 2024-05-29 15:11:33,248 [flink-akka.*.*-dispatcher-17] INFO
	FailoverMessage *string `json:"FailoverMessage,omitempty" xml:"FailoverMessage,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// alter table table1 ***
	SrcSql *string `json:"SrcSql,omitempty" xml:"SrcSql,omitempty"`
	// example:
	//
	// table1
	SrcTable *string `json:"SrcTable,omitempty" xml:"SrcTable,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Delay
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetAction(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Action = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetChannels(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Channels = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetCreateTime(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.CreateTime = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDetail(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Detail = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetFailoverMessage(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.FailoverMessage = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetId(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Id = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSeverity(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Severity = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetStatus(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Status = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetType(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Type = &v
	return s
}

type ListDIJobEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponse) SetHeaders(v map[string]*string) *ListDIJobEventsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobEventsResponse) SetStatusCode(v int32) *ListDIJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobEventsResponse) SetBody(v *ListDIJobEventsResponseBody) *ListDIJobEventsResponse {
	s.Body = v
	return s
}

type ListDIJobMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11265
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	MetricName []*string `json:"MetricName,omitempty" xml:"MetricName,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsRequest) SetDIJobId(v string) *ListDIJobMetricsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetEndTime(v int64) *ListDIJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetMetricName(v []*string) *ListDIJobMetricsRequest {
	s.MetricName = v
	return s
}

func (s *ListDIJobMetricsRequest) SetStartTime(v int64) *ListDIJobMetricsRequest {
	s.StartTime = &v
	return s
}

type ListDIJobMetricsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11265
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	MetricNameShrink *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsShrinkRequest) SetDIJobId(v string) *ListDIJobMetricsShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetEndTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetMetricNameShrink(v string) *ListDIJobMetricsShrinkRequest {
	s.MetricNameShrink = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetStartTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

type ListDIJobMetricsResponseBody struct {
	PagingInfo *ListDIJobMetricsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBody) SetPagingInfo(v *ListDIJobMetricsResponseBodyPagingInfo) *ListDIJobMetricsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobMetricsResponseBody) SetRequestId(v string) *ListDIJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfo struct {
	JobMetrics []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics `json:"JobMetrics,omitempty" xml:"JobMetrics,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfo) SetJobMetrics(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics) *ListDIJobMetricsResponseBodyPagingInfo {
	s.JobMetrics = v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetrics struct {
	// example:
	//
	// JobDelay
	Name       *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	SeriesList []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList `json:"SeriesList,omitempty" xml:"SeriesList,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetName(v string) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.Name = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetSeriesList(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.SeriesList = v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList struct {
	// example:
	//
	// 1716881141
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 10
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetTime(v int64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Time = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetValue(v float64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Value = &v
	return s
}

type ListDIJobMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponse) SetHeaders(v map[string]*string) *ListDIJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobMetricsResponse) SetStatusCode(v int32) *ListDIJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobMetricsResponse) SetBody(v *ListDIJobMetricsResponseBody) *ListDIJobMetricsResponse {
	s.Body = v
	return s
}

type ListDIJobsRequest struct {
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobsRequest) SetDestinationDataSourceType(v string) *ListDIJobsRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) SetMigrationType(v string) *ListDIJobsRequest {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsRequest) SetName(v string) *ListDIJobsRequest {
	s.Name = &v
	return s
}

func (s *ListDIJobsRequest) SetPageNumber(v int64) *ListDIJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsRequest) SetPageSize(v int64) *ListDIJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsRequest) SetProjectId(v int64) *ListDIJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsRequest) SetSourceDataSourceType(v string) *ListDIJobsRequest {
	s.SourceDataSourceType = &v
	return s
}

type ListDIJobsResponseBody struct {
	PagingInfo *ListDIJobsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 7263E4AC-9D2E-5B29-B8AF-7C5012E92A41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBody) SetPagingInfo(v *ListDIJobsResponseBodyPagingInfo) *ListDIJobsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobsResponseBody) SetRequestId(v string) *ListDIJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobsResponseBodyPagingInfo struct {
	DIJobs []*ListDIJobsResponseBodyPagingInfoDIJobs `json:"DIJobs,omitempty" xml:"DIJobs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfo) SetDIJobs(v []*ListDIJobsResponseBodyPagingInfoDIJobs) *ListDIJobsResponseBodyPagingInfo {
	s.DIJobs = v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIJobsResponseBodyPagingInfoDIJobs struct {
	// example:
	//
	// 32599
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// example:
	//
	// mysql_to_holo_sync_35197
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 26442
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDIJobId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDestinationDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobName(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobName = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobStatus(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetMigrationType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetProjectId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetSourceDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.SourceDataSourceType = &v
	return s
}

type ListDIJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponse) SetHeaders(v map[string]*string) *ListDIJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobsResponse) SetStatusCode(v int32) *ListDIJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobsResponse) SetBody(v *ListDIJobsResponseBody) *ListDIJobsResponse {
	s.Body = v
	return s
}

type ListDataSourceSharedRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesRequest) SetDataSourceId(v int64) *ListDataSourceSharedRulesRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesRequest) SetTargetProjectId(v int64) *ListDataSourceSharedRulesRequest {
	s.TargetProjectId = &v
	return s
}

type ListDataSourceSharedRulesResponseBody struct {
	DataSourceSharedRules []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules `json:"DataSourceSharedRules,omitempty" xml:"DataSourceSharedRules,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBody) SetDataSourceSharedRules(v []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) *ListDataSourceSharedRulesResponseBody {
	s.DataSourceSharedRules = v
	return s
}

func (s *ListDataSourceSharedRulesResponseBody) SetRequestId(v string) *ListDataSourceSharedRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourceSharedRulesResponseBodyDataSourceSharedRules struct {
	// example:
	//
	// 1724379762000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// targetProject.datasource
	SharedDataSourceName *string `json:"SharedDataSourceName,omitempty" xml:"SharedDataSourceName,omitempty"`
	// example:
	//
	// 1
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// example:
	//
	// 1
	SourceProjectId *int64 `json:"SourceProjectId,omitempty" xml:"SourceProjectId,omitempty"`
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateTime(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetDataSourceId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetEnvType(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.EnvType = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.Id = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedDataSourceName(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedDataSourceName = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSourceProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SourceProjectId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetTargetProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.TargetProjectId = &v
	return s
}

type ListDataSourceSharedRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceSharedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceSharedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponse) SetHeaders(v map[string]*string) *ListDataSourceSharedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetStatusCode(v int32) *ListDataSourceSharedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetBody(v *ListDataSourceSharedRulesResponseBody) *ListDataSourceSharedRulesResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17820
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ["tag1", "tag2", "tag3"]
	Tags  *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetEnvType(v string) *ListDataSourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesRequest) SetName(v string) *ListDataSourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesRequest) SetOrder(v string) *ListDataSourcesRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageNumber(v int32) *ListDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v int32) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetProjectId(v int64) *ListDataSourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesRequest) SetSortBy(v string) *ListDataSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesRequest) SetTags(v string) *ListDataSourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListDataSourcesRequest) SetTypes(v []*string) *ListDataSourcesRequest {
	s.Types = v
	return s
}

type ListDataSourcesShrinkRequest struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17820
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ["tag1", "tag2", "tag3"]
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) SetEnvType(v string) *ListDataSourcesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetName(v string) *ListDataSourcesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetOrder(v string) *ListDataSourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageNumber(v int32) *ListDataSourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageSize(v int32) *ListDataSourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetProjectId(v int64) *ListDataSourcesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetSortBy(v string) *ListDataSourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTags(v string) *ListDataSourcesShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTypesShrink(v string) *ListDataSourcesShrinkRequest {
	s.TypesShrink = &v
	return s
}

type ListDataSourcesResponseBody struct {
	PagingInfo *ListDataSourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 7BE1433F-6D55-5D86-9344-CA6F7DD19B13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetPagingInfo(v *ListDataSourcesResponseBodyPagingInfo) *ListDataSourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfo struct {
	DataSources []*ListDataSourcesResponseBodyPagingInfoDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 131
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetDataSources(v []*ListDataSourcesResponseBodyPagingInfoDataSources) *ListDataSourcesResponseBodyPagingInfo {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageNumber(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageSize(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfoDataSources struct {
	DataSource []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// mysql
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetDataSource(v []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.DataSource = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetName(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetType(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Type = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource struct {
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// 1648711113000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1624387842781448
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1648711113000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1624387842781448
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// 1648711121000:cn-beijing:odps:yongxunQA_beijing_standard
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionProperties(v interface{}) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionPropertiesMode(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetDescription(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Description = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetId(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetQualifiedName(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.QualifiedName = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListDeploymentsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageNumber(v int32) *ListDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetProjectId(v string) *ListDeploymentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v string) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBody struct {
	PagingInfo *ListDeploymentsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) SetPagingInfo(v *ListDeploymentsResponseBodyPagingInfo) *ListDeploymentsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfo struct {
	Deployments []*ListDeploymentsResponseBodyPagingInfoDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2524
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetDeployments(v []*ListDeploymentsResponseBodyPagingInfoDeployments) *ListDeploymentsResponseBodyPagingInfo {
	s.Deployments = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageNumber(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageSize(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetTotalCount(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeployments struct {
	// 发布包创建时间戳
	//
	// example:
	//
	// 1702736654000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 创建人
	//
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 发布流程Id
	//
	// example:
	//
	// ddf354a5-03df-48fc-94c1-cc973f79XXXX
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1702736654000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 项目Id
	//
	// example:
	//
	// 44683
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 步骤详情
	Stages []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// 发布流程状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreateTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.CreateTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreator(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Id = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetModifyTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ModifyTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetProjectId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStages(v []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Stages = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeploymentsStages struct {
	// 阶段代号
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 阶段描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 阶段详细信息
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// 阶段信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 阶段名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 阶段状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 步骤
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// 阶段类型
	//
	// example:
	//
	// CHECK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetCode(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Code = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDescription(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Description = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDetail(v map[string]interface{}) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Detail = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetName(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Name = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Status = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStep(v int32) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Step = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetType(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Type = &v
	return s
}

type ListDeploymentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponse) SetHeaders(v map[string]*string) *ListDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentsResponse) SetStatusCode(v int32) *ListDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentsResponse) SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetOwner(v string) *ListFunctionsRequest {
	s.Owner = &v
	return s
}

func (s *ListFunctionsRequest) SetPageNumber(v int32) *ListFunctionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsRequest) SetPageSize(v int32) *ListFunctionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsRequest) SetProjectId(v string) *ListFunctionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsRequest) SetType(v string) *ListFunctionsRequest {
	s.Type = &v
	return s
}

type ListFunctionsResponseBody struct {
	PagingInfo *ListFunctionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 89FB2BF0-EB00-5D03-9C34-05931001XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetPagingInfo(v *ListFunctionsResponseBodyPagingInfo) *ListFunctionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfo struct {
	Functions []*ListFunctionsResponseBodyPagingInfoFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfo) SetFunctions(v []*ListFunctionsResponseBodyPagingInfoFunctions) *ListFunctionsResponseBodyPagingInfo {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageSize(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctions struct {
	// ARM集群资源文件列表
	//
	// example:
	//
	// xxx.jar,yyy.jar
	ArmResource *string `json:"ArmResource,omitempty" xml:"ArmResource,omitempty"`
	// 函数实现类名
	//
	// example:
	//
	// com.demo.Main
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// 命令描述
	//
	// example:
	//
	// testUdf(xx,yy)
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 1655953028000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 函数注册到的数据源信息
	DataSource *ListFunctionsResponseBodyPagingInfoFunctionsDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 数据库名，可选
	//
	// example:
	//
	// odps_first
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 对funciotn的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 嵌套函数代码内容
	//
	// example:
	//
	// print(\\"hello,world!\\")
	EmbeddedCode *string `json:"EmbeddedCode,omitempty" xml:"EmbeddedCode,omitempty"`
	// 嵌套代码类型
	//
	// example:
	//
	// Python2
	EmbeddedCodeType *string `json:"EmbeddedCodeType,omitempty" xml:"EmbeddedCodeType,omitempty"`
	// 嵌套资源类型
	//
	// example:
	//
	// File
	EmbeddedResourceType *string `json:"EmbeddedResourceType,omitempty" xml:"EmbeddedResourceType,omitempty"`
	// 示例说明
	ExampleDescription *string `json:"ExampleDescription,omitempty" xml:"ExampleDescription,omitempty"`
	// 函数的实现代码
	//
	// example:
	//
	// xxx.jar,yyy.jar
	FileResource *string `json:"FileResource,omitempty" xml:"FileResource,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1655953028000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 代表资源名称的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 函数责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 命令描述
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// 项目Id
	//
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 返回值说明
	ReturnValueDescription *string `json:"ReturnValueDescription,omitempty" xml:"ReturnValueDescription,omitempty"`
	// 运行时资源组信息
	RuntimeResource *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListFunctionsResponseBodyPagingInfoFunctionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 函数类型
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetArmResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ArmResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetClassName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ClassName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCommandDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CommandDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCreateTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDataSource(v *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DataSource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDatabaseName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Description = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCode(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCode = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCodeType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCodeType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedResourceType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedResourceType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetExampleDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ExampleDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetFileResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.FileResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetModifyTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetOwner(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetParameterDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ParameterDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetProjectId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetReturnValueDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ReturnValueDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetRuntimeResource(v *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.RuntimeResource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetScript(v *ListFunctionsResponseBodyPagingInfoFunctionsScript) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Script = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource struct {
	// 运行时资源组Id
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) SetResourceGroupId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScript struct {
	// 脚本的id
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetPath(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Path = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetRuntime(v *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Runtime = v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_FUNCTION
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) SetCommand(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	s.Command = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListNodeDependenciesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDependenciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesRequest) SetId(v string) *ListNodeDependenciesRequest {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageNumber(v int32) *ListNodeDependenciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageSize(v int32) *ListNodeDependenciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetProjectId(v string) *ListNodeDependenciesRequest {
	s.ProjectId = &v
	return s
}

type ListNodeDependenciesResponseBody struct {
	PagingInfo *ListNodeDependenciesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeDependenciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBody) SetPagingInfo(v *ListNodeDependenciesResponseBodyPagingInfo) *ListNodeDependenciesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodeDependenciesResponseBody) SetRequestId(v string) *ListNodeDependenciesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfo struct {
	Nodes []*ListNodeDependenciesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 90
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetNodes(v []*ListNodeDependenciesResponseBodyPagingInfoNodes) *ListNodeDependenciesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageSize(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodes struct {
	// 节点的创建时间
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 数据源信息
	DataSource *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 723932906364267XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 节点输入
	Inputs *ListNodeDependenciesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// 属性修改时间
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 节点名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 节点输出
	Outputs *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// 节点的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 65133
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// 资源组信息
	RuntimeResource *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListNodeDependenciesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 调度策略
	Strategy *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// 标签信息
	Tags []*ListNodeDependenciesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 调度任务Id
	//
	// example:
	//
	// 580667964888595XXXX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 触发器信息
	Trigger *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetInputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetScript(v *ListNodeDependenciesResponseBodyPagingInfoNodesScript) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTags(v []*ListNodeDependenciesResponseBodyPagingInfoNodesTags) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543218872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 463497880880954XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 463497880880954XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource struct {
	// 资源组id
	//
	// example:
	//
	// S_res_group_XXXX_XXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScript struct {
	// 脚本的id
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesStrategy struct {
	// 生成实例的模式
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// 重试时间间隔
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// 允许重跑的模式
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// 重试次数
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// 超时时间
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTags struct {
	// 标签键
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTrigger struct {
	// 触发器的cron表达式
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// 结束时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 触发器id
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 开始时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// 触发器类型
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDependenciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponse) SetHeaders(v map[string]*string) *ListNodeDependenciesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDependenciesResponse) SetStatusCode(v int32) *ListNodeDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDependenciesResponse) SetBody(v *ListNodeDependenciesResponseBody) *ListNodeDependenciesResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// example:
	//
	// Normal
	Rerurrence *string `json:"Rerurrence,omitempty" xml:"Rerurrence,omitempty"`
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetContainerId(v string) *ListNodesRequest {
	s.ContainerId = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProjectId(v string) *ListNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequest) SetRerunMode(v string) *ListNodesRequest {
	s.RerunMode = &v
	return s
}

func (s *ListNodesRequest) SetRerurrence(v string) *ListNodesRequest {
	s.Rerurrence = &v
	return s
}

func (s *ListNodesRequest) SetScene(v string) *ListNodesRequest {
	s.Scene = &v
	return s
}

type ListNodesResponseBody struct {
	PagingInfo *ListNodesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2197B9C4-39CE-55EA-8EEA-FDBAE52DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetPagingInfo(v *ListNodesResponseBodyPagingInfo) *ListNodesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodesResponseBodyPagingInfo struct {
	Nodes []*ListNodesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfo) SetNodes(v []*ListNodesResponseBodyPagingInfoNodes) *ListNodesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageSize(v string) *ListNodesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodes struct {
	// 节点的创建时间
	//
	// example:
	//
	// 1722910655000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 数据源信息
	DataSource *ListNodesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 节点输入
	Inputs *ListNodesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// 属性修改时间
	//
	// example:
	//
	// 1722910655000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 节点名
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 节点输出
	Outputs *ListNodesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// 节点的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 33233
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// 资源组信息
	RuntimeResource *ListNodesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListNodesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 调度策略
	Strategy *ListNodesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// 标签信息
	Tags []*ListNodesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 调度任务Id
	//
	// example:
	//
	// 88888888888
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 触发器信息
	Trigger *ListNodesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodesResponseBodyPagingInfoNodesDataSource) *ListNodesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetInputs(v *ListNodesResponseBodyPagingInfoNodesInputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetName(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodesResponseBodyPagingInfoNodesOutputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodesResponseBodyPagingInfoNodesRuntimeResource) *ListNodesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetScript(v *ListNodesResponseBodyPagingInfoNodesScript) *ListNodesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodesResponseBodyPagingInfoNodesStrategy) *ListNodesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTags(v []*ListNodesResponseBodyPagingInfoNodesTags) *ListNodesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodesResponseBodyPagingInfoNodesTrigger) *ListNodesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesInputsTables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesInputsVariables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 623731286945488XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543211286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 222
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 623731286945488XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesOutputsTables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesOutputsVariables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 623731286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesRuntimeResource struct {
	// 资源组id
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScript struct {
	// 脚本的id
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListNodesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodesResponseBodyPagingInfoNodesScriptRuntime) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesStrategy struct {
	// 生成实例的模式
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// 重试时间间隔
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// 允许重跑的模式
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// 重试次数
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// 超时时间
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTags struct {
	// 标签键
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTrigger struct {
	// 触发器的cron表达式
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// 结束时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 触发器id
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 开始时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// 触发器类型
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string                                  `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*ListProjectsRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool     `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	Ids             []*int64  `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	Names           []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetAliyunResourceGroupId(v string) *ListProjectsRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsRequest) SetAliyunResourceTags(v []*ListProjectsRequestAliyunResourceTags) *ListProjectsRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetDevRoleDisabled(v bool) *ListProjectsRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsRequest) SetIds(v []*int64) *ListProjectsRequest {
	s.Ids = v
	return s
}

func (s *ListProjectsRequest) SetNames(v []*string) *ListProjectsRequest {
	s.Names = v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetPaiTaskEnabled(v bool) *ListProjectsRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetStatus(v string) *ListProjectsRequest {
	s.Status = &v
	return s
}

type ListProjectsRequestAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsRequestAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestAliyunResourceTags) SetKey(v string) *ListProjectsRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsRequestAliyunResourceTags) SetValue(v string) *ListProjectsRequestAliyunResourceTags {
	s.Value = &v
	return s
}

type ListProjectsShrinkRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId    *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	IdsShrink       *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	NamesShrink     *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceGroupId(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceTagsShrink(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevRoleDisabled(v bool) *ListProjectsShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetIdsShrink(v string) *ListProjectsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetNamesShrink(v string) *ListProjectsShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int32) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int32) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPaiTaskEnabled(v bool) *ListProjectsShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetStatus(v string) *ListProjectsShrinkRequest {
	s.Status = &v
	return s
}

type ListProjectsResponseBody struct {
	PagingInfo *ListProjectsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 6D24AD9A-652F-59E2-AC1F-05029300F8A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetPagingInfo(v *ListProjectsResponseBodyPagingInfo) *ListProjectsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyPagingInfo struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Projects []*ListProjectsResponseBodyPagingInfoProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageNumber(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageSize(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetProjects(v []*ListProjectsResponseBodyPagingInfoProjects) *ListProjectsResponseBodyPagingInfo {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetTotalCount(v int32) *ListProjectsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyPagingInfoProjects struct {
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string                                                         `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	Description           *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123532153125
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfoProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceGroupId(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceTags(v []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDescription(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevEnvironmentEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevRoleDisabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDisplayName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.DisplayName = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetId(v int64) *ListProjectsResponseBodyPagingInfoProjects {
	s.Id = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetOwner(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetPaiTaskEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetStatus(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Status = &v
	return s
}

type ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetKey(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetValue(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Value = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// python
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetOwner(v string) *ListResourcesRequest {
	s.Owner = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProjectId(v string) *ListResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesRequest) SetType(v string) *ListResourcesRequest {
	s.Type = &v
	return s
}

type ListResourcesResponseBody struct {
	PagingInfo *ListResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPagingInfo(v *ListResourcesResponseBodyPagingInfo) *ListResourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resources []*ListResourcesResponseBodyPagingInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 131
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageNumber(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageSize(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetResources(v []*ListResourcesResponseBodyPagingInfoResources) *ListResourcesResponseBodyPagingInfo {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetTotalCount(v int32) *ListResourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResources struct {
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 函数注册到的数据源信息
	DataSource *ListResourcesResponseBodyPagingInfoResourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 资源文件的最近修改时间
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// math.py
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 资源文件的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 资源文件的项目id
	//
	// example:
	//
	// 344247
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 工作流的脚本信息
	Script *ListResourcesResponseBodyPagingInfoResourcesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 文件目标存储路径
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// 文件资源来源存储类型
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// 文件来源路径
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 文件目标存储类型
	//
	// example:
	//
	// oss
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 资源类型
	//
	// example:
	//
	// jar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetCreateTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetDataSource(v *ListResourcesResponseBodyPagingInfoResourcesDataSource) *ListResourcesResponseBodyPagingInfoResources {
	s.DataSource = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetModifyTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.ModifyTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetName(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetOwner(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetProjectId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetScript(v *ListResourcesResponseBodyPagingInfoResourcesScript) *ListResourcesResponseBodyPagingInfoResources {
	s.Script = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourcePath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourcePath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourceType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourceType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetPath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetPath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetName(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetType(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScript struct {
	// 工作流脚本的id
	//
	// example:
	//
	// 123348864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetId(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetPath(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Path = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetRuntime(v *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Runtime = v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_PYTHON
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) SetCommand(v string) *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime {
	s.Command = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListWorkflowDefinitionsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsRequest) SetOwner(v string) *ListWorkflowDefinitionsRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageNumber(v int32) *ListWorkflowDefinitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageSize(v int32) *ListWorkflowDefinitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetProjectId(v string) *ListWorkflowDefinitionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetType(v string) *ListWorkflowDefinitionsRequest {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBody struct {
	PagingInfo *ListWorkflowDefinitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 8C3ED0C5-ABAB-55E1-854B-DAC02B11XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBody) SetPagingInfo(v *ListWorkflowDefinitionsResponseBodyPagingInfo) *ListWorkflowDefinitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBody) SetRequestId(v string) *ListWorkflowDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 227
	TotalCount          *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkflowDefinitions []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions `json:"WorkflowDefinitions,omitempty" xml:"WorkflowDefinitions,omitempty" type:"Repeated"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetWorkflowDefinitions(v []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.WorkflowDefinitions = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions struct {
	// 工作流的创建时间
	//
	// example:
	//
	// 1698057323000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 工作流的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工作流定义的唯一ID
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的最近修改时间
	//
	// example:
	//
	// 1698057323000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 工作流的名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 工作流的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 工作流定义的所属项目空间
	//
	// This parameter is required.
	//
	// example:
	//
	// 4710
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 工作流的脚本信息
	Script *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 工作流类型，可选值：CycleWorkflow、ManualWorkflow，分别表示周期工作流和手动工作流
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetCreateTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetDescription(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Description = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetModifyTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetName(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Name = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetOwner(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetProjectId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetScript(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Script = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetType(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript struct {
	// 工作流脚本的id
	//
	// example:
	//
	// 698002781368644XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的脚本路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetPath(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Path = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetRuntime(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Runtime = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// WORKFLOW
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) SetCommand(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime {
	s.Command = &v
	return s
}

type ListWorkflowDefinitionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowDefinitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponse) SetHeaders(v map[string]*string) *ListWorkflowDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetStatusCode(v int32) *ListWorkflowDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetBody(v *ListWorkflowDefinitionsResponseBody) *ListWorkflowDefinitionsResponse {
	s.Body = v
	return s
}

type MoveFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionRequest) GoString() string {
	return s.String()
}

func (s *MoveFunctionRequest) SetId(v string) *MoveFunctionRequest {
	s.Id = &v
	return s
}

func (s *MoveFunctionRequest) SetPath(v string) *MoveFunctionRequest {
	s.Path = &v
	return s
}

func (s *MoveFunctionRequest) SetProjectId(v string) *MoveFunctionRequest {
	s.ProjectId = &v
	return s
}

type MoveFunctionResponseBody struct {
	// example:
	//
	// 48C0E2F0-52BA-5888-BDFA-28F1B9AFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponseBody) SetRequestId(v string) *MoveFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveFunctionResponseBody) SetSuccess(v bool) *MoveFunctionResponseBody {
	s.Success = &v
	return s
}

type MoveFunctionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponse) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponse) SetHeaders(v map[string]*string) *MoveFunctionResponse {
	s.Headers = v
	return s
}

func (s *MoveFunctionResponse) SetStatusCode(v int32) *MoveFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFunctionResponse) SetBody(v *MoveFunctionResponseBody) *MoveFunctionResponse {
	s.Body = v
	return s
}

type MoveNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeRequest) GoString() string {
	return s.String()
}

func (s *MoveNodeRequest) SetId(v string) *MoveNodeRequest {
	s.Id = &v
	return s
}

func (s *MoveNodeRequest) SetPath(v string) *MoveNodeRequest {
	s.Path = &v
	return s
}

func (s *MoveNodeRequest) SetProjectId(v string) *MoveNodeRequest {
	s.ProjectId = &v
	return s
}

type MoveNodeResponseBody struct {
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponseBody) GoString() string {
	return s.String()
}

func (s *MoveNodeResponseBody) SetRequestId(v string) *MoveNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveNodeResponseBody) SetSuccess(v bool) *MoveNodeResponseBody {
	s.Success = &v
	return s
}

type MoveNodeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponse) GoString() string {
	return s.String()
}

func (s *MoveNodeResponse) SetHeaders(v map[string]*string) *MoveNodeResponse {
	s.Headers = v
	return s
}

func (s *MoveNodeResponse) SetStatusCode(v int32) *MoveNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveNodeResponse) SetBody(v *MoveNodeResponseBody) *MoveNodeResponse {
	s.Body = v
	return s
}

type MoveResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) SetId(v string) *MoveResourceRequest {
	s.Id = &v
	return s
}

func (s *MoveResourceRequest) SetPath(v string) *MoveResourceRequest {
	s.Path = &v
	return s
}

func (s *MoveResourceRequest) SetProjectId(v string) *MoveResourceRequest {
	s.ProjectId = &v
	return s
}

type MoveResourceResponseBody struct {
	// example:
	//
	// F332BED4-DD73-5972-A9C2-642BA6CFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

type MoveResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceResponse) SetHeaders(v map[string]*string) *MoveResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceResponse) SetStatusCode(v int32) *MoveResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceResponse) SetBody(v *MoveResourceResponseBody) *MoveResourceResponse {
	s.Body = v
	return s
}

type MoveWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionRequest) SetId(v string) *MoveWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetPath(v string) *MoveWorkflowDefinitionRequest {
	s.Path = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetProjectId(v string) *MoveWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type MoveWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 05ADAF4F-7709-5FB1-B606-3513483FXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponseBody) SetRequestId(v string) *MoveWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveWorkflowDefinitionResponseBody) SetSuccess(v bool) *MoveWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type MoveWorkflowDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *MoveWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetStatusCode(v int32) *MoveWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetBody(v *MoveWorkflowDefinitionResponseBody) *MoveWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type RenameFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionRequest) GoString() string {
	return s.String()
}

func (s *RenameFunctionRequest) SetId(v string) *RenameFunctionRequest {
	s.Id = &v
	return s
}

func (s *RenameFunctionRequest) SetName(v string) *RenameFunctionRequest {
	s.Name = &v
	return s
}

func (s *RenameFunctionRequest) SetProjectId(v string) *RenameFunctionRequest {
	s.ProjectId = &v
	return s
}

type RenameFunctionResponseBody struct {
	// example:
	//
	// 1ED4C97F-BA2A-57C5-BA7C-8853627EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponseBody) SetRequestId(v string) *RenameFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFunctionResponseBody) SetSuccess(v string) *RenameFunctionResponseBody {
	s.Success = &v
	return s
}

type RenameFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponse) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponse) SetHeaders(v map[string]*string) *RenameFunctionResponse {
	s.Headers = v
	return s
}

func (s *RenameFunctionResponse) SetStatusCode(v int32) *RenameFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameFunctionResponse) SetBody(v *RenameFunctionResponseBody) *RenameFunctionResponse {
	s.Body = v
	return s
}

type RenameNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeRequest) GoString() string {
	return s.String()
}

func (s *RenameNodeRequest) SetId(v string) *RenameNodeRequest {
	s.Id = &v
	return s
}

func (s *RenameNodeRequest) SetName(v string) *RenameNodeRequest {
	s.Name = &v
	return s
}

func (s *RenameNodeRequest) SetProjectId(v string) *RenameNodeRequest {
	s.ProjectId = &v
	return s
}

type RenameNodeResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenameNodeResponseBody) SetRequestId(v string) *RenameNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameNodeResponseBody) SetSuccess(v bool) *RenameNodeResponseBody {
	s.Success = &v
	return s
}

type RenameNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponse) GoString() string {
	return s.String()
}

func (s *RenameNodeResponse) SetHeaders(v map[string]*string) *RenameNodeResponse {
	s.Headers = v
	return s
}

func (s *RenameNodeResponse) SetStatusCode(v int32) *RenameNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameNodeResponse) SetBody(v *RenameNodeResponseBody) *RenameNodeResponse {
	s.Body = v
	return s
}

type RenameResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceRequest) GoString() string {
	return s.String()
}

func (s *RenameResourceRequest) SetId(v string) *RenameResourceRequest {
	s.Id = &v
	return s
}

func (s *RenameResourceRequest) SetName(v string) *RenameResourceRequest {
	s.Name = &v
	return s
}

func (s *RenameResourceRequest) SetProjectId(v string) *RenameResourceRequest {
	s.ProjectId = &v
	return s
}

type RenameResourceResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA8XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameResourceResponseBody) SetRequestId(v string) *RenameResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameResourceResponseBody) SetSuccess(v bool) *RenameResourceResponseBody {
	s.Success = &v
	return s
}

type RenameResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponse) GoString() string {
	return s.String()
}

func (s *RenameResourceResponse) SetHeaders(v map[string]*string) *RenameResourceResponse {
	s.Headers = v
	return s
}

func (s *RenameResourceResponse) SetStatusCode(v int32) *RenameResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameResourceResponse) SetBody(v *RenameResourceResponseBody) *RenameResourceResponse {
	s.Body = v
	return s
}

type RenameWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionRequest) SetId(v string) *RenameWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetName(v string) *RenameWorkflowDefinitionRequest {
	s.Name = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetProjectId(v string) *RenameWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type RenameWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 975BD43D-C421-595C-A29C-565A8AD5XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponseBody) SetRequestId(v string) *RenameWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameWorkflowDefinitionResponseBody) SetSuccess(v bool) *RenameWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type RenameWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *RenameWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetStatusCode(v int32) *RenameWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetBody(v *RenameWorkflowDefinitionResponseBody) *RenameWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type StartDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// false
	ForceToRerun          *bool                                   `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	RealtimeStartSettings *StartDIJobRequestRealtimeStartSettings `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty" type:"Struct"`
}

func (s StartDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobRequest) SetDIJobId(v string) *StartDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobRequest) SetForceToRerun(v bool) *StartDIJobRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobRequest) SetRealtimeStartSettings(v *StartDIJobRequestRealtimeStartSettings) *StartDIJobRequest {
	s.RealtimeStartSettings = v
	return s
}

type StartDIJobRequestRealtimeStartSettings struct {
	FailoverSettings *StartDIJobRequestRealtimeStartSettingsFailoverSettings `json:"FailoverSettings,omitempty" xml:"FailoverSettings,omitempty" type:"Struct"`
	// example:
	//
	// 1671516776
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettings) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettings) SetFailoverSettings(v *StartDIJobRequestRealtimeStartSettingsFailoverSettings) *StartDIJobRequestRealtimeStartSettings {
	s.FailoverSettings = v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettings) SetStartTime(v int64) *StartDIJobRequestRealtimeStartSettings {
	s.StartTime = &v
	return s
}

type StartDIJobRequestRealtimeStartSettingsFailoverSettings struct {
	// example:
	//
	// 10
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 30
	UpperLimit *int64 `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetInterval(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.Interval = &v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetUpperLimit(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.UpperLimit = &v
	return s
}

type StartDIJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// false
	ForceToRerun                *bool   `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	RealtimeStartSettingsShrink *string `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty"`
}

func (s StartDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobShrinkRequest) SetDIJobId(v string) *StartDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetForceToRerun(v bool) *StartDIJobShrinkRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest {
	s.RealtimeStartSettingsShrink = &v
	return s
}

type StartDIJobResponseBody struct {
	// example:
	//
	// 999431B2-6013-577F-B684-36F7433C753B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDIJobResponseBody) SetRequestId(v string) *StartDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDIJobResponseBody) SetSuccess(v bool) *StartDIJobResponseBody {
	s.Success = &v
	return s
}

type StartDIJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobResponse) GoString() string {
	return s.String()
}

func (s *StartDIJobResponse) SetHeaders(v map[string]*string) *StartDIJobResponse {
	s.Headers = v
	return s
}

func (s *StartDIJobResponse) SetStatusCode(v int32) *StartDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDIJobResponse) SetBody(v *StartDIJobResponseBody) *StartDIJobResponse {
	s.Body = v
	return s
}

type UpdateDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId             *int64                                   `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description         *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	JobSettings         *UpdateDIJobRequestJobSettings           `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	ResourceSettings    *UpdateDIJobRequestResourceSettings      `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	TableMappings       []*UpdateDIJobRequestTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules []*UpdateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequest) SetDIJobId(v int64) *UpdateDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobRequest) SetDescription(v string) *UpdateDIJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobRequest) SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetResourceSettings(v *UpdateDIJobRequestResourceSettings) *UpdateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetTableMappings(v []*UpdateDIJobRequestTableMappings) *UpdateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *UpdateDIJobRequest) SetTransformationRules(v []*UpdateDIJobRequestTransformationRules) *UpdateDIJobRequest {
	s.TransformationRules = v
	return s
}

type UpdateDIJobRequestJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *UpdateDIJobRequestJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*UpdateDIJobRequestJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*UpdateDIJobRequestJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestJobSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettings) SetChannelSettings(v string) *UpdateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings) *UpdateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetCycleScheduleSettings(v *UpdateDIJobRequestJobSettingsCycleScheduleSettings) *UpdateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*UpdateDIJobRequestJobSettingsDdlHandlingSettings) *UpdateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetRuntimeSettings(v []*UpdateDIJobRequestJobSettingsRuntimeSettings) *UpdateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

type UpdateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type UpdateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *UpdateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type UpdateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type UpdateDIJobRequestJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type UpdateDIJobRequestResourceSettings struct {
	OfflineResourceSettings  *UpdateDIJobRequestResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *UpdateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s UpdateDIJobRequestResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type UpdateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1721021993437
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestTableMappings struct {
	SourceObjectSelectionRules []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*UpdateDIJobRequestTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestTableMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) *UpdateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *UpdateDIJobRequestTableMappings) SetTransformationRules(v []*UpdateDIJobRequestTableMappingsTransformationRules) *UpdateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

type UpdateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type UpdateDIJobRequestTableMappingsTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type UpdateDIJobRequestTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleExpression(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

type UpdateDIJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId                   *int64  `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	JobSettingsShrink         *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	ResourceSettingsShrink    *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s UpdateDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobShrinkRequest) SetDIJobId(v int64) *UpdateDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetDescription(v string) *UpdateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

type UpdateDIJobResponseBody struct {
	// example:
	//
	// AAC30B35-820D-5F3E-A42C-E96BB6379325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponseBody) SetRequestId(v string) *UpdateDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIJobResponseBody) SetSuccess(v bool) *UpdateDIJobResponseBody {
	s.Success = &v
	return s
}

type UpdateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponse) SetHeaders(v map[string]*string) *UpdateDIJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIJobResponse) SetStatusCode(v int32) *UpdateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIJobResponse) SetBody(v *UpdateDIJobResponseBody) *UpdateDIJobResponse {
	s.Body = v
	return s
}

type UpdateDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5678
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) SetConnectionProperties(v string) *UpdateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *UpdateDataSourceRequest) SetConnectionPropertiesMode(v string) *UpdateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDescription(v string) *UpdateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceRequest) SetId(v int64) *UpdateDataSourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceRequest) SetProjectId(v int64) *UpdateDataSourceRequest {
	s.ProjectId = &v
	return s
}

type UpdateDataSourceResponseBody struct {
	// example:
	//
	// 102E8E24-0387-531D-8A75-1C0AE7DD03E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

type UpdateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponse) SetHeaders(v map[string]*string) *UpdateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceResponse) SetStatusCode(v int32) *UpdateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceResponse) SetBody(v *UpdateDataSourceResponseBody) *UpdateDataSourceResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetId(v string) *UpdateFunctionRequest {
	s.Id = &v
	return s
}

func (s *UpdateFunctionRequest) SetProjectId(v string) *UpdateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpec(v string) *UpdateFunctionRequest {
	s.Spec = &v
	return s
}

type UpdateFunctionResponseBody struct {
	// example:
	//
	// 12123960-CB2C-5086-868E-C6C1D024XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSuccess(v bool) *UpdateFunctionResponseBody {
	s.Success = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeRequest) SetId(v string) *UpdateNodeRequest {
	s.Id = &v
	return s
}

func (s *UpdateNodeRequest) SetProjectId(v string) *UpdateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateNodeRequest) SetSpec(v string) *UpdateNodeRequest {
	s.Spec = &v
	return s
}

type UpdateNodeResponseBody struct {
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponseBody) SetRequestId(v string) *UpdateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeResponseBody) SetSuccess(v bool) *UpdateNodeResponseBody {
	s.Success = &v
	return s
}

type UpdateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponse) SetHeaders(v map[string]*string) *UpdateNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeResponse) SetStatusCode(v int32) *UpdateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeResponse) SetBody(v *UpdateNodeResponseBody) *UpdateNodeResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Forbidden
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetDevEnvironmentEnabled(v bool) *UpdateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDevRoleDisabled(v bool) *UpdateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDisplayName(v string) *UpdateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateProjectRequest) SetId(v int64) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetPaiTaskEnabled(v bool) *UpdateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetStatus(v string) *UpdateProjectRequest {
	s.Status = &v
	return s
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetId(v string) *UpdateResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceRequest) SetProjectId(v string) *UpdateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceRequest) SetSpec(v string) *UpdateResourceRequest {
	s.Spec = &v
	return s
}

type UpdateResourceResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetSuccess(v bool) *UpdateResourceResponseBody {
	s.Success = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionRequest) SetId(v string) *UpdateWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetProjectId(v string) *UpdateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetSpec(v string) *UpdateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type UpdateWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 20BF7E80-668A-5620-8AD8-879B8FEAXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponseBody) SetRequestId(v string) *UpdateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponseBody) SetSuccess(v bool) *UpdateWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetStatusCode(v int32) *UpdateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetBody(v *UpdateWorkflowDefinitionResponseBody) *UpdateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            tea.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           tea.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dataworks.aliyuncs.com"),
		"eu-central-1":          tea.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
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

// Summary:
//
// 终止发布流程
//
// @param request - AbolishDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeploymentWithOptions(request *AbolishDeploymentRequest, runtime *util.RuntimeOptions) (_result *AbolishDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止发布流程
//
// @param request - AbolishDeploymentRequest
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeployment(request *AbolishDeploymentRequest) (_result *AbolishDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.AbolishDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CloneDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSourceWithOptions(request *CloneDataSourceRequest, runtime *util.RuntimeOptions) (_result *CloneDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloneDataSourceName)) {
		query["CloneDataSourceName"] = request.CloneDataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CloneDataSourceRequest
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSource(request *CloneDataSourceRequest) (_result *CloneDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CloneDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集成任务
//
// @param tmpReq - CreateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJobWithOptions(tmpReq *CreateDIJobRequest, runtime *util.RuntimeOptions) (_result *CreateDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestinationDataSourceSettings)) {
		request.DestinationDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationDataSourceSettings, tea.String("DestinationDataSourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.JobSettings)) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, tea.String("JobSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSettings)) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, tea.String("ResourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceDataSourceSettings)) {
		request.SourceDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDataSourceSettings, tea.String("SourceDataSourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TableMappings)) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, tea.String("TableMappings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TransformationRules)) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, tea.String("TransformationRules"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集成任务
//
// @param request - CreateDIJobRequest
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJob(request *CreateDIJobRequest) (_result *CreateDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CreateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPropertiesMode)) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRuleWithOptions(request *CreateDataSourceSharedRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.SharedUser)) {
		query["SharedUser"] = request.SharedUser
	}

	if !tea.BoolValue(util.IsUnset(request.TargetProjectId)) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSourceSharedRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRule(request *CreateDataSourceSharedRuleRequest) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CreateDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发布流程
//
// @param tmpReq - CreateDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeploymentWithOptions(tmpReq *CreateDeploymentRequest, runtime *util.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDeploymentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ObjectIds)) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, tea.String("ObjectIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectIdsShrink)) {
		body["ObjectIds"] = request.ObjectIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发布流程
//
// @param request - CreateDeploymentRequest
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeployment(request *CreateDeploymentRequest) (_result *CreateDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CreateDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建udf函数
//
// @param request - CreateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建udf函数
//
// @param request - CreateFunctionRequest
//
// @return CreateFunctionResponse
func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建节点
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithOptions(request *CreateNodeRequest, runtime *util.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerId)) {
		body["ContainerId"] = request.ContainerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建节点
//
// @param request - CreateNodeRequest
//
// @return CreateNodeResponse
func (client *Client) CreateNode(request *CreateNodeRequest) (_result *CreateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeResponse{}
	_body, _err := client.CreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AliyunResourceTags)) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, tea.String("AliyunResourceTags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunResourceGroupId)) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunResourceTagsShrink)) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源文件
//
// @param request - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源文件
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinitionWithOptions(request *CreateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *CreateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowDefinitionRequest
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinition(request *CreateWorkflowDefinitionRequest) (_result *CreateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CreateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集成任务
//
// @param request - DeleteDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJobWithOptions(request *DeleteDIJobRequest, runtime *util.RuntimeOptions) (_result *DeleteDIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集成任务
//
// @param request - DeleteDIJobRequest
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJob(request *DeleteDIJobRequest) (_result *DeleteDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.DeleteDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRuleWithOptions(request *DeleteDataSourceSharedRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSourceSharedRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRule(request *DeleteDataSourceSharedRuleRequest) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.DeleteDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除udf函数
//
// @param request - DeleteFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除udf函数
//
// @param request - DeleteFunctionRequest
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除节点
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithOptions(request *DeleteNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除节点
//
// @param request - DeleteNodeRequest
//
// @return DeleteNodeResponse
func (client *Client) DeleteNode(request *DeleteNodeRequest) (_result *DeleteNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DeleteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 销毁工作空间
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 销毁工作空间
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源文件
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源文件
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinitionWithOptions(request *DeleteWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinition(request *DeleteWorkflowDefinitionRequest) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.DeleteWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行Deployment一个阶段
//
// @param request - ExecDeploymentStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStageWithOptions(request *ExecDeploymentStageRequest, runtime *util.RuntimeOptions) (_result *ExecDeploymentStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecDeploymentStage"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行Deployment一个阶段
//
// @param request - ExecDeploymentStageRequest
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStage(request *ExecDeploymentStageRequest) (_result *ExecDeploymentStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.ExecDeploymentStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看数据集成任务
//
// @param request - GetDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobResponse
func (client *Client) GetDIJobWithOptions(request *GetDIJobRequest, runtime *util.RuntimeOptions) (_result *GetDIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据集成任务
//
// @param request - GetDIJobRequest
//
// @return GetDIJobResponse
func (client *Client) GetDIJob(request *GetDIJobRequest) (_result *GetDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDIJobResponse{}
	_body, _err := client.GetDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务日志
//
// @param request - GetDIJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLogWithOptions(request *GetDIJobLogRequest, runtime *util.RuntimeOptions) (_result *GetDIJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDIJobLog"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务日志
//
// @param request - GetDIJobLogRequest
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLog(request *GetDIJobLogRequest) (_result *GetDIJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.GetDIJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - GetDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithOptions(request *GetDataSourceRequest, runtime *util.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - GetDataSourceRequest
//
// @return GetDataSourceResponse
func (client *Client) GetDataSource(request *GetDataSourceRequest) (_result *GetDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布流程详情
//
// @param request - GetDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithOptions(request *GetDeploymentRequest, runtime *util.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发布流程详情
//
// @param request - GetDeploymentRequest
//
// @return GetDeploymentResponse
func (client *Client) GetDeployment(request *GetDeploymentRequest) (_result *GetDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取函数信息
//
// @param request - GetFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithOptions(request *GetFunctionRequest, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取函数信息
//
// @param request - GetFunctionRequest
//
// @return GetFunctionResponse
func (client *Client) GetFunction(request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *util.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源文件
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源文件
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - GetWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinitionWithOptions(request *GetWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *GetWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - GetWorkflowDefinitionRequest
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinition(request *GetWorkflowDefinitionRequest) (_result *GetWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.GetWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务事件
//
// @param request - ListDIJobEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEventsWithOptions(request *ListDIJobEventsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobEvents"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务事件
//
// @param request - ListDIJobEventsRequest
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEvents(request *ListDIJobEventsRequest) (_result *ListDIJobEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.ListDIJobEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务指标
//
// @param tmpReq - ListDIJobMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetricsWithOptions(tmpReq *ListDIJobMetricsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDIJobMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MetricName)) {
		request.MetricNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricName, tea.String("MetricName"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobMetrics"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务指标
//
// @param request - ListDIJobMetricsRequest
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetrics(request *ListDIJobMetricsRequest) (_result *ListDIJobMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.ListDIJobMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务
//
// @param request - ListDIJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobsWithOptions(request *ListDIJobsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobs"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务
//
// @param request - ListDIJobsRequest
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobs(request *ListDIJobsRequest) (_result *ListDIJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobsResponse{}
	_body, _err := client.ListDIJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourceSharedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRulesWithOptions(request *ListDataSourceSharedRulesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceSharedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceSharedRules"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourceSharedRulesRequest
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRules(request *ListDataSourceSharedRulesRequest) (_result *ListDataSourceSharedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.ListDataSourceSharedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param tmpReq - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(tmpReq *ListDataSourcesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Types)) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, tea.String("Types"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取发布流程
//
// @param request - ListDeploymentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(request *ListDeploymentsRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployments"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取发布流程
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
func (client *Client) ListDeployments(request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取udf函数列表
//
// @param request - ListFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取udf函数列表
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点依赖列表
//
// @param request - ListNodeDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependenciesWithOptions(request *ListNodeDependenciesRequest, runtime *util.RuntimeOptions) (_result *ListNodeDependenciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeDependencies"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点依赖列表
//
// @param request - ListNodeDependenciesRequest
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependencies(request *ListNodeDependenciesRequest) (_result *ListNodeDependenciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.ListNodeDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点列表
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点列表
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询工作空间详情
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AliyunResourceTags)) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, tea.String("AliyunResourceTags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Ids)) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, tea.String("Ids"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Names)) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, tea.String("Names"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunResourceGroupId)) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunResourceTagsShrink)) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.IdsShrink)) {
		body["Ids"] = request.IdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NamesShrink)) {
		body["Names"] = request.NamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询工作空间详情
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取资源文件
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取资源文件
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取workflowDefinition的列表
//
// @param request - ListWorkflowDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitionsWithOptions(request *ListWorkflowDefinitionsRequest, runtime *util.RuntimeOptions) (_result *ListWorkflowDefinitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflowDefinitions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取workflowDefinition的列表
//
// @param request - ListWorkflowDefinitionsRequest
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitions(request *ListWorkflowDefinitionsRequest) (_result *ListWorkflowDefinitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.ListWorkflowDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动funciton的路径
//
// @param request - MoveFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFunctionResponse
func (client *Client) MoveFunctionWithOptions(request *MoveFunctionRequest, runtime *util.RuntimeOptions) (_result *MoveFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动funciton的路径
//
// @param request - MoveFunctionRequest
//
// @return MoveFunctionResponse
func (client *Client) MoveFunction(request *MoveFunctionRequest) (_result *MoveFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveFunctionResponse{}
	_body, _err := client.MoveFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动节点路径
//
// @param request - MoveNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveNodeResponse
func (client *Client) MoveNodeWithOptions(request *MoveNodeRequest, runtime *util.RuntimeOptions) (_result *MoveNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动节点路径
//
// @param request - MoveNodeRequest
//
// @return MoveNodeResponse
func (client *Client) MoveNode(request *MoveNodeRequest) (_result *MoveNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveNodeResponse{}
	_body, _err := client.MoveNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动资源文件路径
//
// @param request - MoveResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithOptions(request *MoveResourceRequest, runtime *util.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动资源文件路径
//
// @param request - MoveResourceRequest
//
// @return MoveResourceResponse
func (client *Client) MoveResource(request *MoveResourceRequest) (_result *MoveResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceResponse{}
	_body, _err := client.MoveResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动工作流的路径
//
// @param request - MoveWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinitionWithOptions(request *MoveWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *MoveWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动工作流的路径
//
// @param request - MoveWorkflowDefinitionRequest
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinition(request *MoveWorkflowDefinitionRequest) (_result *MoveWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.MoveWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对function重命名
//
// @param request - RenameFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFunctionResponse
func (client *Client) RenameFunctionWithOptions(request *RenameFunctionRequest, runtime *util.RuntimeOptions) (_result *RenameFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对function重命名
//
// @param request - RenameFunctionRequest
//
// @return RenameFunctionResponse
func (client *Client) RenameFunction(request *RenameFunctionRequest) (_result *RenameFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameFunctionResponse{}
	_body, _err := client.RenameFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名节点
//
// @param request - RenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameNodeResponse
func (client *Client) RenameNodeWithOptions(request *RenameNodeRequest, runtime *util.RuntimeOptions) (_result *RenameNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名节点
//
// @param request - RenameNodeRequest
//
// @return RenameNodeResponse
func (client *Client) RenameNode(request *RenameNodeRequest) (_result *RenameNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameNodeResponse{}
	_body, _err := client.RenameNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对资源文件重命名
//
// @param request - RenameResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameResourceResponse
func (client *Client) RenameResourceWithOptions(request *RenameResourceRequest, runtime *util.RuntimeOptions) (_result *RenameResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对资源文件重命名
//
// @param request - RenameResourceRequest
//
// @return RenameResourceResponse
func (client *Client) RenameResource(request *RenameResourceRequest) (_result *RenameResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameResourceResponse{}
	_body, _err := client.RenameResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名工作流
//
// @param request - RenameWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinitionWithOptions(request *RenameWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *RenameWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名工作流
//
// @param request - RenameWorkflowDefinitionRequest
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinition(request *RenameWorkflowDefinitionRequest) (_result *RenameWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.RenameWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动数据集成任务
//
// @param tmpReq - StartDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDIJobResponse
func (client *Client) StartDIJobWithOptions(tmpReq *StartDIJobRequest, runtime *util.RuntimeOptions) (_result *StartDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StartDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RealtimeStartSettings)) {
		request.RealtimeStartSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeStartSettings, tea.String("RealtimeStartSettings"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动数据集成任务
//
// @param request - StartDIJobRequest
//
// @return StartDIJobResponse
func (client *Client) StartDIJob(request *StartDIJobRequest) (_result *StartDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDIJobResponse{}
	_body, _err := client.StartDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集成任务
//
// @param tmpReq - UpdateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJobWithOptions(tmpReq *UpdateDIJobRequest, runtime *util.RuntimeOptions) (_result *UpdateDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobSettings)) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, tea.String("JobSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSettings)) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, tea.String("ResourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TableMappings)) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, tea.String("TableMappings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TransformationRules)) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, tea.String("TransformationRules"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集成任务
//
// @param request - UpdateDIJobRequest
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJob(request *UpdateDIJobRequest) (_result *UpdateDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.UpdateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithOptions(request *UpdateDataSourceRequest, runtime *util.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPropertiesMode)) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - UpdateDataSourceRequest
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSource(request *UpdateDataSourceRequest) (_result *UpdateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.UpdateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新udf函数
//
// @param request - UpdateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新udf函数
//
// @param request - UpdateFunctionRequest
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新节点
//
// @param request - UpdateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeResponse
func (client *Client) UpdateNodeWithOptions(request *UpdateNodeRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新节点
//
// @param request - UpdateNodeRequest
//
// @return UpdateNodeResponse
func (client *Client) UpdateNode(request *UpdateNodeRequest) (_result *UpdateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeResponse{}
	_body, _err := client.UpdateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作空间
//
// @param request - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作空间
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源文件
//
// @param request - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(request *UpdateResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源文件
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinitionWithOptions(request *UpdateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinition(request *UpdateWorkflowDefinitionRequest) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.UpdateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
