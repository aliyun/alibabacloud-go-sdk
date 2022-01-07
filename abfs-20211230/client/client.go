// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateBackFlowRequest struct {
	OdpsPartition *string `json:"odpsPartition,omitempty" xml:"odpsPartition,omitempty"`
	Timestamp     *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s CreateBackFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateBackFlowRequest) SetOdpsPartition(v string) *CreateBackFlowRequest {
	s.OdpsPartition = &v
	return s
}

func (s *CreateBackFlowRequest) SetTimestamp(v string) *CreateBackFlowRequest {
	s.Timestamp = &v
	return s
}

type CreateBackFlowResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateBackFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackFlowResponseBody) SetCode(v string) *CreateBackFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackFlowResponseBody) SetMessage(v string) *CreateBackFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackFlowResponseBody) SetRequestId(v string) *CreateBackFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackFlowResponseBody) SetResult(v map[string]interface{}) *CreateBackFlowResponseBody {
	s.Result = v
	return s
}

type CreateBackFlowResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateBackFlowResponse) SetHeaders(v map[string]*string) *CreateBackFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateBackFlowResponse) SetBody(v *CreateBackFlowResponseBody) *CreateBackFlowResponse {
	s.Body = v
	return s
}

type CreateInstanceTableRequest struct {
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 字段描述
	FieldMap *CreateInstanceTableRequestFieldMap `json:"fieldMap,omitempty" xml:"fieldMap,omitempty" type:"Struct"`
	// 全量数据源描述
	FullDataSourceConfig *CreateInstanceTableRequestFullDataSourceConfig `json:"fullDataSourceConfig,omitempty" xml:"fullDataSourceConfig,omitempty" type:"Struct"`
	// 增量数据源描述
	IncDataSourceConfig *CreateInstanceTableRequestIncDataSourceConfig `json:"incDataSourceConfig,omitempty" xml:"incDataSourceConfig,omitempty" type:"Struct"`
	// 索引类型
	IndexType *string `json:"indexType,omitempty" xml:"indexType,omitempty"`
	// 表名
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// 回流触发模式
	TriggerMode *string `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
	// Topic过期时间
	Ttl *int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateInstanceTableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequest) SetDescription(v string) *CreateInstanceTableRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceTableRequest) SetFieldMap(v *CreateInstanceTableRequestFieldMap) *CreateInstanceTableRequest {
	s.FieldMap = v
	return s
}

func (s *CreateInstanceTableRequest) SetFullDataSourceConfig(v *CreateInstanceTableRequestFullDataSourceConfig) *CreateInstanceTableRequest {
	s.FullDataSourceConfig = v
	return s
}

func (s *CreateInstanceTableRequest) SetIncDataSourceConfig(v *CreateInstanceTableRequestIncDataSourceConfig) *CreateInstanceTableRequest {
	s.IncDataSourceConfig = v
	return s
}

func (s *CreateInstanceTableRequest) SetIndexType(v string) *CreateInstanceTableRequest {
	s.IndexType = &v
	return s
}

func (s *CreateInstanceTableRequest) SetTableName(v string) *CreateInstanceTableRequest {
	s.TableName = &v
	return s
}

func (s *CreateInstanceTableRequest) SetTriggerMode(v string) *CreateInstanceTableRequest {
	s.TriggerMode = &v
	return s
}

func (s *CreateInstanceTableRequest) SetTtl(v int64) *CreateInstanceTableRequest {
	s.Ttl = &v
	return s
}

type CreateInstanceTableRequestFieldMap struct {
	Fields []*CreateInstanceTableRequestFieldMapFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// 索引信息
	Indexs          []*CreateInstanceTableRequestFieldMapIndexs `json:"indexs,omitempty" xml:"indexs,omitempty" type:"Repeated"`
	IsAttributePack *bool                                       `json:"isAttributePack,omitempty" xml:"isAttributePack,omitempty"`
	// 最大外建数
	MaxSkeyNum *int64 `json:"maxSkeyNum,omitempty" xml:"maxSkeyNum,omitempty"`
	// 离线构建线程数
	OfflineBuildProtectionThreshold *int64 `json:"offlineBuildProtectionThreshold,omitempty" xml:"offlineBuildProtectionThreshold,omitempty"`
	// 主键
	Pkey *string `json:"pkey,omitempty" xml:"pkey,omitempty"`
	// 记录类型
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s CreateInstanceTableRequestFieldMap) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestFieldMap) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestFieldMap) SetFields(v []*CreateInstanceTableRequestFieldMapFields) *CreateInstanceTableRequestFieldMap {
	s.Fields = v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetIndexs(v []*CreateInstanceTableRequestFieldMapIndexs) *CreateInstanceTableRequestFieldMap {
	s.Indexs = v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetIsAttributePack(v bool) *CreateInstanceTableRequestFieldMap {
	s.IsAttributePack = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetMaxSkeyNum(v int64) *CreateInstanceTableRequestFieldMap {
	s.MaxSkeyNum = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetOfflineBuildProtectionThreshold(v int64) *CreateInstanceTableRequestFieldMap {
	s.OfflineBuildProtectionThreshold = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetPkey(v string) *CreateInstanceTableRequestFieldMap {
	s.Pkey = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMap) SetRecordType(v string) *CreateInstanceTableRequestFieldMap {
	s.RecordType = &v
	return s
}

type CreateInstanceTableRequestFieldMapFields struct {
	// 默认值
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// 描述
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	IsMultiValue *bool   `json:"isMultiValue,omitempty" xml:"isMultiValue,omitempty"`
	IsVirtual    *bool   `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	OnlineStatus *bool   `json:"onlineStatus,omitempty" xml:"onlineStatus,omitempty"`
	// 是否为主键
	Pkey *bool `json:"pkey,omitempty" xml:"pkey,omitempty"`
	// 插件名
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	// 插件信息
	PluginParam *string `json:"pluginParam,omitempty" xml:"pluginParam,omitempty"`
	// 分隔符
	Seperator *string `json:"seperator,omitempty" xml:"seperator,omitempty"`
	// 源字段名
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// 源字段类型
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 目标字段名
	TargetName *string `json:"targetName,omitempty" xml:"targetName,omitempty"`
	// 目标字段类型
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateInstanceTableRequestFieldMapFields) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestFieldMapFields) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestFieldMapFields) SetDefaultValue(v string) *CreateInstanceTableRequestFieldMapFields {
	s.DefaultValue = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetDescription(v string) *CreateInstanceTableRequestFieldMapFields {
	s.Description = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetIsMultiValue(v bool) *CreateInstanceTableRequestFieldMapFields {
	s.IsMultiValue = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetIsVirtual(v bool) *CreateInstanceTableRequestFieldMapFields {
	s.IsVirtual = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetOnlineStatus(v bool) *CreateInstanceTableRequestFieldMapFields {
	s.OnlineStatus = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetPkey(v bool) *CreateInstanceTableRequestFieldMapFields {
	s.Pkey = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetPluginName(v string) *CreateInstanceTableRequestFieldMapFields {
	s.PluginName = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetPluginParam(v string) *CreateInstanceTableRequestFieldMapFields {
	s.PluginParam = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetSeperator(v string) *CreateInstanceTableRequestFieldMapFields {
	s.Seperator = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetSourceName(v string) *CreateInstanceTableRequestFieldMapFields {
	s.SourceName = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetSourceType(v string) *CreateInstanceTableRequestFieldMapFields {
	s.SourceType = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetTargetName(v string) *CreateInstanceTableRequestFieldMapFields {
	s.TargetName = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapFields) SetTargetType(v string) *CreateInstanceTableRequestFieldMapFields {
	s.TargetType = &v
	return s
}

type CreateInstanceTableRequestFieldMapIndexs struct {
	// 字段名
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// 索引字段名
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s CreateInstanceTableRequestFieldMapIndexs) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestFieldMapIndexs) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestFieldMapIndexs) SetField(v string) *CreateInstanceTableRequestFieldMapIndexs {
	s.Field = &v
	return s
}

func (s *CreateInstanceTableRequestFieldMapIndexs) SetIndexName(v string) *CreateInstanceTableRequestFieldMapIndexs {
	s.IndexName = &v
	return s
}

type CreateInstanceTableRequestFullDataSourceConfig struct {
	// MaxCompute描述
	OdpsDataDesc *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc `json:"odpsDataDesc,omitempty" xml:"odpsDataDesc,omitempty" type:"Struct"`
	// 数据源类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateInstanceTableRequestFullDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestFullDataSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestFullDataSourceConfig) SetOdpsDataDesc(v *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) *CreateInstanceTableRequestFullDataSourceConfig {
	s.OdpsDataDesc = v
	return s
}

func (s *CreateInstanceTableRequestFullDataSourceConfig) SetType(v string) *CreateInstanceTableRequestFullDataSourceConfig {
	s.Type = &v
	return s
}

type CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc struct {
	// 项目名
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// 表名
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) SetProject(v string) *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc {
	s.Project = &v
	return s
}

func (s *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) SetTable(v string) *CreateInstanceTableRequestFullDataSourceConfigOdpsDataDesc {
	s.Table = &v
	return s
}

type CreateInstanceTableRequestIncDataSourceConfig struct {
	SwiftDataDesc *CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc `json:"swiftDataDesc,omitempty" xml:"swiftDataDesc,omitempty" type:"Struct"`
	// 数据源类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateInstanceTableRequestIncDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestIncDataSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestIncDataSourceConfig) SetSwiftDataDesc(v *CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) *CreateInstanceTableRequestIncDataSourceConfig {
	s.SwiftDataDesc = v
	return s
}

func (s *CreateInstanceTableRequestIncDataSourceConfig) SetType(v string) *CreateInstanceTableRequestIncDataSourceConfig {
	s.Type = &v
	return s
}

type CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc struct {
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) SetTopic(v string) *CreateInstanceTableRequestIncDataSourceConfigSwiftDataDesc {
	s.Topic = &v
	return s
}

type CreateInstanceTableResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateInstanceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableResponseBody) SetCode(v string) *CreateInstanceTableResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceTableResponseBody) SetMessage(v string) *CreateInstanceTableResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceTableResponseBody) SetRequestId(v string) *CreateInstanceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceTableResponseBody) SetResult(v map[string]interface{}) *CreateInstanceTableResponseBody {
	s.Result = v
	return s
}

type CreateInstanceTableResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceTableResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceTableResponse) SetHeaders(v map[string]*string) *CreateInstanceTableResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceTableResponse) SetBody(v *CreateInstanceTableResponseBody) *CreateInstanceTableResponse {
	s.Body = v
	return s
}

type DeleteInstanceTableResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// reuslt
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInstanceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceTableResponseBody) SetCode(v string) *DeleteInstanceTableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceTableResponseBody) SetMessage(v string) *DeleteInstanceTableResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceTableResponseBody) SetRequestId(v string) *DeleteInstanceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceTableResponseBody) SetResult(v map[string]interface{}) *DeleteInstanceTableResponseBody {
	s.Result = v
	return s
}

type DeleteInstanceTableResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceTableResponse) SetHeaders(v map[string]*string) *DeleteInstanceTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceTableResponse) SetBody(v *DeleteInstanceTableResponseBody) *DeleteInstanceTableResponse {
	s.Body = v
	return s
}

type GetBackFlowResponseBody struct {
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// error message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetBackFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBackFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackFlowResponseBody) SetCode(v string) *GetBackFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetBackFlowResponseBody) SetMessage(v string) *GetBackFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetBackFlowResponseBody) SetRequestId(v string) *GetBackFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackFlowResponseBody) SetResult(v map[string]interface{}) *GetBackFlowResponseBody {
	s.Result = v
	return s
}

type GetBackFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBackFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBackFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBackFlowResponse) GoString() string {
	return s.String()
}

func (s *GetBackFlowResponse) SetHeaders(v map[string]*string) *GetBackFlowResponse {
	s.Headers = v
	return s
}

func (s *GetBackFlowResponse) SetBody(v *GetBackFlowResponseBody) *GetBackFlowResponse {
	s.Body = v
	return s
}

type GetIndexesResponseBody struct {
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// error message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIndexesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexesResponseBody) SetCode(v string) *GetIndexesResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexesResponseBody) SetMessage(v string) *GetIndexesResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexesResponseBody) SetRequestId(v string) *GetIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexesResponseBody) SetResult(v map[string]interface{}) *GetIndexesResponseBody {
	s.Result = v
	return s
}

type GetIndexesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexesResponse) GoString() string {
	return s.String()
}

func (s *GetIndexesResponse) SetHeaders(v map[string]*string) *GetIndexesResponse {
	s.Headers = v
	return s
}

func (s *GetIndexesResponse) SetBody(v *GetIndexesResponseBody) *GetIndexesResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 结果
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResult(v map[string]interface{}) *GetInstanceResponseBody {
	s.Result = v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceTableResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetInstanceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTableResponseBody) SetCode(v string) *GetInstanceTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceTableResponseBody) SetMessage(v string) *GetInstanceTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceTableResponseBody) SetRequestId(v string) *GetInstanceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTableResponseBody) SetResult(v map[string]interface{}) *GetInstanceTableResponseBody {
	s.Result = v
	return s
}

type GetInstanceTableResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTableResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTableResponse) SetHeaders(v map[string]*string) *GetInstanceTableResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTableResponse) SetBody(v *GetInstanceTableResponseBody) *GetInstanceTableResponse {
	s.Body = v
	return s
}

type ListBackFlowsRequest struct {
	// 查询回流次数(默认值为1,最大值为10)
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s ListBackFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBackFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListBackFlowsRequest) SetLimit(v int64) *ListBackFlowsRequest {
	s.Limit = &v
	return s
}

