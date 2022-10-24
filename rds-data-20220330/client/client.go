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

type BatchExecuteStatementResult struct {
	GeneratedFieldsList [][]*Field `json:"GeneratedFieldsList,omitempty" xml:"GeneratedFieldsList,omitempty" type:"Repeated"`
}

func (s BatchExecuteStatementResult) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResult) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResult) SetGeneratedFieldsList(v [][]*Field) *BatchExecuteStatementResult {
	s.GeneratedFieldsList = v
	return s
}

type BeginTransactionResult struct {
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s BeginTransactionResult) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResult) GoString() string {
	return s.String()
}

func (s *BeginTransactionResult) SetTransactionId(v string) *BeginTransactionResult {
	s.TransactionId = &v
	return s
}

type ColumnMetadata struct {
	ArrayBaseColumnType *int32  `json:"ArrayBaseColumnType,omitempty" xml:"ArrayBaseColumnType,omitempty"`
	IsAutoIncrement     *bool   `json:"IsAutoIncrement,omitempty" xml:"IsAutoIncrement,omitempty"`
	IsCaseSensitive     *bool   `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	IsCurrency          *bool   `json:"IsCurrency,omitempty" xml:"IsCurrency,omitempty"`
	IsSigned            *bool   `json:"IsSigned,omitempty" xml:"IsSigned,omitempty"`
	Label               *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nullable            *int32  `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	Precision           *int32  `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Scale               *int32  `json:"Scale,omitempty" xml:"Scale,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName            *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ColumnMetadata) String() string {
	return tea.Prettify(s)
}

func (s ColumnMetadata) GoString() string {
	return s.String()
}

func (s *ColumnMetadata) SetArrayBaseColumnType(v int32) *ColumnMetadata {
	s.ArrayBaseColumnType = &v
	return s
}

func (s *ColumnMetadata) SetIsAutoIncrement(v bool) *ColumnMetadata {
	s.IsAutoIncrement = &v
	return s
}

func (s *ColumnMetadata) SetIsCaseSensitive(v bool) *ColumnMetadata {
	s.IsCaseSensitive = &v
	return s
}

func (s *ColumnMetadata) SetIsCurrency(v bool) *ColumnMetadata {
	s.IsCurrency = &v
	return s
}

func (s *ColumnMetadata) SetIsSigned(v bool) *ColumnMetadata {
	s.IsSigned = &v
	return s
}

func (s *ColumnMetadata) SetLabel(v string) *ColumnMetadata {
	s.Label = &v
	return s
}

func (s *ColumnMetadata) SetName(v string) *ColumnMetadata {
	s.Name = &v
	return s
}

func (s *ColumnMetadata) SetNullable(v int32) *ColumnMetadata {
	s.Nullable = &v
	return s
}

func (s *ColumnMetadata) SetPrecision(v int32) *ColumnMetadata {
	s.Precision = &v
	return s
}

func (s *ColumnMetadata) SetScale(v int32) *ColumnMetadata {
	s.Scale = &v
	return s
}

func (s *ColumnMetadata) SetSchemaName(v string) *ColumnMetadata {
	s.SchemaName = &v
	return s
}

func (s *ColumnMetadata) SetTableName(v string) *ColumnMetadata {
	s.TableName = &v
	return s
}

func (s *ColumnMetadata) SetType(v int32) *ColumnMetadata {
	s.Type = &v
	return s
}

func (s *ColumnMetadata) SetTypeName(v string) *ColumnMetadata {
	s.TypeName = &v
	return s
}

type CommitTransactionResult struct {
	TransactionStatus *string `json:"TransactionStatus,omitempty" xml:"TransactionStatus,omitempty"`
}

func (s CommitTransactionResult) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResult) GoString() string {
	return s.String()
}

func (s *CommitTransactionResult) SetTransactionStatus(v string) *CommitTransactionResult {
	s.TransactionStatus = &v
	return s
}

