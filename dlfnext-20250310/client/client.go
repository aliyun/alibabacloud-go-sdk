// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Catalog struct {
	CreatedAt *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy *string            `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id        *string            `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options   map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Owner     *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	Status    *string            `json:"status,omitempty" xml:"status,omitempty"`
	Type      *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt *int64             `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy *string            `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Catalog) String() string {
	return tea.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) SetCreatedAt(v int64) *Catalog {
	s.CreatedAt = &v
	return s
}

func (s *Catalog) SetCreatedBy(v string) *Catalog {
	s.CreatedBy = &v
	return s
}

func (s *Catalog) SetId(v string) *Catalog {
	s.Id = &v
	return s
}

func (s *Catalog) SetName(v string) *Catalog {
	s.Name = &v
	return s
}

func (s *Catalog) SetOptions(v map[string]*string) *Catalog {
	s.Options = v
	return s
}

func (s *Catalog) SetOwner(v string) *Catalog {
	s.Owner = &v
	return s
}

func (s *Catalog) SetStatus(v string) *Catalog {
	s.Status = &v
	return s
}

func (s *Catalog) SetType(v string) *Catalog {
	s.Type = &v
	return s
}

func (s *Catalog) SetUpdatedAt(v int64) *Catalog {
	s.UpdatedAt = &v
	return s
}

func (s *Catalog) SetUpdatedBy(v string) *Catalog {
	s.UpdatedBy = &v
	return s
}

type CatalogSummary struct {
	ApiVisitCountMonthly *int64     `json:"apiVisitCountMonthly,omitempty" xml:"apiVisitCountMonthly,omitempty"`
	DatabaseCount        *MoMValues `json:"databaseCount,omitempty" xml:"databaseCount,omitempty"`
	// Update date of the statistics
	GeneratedDate        *string    `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	PartitionCount       *MoMValues `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	TableCount           *MoMValues `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	ThroughputMonthly    *int64     `json:"throughputMonthly,omitempty" xml:"throughputMonthly,omitempty"`
	TotalFileCount       *MoMValues `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	TotalFileSizeInBytes *MoMValues `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
}

func (s CatalogSummary) String() string {
	return tea.Prettify(s)
}

func (s CatalogSummary) GoString() string {
	return s.String()
}

func (s *CatalogSummary) SetApiVisitCountMonthly(v int64) *CatalogSummary {
	s.ApiVisitCountMonthly = &v
	return s
}

func (s *CatalogSummary) SetDatabaseCount(v *MoMValues) *CatalogSummary {
	s.DatabaseCount = v
	return s
}

func (s *CatalogSummary) SetGeneratedDate(v string) *CatalogSummary {
	s.GeneratedDate = &v
	return s
}

func (s *CatalogSummary) SetPartitionCount(v *MoMValues) *CatalogSummary {
	s.PartitionCount = v
	return s
}

func (s *CatalogSummary) SetTableCount(v *MoMValues) *CatalogSummary {
	s.TableCount = v
	return s
}

func (s *CatalogSummary) SetThroughputMonthly(v int64) *CatalogSummary {
	s.ThroughputMonthly = &v
	return s
}

func (s *CatalogSummary) SetTotalFileCount(v *MoMValues) *CatalogSummary {
	s.TotalFileCount = v
	return s
}

func (s *CatalogSummary) SetTotalFileSizeInBytes(v *MoMValues) *CatalogSummary {
	s.TotalFileSizeInBytes = v
	return s
}

type CatalogSummaryTrend struct {
	// API visit count trends
	ApiVisitCount []*DateSummary `json:"apiVisitCount,omitempty" xml:"apiVisitCount,omitempty" type:"Repeated"`
	// Table count trends
	Throughput []*DateSummary `json:"throughput,omitempty" xml:"throughput,omitempty" type:"Repeated"`
	// Historical total file count
	TotalFileCount []*DateSummary `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty" type:"Repeated"`
	// Database count trends
	TotalFileSizeInBytes []*DateSummary `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty" type:"Repeated"`
	// Latest snapshot file count
	TotalMetaCount []*DateSummary `json:"totalMetaCount,omitempty" xml:"totalMetaCount,omitempty" type:"Repeated"`
}

func (s CatalogSummaryTrend) String() string {
	return tea.Prettify(s)
}

func (s CatalogSummaryTrend) GoString() string {
	return s.String()
}

func (s *CatalogSummaryTrend) SetApiVisitCount(v []*DateSummary) *CatalogSummaryTrend {
	s.ApiVisitCount = v
	return s
}

