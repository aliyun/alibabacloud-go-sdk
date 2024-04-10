// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Adb4MysqlSparkDiagnosisInfo struct {
	DiagnosisCode      *string `json:"DiagnosisCode,omitempty" xml:"DiagnosisCode,omitempty"`
	DiagnosisCodeLabel *string `json:"DiagnosisCodeLabel,omitempty" xml:"DiagnosisCodeLabel,omitempty"`
	DiagnosisMsg       *string `json:"DiagnosisMsg,omitempty" xml:"DiagnosisMsg,omitempty"`
	DiagnosisType      *string `json:"DiagnosisType,omitempty" xml:"DiagnosisType,omitempty"`
}

func (s Adb4MysqlSparkDiagnosisInfo) String() string {
	return tea.Prettify(s)
}

func (s Adb4MysqlSparkDiagnosisInfo) GoString() string {
	return s.String()
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisCode(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisCode = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisCodeLabel(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisCodeLabel = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisMsg(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisMsg = &v
	return s
}

func (s *Adb4MysqlSparkDiagnosisInfo) SetDiagnosisType(v string) *Adb4MysqlSparkDiagnosisInfo {
	s.DiagnosisType = &v
	return s
}

type ColDetailModel struct {
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributeKey *bool   `json:"DistributeKey,omitempty" xml:"DistributeKey,omitempty"`
	Nullable      *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	PartitionKey  *bool   `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	PrimaryKey    *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName    *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ColDetailModel) String() string {
	return tea.Prettify(s)
}

func (s ColDetailModel) GoString() string {
	return s.String()
}

func (s *ColDetailModel) SetColumnName(v string) *ColDetailModel {
	s.ColumnName = &v
	return s
}

func (s *ColDetailModel) SetCreateTime(v string) *ColDetailModel {
	s.CreateTime = &v
	return s
}

func (s *ColDetailModel) SetDescription(v string) *ColDetailModel {
	s.Description = &v
	return s
}

func (s *ColDetailModel) SetDistributeKey(v bool) *ColDetailModel {
	s.DistributeKey = &v
	return s
}

func (s *ColDetailModel) SetNullable(v bool) *ColDetailModel {
	s.Nullable = &v
	return s
}

func (s *ColDetailModel) SetPartitionKey(v bool) *ColDetailModel {
	s.PartitionKey = &v
	return s
}

func (s *ColDetailModel) SetPrimaryKey(v bool) *ColDetailModel {
	s.PrimaryKey = &v
	return s
}

func (s *ColDetailModel) SetSchemaName(v string) *ColDetailModel {
	s.SchemaName = &v
	return s
}

func (s *ColDetailModel) SetTableName(v string) *ColDetailModel {
	s.TableName = &v
	return s
}

func (s *ColDetailModel) SetType(v string) *ColDetailModel {
	s.Type = &v
	return s
}

func (s *ColDetailModel) SetUpdateTime(v string) *ColDetailModel {
	s.UpdateTime = &v
	return s
}

type CstoreIndexModel struct {
	ColumnOrds        []*string           `json:"ColumnOrds,omitempty" xml:"ColumnOrds,omitempty" type:"Repeated"`
	CreateTime        *string             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName      *string             `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IndexColumns      []*FieldSchemaModel `json:"IndexColumns,omitempty" xml:"IndexColumns,omitempty" type:"Repeated"`
	IndexName         *string             `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexType         *string             `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	Options           map[string]*string  `json:"Options,omitempty" xml:"Options,omitempty"`
	PhysicalTableName *string             `json:"PhysicalTableName,omitempty" xml:"PhysicalTableName,omitempty"`
	UpdateTime        *string             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CstoreIndexModel) String() string {
	return tea.Prettify(s)
}

func (s CstoreIndexModel) GoString() string {
	return s.String()
}

func (s *CstoreIndexModel) SetColumnOrds(v []*string) *CstoreIndexModel {
	s.ColumnOrds = v
	return s
}

func (s *CstoreIndexModel) SetCreateTime(v string) *CstoreIndexModel {
	s.CreateTime = &v
	return s
}

func (s *CstoreIndexModel) SetDatabaseName(v string) *CstoreIndexModel {
	s.DatabaseName = &v
	return s
}

func (s *CstoreIndexModel) SetIndexColumns(v []*FieldSchemaModel) *CstoreIndexModel {
	s.IndexColumns = v
	return s
}

func (s *CstoreIndexModel) SetIndexName(v string) *CstoreIndexModel {
	s.IndexName = &v
	return s
}

func (s *CstoreIndexModel) SetIndexType(v string) *CstoreIndexModel {
	s.IndexType = &v
	return s
}

func (s *CstoreIndexModel) SetOptions(v map[string]*string) *CstoreIndexModel {
	s.Options = v
	return s
}

func (s *CstoreIndexModel) SetPhysicalTableName(v string) *CstoreIndexModel {
	s.PhysicalTableName = &v
	return s
}

func (s *CstoreIndexModel) SetUpdateTime(v string) *CstoreIndexModel {
	s.UpdateTime = &v
	return s
}

type DatabaseSummaryModel struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DatabaseSummaryModel) String() string {
	return tea.Prettify(s)
}

func (s DatabaseSummaryModel) GoString() string {
	return s.String()
}

func (s *DatabaseSummaryModel) SetCreateTime(v string) *DatabaseSummaryModel {
	s.CreateTime = &v
	return s
}

func (s *DatabaseSummaryModel) SetDescription(v string) *DatabaseSummaryModel {
	s.Description = &v
	return s
}

func (s *DatabaseSummaryModel) SetOwner(v string) *DatabaseSummaryModel {
	s.Owner = &v
	return s
}

func (s *DatabaseSummaryModel) SetSchemaName(v string) *DatabaseSummaryModel {
	s.SchemaName = &v
	return s
}

func (s *DatabaseSummaryModel) SetUpdateTime(v string) *DatabaseSummaryModel {
	s.UpdateTime = &v
	return s
}

type Detail struct {
	AppType                           *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	DBClusterId                       *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Data                              *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DurationInMillis                  *int64  `json:"DurationInMillis,omitempty" xml:"DurationInMillis,omitempty"`
	EstimateExecutionCpuTimeInSeconds *int64  `json:"EstimateExecutionCpuTimeInSeconds,omitempty" xml:"EstimateExecutionCpuTimeInSeconds,omitempty"`
	LastAttemptId                     *string `json:"LastAttemptId,omitempty" xml:"LastAttemptId,omitempty"`
	LastUpdatedTimeInMillis           *int64  `json:"LastUpdatedTimeInMillis,omitempty" xml:"LastUpdatedTimeInMillis,omitempty"`
	LogRootPath                       *string `json:"LogRootPath,omitempty" xml:"LogRootPath,omitempty"`
	ResourceGroupName                 *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	StartedTimeInMillis               *int64  `json:"StartedTimeInMillis,omitempty" xml:"StartedTimeInMillis,omitempty"`
	SubmittedTimeInMillis             *int64  `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
	TerminatedTimeInMillis            *int64  `json:"TerminatedTimeInMillis,omitempty" xml:"TerminatedTimeInMillis,omitempty"`
	WebUiAddress                      *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s Detail) String() string {
	return tea.Prettify(s)
}

func (s Detail) GoString() string {
	return s.String()
}

func (s *Detail) SetAppType(v string) *Detail {
	s.AppType = &v
	return s
}

func (s *Detail) SetDBClusterId(v string) *Detail {
	s.DBClusterId = &v
	return s
}

func (s *Detail) SetData(v string) *Detail {
	s.Data = &v
	return s
}

func (s *Detail) SetDurationInMillis(v int64) *Detail {
	s.DurationInMillis = &v
	return s
}

func (s *Detail) SetEstimateExecutionCpuTimeInSeconds(v int64) *Detail {
	s.EstimateExecutionCpuTimeInSeconds = &v
	return s
}

func (s *Detail) SetLastAttemptId(v string) *Detail {
	s.LastAttemptId = &v
	return s
}

func (s *Detail) SetLastUpdatedTimeInMillis(v int64) *Detail {
	s.LastUpdatedTimeInMillis = &v
	return s
}

func (s *Detail) SetLogRootPath(v string) *Detail {
	s.LogRootPath = &v
	return s
}

func (s *Detail) SetResourceGroupName(v string) *Detail {
	s.ResourceGroupName = &v
	return s
}

func (s *Detail) SetStartedTimeInMillis(v int64) *Detail {
	s.StartedTimeInMillis = &v
	return s
}

func (s *Detail) SetSubmittedTimeInMillis(v int64) *Detail {
	s.SubmittedTimeInMillis = &v
	return s
}

func (s *Detail) SetTerminatedTimeInMillis(v int64) *Detail {
	s.TerminatedTimeInMillis = &v
	return s
}

func (s *Detail) SetWebUiAddress(v string) *Detail {
	s.WebUiAddress = &v
	return s
}

type FieldSchemaModel struct {
	AutoIncrement         *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	ColumnRawName         *string `json:"ColumnRawName,omitempty" xml:"ColumnRawName,omitempty"`
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CompressFloatUseShort *bool   `json:"CompressFloatUseShort,omitempty" xml:"CompressFloatUseShort,omitempty"`
	Compression           *string `json:"Compression,omitempty" xml:"Compression,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataType              *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DatabaseName          *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DefaultValue          *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Delimiter             *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	Encode                *string `json:"Encode,omitempty" xml:"Encode,omitempty"`
	IsPartitionKey        *bool   `json:"IsPartitionKey,omitempty" xml:"IsPartitionKey,omitempty"`
	MappedName            *string `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nullable              *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	OnUpdate              *string `json:"OnUpdate,omitempty" xml:"OnUpdate,omitempty"`
	OrdinalPosition       *int64  `json:"OrdinalPosition,omitempty" xml:"OrdinalPosition,omitempty"`
	PhysicalColumnName    *string `json:"PhysicalColumnName,omitempty" xml:"PhysicalColumnName,omitempty"`
	PkPosition            *int64  `json:"PkPosition,omitempty" xml:"PkPosition,omitempty"`
	Precision             *int64  `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Primarykey            *bool   `json:"Primarykey,omitempty" xml:"Primarykey,omitempty"`
	Scale                 *int64  `json:"Scale,omitempty" xml:"Scale,omitempty"`
	TableName             *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Tokenizer             *string `json:"Tokenizer,omitempty" xml:"Tokenizer,omitempty"`
	Type                  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ValueType             *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s FieldSchemaModel) String() string {
	return tea.Prettify(s)
}

func (s FieldSchemaModel) GoString() string {
	return s.String()
}

func (s *FieldSchemaModel) SetAutoIncrement(v bool) *FieldSchemaModel {
	s.AutoIncrement = &v
	return s
}

func (s *FieldSchemaModel) SetColumnRawName(v string) *FieldSchemaModel {
	s.ColumnRawName = &v
	return s
}

func (s *FieldSchemaModel) SetComment(v string) *FieldSchemaModel {
	s.Comment = &v
	return s
}

func (s *FieldSchemaModel) SetCompressFloatUseShort(v bool) *FieldSchemaModel {
	s.CompressFloatUseShort = &v
	return s
}

func (s *FieldSchemaModel) SetCompression(v string) *FieldSchemaModel {
	s.Compression = &v
	return s
}

func (s *FieldSchemaModel) SetCreateTime(v string) *FieldSchemaModel {
	s.CreateTime = &v
	return s
}

func (s *FieldSchemaModel) SetDataType(v string) *FieldSchemaModel {
	s.DataType = &v
	return s
}

func (s *FieldSchemaModel) SetDatabaseName(v string) *FieldSchemaModel {
	s.DatabaseName = &v
	return s
}

func (s *FieldSchemaModel) SetDefaultValue(v string) *FieldSchemaModel {
	s.DefaultValue = &v
	return s
}

func (s *FieldSchemaModel) SetDelimiter(v string) *FieldSchemaModel {
	s.Delimiter = &v
	return s
}

func (s *FieldSchemaModel) SetEncode(v string) *FieldSchemaModel {
	s.Encode = &v
	return s
}

func (s *FieldSchemaModel) SetIsPartitionKey(v bool) *FieldSchemaModel {
	s.IsPartitionKey = &v
	return s
}

func (s *FieldSchemaModel) SetMappedName(v string) *FieldSchemaModel {
	s.MappedName = &v
	return s
}

func (s *FieldSchemaModel) SetName(v string) *FieldSchemaModel {
	s.Name = &v
	return s
}

func (s *FieldSchemaModel) SetNullable(v bool) *FieldSchemaModel {
	s.Nullable = &v
	return s
}

func (s *FieldSchemaModel) SetOnUpdate(v string) *FieldSchemaModel {
	s.OnUpdate = &v
	return s
}

func (s *FieldSchemaModel) SetOrdinalPosition(v int64) *FieldSchemaModel {
	s.OrdinalPosition = &v
	return s
}

func (s *FieldSchemaModel) SetPhysicalColumnName(v string) *FieldSchemaModel {
	s.PhysicalColumnName = &v
	return s
}

func (s *FieldSchemaModel) SetPkPosition(v int64) *FieldSchemaModel {
	s.PkPosition = &v
	return s
}

func (s *FieldSchemaModel) SetPrecision(v int64) *FieldSchemaModel {
	s.Precision = &v
	return s
}

func (s *FieldSchemaModel) SetPrimarykey(v bool) *FieldSchemaModel {
	s.Primarykey = &v
	return s
}

func (s *FieldSchemaModel) SetScale(v int64) *FieldSchemaModel {
	s.Scale = &v
	return s
}

func (s *FieldSchemaModel) SetTableName(v string) *FieldSchemaModel {
	s.TableName = &v
	return s
}

func (s *FieldSchemaModel) SetTokenizer(v string) *FieldSchemaModel {
	s.Tokenizer = &v
	return s
}

func (s *FieldSchemaModel) SetType(v string) *FieldSchemaModel {
	s.Type = &v
	return s
}

func (s *FieldSchemaModel) SetUpdateTime(v string) *FieldSchemaModel {
	s.UpdateTime = &v
	return s
}

func (s *FieldSchemaModel) SetValueType(v string) *FieldSchemaModel {
	s.ValueType = &v
	return s
}

type Filters struct {
	AppIdRegex         *string                    `json:"AppIdRegex,omitempty" xml:"AppIdRegex,omitempty"`
	AppNameRegex       *string                    `json:"AppNameRegex,omitempty" xml:"AppNameRegex,omitempty"`
	AppState           *string                    `json:"AppState,omitempty" xml:"AppState,omitempty"`
	AppType            *string                    `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ExecutionTimeRange *FiltersExecutionTimeRange `json:"ExecutionTimeRange,omitempty" xml:"ExecutionTimeRange,omitempty" type:"Struct"`
	SubmitTimeRange    *FiltersSubmitTimeRange    `json:"SubmitTimeRange,omitempty" xml:"SubmitTimeRange,omitempty" type:"Struct"`
	TermiatedTimeRange *FiltersTermiatedTimeRange `json:"TermiatedTimeRange,omitempty" xml:"TermiatedTimeRange,omitempty" type:"Struct"`
}

func (s Filters) String() string {
	return tea.Prettify(s)
}

func (s Filters) GoString() string {
	return s.String()
}

func (s *Filters) SetAppIdRegex(v string) *Filters {
	s.AppIdRegex = &v
	return s
}

func (s *Filters) SetAppNameRegex(v string) *Filters {
	s.AppNameRegex = &v
	return s
}

func (s *Filters) SetAppState(v string) *Filters {
	s.AppState = &v
	return s
}

func (s *Filters) SetAppType(v string) *Filters {
	s.AppType = &v
	return s
}

func (s *Filters) SetExecutionTimeRange(v *FiltersExecutionTimeRange) *Filters {
	s.ExecutionTimeRange = v
	return s
}

func (s *Filters) SetSubmitTimeRange(v *FiltersSubmitTimeRange) *Filters {
	s.SubmitTimeRange = v
	return s
}

func (s *Filters) SetTermiatedTimeRange(v *FiltersTermiatedTimeRange) *Filters {
	s.TermiatedTimeRange = v
	return s
}

type FiltersExecutionTimeRange struct {
	MaxTimeInSeconds *int64 `json:"MaxTimeInSeconds,omitempty" xml:"MaxTimeInSeconds,omitempty"`
	MinTimeInSeconds *int64 `json:"MinTimeInSeconds,omitempty" xml:"MinTimeInSeconds,omitempty"`
}

func (s FiltersExecutionTimeRange) String() string {
	return tea.Prettify(s)
}

func (s FiltersExecutionTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersExecutionTimeRange) SetMaxTimeInSeconds(v int64) *FiltersExecutionTimeRange {
	s.MaxTimeInSeconds = &v
	return s
}

func (s *FiltersExecutionTimeRange) SetMinTimeInSeconds(v int64) *FiltersExecutionTimeRange {
	s.MinTimeInSeconds = &v
	return s
}

type FiltersSubmitTimeRange struct {
	MaxTimeInMills *int64 `json:"MaxTimeInMills,omitempty" xml:"MaxTimeInMills,omitempty"`
	MinTimeInMills *int64 `json:"MinTimeInMills,omitempty" xml:"MinTimeInMills,omitempty"`
}

func (s FiltersSubmitTimeRange) String() string {
	return tea.Prettify(s)
}

func (s FiltersSubmitTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersSubmitTimeRange) SetMaxTimeInMills(v int64) *FiltersSubmitTimeRange {
	s.MaxTimeInMills = &v
	return s
}

func (s *FiltersSubmitTimeRange) SetMinTimeInMills(v int64) *FiltersSubmitTimeRange {
	s.MinTimeInMills = &v
	return s
}

type FiltersTermiatedTimeRange struct {
	MaxTimeInMills *int64 `json:"MaxTimeInMills,omitempty" xml:"MaxTimeInMills,omitempty"`
	MinTimeInMills *int64 `json:"MinTimeInMills,omitempty" xml:"MinTimeInMills,omitempty"`
}

func (s FiltersTermiatedTimeRange) String() string {
	return tea.Prettify(s)
}

func (s FiltersTermiatedTimeRange) GoString() string {
	return s.String()
}

func (s *FiltersTermiatedTimeRange) SetMaxTimeInMills(v int64) *FiltersTermiatedTimeRange {
	s.MaxTimeInMills = &v
	return s
}

func (s *FiltersTermiatedTimeRange) SetMinTimeInMills(v int64) *FiltersTermiatedTimeRange {
	s.MinTimeInMills = &v
	return s
}

type LogAnalyzeResult struct {
	AppErrorAdvice *string `json:"AppErrorAdvice,omitempty" xml:"AppErrorAdvice,omitempty"`
	AppErrorCode   *string `json:"AppErrorCode,omitempty" xml:"AppErrorCode,omitempty"`
	AppErrorLog    *string `json:"AppErrorLog,omitempty" xml:"AppErrorLog,omitempty"`
}

func (s LogAnalyzeResult) String() string {
	return tea.Prettify(s)
}

func (s LogAnalyzeResult) GoString() string {
	return s.String()
}

func (s *LogAnalyzeResult) SetAppErrorAdvice(v string) *LogAnalyzeResult {
	s.AppErrorAdvice = &v
	return s
}

func (s *LogAnalyzeResult) SetAppErrorCode(v string) *LogAnalyzeResult {
	s.AppErrorCode = &v
	return s
}

func (s *LogAnalyzeResult) SetAppErrorLog(v string) *LogAnalyzeResult {
	s.AppErrorLog = &v
	return s
}

type OperatorNode struct {
	Children   []*OperatorNode    `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	Id         *int32             `json:"id,omitempty" xml:"id,omitempty"`
	LevelWidth *int32             `json:"levelWidth,omitempty" xml:"levelWidth,omitempty"`
	NodeDepth  *int32             `json:"nodeDepth,omitempty" xml:"nodeDepth,omitempty"`
	NodeName   *string            `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	NodeWidth  *int32             `json:"nodeWidth,omitempty" xml:"nodeWidth,omitempty"`
	ParentId   *int32             `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Stats      *OperatorNodeStats `json:"stats,omitempty" xml:"stats,omitempty" type:"Struct"`
}

func (s OperatorNode) String() string {
	return tea.Prettify(s)
}

func (s OperatorNode) GoString() string {
	return s.String()
}

func (s *OperatorNode) SetChildren(v []*OperatorNode) *OperatorNode {
	s.Children = v
	return s
}

func (s *OperatorNode) SetId(v int32) *OperatorNode {
	s.Id = &v
	return s
}

func (s *OperatorNode) SetLevelWidth(v int32) *OperatorNode {
	s.LevelWidth = &v
	return s
}

func (s *OperatorNode) SetNodeDepth(v int32) *OperatorNode {
	s.NodeDepth = &v
	return s
}

func (s *OperatorNode) SetNodeName(v string) *OperatorNode {
	s.NodeName = &v
	return s
}

func (s *OperatorNode) SetNodeWidth(v int32) *OperatorNode {
	s.NodeWidth = &v
	return s
}

func (s *OperatorNode) SetParentId(v int32) *OperatorNode {
	s.ParentId = &v
	return s
}

func (s *OperatorNode) SetStats(v *OperatorNodeStats) *OperatorNode {
	s.Stats = v
	return s
}

type OperatorNodeStats struct {
	Bytes      *int64  `json:"bytes,omitempty" xml:"bytes,omitempty"`
	OutputRows *int64  `json:"outputRows,omitempty" xml:"outputRows,omitempty"`
	Parameters *string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	PeakMemory *int64  `json:"peakMemory,omitempty" xml:"peakMemory,omitempty"`
	TimeCost   *int64  `json:"timeCost,omitempty" xml:"timeCost,omitempty"`
}

func (s OperatorNodeStats) String() string {
	return tea.Prettify(s)
}

func (s OperatorNodeStats) GoString() string {
	return s.String()
}

func (s *OperatorNodeStats) SetBytes(v int64) *OperatorNodeStats {
	s.Bytes = &v
	return s
}

func (s *OperatorNodeStats) SetOutputRows(v int64) *OperatorNodeStats {
	s.OutputRows = &v
	return s
}

func (s *OperatorNodeStats) SetParameters(v string) *OperatorNodeStats {
	s.Parameters = &v
	return s
}

func (s *OperatorNodeStats) SetPeakMemory(v int64) *OperatorNodeStats {
	s.PeakMemory = &v
	return s
}

func (s *OperatorNodeStats) SetTimeCost(v int64) *OperatorNodeStats {
	s.TimeCost = &v
	return s
}

type SerDeInfoModel struct {
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerDeId          *int64             `json:"SerDeId,omitempty" xml:"SerDeId,omitempty"`
	SerializationLib *string            `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
}

func (s SerDeInfoModel) String() string {
	return tea.Prettify(s)
}

func (s SerDeInfoModel) GoString() string {
	return s.String()
}

func (s *SerDeInfoModel) SetName(v string) *SerDeInfoModel {
	s.Name = &v
	return s
}

func (s *SerDeInfoModel) SetParameters(v map[string]*string) *SerDeInfoModel {
	s.Parameters = v
	return s
}

func (s *SerDeInfoModel) SetSerDeId(v int64) *SerDeInfoModel {
	s.SerDeId = &v
	return s
}

func (s *SerDeInfoModel) SetSerializationLib(v string) *SerDeInfoModel {
	s.SerializationLib = &v
	return s
}

type SparkAnalyzeLogTask struct {
	DBClusterId            *string           `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Result                 *LogAnalyzeResult `json:"Result,omitempty" xml:"Result,omitempty"`
	RuleMatched            *bool             `json:"RuleMatched,omitempty" xml:"RuleMatched,omitempty"`
	StartedTimeInMillis    *int64            `json:"StartedTimeInMillis,omitempty" xml:"StartedTimeInMillis,omitempty"`
	SubmittedTimeInMillis  *int64            `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
	TaskErrMsg             *string           `json:"TaskErrMsg,omitempty" xml:"TaskErrMsg,omitempty"`
	TaskId                 *int64            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskState              *string           `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	TerminatedTimeInMillis *int64            `json:"TerminatedTimeInMillis,omitempty" xml:"TerminatedTimeInMillis,omitempty"`
	UserId                 *int64            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SparkAnalyzeLogTask) String() string {
	return tea.Prettify(s)
}

func (s SparkAnalyzeLogTask) GoString() string {
	return s.String()
}

func (s *SparkAnalyzeLogTask) SetDBClusterId(v string) *SparkAnalyzeLogTask {
	s.DBClusterId = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetResult(v *LogAnalyzeResult) *SparkAnalyzeLogTask {
	s.Result = v
	return s
}

func (s *SparkAnalyzeLogTask) SetRuleMatched(v bool) *SparkAnalyzeLogTask {
	s.RuleMatched = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetStartedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.StartedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetSubmittedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.SubmittedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskErrMsg(v string) *SparkAnalyzeLogTask {
	s.TaskErrMsg = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskId(v int64) *SparkAnalyzeLogTask {
	s.TaskId = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTaskState(v string) *SparkAnalyzeLogTask {
	s.TaskState = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetTerminatedTimeInMillis(v int64) *SparkAnalyzeLogTask {
	s.TerminatedTimeInMillis = &v
	return s
}

func (s *SparkAnalyzeLogTask) SetUserId(v int64) *SparkAnalyzeLogTask {
	s.UserId = &v
	return s
}

type SparkAppInfo struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Detail      *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkAppInfo) String() string {
	return tea.Prettify(s)
}

func (s SparkAppInfo) GoString() string {
	return s.String()
}

func (s *SparkAppInfo) SetAppId(v string) *SparkAppInfo {
	s.AppId = &v
	return s
}

func (s *SparkAppInfo) SetAppName(v string) *SparkAppInfo {
	s.AppName = &v
	return s
}

func (s *SparkAppInfo) SetDBClusterId(v string) *SparkAppInfo {
	s.DBClusterId = &v
	return s
}

func (s *SparkAppInfo) SetDetail(v *Detail) *SparkAppInfo {
	s.Detail = v
	return s
}

func (s *SparkAppInfo) SetMessage(v string) *SparkAppInfo {
	s.Message = &v
	return s
}

func (s *SparkAppInfo) SetPriority(v string) *SparkAppInfo {
	s.Priority = &v
	return s
}

func (s *SparkAppInfo) SetState(v string) *SparkAppInfo {
	s.State = &v
	return s
}

type SparkAttemptInfo struct {
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	Detail    *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Priority  *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkAttemptInfo) String() string {
	return tea.Prettify(s)
}

func (s SparkAttemptInfo) GoString() string {
	return s.String()
}

func (s *SparkAttemptInfo) SetAttemptId(v string) *SparkAttemptInfo {
	s.AttemptId = &v
	return s
}

func (s *SparkAttemptInfo) SetDetail(v *Detail) *SparkAttemptInfo {
	s.Detail = v
	return s
}

func (s *SparkAttemptInfo) SetMessage(v string) *SparkAttemptInfo {
	s.Message = &v
	return s
}

func (s *SparkAttemptInfo) SetPriority(v string) *SparkAttemptInfo {
	s.Priority = &v
	return s
}

func (s *SparkAttemptInfo) SetState(v string) *SparkAttemptInfo {
	s.State = &v
	return s
}

type SparkOperatorInfo struct {
	MetricValue  *int64 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	OperatorName []byte `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s SparkOperatorInfo) String() string {
	return tea.Prettify(s)
}

func (s SparkOperatorInfo) GoString() string {
	return s.String()
}

func (s *SparkOperatorInfo) SetMetricValue(v int64) *SparkOperatorInfo {
	s.MetricValue = &v
	return s
}

func (s *SparkOperatorInfo) SetOperatorName(v []byte) *SparkOperatorInfo {
	s.OperatorName = v
	return s
}

type SparkSession struct {
	Active    *string `json:"Active,omitempty" xml:"Active,omitempty"`
	AliyunUid *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	SessionId *int64  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkSession) String() string {
	return tea.Prettify(s)
}

func (s SparkSession) GoString() string {
	return s.String()
}

func (s *SparkSession) SetActive(v string) *SparkSession {
	s.Active = &v
	return s
}

func (s *SparkSession) SetAliyunUid(v int64) *SparkSession {
	s.AliyunUid = &v
	return s
}

func (s *SparkSession) SetSessionId(v int64) *SparkSession {
	s.SessionId = &v
	return s
}

func (s *SparkSession) SetState(v string) *SparkSession {
	s.State = &v
	return s
}

type Statement struct {
	AliyunUid     *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeState     *string `json:"CodeState,omitempty" xml:"CodeState,omitempty"`
	CodeType      *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Error         *string `json:"Error,omitempty" xml:"Error,omitempty"`
	HaveRows      *bool   `json:"HaveRows,omitempty" xml:"HaveRows,omitempty"`
	Output        *string `json:"Output,omitempty" xml:"Output,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	SessionId     *int64  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatementId   *int64  `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
	TotalCount    *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s Statement) String() string {
	return tea.Prettify(s)
}

func (s Statement) GoString() string {
	return s.String()
}

func (s *Statement) SetAliyunUid(v int64) *Statement {
	s.AliyunUid = &v
	return s
}

func (s *Statement) SetCode(v string) *Statement {
	s.Code = &v
	return s
}

func (s *Statement) SetCodeState(v string) *Statement {
	s.CodeState = &v
	return s
}

func (s *Statement) SetCodeType(v string) *Statement {
	s.CodeType = &v
	return s
}

func (s *Statement) SetEndTime(v int64) *Statement {
	s.EndTime = &v
	return s
}

func (s *Statement) SetError(v string) *Statement {
	s.Error = &v
	return s
}

func (s *Statement) SetHaveRows(v bool) *Statement {
	s.HaveRows = &v
	return s
}

func (s *Statement) SetOutput(v string) *Statement {
	s.Output = &v
	return s
}

func (s *Statement) SetResourceGroup(v string) *Statement {
	s.ResourceGroup = &v
	return s
}

func (s *Statement) SetSessionId(v int64) *Statement {
	s.SessionId = &v
	return s
}

func (s *Statement) SetStartTime(v int64) *Statement {
	s.StartTime = &v
	return s
}

func (s *Statement) SetStatementId(v int64) *Statement {
	s.StatementId = &v
	return s
}

func (s *Statement) SetTotalCount(v int64) *Statement {
	s.TotalCount = &v
	return s
}

type StatementInfo struct {
	Code                 *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	CompletedTimeInMills *int64   `json:"CompletedTimeInMills,omitempty" xml:"CompletedTimeInMills,omitempty"`
	Output               *string  `json:"Output,omitempty" xml:"Output,omitempty"`
	Process              *float32 `json:"Process,omitempty" xml:"Process,omitempty"`
	StartedTimeInMills   *int64   `json:"StartedTimeInMills,omitempty" xml:"StartedTimeInMills,omitempty"`
	State                *string  `json:"State,omitempty" xml:"State,omitempty"`
	StatementId          *string  `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s StatementInfo) String() string {
	return tea.Prettify(s)
}

func (s StatementInfo) GoString() string {
	return s.String()
}

func (s *StatementInfo) SetCode(v string) *StatementInfo {
	s.Code = &v
	return s
}

func (s *StatementInfo) SetCompletedTimeInMills(v int64) *StatementInfo {
	s.CompletedTimeInMills = &v
	return s
}

func (s *StatementInfo) SetOutput(v string) *StatementInfo {
	s.Output = &v
	return s
}

func (s *StatementInfo) SetProcess(v float32) *StatementInfo {
	s.Process = &v
	return s
}

func (s *StatementInfo) SetStartedTimeInMills(v int64) *StatementInfo {
	s.StartedTimeInMills = &v
	return s
}

func (s *StatementInfo) SetState(v string) *StatementInfo {
	s.State = &v
	return s
}

func (s *StatementInfo) SetStatementId(v string) *StatementInfo {
	s.StatementId = &v
	return s
}

type StorageDescriptorModel struct {
	Compressed             *bool              `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	InputFormat            *string            `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	Location               *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	NumBuckets             *int64             `json:"NumBuckets,omitempty" xml:"NumBuckets,omitempty"`
	OutputFormat           *string            `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters             map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SdId                   *int64             `json:"SdId,omitempty" xml:"SdId,omitempty"`
	SerDeInfo              *SerDeInfoModel    `json:"SerDeInfo,omitempty" xml:"SerDeInfo,omitempty"`
	StoredAsSubDirectories *bool              `json:"StoredAsSubDirectories,omitempty" xml:"StoredAsSubDirectories,omitempty"`
}

func (s StorageDescriptorModel) String() string {
	return tea.Prettify(s)
}

func (s StorageDescriptorModel) GoString() string {
	return s.String()
}

func (s *StorageDescriptorModel) SetCompressed(v bool) *StorageDescriptorModel {
	s.Compressed = &v
	return s
}

func (s *StorageDescriptorModel) SetInputFormat(v string) *StorageDescriptorModel {
	s.InputFormat = &v
	return s
}

func (s *StorageDescriptorModel) SetLocation(v string) *StorageDescriptorModel {
	s.Location = &v
	return s
}

func (s *StorageDescriptorModel) SetNumBuckets(v int64) *StorageDescriptorModel {
	s.NumBuckets = &v
	return s
}

func (s *StorageDescriptorModel) SetOutputFormat(v string) *StorageDescriptorModel {
	s.OutputFormat = &v
	return s
}

func (s *StorageDescriptorModel) SetParameters(v map[string]*string) *StorageDescriptorModel {
	s.Parameters = v
	return s
}

func (s *StorageDescriptorModel) SetSdId(v int64) *StorageDescriptorModel {
	s.SdId = &v
	return s
}

func (s *StorageDescriptorModel) SetSerDeInfo(v *SerDeInfoModel) *StorageDescriptorModel {
	s.SerDeInfo = v
	return s
}

func (s *StorageDescriptorModel) SetStoredAsSubDirectories(v bool) *StorageDescriptorModel {
	s.StoredAsSubDirectories = &v
	return s
}

type TableDetailModel struct {
	Catalog     *string           `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	Columns     []*ColDetailModel `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	CreateTime  *string           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string           `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string           `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SchemaName  *string           `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType   *string           `json:"TableType,omitempty" xml:"TableType,omitempty"`
	UpdateTime  *string           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TableDetailModel) String() string {
	return tea.Prettify(s)
}

func (s TableDetailModel) GoString() string {
	return s.String()
}

func (s *TableDetailModel) SetCatalog(v string) *TableDetailModel {
	s.Catalog = &v
	return s
}

func (s *TableDetailModel) SetColumns(v []*ColDetailModel) *TableDetailModel {
	s.Columns = v
	return s
}

func (s *TableDetailModel) SetCreateTime(v string) *TableDetailModel {
	s.CreateTime = &v
	return s
}

func (s *TableDetailModel) SetDescription(v string) *TableDetailModel {
	s.Description = &v
	return s
}

func (s *TableDetailModel) SetOwner(v string) *TableDetailModel {
	s.Owner = &v
	return s
}

func (s *TableDetailModel) SetSchemaName(v string) *TableDetailModel {
	s.SchemaName = &v
	return s
}

func (s *TableDetailModel) SetTableName(v string) *TableDetailModel {
	s.TableName = &v
	return s
}

func (s *TableDetailModel) SetTableType(v string) *TableDetailModel {
	s.TableType = &v
	return s
}

func (s *TableDetailModel) SetUpdateTime(v string) *TableDetailModel {
	s.UpdateTime = &v
	return s
}

type TableModel struct {
	ArchiveType          *string                 `json:"ArchiveType,omitempty" xml:"ArchiveType,omitempty"`
	BlockSize            *int64                  `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	Bucket               *int64                  `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BucketCount          *int64                  `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	Cols                 []*FieldSchemaModel     `json:"Cols,omitempty" xml:"Cols,omitempty" type:"Repeated"`
	Comment              *string                 `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Compression          *string                 `json:"Compression,omitempty" xml:"Compression,omitempty"`
	CreateTime           *string                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentVersion       *int64                  `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	DbName               *string                 `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DictEncode           *bool                   `json:"DictEncode,omitempty" xml:"DictEncode,omitempty"`
	DistributeColumns    []*FieldSchemaModel     `json:"DistributeColumns,omitempty" xml:"DistributeColumns,omitempty" type:"Repeated"`
	DistributeType       *string                 `json:"DistributeType,omitempty" xml:"DistributeType,omitempty"`
	EnableDfs            *bool                   `json:"EnableDfs,omitempty" xml:"EnableDfs,omitempty"`
	HotPartitionCount    *int64                  `json:"HotPartitionCount,omitempty" xml:"HotPartitionCount,omitempty"`
	Indexes              []*CstoreIndexModel     `json:"Indexes,omitempty" xml:"Indexes,omitempty" type:"Repeated"`
	IsAllIndex           *bool                   `json:"IsAllIndex,omitempty" xml:"IsAllIndex,omitempty"`
	IsFulltextDict       *bool                   `json:"IsFulltextDict,omitempty" xml:"IsFulltextDict,omitempty"`
	MaxColumnId          *int64                  `json:"MaxColumnId,omitempty" xml:"MaxColumnId,omitempty"`
	Parameters           map[string]*string      `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionColumn      *string                 `json:"PartitionColumn,omitempty" xml:"PartitionColumn,omitempty"`
	PartitionCount       *int64                  `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	PartitionKeys        []*FieldSchemaModel     `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	PartitionType        *string                 `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	PhysicalDatabaseName *string                 `json:"PhysicalDatabaseName,omitempty" xml:"PhysicalDatabaseName,omitempty"`
	PhysicalTableName    *string                 `json:"PhysicalTableName,omitempty" xml:"PhysicalTableName,omitempty"`
	PreviousVersion      *int64                  `json:"PreviousVersion,omitempty" xml:"PreviousVersion,omitempty"`
	RawTableName         *string                 `json:"RawTableName,omitempty" xml:"RawTableName,omitempty"`
	RouteColumns         []*FieldSchemaModel     `json:"RouteColumns,omitempty" xml:"RouteColumns,omitempty" type:"Repeated"`
	RouteEffectiveColumn *FieldSchemaModel       `json:"RouteEffectiveColumn,omitempty" xml:"RouteEffectiveColumn,omitempty"`
	RouteType            *string                 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	RtEngineType         *string                 `json:"RtEngineType,omitempty" xml:"RtEngineType,omitempty"`
	RtIndexAll           *bool                   `json:"RtIndexAll,omitempty" xml:"RtIndexAll,omitempty"`
	RtModeType           *string                 `json:"RtModeType,omitempty" xml:"RtModeType,omitempty"`
	Sd                   *StorageDescriptorModel `json:"Sd,omitempty" xml:"Sd,omitempty"`
	StoragePolicy        *string                 `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	SubpartitionColumn   *string                 `json:"SubpartitionColumn,omitempty" xml:"SubpartitionColumn,omitempty"`
	SubpartitionCount    *int64                  `json:"SubpartitionCount,omitempty" xml:"SubpartitionCount,omitempty"`
	SubpartitionType     *string                 `json:"SubpartitionType,omitempty" xml:"SubpartitionType,omitempty"`
	TableEngineName      *string                 `json:"TableEngineName,omitempty" xml:"TableEngineName,omitempty"`
	TableName            *string                 `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType            *string                 `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TblId                *int64                  `json:"TblId,omitempty" xml:"TblId,omitempty"`
	Temporary            *bool                   `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	UpdateTime           *string                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ViewExpandedText     *string                 `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText     *string                 `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
	ViewSecurityMode     *string                 `json:"ViewSecurityMode,omitempty" xml:"ViewSecurityMode,omitempty"`
}

func (s TableModel) String() string {
	return tea.Prettify(s)
}

func (s TableModel) GoString() string {
	return s.String()
}

func (s *TableModel) SetArchiveType(v string) *TableModel {
	s.ArchiveType = &v
	return s
}

func (s *TableModel) SetBlockSize(v int64) *TableModel {
	s.BlockSize = &v
	return s
}

func (s *TableModel) SetBucket(v int64) *TableModel {
	s.Bucket = &v
	return s
}

func (s *TableModel) SetBucketCount(v int64) *TableModel {
	s.BucketCount = &v
	return s
}

func (s *TableModel) SetCols(v []*FieldSchemaModel) *TableModel {
	s.Cols = v
	return s
}

func (s *TableModel) SetComment(v string) *TableModel {
	s.Comment = &v
	return s
}

func (s *TableModel) SetCompression(v string) *TableModel {
	s.Compression = &v
	return s
}

func (s *TableModel) SetCreateTime(v string) *TableModel {
	s.CreateTime = &v
	return s
}

func (s *TableModel) SetCurrentVersion(v int64) *TableModel {
	s.CurrentVersion = &v
	return s
}

func (s *TableModel) SetDbName(v string) *TableModel {
	s.DbName = &v
	return s
}

func (s *TableModel) SetDictEncode(v bool) *TableModel {
	s.DictEncode = &v
	return s
}

func (s *TableModel) SetDistributeColumns(v []*FieldSchemaModel) *TableModel {
	s.DistributeColumns = v
	return s
}

func (s *TableModel) SetDistributeType(v string) *TableModel {
	s.DistributeType = &v
	return s
}

func (s *TableModel) SetEnableDfs(v bool) *TableModel {
	s.EnableDfs = &v
	return s
}

func (s *TableModel) SetHotPartitionCount(v int64) *TableModel {
	s.HotPartitionCount = &v
	return s
}

func (s *TableModel) SetIndexes(v []*CstoreIndexModel) *TableModel {
	s.Indexes = v
	return s
}

func (s *TableModel) SetIsAllIndex(v bool) *TableModel {
	s.IsAllIndex = &v
	return s
}

func (s *TableModel) SetIsFulltextDict(v bool) *TableModel {
	s.IsFulltextDict = &v
	return s
}

func (s *TableModel) SetMaxColumnId(v int64) *TableModel {
	s.MaxColumnId = &v
	return s
}

func (s *TableModel) SetParameters(v map[string]*string) *TableModel {
	s.Parameters = v
	return s
}

func (s *TableModel) SetPartitionColumn(v string) *TableModel {
	s.PartitionColumn = &v
	return s
}

func (s *TableModel) SetPartitionCount(v int64) *TableModel {
	s.PartitionCount = &v
	return s
}

func (s *TableModel) SetPartitionKeys(v []*FieldSchemaModel) *TableModel {
	s.PartitionKeys = v
	return s
}

func (s *TableModel) SetPartitionType(v string) *TableModel {
	s.PartitionType = &v
	return s
}

func (s *TableModel) SetPhysicalDatabaseName(v string) *TableModel {
	s.PhysicalDatabaseName = &v
	return s
}

func (s *TableModel) SetPhysicalTableName(v string) *TableModel {
	s.PhysicalTableName = &v
	return s
}

func (s *TableModel) SetPreviousVersion(v int64) *TableModel {
	s.PreviousVersion = &v
	return s
}

func (s *TableModel) SetRawTableName(v string) *TableModel {
	s.RawTableName = &v
	return s
}

func (s *TableModel) SetRouteColumns(v []*FieldSchemaModel) *TableModel {
	s.RouteColumns = v
	return s
}

func (s *TableModel) SetRouteEffectiveColumn(v *FieldSchemaModel) *TableModel {
	s.RouteEffectiveColumn = v
	return s
}

func (s *TableModel) SetRouteType(v string) *TableModel {
	s.RouteType = &v
	return s
}

func (s *TableModel) SetRtEngineType(v string) *TableModel {
	s.RtEngineType = &v
	return s
}

func (s *TableModel) SetRtIndexAll(v bool) *TableModel {
	s.RtIndexAll = &v
	return s
}

func (s *TableModel) SetRtModeType(v string) *TableModel {
	s.RtModeType = &v
	return s
}

func (s *TableModel) SetSd(v *StorageDescriptorModel) *TableModel {
	s.Sd = v
	return s
}

func (s *TableModel) SetStoragePolicy(v string) *TableModel {
	s.StoragePolicy = &v
	return s
}

func (s *TableModel) SetSubpartitionColumn(v string) *TableModel {
	s.SubpartitionColumn = &v
	return s
}

func (s *TableModel) SetSubpartitionCount(v int64) *TableModel {
	s.SubpartitionCount = &v
	return s
}

func (s *TableModel) SetSubpartitionType(v string) *TableModel {
	s.SubpartitionType = &v
	return s
}

func (s *TableModel) SetTableEngineName(v string) *TableModel {
	s.TableEngineName = &v
	return s
}

func (s *TableModel) SetTableName(v string) *TableModel {
	s.TableName = &v
	return s
}

func (s *TableModel) SetTableType(v string) *TableModel {
	s.TableType = &v
	return s
}

func (s *TableModel) SetTblId(v int64) *TableModel {
	s.TblId = &v
	return s
}

func (s *TableModel) SetTemporary(v bool) *TableModel {
	s.Temporary = &v
	return s
}

func (s *TableModel) SetUpdateTime(v string) *TableModel {
	s.UpdateTime = &v
	return s
}

func (s *TableModel) SetViewExpandedText(v string) *TableModel {
	s.ViewExpandedText = &v
	return s
}

func (s *TableModel) SetViewOriginalText(v string) *TableModel {
	s.ViewOriginalText = &v
	return s
}

func (s *TableModel) SetViewSecurityMode(v string) *TableModel {
	s.ViewSecurityMode = &v
	return s
}

type TableSummaryModel struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SQL         *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSize   *int64  `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	TableType   *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TableSummaryModel) String() string {
	return tea.Prettify(s)
}

func (s TableSummaryModel) GoString() string {
	return s.String()
}

func (s *TableSummaryModel) SetCreateTime(v string) *TableSummaryModel {
	s.CreateTime = &v
	return s
}

func (s *TableSummaryModel) SetDescription(v string) *TableSummaryModel {
	s.Description = &v
	return s
}

func (s *TableSummaryModel) SetOwner(v string) *TableSummaryModel {
	s.Owner = &v
	return s
}

func (s *TableSummaryModel) SetSQL(v string) *TableSummaryModel {
	s.SQL = &v
	return s
}

func (s *TableSummaryModel) SetSchemaName(v string) *TableSummaryModel {
	s.SchemaName = &v
	return s
}

func (s *TableSummaryModel) SetTableName(v string) *TableSummaryModel {
	s.TableName = &v
	return s
}

func (s *TableSummaryModel) SetTableSize(v int64) *TableSummaryModel {
	s.TableSize = &v
	return s
}

func (s *TableSummaryModel) SetTableType(v string) *TableSummaryModel {
	s.TableType = &v
	return s
}

func (s *TableSummaryModel) SetUpdateTime(v string) *TableSummaryModel {
	s.UpdateTime = &v
	return s
}

type AllocateClusterPublicConnectionRequest struct {
	// The prefix of the public endpoint.
	//
	// *   The prefix can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	// *   The prefix can be up to 30 characters in length.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

type AllocateClusterPublicConnectionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponseBody) SetRequestId(v string) *AllocateClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateClusterPublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetStatusCode(v int32) *AllocateClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type AttachUserENIRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s AttachUserENIRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIRequest) GoString() string {
	return s.String()
}

func (s *AttachUserENIRequest) SetDBClusterId(v string) *AttachUserENIRequest {
	s.DBClusterId = &v
	return s
}

type AttachUserENIResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachUserENIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponseBody) SetRequestId(v string) *AttachUserENIResponseBody {
	s.RequestId = &v
	return s
}

type AttachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachUserENIResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIResponse) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponse) SetHeaders(v map[string]*string) *AttachUserENIResponse {
	s.Headers = v
	return s
}

func (s *AttachUserENIResponse) SetStatusCode(v int32) *AttachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachUserENIResponse) SetBody(v *AttachUserENIResponseBody) *AttachUserENIResponse {
	s.Body = v
	return s
}

type BindAccountRequest struct {
	// The standard account of the cluster.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the RAM user.
	RamUser *string `json:"RamUser,omitempty" xml:"RamUser,omitempty"`
}

func (s BindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAccountRequest) GoString() string {
	return s.String()
}

func (s *BindAccountRequest) SetAccountName(v string) *BindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *BindAccountRequest) SetDBClusterId(v string) *BindAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindAccountRequest) SetRamUser(v string) *BindAccountRequest {
	s.RamUser = &v
	return s
}

type BindAccountResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBody) SetRequestId(v string) *BindAccountResponseBody {
	s.RequestId = &v
	return s
}

type BindAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAccountResponse) GoString() string {
	return s.String()
}

func (s *BindAccountResponse) SetHeaders(v map[string]*string) *BindAccountResponse {
	s.Headers = v
	return s
}

func (s *BindAccountResponse) SetStatusCode(v int32) *BindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAccountResponse) SetBody(v *BindAccountResponseBody) *BindAccountResponse {
	s.Body = v
	return s
}

type BindDBResourceGroupWithUserRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the database account. It can be a standard account or a privileged account.
	GroupUser *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
}

func (s BindDBResourceGroupWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupName(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupUser(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

type BindDBResourceGroupWithUserResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourceGroupWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *BindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

type BindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDBResourceGroupWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *BindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetBody(v *BindDBResourceGroupWithUserResponseBody) *BindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

type CheckBindRamUserRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckBindRamUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckBindRamUserRequest) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserRequest) SetDBClusterId(v string) *CheckBindRamUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckBindRamUserRequest) SetRegionId(v string) *CheckBindRamUserRequest {
	s.RegionId = &v
	return s
}

type CheckBindRamUserResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result of the request. Valid values:
	//
	// *   **true**: the database account is associated with a RAM user.
	// *   **false**: the database account is not associated with a RAM user.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckBindRamUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckBindRamUserResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserResponseBody) SetRequestId(v string) *CheckBindRamUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBindRamUserResponseBody) SetResult(v bool) *CheckBindRamUserResponseBody {
	s.Result = &v
	return s
}

type CheckBindRamUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBindRamUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBindRamUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckBindRamUserResponse) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserResponse) SetHeaders(v map[string]*string) *CheckBindRamUserResponse {
	s.Headers = v
	return s
}

func (s *CheckBindRamUserResponse) SetStatusCode(v int32) *CheckBindRamUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBindRamUserResponse) SetBody(v *CheckBindRamUserResponseBody) *CheckBindRamUserResponse {
	s.Body = v
	return s
}

type CheckSampleDataSetRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CheckSampleDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSampleDataSetRequest) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetRequest) SetDBClusterId(v string) *CheckSampleDataSetRequest {
	s.DBClusterId = &v
	return s
}

type CheckSampleDataSetResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the built-in dataset. Valid values:
	//
	// *   **SUCCEED**: The dataset is loaded.
	// *   **INIT**: The dataset is being loaded.
	// *   **FAILED**: The dataset failed to be loaded.
	// *   **UNINITIALIZED**: The dataset is not loaded.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckSampleDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSampleDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetResponseBody) SetRequestId(v string) *CheckSampleDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSampleDataSetResponseBody) SetStatus(v string) *CheckSampleDataSetResponseBody {
	s.Status = &v
	return s
}

type CheckSampleDataSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSampleDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSampleDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSampleDataSetResponse) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetResponse) SetHeaders(v map[string]*string) *CheckSampleDataSetResponse {
	s.Headers = v
	return s
}

func (s *CheckSampleDataSetResponse) SetStatusCode(v int32) *CheckSampleDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSampleDataSetResponse) SetBody(v *CheckSampleDataSetResponseBody) *CheckSampleDataSetResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// The description of the account.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description can be up to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// *   The name must start with a lowercase letter and end with a lowercase letter or a digit.
	// *   The name can contain lowercase letters, digits, and underscores (\_).
	// *   The name must be 2 to 16 characters in length.
	// *   Reserved account names such as root, admin, and opsadmin cannot be used.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// *   The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	// *   Special characters include `! @ # $ % ^ & * ( ) _ + - =`
	// *   The password must be 8 to 32 characters in length.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// *   **Normal**: standard account.
	// *   **Super**: privileged account.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

type CreateAccountResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateDBClusterRequest struct {
	// The ID of the backup set that you want to use to restore data.
	//
	// >  You can call the [DescribeBackups](~~612318~~) operation to query the backup sets of the cluster.
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The amount of reserved computing resources. Unit: ACUs. Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Only **VPC** is supported.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The version of the cluster. Set the value to **5.0**.
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// Specifies whether to allocate all reserved computing resources to the user_default resource group. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	EnableDefaultResourcePool *bool `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// *   **Year**: subscription on a yearly basis.
	// *   **Month**: subscription on a monthly basis.
	//
	// >  This parameter must be specified when PayType is set to Prepaid.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The point in time to which you want to restore data from the backup set.
	RestoreToTime *string `json:"RestoreToTime,omitempty" xml:"RestoreToTime,omitempty"`
	// The method that you want to use to restore data. Valid values:
	//
	// *   **backup**: restores data from a backup set. You must also specify the **BackupSetId** and **SourceDBClusterId** parameters.
	// *   **timepoint**: restores data to a point in time. You must also specify the **RestoreToTime** and **SourceDBClusterId** parameters.
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster. If you want to restore a Data Lakehouse Edition cluster from a Data Warehouse Edition cluster, you must specify this parameter.
	SourceDbClusterId *string `json:"SourceDbClusterId,omitempty" xml:"SourceDbClusterId,omitempty"`
	// The amount of reserved storage resources. Unit: AnalyticDB compute units (ACUs). Valid values: 0 to 2064. The value must be in increments of 24 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The tags to add to the cluster.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration of the subscription cluster.
	//
	// *   Valid values when **Period** is set to Year: 1 to 3 (integer).
	// *   Valid values when **Period** is set to Month: 1 to 9 (integer).
	//
	// >  This parameter must be specified when PayType is set to **Prepaid**.
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetBackupSetId(v string) *CreateDBClusterRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterNetworkType(v string) *CreateDBClusterRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterVersion(v string) *CreateDBClusterRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetEnableDefaultResourcePool(v bool) *CreateDBClusterRequest {
	s.EnableDefaultResourcePool = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreToTime(v string) *CreateDBClusterRequest {
	s.RestoreToTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreType(v string) *CreateDBClusterRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceDbClusterId(v string) *CreateDBClusterRequest {
	s.SourceDbClusterId = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageResource(v string) *CreateDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateDBClusterRequestTag struct {
	// The key of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag value can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBClusterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequestTag) SetKey(v string) *CreateDBClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBClusterRequestTag) SetValue(v string) *CreateDBClusterRequestTag {
	s.Value = &v
	return s
}

type CreateDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The default resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

type CreateDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetStatusCode(v int32) *CreateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

type CreateDBResourceGroupRequest struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EnableSpot  *bool   `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The name of the resource group.
	//
	// *   The name can be up to 255 characters in length.
	// *   The name must start with a letter or a digit.
	// *   The name can contain letters, digits, hyphens (\_), and underscores (\_).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum reserved computing resources. Unit: ACU.
	//
	// *   If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16 ACUs.
	// *   If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8 ACUs.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum reserved computing resources. Unit: AnalyticDB Compute Units (ACUs).
	//
	// *   When GroupType is set to Interactive, set this parameter to 16 ACUs.
	// *   When GroupType is set to Job, set this parameter to 0 ACUs.
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~612393~~) operation to query the most recent region list.
	RegionId *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rules    []*CreateDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) SetClusterMode(v string) *CreateDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetDBClusterId(v string) *CreateDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEnableSpot(v bool) *CreateDBResourceGroupRequest {
	s.EnableSpot = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupName(v string) *CreateDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupType(v string) *CreateDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetRegionId(v string) *CreateDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetRules(v []*CreateDBResourceGroupRequestRules) *CreateDBResourceGroupRequest {
	s.Rules = v
	return s
}

type CreateDBResourceGroupRequestRules struct {
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	QueryTime       *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s CreateDBResourceGroupRequestRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupRequestRules) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestRules) SetGroupName(v string) *CreateDBResourceGroupRequestRules {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequestRules) SetQueryTime(v string) *CreateDBResourceGroupRequestRules {
	s.QueryTime = &v
	return s
}

func (s *CreateDBResourceGroupRequestRules) SetTargetGroupName(v string) *CreateDBResourceGroupRequestRules {
	s.TargetGroupName = &v
	return s
}

type CreateDBResourceGroupShrinkRequest struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EnableSpot  *bool   `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The name of the resource group.
	//
	// *   The name can be up to 255 characters in length.
	// *   The name must start with a letter or a digit.
	// *   The name can contain letters, digits, hyphens (\_), and underscores (\_).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum reserved computing resources. Unit: ACU.
	//
	// *   If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16 ACUs.
	// *   If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8 ACUs.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum reserved computing resources. Unit: AnalyticDB Compute Units (ACUs).
	//
	// *   When GroupType is set to Interactive, set this parameter to 16 ACUs.
	// *   When GroupType is set to Job, set this parameter to 0 ACUs.
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~612393~~) operation to query the most recent region list.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateDBResourceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterMode(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetDBClusterId(v string) *CreateDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEnableSpot(v bool) *CreateDBResourceGroupShrinkRequest {
	s.EnableSpot = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupName(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupType(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetRegionId(v string) *CreateDBResourceGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetRulesShrink(v string) *CreateDBResourceGroupShrinkRequest {
	s.RulesShrink = &v
	return s
}

type CreateDBResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponseBody) SetRequestId(v string) *CreateDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourceGroupResponse) SetStatusCode(v int32) *CreateDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResourceGroupResponse) SetBody(v *CreateDBResourceGroupResponseBody) *CreateDBResourceGroupResponse {
	s.Body = v
	return s
}

type CreateElasticPlanRequest struct {
	// Specifies whether to enable **Default Proportional Scaling for EIUs**. Valid values:
	//
	// *   true. In this case, storage resources are scaled along with computing resources, and the TargetSize and CronExpression parameters are not supported.
	// *   false
	//
	// >
	//
	// *   This parameter must be specified when Type is set to WORKER. This parameter is not required when Type is set to EXECUTOR.
	//
	// *   You can enable Default Proportional Scaling for EIUs for only a single scaling plan of a cluster.
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// A CORN expression that specifies the scaling cycle and time for the scaling plan.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  The name must be 2 to 30 characters in length and can contain letters, digits, and underscores (\_). The name must start with a letter.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Specifies whether to immediately enable the scaling plan after the plan is created. Valid values:
	//
	// *   true
	// *   false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The end time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	//
	// >
	//
	// *   If you want to create a scaling plan that uses interactive resource groups, you must specify this parameter. If you want to create a scaling plan that uses elastic I/O units (EIUs), you do not need to specify this parameter.
	//
	// *   You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the resource group name for a cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The desired specifications of elastic resources after scaling.
	//
	// >
	//
	// *   If the scaling plan uses **EIUs** and **Default Proportional Scaling for EIUs** is enabled, you do not need to specify this parameter. In other cases, you must specify this parameter.
	//
	// *   You can call the [DescribeElasticPlanSpecifications](~~601278~~) operation to query the specifications that are supported for scaling plans.
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   EXECUTOR: the interactive resource group type, which indicates the computing resource type.
	// *   WORKER: the EIU type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) SetAutoScale(v bool) *CreateElasticPlanRequest {
	s.AutoScale = &v
	return s
}

func (s *CreateElasticPlanRequest) SetCronExpression(v string) *CreateElasticPlanRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetEnabled(v bool) *CreateElasticPlanRequest {
	s.Enabled = &v
	return s
}

func (s *CreateElasticPlanRequest) SetEndTime(v string) *CreateElasticPlanRequest {
	s.EndTime = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceGroupName(v string) *CreateElasticPlanRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetStartTime(v string) *CreateElasticPlanRequest {
	s.StartTime = &v
	return s
}

func (s *CreateElasticPlanRequest) SetTargetSize(v string) *CreateElasticPlanRequest {
	s.TargetSize = &v
	return s
}

func (s *CreateElasticPlanRequest) SetType(v string) *CreateElasticPlanRequest {
	s.Type = &v
	return s
}

type CreateElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBody) SetRequestId(v string) *CreateElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponse) SetHeaders(v map[string]*string) *CreateElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticPlanResponse) SetStatusCode(v int32) *CreateElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticPlanResponse) SetBody(v *CreateElasticPlanResponseBody) *CreateElasticPlanResponse {
	s.Body = v
	return s
}

type CreateOssSubDirectoryRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The OSS path where you want to create a subdirectory.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateOssSubDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOssSubDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryRequest) SetDBClusterId(v string) *CreateOssSubDirectoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateOssSubDirectoryRequest) SetPath(v string) *CreateOssSubDirectoryRequest {
	s.Path = &v
	return s
}

type CreateOssSubDirectoryResponseBody struct {
	// The returned data.
	Data *CreateOssSubDirectoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response code. The status code 200 indicates that the request was successful.
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// *   If the request was successful, a **success** message is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOssSubDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOssSubDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponseBody) SetData(v *CreateOssSubDirectoryResponseBodyData) *CreateOssSubDirectoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetHttpStatusCode(v int64) *CreateOssSubDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetMessage(v string) *CreateOssSubDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetRequestId(v string) *CreateOssSubDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetSuccess(v bool) *CreateOssSubDirectoryResponseBody {
	s.Success = &v
	return s
}

type CreateOssSubDirectoryResponseBodyData struct {
	// The cyclic redundancy check (CRC) value on the client.
	ClientCRC *int64 `json:"ClientCRC,omitempty" xml:"ClientCRC,omitempty"`
	// The tag of the OSS path.
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CRC-64 value on the OSS bucket.
	ServerCRC *int64 `json:"ServerCRC,omitempty" xml:"ServerCRC,omitempty"`
}

func (s CreateOssSubDirectoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateOssSubDirectoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponseBodyData) SetClientCRC(v int64) *CreateOssSubDirectoryResponseBodyData {
	s.ClientCRC = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetETag(v string) *CreateOssSubDirectoryResponseBodyData {
	s.ETag = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetRequestId(v string) *CreateOssSubDirectoryResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetServerCRC(v int64) *CreateOssSubDirectoryResponseBodyData {
	s.ServerCRC = &v
	return s
}

type CreateOssSubDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssSubDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssSubDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOssSubDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponse) SetHeaders(v map[string]*string) *CreateOssSubDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateOssSubDirectoryResponse) SetStatusCode(v int32) *CreateOssSubDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssSubDirectoryResponse) SetBody(v *CreateOssSubDirectoryResponseBody) *CreateOssSubDirectoryResponse {
	s.Body = v
	return s
}

type CreateSparkTemplateRequest struct {
	// The application type. Valid values:
	//
	// *   **SQL**
	// *   **STREAMING**
	// *   **BATCH**
	//
	// >  You do not need to specify this parameter when Type is set to folder.
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the application template. The name can be up to 64 characters in length.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the directory to which the application template belongs.
	//
	// >  You can call the [GetSparkTemplateFolderTree](~~456218~~) operation to query the directory ID.
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the application template. Valid values:
	//
	// *   **folder**: directory.
	// *   **file**: application.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateSparkTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSparkTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateRequest) SetAppType(v string) *CreateSparkTemplateRequest {
	s.AppType = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetDBClusterId(v string) *CreateSparkTemplateRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetName(v string) *CreateSparkTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetParentId(v int64) *CreateSparkTemplateRequest {
	s.ParentId = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetType(v string) *CreateSparkTemplateRequest {
	s.Type = &v
	return s
}

type CreateSparkTemplateResponseBody struct {
	// The creation result.
	Data *CreateSparkTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSparkTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSparkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponseBody) SetData(v *CreateSparkTemplateResponseBodyData) *CreateSparkTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateSparkTemplateResponseBody) SetRequestId(v string) *CreateSparkTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateSparkTemplateResponseBodyData struct {
	// Indicates whether the application template is created. Valid values:
	//
	// *   **True**
	// *   **False**
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s CreateSparkTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSparkTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponseBodyData) SetSucceeded(v bool) *CreateSparkTemplateResponseBodyData {
	s.Succeeded = &v
	return s
}

type CreateSparkTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSparkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSparkTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSparkTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateResponse) SetHeaders(v map[string]*string) *CreateSparkTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSparkTemplateResponse) SetStatusCode(v int32) *CreateSparkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSparkTemplateResponse) SetBody(v *CreateSparkTemplateResponseBody) *CreateSparkTemplateResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The name of the database account.
	//
	// > You can call the [DescribeAccounts](~~612430~~) operation to query the information about database accounts in a cluster, including the database account name.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

type DeleteAccountResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteDBClusterRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetStatusCode(v int32) *DeleteDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~612410~~) operation to query the information about resource groups of an AnalyticDB for MySQL cluster, including resource group names.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupRequest) SetDBClusterId(v string) *DeleteDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetGroupName(v string) *DeleteDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

type DeleteDBResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponseBody) SetRequestId(v string) *DeleteDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetStatusCode(v int32) *DeleteDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetBody(v *DeleteDBResourceGroupResponseBody) *DeleteDBResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteElasticPlanRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s DeleteElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

type DeleteElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponseBody) SetRequestId(v string) *DeleteElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponse) SetHeaders(v map[string]*string) *DeleteElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticPlanResponse) SetStatusCode(v int32) *DeleteElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticPlanResponse) SetBody(v *DeleteElasticPlanResponseBody) *DeleteElasticPlanResponse {
	s.Body = v
	return s
}

type DeleteProcessInstanceRequest struct {
	// The ID of the Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the workflow instance.
	ProcessInstanceId *int64 `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// The project ID, which is the unique identifier of the project.
	ProjectCode *int64 `json:"ProjectCode,omitempty" xml:"ProjectCode,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProcessInstanceRequest) SetDBClusterId(v string) *DeleteProcessInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteProcessInstanceRequest) SetProcessInstanceId(v int64) *DeleteProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *DeleteProcessInstanceRequest) SetProjectCode(v int64) *DeleteProcessInstanceRequest {
	s.ProjectCode = &v
	return s
}

func (s *DeleteProcessInstanceRequest) SetRegionId(v string) *DeleteProcessInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteProcessInstanceResponseBody struct {
	// Indicates whether the workflow instance is deleted. Valid values:
	//
	// *   **true**
	// *   **false**
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. Valid values:
	//
	// *   If the request was successful, **Success** is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProcessInstanceResponseBody) SetData(v bool) *DeleteProcessInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteProcessInstanceResponseBody) SetMessage(v string) *DeleteProcessInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProcessInstanceResponseBody) SetRequestId(v string) *DeleteProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProcessInstanceResponseBody) SetSuccess(v bool) *DeleteProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteProcessInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProcessInstanceResponse) SetHeaders(v map[string]*string) *DeleteProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProcessInstanceResponse) SetStatusCode(v int32) *DeleteProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProcessInstanceResponse) SetBody(v *DeleteProcessInstanceResponseBody) *DeleteProcessInstanceResponse {
	s.Body = v
	return s
}

type DeleteSparkTemplateRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The directory ID of the template files that you want to delete.
	//
	// >
	//
	// *   You can call the [GetSparkTemplateFullTree](~~456205~~) operation to query the directory ID of template files.
	//
	// *   When you specify a directory ID, the directory and all template files that are included in the directory are deleted.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSparkTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateRequest) SetDBClusterId(v string) *DeleteSparkTemplateRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSparkTemplateRequest) SetId(v int64) *DeleteSparkTemplateRequest {
	s.Id = &v
	return s
}

type DeleteSparkTemplateResponseBody struct {
	// The returned result.
	Data *DeleteSparkTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSparkTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponseBody) SetData(v *DeleteSparkTemplateResponseBodyData) *DeleteSparkTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSparkTemplateResponseBody) SetRequestId(v string) *DeleteSparkTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSparkTemplateResponseBodyData struct {
	// Indicates whether the request was successful. Valid values:
	//
	// *   **True**
	// *   **False**
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s DeleteSparkTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponseBodyData) SetSucceeded(v bool) *DeleteSparkTemplateResponseBodyData {
	s.Succeeded = &v
	return s
}

type DeleteSparkTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSparkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSparkTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponse) SetHeaders(v map[string]*string) *DeleteSparkTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSparkTemplateResponse) SetStatusCode(v int32) *DeleteSparkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSparkTemplateResponse) SetBody(v *DeleteSparkTemplateResponseBody) *DeleteSparkTemplateResponse {
	s.Body = v
	return s
}

type DeleteSparkTemplateFileRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the template file to be deleted.
	//
	// >  You can call the [GetSparkTemplateFullTree](~~456205~~) operation to query all template file IDs.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSparkTemplateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileRequest) SetDBClusterId(v string) *DeleteSparkTemplateFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSparkTemplateFileRequest) SetId(v int64) *DeleteSparkTemplateFileRequest {
	s.Id = &v
	return s
}

type DeleteSparkTemplateFileResponseBody struct {
	// The deletion result.
	Data *DeleteSparkTemplateFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSparkTemplateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponseBody) SetData(v *DeleteSparkTemplateFileResponseBodyData) *DeleteSparkTemplateFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSparkTemplateFileResponseBody) SetRequestId(v string) *DeleteSparkTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSparkTemplateFileResponseBodyData struct {
	// Indicates whether the template file is deleted. Valid values:
	//
	// *   **true**
	// *   **false**
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s DeleteSparkTemplateFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponseBodyData) SetSucceeded(v bool) *DeleteSparkTemplateFileResponseBodyData {
	s.Succeeded = &v
	return s
}

type DeleteSparkTemplateFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSparkTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSparkTemplateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSparkTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponse) SetHeaders(v map[string]*string) *DeleteSparkTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSparkTemplateFileResponse) SetStatusCode(v int32) *DeleteSparkTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSparkTemplateFileResponse) SetBody(v *DeleteSparkTemplateFileResponseBody) *DeleteSparkTemplateFileResponse {
	s.Body = v
	return s
}

type DescribeAccountAllPrivilegesRequest struct {
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies the start position marker from which to return results. If you receive a response indicating that the results are truncated, set this parameter to the value of the `Marker` parameter in the response that you received.
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountAllPrivilegesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesRequest) SetAccountName(v string) *DescribeAccountAllPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetDBClusterId(v string) *DescribeAccountAllPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetMarker(v string) *DescribeAccountAllPrivilegesRequest {
	s.Marker = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetRegionId(v string) *DescribeAccountAllPrivilegesRequest {
	s.RegionId = &v
	return s
}

type DescribeAccountAllPrivilegesResponseBody struct {
	// Details of the permissions.
	Data *DescribeAccountAllPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBody) SetData(v *DescribeAccountAllPrivilegesResponseBodyData) *DescribeAccountAllPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBody) SetRequestId(v string) *DescribeAccountAllPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountAllPrivilegesResponseBodyData struct {
	// Indicates the position where the results are truncated. When a value of `true` is returned for the `Truncated` parameter, this parameter is present and contains the value to use for the Marker parameter in a subsequent call.
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The permissions.
	Result []*DescribeAccountAllPrivilegesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the results are truncated. If the results are truncated, a value of `true` is returned. In this case, you must call this operation again to obtain all the results until a value of `false` is returned for this parameter.
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetMarker(v string) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetResult(v []*DescribeAccountAllPrivilegesResponseBodyDataResult) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetTruncated(v bool) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Truncated = &v
	return s
}

type DescribeAccountAllPrivilegesResponseBodyDataResult struct {
	// The objects on which the permission takes effect, including databases, tables, and columns. If Global is returned for the PrivilegeType parameter, an empty string is returned for this parameter.
	PrivilegeObject *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level of the database account. You can call the `DescribeEnabledPrivileges` operation to query the permission level of the database account.
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The name of the permission, which is the same as the permission name returned by the `DescribeEnabledPrivileges` operation.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivilegeObject(v *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.PrivilegeObject = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivilegeType(v string) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivileges(v []*string) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.Privileges = v
	return s
}

type DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject struct {
	// The name of the column.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description of the permission object.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetColumn(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Column = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetDatabase(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Database = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetDescription(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Description = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetTable(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Table = &v
	return s
}

type DescribeAccountAllPrivilegesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAllPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeAccountAllPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponse) SetStatusCode(v int32) *DescribeAccountAllPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponse) SetBody(v *DescribeAccountAllPrivilegesResponseBody) *DescribeAccountAllPrivilegesResponse {
	s.Body = v
	return s
}