type ExecuteStatementResult struct {
	ColumnMetadata         []*ColumnMetadata `json:"ColumnMetadata,omitempty" xml:"ColumnMetadata,omitempty" type:"Repeated"`
	FormattedRecords       *string           `json:"FormattedRecords,omitempty" xml:"FormattedRecords,omitempty"`
	GeneratedFields        []*Field          `json:"GeneratedFields,omitempty" xml:"GeneratedFields,omitempty" type:"Repeated"`
	NumberOfRecordsUpdated *int64            `json:"NumberOfRecordsUpdated,omitempty" xml:"NumberOfRecordsUpdated,omitempty"`
	Records                [][]*Field        `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ExecuteStatementResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResult) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResult) SetColumnMetadata(v []*ColumnMetadata) *ExecuteStatementResult {
	s.ColumnMetadata = v
	return s
}

func (s *ExecuteStatementResult) SetFormattedRecords(v string) *ExecuteStatementResult {
	s.FormattedRecords = &v
	return s
}

func (s *ExecuteStatementResult) SetGeneratedFields(v []*Field) *ExecuteStatementResult {
	s.GeneratedFields = v
	return s
}

func (s *ExecuteStatementResult) SetNumberOfRecordsUpdated(v int64) *ExecuteStatementResult {
	s.NumberOfRecordsUpdated = &v
	return s
}

func (s *ExecuteStatementResult) SetRecords(v [][]*Field) *ExecuteStatementResult {
	s.Records = v
	return s
}

type Field struct {
	ArrayValue   *string `json:"ArrayValue,omitempty" xml:"ArrayValue,omitempty"`
	BlobValue    *string `json:"BlobValue,omitempty" xml:"BlobValue,omitempty"`
	BooleanValue *bool   `json:"BooleanValue,omitempty" xml:"BooleanValue,omitempty"`
	IsNull       *bool   `json:"IsNull,omitempty" xml:"IsNull,omitempty"`
	LongValue    *int64  `json:"LongValue,omitempty" xml:"LongValue,omitempty"`
	StringValue  *string `json:"StringValue,omitempty" xml:"StringValue,omitempty"`
}

func (s Field) String() string {
	return tea.Prettify(s)
}

func (s Field) GoString() string {
	return s.String()
}

func (s *Field) SetArrayValue(v string) *Field {
	s.ArrayValue = &v
	return s
}

func (s *Field) SetBlobValue(v string) *Field {
	s.BlobValue = &v
	return s
}

func (s *Field) SetBooleanValue(v bool) *Field {
	s.BooleanValue = &v
	return s
}

func (s *Field) SetIsNull(v bool) *Field {
	s.IsNull = &v
	return s
}

func (s *Field) SetLongValue(v int64) *Field {
	s.LongValue = &v
	return s
}

func (s *Field) SetStringValue(v string) *Field {
	s.StringValue = &v
	return s
}

type ResultSetOptions struct {
	DecimalReturnType *string `json:"DecimalReturnType,omitempty" xml:"DecimalReturnType,omitempty"`
	LongReturnType    *string `json:"LongReturnType,omitempty" xml:"LongReturnType,omitempty"`
}

func (s ResultSetOptions) String() string {
	return tea.Prettify(s)
}

func (s ResultSetOptions) GoString() string {
	return s.String()
}

func (s *ResultSetOptions) SetDecimalReturnType(v string) *ResultSetOptions {
	s.DecimalReturnType = &v
	return s
}

func (s *ResultSetOptions) SetLongReturnType(v string) *ResultSetOptions {
	s.LongReturnType = &v
	return s
}

type RollbackTransactionResult struct {
	TransactionStatus *string `json:"TransactionStatus,omitempty" xml:"TransactionStatus,omitempty"`
}

func (s RollbackTransactionResult) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResult) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResult) SetTransactionStatus(v string) *RollbackTransactionResult {
	s.TransactionStatus = &v
	return s
}

type SqlParameter struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TypeHint *string `json:"TypeHint,omitempty" xml:"TypeHint,omitempty"`
	Value    *Field  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SqlParameter) String() string {
	return tea.Prettify(s)
}

func (s SqlParameter) GoString() string {
	return s.String()
}

func (s *SqlParameter) SetName(v string) *SqlParameter {
	s.Name = &v
	return s
}

func (s *SqlParameter) SetTypeHint(v string) *SqlParameter {
	s.TypeHint = &v
	return s
}

func (s *SqlParameter) SetValue(v *Field) *SqlParameter {
	s.Value = v
	return s
}

type BatchExecuteStatementRequest struct {
	Database      *string           `json:"Database,omitempty" xml:"Database,omitempty"`
	ParameterSets [][]*SqlParameter `json:"ParameterSets,omitempty" xml:"ParameterSets,omitempty" type:"Repeated"`
	ResourceArn   *string           `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	SecretArn     *string           `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	Sql           *string           `json:"Sql,omitempty" xml:"Sql,omitempty"`
	TransactionId *string           `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s BatchExecuteStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementRequest) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementRequest) SetDatabase(v string) *BatchExecuteStatementRequest {
	s.Database = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetParameterSets(v [][]*SqlParameter) *BatchExecuteStatementRequest {
	s.ParameterSets = v
	return s
}

func (s *BatchExecuteStatementRequest) SetResourceArn(v string) *BatchExecuteStatementRequest {
	s.ResourceArn = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetSecretArn(v string) *BatchExecuteStatementRequest {
	s.SecretArn = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetSql(v string) *BatchExecuteStatementRequest {
	s.Sql = &v
	return s
}

func (s *BatchExecuteStatementRequest) SetTransactionId(v string) *BatchExecuteStatementRequest {
	s.TransactionId = &v
	return s
}

type BatchExecuteStatementShrinkRequest struct {
	Database            *string `json:"Database,omitempty" xml:"Database,omitempty"`
	ParameterSetsShrink *string `json:"ParameterSets,omitempty" xml:"ParameterSets,omitempty"`
	ResourceArn         *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	SecretArn           *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	Sql                 *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	TransactionId       *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s BatchExecuteStatementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementShrinkRequest) SetDatabase(v string) *BatchExecuteStatementShrinkRequest {
	s.Database = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetParameterSetsShrink(v string) *BatchExecuteStatementShrinkRequest {
	s.ParameterSetsShrink = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetResourceArn(v string) *BatchExecuteStatementShrinkRequest {
	s.ResourceArn = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetSecretArn(v string) *BatchExecuteStatementShrinkRequest {
	s.SecretArn = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetSql(v string) *BatchExecuteStatementShrinkRequest {
	s.Sql = &v
	return s
}

func (s *BatchExecuteStatementShrinkRequest) SetTransactionId(v string) *BatchExecuteStatementShrinkRequest {
	s.TransactionId = &v
	return s
}

type BatchExecuteStatementResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchExecuteStatementResult `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchExecuteStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponseBody) SetCode(v string) *BatchExecuteStatementResponseBody {
	s.Code = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetData(v *BatchExecuteStatementResult) *BatchExecuteStatementResponseBody {
	s.Data = v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetMessage(v string) *BatchExecuteStatementResponseBody {
	s.Message = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetRequestId(v string) *BatchExecuteStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchExecuteStatementResponseBody) SetSuccess(v bool) *BatchExecuteStatementResponseBody {
	s.Success = &v
	return s
}

type BatchExecuteStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchExecuteStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteStatementResponse) GoString() string {
	return s.String()
}

func (s *BatchExecuteStatementResponse) SetHeaders(v map[string]*string) *BatchExecuteStatementResponse {
	s.Headers = v
	return s
}

func (s *BatchExecuteStatementResponse) SetStatusCode(v int32) *BatchExecuteStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExecuteStatementResponse) SetBody(v *BatchExecuteStatementResponseBody) *BatchExecuteStatementResponse {
	s.Body = v
	return s
}

type BeginTransactionRequest struct {
	Database    *string `json:"Database,omitempty" xml:"Database,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	SecretArn   *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
}

func (s BeginTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionRequest) GoString() string {
	return s.String()
}

func (s *BeginTransactionRequest) SetDatabase(v string) *BeginTransactionRequest {
	s.Database = &v
	return s
}

func (s *BeginTransactionRequest) SetResourceArn(v string) *BeginTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *BeginTransactionRequest) SetSecretArn(v string) *BeginTransactionRequest {
	s.SecretArn = &v
	return s
}

type BeginTransactionResponseBody struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BeginTransactionResult `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BeginTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginTransactionResponseBody) SetCode(v string) *BeginTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *BeginTransactionResponseBody) SetData(v *BeginTransactionResult) *BeginTransactionResponseBody {
	s.Data = v
	return s
}

func (s *BeginTransactionResponseBody) SetMessage(v string) *BeginTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *BeginTransactionResponseBody) SetRequestId(v string) *BeginTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginTransactionResponseBody) SetSuccess(v bool) *BeginTransactionResponseBody {
	s.Success = &v
	return s
}

type BeginTransactionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BeginTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BeginTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s BeginTransactionResponse) GoString() string {
	return s.String()
}

func (s *BeginTransactionResponse) SetHeaders(v map[string]*string) *BeginTransactionResponse {
	s.Headers = v
	return s
}

func (s *BeginTransactionResponse) SetStatusCode(v int32) *BeginTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginTransactionResponse) SetBody(v *BeginTransactionResponseBody) *BeginTransactionResponse {
	s.Body = v
	return s
}

type CommitTransactionRequest struct {
	ResourceArn   *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	SecretArn     *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CommitTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionRequest) GoString() string {
	return s.String()
}

func (s *CommitTransactionRequest) SetResourceArn(v string) *CommitTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *CommitTransactionRequest) SetSecretArn(v string) *CommitTransactionRequest {
	s.SecretArn = &v
	return s
}

func (s *CommitTransactionRequest) SetTransactionId(v string) *CommitTransactionRequest {
	s.TransactionId = &v
	return s
}

type CommitTransactionResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CommitTransactionResult `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CommitTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *CommitTransactionResponseBody) SetCode(v string) *CommitTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *CommitTransactionResponseBody) SetData(v *CommitTransactionResult) *CommitTransactionResponseBody {
	s.Data = v
	return s
}

func (s *CommitTransactionResponseBody) SetMessage(v string) *CommitTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *CommitTransactionResponseBody) SetRequestId(v string) *CommitTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitTransactionResponseBody) SetSuccess(v bool) *CommitTransactionResponseBody {
	s.Success = &v
	return s
}

type CommitTransactionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CommitTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitTransactionResponse) GoString() string {
	return s.String()
}