func (s *CatalogSummaryTrend) SetThroughput(v []*DateSummary) *CatalogSummaryTrend {
	s.Throughput = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalFileCount(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalFileCount = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalFileSizeInBytes(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalFileSizeInBytes = v
	return s
}

func (s *CatalogSummaryTrend) SetTotalMetaCount(v []*DateSummary) *CatalogSummaryTrend {
	s.TotalMetaCount = v
	return s
}

type DataField struct {
	Description *string       `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int32        `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string       `json:"name,omitempty" xml:"name,omitempty"`
	Type        *FullDataType `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DataField) String() string {
	return tea.Prettify(s)
}

func (s DataField) GoString() string {
	return s.String()
}

func (s *DataField) SetDescription(v string) *DataField {
	s.Description = &v
	return s
}

func (s *DataField) SetId(v int32) *DataField {
	s.Id = &v
	return s
}

func (s *DataField) SetName(v string) *DataField {
	s.Name = &v
	return s
}

func (s *DataField) SetType(v *FullDataType) *DataField {
	s.Type = v
	return s
}

type Database struct {
	CreatedAt *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy *string            `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id        *string            `json:"id,omitempty" xml:"id,omitempty"`
	Location  *string            `json:"location,omitempty" xml:"location,omitempty"`
	Name      *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options   map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Owner     *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt *int64             `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy *string            `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Database) String() string {
	return tea.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) SetCreatedAt(v int64) *Database {
	s.CreatedAt = &v
	return s
}

func (s *Database) SetCreatedBy(v string) *Database {
	s.CreatedBy = &v
	return s
}

func (s *Database) SetId(v string) *Database {
	s.Id = &v
	return s
}

func (s *Database) SetLocation(v string) *Database {
	s.Location = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetOptions(v map[string]*string) *Database {
	s.Options = v
	return s
}

func (s *Database) SetOwner(v string) *Database {
	s.Owner = &v
	return s
}

func (s *Database) SetUpdatedAt(v int64) *Database {
	s.UpdatedAt = &v
	return s
}

func (s *Database) SetUpdatedBy(v string) *Database {
	s.UpdatedBy = &v
	return s
}

type DatabaseSummary struct {
	// Creation timestamp in milliseconds
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 库名 - Database name
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// Last profile update date in format yyyyMMdd
	GeneratedDate *string `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	// Storage location URI
	Location       *string `json:"location,omitempty" xml:"location,omitempty"`
	PartitionCount *int64  `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// Total storage in bytes
	TableCount     *int64 `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	TotalFileCount *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	// Total file count
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
}

func (s DatabaseSummary) String() string {
	return tea.Prettify(s)
}

func (s DatabaseSummary) GoString() string {
	return s.String()
}

func (s *DatabaseSummary) SetCreatedAt(v int64) *DatabaseSummary {
	s.CreatedAt = &v
	return s
}

func (s *DatabaseSummary) SetDatabaseName(v string) *DatabaseSummary {
	s.DatabaseName = &v
	return s
}

func (s *DatabaseSummary) SetGeneratedDate(v string) *DatabaseSummary {
	s.GeneratedDate = &v
	return s
}

func (s *DatabaseSummary) SetLocation(v string) *DatabaseSummary {
	s.Location = &v
	return s
}

func (s *DatabaseSummary) SetPartitionCount(v int64) *DatabaseSummary {
	s.PartitionCount = &v
	return s
}

func (s *DatabaseSummary) SetTableCount(v int64) *DatabaseSummary {
	s.TableCount = &v
	return s
}

func (s *DatabaseSummary) SetTotalFileCount(v int64) *DatabaseSummary {
	s.TotalFileCount = &v
	return s
}

func (s *DatabaseSummary) SetTotalFileSizeInBytes(v int64) *DatabaseSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

type DateSummary struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// Metric value at corresponding date
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DateSummary) String() string {
	return tea.Prettify(s)
}

func (s DateSummary) GoString() string {
	return s.String()
}

func (s *DateSummary) SetDate(v string) *DateSummary {
	s.Date = &v
	return s
}

func (s *DateSummary) SetValue(v int64) *DateSummary {
	s.Value = &v
	return s
}

type ErrorResponse struct {
	Code         *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Message      *string `json:"message,omitempty" xml:"message,omitempty"`
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s ErrorResponse) GoString() string {
	return s.String()
}

func (s *ErrorResponse) SetCode(v int32) *ErrorResponse {
	s.Code = &v
	return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
	s.Message = &v
	return s
}

func (s *ErrorResponse) SetResourceName(v string) *ErrorResponse {
	s.ResourceName = &v
	return s
}

func (s *ErrorResponse) SetResourceType(v string) *ErrorResponse {
	s.ResourceType = &v
	return s
}

type FailurePermission struct {
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Permission   *Permission `json:"permission,omitempty" xml:"permission,omitempty"`
}

func (s FailurePermission) String() string {
	return tea.Prettify(s)
}

func (s FailurePermission) GoString() string {
	return s.String()
}

func (s *FailurePermission) SetErrorCode(v string) *FailurePermission {
	s.ErrorCode = &v
	return s
}

func (s *FailurePermission) SetErrorMessage(v string) *FailurePermission {
	s.ErrorMessage = &v
	return s
}

func (s *FailurePermission) SetPermission(v *Permission) *FailurePermission {
	s.Permission = v
	return s
}

type FullDataType struct {
	Element *FullDataType `json:"element,omitempty" xml:"element,omitempty"`
	Fields  []*DataField  `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Key     *FullDataType `json:"key,omitempty" xml:"key,omitempty"`
	Type    *string       `json:"type,omitempty" xml:"type,omitempty"`
	Value   *FullDataType `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FullDataType) String() string {
	return tea.Prettify(s)
}

func (s FullDataType) GoString() string {
	return s.String()
}

func (s *FullDataType) SetElement(v *FullDataType) *FullDataType {
	s.Element = v
	return s
}

func (s *FullDataType) SetFields(v []*DataField) *FullDataType {
	s.Fields = v
	return s
}

func (s *FullDataType) SetKey(v *FullDataType) *FullDataType {
	s.Key = v
	return s
}

func (s *FullDataType) SetType(v string) *FullDataType {
	s.Type = &v
	return s
}

func (s *FullDataType) SetValue(v *FullDataType) *FullDataType {
	s.Value = v
	return s
}

type FullInstant struct {
	SnapshotId *int64  `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	TagName    *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FullInstant) String() string {
	return tea.Prettify(s)
}

func (s FullInstant) GoString() string {
	return s.String()
}

func (s *FullInstant) SetSnapshotId(v int64) *FullInstant {
	s.SnapshotId = &v
	return s
}

func (s *FullInstant) SetTagName(v string) *FullInstant {
	s.TagName = &v
	return s
}

func (s *FullInstant) SetType(v string) *FullInstant {
	s.Type = &v
	return s
}

type FullSchemaChange struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// required in UpdateComment/AddColumn
	Comment  *string       `json:"comment,omitempty" xml:"comment,omitempty"`
	DataType *FullDataType `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// required in AddColumn/RenameColumn/DropColumn/UpdateColumnComment/UpdateColumnType/UpdateColumnNullability
	FieldNames []*string `json:"fieldNames,omitempty" xml:"fieldNames,omitempty" type:"Repeated"`
	// required in UpdateColumnType
	KeepNullability *bool `json:"keepNullability,omitempty" xml:"keepNullability,omitempty"`
	// required in SetOption/RemoveOption
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Move *Move   `json:"move,omitempty" xml:"move,omitempty"`
	// required in UpdateColumnComment
	NewComment  *string       `json:"newComment,omitempty" xml:"newComment,omitempty"`
	NewDataType *FullDataType `json:"newDataType,omitempty" xml:"newDataType,omitempty"`
	// required in RenameColumn
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// required in UpdateColumnNullability
	NewNullability *bool `json:"newNullability,omitempty" xml:"newNullability,omitempty"`
	// required in SetOption
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FullSchemaChange) String() string {
	return tea.Prettify(s)
}

func (s FullSchemaChange) GoString() string {
	return s.String()
}

func (s *FullSchemaChange) SetAction(v string) *FullSchemaChange {
	s.Action = &v
	return s
}

func (s *FullSchemaChange) SetComment(v string) *FullSchemaChange {
	s.Comment = &v
	return s
}

func (s *FullSchemaChange) SetDataType(v *FullDataType) *FullSchemaChange {
	s.DataType = v
	return s
}

func (s *FullSchemaChange) SetFieldNames(v []*string) *FullSchemaChange {
	s.FieldNames = v
	return s
}

func (s *FullSchemaChange) SetKeepNullability(v bool) *FullSchemaChange {
	s.KeepNullability = &v
	return s
}

func (s *FullSchemaChange) SetKey(v string) *FullSchemaChange {
	s.Key = &v
	return s
}

func (s *FullSchemaChange) SetMove(v *Move) *FullSchemaChange {
	s.Move = v
	return s
}

func (s *FullSchemaChange) SetNewComment(v string) *FullSchemaChange {
	s.NewComment = &v
	return s
}

func (s *FullSchemaChange) SetNewDataType(v *FullDataType) *FullSchemaChange {
	s.NewDataType = v
	return s
}

func (s *FullSchemaChange) SetNewName(v string) *FullSchemaChange {
	s.NewName = &v
	return s
}

func (s *FullSchemaChange) SetNewNullability(v bool) *FullSchemaChange {
	s.NewNullability = &v
	return s
}

func (s *FullSchemaChange) SetValue(v string) *FullSchemaChange {
	s.Value = &v
	return s
}

type IcebergNestedField struct {
	Doc      *string `json:"doc,omitempty" xml:"doc,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Optional *bool   `json:"optional,omitempty" xml:"optional,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IcebergNestedField) String() string {
	return tea.Prettify(s)
}

func (s IcebergNestedField) GoString() string {
	return s.String()
}

func (s *IcebergNestedField) SetDoc(v string) *IcebergNestedField {
	s.Doc = &v
	return s
}

func (s *IcebergNestedField) SetId(v int64) *IcebergNestedField {
	s.Id = &v
	return s
}

func (s *IcebergNestedField) SetName(v string) *IcebergNestedField {
	s.Name = &v
	return s
}

func (s *IcebergNestedField) SetOptional(v bool) *IcebergNestedField {
	s.Optional = &v
	return s
}

func (s *IcebergNestedField) SetType(v string) *IcebergNestedField {
	s.Type = &v
	return s
}

type IcebergPartitionField struct {
	FieldId   *int64  `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceId  *int64  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Transform *string `json:"transform,omitempty" xml:"transform,omitempty"`
}

func (s IcebergPartitionField) String() string {
	return tea.Prettify(s)
}

func (s IcebergPartitionField) GoString() string {
	return s.String()
}

func (s *IcebergPartitionField) SetFieldId(v int64) *IcebergPartitionField {
	s.FieldId = &v
	return s
}

func (s *IcebergPartitionField) SetName(v string) *IcebergPartitionField {
	s.Name = &v
	return s
}

func (s *IcebergPartitionField) SetSourceId(v int64) *IcebergPartitionField {
	s.SourceId = &v
	return s
}

func (s *IcebergPartitionField) SetTransform(v string) *IcebergPartitionField {
	s.Transform = &v
	return s
}

type IcebergSnapshot struct {
	AddedRows       *int64             `json:"addedRows,omitempty" xml:"addedRows,omitempty"`
	Id              *int64             `json:"id,omitempty" xml:"id,omitempty"`
	Operation       *string            `json:"operation,omitempty" xml:"operation,omitempty"`
	ParentId        *int64             `json:"parentId,omitempty" xml:"parentId,omitempty"`
	SchemaId        *int64             `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	SequenceNumber  *int64             `json:"sequenceNumber,omitempty" xml:"sequenceNumber,omitempty"`
	Summary         map[string]*string `json:"summary,omitempty" xml:"summary,omitempty"`
	TimestampMillis *int64             `json:"timestampMillis,omitempty" xml:"timestampMillis,omitempty"`
}

func (s IcebergSnapshot) String() string {
	return tea.Prettify(s)
}

func (s IcebergSnapshot) GoString() string {
	return s.String()
}

func (s *IcebergSnapshot) SetAddedRows(v int64) *IcebergSnapshot {
	s.AddedRows = &v
	return s
}

func (s *IcebergSnapshot) SetId(v int64) *IcebergSnapshot {
	s.Id = &v
	return s
}

func (s *IcebergSnapshot) SetOperation(v string) *IcebergSnapshot {
	s.Operation = &v
	return s
}

func (s *IcebergSnapshot) SetParentId(v int64) *IcebergSnapshot {
	s.ParentId = &v
	return s
}

func (s *IcebergSnapshot) SetSchemaId(v int64) *IcebergSnapshot {
	s.SchemaId = &v
	return s
}

func (s *IcebergSnapshot) SetSequenceNumber(v int64) *IcebergSnapshot {
	s.SequenceNumber = &v
	return s
}

func (s *IcebergSnapshot) SetSummary(v map[string]*string) *IcebergSnapshot {
	s.Summary = v
	return s
}

func (s *IcebergSnapshot) SetTimestampMillis(v int64) *IcebergSnapshot {
	s.TimestampMillis = &v
	return s
}

type IcebergTable struct {
	CreatedAt            *int64                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy            *string               `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	IcebergTableMetadata *IcebergTableMetadata `json:"icebergTableMetadata,omitempty" xml:"icebergTableMetadata,omitempty"`
	Id                   *string               `json:"id,omitempty" xml:"id,omitempty"`
	Name                 *string               `json:"name,omitempty" xml:"name,omitempty"`
	Owner                *string               `json:"owner,omitempty" xml:"owner,omitempty"`
	Path                 *string               `json:"path,omitempty" xml:"path,omitempty"`
	UpdatedAt            *int64                `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy            *string               `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	Version              *int64                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s IcebergTable) String() string {
	return tea.Prettify(s)
}

func (s IcebergTable) GoString() string {
	return s.String()
}

func (s *IcebergTable) SetCreatedAt(v int64) *IcebergTable {
	s.CreatedAt = &v
	return s
}

func (s *IcebergTable) SetCreatedBy(v string) *IcebergTable {
	s.CreatedBy = &v
	return s
}

func (s *IcebergTable) SetIcebergTableMetadata(v *IcebergTableMetadata) *IcebergTable {
	s.IcebergTableMetadata = v
	return s
}

func (s *IcebergTable) SetId(v string) *IcebergTable {
	s.Id = &v
	return s
}

func (s *IcebergTable) SetName(v string) *IcebergTable {
	s.Name = &v
	return s
}

func (s *IcebergTable) SetOwner(v string) *IcebergTable {
	s.Owner = &v
	return s
}

func (s *IcebergTable) SetPath(v string) *IcebergTable {
	s.Path = &v
	return s
}

func (s *IcebergTable) SetUpdatedAt(v int64) *IcebergTable {
	s.UpdatedAt = &v
	return s
}

func (s *IcebergTable) SetUpdatedBy(v string) *IcebergTable {
	s.UpdatedBy = &v
	return s
}

func (s *IcebergTable) SetVersion(v int64) *IcebergTable {
	s.Version = &v
	return s
}

type IcebergTableMetadata struct {
	CurrentSnapshot *IcebergSnapshot         `json:"currentSnapshot,omitempty" xml:"currentSnapshot,omitempty"`
	Fields          []*IcebergNestedField    `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	PartitionFields []*IcebergPartitionField `json:"partitionFields,omitempty" xml:"partitionFields,omitempty" type:"Repeated"`
	Properties      map[string]*string       `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s IcebergTableMetadata) String() string {
	return tea.Prettify(s)
}

func (s IcebergTableMetadata) GoString() string {
	return s.String()
}

func (s *IcebergTableMetadata) SetCurrentSnapshot(v *IcebergSnapshot) *IcebergTableMetadata {
	s.CurrentSnapshot = v
	return s
}

func (s *IcebergTableMetadata) SetFields(v []*IcebergNestedField) *IcebergTableMetadata {
	s.Fields = v
	return s
}

func (s *IcebergTableMetadata) SetPartitionFields(v []*IcebergPartitionField) *IcebergTableMetadata {
	s.PartitionFields = v
	return s
}

func (s *IcebergTableMetadata) SetProperties(v map[string]*string) *IcebergTableMetadata {
	s.Properties = v
	return s
}

type Identifier struct {
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	Object   *string `json:"object,omitempty" xml:"object,omitempty"`
}

func (s Identifier) String() string {
	return tea.Prettify(s)
}

func (s Identifier) GoString() string {
	return s.String()
}

func (s *Identifier) SetDatabase(v string) *Identifier {
	s.Database = &v
	return s
}

func (s *Identifier) SetObject(v string) *Identifier {
	s.Object = &v
	return s
}

type MoMValues struct {
	// total
	CurrentValue *int64 `json:"currentValue,omitempty" xml:"currentValue,omitempty"`
	// daily addition
	LastDayValue *int64 `json:"lastDayValue,omitempty" xml:"lastDayValue,omitempty"`
	// monthly addition
	LastMonthValue *int64 `json:"lastMonthValue,omitempty" xml:"lastMonthValue,omitempty"`
}

func (s MoMValues) String() string {
	return tea.Prettify(s)
}

func (s MoMValues) GoString() string {
	return s.String()
}

func (s *MoMValues) SetCurrentValue(v int64) *MoMValues {
	s.CurrentValue = &v
	return s
}

func (s *MoMValues) SetLastDayValue(v int64) *MoMValues {
	s.LastDayValue = &v
	return s
}

func (s *MoMValues) SetLastMonthValue(v int64) *MoMValues {
	s.LastMonthValue = &v
	return s
}

type Move struct {
	FieldName          *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	ReferenceFieldName *string `json:"referenceFieldName,omitempty" xml:"referenceFieldName,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Move) String() string {
	return tea.Prettify(s)
}

func (s Move) GoString() string {
	return s.String()
}

func (s *Move) SetFieldName(v string) *Move {
	s.FieldName = &v
	return s
}

func (s *Move) SetReferenceFieldName(v string) *Move {
	s.ReferenceFieldName = &v
	return s
}

func (s *Move) SetType(v string) *Move {
	s.Type = &v
	return s
}

type Namespace struct {
	CreatedAt *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy *string            `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id        *string            `json:"id,omitempty" xml:"id,omitempty"`
	Location  *string            `json:"location,omitempty" xml:"location,omitempty"`
	Name      *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options   map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Owner     *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt *int64             `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy *string            `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Namespace) String() string {
	return tea.Prettify(s)
}

func (s Namespace) GoString() string {
	return s.String()
}

func (s *Namespace) SetCreatedAt(v int64) *Namespace {
	s.CreatedAt = &v
	return s
}

func (s *Namespace) SetCreatedBy(v string) *Namespace {
	s.CreatedBy = &v
	return s
}

func (s *Namespace) SetId(v string) *Namespace {
	s.Id = &v
	return s
}

func (s *Namespace) SetLocation(v string) *Namespace {
	s.Location = &v
	return s
}

func (s *Namespace) SetName(v string) *Namespace {
	s.Name = &v
	return s
}

func (s *Namespace) SetOptions(v map[string]*string) *Namespace {
	s.Options = v
	return s
}

func (s *Namespace) SetOwner(v string) *Namespace {
	s.Owner = &v
	return s
}

func (s *Namespace) SetUpdatedAt(v int64) *Namespace {
	s.UpdatedAt = &v
	return s
}

func (s *Namespace) SetUpdatedBy(v string) *Namespace {
	s.UpdatedBy = &v
	return s
}

type Partition struct {
	CreatedAt            *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy            *string                `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Done                 *bool                  `json:"done,omitempty" xml:"done,omitempty"`
	FileCount            *int64                 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	FileSizeInBytes      *int64                 `json:"fileSizeInBytes,omitempty" xml:"fileSizeInBytes,omitempty"`
	LastFileCreationTime *int64                 `json:"lastFileCreationTime,omitempty" xml:"lastFileCreationTime,omitempty"`
	RecordCount          *int64                 `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	Spec                 map[string]interface{} `json:"spec,omitempty" xml:"spec,omitempty"`
	UpdatedAt            *int64                 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy            *string                `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Partition) String() string {
	return tea.Prettify(s)
}

func (s Partition) GoString() string {
	return s.String()
}

func (s *Partition) SetCreatedAt(v int64) *Partition {
	s.CreatedAt = &v
	return s
}

func (s *Partition) SetCreatedBy(v string) *Partition {
	s.CreatedBy = &v
	return s
}

func (s *Partition) SetDone(v bool) *Partition {
	s.Done = &v
	return s
}

func (s *Partition) SetFileCount(v int64) *Partition {
	s.FileCount = &v
	return s
}

func (s *Partition) SetFileSizeInBytes(v int64) *Partition {
	s.FileSizeInBytes = &v
	return s
}

func (s *Partition) SetLastFileCreationTime(v int64) *Partition {
	s.LastFileCreationTime = &v
	return s
}

func (s *Partition) SetRecordCount(v int64) *Partition {
	s.RecordCount = &v
	return s
}

func (s *Partition) SetSpec(v map[string]interface{}) *Partition {
	s.Spec = v
	return s
}

func (s *Partition) SetUpdatedAt(v int64) *Partition {
	s.UpdatedAt = &v
	return s
}

func (s *Partition) SetUpdatedBy(v string) *Partition {
	s.UpdatedBy = &v
	return s
}

type PartitionSummaries struct {
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	// Current page of partition profiles
	Partitions []*PartitionSummary `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s PartitionSummaries) String() string {
	return tea.Prettify(s)
}

func (s PartitionSummaries) GoString() string {
	return s.String()
}

func (s *PartitionSummaries) SetNextPageToken(v string) *PartitionSummaries {
	s.NextPageToken = &v
	return s
}

func (s *PartitionSummaries) SetPartitions(v []*PartitionSummary) *PartitionSummaries {
	s.Partitions = v
	return s
}

type PartitionSummary struct {
	// Partition creation timestamp in milliseconds
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Database name
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// Total files in partition
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// Partition identifier
	PartitionName *string `json:"partitionName,omitempty" xml:"partitionName,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// 24h access count
	TotalFileCount *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	// Last data access timestamp in milliseconds
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
	UpdatedAt            *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s PartitionSummary) String() string {
	return tea.Prettify(s)
}

func (s PartitionSummary) GoString() string {
	return s.String()
}

func (s *PartitionSummary) SetCreatedAt(v int64) *PartitionSummary {
	s.CreatedAt = &v
	return s
}

func (s *PartitionSummary) SetDatabaseName(v string) *PartitionSummary {
	s.DatabaseName = &v
	return s
}

func (s *PartitionSummary) SetLastAccessTime(v int64) *PartitionSummary {
	s.LastAccessTime = &v
	return s
}

func (s *PartitionSummary) SetPartitionName(v string) *PartitionSummary {
	s.PartitionName = &v
	return s
}

func (s *PartitionSummary) SetTableName(v string) *PartitionSummary {
	s.TableName = &v
	return s
}

func (s *PartitionSummary) SetTotalFileCount(v int64) *PartitionSummary {
	s.TotalFileCount = &v
	return s
}

func (s *PartitionSummary) SetTotalFileSizeInBytes(v int64) *PartitionSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

func (s *PartitionSummary) SetUpdatedAt(v int64) *PartitionSummary {
	s.UpdatedAt = &v
	return s
}

type Permission struct {
	Access       *string `json:"access,omitempty" xml:"access,omitempty"`
	Database     *string `json:"database,omitempty" xml:"database,omitempty"`
	Principal    *string `json:"principal,omitempty" xml:"principal,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Table        *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s Permission) String() string {
	return tea.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) SetAccess(v string) *Permission {
	s.Access = &v
	return s
}

func (s *Permission) SetDatabase(v string) *Permission {
	s.Database = &v
	return s
}

func (s *Permission) SetPrincipal(v string) *Permission {
	s.Principal = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

func (s *Permission) SetTable(v string) *Permission {
	s.Table = &v
	return s
}

type Role struct {
	CreatedAt     *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy     *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName   *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	IsPredefined  *string `json:"isPredefined,omitempty" xml:"isPredefined,omitempty"`
	RoleName      *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UpdatedAt     *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy     *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	Users         []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s Role) String() string {
	return tea.Prettify(s)
}

func (s Role) GoString() string {
	return s.String()
}

func (s *Role) SetCreatedAt(v int64) *Role {
	s.CreatedAt = &v
	return s
}

func (s *Role) SetCreatedBy(v string) *Role {
	s.CreatedBy = &v
	return s
}

func (s *Role) SetDescription(v string) *Role {
	s.Description = &v
	return s
}

func (s *Role) SetDisplayName(v string) *Role {
	s.DisplayName = &v
	return s
}

func (s *Role) SetIsPredefined(v string) *Role {
	s.IsPredefined = &v
	return s
}

func (s *Role) SetRoleName(v string) *Role {
	s.RoleName = &v
	return s
}

func (s *Role) SetRolePrincipal(v string) *Role {
	s.RolePrincipal = &v
	return s
}

func (s *Role) SetUpdatedAt(v int64) *Role {
	s.UpdatedAt = &v
	return s
}

func (s *Role) SetUpdatedBy(v string) *Role {
	s.UpdatedBy = &v
	return s
}

func (s *Role) SetUsers(v []*User) *Role {
	s.Users = v
	return s
}

type Schema struct {
	Comment       *string            `json:"comment,omitempty" xml:"comment,omitempty"`
	Fields        []*DataField       `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Options       map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	PartitionKeys []*string          `json:"partitionKeys,omitempty" xml:"partitionKeys,omitempty" type:"Repeated"`
	PrimaryKeys   []*string          `json:"primaryKeys,omitempty" xml:"primaryKeys,omitempty" type:"Repeated"`
}

func (s Schema) String() string {
	return tea.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) SetComment(v string) *Schema {
	s.Comment = &v
	return s
}

func (s *Schema) SetFields(v []*DataField) *Schema {
	s.Fields = v
	return s
}

func (s *Schema) SetOptions(v map[string]*string) *Schema {
	s.Options = v
	return s
}

func (s *Schema) SetPartitionKeys(v []*string) *Schema {
	s.PartitionKeys = v
	return s
}

func (s *Schema) SetPrimaryKeys(v []*string) *Schema {
	s.PrimaryKeys = v
	return s
}

type Snapshot struct {
	BaseManifestList *string `json:"baseManifestList,omitempty" xml:"baseManifestList,omitempty"`
	// if can be null:
	// true
	ChangelogManifestList *string           `json:"changelogManifestList,omitempty" xml:"changelogManifestList,omitempty"`
	ChangelogRecordCount  *int64            `json:"changelogRecordCount,omitempty" xml:"changelogRecordCount,omitempty"`
	CommitIdentifier      *int64            `json:"commitIdentifier,omitempty" xml:"commitIdentifier,omitempty"`
	CommitKind            *string           `json:"commitKind,omitempty" xml:"commitKind,omitempty"`
	CommitUser            *string           `json:"commitUser,omitempty" xml:"commitUser,omitempty"`
	DeltaManifestList     *string           `json:"deltaManifestList,omitempty" xml:"deltaManifestList,omitempty"`
	DeltaRecordCount      *int64            `json:"deltaRecordCount,omitempty" xml:"deltaRecordCount,omitempty"`
	Id                    *int64            `json:"id,omitempty" xml:"id,omitempty"`
	IndexManifest         *string           `json:"indexManifest,omitempty" xml:"indexManifest,omitempty"`
	LogOffsets            map[string]*int64 `json:"logOffsets,omitempty" xml:"logOffsets,omitempty"`
	SchemaId              *int64            `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	Statistics            *string           `json:"statistics,omitempty" xml:"statistics,omitempty"`
	TimeMillis            *int64            `json:"timeMillis,omitempty" xml:"timeMillis,omitempty"`
	TotalRecordCount      *int64            `json:"totalRecordCount,omitempty" xml:"totalRecordCount,omitempty"`
	// if can be null:
	// true
	Version   *int32 `json:"version,omitempty" xml:"version,omitempty"`
	Watermark *int64 `json:"watermark,omitempty" xml:"watermark,omitempty"`
}

func (s Snapshot) String() string {
	return tea.Prettify(s)
}

func (s Snapshot) GoString() string {
	return s.String()
}

func (s *Snapshot) SetBaseManifestList(v string) *Snapshot {
	s.BaseManifestList = &v
	return s
}

func (s *Snapshot) SetChangelogManifestList(v string) *Snapshot {
	s.ChangelogManifestList = &v
	return s
}

func (s *Snapshot) SetChangelogRecordCount(v int64) *Snapshot {
	s.ChangelogRecordCount = &v
	return s
}

func (s *Snapshot) SetCommitIdentifier(v int64) *Snapshot {
	s.CommitIdentifier = &v
	return s
}

func (s *Snapshot) SetCommitKind(v string) *Snapshot {
	s.CommitKind = &v
	return s
}

func (s *Snapshot) SetCommitUser(v string) *Snapshot {
	s.CommitUser = &v
	return s
}

func (s *Snapshot) SetDeltaManifestList(v string) *Snapshot {
	s.DeltaManifestList = &v
	return s
}

func (s *Snapshot) SetDeltaRecordCount(v int64) *Snapshot {
	s.DeltaRecordCount = &v
	return s
}

func (s *Snapshot) SetId(v int64) *Snapshot {
	s.Id = &v
	return s
}

func (s *Snapshot) SetIndexManifest(v string) *Snapshot {
	s.IndexManifest = &v
	return s
}

func (s *Snapshot) SetLogOffsets(v map[string]*int64) *Snapshot {
	s.LogOffsets = v
	return s
}

func (s *Snapshot) SetSchemaId(v int64) *Snapshot {
	s.SchemaId = &v
	return s
}

func (s *Snapshot) SetStatistics(v string) *Snapshot {
	s.Statistics = &v
	return s
}

func (s *Snapshot) SetTimeMillis(v int64) *Snapshot {
	s.TimeMillis = &v
	return s
}

func (s *Snapshot) SetTotalRecordCount(v int64) *Snapshot {
	s.TotalRecordCount = &v
	return s
}

func (s *Snapshot) SetVersion(v int32) *Snapshot {
	s.Version = &v
	return s
}

func (s *Snapshot) SetWatermark(v int64) *Snapshot {
	s.Watermark = &v
	return s
}

type Table struct {
	CreatedAt  *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy  *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsExternal *bool   `json:"isExternal,omitempty" xml:"isExternal,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Owner      *string `json:"owner,omitempty" xml:"owner,omitempty"`
	Path       *string `json:"path,omitempty" xml:"path,omitempty"`
	Schema     *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
	SchemaId   *int64  `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	UpdatedAt  *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy  *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Table) String() string {
	return tea.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) SetCreatedAt(v int64) *Table {
	s.CreatedAt = &v
	return s
}

func (s *Table) SetCreatedBy(v string) *Table {
	s.CreatedBy = &v
	return s
}

func (s *Table) SetId(v string) *Table {
	s.Id = &v
	return s
}

func (s *Table) SetIsExternal(v bool) *Table {
	s.IsExternal = &v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetOwner(v string) *Table {
	s.Owner = &v
	return s
}

func (s *Table) SetPath(v string) *Table {
	s.Path = &v
	return s
}

func (s *Table) SetSchema(v *Schema) *Table {
	s.Schema = v
	return s
}

func (s *Table) SetSchemaId(v int64) *Table {
	s.SchemaId = &v
	return s
}

func (s *Table) SetUpdatedAt(v int64) *Table {
	s.UpdatedAt = &v
	return s
}

func (s *Table) SetUpdatedBy(v string) *Table {
	s.UpdatedBy = &v
	return s
}

type TableCompaction struct {
	CatalogId             *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	LastCompactedFileTime *int64  `json:"lastCompactedFileTime,omitempty" xml:"lastCompactedFileTime,omitempty"`
	MaxLevel0FileCount    *string `json:"maxLevel0FileCount,omitempty" xml:"maxLevel0FileCount,omitempty"`
	TableId               *string `json:"tableId,omitempty" xml:"tableId,omitempty"`
}

func (s TableCompaction) String() string {
	return tea.Prettify(s)
}

func (s TableCompaction) GoString() string {
	return s.String()
}

func (s *TableCompaction) SetCatalogId(v string) *TableCompaction {
	s.CatalogId = &v
	return s
}

func (s *TableCompaction) SetLastCompactedFileTime(v int64) *TableCompaction {
	s.LastCompactedFileTime = &v
	return s
}

func (s *TableCompaction) SetMaxLevel0FileCount(v string) *TableCompaction {
	s.MaxLevel0FileCount = &v
	return s
}

func (s *TableCompaction) SetTableId(v string) *TableCompaction {
	s.TableId = &v
	return s
}

type TableSnapshot struct {
	FileCount            *int64    `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	FileSizeInBytes      *int64    `json:"fileSizeInBytes,omitempty" xml:"fileSizeInBytes,omitempty"`
	LastFileCreationTime *int64    `json:"lastFileCreationTime,omitempty" xml:"lastFileCreationTime,omitempty"`
	RecordCount          *int64    `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	Snapshot             *Snapshot `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
}

func (s TableSnapshot) String() string {
	return tea.Prettify(s)
}

func (s TableSnapshot) GoString() string {
	return s.String()
}

func (s *TableSnapshot) SetFileCount(v int64) *TableSnapshot {
	s.FileCount = &v
	return s
}

func (s *TableSnapshot) SetFileSizeInBytes(v int64) *TableSnapshot {
	s.FileSizeInBytes = &v
	return s
}

func (s *TableSnapshot) SetLastFileCreationTime(v int64) *TableSnapshot {
	s.LastFileCreationTime = &v
	return s
}

func (s *TableSnapshot) SetRecordCount(v int64) *TableSnapshot {
	s.RecordCount = &v
	return s
}

func (s *TableSnapshot) SetSnapshot(v *Snapshot) *TableSnapshot {
	s.Snapshot = v
	return s
}

type TableSummary struct {
	// Latest snapshot storage size
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Database name
	DatabaseName   *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	GeneratedDate  *string `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	LastAccessTime *int64  `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// Creation timestamp in milliseconds
	PartitionCount *int64  `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// 30-day access count
	TotalFileCount       *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
}

func (s TableSummary) String() string {
	return tea.Prettify(s)
}

func (s TableSummary) GoString() string {
	return s.String()
}

func (s *TableSummary) SetCreatedAt(v int64) *TableSummary {
	s.CreatedAt = &v
	return s
}

func (s *TableSummary) SetDatabaseName(v string) *TableSummary {
	s.DatabaseName = &v
	return s
}

func (s *TableSummary) SetGeneratedDate(v string) *TableSummary {
	s.GeneratedDate = &v
	return s
}

func (s *TableSummary) SetLastAccessTime(v int64) *TableSummary {
	s.LastAccessTime = &v
	return s
}

func (s *TableSummary) SetPartitionCount(v int64) *TableSummary {
	s.PartitionCount = &v
	return s
}

func (s *TableSummary) SetPath(v string) *TableSummary {
	s.Path = &v
	return s
}

func (s *TableSummary) SetTableName(v string) *TableSummary {
	s.TableName = &v
	return s
}

func (s *TableSummary) SetTotalFileCount(v int64) *TableSummary {
	s.TotalFileCount = &v
	return s
}

func (s *TableSummary) SetTotalFileSizeInBytes(v int64) *TableSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

type User struct {
	CreatedAt     *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy     *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	DisplayName   *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt     *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy     *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName      *string `json:"userName,omitempty" xml:"userName,omitempty"`
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s User) String() string {
	return tea.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) SetCreatedAt(v int64) *User {
	s.CreatedAt = &v
	return s
}

func (s *User) SetCreatedBy(v string) *User {
	s.CreatedBy = &v
	return s
}

func (s *User) SetDisplayName(v string) *User {
	s.DisplayName = &v
	return s
}

func (s *User) SetType(v string) *User {
	s.Type = &v
	return s
}

func (s *User) SetUpdatedAt(v int64) *User {
	s.UpdatedAt = &v
	return s
}

func (s *User) SetUpdatedBy(v string) *User {
	s.UpdatedBy = &v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

func (s *User) SetUserPrincipal(v string) *User {
	s.UserPrincipal = &v
	return s
}

type AlterCatalogRequest struct {
	Removals []*string          `json:"removals,omitempty" xml:"removals,omitempty" type:"Repeated"`
	Updates  map[string]*string `json:"updates,omitempty" xml:"updates,omitempty"`
}

func (s AlterCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s AlterCatalogRequest) GoString() string {
	return s.String()
}

func (s *AlterCatalogRequest) SetRemovals(v []*string) *AlterCatalogRequest {
	s.Removals = v
	return s
}

func (s *AlterCatalogRequest) SetUpdates(v map[string]*string) *AlterCatalogRequest {
	s.Updates = v
	return s
}

type AlterCatalogResponseBody struct {
	Missing []*string `json:"missing,omitempty" xml:"missing,omitempty" type:"Repeated"`
	Removed []*string `json:"removed,omitempty" xml:"removed,omitempty" type:"Repeated"`
	Updated []*string `json:"updated,omitempty" xml:"updated,omitempty" type:"Repeated"`
}

func (s AlterCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AlterCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *AlterCatalogResponseBody) SetMissing(v []*string) *AlterCatalogResponseBody {
	s.Missing = v
	return s
}

func (s *AlterCatalogResponseBody) SetRemoved(v []*string) *AlterCatalogResponseBody {
	s.Removed = v
	return s
}

func (s *AlterCatalogResponseBody) SetUpdated(v []*string) *AlterCatalogResponseBody {
	s.Updated = v
	return s
}

type AlterCatalogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlterCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlterCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s AlterCatalogResponse) GoString() string {
	return s.String()
}

func (s *AlterCatalogResponse) SetHeaders(v map[string]*string) *AlterCatalogResponse {
	s.Headers = v
	return s
}

func (s *AlterCatalogResponse) SetStatusCode(v int32) *AlterCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterCatalogResponse) SetBody(v *AlterCatalogResponseBody) *AlterCatalogResponse {
	s.Body = v
	return s
}

type BatchGrantPermissionsRequest struct {
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s BatchGrantPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsRequest) SetPermissions(v []*Permission) *BatchGrantPermissionsRequest {
	s.Permissions = v
	return s
}

type BatchGrantPermissionsResponseBody struct {
	ErrorMessage       *string              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FailurePermissions []*FailurePermission `json:"failurePermissions,omitempty" xml:"failurePermissions,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchGrantPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponseBody) SetErrorMessage(v string) *BatchGrantPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetFailurePermissions(v []*FailurePermission) *BatchGrantPermissionsResponseBody {
	s.FailurePermissions = v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetSuccess(v bool) *BatchGrantPermissionsResponseBody {
	s.Success = &v
	return s
}

type BatchGrantPermissionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGrantPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGrantPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponse) SetHeaders(v map[string]*string) *BatchGrantPermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchGrantPermissionsResponse) SetStatusCode(v int32) *BatchGrantPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGrantPermissionsResponse) SetBody(v *BatchGrantPermissionsResponseBody) *BatchGrantPermissionsResponse {
	s.Body = v
	return s
}

