// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FeatureViewConfigValue struct {
	Partitions map[string]*FeatureViewConfigValuePartitionsValue `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	EventTime  *string                                           `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Equal      *bool                                             `json:"Equal,omitempty" xml:"Equal,omitempty"`
	UseMock    *bool                                             `json:"UseMock,omitempty" xml:"UseMock,omitempty"`
}

func (s FeatureViewConfigValue) String() string {
	return tea.Prettify(s)
}

func (s FeatureViewConfigValue) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValue) SetPartitions(v map[string]*FeatureViewConfigValuePartitionsValue) *FeatureViewConfigValue {
	s.Partitions = v
	return s
}

func (s *FeatureViewConfigValue) SetEventTime(v string) *FeatureViewConfigValue {
	s.EventTime = &v
	return s
}

func (s *FeatureViewConfigValue) SetEqual(v bool) *FeatureViewConfigValue {
	s.Equal = &v
	return s
}

func (s *FeatureViewConfigValue) SetUseMock(v bool) *FeatureViewConfigValue {
	s.UseMock = &v
	return s
}

type FeatureViewConfigValuePartitionsValue struct {
	Value      *string   `json:"Value,omitempty" xml:"Value,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	StartValue *string   `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	EndValue   *string   `json:"EndValue,omitempty" xml:"EndValue,omitempty"`
}

func (s FeatureViewConfigValuePartitionsValue) String() string {
	return tea.Prettify(s)
}

func (s FeatureViewConfigValuePartitionsValue) GoString() string {
	return s.String()
}

func (s *FeatureViewConfigValuePartitionsValue) SetValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.Value = &v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetValues(v []*string) *FeatureViewConfigValuePartitionsValue {
	s.Values = v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetStartValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.StartValue = &v
	return s
}

func (s *FeatureViewConfigValuePartitionsValue) SetEndValue(v string) *FeatureViewConfigValuePartitionsValue {
	s.EndValue = &v
	return s
}

type CheckInstanceDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceDatasourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceDatasourceRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceRequest) SetConfig(v string) *CheckInstanceDatasourceRequest {
	s.Config = &v
	return s
}

func (s *CheckInstanceDatasourceRequest) SetType(v string) *CheckInstanceDatasourceRequest {
	s.Type = &v
	return s
}

func (s *CheckInstanceDatasourceRequest) SetUri(v string) *CheckInstanceDatasourceRequest {
	s.Uri = &v
	return s
}

type CheckInstanceDatasourceResponseBody struct {
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckInstanceDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceResponseBody) SetRequestId(v string) *CheckInstanceDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceDatasourceResponseBody) SetStatus(v string) *CheckInstanceDatasourceResponseBody {
	s.Status = &v
	return s
}

type CheckInstanceDatasourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInstanceDatasourceResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceResponse) SetHeaders(v map[string]*string) *CheckInstanceDatasourceResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceDatasourceResponse) SetStatusCode(v int32) *CheckInstanceDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceDatasourceResponse) SetBody(v *CheckInstanceDatasourceResponseBody) *CheckInstanceDatasourceResponse {
	s.Body = v
	return s
}

type CreateDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasourceRequest) SetConfig(v string) *CreateDatasourceRequest {
	s.Config = &v
	return s
}

func (s *CreateDatasourceRequest) SetName(v string) *CreateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasourceRequest) SetType(v string) *CreateDatasourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDatasourceRequest) SetUri(v string) *CreateDatasourceRequest {
	s.Uri = &v
	return s
}

func (s *CreateDatasourceRequest) SetWorkspaceId(v string) *CreateDatasourceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDatasourceResponseBody struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 1C5B1511-8A5B-59C3-90AF-513F9210E882
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponseBody) SetDatasourceId(v string) *CreateDatasourceResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *CreateDatasourceResponseBody) SetRequestId(v string) *CreateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponse) SetHeaders(v map[string]*string) *CreateDatasourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasourceResponse) SetStatusCode(v int32) *CreateDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasourceResponse) SetBody(v *CreateDatasourceResponseBody) *CreateDatasourceResponse {
	s.Body = v
	return s
}

type CreateFeatureEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateFeatureEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityRequest) SetJoinId(v string) *CreateFeatureEntityRequest {
	s.JoinId = &v
	return s
}

func (s *CreateFeatureEntityRequest) SetName(v string) *CreateFeatureEntityRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureEntityRequest) SetProjectId(v string) *CreateFeatureEntityRequest {
	s.ProjectId = &v
	return s
}

type CreateFeatureEntityResponseBody struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityResponseBody) SetFeatureEntityId(v string) *CreateFeatureEntityResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *CreateFeatureEntityResponseBody) SetRequestId(v string) *CreateFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

type CreateFeatureEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityResponse) SetHeaders(v map[string]*string) *CreateFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureEntityResponse) SetStatusCode(v int32) *CreateFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureEntityResponse) SetBody(v *CreateFeatureEntityResponseBody) *CreateFeatureEntityResponse {
	s.Body = v
	return s
}

type CreateFeatureViewRequest struct {
	// example:
	//
	// {"save_original_field":true}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 4
	FeatureEntityId *string                           `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	Fields          []*CreateFeatureViewRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// FeatureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 5
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SyncOnlineTable *bool `json:"SyncOnlineTable,omitempty" xml:"SyncOnlineTable,omitempty"`
	// example:
	//
	// 90
	TTL  *int32    `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Custom
	WriteMethod      *string `json:"WriteMethod,omitempty" xml:"WriteMethod,omitempty"`
	WriteToFeatureDB *bool   `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s CreateFeatureViewRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureViewRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequest) SetConfig(v string) *CreateFeatureViewRequest {
	s.Config = &v
	return s
}

func (s *CreateFeatureViewRequest) SetFeatureEntityId(v string) *CreateFeatureViewRequest {
	s.FeatureEntityId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetFields(v []*CreateFeatureViewRequestFields) *CreateFeatureViewRequest {
	s.Fields = v
	return s
}

func (s *CreateFeatureViewRequest) SetName(v string) *CreateFeatureViewRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureViewRequest) SetProjectId(v string) *CreateFeatureViewRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetRegisterDatasourceId(v string) *CreateFeatureViewRequest {
	s.RegisterDatasourceId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetRegisterTable(v string) *CreateFeatureViewRequest {
	s.RegisterTable = &v
	return s
}

func (s *CreateFeatureViewRequest) SetSyncOnlineTable(v bool) *CreateFeatureViewRequest {
	s.SyncOnlineTable = &v
	return s
}

func (s *CreateFeatureViewRequest) SetTTL(v int32) *CreateFeatureViewRequest {
	s.TTL = &v
	return s
}

func (s *CreateFeatureViewRequest) SetTags(v []*string) *CreateFeatureViewRequest {
	s.Tags = v
	return s
}

func (s *CreateFeatureViewRequest) SetType(v string) *CreateFeatureViewRequest {
	s.Type = &v
	return s
}

func (s *CreateFeatureViewRequest) SetWriteMethod(v string) *CreateFeatureViewRequest {
	s.WriteMethod = &v
	return s
}

func (s *CreateFeatureViewRequest) SetWriteToFeatureDB(v bool) *CreateFeatureViewRequest {
	s.WriteToFeatureDB = &v
	return s
}

type CreateFeatureViewRequestFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFeatureViewRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureViewRequestFields) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequestFields) SetAttributes(v []*string) *CreateFeatureViewRequestFields {
	s.Attributes = v
	return s
}

func (s *CreateFeatureViewRequestFields) SetName(v string) *CreateFeatureViewRequestFields {
	s.Name = &v
	return s
}

func (s *CreateFeatureViewRequestFields) SetType(v string) *CreateFeatureViewRequestFields {
	s.Type = &v
	return s
}

type CreateFeatureViewResponseBody struct {
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFeatureViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewResponseBody) SetFeatureViewId(v string) *CreateFeatureViewResponseBody {
	s.FeatureViewId = &v
	return s
}

func (s *CreateFeatureViewResponseBody) SetRequestId(v string) *CreateFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

type CreateFeatureViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureViewResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewResponse) SetHeaders(v map[string]*string) *CreateFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureViewResponse) SetStatusCode(v int32) *CreateFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureViewResponse) SetBody(v *CreateFeatureViewResponseBody) *CreateFeatureViewResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetType(v string) *CreateInstanceRequest {
	s.Type = &v
	return s
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// InstanceAlreadyExistsErrorProblem
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// featureStore-cn-7mz2xfu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateLabelTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// This parameter is required.
	Fields []*CreateLabelTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rec_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateLabelTableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTableRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelTableRequest) SetDatasourceId(v string) *CreateLabelTableRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateLabelTableRequest) SetFields(v []*CreateLabelTableRequestFields) *CreateLabelTableRequest {
	s.Fields = v
	return s
}

func (s *CreateLabelTableRequest) SetName(v string) *CreateLabelTableRequest {
	s.Name = &v
	return s
}

func (s *CreateLabelTableRequest) SetProjectId(v string) *CreateLabelTableRequest {
	s.ProjectId = &v
	return s
}

type CreateLabelTableRequestFields struct {
	// This parameter is required.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// lat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLabelTableRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTableRequestFields) GoString() string {
	return s.String()
}

func (s *CreateLabelTableRequestFields) SetAttributes(v []*string) *CreateLabelTableRequestFields {
	s.Attributes = v
	return s
}

func (s *CreateLabelTableRequestFields) SetName(v string) *CreateLabelTableRequestFields {
	s.Name = &v
	return s
}

func (s *CreateLabelTableRequestFields) SetType(v string) *CreateLabelTableRequestFields {
	s.Type = &v
	return s
}

type CreateLabelTableResponseBody struct {
	// example:
	//
	// 1
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// 0FA90B3B-F30A-5C9D-A9FD-8114F8868062
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLabelTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelTableResponseBody) SetLabelTableId(v string) *CreateLabelTableResponseBody {
	s.LabelTableId = &v
	return s
}

func (s *CreateLabelTableResponseBody) SetRequestId(v string) *CreateLabelTableResponseBody {
	s.RequestId = &v
	return s
}

type CreateLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLabelTableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTableResponse) GoString() string {
	return s.String()
}

func (s *CreateLabelTableResponse) SetHeaders(v map[string]*string) *CreateLabelTableResponse {
	s.Headers = v
	return s
}

func (s *CreateLabelTableResponse) SetStatusCode(v int32) *CreateLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabelTableResponse) SetBody(v *CreateLabelTableResponseBody) *CreateLabelTableResponse {
	s.Body = v
	return s
}

type CreateModelFeatureRequest struct {
	// This parameter is required.
	Features []*CreateModelFeatureRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model_feature_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId              *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SequenceFeatureViewIds []*string `json:"SequenceFeatureViewIds,omitempty" xml:"SequenceFeatureViewIds,omitempty" type:"Repeated"`
}

func (s CreateModelFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelFeatureRequest) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureRequest) SetFeatures(v []*CreateModelFeatureRequestFeatures) *CreateModelFeatureRequest {
	s.Features = v
	return s
}

func (s *CreateModelFeatureRequest) SetLabelPriorityLevel(v int64) *CreateModelFeatureRequest {
	s.LabelPriorityLevel = &v
	return s
}

func (s *CreateModelFeatureRequest) SetLabelTableId(v string) *CreateModelFeatureRequest {
	s.LabelTableId = &v
	return s
}

func (s *CreateModelFeatureRequest) SetName(v string) *CreateModelFeatureRequest {
	s.Name = &v
	return s
}

func (s *CreateModelFeatureRequest) SetProjectId(v string) *CreateModelFeatureRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateModelFeatureRequest) SetSequenceFeatureViewIds(v []*string) *CreateModelFeatureRequest {
	s.SequenceFeatureViewIds = v
	return s
}

type CreateModelFeatureRequestFeatures struct {
	// example:
	//
	// userid
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateModelFeatureRequestFeatures) String() string {
	return tea.Prettify(s)
}

func (s CreateModelFeatureRequestFeatures) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureRequestFeatures) SetAliasName(v string) *CreateModelFeatureRequestFeatures {
	s.AliasName = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetFeatureViewId(v string) *CreateModelFeatureRequestFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetName(v string) *CreateModelFeatureRequestFeatures {
	s.Name = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetType(v string) *CreateModelFeatureRequestFeatures {
	s.Type = &v
	return s
}

type CreateModelFeatureResponseBody struct {
	// example:
	//
	// 3
	ModelFeatureId *string `json:"ModelFeatureId,omitempty" xml:"ModelFeatureId,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureResponseBody) SetModelFeatureId(v string) *CreateModelFeatureResponseBody {
	s.ModelFeatureId = &v
	return s
}

func (s *CreateModelFeatureResponseBody) SetRequestId(v string) *CreateModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

type CreateModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureResponse) SetHeaders(v map[string]*string) *CreateModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *CreateModelFeatureResponse) SetStatusCode(v int32) *CreateModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelFeatureResponse) SetBody(v *CreateModelFeatureResponseBody) *CreateModelFeatureResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	OfflineDatasourceId *string `json:"OfflineDatasourceId,omitempty" xml:"OfflineDatasourceId,omitempty"`
	// example:
	//
	// 90
	OfflineLifeCycle *int32 `json:"OfflineLifeCycle,omitempty" xml:"OfflineLifeCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	OnlineDatasourceId *string `json:"OnlineDatasourceId,omitempty" xml:"OnlineDatasourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 324
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetOfflineDatasourceId(v string) *CreateProjectRequest {
	s.OfflineDatasourceId = &v
	return s
}

func (s *CreateProjectRequest) SetOfflineLifeCycle(v int32) *CreateProjectRequest {
	s.OfflineLifeCycle = &v
	return s
}

func (s *CreateProjectRequest) SetOnlineDatasourceId(v string) *CreateProjectRequest {
	s.OnlineDatasourceId = &v
	return s
}

func (s *CreateProjectRequest) SetWorkspaceId(v string) *CreateProjectRequest {
	s.WorkspaceId = &v
	return s
}

type CreateProjectResponseBody struct {
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 977ADE3A-D69D-58AD-8323-96E2FB898E99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
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

type CreateServiceIdentityRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForFeatureStore
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleRequest) SetRoleName(v string) *CreateServiceIdentityRoleRequest {
	s.RoleName = &v
	return s
}

type CreateServiceIdentityRoleResponseBody struct {
	// example:
	//
	// ServiceLinkedRoleAlreadyExistsErrorProblem
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForFeatureStore
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponseBody) SetCode(v string) *CreateServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRequestId(v string) *CreateServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleName(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type CreateServiceIdentityRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *CreateServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetStatusCode(v int32) *CreateServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetBody(v *CreateServiceIdentityRoleResponseBody) *CreateServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type DeleteDatasourceResponseBody struct {
	// example:
	//
	// E2E1575F-29D1-5579-B649-B7883A793562
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceResponseBody) SetRequestId(v string) *DeleteDatasourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceResponse) SetHeaders(v map[string]*string) *DeleteDatasourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasourceResponse) SetStatusCode(v int32) *DeleteDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasourceResponse) SetBody(v *DeleteDatasourceResponseBody) *DeleteDatasourceResponse {
	s.Body = v
	return s
}