func (s *CommitTransactionResponse) SetHeaders(v map[string]*string) *CommitTransactionResponse {
	s.Headers = v
	return s
}

func (s *CommitTransactionResponse) SetStatusCode(v int32) *CommitTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitTransactionResponse) SetBody(v *CommitTransactionResponseBody) *CommitTransactionResponse {
	s.Body = v
	return s
}

type ExecuteStatementRequest struct {
	ContinueAfterTimeout  *bool                                    `json:"ContinueAfterTimeout,omitempty" xml:"ContinueAfterTimeout,omitempty"`
	Database              *string                                  `json:"Database,omitempty" xml:"Database,omitempty"`
	FormatRecordsAs       *string                                  `json:"FormatRecordsAs,omitempty" xml:"FormatRecordsAs,omitempty"`
	IncludeResultMetadata *bool                                    `json:"IncludeResultMetadata,omitempty" xml:"IncludeResultMetadata,omitempty"`
	Parameters            []*SqlParameter                          `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	ResourceArn           *string                                  `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	ResultSetOptions      *ExecuteStatementRequestResultSetOptions `json:"ResultSetOptions,omitempty" xml:"ResultSetOptions,omitempty" type:"Struct"`
	SecretArn             *string                                  `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	Sql                   *string                                  `json:"Sql,omitempty" xml:"Sql,omitempty"`
	TransactionId         *string                                  `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s ExecuteStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequest) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequest) SetContinueAfterTimeout(v bool) *ExecuteStatementRequest {
	s.ContinueAfterTimeout = &v
	return s
}