type BatchRevokePermissionsRequest struct {
	Permissions []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s BatchRevokePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsRequest) SetPermissions(v []*Permission) *BatchRevokePermissionsRequest {
	s.Permissions = v
	return s
}

type BatchRevokePermissionsResponseBody struct {
	ErrorMessage       *string              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FailurePermissions []*FailurePermission `json:"failurePermissions,omitempty" xml:"failurePermissions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchRevokePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponseBody) SetErrorMessage(v string) *BatchRevokePermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetFailurePermissions(v []*FailurePermission) *BatchRevokePermissionsResponseBody {
	s.FailurePermissions = v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetSuccess(v bool) *BatchRevokePermissionsResponseBody {
	s.Success = &v
	return s
}

type BatchRevokePermissionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRevokePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRevokePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponse) SetHeaders(v map[string]*string) *BatchRevokePermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchRevokePermissionsResponse) SetStatusCode(v int32) *BatchRevokePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRevokePermissionsResponse) SetBody(v *BatchRevokePermissionsResponseBody) *BatchRevokePermissionsResponse {
	s.Body = v
	return s
}

type CreateCatalogRequest struct {
	// example:
	//
	// catalog_demo
	Name               *string            `json:"name,omitempty" xml:"name,omitempty"`
	OptimizationConfig map[string]*string `json:"optimizationConfig,omitempty" xml:"optimizationConfig,omitempty"`
	Options            map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Type               *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCatalogRequest) GoString() string {
	return s.String()
}

func (s *CreateCatalogRequest) SetName(v string) *CreateCatalogRequest {
	s.Name = &v
	return s
}

func (s *CreateCatalogRequest) SetOptimizationConfig(v map[string]*string) *CreateCatalogRequest {
	s.OptimizationConfig = v
	return s
}

func (s *CreateCatalogRequest) SetOptions(v map[string]*string) *CreateCatalogRequest {
	s.Options = v
	return s
}

func (s *CreateCatalogRequest) SetType(v string) *CreateCatalogRequest {
	s.Type = &v
	return s
}

type CreateCatalogResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCatalogResponse) GoString() string {
	return s.String()
}

func (s *CreateCatalogResponse) SetHeaders(v map[string]*string) *CreateCatalogResponse {
	s.Headers = v
	return s
}

func (s *CreateCatalogResponse) SetStatusCode(v int32) *CreateCatalogResponse {
	s.StatusCode = &v
	return s
}

type CreateRoleRequest struct {
	// example:
	//
	// role_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// role_display_name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// role_name
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetDisplayName(v string) *CreateRoleRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

type DeleteRoleRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) SetRolePrincipal(v string) *DeleteRoleRequest {
	s.RolePrincipal = &v
	return s
}

type DeleteRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponse) SetHeaders(v map[string]*string) *DeleteRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoleResponse) SetStatusCode(v int32) *DeleteRoleResponse {
	s.StatusCode = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions []*DescribeRegionsResponseBodyRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The region description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The region name
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region show name
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// The region type
	//
	// example:
	//
	// region
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetDescription(v string) *DescribeRegionsResponseBodyRegions {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetName(v string) *DescribeRegionsResponseBodyRegions {
	s.Name = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetShowName(v string) *DescribeRegionsResponseBodyRegions {
	s.ShowName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetType(v string) *DescribeRegionsResponseBodyRegions {
	s.Type = &v
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

type DropCatalogResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s DropCatalogResponse) GoString() string {
	return s.String()
}

func (s *DropCatalogResponse) SetHeaders(v map[string]*string) *DropCatalogResponse {
	s.Headers = v
	return s
}

func (s *DropCatalogResponse) SetStatusCode(v int32) *DropCatalogResponse {
	s.StatusCode = &v
	return s
}

type GetCatalogResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Catalog           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogResponse) SetHeaders(v map[string]*string) *GetCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogResponse) SetStatusCode(v int32) *GetCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogResponse) SetBody(v *Catalog) *GetCatalogResponse {
	s.Body = v
	return s
}

type GetCatalogSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogSummary    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryResponse) SetHeaders(v map[string]*string) *GetCatalogSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogSummaryResponse) SetStatusCode(v int32) *GetCatalogSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogSummaryResponse) SetBody(v *CatalogSummary) *GetCatalogSummaryResponse {
	s.Body = v
	return s
}

type GetCatalogSummaryTrendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-01
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s GetCatalogSummaryTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSummaryTrendRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryTrendRequest) SetEndDate(v string) *GetCatalogSummaryTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetCatalogSummaryTrendRequest) SetStartDate(v string) *GetCatalogSummaryTrendRequest {
	s.StartDate = &v
	return s
}

type GetCatalogSummaryTrendResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogSummaryTrend `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogSummaryTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSummaryTrendResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogSummaryTrendResponse) SetHeaders(v map[string]*string) *GetCatalogSummaryTrendResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogSummaryTrendResponse) SetStatusCode(v int32) *GetCatalogSummaryTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogSummaryTrendResponse) SetBody(v *CatalogSummaryTrend) *GetCatalogSummaryTrendResponse {
	s.Body = v
	return s
}