type DeleteFeatureEntityResponseBody struct {
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFeatureEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureEntityResponseBody) SetRequestId(v string) *DeleteFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFeatureEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureEntityResponse) SetHeaders(v map[string]*string) *DeleteFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureEntityResponse) SetStatusCode(v int32) *DeleteFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureEntityResponse) SetBody(v *DeleteFeatureEntityResponseBody) *DeleteFeatureEntityResponse {
	s.Body = v
	return s
}

type DeleteFeatureViewResponseBody struct {
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFeatureViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureViewResponseBody) SetRequestId(v string) *DeleteFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFeatureViewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureViewResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureViewResponse) SetHeaders(v map[string]*string) *DeleteFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureViewResponse) SetStatusCode(v int32) *DeleteFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureViewResponse) SetBody(v *DeleteFeatureViewResponseBody) *DeleteFeatureViewResponse {
	s.Body = v
	return s
}

type DeleteLabelTableResponseBody struct {
	// example:
	//
	// FFD39C0F-DD8D-51B2-864E-2842206DB0E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelTableResponseBody) SetRequestId(v string) *DeleteLabelTableResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLabelTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelTableResponse) SetHeaders(v map[string]*string) *DeleteLabelTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelTableResponse) SetStatusCode(v int32) *DeleteLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelTableResponse) SetBody(v *DeleteLabelTableResponseBody) *DeleteLabelTableResponse {
	s.Body = v
	return s
}