type DescribeAccountPrivilegeObjectsRequest struct {
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The column name that is used to filter columns.
	ColumnPrivilegeObject *string `json:"ColumnPrivilegeObject,omitempty" xml:"ColumnPrivilegeObject,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name that is used to filter databases.
	DatabasePrivilegeObject *string `json:"DatabasePrivilegeObject,omitempty" xml:"DatabasePrivilegeObject,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The permission level. Valid values: Database, Table, and Column. Global is not supported.
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table name that is used to filter tables.
	TablePrivilegeObject *string `json:"TablePrivilegeObject,omitempty" xml:"TablePrivilegeObject,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetAccountName(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.ColumnPrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetDBClusterId(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.DatabasePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPageNumber(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPageSize(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPrivilegeType(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetRegionId(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetTablePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.TablePrivilegeObject = &v
	return s
}

type DescribeAccountPrivilegeObjectsResponseBody struct {
	// The permissions.
	Data []*DescribeAccountPrivilegeObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetData(v []*DescribeAccountPrivilegeObjectsResponseBodyData) *DescribeAccountPrivilegeObjectsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetPageNumber(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetPageSize(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetRequestId(v string) *DescribeAccountPrivilegeObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBody) SetTotalCount(v int64) *DescribeAccountPrivilegeObjectsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccountPrivilegeObjectsResponseBodyData struct {
	// The name of the column. This parameter is returned when PrivilegeType is set to Column.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database. This parameter is returned when PrivilegeType is set to Database, Table, or Column.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description that is specified when you create a table or column. This parameter is returned only when PrivilegeType is set to Database or Table, indicating the database description or table description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table. This parameter is returned when PrivilegeType is set to Table or Column.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetColumn(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Column = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetDatabase(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Database = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetDescription(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponseBodyData) SetTable(v string) *DescribeAccountPrivilegeObjectsResponseBodyData {
	s.Table = &v
	return s
}

type DescribeAccountPrivilegeObjectsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountPrivilegeObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetHeaders(v map[string]*string) *DescribeAccountPrivilegeObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetStatusCode(v int32) *DescribeAccountPrivilegeObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetBody(v *DescribeAccountPrivilegeObjectsResponseBody) *DescribeAccountPrivilegeObjectsResponse {
	s.Body = v
	return s
}

type DescribeAccountPrivilegesRequest struct {
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The columns that you want to query. You can use this parameter to query the permissions of the database account on specific columns. This parameter is available only if the PrivilegeType parameter is set to Column.
	ColumnPrivilegeObject *string `json:"ColumnPrivilegeObject,omitempty" xml:"ColumnPrivilegeObject,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The databases that you want to query. You can use this parameter to query the permissions of the database account on specific databases. This parameter is available only if the PrivilegeType parameter is set to Database, Table, or Column.
	DatabasePrivilegeObject *string `json:"DatabasePrivilegeObject,omitempty" xml:"DatabasePrivilegeObject,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The permission level that you want to query. You can call the `DescribeEnabledPrivileges` operation to query the permission level of the database account.
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tables that you want to query. You can use this parameter to query the permissions of the database account on specific tables. This parameter can be used together with the DatabasePrivilegeObject parameter. This parameter is available only if the PrivilegeType parameter is set to Table or Column.
	TablePrivilegeObject *string `json:"TablePrivilegeObject,omitempty" xml:"TablePrivilegeObject,omitempty"`
}

func (s DescribeAccountPrivilegesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesRequest) SetAccountName(v string) *DescribeAccountPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.ColumnPrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetDBClusterId(v string) *DescribeAccountPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.DatabasePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPageNumber(v string) *DescribeAccountPrivilegesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPageSize(v string) *DescribeAccountPrivilegesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPrivilegeType(v string) *DescribeAccountPrivilegesRequest {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetRegionId(v string) *DescribeAccountPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetTablePrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.TablePrivilegeObject = &v
	return s
}

type DescribeAccountPrivilegesResponseBody struct {
	// Details of the permissions.
	Data []*DescribeAccountPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountPrivilegesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBody) SetData(v []*DescribeAccountPrivilegesResponseBodyData) *DescribeAccountPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetPageNumber(v int64) *DescribeAccountPrivilegesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetPageSize(v int64) *DescribeAccountPrivilegesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetRequestId(v string) *DescribeAccountPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetTotalCount(v int64) *DescribeAccountPrivilegesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccountPrivilegesResponseBodyData struct {
	// The objects on which the permission takes effect, including databases, tables, columns, and additional descriptions.
	PrivilegeObject *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level of the permission. Valid values: `Global`, `Database`, `Table`, and `Column`. You can call the `DescribeEnabledPrivileges` parameter to query the permission level of a specific permission.
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The name of the permission. You can call the `DescribeEnabledPrivileges` operation to query the name of the permission.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountPrivilegesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivilegeObject(v *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) *DescribeAccountPrivilegesResponseBodyData {
	s.PrivilegeObject = v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivilegeType(v string) *DescribeAccountPrivilegesResponseBodyData {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivileges(v []*string) *DescribeAccountPrivilegesResponseBodyData {
	s.Privileges = v
	return s
}

type DescribeAccountPrivilegesResponseBodyDataPrivilegeObject struct {
	// The name of the column.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description of the permission object.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetColumn(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Column = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetDatabase(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Database = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetDescription(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Description = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetTable(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Table = &v
	return s
}

type DescribeAccountPrivilegesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountPrivilegesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeAccountPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountPrivilegesResponse) SetStatusCode(v int32) *DescribeAccountPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountPrivilegesResponse) SetBody(v *DescribeAccountPrivilegesResponseBody) *DescribeAccountPrivilegesResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// The name of the database account.
	//
	// > If you do not specify this parameter, the information about all database accounts in the cluster is returned.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v string) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	// The queried database accounts.
	AccountList *DescribeAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccountList struct {
	DBAccount []*DescribeAccountsResponseBodyAccountListDBAccount `json:"DBAccount,omitempty" xml:"DBAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountList) SetDBAccount(v []*DescribeAccountsResponseBodyAccountListDBAccount) *DescribeAccountsResponseBodyAccountList {
	s.DBAccount = v
	return s
}

type DescribeAccountsResponseBodyAccountListDBAccount struct {
	// The description of the database account.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the database account. Valid values:
	//
	// *   **Creating**
	// *   **Available**
	// *   **Deleting**
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the database account. Valid values:
	//
	// *   **Normal**: standard account.
	// *   **Super**: privileged account.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the RAM user.
	RamUsers *string `json:"RamUsers,omitempty" xml:"RamUsers,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetRamUsers(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.RamUsers = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetStatusCode(v int32) *DescribeAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAdbMySqlColumnsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsRequest) SetDBClusterId(v string) *DescribeAdbMySqlColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetRegionId(v string) *DescribeAdbMySqlColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetSchema(v string) *DescribeAdbMySqlColumnsRequest {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetTableName(v string) *DescribeAdbMySqlColumnsRequest {
	s.TableName = &v
	return s
}

type DescribeAdbMySqlColumnsResponseBody struct {
	// The total number of columns.
	ColumnCount *int32 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// Details of the columns.
	Columns []*DescribeAdbMySqlColumnsResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The message returned for the operation. Valid values:
	//
	// *   **Success** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// *   **true**: The operation is successful.
	// *   **false**: The operation fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetColumnCount(v int32) *DescribeAdbMySqlColumnsResponseBody {
	s.ColumnCount = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetColumns(v []*DescribeAdbMySqlColumnsResponseBodyColumns) *DescribeAdbMySqlColumnsResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetMessage(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetRequestId(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetSchema(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetSuccess(v bool) *DescribeAdbMySqlColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBody) SetTableName(v string) *DescribeAdbMySqlColumnsResponseBody {
	s.TableName = &v
	return s
}

type DescribeAdbMySqlColumnsResponseBodyColumns struct {
	// The comments of the column.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the column.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the column.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetComment(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Comment = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetName(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponseBodyColumns) SetType(v string) *DescribeAdbMySqlColumnsResponseBodyColumns {
	s.Type = &v
	return s
}

type DescribeAdbMySqlColumnsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlColumnsResponse) SetStatusCode(v int32) *DescribeAdbMySqlColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponse) SetBody(v *DescribeAdbMySqlColumnsResponseBody) *DescribeAdbMySqlColumnsResponse {
	s.Body = v
	return s
}

type DescribeAdbMySqlSchemasRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAdbMySqlSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasRequest) SetDBClusterId(v string) *DescribeAdbMySqlSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlSchemasRequest) SetRegionId(v string) *DescribeAdbMySqlSchemasRequest {
	s.RegionId = &v
	return s
}

type DescribeAdbMySqlSchemasResponseBody struct {
	// The returned message.
	//
	// *   If the request was successful, a **success** message is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried databases.
	Schemas []*string `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAdbMySqlSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetMessage(v string) *DescribeAdbMySqlSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetRequestId(v string) *DescribeAdbMySqlSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetSchemas(v []*string) *DescribeAdbMySqlSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetSuccess(v bool) *DescribeAdbMySqlSchemasResponseBody {
	s.Success = &v
	return s
}

type DescribeAdbMySqlSchemasResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlSchemasResponse) SetStatusCode(v int32) *DescribeAdbMySqlSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponse) SetBody(v *DescribeAdbMySqlSchemasResponseBody) *DescribeAdbMySqlSchemasResponse {
	s.Body = v
	return s
}

type DescribeAdbMySqlTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeAdbMySqlTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesRequest) SetDBClusterId(v string) *DescribeAdbMySqlTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlTablesRequest) SetRegionId(v string) *DescribeAdbMySqlTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlTablesRequest) SetSchema(v string) *DescribeAdbMySqlTablesRequest {
	s.Schema = &v
	return s
}

type DescribeAdbMySqlTablesResponseBody struct {
	// The message returned for the operation. Valid values:
	//
	// *   **Success** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// *   **true**: The operation is successful.
	// *   **false**: The operation fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The names of tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeAdbMySqlTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesResponseBody) SetMessage(v string) *DescribeAdbMySqlTablesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetRequestId(v string) *DescribeAdbMySqlTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetSchema(v string) *DescribeAdbMySqlTablesResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetSuccess(v bool) *DescribeAdbMySqlTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponseBody) SetTables(v []*string) *DescribeAdbMySqlTablesResponseBody {
	s.Tables = v
	return s
}

type DescribeAdbMySqlTablesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdbMySqlTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlTablesResponse) SetStatusCode(v int32) *DescribeAdbMySqlTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponse) SetBody(v *DescribeAdbMySqlTablesResponseBody) *DescribeAdbMySqlTablesResponse {
	s.Body = v
	return s
}

type DescribeAllDataSourceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetRegionId(v string) *DescribeAllDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetSchemaName(v string) *DescribeAllDataSourceRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetTableName(v string) *DescribeAllDataSourceRequest {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponseBody struct {
	// The queried columns.
	Columns *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried databases.
	Schemas *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The queried tables.
	Tables *DescribeAllDataSourceResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody {
	s.Tables = v
	return s
}

type DescribeAllDataSourceResponseBodyColumns struct {
	Column []*DescribeAllDataSourceResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumns) SetColumn(v []*DescribeAllDataSourceResponseBodyColumnsColumn) *DescribeAllDataSourceResponseBodyColumns {
	s.Column = v
	return s
}

type DescribeAllDataSourceResponseBodyColumnsColumn struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// *   **true**
	// *   **false**
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The name of the column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// *   **true**
	// *   **false**
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The logical name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The logical name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The data type of the column.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

type DescribeAllDataSourceResponseBodySchemas struct {
	Schema []*DescribeAllDataSourceResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemas) SetSchema(v []*DescribeAllDataSourceResponseBodySchemasSchema) *DescribeAllDataSourceResponseBodySchemas {
	s.Schema = v
	return s
}

type DescribeAllDataSourceResponseBodySchemasSchema struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The logical name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

type DescribeAllDataSourceResponseBodyTables struct {
	Table []*DescribeAllDataSourceResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTables) SetTable(v []*DescribeAllDataSourceResponseBodyTablesTable) *DescribeAllDataSourceResponseBodyTables {
	s.Table = v
	return s
}

type DescribeAllDataSourceResponseBodyTablesTable struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The logical name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyTablesTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourceResponse) SetStatusCode(v int32) *DescribeAllDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDataSourceResponse) SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse {
	s.Body = v
	return s
}

type DescribeApsActionLogsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the logs to be queried. Specify the time in the ISO 8601 standard in the **yyyy-MM-ddTHH:mmZ** format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. Their interval cannot be longer than 30 days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that you want to use for fuzzy match in the query.
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task phase during which the logs to be queried are generated. Valid values:
	//
	// *   **StructureMigrate**: schema migration.
	// *   **FullDataSync**: full data synchronization.
	// *   **IncrementalSync**: incremental data synchronization.
	//
	// >  If you do not specify this parameter, logs of all the task phases are queried.
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The start time of the logs to be queried. Specify the time in the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ** format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the log. Separate multiple log types with commas (,). Valid values:
	//
	// *   **INFO**
	// *   **WARN**
	// *   **ERROR**
	//
	// >  If you do not specify this parameter, logs of all types are queried.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the real-time data ingestion task.
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsActionLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsActionLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsRequest) SetDBClusterId(v string) *DescribeApsActionLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetEndTime(v string) *DescribeApsActionLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetKeyword(v string) *DescribeApsActionLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetOwnerAccount(v string) *DescribeApsActionLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetOwnerId(v int64) *DescribeApsActionLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetPageNumber(v int32) *DescribeApsActionLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetPageSize(v int32) *DescribeApsActionLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetRegionId(v string) *DescribeApsActionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetResourceOwnerAccount(v string) *DescribeApsActionLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetResourceOwnerId(v int64) *DescribeApsActionLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetStage(v string) *DescribeApsActionLogsRequest {
	s.Stage = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetStartTime(v string) *DescribeApsActionLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetState(v string) *DescribeApsActionLogsRequest {
	s.State = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetWorkloadId(v string) *DescribeApsActionLogsRequest {
	s.WorkloadId = &v
	return s
}

type DescribeApsActionLogsResponseBody struct {
	// Details of the logs.
	ActionLogs []*DescribeApsActionLogsResponseBodyActionLogs `json:"ActionLogs,omitempty" xml:"ActionLogs,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the real-time data ingestion task.
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsActionLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsActionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponseBody) SetActionLogs(v []*DescribeApsActionLogsResponseBodyActionLogs) *DescribeApsActionLogsResponseBody {
	s.ActionLogs = v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetDBClusterId(v string) *DescribeApsActionLogsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetPageNumber(v string) *DescribeApsActionLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetPageSize(v string) *DescribeApsActionLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetRequestId(v string) *DescribeApsActionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetTotalCount(v string) *DescribeApsActionLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApsActionLogsResponseBody) SetWorkloadId(v string) *DescribeApsActionLogsResponseBody {
	s.WorkloadId = &v
	return s
}

type DescribeApsActionLogsResponseBodyActionLogs struct {
	// The content of the log.
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The task phase during which the logs are generated. Valid values:
	//
	// *   **StructureMigrate**: schema migration.
	// *   **FullDataSync**: full data synchronization.
	// *   **IncrementalSync**: incremental data synchronization.
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The type of the log. Multiple log types are separated by commas (,). Valid values:
	//
	// *   **INFO**
	// *   **WARN**
	// *   **ERROR**
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the log is generated. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ** format. The time is displayed in UTC.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeApsActionLogsResponseBodyActionLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsActionLogsResponseBodyActionLogs) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetContext(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Context = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetStage(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Stage = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetState(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.State = &v
	return s
}

func (s *DescribeApsActionLogsResponseBodyActionLogs) SetTime(v string) *DescribeApsActionLogsResponseBodyActionLogs {
	s.Time = &v
	return s
}

type DescribeApsActionLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsActionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsActionLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsActionLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponse) SetHeaders(v map[string]*string) *DescribeApsActionLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsActionLogsResponse) SetStatusCode(v int32) *DescribeApsActionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsActionLogsResponse) SetBody(v *DescribeApsActionLogsResponseBody) *DescribeApsActionLogsResponse {
	s.Body = v
	return s
}

type DescribeApsResourceGroupsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WorkloadId  *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsRequest) SetDBClusterId(v string) *DescribeApsResourceGroupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsResourceGroupsRequest) SetRegionId(v string) *DescribeApsResourceGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsResourceGroupsRequest) SetWorkloadId(v string) *DescribeApsResourceGroupsRequest {
	s.WorkloadId = &v
	return s
}

type DescribeApsResourceGroupsResponseBody struct {
	// The queried resource groups.
	Data *DescribeApsResourceGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// *   If the request was successful, a success message is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBody) SetData(v *DescribeApsResourceGroupsResponseBodyData) *DescribeApsResourceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetHttpStatusCode(v int64) *DescribeApsResourceGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetMessage(v string) *DescribeApsResourceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetRequestId(v string) *DescribeApsResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetSuccess(v bool) *DescribeApsResourceGroupsResponseBody {
	s.Success = &v
	return s
}

type DescribeApsResourceGroupsResponseBodyData struct {
	// The queried resource groups.
	ResourceGroups []*DescribeApsResourceGroupsResponseBodyDataResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The step size of resources. Unit: AnalyticDB compute units (ACUs).
	//
	// *   If the value of GroupType is **Interactive**, 16 is returned.
	// *   If the value of GroupType is **Job**, 8 is returned.
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBodyData) SetResourceGroups(v []*DescribeApsResourceGroupsResponseBodyDataResourceGroups) *DescribeApsResourceGroupsResponseBodyData {
	s.ResourceGroups = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyData) SetStep(v int64) *DescribeApsResourceGroupsResponseBodyData {
	s.Step = &v
	return s
}

type DescribeApsResourceGroupsResponseBodyDataResourceGroups struct {
	// Indicates whether the resource group is available. Valid values:
	//
	// *   **true**
	// *   **false**
	Available *bool    `json:"Available,omitempty" xml:"Available,omitempty"`
	CuOptions []*int64 `json:"CuOptions,omitempty" xml:"CuOptions,omitempty" type:"Repeated"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// >  For more information about resource groups, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The amount of remaining computing resources. Unit: ACUs.
	LeftComputeResource *int32 `json:"LeftComputeResource,omitempty" xml:"LeftComputeResource,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACUs.
	//
	// *   If the value of GroupType is **Interactive**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 16 ACUs.
	// *   If the value of GroupType is **Job**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 8 ACUs.
	MaxComputeResource *int32 `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// The minimum amount of reserved computing resources. Unit: ACUs.
	//
	// *   If the value of GroupType is **Interactive**, 16 is returned.
	// *   If the value of GroupType is **Job**, 0 is returned.
	MinComputeResource *int32 `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBodyDataResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBodyDataResourceGroups) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetAvailable(v bool) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.Available = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetCuOptions(v []*int64) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.CuOptions = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetGroupName(v string) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetGroupType(v string) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetLeftComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.LeftComputeResource = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetMaxComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetMinComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.MinComputeResource = &v
	return s
}

type DescribeApsResourceGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApsResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeApsResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsResourceGroupsResponse) SetStatusCode(v int32) *DescribeApsResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsResourceGroupsResponse) SetBody(v *DescribeApsResourceGroupsResponseBody) *DescribeApsResourceGroupsResponse {
	s.Body = v
	return s
}

type DescribeAuditLogRecordsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statement was executed.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The order in which to sort the retrieved entries by field. Specify this parameter in the JSON format. The value is an ordered array that uses the order of the input array and contains `Field` and `Type`. Example: `[{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]`. Fields:
	//
	// *   `Field`: the field that is used to sort the retrieved entries. Valid values:
	//
	//     *   **HostAddress**: the IP address of the client that is used to connect to the database.
	//     *   **UserName**: the username.
	//     *   **ExecutionStartTime**: the start time of the query execution.
	//     *   **QueryTime**: the amount of time consumed to execute the SQL statement.
	//     *   **PeakMemoryUsage**: the maximum memory usage when the SQL statement is executed.
	//     *   **ScanRows**: the number of rows to be scanned from a data source in the task.
	//     *   **ScanSize**: the amount of data to be scanned.
	//     *   **ScanTime**: the total amount of time consumed to scan data.
	//     *   **PlanningTime**: the amount of time consumed to generate execution plans.
	//     *   **WallTime**: the accumulated CPU Time values of all operators in the query on each node.
	//     *   **ProcessID**: the process ID.
	//
	// *   `Type`: the sorting type of the retrieved entries. Valid values:
	//
	//     *   **Desc**: descending order.
	//     *   **Asc**: ascending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sorting order of the retrieved entries. Valid values:
	//
	// *   **asc**: sorts the retrieved entries by time in ascending order.
	// *   **desc**: sorts the retrieved entries by time in descending order.
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **10** (default)
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A reserved parameter.
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// The keyword based on which audit logs are queried. You can set this parameter to a value of the STRING type.
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// *   **DELETE**
	// *   **SELECT**
	// *   **UPDATE**
	// *   **INSERT INTO SELECT**
	// *   **ALTER**
	// *   **DROP**
	// *   **CREATE**
	//
	// >  You can query only a single type of SQL statements at a time. If you leave this parameter empty, the **SELECT** statements are queried.
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether the execution of the SQL statement succeeds. Valid values:
	//
	// *   **true**
	// *   **false**
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The username that is used to execute the SQL statement.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBName(v string) *DescribeAuditLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEndTime(v string) *DescribeAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetHostAddress(v string) *DescribeAuditLogRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrder(v string) *DescribeAuditLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrderType(v string) *DescribeAuditLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageSize(v int32) *DescribeAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetProxyUser(v string) *DescribeAuditLogRecordsRequest {
	s.ProxyUser = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetRegionId(v string) *DescribeAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSqlType(v string) *DescribeAuditLogRecordsRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetStartTime(v string) *DescribeAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSucceed(v string) *DescribeAuditLogRecordsRequest {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetUser(v string) *DescribeAuditLogRecordsRequest {
	s.User = &v
	return s
}

type DescribeAuditLogRecordsResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried SQL audit logs.
	Items []*DescribeAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageSize(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetRequestId(v string) *DescribeAuditLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditLogRecordsResponseBodyItems struct {
	// The connection ID.
	ConnId *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	// The name of the database on which the SQL statement was executed.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The start time of the execution of the SQL statement. The time is displayed in the ISO 8601 standard in the yyyy-MM-dd HH:mm:ss format. The time must be in UTC.
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The task ID.
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The SQL statement.
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The type of the SQL statement.
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// Indicates whether the SQL statement was successfully executed. Valid values:
	//
	// *   **true**
	// *   **false**
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The amount of time that is consumed to execute the SQL statement. Unit: milliseconds.
	TotalTime *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The username that is used to execute the SQL statement.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetConnId(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ConnId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetDBName(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetProcessID(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ProcessID = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLType(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSucceed(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetTotalTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetUser(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.User = &v
	return s
}

type DescribeAuditLogRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetStatusCode(v int32) *DescribeAuditLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetBody(v *DescribeAuditLogRecordsResponseBody) *DescribeAuditLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// The number of days for which data backup files are retained.
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether log backup is enabled. Valid values:
	//
	// *   **Enable**
	// *   **Disable**
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which the log backup files are retained.
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The cycle based on which backups are performed. If more than one day of the week are specified, the days of the week are separated by commas (,). Valid value:
	//
	// *   Monday
	// *   Tuesday
	// *   Wednesday
	// *   Thursday
	// *   Friday
	// *   Saturday
	// *   Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The data backup time. The time is in the HH:mmZ-HH:mmZ format. The time is displayed in UTC.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	// The backup set ID.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupsResponseBody struct {
	// The queried backup sets.
	Items *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v string) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyItemsBackup struct {
	// The end time of the backup.
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The backup set ID.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Snapshot is returned.
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The size of the backup set. Unit: bytes.
	BackupSize *int32 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup.
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup type. Valid values:
	//
	// *   **FullBackup**
	// *   **IncrementalBackup**
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int32) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetStatusCode(v int32) *DescribeBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeClusterAccessWhiteListRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DescribeClusterAccessWhiteListResponseBody struct {
	// The queried IP address whitelists.
	Items *DescribeClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBody) SetItems(v *DescribeClusterAccessWhiteListResponseBodyItems) *DescribeClusterAccessWhiteListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterAccessWhiteListResponseBodyItems struct {
	IPArray []*DescribeClusterAccessWhiteListResponseBodyItemsIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeClusterAccessWhiteListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBodyItems) SetIPArray(v []*DescribeClusterAccessWhiteListResponseBodyItemsIPArray) *DescribeClusterAccessWhiteListResponseBodyItems {
	s.IPArray = v
	return s
}

type DescribeClusterAccessWhiteListResponseBodyItemsIPArray struct {
	// The attribute of the whitelist.
	//
	// > Whitelists with the **hidden** attribute are not displayed in the console. Those whitelists are used to access Data Transmission Service (DTS) and PolarDB.
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// Each cluster supports up to 50 IP address whitelists.
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses in the IP address whitelist. Up to 500 IP addresses can be returned. Multiple IP addresses are separated by commas (,).
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponseBodyItemsIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBodyItemsIPArray) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponse) SetStatusCode(v int32) *DescribeClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponse) SetBody(v *DescribeClusterAccessWhiteListResponseBody) *DescribeClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeClusterNetInfoRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeClusterNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoRequest) SetDBClusterId(v string) *DescribeClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

type DescribeClusterNetInfoResponseBody struct {
	// The network type of the cluster. Only the Virtual Private Cloud (VPC) network type is supported. **VPC** is returned.
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The network information about the cluster.
	Items *DescribeClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBody) SetItems(v *DescribeClusterNetInfoResponseBodyItems) *DescribeClusterNetInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeClusterNetInfoResponseBody) SetRequestId(v string) *DescribeClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterNetInfoResponseBodyItems struct {
	Address []*DescribeClusterNetInfoResponseBodyItemsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeClusterNetInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItems) SetAddress(v []*DescribeClusterNetInfoResponseBodyItemsAddress) *DescribeClusterNetInfoResponseBodyItems {
	s.Address = v
	return s
}

type DescribeClusterNetInfoResponseBodyItemsAddress struct {
	// The endpoint of the cluster.
	//
	// *   If the network type of the cluster is VPC, the VPC endpoint of the cluster is returned.
	// *   If the network type of the cluster is Public, the public endpoint of the cluster is returned.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the endpoint.
	//
	// *   If the network type of the cluster is VPC, the prefix of the private endpoint is returned.
	// *   If the network type of the cluster is Public, the prefix of the public endpoint is returned.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The IP address of the endpoint.
	//
	// *   If the network type of the cluster is VPC, the IP address of the private endpoint is returned.
	// *   If the network type of the cluster is Public, the IP address of the public endpoint is returned.
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// *   **Public**: Internet.
	// *   **VPC**: VPC.
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number that is used to connect to the cluster. **3306** is returned.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The VPC ID.
	//
	// > If NetType is set to Public, an empty string is returned for this parameter.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// > If NetType is set to Public, an empty string is returned for this parameter.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeClusterNetInfoResponseBodyItemsAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetNetType(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.NetType = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetPort(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.Port = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetVPCId(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

type DescribeClusterNetInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNetInfoResponse) SetStatusCode(v int32) *DescribeClusterNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNetInfoResponse) SetBody(v *DescribeClusterNetInfoResponseBody) *DescribeClusterNetInfoResponse {
	s.Body = v
	return s
}

type DescribeClusterResourceDetailRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeClusterResourceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailRequest) SetDBClusterId(v string) *DescribeClusterResourceDetailRequest {
	s.DBClusterId = &v
	return s
}

type DescribeClusterResourceDetailResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the cluster resource usage.
	Data *DescribeClusterResourceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBody) SetCode(v int32) *DescribeClusterResourceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBody) SetData(v *DescribeClusterResourceDetailResponseBodyData) *DescribeClusterResourceDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeClusterResourceDetailResponseBody) SetRequestId(v string) *DescribeClusterResourceDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResourceDetailResponseBodyData struct {
	// The amount of reserved computing resources. Unit: AnalyticDB compute units (ACUs). Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The amount of idle reserved computing resources. Unit: ACUs. Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	FreeComputeResource *string `json:"FreeComputeResource,omitempty" xml:"FreeComputeResource,omitempty"`
	// The resource groups.
	ResourceGroupList []*DescribeClusterResourceDetailResponseBodyDataResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// The amount of reserved storage resources. Unit: ACUs. Valid values: 0 to 2064. The value must be in increments of 24 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetComputeResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.ComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetDBClusterId(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetFreeComputeResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.FreeComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetResourceGroupList(v []*DescribeClusterResourceDetailResponseBodyDataResourceGroupList) *DescribeClusterResourceDetailResponseBodyData {
	s.ResourceGroupList = v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetStorageResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.StorageResource = &v
	return s
}

type DescribeClusterResourceDetailResponseBodyDataResourceGroupList struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// Indicates whether the preemptible instance feature is enabled for the resource group. After the preemptible instance feature is enabled, you are charged for resources at a lower unit price but the resources are probably released. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// The True value is returned only for job resource groups.
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACUs.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: ACUs.
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The resource group ID.
	PoolId *int64 `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// The name of the resource group.
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The type of the resource group.
	PoolType *string `json:"PoolType,omitempty" xml:"PoolType,omitempty"`
	// The user of the resource group.
	PoolUsers *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	// A reserved parameter.
	RunningClusterCount *int32 `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   **running**
	// *   **deleting**
	// *   **scaling**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBodyDataResourceGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetClusterMode(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.ClusterMode = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetClusterSizeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.ClusterSizeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetEnableSpot(v bool) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.EnableSpot = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMaxClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MaxClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMaxComputeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMinClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MinClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMinComputeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolId(v int64) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolId = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolName(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolName = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolType(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolType = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolUsers(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolUsers = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetRunningClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.RunningClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetStatus(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.Status = &v
	return s
}

type DescribeClusterResourceDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResourceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResourceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterResourceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourceDetailResponse) SetStatusCode(v int32) *DescribeClusterResourceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResourceDetailResponse) SetBody(v *DescribeClusterResourceDetailResponseBody) *DescribeClusterResourceDetailResponse {
	s.Body = v
	return s
}

type DescribeClusterResourceUsageRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageRequest) SetDBClusterId(v string) *DescribeClusterResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceUsageRequest) SetEndTime(v string) *DescribeClusterResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterResourceUsageRequest) SetStartTime(v string) *DescribeClusterResourceUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeClusterResourceUsageResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeClusterResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBody) SetCode(v int32) *DescribeClusterResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBody) SetData(v *DescribeClusterResourceUsageResponseBodyData) *DescribeClusterResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeClusterResourceUsageResponseBody) SetRequestId(v string) *DescribeClusterResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeClusterResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeClusterResourceUsageResponseBodyDataAcuInfo) *DescribeClusterResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetEndTime(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetStartTime(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeClusterResourceUsageResponseBodyDataAcuInfo struct {
	// The resource usage metric. Valid values:
	//
	// *   `TotalAcuNumber`: the total number of ACUs.
	// *   `ReservedAcuNumber`: the number of ACUs for the reserved resources.
	// *   `ReservedAcuUsageNumber`: the number of ACUs for the reserved resources that are used.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the metric at specific points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeClusterResourceUsageResponseBodyDataAcuInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeClusterResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeClusterResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

type DescribeClusterResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeClusterResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourceUsageResponse) SetStatusCode(v int32) *DescribeClusterResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResourceUsageResponse) SetBody(v *DescribeClusterResourceUsageResponseBody) *DescribeClusterResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeColumnsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRegionId(v string) *DescribeColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnsRequest) SetSchemaName(v string) *DescribeColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

type DescribeColumnsResponseBody struct {
	// The queried columns.
	Items *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColumnsResponseBodyItems struct {
	Column []*DescribeColumnsResponseBodyItemsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) SetColumn(v []*DescribeColumnsResponseBodyItemsColumn) *DescribeColumnsResponseBodyItems {
	s.Column = v
	return s
}

type DescribeColumnsResponseBodyItemsColumn struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// *   **true**
	// *   **false**
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The name of the column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// *   **true**
	// *   **false**
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The data type of the column.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsColumn) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetAutoIncrementColumn(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetColumnName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetDBClusterId(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetPrimaryKey(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetSchemaName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetTableName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetType(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.Type = &v
	return s
}

type DescribeColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponse) SetHeaders(v map[string]*string) *DescribeColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnsResponse) SetStatusCode(v int32) *DescribeColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeComputeResourceUsageRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeComputeResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageRequest) SetDBClusterId(v string) *DescribeComputeResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetEndTime(v string) *DescribeComputeResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetResourceGroupName(v string) *DescribeComputeResourceUsageRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeComputeResourceUsageRequest) SetStartTime(v string) *DescribeComputeResourceUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeComputeResourceUsageResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeComputeResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComputeResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBody) SetCode(v int32) *DescribeComputeResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBody) SetData(v *DescribeComputeResourceUsageResponseBodyData) *DescribeComputeResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeComputeResourceUsageResponseBody) SetRequestId(v string) *DescribeComputeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComputeResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeComputeResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the resource group.
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeComputeResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeComputeResourceUsageResponseBodyDataAcuInfo) *DescribeComputeResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetEndTime(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetResourceGroupName(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetResourceGroupType(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.ResourceGroupType = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetStartTime(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeComputeResourceUsageResponseBodyDataAcuInfo struct {
	// The resource usage metric. Valid values:
	//
	// *   `TotalAcuNumber`: the total number of ACUs.
	// *   `ReservedAcuNumber`: the number of ACUs for the reserved resources.
	// *   `ReservedAcuUsageNumber`: the number of ACUs for the reserved resources that are used.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the metric at specific points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeComputeResourceUsageResponseBodyDataAcuInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeComputeResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeComputeResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

type DescribeComputeResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComputeResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComputeResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeComputeResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeResourceUsageResponse) SetStatusCode(v int32) *DescribeComputeResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComputeResourceUsageResponse) SetBody(v *DescribeComputeResourceUsageResponseBody) *DescribeComputeResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAttributeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAttributeResponseBody struct {
	// The queried AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster information.
	Items *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItems struct {
	DBCluster []*DescribeDBClusterAttributeResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItems) SetDBCluster(v []*DescribeDBClusterAttributeResponseBodyItemsDBCluster) *DescribeDBClusterAttributeResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBCluster struct {
	// The billing method of the cluster. Valid values:
	//
	// *   **ads**: pay-as-you-go.
	// *   **ads_pre**: subscription.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The specifications of reserved computing resources. Each ACU is equivalent to 1 core and 4 GB memory. Computing resources are used to compute data. The increase in the computing resources can accelerate queries. You can scale computing resources based on your business requirements.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The total amount of computing resources in the cluster. Each ACU is equivalent to 1 core and 4 GB memory.
	ComputeResourceTotal *string `json:"ComputeResourceTotal,omitempty" xml:"ComputeResourceTotal,omitempty"`
	// The public endpoint that is used to connect to the cluster.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddThh:mm:ssZ` format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the cluster.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. **VPC** is returned.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. Valid values:
	//
	// *   **Preparing**
	// *   **Creating**
	// *   **Running**
	// *   **Deleting**
	// *   **Restoring**
	// *   **ClassChanging**
	// *   **NetAddressCreating**
	// *   **NetAddressDeleting**
	// *   **NetAddressModifying**
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The cluster type. By default, **Common** is returned, which indicates a common cluster.
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The engine version of the AnalyticDB for MySQL Data Lakehouse Edition cluster. **5.0** is returned.
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The engine of the cluster. **AnalyticDB** is returned.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor version of the cluster.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the cluster.
	//
	// *   If the billing method of the cluster is subscription, the actual expiration time is returned.
	// *   If the billing method of the cluster is pay-as-you-go, null is returned.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the subscription cluster has expired. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// >
	//
	// *   If the cluster has expired, the system locks or releases the cluster within a period of time. We recommend that you renew the expired cluster. For more information, see [Renewal policy](~~135248~~).
	//
	// *   This parameter is not returned for pay-as-you-go clusters.
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// *   **Unlock**: The cluster is not locked.
	// *   **ManualLock**: The cluster is manually locked.
	// *   **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire** is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maintenance window of the cluster. The time is displayed in the `HH:mmZ-HH:mmZ` format in UTC.
	//
	// >  For more information about maintenance windows, see [Configure a maintenance window](~~122569~~).
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The mode of the cluster. By default, **flexible** is returned, which indicates that the cluster is in elastic mode.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the cluster.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remaining reserved computing resources that are available in the cluster. Each ACU is equivalent to 1 core and 4 GB memory.
	ReservedACU *string `json:"ReservedACU,omitempty" xml:"ReservedACU,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of reserved storage resources. Each AnalyticDB compute unit (ACU) is equivalent to 1 core and 4 GB memory. Storage resources are used to read and write data. The increase in the storage resources can improve the read and write performance of the cluster.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The total amount of storage resources in the cluster. Each ACU is equivalent to 1 core and 4 GB memory.
	StorageResourceTotal *string `json:"StorageResourceTotal,omitempty" xml:"StorageResourceTotal,omitempty"`
	// A reserved parameter.
	SupportedFeatures map[string]*string `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// Indicates whether Elastic Network Interface (ENI) is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	UserENIStatus *bool `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResourceTotal(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResourceTotal = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetReservedACU(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ReservedACU = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResourceTotal(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResourceTotal = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetSupportedFeatures(v map[string]*string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIStatus(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag struct {
	// The tag key.
	//
	// >  You can call the [TagResources](~~179253~~) operation to add tags to a cluster.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetStatusCode(v int32) *DescribeDBClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterHealthStatusRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusRequest) SetDBClusterId(v string) *DescribeDBClusterHealthStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusRequest) SetRegionId(v string) *DescribeDBClusterHealthStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBody struct {
	// The access nodes of the queried cluster.
	CS *DescribeDBClusterHealthStatusResponseBodyCS `json:"CS,omitempty" xml:"CS,omitempty" type:"Struct"`
	// The compute node groups of the queried cluster.
	Executor *DescribeDBClusterHealthStatusResponseBodyExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// The health state of the cluster. Valid values:
	//
	// *   **RISK**
	// *   **NORMAL**
	// *   **UNAVAILABLE**
	//
	// >  When the states of the access nodes, compute node groups, and storage node groups of a cluster are all **NORMAL** and a connection to the cluster is established, the state of the cluster is **NORMAL**. When the state of the access nodes, compute node groups, or storage node groups of the cluster is **RISK**, the state of the cluster is **RISK**. When the state of the access nodes, compute node groups, or storage node groups of the cluster is **UNAVAILABLE**, the state of the cluster is **UNAVAILABLE**.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage node groups of the queried cluster.
	Worker *DescribeDBClusterHealthStatusResponseBodyWorker `json:"Worker,omitempty" xml:"Worker,omitempty" type:"Struct"`
}

func (s DescribeDBClusterHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetCS(v *DescribeDBClusterHealthStatusResponseBodyCS) *DescribeDBClusterHealthStatusResponseBody {
	s.CS = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetExecutor(v *DescribeDBClusterHealthStatusResponseBodyExecutor) *DescribeDBClusterHealthStatusResponseBody {
	s.Executor = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetInstanceStatus(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetRequestId(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetWorker(v *DescribeDBClusterHealthStatusResponseBodyWorker) *DescribeDBClusterHealthStatusResponseBody {
	s.Worker = v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyCS struct {
	// The number of healthy access nodes.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of access nodes.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky nodes.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of access nodes. Valid values:
	//
	// *   **RISK**
	// *   **NORMAL**
	// *   **UNAVAILABLE**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable access nodes.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyExecutor struct {
	// The number of healthy access nodes.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of compute nodes.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky nodes.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of compute node groups. Valid values:
	//
	// *   **RISK**
	// *   **NORMAL**
	// *   **UNAVAILABLE**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable access nodes.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyWorker struct {
	// The number of healthy storage node groups.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of storage node groups.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky storage node groups.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of storage node groups. Valid values:
	//
	// *   **RISK**
	// *   **NORMAL**
	// *   **UNAVAILABLE**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable storage node groups.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetStatusCode(v int32) *DescribeDBClusterHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetBody(v *DescribeDBClusterHealthStatusResponseBody) *DescribeDBClusterHealthStatusResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~~612397~~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is two days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics to be queried. Separate multiple values with commas (,). Valid values:
	//
	// *   CPU
	//
	//     *   **AnalyticDB_CPU_Usage_Percentage**: the average CPU utilization.
	//
	// *   Connections
	//
	//     *   **AnalyticDB_Instance_Connection_Count**: the number of database connections.
	//
	// *   Writes
	//
	//     *   **AnalyticDB_TPS**: the write transactions per second (TPS).
	//     *   **AnalyticDB_InsertRT**: the write response time.
	//     *   **AnalyticDB_InsertBytes**: the write throughput.
	//
	// *   Queries
	//
	//     *   **AnalyticDB_QPS**: the queries per second (QPS).
	//     *   **AnalyticDB_QueryRT**: the query response time.
	//     *   **AnalyticDB_QueryWaitTime**: the query wait time.
	//
	// *   Disks
	//
	//     *   **AnalyticDB_Disk_IO_Avg_Usage_Percentage**: the average I/O utilization.
	//     *   **AnalyticDB_Disk_IO_Avg_Waiting_Time**: the average I/O wait time.
	//     *   **AnalyticDB_IO_Throughput**: the disk throughput.
	//     *   **AnalyticDB_IOPS**: the disk IOPS.
	//     *   **AnalyticDB_Disk_Usage**: the disk space that is used.
	//     *   **AnalyticDB_Disk_Usage_Percentage**: the disk usage.
	//     *   **AnalyticDB_HotDataDiskUsage**: the disk space that is used by hot data.
	//     *   **AnalyticDB_ColdDataDiskUsage**: the disk space that is used by hot data.
	//
	// > This parameter must be specified.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~612393~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourcePools *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetRegionId(v string) *DescribeDBClusterPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	// The name of the performance metric.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried performance metric data.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag value.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The values of the performance metric at different points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetTags(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Tags = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClusterSpaceSummaryRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetDBClusterId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetRegionId(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterSpaceSummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryRequest) SetResourceOwnerId(v int64) *DescribeDBClusterSpaceSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponseBody struct {
	Data      *DescribeDBClusterSpaceSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) SetData(v *DescribeDBClusterSpaceSummaryResponseBodyData) *DescribeDBClusterSpaceSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) SetRequestId(v string) *DescribeDBClusterSpaceSummaryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponseBodyData struct {
	ColdData   *DescribeDBClusterSpaceSummaryResponseBodyDataColdData   `json:"ColdData,omitempty" xml:"ColdData,omitempty" type:"Struct"`
	DataGrowth *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth `json:"DataGrowth,omitempty" xml:"DataGrowth,omitempty" type:"Struct"`
	HotData    *DescribeDBClusterSpaceSummaryResponseBodyDataHotData    `json:"HotData,omitempty" xml:"HotData,omitempty" type:"Struct"`
	TotalSize  *string                                                  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetColdData(v *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.ColdData = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetDataGrowth(v *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.DataGrowth = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetHotData(v *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.HotData = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetTotalSize(v string) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.TotalSize = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponseBodyDataColdData struct {
	DataSize            *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	IndexSize           *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	OtherSize           *int64 `json:"OtherSize,omitempty" xml:"OtherSize,omitempty"`
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	TotalSize           *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataColdData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetDataSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.DataSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetOtherSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.OtherSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetPrimaryKeyIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetTotalSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.TotalSize = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth struct {
	DayGrowth  *int64 `json:"DayGrowth,omitempty" xml:"DayGrowth,omitempty"`
	WeekGrowth *int64 `json:"WeekGrowth,omitempty" xml:"WeekGrowth,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) SetDayGrowth(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth {
	s.DayGrowth = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) SetWeekGrowth(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth {
	s.WeekGrowth = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponseBodyDataHotData struct {
	DataSize            *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	IndexSize           *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	OtherSize           *int64 `json:"OtherSize,omitempty" xml:"OtherSize,omitempty"`
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	TotalSize           *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataHotData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetDataSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.DataSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetOtherSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.OtherSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetPrimaryKeyIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetTotalSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.TotalSize = &v
	return s
}

type DescribeDBClusterSpaceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterSpaceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetHeaders(v map[string]*string) *DescribeDBClusterSpaceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetStatusCode(v int32) *DescribeDBClusterSpaceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetBody(v *DescribeDBClusterSpaceSummaryResponseBody) *DescribeDBClusterSpaceSummaryResponse {
	s.Body = v
	return s
}

type DescribeDBClusterStatusRequest struct {
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusRequest) SetRegionId(v string) *DescribeDBClusterStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeDBClusterStatusResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried cluster states.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponseBody) SetRequestId(v string) *DescribeDBClusterStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterStatusResponseBody) SetStatus(v []*string) *DescribeDBClusterStatusResponseBody {
	s.Status = v
	return s
}

type DescribeDBClusterStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetStatusCode(v int32) *DescribeDBClusterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetBody(v *DescribeDBClusterStatusResponseBody) *DescribeDBClusterStatusResponse {
	s.Body = v
	return s
}

type DescribeDBClustersRequest struct {
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// If you do not specify this parameter, the information about all clusters that reside in the region is returned.
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The state of the cluster. Valid values:
	//
	// *   **Preparing**
	//
	// <!---->
	//
	// *   **Creating**
	// *   **Running**
	// *   **Deleting**
	//
	// <!---->
	//
	// *   **Restoring**
	//
	// <!---->
	//
	// *   **ClassChanging**
	// *   **NetAddressCreating**
	// *   **NetAddressDeleting**
	// *   **NetAddressModifying**
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. If you do not specify this parameter, the information about all resource groups in the cluster is returned.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are added to the cluster.
	Tag []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

type DescribeDBClustersRequestTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponseBody struct {
	// The queried cluster.
	Items *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBClustersResponseBodyItems struct {
	DBCluster []*DescribeDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItems) SetDBCluster(v []*DescribeDBClustersResponseBodyItemsDBCluster) *DescribeDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	// The billing method of the cluster. Valid values:
	//
	// *   **ads**: pay-as-you-go.
	// *   **ads_pre**: subscription.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The specifications of reserved computing resources. Each ACU is equivalent to 1 core and 4 GB memory. Computing resources are used to compute data. The increase in the computing resources can accelerate queries. You can scale computing resources based on your business requirements.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The public endpoint that is used to connect to the cluster.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the *yyyy-mm-ddThh:mm:ssZ* format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. **VPC** is returned.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. Valid values:
	//
	// *   **Preparing**
	//
	// <!---->
	//
	// *   **Creating**
	// *   **Running**
	// *   **Deleting**
	//
	// <!---->
	//
	// *   **Restoring**
	//
	// <!---->
	//
	// *   **ClassChanging**
	// *   **NetAddressCreating**
	// *   **NetAddressDeleting**
	// *   **NetAddressModifying**
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The type of the cluster. By default, **Common** is returned, which indicates a common cluster.
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The version of AnalyticDB for MySQL Data Lakehouse Edition. **5.0** is returned.
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The database engine of the cluster. **AnalyticDB** is returned.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The time when the cluster expired. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	//
	// > - The expiration time is returned for a subscription cluster.
	// > - Anempty string is returned for a pay-as-you-go cluster.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the subscription cluster has expired. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > - If the cluster has expired, the system locks or releases the cluster within a period of time. We recommend that you renew expired clusters. For more information, see [Renewal policy](~~135246~~).
	// > - This parameter is not returned for pay-as-you-go clusters.
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The lock state of the cluster. Valid values:
	//
	// *   **Unlock**: The cluster is not locked.
	// *   **ManualLock**: The cluster is manually locked.
	// *   **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire** is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The mode of the cluster. By default, **flexible** is returned, which indicates that the cluster is in elastic mode.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the cluster.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The amount of remaining reserved computing resources that are available in the cluster. Each ACU is equivalent to 1 core and 4 GB memory.
	ReservedACU *string `json:"ReservedACU,omitempty" xml:"ReservedACU,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of reserved storage resources. Each AnalyticDB compute unit (ACU) is equivalent to 1 core and 4 GB memory. Storage resources are used to read and write data. The increase in the storage resources can improve the read and write performance of the cluster.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The virtual private cloud (VPC) ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetReservedACU(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ReservedACU = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	// The tag key.
	//
	// >  You can call the [TagResources](~~179253~~) operation to add tags to a cluster.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) SetHeaders(v map[string]*string) *DescribeDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersResponse) SetStatusCode(v int32) *DescribeDBClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

type DescribeDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// > If you do not specify this parameter, the information about all resource groups in the cluster is returned.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](~~428610~~).
	GroupType            *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupRequest) SetDBClusterId(v string) *DescribeDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupName(v string) *DescribeDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupType(v string) *DescribeDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetRegionId(v string) *DescribeDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DescribeDBResourceGroupResponseBody struct {
	// The queried resource group.
	GroupsInfo []*DescribeDBResourceGroupResponseBodyGroupsInfo `json:"GroupsInfo,omitempty" xml:"GroupsInfo,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBody) SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody {
	s.GroupsInfo = v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetRequestId(v string) *DescribeDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBResourceGroupResponseBodyGroupsInfo struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The time when the resource group was created. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The minimum amount of elastic computing resources. Unit: ACU.
	ElasticMinComputeResource *string `json:"ElasticMinComputeResource,omitempty" xml:"ElasticMinComputeResource,omitempty"`
	EnableSpot                *string `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// >  For more information about resource groups, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The Resource Access Management (RAM) user that is associated with the resource group.
	GroupUsers *string `json:"GroupUsers,omitempty" xml:"GroupUsers,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACU.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: AnalyticDB compute unit (ACU).
	MinComputeResource *string                                               `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	Rules              []*DescribeDBResourceGroupResponseBodyGroupsInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// A reserved parameter.
	RunningClusterCount *int32 `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	// The status of the resource group. Valid values:
	//
	// *   **creating**: The resource group is being created.
	// *   **ok**: The resource group is created.
	// *   **pendingdelete**: The resource group is pending to be deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource group was updated. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterMode(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterMode = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterSizeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterSizeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetCreateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetElasticMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ElasticMinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEnableSpot(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.EnableSpot = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupUsers(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupUsers = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRules(v []*DescribeDBResourceGroupResponseBodyGroupsInfoRules) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Rules = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRunningClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.RunningClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetStatus(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Status = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetUpdateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.UpdateTime = &v
	return s
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRules struct {
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	QueryTime       *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetQueryTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.QueryTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetTargetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.TargetGroupName = &v
	return s
}

type DescribeDBResourceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetStatusCode(v int32) *DescribeDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetBody(v *DescribeDBResourceGroupResponseBody) *DescribeDBResourceGroupResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisDimensionsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Valid values:
	//
	// *   **zh-CN** (default): simplified Chinese.
	// *   **en-US**: English.
	// *   **ja**: Japanese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, `Min`, and `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries the executed SQL statements. You can set `Value` to `running` to query the SQL statements that are being executed. You can also set Value to `failed` to query the SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries the SQL statements whose execution duration is in the range of 10 to 200 milliseconds. You can also specify custom values for the Min and Max fields.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can query data only within the last 14 days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBClusterId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetEndTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetLang(v string) *DescribeDiagnosisDimensionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetQueryCondition(v string) *DescribeDiagnosisDimensionsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetRegionId(v string) *DescribeDiagnosisDimensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetStartTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.StartTime = &v
	return s
}

type DescribeDiagnosisDimensionsResponseBody struct {
	// The queried source IP addresses.
	ClientIps []*string `json:"ClientIps,omitempty" xml:"ClientIps,omitempty" type:"Repeated"`
	// The queried database names.
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried resource group names.
	ResourceGroups []*string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The queried usernames.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetClientIps(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ClientIps = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetResourceGroups(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisRecordsRequest struct {
	// The source IP address.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statements are executed.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query keyword of the SQL statements.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statements. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the SQL statements. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statements. Unit: bytes.
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statements. Unit: bytes.
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The order in which to sort the SQL statements by field, which contains the `Field` and `Type` fields. Specify the order in the JSON format. Example: `[{"Field":"StartTime", "Type": "desc"}]`. Fields:
	//
	// *   `Field` specifies the field that is used to sort the SQL statements. Valid values:
	//
	//     *   `StartTime`: the execution start time.
	//     *   `Status`: the execution status.
	//     *   `UserName`: the username.
	//     *   `Cost`: the execution duration.
	//     *   `PeakMemory`: the peak memory.
	//     *   `ScanSize`: the amount of data that is scanned.
	//     *   `Database`: the name of the database.
	//     *   `ClientIp`: the source IP address.
	//     *   `ResourceGroup`: the name of the resource group.
	//     *   `QueueTime`: the amount of time that is consumed for queuing.
	//     *   `OutputRows`: the number of output rows.
	//     *   `OutputDataSize`: the amount of output data.
	//     *   `ResourceCostRank`: the execution duration rank of operators that are used in the SQL statements. This value takes effect only when `QueryCondition` is set to `{"Type":"status","Value":"running"}`.
	//
	// *   `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     *   `Desc`: descending order.
	//     *   `Asc`: ascending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The SQL pattern ID.
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, `Min`, and `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries the executed SQL statements. You can set `Value` to `running` to query the SQL statements that are being executed. You can also set Value to `failed` to query the SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries the SQL statements whose execution duration is in the range of 10 to 200 milliseconds. You can also specify custom values for the Min and Max fields.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statements belong.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can query data only within the last 14 days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to execute the SQL statements. You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) SetClientIp(v string) *DescribeDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDBClusterId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDatabase(v string) *DescribeDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetEndTime(v string) *DescribeDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetKeyword(v string) *DescribeDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetLang(v string) *DescribeDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetOrder(v string) *DescribeDiagnosisRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageSize(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPatternId(v string) *DescribeDiagnosisRecordsRequest {
	s.PatternId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetRegionId(v string) *DescribeDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetResourceGroup(v string) *DescribeDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUserName(v string) *DescribeDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponseBody struct {
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried SQL statements.
	Querys []*DescribeDiagnosisRecordsResponseBodyQuerys `json:"Querys,omitempty" xml:"Querys,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageSize(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetQuerys(v []*DescribeDiagnosisRecordsResponseBodyQuerys) *DescribeDiagnosisRecordsResponseBody {
	s.Querys = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisRecordsResponseBodyQuerys struct {
	// The source IP address.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The total execution duration. Unit: milliseconds.
	//
	// >  This value is the cumulative value of the `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime` parameters.
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The name of the database on which the SQL statement is executed.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of rows written to the table by an extract-transform-load (ETL) job.
	EtlWriteRows *int64 `json:"EtlWriteRows,omitempty" xml:"EtlWriteRows,omitempty"`
	// The execution duration. Unit: milliseconds.
	ExecutionTime *int64 `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The amount of returned data. Unit: bytes.
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The number of rows returned.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The peak memory. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The amount of time that is consumed for queuing. Unit: milliseconds.
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	RcHost *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	// The execution duration rank of operators that are used in the SQL statement.
	//
	// >  This parameter is returned only for SQL statements whose `Status` parameter is `running`.
	ResourceCostRank *int32 `json:"ResourceCostRank,omitempty" xml:"ResourceCostRank,omitempty"`
	// The resource group to which the SQL statement belongs.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The queried SQL statement.
	//
	// >  For performance considerations, an SQL statement cannot exceed 5,120 characters in length. Otherwise, the SQL statement is truncated. You can call the [DownloadDiagnosisRecords](~~308212~~) operation to download the information about SQL statements that meet a query condition for an AnalyticDB for MySQL cluster, including the complete SQL statements.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// Indicates whether the SQL statement is truncated. Valid values:
	//
	// *   **true**
	// *   **false**
	SQLTruncated *bool `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	// The maximum length of the SQL statement. 5120 is returned. Unit: characters. SQL statements that exceed this limit are truncated.
	SQLTruncatedThreshold *int64 `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	// The number of rows scanned.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: bytes.
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The execution start time of the SQL statement. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// *   **running**
	// *   **finished**
	// *   **failed**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The amount of time that is consumed to generate an execution plan. Unit: milliseconds.
	TotalPlanningTime *int64 `json:"TotalPlanningTime,omitempty" xml:"TotalPlanningTime,omitempty"`
	// The total number of stages generated.
	TotalStages *int32 `json:"TotalStages,omitempty" xml:"TotalStages,omitempty"`
	// The username that is used to execute the SQL statements.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetClientIp(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetCost(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetEtlWriteRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.EtlWriteRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetExecutionTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputDataSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetPeakMemory(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetProcessId(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetQueueTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.QueueTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetRcHost(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceCostRank(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceCostRank = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceGroup(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQL(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQL = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncatedThreshold(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalPlanningTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalPlanningTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalStages(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalStages = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetUserName(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisSQLInfoRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh**: simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query ID.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the diagnostic information about SQL statements for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster, including the query ID.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the diagnostic information about SQL statements for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster, including the IP address and port number of the frontend node.
	ProcessRcHost *string `json:"ProcessRcHost,omitempty" xml:"ProcessRcHost,omitempty"`
	// The execution start time of the SQL statement. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the diagnostic information about SQL statements for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster, including the execution start time of the SQL statement.
	ProcessStartTime *int64 `json:"ProcessStartTime,omitempty" xml:"ProcessStartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// *   **running**
	// *   **finished**
	// *   **failed**
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the diagnostic information about SQL statements for an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster, including the status of the SQL statement.
	ProcessState *string `json:"ProcessState,omitempty" xml:"ProcessState,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBClusterId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetLang(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessRcHost(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessRcHost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessStartTime(v int64) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessStartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessState(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessState = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetRegionId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.RegionId = &v
	return s
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	// The queried execution information, including the SQL statement, statistics, execution plan, and operator information.
	DiagnosisSQLInfo *string `json:"DiagnosisSQLInfo,omitempty" xml:"DiagnosisSQLInfo,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried execution information by stage.
	StageInfos []*DescribeDiagnosisSQLInfoResponseBodyStageInfos `json:"StageInfos,omitempty" xml:"StageInfos,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDiagnosisSQLInfo(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.DiagnosisSQLInfo = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStageInfos(v []*DescribeDiagnosisSQLInfoResponseBodyStageInfos) *DescribeDiagnosisSQLInfoResponseBody {
	s.StageInfos = v
	return s
}

type DescribeDiagnosisSQLInfoResponseBodyStageInfos struct {
	// The total amount of input data in the stage. Unit: bytes.
	InputDataSize *int64 `json:"InputDataSize,omitempty" xml:"InputDataSize,omitempty"`
	// The total number of input rows in the stage.
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The total amount of time consumed by all operators in the stage. Unit: milliseconds.
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The total amount of output data in the stage. Unit: bytes.
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The total number of output rows in the stage.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The total peak memory of the stage. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The execution progress of the stage.
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The stage ID.
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The state of the stage.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOperatorCost(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OperatorCost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetPeakMemory(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetProgress(v float64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.Progress = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetStageId(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.StageId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetState(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.State = &v
	return s
}

type DescribeDiagnosisSQLInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeDownloadRecordsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of the returned data. Valid values:
	//
	// *   **zh**: simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) SetDBClusterId(v string) *DescribeDownloadRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetLang(v string) *DescribeDownloadRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetRegionId(v string) *DescribeDownloadRecordsRequest {
	s.RegionId = &v
	return s
}

type DescribeDownloadRecordsResponseBody struct {
	// The queried download tasks.
	Records []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	// The download task ID.
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The error message returned if the download task failed.
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// The name of the downloaded file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The state of the download task. Valid values:
	//
	// *   **running**
	// *   **finished**
	// *   **failed**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The download URL of the file.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

type DescribeDownloadRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetStatusCode(v int32) *DescribeDownloadRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

type DescribeElasticPlanAttributeRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the ID of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](~~601334~~) operation to query the name of a scaling plan.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s DescribeElasticPlanAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeRequest) SetDBClusterId(v string) *DescribeElasticPlanAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanAttributeRequest) SetElasticPlanName(v string) *DescribeElasticPlanAttributeRequest {
	s.ElasticPlanName = &v
	return s
}

type DescribeElasticPlanAttributeResponseBody struct {
	// Details of the scaling plan.
	ElasticPlan *DescribeElasticPlanAttributeResponseBodyElasticPlan `json:"ElasticPlan,omitempty" xml:"ElasticPlan,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticPlanAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponseBody) SetElasticPlan(v *DescribeElasticPlanAttributeResponseBodyElasticPlan) *DescribeElasticPlanAttributeResponseBody {
	s.ElasticPlan = v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBody) SetRequestId(v string) *DescribeElasticPlanAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticPlanAttributeResponseBodyElasticPlan struct {
	// Indicates whether **Proportional Default Scaling for EIUs** is enabled.
	//
	// Valid values:
	//
	// true: Proportional Default Scaling for EIUs is enabled. If you set this parameter to true, the amount of storage resources scales along with the computing resources.
	//
	// false: Proportional Default Scaling for EIUs is not enabled.
	//
	// >  You can enable Proportional Default Scaling for EIUs for only a single scaling plan of a cluster.
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// A CORN expression that indicates the scaling cycle and time for the scaling plan.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The name of the scaling plan.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the scaling plan was enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The end time of the scaling plan.
	//
	// >  The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time of the scaling plan.
	//
	// >  The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The amount of elastic resources after scaling.
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the scaling plan.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanAttributeResponseBodyElasticPlan) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponseBodyElasticPlan) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetAutoScale(v bool) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.AutoScale = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetCronExpression(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.CronExpression = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetElasticPlanName(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetEnabled(v bool) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetEndTime(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetResourceGroupName(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetStartTime(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetTargetSize(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetType(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.Type = &v
	return s
}

type DescribeElasticPlanAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanAttributeResponse) SetStatusCode(v int32) *DescribeElasticPlanAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponse) SetBody(v *DescribeElasticPlanAttributeResponseBody) *DescribeElasticPlanAttributeResponse {
	s.Body = v
	return s
}

type DescribeElasticPlanJobsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >
	//
	// *   If you do not specify this parameter, all scaling plans of the cluster are queried.
	//
	// *   You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource group.
	//
	// >
	//
	// *   If you do not specify this parameter, the scaling plans of all resource groups are queried, including the interactive resource group and elastic I/O unit (EIU) types.
	//
	// *   You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the resource group name for a cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the scaling plan job. Valid values:
	//
	// *   RUNNING
	// *   SUCCESSFUL
	// *   FAILED
	//
	// >  If you do not specify this parameter, the scaling plans in all states are queried.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticPlanJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsRequest) SetDBClusterId(v string) *DescribeElasticPlanJobsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetElasticPlanName(v string) *DescribeElasticPlanJobsRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetPageNumber(v int32) *DescribeElasticPlanJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetPageSize(v int32) *DescribeElasticPlanJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetResourceGroupName(v string) *DescribeElasticPlanJobsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetStartTime(v string) *DescribeElasticPlanJobsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetStatus(v string) *DescribeElasticPlanJobsRequest {
	s.Status = &v
	return s
}

type DescribeElasticPlanJobsResponseBody struct {
	// The queried scaling plan jobs.
	Jobs []*DescribeElasticPlanJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of scaling plan jobs.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlanJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponseBody) SetJobs(v []*DescribeElasticPlanJobsResponseBodyJobs) *DescribeElasticPlanJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetPageNumber(v int32) *DescribeElasticPlanJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetPageSize(v int32) *DescribeElasticPlanJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetRequestId(v string) *DescribeElasticPlanJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetTotalCount(v int32) *DescribeElasticPlanJobsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeElasticPlanJobsResponseBodyJobs struct {
	// The amount of elastic resources.
	//
	// >
	//
	// *   If Type is set to EXECUTOR, ElasticAcu indicates the amount of elastic resources in the current resource group.
	// *   If Type is set to WORKER, ElasticAcu indicates the total amount of elastic storage resources in the current cluster.
	ElasticAcu *string `json:"ElasticAcu,omitempty" xml:"ElasticAcu,omitempty"`
	// The name of the scaling plan.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The end time of the scaling plan job.
	//
	// >  The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of compute nodes or storage replica sets.
	//
	// >
	//
	// *   If Type is set to EXECUTOR, InstanceSize indicates the number of compute nodes in the cluster.
	// *   If Type is set to EXECUTOR, InstanceSize indicates the number of storage replica sets in the cluster.
	InstanceSize *int32 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// The amount of reserved resources.
	//
	// >
	//
	// *   If Type is set to EXECUTOR, ReserveAcu indicates the amount of reserved resources in the current resource group.
	// *   If Type is set to WORKER, ReserveAcu indicates the total amount of reserved storage resources in the current cluster.
	ReserveAcu *string `json:"ReserveAcu,omitempty" xml:"ReserveAcu,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time of the scaling plan job.
	//
	// >  The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the scaling plan job. Valid values:
	//
	// *   RUNNING
	// *   SUCCESSFUL
	// *   FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The desired specifications of elastic resources after scaling.
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The total amount of resources.
	//
	// >
	//
	// *   If Type is set to EXECUTOR, TotalAcu indicates the total amount of computing resources in the current resource group.
	// *   If Type is set to WORKER, TotalAcu indicates the total amount of storage resources in the cluster.
	TotalAcu *string `json:"TotalAcu,omitempty" xml:"TotalAcu,omitempty"`
	// The type of the scaling plan job. Valid values:
	//
	// *   EXECUTOR: the interactive resource group type, which indicates the computing resource type.
	// *   WORKER: the EIU type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetElasticAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ElasticAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetElasticPlanName(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetEndTime(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetInstanceSize(v int32) *DescribeElasticPlanJobsResponseBodyJobs {
	s.InstanceSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetReserveAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ReserveAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetResourceGroupName(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetStartTime(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetStatus(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetTargetSize(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetTotalAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.TotalAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetType(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.Type = &v
	return s
}

type DescribeElasticPlanJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanJobsResponse) SetStatusCode(v int32) *DescribeElasticPlanJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanJobsResponse) SetBody(v *DescribeElasticPlanJobsResponseBody) *DescribeElasticPlanJobsResponse {
	s.Body = v
	return s
}

type DescribeElasticPlanSpecificationsRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the ID of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// > *   This parameter is required only when you query the resource specifications that can be scaled for an interactive resource group.
	// > *   You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of a resource group within a specific cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   EXECUTOR: interactive resource groups, which fall into the computing resource category.
	// *   WORKER: EIUs.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanSpecificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsRequest) SetDBClusterId(v string) *DescribeElasticPlanSpecificationsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsRequest) SetResourceGroupName(v string) *DescribeElasticPlanSpecificationsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsRequest) SetType(v string) *DescribeElasticPlanSpecificationsRequest {
	s.Type = &v
	return s
}

type DescribeElasticPlanSpecificationsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of resource specifications returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource specifications that can be scaled.
	Specifications []*string `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	// The number of resource specifications that can be scaled.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlanSpecificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetPageNumber(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetPageSize(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetRequestId(v string) *DescribeElasticPlanSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetSpecifications(v []*string) *DescribeElasticPlanSpecificationsResponseBody {
	s.Specifications = v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetTotalCount(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeElasticPlanSpecificationsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanSpecificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponse) SetStatusCode(v int32) *DescribeElasticPlanSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponse) SetBody(v *DescribeElasticPlanSpecificationsResponseBody) *DescribeElasticPlanSpecificationsResponse {
	s.Body = v
	return s
}

type DescribeElasticPlansRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// > If you do not specify this parameter, all scaling plans are queried.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the scaling plan was immediately enabled after the plan is created. Valid values:
	//
	// *   true
	// *   false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource group.
	//
	// > *   If you do not specify this parameter, the scaling plans of all resource groups are queried, covering the interactive resource group type and the elastic I/O unit (EIU) type.
	// >*   You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of a resource group within a cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   EXECUTOR: interactive resource groups, which fall into the computing resource category.
	// *   WORKER: EIUs.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansRequest) SetDBClusterId(v string) *DescribeElasticPlansRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetElasticPlanName(v string) *DescribeElasticPlansRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetEnabled(v bool) *DescribeElasticPlansRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetPageNumber(v int32) *DescribeElasticPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetPageSize(v int32) *DescribeElasticPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetResourceGroupName(v string) *DescribeElasticPlansRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetType(v string) *DescribeElasticPlansRequest {
	s.Type = &v
	return s
}

type DescribeElasticPlansResponseBody struct {
	// The scaling plans.
	ElasticPlans []*DescribeElasticPlansResponseBodyElasticPlans `json:"ElasticPlans,omitempty" xml:"ElasticPlans,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponseBody) SetElasticPlans(v []*DescribeElasticPlansResponseBodyElasticPlans) *DescribeElasticPlansResponseBody {
	s.ElasticPlans = v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetPageNumber(v int32) *DescribeElasticPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetPageSize(v int32) *DescribeElasticPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetRequestId(v string) *DescribeElasticPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetTotalCount(v int32) *DescribeElasticPlansResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeElasticPlansResponseBodyElasticPlans struct {
	// Indicates whether **Proportional Default Scaling for EIUs** is enabled. Valid values:
	//
	// *   true
	// *   false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The name of the scaling plan.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the scaling plan was immediately enabled after the plan is created. Valid values:
	//
	// *   true
	// *   false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when the next scheduling is performed.
	//
	// > The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	NextScheduleTime *string `json:"NextScheduleTime,omitempty" xml:"NextScheduleTime,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of a resource group within a cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The amount of elastic resources after scaling.
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   EXECUTOR: interactive resource group.
	// *   WORKER: EIU.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlansResponseBodyElasticPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlansResponseBodyElasticPlans) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetAutoScale(v bool) *DescribeElasticPlansResponseBodyElasticPlans {
	s.AutoScale = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetElasticPlanName(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetEnabled(v bool) *DescribeElasticPlansResponseBodyElasticPlans {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetNextScheduleTime(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.NextScheduleTime = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetResourceGroupName(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetTargetSize(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetType(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.Type = &v
	return s
}

type DescribeElasticPlansResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponse) SetHeaders(v map[string]*string) *DescribeElasticPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlansResponse) SetStatusCode(v int32) *DescribeElasticPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlansResponse) SetBody(v *DescribeElasticPlansResponseBody) *DescribeElasticPlansResponse {
	s.Body = v
	return s
}

type DescribeEnabledPrivilegesRequest struct {
	// The name of the database account.
	//
	// >  You can call the [DescribeAccounts](~~612430~~) operation to query the information about database accounts for a cluster, including the account name.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnabledPrivilegesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnabledPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesRequest) SetAccountName(v string) *DescribeEnabledPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeEnabledPrivilegesRequest) SetDBClusterId(v string) *DescribeEnabledPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEnabledPrivilegesRequest) SetRegionId(v string) *DescribeEnabledPrivilegesRequest {
	s.RegionId = &v
	return s
}

type DescribeEnabledPrivilegesResponseBody struct {
	// The queried permission level and permissions.
	Data []*DescribeEnabledPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBody) SetData(v []*DescribeEnabledPrivilegesResponseBodyData) *DescribeEnabledPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBody) SetRequestId(v string) *DescribeEnabledPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnabledPrivilegesResponseBodyData struct {
	// The description of the permission level.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The queried permissions.
	Privileges []*DescribeEnabledPrivilegesResponseBodyDataPrivileges `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
	// The permission level.
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetDescription(v string) *DescribeEnabledPrivilegesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetPrivileges(v []*DescribeEnabledPrivilegesResponseBodyDataPrivileges) *DescribeEnabledPrivilegesResponseBodyData {
	s.Privileges = v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyData) SetScope(v string) *DescribeEnabledPrivilegesResponseBodyData {
	s.Scope = &v
	return s
}

type DescribeEnabledPrivilegesResponseBodyDataPrivileges struct {
	// The description of the permission.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeEnabledPrivilegesResponseBodyDataPrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponseBodyDataPrivileges) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) SetDescription(v string) *DescribeEnabledPrivilegesResponseBodyDataPrivileges {
	s.Description = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponseBodyDataPrivileges) SetKey(v string) *DescribeEnabledPrivilegesResponseBodyDataPrivileges {
	s.Key = &v
	return s
}

type DescribeEnabledPrivilegesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnabledPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnabledPrivilegesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnabledPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeEnabledPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnabledPrivilegesResponse) SetStatusCode(v int32) *DescribeEnabledPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnabledPrivilegesResponse) SetBody(v *DescribeEnabledPrivilegesResponseBody) *DescribeEnabledPrivilegesResponse {
	s.Body = v
	return s
}

type DescribeJobResourceUsageRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeJobResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageRequest) SetDBClusterId(v string) *DescribeJobResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetEndTime(v string) *DescribeJobResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetStartTime(v string) *DescribeJobResourceUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeJobResourceUsageResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeJobResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBody) SetCode(v int32) *DescribeJobResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBody) SetData(v *DescribeJobResourceUsageResponseBodyData) *DescribeJobResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobResourceUsageResponseBody) SetRequestId(v string) *DescribeJobResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeJobResourceUsageResponseBodyData struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The AnalyticDB compute unit (ACU) usage of the job resource group.
	JobAcuUsage []*DescribeJobResourceUsageResponseBodyDataJobAcuUsage `json:"JobAcuUsage,omitempty" xml:"JobAcuUsage,omitempty" type:"Repeated"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeJobResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetEndTime(v string) *DescribeJobResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetJobAcuUsage(v []*DescribeJobResourceUsageResponseBodyDataJobAcuUsage) *DescribeJobResourceUsageResponseBodyData {
	s.JobAcuUsage = v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetStartTime(v string) *DescribeJobResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeJobResourceUsageResponseBodyDataJobAcuUsage struct {
	// The ACU usage.
	AcuUsageDetail *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail `json:"AcuUsageDetail,omitempty" xml:"AcuUsageDetail,omitempty" type:"Struct"`
	// The end time of the job. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// The job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The start time of the job. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	JobStartTime *string `json:"JobStartTime,omitempty" xml:"JobStartTime,omitempty"`
	// The name of the job resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetAcuUsageDetail(v *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.AcuUsageDetail = v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobEndTime(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobEndTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobId(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobId = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobStartTime(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobStartTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetResourceGroupName(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.ResourceGroupName = &v
	return s
}

type DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail struct {
	// The number of ACUs for the elastic resources.
	ElasticAcuNumber *float32 `json:"ElasticAcuNumber,omitempty" xml:"ElasticAcuNumber,omitempty"`
	// The number of ACUs for the reserved resources.
	ReservedAcuNumber *float32 `json:"ReservedAcuNumber,omitempty" xml:"ReservedAcuNumber,omitempty"`
	// The total number of ACUs.
	TotalAcuNumber *float32 `json:"TotalAcuNumber,omitempty" xml:"TotalAcuNumber,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetElasticAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.ElasticAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetReservedAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.ReservedAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetTotalAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.TotalAcuNumber = &v
	return s
}

type DescribeJobResourceUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeJobResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResourceUsageResponse) SetStatusCode(v int32) *DescribeJobResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobResourceUsageResponse) SetBody(v *DescribeJobResourceUsageResponseBody) *DescribeJobResourceUsageResponse {
	s.Body = v
	return s
}

type DescribePatternPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SQL pattern ID.
	//
	// >  You can call the [DescribeSQLPatterns](~~321868~~) operation to query the information about all SQL patterns in an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster within a period of time, including SQL pattern IDs.
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// >
	//
	// *   If the current date is August 22, 2022 (UTC+8), you can query the data of August 9, 2022 (2022-08-08T16:00:00Z) to the earliest extent. If you want to query the data that is earlier than August 9, 2022 (2022-08-08T16:00:00Z), null is returned.
	//
	// *   The maximum time range that can be specified is 24 hours.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceRequest) SetDBClusterId(v string) *DescribePatternPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetEndTime(v string) *DescribePatternPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetPatternId(v string) *DescribePatternPerformanceRequest {
	s.PatternId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetRegionId(v string) *DescribePatternPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetStartTime(v string) *DescribePatternPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBody struct {
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribePatternPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBody) SetEndTime(v string) *DescribePatternPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetPerformances(v []*DescribePatternPerformanceResponseBodyPerformances) *DescribePatternPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetRequestId(v string) *DescribePatternPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetStartTime(v string) *DescribePatternPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformances struct {
	// The queried performance metric. Valid values:
	//
	// *   **AnalyticDB_PatternQueryCount**: the total number of queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternQueryTime**: the total amount of time consumed by the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternExecutionTime**: the execution duration of the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternPeakMemory**: the peak memory usage of the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternScanSize**: the amount of data scanned in the queries executed in association with the SQL pattern.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the performance metrics.
	Series []*DescribePatternPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric. Valid values:
	//
	// *   If the performance metric is related to the query time (the value of `Key` is `AnalyticDB_PatternQueryTime` or `AnalyticDB_PatternExecutionTime`), **ms** is returned.
	// *   If the performance metric is related to the peak memory usage (the value of `Key` is `AnalyticDB_PatternPeakMemory`), **MB** is returned.
	// *   If the performance metric is related to the amount of data scanned (the value of `Key` is `AnalyticDB_PatternScanSize`), **MB** is returned.
	// *   If the performance metric is related to the number of queries (the value of `Key` is `AnalyticDB_PatternQueryCount`), null is returned.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribePatternPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetKey(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetSeries(v []*DescribePatternPerformanceResponseBodyPerformancesSeries) *DescribePatternPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetUnit(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value. Valid values:
	//
	// *   If the value of `Key` is `AnalyticDB_PatternQueryCount`, `pattern_query_count` is returned, which indicates the number of executions of the SQL statements in association with the SQL pattern.
	//
	// *   If the value of `Key` is `AnalyticDB_PatternQueryTime`, the following values are returned:
	//
	//     *   `average_query_time`, which indicates the average total amount of time consumed by the SQL statements in association with the SQL pattern.
	//     *   `max_query_time`, which indicates the maximum total amount of time consumed by the SQL statements in association with the SQL pattern.
	//
	// *   If the value of `Key` is `AnalyticDB_PatternExecutionTime`, the following values are returned:
	//
	//     *   `average_execution_time`, which indicates the average execution duration of the SQL statements in association with the SQL pattern.
	//     *   `max_execution_time`, which indicates the maximum execution duration of the SQL statements in association with the SQL pattern.
	//
	// *   If the value of `Key` is `AnalyticDB_PatternPeakMemory`, the following values are returned:
	//
	//     *   `average_peak_memory`, which indicates the average peak memory usage of the SQL statements in association with the SQL pattern.
	//     *   `max_peak_memory`, which indicates the maximum peak memory usage of the SQL statements in association with the SQL pattern.
	//
	// *   If the value of `Key` is `AnalyticDB_PatternScanSize`, the following values are returned:
	//
	//     *   `average_scan_size`, which indicates the average amount of data scanned by the SQL statements in association with the SQL pattern.
	//     *   `max_scan_size`, which indicates the maximum amount of data scanned by the SQL statements in association with the SQL pattern.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the performance metric.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribePatternPerformanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePatternPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePatternPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponse) SetHeaders(v map[string]*string) *DescribePatternPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribePatternPerformanceResponse) SetStatusCode(v int32) *DescribePatternPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePatternPerformanceResponse) SetBody(v *DescribePatternPerformanceResponseBody) *DescribePatternPerformanceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The language used for the region and zone names specified by the LocalName parameter. Default value: zh-CN. Valid values:
	//
	// *   **zh-CN**: simplified Chinese
	// *   **en-US**: English
	// *   **ja**: Japanese
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// Details of the regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Details of the zones.
	Zones *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	// The name of the zone.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Indicates whether Virtual Private Cloud (VPC) is supported in the zone. Valid values:
	//
	// *   **true**: VPC is supported.
	// *   **false**: VPC is not supported.
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSQLPatternsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"AverageQueryTime","Type":"Asc"}]`.
	//
	// *   `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     *   `PatternCreationTime`: the earliest commit time of the SQL pattern within the time range to query.
	//     *   `AverageQueryTime`: the average total amount of time consumed by the SQL pattern within the time range to query.
	//     *   `MaxQueryTime`: the maximum total amount of time consumed by the SQL pattern within the time range to query.
	//     *   `AverageExecutionTime`: the average execution duration of the SQL pattern within the time range to query.
	//     *   `MaxExecutionTime`: the maximum execution duration of the SQL pattern within the time range to query.
	//     *   `AveragePeakMemory`: the average peak memory usage of the SQL pattern within the time range to query.
	//     *   `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query.
	//     *   `AverageScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query.
	//     *   `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query.
	//     *   `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//     *   `FailedCount`: the number of failed queries performed in association with the SQL pattern within the time range to query.
	//
	// *   `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **10** (default)
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// > *   Only data within the last 14 days can be queried.
	// > * The maximum time range that can be specified is 24 hours.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsRequest) SetDBClusterId(v string) *DescribeSQLPatternsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetEndTime(v string) *DescribeSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetKeyword(v string) *DescribeSQLPatternsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetLang(v string) *DescribeSQLPatternsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetOrder(v string) *DescribeSQLPatternsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageNumber(v int32) *DescribeSQLPatternsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageSize(v int32) *DescribeSQLPatternsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetRegionId(v string) *DescribeSQLPatternsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetStartTime(v string) *DescribeSQLPatternsRequest {
	s.StartTime = &v
	return s
}

type DescribeSQLPatternsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried SQL pattern.
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSQLPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBody) SetPageNumber(v int32) *DescribeSQLPatternsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPageSize(v int32) *DescribeSQLPatternsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPatternDetails(v []*DescribeSQLPatternsResponseBodyPatternDetails) *DescribeSQLPatternsResponseBody {
	s.PatternDetails = v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetRequestId(v string) *DescribeSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetTotalCount(v int32) *DescribeSQLPatternsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSQLPatternsResponseBodyPatternDetails struct {
	// The IP address of the SQL client that commits the SQL pattern.
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The average execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	AverageExecutionTime *float64 `json:"AverageExecutionTime,omitempty" xml:"AverageExecutionTime,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	AveragePeakMemory *float64 `json:"AveragePeakMemory,omitempty" xml:"AveragePeakMemory,omitempty"`
	// The average total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	AverageQueryTime *float64 `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	AverageScanSize *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	// Indicates whether the execution of the SQL pattern can be blocked. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// >  Only SELECT and INSERT statements can be blocked.
	Blockable *bool `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	// The number of failed queries executed in association with the SQL pattern within the query time range.
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The maximum execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	MaxQueryTime *int64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The earliest commit time of the SQL pattern within the query time range.
	PatternCreationTime *string `json:"PatternCreationTime,omitempty" xml:"PatternCreationTime,omitempty"`
	// The ID of the SQL pattern.
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The number of queries executed in association with the SQL pattern within the query time range.
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The statement of the SQL pattern.
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The tables scanned based on the SQL pattern.
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The database username that is used to commit the SQL pattern.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAccessIp(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AccessIp = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageExecutionTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAveragePeakMemory(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AveragePeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageQueryTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageScanSize(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetBlockable(v bool) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Blockable = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetFailedCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.FailedCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxExecutionTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxPeakMemory(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxQueryTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxScanSize(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternCreationTime(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternCreationTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternId(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetSQLPattern(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.SQLPattern = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetTables(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Tables = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetUser(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.User = &v
	return s
}

type DescribeSQLPatternsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPatternsResponse) SetStatusCode(v int32) *DescribeSQLPatternsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPatternsResponse) SetBody(v *DescribeSQLPatternsResponseBody) *DescribeSQLPatternsResponse {
	s.Body = v
	return s
}

type DescribeSchemasRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasRequest) SetRegionId(v string) *DescribeSchemasRequest {
	s.RegionId = &v
	return s
}

type DescribeSchemasResponseBody struct {
	// The queried databases.
	Items *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBody) SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSchemasResponseBody) SetRequestId(v string) *DescribeSchemasResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSchemasResponseBodyItems struct {
	Schema []*DescribeSchemasResponseBodyItemsSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeSchemasResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItems) SetSchema(v []*DescribeSchemasResponseBodyItemsSchema) *DescribeSchemasResponseBodyItems {
	s.Schema = v
	return s
}

type DescribeSchemasResponseBodyItemsSchema struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeSchemasResponseBodyItemsSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItemsSchema) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetDBClusterId(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetSchemaName(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.SchemaName = &v
	return s
}

type DescribeSchemasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponse) SetHeaders(v map[string]*string) *DescribeSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchemasResponse) SetStatusCode(v int32) *DescribeSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSchemasResponse) SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse {
	s.Body = v
	return s
}

type DescribeSparkCodeLogRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Spark job.
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkCodeLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogRequest) SetDBClusterId(v string) *DescribeSparkCodeLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeLogRequest) SetJobId(v int64) *DescribeSparkCodeLogRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeLogRequest) SetRegionId(v string) *DescribeSparkCodeLogRequest {
	s.RegionId = &v
	return s
}

type DescribeSparkCodeLogResponseBody struct {
	// The content of the log.
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The returned message.
	//
	// *   If the request was successful, **Success** is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSparkCodeLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogResponseBody) SetLog(v string) *DescribeSparkCodeLogResponseBody {
	s.Log = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetMessage(v string) *DescribeSparkCodeLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetRequestId(v string) *DescribeSparkCodeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetSuccess(v bool) *DescribeSparkCodeLogResponseBody {
	s.Success = &v
	return s
}

type DescribeSparkCodeLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeLogResponse) SetStatusCode(v int32) *DescribeSparkCodeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeLogResponse) SetBody(v *DescribeSparkCodeLogResponseBody) *DescribeSparkCodeLogResponse {
	s.Body = v
	return s
}

type DescribeSparkCodeOutputRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Spark job.
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkCodeOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputRequest) SetDBClusterId(v string) *DescribeSparkCodeOutputRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeOutputRequest) SetJobId(v int64) *DescribeSparkCodeOutputRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeOutputRequest) SetRegionId(v string) *DescribeSparkCodeOutputRequest {
	s.RegionId = &v
	return s
}

type DescribeSparkCodeOutputResponseBody struct {
	// The returned message.
	//
	// *   If the request was successful, **Success** is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result, which is in the format of JSON objects.
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSparkCodeOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputResponseBody) SetMessage(v string) *DescribeSparkCodeOutputResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetOutput(v string) *DescribeSparkCodeOutputResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetRequestId(v string) *DescribeSparkCodeOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetSuccess(v bool) *DescribeSparkCodeOutputResponseBody {
	s.Success = &v
	return s
}

type DescribeSparkCodeOutputResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeOutputResponse) SetStatusCode(v int32) *DescribeSparkCodeOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeOutputResponse) SetBody(v *DescribeSparkCodeOutputResponseBody) *DescribeSparkCodeOutputResponse {
	s.Body = v
	return s
}

type DescribeSparkCodeWebUiRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Spark job.
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkCodeWebUiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeWebUiRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiRequest) SetDBClusterId(v string) *DescribeSparkCodeWebUiRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeWebUiRequest) SetJobId(v int64) *DescribeSparkCodeWebUiRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeWebUiRequest) SetRegionId(v string) *DescribeSparkCodeWebUiRequest {
	s.RegionId = &v
	return s
}

type DescribeSparkCodeWebUiResponseBody struct {
	// The returned message.
	//
	// *   If the request was successful, **SUCCESS** is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**
	// *   **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The URL of the web UI for the Spark application.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeSparkCodeWebUiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeWebUiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiResponseBody) SetMessage(v string) *DescribeSparkCodeWebUiResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetRequestId(v string) *DescribeSparkCodeWebUiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetSuccess(v bool) *DescribeSparkCodeWebUiResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetUrl(v string) *DescribeSparkCodeWebUiResponseBody {
	s.Url = &v
	return s
}

type DescribeSparkCodeWebUiResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeWebUiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeWebUiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSparkCodeWebUiResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeWebUiResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeWebUiResponse) SetStatusCode(v int32) *DescribeSparkCodeWebUiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponse) SetBody(v *DescribeSparkCodeWebUiResponseBody) *DescribeSparkCodeWebUiResponse {
	s.Body = v
	return s
}

type DescribeSqlPatternRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"Pattern","Type":"Asc"}]`. Parameters:
	//
	// *   `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     *   `Pattern`: the SQL pattern.
	//     *   `AccessIP`: the IP address of the client.
	//     *   `User`: the username.
	//     *   `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//     *   `AvgPeakMemory`: the average peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//     *   `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//     *   `AvgCpuTime`: the average execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//     *   `MaxCpuTime`: the maximum execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//     *   `AvgStageCount`: the average number of stages.
	//     *   `MaxStageCount`: the maximum number of stages.
	//     *   `AvgTaskCount`: the average number of tasks.
	//     *   `MaxTaskCount`: the maximum number of tasks.
	//     *   `AvgScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//     *   `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//
	// *   `Type` specifies the sorting order. Valid values:
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	//
	// >
	//
	// *   If you do not specify this parameter, query results are sorted in ascending order of `Pattern`.
	//
	// *   If you want to sort query results by `AccessIP`, you must set the `Type` parameter to `accessip`. If you want to sort query results by `User`, you must leave the `Type` parameter empty or set it to `user`.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **10** (default)
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that is used for the query.
	//
	// > If you do not specify this parameter, all SQL patterns of the AnalyticDB for MySQL cluster within the time period specified by `StartTime` are returned.
	SqlPattern *string `json:"SqlPattern,omitempty" xml:"SqlPattern,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd format. The time must be in UTC.
	//
	// > Only data within the last 30 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The dimension by which to aggregate the SQL patterns. Valid values:
	//
	// *   `user`: aggregates the SQL patterns by user.
	// *   `accessip`: aggregates the SQL patterns by client IP address.
	//
	// > If you do not specify this parameter, the SQL patterns are aggregated by `user`.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSqlPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternRequest) SetDBClusterId(v string) *DescribeSqlPatternRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetOrder(v string) *DescribeSqlPatternRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageNumber(v int32) *DescribeSqlPatternRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageSize(v int32) *DescribeSqlPatternRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetRegionId(v string) *DescribeSqlPatternRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetSqlPattern(v string) *DescribeSqlPatternRequest {
	s.SqlPattern = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetStartTime(v string) *DescribeSqlPatternRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetType(v string) *DescribeSqlPatternRequest {
	s.Type = &v
	return s
}

type DescribeSqlPatternResponseBody struct {
	// The queried SQL pattern.
	Items []*DescribeSqlPatternResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBody) SetItems(v []*DescribeSqlPatternResponseBodyItems) *DescribeSqlPatternResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageNumber(v int32) *DescribeSqlPatternResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageSize(v int32) *DescribeSqlPatternResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetRequestId(v string) *DescribeSqlPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetTotalCount(v int32) *DescribeSqlPatternResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSqlPatternResponseBodyItems struct {
	// The IP address of the client.
	//
	// >  This parameter is returned only when **Type** is set to **accessip**.
	AccessIP *string `json:"AccessIP,omitempty" xml:"AccessIP,omitempty"`
	// The average execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	AvgCpuTime *string `json:"AvgCpuTime,omitempty" xml:"AvgCpuTime,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: KB.
	AvgPeakMemory *string `json:"AvgPeakMemory,omitempty" xml:"AvgPeakMemory,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	AvgScanSize *string `json:"AvgScanSize,omitempty" xml:"AvgScanSize,omitempty"`
	// The average number of scanned rows.
	AvgStageCount *string `json:"AvgStageCount,omitempty" xml:"AvgStageCount,omitempty"`
	// The average number of tasks.
	AvgTaskCount *string `json:"AvgTaskCount,omitempty" xml:"AvgTaskCount,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	MaxCpuTime *string `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: KB.
	MaxPeakMemory *string `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	MaxScanSize *string `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The maximum number of stages.
	MaxStageCount *string `json:"MaxStageCount,omitempty" xml:"MaxStageCount,omitempty"`
	// The maximum number of tasks.
	MaxTaskCount *string `json:"MaxTaskCount,omitempty" xml:"MaxTaskCount,omitempty"`
	// The SQL pattern.
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The number of queries performed in association with the SQL pattern within the query time range.
	QueryCount *string `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The start date of the query.
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The username.
	//
	// >  This parameter is returned only when **Type** is left empty or set to **user**.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSqlPatternResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBodyItems) SetAccessIP(v string) *DescribeSqlPatternResponseBodyItems {
	s.AccessIP = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetInstanceName(v string) *DescribeSqlPatternResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetPattern(v string) *DescribeSqlPatternResponseBodyItems {
	s.Pattern = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetQueryCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.QueryCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetReportDate(v string) *DescribeSqlPatternResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetUser(v string) *DescribeSqlPatternResponseBodyItems {
	s.User = &v
	return s
}

type DescribeSqlPatternResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternResponse) SetStatusCode(v int32) *DescribeSqlPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlPatternResponse) SetBody(v *DescribeSqlPatternResponseBody) *DescribeSqlPatternResponse {
	s.Body = v
	return s
}

type DescribeStorageResourceUsageRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStorageResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageRequest) SetDBClusterId(v string) *DescribeStorageResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeStorageResourceUsageRequest) SetEndTime(v string) *DescribeStorageResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStorageResourceUsageRequest) SetStartTime(v string) *DescribeStorageResourceUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeStorageResourceUsageResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeStorageResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStorageResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBody) SetCode(v int32) *DescribeStorageResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBody) SetData(v *DescribeStorageResourceUsageResponseBodyData) *DescribeStorageResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStorageResourceUsageResponseBody) SetRequestId(v string) *DescribeStorageResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeStorageResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeStorageResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStorageResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeStorageResourceUsageResponseBodyDataAcuInfo) *DescribeStorageResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetEndTime(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetStartTime(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeStorageResourceUsageResponseBodyDataAcuInfo struct {
	// The resource usage metric. Valid values:
	//
	// *   `TotalAcuNumber`: the total number of ACUs.
	// *   `ReservedAcuNumber`: the number of ACUs for the reserved resources.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the metric at specific points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeStorageResourceUsageResponseBodyDataAcuInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeStorageResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeStorageResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

type DescribeStorageResourceUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeStorageResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResourceUsageResponse) SetStatusCode(v int32) *DescribeStorageResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageResourceUsageResponse) SetBody(v *DescribeStorageResourceUsageResponseBody) *DescribeStorageResourceUsageResponse {
	s.Body = v
	return s
}

type DescribeTableAccessCountRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"TableSchema","Type":"Asc"}]`. Fields in the request parameter:
	//
	// *   `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     *   `TableSchema`: the name of the database to which the table belongs.
	//     *   `TableName`: the name of the table.
	//     *   `AccessCount`: the number of accesses to the table.
	//
	// *   `Type` specifies the sorting order. Valid values:
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	//
	// >  If you do not specify this parameter, query results are sorted in ascending order based on the database and the table.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **10** (default)
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  Only data within the last 30 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table.
	//
	// >  If you leave this parameter empty, the number of accesses to all tables in the cluster on a date is returned.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableAccessCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountRequest) SetDBClusterId(v string) *DescribeTableAccessCountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetOrder(v string) *DescribeTableAccessCountRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageNumber(v int32) *DescribeTableAccessCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageSize(v int32) *DescribeTableAccessCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetRegionId(v string) *DescribeTableAccessCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetStartTime(v string) *DescribeTableAccessCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetTableName(v string) *DescribeTableAccessCountRequest {
	s.TableName = &v
	return s
}

type DescribeTableAccessCountResponseBody struct {
	// The queried tables.
	Items []*DescribeTableAccessCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableAccessCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBody) SetItems(v []*DescribeTableAccessCountResponseBodyItems) *DescribeTableAccessCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageNumber(v int32) *DescribeTableAccessCountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageSize(v int32) *DescribeTableAccessCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetRequestId(v string) *DescribeTableAccessCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetTotalCount(v int32) *DescribeTableAccessCountResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTableAccessCountResponseBodyItems struct {
	// The number of accesses to the table.
	AccessCount *string `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The ID of the cluster to which the table belongs.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The date when the table was accessed.
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The database to which the table belongs.
	TableSchema *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
}

func (s DescribeTableAccessCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBodyItems) SetAccessCount(v string) *DescribeTableAccessCountResponseBodyItems {
	s.AccessCount = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetInstanceName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetReportDate(v string) *DescribeTableAccessCountResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableSchema(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableSchema = &v
	return s
}

type DescribeTableAccessCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableAccessCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableAccessCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponse) SetHeaders(v map[string]*string) *DescribeTableAccessCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableAccessCountResponse) SetStatusCode(v int32) *DescribeTableAccessCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableAccessCountResponse) SetBody(v *DescribeTableAccessCountResponseBody) *DescribeTableAccessCountResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

type DescribeTablesResponseBody struct {
	// The queried tables.
	Items *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTablesResponseBodyItems struct {
	Table []*DescribeTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) SetTable(v []*DescribeTablesResponseBodyItemsTable) *DescribeTablesResponseBodyItems {
	s.Table = v
	return s
}

type DescribeTablesResponseBodyItemsTable struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsTable) SetDBClusterId(v string) *DescribeTablesResponseBodyItemsTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetSchemaName(v string) *DescribeTablesResponseBodyItemsTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetTableName(v string) *DescribeTablesResponseBodyItemsTable {
	s.TableName = &v
	return s
}

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeUserQuotaRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaRequest) SetDBClusterId(v string) *DescribeUserQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeUserQuotaRequest) SetRegionId(v string) *DescribeUserQuotaRequest {
	s.RegionId = &v
	return s
}

type DescribeUserQuotaResponseBody struct {
	// The available elastic AnalyticDB compute units (ACUs).
	ElasticACU *string `json:"ElasticACU,omitempty" xml:"ElasticACU,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The available reserved computing resources.
	ReserverdCompteACU *string `json:"ReserverdCompteACU,omitempty" xml:"ReserverdCompteACU,omitempty"`
	// The available reserved storage resources.
	ReserverdStorageACU *string `json:"ReserverdStorageACU,omitempty" xml:"ReserverdStorageACU,omitempty"`
	// The number of available resource groups.
	ResourceGroupCount *string `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
}

func (s DescribeUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBody) SetElasticACU(v string) *DescribeUserQuotaResponseBody {
	s.ElasticACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetRequestId(v string) *DescribeUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetReserverdCompteACU(v string) *DescribeUserQuotaResponseBody {
	s.ReserverdCompteACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetReserverdStorageACU(v string) *DescribeUserQuotaResponseBody {
	s.ReserverdStorageACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetResourceGroupCount(v string) *DescribeUserQuotaResponseBody {
	s.ResourceGroupCount = &v
	return s
}

type DescribeUserQuotaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeUserQuotaResponse) SetStatusCode(v int32) *DescribeUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserQuotaResponse) SetBody(v *DescribeUserQuotaResponseBody) *DescribeUserQuotaResponse {
	s.Body = v
	return s
}

type DetachUserENIRequest struct {
	// The instance ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DetachUserENIRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIRequest) GoString() string {
	return s.String()
}

func (s *DetachUserENIRequest) SetDBClusterId(v string) *DetachUserENIRequest {
	s.DBClusterId = &v
	return s
}

type DetachUserENIResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachUserENIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponseBody) SetRequestId(v string) *DetachUserENIResponseBody {
	s.RequestId = &v
	return s
}

type DetachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachUserENIResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIResponse) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponse) SetHeaders(v map[string]*string) *DetachUserENIResponse {
	s.Headers = v
	return s
}

func (s *DetachUserENIResponse) SetStatusCode(v int32) *DetachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachUserENIResponse) SetBody(v *DetachUserENIResponseBody) *DetachUserENIResponse {
	s.Body = v
	return s
}

type DisableElasticPlanRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s DisableElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableElasticPlanRequest) SetDBClusterId(v string) *DisableElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableElasticPlanRequest) SetElasticPlanName(v string) *DisableElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

type DisableElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableElasticPlanResponseBody) SetRequestId(v string) *DisableElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DisableElasticPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableElasticPlanResponse) SetHeaders(v map[string]*string) *DisableElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableElasticPlanResponse) SetStatusCode(v int32) *DisableElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableElasticPlanResponse) SetBody(v *DisableElasticPlanResponseBody) *DisableElasticPlanResponse {
	s.Body = v
	return s
}

type DownloadDiagnosisRecordsRequest struct {
	// The source IP address.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statements are executed.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query keyword of the SQL statements.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language. Valid values:
	//
	// *   **zh**: simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statements. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the SQL statements. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statements. Unit: bytes.
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statements. Unit: bytes.
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, `Min`, and `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries the executed SQL statements. You can set `Value` to `running` to query the SQL statements that are being executed. You can also set Value to `failed` to query the SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries the SQL statements whose execution duration is in the range of 10 to 200 milliseconds. You can also specify custom values for the Min and Max fields.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statements belong.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  You can query data only within the last 14 days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to execute the SQL statements.
	//
	// >  You can call the [DescribeDiagnosisDimensions](~~~~) operation to query the resource groups, database names, usernames, and source IP addresses of the SQL statements that meet a query condition.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) SetClientIp(v string) *DownloadDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDBClusterId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetKeyword(v string) *DownloadDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetRegionId(v string) *DownloadDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroup(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUserName(v string) *DownloadDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DownloadDiagnosisRecordsResponseBody struct {
	// The download ID.
	DownloadId *int32 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v int32) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type EnableElasticPlanRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s EnableElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *EnableElasticPlanRequest) SetDBClusterId(v string) *EnableElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *EnableElasticPlanRequest) SetElasticPlanName(v string) *EnableElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

type EnableElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *EnableElasticPlanResponseBody) SetRequestId(v string) *EnableElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type EnableElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *EnableElasticPlanResponse) SetHeaders(v map[string]*string) *EnableElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *EnableElasticPlanResponse) SetStatusCode(v int32) *EnableElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableElasticPlanResponse) SetBody(v *EnableElasticPlanResponseBody) *EnableElasticPlanResponse {
	s.Body = v
	return s
}

type ExistRunningSQLEngineRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of the resource group for a cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ExistRunningSQLEngineRequest) String() string {
	return tea.Prettify(s)
}

func (s ExistRunningSQLEngineRequest) GoString() string {
	return s.String()
}

func (s *ExistRunningSQLEngineRequest) SetDBClusterId(v string) *ExistRunningSQLEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *ExistRunningSQLEngineRequest) SetResourceGroupName(v string) *ExistRunningSQLEngineRequest {
	s.ResourceGroupName = &v
	return s
}

type ExistRunningSQLEngineResponseBody struct {
	// Indicates whether a running SQL engine exists in the resource group.
	//
	// Valid values:
	//
	// *   **True**
	// *   **False**
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExistRunningSQLEngineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExistRunningSQLEngineResponseBody) GoString() string {
	return s.String()
}

func (s *ExistRunningSQLEngineResponseBody) SetData(v bool) *ExistRunningSQLEngineResponseBody {
	s.Data = &v
	return s
}

func (s *ExistRunningSQLEngineResponseBody) SetRequestId(v string) *ExistRunningSQLEngineResponseBody {
	s.RequestId = &v
	return s
}

type ExistRunningSQLEngineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExistRunningSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExistRunningSQLEngineResponse) String() string {
	return tea.Prettify(s)
}

func (s ExistRunningSQLEngineResponse) GoString() string {
	return s.String()
}

func (s *ExistRunningSQLEngineResponse) SetHeaders(v map[string]*string) *ExistRunningSQLEngineResponse {
	s.Headers = v
	return s
}

func (s *ExistRunningSQLEngineResponse) SetStatusCode(v int32) *ExistRunningSQLEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *ExistRunningSQLEngineResponse) SetBody(v *ExistRunningSQLEngineResponseBody) *ExistRunningSQLEngineResponse {
	s.Body = v
	return s
}

type GetDatabaseObjectsRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The owner of the database.
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the database.
	FilterSchemaName *string `json:"FilterSchemaName,omitempty" xml:"FilterSchemaName,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   Asc
	// *   Desc
	//
	// Valid values for Field: DatabaseName, CreateTime, and UpdateTime. -CreateTime; -UpdateTime;
	//
	// Default value: {"Type": "Desc","Field": "DatabaseName"}.
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the database.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDatabaseObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsRequest) SetDBClusterId(v string) *GetDatabaseObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetFilterOwner(v string) *GetDatabaseObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetFilterSchemaName(v string) *GetDatabaseObjectsRequest {
	s.FilterSchemaName = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetOrderBy(v string) *GetDatabaseObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetPageNumber(v int64) *GetDatabaseObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetPageSize(v int64) *GetDatabaseObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetRegionId(v string) *GetDatabaseObjectsRequest {
	s.RegionId = &v
	return s
}

type GetDatabaseObjectsResponseBody struct {
	// The queried databases.
	Data *GetDatabaseObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDatabaseObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponseBody) SetData(v *GetDatabaseObjectsResponseBodyData) *GetDatabaseObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetPageNumber(v int64) *GetDatabaseObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetPageSize(v int64) *GetDatabaseObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetRequestId(v string) *GetDatabaseObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetTotalCount(v int64) *GetDatabaseObjectsResponseBody {
	s.TotalCount = &v
	return s
}

type GetDatabaseObjectsResponseBodyData struct {
	// The queried database.
	DatabaseSummaryModels []*DatabaseSummaryModel `json:"DatabaseSummaryModels,omitempty" xml:"DatabaseSummaryModels,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDatabaseObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponseBodyData) SetDatabaseSummaryModels(v []*DatabaseSummaryModel) *GetDatabaseObjectsResponseBodyData {
	s.DatabaseSummaryModels = v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetPageNumber(v int64) *GetDatabaseObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetPageSize(v int64) *GetDatabaseObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetTotalCount(v int64) *GetDatabaseObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetDatabaseObjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponse) SetHeaders(v map[string]*string) *GetDatabaseObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseObjectsResponse) SetStatusCode(v int32) *GetDatabaseObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseObjectsResponse) SetBody(v *GetDatabaseObjectsResponseBody) *GetDatabaseObjectsResponse {
	s.Body = v
	return s
}

type GetSparkAppAttemptLogRequest struct {
	// The ID of the log.
	//
	// > You can call the [ListSparkAppAttempts](~~455887~~) operation to query the information about the retry attempts of a Spark application, including the retry log IDs.
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The number of log entries to return. Valid values: 1 to 500. Default value: 300.
	LogLength *int64 `json:"LogLength,omitempty" xml:"LogLength,omitempty"`
	// The log offset.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSparkAppAttemptLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppAttemptLogRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogRequest) SetAttemptId(v string) *GetSparkAppAttemptLogRequest {
	s.AttemptId = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetLogLength(v int64) *GetSparkAppAttemptLogRequest {
	s.LogLength = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetPageNumber(v int32) *GetSparkAppAttemptLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetPageSize(v string) *GetSparkAppAttemptLogRequest {
	s.PageSize = &v
	return s
}

type GetSparkAppAttemptLogResponseBody struct {
	// The queried log.
	Data *GetSparkAppAttemptLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppAttemptLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppAttemptLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponseBody) SetData(v *GetSparkAppAttemptLogResponseBodyData) *GetSparkAppAttemptLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppAttemptLogResponseBody) SetRequestId(v string) *GetSparkAppAttemptLogResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppAttemptLogResponseBodyData struct {
	// The application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The content of the log.
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The number of log entries. A value of 0 indicates that no valid logs are returned.
	LogSize *int32 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The alert message returned for the request, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSparkAppAttemptLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppAttemptLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetAppId(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetDBClusterId(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetLogContent(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetLogSize(v int32) *GetSparkAppAttemptLogResponseBodyData {
	s.LogSize = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetMessage(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.Message = &v
	return s
}

type GetSparkAppAttemptLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppAttemptLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppAttemptLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppAttemptLogResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponse) SetHeaders(v map[string]*string) *GetSparkAppAttemptLogResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppAttemptLogResponse) SetStatusCode(v int32) *GetSparkAppAttemptLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppAttemptLogResponse) SetBody(v *GetSparkAppAttemptLogResponseBody) *GetSparkAppAttemptLogResponse {
	s.Body = v
	return s
}

type GetSparkAppInfoRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](~~455888~~) operation to query the Spark application IDs.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoRequest) SetAppId(v string) *GetSparkAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppInfoRequest) SetDBClusterId(v string) *GetSparkAppInfoRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkAppInfoResponseBody struct {
	// The queried Spark application. Fields in the response parameter:
	//
	// *   **Data**: the data of the Spark application template.
	// *   **EstimateExecutionCpuTimeInSeconds**: the amount of time that is required to consume CPU resources for running the Spark application. Unit: milliseconds.
	// *   **LogRootPath**: the storage path of log files.
	// *   **LastAttemptId**: the most recent attempt ID.
	// *   **WebUiAddress**: the web UI URL.
	// *   **SubmittedTimeInMillis**: the time when the Spark application was submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	// *   **StartedTimeInMillis**: the time when the Spark application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	// *   **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	// *   **TerminatedTimeInMillis**: the time when the Spark application was terminated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	// *   **DBClusterId**: the ID of the cluster on which the Spark application runs.
	// *   **ResourceGroupName**: the name of the job resource group.
	// *   **DurationInMillis**: the amount of time that is required to run the Spark application. Unit: milliseconds.
	Data *SparkAppInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoResponseBody) SetData(v *SparkAppInfo) *GetSparkAppInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppInfoResponseBody) SetRequestId(v string) *GetSparkAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoResponse) SetHeaders(v map[string]*string) *GetSparkAppInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppInfoResponse) SetStatusCode(v int32) *GetSparkAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppInfoResponse) SetBody(v *GetSparkAppInfoResponseBody) *GetSparkAppInfoResponse {
	s.Body = v
	return s
}

type GetSparkAppLogRequest struct {
	// The Spark application ID.
	//
	// > You can call the [ListSparkApps](~~612475~~) operation to query the Spark application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of log entries to return. Valid values: 1 to 500. Default value: 300.
	LogLength *int64 `json:"LogLength,omitempty" xml:"LogLength,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSparkAppLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppLogRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogRequest) SetAppId(v string) *GetSparkAppLogRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppLogRequest) SetDBClusterId(v string) *GetSparkAppLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppLogRequest) SetLogLength(v int64) *GetSparkAppLogRequest {
	s.LogLength = &v
	return s
}

func (s *GetSparkAppLogRequest) SetPageNumber(v int32) *GetSparkAppLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSparkAppLogRequest) SetPageSize(v int32) *GetSparkAppLogRequest {
	s.PageSize = &v
	return s
}

type GetSparkAppLogResponseBody struct {
	// The queried log.
	Data *GetSparkAppLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponseBody) SetData(v *GetSparkAppLogResponseBodyData) *GetSparkAppLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppLogResponseBody) SetRequestId(v string) *GetSparkAppLogResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppLogResponseBodyData struct {
	// The ID of the Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The content of the log.
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The number of log entries. A value of 0 indicates that no valid logs are returned.
	LogSize *int32 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The alert message returned for the request, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSparkAppLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponseBodyData) SetDBClusterId(v string) *GetSparkAppLogResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetLogContent(v string) *GetSparkAppLogResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetLogSize(v int32) *GetSparkAppLogResponseBodyData {
	s.LogSize = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetMessage(v string) *GetSparkAppLogResponseBodyData {
	s.Message = &v
	return s
}

type GetSparkAppLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppLogResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponse) SetHeaders(v map[string]*string) *GetSparkAppLogResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppLogResponse) SetStatusCode(v int32) *GetSparkAppLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppLogResponse) SetBody(v *GetSparkAppLogResponseBody) *GetSparkAppLogResponse {
	s.Body = v
	return s
}

type GetSparkAppMetricsRequest struct {
	// The ID of the Spark application.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsRequest) SetAppId(v string) *GetSparkAppMetricsRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppMetricsRequest) SetDBClusterId(v string) *GetSparkAppMetricsRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkAppMetricsResponseBody struct {
	// The returned data.
	Data *GetSparkAppMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBody) SetData(v *GetSparkAppMetricsResponseBodyData) *GetSparkAppMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppMetricsResponseBody) SetRequestId(v string) *GetSparkAppMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppMetricsResponseBodyData struct {
	// The ID of the Spark application.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The attempt ID of the Spark application.
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The path of the event log.
	EventLogPath *string `json:"EventLogPath,omitempty" xml:"EventLogPath,omitempty"`
	// Indicates whether parsing is complete. Valid values:
	//
	// *   true
	// *   false
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The metrics.
	ScanMetrics *GetSparkAppMetricsResponseBodyDataScanMetrics `json:"ScanMetrics,omitempty" xml:"ScanMetrics,omitempty" type:"Struct"`
}

func (s GetSparkAppMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBodyData) SetAppId(v string) *GetSparkAppMetricsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetAttemptId(v string) *GetSparkAppMetricsResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetEventLogPath(v string) *GetSparkAppMetricsResponseBodyData {
	s.EventLogPath = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetFinished(v bool) *GetSparkAppMetricsResponseBodyData {
	s.Finished = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyData) SetScanMetrics(v *GetSparkAppMetricsResponseBodyDataScanMetrics) *GetSparkAppMetricsResponseBodyData {
	s.ScanMetrics = v
	return s
}

type GetSparkAppMetricsResponseBodyDataScanMetrics struct {
	// The number of scanned rows.
	OutputRowsCount *int64 `json:"OutputRowsCount,omitempty" xml:"OutputRowsCount,omitempty"`
	// The number of scanned bytes.
	TotalReadFileSizeInByte *int64 `json:"TotalReadFileSizeInByte,omitempty" xml:"TotalReadFileSizeInByte,omitempty"`
}

func (s GetSparkAppMetricsResponseBodyDataScanMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppMetricsResponseBodyDataScanMetrics) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) SetOutputRowsCount(v int64) *GetSparkAppMetricsResponseBodyDataScanMetrics {
	s.OutputRowsCount = &v
	return s
}

func (s *GetSparkAppMetricsResponseBodyDataScanMetrics) SetTotalReadFileSizeInByte(v int64) *GetSparkAppMetricsResponseBodyDataScanMetrics {
	s.TotalReadFileSizeInByte = &v
	return s
}

type GetSparkAppMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponse) SetHeaders(v map[string]*string) *GetSparkAppMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppMetricsResponse) SetStatusCode(v int32) *GetSparkAppMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppMetricsResponse) SetBody(v *GetSparkAppMetricsResponseBody) *GetSparkAppMetricsResponse {
	s.Body = v
	return s
}

type GetSparkAppStateRequest struct {
	// The Spark application ID.
	//
	// >  You can call the [ListSparkApps](~~455888~~) operation to query Spark application IDs.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~612397~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppStateRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateRequest) SetAppId(v string) *GetSparkAppStateRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppStateRequest) SetDBClusterId(v string) *GetSparkAppStateRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkAppStateResponseBody struct {
	// The returned data.
	Data *GetSparkAppStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponseBody) SetData(v *GetSparkAppStateResponseBodyData) *GetSparkAppStateResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppStateResponseBody) SetRequestId(v string) *GetSparkAppStateResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppStateResponseBodyData struct {
	// The Spark application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The database ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The alert message returned for the operation, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the application. Valid values:
	//
	// *   **SUBMITTED**
	// *   **STARTING**
	// *   **RUNNING**
	// *   **FAILING**
	// *   **FAILED**
	// *   **KILLING**
	// *   **KILLED**
	// *   **SUCCEEDING**
	// *   **COMPLETED**
	// *   **FATAL**
	// *   **UNKNOWN**
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSparkAppStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponseBodyData) SetAppId(v string) *GetSparkAppStateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetAppName(v string) *GetSparkAppStateResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetDBClusterId(v string) *GetSparkAppStateResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetMessage(v string) *GetSparkAppStateResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetState(v string) *GetSparkAppStateResponseBodyData {
	s.State = &v
	return s
}

type GetSparkAppStateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppStateResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponse) SetHeaders(v map[string]*string) *GetSparkAppStateResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppStateResponse) SetStatusCode(v int32) *GetSparkAppStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppStateResponse) SetBody(v *GetSparkAppStateResponseBody) *GetSparkAppStateResponse {
	s.Body = v
	return s
}

type GetSparkAppWebUiAddressRequest struct {
	// The Spark application ID.
	//
	// >  You can call the [ListSparkApps](~~455888~~) operation to query Spark application IDs.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkAppWebUiAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppWebUiAddressRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressRequest) SetAppId(v string) *GetSparkAppWebUiAddressRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppWebUiAddressRequest) SetDBClusterId(v string) *GetSparkAppWebUiAddressRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkAppWebUiAddressResponseBody struct {
	// The returned data.
	Data *GetSparkAppWebUiAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppWebUiAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponseBody) SetData(v *GetSparkAppWebUiAddressResponseBodyData) *GetSparkAppWebUiAddressResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBody) SetRequestId(v string) *GetSparkAppWebUiAddressResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkAppWebUiAddressResponseBodyData struct {
	// The Spark application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The database ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The expiration time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	ExpirationTimeInMillis *int64 `json:"ExpirationTimeInMillis,omitempty" xml:"ExpirationTimeInMillis,omitempty"`
	// The URL of the web UI for the Spark application.
	WebUiAddress *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s GetSparkAppWebUiAddressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetAppId(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetDBClusterId(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetExpirationTimeInMillis(v int64) *GetSparkAppWebUiAddressResponseBodyData {
	s.ExpirationTimeInMillis = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetWebUiAddress(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.WebUiAddress = &v
	return s
}

type GetSparkAppWebUiAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppWebUiAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppWebUiAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponse) SetHeaders(v map[string]*string) *GetSparkAppWebUiAddressResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppWebUiAddressResponse) SetStatusCode(v int32) *GetSparkAppWebUiAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponse) SetBody(v *GetSparkAppWebUiAddressResponseBody) *GetSparkAppWebUiAddressResponse {
	s.Body = v
	return s
}

type GetSparkConfigLogPathRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkConfigLogPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkConfigLogPathRequest) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathRequest) SetDBClusterId(v string) *GetSparkConfigLogPathRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkConfigLogPathResponseBody struct {
	// The queried Spark log configuration.
	Data *GetSparkConfigLogPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkConfigLogPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkConfigLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponseBody) SetData(v *GetSparkConfigLogPathResponseBodyData) *GetSparkConfigLogPathResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkConfigLogPathResponseBody) SetRequestId(v string) *GetSparkConfigLogPathResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkConfigLogPathResponseBodyData struct {
	// The default log path.
	DefaultLogPath *string `json:"DefaultLogPath,omitempty" xml:"DefaultLogPath,omitempty"`
	// Indicates whether a log path exists.
	IsLogPathExists *bool `json:"IsLogPathExists,omitempty" xml:"IsLogPathExists,omitempty"`
	// The last modification time.
	ModifiedTimestamp *string `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The account ID of the modifier.
	ModifiedUid *string `json:"ModifiedUid,omitempty" xml:"ModifiedUid,omitempty"`
	// The recorded log path.
	RecordedLogPath *string `json:"RecordedLogPath,omitempty" xml:"RecordedLogPath,omitempty"`
}

func (s GetSparkConfigLogPathResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkConfigLogPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponseBodyData) SetDefaultLogPath(v string) *GetSparkConfigLogPathResponseBodyData {
	s.DefaultLogPath = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetIsLogPathExists(v bool) *GetSparkConfigLogPathResponseBodyData {
	s.IsLogPathExists = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetModifiedTimestamp(v string) *GetSparkConfigLogPathResponseBodyData {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetModifiedUid(v string) *GetSparkConfigLogPathResponseBodyData {
	s.ModifiedUid = &v
	return s
}

func (s *GetSparkConfigLogPathResponseBodyData) SetRecordedLogPath(v string) *GetSparkConfigLogPathResponseBodyData {
	s.RecordedLogPath = &v
	return s
}

type GetSparkConfigLogPathResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkConfigLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkConfigLogPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkConfigLogPathResponse) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponse) SetHeaders(v map[string]*string) *GetSparkConfigLogPathResponse {
	s.Headers = v
	return s
}

func (s *GetSparkConfigLogPathResponse) SetStatusCode(v int32) *GetSparkConfigLogPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkConfigLogPathResponse) SetBody(v *GetSparkConfigLogPathResponseBody) *GetSparkConfigLogPathResponse {
	s.Body = v
	return s
}

type GetSparkDefinitionsRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkDefinitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *GetSparkDefinitionsRequest) SetDBClusterId(v string) *GetSparkDefinitionsRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkDefinitionsResponseBody struct {
	// The common definitions of Spark applications.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkDefinitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkDefinitionsResponseBody) SetData(v string) *GetSparkDefinitionsResponseBody {
	s.Data = &v
	return s
}

func (s *GetSparkDefinitionsResponseBody) SetRequestId(v string) *GetSparkDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkDefinitionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkDefinitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *GetSparkDefinitionsResponse) SetHeaders(v map[string]*string) *GetSparkDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *GetSparkDefinitionsResponse) SetStatusCode(v int32) *GetSparkDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkDefinitionsResponse) SetBody(v *GetSparkDefinitionsResponseBody) *GetSparkDefinitionsResponse {
	s.Body = v
	return s
}

type GetSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark log analysis task. You can call the ListSparkLogAnalyzeTasks operation to query the IDs of all Spark log analysis tasks that are submitted in the current cluster.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSparkLogAnalyzeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskRequest) SetTaskId(v int64) *GetSparkLogAnalyzeTaskRequest {
	s.TaskId = &v
	return s
}

type GetSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkLogAnalyzeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *GetSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *GetSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkLogAnalyzeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *GetSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *GetSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponse) SetBody(v *GetSparkLogAnalyzeTaskResponseBody) *GetSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

type GetSparkSQLEngineStateRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the job resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s GetSparkSQLEngineStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSQLEngineStateRequest) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateRequest) SetDBClusterId(v string) *GetSparkSQLEngineStateRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkSQLEngineStateRequest) SetResourceGroupName(v string) *GetSparkSQLEngineStateRequest {
	s.ResourceGroupName = &v
	return s
}

type GetSparkSQLEngineStateResponseBody struct {
	// The state information about the Spark SQL engine.
	Data *GetSparkSQLEngineStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkSQLEngineStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSQLEngineStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponseBody) SetData(v *GetSparkSQLEngineStateResponseBodyData) *GetSparkSQLEngineStateResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkSQLEngineStateResponseBody) SetRequestId(v string) *GetSparkSQLEngineStateResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkSQLEngineStateResponseBodyData struct {
	// The ID of the Spark application.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configuration of the Spark application.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The third-party JAR package.
	Jars *string `json:"Jars,omitempty" xml:"Jars,omitempty"`
	// The maximum number of started Spark executors.
	MaxExecutor *string `json:"MaxExecutor,omitempty" xml:"MaxExecutor,omitempty"`
	// The minimum number of started Spark executors.
	MinExecutor *string `json:"MinExecutor,omitempty" xml:"MinExecutor,omitempty"`
	// The slot number of the Spark application.
	SlotNum *string `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
	// The execution state of the application. Valid values:
	//
	// *   SUBMITTED
	// *   STARTING
	// *   RUNNING
	// *   FAILING
	// *   FAILED
	// *   KILLING
	// *   KILLED
	// *   SUCCEEDING
	// *   COMPLETED
	// *   FATAL
	// *   UNKNOWN
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timestamp when the Spark SQL application was submitted. Unit: milliseconds.
	SubmittedTimeInMillis *string `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
}