func (s *ExecuteStatementRequest) SetDatabase(v string) *ExecuteStatementRequest {
	s.Database = &v
	return s
}

func (s *ExecuteStatementRequest) SetFormatRecordsAs(v string) *ExecuteStatementRequest {
	s.FormatRecordsAs = &v
	return s
}

func (s *ExecuteStatementRequest) SetIncludeResultMetadata(v bool) *ExecuteStatementRequest {
	s.IncludeResultMetadata = &v
	return s
}

func (s *ExecuteStatementRequest) SetParameters(v []*SqlParameter) *ExecuteStatementRequest {
	s.Parameters = v
	return s
}

func (s *ExecuteStatementRequest) SetResourceArn(v string) *ExecuteStatementRequest {
	s.ResourceArn = &v
	return s
}

func (s *ExecuteStatementRequest) SetResultSetOptions(v *ExecuteStatementRequestResultSetOptions) *ExecuteStatementRequest {
	s.ResultSetOptions = v
	return s
}

func (s *ExecuteStatementRequest) SetSecretArn(v string) *ExecuteStatementRequest {
	s.SecretArn = &v
	return s
}

func (s *ExecuteStatementRequest) SetSql(v string) *ExecuteStatementRequest {
	s.Sql = &v
	return s
}

func (s *ExecuteStatementRequest) SetTransactionId(v string) *ExecuteStatementRequest {
	s.TransactionId = &v
	return s
}