type DeleteModelFeatureResponseBody struct {
	// example:
	//
	// 6B662A64-E4BF-56F8-BF5F-4C63F34EC0A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelFeatureResponseBody) SetRequestId(v string) *DeleteModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelFeatureResponse) SetHeaders(v map[string]*string) *DeleteModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelFeatureResponse) SetStatusCode(v int32) *DeleteModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelFeatureResponse) SetBody(v *DeleteModelFeatureResponseBody) *DeleteModelFeatureResponse {
	s.Body = v
	return s
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// 0DA35264-0877-5852-8971-7735B547C969
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

type ExportModelFeatureTrainingSetTableRequest struct {
	FeatureViewConfig           map[string]*FeatureViewConfigValue                          `json:"FeatureViewConfig,omitempty" xml:"FeatureViewConfig,omitempty"`
	LabelInputConfig            *ExportModelFeatureTrainingSetTableRequestLabelInputConfig  `json:"LabelInputConfig,omitempty" xml:"LabelInputConfig,omitempty" type:"Struct"`
	RealTimeIterateInterval     *int64                                                      `json:"RealTimeIterateInterval,omitempty" xml:"RealTimeIterateInterval,omitempty"`
	RealTimePartitionCountValue *int64                                                      `json:"RealTimePartitionCountValue,omitempty" xml:"RealTimePartitionCountValue,omitempty"`
	TrainingSetConfig           *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig `json:"TrainingSetConfig,omitempty" xml:"TrainingSetConfig,omitempty" type:"Struct"`
}

func (s ExportModelFeatureTrainingSetTableRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequest) GoString() string {
	return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetFeatureViewConfig(v map[string]*FeatureViewConfigValue) *ExportModelFeatureTrainingSetTableRequest {
	s.FeatureViewConfig = v
	return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetLabelInputConfig(v *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) *ExportModelFeatureTrainingSetTableRequest {
	s.LabelInputConfig = v
	return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetRealTimeIterateInterval(v int64) *ExportModelFeatureTrainingSetTableRequest {
	s.RealTimeIterateInterval = &v
	return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetRealTimePartitionCountValue(v int64) *ExportModelFeatureTrainingSetTableRequest {
	s.RealTimePartitionCountValue = &v
	return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetTrainingSetConfig(v *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) *ExportModelFeatureTrainingSetTableRequest {
	s.TrainingSetConfig = v
	return s
}

type ExportModelFeatureTrainingSetTableRequestLabelInputConfig struct {
	// example:
	//
	// 2022-07-02 00:00:00
	EventTime  *string                           `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Partitions map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableRequestLabelInputConfig) String() string {
	return tea.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequestLabelInputConfig) GoString() string {
	return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) SetEventTime(v string) *ExportModelFeatureTrainingSetTableRequestLabelInputConfig {
	s.EventTime = &v
	return s
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) SetPartitions(v map[string]map[string]interface{}) *ExportModelFeatureTrainingSetTableRequestLabelInputConfig {
	s.Partitions = v
	return s
}

type ExportModelFeatureTrainingSetTableRequestTrainingSetConfig struct {
	Partitions map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) String() string {
	return tea.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) GoString() string {
	return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) SetPartitions(v map[string]map[string]interface{}) *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig {
	s.Partitions = v
	return s
}

type ExportModelFeatureTrainingSetTableResponseBody struct {
	// example:
	//
	// 0FBBE454-9BD1-5D8F-9129-D14DB7FAFE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableResponseBody) GoString() string {
	return s.String()
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) SetRequestId(v string) *ExportModelFeatureTrainingSetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportModelFeatureTrainingSetTableResponseBody) SetTaskId(v string) *ExportModelFeatureTrainingSetTableResponseBody {
	s.TaskId = &v
	return s
}

type ExportModelFeatureTrainingSetTableResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportModelFeatureTrainingSetTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableResponse) GoString() string {
	return s.String()
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetHeaders(v map[string]*string) *ExportModelFeatureTrainingSetTableResponse {
	s.Headers = v
	return s
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetStatusCode(v int32) *ExportModelFeatureTrainingSetTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportModelFeatureTrainingSetTableResponse) SetBody(v *ExportModelFeatureTrainingSetTableResponseBody) *ExportModelFeatureTrainingSetTableResponse {
	s.Body = v
	return s
}

type GetDatasourceResponseBody struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AD7D9E95-BD31-53F2-B710-6C01866FCB05
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 32244
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasourceResponseBody) SetConfig(v string) *GetDatasourceResponseBody {
	s.Config = &v
	return s
}

func (s *GetDatasourceResponseBody) SetDatasourceId(v string) *GetDatasourceResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *GetDatasourceResponseBody) SetGmtCreateTime(v string) *GetDatasourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasourceResponseBody) SetGmtModifiedTime(v string) *GetDatasourceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasourceResponseBody) SetName(v string) *GetDatasourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetDatasourceResponseBody) SetRequestId(v string) *GetDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasourceResponseBody) SetType(v string) *GetDatasourceResponseBody {
	s.Type = &v
	return s
}

func (s *GetDatasourceResponseBody) SetUri(v string) *GetDatasourceResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasourceResponseBody) SetWorkspaceId(v string) *GetDatasourceResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetDatasourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasourceResponse) GoString() string {
	return s.String()
}

func (s *GetDatasourceResponse) SetHeaders(v map[string]*string) *GetDatasourceResponse {
	s.Headers = v
	return s
}

func (s *GetDatasourceResponse) SetStatusCode(v int32) *GetDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasourceResponse) SetBody(v *GetDatasourceResponseBody) *GetDatasourceResponse {
	s.Body = v
	return s
}

type GetDatasourceTableResponseBody struct {
	Fields []*GetDatasourceTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// D7B2F8C4-49C7-5CFA-8075-9D715A114873
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDatasourceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasourceTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponseBody) SetFields(v []*GetDatasourceTableResponseBodyFields) *GetDatasourceTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetDatasourceTableResponseBody) SetRequestId(v string) *GetDatasourceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasourceTableResponseBody) SetTableName(v string) *GetDatasourceTableResponseBody {
	s.TableName = &v
	return s
}

type GetDatasourceTableResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// field1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDatasourceTableResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetDatasourceTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponseBodyFields) SetAttributes(v []*string) *GetDatasourceTableResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetDatasourceTableResponseBodyFields) SetName(v string) *GetDatasourceTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetDatasourceTableResponseBodyFields) SetType(v string) *GetDatasourceTableResponseBodyFields {
	s.Type = &v
	return s
}

type GetDatasourceTableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasourceTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasourceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasourceTableResponse) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponse) SetHeaders(v map[string]*string) *GetDatasourceTableResponse {
	s.Headers = v
	return s
}

func (s *GetDatasourceTableResponse) SetStatusCode(v int32) *GetDatasourceTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasourceTableResponse) SetBody(v *GetDatasourceTableResponseBody) *GetDatasourceTableResponse {
	s.Body = v
	return s
}

type GetFeatureEntityResponseBody struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFeatureEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureEntityResponseBody) SetGmtCreateTime(v string) *GetFeatureEntityResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetJoinId(v string) *GetFeatureEntityResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetName(v string) *GetFeatureEntityResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetOwner(v string) *GetFeatureEntityResponseBody {
	s.Owner = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetProjectId(v string) *GetFeatureEntityResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetProjectName(v string) *GetFeatureEntityResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetRequestId(v string) *GetFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

type GetFeatureEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureEntityResponse) SetHeaders(v map[string]*string) *GetFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureEntityResponse) SetStatusCode(v int32) *GetFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureEntityResponse) SetBody(v *GetFeatureEntityResponseBody) *GetFeatureEntityResponse {
	s.Body = v
	return s
}

type GetFeatureViewResponseBody struct {
	// example:
	//
	// {"save_original_field":true}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// featureEntity1
	FeatureEntityName *string                             `json:"FeatureEntityName,omitempty" xml:"FeatureEntityName,omitempty"`
	Fields            []*GetFeatureViewResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtSyncTime *string `json:"GmtSyncTime,omitempty" xml:"GmtSyncTime,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// {
	//
	// 	"mode": "overwrite",
	//
	// 	"partitions": {
	//
	// 		"ds": {
	//
	// 			"value": "20230820"
	//
	// 		}
	//
	// 	},
	//
	// 	"event_time": "",
	//
	// 	"config": {},
	//
	// 	"offline_to_online": true
	//
	// }
	LastSyncConfig *string `json:"LastSyncConfig,omitempty" xml:"LastSyncConfig,omitempty"`
	// example:
	//
	// item_table_mock_1
	MockTableName *string `json:"MockTableName,omitempty" xml:"MockTableName,omitempty"`
	// example:
	//
	// featureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321421412****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// from feature_store_py.fs_client import FeatureStoreClient\nimport datetime\nfrom feature_store_py.fs_datasource import MaxComputeDataSource\nimport sys\n\ncur_day = args[\"dt\"]\nprint(\"cur_day = \", cur_day)\n\naccess_key_id = o.account.access_id\naccess_key_secret = o.account.secret_access_key\nfs = FeatureStoreClient(access_key_id=access_key_id, access_key_secret=access_key_secret, region=\"cn-beijing\")\ncur_project_name = \"p1\"\nproject = fs.get_project(cur_project_name)\n\nfeature_view_name = \"user_fea\"\nbatch_feature_view = project.get_feature_view(feature_view_name)\ntask = batch_feature_view.publish_table(partitions={\"ds\":cur_day}, mode=\"Overwrite\")\ntask.wait()\ntask.print_summary()\n
	PublishTableScript *string `json:"PublishTableScript,omitempty" xml:"PublishTableScript,omitempty"`
	// example:
	//
	// 4
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	RegisterDatasourceName *string `json:"RegisterDatasourceName,omitempty" xml:"RegisterDatasourceName,omitempty"`
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// example:
	//
	// 72F15A8A-5A28-5B18-A0DE-0EABD7D3245A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SyncOnlineTable *bool `json:"SyncOnlineTable,omitempty" xml:"SyncOnlineTable,omitempty"`
	// example:
	//
	// 90
	TTL  *int32    `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// custom
	WriteMethod      *string `json:"WriteMethod,omitempty" xml:"WriteMethod,omitempty"`
	WriteToFeatureDB *bool   `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s GetFeatureViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBody) SetConfig(v string) *GetFeatureViewResponseBody {
	s.Config = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFeatureEntityId(v string) *GetFeatureViewResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFeatureEntityName(v string) *GetFeatureViewResponseBody {
	s.FeatureEntityName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetFields(v []*GetFeatureViewResponseBodyFields) *GetFeatureViewResponseBody {
	s.Fields = v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtCreateTime(v string) *GetFeatureViewResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtModifiedTime(v string) *GetFeatureViewResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetGmtSyncTime(v string) *GetFeatureViewResponseBody {
	s.GmtSyncTime = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetJoinId(v string) *GetFeatureViewResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetLastSyncConfig(v string) *GetFeatureViewResponseBody {
	s.LastSyncConfig = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetMockTableName(v string) *GetFeatureViewResponseBody {
	s.MockTableName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetName(v string) *GetFeatureViewResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetOwner(v string) *GetFeatureViewResponseBody {
	s.Owner = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetProjectId(v string) *GetFeatureViewResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetProjectName(v string) *GetFeatureViewResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetPublishTableScript(v string) *GetFeatureViewResponseBody {
	s.PublishTableScript = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterDatasourceId(v string) *GetFeatureViewResponseBody {
	s.RegisterDatasourceId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterDatasourceName(v string) *GetFeatureViewResponseBody {
	s.RegisterDatasourceName = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRegisterTable(v string) *GetFeatureViewResponseBody {
	s.RegisterTable = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetRequestId(v string) *GetFeatureViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetSyncOnlineTable(v bool) *GetFeatureViewResponseBody {
	s.SyncOnlineTable = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetTTL(v int32) *GetFeatureViewResponseBody {
	s.TTL = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetTags(v []*string) *GetFeatureViewResponseBody {
	s.Tags = v
	return s
}

func (s *GetFeatureViewResponseBody) SetType(v string) *GetFeatureViewResponseBody {
	s.Type = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetWriteMethod(v string) *GetFeatureViewResponseBody {
	s.WriteMethod = &v
	return s
}

func (s *GetFeatureViewResponseBody) SetWriteToFeatureDB(v bool) *GetFeatureViewResponseBody {
	s.WriteToFeatureDB = &v
	return s
}

type GetFeatureViewResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// int
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFeatureViewResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureViewResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponseBodyFields) SetAttributes(v []*string) *GetFeatureViewResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetFeatureViewResponseBodyFields) SetName(v string) *GetFeatureViewResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetFeatureViewResponseBodyFields) SetType(v string) *GetFeatureViewResponseBodyFields {
	s.Type = &v
	return s
}

type GetFeatureViewResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureViewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponse) SetHeaders(v map[string]*string) *GetFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureViewResponse) SetStatusCode(v int32) *GetFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureViewResponse) SetBody(v *GetFeatureViewResponseBody) *GetFeatureViewResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	FeatureDBInstanceInfo *GetInstanceResponseBodyFeatureDBInstanceInfo `json:"FeatureDBInstanceInfo,omitempty" xml:"FeatureDBInstanceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0.8
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1C5B1511-8A5B-59C3-90AF-513F9210E882
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetFeatureDBInstanceInfo(v *GetInstanceResponseBodyFeatureDBInstanceInfo) *GetInstanceResponseBody {
	s.FeatureDBInstanceInfo = v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetProgress(v float64) *GetInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetType(v string) *GetInstanceResponseBody {
	s.Type = &v
	return s
}

type GetInstanceResponseBodyFeatureDBInstanceInfo struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyFeatureDBInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyFeatureDBInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyFeatureDBInstanceInfo) SetStatus(v string) *GetInstanceResponseBodyFeatureDBInstanceInfo {
	s.Status = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetLabelTableResponseBody struct {
	// example:
	//
	// 1
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// datasource1
	DatasourceName *string                            `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	Fields         []*GetLabelTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321312*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName          *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RelatedModelFeatures []*string `json:"RelatedModelFeatures,omitempty" xml:"RelatedModelFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLabelTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponseBody) SetDatasourceId(v string) *GetLabelTableResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *GetLabelTableResponseBody) SetDatasourceName(v string) *GetLabelTableResponseBody {
	s.DatasourceName = &v
	return s
}

func (s *GetLabelTableResponseBody) SetFields(v []*GetLabelTableResponseBodyFields) *GetLabelTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetLabelTableResponseBody) SetGmtCreateTime(v string) *GetLabelTableResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLabelTableResponseBody) SetGmtModifiedTime(v string) *GetLabelTableResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetLabelTableResponseBody) SetName(v string) *GetLabelTableResponseBody {
	s.Name = &v
	return s
}

func (s *GetLabelTableResponseBody) SetOwner(v string) *GetLabelTableResponseBody {
	s.Owner = &v
	return s
}

func (s *GetLabelTableResponseBody) SetProjectId(v string) *GetLabelTableResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetLabelTableResponseBody) SetProjectName(v string) *GetLabelTableResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetLabelTableResponseBody) SetRelatedModelFeatures(v []*string) *GetLabelTableResponseBody {
	s.RelatedModelFeatures = v
	return s
}

func (s *GetLabelTableResponseBody) SetRequestId(v string) *GetLabelTableResponseBody {
	s.RequestId = &v
	return s
}

type GetLabelTableResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// field1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLabelTableResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponseBodyFields) SetAttributes(v []*string) *GetLabelTableResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetLabelTableResponseBodyFields) SetName(v string) *GetLabelTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetLabelTableResponseBodyFields) SetType(v string) *GetLabelTableResponseBodyFields {
	s.Type = &v
	return s
}

type GetLabelTableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLabelTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTableResponse) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponse) SetHeaders(v map[string]*string) *GetLabelTableResponse {
	s.Headers = v
	return s
}

func (s *GetLabelTableResponse) SetStatusCode(v int32) *GetLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelTableResponse) SetBody(v *GetLabelTableResponseBody) *GetLabelTableResponse {
	s.Body = v
	return s
}

type GetModelFeatureResponseBody struct {
	// example:
	//
	// from feature_store_py.fs_client import FeatureStoreClient\nfrom feature_store_py.fs_project import FeatureStoreProject\nfrom feature_store_py.fs_datasource import LabelInput, MaxComputeDataSource, TrainingSetOutput\nfrom feature_store_py.fs_features import FeatureSelector\nfrom feature_store_py.fs_config import LabelInputConfig, PartitionConfig, FeatureViewConfig\nfrom feature_store_py.fs_config import TrainSetOutputConfig, EASDeployConfig\nimport datetime\nimport sys\n\ncur_day = args[\"dt\"]\nprint(\"cur_day = \", cur_day)\noffset = datetime.timedelta(days=-1)\npre_day = (datetime.datetime.strptime(cur_day, \"%Y%m%d\") + offset).strftime(\"%Y%m%d\")\nprint(\"pre_day = \", pre_day)\n\n\naccess_key_id = o.account.access_id\naccess_key_secret = o.account.secret_access_key\nfs = FeatureStoreClient(access_key_id=access_key_id, access_key_secret=access_key_secret, region=\"cn-beijing\")\ncur_project_name = \"p1\"\nproject = fs.get_project(cur_project_name)\n\nlabel_partitions = PartitionConfig(name = \"ds\", value = cur_day)\nlabel_input_config = LabelInputConfig(partition_config=label_partitions)\n\nfeature_view_1_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_1_config = FeatureViewConfig(name = \"user_fea\",\npartition_config=feature_view_1_partitions)\n\nfeature_view_2_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_2_config = FeatureViewConfig(name = \"seq_fea\",\npartition_config=feature_view_2_partitions)\n\nfeature_view_3_partitions = PartitionConfig(name = \"ds\", value = pre_day)\nfeature_view_3_config = FeatureViewConfig(name = \"item_fea\",\npartition_config=feature_view_3_partitions)\n\nfeature_view_config_list = [feature_view_1_config,feature_view_2_config,feature_view_3_config]\ntrain_set_partitions = PartitionConfig(name = \"ds\", value = cur_day)\ntrain_set_output_config = TrainSetOutputConfig(partition_config=train_set_partitions)\n\n\nmodel_name = \"rank_v1\"\ncur_model = project.get_model(model_name)\ntask = cur_model.export_train_set(label_input_config, feature_view_config_list, train_set_output_config)\ntask.wait()\nprint(\"task_summary = \", task.task_summary)\n
	ExportTrainingSetTableScript *string                                `json:"ExportTrainingSetTableScript,omitempty" xml:"ExportTrainingSetTableScript,omitempty"`
	Features                     []*GetModelFeatureResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// example:
	//
	// 3
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// label_table1
	LabelTableName *string `json:"LabelTableName,omitempty" xml:"LabelTableName,omitempty"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1231243253****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 5
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string                               `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Relations   *GetModelFeatureResponseBodyRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Struct"`
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// table2
	TrainingSetFGTable *string `json:"TrainingSetFGTable,omitempty" xml:"TrainingSetFGTable,omitempty"`
	// example:
	//
	// table1
	TrainingSetTable *string `json:"TrainingSetTable,omitempty" xml:"TrainingSetTable,omitempty"`
}

func (s GetModelFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBody) SetExportTrainingSetTableScript(v string) *GetModelFeatureResponseBody {
	s.ExportTrainingSetTableScript = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetFeatures(v []*GetModelFeatureResponseBodyFeatures) *GetModelFeatureResponseBody {
	s.Features = v
	return s
}

func (s *GetModelFeatureResponseBody) SetGmtCreateTime(v string) *GetModelFeatureResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetGmtModifiedTime(v string) *GetModelFeatureResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelPriorityLevel(v int64) *GetModelFeatureResponseBody {
	s.LabelPriorityLevel = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelTableId(v string) *GetModelFeatureResponseBody {
	s.LabelTableId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetLabelTableName(v string) *GetModelFeatureResponseBody {
	s.LabelTableName = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetName(v string) *GetModelFeatureResponseBody {
	s.Name = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetOwner(v string) *GetModelFeatureResponseBody {
	s.Owner = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetProjectId(v string) *GetModelFeatureResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetProjectName(v string) *GetModelFeatureResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetRelations(v *GetModelFeatureResponseBodyRelations) *GetModelFeatureResponseBody {
	s.Relations = v
	return s
}

func (s *GetModelFeatureResponseBody) SetRequestId(v string) *GetModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetTrainingSetFGTable(v string) *GetModelFeatureResponseBody {
	s.TrainingSetFGTable = &v
	return s
}

func (s *GetModelFeatureResponseBody) SetTrainingSetTable(v string) *GetModelFeatureResponseBody {
	s.TrainingSetTable = &v
	return s
}

type GetModelFeatureResponseBodyFeatures struct {
	// example:
	//
	// feature2
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// feature_view_1
	FeatureViewName *string `json:"FeatureViewName,omitempty" xml:"FeatureViewName,omitempty"`
	// example:
	//
	// feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetModelFeatureResponseBodyFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyFeatures) SetAliasName(v string) *GetModelFeatureResponseBodyFeatures {
	s.AliasName = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetFeatureViewId(v string) *GetModelFeatureResponseBodyFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetFeatureViewName(v string) *GetModelFeatureResponseBodyFeatures {
	s.FeatureViewName = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetName(v string) *GetModelFeatureResponseBodyFeatures {
	s.Name = &v
	return s
}

func (s *GetModelFeatureResponseBodyFeatures) SetType(v string) *GetModelFeatureResponseBodyFeatures {
	s.Type = &v
	return s
}

type GetModelFeatureResponseBodyRelations struct {
	Domains []*GetModelFeatureResponseBodyRelationsDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	Links   []*GetModelFeatureResponseBodyRelationsLinks   `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
}

func (s GetModelFeatureResponseBodyRelations) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelations) SetDomains(v []*GetModelFeatureResponseBodyRelationsDomains) *GetModelFeatureResponseBodyRelations {
	s.Domains = v
	return s
}

func (s *GetModelFeatureResponseBodyRelations) SetLinks(v []*GetModelFeatureResponseBodyRelationsLinks) *GetModelFeatureResponseBodyRelations {
	s.Links = v
	return s
}

type GetModelFeatureResponseBodyRelationsDomains struct {
	// example:
	//
	// FeatureEntity
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// Domain ID。
	//
	// example:
	//
	// 3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetModelFeatureResponseBodyRelationsDomains) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelationsDomains) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetDomainType(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.DomainType = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetId(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.Id = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsDomains) SetName(v string) *GetModelFeatureResponseBodyRelationsDomains {
	s.Name = &v
	return s
}

type GetModelFeatureResponseBodyRelationsLinks struct {
	// example:
	//
	// model_feature_2
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// user_id
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// feature_entity_3
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetModelFeatureResponseBodyRelationsLinks) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponseBodyRelationsLinks) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetFrom(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.From = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetLink(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.Link = &v
	return s
}

func (s *GetModelFeatureResponseBodyRelationsLinks) SetTo(v string) *GetModelFeatureResponseBodyRelationsLinks {
	s.To = &v
	return s
}

type GetModelFeatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponse) SetHeaders(v map[string]*string) *GetModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureResponse) SetStatusCode(v int32) *GetModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureResponse) SetBody(v *GetModelFeatureResponseBody) *GetModelFeatureResponse {
	s.Body = v
	return s
}

type GetModelFeatureFGFeatureResponseBody struct {
	LookupFeatures []*GetModelFeatureFGFeatureResponseBodyLookupFeatures `json:"LookupFeatures,omitempty" xml:"LookupFeatures,omitempty" type:"Repeated"`
	RawFeatures    []*GetModelFeatureFGFeatureResponseBodyRawFeatures    `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Reserves         []*string                                               `json:"Reserves,omitempty" xml:"Reserves,omitempty" type:"Repeated"`
	SequenceFeatures []*GetModelFeatureFGFeatureResponseBodySequenceFeatures `json:"SequenceFeatures,omitempty" xml:"SequenceFeatures,omitempty" type:"Repeated"`
}

func (s GetModelFeatureFGFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBody) SetLookupFeatures(v []*GetModelFeatureFGFeatureResponseBodyLookupFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.LookupFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetRawFeatures(v []*GetModelFeatureFGFeatureResponseBodyRawFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.RawFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetRequestId(v string) *GetModelFeatureFGFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetReserves(v []*string) *GetModelFeatureFGFeatureResponseBody {
	s.Reserves = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetSequenceFeatures(v []*GetModelFeatureFGFeatureResponseBodySequenceFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.SequenceFeatures = v
	return s
}

type GetModelFeatureFGFeatureResponseBodyLookupFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// Item
	KeyFeatureDomain *string `json:"KeyFeatureDomain,omitempty" xml:"KeyFeatureDomain,omitempty"`
	// example:
	//
	// 1
	KeyFeatureName *string `json:"KeyFeatureName,omitempty" xml:"KeyFeatureName,omitempty"`
	// example:
	//
	// User
	MapFeatureDomain *string `json:"MapFeatureDomain,omitempty" xml:"MapFeatureDomain,omitempty"`
	// example:
	//
	// item_id
	MapFeatureName *string `json:"MapFeatureName,omitempty" xml:"MapFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodyLookupFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodyLookupFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetKeyFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.KeyFeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetKeyFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.KeyFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetMapFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.MapFeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetMapFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.MapFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.ValueType = &v
	return s
}

type GetModelFeatureFGFeatureResponseBodyRawFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodyRawFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodyRawFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureType(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetInputFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.ValueType = &v
	return s
}

type GetModelFeatureFGFeatureResponseBodySequenceFeatures struct {
	// example:
	//
	// #
	AttributeDelim *string `json:"AttributeDelim,omitempty" xml:"AttributeDelim,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// ;
	SequenceDelim *string `json:"SequenceDelim,omitempty" xml:"SequenceDelim,omitempty"`
	// example:
	//
	// 50
	SequenceLength *int64                                                             `json:"SequenceLength,omitempty" xml:"SequenceLength,omitempty"`
	SubFeatures    []*GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures `json:"SubFeatures,omitempty" xml:"SubFeatures,omitempty" type:"Repeated"`
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetAttributeDelim(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.AttributeDelim = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSequenceDelim(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SequenceDelim = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSequenceLength(v int64) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SequenceLength = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSubFeatures(v []*GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SubFeatures = v
	return s
}

type GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureType(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetInputFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.ValueType = &v
	return s
}

type GetModelFeatureFGFeatureResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureFGFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureFGFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponse) SetHeaders(v map[string]*string) *GetModelFeatureFGFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureFGFeatureResponse) SetStatusCode(v int32) *GetModelFeatureFGFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponse) SetBody(v *GetModelFeatureFGFeatureResponseBody) *GetModelFeatureFGFeatureResponse {
	s.Body = v
	return s
}

type GetModelFeatureFGInfoResponseBody struct {
	// example:
	//
	// {"features": [{"feature_name": "item_id","feature_type": "id_feature","value_type": "String","expression": "item:item_id","default_value": "-1024","combiner": "mean","need_prefix": false},{"feature_name": "f1","feature_type": "lookup_feature","value_type": "Integer","map": "item:f1","key": "user:1","default_value": "0","combiner": "mean","need_prefix": false,"needDiscrete": false,"needWeighting": false,"needKey": false}],"reserves": ["f1"]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 6B662A64-E4BF-56F8-BF5F-4C63F34EC0A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelFeatureFGInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGInfoResponseBody) SetContent(v string) *GetModelFeatureFGInfoResponseBody {
	s.Content = &v
	return s
}

func (s *GetModelFeatureFGInfoResponseBody) SetRequestId(v string) *GetModelFeatureFGInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetModelFeatureFGInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureFGInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureFGInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelFeatureFGInfoResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGInfoResponse) SetHeaders(v map[string]*string) *GetModelFeatureFGInfoResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureFGInfoResponse) SetStatusCode(v int32) *GetModelFeatureFGInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureFGInfoResponse) SetBody(v *GetModelFeatureFGInfoResponseBody) *GetModelFeatureFGInfoResponse {
	s.Body = v
	return s
}

type GetProjectResponseBody struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10
	FeatureEntityCount *int32 `json:"FeatureEntityCount,omitempty" xml:"FeatureEntityCount,omitempty"`
	// example:
	//
	// 10
	FeatureViewCount *int32 `json:"FeatureViewCount,omitempty" xml:"FeatureViewCount,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 5
	ModelCount *int32 `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	OfflineDatasourceId *string `json:"OfflineDatasourceId,omitempty" xml:"OfflineDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	OfflineDatasourceName *string `json:"OfflineDatasourceName,omitempty" xml:"OfflineDatasourceName,omitempty"`
	// example:
	//
	// MaxCompute
	OfflineDatasourceType *string `json:"OfflineDatasourceType,omitempty" xml:"OfflineDatasourceType,omitempty"`
	// example:
	//
	// 90
	OfflineLifecycle *int32 `json:"OfflineLifecycle,omitempty" xml:"OfflineLifecycle,omitempty"`
	// example:
	//
	// 5
	OnlineDatasourceId *string `json:"OnlineDatasourceId,omitempty" xml:"OnlineDatasourceId,omitempty"`
	// example:
	//
	// datasource2
	OnlineDatasourceName *string `json:"OnlineDatasourceName,omitempty" xml:"OnlineDatasourceName,omitempty"`
	// example:
	//
	// Hologres
	OnlineDatasourceType *string `json:"OnlineDatasourceType,omitempty" xml:"OnlineDatasourceType,omitempty"`
	// example:
	//
	// 1232132543543****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetFeatureEntityCount(v int32) *GetProjectResponseBody {
	s.FeatureEntityCount = &v
	return s
}

func (s *GetProjectResponseBody) SetFeatureViewCount(v int32) *GetProjectResponseBody {
	s.FeatureViewCount = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtCreateTime(v string) *GetProjectResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtModifiedTime(v string) *GetProjectResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetProjectResponseBody) SetModelCount(v int32) *GetProjectResponseBody {
	s.ModelCount = &v
	return s
}

func (s *GetProjectResponseBody) SetName(v string) *GetProjectResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceId(v string) *GetProjectResponseBody {
	s.OfflineDatasourceId = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceName(v string) *GetProjectResponseBody {
	s.OfflineDatasourceName = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceType(v string) *GetProjectResponseBody {
	s.OfflineDatasourceType = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineLifecycle(v int32) *GetProjectResponseBody {
	s.OfflineLifecycle = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceId(v string) *GetProjectResponseBody {
	s.OnlineDatasourceId = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceName(v string) *GetProjectResponseBody {
	s.OnlineDatasourceName = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceType(v string) *GetProjectResponseBody {
	s.OnlineDatasourceType = &v
	return s
}

func (s *GetProjectResponseBody) SetOwner(v string) *GetProjectResponseBody {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
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

type GetProjectFeatureEntityResponseBody struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34245
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetProjectFeatureEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectFeatureEntityResponseBody) SetFeatureEntityId(v string) *GetProjectFeatureEntityResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetJoinId(v string) *GetProjectFeatureEntityResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetName(v string) *GetProjectFeatureEntityResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetProjectName(v string) *GetProjectFeatureEntityResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetRequestId(v string) *GetProjectFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetWorkspaceId(v string) *GetProjectFeatureEntityResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetProjectFeatureEntityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectFeatureEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectFeatureEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectFeatureEntityResponse) GoString() string {
	return s.String()
}

func (s *GetProjectFeatureEntityResponse) SetHeaders(v map[string]*string) *GetProjectFeatureEntityResponse {
	s.Headers = v
	return s
}

func (s *GetProjectFeatureEntityResponse) SetStatusCode(v int32) *GetProjectFeatureEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectFeatureEntityResponse) SetBody(v *GetProjectFeatureEntityResponseBody) *GetProjectFeatureEntityResponse {
	s.Body = v
	return s
}

type GetServiceIdentityRoleResponseBody struct {
	// example:
	//
	// {
	//
	// "Version": "1",
	//
	// "Statement":[]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 6F629E92-F64D-502D-85AA-A9C54894CA3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AliyunServiceRoleForPaiFeatureStore
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponseBody) SetPolicy(v string) *GetServiceIdentityRoleResponseBody {
	s.Policy = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRequestId(v string) *GetServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleName(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type GetServiceIdentityRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *GetServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetStatusCode(v int32) *GetServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetBody(v *GetServiceIdentityRoleResponseBody) *GetServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type GetTaskResponseBody struct {
	// example:
	//
	// {
	//
	// 	"mode": "overwrite",
	//
	// 	"partitions": {
	//
	// 		"dt": "20230820"
	//
	// 	},
	//
	// 	"event_time": "",
	//
	// 	"config": {},
	//
	// 	"offline_to_online": true
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtExecutedTime *string `json:"GmtExecutedTime,omitempty" xml:"GmtExecutedTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtFinishedTime *string `json:"GmtFinishedTime,omitempty" xml:"GmtFinishedTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 3
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// DROP TABLE IF EXISTS public.fsxxx
	RunningConfig *string `json:"RunningConfig,omitempty" xml:"RunningConfig,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetConfig(v string) *GetTaskResponseBody {
	s.Config = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtCreateTime(v string) *GetTaskResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtExecutedTime(v string) *GetTaskResponseBody {
	s.GmtExecutedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtFinishedTime(v string) *GetTaskResponseBody {
	s.GmtFinishedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtModifiedTime(v string) *GetTaskResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetObjectId(v string) *GetTaskResponseBody {
	s.ObjectId = &v
	return s
}

func (s *GetTaskResponseBody) SetObjectType(v string) *GetTaskResponseBody {
	s.ObjectType = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectId(v string) *GetTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectName(v string) *GetTaskResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetRunningConfig(v string) *GetTaskResponseBody {
	s.RunningConfig = &v
	return s
}

func (s *GetTaskResponseBody) SetStatus(v string) *GetTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBody) SetType(v string) *GetTaskResponseBody {
	s.Type = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type ListDatasourceTablesRequest struct {
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDatasourceTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesRequest) SetTableName(v string) *ListDatasourceTablesRequest {
	s.TableName = &v
	return s
}

type ListDatasourceTablesResponseBody struct {
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasourceTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesResponseBody) SetRequestId(v string) *ListDatasourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourceTablesResponseBody) SetTables(v []*string) *ListDatasourceTablesResponseBody {
	s.Tables = v
	return s
}

func (s *ListDatasourceTablesResponseBody) SetTotalCount(v int64) *ListDatasourceTablesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDatasourceTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasourceTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesResponse) SetHeaders(v map[string]*string) *ListDatasourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourceTablesResponse) SetStatusCode(v int32) *ListDatasourceTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasourceTablesResponse) SetBody(v *ListDatasourceTablesResponseBody) *ListDatasourceTablesResponse {
	s.Body = v
	return s
}

type ListDatasourcesRequest struct {
	// example:
	//
	// datasource1
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
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourcesRequest) SetName(v string) *ListDatasourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDatasourcesRequest) SetOrder(v string) *ListDatasourcesRequest {
	s.Order = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageNumber(v int32) *ListDatasourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageSize(v int32) *ListDatasourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasourcesRequest) SetSortBy(v string) *ListDatasourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasourcesRequest) SetType(v string) *ListDatasourcesRequest {
	s.Type = &v
	return s
}

func (s *ListDatasourcesRequest) SetWorkspaceId(v string) *ListDatasourcesRequest {
	s.WorkspaceId = &v
	return s
}

type ListDatasourcesResponseBody struct {
	Datasources []*ListDatasourcesResponseBodyDatasources `json:"Datasources,omitempty" xml:"Datasources,omitempty" type:"Repeated"`
	// example:
	//
	// 44933189-493B-5C43-A5C6-11EEC2A43520
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBody) SetDatasources(v []*ListDatasourcesResponseBodyDatasources) *ListDatasourcesResponseBody {
	s.Datasources = v
	return s
}

func (s *ListDatasourcesResponseBody) SetRequestId(v string) *ListDatasourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourcesResponseBody) SetTotalCount(v int64) *ListDatasourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDatasourcesResponseBodyDatasources struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 32324
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasourcesResponseBodyDatasources) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourcesResponseBodyDatasources) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBodyDatasources) SetConfig(v string) *ListDatasourcesResponseBodyDatasources {
	s.Config = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetDatasourceId(v string) *ListDatasourcesResponseBodyDatasources {
	s.DatasourceId = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetGmtCreateTime(v string) *ListDatasourcesResponseBodyDatasources {
	s.GmtCreateTime = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetGmtModifiedTime(v string) *ListDatasourcesResponseBodyDatasources {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetName(v string) *ListDatasourcesResponseBodyDatasources {
	s.Name = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetType(v string) *ListDatasourcesResponseBodyDatasources {
	s.Type = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetUri(v string) *ListDatasourcesResponseBodyDatasources {
	s.Uri = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetWorkspaceId(v string) *ListDatasourcesResponseBodyDatasources {
	s.WorkspaceId = &v
	return s
}

type ListDatasourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponse) SetHeaders(v map[string]*string) *ListDatasourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourcesResponse) SetStatusCode(v int32) *ListDatasourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasourcesResponse) SetBody(v *ListDatasourcesResponseBody) *ListDatasourcesResponse {
	s.Body = v
	return s
}

type ListFeatureEntitiesRequest struct {
	FeatureEntityIds []*string `json:"FeatureEntityIds,omitempty" xml:"FeatureEntityIds,omitempty" type:"Repeated"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListFeatureEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesRequest) SetFeatureEntityIds(v []*string) *ListFeatureEntitiesRequest {
	s.FeatureEntityIds = v
	return s
}

func (s *ListFeatureEntitiesRequest) SetName(v string) *ListFeatureEntitiesRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetOrder(v string) *ListFeatureEntitiesRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetOwner(v string) *ListFeatureEntitiesRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetPageNumber(v int32) *ListFeatureEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetPageSize(v int32) *ListFeatureEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetProjectId(v string) *ListFeatureEntitiesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetSortBy(v string) *ListFeatureEntitiesRequest {
	s.SortBy = &v
	return s
}

type ListFeatureEntitiesShrinkRequest struct {
	FeatureEntityIdsShrink *string `json:"FeatureEntityIds,omitempty" xml:"FeatureEntityIds,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListFeatureEntitiesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesShrinkRequest) SetFeatureEntityIdsShrink(v string) *ListFeatureEntitiesShrinkRequest {
	s.FeatureEntityIdsShrink = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetName(v string) *ListFeatureEntitiesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetOrder(v string) *ListFeatureEntitiesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetOwner(v string) *ListFeatureEntitiesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetPageNumber(v int32) *ListFeatureEntitiesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetPageSize(v int32) *ListFeatureEntitiesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetProjectId(v string) *ListFeatureEntitiesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetSortBy(v string) *ListFeatureEntitiesShrinkRequest {
	s.SortBy = &v
	return s
}

type ListFeatureEntitiesResponseBody struct {
	FeatureEntities []*ListFeatureEntitiesResponseBodyFeatureEntities `json:"FeatureEntities,omitempty" xml:"FeatureEntities,omitempty" type:"Repeated"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponseBody) SetFeatureEntities(v []*ListFeatureEntitiesResponseBodyFeatureEntities) *ListFeatureEntitiesResponseBody {
	s.FeatureEntities = v
	return s
}

func (s *ListFeatureEntitiesResponseBody) SetRequestId(v string) *ListFeatureEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBody) SetTotalCount(v int32) *ListFeatureEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFeatureEntitiesResponseBodyFeatureEntities struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListFeatureEntitiesResponseBodyFeatureEntities) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureEntitiesResponseBodyFeatureEntities) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetFeatureEntityId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.FeatureEntityId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetGmtCreateTime(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetJoinId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.JoinId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetName(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetOwner(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetProjectId(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesResponseBodyFeatureEntities) SetProjectName(v string) *ListFeatureEntitiesResponseBodyFeatureEntities {
	s.ProjectName = &v
	return s
}

type ListFeatureEntitiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponse) SetHeaders(v map[string]*string) *ListFeatureEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureEntitiesResponse) SetStatusCode(v int32) *ListFeatureEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureEntitiesResponse) SetBody(v *ListFeatureEntitiesResponseBody) *ListFeatureEntitiesResponse {
	s.Body = v
	return s
}

type ListFeatureViewFieldRelationshipsResponseBody struct {
	Relationships []*ListFeatureViewFieldRelationshipsResponseBodyRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) SetRelationships(v []*ListFeatureViewFieldRelationshipsResponseBodyRelationships) *ListFeatureViewFieldRelationshipsResponseBody {
	s.Relationships = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) SetRequestId(v string) *ListFeatureViewFieldRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

type ListFeatureViewFieldRelationshipsResponseBodyRelationships struct {
	// example:
	//
	// featureView1
	FeatureName *string                                                             `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	Models      []*ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// table2
	OfflineTableName *string `json:"OfflineTableName,omitempty" xml:"OfflineTableName,omitempty"`
	// example:
	//
	// table1
	OnlineTableName *string `json:"OnlineTableName,omitempty" xml:"OnlineTableName,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationships) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationships) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetFeatureName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetModels(v []*ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.Models = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetOfflineTableName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.OfflineTableName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetOnlineTableName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.OnlineTableName = &v
	return s
}

type ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels struct {
	// example:
	//
	// 3
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// dbmtl
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) SetModelId(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	s.ModelId = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) SetModelName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	s.ModelName = &v
	return s
}

type ListFeatureViewFieldRelationshipsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewFieldRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetHeaders(v map[string]*string) *ListFeatureViewFieldRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetStatusCode(v int32) *ListFeatureViewFieldRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponse) SetBody(v *ListFeatureViewFieldRelationshipsResponseBody) *ListFeatureViewFieldRelationshipsResponse {
	s.Body = v
	return s
}

type ListFeatureViewOnlineFeaturesRequest struct {
	// This parameter is required.
	JoinIds []*string `json:"JoinIds,omitempty" xml:"JoinIds,omitempty" type:"Repeated"`
}

func (s ListFeatureViewOnlineFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesRequest) SetJoinIds(v []*string) *ListFeatureViewOnlineFeaturesRequest {
	s.JoinIds = v
	return s
}

type ListFeatureViewOnlineFeaturesShrinkRequest struct {
	// This parameter is required.
	JoinIdsShrink *string `json:"JoinIds,omitempty" xml:"JoinIds,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesShrinkRequest) SetJoinIdsShrink(v string) *ListFeatureViewOnlineFeaturesShrinkRequest {
	s.JoinIdsShrink = &v
	return s
}

type ListFeatureViewOnlineFeaturesResponseBody struct {
	OnlineFeatures []*string `json:"OnlineFeatures,omitempty" xml:"OnlineFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) SetOnlineFeatures(v []*string) *ListFeatureViewOnlineFeaturesResponseBody {
	s.OnlineFeatures = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) SetRequestId(v string) *ListFeatureViewOnlineFeaturesResponseBody {
	s.RequestId = &v
	return s
}

type ListFeatureViewOnlineFeaturesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewOnlineFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetHeaders(v map[string]*string) *ListFeatureViewOnlineFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetStatusCode(v int32) *ListFeatureViewOnlineFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetBody(v *ListFeatureViewOnlineFeaturesResponseBody) *ListFeatureViewOnlineFeaturesResponse {
	s.Body = v
	return s
}

type ListFeatureViewRelationshipsResponseBody struct {
	Relationships []*ListFeatureViewRelationshipsResponseBodyRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
	// example:
	//
	// 0FBBE454-9BD1-5D8F-9129-D14DB7FAFE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBody) SetRelationships(v []*ListFeatureViewRelationshipsResponseBodyRelationships) *ListFeatureViewRelationshipsResponseBody {
	s.Relationships = v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBody) SetRequestId(v string) *ListFeatureViewRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

type ListFeatureViewRelationshipsResponseBodyRelationships struct {
	// example:
	//
	// fv1
	FeatureViewName *string                                                        `json:"FeatureViewName,omitempty" xml:"FeatureViewName,omitempty"`
	Models          []*ListFeatureViewRelationshipsResponseBodyRelationshipsModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBodyRelationships) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBodyRelationships) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetFeatureViewName(v string) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.FeatureViewName = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetModels(v []*ListFeatureViewRelationshipsResponseBodyRelationshipsModels) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.Models = v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetProjectName(v string) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.ProjectName = &v
	return s
}

type ListFeatureViewRelationshipsResponseBodyRelationshipsModels struct {
	// example:
	//
	// 3
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// dbmtl
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBodyRelationshipsModels) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBodyRelationshipsModels) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) SetModelId(v string) *ListFeatureViewRelationshipsResponseBodyRelationshipsModels {
	s.ModelId = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) SetModelName(v string) *ListFeatureViewRelationshipsResponseBodyRelationshipsModels {
	s.ModelName = &v
	return s
}

type ListFeatureViewRelationshipsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewRelationshipsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponse) SetHeaders(v map[string]*string) *ListFeatureViewRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewRelationshipsResponse) SetStatusCode(v int32) *ListFeatureViewRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponse) SetBody(v *ListFeatureViewRelationshipsResponseBody) *ListFeatureViewRelationshipsResponse {
	s.Body = v
	return s
}

type ListFeatureViewsRequest struct {
	// example:
	//
	// feature1
	FeatureName    *string   `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureViewIds []*string `json:"FeatureViewIds,omitempty" xml:"FeatureViewIds,omitempty" type:"Repeated"`
	// example:
	//
	// fv1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1232143243242****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFeatureViewsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsRequest) SetFeatureName(v string) *ListFeatureViewsRequest {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewsRequest) SetFeatureViewIds(v []*string) *ListFeatureViewsRequest {
	s.FeatureViewIds = v
	return s
}

func (s *ListFeatureViewsRequest) SetName(v string) *ListFeatureViewsRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsRequest) SetOrder(v string) *ListFeatureViewsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureViewsRequest) SetOwner(v string) *ListFeatureViewsRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsRequest) SetPageNumber(v int32) *ListFeatureViewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureViewsRequest) SetPageSize(v int32) *ListFeatureViewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureViewsRequest) SetProjectId(v string) *ListFeatureViewsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsRequest) SetSortBy(v string) *ListFeatureViewsRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureViewsRequest) SetTag(v string) *ListFeatureViewsRequest {
	s.Tag = &v
	return s
}

func (s *ListFeatureViewsRequest) SetType(v string) *ListFeatureViewsRequest {
	s.Type = &v
	return s
}

type ListFeatureViewsShrinkRequest struct {
	// example:
	//
	// feature1
	FeatureName          *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureViewIdsShrink *string `json:"FeatureViewIds,omitempty" xml:"FeatureViewIds,omitempty"`
	// example:
	//
	// fv1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1232143243242****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFeatureViewsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsShrinkRequest) SetFeatureName(v string) *ListFeatureViewsShrinkRequest {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetFeatureViewIdsShrink(v string) *ListFeatureViewsShrinkRequest {
	s.FeatureViewIdsShrink = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetName(v string) *ListFeatureViewsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetOrder(v string) *ListFeatureViewsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetOwner(v string) *ListFeatureViewsShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetPageNumber(v int32) *ListFeatureViewsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetPageSize(v int32) *ListFeatureViewsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetProjectId(v string) *ListFeatureViewsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetSortBy(v string) *ListFeatureViewsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetTag(v string) *ListFeatureViewsShrinkRequest {
	s.Tag = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetType(v string) *ListFeatureViewsShrinkRequest {
	s.Type = &v
	return s
}

type ListFeatureViewsResponseBody struct {
	FeatureViews []*ListFeatureViewsResponseBodyFeatureViews `json:"FeatureViews,omitempty" xml:"FeatureViews,omitempty" type:"Repeated"`
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureViewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponseBody) SetFeatureViews(v []*ListFeatureViewsResponseBodyFeatureViews) *ListFeatureViewsResponseBody {
	s.FeatureViews = v
	return s
}

func (s *ListFeatureViewsResponseBody) SetRequestId(v string) *ListFeatureViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureViewsResponseBody) SetTotalCount(v int64) *ListFeatureViewsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFeatureViewsResponseBodyFeatureViews struct {
	// example:
	//
	// featureEntity1
	FeatureEntityName *string `json:"FeatureEntityName,omitempty" xml:"FeatureEntityName,omitempty"`
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// featureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321421412****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 4
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	RegisterDatasourceName *string `json:"RegisterDatasourceName,omitempty" xml:"RegisterDatasourceName,omitempty"`
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// example:
	//
	// 90
	TTL *int32 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// example:
	//
	// Batch
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WriteToFeatureDB *bool   `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s ListFeatureViewsResponseBodyFeatureViews) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewsResponseBodyFeatureViews) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetFeatureEntityName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.FeatureEntityName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetFeatureViewId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.FeatureViewId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetGmtCreateTime(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetGmtModifiedTime(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetOwner(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetProjectId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetProjectName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.ProjectName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterDatasourceId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterDatasourceId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterDatasourceName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterDatasourceName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterTable(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterTable = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetTTL(v int32) *ListFeatureViewsResponseBodyFeatureViews {
	s.TTL = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetType(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Type = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetWriteToFeatureDB(v bool) *ListFeatureViewsResponseBodyFeatureViews {
	s.WriteToFeatureDB = &v
	return s
}

type ListFeatureViewsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeatureViewsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponse) SetHeaders(v map[string]*string) *ListFeatureViewsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewsResponse) SetStatusCode(v int32) *ListFeatureViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewsResponse) SetBody(v *ListFeatureViewsResponseBody) *ListFeatureViewsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 2CA11923-2A3D-5E5A-8314-E699D2DD15DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	FeatureDBInstanceInfo *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo `json:"FeatureDBInstanceInfo,omitempty" xml:"FeatureDBInstanceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// featureStore-cn-7mz2xfu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Initializing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetFeatureDBInstanceInfo(v *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) *ListInstancesResponseBodyInstances {
	s.FeatureDBInstanceInfo = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyInstancesFeatureDBInstanceInfo struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo) SetStatus(v string) *ListInstancesResponseBodyInstancesFeatureDBInstanceInfo {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListLabelTablesRequest struct {
	LabelTableIds []*string `json:"LabelTableIds,omitempty" xml:"LabelTableIds,omitempty" type:"Repeated"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432432****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// project1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLabelTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLabelTablesRequest) GoString() string {
	return s.String()
}

func (s *ListLabelTablesRequest) SetLabelTableIds(v []*string) *ListLabelTablesRequest {
	s.LabelTableIds = v
	return s
}

func (s *ListLabelTablesRequest) SetName(v string) *ListLabelTablesRequest {
	s.Name = &v
	return s
}

func (s *ListLabelTablesRequest) SetOrder(v string) *ListLabelTablesRequest {
	s.Order = &v
	return s
}

func (s *ListLabelTablesRequest) SetOwner(v string) *ListLabelTablesRequest {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesRequest) SetPageNumber(v int64) *ListLabelTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLabelTablesRequest) SetPageSize(v int64) *ListLabelTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLabelTablesRequest) SetProjectId(v string) *ListLabelTablesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesRequest) SetSortBy(v string) *ListLabelTablesRequest {
	s.SortBy = &v
	return s
}

type ListLabelTablesShrinkRequest struct {
	LabelTableIdsShrink *string `json:"LabelTableIds,omitempty" xml:"LabelTableIds,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432432****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// project1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLabelTablesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLabelTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListLabelTablesShrinkRequest) SetLabelTableIdsShrink(v string) *ListLabelTablesShrinkRequest {
	s.LabelTableIdsShrink = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetName(v string) *ListLabelTablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetOrder(v string) *ListLabelTablesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetOwner(v string) *ListLabelTablesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetPageNumber(v int64) *ListLabelTablesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetPageSize(v int64) *ListLabelTablesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetProjectId(v string) *ListLabelTablesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetSortBy(v string) *ListLabelTablesShrinkRequest {
	s.SortBy = &v
	return s
}

type ListLabelTablesResponseBody struct {
	LabelTables []*ListLabelTablesResponseBodyLabelTables `json:"LabelTables,omitempty" xml:"LabelTables,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLabelTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLabelTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponseBody) SetLabelTables(v []*ListLabelTablesResponseBodyLabelTables) *ListLabelTablesResponseBody {
	s.LabelTables = v
	return s
}

func (s *ListLabelTablesResponseBody) SetRequestId(v string) *ListLabelTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLabelTablesResponseBody) SetTotalCount(v int64) *ListLabelTablesResponseBody {
	s.TotalCount = &v
	return s
}

type ListLabelTablesResponseBodyLabelTables struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// datasource1
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 3
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123214213214
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListLabelTablesResponseBodyLabelTables) String() string {
	return tea.Prettify(s)
}

func (s ListLabelTablesResponseBodyLabelTables) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponseBodyLabelTables) SetDatasourceId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.DatasourceId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetDatasourceName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.DatasourceName = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetGmtCreateTime(v string) *ListLabelTablesResponseBodyLabelTables {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetGmtModifiedTime(v string) *ListLabelTablesResponseBodyLabelTables {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetLabelTableId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.LabelTableId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.Name = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetOwner(v string) *ListLabelTablesResponseBodyLabelTables {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetProjectId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetProjectName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.ProjectName = &v
	return s
}

type ListLabelTablesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLabelTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLabelTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLabelTablesResponse) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponse) SetHeaders(v map[string]*string) *ListLabelTablesResponse {
	s.Headers = v
	return s
}

func (s *ListLabelTablesResponse) SetStatusCode(v int32) *ListLabelTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelTablesResponse) SetBody(v *ListLabelTablesResponseBody) *ListLabelTablesResponse {
	s.Body = v
	return s
}

type ListModelFeatureAvailableFeaturesRequest struct {
	// example:
	//
	// f1
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesRequest) SetFeatureName(v string) *ListModelFeatureAvailableFeaturesRequest {
	s.FeatureName = &v
	return s
}

type ListModelFeatureAvailableFeaturesResponseBody struct {
	AvaliableFeatures []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures `json:"AvaliableFeatures,omitempty" xml:"AvaliableFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// ED4DEA2F-F216-57F0-AE28-08D791233280
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetAvaliableFeatures(v []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) *ListModelFeatureAvailableFeaturesResponseBody {
	s.AvaliableFeatures = v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetTotalCount(v int64) *ListModelFeatureAvailableFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetRequestId(v string) *ListModelFeatureAvailableFeaturesResponseBody {
	s.RequestId = &v
	return s
}

type ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures struct {
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// user_fea
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// example:
	//
	// FeatureView
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetName(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.Name = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetSourceName(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.SourceName = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetSourceType(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.SourceType = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetType(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.Type = &v
	return s
}

type ListModelFeatureAvailableFeaturesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelFeatureAvailableFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetHeaders(v map[string]*string) *ListModelFeatureAvailableFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetStatusCode(v int32) *ListModelFeatureAvailableFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetBody(v *ListModelFeatureAvailableFeaturesResponseBody) *ListModelFeatureAvailableFeaturesResponse {
	s.Body = v
	return s
}

type ListModelFeaturesRequest struct {
	ModelFeatureIds []*string `json:"ModelFeatureIds,omitempty" xml:"ModelFeatureIds,omitempty" type:"Repeated"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 12323143****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// DESC
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListModelFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesRequest) SetModelFeatureIds(v []*string) *ListModelFeaturesRequest {
	s.ModelFeatureIds = v
	return s
}

func (s *ListModelFeaturesRequest) SetName(v string) *ListModelFeaturesRequest {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesRequest) SetOrder(v string) *ListModelFeaturesRequest {
	s.Order = &v
	return s
}

func (s *ListModelFeaturesRequest) SetOwner(v string) *ListModelFeaturesRequest {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesRequest) SetPageNumber(v int32) *ListModelFeaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelFeaturesRequest) SetPageSize(v int32) *ListModelFeaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelFeaturesRequest) SetProjectId(v string) *ListModelFeaturesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesRequest) SetSortBy(v string) *ListModelFeaturesRequest {
	s.SortBy = &v
	return s
}