func (s GetSparkSQLEngineStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSQLEngineStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetAppId(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetConfig(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetJars(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.Jars = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetMaxExecutor(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.MaxExecutor = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetMinExecutor(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.MinExecutor = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetSlotNum(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.SlotNum = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetState(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetSubmittedTimeInMillis(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.SubmittedTimeInMillis = &v
	return s
}

type GetSparkSQLEngineStateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkSQLEngineStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkSQLEngineStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSQLEngineStateResponse) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponse) SetHeaders(v map[string]*string) *GetSparkSQLEngineStateResponse {
	s.Headers = v
	return s
}

func (s *GetSparkSQLEngineStateResponse) SetStatusCode(v int32) *GetSparkSQLEngineStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkSQLEngineStateResponse) SetBody(v *GetSparkSQLEngineStateResponseBody) *GetSparkSQLEngineStateResponse {
	s.Body = v
	return s
}

type GetSparkTemplateFileContentRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application template ID.
	//
	// >  You can call the [GetSparkTemplateFullTree](~~456205~~) operation to query the application template ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSparkTemplateFileContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFileContentRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentRequest) SetDBClusterId(v string) *GetSparkTemplateFileContentRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkTemplateFileContentRequest) SetId(v int64) *GetSparkTemplateFileContentRequest {
	s.Id = &v
	return s
}

type GetSparkTemplateFileContentResponseBody struct {
	// The returned data.
	Data *GetSparkTemplateFileContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFileContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFileContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponseBody) SetData(v *GetSparkTemplateFileContentResponseBodyData) *GetSparkTemplateFileContentResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkTemplateFileContentResponseBody) SetRequestId(v string) *GetSparkTemplateFileContentResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkTemplateFileContentResponseBodyData struct {
	// The application type. Valid values:
	//
	// *   **SQL**
	// *   **STREAMING**
	// *   **BATCH**
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The content of the application template.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The application template ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The file type. Valid values:
	//
	// *   **folder**
	// *   **file**
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSparkTemplateFileContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFileContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetAppType(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetContent(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetId(v int64) *GetSparkTemplateFileContentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetResourceGroupName(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetType(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.Type = &v
	return s
}

type GetSparkTemplateFileContentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFileContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFileContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFileContentResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFileContentResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFileContentResponse) SetStatusCode(v int32) *GetSparkTemplateFileContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFileContentResponse) SetBody(v *GetSparkTemplateFileContentResponseBody) *GetSparkTemplateFileContentResponse {
	s.Body = v
	return s
}

type GetSparkTemplateFolderTreeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkTemplateFolderTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFolderTreeRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeRequest) SetDBClusterId(v string) *GetSparkTemplateFolderTreeRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkTemplateFolderTreeResponseBody struct {
	// The directory structure of Spark applications, which is in the tree format. Fields in the response parameter:
	//
	// *   **Uid**: the UID of the Alibaba Cloud account.
	//
	// *   **Type**: the application template type. Valid values: **FOLDER**
	//
	// *   **Parent**: indicates whether a child directory exists. Valid values:
	//
	//     *   **0**: no.
	//     *   **-1**: yes.
	//
	// *   **Children**: the child directory.
	//
	// *   **LastModified**: the time when applications in the directory are last modified. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **Name**: the name of the directory.
	//
	// *   **Id**: the directory ID.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFolderTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFolderTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeResponseBody) SetData(v string) *GetSparkTemplateFolderTreeResponseBody {
	s.Data = &v
	return s
}

func (s *GetSparkTemplateFolderTreeResponseBody) SetRequestId(v string) *GetSparkTemplateFolderTreeResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkTemplateFolderTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFolderTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFolderTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFolderTreeResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFolderTreeResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFolderTreeResponse) SetStatusCode(v int32) *GetSparkTemplateFolderTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFolderTreeResponse) SetBody(v *GetSparkTemplateFolderTreeResponseBody) *GetSparkTemplateFolderTreeResponse {
	s.Body = v
	return s
}

type GetSparkTemplateFullTreeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkTemplateFullTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFullTreeRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeRequest) SetDBClusterId(v string) *GetSparkTemplateFullTreeRequest {
	s.DBClusterId = &v
	return s
}

type GetSparkTemplateFullTreeResponseBody struct {
	// The directory structure of Spark applications. Fields in the response parameter:
	//
	// *   **Uid**: the UID of the Alibaba Cloud account.
	//
	// *   **Type**: the application template type. Valid values:
	//
	//     *   **FOLDER**
	//     *   **FILE**
	//
	// *   **Parent**: indicates whether a child directory exists. Valid values:
	//
	//     *   **0**: no.
	//     *   **-1**: yes.
	//
	// *   **Children**: the child directory.
	//
	// *   **LastModified**: the time when applications are last modified. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **AppType**: the application type. Valid values:
	//
	//     *   **SQL**
	//     *   **STREAMING**
	//     *   **BATCH**
	//
	// *   **Name**: the name of the directory or application.
	//
	// *   **Id**: the directory ID or application ID.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFullTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFullTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeResponseBody) SetData(v string) *GetSparkTemplateFullTreeResponseBody {
	s.Data = &v
	return s
}

func (s *GetSparkTemplateFullTreeResponseBody) SetRequestId(v string) *GetSparkTemplateFullTreeResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkTemplateFullTreeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFullTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFullTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkTemplateFullTreeResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFullTreeResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFullTreeResponse) SetStatusCode(v int32) *GetSparkTemplateFullTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFullTreeResponse) SetBody(v *GetSparkTemplateFullTreeResponseBody) *GetSparkTemplateFullTreeResponse {
	s.Body = v
	return s
}

type GetTableRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the region in which the cluster resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) SetDBClusterId(v string) *GetTableRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableRequest) SetDbName(v string) *GetTableRequest {
	s.DbName = &v
	return s
}

func (s *GetTableRequest) SetRegionId(v string) *GetTableRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableRequest) SetTableName(v string) *GetTableRequest {
	s.TableName = &v
	return s
}

type GetTableResponseBody struct {
	// The error code returned.
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the query succeeded.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Details of the table.
	Table *TableModel `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) SetCode(v int64) *GetTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableResponseBody) SetMessage(v string) *GetTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) SetTable(v *TableModel) *GetTableResponseBody {
	s.Table = v
	return s
}

type GetTableResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponse) GoString() string {
	return s.String()
}

func (s *GetTableResponse) SetHeaders(v map[string]*string) *GetTableResponse {
	s.Headers = v
	return s
}

func (s *GetTableResponse) SetStatusCode(v int32) *GetTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableResponse) SetBody(v *GetTableResponseBody) *GetTableResponse {
	s.Body = v
	return s
}

type GetTableColumnsRequest struct {
	// The name of the column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnsRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnsRequest) SetColumnName(v string) *GetTableColumnsRequest {
	s.ColumnName = &v
	return s
}

func (s *GetTableColumnsRequest) SetDBClusterId(v string) *GetTableColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableColumnsRequest) SetPageNumber(v int64) *GetTableColumnsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsRequest) SetPageSize(v int64) *GetTableColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsRequest) SetRegionId(v string) *GetTableColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableColumnsRequest) SetSchemaName(v string) *GetTableColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableColumnsRequest) SetTableName(v string) *GetTableColumnsRequest {
	s.TableName = &v
	return s
}

type GetTableColumnsResponseBody struct {
	// The queried data.
	Data *GetTableColumnsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBody) SetData(v *GetTableColumnsResponseBodyData) *GetTableColumnsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableColumnsResponseBody) SetPageNumber(v int64) *GetTableColumnsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetPageSize(v int64) *GetTableColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetRequestId(v string) *GetTableColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetTotalCount(v int64) *GetTableColumnsResponseBody {
	s.TotalCount = &v
	return s
}

type GetTableColumnsResponseBodyData struct {
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the table.
	Table *TableDetailModel `json:"Table,omitempty" xml:"Table,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableColumnsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBodyData) SetPageNumber(v int64) *GetTableColumnsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetPageSize(v int64) *GetTableColumnsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetTable(v *TableDetailModel) *GetTableColumnsResponseBodyData {
	s.Table = v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetTotalCount(v int64) *GetTableColumnsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetTableColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnsResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponse) SetHeaders(v map[string]*string) *GetTableColumnsResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnsResponse) SetStatusCode(v int32) *GetTableColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnsResponse) SetBody(v *GetTableColumnsResponseBody) *GetTableColumnsResponse {
	s.Body = v
	return s
}

type GetTableDDLRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableDDLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableDDLRequest) GoString() string {
	return s.String()
}

func (s *GetTableDDLRequest) SetDBClusterId(v string) *GetTableDDLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableDDLRequest) SetRegionId(v string) *GetTableDDLRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableDDLRequest) SetSchemaName(v string) *GetTableDDLRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableDDLRequest) SetTableName(v string) *GetTableDDLRequest {
	s.TableName = &v
	return s
}

type GetTableDDLResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s GetTableDDLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableDDLResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDDLResponseBody) SetRequestId(v string) *GetTableDDLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDDLResponseBody) SetSQL(v string) *GetTableDDLResponseBody {
	s.SQL = &v
	return s
}

type GetTableDDLResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableDDLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableDDLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableDDLResponse) GoString() string {
	return s.String()
}

func (s *GetTableDDLResponse) SetHeaders(v map[string]*string) *GetTableDDLResponse {
	s.Headers = v
	return s
}

func (s *GetTableDDLResponse) SetStatusCode(v int32) *GetTableDDLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDDLResponse) SetBody(v *GetTableDDLResponseBody) *GetTableDDLResponse {
	s.Body = v
	return s
}

type GetTableObjectsRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the table.
	FilterDescription *string `json:"FilterDescription,omitempty" xml:"FilterDescription,omitempty"`
	// The owner of the table.
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the table.
	FilterTblName *string `json:"FilterTblName,omitempty" xml:"FilterTblName,omitempty"`
	// The type of the table.
	//
	// Valid values:
	//
	// DIMENSION_TABLE
	//
	// FACT_TABLE
	//
	// EXTERNAL_TABLE
	//
	// Default value: null.
	FilterTblType *string `json:"FilterTblType,omitempty" xml:"FilterTblType,omitempty"`
	// The order in which the fields to be returned are sorted.
	//
	// Valid values:
	//
	// *   Asc
	// *   Desc
	//
	// Values for fields:
	//
	// TableName
	//
	// TableSize
	//
	// CreateTime
	//
	// UpdateTime
	//
	// Default value: {"Type": "Desc","Field": "TableName"};
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The value is an integer that is greater than 0. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the cluster resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s GetTableObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetTableObjectsRequest) SetDBClusterId(v string) *GetTableObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterDescription(v string) *GetTableObjectsRequest {
	s.FilterDescription = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterOwner(v string) *GetTableObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterTblName(v string) *GetTableObjectsRequest {
	s.FilterTblName = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterTblType(v string) *GetTableObjectsRequest {
	s.FilterTblType = &v
	return s
}

func (s *GetTableObjectsRequest) SetOrderBy(v string) *GetTableObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTableObjectsRequest) SetPageNumber(v int64) *GetTableObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsRequest) SetPageSize(v int64) *GetTableObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsRequest) SetRegionId(v string) *GetTableObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableObjectsRequest) SetSchemaName(v string) *GetTableObjectsRequest {
	s.SchemaName = &v
	return s
}

type GetTableObjectsResponseBody struct {
	// The data returned.
	Data *GetTableObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of the returned page. The value is an integer that is greater than 0. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponseBody) SetData(v *GetTableObjectsResponseBodyData) *GetTableObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableObjectsResponseBody) SetPageNumber(v int64) *GetTableObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetPageSize(v int64) *GetTableObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetRequestId(v string) *GetTableObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetTotalCount(v int64) *GetTableObjectsResponseBody {
	s.TotalCount = &v
	return s
}

type GetTableObjectsResponseBodyData struct {
	// The number of the returned page. The value is an integer that is greater than 0. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details of the tables.
	TableSummaryModels []*TableSummaryModel `json:"TableSummaryModels,omitempty" xml:"TableSummaryModels,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTableObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponseBodyData) SetPageNumber(v int64) *GetTableObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetPageSize(v int64) *GetTableObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetTableSummaryModels(v []*TableSummaryModel) *GetTableObjectsResponseBodyData {
	s.TableSummaryModels = v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetTotalCount(v int64) *GetTableObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetTableObjectsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponse) SetHeaders(v map[string]*string) *GetTableObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetTableObjectsResponse) SetStatusCode(v int32) *GetTableObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableObjectsResponse) SetBody(v *GetTableObjectsResponseBody) *GetTableObjectsResponse {
	s.Body = v
	return s
}

type GetViewDDLRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the view.
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s GetViewDDLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetViewDDLRequest) GoString() string {
	return s.String()
}

func (s *GetViewDDLRequest) SetDBClusterId(v string) *GetViewDDLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetViewDDLRequest) SetRegionId(v string) *GetViewDDLRequest {
	s.RegionId = &v
	return s
}

func (s *GetViewDDLRequest) SetSchemaName(v string) *GetViewDDLRequest {
	s.SchemaName = &v
	return s
}

func (s *GetViewDDLRequest) SetViewName(v string) *GetViewDDLRequest {
	s.ViewName = &v
	return s
}

type GetViewDDLResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s GetViewDDLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetViewDDLResponseBody) GoString() string {
	return s.String()
}

func (s *GetViewDDLResponseBody) SetRequestId(v string) *GetViewDDLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViewDDLResponseBody) SetSQL(v string) *GetViewDDLResponseBody {
	s.SQL = &v
	return s
}

type GetViewDDLResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViewDDLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViewDDLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetViewDDLResponse) GoString() string {
	return s.String()
}

func (s *GetViewDDLResponse) SetHeaders(v map[string]*string) *GetViewDDLResponse {
	s.Headers = v
	return s
}

func (s *GetViewDDLResponse) SetStatusCode(v int32) *GetViewDDLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViewDDLResponse) SetBody(v *GetViewDDLResponseBody) *GetViewDDLResponse {
	s.Body = v
	return s
}

type GetViewObjectsRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The owner of the view.
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the view.
	FilterViewName *string `json:"FilterViewName,omitempty" xml:"FilterViewName,omitempty"`
	// The type of the view.
	//
	// Valid values:
	//
	// \-VIRTUAL_VIEW
	//
	// \-MATERIALIZED_VIEW
	//
	// Default value: null.
	FilterViewType *string `json:"FilterViewType,omitempty" xml:"FilterViewType,omitempty"`
	// The order in which you want to sort the query results. Valid values for Type:
	//
	// *   Asc
	// *   Desc
	//
	// Valid values for Field: -ViewName
	//
	// \-CreateTime
	//
	// \-UpdateTime
	//
	// Default value: {"Type": "Desc","Field": "ViewName"}.
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s GetViewObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetViewObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetViewObjectsRequest) SetDBClusterId(v string) *GetViewObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterOwner(v string) *GetViewObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterViewName(v string) *GetViewObjectsRequest {
	s.FilterViewName = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterViewType(v string) *GetViewObjectsRequest {
	s.FilterViewType = &v
	return s
}

func (s *GetViewObjectsRequest) SetOrderBy(v string) *GetViewObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetViewObjectsRequest) SetPageNumber(v int64) *GetViewObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsRequest) SetPageSize(v int64) *GetViewObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsRequest) SetRegionId(v string) *GetViewObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *GetViewObjectsRequest) SetSchemaName(v string) *GetViewObjectsRequest {
	s.SchemaName = &v
	return s
}

type GetViewObjectsResponseBody struct {
	// The returned data.
	Data *GetViewObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetViewObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetViewObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponseBody) SetData(v *GetViewObjectsResponseBodyData) *GetViewObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetViewObjectsResponseBody) SetPageNumber(v int64) *GetViewObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetPageSize(v int64) *GetViewObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetRequestId(v string) *GetViewObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetTotalCount(v int64) *GetViewObjectsResponseBody {
	s.TotalCount = &v
	return s
}

type GetViewObjectsResponseBodyData struct {
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried views.
	TableSummaryModels []*TableSummaryModel `json:"TableSummaryModels,omitempty" xml:"TableSummaryModels,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetViewObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetViewObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponseBodyData) SetPageNumber(v int64) *GetViewObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetPageSize(v int64) *GetViewObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetTableSummaryModels(v []*TableSummaryModel) *GetViewObjectsResponseBodyData {
	s.TableSummaryModels = v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetTotalCount(v int64) *GetViewObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetViewObjectsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViewObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViewObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetViewObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponse) SetHeaders(v map[string]*string) *GetViewObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetViewObjectsResponse) SetStatusCode(v int32) *GetViewObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViewObjectsResponse) SetBody(v *GetViewObjectsResponseBody) *GetViewObjectsResponse {
	s.Body = v
	return s
}

type KillSparkAppRequest struct {
	// The ID of the Spark application that you want to terminate.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s KillSparkAppRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkAppRequest) GoString() string {
	return s.String()
}

func (s *KillSparkAppRequest) SetAppId(v string) *KillSparkAppRequest {
	s.AppId = &v
	return s
}

func (s *KillSparkAppRequest) SetDBClusterId(v string) *KillSparkAppRequest {
	s.DBClusterId = &v
	return s
}

type KillSparkAppResponseBody struct {
	// The returned data.
	Data *KillSparkAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillSparkAppResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponseBody) SetData(v *KillSparkAppResponseBodyData) *KillSparkAppResponseBody {
	s.Data = v
	return s
}

func (s *KillSparkAppResponseBody) SetRequestId(v string) *KillSparkAppResponseBody {
	s.RequestId = &v
	return s
}

type KillSparkAppResponseBodyData struct {
	// The Spark application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the Spark application. Valid values:
	//
	// *   **SUBMITTED**
	// *   **STARTING**
	// *   **RUNNING**
	// *   **FAILING**
	// *   **FAILED**
	// *   **KILLING**
	// *   **KILLED**
	// *   **SUCCEEDING**
	// *   **COMPLETED**
	// *   **FATAL**
	// *   **UNKNOWN**
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s KillSparkAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s KillSparkAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponseBodyData) SetAppId(v string) *KillSparkAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetAppName(v string) *KillSparkAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetDBClusterId(v string) *KillSparkAppResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetMessage(v string) *KillSparkAppResponseBodyData {
	s.Message = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetState(v string) *KillSparkAppResponseBodyData {
	s.State = &v
	return s
}

type KillSparkAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkAppResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkAppResponse) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponse) SetHeaders(v map[string]*string) *KillSparkAppResponse {
	s.Headers = v
	return s
}

func (s *KillSparkAppResponse) SetStatusCode(v int32) *KillSparkAppResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkAppResponse) SetBody(v *KillSparkAppResponseBody) *KillSparkAppResponse {
	s.Body = v
	return s
}

type KillSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark log analysis task. You can call the ListSparkLogAnalyzeTasks operation to query the IDs and states of all analysis tasks in the current cluster.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s KillSparkLogAnalyzeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskRequest) SetTaskId(v int64) *KillSparkLogAnalyzeTaskRequest {
	s.TaskId = &v
	return s
}

type KillSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkLogAnalyzeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *KillSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *KillSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

type KillSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkLogAnalyzeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *KillSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *KillSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponse) SetBody(v *KillSparkLogAnalyzeTaskResponseBody) *KillSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

type KillSparkSQLEngineRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s KillSparkSQLEngineRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkSQLEngineRequest) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineRequest) SetDBClusterId(v string) *KillSparkSQLEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillSparkSQLEngineRequest) SetResourceGroupName(v string) *KillSparkSQLEngineRequest {
	s.ResourceGroupName = &v
	return s
}

type KillSparkSQLEngineResponseBody struct {
	// Indicates whether the request was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkSQLEngineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillSparkSQLEngineResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineResponseBody) SetData(v bool) *KillSparkSQLEngineResponseBody {
	s.Data = &v
	return s
}

func (s *KillSparkSQLEngineResponseBody) SetRequestId(v string) *KillSparkSQLEngineResponseBody {
	s.RequestId = &v
	return s
}

type KillSparkSQLEngineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkSQLEngineResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkSQLEngineResponse) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineResponse) SetHeaders(v map[string]*string) *KillSparkSQLEngineResponse {
	s.Headers = v
	return s
}

func (s *KillSparkSQLEngineResponse) SetStatusCode(v int32) *KillSparkSQLEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkSQLEngineResponse) SetBody(v *KillSparkSQLEngineResponseBody) *KillSparkSQLEngineResponse {
	s.Body = v
	return s
}

type ListSparkAppAttemptsRequest struct {
	// The ID of the Spark application.
	//
	// > You can call the [ListSparkApps](~~455888~~) operation to query all application IDs.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **10** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSparkAppAttemptsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppAttemptsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsRequest) SetAppId(v string) *ListSparkAppAttemptsRequest {
	s.AppId = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetDBClusterId(v string) *ListSparkAppAttemptsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetPageNumber(v int64) *ListSparkAppAttemptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetPageSize(v int64) *ListSparkAppAttemptsRequest {
	s.PageSize = &v
	return s
}

type ListSparkAppAttemptsResponseBody struct {
	// The returned data.
	Data *ListSparkAppAttemptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkAppAttemptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppAttemptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponseBody) SetData(v *ListSparkAppAttemptsResponseBodyData) *ListSparkAppAttemptsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkAppAttemptsResponseBody) SetRequestId(v string) *ListSparkAppAttemptsResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkAppAttemptsResponseBodyData struct {
	// The information about the attempts. Fields in the response parameter:
	//
	// *   **AttemptId**: the attempt ID.
	//
	// *   **State**: the state of the Spark application. Valid values:
	//
	//     *   **SUBMITTED**
	//     *   **STARTING**
	//     *   **RUNNING**
	//     *   **FAILING**
	//     *   **FAILED**
	//     *   **KILLING**
	//     *   **KILLED**
	//     *   **SUCCEEDING**
	//     *   **COMPLETED**
	//     *   **FATAL**
	//     *   **UNKNOWN**
	//
	// *   **Message**: the alert message that is returned. If no alert is generated, null is returned.
	//
	// *   **Data**: the data of the Spark application template.
	//
	// *   **EstimateExecutionCpuTimeInSeconds**: the amount of time it takes to consume CPU resources for running the Spark application. Unit: milliseconds.
	//
	// *   **LogRootPath**: the storage path of log files.
	//
	// *   **LastAttemptId**: the ID of the last attempt.
	//
	// *   **WebUiAddress**: the web UI address.
	//
	// *   **SubmittedTimeInMillis**: the time when the Spark application was submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **StartedTimeInMillis**: the time when the Spark application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **TerminatedTimeInMillis**: the time when the Spark application task was terminated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// *   **DBClusterId**: the ID of the cluster on which the Spark application runs.
	//
	// *   **ResourceGroupName**: the name of the job resource group.
	//
	// *   **DurationInMillis**: the amount of time it takes to run the Spark application. Unit: milliseconds.
	AttemptInfoList []*SparkAttemptInfo `json:"AttemptInfoList,omitempty" xml:"AttemptInfoList,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppAttemptsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppAttemptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponseBodyData) SetAttemptInfoList(v []*SparkAttemptInfo) *ListSparkAppAttemptsResponseBodyData {
	s.AttemptInfoList = v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetPageNumber(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetPageSize(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppAttemptsResponseBodyData) SetTotalCount(v int64) *ListSparkAppAttemptsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSparkAppAttemptsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkAppAttemptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkAppAttemptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppAttemptsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponse) SetHeaders(v map[string]*string) *ListSparkAppAttemptsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkAppAttemptsResponse) SetStatusCode(v int32) *ListSparkAppAttemptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkAppAttemptsResponse) SetBody(v *ListSparkAppAttemptsResponseBody) *ListSparkAppAttemptsResponse {
	s.Body = v
	return s
}

type ListSparkAppsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Valid values:
	//
	// - **10**
	// - **50**
	// - **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the job resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ListSparkAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkAppsRequest) SetDBClusterId(v string) *ListSparkAppsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkAppsRequest) SetPageNumber(v int64) *ListSparkAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsRequest) SetPageSize(v int64) *ListSparkAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsRequest) SetResourceGroupName(v string) *ListSparkAppsRequest {
	s.ResourceGroupName = &v
	return s
}

type ListSparkAppsResponseBody struct {
	// The data returned.
	Data *ListSparkAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponseBody) SetData(v *ListSparkAppsResponseBodyData) *ListSparkAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkAppsResponseBody) SetPageNumber(v int64) *ListSparkAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetPageSize(v int64) *ListSparkAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetRequestId(v string) *ListSparkAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkAppsResponseBody) SetTotalCount(v int64) *ListSparkAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSparkAppsResponseBodyData struct {
	// Details of the applications. Fields in the response parameter:
	//
	// - **Data**: the data of the Spark application template.
	// - **EstimateExecutionCpuTimeInSeconds**: the amount of time it takes to consume CPU resources for running the Spark application. Unit: milliseconds.
	// - **LogRootPath**: the storage path of log files.
	// - **LastAttemptId**: the most recent attempt ID.
	// - **WebUiAddress**: the web UI URL.
	// - **SubmittedTimeInMillis**: the time when the Spark application was submitted. The time is displayed in the UNIX timestamp format. Unit: milliseconds.
	// - **StartedTimeInMillis**: the time when the Spark application was created. The time is displayed in the UNIX timestamp format. Unit: milliseconds.
	// - **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. The time is displayed in the UNIX timestamp format. Unit: milliseconds.
	// - **TerminatedTimeInMillis**: the time when the Spark application task was terminated. The time is displayed in the UNIX timestamp format. Unit: milliseconds.
	// - **DBClusterId**: the ID of the cluster on which the Spark application runs.
	// - **ResourceGroupName**: the name of the job resource group.
	// - **DurationInMillis**: the amount of time it takes to run the Spark application. Unit: milliseconds.
	AppInfoList []*SparkAppInfo `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponseBodyData) SetAppInfoList(v []*SparkAppInfo) *ListSparkAppsResponseBodyData {
	s.AppInfoList = v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetPageNumber(v int64) *ListSparkAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetPageSize(v int64) *ListSparkAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsResponseBodyData) SetTotalCount(v int64) *ListSparkAppsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSparkAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkAppsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponse) SetHeaders(v map[string]*string) *ListSparkAppsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkAppsResponse) SetStatusCode(v int32) *ListSparkAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkAppsResponse) SetBody(v *ListSparkAppsResponseBody) *ListSparkAppsResponse {
	s.Body = v
	return s
}

type ListSparkLogAnalyzeTasksRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSparkLogAnalyzeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksRequest) SetDBClusterId(v string) *ListSparkLogAnalyzeTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksRequest) SetPageNumber(v int64) *ListSparkLogAnalyzeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksRequest) SetPageSize(v int64) *ListSparkLogAnalyzeTasksRequest {
	s.PageSize = &v
	return s
}

type ListSparkLogAnalyzeTasksResponseBody struct {
	// The returned data.
	Data *ListSparkLogAnalyzeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponseBody) SetData(v *ListSparkLogAnalyzeTasksResponseBodyData) *ListSparkLogAnalyzeTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBody) SetRequestId(v string) *ListSparkLogAnalyzeTasksResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkLogAnalyzeTasksResponseBodyData struct {
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried Spark log analysis tasks.
	TaskList []*SparkAnalyzeLogTask `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetPageNumber(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetPageSize(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetTaskList(v []*SparkAnalyzeLogTask) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.TaskList = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponseBodyData) SetTotalCount(v int64) *ListSparkLogAnalyzeTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSparkLogAnalyzeTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkLogAnalyzeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponse) SetHeaders(v map[string]*string) *ListSparkLogAnalyzeTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponse) SetStatusCode(v int32) *ListSparkLogAnalyzeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponse) SetBody(v *ListSparkLogAnalyzeTasksResponseBody) *ListSparkLogAnalyzeTasksResponse {
	s.Body = v
	return s
}

type ListSparkTemplateFileIdsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ListSparkTemplateFileIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkTemplateFileIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsRequest) SetDBClusterId(v string) *ListSparkTemplateFileIdsRequest {
	s.DBClusterId = &v
	return s
}

type ListSparkTemplateFileIdsResponseBody struct {
	// The IDs of Spark template files.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkTemplateFileIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkTemplateFileIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsResponseBody) SetData(v []*int64) *ListSparkTemplateFileIdsResponseBody {
	s.Data = v
	return s
}

func (s *ListSparkTemplateFileIdsResponseBody) SetRequestId(v string) *ListSparkTemplateFileIdsResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkTemplateFileIdsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkTemplateFileIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkTemplateFileIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkTemplateFileIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsResponse) SetHeaders(v map[string]*string) *ListSparkTemplateFileIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkTemplateFileIdsResponse) SetStatusCode(v int32) *ListSparkTemplateFileIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkTemplateFileIdsResponse) SetBody(v *ListSparkTemplateFileIdsResponseBody) *ListSparkTemplateFileIdsResponse {
	s.Body = v
	return s
}

type LoadSampleDataSetRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s LoadSampleDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadSampleDataSetRequest) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetRequest) SetDBClusterId(v string) *LoadSampleDataSetRequest {
	s.DBClusterId = &v
	return s
}

type LoadSampleDataSetResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadSampleDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadSampleDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetResponseBody) SetDBClusterId(v string) *LoadSampleDataSetResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *LoadSampleDataSetResponseBody) SetRequestId(v string) *LoadSampleDataSetResponseBody {
	s.RequestId = &v
	return s
}

type LoadSampleDataSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoadSampleDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoadSampleDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadSampleDataSetResponse) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetResponse) SetHeaders(v map[string]*string) *LoadSampleDataSetResponse {
	s.Headers = v
	return s
}

func (s *LoadSampleDataSetResponse) SetStatusCode(v int32) *LoadSampleDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadSampleDataSetResponse) SetBody(v *LoadSampleDataSetResponseBody) *LoadSampleDataSetResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// The description of the database account.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// > You can call the [DescribeAccounts](~~612430~~) operation to query the information about database accounts in a cluster, including the database account name.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAccountPrivilegesRequest struct {
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions of the database account.
	AccountPrivileges []*ModifyAccountPrivilegesRequestAccountPrivileges `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty" type:"Repeated"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequest) SetAccountName(v string) *ModifyAccountPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetAccountPrivileges(v []*ModifyAccountPrivilegesRequestAccountPrivileges) *ModifyAccountPrivilegesRequest {
	s.AccountPrivileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetRegionId(v string) *ModifyAccountPrivilegesRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountPrivilegesRequestAccountPrivileges struct {
	// The objects on which the permission takes effect, including databases, tables, and columns.
	PrivilegeObject *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level of the database account. You can call the `DescribeEnabledPrivileges` operation to query the permission level of the database account.
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The permissions that you want to modify. You can call the `DescribeEnabledPrivileges` operation to query the permissions of the database account.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeObject(v *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeObject = v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeType(v string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeType = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivileges(v []*string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.Privileges = v
	return s
}

type ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject struct {
	// The columns on which the database account has permissions. This parameter is required if the PrivilegeType parameter is set to Column.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The databases on which the database account has permissions. This parameter is required if the PrivilegeType parameter is set to Database, Table, or Column.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The tables on which the database account has permissions. This parameter is required if the PrivilegeType parameter is set to Table or Column.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetColumn(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Column = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetDatabase(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Database = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetTable(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Table = &v
	return s
}

type ModifyAccountPrivilegesShrinkRequest struct {
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions of the database account.
	AccountPrivilegesShrink *string `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountName(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountPrivilegesShrink = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetRegionId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.RegionId = &v
	return s
}

type ModifyAccountPrivilegesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPrivilegesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesResponseBody) SetRequestId(v string) *ModifyAccountPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountPrivilegesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPrivilegesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegesResponse) SetStatusCode(v int32) *ModifyAccountPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegesResponse) SetBody(v *ModifyAccountPrivilegesResponseBody) *ModifyAccountPrivilegesResponse {
	s.Body = v
	return s
}

type ModifyAuditLogConfigRequest struct {
	// Modifies the status of SQL audit. Valid values:
	//
	// *   **on**: enables SQL audit.
	// *   **off**: disables SQL audit.
	//
	// >  After you disable the SQL audit feature, all SQL audit logs are deleted. You must query and export SQL audit logs before you disable SQL audit. For more information, see Query and export SQL audit logs. When you re-enable SQL audit, audit logs that are generated from the last time when SQL audit was enabled are available for queries.
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogStatus = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetDBClusterId(v string) *ModifyAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRegionId(v string) *ModifyAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyAuditLogConfigResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the status of SQL audit is updated. Valid values:
	//
	// *   **true**
	// *   **false**
	UpdateSucceed *bool `json:"UpdateSucceed,omitempty" xml:"UpdateSucceed,omitempty"`
}

func (s ModifyAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponseBody) SetRequestId(v string) *ModifyAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditLogConfigResponseBody) SetUpdateSucceed(v bool) *ModifyAuditLogConfigResponseBody {
	s.UpdateSucceed = &v
	return s
}

type ModifyAuditLogConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAuditLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponse) SetHeaders(v map[string]*string) *ModifyAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetStatusCode(v int32) *ModifyAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetBody(v *ModifyAuditLogConfigResponseBody) *ModifyAuditLogConfigResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	// The number of days for which to retain full backup files. Valid values: 7 to 730.
	//
	// >  If you do not specify this parameter, the default value 7 is used.
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// *   **Enable**
	// *   **Disable**
	//
	// >  If you do not specify this parameter, the default value Enable is used.
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which to retain log backup files. Valid values: 7 to 730.
	//
	// >  If you do not specify this parameter, the default value 7 is used.
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The days of the week on which to perform a full backup. Separate multiple values with commas (,). Valid values:
	//
	// *   **Monday**
	// *   **Tuesday**
	// *   **Wednesday**
	// *   **Thursday**
	// *   **Friday**
	// *   **Saturday**
	// *   **Sunday**
	//
	// >  To ensure data security, we recommend that you specify at least two values.
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time to perform a full backup. Specify the time in the HH:mmZ-HH:mmZ format. The time must be in UTC.
	//
	// >  The time range must be 1 hour.
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyClusterAccessWhiteListRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty.
	//
	// > Whitelists with the hidden attribute are not displayed in the console. Those whitelists are used to access Data Transmission Service (DTS) and PolarDB.
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist. If you do not specify this parameter, the Default whitelist is modified.
	//
	// *   The whitelist name must be 2 to 32 characters in length. The name can contain lowercase letters, digits, and underscores (\_). The name must start with a lowercase letter and end with a lowercase letter or a digit.
	// *   Each cluster supports up to 50 IP address whitelists.
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The method used to modify the IP address whitelist. Valid values:
	//
	// *   **Cover** (default)
	// *   **Append**
	// *   **Delete**
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	//
	// *   IP addresses, such as 10.23.XX.XX.
	// *   CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifyClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

type ModifyClusterAccessWhiteListResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

type ModifyClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAccessWhiteListResponse) SetStatusCode(v int32) *ModifyClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterAccessWhiteListResponse) SetBody(v *ModifyClusterAccessWhiteListResponseBody) *ModifyClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type ModifyClusterConnectionStringRequest struct {
	// The prefix of the public endpoint.
	//
	// *   The prefix can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	// *   The prefix can be up to 30 characters in length.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The public endpoint of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The port number. Set the value to **3306**.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
	return s
}

type ModifyClusterConnectionStringResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponseBody) SetRequestId(v string) *ModifyClusterConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterConnectionStringResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyClusterConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetStatusCode(v int32) *ModifyClusterConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetBody(v *ModifyClusterConnectionStringResponseBody) *ModifyClusterConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBClusterRequest struct {
	// The reserved computing resources. Unit: ACUs. Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~454250~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to allocate all reserved computing resources to the user_default resource group. Valid values:
	//
	// *   true (default)
	// *   false
	EnableDefaultResourcePool *bool   `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The reserved storage resources. Unit: ACUs. Valid values: 0 to 2064. The value must be in increments of 24 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetEnableDefaultResourcePool(v bool) *ModifyDBClusterRequest {
	s.EnableDefaultResourcePool = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

type ModifyDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetDBClusterId(v string) *ModifyDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetOrderId(v string) *ModifyDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResponse) SetStatusCode(v int32) *ModifyDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https`.
	// *   The description must be 2 to 256 characters in length.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

type ModifyDBClusterDescriptionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterDescriptionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetStatusCode(v int32) *ModifyDBClusterDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. It must be in the hh:mmZ-hh:mmZ format.
	//
	// > The interval must be 1 hour on the hour.
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

type ModifyDBClusterMaintainTimeResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBClusterMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBResourceGroupRequest struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the preemptible instance feature for the resource group. This feature can be enabled only for job resource groups. Valid values:
	//
	// *   **True**
	// *   **False**
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of a resource group in a cluster.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACU.
	//
	// *   If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16 ACUs.
	// *   If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8 ACUs.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: AnalyticDB compute units (ACUs).
	//
	// *   If the GroupType parameter is set to Interactive, set the value to 16ACU.
	// *   If GroupType is set to Job, set the value to 0ACU.
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Rules    []*ModifyDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequest) SetClusterMode(v string) *ModifyDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetDBClusterId(v string) *ModifyDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetEnableSpot(v bool) *ModifyDBResourceGroupRequest {
	s.EnableSpot = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupName(v string) *ModifyDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupType(v string) *ModifyDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetRegionId(v string) *ModifyDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetRules(v []*ModifyDBResourceGroupRequestRules) *ModifyDBResourceGroupRequest {
	s.Rules = v
	return s
}

type ModifyDBResourceGroupRequestRules struct {
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	QueryTime       *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s ModifyDBResourceGroupRequestRules) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRules) SetGroupName(v string) *ModifyDBResourceGroupRequestRules {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRules) SetQueryTime(v string) *ModifyDBResourceGroupRequestRules {
	s.QueryTime = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRules) SetTargetGroupName(v string) *ModifyDBResourceGroupRequestRules {
	s.TargetGroupName = &v
	return s
}

type ModifyDBResourceGroupShrinkRequest struct {
	// A reserved parameter.
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the preemptible instance feature for the resource group. This feature can be enabled only for job resource groups. Valid values:
	//
	// *   **True**
	// *   **False**
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the name of a resource group in a cluster.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// *   **Interactive**
	// *   **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](~~428610~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACU.
	//
	// *   If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16 ACUs.
	// *   If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8 ACUs.
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: AnalyticDB compute units (ACUs).
	//
	// *   If the GroupType parameter is set to Interactive, set the value to 16ACU.
	// *   If GroupType is set to Job, set the value to 0ACU.
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~454314~~) operation to query the most recent region list.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyDBResourceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterMode(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetDBClusterId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetEnableSpot(v bool) *ModifyDBResourceGroupShrinkRequest {
	s.EnableSpot = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupName(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupType(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetRegionId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetRulesShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.RulesShrink = &v
	return s
}

type ModifyDBResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponseBody) SetRequestId(v string) *ModifyDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetStatusCode(v int32) *ModifyDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetBody(v *ModifyDBResourceGroupResponseBody) *ModifyDBResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyElasticPlanRequest struct {
	// A CORN expression that specifies the scaling cycle and time for the scaling plan.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The end time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The desired specifications of elastic resources after scaling.
	//
	// >
	//
	// *   If the scaling plan uses **EIUs** and **Default Proportional Scaling for EIUs** is enabled, you do not need to specify this parameter. In other cases, you must specify this parameter.
	//
	// *   You can call the [DescribeElasticPlanSpecifications](~~601278~~) operation to query the specifications that are supported for scaling plans.
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) SetCronExpression(v string) *ModifyElasticPlanRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetEndTime(v string) *ModifyElasticPlanRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetStartTime(v string) *ModifyElasticPlanRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetTargetSize(v string) *ModifyElasticPlanRequest {
	s.TargetSize = &v
	return s
}

type ModifyElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBody) SetRequestId(v string) *ModifyElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponse) SetHeaders(v map[string]*string) *ModifyElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticPlanResponse) SetStatusCode(v int32) *ModifyElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticPlanResponse) SetBody(v *ModifyElasticPlanResponseBody) *ModifyElasticPlanResponse {
	s.Body = v
	return s
}

type PreloadSparkAppMetricsRequest struct {
	// The Spark application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s PreloadSparkAppMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadSparkAppMetricsRequest) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsRequest) SetAppId(v string) *PreloadSparkAppMetricsRequest {
	s.AppId = &v
	return s
}

func (s *PreloadSparkAppMetricsRequest) SetDBClusterId(v string) *PreloadSparkAppMetricsRequest {
	s.DBClusterId = &v
	return s
}

type PreloadSparkAppMetricsResponseBody struct {
	// The returned data.
	Data *PreloadSparkAppMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadSparkAppMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBody) SetData(v *PreloadSparkAppMetricsResponseBodyData) *PreloadSparkAppMetricsResponseBody {
	s.Data = v
	return s
}

func (s *PreloadSparkAppMetricsResponseBody) SetRequestId(v string) *PreloadSparkAppMetricsResponseBody {
	s.RequestId = &v
	return s
}

type PreloadSparkAppMetricsResponseBodyData struct {
	// The ID of the Spark application.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The retry ID of the Spark application.
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The event log path.
	EventLogPath *string `json:"EventLogPath,omitempty" xml:"EventLogPath,omitempty"`
	// Indicates whether parsing is complete. Valid values:
	//
	// *   true
	// *   false
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The metrics.
	ScanMetrics *PreloadSparkAppMetricsResponseBodyDataScanMetrics `json:"ScanMetrics,omitempty" xml:"ScanMetrics,omitempty" type:"Struct"`
}

func (s PreloadSparkAppMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetAppId(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetAttemptId(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetEventLogPath(v string) *PreloadSparkAppMetricsResponseBodyData {
	s.EventLogPath = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetFinished(v bool) *PreloadSparkAppMetricsResponseBodyData {
	s.Finished = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyData) SetScanMetrics(v *PreloadSparkAppMetricsResponseBodyDataScanMetrics) *PreloadSparkAppMetricsResponseBodyData {
	s.ScanMetrics = v
	return s
}

type PreloadSparkAppMetricsResponseBodyDataScanMetrics struct {
	// The number of rows scanned.
	OutputRowsCount *int64 `json:"OutputRowsCount,omitempty" xml:"OutputRowsCount,omitempty"`
	// The size of the scanned data. Unit: bytes.
	TotalReadFileSizeInByte *int64 `json:"TotalReadFileSizeInByte,omitempty" xml:"TotalReadFileSizeInByte,omitempty"`
}

func (s PreloadSparkAppMetricsResponseBodyDataScanMetrics) String() string {
	return tea.Prettify(s)
}

func (s PreloadSparkAppMetricsResponseBodyDataScanMetrics) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) SetOutputRowsCount(v int64) *PreloadSparkAppMetricsResponseBodyDataScanMetrics {
	s.OutputRowsCount = &v
	return s
}

func (s *PreloadSparkAppMetricsResponseBodyDataScanMetrics) SetTotalReadFileSizeInByte(v int64) *PreloadSparkAppMetricsResponseBodyDataScanMetrics {
	s.TotalReadFileSizeInByte = &v
	return s
}

type PreloadSparkAppMetricsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadSparkAppMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadSparkAppMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s PreloadSparkAppMetricsResponse) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponse) SetHeaders(v map[string]*string) *PreloadSparkAppMetricsResponse {
	s.Headers = v
	return s
}

func (s *PreloadSparkAppMetricsResponse) SetStatusCode(v int32) *PreloadSparkAppMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadSparkAppMetricsResponse) SetBody(v *PreloadSparkAppMetricsResponseBody) *PreloadSparkAppMetricsResponse {
	s.Body = v
	return s
}

type ReleaseClusterPublicConnectionRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

type ReleaseClusterPublicConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponseBody) SetRequestId(v string) *ReleaseClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseClusterPublicConnectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetStatusCode(v int32) *ReleaseClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type RenameSparkTemplateFileRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The template file ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the template file that you want to rename.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RenameSparkTemplateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameSparkTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *RenameSparkTemplateFileRequest) SetDBClusterId(v string) *RenameSparkTemplateFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *RenameSparkTemplateFileRequest) SetId(v int64) *RenameSparkTemplateFileRequest {
	s.Id = &v
	return s
}

func (s *RenameSparkTemplateFileRequest) SetName(v string) *RenameSparkTemplateFileRequest {
	s.Name = &v
	return s
}

type RenameSparkTemplateFileResponseBody struct {
	// The data returned.
	Data *RenameSparkTemplateFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameSparkTemplateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameSparkTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *RenameSparkTemplateFileResponseBody) SetData(v *RenameSparkTemplateFileResponseBodyData) *RenameSparkTemplateFileResponseBody {
	s.Data = v
	return s
}

func (s *RenameSparkTemplateFileResponseBody) SetRequestId(v string) *RenameSparkTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

type RenameSparkTemplateFileResponseBodyData struct {
	// Indicates whether the request was successful. Valid values:
	//
	// *   True
	// *   False
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s RenameSparkTemplateFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenameSparkTemplateFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenameSparkTemplateFileResponseBodyData) SetSucceeded(v bool) *RenameSparkTemplateFileResponseBodyData {
	s.Succeeded = &v
	return s
}

type RenameSparkTemplateFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameSparkTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameSparkTemplateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameSparkTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *RenameSparkTemplateFileResponse) SetHeaders(v map[string]*string) *RenameSparkTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *RenameSparkTemplateFileResponse) SetStatusCode(v int32) *RenameSparkTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameSparkTemplateFileResponse) SetBody(v *RenameSparkTemplateFileResponseBody) *RenameSparkTemplateFileResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	// The description of the database account.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// > You can call the [DescribeAccounts](~~612430~~) operation to query the information about database accounts in a cluster, including the database account name.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// *   The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	// *   Special characters include `! @ # $ % ^ & * ( ) _ + - =`
	// *   The password must be 8 to 32 characters in length.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccountDescription(v string) *ResetAccountPasswordRequest {
	s.AccountDescription = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type SetSparkAppLogRootPathRequest struct {
	// The database ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Object Storage Service (OSS) log path.
	OssLogPath *string `json:"OssLogPath,omitempty" xml:"OssLogPath,omitempty"`
	// Specifies whether to use the default OSS log path.
	UseDefaultOss *bool `json:"UseDefaultOss,omitempty" xml:"UseDefaultOss,omitempty"`
}

func (s SetSparkAppLogRootPathRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSparkAppLogRootPathRequest) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathRequest) SetDBClusterId(v string) *SetSparkAppLogRootPathRequest {
	s.DBClusterId = &v
	return s
}

func (s *SetSparkAppLogRootPathRequest) SetOssLogPath(v string) *SetSparkAppLogRootPathRequest {
	s.OssLogPath = &v
	return s
}

func (s *SetSparkAppLogRootPathRequest) SetUseDefaultOss(v bool) *SetSparkAppLogRootPathRequest {
	s.UseDefaultOss = &v
	return s
}

type SetSparkAppLogRootPathResponseBody struct {
	// The returned data.
	Data *SetSparkAppLogRootPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSparkAppLogRootPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSparkAppLogRootPathResponseBody) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponseBody) SetData(v *SetSparkAppLogRootPathResponseBodyData) *SetSparkAppLogRootPathResponseBody {
	s.Data = v
	return s
}

func (s *SetSparkAppLogRootPathResponseBody) SetRequestId(v string) *SetSparkAppLogRootPathResponseBody {
	s.RequestId = &v
	return s
}

type SetSparkAppLogRootPathResponseBodyData struct {
	// The recommended default OSS log path.
	DefaultLogPath *string `json:"DefaultLogPath,omitempty" xml:"DefaultLogPath,omitempty"`
	// Indicates whether an OSS log path exists.
	IsLogPathExists *bool `json:"IsLogPathExists,omitempty" xml:"IsLogPathExists,omitempty"`
	// The time when the modification was last modified.
	ModifiedTimestamp *string `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The modifier ID.
	ModifiedUid *string `json:"ModifiedUid,omitempty" xml:"ModifiedUid,omitempty"`
	// The OSS log path.
	RecordedLogPath *string `json:"RecordedLogPath,omitempty" xml:"RecordedLogPath,omitempty"`
}

func (s SetSparkAppLogRootPathResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetSparkAppLogRootPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetDefaultLogPath(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.DefaultLogPath = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetIsLogPathExists(v bool) *SetSparkAppLogRootPathResponseBodyData {
	s.IsLogPathExists = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetModifiedTimestamp(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.ModifiedTimestamp = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetModifiedUid(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.ModifiedUid = &v
	return s
}

func (s *SetSparkAppLogRootPathResponseBodyData) SetRecordedLogPath(v string) *SetSparkAppLogRootPathResponseBodyData {
	s.RecordedLogPath = &v
	return s
}

type SetSparkAppLogRootPathResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSparkAppLogRootPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSparkAppLogRootPathResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSparkAppLogRootPathResponse) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponse) SetHeaders(v map[string]*string) *SetSparkAppLogRootPathResponse {
	s.Headers = v
	return s
}

func (s *SetSparkAppLogRootPathResponse) SetStatusCode(v int32) *SetSparkAppLogRootPathResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSparkAppLogRootPathResponse) SetBody(v *SetSparkAppLogRootPathResponseBody) *SetSparkAppLogRootPathResponse {
	s.Body = v
	return s
}

type StartSparkSQLEngineRequest struct {
	// The configuration that is required to start the Spark SQL engine. Specify this value in the JSON format. For more information, see [Conf configuration parameters](~~471203~~).
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Object Storage Service (OSS) paths of third-party JAR packages that are required to start the Spark SQL engine. Separate multiple OSS paths with commas (,).
	Jars *string `json:"Jars,omitempty" xml:"Jars,omitempty"`
	// The maximum number of executors that are required to execute SQL statements. Valid values: 1 to 2000. If this value exceeds the total number of executes that are supported by the resource group, the Spark SQL engine fails to be started.
	MaxExecutor *int64 `json:"MaxExecutor,omitempty" xml:"MaxExecutor,omitempty"`
	// The minimum number of executors that are required to execute SQL statements. Valid values: 0 to 2000. A value of 0 indicates that no executors are permanent if no SQL statements are executed. If this value exceeds the total number of executes that are supported by the resource group, the Spark SQL engine fails to be started. The value must be less than the value of MaxExecutor.
	MinExecutor *int64 `json:"MinExecutor,omitempty" xml:"MinExecutor,omitempty"`
	// The name of the resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The maximum number of slots that are required to maintain Spark sessions for executing SQL statements. Valid values: 1 to 500.
	SlotNum *int64 `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
}

func (s StartSparkSQLEngineRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSparkSQLEngineRequest) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineRequest) SetConfig(v string) *StartSparkSQLEngineRequest {
	s.Config = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetDBClusterId(v string) *StartSparkSQLEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetJars(v string) *StartSparkSQLEngineRequest {
	s.Jars = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetMaxExecutor(v int64) *StartSparkSQLEngineRequest {
	s.MaxExecutor = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetMinExecutor(v int64) *StartSparkSQLEngineRequest {
	s.MinExecutor = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetResourceGroupName(v string) *StartSparkSQLEngineRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetSlotNum(v int64) *StartSparkSQLEngineRequest {
	s.SlotNum = &v
	return s
}

type StartSparkSQLEngineResponseBody struct {
	// The returned data.
	Data *StartSparkSQLEngineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartSparkSQLEngineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSparkSQLEngineResponseBody) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponseBody) SetData(v *StartSparkSQLEngineResponseBodyData) *StartSparkSQLEngineResponseBody {
	s.Data = v
	return s
}

func (s *StartSparkSQLEngineResponseBody) SetRequestId(v string) *StartSparkSQLEngineResponseBody {
	s.RequestId = &v
	return s
}

type StartSparkSQLEngineResponseBodyData struct {
	// The ID of the Spark job.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the Spark application.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The state of the Spark SQL engine. Valid values:
	//
	// *   SUBMITTED
	// *   STARTING
	// *   RUNNING
	// *   FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s StartSparkSQLEngineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartSparkSQLEngineResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponseBodyData) SetAppId(v string) *StartSparkSQLEngineResponseBodyData {
	s.AppId = &v
	return s
}

func (s *StartSparkSQLEngineResponseBodyData) SetAppName(v string) *StartSparkSQLEngineResponseBodyData {
	s.AppName = &v
	return s
}

func (s *StartSparkSQLEngineResponseBodyData) SetState(v string) *StartSparkSQLEngineResponseBodyData {
	s.State = &v
	return s
}

type StartSparkSQLEngineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSparkSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSparkSQLEngineResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSparkSQLEngineResponse) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponse) SetHeaders(v map[string]*string) *StartSparkSQLEngineResponse {
	s.Headers = v
	return s
}

func (s *StartSparkSQLEngineResponse) SetStatusCode(v int32) *StartSparkSQLEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSparkSQLEngineResponse) SetBody(v *StartSparkSQLEngineResponseBody) *StartSparkSQLEngineResponse {
	s.Body = v
	return s
}

type SubmitSparkAppRequest struct {
	// The type of the client. The value can be up to 64 characters in length.
	AgentSource *string `json:"AgentSource,omitempty" xml:"AgentSource,omitempty"`
	// The version of the client. The value can be up to 64 characters in length.
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The name of the application. The value can be up to 64 characters in length.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// *   **SQL**
	// *   **STREAMING**
	// *   **BATCH** (default)
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~454250~~) operation to query cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data of the application template.
	//
	// > For information about the application template configuration, see [Spark application configuration guide](~~452402~~).
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The name of the job resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~612410~~) operation to query the resource group ID of an AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The ID of the application template.
	//
	// > You can call the [GetSparkTemplateFullTree](~~456205~~) operation to query the application template ID.
	TemplateFileId *int64 `json:"TemplateFileId,omitempty" xml:"TemplateFileId,omitempty"`
}

func (s SubmitSparkAppRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkAppRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppRequest) SetAgentSource(v string) *SubmitSparkAppRequest {
	s.AgentSource = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAgentVersion(v string) *SubmitSparkAppRequest {
	s.AgentVersion = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAppName(v string) *SubmitSparkAppRequest {
	s.AppName = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAppType(v string) *SubmitSparkAppRequest {
	s.AppType = &v
	return s
}

func (s *SubmitSparkAppRequest) SetDBClusterId(v string) *SubmitSparkAppRequest {
	s.DBClusterId = &v
	return s
}

func (s *SubmitSparkAppRequest) SetData(v string) *SubmitSparkAppRequest {
	s.Data = &v
	return s
}

func (s *SubmitSparkAppRequest) SetResourceGroupName(v string) *SubmitSparkAppRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *SubmitSparkAppRequest) SetTemplateFileId(v int64) *SubmitSparkAppRequest {
	s.TemplateFileId = &v
	return s
}

type SubmitSparkAppResponseBody struct {
	// The returned data.
	Data *SubmitSparkAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkAppResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponseBody) SetData(v *SubmitSparkAppResponseBodyData) *SubmitSparkAppResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSparkAppResponseBody) SetRequestId(v string) *SubmitSparkAppResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSparkAppResponseBodyData struct {
	// The application ID.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The alert message returned for the operation, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the application. Valid values:
	//
	// *   **SUBMITTED**
	// *   **STARTING**
	// *   **RUNNING**
	// *   **FAILING**
	// *   **FAILED**
	// *   **KILLING**
	// *   **KILLED**
	// *   **SUCCEEDING**
	// *   **COMPLETED**
	// *   **FATAL**
	// *   **UNKNOWN**
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitSparkAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponseBodyData) SetAppId(v string) *SubmitSparkAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetAppName(v string) *SubmitSparkAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetMessage(v string) *SubmitSparkAppResponseBodyData {
	s.Message = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetState(v string) *SubmitSparkAppResponseBodyData {
	s.State = &v
	return s
}

type SubmitSparkAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSparkAppResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkAppResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponse) SetHeaders(v map[string]*string) *SubmitSparkAppResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkAppResponse) SetStatusCode(v int32) *SubmitSparkAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkAppResponse) SetBody(v *SubmitSparkAppResponseBody) *SubmitSparkAppResponse {
	s.Body = v
	return s
}

type SubmitSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark application.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskRequest) SetAppId(v string) *SubmitSparkLogAnalyzeTaskRequest {
	s.AppId = &v
	return s
}

type SubmitSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *SubmitSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *SubmitSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSparkLogAnalyzeTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkLogAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetHeaders(v map[string]*string) *SubmitSparkLogAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetStatusCode(v int32) *SubmitSparkLogAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponse) SetBody(v *SubmitSparkLogAnalyzeTaskResponseBody) *SubmitSparkLogAnalyzeTaskResponse {
	s.Body = v
	return s
}

type UnbindAccountRequest struct {
	// The name of the database account.
	//
	// >  You can call the [DescribeAccounts](~~612430~~) operation to query the information about database accounts of an AnalyticDB for MySQL cluster, including database account names.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s UnbindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountRequest) GoString() string {
	return s.String()
}

func (s *UnbindAccountRequest) SetAccountName(v string) *UnbindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *UnbindAccountRequest) SetDBClusterId(v string) *UnbindAccountRequest {
	s.DBClusterId = &v
	return s
}

type UnbindAccountResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAccountResponseBody) SetRequestId(v string) *UnbindAccountResponseBody {
	s.RequestId = &v
	return s
}

type UnbindAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountResponse) GoString() string {
	return s.String()
}

func (s *UnbindAccountResponse) SetHeaders(v map[string]*string) *UnbindAccountResponse {
	s.Headers = v
	return s
}

func (s *UnbindAccountResponse) SetStatusCode(v int32) *UnbindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAccountResponse) SetBody(v *UnbindAccountResponseBody) *UnbindAccountResponse {
	s.Body = v
	return s
}

type UnbindDBResourceGroupWithUserRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the database account.
	GroupUser *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
}

func (s UnbindDBResourceGroupWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

type UnbindDBResourceGroupWithUserResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourceGroupWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *UnbindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDBResourceGroupWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *UnbindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetBody(v *UnbindDBResourceGroupWithUserResponseBody) *UnbindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

type UpdateSparkTemplateFileRequest struct {
	// The template data to be updated.
	//
	// >  If you do not specify this parameter, the application template is not updated. For information about how to configure a Spark application template, see [Configure a Spark application](~~452402~~).
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application template ID.
	//
	// >  You can call the [GetSparkTemplateFullTree](~~456205~~) operation to query the application template ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the job resource group.
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s UpdateSparkTemplateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSparkTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileRequest) SetContent(v string) *UpdateSparkTemplateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetDBClusterId(v string) *UpdateSparkTemplateFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetId(v int64) *UpdateSparkTemplateFileRequest {
	s.Id = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetResourceGroupName(v string) *UpdateSparkTemplateFileRequest {
	s.ResourceGroupName = &v
	return s
}

type UpdateSparkTemplateFileResponseBody struct {
	// The update result.
	Data *UpdateSparkTemplateFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSparkTemplateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSparkTemplateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponseBody) SetData(v *UpdateSparkTemplateFileResponseBodyData) *UpdateSparkTemplateFileResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSparkTemplateFileResponseBody) SetRequestId(v string) *UpdateSparkTemplateFileResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSparkTemplateFileResponseBodyData struct {
	// Indicates whether the application template is updated.
	//
	// *   **True**
	// *   **False**
	Succeeded *bool `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s UpdateSparkTemplateFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateSparkTemplateFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponseBodyData) SetSucceeded(v bool) *UpdateSparkTemplateFileResponseBodyData {
	s.Succeeded = &v
	return s
}

type UpdateSparkTemplateFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSparkTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSparkTemplateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSparkTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponse) SetHeaders(v map[string]*string) *UpdateSparkTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateSparkTemplateFileResponse) SetStatusCode(v int32) *UpdateSparkTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSparkTemplateFileResponse) SetBody(v *UpdateSparkTemplateFileResponseBody) *UpdateSparkTemplateFileResponse {
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
		"cn-qingdao":                  tea.String("adb.aliyuncs.com"),
		"cn-beijing":                  tea.String("adb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("adb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("adb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("adb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("adb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("adb.aliyuncs.com"),
		"us-west-1":                   tea.String("adb.aliyuncs.com"),
		"us-east-1":                   tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("adb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("adb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("adb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("adb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("adb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("adb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("adb.aliyuncs.com"),
		"cn-fujian":                   tea.String("adb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("adb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("adb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("adb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("adb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("adb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("adb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("adb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("adb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("adb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("adb.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("adb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateClusterPublicConnectionWithOptions(request *AllocateClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateClusterPublicConnection"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateClusterPublicConnection(request *AllocateClusterPublicConnectionRequest) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.AllocateClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachUserENIWithOptions(request *AttachUserENIRequest, runtime *util.RuntimeOptions) (_result *AttachUserENIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachUserENI"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachUserENI(request *AttachUserENIRequest) (_result *AttachUserENIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachUserENIResponse{}
	_body, _err := client.AttachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAccountWithOptions(request *BindAccountRequest, runtime *util.RuntimeOptions) (_result *BindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RamUser)) {
		query["RamUser"] = request.RamUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAccount"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAccount(request *BindAccountRequest) (_result *BindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAccountResponse{}
	_body, _err := client.BindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindDBResourceGroupWithUserWithOptions(request *BindDBResourceGroupWithUserRequest, runtime *util.RuntimeOptions) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUser)) {
		query["GroupUser"] = request.GroupUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindDBResourceGroupWithUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindDBResourceGroupWithUser(request *BindDBResourceGroupWithUserRequest) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.BindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckBindRamUserWithOptions(request *CheckBindRamUserRequest, runtime *util.RuntimeOptions) (_result *CheckBindRamUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckBindRamUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckBindRamUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckBindRamUser(request *CheckBindRamUserRequest) (_result *CheckBindRamUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckBindRamUserResponse{}
	_body, _err := client.CheckBindRamUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSampleDataSetWithOptions(request *CheckSampleDataSetRequest, runtime *util.RuntimeOptions) (_result *CheckSampleDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSampleDataSet"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSampleDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSampleDataSet(request *CheckSampleDataSetRequest) (_result *CheckSampleDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSampleDataSetResponse{}
	_body, _err := client.CheckSampleDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterNetworkType)) {
		query["DBClusterNetworkType"] = request.DBClusterNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultResourcePool)) {
		query["EnableDefaultResourcePool"] = request.EnableDefaultResourcePool
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreToTime)) {
		query["RestoreToTime"] = request.RestoreToTime
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreType)) {
		query["RestoreType"] = request.RestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDbClusterId)) {
		query["SourceDbClusterId"] = request.SourceDbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBCluster"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBResourceGroupWithOptions(tmpReq *CreateDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rules)) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, tea.String("Rules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterMode)) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterSizeResource)) {
		query["ClusterSizeResource"] = request.ClusterSizeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSpot)) {
		query["EnableSpot"] = request.EnableSpot
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxClusterCount)) {
		query["MaxClusterCount"] = request.MaxClusterCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxComputeResource)) {
		query["MaxComputeResource"] = request.MaxComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.MinClusterCount)) {
		query["MinClusterCount"] = request.MinClusterCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinComputeResource)) {
		query["MinComputeResource"] = request.MinComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RulesShrink)) {
		query["Rules"] = request.RulesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBResourceGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBResourceGroup(request *CreateDBResourceGroupRequest) (_result *CreateDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CreateDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateElasticPlanWithOptions(request *CreateElasticPlanRequest, runtime *util.RuntimeOptions) (_result *CreateElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoScale)) {
		query["AutoScale"] = request.AutoScale
	}

	if !tea.BoolValue(util.IsUnset(request.CronExpression)) {
		query["CronExpression"] = request.CronExpression
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSize)) {
		query["TargetSize"] = request.TargetSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateElasticPlan"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateElasticPlan(request *CreateElasticPlanRequest) (_result *CreateElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CreateElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOssSubDirectoryWithOptions(request *CreateOssSubDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateOssSubDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOssSubDirectory"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOssSubDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOssSubDirectory(request *CreateOssSubDirectoryRequest) (_result *CreateOssSubDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOssSubDirectoryResponse{}
	_body, _err := client.CreateOssSubDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSparkTemplateWithOptions(request *CreateSparkTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateSparkTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSparkTemplate"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSparkTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSparkTemplate(request *CreateSparkTemplateRequest) (_result *CreateSparkTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSparkTemplateResponse{}
	_body, _err := client.CreateSparkTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * You can call this operation to delete only subscription clusters.
 *
 * @param request DeleteDBClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBClusterResponse
 */
func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCluster"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * You can call this operation to delete only subscription clusters.
 *
 * @param request DeleteDBClusterRequest
 * @return DeleteDBClusterResponse
 */
func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBResourceGroupWithOptions(request *DeleteDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBResourceGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBResourceGroup(request *DeleteDBResourceGroupRequest) (_result *DeleteDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.DeleteDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteElasticPlanWithOptions(request *DeleteElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteElasticPlan"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteElasticPlan(request *DeleteElasticPlanRequest) (_result *DeleteElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.DeleteElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProcessInstanceWithOptions(request *DeleteProcessInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		query["ProjectCode"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProcessInstance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProcessInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProcessInstance(request *DeleteProcessInstanceRequest) (_result *DeleteProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProcessInstanceResponse{}
	_body, _err := client.DeleteProcessInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSparkTemplateWithOptions(request *DeleteSparkTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteSparkTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSparkTemplate"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSparkTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSparkTemplate(request *DeleteSparkTemplateRequest) (_result *DeleteSparkTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSparkTemplateResponse{}
	_body, _err := client.DeleteSparkTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSparkTemplateFileWithOptions(request *DeleteSparkTemplateFileRequest, runtime *util.RuntimeOptions) (_result *DeleteSparkTemplateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSparkTemplateFile"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSparkTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSparkTemplateFile(request *DeleteSparkTemplateFileRequest) (_result *DeleteSparkTemplateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSparkTemplateFileResponse{}
	_body, _err := client.DeleteSparkTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountAllPrivilegesWithOptions(request *DescribeAccountAllPrivilegesRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountAllPrivilegesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountAllPrivileges"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountAllPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountAllPrivileges(request *DescribeAccountAllPrivilegesRequest) (_result *DescribeAccountAllPrivilegesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountAllPrivilegesResponse{}
	_body, _err := client.DescribeAccountAllPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountPrivilegeObjectsWithOptions(request *DescribeAccountPrivilegeObjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountPrivilegeObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnPrivilegeObject)) {
		query["ColumnPrivilegeObject"] = request.ColumnPrivilegeObject
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabasePrivilegeObject)) {
		query["DatabasePrivilegeObject"] = request.DatabasePrivilegeObject
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrivilegeType)) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TablePrivilegeObject)) {
		query["TablePrivilegeObject"] = request.TablePrivilegeObject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountPrivilegeObjects"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountPrivilegeObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountPrivilegeObjects(request *DescribeAccountPrivilegeObjectsRequest) (_result *DescribeAccountPrivilegeObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountPrivilegeObjectsResponse{}
	_body, _err := client.DescribeAccountPrivilegeObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountPrivilegesWithOptions(request *DescribeAccountPrivilegesRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountPrivilegesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnPrivilegeObject)) {
		query["ColumnPrivilegeObject"] = request.ColumnPrivilegeObject
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabasePrivilegeObject)) {
		query["DatabasePrivilegeObject"] = request.DatabasePrivilegeObject
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrivilegeType)) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TablePrivilegeObject)) {
		query["TablePrivilegeObject"] = request.TablePrivilegeObject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountPrivileges"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (_result *DescribeAccountPrivilegesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountPrivilegesResponse{}
	_body, _err := client.DescribeAccountPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdbMySqlColumnsWithOptions(request *DescribeAdbMySqlColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeAdbMySqlColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["Schema"] = request.Schema
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdbMySqlColumns"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdbMySqlColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdbMySqlColumns(request *DescribeAdbMySqlColumnsRequest) (_result *DescribeAdbMySqlColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdbMySqlColumnsResponse{}
	_body, _err := client.DescribeAdbMySqlColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdbMySqlSchemasWithOptions(request *DescribeAdbMySqlSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeAdbMySqlSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdbMySqlSchemas"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdbMySqlSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdbMySqlSchemas(request *DescribeAdbMySqlSchemasRequest) (_result *DescribeAdbMySqlSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdbMySqlSchemasResponse{}
	_body, _err := client.DescribeAdbMySqlSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdbMySqlTablesWithOptions(request *DescribeAdbMySqlTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeAdbMySqlTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		query["Schema"] = request.Schema
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdbMySqlTables"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdbMySqlTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdbMySqlTables(request *DescribeAdbMySqlTablesRequest) (_result *DescribeAdbMySqlTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdbMySqlTablesResponse{}
	_body, _err := client.DescribeAdbMySqlTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllDataSource"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (_result *DescribeAllDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DescribeAllDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApsActionLogsWithOptions(request *DescribeApsActionLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeApsActionLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Stage)) {
		query["Stage"] = request.Stage
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadId)) {
		query["WorkloadId"] = request.WorkloadId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApsActionLogs"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApsActionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApsActionLogs(request *DescribeApsActionLogsRequest) (_result *DescribeApsActionLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApsActionLogsResponse{}
	_body, _err := client.DescribeApsActionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApsResourceGroupsWithOptions(request *DescribeApsResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeApsResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadId)) {
		body["WorkloadId"] = request.WorkloadId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApsResourceGroups"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApsResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApsResourceGroups(request *DescribeApsResourceGroupsRequest) (_result *DescribeApsResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApsResourceGroupsResponse{}
	_body, _err := client.DescribeApsResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
 *
 * @param request DescribeAuditLogRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAuditLogRecordsResponse
 */
func (client *Client) DescribeAuditLogRecordsWithOptions(request *DescribeAuditLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUser)) {
		query["ProxyUser"] = request.ProxyUser
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Succeed)) {
		query["Succeed"] = request.Succeed
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditLogRecords"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
 *
 * @param request DescribeAuditLogRecordsRequest
 * @return DescribeAuditLogRecordsResponse
 */
func (client *Client) DescribeAuditLogRecords(request *DescribeAuditLogRecordsRequest) (_result *DescribeAuditLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.DescribeAuditLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterAccessWhiteListWithOptions(request *DescribeClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterAccessWhiteList"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAccessWhiteList(request *DescribeClusterAccessWhiteListRequest) (_result *DescribeClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterAccessWhiteListResponse{}
	_body, _err := client.DescribeClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNetInfoWithOptions(request *DescribeClusterNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterNetInfo"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNetInfo(request *DescribeClusterNetInfoRequest) (_result *DescribeClusterNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNetInfoResponse{}
	_body, _err := client.DescribeClusterNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterResourceDetailWithOptions(request *DescribeClusterResourceDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResourceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterResourceDetail"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResourceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterResourceDetail(request *DescribeClusterResourceDetailRequest) (_result *DescribeClusterResourceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResourceDetailResponse{}
	_body, _err := client.DescribeClusterResourceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterResourceUsageWithOptions(request *DescribeClusterResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterResourceUsage"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterResourceUsage(request *DescribeClusterResourceUsageRequest) (_result *DescribeClusterResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResourceUsageResponse{}
	_body, _err := client.DescribeClusterResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumns"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComputeResourceUsageWithOptions(request *DescribeComputeResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeComputeResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComputeResourceUsage"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComputeResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComputeResourceUsage(request *DescribeComputeResourceUsageRequest) (_result *DescribeComputeResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComputeResourceUsageResponse{}
	_body, _err := client.DescribeComputeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAttribute"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthStatusWithOptions(request *DescribeDBClusterHealthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterHealthStatus"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthStatus(request *DescribeDBClusterHealthStatusRequest) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.DescribeDBClusterHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePools)) {
		query["ResourcePools"] = request.ResourcePools
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterSpaceSummaryWithOptions(request *DescribeDBClusterSpaceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterSpaceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterSpaceSummary"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterSpaceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterSpaceSummary(request *DescribeDBClusterSpaceSummaryRequest) (_result *DescribeDBClusterSpaceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterSpaceSummaryResponse{}
	_body, _err := client.DescribeDBClusterSpaceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterStatusWithOptions(request *DescribeDBClusterStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterStatus"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterStatus(request *DescribeDBClusterStatusRequest) (_result *DescribeDBClusterStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.DescribeDBClusterStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIds)) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterStatus)) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusters"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBResourceGroupWithOptions(request *DescribeDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBResourceGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBResourceGroup(request *DescribeDBResourceGroupRequest) (_result *DescribeDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.DescribeDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisDimensions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
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

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisRecords"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisSQLInfo"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfo(request *DescribeDiagnosisSQLInfoRequest) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.DescribeDiagnosisSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadRecords"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticPlanAttributeWithOptions(request *DescribeElasticPlanAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlanAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlanAttribute"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlanAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticPlanAttribute(request *DescribeElasticPlanAttributeRequest) (_result *DescribeElasticPlanAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlanAttributeResponse{}
	_body, _err := client.DescribeElasticPlanAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticPlanJobsWithOptions(request *DescribeElasticPlanJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlanJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlanJobs"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlanJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticPlanJobs(request *DescribeElasticPlanJobsRequest) (_result *DescribeElasticPlanJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlanJobsResponse{}
	_body, _err := client.DescribeElasticPlanJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticPlanSpecificationsWithOptions(request *DescribeElasticPlanSpecificationsRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlanSpecificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlanSpecifications"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlanSpecificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticPlanSpecifications(request *DescribeElasticPlanSpecificationsRequest) (_result *DescribeElasticPlanSpecificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlanSpecificationsResponse{}
	_body, _err := client.DescribeElasticPlanSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticPlansWithOptions(request *DescribeElasticPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlans"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticPlans(request *DescribeElasticPlansRequest) (_result *DescribeElasticPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlansResponse{}
	_body, _err := client.DescribeElasticPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnabledPrivilegesWithOptions(request *DescribeEnabledPrivilegesRequest, runtime *util.RuntimeOptions) (_result *DescribeEnabledPrivilegesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEnabledPrivileges"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEnabledPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnabledPrivileges(request *DescribeEnabledPrivilegesRequest) (_result *DescribeEnabledPrivilegesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnabledPrivilegesResponse{}
	_body, _err := client.DescribeEnabledPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobResourceUsageWithOptions(request *DescribeJobResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeJobResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobResourceUsage"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobResourceUsage(request *DescribeJobResourceUsageRequest) (_result *DescribeJobResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobResourceUsageResponse{}
	_body, _err := client.DescribeJobResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePatternPerformanceWithOptions(request *DescribePatternPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribePatternPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePatternPerformance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePatternPerformance(request *DescribePatternPerformanceRequest) (_result *DescribePatternPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.DescribePatternPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPatternsWithOptions(request *DescribeSQLPatternsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPatterns"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPatterns(request *DescribeSQLPatternsRequest) (_result *DescribeSQLPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.DescribeSQLPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSchemas"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSchemas(request *DescribeSchemasRequest) (_result *DescribeSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.DescribeSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSparkCodeLogWithOptions(request *DescribeSparkCodeLogRequest, runtime *util.RuntimeOptions) (_result *DescribeSparkCodeLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSparkCodeLog"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSparkCodeLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSparkCodeLog(request *DescribeSparkCodeLogRequest) (_result *DescribeSparkCodeLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSparkCodeLogResponse{}
	_body, _err := client.DescribeSparkCodeLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSparkCodeOutputWithOptions(request *DescribeSparkCodeOutputRequest, runtime *util.RuntimeOptions) (_result *DescribeSparkCodeOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSparkCodeOutput"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSparkCodeOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSparkCodeOutput(request *DescribeSparkCodeOutputRequest) (_result *DescribeSparkCodeOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSparkCodeOutputResponse{}
	_body, _err := client.DescribeSparkCodeOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSparkCodeWebUiWithOptions(request *DescribeSparkCodeWebUiRequest, runtime *util.RuntimeOptions) (_result *DescribeSparkCodeWebUiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSparkCodeWebUi"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSparkCodeWebUiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSparkCodeWebUi(request *DescribeSparkCodeWebUiRequest) (_result *DescribeSparkCodeWebUiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSparkCodeWebUiResponse{}
	_body, _err := client.DescribeSparkCodeWebUiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlPatternWithOptions(request *DescribeSqlPatternRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlPattern)) {
		query["SqlPattern"] = request.SqlPattern
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlPattern"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlPattern(request *DescribeSqlPatternRequest) (_result *DescribeSqlPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.DescribeSqlPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStorageResourceUsageWithOptions(request *DescribeStorageResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeStorageResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStorageResourceUsage"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStorageResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStorageResourceUsage(request *DescribeStorageResourceUsageRequest) (_result *DescribeStorageResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStorageResourceUsageResponse{}
	_body, _err := client.DescribeStorageResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableAccessCountWithOptions(request *DescribeTableAccessCountRequest, runtime *util.RuntimeOptions) (_result *DescribeTableAccessCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableAccessCount"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableAccessCount(request *DescribeTableAccessCountRequest) (_result *DescribeTableAccessCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.DescribeTableAccessCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserQuota"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DetachUserENIWithOptions(request *DetachUserENIRequest, runtime *util.RuntimeOptions) (_result *DetachUserENIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachUserENI"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachUserENI(request *DetachUserENIRequest) (_result *DetachUserENIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachUserENIResponse{}
	_body, _err := client.DetachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableElasticPlanWithOptions(request *DisableElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DisableElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableElasticPlan"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableElasticPlan(request *DisableElasticPlanRequest) (_result *DisableElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableElasticPlanResponse{}
	_body, _err := client.DisableElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadDiagnosisRecords"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableElasticPlanWithOptions(request *EnableElasticPlanRequest, runtime *util.RuntimeOptions) (_result *EnableElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableElasticPlan"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableElasticPlan(request *EnableElasticPlanRequest) (_result *EnableElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableElasticPlanResponse{}
	_body, _err := client.EnableElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExistRunningSQLEngineWithOptions(request *ExistRunningSQLEngineRequest, runtime *util.RuntimeOptions) (_result *ExistRunningSQLEngineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExistRunningSQLEngine"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExistRunningSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExistRunningSQLEngine(request *ExistRunningSQLEngineRequest) (_result *ExistRunningSQLEngineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExistRunningSQLEngineResponse{}
	_body, _err := client.ExistRunningSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatabaseObjectsWithOptions(request *GetDatabaseObjectsRequest, runtime *util.RuntimeOptions) (_result *GetDatabaseObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterOwner)) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !tea.BoolValue(util.IsUnset(request.FilterSchemaName)) {
		query["FilterSchemaName"] = request.FilterSchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatabaseObjects"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatabaseObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDatabaseObjects(request *GetDatabaseObjectsRequest) (_result *GetDatabaseObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDatabaseObjectsResponse{}
	_body, _err := client.GetDatabaseObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppAttemptLogWithOptions(request *GetSparkAppAttemptLogRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppAttemptLogResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttemptId)) {
		body["AttemptId"] = request.AttemptId
	}

	if !tea.BoolValue(util.IsUnset(request.LogLength)) {
		body["LogLength"] = request.LogLength
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppAttemptLog"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppAttemptLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppAttemptLog(request *GetSparkAppAttemptLogRequest) (_result *GetSparkAppAttemptLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppAttemptLogResponse{}
	_body, _err := client.GetSparkAppAttemptLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppInfoWithOptions(request *GetSparkAppInfoRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppInfo"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppInfo(request *GetSparkAppInfoRequest) (_result *GetSparkAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppInfoResponse{}
	_body, _err := client.GetSparkAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppLogWithOptions(request *GetSparkAppLogRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LogLength)) {
		body["LogLength"] = request.LogLength
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppLog"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppLog(request *GetSparkAppLogRequest) (_result *GetSparkAppLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppLogResponse{}
	_body, _err := client.GetSparkAppLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppMetricsWithOptions(request *GetSparkAppMetricsRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppMetrics"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppMetrics(request *GetSparkAppMetricsRequest) (_result *GetSparkAppMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppMetricsResponse{}
	_body, _err := client.GetSparkAppMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppStateWithOptions(request *GetSparkAppStateRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppState"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppState(request *GetSparkAppStateRequest) (_result *GetSparkAppStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppStateResponse{}
	_body, _err := client.GetSparkAppStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkAppWebUiAddressWithOptions(request *GetSparkAppWebUiAddressRequest, runtime *util.RuntimeOptions) (_result *GetSparkAppWebUiAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkAppWebUiAddress"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkAppWebUiAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkAppWebUiAddress(request *GetSparkAppWebUiAddressRequest) (_result *GetSparkAppWebUiAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkAppWebUiAddressResponse{}
	_body, _err := client.GetSparkAppWebUiAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkConfigLogPathWithOptions(request *GetSparkConfigLogPathRequest, runtime *util.RuntimeOptions) (_result *GetSparkConfigLogPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkConfigLogPath"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkConfigLogPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkConfigLogPath(request *GetSparkConfigLogPathRequest) (_result *GetSparkConfigLogPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkConfigLogPathResponse{}
	_body, _err := client.GetSparkConfigLogPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkDefinitionsWithOptions(request *GetSparkDefinitionsRequest, runtime *util.RuntimeOptions) (_result *GetSparkDefinitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkDefinitions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkDefinitions(request *GetSparkDefinitionsRequest) (_result *GetSparkDefinitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkDefinitionsResponse{}
	_body, _err := client.GetSparkDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkLogAnalyzeTaskWithOptions(request *GetSparkLogAnalyzeTaskRequest, runtime *util.RuntimeOptions) (_result *GetSparkLogAnalyzeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkLogAnalyzeTask"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkLogAnalyzeTask(request *GetSparkLogAnalyzeTaskRequest) (_result *GetSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkLogAnalyzeTaskResponse{}
	_body, _err := client.GetSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkSQLEngineStateWithOptions(request *GetSparkSQLEngineStateRequest, runtime *util.RuntimeOptions) (_result *GetSparkSQLEngineStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkSQLEngineState"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkSQLEngineStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkSQLEngineState(request *GetSparkSQLEngineStateRequest) (_result *GetSparkSQLEngineStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkSQLEngineStateResponse{}
	_body, _err := client.GetSparkSQLEngineStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkTemplateFileContentWithOptions(request *GetSparkTemplateFileContentRequest, runtime *util.RuntimeOptions) (_result *GetSparkTemplateFileContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkTemplateFileContent"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkTemplateFileContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkTemplateFileContent(request *GetSparkTemplateFileContentRequest) (_result *GetSparkTemplateFileContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkTemplateFileContentResponse{}
	_body, _err := client.GetSparkTemplateFileContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### [](#)Usage notes
 * You can call this operation to query the directory structure but not application data in the directory. To query the directory structure that contains application data, call the [GetSparkTemplateFullTree](~~612467~~) operation.
 *
 * @param request GetSparkTemplateFolderTreeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetSparkTemplateFolderTreeResponse
 */
func (client *Client) GetSparkTemplateFolderTreeWithOptions(request *GetSparkTemplateFolderTreeRequest, runtime *util.RuntimeOptions) (_result *GetSparkTemplateFolderTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkTemplateFolderTree"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkTemplateFolderTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### [](#)Usage notes
 * You can call this operation to query the directory structure but not application data in the directory. To query the directory structure that contains application data, call the [GetSparkTemplateFullTree](~~612467~~) operation.
 *
 * @param request GetSparkTemplateFolderTreeRequest
 * @return GetSparkTemplateFolderTreeResponse
 */
func (client *Client) GetSparkTemplateFolderTree(request *GetSparkTemplateFolderTreeRequest) (_result *GetSparkTemplateFolderTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkTemplateFolderTreeResponse{}
	_body, _err := client.GetSparkTemplateFolderTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkTemplateFullTreeWithOptions(request *GetSparkTemplateFullTreeRequest, runtime *util.RuntimeOptions) (_result *GetSparkTemplateFullTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkTemplateFullTree"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkTemplateFullTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkTemplateFullTree(request *GetSparkTemplateFullTreeRequest) (_result *GetSparkTemplateFullTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkTemplateFullTreeResponse{}
	_body, _err := client.GetSparkTemplateFullTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetTableRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTableResponse
 */
// Deprecated
func (client *Client) GetTableWithOptions(request *GetTableRequest, runtime *util.RuntimeOptions) (_result *GetTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTable"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetTableRequest
 * @return GetTableResponse
 */
// Deprecated
func (client *Client) GetTable(request *GetTableRequest) (_result *GetTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableColumnsWithOptions(request *GetTableColumnsRequest, runtime *util.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnName)) {
		query["ColumnName"] = request.ColumnName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableColumns"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableColumns(request *GetTableColumnsRequest) (_result *GetTableColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.GetTableColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableDDLWithOptions(request *GetTableDDLRequest, runtime *util.RuntimeOptions) (_result *GetTableDDLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableDDL"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableDDLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableDDL(request *GetTableDDLRequest) (_result *GetTableDDLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableDDLResponse{}
	_body, _err := client.GetTableDDLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableObjectsWithOptions(request *GetTableObjectsRequest, runtime *util.RuntimeOptions) (_result *GetTableObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterDescription)) {
		query["FilterDescription"] = request.FilterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.FilterOwner)) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !tea.BoolValue(util.IsUnset(request.FilterTblName)) {
		query["FilterTblName"] = request.FilterTblName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterTblType)) {
		query["FilterTblType"] = request.FilterTblType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableObjects"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableObjects(request *GetTableObjectsRequest) (_result *GetTableObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableObjectsResponse{}
	_body, _err := client.GetTableObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetViewDDLWithOptions(request *GetViewDDLRequest, runtime *util.RuntimeOptions) (_result *GetViewDDLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.ViewName)) {
		query["ViewName"] = request.ViewName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetViewDDL"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetViewDDLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetViewDDL(request *GetViewDDLRequest) (_result *GetViewDDLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetViewDDLResponse{}
	_body, _err := client.GetViewDDLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetViewObjectsWithOptions(request *GetViewObjectsRequest, runtime *util.RuntimeOptions) (_result *GetViewObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterOwner)) {
		query["FilterOwner"] = request.FilterOwner
	}

	if !tea.BoolValue(util.IsUnset(request.FilterViewName)) {
		query["FilterViewName"] = request.FilterViewName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterViewType)) {
		query["FilterViewType"] = request.FilterViewType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetViewObjects"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetViewObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetViewObjects(request *GetViewObjectsRequest) (_result *GetViewObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetViewObjectsResponse{}
	_body, _err := client.GetViewObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkAppWithOptions(request *KillSparkAppRequest, runtime *util.RuntimeOptions) (_result *KillSparkAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KillSparkApp"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillSparkAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkApp(request *KillSparkAppRequest) (_result *KillSparkAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkAppResponse{}
	_body, _err := client.KillSparkAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkLogAnalyzeTaskWithOptions(request *KillSparkLogAnalyzeTaskRequest, runtime *util.RuntimeOptions) (_result *KillSparkLogAnalyzeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KillSparkLogAnalyzeTask"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkLogAnalyzeTask(request *KillSparkLogAnalyzeTaskRequest) (_result *KillSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkLogAnalyzeTaskResponse{}
	_body, _err := client.KillSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkSQLEngineWithOptions(request *KillSparkSQLEngineRequest, runtime *util.RuntimeOptions) (_result *KillSparkSQLEngineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KillSparkSQLEngine"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillSparkSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkSQLEngine(request *KillSparkSQLEngineRequest) (_result *KillSparkSQLEngineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkSQLEngineResponse{}
	_body, _err := client.KillSparkSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkAppAttemptsWithOptions(request *ListSparkAppAttemptsRequest, runtime *util.RuntimeOptions) (_result *ListSparkAppAttemptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkAppAttempts"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkAppAttemptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkAppAttempts(request *ListSparkAppAttemptsRequest) (_result *ListSparkAppAttemptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkAppAttemptsResponse{}
	_body, _err := client.ListSparkAppAttemptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkAppsWithOptions(request *ListSparkAppsRequest, runtime *util.RuntimeOptions) (_result *ListSparkAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkApps"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkApps(request *ListSparkAppsRequest) (_result *ListSparkAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkAppsResponse{}
	_body, _err := client.ListSparkAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkLogAnalyzeTasksWithOptions(request *ListSparkLogAnalyzeTasksRequest, runtime *util.RuntimeOptions) (_result *ListSparkLogAnalyzeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkLogAnalyzeTasks"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkLogAnalyzeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkLogAnalyzeTasks(request *ListSparkLogAnalyzeTasksRequest) (_result *ListSparkLogAnalyzeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkLogAnalyzeTasksResponse{}
	_body, _err := client.ListSparkLogAnalyzeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkTemplateFileIdsWithOptions(request *ListSparkTemplateFileIdsRequest, runtime *util.RuntimeOptions) (_result *ListSparkTemplateFileIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkTemplateFileIds"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkTemplateFileIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkTemplateFileIds(request *ListSparkTemplateFileIdsRequest) (_result *ListSparkTemplateFileIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkTemplateFileIdsResponse{}
	_body, _err := client.ListSparkTemplateFileIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadSampleDataSetWithOptions(request *LoadSampleDataSetRequest, runtime *util.RuntimeOptions) (_result *LoadSampleDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LoadSampleDataSet"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadSampleDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadSampleDataSet(request *LoadSampleDataSetRequest) (_result *LoadSampleDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoadSampleDataSetResponse{}
	_body, _err := client.LoadSampleDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPrivilegesWithOptions(tmpReq *ModifyAccountPrivilegesRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPrivilegesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAccountPrivilegesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountPrivileges)) {
		request.AccountPrivilegesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountPrivileges, tea.String("AccountPrivileges"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPrivilegesShrink)) {
		query["AccountPrivileges"] = request.AccountPrivilegesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountPrivileges"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountPrivilegesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (_result *ModifyAccountPrivilegesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPrivilegesResponse{}
	_body, _err := client.ModifyAccountPrivilegesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditLogStatus)) {
		query["AuditLogStatus"] = request.AuditLogStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAuditLogConfig"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAuditLogConfig(request *ModifyAuditLogConfigRequest) (_result *ModifyAuditLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.ModifyAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetentionPeriod)) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBackupLog)) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterAccessWhiteListWithOptions(request *ModifyClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayAttribute)) {
		query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayName)) {
		query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		query["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterAccessWhiteList"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterAccessWhiteList(request *ModifyClusterAccessWhiteListRequest) (_result *ModifyClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterAccessWhiteListResponse{}
	_body, _err := client.ModifyClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterConnectionStringWithOptions(request *ModifyClusterConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterConnectionString"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterConnectionString(request *ModifyClusterConnectionStringRequest) (_result *ModifyClusterConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.ModifyClusterConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### [](#)
 * *   During a scaling event, you are not allowed to execute the `SUBMIT JOB` statement to submit asynchronous jobs. If your business requires asynchronous jobs, perform scaling during appropriate periods.
 * *   When cluster specifications are scaled up or down, data in the cluster is migrated for redistribution. The amount of time that is required for data migration is proportional to the volume of data. During a scaling event, the services provided by the cluster are not interrupted. During a scale-down event, data migration can take up to dozens of hours to complete. Proceed with caution especially when your cluster contains a large amount of data.
 * *   If the cluster has a built-in dataset loaded, make sure that the cluster has reserved storage resources of at least 24 AnalyticDB compute units (ACUs). Otherwise, the built-in dataset cannot be used.
 * *   When the scaling process is about to end, your service may encounter transient connections. We recommend that you scale your cluster during off-peak hours or make sure that your application is configured to automatically reconnect to your cluster.
 * *   You can change an AnalyticDB for MySQL cluster from Data Warehouse Edition (V3.0) to Data Lakehouse Edition (V3.0), but not the other way around. For more information, see Change a cluster from Data Warehouse Edition to Data Lakehouse Edition.
 *
 * @param request ModifyDBClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBClusterResponse
 */
func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultResourcePool)) {
		query["EnableDefaultResourcePool"] = request.EnableDefaultResourcePool
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### [](#)
 * *   During a scaling event, you are not allowed to execute the `SUBMIT JOB` statement to submit asynchronous jobs. If your business requires asynchronous jobs, perform scaling during appropriate periods.
 * *   When cluster specifications are scaled up or down, data in the cluster is migrated for redistribution. The amount of time that is required for data migration is proportional to the volume of data. During a scaling event, the services provided by the cluster are not interrupted. During a scale-down event, data migration can take up to dozens of hours to complete. Proceed with caution especially when your cluster contains a large amount of data.
 * *   If the cluster has a built-in dataset loaded, make sure that the cluster has reserved storage resources of at least 24 AnalyticDB compute units (ACUs). Otherwise, the built-in dataset cannot be used.
 * *   When the scaling process is about to end, your service may encounter transient connections. We recommend that you scale your cluster during off-peak hours or make sure that your application is configured to automatically reconnect to your cluster.
 * *   You can change an AnalyticDB for MySQL cluster from Data Warehouse Edition (V3.0) to Data Lakehouse Edition (V3.0), but not the other way around. For more information, see Change a cluster from Data Warehouse Edition to Data Lakehouse Edition.
 *
 * @param request ModifyDBClusterRequest
 * @return ModifyDBClusterResponse
 */
func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainTime)) {
		query["MaintainTime"] = request.MaintainTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterMaintainTime"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBResourceGroupWithOptions(tmpReq *ModifyDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDBResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rules)) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, tea.String("Rules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterMode)) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterSizeResource)) {
		query["ClusterSizeResource"] = request.ClusterSizeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSpot)) {
		query["EnableSpot"] = request.EnableSpot
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxClusterCount)) {
		query["MaxClusterCount"] = request.MaxClusterCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxComputeResource)) {
		query["MaxComputeResource"] = request.MaxComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.MinClusterCount)) {
		query["MinClusterCount"] = request.MinClusterCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinComputeResource)) {
		query["MinComputeResource"] = request.MinComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RulesShrink)) {
		query["Rules"] = request.RulesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBResourceGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBResourceGroup(request *ModifyDBResourceGroupRequest) (_result *ModifyDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.ModifyDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticPlanWithOptions(request *ModifyElasticPlanRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CronExpression)) {
		query["CronExpression"] = request.CronExpression
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSize)) {
		query["TargetSize"] = request.TargetSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyElasticPlan"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticPlan(request *ModifyElasticPlanRequest) (_result *ModifyElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.ModifyElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreloadSparkAppMetricsWithOptions(request *PreloadSparkAppMetricsRequest, runtime *util.RuntimeOptions) (_result *PreloadSparkAppMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PreloadSparkAppMetrics"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreloadSparkAppMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreloadSparkAppMetrics(request *PreloadSparkAppMetricsRequest) (_result *PreloadSparkAppMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreloadSparkAppMetricsResponse{}
	_body, _err := client.PreloadSparkAppMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnectionWithOptions(request *ReleaseClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseClusterPublicConnection"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnection(request *ReleaseClusterPublicConnectionRequest) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.ReleaseClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameSparkTemplateFileWithOptions(request *RenameSparkTemplateFileRequest, runtime *util.RuntimeOptions) (_result *RenameSparkTemplateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameSparkTemplateFile"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameSparkTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameSparkTemplateFile(request *RenameSparkTemplateFileRequest) (_result *RenameSparkTemplateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameSparkTemplateFileResponse{}
	_body, _err := client.RenameSparkTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSparkAppLogRootPathWithOptions(request *SetSparkAppLogRootPathRequest, runtime *util.RuntimeOptions) (_result *SetSparkAppLogRootPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OssLogPath)) {
		body["OssLogPath"] = request.OssLogPath
	}

	if !tea.BoolValue(util.IsUnset(request.UseDefaultOss)) {
		body["UseDefaultOss"] = request.UseDefaultOss
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSparkAppLogRootPath"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSparkAppLogRootPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSparkAppLogRootPath(request *SetSparkAppLogRootPathRequest) (_result *SetSparkAppLogRootPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSparkAppLogRootPathResponse{}
	_body, _err := client.SetSparkAppLogRootPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSparkSQLEngineWithOptions(request *StartSparkSQLEngineRequest, runtime *util.RuntimeOptions) (_result *StartSparkSQLEngineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Jars)) {
		body["Jars"] = request.Jars
	}

	if !tea.BoolValue(util.IsUnset(request.MaxExecutor)) {
		body["MaxExecutor"] = request.MaxExecutor
	}

	if !tea.BoolValue(util.IsUnset(request.MinExecutor)) {
		body["MinExecutor"] = request.MinExecutor
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SlotNum)) {
		body["SlotNum"] = request.SlotNum
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSparkSQLEngine"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSparkSQLEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSparkSQLEngine(request *StartSparkSQLEngineRequest) (_result *StartSparkSQLEngineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSparkSQLEngineResponse{}
	_body, _err := client.StartSparkSQLEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkAppWithOptions(request *SubmitSparkAppRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentSource)) {
		body["AgentSource"] = request.AgentSource
	}

	if !tea.BoolValue(util.IsUnset(request.AgentVersion)) {
		body["AgentVersion"] = request.AgentVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateFileId)) {
		body["TemplateFileId"] = request.TemplateFileId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSparkApp"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSparkAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkApp(request *SubmitSparkAppRequest) (_result *SubmitSparkAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkAppResponse{}
	_body, _err := client.SubmitSparkAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkLogAnalyzeTaskWithOptions(request *SubmitSparkLogAnalyzeTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkLogAnalyzeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSparkLogAnalyzeTask"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSparkLogAnalyzeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkLogAnalyzeTask(request *SubmitSparkLogAnalyzeTaskRequest) (_result *SubmitSparkLogAnalyzeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkLogAnalyzeTaskResponse{}
	_body, _err := client.SubmitSparkLogAnalyzeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindAccountWithOptions(request *UnbindAccountRequest, runtime *util.RuntimeOptions) (_result *UnbindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindAccount"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAccount(request *UnbindAccountRequest) (_result *UnbindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAccountResponse{}
	_body, _err := client.UnbindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDBResourceGroupWithUserWithOptions(request *UnbindDBResourceGroupWithUserRequest, runtime *util.RuntimeOptions) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUser)) {
		query["GroupUser"] = request.GroupUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDBResourceGroupWithUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDBResourceGroupWithUser(request *UnbindDBResourceGroupWithUserRequest) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.UnbindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSparkTemplateFileWithOptions(request *UpdateSparkTemplateFileRequest, runtime *util.RuntimeOptions) (_result *UpdateSparkTemplateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		body["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		body["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSparkTemplateFile"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSparkTemplateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSparkTemplateFile(request *UpdateSparkTemplateFileRequest) (_result *UpdateSparkTemplateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSparkTemplateFileResponse{}
	_body, _err := client.UpdateSparkTemplateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