type ExecuteStatementRequestResultSetOptions struct {
	DecimalReturnType *string `json:"DecimalReturnType,omitempty" xml:"DecimalReturnType,omitempty"`
	LongReturnType    *string `json:"LongReturnType,omitempty" xml:"LongReturnType,omitempty"`
}

func (s ExecuteStatementRequestResultSetOptions) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementRequestResultSetOptions) GoString() string {
	return s.String()
}

func (s *ExecuteStatementRequestResultSetOptions) SetDecimalReturnType(v string) *ExecuteStatementRequestResultSetOptions {
	s.DecimalReturnType = &v
	return s
}

func (s *ExecuteStatementRequestResultSetOptions) SetLongReturnType(v string) *ExecuteStatementRequestResultSetOptions {
	s.LongReturnType = &v
	return s
}

type ExecuteStatementShrinkRequest struct {
	ContinueAfterTimeout   *bool   `json:"ContinueAfterTimeout,omitempty" xml:"ContinueAfterTimeout,omitempty"`
	Database               *string `json:"Database,omitempty" xml:"Database,omitempty"`
	FormatRecordsAs        *string `json:"FormatRecordsAs,omitempty" xml:"FormatRecordsAs,omitempty"`
	IncludeResultMetadata  *bool   `json:"IncludeResultMetadata,omitempty" xml:"IncludeResultMetadata,omitempty"`
	ParametersShrink       *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ResourceArn            *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	ResultSetOptionsShrink *string `json:"ResultSetOptions,omitempty" xml:"ResultSetOptions,omitempty"`
	SecretArn              *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	Sql                    *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	TransactionId          *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s ExecuteStatementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteStatementShrinkRequest) SetContinueAfterTimeout(v bool) *ExecuteStatementShrinkRequest {
	s.ContinueAfterTimeout = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetDatabase(v string) *ExecuteStatementShrinkRequest {
	s.Database = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetFormatRecordsAs(v string) *ExecuteStatementShrinkRequest {
	s.FormatRecordsAs = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetIncludeResultMetadata(v bool) *ExecuteStatementShrinkRequest {
	s.IncludeResultMetadata = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetParametersShrink(v string) *ExecuteStatementShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetResourceArn(v string) *ExecuteStatementShrinkRequest {
	s.ResourceArn = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetResultSetOptionsShrink(v string) *ExecuteStatementShrinkRequest {
	s.ResultSetOptionsShrink = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetSecretArn(v string) *ExecuteStatementShrinkRequest {
	s.SecretArn = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetSql(v string) *ExecuteStatementShrinkRequest {
	s.Sql = &v
	return s
}

func (s *ExecuteStatementShrinkRequest) SetTransactionId(v string) *ExecuteStatementShrinkRequest {
	s.TransactionId = &v
	return s
}

type ExecuteStatementResponseBody struct {
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ExecuteStatementResult `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponseBody) SetCode(v string) *ExecuteStatementResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetData(v *ExecuteStatementResult) *ExecuteStatementResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteStatementResponseBody) SetMessage(v string) *ExecuteStatementResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetRequestId(v string) *ExecuteStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteStatementResponseBody) SetSuccess(v bool) *ExecuteStatementResponseBody {
	s.Success = &v
	return s
}

type ExecuteStatementResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStatementResponse) GoString() string {
	return s.String()
}

func (s *ExecuteStatementResponse) SetHeaders(v map[string]*string) *ExecuteStatementResponse {
	s.Headers = v
	return s
}

func (s *ExecuteStatementResponse) SetStatusCode(v int32) *ExecuteStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteStatementResponse) SetBody(v *ExecuteStatementResponseBody) *ExecuteStatementResponse {
	s.Body = v
	return s
}

type RollbackTransactionRequest struct {
	ResourceArn   *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	SecretArn     *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s RollbackTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionRequest) GoString() string {
	return s.String()
}

func (s *RollbackTransactionRequest) SetResourceArn(v string) *RollbackTransactionRequest {
	s.ResourceArn = &v
	return s
}

func (s *RollbackTransactionRequest) SetSecretArn(v string) *RollbackTransactionRequest {
	s.SecretArn = &v
	return s
}

func (s *RollbackTransactionRequest) SetTransactionId(v string) *RollbackTransactionRequest {
	s.TransactionId = &v
	return s
}

type RollbackTransactionResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RollbackTransactionResult `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResponseBody) SetCode(v string) *RollbackTransactionResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetData(v *RollbackTransactionResult) *RollbackTransactionResponseBody {
	s.Data = v
	return s
}

func (s *RollbackTransactionResponseBody) SetMessage(v string) *RollbackTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetRequestId(v string) *RollbackTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackTransactionResponseBody) SetSuccess(v bool) *RollbackTransactionResponseBody {
	s.Success = &v
	return s
}

type RollbackTransactionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackTransactionResponse) GoString() string {
	return s.String()
}

func (s *RollbackTransactionResponse) SetHeaders(v map[string]*string) *RollbackTransactionResponse {
	s.Headers = v
	return s
}

func (s *RollbackTransactionResponse) SetStatusCode(v int32) *RollbackTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackTransactionResponse) SetBody(v *RollbackTransactionResponseBody) *RollbackTransactionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("rds-data"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchExecuteStatementWithOptions(tmpReq *BatchExecuteStatementRequest, runtime *util.RuntimeOptions) (_result *BatchExecuteStatementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchExecuteStatementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParameterSets)) {
		request.ParameterSetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParameterSets, tea.String("ParameterSets"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterSetsShrink)) {
		body["ParameterSets"] = request.ParameterSetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["SecretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["Sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchExecuteStatement"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchExecuteStatement(request *BatchExecuteStatementRequest) (_result *BatchExecuteStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchExecuteStatementResponse{}
	_body, _err := client.BatchExecuteStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BeginTransactionWithOptions(request *BeginTransactionRequest, runtime *util.RuntimeOptions) (_result *BeginTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["SecretArn"] = request.SecretArn
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BeginTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeginTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BeginTransaction(request *BeginTransactionRequest) (_result *BeginTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeginTransactionResponse{}
	_body, _err := client.BeginTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommitTransactionWithOptions(request *CommitTransactionRequest, runtime *util.RuntimeOptions) (_result *CommitTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["SecretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitTransaction(request *CommitTransactionRequest) (_result *CommitTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitTransactionResponse{}
	_body, _err := client.CommitTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteStatementWithOptions(tmpReq *ExecuteStatementRequest, runtime *util.RuntimeOptions) (_result *ExecuteStatementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteStatementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ResultSetOptions))) {
		request.ResultSetOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ResultSetOptions), tea.String("ResultSetOptions"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContinueAfterTimeout)) {
		body["ContinueAfterTimeout"] = request.ContinueAfterTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		body["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.FormatRecordsAs)) {
		body["FormatRecordsAs"] = request.FormatRecordsAs
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeResultMetadata)) {
		body["IncludeResultMetadata"] = request.IncludeResultMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSetOptionsShrink)) {
		body["ResultSetOptions"] = request.ResultSetOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["SecretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["Sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteStatement"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteStatement(request *ExecuteStatementRequest) (_result *ExecuteStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.ExecuteStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackTransactionWithOptions(request *RollbackTransactionRequest, runtime *util.RuntimeOptions) (_result *RollbackTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.SecretArn)) {
		body["SecretArn"] = request.SecretArn
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackTransaction"),
		Version:     tea.String("2022-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackTransaction(request *RollbackTransactionRequest) (_result *RollbackTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackTransactionResponse{}
	_body, _err := client.RollbackTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