type ListModelFeaturesShrinkRequest struct {
	ModelFeatureIdsShrink *string `json:"ModelFeatureIds,omitempty" xml:"ModelFeatureIds,omitempty"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 12323143****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// DESC
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListModelFeaturesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeaturesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesShrinkRequest) SetModelFeatureIdsShrink(v string) *ListModelFeaturesShrinkRequest {
	s.ModelFeatureIdsShrink = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetName(v string) *ListModelFeaturesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetOrder(v string) *ListModelFeaturesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetOwner(v string) *ListModelFeaturesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetPageNumber(v int32) *ListModelFeaturesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetPageSize(v int32) *ListModelFeaturesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetProjectId(v string) *ListModelFeaturesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetSortBy(v string) *ListModelFeaturesShrinkRequest {
	s.SortBy = &v
	return s
}

type ListModelFeaturesResponseBody struct {
	ModelFeatures []*ListModelFeaturesResponseBodyModelFeatures `json:"ModelFeatures,omitempty" xml:"ModelFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 2CA11923-2A3D-5E5A-8314-E699D2DD15DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponseBody) SetModelFeatures(v []*ListModelFeaturesResponseBodyModelFeatures) *ListModelFeaturesResponseBody {
	s.ModelFeatures = v
	return s
}

func (s *ListModelFeaturesResponseBody) SetRequestId(v string) *ListModelFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelFeaturesResponseBody) SetTotalCount(v int64) *ListModelFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