type GetCatalogTokenResponseBody struct {
	// example:
	//
	// 1749160909000
	ExpiresAtMillis *int64             `json:"expiresAtMillis,omitempty" xml:"expiresAtMillis,omitempty"`
	Token           map[string]*string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetCatalogTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogTokenResponseBody) SetExpiresAtMillis(v int64) *GetCatalogTokenResponseBody {
	s.ExpiresAtMillis = &v
	return s
}

func (s *GetCatalogTokenResponseBody) SetToken(v map[string]*string) *GetCatalogTokenResponseBody {
	s.Token = v
	return s
}

type GetCatalogTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogTokenResponse) SetHeaders(v map[string]*string) *GetCatalogTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogTokenResponse) SetStatusCode(v int32) *GetCatalogTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogTokenResponse) SetBody(v *GetCatalogTokenResponseBody) *GetCatalogTokenResponse {
	s.Body = v
	return s
}

type GetDatabaseSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DatabaseSummary   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseSummaryResponse) SetHeaders(v map[string]*string) *GetDatabaseSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseSummaryResponse) SetStatusCode(v int32) *GetDatabaseSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseSummaryResponse) SetBody(v *DatabaseSummary) *GetDatabaseSummaryResponse {
	s.Body = v
	return s
}

type GetRegionStatusResponseBody struct {
	// example:
	//
	// true
	ServiceRoleExists *bool `json:"serviceRoleExists,omitempty" xml:"serviceRoleExists,omitempty"`
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetRegionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponseBody) SetServiceRoleExists(v bool) *GetRegionStatusResponseBody {
	s.ServiceRoleExists = &v
	return s
}