type ListBackFlowsResponseBody struct {
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// error message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListBackFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBackFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackFlowsResponseBody) SetCode(v string) *ListBackFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBackFlowsResponseBody) SetMessage(v string) *ListBackFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBackFlowsResponseBody) SetRequestId(v string) *ListBackFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackFlowsResponseBody) SetResult(v []map[string]interface{}) *ListBackFlowsResponseBody {
	s.Result = v
	return s
}

type ListBackFlowsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBackFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBackFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBackFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListBackFlowsResponse) SetHeaders(v map[string]*string) *ListBackFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListBackFlowsResponse) SetBody(v *ListBackFlowsResponseBody) *ListBackFlowsResponse {
	s.Body = v
	return s
}

type ListInstanceTableRequest struct {
	// 起始页
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页码大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstanceTableRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTableRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceTableRequest) SetPageNumber(v int64) *ListInstanceTableRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceTableRequest) SetPageSize(v int64) *ListInstanceTableRequest {
	s.PageSize = &v
	return s
}

type ListInstanceTableResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListInstanceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTableResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceTableResponseBody) SetCode(v string) *ListInstanceTableResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceTableResponseBody) SetMessage(v string) *ListInstanceTableResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceTableResponseBody) SetRequestId(v string) *ListInstanceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceTableResponseBody) SetResult(v map[string]interface{}) *ListInstanceTableResponseBody {
	s.Result = v
	return s
}

type ListInstanceTableResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTableResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceTableResponse) SetHeaders(v map[string]*string) *ListInstanceTableResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceTableResponse) SetBody(v *ListInstanceTableResponseBody) *ListInstanceTableResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// 起始页
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页码大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetPageNumber(v int64) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