type ListModelFeaturesResponseBodyModelFeatures struct {
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// label_table_1
	LabelTableName *string `json:"LabelTableName,omitempty" xml:"LabelTableName,omitempty"`
	// example:
	//
	// 3
	ModelFeatureId *string `json:"ModelFeatureId,omitempty" xml:"ModelFeatureId,omitempty"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1231243253****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 5
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListModelFeaturesResponseBodyModelFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeaturesResponseBodyModelFeatures) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetGmtCreateTime(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.GmtCreateTime = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetGmtModifiedTime(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetLabelTableName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.LabelTableName = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetModelFeatureId(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ModelFeatureId = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetOwner(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetProjectId(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetProjectName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ProjectName = &v
	return s
}

type ListModelFeaturesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponse) SetHeaders(v map[string]*string) *ListModelFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListModelFeaturesResponse) SetStatusCode(v int32) *ListModelFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelFeaturesResponse) SetBody(v *ListModelFeaturesResponseBody) *ListModelFeaturesResponse {
	s.Body = v
	return s
}

type ListProjectFeatureViewsResponseBody struct {
	FeatureViews []*ListProjectFeatureViewsResponseBodyFeatureViews `json:"FeatureViews,omitempty" xml:"FeatureViews,omitempty" type:"Repeated"`
	// example:
	//
	// AE2AF33E-0C0D-51A8-B89B-C5F1DF261D92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectFeatureViewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBody) SetFeatureViews(v []*ListProjectFeatureViewsResponseBodyFeatureViews) *ListProjectFeatureViewsResponseBody {
	s.FeatureViews = v
	return s
}

func (s *ListProjectFeatureViewsResponseBody) SetRequestId(v string) *ListProjectFeatureViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBody) SetTotalCount(v int64) *ListProjectFeatureViewsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectFeatureViewsResponseBodyFeatureViews struct {
	// example:
	//
	// 3
	FeatureViewId *string                                                    `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	Features      []*ListProjectFeatureViewsResponseBodyFeatureViewsFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// feature_view1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectFeatureViewsResponseBodyFeatureViews) String() string {
	return tea.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBodyFeatureViews) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetFeatureViewId(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.FeatureViewId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetFeatures(v []*ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Features = v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetName(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Name = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetType(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Type = &v
	return s
}

type ListProjectFeatureViewsResponseBodyFeatureViewsFeatures struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetAttributes(v []*string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Attributes = v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetName(v string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Name = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetType(v string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Type = &v
	return s
}

type ListProjectFeatureViewsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectFeatureViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectFeatureViewsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectFeatureViewsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponse) SetHeaders(v map[string]*string) *ListProjectFeatureViewsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectFeatureViewsResponse) SetStatusCode(v int32) *ListProjectFeatureViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectFeatureViewsResponse) SetBody(v *ListProjectFeatureViewsResponseBody) *ListProjectFeatureViewsResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// example:
	//
	// fs1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 134324352****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetName(v string) *ListProjectsRequest {
	s.Name = &v
	return s
}

func (s *ListProjectsRequest) SetOrder(v string) *ListProjectsRequest {
	s.Order = &v
	return s
}

func (s *ListProjectsRequest) SetOwner(v string) *ListProjectsRequest {
	s.Owner = &v
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

func (s *ListProjectsRequest) SetProjectIds(v []*string) *ListProjectsRequest {
	s.ProjectIds = v
	return s
}

func (s *ListProjectsRequest) SetSortBy(v string) *ListProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsShrinkRequest struct {
	// example:
	//
	// fs1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 134324352****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectIdsShrink *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) SetName(v string) *ListProjectsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOrder(v string) *ListProjectsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOwner(v string) *ListProjectsShrinkRequest {
	s.Owner = &v
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

func (s *ListProjectsShrinkRequest) SetProjectIdsShrink(v string) *ListProjectsShrinkRequest {
	s.ProjectIdsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetSortBy(v string) *ListProjectsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetWorkspaceId(v string) *ListProjectsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponseBody struct {
	Projects []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// example:
	//
	// 44933189-493B-5C43-A5C6-11EEC2A43520
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10
	FeatureEntityCount *int32 `json:"FeatureEntityCount,omitempty" xml:"FeatureEntityCount,omitempty"`
	// example:
	//
	// 10
	FeatureViewCount *int32 `json:"FeatureViewCount,omitempty" xml:"FeatureViewCount,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 5
	ModelCount *int32 `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	OfflineDatasourceId *string `json:"OfflineDatasourceId,omitempty" xml:"OfflineDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	OfflineDatasourceName *string `json:"OfflineDatasourceName,omitempty" xml:"OfflineDatasourceName,omitempty"`
	// example:
	//
	// MaxCompute
	OfflineDatasourceType *string `json:"OfflineDatasourceType,omitempty" xml:"OfflineDatasourceType,omitempty"`
	// example:
	//
	// 10
	OfflineLifecycle *int32 `json:"OfflineLifecycle,omitempty" xml:"OfflineLifecycle,omitempty"`
	// example:
	//
	// 5
	OnlineDatasourceId *string `json:"OnlineDatasourceId,omitempty" xml:"OnlineDatasourceId,omitempty"`
	// example:
	//
	// datasource2
	OnlineDatasourceName *string `json:"OnlineDatasourceName,omitempty" xml:"OnlineDatasourceName,omitempty"`
	// example:
	//
	// Hologres
	OnlineDatasourceType *string `json:"OnlineDatasourceType,omitempty" xml:"OnlineDatasourceType,omitempty"`
	// example:
	//
	// 1232132543543****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetFeatureEntityCount(v int32) *ListProjectsResponseBodyProjects {
	s.FeatureEntityCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetFeatureViewCount(v int32) *ListProjectsResponseBodyProjects {
	s.FeatureViewCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreateTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtCreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtModifiedTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModelCount(v int32) *ListProjectsResponseBodyProjects {
	s.ModelCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceId(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceName(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceType(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineLifecycle(v int32) *ListProjectsResponseBodyProjects {
	s.OfflineLifecycle = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceId(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceName(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceType(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOwner(v string) *ListProjectsResponseBodyProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectId(v string) *ListProjectsResponseBodyProjects {
	s.ProjectId = &v
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

type ListTaskLogsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTaskLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskLogsRequest) SetPageNumber(v int32) *ListTaskLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskLogsRequest) SetPageSize(v int32) *ListTaskLogsRequest {
	s.PageSize = &v
	return s
}

type ListTaskLogsResponseBody struct {
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 72F15A8A-5A28-5B18-A0DE-0EABD7D3245A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskLogsResponseBody) SetLogs(v []*string) *ListTaskLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListTaskLogsResponseBody) SetRequestId(v string) *ListTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskLogsResponseBody) SetTotalCount(v int32) *ListTaskLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTaskLogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskLogsResponse) SetHeaders(v map[string]*string) *ListTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskLogsResponse) SetStatusCode(v int32) *ListTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskLogsResponse) SetBody(v *ListTaskLogsResponseBody) *ListTaskLogsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// example:
	//
	// 4
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
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
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Running
	Status  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetObjectId(v string) *ListTasksRequest {
	s.ObjectId = &v
	return s
}

func (s *ListTasksRequest) SetObjectType(v string) *ListTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v string) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTaskIds(v []*string) *ListTasksRequest {
	s.TaskIds = v
	return s
}

func (s *ListTasksRequest) SetType(v string) *ListTasksRequest {
	s.Type = &v
	return s
}

type ListTasksShrinkRequest struct {
	// example:
	//
	// 4
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
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
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Running
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) SetObjectId(v string) *ListTasksShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetObjectType(v string) *ListTasksShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageNumber(v int32) *ListTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageSize(v int32) *ListTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectId(v string) *ListTasksShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskIdsShrink(v string) *ListTasksShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetType(v string) *ListTasksShrinkRequest {
	s.Type = &v
	return s
}

type ListTasksResponseBody struct {
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtExecutedTime *string `json:"GmtExecutedTime,omitempty" xml:"GmtExecutedTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtFinishedTime *string `json:"GmtFinishedTime,omitempty" xml:"GmtFinishedTime,omitempty"`
	// example:
	//
	// 5
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 4
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetGmtCreateTime(v string) *ListTasksResponseBodyTasks {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGmtExecutedTime(v string) *ListTasksResponseBodyTasks {
	s.GmtExecutedTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGmtFinishedTime(v string) *ListTasksResponseBodyTasks {
	s.GmtFinishedTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetObjectId(v string) *ListTasksResponseBodyTasks {
	s.ObjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetObjectType(v string) *ListTasksResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetProjectId(v string) *ListTasksResponseBodyTasks {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetProjectName(v string) *ListTasksResponseBodyTasks {
	s.ProjectName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetType(v string) *ListTasksResponseBodyTasks {
	s.Type = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type PublishFeatureViewTableRequest struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Overwrite
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OfflineToOnline *bool                             `json:"OfflineToOnline,omitempty" xml:"OfflineToOnline,omitempty"`
	Partitions      map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s PublishFeatureViewTableRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishFeatureViewTableRequest) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableRequest) SetConfig(v string) *PublishFeatureViewTableRequest {
	s.Config = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetEventTime(v string) *PublishFeatureViewTableRequest {
	s.EventTime = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetMode(v string) *PublishFeatureViewTableRequest {
	s.Mode = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetOfflineToOnline(v bool) *PublishFeatureViewTableRequest {
	s.OfflineToOnline = &v
	return s
}

func (s *PublishFeatureViewTableRequest) SetPartitions(v map[string]map[string]interface{}) *PublishFeatureViewTableRequest {
	s.Partitions = v
	return s
}

type PublishFeatureViewTableResponseBody struct {
	// example:
	//
	// 627B5776-4D06-5A49-8A04-508AA39653F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PublishFeatureViewTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishFeatureViewTableResponseBody) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableResponseBody) SetRequestId(v string) *PublishFeatureViewTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishFeatureViewTableResponseBody) SetTaskId(v string) *PublishFeatureViewTableResponseBody {
	s.TaskId = &v
	return s
}

type PublishFeatureViewTableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishFeatureViewTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFeatureViewTableResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishFeatureViewTableResponse) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableResponse) SetHeaders(v map[string]*string) *PublishFeatureViewTableResponse {
	s.Headers = v
	return s
}

func (s *PublishFeatureViewTableResponse) SetStatusCode(v int32) *PublishFeatureViewTableResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFeatureViewTableResponse) SetBody(v *PublishFeatureViewTableResponseBody) *PublishFeatureViewTableResponse {
	s.Body = v
	return s
}

type UpdateDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateDatasourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceRequest) SetConfig(v string) *UpdateDatasourceRequest {
	s.Config = &v
	return s
}

func (s *UpdateDatasourceRequest) SetName(v string) *UpdateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasourceRequest) SetUri(v string) *UpdateDatasourceRequest {
	s.Uri = &v
	return s
}

type UpdateDatasourceResponseBody struct {
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceResponseBody) SetRequestId(v string) *UpdateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceResponse) SetHeaders(v map[string]*string) *UpdateDatasourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasourceResponse) SetStatusCode(v int32) *UpdateDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasourceResponse) SetBody(v *UpdateDatasourceResponseBody) *UpdateDatasourceResponse {
	s.Body = v
	return s
}

type UpdateLabelTableRequest struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// This parameter is required.
	Fields []*UpdateLabelTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// rec_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLabelTableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableRequest) SetDatasourceId(v string) *UpdateLabelTableRequest {
	s.DatasourceId = &v
	return s
}

func (s *UpdateLabelTableRequest) SetFields(v []*UpdateLabelTableRequestFields) *UpdateLabelTableRequest {
	s.Fields = v
	return s
}

func (s *UpdateLabelTableRequest) SetName(v string) *UpdateLabelTableRequest {
	s.Name = &v
	return s
}

type UpdateLabelTableRequestFields struct {
	// This parameter is required.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// lat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DOUBLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateLabelTableRequestFields) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelTableRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableRequestFields) SetAttributes(v []*string) *UpdateLabelTableRequestFields {
	s.Attributes = v
	return s
}

func (s *UpdateLabelTableRequestFields) SetName(v string) *UpdateLabelTableRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateLabelTableRequestFields) SetType(v string) *UpdateLabelTableRequestFields {
	s.Type = &v
	return s
}

type UpdateLabelTableResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLabelTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableResponseBody) SetRequestId(v string) *UpdateLabelTableResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLabelTableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableResponse) SetHeaders(v map[string]*string) *UpdateLabelTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateLabelTableResponse) SetStatusCode(v int32) *UpdateLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLabelTableResponse) SetBody(v *UpdateLabelTableResponseBody) *UpdateLabelTableResponse {
	s.Body = v
	return s
}