func (s *GetRegionStatusResponseBody) SetStatus(v string) *GetRegionStatusResponseBody {
	s.Status = &v
	return s
}

type GetRegionStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponse) SetHeaders(v map[string]*string) *GetRegionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRegionStatusResponse) SetStatusCode(v int32) *GetRegionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionStatusResponse) SetBody(v *GetRegionStatusResponseBody) *GetRegionStatusResponse {
	s.Body = v
	return s
}

type GetRoleRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s GetRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) SetRolePrincipal(v string) *GetRoleRequest {
	s.RolePrincipal = &v
	return s
}

type GetRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Role              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponse) GoString() string {
	return s.String()
}

func (s *GetRoleResponse) SetHeaders(v map[string]*string) *GetRoleResponse {
	s.Headers = v
	return s
}

func (s *GetRoleResponse) SetStatusCode(v int32) *GetRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleResponse) SetBody(v *Role) *GetRoleResponse {
	s.Body = v
	return s
}

type GetTableSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TableSummary      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTableSummaryResponse) SetHeaders(v map[string]*string) *GetTableSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTableSummaryResponse) SetStatusCode(v int32) *GetTableSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableSummaryResponse) SetBody(v *TableSummary) *GetTableSummaryResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// example:
	//
	// acs:ram::[accountId]:user/user_name
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserPrincipal(v string) *GetUserRequest {
	s.UserPrincipal = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *User) *GetUserResponse {
	s.Body = v
	return s
}

type GrantRoleToUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s GrantRoleToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersRequest) SetRolePrincipal(v string) *GrantRoleToUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *GrantRoleToUsersRequest) SetUserPrincipals(v []*string) *GrantRoleToUsersRequest {
	s.UserPrincipals = v
	return s
}

type GrantRoleToUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GrantRoleToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersResponse) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponse) SetHeaders(v map[string]*string) *GrantRoleToUsersResponse {
	s.Headers = v
	return s
}

func (s *GrantRoleToUsersResponse) SetStatusCode(v int32) *GrantRoleToUsersResponse {
	s.StatusCode = &v
	return s
}

type ListCatalogsRequest struct {
	CatalogNamePattern *string `json:"catalogNamePattern,omitempty" xml:"catalogNamePattern,omitempty"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListCatalogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) SetCatalogNamePattern(v string) *ListCatalogsRequest {
	s.CatalogNamePattern = &v
	return s
}

func (s *ListCatalogsRequest) SetMaxResults(v int32) *ListCatalogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCatalogsRequest) SetPageToken(v string) *ListCatalogsRequest {
	s.PageToken = &v
	return s
}

type ListCatalogsResponseBody struct {
	Catalogs []*Catalog `json:"catalogs,omitempty" xml:"catalogs,omitempty" type:"Repeated"`
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) SetCatalogs(v []*Catalog) *ListCatalogsResponseBody {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBody) SetNextPageToken(v string) *ListCatalogsResponseBody {
	s.NextPageToken = &v
	return s
}

type ListCatalogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCatalogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsResponse) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponse) SetHeaders(v map[string]*string) *ListCatalogsResponse {
	s.Headers = v
	return s
}

func (s *ListCatalogsResponse) SetStatusCode(v int32) *ListCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCatalogsResponse) SetBody(v *ListCatalogsResponseBody) *ListCatalogsResponse {
	s.Body = v
	return s
}

type ListPartitionSummariesRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// hh=10
	PartitionNamePattern *string `json:"partitionNamePattern,omitempty" xml:"partitionNamePattern,omitempty"`
}

func (s ListPartitionSummariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionSummariesRequest) SetMaxResults(v int32) *ListPartitionSummariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPartitionSummariesRequest) SetPageToken(v string) *ListPartitionSummariesRequest {
	s.PageToken = &v
	return s
}

func (s *ListPartitionSummariesRequest) SetPartitionNamePattern(v string) *ListPartitionSummariesRequest {
	s.PartitionNamePattern = &v
	return s
}

type ListPartitionSummariesResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PartitionSummaries `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPartitionSummariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionSummariesResponse) SetHeaders(v map[string]*string) *ListPartitionSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionSummariesResponse) SetStatusCode(v int32) *ListPartitionSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionSummariesResponse) SetBody(v *PartitionSummaries) *ListPartitionSummariesResponse {
	s.Body = v
	return s
}