type ListInstancesResponseBody struct {
	// 错误码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 结果数据
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetResult(v map[string]interface{}) *ListInstancesResponseBody {
	s.Result = v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type UpdateInstanceTableRequest struct {
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 字段描述
	FieldMap *UpdateInstanceTableRequestFieldMap `json:"fieldMap,omitempty" xml:"fieldMap,omitempty" type:"Struct"`
	// 全量数据源描述
	FullDataSourceConfig *UpdateInstanceTableRequestFullDataSourceConfig `json:"fullDataSourceConfig,omitempty" xml:"fullDataSourceConfig,omitempty" type:"Struct"`
	// 增量数据源描述
	IncDataSourceConfig *UpdateInstanceTableRequestIncDataSourceConfig `json:"incDataSourceConfig,omitempty" xml:"incDataSourceConfig,omitempty" type:"Struct"`
	// 索引类型
	IndexType *string `json:"indexType,omitempty" xml:"indexType,omitempty"`
	// 回流触发模式
	TriggerMode *string `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
	// Topic过期时间
	Ttl *int64 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s UpdateInstanceTableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequest) SetDescription(v string) *UpdateInstanceTableRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceTableRequest) SetFieldMap(v *UpdateInstanceTableRequestFieldMap) *UpdateInstanceTableRequest {
	s.FieldMap = v
	return s
}

func (s *UpdateInstanceTableRequest) SetFullDataSourceConfig(v *UpdateInstanceTableRequestFullDataSourceConfig) *UpdateInstanceTableRequest {
	s.FullDataSourceConfig = v
	return s
}

func (s *UpdateInstanceTableRequest) SetIncDataSourceConfig(v *UpdateInstanceTableRequestIncDataSourceConfig) *UpdateInstanceTableRequest {
	s.IncDataSourceConfig = v
	return s
}

func (s *UpdateInstanceTableRequest) SetIndexType(v string) *UpdateInstanceTableRequest {
	s.IndexType = &v
	return s
}

func (s *UpdateInstanceTableRequest) SetTriggerMode(v string) *UpdateInstanceTableRequest {
	s.TriggerMode = &v
	return s
}

func (s *UpdateInstanceTableRequest) SetTtl(v int64) *UpdateInstanceTableRequest {
	s.Ttl = &v
	return s
}

type UpdateInstanceTableRequestFieldMap struct {
	Fields []*UpdateInstanceTableRequestFieldMapFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// 索引信息
	Indexs          []*UpdateInstanceTableRequestFieldMapIndexs `json:"indexs,omitempty" xml:"indexs,omitempty" type:"Repeated"`
	IsAttributePack *bool                                       `json:"isAttributePack,omitempty" xml:"isAttributePack,omitempty"`
	// 最大外建数
	MaxSkeyNum *int64 `json:"maxSkeyNum,omitempty" xml:"maxSkeyNum,omitempty"`
	// 离线构建线程数
	OfflineBuildProtectionThreshold *int64 `json:"offlineBuildProtectionThreshold,omitempty" xml:"offlineBuildProtectionThreshold,omitempty"`
	// 主键
	Pkey *string `json:"pkey,omitempty" xml:"pkey,omitempty"`
	// 记录类型
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s UpdateInstanceTableRequestFieldMap) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestFieldMap) SetFields(v []*UpdateInstanceTableRequestFieldMapFields) *UpdateInstanceTableRequestFieldMap {
	s.Fields = v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetIndexs(v []*UpdateInstanceTableRequestFieldMapIndexs) *UpdateInstanceTableRequestFieldMap {
	s.Indexs = v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetIsAttributePack(v bool) *UpdateInstanceTableRequestFieldMap {
	s.IsAttributePack = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetMaxSkeyNum(v int64) *UpdateInstanceTableRequestFieldMap {
	s.MaxSkeyNum = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetOfflineBuildProtectionThreshold(v int64) *UpdateInstanceTableRequestFieldMap {
	s.OfflineBuildProtectionThreshold = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetPkey(v string) *UpdateInstanceTableRequestFieldMap {
	s.Pkey = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMap) SetRecordType(v string) *UpdateInstanceTableRequestFieldMap {
	s.RecordType = &v
	return s
}

type UpdateInstanceTableRequestFieldMapFields struct {
	// 默认值
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// 描述
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	IsMultiValue *bool   `json:"isMultiValue,omitempty" xml:"isMultiValue,omitempty"`
	IsVirtual    *bool   `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	OnlineStatus *bool   `json:"onlineStatus,omitempty" xml:"onlineStatus,omitempty"`
	// 是否为主键
	Pkey *bool `json:"pkey,omitempty" xml:"pkey,omitempty"`
	// 插件名
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	// 插件信息
	PluginParam *string `json:"pluginParam,omitempty" xml:"pluginParam,omitempty"`
	// 分隔符
	Seperator *string `json:"seperator,omitempty" xml:"seperator,omitempty"`
	// 源字段名
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// 源字段类型
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 目标字段名
	TargetName *string `json:"targetName,omitempty" xml:"targetName,omitempty"`
	// 目标字段类型
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s UpdateInstanceTableRequestFieldMapFields) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestFieldMapFields) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetDefaultValue(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.DefaultValue = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetDescription(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.Description = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetIsMultiValue(v bool) *UpdateInstanceTableRequestFieldMapFields {
	s.IsMultiValue = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetIsVirtual(v bool) *UpdateInstanceTableRequestFieldMapFields {
	s.IsVirtual = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetOnlineStatus(v bool) *UpdateInstanceTableRequestFieldMapFields {
	s.OnlineStatus = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetPkey(v bool) *UpdateInstanceTableRequestFieldMapFields {
	s.Pkey = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetPluginName(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.PluginName = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetPluginParam(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.PluginParam = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetSeperator(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.Seperator = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetSourceName(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.SourceName = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetSourceType(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.SourceType = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetTargetName(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.TargetName = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapFields) SetTargetType(v string) *UpdateInstanceTableRequestFieldMapFields {
	s.TargetType = &v
	return s
}

type UpdateInstanceTableRequestFieldMapIndexs struct {
	// 字段名
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// 索引字段名
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s UpdateInstanceTableRequestFieldMapIndexs) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestFieldMapIndexs) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestFieldMapIndexs) SetField(v string) *UpdateInstanceTableRequestFieldMapIndexs {
	s.Field = &v
	return s
}

func (s *UpdateInstanceTableRequestFieldMapIndexs) SetIndexName(v string) *UpdateInstanceTableRequestFieldMapIndexs {
	s.IndexName = &v
	return s
}

type UpdateInstanceTableRequestFullDataSourceConfig struct {
	// MaxCompute描述
	OdpsDataDesc *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc `json:"odpsDataDesc,omitempty" xml:"odpsDataDesc,omitempty" type:"Struct"`
	// 数据源类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateInstanceTableRequestFullDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestFullDataSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestFullDataSourceConfig) SetOdpsDataDesc(v *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) *UpdateInstanceTableRequestFullDataSourceConfig {
	s.OdpsDataDesc = v
	return s
}

func (s *UpdateInstanceTableRequestFullDataSourceConfig) SetType(v string) *UpdateInstanceTableRequestFullDataSourceConfig {
	s.Type = &v
	return s
}

type UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc struct {
	// 项目名
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// 表名
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) SetProject(v string) *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc {
	s.Project = &v
	return s
}

func (s *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc) SetTable(v string) *UpdateInstanceTableRequestFullDataSourceConfigOdpsDataDesc {
	s.Table = &v
	return s
}

type UpdateInstanceTableRequestIncDataSourceConfig struct {
	SwiftDataDesc *UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc `json:"swiftDataDesc,omitempty" xml:"swiftDataDesc,omitempty" type:"Struct"`
	// 数据源类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateInstanceTableRequestIncDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestIncDataSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestIncDataSourceConfig) SetSwiftDataDesc(v *UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) *UpdateInstanceTableRequestIncDataSourceConfig {
	s.SwiftDataDesc = v
	return s
}

func (s *UpdateInstanceTableRequestIncDataSourceConfig) SetType(v string) *UpdateInstanceTableRequestIncDataSourceConfig {
	s.Type = &v
	return s
}

type UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc struct {
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc) SetTopic(v string) *UpdateInstanceTableRequestIncDataSourceConfigSwiftDataDesc {
	s.Topic = &v
	return s
}

type UpdateInstanceTableResponseBody struct {
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInstanceTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableResponseBody) SetCode(v string) *UpdateInstanceTableResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceTableResponseBody) SetMessage(v string) *UpdateInstanceTableResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceTableResponseBody) SetRequestId(v string) *UpdateInstanceTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceTableResponseBody) SetResult(v map[string]interface{}) *UpdateInstanceTableResponseBody {
	s.Result = v
	return s
}

type UpdateInstanceTableResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceTableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceTableResponse) SetHeaders(v map[string]*string) *UpdateInstanceTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceTableResponse) SetBody(v *UpdateInstanceTableResponseBody) *UpdateInstanceTableResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("abfs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateBackFlow(tableName *string, instanceId *string, request *CreateBackFlowRequest) (_result *CreateBackFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBackFlowResponse{}
	_body, _err := client.CreateBackFlowWithOptions(tableName, instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackFlowWithOptions(tableName *string, instanceId *string, request *CreateBackFlowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBackFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tableName = openapiutil.GetEncodeParam(tableName)
	instanceId = openapiutil.GetEncodeParam(instanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OdpsPartition)) {
		body["odpsPartition"] = request.OdpsPartition
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackFlow"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName) + "/backflows"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBackFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceTable(instanceId *string, request *CreateInstanceTableRequest) (_result *CreateInstanceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceTableResponse{}
	_body, _err := client.CreateInstanceTableWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceTableWithOptions(instanceId *string, request *CreateInstanceTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FieldMap))) {
		body["fieldMap"] = request.FieldMap
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FullDataSourceConfig))) {
		body["fullDataSourceConfig"] = request.FullDataSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.IncDataSourceConfig))) {
		body["incDataSourceConfig"] = request.IncDataSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.IndexType)) {
		body["indexType"] = request.IndexType
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["tableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerMode)) {
		body["triggerMode"] = request.TriggerMode
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceTable"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceTable(tableName *string, instanceId *string) (_result *DeleteInstanceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceTableResponse{}
	_body, _err := client.DeleteInstanceTableWithOptions(tableName, instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceTableWithOptions(tableName *string, instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceTableResponse, _err error) {
	tableName = openapiutil.GetEncodeParam(tableName)
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceTable"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBackFlow(tableName *string, instanceId *string, id *string) (_result *GetBackFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBackFlowResponse{}
	_body, _err := client.GetBackFlowWithOptions(tableName, instanceId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBackFlowWithOptions(tableName *string, instanceId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBackFlowResponse, _err error) {
	tableName = openapiutil.GetEncodeParam(tableName)
	instanceId = openapiutil.GetEncodeParam(instanceId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBackFlow"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName) + "/backflows/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBackFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIndexes(instanceId *string) (_result *GetIndexesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexesResponse{}
	_body, _err := client.GetIndexesWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIndexesWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexesResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexes"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/indexes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId)),
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

func (client *Client) GetInstanceTable(instanceId *string, tableName *string) (_result *GetInstanceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceTableResponse{}
	_body, _err := client.GetInstanceTableWithOptions(instanceId, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceTableWithOptions(instanceId *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceTableResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	tableName = openapiutil.GetEncodeParam(tableName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceTable"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBackFlows(tableName *string, instanceId *string, request *ListBackFlowsRequest) (_result *ListBackFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBackFlowsResponse{}
	_body, _err := client.ListBackFlowsWithOptions(tableName, instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBackFlowsWithOptions(tableName *string, instanceId *string, request *ListBackFlowsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBackFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tableName = openapiutil.GetEncodeParam(tableName)
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBackFlows"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName) + "/backflows"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBackFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceTable(instanceId *string, request *ListInstanceTableRequest) (_result *ListInstanceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceTableResponse{}
	_body, _err := client.ListInstanceTableWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceTableWithOptions(instanceId *string, request *ListInstanceTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceTable"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-12-30"),
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

func (client *Client) UpdateInstanceTable(instanceId *string, tableName *string, request *UpdateInstanceTableRequest) (_result *UpdateInstanceTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceTableResponse{}
	_body, _err := client.UpdateInstanceTableWithOptions(instanceId, tableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceTableWithOptions(instanceId *string, tableName *string, request *UpdateInstanceTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	tableName = openapiutil.GetEncodeParam(tableName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FieldMap))) {
		body["fieldMap"] = request.FieldMap
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FullDataSourceConfig))) {
		body["fullDataSourceConfig"] = request.FullDataSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.IncDataSourceConfig))) {
		body["incDataSourceConfig"] = request.IncDataSourceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.IndexType)) {
		body["indexType"] = request.IndexType
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerMode)) {
		body["triggerMode"] = request.TriggerMode
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceTable"),
		Version:     tea.String("2021-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(instanceId) + "/tables/" + tea.StringValue(tableName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