type UpdateModelFeatureRequest struct {
	Features []*UpdateModelFeatureRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// example:
	//
	// 4
	LabelTableId           *string   `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	SequenceFeatureViewIds []*string `json:"SequenceFeatureViewIds,omitempty" xml:"SequenceFeatureViewIds,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureRequest) SetFeatures(v []*UpdateModelFeatureRequestFeatures) *UpdateModelFeatureRequest {
	s.Features = v
	return s
}

func (s *UpdateModelFeatureRequest) SetLabelPriorityLevel(v int64) *UpdateModelFeatureRequest {
	s.LabelPriorityLevel = &v
	return s
}

func (s *UpdateModelFeatureRequest) SetLabelTableId(v string) *UpdateModelFeatureRequest {
	s.LabelTableId = &v
	return s
}

func (s *UpdateModelFeatureRequest) SetSequenceFeatureViewIds(v []*string) *UpdateModelFeatureRequest {
	s.SequenceFeatureViewIds = v
	return s
}

type UpdateModelFeatureRequestFeatures struct {
	// example:
	//
	// sex
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateModelFeatureRequestFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureRequestFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureRequestFeatures) SetAliasName(v string) *UpdateModelFeatureRequestFeatures {
	s.AliasName = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetFeatureViewId(v string) *UpdateModelFeatureRequestFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetName(v string) *UpdateModelFeatureRequestFeatures {
	s.Name = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetType(v string) *UpdateModelFeatureRequestFeatures {
	s.Type = &v
	return s
}

type UpdateModelFeatureResponseBody struct {
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureResponseBody) SetRequestId(v string) *UpdateModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureResponse) SetHeaders(v map[string]*string) *UpdateModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelFeatureResponse) SetStatusCode(v int32) *UpdateModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelFeatureResponse) SetBody(v *UpdateModelFeatureResponseBody) *UpdateModelFeatureResponse {
	s.Body = v
	return s
}

type UpdateModelFeatureFGFeatureRequest struct {
	LookupFeatures   []*UpdateModelFeatureFGFeatureRequestLookupFeatures   `json:"LookupFeatures,omitempty" xml:"LookupFeatures,omitempty" type:"Repeated"`
	RawFeatures      []*UpdateModelFeatureFGFeatureRequestRawFeatures      `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty" type:"Repeated"`
	Reserves         []*string                                             `json:"Reserves,omitempty" xml:"Reserves,omitempty" type:"Repeated"`
	SequenceFeatures []*UpdateModelFeatureFGFeatureRequestSequenceFeatures `json:"SequenceFeatures,omitempty" xml:"SequenceFeatures,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureFGFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequest) SetLookupFeatures(v []*UpdateModelFeatureFGFeatureRequestLookupFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.LookupFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetRawFeatures(v []*UpdateModelFeatureFGFeatureRequestRawFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.RawFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetReserves(v []*string) *UpdateModelFeatureFGFeatureRequest {
	s.Reserves = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetSequenceFeatures(v []*UpdateModelFeatureFGFeatureRequestSequenceFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.SequenceFeatures = v
	return s
}

type UpdateModelFeatureFGFeatureRequestLookupFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Item
	KeyFeatureDomain *string `json:"KeyFeatureDomain,omitempty" xml:"KeyFeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	KeyFeatureName *string `json:"KeyFeatureName,omitempty" xml:"KeyFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	MapFeatureDomain *string `json:"MapFeatureDomain,omitempty" xml:"MapFeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	MapFeatureName *string `json:"MapFeatureName,omitempty" xml:"MapFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestLookupFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestLookupFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetKeyFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.KeyFeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetKeyFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.KeyFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetMapFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.MapFeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetMapFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.MapFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.ValueType = &v
	return s
}

type UpdateModelFeatureFGFeatureRequestRawFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestRawFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestRawFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureType(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetInputFeatureName(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.ValueType = &v
	return s
}

type UpdateModelFeatureFGFeatureRequestSequenceFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// #
	AttributeDelim *string `json:"AttributeDelim,omitempty" xml:"AttributeDelim,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ;
	SequenceDelim *string `json:"SequenceDelim,omitempty" xml:"SequenceDelim,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	SequenceLength *int64                                                           `json:"SequenceLength,omitempty" xml:"SequenceLength,omitempty"`
	SubFeatures    []*UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures `json:"SubFeatures,omitempty" xml:"SubFeatures,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetAttributeDelim(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.AttributeDelim = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSequenceDelim(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SequenceDelim = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSequenceLength(v int64) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SequenceLength = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSubFeatures(v []*UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SubFeatures = v
	return s
}

type UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RawFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureType(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetInputFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.ValueType = &v
	return s
}

type UpdateModelFeatureFGFeatureResponseBody struct {
	// example:
	//
	// 7D497816-607C-5B67-97B1-61354B6ACB2B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateModelFeatureFGFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureResponseBody) SetRequestId(v string) *UpdateModelFeatureFGFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModelFeatureFGFeatureResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelFeatureFGFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelFeatureFGFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureResponse) SetHeaders(v map[string]*string) *UpdateModelFeatureFGFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponse) SetStatusCode(v int32) *UpdateModelFeatureFGFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponse) SetBody(v *UpdateModelFeatureFGFeatureResponseBody) *UpdateModelFeatureFGFeatureResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// 2150233F-A1F7-54D2-B5B5-8A70567549BD
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

type WriteFeatureViewTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Merge
	Mode          *string                                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Partitions    map[string]map[string]interface{}          `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	UrlDatasource *WriteFeatureViewTableRequestUrlDatasource `json:"UrlDatasource,omitempty" xml:"UrlDatasource,omitempty" type:"Struct"`
}