type ListPermissionsRequest struct {
	// example:
	//
	// database_name
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:user/user_name
	Principal *string `json:"principal,omitempty" xml:"principal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CATALOG
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// table_name
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) SetDatabase(v string) *ListPermissionsRequest {
	s.Database = &v
	return s
}

func (s *ListPermissionsRequest) SetMaxResults(v int32) *ListPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequest) SetPageToken(v string) *ListPermissionsRequest {
	s.PageToken = &v
	return s
}

func (s *ListPermissionsRequest) SetPrincipal(v string) *ListPermissionsRequest {
	s.Principal = &v
	return s
}

func (s *ListPermissionsRequest) SetResourceType(v string) *ListPermissionsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionsRequest) SetTable(v string) *ListPermissionsRequest {
	s.Table = &v
	return s
}

type ListPermissionsResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string       `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Permissions   []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetNextPageToken(v string) *ListPermissionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*Permission) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListRoleUsersRequest struct {
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s ListRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoleUsersRequest) SetMaxResults(v string) *ListRoleUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRoleUsersRequest) SetPageToken(v string) *ListRoleUsersRequest {
	s.PageToken = &v
	return s
}

func (s *ListRoleUsersRequest) SetRolePrincipal(v string) *ListRoleUsersRequest {
	s.RolePrincipal = &v
	return s
}

type ListRoleUsersResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Users         []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponseBody) SetNextPageToken(v string) *ListRoleUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetUsers(v []*User) *ListRoleUsersResponseBody {
	s.Users = v
	return s
}

type ListRoleUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponse) SetHeaders(v map[string]*string) *ListRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRoleUsersResponse) SetStatusCode(v int32) *ListRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleUsersResponse) SetBody(v *ListRoleUsersResponseBody) *ListRoleUsersResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// role_name
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetMaxResults(v int32) *ListRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRolesRequest) SetPageToken(v string) *ListRolesRequest {
	s.PageToken = &v
	return s
}

func (s *ListRolesRequest) SetRoleName(v string) *ListRolesRequest {
	s.RoleName = &v
	return s
}

type ListRolesResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Roles         []*Role `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetNextPageToken(v string) *ListRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v []*Role) *ListRolesResponseBody {
	s.Roles = v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListUserRolesRequest struct {
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:user/user_name
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s ListUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ListUserRolesRequest) SetMaxResults(v string) *ListUserRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserRolesRequest) SetPageToken(v string) *ListUserRolesRequest {
	s.PageToken = &v
	return s
}

func (s *ListUserRolesRequest) SetUserPrincipal(v string) *ListUserRolesRequest {
	s.UserPrincipal = &v
	return s
}

type ListUserRolesResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Roles         []*Role `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponseBody) SetNextPageToken(v string) *ListUserRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListUserRolesResponseBody) SetRoles(v []*Role) *ListUserRolesResponseBody {
	s.Roles = v
	return s
}

type ListUserRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponse) SetHeaders(v map[string]*string) *ListUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ListUserRolesResponse) SetStatusCode(v int32) *ListUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRolesResponse) SetBody(v *ListUserRolesResponseBody) *ListUserRolesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// RAM_USER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// user_name
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetPageToken(v string) *ListUsersRequest {
	s.PageToken = &v
	return s
}

func (s *ListUsersRequest) SetType(v string) *ListUsersRequest {
	s.Type = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// token!
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Users         []*User `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetNextPageToken(v string) *ListUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*User) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type RevokeRoleFromUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s RevokeRoleFromUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeRoleFromUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersRequest) SetRolePrincipal(v string) *RevokeRoleFromUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *RevokeRoleFromUsersRequest) SetUserPrincipals(v []*string) *RevokeRoleFromUsersRequest {
	s.UserPrincipals = v
	return s
}

type RevokeRoleFromUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RevokeRoleFromUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeRoleFromUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersResponse) SetHeaders(v map[string]*string) *RevokeRoleFromUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeRoleFromUsersResponse) SetStatusCode(v int32) *RevokeRoleFromUsersResponse {
	s.StatusCode = &v
	return s
}

type SubscribeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeResponse) GoString() string {
	return s.String()
}

func (s *SubscribeResponse) SetHeaders(v map[string]*string) *SubscribeResponse {
	s.Headers = v
	return s
}

func (s *SubscribeResponse) SetStatusCode(v int32) *SubscribeResponse {
	s.StatusCode = &v
	return s
}

type UpdateRoleRequest struct {
	// example:
	//
	// role_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// role_display_name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal *string `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetDescription(v string) *UpdateRoleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRoleRequest) SetDisplayName(v string) *UpdateRoleRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateRoleRequest) SetRolePrincipal(v string) *UpdateRoleRequest {
	s.RolePrincipal = &v
	return s
}

type UpdateRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetStatusCode(v int32) *UpdateRoleResponse {
	s.StatusCode = &v
	return s
}

type UpdateRoleUsersRequest struct {
	// example:
	//
	// acs:dlf::[accountId]:role/role_name
	RolePrincipal  *string   `json:"rolePrincipal,omitempty" xml:"rolePrincipal,omitempty"`
	UserPrincipals []*string `json:"userPrincipals,omitempty" xml:"userPrincipals,omitempty" type:"Repeated"`
}

func (s UpdateRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersRequest) SetRolePrincipal(v string) *UpdateRoleUsersRequest {
	s.RolePrincipal = &v
	return s
}

func (s *UpdateRoleUsersRequest) SetUserPrincipals(v []*string) *UpdateRoleUsersRequest {
	s.UserPrincipals = v
	return s
}

type UpdateRoleUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersResponse) SetHeaders(v map[string]*string) *UpdateRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleUsersResponse) SetStatusCode(v int32) *UpdateRoleUsersResponse {
	s.StatusCode = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dlfnext"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 更新数据目录
//
// @param request - AlterCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterCatalogResponse
func (client *Client) AlterCatalogWithOptions(catalog *string, request *AlterCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AlterCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Removals)) {
		body["removals"] = request.Removals
	}

	if !tea.BoolValue(util.IsUnset(request.Updates)) {
		body["updates"] = request.Updates
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AlterCatalog"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalog))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AlterCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据目录
//
// @param request - AlterCatalogRequest
//
// @return AlterCatalogResponse
func (client *Client) AlterCatalog(catalog *string, request *AlterCatalogRequest) (_result *AlterCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterCatalogResponse{}
	_body, _err := client.AlterCatalogWithOptions(catalog, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量授权
//
// @param request - BatchGrantPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGrantPermissionsResponse
func (client *Client) BatchGrantPermissionsWithOptions(catalogId *string, request *BatchGrantPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGrantPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGrantPermissions"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/permissions/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/batchgrant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量授权
//
// @param request - BatchGrantPermissionsRequest
//
// @return BatchGrantPermissionsResponse
func (client *Client) BatchGrantPermissions(catalogId *string, request *BatchGrantPermissionsRequest) (_result *BatchGrantPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.BatchGrantPermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量取消授权
//
// @param request - BatchRevokePermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRevokePermissionsResponse
func (client *Client) BatchRevokePermissionsWithOptions(catalogId *string, request *BatchRevokePermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchRevokePermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["permissions"] = request.Permissions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRevokePermissions"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/permissions/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/batchrevoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量取消授权
//
// @param request - BatchRevokePermissionsRequest
//
// @return BatchRevokePermissionsResponse
func (client *Client) BatchRevokePermissions(catalogId *string, request *BatchRevokePermissionsRequest) (_result *BatchRevokePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.BatchRevokePermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据目录
//
// @param request - CreateCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCatalogResponse
func (client *Client) CreateCatalogWithOptions(request *CreateCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptimizationConfig)) {
		body["optimizationConfig"] = request.OptimizationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCatalog"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据目录
//
// @param request - CreateCatalogRequest
//
// @return CreateCatalogResponse
func (client *Client) CreateCatalog(request *CreateCatalogRequest) (_result *CreateCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCatalogResponse{}
	_body, _err := client.CreateCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建角色
//
// @param request - CreateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建角色
//
// @param request - CreateRoleRequest
//
// @return CreateRoleResponse
func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除角色
//
// @param request - DeleteRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRole"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除角色
//
// @param request - DeleteRoleRequest
//
// @return DeleteRoleResponse
func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 DLF 开通地域
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/service/regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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

// Summary:
//
// 查询 DLF 开通地域
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据湖Catalog
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropCatalogResponse
func (client *Client) DropCatalogWithOptions(catalog *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DropCatalogResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DropCatalog"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalog))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DropCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据湖Catalog
//
// @return DropCatalogResponse
func (client *Client) DropCatalog(catalog *string) (_result *DropCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropCatalogResponse{}
	_body, _err := client.DropCatalogWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看数据湖Catalog
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogResponse
func (client *Client) GetCatalogWithOptions(catalog *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalog"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalog))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据湖Catalog
//
// @return GetCatalogResponse
func (client *Client) GetCatalog(catalog *string) (_result *GetCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogResponse{}
	_body, _err := client.GetCatalogWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogSummaryResponse
func (client *Client) GetCatalogSummaryWithOptions(catalogId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalogSummary"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/storage-summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @return GetCatalogSummaryResponse
func (client *Client) GetCatalogSummary(catalogId *string) (_result *GetCatalogSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogSummaryResponse{}
	_body, _err := client.GetCatalogSummaryWithOptions(catalogId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetCatalogSummaryTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogSummaryTrendResponse
func (client *Client) GetCatalogSummaryTrendWithOptions(catalogId *string, request *GetCatalogSummaryTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogSummaryTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalogSummaryTrend"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/storage-summary/trend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogSummaryTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetCatalogSummaryTrendRequest
//
// @return GetCatalogSummaryTrendResponse
func (client *Client) GetCatalogSummaryTrend(catalogId *string, request *GetCatalogSummaryTrendRequest) (_result *GetCatalogSummaryTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogSummaryTrendResponse{}
	_body, _err := client.GetCatalogSummaryTrendWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖Catalog的临时访问凭证
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogTokenResponse
func (client *Client) GetCatalogTokenWithOptions(catalog *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalogToken"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs/" + tea.StringValue(openapiutil.GetEncodeParam(catalog)) + "/token"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖Catalog的临时访问凭证
//
// @return GetCatalogTokenResponse
func (client *Client) GetCatalogToken(catalog *string) (_result *GetCatalogTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogTokenResponse{}
	_body, _err := client.GetCatalogTokenWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseSummaryResponse
func (client *Client) GetDatabaseSummaryWithOptions(catalogId *string, database *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatabaseSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatabaseSummary"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/databases/" + tea.StringValue(openapiutil.GetEncodeParam(database)) + "/storage-summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatabaseSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @return GetDatabaseSummaryResponse
func (client *Client) GetDatabaseSummary(catalogId *string, database *string) (_result *GetDatabaseSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatabaseSummaryResponse{}
	_body, _err := client.GetDatabaseSummaryWithOptions(catalogId, database, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 DLF 当前地域开通状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionStatusResponse
func (client *Client) GetRegionStatusWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRegionStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegionStatus"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/service/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 DLF 当前地域开通状态
//
// @return GetRegionStatusResponse
func (client *Client) GetRegionStatus() (_result *GetRegionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionStatusResponse{}
	_body, _err := client.GetRegionStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取角色
//
// @param request - GetRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithOptions(request *GetRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRole"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色
//
// @param request - GetRoleRequest
//
// @return GetRoleResponse
func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableSummaryResponse
func (client *Client) GetTableSummaryWithOptions(catalogId *string, database *string, table *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableSummary"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/databases/" + tea.StringValue(openapiutil.GetEncodeParam(database)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(table)) + "/storage-summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @return GetTableSummaryResponse
func (client *Client) GetTableSummary(catalogId *string, database *string, table *string) (_result *GetTableSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableSummaryResponse{}
	_body, _err := client.GetTableSummaryWithOptions(catalogId, database, table, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipal)) {
		query["userPrincipal"] = request.UserPrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量授予角色权限给用户
//
// @param request - GrantRoleToUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsersWithOptions(request *GrantRoleToUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipals)) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantRoleToUsers"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles/grant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量授予角色权限给用户
//
// @param request - GrantRoleToUsersRequest
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsers(request *GrantRoleToUsersRequest) (_result *GrantRoleToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.GrantRoleToUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看数据目录列表
//
// @param request - ListCatalogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogsWithOptions(request *ListCatalogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogNamePattern)) {
		query["catalogNamePattern"] = request.CatalogNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCatalogs"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/catalogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据目录列表
//
// @param request - ListCatalogsRequest
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogs(request *ListCatalogsRequest) (_result *ListCatalogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCatalogsResponse{}
	_body, _err := client.ListCatalogsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - ListPartitionSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPartitionSummariesResponse
func (client *Client) ListPartitionSummariesWithOptions(catalogId *string, database *string, table *string, request *ListPartitionSummariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartitionSummariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNamePattern)) {
		query["partitionNamePattern"] = request.PartitionNamePattern
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPartitionSummaries"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/databases/" + tea.StringValue(openapiutil.GetEncodeParam(database)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(table)) + "/partitions/storage-summary"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPartitionSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - ListPartitionSummariesRequest
//
// @return ListPartitionSummariesResponse
func (client *Client) ListPartitionSummaries(catalogId *string, database *string, table *string, request *ListPartitionSummariesRequest) (_result *ListPartitionSummariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionSummariesResponse{}
	_body, _err := client.ListPartitionSummariesWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定资源或指定Principal的权限信息
//
// @param request - ListPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithOptions(catalogId *string, request *ListPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.Principal)) {
		query["principal"] = request.Principal
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		query["table"] = request.Table
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/permissions/" + tea.StringValue(openapiutil.GetEncodeParam(catalogId)) + "/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定资源或指定Principal的权限信息
//
// @param request - ListPermissionsRequest
//
// @return ListPermissionsResponse
func (client *Client) ListPermissions(catalogId *string, request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取角色用户列表
//
// @param request - ListRoleUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoleUsersResponse
func (client *Client) ListRoleUsersWithOptions(request *ListRoleUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoleUsers"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles/users/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色用户列表
//
// @param request - ListRoleUsersRequest
//
// @return ListRoleUsersResponse
func (client *Client) ListRoleUsers(request *ListRoleUsersRequest) (_result *ListRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRoleUsersResponse{}
	_body, _err := client.ListRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取角色列表
//
// @param request - ListRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["roleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色列表
//
// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListUserRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserRolesResponse
func (client *Client) ListUserRolesWithOptions(request *ListUserRolesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipal)) {
		query["userPrincipal"] = request.UserPrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserRoles"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/users/roles/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListUserRolesRequest
//
// @return ListUserRolesResponse
func (client *Client) ListUserRoles(request *ListUserRolesRequest) (_result *ListUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserRolesResponse{}
	_body, _err := client.ListUserRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户列表
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["pageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/users/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户列表
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量取消授予角色权限给用户
//
// @param request - RevokeRoleFromUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeRoleFromUsersResponse
func (client *Client) RevokeRoleFromUsersWithOptions(request *RevokeRoleFromUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeRoleFromUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipals)) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeRoleFromUsers"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量取消授予角色权限给用户
//
// @param request - RevokeRoleFromUsersRequest
//
// @return RevokeRoleFromUsersResponse
func (client *Client) RevokeRoleFromUsers(request *RevokeRoleFromUsersRequest) (_result *RevokeRoleFromUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.RevokeRoleFromUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订阅当前地域的 DLF
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeResponse
func (client *Client) SubscribeWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubscribeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("Subscribe"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/service/subscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &SubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅当前地域的 DLF
//
// @return SubscribeResponse
func (client *Client) Subscribe() (_result *SubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubscribeResponse{}
	_body, _err := client.SubscribeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新角色
//
// @param request - UpdateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRole"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新角色
//
// @param request - UpdateRoleRequest
//
// @return UpdateRoleResponse
func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新角色用户
//
// @param request - UpdateRoleUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleUsersResponse
func (client *Client) UpdateRoleUsersWithOptions(request *UpdateRoleUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RolePrincipal)) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipals)) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoleUsers"),
		Version:     tea.String("2025-03-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dlf/v1/auth/roles/users"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新角色用户
//
// @param request - UpdateRoleUsersRequest
//
// @return UpdateRoleUsersResponse
func (client *Client) UpdateRoleUsers(request *UpdateRoleUsersRequest) (_result *UpdateRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.UpdateRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