func (s WriteFeatureViewTableRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteFeatureViewTableRequest) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableRequest) SetMode(v string) *WriteFeatureViewTableRequest {
	s.Mode = &v
	return s
}

func (s *WriteFeatureViewTableRequest) SetPartitions(v map[string]map[string]interface{}) *WriteFeatureViewTableRequest {
	s.Partitions = v
	return s
}

func (s *WriteFeatureViewTableRequest) SetUrlDatasource(v *WriteFeatureViewTableRequestUrlDatasource) *WriteFeatureViewTableRequest {
	s.UrlDatasource = v
	return s
}

type WriteFeatureViewTableRequestUrlDatasource struct {
	// example:
	//
	// ,
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// example:
	//
	// true
	OmitHeader *bool `json:"OmitHeader,omitempty" xml:"OmitHeader,omitempty"`
	// example:
	//
	// xxx.xxx.com/file.csv
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s WriteFeatureViewTableRequestUrlDatasource) String() string {
	return tea.Prettify(s)
}

func (s WriteFeatureViewTableRequestUrlDatasource) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetDelimiter(v string) *WriteFeatureViewTableRequestUrlDatasource {
	s.Delimiter = &v
	return s
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetOmitHeader(v bool) *WriteFeatureViewTableRequestUrlDatasource {
	s.OmitHeader = &v
	return s
}

func (s *WriteFeatureViewTableRequestUrlDatasource) SetPath(v string) *WriteFeatureViewTableRequestUrlDatasource {
	s.Path = &v
	return s
}

type WriteFeatureViewTableResponseBody struct {
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s WriteFeatureViewTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteFeatureViewTableResponseBody) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableResponseBody) SetRequestId(v string) *WriteFeatureViewTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *WriteFeatureViewTableResponseBody) SetTaskId(v string) *WriteFeatureViewTableResponseBody {
	s.TaskId = &v
	return s
}

type WriteFeatureViewTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WriteFeatureViewTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WriteFeatureViewTableResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteFeatureViewTableResponse) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableResponse) SetHeaders(v map[string]*string) *WriteFeatureViewTableResponse {
	s.Headers = v
	return s
}

func (s *WriteFeatureViewTableResponse) SetStatusCode(v int32) *WriteFeatureViewTableResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteFeatureViewTableResponse) SetBody(v *WriteFeatureViewTableResponseBody) *WriteFeatureViewTableResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paifeaturestore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检测资源连接状态。
//
// @param request - CheckInstanceDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceDatasourceResponse
func (client *Client) CheckInstanceDatasourceWithOptions(InstanceId *string, request *CheckInstanceDatasourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInstanceDatasourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInstanceDatasource"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/action/checkdatasource"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckInstanceDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测资源连接状态。
//
// @param request - CheckInstanceDatasourceRequest
//
// @return CheckInstanceDatasourceResponse
func (client *Client) CheckInstanceDatasource(InstanceId *string, request *CheckInstanceDatasourceRequest) (_result *CheckInstanceDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceDatasourceResponse{}
	_body, _err := client.CheckInstanceDatasourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据源。
//
// @param request - CreateDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasourceResponse
func (client *Client) CreateDatasourceWithOptions(InstanceId *string, request *CreateDatasourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatasource"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源。
//
// @param request - CreateDatasourceRequest
//
// @return CreateDatasourceResponse
func (client *Client) CreateDatasource(InstanceId *string, request *CreateDatasourceRequest) (_result *CreateDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasourceResponse{}
	_body, _err := client.CreateDatasourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建特征实体
//
// @param request - CreateFeatureEntityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureEntityResponse
func (client *Client) CreateFeatureEntityWithOptions(InstanceId *string, request *CreateFeatureEntityRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinId)) {
		body["JoinId"] = request.JoinId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFeatureEntity"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureentities"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFeatureEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征实体
//
// @param request - CreateFeatureEntityRequest
//
// @return CreateFeatureEntityResponse
func (client *Client) CreateFeatureEntity(InstanceId *string, request *CreateFeatureEntityRequest) (_result *CreateFeatureEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureEntityResponse{}
	_body, _err := client.CreateFeatureEntityWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建特征视图。
//
// @param request - CreateFeatureViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureViewResponse
func (client *Client) CreateFeatureViewWithOptions(InstanceId *string, request *CreateFeatureViewRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFeatureViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureEntityId)) {
		body["FeatureEntityId"] = request.FeatureEntityId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterDatasourceId)) {
		body["RegisterDatasourceId"] = request.RegisterDatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterTable)) {
		body["RegisterTable"] = request.RegisterTable
	}

	if !tea.BoolValue(util.IsUnset(request.SyncOnlineTable)) {
		body["SyncOnlineTable"] = request.SyncOnlineTable
	}

	if !tea.BoolValue(util.IsUnset(request.TTL)) {
		body["TTL"] = request.TTL
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WriteMethod)) {
		body["WriteMethod"] = request.WriteMethod
	}

	if !tea.BoolValue(util.IsUnset(request.WriteToFeatureDB)) {
		body["WriteToFeatureDB"] = request.WriteToFeatureDB
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFeatureView"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFeatureViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征视图。
//
// @param request - CreateFeatureViewRequest
//
// @return CreateFeatureViewResponse
func (client *Client) CreateFeatureView(InstanceId *string, request *CreateFeatureViewRequest) (_result *CreateFeatureViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureViewResponse{}
	_body, _err := client.CreateFeatureViewWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Feature Store实例。
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Feature Store实例。
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建label表
//
// @param request - CreateLabelTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLabelTableResponse
func (client *Client) CreateLabelTableWithOptions(InstanceId *string, request *CreateLabelTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLabelTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasourceId)) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabelTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/labeltables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLabelTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建label表
//
// @param request - CreateLabelTableRequest
//
// @return CreateLabelTableResponse
func (client *Client) CreateLabelTable(InstanceId *string, request *CreateLabelTableRequest) (_result *CreateLabelTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLabelTableResponse{}
	_body, _err := client.CreateLabelTableWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型特征。
//
// @param request - CreateModelFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelFeatureResponse
func (client *Client) CreateModelFeatureWithOptions(InstanceId *string, request *CreateModelFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Features)) {
		body["Features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.LabelPriorityLevel)) {
		body["LabelPriorityLevel"] = request.LabelPriorityLevel
	}

	if !tea.BoolValue(util.IsUnset(request.LabelTableId)) {
		body["LabelTableId"] = request.LabelTableId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceFeatureViewIds)) {
		body["SequenceFeatureViewIds"] = request.SequenceFeatureViewIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModelFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型特征。
//
// @param request - CreateModelFeatureRequest
//
// @return CreateModelFeatureResponse
func (client *Client) CreateModelFeature(InstanceId *string, request *CreateModelFeatureRequest) (_result *CreateModelFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelFeatureResponse{}
	_body, _err := client.CreateModelFeatureWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建FeatureStore项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(InstanceId *string, request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OfflineDatasourceId)) {
		body["OfflineDatasourceId"] = request.OfflineDatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.OfflineLifeCycle)) {
		body["OfflineLifeCycle"] = request.OfflineLifeCycle
	}

	if !tea.BoolValue(util.IsUnset(request.OnlineDatasourceId)) {
		body["OnlineDatasourceId"] = request.OnlineDatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// 创建FeatureStore项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(InstanceId *string, request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建feature store服务账户角色
//
// @param request - CreateServiceIdentityRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRoleWithOptions(request *CreateServiceIdentityRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceIdentityRole"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/serviceidentityroles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建feature store服务账户角色
//
// @param request - CreateServiceIdentityRoleRequest
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRole(request *CreateServiceIdentityRoleRequest) (_result *CreateServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CreateServiceIdentityRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定数据源。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasourceResponse
func (client *Client) DeleteDatasourceWithOptions(InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatasource"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(DatasourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定数据源。
//
// @return DeleteDatasourceResponse
func (client *Client) DeleteDatasource(InstanceId *string, DatasourceId *string) (_result *DeleteDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasourceResponse{}
	_body, _err := client.DeleteDatasourceWithOptions(InstanceId, DatasourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定特征实体
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureEntityResponse
func (client *Client) DeleteFeatureEntityWithOptions(InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFeatureEntityResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFeatureEntity"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureentities/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureEntityId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFeatureEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定特征实体
//
// @return DeleteFeatureEntityResponse
func (client *Client) DeleteFeatureEntity(InstanceId *string, FeatureEntityId *string) (_result *DeleteFeatureEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFeatureEntityResponse{}
	_body, _err := client.DeleteFeatureEntityWithOptions(InstanceId, FeatureEntityId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定特征视图。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureViewResponse
func (client *Client) DeleteFeatureViewWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFeatureViewResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFeatureView"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFeatureViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定特征视图。
//
// @return DeleteFeatureViewResponse
func (client *Client) DeleteFeatureView(InstanceId *string, FeatureViewId *string) (_result *DeleteFeatureViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFeatureViewResponse{}
	_body, _err := client.DeleteFeatureViewWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除label表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLabelTableResponse
func (client *Client) DeleteLabelTableWithOptions(InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLabelTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLabelTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/labeltables/" + tea.StringValue(openapiutil.GetEncodeParam(LabelTableId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLabelTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除label表
//
// @return DeleteLabelTableResponse
func (client *Client) DeleteLabelTable(InstanceId *string, LabelTableId *string) (_result *DeleteLabelTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLabelTableResponse{}
	_body, _err := client.DeleteLabelTableWithOptions(InstanceId, LabelTableId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定模型特征。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelFeatureResponse
func (client *Client) DeleteModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelFeatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModelFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定模型特征。
//
// @return DeleteModelFeatureResponse
func (client *Client) DeleteModelFeature(InstanceId *string, ModelFeatureId *string) (_result *DeleteModelFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelFeatureResponse{}
	_body, _err := client.DeleteModelFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定Feature Store项目。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// 删除指定Feature Store项目。
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(InstanceId *string, ProjectId *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出训练集表。
//
// @param request - ExportModelFeatureTrainingSetTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportModelFeatureTrainingSetTableResponse
func (client *Client) ExportModelFeatureTrainingSetTableWithOptions(InstanceId *string, ModelFeatureId *string, request *ExportModelFeatureTrainingSetTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExportModelFeatureTrainingSetTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureViewConfig)) {
		body["FeatureViewConfig"] = request.FeatureViewConfig
	}

	if !tea.BoolValue(util.IsUnset(request.LabelInputConfig)) {
		body["LabelInputConfig"] = request.LabelInputConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RealTimeIterateInterval)) {
		body["RealTimeIterateInterval"] = request.RealTimeIterateInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RealTimePartitionCountValue)) {
		body["RealTimePartitionCountValue"] = request.RealTimePartitionCountValue
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingSetConfig)) {
		body["TrainingSetConfig"] = request.TrainingSetConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportModelFeatureTrainingSetTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId)) + "/action/exporttrainingsettable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportModelFeatureTrainingSetTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出训练集表。
//
// @param request - ExportModelFeatureTrainingSetTableRequest
//
// @return ExportModelFeatureTrainingSetTableResponse
func (client *Client) ExportModelFeatureTrainingSetTable(InstanceId *string, ModelFeatureId *string, request *ExportModelFeatureTrainingSetTableRequest) (_result *ExportModelFeatureTrainingSetTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportModelFeatureTrainingSetTableResponse{}
	_body, _err := client.ExportModelFeatureTrainingSetTableWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasourceResponse
func (client *Client) GetDatasourceWithOptions(InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatasource"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(DatasourceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源详细信息。
//
// @return GetDatasourceResponse
func (client *Client) GetDatasource(InstanceId *string, DatasourceId *string) (_result *GetDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasourceResponse{}
	_body, _err := client.GetDatasourceWithOptions(InstanceId, DatasourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasourceTableResponse
func (client *Client) GetDatasourceTableWithOptions(InstanceId *string, DatasourceId *string, TableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasourceTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatasourceTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(DatasourceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(TableName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasourceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @return GetDatasourceTableResponse
func (client *Client) GetDatasourceTable(InstanceId *string, DatasourceId *string, TableName *string) (_result *GetDatasourceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasourceTableResponse{}
	_body, _err := client.GetDatasourceTableWithOptions(InstanceId, DatasourceId, TableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征实体详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureEntityResponse
func (client *Client) GetFeatureEntityWithOptions(InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFeatureEntityResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureEntity"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureentities/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureEntityId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeatureEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征实体详细信息
//
// @return GetFeatureEntityResponse
func (client *Client) GetFeatureEntity(InstanceId *string, FeatureEntityId *string) (_result *GetFeatureEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureEntityResponse{}
	_body, _err := client.GetFeatureEntityWithOptions(InstanceId, FeatureEntityId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征视图详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureViewResponse
func (client *Client) GetFeatureViewWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFeatureViewResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureView"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeatureViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图详细信息。
//
// @return GetFeatureViewResponse
func (client *Client) GetFeatureView(InstanceId *string, FeatureViewId *string) (_result *GetFeatureViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureViewResponse{}
	_body, _err := client.GetFeatureViewWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详细信息
//
// @return GetInstanceResponse
func (client *Client) GetInstance(InstanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Label表详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLabelTableResponse
func (client *Client) GetLabelTableWithOptions(InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLabelTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLabelTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/labeltables/" + tea.StringValue(openapiutil.GetEncodeParam(LabelTableId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLabelTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Label表详细信息。
//
// @return GetLabelTableResponse
func (client *Client) GetLabelTable(InstanceId *string, LabelTableId *string) (_result *GetLabelTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLabelTableResponse{}
	_body, _err := client.GetLabelTableWithOptions(InstanceId, LabelTableId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模型特征详情。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureResponse
func (client *Client) GetModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelFeatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征详情。
//
// @return GetModelFeatureResponse
func (client *Client) GetModelFeature(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureResponse{}
	_body, _err := client.GetModelFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模型特征的FG特征配置信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureFGFeatureResponse
func (client *Client) GetModelFeatureFGFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelFeatureFGFeatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelFeatureFGFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId)) + "/fgfeature"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelFeatureFGFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征的FG特征配置信息。
//
// @return GetModelFeatureFGFeatureResponse
func (client *Client) GetModelFeatureFGFeature(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureFGFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureFGFeatureResponse{}
	_body, _err := client.GetModelFeatureFGFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模型特征的fg.json文件配置信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureFGInfoResponse
func (client *Client) GetModelFeatureFGInfoWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelFeatureFGInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelFeatureFGInfo"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId)) + "/fginfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelFeatureFGInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征的fg.json文件配置信息。
//
// @return GetModelFeatureFGInfoResponse
func (client *Client) GetModelFeatureFGInfo(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureFGInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureFGInfoResponse{}
	_body, _err := client.GetModelFeatureFGInfoWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定Feature Store项目详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// 获取指定Feature Store项目详细信息。
//
// @return GetProjectResponse
func (client *Client) GetProject(InstanceId *string, ProjectId *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目下特征实体详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectFeatureEntityResponse
func (client *Client) GetProjectFeatureEntityWithOptions(InstanceId *string, ProjectId *string, FeatureEntityName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectFeatureEntityResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectFeatureEntity"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/featureentities/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureEntityName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectFeatureEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下特征实体详细信息
//
// @return GetProjectFeatureEntityResponse
func (client *Client) GetProjectFeatureEntity(InstanceId *string, ProjectId *string, FeatureEntityName *string) (_result *GetProjectFeatureEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectFeatureEntityResponse{}
	_body, _err := client.GetProjectFeatureEntityWithOptions(InstanceId, ProjectId, FeatureEntityName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取feature store服务账户角色。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRoleWithOptions(RoleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceIdentityRole"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/serviceidentityroles/" + tea.StringValue(openapiutil.GetEncodeParam(RoleName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取feature store服务账户角色。
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRole(RoleName *string) (_result *GetServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.GetServiceIdentityRoleWithOptions(RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(InstanceId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务详情
//
// @return GetTaskResponse
func (client *Client) GetTask(InstanceId *string, TaskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(InstanceId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源下所有表。
//
// @param request - ListDatasourceTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasourceTablesResponse
func (client *Client) ListDatasourceTablesWithOptions(InstanceId *string, DatasourceId *string, request *ListDatasourceTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatasourceTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasourceTables"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(DatasourceId)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasourceTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源下所有表。
//
// @param request - ListDatasourceTablesRequest
//
// @return ListDatasourceTablesResponse
func (client *Client) ListDatasourceTables(InstanceId *string, DatasourceId *string, request *ListDatasourceTablesRequest) (_result *ListDatasourceTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasourceTablesResponse{}
	_body, _err := client.ListDatasourceTablesWithOptions(InstanceId, DatasourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源列表。
//
// @param request - ListDatasourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasourcesResponse
func (client *Client) ListDatasourcesWithOptions(InstanceId *string, request *ListDatasourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatasourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasources"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源列表。
//
// @param request - ListDatasourcesRequest
//
// @return ListDatasourcesResponse
func (client *Client) ListDatasources(InstanceId *string, request *ListDatasourcesRequest) (_result *ListDatasourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasourcesResponse{}
	_body, _err := client.ListDatasourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建特征实体列表
//
// @param tmpReq - ListFeatureEntitiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureEntitiesResponse
func (client *Client) ListFeatureEntitiesWithOptions(InstanceId *string, tmpReq *ListFeatureEntitiesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureEntitiesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FeatureEntityIds)) {
		request.FeatureEntityIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureEntityIds, tea.String("FeatureEntityIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureEntityIdsShrink)) {
		query["FeatureEntityIds"] = request.FeatureEntityIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureEntities"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureentities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeatureEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征实体列表
//
// @param request - ListFeatureEntitiesRequest
//
// @return ListFeatureEntitiesResponse
func (client *Client) ListFeatureEntities(InstanceId *string, request *ListFeatureEntitiesRequest) (_result *ListFeatureEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureEntitiesResponse{}
	_body, _err := client.ListFeatureEntitiesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征字段血缘关系。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewFieldRelationshipsResponse
func (client *Client) ListFeatureViewFieldRelationshipsWithOptions(InstanceId *string, FeatureViewId *string, FieldName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureViewFieldRelationshipsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureViewFieldRelationships"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId)) + "/fields/" + tea.StringValue(openapiutil.GetEncodeParam(FieldName)) + "/relationships"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeatureViewFieldRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征字段血缘关系。
//
// @return ListFeatureViewFieldRelationshipsResponse
func (client *Client) ListFeatureViewFieldRelationships(InstanceId *string, FeatureViewId *string, FieldName *string) (_result *ListFeatureViewFieldRelationshipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewFieldRelationshipsResponse{}
	_body, _err := client.ListFeatureViewFieldRelationshipsWithOptions(InstanceId, FeatureViewId, FieldName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征视图下的在线特征数据。
//
// @param tmpReq - ListFeatureViewOnlineFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewOnlineFeaturesResponse
func (client *Client) ListFeatureViewOnlineFeaturesWithOptions(InstanceId *string, FeatureViewId *string, tmpReq *ListFeatureViewOnlineFeaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureViewOnlineFeaturesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureViewOnlineFeaturesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JoinIds)) {
		request.JoinIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JoinIds, tea.String("JoinIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinIdsShrink)) {
		query["JoinIds"] = request.JoinIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureViewOnlineFeatures"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId)) + "/onlinefeatures"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeatureViewOnlineFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图下的在线特征数据。
//
// @param request - ListFeatureViewOnlineFeaturesRequest
//
// @return ListFeatureViewOnlineFeaturesResponse
func (client *Client) ListFeatureViewOnlineFeatures(InstanceId *string, FeatureViewId *string, request *ListFeatureViewOnlineFeaturesRequest) (_result *ListFeatureViewOnlineFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewOnlineFeaturesResponse{}
	_body, _err := client.ListFeatureViewOnlineFeaturesWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewRelationshipsResponse
func (client *Client) ListFeatureViewRelationshipsWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureViewRelationshipsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureViewRelationships"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId)) + "/relationships"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeatureViewRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @return ListFeatureViewRelationshipsResponse
func (client *Client) ListFeatureViewRelationships(InstanceId *string, FeatureViewId *string) (_result *ListFeatureViewRelationshipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewRelationshipsResponse{}
	_body, _err := client.ListFeatureViewRelationshipsWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征视图列表。
//
// @param tmpReq - ListFeatureViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewsResponse
func (client *Client) ListFeatureViewsWithOptions(InstanceId *string, tmpReq *ListFeatureViewsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeatureViewsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureViewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FeatureViewIds)) {
		request.FeatureViewIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureViewIds, tea.String("FeatureViewIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureName)) {
		query["FeatureName"] = request.FeatureName
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureViewIdsShrink)) {
		query["FeatureViewIds"] = request.FeatureViewIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatureViews"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeatureViewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图列表。
//
// @param request - ListFeatureViewsRequest
//
// @return ListFeatureViewsResponse
func (client *Client) ListFeatureViews(InstanceId *string, request *ListFeatureViewsRequest) (_result *ListFeatureViewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewsResponse{}
	_body, _err := client.ListFeatureViewsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Feature Store实例列表。
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Feature Store实例列表。
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Label表列表。
//
// @param tmpReq - ListLabelTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLabelTablesResponse
func (client *Client) ListLabelTablesWithOptions(InstanceId *string, tmpReq *ListLabelTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLabelTablesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListLabelTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelTableIds)) {
		request.LabelTableIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelTableIds, tea.String("LabelTableIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelTableIdsShrink)) {
		query["LabelTableIds"] = request.LabelTableIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLabelTables"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/labeltables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLabelTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Label表列表。
//
// @param request - ListLabelTablesRequest
//
// @return ListLabelTablesResponse
func (client *Client) ListLabelTables(InstanceId *string, request *ListLabelTablesRequest) (_result *ListLabelTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLabelTablesResponse{}
	_body, _err := client.ListLabelTablesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取注册FG特征时模型特征下可选的所有特征。
//
// @param request - ListModelFeatureAvailableFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelFeatureAvailableFeaturesResponse
func (client *Client) ListModelFeatureAvailableFeaturesWithOptions(InstanceId *string, ModelFeatureId *string, request *ListModelFeatureAvailableFeaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelFeatureAvailableFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureName)) {
		query["FeatureName"] = request.FeatureName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModelFeatureAvailableFeatures"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId)) + "/availablefeatures"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelFeatureAvailableFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取注册FG特征时模型特征下可选的所有特征。
//
// @param request - ListModelFeatureAvailableFeaturesRequest
//
// @return ListModelFeatureAvailableFeaturesResponse
func (client *Client) ListModelFeatureAvailableFeatures(InstanceId *string, ModelFeatureId *string, request *ListModelFeatureAvailableFeaturesRequest) (_result *ListModelFeatureAvailableFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelFeatureAvailableFeaturesResponse{}
	_body, _err := client.ListModelFeatureAvailableFeaturesWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模型特征列表。
//
// @param tmpReq - ListModelFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelFeaturesResponse
func (client *Client) ListModelFeaturesWithOptions(InstanceId *string, tmpReq *ListModelFeaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelFeaturesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListModelFeaturesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ModelFeatureIds)) {
		request.ModelFeatureIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelFeatureIds, tea.String("ModelFeatureIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelFeatureIdsShrink)) {
		query["ModelFeatureIds"] = request.ModelFeatureIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModelFeatures"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征列表。
//
// @param request - ListModelFeaturesRequest
//
// @return ListModelFeaturesResponse
func (client *Client) ListModelFeatures(InstanceId *string, request *ListModelFeaturesRequest) (_result *ListModelFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelFeaturesResponse{}
	_body, _err := client.ListModelFeaturesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目下的所有特征视图、特征信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectFeatureViewsResponse
func (client *Client) ListProjectFeatureViewsWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectFeatureViewsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectFeatureViews"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/featureviews"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectFeatureViewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下的所有特征视图、特征信息。
//
// @return ListProjectFeatureViewsResponse
func (client *Client) ListProjectFeatureViews(InstanceId *string, ProjectId *string) (_result *ListProjectFeatureViewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectFeatureViewsResponse{}
	_body, _err := client.ListProjectFeatureViewsWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Feature Store项目列表。
//
// @param tmpReq - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(InstanceId *string, tmpReq *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ProjectIds)) {
		request.ProjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProjectIds, tea.String("ProjectIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIdsShrink)) {
		query["ProjectIds"] = request.ProjectIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// 获取Feature Store项目列表。
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(InstanceId *string, request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务日志列表
//
// @param request - ListTaskLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskLogsResponse
func (client *Client) ListTaskLogsWithOptions(InstanceId *string, TaskId *string, request *ListTaskLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTaskLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskLogs"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务日志列表
//
// @param request - ListTaskLogsRequest
//
// @return ListTaskLogsResponse
func (client *Client) ListTaskLogs(InstanceId *string, TaskId *string, request *ListTaskLogsRequest) (_result *ListTaskLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTaskLogsResponse{}
	_body, _err := client.ListTaskLogsWithOptions(InstanceId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务列表
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(InstanceId *string, tmpReq *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskIds)) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, tea.String("TaskIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdsShrink)) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务列表
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(InstanceId *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将特征视图的离线数据发布/同步到线上。
//
// @param request - PublishFeatureViewTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFeatureViewTableResponse
func (client *Client) PublishFeatureViewTableWithOptions(InstanceId *string, FeatureViewId *string, request *PublishFeatureViewTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishFeatureViewTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.EventTime)) {
		body["EventTime"] = request.EventTime
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OfflineToOnline)) {
		body["OfflineToOnline"] = request.OfflineToOnline
	}

	if !tea.BoolValue(util.IsUnset(request.Partitions)) {
		body["Partitions"] = request.Partitions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishFeatureViewTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId)) + "/action/publishtable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishFeatureViewTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将特征视图的离线数据发布/同步到线上。
//
// @param request - PublishFeatureViewTableRequest
//
// @return PublishFeatureViewTableResponse
func (client *Client) PublishFeatureViewTable(InstanceId *string, FeatureViewId *string, request *PublishFeatureViewTableRequest) (_result *PublishFeatureViewTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishFeatureViewTableResponse{}
	_body, _err := client.PublishFeatureViewTableWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据源信息。
//
// @param request - UpdateDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasourceResponse
func (client *Client) UpdateDatasourceWithOptions(InstanceId *string, DatasourceId *string, request *UpdateDatasourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDatasourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDatasource"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/datasources/" + tea.StringValue(openapiutil.GetEncodeParam(DatasourceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据源信息。
//
// @param request - UpdateDatasourceRequest
//
// @return UpdateDatasourceResponse
func (client *Client) UpdateDatasource(InstanceId *string, DatasourceId *string, request *UpdateDatasourceRequest) (_result *UpdateDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasourceResponse{}
	_body, _err := client.UpdateDatasourceWithOptions(InstanceId, DatasourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新label表。
//
// @param request - UpdateLabelTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLabelTableResponse
func (client *Client) UpdateLabelTableWithOptions(InstanceId *string, LabelTableId *string, request *UpdateLabelTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLabelTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasourceId)) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLabelTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/labeltables/" + tea.StringValue(openapiutil.GetEncodeParam(LabelTableId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLabelTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新label表。
//
// @param request - UpdateLabelTableRequest
//
// @return UpdateLabelTableResponse
func (client *Client) UpdateLabelTable(InstanceId *string, LabelTableId *string, request *UpdateLabelTableRequest) (_result *UpdateLabelTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLabelTableResponse{}
	_body, _err := client.UpdateLabelTableWithOptions(InstanceId, LabelTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新模型特征。
//
// @param request - UpdateModelFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelFeatureResponse
func (client *Client) UpdateModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModelFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Features)) {
		body["Features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.LabelPriorityLevel)) {
		body["LabelPriorityLevel"] = request.LabelPriorityLevel
	}

	if !tea.BoolValue(util.IsUnset(request.LabelTableId)) {
		body["LabelTableId"] = request.LabelTableId
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceFeatureViewIds)) {
		body["SequenceFeatureViewIds"] = request.SequenceFeatureViewIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModelFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModelFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型特征。
//
// @param request - UpdateModelFeatureRequest
//
// @return UpdateModelFeatureResponse
func (client *Client) UpdateModelFeature(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureRequest) (_result *UpdateModelFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelFeatureResponse{}
	_body, _err := client.UpdateModelFeatureWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新模型特征的FG特征配置信息。
//
// @param request - UpdateModelFeatureFGFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelFeatureFGFeatureResponse
func (client *Client) UpdateModelFeatureFGFeatureWithOptions(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureFGFeatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModelFeatureFGFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LookupFeatures)) {
		body["LookupFeatures"] = request.LookupFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.RawFeatures)) {
		body["RawFeatures"] = request.RawFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.Reserves)) {
		body["Reserves"] = request.Reserves
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceFeatures)) {
		body["SequenceFeatures"] = request.SequenceFeatures
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModelFeatureFGFeature"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/modelfeatures/" + tea.StringValue(openapiutil.GetEncodeParam(ModelFeatureId)) + "/fgfeature"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModelFeatureFGFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型特征的FG特征配置信息。
//
// @param request - UpdateModelFeatureFGFeatureRequest
//
// @return UpdateModelFeatureFGFeatureResponse
func (client *Client) UpdateModelFeatureFGFeature(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureFGFeatureRequest) (_result *UpdateModelFeatureFGFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelFeatureFGFeatureResponse{}
	_body, _err := client.UpdateModelFeatureFGFeatureWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新指定Feature Store项目信息。
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(InstanceId *string, ProjectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// 更新指定Feature Store项目信息。
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(InstanceId *string, ProjectId *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(InstanceId, ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @param request - WriteFeatureViewTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WriteFeatureViewTableResponse
func (client *Client) WriteFeatureViewTableWithOptions(InstanceId *string, FeatureViewId *string, request *WriteFeatureViewTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WriteFeatureViewTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Partitions)) {
		body["Partitions"] = request.Partitions
	}

	if !tea.BoolValue(util.IsUnset(request.UrlDatasource)) {
		body["UrlDatasource"] = request.UrlDatasource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("WriteFeatureViewTable"),
		Version:     tea.String("2023-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/featureviews/" + tea.StringValue(openapiutil.GetEncodeParam(FeatureViewId)) + "/action/writetable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteFeatureViewTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @param request - WriteFeatureViewTableRequest
//
// @return WriteFeatureViewTableResponse
func (client *Client) WriteFeatureViewTable(InstanceId *string, FeatureViewId *string, request *WriteFeatureViewTableRequest) (_result *WriteFeatureViewTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WriteFeatureViewTableResponse{}
	_body, _err := client.WriteFeatureViewTableWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
