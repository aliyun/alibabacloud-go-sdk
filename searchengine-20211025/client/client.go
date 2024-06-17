// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ErrorResponse struct {
	// example:
	//
	// InternalServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// internal server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 90D6B8F5-FE97-4509-9AAB-367836C51818
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s ErrorResponse) GoString() string {
	return s.String()
}

func (s *ErrorResponse) SetCode(v string) *ErrorResponse {
	s.Code = &v
	return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
	s.Message = &v
	return s
}

func (s *ErrorResponse) SetRequestId(v string) *ErrorResponse {
	s.RequestId = &v
	return s
}

type ResultClusterValue struct {
	BuildParallelNum *int32 `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	MergeParallelNum *int32 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
}

func (s ResultClusterValue) String() string {
	return tea.Prettify(s)
}

func (s ResultClusterValue) GoString() string {
	return s.String()
}

func (s *ResultClusterValue) SetBuildParallelNum(v int32) *ResultClusterValue {
	s.BuildParallelNum = &v
	return s
}

func (s *ResultClusterValue) SetMergeParallelNum(v int32) *ResultClusterValue {
	s.MergeParallelNum = &v
	return s
}

type ResultValue struct {
	// example:
	//
	// true
	PauseAll *bool `json:"pauseAll,omitempty" xml:"pauseAll,omitempty"`
	// example:
	//
	// true
	PauseIndex *bool `json:"pauseIndex,omitempty" xml:"pauseIndex,omitempty"`
	// example:
	//
	// true
	PauseIndexBatch *bool `json:"pauseIndexBatch,omitempty" xml:"pauseIndexBatch,omitempty"`
	// example:
	//
	// true
	PauseBiz *bool `json:"pauseBiz,omitempty" xml:"pauseBiz,omitempty"`
	// example:
	//
	// true
	PauseRuntime *bool `json:"pauseRuntime,omitempty" xml:"pauseRuntime,omitempty"`
}

func (s ResultValue) String() string {
	return tea.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) SetPauseAll(v bool) *ResultValue {
	s.PauseAll = &v
	return s
}

func (s *ResultValue) SetPauseIndex(v bool) *ResultValue {
	s.PauseIndex = &v
	return s
}

func (s *ResultValue) SetPauseIndexBatch(v bool) *ResultValue {
	s.PauseIndexBatch = &v
	return s
}

func (s *ResultValue) SetPauseBiz(v bool) *ResultValue {
	s.PauseBiz = &v
	return s
}

func (s *ResultValue) SetPauseRuntime(v bool) *ResultValue {
	s.PauseRuntime = &v
	return s
}

type VariablesValue struct {
	// Specifies whether the variable cannot be modified.
	//
	// example:
	//
	// false
	DisableModify *bool `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	// Specifies whether the variable is modified.
	//
	// example:
	//
	// false
	IsModify *bool `json:"isModify,omitempty" xml:"isModify,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The description about the variable.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The value of the template.
	//
	// example:
	//
	// ""
	TemplateValue *string `json:"templateValue,omitempty" xml:"templateValue,omitempty"`
	// The type of the variable. Valid values:
	//
	// 	- NORMAL: a normal variable
	//
	// 	- FUNCTION: a function variable
	//
	// example:
	//
	// NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The function variable.
	FuncValue *VariablesValueFuncValue `json:"funcValue,omitempty" xml:"funcValue,omitempty" type:"Struct"`
}

func (s VariablesValue) String() string {
	return tea.Prettify(s)
}

func (s VariablesValue) GoString() string {
	return s.String()
}

func (s *VariablesValue) SetDisableModify(v bool) *VariablesValue {
	s.DisableModify = &v
	return s
}

func (s *VariablesValue) SetIsModify(v bool) *VariablesValue {
	s.IsModify = &v
	return s
}

func (s *VariablesValue) SetValue(v string) *VariablesValue {
	s.Value = &v
	return s
}

func (s *VariablesValue) SetDescription(v string) *VariablesValue {
	s.Description = &v
	return s
}

func (s *VariablesValue) SetTemplateValue(v string) *VariablesValue {
	s.TemplateValue = &v
	return s
}

func (s *VariablesValue) SetType(v string) *VariablesValue {
	s.Type = &v
	return s
}

func (s *VariablesValue) SetFuncValue(v *VariablesValueFuncValue) *VariablesValue {
	s.FuncValue = v
	return s
}

type VariablesValueFuncValue struct {
	// The class name.
	//
	// example:
	//
	// ""
	FuncClassName *string `json:"funcClassName,omitempty" xml:"funcClassName,omitempty"`
	// The template of the variable.
	//
	// example:
	//
	// ""
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s VariablesValueFuncValue) String() string {
	return tea.Prettify(s)
}

func (s VariablesValueFuncValue) GoString() string {
	return s.String()
}

func (s *VariablesValueFuncValue) SetFuncClassName(v string) *VariablesValueFuncValue {
	s.FuncClassName = &v
	return s
}

func (s *VariablesValueFuncValue) SetTemplate(v string) *VariablesValueFuncValue {
	s.Template = &v
	return s
}

type ConfigValue struct {
	Desc  *string             `json:"desc,omitempty" xml:"desc,omitempty"`
	Files []*ConfigValueFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s ConfigValue) String() string {
	return tea.Prettify(s)
}

func (s ConfigValue) GoString() string {
	return s.String()
}

func (s *ConfigValue) SetDesc(v string) *ConfigValue {
	s.Desc = &v
	return s
}

func (s *ConfigValue) SetFiles(v []*ConfigValueFiles) *ConfigValue {
	s.Files = v
	return s
}

type ConfigValueFiles struct {
	// example:
	//
	// UPDATE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// /
	ParentFullPath *string                 `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
	FileName       *string                 `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Config         *ConfigValueFilesConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// /test
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
}

func (s ConfigValueFiles) String() string {
	return tea.Prettify(s)
}

func (s ConfigValueFiles) GoString() string {
	return s.String()
}

func (s *ConfigValueFiles) SetOperateType(v string) *ConfigValueFiles {
	s.OperateType = &v
	return s
}

func (s *ConfigValueFiles) SetParentFullPath(v string) *ConfigValueFiles {
	s.ParentFullPath = &v
	return s
}

func (s *ConfigValueFiles) SetFileName(v string) *ConfigValueFiles {
	s.FileName = &v
	return s
}

func (s *ConfigValueFiles) SetConfig(v *ConfigValueFilesConfig) *ConfigValueFiles {
	s.Config = v
	return s
}

func (s *ConfigValueFiles) SetDirName(v string) *ConfigValueFiles {
	s.DirName = &v
	return s
}

type ConfigValueFilesConfig struct {
	// example:
	//
	// $dictContent
	Content   *string                                          `json:"content,omitempty" xml:"content,omitempty"`
	Variables map[string]*ConfigValueFilesConfigVariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ConfigValueFilesConfig) String() string {
	return tea.Prettify(s)
}

func (s ConfigValueFilesConfig) GoString() string {
	return s.String()
}

func (s *ConfigValueFilesConfig) SetContent(v string) *ConfigValueFilesConfig {
	s.Content = &v
	return s
}

func (s *ConfigValueFilesConfig) SetVariables(v map[string]*ConfigValueFilesConfigVariablesValue) *ConfigValueFilesConfig {
	s.Variables = v
	return s
}

type BodyValue struct {
	// example:
	//
	// true
	PauseAll *bool `json:"pauseAll,omitempty" xml:"pauseAll,omitempty"`
	// example:
	//
	// true
	PauseIndex *bool `json:"pauseIndex,omitempty" xml:"pauseIndex,omitempty"`
	// example:
	//
	// true
	PauseIndexBatch *bool `json:"pauseIndexBatch,omitempty" xml:"pauseIndexBatch,omitempty"`
	// example:
	//
	// true
	PauseBiz *bool `json:"pauseBiz,omitempty" xml:"pauseBiz,omitempty"`
	// example:
	//
	// true
	PauseRuntime *bool `json:"pauseRuntime,omitempty" xml:"pauseRuntime,omitempty"`
}

func (s BodyValue) String() string {
	return tea.Prettify(s)
}

func (s BodyValue) GoString() string {
	return s.String()
}

func (s *BodyValue) SetPauseAll(v bool) *BodyValue {
	s.PauseAll = &v
	return s
}

func (s *BodyValue) SetPauseIndex(v bool) *BodyValue {
	s.PauseIndex = &v
	return s
}

func (s *BodyValue) SetPauseIndexBatch(v bool) *BodyValue {
	s.PauseIndexBatch = &v
	return s
}

func (s *BodyValue) SetPauseBiz(v bool) *BodyValue {
	s.PauseBiz = &v
	return s
}

func (s *BodyValue) SetPauseRuntime(v bool) *BodyValue {
	s.PauseRuntime = &v
	return s
}

type FilesConfigVariablesValue struct {
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	DisableModify *bool   `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	IsModify      *bool   `json:"isModify,omitempty" xml:"isModify,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilesConfigVariablesValue) String() string {
	return tea.Prettify(s)
}

func (s FilesConfigVariablesValue) GoString() string {
	return s.String()
}

func (s *FilesConfigVariablesValue) SetDescription(v string) *FilesConfigVariablesValue {
	s.Description = &v
	return s
}

func (s *FilesConfigVariablesValue) SetDisableModify(v bool) *FilesConfigVariablesValue {
	s.DisableModify = &v
	return s
}

func (s *FilesConfigVariablesValue) SetIsModify(v bool) *FilesConfigVariablesValue {
	s.IsModify = &v
	return s
}

func (s *FilesConfigVariablesValue) SetType(v string) *FilesConfigVariablesValue {
	s.Type = &v
	return s
}

func (s *FilesConfigVariablesValue) SetValue(v string) *FilesConfigVariablesValue {
	s.Value = &v
	return s
}

type ConfigValueFilesConfigVariablesValue struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	DisableModify *bool `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	// example:
	//
	// false
	IsModify *bool `json:"isModify,omitempty" xml:"isModify,omitempty"`
	// example:
	//
	// NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfigValueFilesConfigVariablesValue) String() string {
	return tea.Prettify(s)
}

func (s ConfigValueFilesConfigVariablesValue) GoString() string {
	return s.String()
}

func (s *ConfigValueFilesConfigVariablesValue) SetDescription(v string) *ConfigValueFilesConfigVariablesValue {
	s.Description = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetDisableModify(v bool) *ConfigValueFilesConfigVariablesValue {
	s.DisableModify = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetIsModify(v bool) *ConfigValueFilesConfigVariablesValue {
	s.IsModify = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetType(v string) *ConfigValueFilesConfigVariablesValue {
	s.Type = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetValue(v string) *ConfigValueFilesConfigVariablesValue {
	s.Value = &v
	return s
}

type BuildIndexRequest struct {
	// The reindexing mode.
	//
	// example:
	//
	// indexRecover
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// my_data_source
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// swift
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// The timestamp in seconds. The value must be of the INTEGER type. This parameter is required if you specify an API data source.
	//
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center in which the data source resides.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The data restoration version.
	//
	// example:
	//
	// 160131146
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The data partition. This parameter is required if dataSourceType is set to odps.
	//
	// example:
	//
	// 20201010
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s BuildIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexRequest) GoString() string {
	return s.String()
}

func (s *BuildIndexRequest) SetBuildMode(v string) *BuildIndexRequest {
	s.BuildMode = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceName(v string) *BuildIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceType(v string) *BuildIndexRequest {
	s.DataSourceType = &v
	return s
}

func (s *BuildIndexRequest) SetDataTimeSec(v int32) *BuildIndexRequest {
	s.DataTimeSec = &v
	return s
}

func (s *BuildIndexRequest) SetDomain(v string) *BuildIndexRequest {
	s.Domain = &v
	return s
}

func (s *BuildIndexRequest) SetGeneration(v int64) *BuildIndexRequest {
	s.Generation = &v
	return s
}

func (s *BuildIndexRequest) SetPartition(v string) *BuildIndexRequest {
	s.Partition = &v
	return s
}

type BuildIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of clusters
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BuildIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexResponseBody) GoString() string {
	return s.String()
}

func (s *BuildIndexResponseBody) SetRequestId(v string) *BuildIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildIndexResponseBody) SetResult(v map[string]interface{}) *BuildIndexResponseBody {
	s.Result = v
	return s
}

type BuildIndexResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s BuildIndexResponse) GoString() string {
	return s.String()
}

func (s *BuildIndexResponse) SetHeaders(v map[string]*string) *BuildIndexResponse {
	s.Headers = v
	return s
}

func (s *BuildIndexResponse) SetStatusCode(v int32) *BuildIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildIndexResponse) SetBody(v *BuildIndexResponseBody) *BuildIndexResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// The remarks of the query node
	//
	// example:
	//
	// true
	AutoLoad *bool `json:"autoLoad,omitempty" xml:"autoLoad,omitempty"`
	// The description of the data node
	DataNode *CreateClusterRequestDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster
	//
	// example:
	//
	// "ha-tets"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the node
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The description of the query node
	QueryNode *CreateClusterRequestQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetAutoLoad(v bool) *CreateClusterRequest {
	s.AutoLoad = &v
	return s
}

func (s *CreateClusterRequest) SetDataNode(v *CreateClusterRequestDataNode) *CreateClusterRequest {
	s.DataNode = v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetQueryNode(v *CreateClusterRequestQueryNode) *CreateClusterRequest {
	s.QueryNode = v
	return s
}

type CreateClusterRequestDataNode struct {
	// The number of data nodes
	//
	// example:
	//
	// 2
	Number    *int32  `json:"number,omitempty" xml:"number,omitempty"`
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s CreateClusterRequestDataNode) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestDataNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestDataNode) SetNumber(v int32) *CreateClusterRequestDataNode {
	s.Number = &v
	return s
}

func (s *CreateClusterRequestDataNode) SetPartition(v string) *CreateClusterRequestDataNode {
	s.Partition = &v
	return s
}

type CreateClusterRequestQueryNode struct {
	// The number of nodes to query
	//
	// example:
	//
	// 2
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s CreateClusterRequestQueryNode) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestQueryNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestQueryNode) SetNumber(v int32) *CreateClusterRequestQueryNode {
	s.Number = &v
	return s
}

type CreateClusterResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetResult(v map[string]interface{}) *CreateClusterResponseBody {
	s.Result = v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateConfigDirRequest struct {
	// example:
	//
	// /clusters
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s CreateConfigDirRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDirRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigDirRequest) SetDirName(v string) *CreateConfigDirRequest {
	s.DirName = &v
	return s
}

func (s *CreateConfigDirRequest) SetParentFullPath(v string) *CreateConfigDirRequest {
	s.ParentFullPath = &v
	return s
}

type CreateConfigDirResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateConfigDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigDirResponseBody) SetRequestId(v string) *CreateConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigDirResponseBody) SetResult(v map[string]interface{}) *CreateConfigDirResponseBody {
	s.Result = v
	return s
}

type CreateConfigDirResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigDirResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDirResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigDirResponse) SetHeaders(v map[string]*string) *CreateConfigDirResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigDirResponse) SetStatusCode(v int32) *CreateConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigDirResponse) SetBody(v *CreateConfigDirResponseBody) *CreateConfigDirResponse {
	s.Body = v
	return s
}

type CreateConfigFileRequest struct {
	// example:
	//
	// /schemas/device_event_xt_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// oss://xxx/xxxx/xxx
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s CreateConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigFileRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigFileRequest) SetFileName(v string) *CreateConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateConfigFileRequest) SetOssPath(v string) *CreateConfigFileRequest {
	s.OssPath = &v
	return s
}

func (s *CreateConfigFileRequest) SetParentFullPath(v string) *CreateConfigFileRequest {
	s.ParentFullPath = &v
	return s
}

type CreateConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigFileResponseBody) SetRequestId(v string) *CreateConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigFileResponseBody) SetResult(v map[string]interface{}) *CreateConfigFileResponseBody {
	s.Result = v
	return s
}

type CreateConfigFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigFileResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigFileResponse) SetHeaders(v map[string]*string) *CreateConfigFileResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigFileResponse) SetStatusCode(v int32) *CreateConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigFileResponse) SetBody(v *CreateConfigFileResponseBody) *CreateConfigFileResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	AutoBuildIndex *bool                              `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *CreateDataSourceRequestConfig     `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	Domain         *string                            `json:"domain,omitempty" xml:"domain,omitempty"`
	Name           *string                            `json:"name,omitempty" xml:"name,omitempty"`
	SaroConfig     *CreateDataSourceRequestSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	Type           *string                            `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetAutoBuildIndex(v bool) *CreateDataSourceRequest {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateDataSourceRequest) SetConfig(v *CreateDataSourceRequestConfig) *CreateDataSourceRequest {
	s.Config = v
	return s
}

func (s *CreateDataSourceRequest) SetDomain(v string) *CreateDataSourceRequest {
	s.Domain = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetSaroConfig(v *CreateDataSourceRequestSaroConfig) *CreateDataSourceRequest {
	s.SaroConfig = v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequest) SetDryRun(v bool) *CreateDataSourceRequest {
	s.DryRun = &v
	return s
}

type CreateDataSourceRequestConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Namespace    *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OssPath      *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition    *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Path         *string `json:"path,omitempty" xml:"path,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	Table        *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateDataSourceRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestConfig) SetAccessKey(v string) *CreateDataSourceRequestConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetAccessSecret(v string) *CreateDataSourceRequestConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetBucket(v string) *CreateDataSourceRequestConfig {
	s.Bucket = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetEndpoint(v string) *CreateDataSourceRequestConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetNamespace(v string) *CreateDataSourceRequestConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetOssPath(v string) *CreateDataSourceRequestConfig {
	s.OssPath = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPartition(v string) *CreateDataSourceRequestConfig {
	s.Partition = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPath(v string) *CreateDataSourceRequestConfig {
	s.Path = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetProject(v string) *CreateDataSourceRequestConfig {
	s.Project = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetTable(v string) *CreateDataSourceRequestConfig {
	s.Table = &v
	return s
}

type CreateDataSourceRequestSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s CreateDataSourceRequestSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestSaroConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestSaroConfig) SetNamespace(v string) *CreateDataSourceRequestSaroConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestSaroConfig) SetTableName(v string) *CreateDataSourceRequestSaroConfig {
	s.TableName = &v
	return s
}

type CreateDataSourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetResult(v map[string]interface{}) *CreateDataSourceResponseBody {
	s.Result = v
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

type CreateIndexRequest struct {
	BuildParallelNum *int32 `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	// The content of the index.
	//
	// example:
	//
	// {\\"summarys\\":{\\"summary_fields\\":[\\"id\\"]},\\"indexs\\":[{\\"index_name\\":\\"index_id\\",\\"index_type\\":\\"PRIMARYKEY64\\",\\"index_fields\\":\\"id\\",\\"has_primary_key_attribute\\":true,\\"is_primary_key_sorted\\":false}],\\"attributes\\":[\\"id\\"],\\"fields\\":[{\\"field_name\\":\\"id\\",\\"field_type\\":\\"UINT16\\"}],\\"table_name\\":\\"index_2\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Optional. The data source, which can be MaxCompute, Message Service (MNS), Realtime Compute for Apache Flink, or StreamCompute.
	//
	// example:
	//
	// test1
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source, which is required for the new version of OpenSearch Vector Search Edition.
	DataSourceInfo *CreateIndexRequestDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The data center in which the data source resides.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The extended configurations of the field. Keys such as vector and embedding are included. Vector indicates the vector field. Embedding indicates the field that requires embedding.
	//
	// example:
	//
	// {
	//
	//         "vector":
	//
	//         [
	//
	//             "source_image_vector"
	//
	//         ],
	//
	//         "embeding":
	//
	//         [
	//
	//             "source_image"
	//
	//         ],
	//
	//         "description":
	//
	//         []
	//
	//     }
	Extend           map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	MergeParallelNum *int32                 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data partition.
	//
	// example:
	//
	// 20211202
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetBuildParallelNum(v int32) *CreateIndexRequest {
	s.BuildParallelNum = &v
	return s
}

func (s *CreateIndexRequest) SetContent(v string) *CreateIndexRequest {
	s.Content = &v
	return s
}

func (s *CreateIndexRequest) SetDataSource(v string) *CreateIndexRequest {
	s.DataSource = &v
	return s
}

func (s *CreateIndexRequest) SetDataSourceInfo(v *CreateIndexRequestDataSourceInfo) *CreateIndexRequest {
	s.DataSourceInfo = v
	return s
}

func (s *CreateIndexRequest) SetDomain(v string) *CreateIndexRequest {
	s.Domain = &v
	return s
}

func (s *CreateIndexRequest) SetExtend(v map[string]interface{}) *CreateIndexRequest {
	s.Extend = v
	return s
}

func (s *CreateIndexRequest) SetMergeParallelNum(v int32) *CreateIndexRequest {
	s.MergeParallelNum = &v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetPartition(v int32) *CreateIndexRequest {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequest) SetDryRun(v bool) *CreateIndexRequest {
	s.DryRun = &v
	return s
}

type CreateIndexRequestDataSourceInfo struct {
	// Specifies whether to enable the automatic full indexing feature.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The information about the MaxCompute data source.
	Config *CreateIndexRequestDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1709715164
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center where the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessParallelNum *int32  `json:"processParallelNum,omitempty" xml:"processParallelNum,omitempty"`
	// The number of resources used for data update.
	//
	// example:
	//
	// 4
	ProcessPartitionCount *int32                                      `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	SaroConfig            *CreateIndexRequestDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, and oss.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateIndexRequestDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfo) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfo) SetAutoBuildIndex(v bool) *CreateIndexRequestDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetConfig(v *CreateIndexRequestDataSourceInfoConfig) *CreateIndexRequestDataSourceInfo {
	s.Config = v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetDataTimeSec(v int32) *CreateIndexRequestDataSourceInfo {
	s.DataTimeSec = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetDomain(v string) *CreateIndexRequestDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetName(v string) *CreateIndexRequestDataSourceInfo {
	s.Name = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetProcessParallelNum(v int32) *CreateIndexRequestDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetProcessPartitionCount(v int32) *CreateIndexRequestDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetSaroConfig(v *CreateIndexRequestDataSourceInfoSaroConfig) *CreateIndexRequestDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetType(v string) *CreateIndexRequestDataSourceInfo {
	s.Type = &v
	return s
}

type CreateIndexRequestDataSourceInfoConfig struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// L***p
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// 5**9a6
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// oss bucket
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The endpoint of the MaxCompute or Object Storage Service (OSS) data source.
	//
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// test-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// /opensearch/oss.json
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The data partition. This parameter is required if type is set to odps.
	//
	// example:
	//
	// ds=20230114
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// test-hdfs-path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// bbt_algo_pai
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// bbt_rec_swing_u2i2i_score_be_v1
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateIndexRequestDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessKey(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessSecret(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetBucket(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetEndpoint(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetNamespace(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetOssPath(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetPartition(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetPath(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetProject(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetTable(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Table = &v
	return s
}

type CreateIndexRequestDataSourceInfoSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s CreateIndexRequestDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) SetNamespace(v string) *CreateIndexRequestDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) SetTableName(v string) *CreateIndexRequestDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type CreateIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetResult(v map[string]interface{}) *CreateIndexResponseBody {
	s.Result = v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexResponse) SetBody(v *CreateIndexResponseBody) *CreateIndexResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// The billing method of the instance. Valid values: PREPAY and POSTPAY. PREPAY indicates the instance is a subscription instance. When you set this parameter to PREPAY, make sure that your Alibaba Cloud account supports balance payment or credit card payment. Otherwise, the system returns the InvalidPayMethod error message. If you set this parameter to PREPAY, you must also specify the paymentInfo parameter. POSTPAY indicates that the instance is a pay-as-you-go instance. This billing method is not supported.
	//
	// example:
	//
	// ""
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// A list of instance-related specifications.
	Components []*CreateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The billing information.
	Order *CreateInstanceRequestOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateInstanceRequest) SetOrder(v *CreateInstanceRequestOrder) *CreateInstanceRequest {
	s.Order = v
	return s
}

type CreateInstanceRequestComponents struct {
	// The specification code, which must be consistent with values of the corresponding module parameters.
	//
	// example:
	//
	// ""
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Values that you specify for the corresponding module components on the buy page.
	//
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponents) SetCode(v string) *CreateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetValue(v string) *CreateInstanceRequestComponents {
	s.Value = &v
	return s
}

type CreateInstanceRequestOrder struct {
	// Specifies whether to enable auto-renewal. Valid values: true and false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The billing duration. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, and 12.
	//
	// example:
	//
	// 29
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The unit of the billing duration. Valid values: Month and Year.
	//
	// example:
	//
	// ""
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s CreateInstanceRequestOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestOrder) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestOrder) SetAutoRenew(v bool) *CreateInstanceRequestOrder {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetDuration(v int64) *CreateInstanceRequestOrder {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetPricingCycle(v string) *CreateInstanceRequestOrder {
	s.PricingCycle = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result *CreateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

type CreateInstanceResponseBodyResult struct {
	// The ID of the instance
	//
	// example:
	//
	// ha-cn-2r42ppr7901
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
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

type CreatePublicUrlResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreatePublicUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicUrlResponseBody) SetRequestId(v string) *CreatePublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublicUrlResponseBody) SetResult(v map[string]interface{}) *CreatePublicUrlResponseBody {
	s.Result = v
	return s
}

type CreatePublicUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicUrlResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicUrlResponse) SetHeaders(v map[string]*string) *CreatePublicUrlResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicUrlResponse) SetStatusCode(v int32) *CreatePublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicUrlResponse) SetBody(v *CreatePublicUrlResponseBody) *CreatePublicUrlResponse {
	s.Body = v
	return s
}

type CreateTableRequest struct {
	DataProcessConfig []*CreateTableRequestDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DataProcessorCount *int32                        `json:"dataProcessorCount,omitempty" xml:"dataProcessorCount,omitempty"`
	DataSource         *CreateTableRequestDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	FieldSchema        map[string]*string            `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// example:
	//
	// index_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// example:
	//
	// {}
	RawSchema   *string                          `json:"rawSchema,omitempty" xml:"rawSchema,omitempty"`
	VectorIndex []*CreateTableRequestVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateTableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) SetDataProcessConfig(v []*CreateTableRequestDataProcessConfig) *CreateTableRequest {
	s.DataProcessConfig = v
	return s
}

func (s *CreateTableRequest) SetDataProcessorCount(v int32) *CreateTableRequest {
	s.DataProcessorCount = &v
	return s
}

func (s *CreateTableRequest) SetDataSource(v *CreateTableRequestDataSource) *CreateTableRequest {
	s.DataSource = v
	return s
}

func (s *CreateTableRequest) SetFieldSchema(v map[string]*string) *CreateTableRequest {
	s.FieldSchema = v
	return s
}

func (s *CreateTableRequest) SetName(v string) *CreateTableRequest {
	s.Name = &v
	return s
}

func (s *CreateTableRequest) SetPartitionCount(v int32) *CreateTableRequest {
	s.PartitionCount = &v
	return s
}

func (s *CreateTableRequest) SetPrimaryKey(v string) *CreateTableRequest {
	s.PrimaryKey = &v
	return s
}

func (s *CreateTableRequest) SetRawSchema(v string) *CreateTableRequest {
	s.RawSchema = &v
	return s
}

func (s *CreateTableRequest) SetVectorIndex(v []*CreateTableRequestVectorIndex) *CreateTableRequest {
	s.VectorIndex = v
	return s
}

func (s *CreateTableRequest) SetDryRun(v bool) *CreateTableRequest {
	s.DryRun = &v
	return s
}

type CreateTableRequestDataProcessConfig struct {
	// example:
	//
	// source_image_vector
	DstField *string `json:"dstField,omitempty" xml:"dstField,omitempty"`
	// example:
	//
	// vectorize
	Operator *string                                    `json:"operator,omitempty" xml:"operator,omitempty"`
	Params   *CreateTableRequestDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s CreateTableRequestDataProcessConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestDataProcessConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfig) SetDstField(v string) *CreateTableRequestDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetOperator(v string) *CreateTableRequestDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetParams(v *CreateTableRequestDataProcessConfigParams) *CreateTableRequestDataProcessConfig {
	s.Params = v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetSrcField(v string) *CreateTableRequestDataProcessConfig {
	s.SrcField = &v
	return s
}

type CreateTableRequestDataProcessConfigParams struct {
	SrcFieldConfig *CreateTableRequestDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
	// example:
	//
	// image
	VectorModal *string `json:"vectorModal,omitempty" xml:"vectorModal,omitempty"`
	// example:
	//
	// clip
	VectorModel *string `json:"vectorModel,omitempty" xml:"vectorModel,omitempty"`
}

func (s CreateTableRequestDataProcessConfigParams) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfigParams) SetSrcFieldConfig(v *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) *CreateTableRequestDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *CreateTableRequestDataProcessConfigParams) SetVectorModal(v string) *CreateTableRequestDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParams) SetVectorModel(v string) *CreateTableRequestDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

type CreateTableRequestDataProcessConfigParamsSrcFieldConfig struct {
	// OSSBucket
	//
	// example:
	//
	// test
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// example:
	//
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateTableRequestDataProcessConfigParamsSrcFieldConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

type CreateTableRequestDataSource struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                               `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *CreateTableRequestDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTableRequestDataSource) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataSource) SetAutoBuildIndex(v bool) *CreateTableRequestDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateTableRequestDataSource) SetConfig(v *CreateTableRequestDataSourceConfig) *CreateTableRequestDataSource {
	s.Config = v
	return s
}

func (s *CreateTableRequestDataSource) SetDataTimeSec(v int32) *CreateTableRequestDataSource {
	s.DataTimeSec = &v
	return s
}

func (s *CreateTableRequestDataSource) SetType(v string) *CreateTableRequestDataSource {
	s.Type = &v
	return s
}

type CreateTableRequestDataSourceConfig struct {
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// oss bucket
	//
	// example:
	//
	// antsys-flytest-ci
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// project_20210220122847_3218
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// test56
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateTableRequestDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestDataSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataSourceConfig) SetAccessKey(v string) *CreateTableRequestDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetAccessSecret(v string) *CreateTableRequestDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetBucket(v string) *CreateTableRequestDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetEndpoint(v string) *CreateTableRequestDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetOssPath(v string) *CreateTableRequestDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetPartition(v string) *CreateTableRequestDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetProject(v string) *CreateTableRequestDataSourceConfig {
	s.Project = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetTable(v string) *CreateTableRequestDataSourceConfig {
	s.Table = &v
	return s
}

type CreateTableRequestVectorIndex struct {
	AdvanceParams *CreateTableRequestVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
	// example:
	//
	// 128
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// SquaredEuclidean
	DistanceType *string `json:"distanceType,omitempty" xml:"distanceType,omitempty"`
	// example:
	//
	// case_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// sparse_indices
	SparseIndexField *string `json:"sparseIndexField,omitempty" xml:"sparseIndexField,omitempty"`
	// example:
	//
	// sparse_values
	SparseValueField *string `json:"sparseValueField,omitempty" xml:"sparseValueField,omitempty"`
	// example:
	//
	// source_image_vector
	VectorField *string `json:"vectorField,omitempty" xml:"vectorField,omitempty"`
	// example:
	//
	// Qc
	VectorIndexType *string `json:"vectorIndexType,omitempty" xml:"vectorIndexType,omitempty"`
}

func (s CreateTableRequestVectorIndex) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestVectorIndex) GoString() string {
	return s.String()
}

func (s *CreateTableRequestVectorIndex) SetAdvanceParams(v *CreateTableRequestVectorIndexAdvanceParams) *CreateTableRequestVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *CreateTableRequestVectorIndex) SetDimension(v string) *CreateTableRequestVectorIndex {
	s.Dimension = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetDistanceType(v string) *CreateTableRequestVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetIndexName(v string) *CreateTableRequestVectorIndex {
	s.IndexName = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetNamespace(v string) *CreateTableRequestVectorIndex {
	s.Namespace = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetSparseIndexField(v string) *CreateTableRequestVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetSparseValueField(v string) *CreateTableRequestVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetVectorField(v string) *CreateTableRequestVectorIndex {
	s.VectorField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetVectorIndexType(v string) *CreateTableRequestVectorIndex {
	s.VectorIndexType = &v
	return s
}

type CreateTableRequestVectorIndexAdvanceParams struct {
	// example:
	//
	// {}
	BuildIndexParams *string `json:"buildIndexParams,omitempty" xml:"buildIndexParams,omitempty"`
	// example:
	//
	// 5000
	LinearBuildThreshold *string `json:"linearBuildThreshold,omitempty" xml:"linearBuildThreshold,omitempty"`
	// example:
	//
	// 20000
	MinScanDocCnt *string `json:"minScanDocCnt,omitempty" xml:"minScanDocCnt,omitempty"`
	// example:
	//
	// {}
	SearchIndexParams *string `json:"searchIndexParams,omitempty" xml:"searchIndexParams,omitempty"`
}

func (s CreateTableRequestVectorIndexAdvanceParams) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequestVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetBuildIndexParams(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetSearchIndexParams(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
	return s
}

type CreateTableResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBody) SetRequestId(v string) *CreateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableResponseBody) SetResult(v map[string]interface{}) *CreateTableResponseBody {
	s.Result = v
	return s
}

type CreateTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTableResponse) SetHeaders(v map[string]*string) *CreateTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTableResponse) SetStatusCode(v int32) *CreateTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableResponse) SetBody(v *CreateTableResponseBody) *CreateTableResponse {
	s.Body = v
	return s
}

type DeleteAdvanceConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponseBody) SetRequestId(v string) *DeleteAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *DeleteAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type DeleteAdvanceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponse) SetHeaders(v map[string]*string) *DeleteAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetStatusCode(v int32) *DeleteAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetBody(v *DeleteAdvanceConfigResponseBody) *DeleteAdvanceConfigResponse {
	s.Body = v
	return s
}

type DeleteConfigDirRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /clusters
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s DeleteConfigDirRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigDirRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirRequest) SetDirName(v string) *DeleteConfigDirRequest {
	s.DirName = &v
	return s
}

func (s *DeleteConfigDirRequest) SetParentFullPath(v string) *DeleteConfigDirRequest {
	s.ParentFullPath = &v
	return s
}

type DeleteConfigDirResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteConfigDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirResponseBody) SetRequestId(v string) *DeleteConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigDirResponseBody) SetResult(v map[string]interface{}) *DeleteConfigDirResponseBody {
	s.Result = v
	return s
}

type DeleteConfigDirResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigDirResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigDirResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirResponse) SetHeaders(v map[string]*string) *DeleteConfigDirResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigDirResponse) SetStatusCode(v int32) *DeleteConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigDirResponse) SetBody(v *DeleteConfigDirResponseBody) *DeleteConfigDirResponse {
	s.Body = v
	return s
}

type DeleteConfigFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /schemas/automobile_vector_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s DeleteConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileRequest) SetFileName(v string) *DeleteConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *DeleteConfigFileRequest) SetParentFullPath(v string) *DeleteConfigFileRequest {
	s.ParentFullPath = &v
	return s
}

type DeleteConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileResponseBody) SetRequestId(v string) *DeleteConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigFileResponseBody) SetResult(v map[string]interface{}) *DeleteConfigFileResponseBody {
	s.Result = v
	return s
}

type DeleteConfigFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileResponse) SetHeaders(v map[string]*string) *DeleteConfigFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigFileResponse) SetStatusCode(v int32) *DeleteConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigFileResponse) SetBody(v *DeleteConfigFileResponseBody) *DeleteConfigFileResponse {
	s.Body = v
	return s
}

type DeleteDataSourceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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

func (s *DeleteDataSourceResponseBody) SetResult(v map[string]interface{}) *DeleteDataSourceResponseBody {
	s.Result = v
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

type DeleteIndexRequest struct {
	// The data source
	//
	// This parameter is required.
	//
	// example:
	//
	// ha-cn-pl32rf0js04_test
	DataSource       *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	DeleteDataSource *bool   `json:"deleteDataSource,omitempty" xml:"deleteDataSource,omitempty"`
}

func (s DeleteIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) SetDataSource(v string) *DeleteIndexRequest {
	s.DataSource = &v
	return s
}

func (s *DeleteIndexRequest) SetDeleteDataSource(v bool) *DeleteIndexRequest {
	s.DeleteDataSource = &v
	return s
}

type DeleteIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetResult(v map[string]interface{}) *DeleteIndexResponseBody {
	s.Result = v
	return s
}

type DeleteIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexResponse) SetBody(v *DeleteIndexResponseBody) *DeleteIndexResponse {
	s.Body = v
	return s
}

type DeleteIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponseBody) SetRequestId(v string) *DeleteIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexVersionResponseBody) SetResult(v map[string]interface{}) *DeleteIndexVersionResponseBody {
	s.Result = v
	return s
}

type DeleteIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexVersionResponse) SetHeaders(v map[string]*string) *DeleteIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexVersionResponse) SetStatusCode(v int32) *DeleteIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexVersionResponse) SetBody(v *DeleteIndexVersionResponseBody) *DeleteIndexVersionResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetResult(v map[string]interface{}) *DeleteInstanceResponseBody {
	s.Result = v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeletePublicUrlResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeletePublicUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicUrlResponseBody) SetRequestId(v string) *DeletePublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePublicUrlResponseBody) SetResult(v map[string]interface{}) *DeletePublicUrlResponseBody {
	s.Result = v
	return s
}

type DeletePublicUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicUrlResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicUrlResponse) SetHeaders(v map[string]*string) *DeletePublicUrlResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicUrlResponse) SetStatusCode(v int32) *DeletePublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicUrlResponse) SetBody(v *DeletePublicUrlResponseBody) *DeletePublicUrlResponse {
	s.Body = v
	return s
}

type DeleteTableResponseBody struct {
	// requestId
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBody) SetRequestId(v string) *DeleteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableResponseBody) SetResult(v map[string]interface{}) *DeleteTableResponseBody {
	s.Result = v
	return s
}

type DeleteTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableResponse) SetHeaders(v map[string]*string) *DeleteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableResponse) SetStatusCode(v int32) *DeleteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableResponse) SetBody(v *DeleteTableResponseBody) *DeleteTableResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// example:
	//
	// zh-cn
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
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

type DescribeRegionsResponseBody struct {
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*DescribeRegionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

type DescribeRegionsResponseBodyResult struct {
	// example:
	//
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) SetEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
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

type ForceSwitchResponseBody struct {
	// id of request
	//
	// example:
	//
	// 0B1FF998-BB8D-5182-BFC0-E471AA77095A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ForceSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponseBody) SetRequestId(v string) *ForceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceSwitchResponseBody) SetResult(v map[string]interface{}) *ForceSwitchResponseBody {
	s.Result = v
	return s
}

type ForceSwitchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForceSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceSwitchResponse) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponse) SetHeaders(v map[string]*string) *ForceSwitchResponse {
	s.Headers = v
	return s
}

func (s *ForceSwitchResponse) SetStatusCode(v int32) *ForceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceSwitchResponse) SetBody(v *ForceSwitchResponseBody) *ForceSwitchResponse {
	s.Body = v
	return s
}

type GetAdvanceConfigRequest struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAdvanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigRequest) SetType(v string) *GetAdvanceConfigRequest {
	s.Type = &v
	return s
}

type GetAdvanceConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result *GetAdvanceConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBody) SetRequestId(v string) *GetAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigResponseBody) SetResult(v *GetAdvanceConfigResponseBodyResult) *GetAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type GetAdvanceConfigResponseBodyResult struct {
	// The content of the configuration that is returned.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The type of the configuration content. Valid values: FILE, GIT, HTTP, and ODPS.
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The description.
	//
	// example:
	//
	// close alarm, chiji id 37080
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about files.
	Files []*GetAdvanceConfigResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// ha-cn-0ju2s170b03_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status.
	//
	// example:
	//
	// 0,1,3,6,8
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// ""
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResult) SetContent(v string) *GetAdvanceConfigResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetContentType(v string) *GetAdvanceConfigResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetDesc(v string) *GetAdvanceConfigResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetFiles(v []*GetAdvanceConfigResponseBodyResultFiles) *GetAdvanceConfigResponseBodyResult {
	s.Files = v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetName(v string) *GetAdvanceConfigResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetStatus(v string) *GetAdvanceConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetUpdateTime(v int64) *GetAdvanceConfigResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type GetAdvanceConfigResponseBodyResultFiles struct {
	// The name of the file path.
	//
	// example:
	//
	// ""
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether it is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether it is a template.
	//
	// example:
	//
	// True
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResultFiles) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetFullPathName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsDir(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsTemplate(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.Name = &v
	return s
}

type GetAdvanceConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigResponse) SetStatusCode(v int32) *GetAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigResponse) SetBody(v *GetAdvanceConfigResponseBody) *GetAdvanceConfigResponse {
	s.Body = v
	return s
}

type GetAdvanceConfigFileRequest struct {
	// The name of the file
	//
	// This parameter is required.
	//
	// example:
	//
	// /intervene_dict/chn_ecommerce_general.dict
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetAdvanceConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileRequest) SetFileName(v string) *GetAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

type GetAdvanceConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	Result *GetAdvanceConfigFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBody) SetRequestId(v string) *GetAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigFileResponseBody) SetResult(v *GetAdvanceConfigFileResponseBodyResult) *GetAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

type GetAdvanceConfigFileResponseBodyResult struct {
	// The content of the file
	//
	// example:
	//
	// {"summarys":{"parameter":{"file_compressor":"zstd"},"summary_fields":["id"]},"file_compress":[{"name":"file_compressor","type":"zstd"},{"name":"no_compressor","type":""}],"indexs":[{"index_fields":"name","index_name":"ids","index_type":"STRING"},{"has_primary_key_attribute":true,"index_fields":"id","is_primary_key_sorted":false,"index_name":"id","index_type":"PRIMARYKEY64"}],"attributes":[{"file_compress":"no_compressor","field_name":"id"}],"fields":[{"user_defined_param":{},"compress_type":"uniq","field_type":"STRING","field_name":"id"},{"compress_type":"uniq","field_type":"STRING","field_name":"name"}],"table_name":"api"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetAdvanceConfigFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponseBodyResult) SetContent(v string) *GetAdvanceConfigFileResponseBodyResult {
	s.Content = &v
	return s
}

type GetAdvanceConfigFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvanceConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *GetAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetStatusCode(v int32) *GetAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvanceConfigFileResponse) SetBody(v *GetAdvanceConfigFileResponseBody) *GetAdvanceConfigFileResponse {
	s.Body = v
	return s
}

type GetClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of the cluster details.
	Result *GetClusterResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetResult(v *GetClusterResponseBodyResult) *GetClusterResponseBody {
	s.Result = v
	return s
}

type GetClusterResponseBodyResult struct {
	Config map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2021-08-09 00:01:02
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	CreateTime       *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The effective advanced configuration version.
	//
	// example:
	//
	// test_yyds_data1
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// The effective online configuration version.
	//
	// example:
	//
	// test_yyds_data1
	CurrentOnlineConfigVersion *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// The specifications of the data node.
	DataNode *GetClusterResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	//
	// example:
	//
	// fzz_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest advanced configuration version.
	//
	// example:
	//
	// test_yyds_data1
	LatestAdvanceConfigVersion *string `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// The latest online configuration version.
	//
	// example:
	//
	// test_yyds_data1
	LatestOnlineConfigVersion *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The specifications of the query node.
	QueryNode *GetClusterResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	// The creation status of the cluster. Valid values: NEW and PUBLISH. NEW indicates that the cluster is being created. PUBLISH indicates that the cluster is created.
	//
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetClusterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *GetClusterResponseBodyResult {
	s.Config = v
	return s
}

func (s *GetClusterResponseBodyResult) SetConfigUpdateTime(v string) *GetClusterResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCreateTime(v string) *GetClusterResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetDataNode(v *GetClusterResponseBodyResultDataNode) *GetClusterResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetDescription(v string) *GetClusterResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetLatestOnlineConfigVersion(v string) *GetClusterResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetName(v string) *GetClusterResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResult) SetQueryNode(v *GetClusterResponseBodyResultQueryNode) *GetClusterResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *GetClusterResponseBodyResult) SetStatus(v string) *GetClusterResponseBodyResult {
	s.Status = &v
	return s
}

type GetClusterResponseBodyResultDataNode struct {
	// The name of the node.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultDataNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultDataNode) SetName(v string) *GetClusterResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetNumber(v int32) *GetClusterResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultDataNode) SetPartition(v int32) *GetClusterResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

type GetClusterResponseBodyResultQueryNode struct {
	// The name of the node.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetClusterResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyResultQueryNode) SetName(v string) *GetClusterResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetNumber(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *GetClusterResponseBodyResultQueryNode) SetPartition(v int32) *GetClusterResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetClusterRunTimeInfoResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result set.
	Result []*GetClusterRunTimeInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetClusterRunTimeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBody) SetRequestId(v string) *GetClusterRunTimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBody) SetResult(v []*GetClusterRunTimeInfoResponseBodyResult) *GetClusterRunTimeInfoResponseBody {
	s.Result = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResult struct {
	// The name of the cluster.
	//
	// example:
	//
	// vpc_hz_domain_1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The information about the data node.
	DataNodes []*GetClusterRunTimeInfoResponseBodyResultDataNodes `json:"dataNodes,omitempty" xml:"dataNodes,omitempty" type:"Repeated"`
	// The information about the query node.
	QueryNode *GetClusterRunTimeInfoResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetClusterName(v string) *GetClusterRunTimeInfoResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetDataNodes(v []*GetClusterRunTimeInfoResponseBodyResultDataNodes) *GetClusterRunTimeInfoResponseBodyResult {
	s.DataNodes = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResult) SetQueryNode(v *GetClusterRunTimeInfoResponseBodyResultQueryNode) *GetClusterRunTimeInfoResponseBodyResult {
	s.QueryNode = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodes struct {
	// The configuration status list.
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	// The dataStatusList.
	DataStatusList []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList `json:"dataStatusList,omitempty" xml:"dataStatusList,omitempty" type:"Repeated"`
	// The service status.
	ServiceStatus *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodes) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetDataStatusList(v []*GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.DataStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodes) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) *GetClusterRunTimeInfoResponseBodyResultDataNodes {
	s.ServiceStatus = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList struct {
	// The time when the cluster was updated.
	//
	// example:
	//
	// ""
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The overall progress.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of nodes that are configured.
	//
	// example:
	//
	// 1
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test_0704
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of nodes that you specify when you create the cluster.
	//
	// example:
	//
	// 0
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesConfigStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList struct {
	// The information about advanced configurations.
	AdvanceConfigInfo *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo `json:"advanceConfigInfo,omitempty" xml:"advanceConfigInfo,omitempty" type:"Struct"`
	// The name of the worker that failed because of a deployment failure.
	DeployFailedWorker []*string `json:"deployFailedWorker,omitempty" xml:"deployFailedWorker,omitempty" type:"Repeated"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 2
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The overall progress.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of nodes that are configured.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0A3B1C48006A6C0905F6375F4821EB50
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The time when the full data was updated.
	//
	// example:
	//
	// " "
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The full version.
	//
	// example:
	//
	// 123423
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The time when the incremental data was updated.
	//
	// example:
	//
	// ""
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The incremental version.
	//
	// example:
	//
	// 123423
	IncVersion *int64 `json:"incVersion,omitempty" xml:"incVersion,omitempty"`
	// The configuration information of the index.
	IndexConfigInfo *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo `json:"indexConfigInfo,omitempty" xml:"indexConfigInfo,omitempty" type:"Struct"`
	// The size of the index.
	//
	// example:
	//
	// 100
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The name of the worker that failed because of insufficient disk space.
	LackDiskWorker []*string `json:"lackDiskWorker,omitempty" xml:"lackDiskWorker,omitempty" type:"Repeated"`
	// The name of the worker that failed because of insufficient memory.
	LackMemWorker []*string `json:"lackMemWorker,omitempty" xml:"lackMemWorker,omitempty" type:"Repeated"`
	// The name of the node.
	//
	// example:
	//
	// ha-cn-c4d2rq7nt04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of nodes that you specify when you create the cluster.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetAdvanceConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.AdvanceConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDeployFailedWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DeployFailedWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDocSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DocSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetErrorMsg(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.ErrorMsg = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetFullVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.FullVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIncVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IncVersion = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexConfigInfo(v *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexConfigInfo = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetIndexSize(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.IndexSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackDiskWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackDiskWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetLackMemWorker(v []*string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.LackMemWorker = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo struct {
	// The name of the index configuration.
	//
	// example:
	//
	// index_meta_name
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListAdvanceConfigInfo {
	s.Version = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo struct {
	// The name of the index configuration.
	//
	// example:
	//
	// index_meta_name
	ConfigMetaName *string `json:"configMetaName,omitempty" xml:"configMetaName,omitempty"`
	// The version of the index template.
	//
	// example:
	//
	// 1.0.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetConfigMetaName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.ConfigMetaName = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo) SetVersion(v int64) *GetClusterRunTimeInfoResponseBodyResultDataNodesDataStatusListIndexConfigInfo {
	s.Version = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus struct {
	// The overall progress.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of nodes being processed in the cluster.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name.
	//
	// example:
	//
	// ha-cn-0ju2s170b03_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of nodes in the cluster.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultDataNodesServiceStatus {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNode struct {
	// The dataStatusList.
	ConfigStatusList []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList `json:"configStatusList,omitempty" xml:"configStatusList,omitempty" type:"Repeated"`
	// The service status.
	ServiceStatus *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty" type:"Struct"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetConfigStatusList(v []*GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ConfigStatusList = v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNode) SetServiceStatus(v *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) *GetClusterRunTimeInfoResponseBodyResultQueryNode {
	s.ServiceStatus = v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList struct {
	// The time when the cluster was updated.
	//
	// example:
	//
	// " "
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	// The progress.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of nodes that are configured.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of nodes that you specify when you create the cluster.
	//
	// example:
	//
	// 6
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetConfigUpdateTime(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.ConfigUpdateTime = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeConfigStatusList {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus struct {
	// The progress.
	//
	// example:
	//
	// 100
	DonePercent *int32 `json:"donePercent,omitempty" xml:"donePercent,omitempty"`
	// The number of nodes that are configured.
	//
	// example:
	//
	// 100
	DoneSize *int32 `json:"doneSize,omitempty" xml:"doneSize,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ha-cn-c4d2rq7nt04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The total number of nodes that you specify when you create the cluster.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDonePercent(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DonePercent = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetDoneSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.DoneSize = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetName(v string) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.Name = &v
	return s
}

func (s *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus) SetTotalSize(v int32) *GetClusterRunTimeInfoResponseBodyResultQueryNodeServiceStatus {
	s.TotalSize = &v
	return s
}

type GetClusterRunTimeInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterRunTimeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterRunTimeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRunTimeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponse) SetHeaders(v map[string]*string) *GetClusterRunTimeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetStatusCode(v int32) *GetClusterRunTimeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetBody(v *GetClusterRunTimeInfoResponseBody) *GetClusterRunTimeInfoResponse {
	s.Body = v
	return s
}

type GetDataSourceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of information about the data source
	Result *GetDataSourceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetResult(v *GetDataSourceResponseBodyResult) *GetDataSourceResponseBody {
	s.Result = v
	return s
}

type GetDataSourceResponseBodyResult struct {
	// The data center where the data source is deployed in offline mode
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The list of index information
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The time when an index for full data was last built
	LastFulTime *int64 `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	// The name of the data source
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The state of the data source
	//
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDataSourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyResult) SetDomain(v string) *GetDataSourceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetIndexes(v []*string) *GetDataSourceResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetLastFulTime(v int64) *GetDataSourceResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetName(v string) *GetDataSourceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetStatus(v string) *GetDataSourceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetType(v string) *GetDataSourceResponseBodyResult {
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

type GetDataSourceDeployResponseBody struct {
	// requestId
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Deploy
	Result *GetDataSourceDeployResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBody) SetRequestId(v string) *GetDataSourceDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceDeployResponseBody) SetResult(v *GetDataSourceDeployResponseBodyResult) *GetDataSourceDeployResponseBody {
	s.Result = v
	return s
}

type GetDataSourceDeployResponseBodyResult struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                                           `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Extend         *GetDataSourceDeployResponseBodyResultExtend    `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	Processor      *GetDataSourceDeployResponseBodyResultProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Storage        *GetDataSourceDeployResponseBodyResultStorage   `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	Swift          *GetDataSourceDeployResponseBodyResultSwift     `json:"swift,omitempty" xml:"swift,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResult) SetAutoBuildIndex(v bool) *GetDataSourceDeployResponseBodyResult {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetExtend(v *GetDataSourceDeployResponseBodyResultExtend) *GetDataSourceDeployResponseBodyResult {
	s.Extend = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetProcessor(v *GetDataSourceDeployResponseBodyResultProcessor) *GetDataSourceDeployResponseBodyResult {
	s.Processor = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetStorage(v *GetDataSourceDeployResponseBodyResultStorage) *GetDataSourceDeployResponseBodyResult {
	s.Storage = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetSwift(v *GetDataSourceDeployResponseBodyResultSwift) *GetDataSourceDeployResponseBodyResult {
	s.Swift = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtend struct {
	Hdfs *GetDataSourceDeployResponseBodyResultExtendHdfs `json:"hdfs,omitempty" xml:"hdfs,omitempty" type:"Struct"`
	Odps *GetDataSourceDeployResponseBodyResultExtendOdps `json:"odps,omitempty" xml:"odps,omitempty" type:"Struct"`
	Oss  *GetDataSourceDeployResponseBodyResultExtendOss  `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
	Saro *GetDataSourceDeployResponseBodyResultExtendSaro `json:"saro,omitempty" xml:"saro,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResultExtend) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtend) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetHdfs(v *GetDataSourceDeployResponseBodyResultExtendHdfs) *GetDataSourceDeployResponseBodyResultExtend {
	s.Hdfs = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOdps(v *GetDataSourceDeployResponseBodyResultExtendOdps) *GetDataSourceDeployResponseBodyResultExtend {
	s.Odps = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOss(v *GetDataSourceDeployResponseBodyResultExtendOss) *GetDataSourceDeployResponseBodyResultExtend {
	s.Oss = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetSaro(v *GetDataSourceDeployResponseBodyResultExtendSaro) *GetDataSourceDeployResponseBodyResultExtend {
	s.Saro = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendHdfs struct {
	// example:
	//
	// dist-dmj-job/src/main/java
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendHdfs) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendHdfs {
	s.Path = &v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendOdps struct {
	Partitions map[string]*string `json:"partitions,omitempty" xml:"partitions,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOdps) SetPartitions(v map[string]*string) *GetDataSourceDeployResponseBodyResultExtendOdps {
	s.Partitions = v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendOss struct {
	// example:
	//
	// oss://opensearch
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOss) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendOss {
	s.Path = &v
	return s
}

type GetDataSourceDeployResponseBodyResultExtendSaro struct {
	// example:
	//
	// dist-dmj-job/src/main/java
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 0.6.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetVersion(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Version = &v
	return s
}

type GetDataSourceDeployResponseBodyResultProcessor struct {
	// example:
	//
	// {}
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// example:
	//
	// {}
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultProcessor) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultProcessor) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetArgs(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Args = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetResource(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Resource = &v
	return s
}

type GetDataSourceDeployResponseBodyResultStorage struct {
	// AK
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// oss bucket
	//
	// example:
	//
	// antsys-miniapp-chongwen-static
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// odps endpoint
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// lazada-campaign-flink
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20220926
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// /beiming_xobject/dwd_xobjectsandbox__list_create_action_by_new/
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// wireless_1688_personal_rec
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// behavior
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultStorage) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultStorage) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessKey(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessKey = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessSecret(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessSecret = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetBucket(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Bucket = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetEndpoint(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Endpoint = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetNamespace(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Namespace = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetOssPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.OssPath = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPartition(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Partition = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetProject(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Project = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetTable(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Table = &v
	return s
}

type GetDataSourceDeployResponseBodyResultSwift struct {
	// topic
	//
	// example:
	//
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// zk
	//
	// example:
	//
	// zk
	Zk *string `json:"zk,omitempty" xml:"zk,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultSwift) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultSwift) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetTopic(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Topic = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetZk(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Zk = &v
	return s
}

type GetDataSourceDeployResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceDeployResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceDeployResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponse) SetHeaders(v map[string]*string) *GetDataSourceDeployResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceDeployResponse) SetStatusCode(v int32) *GetDataSourceDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceDeployResponse) SetBody(v *GetDataSourceDeployResponseBody) *GetDataSourceDeployResponse {
	s.Body = v
	return s
}

type GetDeployGraphResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result *GetDeployGraphResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBody) SetRequestId(v string) *GetDeployGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployGraphResponseBody) SetResult(v *GetDeployGraphResponseBodyResult) *GetDeployGraphResponseBody {
	s.Result = v
	return s
}

type GetDeployGraphResponseBodyResult struct {
	Graph *GetDeployGraphResponseBodyResultGraph `json:"graph,omitempty" xml:"graph,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResult) SetGraph(v *GetDeployGraphResponseBodyResultGraph) *GetDeployGraphResponseBodyResult {
	s.Graph = v
	return s
}

type GetDeployGraphResponseBodyResultGraph struct {
	// 索引元信息
	IndexMetas []*GetDeployGraphResponseBodyResultGraphIndexMetas `json:"indexMetas,omitempty" xml:"indexMetas,omitempty" type:"Repeated"`
	// 在线集群元信息
	OnlineMaster []*GetDeployGraphResponseBodyResultGraphOnlineMaster `json:"onlineMaster,omitempty" xml:"onlineMaster,omitempty" type:"Repeated"`
	// 数据源和索引关联关系
	TableIndexRelation map[string][]*string `json:"tableIndexRelation,omitempty" xml:"tableIndexRelation,omitempty"`
	// 数据源元信息
	TableMetas []*GetDeployGraphResponseBodyResultGraphTableMetas `json:"tableMetas,omitempty" xml:"tableMetas,omitempty" type:"Repeated"`
	// zone和索引关联关系
	ZoneIndexRelation map[string][]*string `json:"zoneIndexRelation,omitempty" xml:"zoneIndexRelation,omitempty"`
	// zone元信息
	ZoneMetas []*GetDeployGraphResponseBodyResultGraphZoneMetas `json:"zoneMetas,omitempty" xml:"zoneMetas,omitempty" type:"Repeated"`
}

func (s GetDeployGraphResponseBodyResultGraph) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraph) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraph) SetIndexMetas(v []*GetDeployGraphResponseBodyResultGraphIndexMetas) *GetDeployGraphResponseBodyResultGraph {
	s.IndexMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetOnlineMaster(v []*GetDeployGraphResponseBodyResultGraphOnlineMaster) *GetDeployGraphResponseBodyResultGraph {
	s.OnlineMaster = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.TableIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableMetas(v []*GetDeployGraphResponseBodyResultGraphTableMetas) *GetDeployGraphResponseBodyResultGraph {
	s.TableMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneMetas(v []*GetDeployGraphResponseBodyResultGraphZoneMetas) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneMetas = v
	return s
}

type GetDeployGraphResponseBodyResultGraphIndexMetas struct {
	DomainName    *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	TableDeployId *int64  `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	TableName     *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	ZoneName      *string `json:"zoneName,omitempty" xml:"zoneName,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetZoneName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.ZoneName = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphOnlineMaster struct {
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	HippoId    *string `json:"hippoId,omitempty" xml:"hippoId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetHippoId(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.HippoId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetId(v int64) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Id = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Name = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphTableMetas struct {
	BuildDeployId *int64  `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	DomainName    *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	TableDeployId *int64  `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetBuildDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.BuildDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Type = &v
	return s
}

type GetDeployGraphResponseBodyResultGraphZoneMetas struct {
	DomainInfo    *string `json:"domainInfo,omitempty" xml:"domainInfo,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	SuezAdminName *string `json:"suezAdminName,omitempty" xml:"suezAdminName,omitempty"`
	Tag           *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetDomainInfo(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.DomainInfo = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetSuezAdminName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.SuezAdminName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Type = &v
	return s
}

type GetDeployGraphResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeployGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeployGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeployGraphResponse) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponse) SetHeaders(v map[string]*string) *GetDeployGraphResponse {
	s.Headers = v
	return s
}

func (s *GetDeployGraphResponse) SetStatusCode(v int32) *GetDeployGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployGraphResponse) SetBody(v *GetDeployGraphResponseBody) *GetDeployGraphResponse {
	s.Body = v
	return s
}

type GetFileRequest struct {
	// The name of the file in full path
	//
	// This parameter is required.
	//
	// example:
	//
	// /schemas/automobile_vector_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) SetFileName(v string) *GetFileRequest {
	s.FileName = &v
	return s
}

type GetFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result *GetFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileResponseBody) SetRequestId(v string) *GetFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileResponseBody) SetResult(v *GetFileResponseBodyResult) *GetFileResponseBody {
	s.Result = v
	return s
}

type GetFileResponseBodyResult struct {
	// The content of the file.
	//
	// example:
	//
	// None
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The data source.
	//
	// example:
	//
	// ha-cn-2r42p5oi202_xijie_test
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The name of the full path.
	//
	// example:
	//
	// ""
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether it is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// ds=20210828
	Partition *int64 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyResult) SetContent(v string) *GetFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileResponseBodyResult) SetDataSource(v string) *GetFileResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetFileResponseBodyResult) SetFullPathName(v string) *GetFileResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *GetFileResponseBodyResult) SetIsDir(v bool) *GetFileResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *GetFileResponseBodyResult) SetName(v string) *GetFileResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetFileResponseBodyResult) SetPartition(v int64) *GetFileResponseBodyResult {
	s.Partition = &v
	return s
}

type GetFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetStatusCode(v int32) *GetFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileResponse) SetBody(v *GetFileResponseBody) *GetFileResponse {
	s.Body = v
	return s
}

type GetIndexResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result *GetIndexResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBody) SetRequestId(v string) *GetIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexResponseBody) SetResult(v *GetIndexResponseBodyResult) *GetIndexResponseBody {
	s.Result = v
	return s
}

type GetIndexResponseBodyResult struct {
	Cluster         map[string]*ResultClusterValue    `json:"cluster,omitempty" xml:"cluster,omitempty"`
	Config          map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	ConfigWhenBuild map[string]map[string]interface{} `json:"configWhenBuild,omitempty" xml:"configWhenBuild,omitempty"`
	// The content of the index.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// ha-cn-tl32nd2nq01_00
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source.
	DataSourceInfo *GetIndexResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The remarks.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The last time when full data in the index was updated.
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The version of the data.
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The last time when incremental data in the index was updated.
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The index size.
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The status of the index. Valid values: NEW, PUBLISH, IN_USE, NOT_USE, STOP_USE, and RESTORE_USE. After a Retrieval Engine Edition instance is created, it enters the IN_USE state.
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The information about the versions.
	Versions []*GetIndexResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResult) SetCluster(v map[string]*ResultClusterValue) *GetIndexResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *GetIndexResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *GetIndexResponseBodyResult {
	s.Config = v
	return s
}

func (s *GetIndexResponseBodyResult) SetConfigWhenBuild(v map[string]map[string]interface{}) *GetIndexResponseBodyResult {
	s.ConfigWhenBuild = v
	return s
}

func (s *GetIndexResponseBodyResult) SetContent(v string) *GetIndexResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSource(v string) *GetIndexResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSourceInfo(v *GetIndexResponseBodyResultDataSourceInfo) *GetIndexResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *GetIndexResponseBodyResult) SetDescription(v string) *GetIndexResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDomain(v string) *GetIndexResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullUpdateTime(v string) *GetIndexResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullVersion(v int64) *GetIndexResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIncUpdateTime(v string) *GetIndexResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexSize(v int64) *GetIndexResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexStatus(v string) *GetIndexResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetName(v string) *GetIndexResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetPartition(v int32) *GetIndexResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetVersions(v []*GetIndexResponseBodyResultVersions) *GetIndexResponseBodyResult {
	s.Versions = v
	return s
}

type GetIndexResponseBodyResultDataSourceInfo struct {
	// Indicates whether the automatic full indexing feature is enabled.
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration of MaxCompute data sources.
	Config *GetIndexResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The offline deployment name of the data source.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the data source.
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessParallelNum *int32  `json:"processParallelNum,omitempty" xml:"processParallelNum,omitempty"`
	// The number of resources used for data update.
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configuration of SARO data sources.
	SaroConfig *GetIndexResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *GetIndexResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetConfig(v *GetIndexResponseBodyResultDataSourceInfoConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetDomain(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetName(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetProcessParallelNum(v int32) *GetIndexResponseBodyResultDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *GetIndexResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetSaroConfig(v *GetIndexResponseBodyResultDataSourceInfoSaroConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetType(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

type GetIndexResponseBodyResultDataSourceInfoConfig struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	Bucket       *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// A parameter related to MaxCompute.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// A parameter related to SARO.
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// A parameter related to OSS.
	OssPath   *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// A parameter related to Apsara File Storage for HDFS.
	Path    *string `json:"path,omitempty" xml:"path,omitempty"`
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// A parameter related to SARO and MaxCompute.
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetProject(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetTable(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

type GetIndexResponseBodyResultDataSourceInfoSaroConfig struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type GetIndexResponseBodyResultVersions struct {
	// The description of the version.
	//
	// example:
	//
	// close alarm, by 3.9.2 hotfix workflow
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about the files.
	Files []*GetIndexResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the version.
	//
	// example:
	//
	// ha-cn-7pp2ngv4s02_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the version.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last time when the version was updated.
	//
	// example:
	//
	// " "
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the version.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetIndexResponseBodyResultVersions) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersions) SetDesc(v string) *GetIndexResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetFiles(v []*GetIndexResponseBodyResultVersionsFiles) *GetIndexResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetName(v string) *GetIndexResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetStatus(v string) *GetIndexResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetUpdateTime(v int64) *GetIndexResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetVersionId(v int32) *GetIndexResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

type GetIndexResponseBodyResultVersionsFiles struct {
	// The full path of the file.
	//
	// example:
	//
	// " "
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template.
	//
	// example:
	//
	// True
	IsTemplate *bool   `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetIndexResponseBodyResultVersionsFiles) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetFullPathName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsDir(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

type GetIndexResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponse) GoString() string {
	return s.String()
}

func (s *GetIndexResponse) SetHeaders(v map[string]*string) *GetIndexResponse {
	s.Headers = v
	return s
}

func (s *GetIndexResponse) SetStatusCode(v int32) *GetIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexResponse) SetBody(v *GetIndexResponseBody) *GetIndexResponse {
	s.Body = v
	return s
}

type GetIndexOnlineStrategyResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// IndexOnlineStrategy
	Result *GetIndexOnlineStrategyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexOnlineStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexOnlineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponseBody) SetRequestId(v string) *GetIndexOnlineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexOnlineStrategyResponseBody) SetResult(v *GetIndexOnlineStrategyResponseBodyResult) *GetIndexOnlineStrategyResponseBody {
	s.Result = v
	return s
}

type GetIndexOnlineStrategyResponseBodyResult struct {
	// example:
	//
	// 20
	ChangeRate *int32 `json:"changeRate,omitempty" xml:"changeRate,omitempty"`
}

func (s GetIndexOnlineStrategyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIndexOnlineStrategyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponseBodyResult) SetChangeRate(v int32) *GetIndexOnlineStrategyResponseBodyResult {
	s.ChangeRate = &v
	return s
}

type GetIndexOnlineStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexOnlineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexOnlineStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexOnlineStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponse) SetHeaders(v map[string]*string) *GetIndexOnlineStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetIndexOnlineStrategyResponse) SetStatusCode(v int32) *GetIndexOnlineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexOnlineStrategyResponse) SetBody(v *GetIndexOnlineStrategyResponseBody) *GetIndexOnlineStrategyResponse {
	s.Body = v
	return s
}

type GetIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of cluster details
	Result *GetIndexVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBody) SetRequestId(v string) *GetIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexVersionResponseBody) SetResult(v *GetIndexVersionResponseBodyResult) *GetIndexVersionResponseBody {
	s.Result = v
	return s
}

type GetIndexVersionResponseBodyResult struct {
	// The name of the cluster
	//
	// example:
	//
	// ayoss-cn-zhangjiakou-b
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the cluster was updated
	IndexVersions []*GetIndexVersionResponseBodyResultIndexVersions `json:"indexVersions,omitempty" xml:"indexVersions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResult) SetCluster(v string) *GetIndexVersionResponseBodyResult {
	s.Cluster = &v
	return s
}

func (s *GetIndexVersionResponseBodyResult) SetIndexVersions(v []*GetIndexVersionResponseBodyResultIndexVersions) *GetIndexVersionResponseBodyResult {
	s.IndexVersions = v
	return s
}

type GetIndexVersionResponseBodyResultIndexVersions struct {
	// The ID of the index deployed in offline mode
	//
	// example:
	//
	// " "
	BuildDeployId  *string `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	CurrentVersion *int64  `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The name of the index table
	//
	// example:
	//
	// table4
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The version of the index
	Versions []*int64 `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResultIndexVersions) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponseBodyResultIndexVersions) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetBuildDeployId(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.BuildDeployId = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetCurrentVersion(v int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.CurrentVersion = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetIndexName(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.IndexName = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetVersions(v []*int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.Versions = v
	return s
}

type GetIndexVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponse) SetHeaders(v map[string]*string) *GetIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *GetIndexVersionResponse) SetStatusCode(v int32) *GetIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexVersionResponse) SetBody(v *GetIndexVersionResponseBody) *GetIndexVersionResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result *GetInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResult(v *GetInstanceResponseBodyResult) *GetInstanceResponseBody {
	s.Result = v
	return s
}

type GetInstanceResponseBodyResult struct {
	// The billing method.
	//
	// example:
	//
	// POSYPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The product code.
	//
	// example:
	//
	// commodityCode
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-06-17T02:01:26Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ha3_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1634609702
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved.
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// ha-cn-7mz2qsgq301
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzjvw24el5lma
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- INIT: being initialized
	//
	// 	- WAIT_CONFIG: to be configured
	//
	// 	- CONFIG_UPDATING: configuration taking effect
	//
	// 	- READY: normal
	//
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 标签。
	Tags []*GetInstanceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 1634609702
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResult) SetChargeType(v string) *GetInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCommodityCode(v string) *GetInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCreateTime(v string) *GetInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetDescription(v string) *GetInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetExpiredTime(v string) *GetInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInDebt(v bool) *GetInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInstanceId(v string) *GetInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetLockMode(v string) *GetInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetResourceGroupId(v string) *GetInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetStatus(v string) *GetInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetTags(v []*GetInstanceResponseBodyResultTags) *GetInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyResult) SetUpdateTime(v string) *GetInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type GetInstanceResponseBodyResultTags struct {
	// 标签键
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签值
	//
	// example:
	//
	// prod
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetInstanceResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultTags) SetKey(v string) *GetInstanceResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyResultTags) SetValue(v string) *GetInstanceResponseBodyResultTags {
	s.Value = &v
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

type GetNodeConfigRequest struct {
	// The name of the cluster
	//
	// example:
	//
	// vpc_sh_domain_2
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ha-cn-30174dhoz53_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the node. Valid values: qrs, search, index, and cluster. qrs indicates a query node, search indicates a data node, index indicates an index node, and cluster indicates a cluster node.
	//
	// example:
	//
	// index
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetNodeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNodeConfigRequest) SetClusterName(v string) *GetNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *GetNodeConfigRequest) SetName(v string) *GetNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *GetNodeConfigRequest) SetType(v string) *GetNodeConfigRequest {
	s.Type = &v
	return s
}

type GetNodeConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The configurations of the node.
	Result *GetNodeConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetNodeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBody) SetRequestId(v string) *GetNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeConfigResponseBody) SetResult(v *GetNodeConfigResponseBodyResult) *GetNodeConfigResponseBody {
	s.Result = v
	return s
}

type GetNodeConfigResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// 1
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The number of data replicas.
	//
	// example:
	//
	// 2
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 2
	DataFragmentNumber *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	// The traffic percentage.
	//
	// example:
	//
	// 0
	FlowRatio *int32 `json:"flowRatio,omitempty" xml:"flowRatio,omitempty"`
	// The minimum service ratio.
	//
	// example:
	//
	// 100
	MinServicePercent *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	// Indicates whether the node is associated with the cluster.
	//
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
}

func (s GetNodeConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponseBodyResult) SetActive(v bool) *GetNodeConfigResponseBodyResult {
	s.Active = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataDuplicateNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataDuplicateNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetDataFragmentNumber(v int32) *GetNodeConfigResponseBodyResult {
	s.DataFragmentNumber = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetFlowRatio(v int32) *GetNodeConfigResponseBodyResult {
	s.FlowRatio = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetMinServicePercent(v int32) *GetNodeConfigResponseBodyResult {
	s.MinServicePercent = &v
	return s
}

func (s *GetNodeConfigResponseBodyResult) SetPublished(v bool) *GetNodeConfigResponseBodyResult {
	s.Published = &v
	return s
}

type GetNodeConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponse) SetHeaders(v map[string]*string) *GetNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNodeConfigResponse) SetStatusCode(v int32) *GetNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeConfigResponse) SetBody(v *GetNodeConfigResponseBody) *GetNodeConfigResponse {
	s.Body = v
	return s
}

type GetTableResponseBody struct {
	// requestId
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetResult(v *GetTableResponseBodyResult) *GetTableResponseBody {
	s.Result = v
	return s
}

type GetTableResponseBodyResult struct {
	DataProcessConfig []*GetTableResponseBodyResultDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DataProcessorCount *int32                                `json:"dataProcessorCount,omitempty" xml:"dataProcessorCount,omitempty"`
	DataSource         *GetTableResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	FieldSchema        map[string]*string                    `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// example:
	//
	// test_oss
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// example:
	//
	// {}
	RawSchema   *string                                  `json:"rawSchema,omitempty" xml:"rawSchema,omitempty"`
	Status      *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	VectorIndex []*GetTableResponseBodyResultVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
}

func (s GetTableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResult) SetDataProcessConfig(v []*GetTableResponseBodyResultDataProcessConfig) *GetTableResponseBodyResult {
	s.DataProcessConfig = v
	return s
}

func (s *GetTableResponseBodyResult) SetDataProcessorCount(v int32) *GetTableResponseBodyResult {
	s.DataProcessorCount = &v
	return s
}

func (s *GetTableResponseBodyResult) SetDataSource(v *GetTableResponseBodyResultDataSource) *GetTableResponseBodyResult {
	s.DataSource = v
	return s
}

func (s *GetTableResponseBodyResult) SetFieldSchema(v map[string]*string) *GetTableResponseBodyResult {
	s.FieldSchema = v
	return s
}

func (s *GetTableResponseBodyResult) SetName(v string) *GetTableResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetTableResponseBodyResult) SetPartitionCount(v int32) *GetTableResponseBodyResult {
	s.PartitionCount = &v
	return s
}

func (s *GetTableResponseBodyResult) SetPrimaryKey(v string) *GetTableResponseBodyResult {
	s.PrimaryKey = &v
	return s
}

func (s *GetTableResponseBodyResult) SetRawSchema(v string) *GetTableResponseBodyResult {
	s.RawSchema = &v
	return s
}

func (s *GetTableResponseBodyResult) SetStatus(v string) *GetTableResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetTableResponseBodyResult) SetVectorIndex(v []*GetTableResponseBodyResultVectorIndex) *GetTableResponseBodyResult {
	s.VectorIndex = v
	return s
}

type GetTableResponseBodyResultDataProcessConfig struct {
	// example:
	//
	// source_image_vector
	DstField *string `json:"dstField,omitempty" xml:"dstField,omitempty"`
	// example:
	//
	// vectorize
	Operator *string                                            `json:"operator,omitempty" xml:"operator,omitempty"`
	Params   *GetTableResponseBodyResultDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s GetTableResponseBodyResultDataProcessConfig) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetDstField(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetOperator(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetParams(v *GetTableResponseBodyResultDataProcessConfigParams) *GetTableResponseBodyResultDataProcessConfig {
	s.Params = v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetSrcField(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.SrcField = &v
	return s
}

type GetTableResponseBodyResultDataProcessConfigParams struct {
	SrcFieldConfig *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
	// example:
	//
	// image
	VectorModal *string `json:"vectorModal,omitempty" xml:"vectorModal,omitempty"`
	// example:
	//
	// clip
	VectorModel *string `json:"vectorModel,omitempty" xml:"vectorModel,omitempty"`
}

func (s GetTableResponseBodyResultDataProcessConfigParams) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetSrcFieldConfig(v *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) *GetTableResponseBodyResultDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetVectorModal(v string) *GetTableResponseBodyResultDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetVectorModel(v string) *GetTableResponseBodyResultDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

type GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig struct {
	// OSS Bucket
	//
	// example:
	//
	// test
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// example:
	//
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

type GetTableResponseBodyResultDataSource struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                                       `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *GetTableResponseBodyResultDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableResponseBodyResultDataSource) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultDataSource) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataSource) SetAutoBuildIndex(v bool) *GetTableResponseBodyResultDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetConfig(v *GetTableResponseBodyResultDataSourceConfig) *GetTableResponseBodyResultDataSource {
	s.Config = v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetDataTimeSec(v int32) *GetTableResponseBodyResultDataSource {
	s.DataTimeSec = &v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetType(v string) *GetTableResponseBodyResultDataSource {
	s.Type = &v
	return s
}

type GetTableResponseBodyResultDataSourceConfig struct {
	// AK
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// example:
	//
	// heytea-ops-oss
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// /opensearch_index_data/sift_oss_test.data
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20220808
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// vendor/sebastian/comparator/src/exceptions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// dp_pdm_marketing_prod
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// test_add
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetTableResponseBodyResultDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultDataSourceConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetAccessKey(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetAccessSecret(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetBucket(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetEndpoint(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetNamespace(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Namespace = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetOssPath(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetPartition(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetPath(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Path = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetProject(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Project = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetTable(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Table = &v
	return s
}

type GetTableResponseBodyResultVectorIndex struct {
	AdvanceParams *GetTableResponseBodyResultVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
	// example:
	//
	// 128
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// SquaredEuclidean
	DistanceType *string `json:"distanceType,omitempty" xml:"distanceType,omitempty"`
	// example:
	//
	// test_odps
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// sparse_indices
	SparseIndexField *string `json:"sparseIndexField,omitempty" xml:"sparseIndexField,omitempty"`
	// example:
	//
	// sparse_values
	SparseValueField *string `json:"sparseValueField,omitempty" xml:"sparseValueField,omitempty"`
	// example:
	//
	// source_image_vector
	VectorField *string `json:"vectorField,omitempty" xml:"vectorField,omitempty"`
	// example:
	//
	// Qc
	VectorIndexType *string `json:"vectorIndexType,omitempty" xml:"vectorIndexType,omitempty"`
}

func (s GetTableResponseBodyResultVectorIndex) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultVectorIndex) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultVectorIndex) SetAdvanceParams(v *GetTableResponseBodyResultVectorIndexAdvanceParams) *GetTableResponseBodyResultVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetDimension(v string) *GetTableResponseBodyResultVectorIndex {
	s.Dimension = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetDistanceType(v string) *GetTableResponseBodyResultVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetIndexName(v string) *GetTableResponseBodyResultVectorIndex {
	s.IndexName = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetNamespace(v string) *GetTableResponseBodyResultVectorIndex {
	s.Namespace = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetSparseIndexField(v string) *GetTableResponseBodyResultVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetSparseValueField(v string) *GetTableResponseBodyResultVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetVectorField(v string) *GetTableResponseBodyResultVectorIndex {
	s.VectorField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetVectorIndexType(v string) *GetTableResponseBodyResultVectorIndex {
	s.VectorIndexType = &v
	return s
}

type GetTableResponseBodyResultVectorIndexAdvanceParams struct {
	// example:
	//
	// {}
	BuildIndexParams *string `json:"buildIndexParams,omitempty" xml:"buildIndexParams,omitempty"`
	// example:
	//
	// 5000
	LinearBuildThreshold *string `json:"linearBuildThreshold,omitempty" xml:"linearBuildThreshold,omitempty"`
	// example:
	//
	// 20000
	MinScanDocCnt *string `json:"minScanDocCnt,omitempty" xml:"minScanDocCnt,omitempty"`
	// example:
	//
	// {}
	SearchIndexParams *string `json:"searchIndexParams,omitempty" xml:"searchIndexParams,omitempty"`
}

func (s GetTableResponseBodyResultVectorIndexAdvanceParams) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBodyResultVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetBuildIndexParams(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetSearchIndexParams(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
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

type GetTableGenerationResponseBody struct {
	// requestId
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetTableGenerationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTableGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponseBody) SetRequestId(v string) *GetTableGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableGenerationResponseBody) SetResult(v *GetTableGenerationResponseBodyResult) *GetTableGenerationResponseBody {
	s.Result = v
	return s
}

type GetTableGenerationResponseBodyResult struct {
	// generationId
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
	// starting, building, ready, stopped, failed
	//
	// example:
	//
	// ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTableGenerationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTableGenerationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponseBodyResult) SetGenerationId(v int64) *GetTableGenerationResponseBodyResult {
	s.GenerationId = &v
	return s
}

func (s *GetTableGenerationResponseBodyResult) SetStatus(v string) *GetTableGenerationResponseBodyResult {
	s.Status = &v
	return s
}

type GetTableGenerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableGenerationResponse) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponse) SetHeaders(v map[string]*string) *GetTableGenerationResponse {
	s.Headers = v
	return s
}

func (s *GetTableGenerationResponse) SetStatusCode(v int32) *GetTableGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableGenerationResponse) SetBody(v *GetTableGenerationResponseBody) *GetTableGenerationResponse {
	s.Body = v
	return s
}

type ListAdvanceConfigDirRequest struct {
	// The name of the directory
	//
	// This parameter is required.
	//
	// example:
	//
	// /zones/general
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
}

func (s ListAdvanceConfigDirRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirRequest) SetDirName(v string) *ListAdvanceConfigDirRequest {
	s.DirName = &v
	return s
}

type ListAdvanceConfigDirResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The file list in the advanced configuration directory.
	Result []*ListAdvanceConfigDirResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBody) SetRequestId(v string) *ListAdvanceConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBody) SetResult(v []*ListAdvanceConfigDirResponseBodyResult) *ListAdvanceConfigDirResponseBody {
	s.Result = v
	return s
}

type ListAdvanceConfigDirResponseBodyResult struct {
	// The name of the absolute path.
	//
	// example:
	//
	// "/path/wpd/nae"
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether it is a directory. Valid values: true and false. true indicates that it is a directory, and false indicates that it is not a directory.
	//
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether it is a template. Valid values: **true*	- and **false**. true indicates that it is a template, and false indicates that it is not a template.
	//
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// file_name_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigDirResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetFullPathName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsDir(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsTemplate(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.Name = &v
	return s
}

type ListAdvanceConfigDirResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdvanceConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdvanceConfigDirResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigDirResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigDirResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetStatusCode(v int32) *ListAdvanceConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetBody(v *ListAdvanceConfigDirResponseBody) *ListAdvanceConfigDirResponse {
	s.Body = v
	return s
}

type ListAdvanceConfigsRequest struct {
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	IndexName      *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	NewMode        *bool   `json:"newMode,omitempty" xml:"newMode,omitempty"`
	// The type of the advanced configurations. Valid values: online and offline. - online The default value is offline.
	//
	// example:
	//
	// online
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAdvanceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsRequest) SetDataSourceName(v string) *ListAdvanceConfigsRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetIndexName(v string) *ListAdvanceConfigsRequest {
	s.IndexName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetNewMode(v bool) *ListAdvanceConfigsRequest {
	s.NewMode = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetType(v string) *ListAdvanceConfigsRequest {
	s.Type = &v
	return s
}

type ListAdvanceConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of advanced configurations.
	Result []*ListAdvanceConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBody) SetRequestId(v string) *ListAdvanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigsResponseBody) SetResult(v []*ListAdvanceConfigsResponseBodyResult) *ListAdvanceConfigsResponseBody {
	s.Result = v
	return s
}

type ListAdvanceConfigsResponseBodyResult struct {
	AdvanceConfigType *string `json:"advanceConfigType,omitempty" xml:"advanceConfigType,omitempty"`
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType       *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Creator           *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The list of file names.
	Files []*ListAdvanceConfigsResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the advanced configuration.
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The state of the advanced configuration. Valid values: drafting, used, unused, and trash. drafting indicates that the advanced configuration is a draft. used indicates that the advanced configuration is in use. unused indicates that the advanced configuration is unused. trash indicates that the advanced configuration is being deleted.
	//
	// example:
	//
	// drafting
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1631070464000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResult) SetAdvanceConfigType(v string) *ListAdvanceConfigsResponseBodyResult {
	s.AdvanceConfigType = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContent(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContentType(v string) *ListAdvanceConfigsResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetCreator(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetDesc(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetFiles(v []*ListAdvanceConfigsResponseBodyResultFiles) *ListAdvanceConfigsResponseBodyResult {
	s.Files = v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetName(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetStatus(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetUpdateTime(v int64) *ListAdvanceConfigsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type ListAdvanceConfigsResponseBodyResultFiles struct {
	// The name of the absolute path.
	//
	// example:
	//
	// /path/wpd/nae
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether it is a directory. Valid values: true and false. true indicates that it is a directory, and false indicates that it is not a directory.
	//
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether it is a template. Valid values: true and false. true indicates that it is a directory, and false indicates that it is not a directory.
	//
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// file_name_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResultFiles) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetFullPathName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsDir(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsTemplate(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.Name = &v
	return s
}

type ListAdvanceConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdvanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdvanceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdvanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigsResponse) SetStatusCode(v int32) *ListAdvanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigsResponse) SetBody(v *ListAdvanceConfigsResponseBody) *ListAdvanceConfigsResponse {
	s.Body = v
	return s
}

type ListClusterNamesResponseBody struct {
	// id of request
	//
	// example:
	//
	// F6E3D968-529C-5C40-AFDD-133A8B8FD930
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result set
	Result *ListClusterNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListClusterNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBody) SetRequestId(v string) *ListClusterNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterNamesResponseBody) SetResult(v *ListClusterNamesResponseBodyResult) *ListClusterNamesResponseBody {
	s.Result = v
	return s
}

type ListClusterNamesResponseBodyResult struct {
	// The description of the cluster
	//
	// example:
	//
	// ha3_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the cluster
	//
	// example:
	//
	// 25030
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the cluster
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListClusterNamesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBodyResult) SetDescription(v string) *ListClusterNamesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetId(v int64) *ListClusterNamesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetName(v string) *ListClusterNamesResponseBodyResult {
	s.Name = &v
	return s
}

type ListClusterNamesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNamesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponse) SetHeaders(v map[string]*string) *ListClusterNamesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNamesResponse) SetStatusCode(v int32) *ListClusterNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNamesResponse) SetBody(v *ListClusterNamesResponseBody) *ListClusterNamesResponse {
	s.Body = v
	return s
}

type ListClusterTasksResponseBody struct {
	// id of request
	//
	// example:
	//
	// CC5EC8FA-5C0D-56AF-BEF4-6FCCEABD0511
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result []*ListClusterTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClusterTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBody) SetRequestId(v string) *ListClusterTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTasksResponseBody) SetResult(v []*ListClusterTasksResponseBodyResult) *ListClusterTasksResponseBody {
	s.Result = v
	return s
}

type ListClusterTasksResponseBodyResult struct {
	// The additional attributes of the card.
	//
	// example:
	//
	// " "
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	// The field3 field that is passed through when you create a state machine.
	//
	// example:
	//
	// " "
	Field3 *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// fsmId
	//
	// example:
	//
	// tisplus_opensearch@datasource_flow_fsm@1865410598556969-ha-cn-zvp2ljiwe01_api2@bj_vpc_domain_1@null@MANUAL-ha-cn-zvp2ljiwe01_api2@1649729867698@028315
	FsmId *string `json:"fsmId,omitempty" xml:"fsmId,omitempty"`
	// Indicates whether the change is a data source task change or a cluster task change.
	//
	// example:
	//
	// " "
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The task name on the card.
	//
	// example:
	//
	// ha-cn-pl32rf0js04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The overall status of FSM.
	//
	// example:
	//
	// onlyPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status tag of the progress bar chart.
	Tags []*ListClusterTasksResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The information about the task.
	TaskNodes []*ListClusterTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	// The timestamp of the task on the card.
	//
	// example:
	//
	// 1657610520
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The type of the task on the card.
	//
	// example:
	//
	// qrs
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The user who triggered the FSM process.
	//
	// example:
	//
	// " "
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListClusterTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResult) SetExtraAttribute(v string) *ListClusterTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetField3(v string) *ListClusterTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetFsmId(v string) *ListClusterTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetGroupType(v string) *ListClusterTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetName(v string) *ListClusterTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetStatus(v string) *ListClusterTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTags(v []*ListClusterTasksResponseBodyResultTags) *ListClusterTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTaskNodes(v []*ListClusterTasksResponseBodyResultTaskNodes) *ListClusterTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTime(v string) *ListClusterTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetType(v string) *ListClusterTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetUser(v string) *ListClusterTasksResponseBodyResult {
	s.User = &v
	return s
}

type ListClusterTasksResponseBodyResultTags struct {
	// The content of the tag.
	//
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The level of the tag.
	//
	// example:
	//
	// " "
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTags) SetMsg(v string) *ListClusterTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTags) SetTagLevel(v string) *ListClusterTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

type ListClusterTasksResponseBodyResultTaskNodes struct {
	// The date when the task was completed.
	//
	// example:
	//
	// " "
	FinishDate *string `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
	// The sequence number of the task.
	//
	// example:
	//
	// 100
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTaskNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetName(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

type ListClusterTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTasksResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponse) SetHeaders(v map[string]*string) *ListClusterTasksResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTasksResponse) SetStatusCode(v int32) *ListClusterTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterTasksResponse) SetBody(v *ListClusterTasksResponseBody) *ListClusterTasksResponse {
	s.Body = v
	return s
}

type ListClustersResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of clusters.
	Result []*ListClustersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetResult(v []*ListClustersResponseBodyResult) *ListClustersResponseBody {
	s.Result = v
	return s
}

type ListClustersResponseBodyResult struct {
	Config map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// " "
	ConfigUpdateTime *string `json:"configUpdateTime,omitempty" xml:"configUpdateTime,omitempty"`
	CreateTime       *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The effective advanced version.
	//
	// example:
	//
	// " "
	CurrentAdvanceConfigVersion *string `json:"currentAdvanceConfigVersion,omitempty" xml:"currentAdvanceConfigVersion,omitempty"`
	// 词典配置生效版本
	CurrentOfflineDictConfigVersion *string `json:"currentOfflineDictConfigVersion,omitempty" xml:"currentOfflineDictConfigVersion,omitempty"`
	// The effective online configuration version.
	//
	// example:
	//
	// " "
	CurrentOnlineConfigVersion *string `json:"currentOnlineConfigVersion,omitempty" xml:"currentOnlineConfigVersion,omitempty"`
	// 查询配置生效版本
	CurrentOnlineQueryConfigVersion *string `json:"currentOnlineQueryConfigVersion,omitempty" xml:"currentOnlineQueryConfigVersion,omitempty"`
	// The information about the node in the cluster.
	DataNode *ListClustersResponseBodyResultDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	//
	// example:
	//
	// fzz_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest advanced configuration version.
	//
	// example:
	//
	// " "
	LatestAdvanceConfigVersion *string `json:"latestAdvanceConfigVersion,omitempty" xml:"latestAdvanceConfigVersion,omitempty"`
	// 词典配置最新版本
	LatestOfflineDictConfigVersion *string `json:"latestOfflineDictConfigVersion,omitempty" xml:"latestOfflineDictConfigVersion,omitempty"`
	// The latest online configuration version.
	//
	// example:
	//
	// " "
	LatestOnlineConfigVersion *string `json:"latestOnlineConfigVersion,omitempty" xml:"latestOnlineConfigVersion,omitempty"`
	// 查询配置最新版本
	LatestOnlineQueryConfigVersion *string `json:"latestOnlineQueryConfigVersion,omitempty" xml:"latestOnlineQueryConfigVersion,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ha-cn-7pp2pcna701_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The query node of the cluster.
	QueryNode *ListClustersResponseBodyResultQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
	// The status of the cluster. Valid values: running, starting, stopping, and stopped. running indicates the cluster is running, starting indicates the cluster is starting, stopping indicates the cluster is stopping, and stopped indicates the cluster has stopped.
	//
	// example:
	//
	// "starting"
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClustersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *ListClustersResponseBodyResult {
	s.Config = v
	return s
}

func (s *ListClustersResponseBodyResult) SetConfigUpdateTime(v string) *ListClustersResponseBodyResult {
	s.ConfigUpdateTime = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCreateTime(v string) *ListClustersResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetCurrentOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.CurrentOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetDataNode(v *ListClustersResponseBodyResultDataNode) *ListClustersResponseBodyResult {
	s.DataNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetDescription(v string) *ListClustersResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestAdvanceConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestAdvanceConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOfflineDictConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOfflineDictConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetLatestOnlineQueryConfigVersion(v string) *ListClustersResponseBodyResult {
	s.LatestOnlineQueryConfigVersion = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetName(v string) *ListClustersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResult) SetQueryNode(v *ListClustersResponseBodyResultQueryNode) *ListClustersResponseBodyResult {
	s.QueryNode = v
	return s
}

func (s *ListClustersResponseBodyResult) SetStatus(v string) *ListClustersResponseBodyResult {
	s.Status = &v
	return s
}

type ListClustersResponseBodyResultDataNode struct {
	// The name of the node.
	//
	// example:
	//
	// ha-cn-8ed2k7brm05_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The partition ID of the node.
	//
	// example:
	//
	// dt=20220216
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultDataNode) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResultDataNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultDataNode) SetName(v string) *ListClustersResponseBodyResultDataNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetNumber(v int32) *ListClustersResponseBodyResultDataNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultDataNode) SetPartition(v int32) *ListClustersResponseBodyResultDataNode {
	s.Partition = &v
	return s
}

type ListClustersResponseBodyResultQueryNode struct {
	// The name of the node.
	//
	// example:
	//
	// test_0704
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number o replicas.
	//
	// example:
	//
	// dt=20211216
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ListClustersResponseBodyResultQueryNode) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyResultQueryNode) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyResultQueryNode) SetName(v string) *ListClustersResponseBodyResultQueryNode {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetNumber(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Number = &v
	return s
}

func (s *ListClustersResponseBodyResultQueryNode) SetPartition(v int32) *ListClustersResponseBodyResultQueryNode {
	s.Partition = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListDataSourceSchemasResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result []*ListDataSourceSchemasResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBody) SetRequestId(v string) *ListDataSourceSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceSchemasResponseBody) SetResult(v []*ListDataSourceSchemasResponseBodyResult) *ListDataSourceSchemasResponseBody {
	s.Result = v
	return s
}

type ListDataSourceSchemasResponseBodyResult struct {
	// Indicates whether the index properties are added. Valid values: true and false. **true*	- indicates that the index properties are added, and **false*	- indicates that the index properties are not added.
	//
	// example:
	//
	// false
	AddIndex *bool `json:"addIndex,omitempty" xml:"addIndex,omitempty"`
	// Indicates whether it is an attribute field. Valid values: true and false. **true*	- indicates that it is an attribute field, and **false*	- indicates that it is not an attribute field.
	//
	// example:
	//
	// false
	Attribute *bool `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// Indicates whether it is a custom field. Valid values: true and false. **true*	- indicates that it is a custom field, and **false*	- indicates that it is not a custom field.
	//
	// example:
	//
	// false
	Custom *bool `json:"custom,omitempty" xml:"custom,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The primary key.
	PrimaryKey *ListDataSourceSchemasResponseBodyResultPrimaryKey `json:"primaryKey,omitempty" xml:"primaryKey,omitempty" type:"Struct"`
	// Indicates whether the information can be displayed. Valid values: true and false. **true*	- indicates that the information can be displayed, and **false*	- indicates that the information cannot be displayed.
	//
	// example:
	//
	// false
	Summary *bool `json:"summary,omitempty" xml:"summary,omitempty"`
	// The type of the field.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAddIndex(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.AddIndex = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAttribute(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Attribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetCustom(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Custom = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetName(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetPrimaryKey(v *ListDataSourceSchemasResponseBodyResultPrimaryKey) *ListDataSourceSchemasResponseBodyResult {
	s.PrimaryKey = v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetSummary(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetType(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Type = &v
	return s
}

type ListDataSourceSchemasResponseBodyResultPrimaryKey struct {
	// Indicates whether it has the primary key property. **true*	- indicates that it has the primary key property, and **false*	- indicates that it does not have the primary key property.
	//
	// example:
	//
	// false
	HasPrimaryKeyAttribute *bool `json:"hasPrimaryKeyAttribute,omitempty" xml:"hasPrimaryKeyAttribute,omitempty"`
	// Indicates whether it is the primary key. Valid values: true and false. **true*	- indicates that it is the primary key, and **false*	- indicates that it is not the primary key.
	//
	// example:
	//
	// false
	IsPrimaryKey *bool `json:"isPrimaryKey,omitempty" xml:"isPrimaryKey,omitempty"`
	// Indicates whether data is sorted based on the primary key. Valid values: true and false. **true*	- indicates that data is sorted based on the primary key, and **false*	- indicates that data is not sorted based on the primary key.
	//
	// example:
	//
	// false
	IsPrimaryKeySorted *bool `json:"isPrimaryKeySorted,omitempty" xml:"isPrimaryKeySorted,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetHasPrimaryKeyAttribute(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.HasPrimaryKeyAttribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKey(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKey = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKeySorted(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKeySorted = &v
	return s
}

type ListDataSourceSchemasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponse) SetHeaders(v map[string]*string) *ListDataSourceSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSchemasResponse) SetStatusCode(v int32) *ListDataSourceSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSchemasResponse) SetBody(v *ListDataSourceSchemasResponseBody) *ListDataSourceSchemasResponse {
	s.Body = v
	return s
}

type ListDataSourceTasksResponseBody struct {
	// id of request
	//
	// example:
	//
	// CC5EC8FA-5C0D-56AF-BEF4-6FCCEABD0511
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index.
	Result []*ListDataSourceTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBody) SetRequestId(v string) *ListDataSourceTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTasksResponseBody) SetResult(v []*ListDataSourceTasksResponseBodyResult) *ListDataSourceTasksResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTasksResponseBodyResult struct {
	// The additional attributes of the card.
	//
	// example:
	//
	// ""
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	// The field3 field that is passed through when you create a state machine.
	//
	// example:
	//
	// ""
	Field3 *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// fsmId
	//
	// example:
	//
	// tisplus_opensearch@datasource_flow_fsm@1062017779051424-ha-cn-2r42ostoc01_ecom_table@vpc_hz_domain_1@null@MANUAL-ha-cn-2r42ostoc01_ecom_table@1655974525756@006754
	FsmId *string `json:"fsmId,omitempty" xml:"fsmId,omitempty"`
	// Indicates whether the change is a data source task change or a cluster task change.
	//
	// example:
	//
	// " "
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The task name on the card.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The overall status of FSM.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status tag of the progress bar chart.
	Tags []*ListDataSourceTasksResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The information about the task.
	TaskNodes []*ListDataSourceTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	// The timestamp of the task on the card.
	//
	// example:
	//
	// 1646279473
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The type of the task on the card.
	//
	// example:
	//
	// search
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The user who triggered the finite-state machine (FSM) process.
	//
	// example:
	//
	// ""
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResult) SetExtraAttribute(v string) *ListDataSourceTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetField3(v string) *ListDataSourceTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetFsmId(v string) *ListDataSourceTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetGroupType(v string) *ListDataSourceTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetName(v string) *ListDataSourceTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetStatus(v string) *ListDataSourceTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTags(v []*ListDataSourceTasksResponseBodyResultTags) *ListDataSourceTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTaskNodes(v []*ListDataSourceTasksResponseBodyResultTaskNodes) *ListDataSourceTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTime(v string) *ListDataSourceTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetType(v string) *ListDataSourceTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetUser(v string) *ListDataSourceTasksResponseBodyResult {
	s.User = &v
	return s
}

type ListDataSourceTasksResponseBodyResultTags struct {
	// The content of the tag.
	//
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The level of the tag.
	//
	// example:
	//
	// ""
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetMsg(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetTagLevel(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

type ListDataSourceTasksResponseBodyResultTaskNodes struct {
	// The date when the task was completed.
	//
	// example:
	//
	// ""
	FinishDate *string `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
	// The sequence number of the task.
	//
	// example:
	//
	// 100
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// ha-cn-7pp2ngv4s02_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// onlyPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetName(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

type ListDataSourceTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponse) SetHeaders(v map[string]*string) *ListDataSourceTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTasksResponse) SetStatusCode(v int32) *ListDataSourceTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTasksResponse) SetBody(v *ListDataSourceTasksResponseBody) *ListDataSourceTasksResponse {
	s.Body = v
	return s
}

type ListDataSourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result []*ListDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetResult(v []*ListDataSourcesResponseBodyResult) *ListDataSourcesResponseBody {
	s.Result = v
	return s
}

type ListDataSourcesResponseBodyResult struct {
	// The data center where the data source is deployed in offline mode.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The information about indexes.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The time when an index for full data was last built.
	LastFulTime *int64 `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// data_source_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The state of the data source.
	//
	// example:
	//
	// new
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyResult) SetDomain(v string) *ListDataSourcesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetIndexes(v []*string) *ListDataSourcesResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetLastFulTime(v int64) *ListDataSourcesResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetName(v string) *ListDataSourcesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetStatus(v string) *ListDataSourcesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetType(v string) *ListDataSourcesResponseBodyResult {
	s.Type = &v
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

type ListDateSourceGenerationsRequest struct {
	// The data center where the data source is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj_vpc_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The valid state of the data source. Valid values: true and false. The default value of this parameter is true.
	//
	// 1.  true indicates that the generations that have not expired and of which the tasks have been executed are returned.
	//
	// 2.  false indicates that all generations are returned.
	//
	// example:
	//
	// true
	ValidStatus *bool `json:"validStatus,omitempty" xml:"validStatus,omitempty"`
}

func (s ListDateSourceGenerationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsRequest) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsRequest) SetDomainName(v string) *ListDateSourceGenerationsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDateSourceGenerationsRequest) SetValidStatus(v bool) *ListDateSourceGenerationsRequest {
	s.ValidStatus = &v
	return s
}

type ListDateSourceGenerationsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListDateSourceGenerationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDateSourceGenerationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBody) SetRequestId(v string) *ListDateSourceGenerationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBody) SetResult(v []*ListDateSourceGenerationsResponseBodyResult) *ListDateSourceGenerationsResponseBody {
	s.Result = v
	return s
}

type ListDateSourceGenerationsResponseBodyResult struct {
	// buildDeployId
	//
	// example:
	//
	// 122
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The time to start index building.
	//
	// example:
	//
	// 1626143673
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The directory where the index file created by using the dump table is saved.
	//
	// example:
	//
	// ""
	DataDumpRoot *string `json:"dataDumpRoot,omitempty" xml:"dataDumpRoot,omitempty"`
	// The primary key of the generation.
	//
	// example:
	//
	// 1626143930
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// Key indicates the name of the index. value indicates the number of shards.
	Partition map[string]*int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The status.
	//
	// example:
	//
	// STOPPED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the offline indexing was initiated.
	//
	// example:
	//
	// 1626143673
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListDateSourceGenerationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetBuildDeployId(v int32) *ListDateSourceGenerationsResponseBodyResult {
	s.BuildDeployId = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetCreateTime(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetDataDumpRoot(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.DataDumpRoot = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetGeneration(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetPartition(v map[string]*int32) *ListDateSourceGenerationsResponseBodyResult {
	s.Partition = v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetStatus(v string) *ListDateSourceGenerationsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDateSourceGenerationsResponseBodyResult) SetTimestamp(v int64) *ListDateSourceGenerationsResponseBodyResult {
	s.Timestamp = &v
	return s
}

type ListDateSourceGenerationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDateSourceGenerationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDateSourceGenerationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDateSourceGenerationsResponse) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsResponse) SetHeaders(v map[string]*string) *ListDateSourceGenerationsResponse {
	s.Headers = v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetStatusCode(v int32) *ListDateSourceGenerationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDateSourceGenerationsResponse) SetBody(v *ListDateSourceGenerationsResponseBody) *ListDateSourceGenerationsResponse {
	s.Body = v
	return s
}

type ListIndexRecoverRecordsResponseBody struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// 2024-06-07 16:43:00
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// generationId
	//
	// example:
	//
	// 1708674867
	GenerationId *string `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ListIndexRecoverRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexRecoverRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexRecoverRecordsResponseBody) SetDesc(v string) *ListIndexRecoverRecordsResponseBody {
	s.Desc = &v
	return s
}

func (s *ListIndexRecoverRecordsResponseBody) SetFinishedTime(v string) *ListIndexRecoverRecordsResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *ListIndexRecoverRecordsResponseBody) SetGenerationId(v string) *ListIndexRecoverRecordsResponseBody {
	s.GenerationId = &v
	return s
}

type ListIndexRecoverRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexRecoverRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexRecoverRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexRecoverRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListIndexRecoverRecordsResponse) SetHeaders(v map[string]*string) *ListIndexRecoverRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListIndexRecoverRecordsResponse) SetStatusCode(v int32) *ListIndexRecoverRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexRecoverRecordsResponse) SetBody(v *ListIndexRecoverRecordsResponseBody) *ListIndexRecoverRecordsResponse {
	s.Body = v
	return s
}

type ListIndexesRequest struct {
	// 是否为新版本控制台页面
	//
	// example:
	//
	// true
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
}

func (s ListIndexesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexesRequest) SetNewMode(v bool) *ListIndexesRequest {
	s.NewMode = &v
	return s
}

type ListIndexesResponseBody struct {
	// id of request
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index list.
	Result []*ListIndexesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBody) SetRequestId(v string) *ListIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexesResponseBody) SetResult(v []*ListIndexesResponseBodyResult) *ListIndexesResponseBody {
	s.Result = v
	return s
}

type ListIndexesResponseBodyResult struct {
	// schema JSON
	//
	// example:
	//
	// {"summarys":{"parameter":{"file_compressor":"zstd"},"summary_fields":["id"]},"file_compress":[{"name":"file_compressor","type":"zstd"},{"name":"no_compressor","type":""}],"indexs":[{"index_fields":"name","index_name":"ids","index_type":"STRING"},{"has_primary_key_attribute":true,"index_fields":"id","is_primary_key_sorted":false,"index_name":"id","index_type":"PRIMARYKEY64"}],"attributes":[{"file_compress":"no_compressor","field_name":"id"}],"fields":[{"user_defined_param":{},"compress_type":"uniq","field_type":"STRING","field_name":"id"},{"compress_type":"uniq","field_type":"STRING","field_name":"name"}],"table_name":"api"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The data source.
	//
	// example:
	//
	// ha-cn-7mz2kvu2c01_table4
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// 数据源相关信息
	DataSourceInfo *ListIndexesResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// 备注
	//
	// example:
	//
	// 备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the data center where the data source is deployed.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 全量切换时间
	//
	// example:
	//
	// 2023-07-05 10:40:38
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// 全量版本  即：索引版本
	//
	// example:
	//
	// 1688523414
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// 增量更新时间
	//
	// example:
	//
	// 2023-07-05 10:58:33
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// 索引大小
	//
	// example:
	//
	// 4689
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// NEW, PUBLISH
	//
	// example:
	//
	// " "
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据分片
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The list of version information.
	Versions []*ListIndexesResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResult) SetContent(v string) *ListIndexesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSource(v string) *ListIndexesResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSourceInfo(v *ListIndexesResponseBodyResultDataSourceInfo) *ListIndexesResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDescription(v string) *ListIndexesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDomain(v string) *ListIndexesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullVersion(v int64) *ListIndexesResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIncUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexSize(v int64) *ListIndexesResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexStatus(v string) *ListIndexesResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetName(v string) *ListIndexesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetPartition(v int32) *ListIndexesResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetVersions(v []*ListIndexesResponseBodyResultVersions) *ListIndexesResponseBodyResult {
	s.Versions = v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfo struct {
	// 是否开启自动全量
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// odps 数据源配置
	Config *ListIndexesResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// 离线部署
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 数据源名
	//
	// example:
	//
	// index1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据更新资源数
	//
	// example:
	//
	// 2
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// saro数据源配置
	SaroConfig *ListIndexesResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// 数据源类型 (odps, swift, saro, oss, unKnow)
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *ListIndexesResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetConfig(v *ListIndexesResponseBodyResultDataSourceInfoConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetDomain(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetName(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *ListIndexesResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetSaroConfig(v *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetType(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfoConfig struct {
	// odps数据源ak
	//
	// example:
	//
	// root
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// odps数据源ak secret
	//
	// example:
	//
	// root123
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// oss命名空间
	//
	// example:
	//
	// ha3test-oss
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// odps相关
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// saro相关
	//
	// example:
	//
	// TEST_dump_demo_sj_na61hunbu2_share_holo
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// oss数据源相关
	//
	// example:
	//
	// /test_opensearch/sift_oss_test.data
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// 数据分片
	//
	// example:
	//
	// ds=12345
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// hdfs相关
	//
	// example:
	//
	// http://test_opensearch/sift_oss_test.data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// odps数据源项目名称
	//
	// example:
	//
	// tisplus_dev
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// saro、odps相关
	//
	// example:
	//
	// dump_odps_demo
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetProject(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetTable(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

type ListIndexesResponseBodyResultDataSourceInfoSaroConfig struct {
	// saro数据源的namespace
	//
	// example:
	//
	// TEST_dump_demo_sj_na61hunbu2_share_holo
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// saro数据表名称
	//
	// example:
	//
	// dump_odps_demo
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type ListIndexesResponseBodyResultVersions struct {
	// The description.
	//
	// example:
	//
	// close alarm, by 3.9.2 hotfix workflow
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The list of file names.
	Files []*ListIndexesResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the version.
	//
	// example:
	//
	// ha-cn-7pp2ngv4s02_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The state of the version. Valid values: drafting, used, unused and trash. drafting indicates that the version is a draft, used indicates that the version is used online, unused indicates that the version is not used, and trash indicates that the version is being deleted.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the version was updated.
	//
	// example:
	//
	// " "
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the version. The value of this parameter is null for the edit version.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListIndexesResponseBodyResultVersions) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersions) SetDesc(v string) *ListIndexesResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetFiles(v []*ListIndexesResponseBodyResultVersionsFiles) *ListIndexesResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetName(v string) *ListIndexesResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetStatus(v string) *ListIndexesResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetUpdateTime(v int64) *ListIndexesResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetVersionId(v int32) *ListIndexesResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

type ListIndexesResponseBodyResultVersionsFiles struct {
	// The name of the directory for the index.
	//
	// example:
	//
	// " "
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether a directory exists.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether it is a template.
	//
	// example:
	//
	// True
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// ha-cn-7mz2iv7sq01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIndexesResponseBodyResultVersionsFiles) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetFullPathName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsDir(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

type ListIndexesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexesResponse) SetHeaders(v map[string]*string) *ListIndexesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexesResponse) SetStatusCode(v int32) *ListIndexesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexesResponse) SetBody(v *ListIndexesResponseBody) *ListIndexesResponse {
	s.Body = v
	return s
}

type ListInstanceSpecsRequest struct {
	// The node type. Valid values: qrs, search, index, and cluster. qrs indicates a query node, search indicates a data node, index indicates an index node, and cluster indicates a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// search
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsRequest) SetType(v string) *ListInstanceSpecsRequest {
	s.Type = &v
	return s
}

type ListInstanceSpecsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListInstanceSpecsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBody) SetRequestId(v string) *ListInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSpecsResponseBody) SetResult(v []*ListInstanceSpecsResponseBodyResult) *ListInstanceSpecsResponseBody {
	s.Result = v
	return s
}

type ListInstanceSpecsResponseBodyResult struct {
	// None
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 单数据节点存储空间最大值
	MaxDisk *int32 `json:"maxDisk,omitempty" xml:"maxDisk,omitempty"`
	// Unit: GB
	//
	// example:
	//
	// 4
	Mem *int32 `json:"mem,omitempty" xml:"mem,omitempty"`
	// 单数据节点存储空间最小值
	MinDisk *int32 `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
}

func (s ListInstanceSpecsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponseBodyResult) SetCpu(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Cpu = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMaxDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MaxDisk = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMem(v int32) *ListInstanceSpecsResponseBodyResult {
	s.Mem = &v
	return s
}

func (s *ListInstanceSpecsResponseBodyResult) SetMinDisk(v int32) *ListInstanceSpecsResponseBodyResult {
	s.MinDisk = &v
	return s
}

type ListInstanceSpecsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsResponse) SetHeaders(v map[string]*string) *ListInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSpecsResponse) SetStatusCode(v int32) *ListInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSpecsResponse) SetBody(v *ListInstanceSpecsResponseBody) *ListInstanceSpecsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The description of the instance. You can use this description to filter instances. Fuzzy match is supported.
	//
	// example:
	//
	// Havenask instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Instance type, vector (vector index version),engine (recall engine version)
	//
	// example:
	//
	// vector
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ha-cn-83570439y0n
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzgpiswzbksdi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tag dictionary.
	Tags []*ListInstancesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetDescription(v string) *ListInstancesRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesRequest) SetEdition(v string) *ListInstancesRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
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

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v []*ListInstancesRequestTags) *ListInstancesRequest {
	s.Tags = v
	return s
}

type ListInstancesRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// backup
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// oboms-disk
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTags) SetKey(v string) *ListInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTags) SetValue(v string) *ListInstancesRequestTags {
	s.Value = &v
	return s
}

type ListInstancesShrinkRequest struct {
	// The description of the instance. You can use this description to filter instances. Fuzzy match is supported.
	//
	// example:
	//
	// Havenask instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Instance type, vector (vector index version),engine (recall engine version)
	//
	// example:
	//
	// vector
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ha-cn-83570439y0n
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzgpiswzbksdi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tag dictionary.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) SetDescription(v string) *ListInstancesShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetEdition(v string) *ListInstancesShrinkRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceId(v string) *ListInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagsShrink(v string) *ListInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListInstancesResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 89B968E6-1E41-58DF-BB25-5F98ECC759CE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result []*ListInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetResult(v []*ListInstancesResponseBodyResult) *ListInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyResult struct {
	// The billing method
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The product code
	//
	// example:
	//
	// ""
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created
	//
	// example:
	//
	// 2022-06-04T02:03:21Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance
	//
	// example:
	//
	// Emergency test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time
	//
	// example:
	//
	// 1634885083
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The ID of the resource
	//
	// example:
	//
	// ha-cn-2r42n8oh001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Information about the instance of the network search engine
	Network *ListInstancesResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The ID of the resource group
	//
	// example:
	//
	// rg-aekzgpiswzbksdi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the instance
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The result returned.
	Tags []*ListInstancesResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was last updated
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResult) SetChargeType(v string) *ListInstancesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCommodityCode(v string) *ListInstancesResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCreateTime(v string) *ListInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetDescription(v string) *ListInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetExpiredTime(v string) *ListInstancesResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInDebt(v bool) *ListInstancesResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInstanceId(v string) *ListInstancesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetLockMode(v string) *ListInstancesResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetNetwork(v *ListInstancesResponseBodyResultNetwork) *ListInstancesResponseBodyResult {
	s.Network = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetResourceGroupId(v string) *ListInstancesResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetStatus(v string) *ListInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetTags(v []*ListInstancesResponseBodyResultTags) *ListInstancesResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetUpdateTime(v string) *ListInstancesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type ListInstancesResponseBodyResultNetwork struct {
	// The access point of the gateway
	//
	// example:
	//
	// ""
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The ID of the virtual switch
	//
	// example:
	//
	// vsw-bp11ldcf59q2nbwkqgj6z
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) network
	//
	// example:
	//
	// vpc-wz9axk41d9vffoc79x0oe
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListInstancesResponseBodyResultNetwork) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultNetwork) SetEndpoint(v string) *ListInstancesResponseBodyResultNetwork {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVSwitchId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVpcId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VpcId = &v
	return s
}

type ListInstancesResponseBodyResultTags struct {
	// The key of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// oboms-disk
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultTags) SetKey(v string) *ListInstancesResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyResultTags) SetValue(v string) *ListInstancesResponseBodyResultTags {
	s.Value = &v
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

type ListOnlineConfigsRequest struct {
	// The name of the domain
	//
	// This parameter is required.
	//
	// example:
	//
	// sz_vpc_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListOnlineConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsRequest) SetDomain(v string) *ListOnlineConfigsRequest {
	s.Domain = &v
	return s
}

type ListOnlineConfigsResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListOnlineConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListOnlineConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBody) SetRequestId(v string) *ListOnlineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineConfigsResponseBody) SetResult(v []*ListOnlineConfigsResponseBodyResult) *ListOnlineConfigsResponseBody {
	s.Result = v
	return s
}

type ListOnlineConfigsResponseBodyResult struct {
	// The configuration information
	//
	// example:
	//
	// {\\"specItems\\":[{\\"specKey\\":\\"YQ_KEYWORD_NUMBER_PLUS\\",\\"value\\":\\"1\\"}]}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the index
	//
	// example:
	//
	// generation
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s ListOnlineConfigsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBodyResult) SetConfig(v string) *ListOnlineConfigsResponseBodyResult {
	s.Config = &v
	return s
}

func (s *ListOnlineConfigsResponseBodyResult) SetIndexName(v string) *ListOnlineConfigsResponseBodyResult {
	s.IndexName = &v
	return s
}

type ListOnlineConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponse) SetHeaders(v map[string]*string) *ListOnlineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineConfigsResponse) SetStatusCode(v int32) *ListOnlineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineConfigsResponse) SetBody(v *ListOnlineConfigsResponseBody) *ListOnlineConfigsResponse {
	s.Body = v
	return s
}

type ListPausePolicysResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	Result map[string]*ResultValue `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPausePolicysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPausePolicysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPausePolicysResponseBody) SetRequestId(v string) *ListPausePolicysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPausePolicysResponseBody) SetResult(v map[string]*ResultValue) *ListPausePolicysResponseBody {
	s.Result = v
	return s
}

type ListPausePolicysResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPausePolicysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPausePolicysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPausePolicysResponse) GoString() string {
	return s.String()
}

func (s *ListPausePolicysResponse) SetHeaders(v map[string]*string) *ListPausePolicysResponse {
	s.Headers = v
	return s
}

func (s *ListPausePolicysResponse) SetStatusCode(v int32) *ListPausePolicysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPausePolicysResponse) SetBody(v *ListPausePolicysResponseBody) *ListPausePolicysResponse {
	s.Body = v
	return s
}

type ListPostQueryResultRequest struct {
	// example:
	//
	// 127.0.0.1
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// ha3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPostQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPostQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultRequest) SetAddress(v string) *ListPostQueryResultRequest {
	s.Address = &v
	return s
}

func (s *ListPostQueryResultRequest) SetBody(v map[string]interface{}) *ListPostQueryResultRequest {
	s.Body = v
	return s
}

func (s *ListPostQueryResultRequest) SetType(v string) *ListPostQueryResultRequest {
	s.Type = &v
	return s
}

type ListPostQueryResultResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListPostQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPostQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultResponseBody) SetRequestId(v string) *ListPostQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPostQueryResultResponseBody) SetResult(v interface{}) *ListPostQueryResultResponseBody {
	s.Result = v
	return s
}

type ListPostQueryResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPostQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPostQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPostQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultResponse) SetHeaders(v map[string]*string) *ListPostQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListPostQueryResultResponse) SetStatusCode(v int32) *ListPostQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPostQueryResultResponse) SetBody(v *ListPostQueryResultResponseBody) *ListPostQueryResultResponse {
	s.Body = v
	return s
}

type ListQueryResultRequest struct {
	// The query statement
	//
	// example:
	//
	// query%3D1%26%26config%3Dstart%3A0%2Chit%3A10%2Cformat%3Ajson%26%26cluster%3Dgeneral
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The SQL statement that is executed in the query
	//
	// example:
	//
	// query%3Dselect%20max(content_id)%20from%20generation
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s ListQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListQueryResultRequest) SetQuery(v string) *ListQueryResultRequest {
	s.Query = &v
	return s
}

func (s *ListQueryResultRequest) SetSql(v string) *ListQueryResultRequest {
	s.Sql = &v
	return s
}

type ListQueryResultResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 9E5BCFAA-98B3-51D0-9188-B1BC07589337
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponseBody) SetRequestId(v string) *ListQueryResultResponseBody {
	s.RequestId = &v
	return s
}

type ListQueryResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListQueryResultResponse) SetHeaders(v map[string]*string) *ListQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListQueryResultResponse) SetStatusCode(v int32) *ListQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryResultResponse) SetBody(v *ListQueryResultResponseBody) *ListQueryResultResponse {
	s.Body = v
	return s
}

type ListRestQueryResultRequest struct {
	// example:
	//
	// 127.0.0.1
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// example:
	//
	// query%3Drelation_id%3A%221151274675_2%22%26%26cluster%3Dgeneral%26%26config%3Dstart%3A0%2Chit%3A10%2Cformat%3Ajson
	Query map[string]interface{} `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListRestQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRestQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultRequest) SetAddress(v string) *ListRestQueryResultRequest {
	s.Address = &v
	return s
}

func (s *ListRestQueryResultRequest) SetIndexName(v string) *ListRestQueryResultRequest {
	s.IndexName = &v
	return s
}

func (s *ListRestQueryResultRequest) SetQuery(v map[string]interface{}) *ListRestQueryResultRequest {
	s.Query = v
	return s
}

type ListRestQueryResultResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListRestQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRestQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultResponseBody) SetRequestId(v string) *ListRestQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRestQueryResultResponseBody) SetResult(v interface{}) *ListRestQueryResultResponseBody {
	s.Result = v
	return s
}

type ListRestQueryResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRestQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRestQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRestQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultResponse) SetHeaders(v map[string]*string) *ListRestQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListRestQueryResultResponse) SetStatusCode(v int32) *ListRestQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRestQueryResultResponse) SetBody(v *ListRestQueryResultResponseBody) *ListRestQueryResultResponse {
	s.Body = v
	return s
}

type ListSchemasRequest struct {
	// AK
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// odps endpoint
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// igraph-cn-tl32wnrhi04
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// dt=20230520
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// start-flask-v3-obcc
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// item
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// odps, swift, saro, oss, unKnow
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasRequest) SetAccessKey(v string) *ListSchemasRequest {
	s.AccessKey = &v
	return s
}

func (s *ListSchemasRequest) SetAccessSecret(v string) *ListSchemasRequest {
	s.AccessSecret = &v
	return s
}

func (s *ListSchemasRequest) SetEndpoint(v string) *ListSchemasRequest {
	s.Endpoint = &v
	return s
}

func (s *ListSchemasRequest) SetNamespace(v string) *ListSchemasRequest {
	s.Namespace = &v
	return s
}

func (s *ListSchemasRequest) SetPartition(v string) *ListSchemasRequest {
	s.Partition = &v
	return s
}

func (s *ListSchemasRequest) SetProject(v string) *ListSchemasRequest {
	s.Project = &v
	return s
}

func (s *ListSchemasRequest) SetTable(v string) *ListSchemasRequest {
	s.Table = &v
	return s
}

func (s *ListSchemasRequest) SetType(v string) *ListSchemasRequest {
	s.Type = &v
	return s
}

type ListSchemasResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBody) SetRequestId(v string) *ListSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemasResponseBody) SetResult(v interface{}) *ListSchemasResponseBody {
	s.Result = v
	return s
}

type ListSchemasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListSchemasResponse) SetHeaders(v map[string]*string) *ListSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListSchemasResponse) SetStatusCode(v int32) *ListSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemasResponse) SetBody(v *ListSchemasResponseBody) *ListSchemasResponse {
	s.Body = v
	return s
}

type ListTableGenerationsResponseBody struct {
	// requestId
	//
	// example:
	//
	// F6E3D968-529C-5C40-AFDD-133A8B8FD930
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListTableGenerationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListTableGenerationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTableGenerationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponseBody) SetRequestId(v string) *ListTableGenerationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableGenerationsResponseBody) SetResult(v []*ListTableGenerationsResponseBodyResult) *ListTableGenerationsResponseBody {
	s.Result = v
	return s
}

type ListTableGenerationsResponseBodyResult struct {
	// generationId
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ListTableGenerationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTableGenerationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponseBodyResult) SetGenerationId(v int64) *ListTableGenerationsResponseBodyResult {
	s.GenerationId = &v
	return s
}

type ListTableGenerationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableGenerationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableGenerationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTableGenerationsResponse) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponse) SetHeaders(v map[string]*string) *ListTableGenerationsResponse {
	s.Headers = v
	return s
}

func (s *ListTableGenerationsResponse) SetStatusCode(v int32) *ListTableGenerationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableGenerationsResponse) SetBody(v *ListTableGenerationsResponseBody) *ListTableGenerationsResponse {
	s.Body = v
	return s
}

type ListTablesRequest struct {
	// example:
	//
	// true
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
}

func (s ListTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) SetNewMode(v bool) *ListTablesRequest {
	s.NewMode = &v
	return s
}

type ListTablesResponseBody struct {
	// requestId
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListTablesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetResult(v []*ListTablesResponseBodyResult) *ListTablesResponseBody {
	s.Result = v
	return s
}

type ListTablesResponseBodyResult struct {
	// example:
	//
	// IN_USE
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// example:
	//
	// es_test_1b
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListTablesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyResult) SetIndexStatus(v string) *ListTablesResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *ListTablesResponseBodyResult) SetName(v string) *ListTablesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyResult) SetStatus(v string) *ListTablesResponseBodyResult {
	s.Status = &v
	return s
}

type ListTablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetStatusCode(v int32) *ListTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// example:
	//
	// 1690423741577
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1687238865434
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetEnd(v int64) *ListTasksRequest {
	s.End = &v
	return s
}

func (s *ListTasksRequest) SetStart(v int64) *ListTasksRequest {
	s.Start = &v
	return s
}

type ListTasksResponseBody struct {
	// id of request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
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

func (s *ListTasksResponseBody) SetResult(v interface{}) *ListTasksResponseBody {
	s.Result = v
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

type ListVectorQueryResultRequest struct {
	// example:
	//
	// 127.0.0.1
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// primary_key
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// example:
	//
	// image
	VectorQueryType *string `json:"vectorQueryType,omitempty" xml:"vectorQueryType,omitempty"`
}

func (s ListVectorQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVectorQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultRequest) SetAddress(v string) *ListVectorQueryResultRequest {
	s.Address = &v
	return s
}

func (s *ListVectorQueryResultRequest) SetBody(v map[string]interface{}) *ListVectorQueryResultRequest {
	s.Body = v
	return s
}

func (s *ListVectorQueryResultRequest) SetQueryType(v string) *ListVectorQueryResultRequest {
	s.QueryType = &v
	return s
}

func (s *ListVectorQueryResultRequest) SetVectorQueryType(v string) *ListVectorQueryResultRequest {
	s.VectorQueryType = &v
	return s
}

type ListVectorQueryResultResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListVectorQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVectorQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultResponseBody) SetRequestId(v string) *ListVectorQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVectorQueryResultResponseBody) SetResult(v interface{}) *ListVectorQueryResultResponseBody {
	s.Result = v
	return s
}

type ListVectorQueryResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVectorQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVectorQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVectorQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultResponse) SetHeaders(v map[string]*string) *ListVectorQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListVectorQueryResultResponse) SetStatusCode(v int32) *ListVectorQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVectorQueryResultResponse) SetBody(v *ListVectorQueryResultResponseBody) *ListVectorQueryResultResponse {
	s.Body = v
	return s
}

type ModifyAdvanceConfigRequest struct {
	// example:
	//
	// ""
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// FILE
	ContentType *string                            `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Desc        *string                            `json:"desc,omitempty" xml:"desc,omitempty"`
	Files       []*ModifyAdvanceConfigRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// used
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024-02-27T07:50:55Z
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ModifyAdvanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigRequest) SetContent(v string) *ModifyAdvanceConfigRequest {
	s.Content = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetContentType(v string) *ModifyAdvanceConfigRequest {
	s.ContentType = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetDesc(v string) *ModifyAdvanceConfigRequest {
	s.Desc = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetFiles(v []*ModifyAdvanceConfigRequestFiles) *ModifyAdvanceConfigRequest {
	s.Files = v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetName(v string) *ModifyAdvanceConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetStatus(v string) *ModifyAdvanceConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetUpdateTime(v int64) *ModifyAdvanceConfigRequest {
	s.UpdateTime = &v
	return s
}

type ModifyAdvanceConfigRequestFiles struct {
	// example:
	//
	// /cluster.json
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyAdvanceConfigRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigRequestFiles) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigRequestFiles) SetFullPathName(v string) *ModifyAdvanceConfigRequestFiles {
	s.FullPathName = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetIsDir(v bool) *ModifyAdvanceConfigRequestFiles {
	s.IsDir = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetIsTemplate(v bool) *ModifyAdvanceConfigRequestFiles {
	s.IsTemplate = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetName(v string) *ModifyAdvanceConfigRequestFiles {
	s.Name = &v
	return s
}

type ModifyAdvanceConfigResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigResponseBody) SetRequestId(v string) *ModifyAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *ModifyAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type ModifyAdvanceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigResponse) SetHeaders(v map[string]*string) *ModifyAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvanceConfigResponse) SetStatusCode(v int32) *ModifyAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvanceConfigResponse) SetBody(v *ModifyAdvanceConfigResponseBody) *ModifyAdvanceConfigResponse {
	s.Body = v
	return s
}

type ModifyAdvanceConfigFileRequest struct {
	// The content of the file.
	//
	// example:
	//
	// "ha3"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The variable.
	Variables map[string]*VariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// /qrs.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyAdvanceConfigFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileRequest) SetContent(v string) *ModifyAdvanceConfigFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetVariables(v map[string]*VariablesValue) *ModifyAdvanceConfigFileRequest {
	s.Variables = v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetFileName(v string) *ModifyAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

type ModifyAdvanceConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAdvanceConfigFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponseBody) SetRequestId(v string) *ModifyAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponseBody) SetResult(v map[string]interface{}) *ModifyAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

type ModifyAdvanceConfigFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvanceConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvanceConfigFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponse) SetHeaders(v map[string]*string) *ModifyAdvanceConfigFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetStatusCode(v int32) *ModifyAdvanceConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponse) SetBody(v *ModifyAdvanceConfigFileResponseBody) *ModifyAdvanceConfigFileResponse {
	s.Body = v
	return s
}

type ModifyClusterDescRequest struct {
	// The parameters in the request body
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDescRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescRequest) SetBody(v map[string]interface{}) *ModifyClusterDescRequest {
	s.Body = v
	return s
}

type ModifyClusterDescResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterDescResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponseBody) SetRequestId(v string) *ModifyClusterDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterDescResponseBody) SetResult(v map[string]interface{}) *ModifyClusterDescResponseBody {
	s.Result = v
	return s
}

type ModifyClusterDescResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterDescResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDescResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDescResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescResponse) SetHeaders(v map[string]*string) *ModifyClusterDescResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDescResponse) SetStatusCode(v int32) *ModifyClusterDescResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterDescResponse) SetBody(v *ModifyClusterDescResponseBody) *ModifyClusterDescResponse {
	s.Body = v
	return s
}

type ModifyClusterOfflineConfigRequest struct {
	// The mode of reindexing. Valid values: api and indexRecover. api indicates to push incremental data to a data source by calling the API operations. indexRecover indicates that the data source is restored from the index.
	//
	// example:
	//
	// indexRecover
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The configuration name, which is stored as a key.
	Config map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test1
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source. Valid values: odps, swift, saro, and unKnow.
	//
	// example:
	//
	// maxComputer
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// This parameter is required when index building by using API data sources is triggered.
	//
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The domain where the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The data restoration version.
	//
	// example:
	//
	// 160142641
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// This parameter is required when index building for full data in a MaxCompute data source is triggered.
	//
	// example:
	//
	// 20211202
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	PushMode  *string `json:"pushMode,omitempty" xml:"pushMode,omitempty"`
}

func (s ModifyClusterOfflineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigRequest) SetBuildMode(v string) *ModifyClusterOfflineConfigRequest {
	s.BuildMode = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOfflineConfigRequest {
	s.Config = v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceName(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceType(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataTimeSec(v int32) *ModifyClusterOfflineConfigRequest {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDomain(v string) *ModifyClusterOfflineConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetGeneration(v int64) *ModifyClusterOfflineConfigRequest {
	s.Generation = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPartition(v string) *ModifyClusterOfflineConfigRequest {
	s.Partition = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPushMode(v string) *ModifyClusterOfflineConfigRequest {
	s.PushMode = &v
	return s
}

type ModifyClusterOfflineConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOfflineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponseBody) SetRequestId(v string) *ModifyClusterOfflineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOfflineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyClusterOfflineConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterOfflineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterOfflineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOfflineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOfflineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetStatusCode(v int32) *ModifyClusterOfflineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOfflineConfigResponse) SetBody(v *ModifyClusterOfflineConfigResponseBody) *ModifyClusterOfflineConfigResponse {
	s.Body = v
	return s
}

type ModifyClusterOnlineConfigRequest struct {
	// The information about the cluster
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// 配置信息
	Config map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
}

func (s ModifyClusterOnlineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigRequest) SetClusters(v []*string) *ModifyClusterOnlineConfigRequest {
	s.Clusters = v
	return s
}

func (s *ModifyClusterOnlineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOnlineConfigRequest {
	s.Config = v
	return s
}

type ModifyClusterOnlineConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyClusterOnlineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponseBody) SetRequestId(v string) *ModifyClusterOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyClusterOnlineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyClusterOnlineConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterOnlineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetStatusCode(v int32) *ModifyClusterOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterOnlineConfigResponse) SetBody(v *ModifyClusterOnlineConfigResponseBody) *ModifyClusterOnlineConfigResponse {
	s.Body = v
	return s
}

type ModifyDataSourceRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the data source is created by using the dryRun feature. This parameter only checks whether the data source is valid. Valid values: true and false. true indicates that the data source is created by using the dryRun feature, and false indicates that the data source is not created by using the dryRun feature.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceRequest) SetBody(v map[string]interface{}) *ModifyDataSourceRequest {
	s.Body = v
	return s
}

func (s *ModifyDataSourceRequest) SetDryRun(v bool) *ModifyDataSourceRequest {
	s.DryRun = &v
	return s
}

type ModifyDataSourceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceResponseBody) SetResult(v map[string]interface{}) *ModifyDataSourceResponseBody {
	s.Result = v
	return s
}

type ModifyDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetStatusCode(v int32) *ModifyDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
}

type ModifyDataSourceDeployRequest struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                                   `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Extend         *ModifyDataSourceDeployRequestExtend    `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	Processor      *ModifyDataSourceDeployRequestProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	Storage        *ModifyDataSourceDeployRequestStorage   `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	Swift          *ModifyDataSourceDeployRequestSwift     `json:"swift,omitempty" xml:"swift,omitempty" type:"Struct"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// generationId
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ModifyDataSourceDeployRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequest) SetAutoBuildIndex(v bool) *ModifyDataSourceDeployRequest {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetExtend(v *ModifyDataSourceDeployRequestExtend) *ModifyDataSourceDeployRequest {
	s.Extend = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetProcessor(v *ModifyDataSourceDeployRequestProcessor) *ModifyDataSourceDeployRequest {
	s.Processor = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetStorage(v *ModifyDataSourceDeployRequestStorage) *ModifyDataSourceDeployRequest {
	s.Storage = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetSwift(v *ModifyDataSourceDeployRequestSwift) *ModifyDataSourceDeployRequest {
	s.Swift = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetDryRun(v bool) *ModifyDataSourceDeployRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetGenerationId(v int64) *ModifyDataSourceDeployRequest {
	s.GenerationId = &v
	return s
}

type ModifyDataSourceDeployRequestExtend struct {
	Hdfs *ModifyDataSourceDeployRequestExtendHdfs `json:"hdfs,omitempty" xml:"hdfs,omitempty" type:"Struct"`
	Odps *ModifyDataSourceDeployRequestExtendOdps `json:"odps,omitempty" xml:"odps,omitempty" type:"Struct"`
	Oss  *ModifyDataSourceDeployRequestExtendOss  `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
	Saro *ModifyDataSourceDeployRequestExtendSaro `json:"saro,omitempty" xml:"saro,omitempty" type:"Struct"`
}

func (s ModifyDataSourceDeployRequestExtend) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtend) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtend) SetHdfs(v *ModifyDataSourceDeployRequestExtendHdfs) *ModifyDataSourceDeployRequestExtend {
	s.Hdfs = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetOdps(v *ModifyDataSourceDeployRequestExtendOdps) *ModifyDataSourceDeployRequestExtend {
	s.Odps = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetOss(v *ModifyDataSourceDeployRequestExtendOss) *ModifyDataSourceDeployRequestExtend {
	s.Oss = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetSaro(v *ModifyDataSourceDeployRequestExtendSaro) *ModifyDataSourceDeployRequestExtend {
	s.Saro = v
	return s
}

type ModifyDataSourceDeployRequestExtendHdfs struct {
	// example:
	//
	// ymsh-service/src/main/java/cn/ymsh/util/jd
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendHdfs) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendHdfs) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendHdfs) SetPath(v string) *ModifyDataSourceDeployRequestExtendHdfs {
	s.Path = &v
	return s
}

type ModifyDataSourceDeployRequestExtendOdps struct {
	Partitions map[string]*string `json:"partitions,omitempty" xml:"partitions,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendOdps) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendOdps) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendOdps) SetPartitions(v map[string]*string) *ModifyDataSourceDeployRequestExtendOdps {
	s.Partitions = v
	return s
}

type ModifyDataSourceDeployRequestExtendOss struct {
	// example:
	//
	// oss://test
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendOss) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendOss) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendOss) SetPath(v string) *ModifyDataSourceDeployRequestExtendOss {
	s.Path = &v
	return s
}

type ModifyDataSourceDeployRequestExtendSaro struct {
	// example:
	//
	// /
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendSaro) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendSaro) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendSaro) SetPath(v string) *ModifyDataSourceDeployRequestExtendSaro {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendSaro) SetVersion(v string) *ModifyDataSourceDeployRequestExtendSaro {
	s.Version = &v
	return s
}

type ModifyDataSourceDeployRequestProcessor struct {
	// example:
	//
	// {}
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// example:
	//
	// {}
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s ModifyDataSourceDeployRequestProcessor) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestProcessor) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestProcessor) SetArgs(v string) *ModifyDataSourceDeployRequestProcessor {
	s.Args = &v
	return s
}

func (s *ModifyDataSourceDeployRequestProcessor) SetResource(v string) *ModifyDataSourceDeployRequestProcessor {
	s.Resource = &v
	return s
}

type ModifyDataSourceDeployRequestStorage struct {
	// AK
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// example:
	//
	// test-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// dp-dev
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// /opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// /ude_jobs/iflow_offline_data_access
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// kubenest
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// item
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ModifyDataSourceDeployRequestStorage) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestStorage) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestStorage) SetAccessKey(v string) *ModifyDataSourceDeployRequestStorage {
	s.AccessKey = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetAccessSecret(v string) *ModifyDataSourceDeployRequestStorage {
	s.AccessSecret = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetBucket(v string) *ModifyDataSourceDeployRequestStorage {
	s.Bucket = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetEndpoint(v string) *ModifyDataSourceDeployRequestStorage {
	s.Endpoint = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetNamespace(v string) *ModifyDataSourceDeployRequestStorage {
	s.Namespace = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetOssPath(v string) *ModifyDataSourceDeployRequestStorage {
	s.OssPath = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetPartition(v string) *ModifyDataSourceDeployRequestStorage {
	s.Partition = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetPath(v string) *ModifyDataSourceDeployRequestStorage {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetProject(v string) *ModifyDataSourceDeployRequestStorage {
	s.Project = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetTable(v string) *ModifyDataSourceDeployRequestStorage {
	s.Table = &v
	return s
}

type ModifyDataSourceDeployRequestSwift struct {
	// topic
	//
	// example:
	//
	// ha-cn-0ju2rps6c08_api
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// zk
	//
	// example:
	//
	// zk
	Zk *string `json:"zk,omitempty" xml:"zk,omitempty"`
}

func (s ModifyDataSourceDeployRequestSwift) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployRequestSwift) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestSwift) SetTopic(v string) *ModifyDataSourceDeployRequestSwift {
	s.Topic = &v
	return s
}

func (s *ModifyDataSourceDeployRequestSwift) SetZk(v string) *ModifyDataSourceDeployRequestSwift {
	s.Zk = &v
	return s
}

type ModifyDataSourceDeployResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyDataSourceDeployResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployResponseBody) SetRequestId(v string) *ModifyDataSourceDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceDeployResponseBody) SetResult(v map[string]interface{}) *ModifyDataSourceDeployResponseBody {
	s.Result = v
	return s
}

type ModifyDataSourceDeployResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceDeployResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceDeployResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployResponse) SetHeaders(v map[string]*string) *ModifyDataSourceDeployResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceDeployResponse) SetStatusCode(v int32) *ModifyDataSourceDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceDeployResponse) SetBody(v *ModifyDataSourceDeployResponseBody) *ModifyDataSourceDeployResponse {
	s.Body = v
	return s
}

type ModifyFileRequest struct {
	// The content of the file.
	//
	// example:
	//
	// ""
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required when index building for full data in a MaxCompute data source is triggered.
	//
	// example:
	//
	// ds=20220713
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The name of the file in the full path
	//
	// This parameter is required.
	//
	// example:
	//
	// /schemas/generation_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileRequest) SetContent(v string) *ModifyFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyFileRequest) SetPartition(v int32) *ModifyFileRequest {
	s.Partition = &v
	return s
}

func (s *ModifyFileRequest) SetFileName(v string) *ModifyFileRequest {
	s.FileName = &v
	return s
}

type ModifyFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 89B968E6-1E41-58DF-BB25-5F98ECC759CE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileResponseBody) SetRequestId(v string) *ModifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFileResponseBody) SetResult(v map[string]interface{}) *ModifyFileResponseBody {
	s.Result = v
	return s
}

type ModifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileResponse) SetHeaders(v map[string]*string) *ModifyFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileResponse) SetStatusCode(v int32) *ModifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileResponse) SetBody(v *ModifyFileResponseBody) *ModifyFileResponse {
	s.Body = v
	return s
}

type ModifyIndexRequest struct {
	// example:
	//
	// 2
	BuildParallelNum *int32                            `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	Cluster          map[string]map[string]interface{} `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// ha-cn-35t3r02iq03@ha-cn-35t3r02iq03_test_api@hz_pre_vpc_domain_1@test_api@index_config_v1
	ClusterConfigName *string                 `json:"clusterConfigName,omitempty" xml:"clusterConfigName,omitempty"`
	Config            map[string]*ConfigValue `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// {\\"summarys\\":{\\"summary_fields\\":[\\"id\\"]},\\"indexs\\":[{\\"index_name\\":\\"index_id\\",\\"index_type\\":\\"PRIMARYKEY64\\",\\"index_fields\\":\\"id\\",\\"has_primary_key_attribute\\":true,\\"is_primary_key_sorted\\":false}],\\"attributes\\":[\\"id\\"],\\"fields\\":[{\\"field_name\\":\\"id\\",\\"field_type\\":\\"UINT16\\"}],\\"table_name\\":\\"index_2\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	DataSource     *string                           `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	DataSourceInfo *ModifyIndexRequestDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// 2
	MergeParallelNum *int32 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// PUSH_ONLY
	PushMode *string `json:"pushMode,omitempty" xml:"pushMode,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequest) SetBuildParallelNum(v int32) *ModifyIndexRequest {
	s.BuildParallelNum = &v
	return s
}

func (s *ModifyIndexRequest) SetCluster(v map[string]map[string]interface{}) *ModifyIndexRequest {
	s.Cluster = v
	return s
}

func (s *ModifyIndexRequest) SetClusterConfigName(v string) *ModifyIndexRequest {
	s.ClusterConfigName = &v
	return s
}

func (s *ModifyIndexRequest) SetConfig(v map[string]*ConfigValue) *ModifyIndexRequest {
	s.Config = v
	return s
}

func (s *ModifyIndexRequest) SetContent(v string) *ModifyIndexRequest {
	s.Content = &v
	return s
}

func (s *ModifyIndexRequest) SetDataSource(v string) *ModifyIndexRequest {
	s.DataSource = &v
	return s
}

func (s *ModifyIndexRequest) SetDataSourceInfo(v *ModifyIndexRequestDataSourceInfo) *ModifyIndexRequest {
	s.DataSourceInfo = v
	return s
}

func (s *ModifyIndexRequest) SetDescription(v string) *ModifyIndexRequest {
	s.Description = &v
	return s
}

func (s *ModifyIndexRequest) SetDomain(v string) *ModifyIndexRequest {
	s.Domain = &v
	return s
}

func (s *ModifyIndexRequest) SetMergeParallelNum(v int32) *ModifyIndexRequest {
	s.MergeParallelNum = &v
	return s
}

func (s *ModifyIndexRequest) SetPartition(v int32) *ModifyIndexRequest {
	s.Partition = &v
	return s
}

func (s *ModifyIndexRequest) SetPushMode(v string) *ModifyIndexRequest {
	s.PushMode = &v
	return s
}

func (s *ModifyIndexRequest) SetDryRun(v bool) *ModifyIndexRequest {
	s.DryRun = &v
	return s
}

type ModifyIndexRequestDataSourceInfo struct {
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// example:
	//
	// api
	BuildMode *string                                 `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	Config    *ModifyIndexRequestDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1709715164
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// 4
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	ProcessParallelNum *int32 `json:"processParallelNum,omitempty" xml:"processParallelNum,omitempty"`
	// example:
	//
	// 4
	ProcessPartitionCount *int32                                      `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	SaroConfig            *ModifyIndexRequestDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyIndexRequestDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfo) SetAutoBuildIndex(v bool) *ModifyIndexRequestDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetBuildMode(v string) *ModifyIndexRequestDataSourceInfo {
	s.BuildMode = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetConfig(v *ModifyIndexRequestDataSourceInfoConfig) *ModifyIndexRequestDataSourceInfo {
	s.Config = v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetDataTimeSec(v int32) *ModifyIndexRequestDataSourceInfo {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetDomain(v string) *ModifyIndexRequestDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetGeneration(v int64) *ModifyIndexRequestDataSourceInfo {
	s.Generation = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetName(v string) *ModifyIndexRequestDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetProcessParallelNum(v int32) *ModifyIndexRequestDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetProcessPartitionCount(v int32) *ModifyIndexRequestDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetSaroConfig(v *ModifyIndexRequestDataSourceInfoSaroConfig) *ModifyIndexRequestDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetType(v string) *ModifyIndexRequestDataSourceInfo {
	s.Type = &v
	return s
}

type ModifyIndexRequestDataSourceInfoConfig struct {
	// example:
	//
	// L***p
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// example:
	//
	// 5**9a6
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// example:
	//
	// test-bucket
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// test-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// /opensearch/oss.json
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20230114
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// test-hdfs-path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// bbt_algo_pai
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// item
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ModifyIndexRequestDataSourceInfoConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetAccessKey(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetAccessSecret(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetBucket(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetEndpoint(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetNamespace(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetOssPath(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetPartition(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetPath(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetProject(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetTable(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Table = &v
	return s
}

type ModifyIndexRequestDataSourceInfoSaroConfig struct {
	// example:
	//
	// flink-test-fjx-default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// device_event_shy_summary_
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s ModifyIndexRequestDataSourceInfoSaroConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) SetNamespace(v string) *ModifyIndexRequestDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) SetTableName(v string) *ModifyIndexRequestDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

type ModifyIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Object
	//
	// example:
	//
	// {}
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexResponseBody) SetRequestId(v string) *ModifyIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexResponseBody) SetResult(v interface{}) *ModifyIndexResponseBody {
	s.Result = v
	return s
}

type ModifyIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexResponse) SetHeaders(v map[string]*string) *ModifyIndexResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexResponse) SetStatusCode(v int32) *ModifyIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexResponse) SetBody(v *ModifyIndexResponseBody) *ModifyIndexResponse {
	s.Body = v
	return s
}

type ModifyIndexOnlineStrategyRequest struct {
	// example:
	//
	// 20
	ChangeRate *int32 `json:"changeRate,omitempty" xml:"changeRate,omitempty"`
}

func (s ModifyIndexOnlineStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexOnlineStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyRequest) SetChangeRate(v int32) *ModifyIndexOnlineStrategyRequest {
	s.ChangeRate = &v
	return s
}

type ModifyIndexOnlineStrategyResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexOnlineStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexOnlineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyResponseBody) SetRequestId(v string) *ModifyIndexOnlineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexOnlineStrategyResponseBody) SetResult(v map[string]interface{}) *ModifyIndexOnlineStrategyResponseBody {
	s.Result = v
	return s
}

type ModifyIndexOnlineStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexOnlineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexOnlineStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexOnlineStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyResponse) SetHeaders(v map[string]*string) *ModifyIndexOnlineStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexOnlineStrategyResponse) SetStatusCode(v int32) *ModifyIndexOnlineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexOnlineStrategyResponse) SetBody(v *ModifyIndexOnlineStrategyResponseBody) *ModifyIndexOnlineStrategyResponse {
	s.Body = v
	return s
}

type ModifyIndexPartitionRequest struct {
	// The name of the data source.
	//
	// example:
	//
	// test1
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The name of the data center.
	//
	// example:
	//
	// pre_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The primary key of generation.
	//
	// example:
	//
	// 1633293829
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The information about shards of the index.
	IndexInfos []*ModifyIndexPartitionRequestIndexInfos `json:"indexInfos,omitempty" xml:"indexInfos,omitempty" type:"Repeated"`
}

func (s ModifyIndexPartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequest) SetDataSourceName(v string) *ModifyIndexPartitionRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetDomainName(v string) *ModifyIndexPartitionRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetGeneration(v int64) *ModifyIndexPartitionRequest {
	s.Generation = &v
	return s
}

func (s *ModifyIndexPartitionRequest) SetIndexInfos(v []*ModifyIndexPartitionRequestIndexInfos) *ModifyIndexPartitionRequest {
	s.IndexInfos = v
	return s
}

type ModifyIndexPartitionRequestIndexInfos struct {
	// The name of the index.
	//
	// example:
	//
	// atest2
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The number of concurrency. The default value is 1.
	//
	// example:
	//
	// 1
	ParallelNum *int32 `json:"parallelNum,omitempty" xml:"parallelNum,omitempty"`
	// The number of shards of the index.
	//
	// example:
	//
	// 3
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
}

func (s ModifyIndexPartitionRequestIndexInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionRequestIndexInfos) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetIndexName(v string) *ModifyIndexPartitionRequestIndexInfos {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetParallelNum(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.ParallelNum = &v
	return s
}

func (s *ModifyIndexPartitionRequestIndexInfos) SetPartitionCount(v int32) *ModifyIndexPartitionRequestIndexInfos {
	s.PartitionCount = &v
	return s
}

type ModifyIndexPartitionResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexPartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponseBody) SetRequestId(v string) *ModifyIndexPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexPartitionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexPartitionResponseBody {
	s.Result = v
	return s
}

type ModifyIndexPartitionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexPartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexPartitionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponse) SetHeaders(v map[string]*string) *ModifyIndexPartitionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexPartitionResponse) SetStatusCode(v int32) *ModifyIndexPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexPartitionResponse) SetBody(v *ModifyIndexPartitionResponseBody) *ModifyIndexPartitionResponse {
	s.Body = v
	return s
}

type ModifyIndexVersionRequest struct {
	// The keyword used to search for a version. Fuzzy match is supported.
	//
	// example:
	//
	// {}
	Body []*ModifyIndexVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ModifyIndexVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequest) SetBody(v []*ModifyIndexVersionRequestBody) *ModifyIndexVersionRequest {
	s.Body = v
	return s
}

type ModifyIndexVersionRequestBody struct {
	// The ID of the index deployed in offline mode.
	//
	// example:
	//
	// 277
	BuildDeployId *string `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The version of the index.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModifyIndexVersionRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionRequestBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionRequestBody) SetBuildDeployId(v string) *ModifyIndexVersionRequestBody {
	s.BuildDeployId = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetIndexName(v string) *ModifyIndexVersionRequestBody {
	s.IndexName = &v
	return s
}

func (s *ModifyIndexVersionRequestBody) SetVersion(v string) *ModifyIndexVersionRequestBody {
	s.Version = &v
	return s
}

type ModifyIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponseBody) SetRequestId(v string) *ModifyIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexVersionResponseBody) SetResult(v map[string]interface{}) *ModifyIndexVersionResponseBody {
	s.Result = v
	return s
}

type ModifyIndexVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexVersionResponse) SetHeaders(v map[string]*string) *ModifyIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexVersionResponse) SetStatusCode(v int32) *ModifyIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexVersionResponse) SetBody(v *ModifyIndexVersionResponseBody) *ModifyIndexVersionResponse {
	s.Body = v
	return s
}

type ModifyNodeConfigRequest struct {
	// Specifies whether to enable the index.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The number of data replicas.
	//
	// example:
	//
	// 1
	DataDuplicateNumber *int32 `json:"dataDuplicateNumber,omitempty" xml:"dataDuplicateNumber,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 1
	DataFragmentNumber *int32 `json:"dataFragmentNumber,omitempty" xml:"dataFragmentNumber,omitempty"`
	// The traffic percentage.
	//
	// example:
	//
	// -100
	FlowRatio *int32 `json:"flowRatio,omitempty" xml:"flowRatio,omitempty"`
	// The minimum service ratio.
	//
	// example:
	//
	// 10
	MinServicePercent *int32 `json:"minServicePercent,omitempty" xml:"minServicePercent,omitempty"`
	// Specifies whether to mount the cluster.
	//
	// example:
	//
	// true
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// vpc_sh_domain_2
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The name of the data source. Valid values: search and not_search. search indicates to search data. not_search indicates not to search data.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_0704
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The original name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// ha-cn-zvp2iv9a401_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the algorithm. Valid values: pop, cp, hot, hint, and suggest.
	//
	// 	- pop indicates the popularity model.
	//
	// 	- cp indicates the category prediction model.
	//
	// 	- hot indicates the top search model.
	//
	// 	- hint indicates the hint model.
	//
	// 	- suggest indicates the drop-down suggestion model.
	//
	// This parameter is required.
	//
	// example:
	//
	// " "
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyNodeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigRequest) SetActive(v bool) *ModifyNodeConfigRequest {
	s.Active = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataDuplicateNumber(v int32) *ModifyNodeConfigRequest {
	s.DataDuplicateNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataFragmentNumber(v int32) *ModifyNodeConfigRequest {
	s.DataFragmentNumber = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetFlowRatio(v int32) *ModifyNodeConfigRequest {
	s.FlowRatio = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetMinServicePercent(v int32) *ModifyNodeConfigRequest {
	s.MinServicePercent = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetPublished(v bool) *ModifyNodeConfigRequest {
	s.Published = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetClusterName(v string) *ModifyNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetDataSourceName(v string) *ModifyNodeConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetName(v string) *ModifyNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyNodeConfigRequest) SetType(v string) *ModifyNodeConfigRequest {
	s.Type = &v
	return s
}

type ModifyNodeConfigResponseBody struct {
	// id of request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyNodeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponseBody) SetRequestId(v string) *ModifyNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeConfigResponseBody) SetResult(v map[string]interface{}) *ModifyNodeConfigResponseBody {
	s.Result = v
	return s
}

type ModifyNodeConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeConfigResponse) SetHeaders(v map[string]*string) *ModifyNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeConfigResponse) SetStatusCode(v int32) *ModifyNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeConfigResponse) SetBody(v *ModifyNodeConfigResponseBody) *ModifyNodeConfigResponse {
	s.Body = v
	return s
}

type ModifyOnlineConfigRequest struct {
	// ashortdescriptionofstruct
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOnlineConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigRequest) SetBody(v map[string]*string) *ModifyOnlineConfigRequest {
	s.Body = v
	return s
}

type ModifyOnlineConfigResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyOnlineConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponseBody) SetRequestId(v string) *ModifyOnlineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOnlineConfigResponseBody) SetResult(v map[string]interface{}) *ModifyOnlineConfigResponseBody {
	s.Result = v
	return s
}

type ModifyOnlineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOnlineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOnlineConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnlineConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigResponse) SetHeaders(v map[string]*string) *ModifyOnlineConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyOnlineConfigResponse) SetStatusCode(v int32) *ModifyOnlineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOnlineConfigResponse) SetBody(v *ModifyOnlineConfigResponseBody) *ModifyOnlineConfigResponse {
	s.Body = v
	return s
}

type ModifyPasswordRequest struct {
	// The password
	//
	// example:
	//
	// ******************************
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username
	//
	// example:
	//
	// "username"
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ModifyPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyPasswordRequest) SetPassword(v string) *ModifyPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyPasswordRequest) SetUsername(v string) *ModifyPasswordRequest {
	s.Username = &v
	return s
}

type ModifyPasswordResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 407BFD91-DE7D-50BA-8F88-CDE52A3B5E46
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponseBody) SetRequestId(v string) *ModifyPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPasswordResponseBody) SetResult(v map[string]interface{}) *ModifyPasswordResponseBody {
	s.Result = v
	return s
}

type ModifyPasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponse) SetHeaders(v map[string]*string) *ModifyPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyPasswordResponse) SetStatusCode(v int32) *ModifyPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPasswordResponse) SetBody(v *ModifyPasswordResponseBody) *ModifyPasswordResponse {
	s.Body = v
	return s
}

type ModifyPausePolicyRequest struct {
	Body map[string]*BodyValue `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPausePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPausePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyRequest) SetBody(v map[string]*BodyValue) *ModifyPausePolicyRequest {
	s.Body = v
	return s
}

type ModifyPausePolicyResponseBody struct {
	// id of request
	//
	// example:
	//
	// 0B1FF998-BB8D-5182-BFC0-E471AA77095A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPausePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPausePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyResponseBody) SetRequestId(v string) *ModifyPausePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPausePolicyResponseBody) SetResult(v map[string]interface{}) *ModifyPausePolicyResponseBody {
	s.Result = v
	return s
}

type ModifyPausePolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPausePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPausePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPausePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyPausePolicyResponse) SetHeaders(v map[string]*string) *ModifyPausePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyPausePolicyResponse) SetStatusCode(v int32) *ModifyPausePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPausePolicyResponse) SetBody(v *ModifyPausePolicyResponseBody) *ModifyPausePolicyResponse {
	s.Body = v
	return s
}

type ModifyPublicUrlIpListRequest struct {
	// body
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPublicUrlIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPublicUrlIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListRequest) SetBody(v map[string]*string) *ModifyPublicUrlIpListRequest {
	s.Body = v
	return s
}

type ModifyPublicUrlIpListResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPublicUrlIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPublicUrlIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListResponseBody) SetRequestId(v string) *ModifyPublicUrlIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPublicUrlIpListResponseBody) SetResult(v map[string]interface{}) *ModifyPublicUrlIpListResponseBody {
	s.Result = v
	return s
}

type ModifyPublicUrlIpListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPublicUrlIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPublicUrlIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPublicUrlIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListResponse) SetHeaders(v map[string]*string) *ModifyPublicUrlIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyPublicUrlIpListResponse) SetStatusCode(v int32) *ModifyPublicUrlIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPublicUrlIpListResponse) SetBody(v *ModifyPublicUrlIpListResponseBody) *ModifyPublicUrlIpListResponse {
	s.Body = v
	return s
}

type ModifyTableRequest struct {
	DataProcessConfig []*ModifyTableRequestDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	DataSource        *ModifyTableRequestDataSource          `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	FieldSchema       map[string]*string                     `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// example:
	//
	// 1
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// example:
	//
	// {}
	RawSchema   *string                          `json:"rawSchema,omitempty" xml:"rawSchema,omitempty"`
	VectorIndex []*ModifyTableRequestVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyTableRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequest) GoString() string {
	return s.String()
}

func (s *ModifyTableRequest) SetDataProcessConfig(v []*ModifyTableRequestDataProcessConfig) *ModifyTableRequest {
	s.DataProcessConfig = v
	return s
}

func (s *ModifyTableRequest) SetDataSource(v *ModifyTableRequestDataSource) *ModifyTableRequest {
	s.DataSource = v
	return s
}

func (s *ModifyTableRequest) SetFieldSchema(v map[string]*string) *ModifyTableRequest {
	s.FieldSchema = v
	return s
}

func (s *ModifyTableRequest) SetPartitionCount(v int32) *ModifyTableRequest {
	s.PartitionCount = &v
	return s
}

func (s *ModifyTableRequest) SetPrimaryKey(v string) *ModifyTableRequest {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyTableRequest) SetRawSchema(v string) *ModifyTableRequest {
	s.RawSchema = &v
	return s
}

func (s *ModifyTableRequest) SetVectorIndex(v []*ModifyTableRequestVectorIndex) *ModifyTableRequest {
	s.VectorIndex = v
	return s
}

func (s *ModifyTableRequest) SetDryRun(v bool) *ModifyTableRequest {
	s.DryRun = &v
	return s
}

type ModifyTableRequestDataProcessConfig struct {
	// example:
	//
	// source_image_vector
	DstField *string `json:"dstField,omitempty" xml:"dstField,omitempty"`
	// example:
	//
	// vectorize
	Operator *string                                    `json:"operator,omitempty" xml:"operator,omitempty"`
	Params   *ModifyTableRequestDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s ModifyTableRequestDataProcessConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfig) SetDstField(v string) *ModifyTableRequestDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetOperator(v string) *ModifyTableRequestDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetParams(v *ModifyTableRequestDataProcessConfigParams) *ModifyTableRequestDataProcessConfig {
	s.Params = v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetSrcField(v string) *ModifyTableRequestDataProcessConfig {
	s.SrcField = &v
	return s
}

type ModifyTableRequestDataProcessConfigParams struct {
	SrcFieldConfig *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
	// example:
	//
	// image
	VectorModal *string `json:"vectorModal,omitempty" xml:"vectorModal,omitempty"`
	// example:
	//
	// clip
	VectorModel *string `json:"vectorModel,omitempty" xml:"vectorModel,omitempty"`
}

func (s ModifyTableRequestDataProcessConfigParams) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfigParams) SetSrcFieldConfig(v *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) *ModifyTableRequestDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParams) SetVectorModal(v string) *ModifyTableRequestDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParams) SetVectorModel(v string) *ModifyTableRequestDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

type ModifyTableRequestDataProcessConfigParamsSrcFieldConfig struct {
	// example:
	//
	// test
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// example:
	//
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

type ModifyTableRequestDataSource struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                               `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *ModifyTableRequestDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
}

func (s ModifyTableRequestDataSource) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestDataSource) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataSource) SetAutoBuildIndex(v bool) *ModifyTableRequestDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyTableRequestDataSource) SetConfig(v *ModifyTableRequestDataSourceConfig) *ModifyTableRequestDataSource {
	s.Config = v
	return s
}

func (s *ModifyTableRequestDataSource) SetDataTimeSec(v int32) *ModifyTableRequestDataSource {
	s.DataTimeSec = &v
	return s
}

type ModifyTableRequestDataSourceConfig struct {
	// AK
	//
	// example:
	//
	// AK
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// AS
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// oss bucket
	//
	// example:
	//
	// antsys-shujiang-osstest
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20231220
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// yw_dw_rpt
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// behavior
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ModifyTableRequestDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestDataSourceConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataSourceConfig) SetAccessKey(v string) *ModifyTableRequestDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetAccessSecret(v string) *ModifyTableRequestDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetBucket(v string) *ModifyTableRequestDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetEndpoint(v string) *ModifyTableRequestDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetOssPath(v string) *ModifyTableRequestDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetPartition(v string) *ModifyTableRequestDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetProject(v string) *ModifyTableRequestDataSourceConfig {
	s.Project = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetTable(v string) *ModifyTableRequestDataSourceConfig {
	s.Table = &v
	return s
}

type ModifyTableRequestVectorIndex struct {
	AdvanceParams *ModifyTableRequestVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
	// example:
	//
	// 128
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// SquaredEuclidean
	DistanceType *string `json:"distanceType,omitempty" xml:"distanceType,omitempty"`
	// example:
	//
	// test_api
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// sparse_indices
	SparseIndexField *string `json:"sparseIndexField,omitempty" xml:"sparseIndexField,omitempty"`
	// example:
	//
	// sparse_values
	SparseValueField *string `json:"sparseValueField,omitempty" xml:"sparseValueField,omitempty"`
	// example:
	//
	// source_image_vector
	VectorField *string `json:"vectorField,omitempty" xml:"vectorField,omitempty"`
	// example:
	//
	// Qc
	VectorIndexType *string `json:"vectorIndexType,omitempty" xml:"vectorIndexType,omitempty"`
}

func (s ModifyTableRequestVectorIndex) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestVectorIndex) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestVectorIndex) SetAdvanceParams(v *ModifyTableRequestVectorIndexAdvanceParams) *ModifyTableRequestVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetDimension(v string) *ModifyTableRequestVectorIndex {
	s.Dimension = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetDistanceType(v string) *ModifyTableRequestVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetIndexName(v string) *ModifyTableRequestVectorIndex {
	s.IndexName = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetNamespace(v string) *ModifyTableRequestVectorIndex {
	s.Namespace = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetSparseIndexField(v string) *ModifyTableRequestVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetSparseValueField(v string) *ModifyTableRequestVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetVectorField(v string) *ModifyTableRequestVectorIndex {
	s.VectorField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetVectorIndexType(v string) *ModifyTableRequestVectorIndex {
	s.VectorIndexType = &v
	return s
}

type ModifyTableRequestVectorIndexAdvanceParams struct {
	// example:
	//
	// {}
	BuildIndexParams *string `json:"buildIndexParams,omitempty" xml:"buildIndexParams,omitempty"`
	// example:
	//
	// 5000
	LinearBuildThreshold *string `json:"linearBuildThreshold,omitempty" xml:"linearBuildThreshold,omitempty"`
	// example:
	//
	// 20000
	MinScanDocCnt *string `json:"minScanDocCnt,omitempty" xml:"minScanDocCnt,omitempty"`
	// example:
	//
	// {}
	SearchIndexParams *string `json:"searchIndexParams,omitempty" xml:"searchIndexParams,omitempty"`
}

func (s ModifyTableRequestVectorIndexAdvanceParams) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableRequestVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetBuildIndexParams(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetSearchIndexParams(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
	return s
}

type ModifyTableResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTableResponseBody) SetRequestId(v string) *ModifyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTableResponseBody) SetResult(v map[string]interface{}) *ModifyTableResponseBody {
	s.Result = v
	return s
}

type ModifyTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTableResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTableResponse) GoString() string {
	return s.String()
}

func (s *ModifyTableResponse) SetHeaders(v map[string]*string) *ModifyTableResponse {
	s.Headers = v
	return s
}

func (s *ModifyTableResponse) SetStatusCode(v int32) *ModifyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTableResponse) SetBody(v *ModifyTableResponseBody) *ModifyTableResponse {
	s.Body = v
	return s
}

type PublishAdvanceConfigRequest struct {
	Desc  *string                             `json:"desc,omitempty" xml:"desc,omitempty"`
	Files []*PublishAdvanceConfigRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s PublishAdvanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequest) SetDesc(v string) *PublishAdvanceConfigRequest {
	s.Desc = &v
	return s
}

func (s *PublishAdvanceConfigRequest) SetFiles(v []*PublishAdvanceConfigRequestFiles) *PublishAdvanceConfigRequest {
	s.Files = v
	return s
}

type PublishAdvanceConfigRequestFiles struct {
	Config         *PublishAdvanceConfigRequestFilesConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	DirName        *string                                 `json:"dirName,omitempty" xml:"dirName,omitempty"`
	FileName       *string                                 `json:"fileName,omitempty" xml:"fileName,omitempty"`
	OperateType    *string                                 `json:"operateType,omitempty" xml:"operateType,omitempty"`
	OssPath        *string                                 `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	ParentFullPath *string                                 `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s PublishAdvanceConfigRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigRequestFiles) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequestFiles) SetConfig(v *PublishAdvanceConfigRequestFilesConfig) *PublishAdvanceConfigRequestFiles {
	s.Config = v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetDirName(v string) *PublishAdvanceConfigRequestFiles {
	s.DirName = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetFileName(v string) *PublishAdvanceConfigRequestFiles {
	s.FileName = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetOperateType(v string) *PublishAdvanceConfigRequestFiles {
	s.OperateType = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetOssPath(v string) *PublishAdvanceConfigRequestFiles {
	s.OssPath = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetParentFullPath(v string) *PublishAdvanceConfigRequestFiles {
	s.ParentFullPath = &v
	return s
}

type PublishAdvanceConfigRequestFilesConfig struct {
	Content   *string                               `json:"content,omitempty" xml:"content,omitempty"`
	Variables map[string]*FilesConfigVariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s PublishAdvanceConfigRequestFilesConfig) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigRequestFilesConfig) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequestFilesConfig) SetContent(v string) *PublishAdvanceConfigRequestFilesConfig {
	s.Content = &v
	return s
}

func (s *PublishAdvanceConfigRequestFilesConfig) SetVariables(v map[string]*FilesConfigVariablesValue) *PublishAdvanceConfigRequestFilesConfig {
	s.Variables = v
	return s
}

type PublishAdvanceConfigResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishAdvanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponseBody) SetRequestId(v string) *PublishAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAdvanceConfigResponseBody) SetResult(v map[string]interface{}) *PublishAdvanceConfigResponseBody {
	s.Result = v
	return s
}

type PublishAdvanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAdvanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigResponse) SetHeaders(v map[string]*string) *PublishAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishAdvanceConfigResponse) SetStatusCode(v int32) *PublishAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAdvanceConfigResponse) SetBody(v *PublishAdvanceConfigResponseBody) *PublishAdvanceConfigResponse {
	s.Body = v
	return s
}

type PublishIndexVersionRequest struct {
	// The query result
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishIndexVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionRequest) SetBody(v map[string]interface{}) *PublishIndexVersionRequest {
	s.Body = v
	return s
}

type PublishIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishIndexVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponseBody) SetRequestId(v string) *PublishIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishIndexVersionResponseBody) SetResult(v map[string]interface{}) *PublishIndexVersionResponseBody {
	s.Result = v
	return s
}

type PublishIndexVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishIndexVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponse) SetHeaders(v map[string]*string) *PublishIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishIndexVersionResponse) SetStatusCode(v int32) *PublishIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishIndexVersionResponse) SetBody(v *PublishIndexVersionResponseBody) *PublishIndexVersionResponse {
	s.Body = v
	return s
}

type PushDocumentsRequest struct {
	Body []interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// id
	PkField *string `json:"pkField,omitempty" xml:"pkField,omitempty"`
}

func (s PushDocumentsRequest) String() string {
	return tea.Prettify(s)
}

func (s PushDocumentsRequest) GoString() string {
	return s.String()
}

func (s *PushDocumentsRequest) SetBody(v []interface{}) *PushDocumentsRequest {
	s.Body = v
	return s
}

func (s *PushDocumentsRequest) SetPkField(v string) *PushDocumentsRequest {
	s.PkField = &v
	return s
}

type PushDocumentsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushDocumentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *PushDocumentsResponseBody) SetRequestId(v string) *PushDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushDocumentsResponseBody) SetResult(v map[string]interface{}) *PushDocumentsResponseBody {
	s.Result = v
	return s
}

type PushDocumentsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDocumentsResponse) String() string {
	return tea.Prettify(s)
}

func (s PushDocumentsResponse) GoString() string {
	return s.String()
}

func (s *PushDocumentsResponse) SetHeaders(v map[string]*string) *PushDocumentsResponse {
	s.Headers = v
	return s
}

func (s *PushDocumentsResponse) SetStatusCode(v int32) *PushDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDocumentsResponse) SetBody(v *PushDocumentsResponseBody) *PushDocumentsResponse {
	s.Body = v
	return s
}

type RecoverIndexRequest struct {
	// The ID of the index deployed in offline mode.
	//
	// example:
	//
	// 277
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the data source
	//
	// example:
	//
	// ha-cn-pl32rf0js04_odps_first
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The primary key of generation.
	//
	// example:
	//
	// 1653018575
	Generation *string `json:"generation,omitempty" xml:"generation,omitempty"`
	// The name of the index
	//
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s RecoverIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexRequest) GoString() string {
	return s.String()
}

func (s *RecoverIndexRequest) SetBuildDeployId(v int32) *RecoverIndexRequest {
	s.BuildDeployId = &v
	return s
}

func (s *RecoverIndexRequest) SetDataSourceName(v string) *RecoverIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *RecoverIndexRequest) SetGeneration(v string) *RecoverIndexRequest {
	s.Generation = &v
	return s
}

func (s *RecoverIndexRequest) SetIndexName(v string) *RecoverIndexRequest {
	s.IndexName = &v
	return s
}

type RecoverIndexResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned by data search.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RecoverIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponseBody) SetRequestId(v string) *RecoverIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverIndexResponseBody) SetResult(v map[string]interface{}) *RecoverIndexResponseBody {
	s.Result = v
	return s
}

type RecoverIndexResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverIndexResponse) GoString() string {
	return s.String()
}

func (s *RecoverIndexResponse) SetHeaders(v map[string]*string) *RecoverIndexResponse {
	s.Headers = v
	return s
}

func (s *RecoverIndexResponse) SetStatusCode(v int32) *RecoverIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverIndexResponse) SetBody(v *RecoverIndexResponseBody) *RecoverIndexResponse {
	s.Body = v
	return s
}

type ReindexRequest struct {
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// example:
	//
	// oss://opensearch
	OssDataPath *string `json:"ossDataPath,omitempty" xml:"ossDataPath,omitempty"`
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s ReindexRequest) String() string {
	return tea.Prettify(s)
}

func (s ReindexRequest) GoString() string {
	return s.String()
}

func (s *ReindexRequest) SetDataTimeSec(v int32) *ReindexRequest {
	s.DataTimeSec = &v
	return s
}

func (s *ReindexRequest) SetOssDataPath(v string) *ReindexRequest {
	s.OssDataPath = &v
	return s
}

func (s *ReindexRequest) SetPartition(v string) *ReindexRequest {
	s.Partition = &v
	return s
}

type ReindexResponseBody struct {
	// requestId
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ReindexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReindexResponseBody) GoString() string {
	return s.String()
}

func (s *ReindexResponseBody) SetRequestId(v string) *ReindexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReindexResponseBody) SetResult(v map[string]interface{}) *ReindexResponseBody {
	s.Result = v
	return s
}

type ReindexResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReindexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReindexResponse) String() string {
	return tea.Prettify(s)
}

func (s ReindexResponse) GoString() string {
	return s.String()
}

func (s *ReindexResponse) SetHeaders(v map[string]*string) *ReindexResponse {
	s.Headers = v
	return s
}

func (s *ReindexResponse) SetStatusCode(v int32) *ReindexResponse {
	s.StatusCode = &v
	return s
}

func (s *ReindexResponse) SetBody(v *ReindexResponseBody) *ReindexResponse {
	s.Body = v
	return s
}

type RemoveClusterResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponseBody) SetRequestId(v string) *RemoveClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClusterResponseBody) SetResult(v map[string]interface{}) *RemoveClusterResponseBody {
	s.Result = v
	return s
}

type RemoveClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterResponse) SetHeaders(v map[string]*string) *RemoveClusterResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterResponse) SetStatusCode(v int32) *RemoveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterResponse) SetBody(v *RemoveClusterResponseBody) *RemoveClusterResponse {
	s.Body = v
	return s
}

type StartIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartIndexResponseBody) GoString() string {
	return s.String()
}

func (s *StartIndexResponseBody) SetRequestId(v string) *StartIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartIndexResponseBody) SetResult(v map[string]interface{}) *StartIndexResponseBody {
	s.Result = v
	return s
}

type StartIndexResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s StartIndexResponse) GoString() string {
	return s.String()
}

func (s *StartIndexResponse) SetHeaders(v map[string]*string) *StartIndexResponse {
	s.Headers = v
	return s
}

func (s *StartIndexResponse) SetStatusCode(v int32) *StartIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *StartIndexResponse) SetBody(v *StartIndexResponseBody) *StartIndexResponse {
	s.Body = v
	return s
}

type StopIndexResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Map
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopIndexResponseBody) GoString() string {
	return s.String()
}

func (s *StopIndexResponseBody) SetRequestId(v string) *StopIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopIndexResponseBody) SetResult(v map[string]interface{}) *StopIndexResponseBody {
	s.Result = v
	return s
}

type StopIndexResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s StopIndexResponse) GoString() string {
	return s.String()
}

func (s *StopIndexResponse) SetHeaders(v map[string]*string) *StopIndexResponse {
	s.Headers = v
	return s
}

func (s *StopIndexResponse) SetStatusCode(v int32) *StopIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *StopIndexResponse) SetBody(v *StopIndexResponseBody) *StopIndexResponse {
	s.Body = v
	return s
}

type StopTaskResponseBody struct {
	// id of request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) SetResult(v map[string]interface{}) *StopTaskResponseBody {
	s.Result = v
	return s
}

type StopTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTaskResponse) SetHeaders(v map[string]*string) *StopTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTaskResponse) SetStatusCode(v int32) *StopTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskResponse) SetBody(v *StopTaskResponseBody) *StopTaskResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// A list of instance-related specifications.
	Components []*UpdateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The description of the instance.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Valid values: UPGRADE and DOWNGRADE. UPGRADE indicates to upgrade the instance specifications. DOWNGRADE indicates to downgrade the instance specifications.
	//
	// example:
	//
	// ""
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest {
	s.Components = v
	return s
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetOrderType(v string) *UpdateInstanceRequest {
	s.OrderType = &v
	return s
}

type UpdateInstanceRequestComponents struct {
	// The specification code, which must be consistent with the values of the corresponding module parameters.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateInstanceRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestComponents) SetCode(v string) *UpdateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *UpdateInstanceRequestComponents) SetValue(v string) *UpdateInstanceRequestComponents {
	s.Value = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 90D6B8F5-FE97-4509-9AAB-367836C51818
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned
	Result *UpdateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody {
	s.Result = v
	return s
}

type UpdateInstanceResponseBodyResult struct {
	// The billing method
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The product code
	//
	// example:
	//
	// ha3-code
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance
	//
	// example:
	//
	// Test instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the instance expires
	//
	// example:
	//
	// 2019-01-06T16:00:00.0
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The ID of the resource
	//
	// example:
	//
	// ha-cn-0ju2s170b03
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The ID of the resource group
	//
	// example:
	//
	// rg-aeky6hthboewpuy
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the instance
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the instance was last updated
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UpdateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResult) SetChargeType(v string) *UpdateInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCommodityCode(v string) *UpdateInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCreateTime(v string) *UpdateInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDescription(v string) *UpdateInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetExpiredTime(v string) *UpdateInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInDebt(v bool) *UpdateInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInstanceId(v string) *UpdateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetLockMode(v string) *UpdateInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetResourceGroupId(v string) *UpdateInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetStatus(v string) *UpdateInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetUpdateTime(v string) *UpdateInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("searchengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Triggers reindexing.
//
// Description:
//
// ## Method
//
//     POST
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/actions/build-index
//
// @param request - BuildIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildIndexResponse
func (client *Client) BuildIndexWithOptions(instanceId *string, request *BuildIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BuildIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildMode)) {
		body["buildMode"] = request.BuildMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataTimeSec)) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BuildIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/actions/build-index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BuildIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers reindexing.
//
// Description:
//
// ## Method
//
//     POST
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/actions/build-index
//
// @param request - BuildIndexRequest
//
// @return BuildIndexResponse
func (client *Client) BuildIndex(instanceId *string, request *BuildIndexRequest) (_result *BuildIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildIndexResponse{}
	_body, _err := client.BuildIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/clusters`
//
// @param request - CreateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(instanceId *string, request *CreateClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoLoad)) {
		body["autoLoad"] = request.AutoLoad
	}

	if !tea.BoolValue(util.IsUnset(request.DataNode)) {
		body["dataNode"] = request.DataNode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.QueryNode)) {
		body["queryNode"] = request.QueryNode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/clusters`
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(instanceId *string, request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigDirResponse
func (client *Client) CreateConfigDirWithOptions(instanceId *string, configName *string, request *CreateConfigDirRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirName)) {
		body["dirName"] = request.DirName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFullPath)) {
		body["parentFullPath"] = request.ParentFullPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfigDir"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/dir"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateConfigDirRequest
//
// @return CreateConfigDirResponse
func (client *Client) CreateConfigDir(instanceId *string, configName *string, request *CreateConfigDirRequest) (_result *CreateConfigDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigDirResponse{}
	_body, _err := client.CreateConfigDirWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigFileResponse
func (client *Client) CreateConfigFileWithOptions(instanceId *string, configName *string, request *CreateConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		body["ossPath"] = request.OssPath
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFullPath)) {
		body["parentFullPath"] = request.ParentFullPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateConfigFileRequest
//
// @return CreateConfigFileResponse
func (client *Client) CreateConfigFile(instanceId *string, configName *string, request *CreateConfigFileRequest) (_result *CreateConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigFileResponse{}
	_body, _err := client.CreateConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates data sources.
//
// @param request - CreateDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(instanceId *string, request *CreateDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuildIndex)) {
		body["autoBuildIndex"] = request.AutoBuildIndex
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SaroConfig)) {
		body["saroConfig"] = request.SaroConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// Creates data sources.
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(instanceId *string, request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an index.
//
// Description:
//
// ### Method
//
// ```java
//
// POST
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/indexes
//
// ```
//
// @param request - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithOptions(instanceId *string, request *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildParallelNum)) {
		body["buildParallelNum"] = request.BuildParallelNum
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInfo)) {
		body["dataSourceInfo"] = request.DataSourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.MergeParallelNum)) {
		body["mergeParallelNum"] = request.MergeParallelNum
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an index.
//
// Description:
//
// ### Method
//
// ```java
//
// POST
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/indexes
//
// ```
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
func (client *Client) CreateIndex(instanceId *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Havenask instance.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/api/instances?dryRun=false`
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
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		body["components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances"),
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
// Creates a Havenask instance.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/api/instances?dryRun=false`
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
// 开通公网域名
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublicUrlResponse
func (client *Client) CreatePublicUrlWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePublicUrlResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublicUrl"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/public-url"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublicUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通公网域名
//
// @return CreatePublicUrlResponse
func (client *Client) CreatePublicUrl(instanceId *string) (_result *CreatePublicUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePublicUrlResponse{}
	_body, _err := client.CreatePublicUrlWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建索引V2
//
// @param request - CreateTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableResponse
func (client *Client) CreateTableWithOptions(instanceId *string, request *CreateTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataProcessConfig)) {
		body["dataProcessConfig"] = request.DataProcessConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DataProcessorCount)) {
		body["dataProcessorCount"] = request.DataProcessorCount
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.FieldSchema)) {
		body["fieldSchema"] = request.FieldSchema
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionCount)) {
		body["partitionCount"] = request.PartitionCount
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryKey)) {
		body["primaryKey"] = request.PrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.RawSchema)) {
		body["rawSchema"] = request.RawSchema
	}

	if !tea.BoolValue(util.IsUnset(request.VectorIndex)) {
		body["vectorIndex"] = request.VectorIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTable"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建索引V2
//
// @param request - CreateTableRequest
//
// @return CreateTableResponse
func (client *Client) CreateTable(instanceId *string, request *CreateTableRequest) (_result *CreateTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTableResponse{}
	_body, _err := client.CreateTableWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the details about advanced configurations.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvanceConfigResponse
func (client *Client) DeleteAdvanceConfigWithOptions(instanceId *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAdvanceConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the details about advanced configurations.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @return DeleteAdvanceConfigResponse
func (client *Client) DeleteAdvanceConfig(instanceId *string, configName *string) (_result *DeleteAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAdvanceConfigResponse{}
	_body, _err := client.DeleteAdvanceConfigWithOptions(instanceId, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigDirResponse
func (client *Client) DeleteConfigDirWithOptions(instanceId *string, configName *string, request *DeleteConfigDirRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirName)) {
		query["dirName"] = request.DirName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFullPath)) {
		query["parentFullPath"] = request.ParentFullPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfigDir"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/dir"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteConfigDirRequest
//
// @return DeleteConfigDirResponse
func (client *Client) DeleteConfigDir(instanceId *string, configName *string, request *DeleteConfigDirRequest) (_result *DeleteConfigDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigDirResponse{}
	_body, _err := client.DeleteConfigDirWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigFileResponse
func (client *Client) DeleteConfigFileWithOptions(instanceId *string, configName *string, request *DeleteConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFullPath)) {
		query["parentFullPath"] = request.ParentFullPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteConfigFileRequest
//
// @return DeleteConfigFileResponse
func (client *Client) DeleteConfigFile(instanceId *string, configName *string, request *DeleteConfigFileRequest) (_result *DeleteConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigFileResponse{}
	_body, _err := client.DeleteConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified data source.
//
// Description:
//
// ## Method
//
// `DELETE`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// Deletes a specified data source.
//
// Description:
//
// ## Method
//
// `DELETE`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(instanceId *string, dataSourceName *string) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an index.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}?dataSource=xxx
//
// @param request - DeleteIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndexWithOptions(instanceId *string, indexName *string, request *DeleteIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		query["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteDataSource)) {
		query["deleteDataSource"] = request.DeleteDataSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an index.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}?dataSource=xxx
//
// @param request - DeleteIndexRequest
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndex(instanceId *string, indexName *string, request *DeleteIndexRequest) (_result *DeleteIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(instanceId, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the version of an index.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexVersionResponse
func (client *Client) DeleteIndexVersionWithOptions(instanceId *string, indexName *string, versionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the version of an index.
//
// Description:
//
// ## Method
//
//     DELETE
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}
//
// @return DeleteIndexVersionResponse
func (client *Client) DeleteIndexVersion(instanceId *string, indexName *string, versionName *string) (_result *DeleteIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexVersionResponse{}
	_body, _err := client.DeleteIndexVersionWithOptions(instanceId, indexName, versionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified instance.
//
// Description:
//
// ### Method
//
// `DELETE`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified instance.
//
// Description:
//
// ### Method
//
// `DELETE`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(instanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除公网域名
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicUrlResponse
func (client *Client) DeletePublicUrlWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePublicUrlResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePublicUrl"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/public-url"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePublicUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除公网域名
//
// @return DeletePublicUrlResponse
func (client *Client) DeletePublicUrl(instanceId *string) (_result *DeletePublicUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePublicUrlResponse{}
	_body, _err := client.DeletePublicUrlWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除索引表V2
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableResponse
func (client *Client) DeleteTableWithOptions(instanceId *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTable"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除索引表V2
//
// @return DeleteTableResponse
func (client *Client) DeleteTable(instanceId *string, tableName *string) (_result *DeleteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableResponse{}
	_body, _err := client.DeleteTableWithOptions(instanceId, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/regions"),
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

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a forced switchover.
//
// Description:
//
// \\### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/force-switch/{fsmId}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceSwitchResponse
func (client *Client) ForceSwitchWithOptions(instanceId *string, fsmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ForceSwitchResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ForceSwitch"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/force-switch/" + tea.StringValue(openapiutil.GetEncodeParam(fsmId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ForceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a forced switchover.
//
// Description:
//
// \\### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/force-switch/{fsmId}
//
// ```
//
// @return ForceSwitchResponse
func (client *Client) ForceSwitch(instanceId *string, fsmId *string) (_result *ForceSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ForceSwitchResponse{}
	_body, _err := client.ForceSwitchWithOptions(instanceId, fsmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a dictionary.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @param request - GetAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvanceConfigResponse
func (client *Client) GetAdvanceConfigWithOptions(instanceId *string, configName *string, request *GetAdvanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAdvanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a dictionary.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @param request - GetAdvanceConfigRequest
//
// @return GetAdvanceConfigResponse
func (client *Client) GetAdvanceConfig(instanceId *string, configName *string, request *GetAdvanceConfigRequest) (_result *GetAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAdvanceConfigResponse{}
	_body, _err := client.GetAdvanceConfigWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information in a specified advanced configuration file.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - GetAdvanceConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvanceConfigFileResponse
func (client *Client) GetAdvanceConfigFileWithOptions(instanceId *string, configName *string, request *GetAdvanceConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAdvanceConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdvanceConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdvanceConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information in a specified advanced configuration file.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - GetAdvanceConfigFileRequest
//
// @return GetAdvanceConfigFileResponse
func (client *Client) GetAdvanceConfigFile(instanceId *string, configName *string, request *GetAdvanceConfigFileRequest) (_result *GetAdvanceConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAdvanceConfigFileResponse{}
	_body, _err := client.GetAdvanceConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cluster.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instance/{instanceId}/clusters/{clusterName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cluster.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instance/{instanceId}/clusters/{clusterName}`
//
// @return GetClusterResponse
func (client *Client) GetCluster(instanceId *string, clusterName *string) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the runtime information about a specified cluster.
//
// Description:
//
// ### Method
//
// GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-run-time-info
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterRunTimeInfoResponse
func (client *Client) GetClusterRunTimeInfoWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterRunTimeInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterRunTimeInfo"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-run-time-info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterRunTimeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the runtime information about a specified cluster.
//
// Description:
//
// ### Method
//
// GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-run-time-info
//
// @return GetClusterRunTimeInfoResponse
func (client *Client) GetClusterRunTimeInfo(instanceId *string) (_result *GetClusterRunTimeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterRunTimeInfoResponse{}
	_body, _err := client.GetClusterRunTimeInfoWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// Obtains a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @return GetDataSourceResponse
func (client *Client) GetDataSource(instanceId *string, dataSourceName *string) (_result *GetDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据源部署信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDeployResponse
func (client *Client) GetDataSourceDeployWithOptions(instanceId *string, deployName *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSourceDeployResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSourceDeploy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/deploys/" + tea.StringValue(openapiutil.GetEncodeParam(deployName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceDeployResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源部署信息
//
// @return GetDataSourceDeployResponse
func (client *Client) GetDataSourceDeploy(instanceId *string, deployName *string, dataSourceName *string) (_result *GetDataSourceDeployResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSourceDeployResponse{}
	_body, _err := client.GetDataSourceDeployWithOptions(instanceId, deployName, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays the overview of the deployment.
//
// Description:
//
// ## Method
//
// GET
//
// ## URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/deploy-graph
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployGraphResponse
func (client *Client) GetDeployGraphWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeployGraphResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployGraph"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/deploy-graph"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeployGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the overview of the deployment.
//
// Description:
//
// ## Method
//
// GET
//
// ## URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/deploy-graph
//
// ```
//
// @return GetDeployGraphResponse
func (client *Client) GetDeployGraph(instanceId *string) (_result *GetDeployGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeployGraphResponse{}
	_body, _err := client.GetDeployGraphWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about the version information of an index table.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
//
// @param request - GetFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithOptions(instanceId *string, indexName *string, versionName *string, request *GetFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName)) + "/file"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about the version information of an index table.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
//
// @param request - GetFileRequest
//
// @return GetFileResponse
func (client *Client) GetFile(instanceId *string, indexName *string, versionName *string, request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(instanceId, indexName, versionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an index version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexResponse
func (client *Client) GetIndexWithOptions(instanceId *string, indexName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an index version.
//
// @return GetIndexResponse
func (client *Client) GetIndex(instanceId *string, indexName *string) (_result *GetIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexResponse{}
	_body, _err := client.GetIndexWithOptions(instanceId, indexName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引在线生效策略
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexOnlineStrategyResponse
func (client *Client) GetIndexOnlineStrategyWithOptions(instanceId *string, dataSourceName *string, deployName *string, indexName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexOnlineStrategyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexOnlineStrategy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/deploys/" + tea.StringValue(openapiutil.GetEncodeParam(deployName)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/online-strategy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexOnlineStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取索引在线生效策略
//
// @return GetIndexOnlineStrategyResponse
func (client *Client) GetIndexOnlineStrategy(instanceId *string, dataSourceName *string, deployName *string, indexName *string) (_result *GetIndexOnlineStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexOnlineStrategyResponse{}
	_body, _err := client.GetIndexOnlineStrategyWithOptions(instanceId, dataSourceName, deployName, indexName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about index versions that the current index version can be rolled back to.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexVersionResponse
func (client *Client) GetIndexVersionWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/index-version"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about index versions that the current index version can be rolled back to.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @return GetIndexVersionResponse
func (client *Client) GetIndexVersion(instanceId *string, clusterName *string) (_result *GetIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexVersionResponse{}
	_body, _err := client.GetIndexVersionWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on a specified instance ID.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
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
// Queries the details of an instance based on a specified instance ID.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @return GetInstanceResponse
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

// Summary:
//
// Gets the configuration information of a node.
//
// @param request - GetNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeConfigResponse
func (client *Client) GetNodeConfigWithOptions(instanceId *string, request *GetNodeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["clusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the configuration information of a node.
//
// @param request - GetNodeConfigRequest
//
// @return GetNodeConfigResponse
func (client *Client) GetNodeConfig(instanceId *string, request *GetNodeConfigRequest) (_result *GetNodeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeConfigResponse{}
	_body, _err := client.GetNodeConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引表信息V2
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithOptions(instanceId *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTable"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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

// Summary:
//
// 获取索引表信息V2
//
// @return GetTableResponse
func (client *Client) GetTable(instanceId *string, tableName *string) (_result *GetTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(instanceId, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据generationId获取某个索引版本状态V2
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableGenerationResponse
func (client *Client) GetTableGenerationWithOptions(instanceId *string, tableName *string, generationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableGenerationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableGeneration"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/index_versions/" + tea.StringValue(openapiutil.GetEncodeParam(generationId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据generationId获取某个索引版本状态V2
//
// @return GetTableGenerationResponse
func (client *Client) GetTableGeneration(instanceId *string, tableName *string, generationId *string) (_result *GetTableGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableGenerationResponse{}
	_body, _err := client.GetTableGenerationWithOptions(instanceId, tableName, generationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the file list in an advanced configuration directory.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/dir?dirName={dirName}`
//
// @param request - ListAdvanceConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdvanceConfigDirResponse
func (client *Client) ListAdvanceConfigDirWithOptions(instanceId *string, configName *string, request *ListAdvanceConfigDirRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAdvanceConfigDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirName)) {
		query["dirName"] = request.DirName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdvanceConfigDir"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/dir"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdvanceConfigDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the file list in an advanced configuration directory.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/dir?dirName={dirName}`
//
// @param request - ListAdvanceConfigDirRequest
//
// @return ListAdvanceConfigDirResponse
func (client *Client) ListAdvanceConfigDir(instanceId *string, configName *string, request *ListAdvanceConfigDirRequest) (_result *ListAdvanceConfigDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAdvanceConfigDirResponse{}
	_body, _err := client.ListAdvanceConfigDirWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of advanced configurations.
//
// Description:
//
// ## Sample requests
//
// `GET /openapi/ha3/instances/ose-test1/advanced-configs`
//
// @param request - ListAdvanceConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdvanceConfigsResponse
func (client *Client) ListAdvanceConfigsWithOptions(instanceId *string, request *ListAdvanceConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAdvanceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		query["indexName"] = request.IndexName
	}

	if !tea.BoolValue(util.IsUnset(request.NewMode)) {
		query["newMode"] = request.NewMode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdvanceConfigs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdvanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of advanced configurations.
//
// Description:
//
// ## Sample requests
//
// `GET /openapi/ha3/instances/ose-test1/advanced-configs`
//
// @param request - ListAdvanceConfigsRequest
//
// @return ListAdvanceConfigsResponse
func (client *Client) ListAdvanceConfigs(instanceId *string, request *ListAdvanceConfigsRequest) (_result *ListAdvanceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAdvanceConfigsResponse{}
	_body, _err := client.ListAdvanceConfigsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cluster names.
//
// Description:
//
// ### Method
//
// GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-names
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterNamesResponse
func (client *Client) ListClusterNamesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterNamesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterNames"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/cluster-names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster names.
//
// Description:
//
// ### Method
//
// GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-names
//
// @return ListClusterNamesResponse
func (client *Client) ListClusterNames() (_result *ListClusterNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterNamesResponse{}
	_body, _err := client.ListClusterNamesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays cluster tasks .
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/cluster-tasks
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterTasksResponse
func (client *Client) ListClusterTasksWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterTasks"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays cluster tasks .
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/cluster-tasks
//
// ```
//
// @return ListClusterTasksResponse
func (client *Client) ListClusterTasks(instanceId *string) (_result *ListClusterTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterTasksResponse{}
	_body, _err := client.ListClusterTasksWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of clusters.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of clusters.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters
//
// ```
//
// @return ListClustersResponse
func (client *Client) ListClusters(instanceId *string) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the schema information of a specified data source.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/schemas`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSchemasResponse
func (client *Client) ListDataSourceSchemasWithOptions(instanceId *string, dataSourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceSchemasResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceSchemas"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the schema information of a specified data source.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/schemas`
//
// @return ListDataSourceSchemasResponse
func (client *Client) ListDataSourceSchemas(instanceId *string, dataSourceName *string) (_result *ListDataSourceSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceSchemasResponse{}
	_body, _err := client.ListDataSourceSchemasWithOptions(instanceId, dataSourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays data source tasks.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/data-source-tasks
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTasksResponse
func (client *Client) ListDataSourceTasksWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTasksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTasks"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-source-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays data source tasks.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/data-source-tasks
//
// ```
//
// @return ListDataSourceTasksResponse
func (client *Client) ListDataSourceTasks(instanceId *string) (_result *ListDataSourceTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTasksResponse{}
	_body, _err := client.ListDataSourceTasksWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of data sources.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
// Obtains the list of data sources.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources`
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(instanceId *string) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the data restoration version of a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/generations?domainName={domainName}`
//
// @param request - ListDateSourceGenerationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDateSourceGenerationsResponse
func (client *Client) ListDateSourceGenerationsWithOptions(instanceId *string, dataSourceName *string, request *ListDateSourceGenerationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDateSourceGenerationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ValidStatus)) {
		query["validStatus"] = request.ValidStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDateSourceGenerations"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/generations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDateSourceGenerationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the data restoration version of a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/generations?domainName={domainName}`
//
// @param request - ListDateSourceGenerationsRequest
//
// @return ListDateSourceGenerationsResponse
func (client *Client) ListDateSourceGenerations(instanceId *string, dataSourceName *string, request *ListDateSourceGenerationsRequest) (_result *ListDateSourceGenerationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDateSourceGenerationsResponse{}
	_body, _err := client.ListDateSourceGenerationsWithOptions(instanceId, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexRecoverRecordsResponse
func (client *Client) ListIndexRecoverRecordsWithOptions(indexName *string, instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndexRecoverRecordsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndexRecoverRecords"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/actions/list-recover-records"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndexRecoverRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return ListIndexRecoverRecordsResponse
func (client *Client) ListIndexRecoverRecords(indexName *string, instanceId *string) (_result *ListIndexRecoverRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexRecoverRecordsResponse{}
	_body, _err := client.ListIndexRecoverRecordsWithOptions(indexName, instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the index list.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes
//
// @param request - ListIndexesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexesResponse
func (client *Client) ListIndexesWithOptions(instanceId *string, request *ListIndexesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndexesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewMode)) {
		query["newMode"] = request.NewMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndexes"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndexesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the index list.
//
// Description:
//
// ## Method
//
//     GET
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes
//
// @param request - ListIndexesRequest
//
// @return ListIndexesResponse
func (client *Client) ListIndexes(instanceId *string, request *ListIndexesRequest) (_result *ListIndexesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexesResponse{}
	_body, _err := client.ListIndexesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instances.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/specs?type=qrs`
//
// @param request - ListInstanceSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceSpecsResponse
func (client *Client) ListInstanceSpecsWithOptions(instanceId *string, request *ListInstanceSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceSpecs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/specs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instances.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/specs?type=qrs`
//
// @param request - ListInstanceSpecsRequest
//
// @return ListInstanceSpecsResponse
func (client *Client) ListInstanceSpecs(instanceId *string, request *ListInstanceSpecsRequest) (_result *ListInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSpecsResponse{}
	_body, _err := client.ListInstanceSpecsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instances.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/[code]/instances`
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances"),
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
// Queries instances.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/[code]/instances`
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
// Obtains the details of online configurations.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs?domain={domain}
//
// ```
//
// @param request - ListOnlineConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineConfigsResponse
func (client *Client) ListOnlineConfigsWithOptions(instanceId *string, nodeName *string, request *ListOnlineConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOnlineConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnlineConfigs"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node/" + tea.StringValue(openapiutil.GetEncodeParam(nodeName)) + "/online-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnlineConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of online configurations.
//
// Description:
//
// ### Method
//
// ```java
//
// GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs?domain={domain}
//
// ```
//
// @param request - ListOnlineConfigsRequest
//
// @return ListOnlineConfigsResponse
func (client *Client) ListOnlineConfigs(instanceId *string, nodeName *string, request *ListOnlineConfigsRequest) (_result *ListOnlineConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOnlineConfigsResponse{}
	_body, _err := client.ListOnlineConfigsWithOptions(instanceId, nodeName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPausePolicysResponse
func (client *Client) ListPausePolicysWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPausePolicysResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPausePolicys"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/pause-policies"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPausePolicysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return ListPausePolicysResponse
func (client *Client) ListPausePolicys(instanceId *string) (_result *ListPausePolicysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPausePolicysResponse{}
	_body, _err := client.ListPausePolicysWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPostQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPostQueryResultResponse
func (client *Client) ListPostQueryResultWithOptions(instanceId *string, request *ListPostQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPostQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPostQueryResult"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPostQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPostQueryResultRequest
//
// @return ListPostQueryResultResponse
func (client *Client) ListPostQueryResult(instanceId *string, request *ListPostQueryResultRequest) (_result *ListPostQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPostQueryResultResponse{}
	_body, _err := client.ListPostQueryResultWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the query result.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/query?query=xxxx`
//
// @param request - ListQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryResultResponse
func (client *Client) ListQueryResultWithOptions(instanceId *string, request *ListQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		query["sql"] = request.Sql
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryResult"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the query result.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/query?query=xxxx`
//
// @param request - ListQueryResultRequest
//
// @return ListQueryResultResponse
func (client *Client) ListQueryResult(instanceId *string, request *ListQueryResultRequest) (_result *ListQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryResultResponse{}
	_body, _err := client.ListQueryResultWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRestQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRestQueryResultResponse
func (client *Client) ListRestQueryResultWithOptions(instanceId *string, request *ListRestQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRestQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		body["indexName"] = request.IndexName
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRestQueryResult"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/rest-query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRestQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRestQueryResultRequest
//
// @return ListRestQueryResultResponse
func (client *Client) ListRestQueryResult(instanceId *string, request *ListRestQueryResultRequest) (_result *ListRestQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRestQueryResultResponse{}
	_body, _err := client.ListRestQueryResultWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过数据源配置获取schema信息
//
// @param request - ListSchemasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemasResponse
func (client *Client) ListSchemasWithOptions(instanceId *string, request *ListSchemasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccessSecret)) {
		query["accessSecret"] = request.AccessSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		query["partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		query["table"] = request.Table
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSchemas"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过数据源配置获取schema信息
//
// @param request - ListSchemasRequest
//
// @return ListSchemasResponse
func (client *Client) ListSchemas(instanceId *string, request *ListSchemasRequest) (_result *ListSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSchemasResponse{}
	_body, _err := client.ListSchemasWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引generation列表V2
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableGenerationsResponse
func (client *Client) ListTableGenerationsWithOptions(instanceId *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTableGenerationsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTableGenerations"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/index_versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTableGenerationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取索引generation列表V2
//
// @return ListTableGenerationsResponse
func (client *Client) ListTableGenerations(instanceId *string, tableName *string) (_result *ListTableGenerationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableGenerationsResponse{}
	_body, _err := client.ListTableGenerationsWithOptions(instanceId, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取索引列表V2
//
// @param request - ListTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(instanceId *string, request *ListTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewMode)) {
		query["newMode"] = request.NewMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTables"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取索引列表V2
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(instanceId *string, request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群任务列表（数据源+集群）
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(instanceId *string, request *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tasks"),
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
// 获取集群任务列表（数据源+集群）
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(instanceId *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListVectorQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVectorQueryResultResponse
func (client *Client) ListVectorQueryResultWithOptions(instanceId *string, request *ListVectorQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVectorQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.VectorQueryType)) {
		query["vectorQueryType"] = request.VectorQueryType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVectorQueryResult"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/vector-query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVectorQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListVectorQueryResultRequest
//
// @return ListVectorQueryResultResponse
func (client *Client) ListVectorQueryResult(instanceId *string, request *ListVectorQueryResultRequest) (_result *ListVectorQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVectorQueryResultResponse{}
	_body, _err := client.ListVectorQueryResultWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAdvanceConfigResponse
func (client *Client) ModifyAdvanceConfigWithOptions(instanceId *string, configName *string, request *ModifyAdvanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAdvanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["contentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTime)) {
		body["updateTime"] = request.UpdateTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAdvanceConfigRequest
//
// @return ModifyAdvanceConfigResponse
func (client *Client) ModifyAdvanceConfig(instanceId *string, configName *string, request *ModifyAdvanceConfigRequest) (_result *ModifyAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAdvanceConfigResponse{}
	_body, _err := client.ModifyAdvanceConfigWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the advanced configurations.
//
// Description:
//
// ## Method
//
//     put
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - ModifyAdvanceConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAdvanceConfigFileResponse
func (client *Client) ModifyAdvanceConfigFileWithOptions(instanceId *string, configName *string, request *ModifyAdvanceConfigFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAdvanceConfigFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAdvanceConfigFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAdvanceConfigFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the advanced configurations.
//
// Description:
//
// ## Method
//
//     put
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - ModifyAdvanceConfigFileRequest
//
// @return ModifyAdvanceConfigFileResponse
func (client *Client) ModifyAdvanceConfigFile(instanceId *string, configName *string, request *ModifyAdvanceConfigFileRequest) (_result *ModifyAdvanceConfigFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAdvanceConfigFileResponse{}
	_body, _err := client.ModifyAdvanceConfigFileWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a specified cluster.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/desc`
//
// @param request - ModifyClusterDescRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterDescResponse
func (client *Client) ModifyClusterDescWithOptions(instanceId *string, clusterName *string, request *ModifyClusterDescRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterDescResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterDesc"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/desc"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterDescResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a specified cluster.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/desc`
//
// @param request - ModifyClusterDescRequest
//
// @return ModifyClusterDescResponse
func (client *Client) ModifyClusterDesc(instanceId *string, clusterName *string, request *ModifyClusterDescRequest) (_result *ModifyClusterDescResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterDescResponse{}
	_body, _err := client.ModifyClusterDescWithOptions(instanceId, clusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration information of a cluster.
//
// Description:
//
// ## Request syntax
//
//     PUT /openapi/ha3/instances/{instanceId}/cluster-offline-config
//
// @param request - ModifyClusterOfflineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterOfflineConfigResponse
func (client *Client) ModifyClusterOfflineConfigWithOptions(instanceId *string, request *ModifyClusterOfflineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterOfflineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildMode)) {
		body["buildMode"] = request.BuildMode
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataTimeSec)) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.PushMode)) {
		body["pushMode"] = request.PushMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterOfflineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-offline-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterOfflineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration information of a cluster.
//
// Description:
//
// ## Request syntax
//
//     PUT /openapi/ha3/instances/{instanceId}/cluster-offline-config
//
// @param request - ModifyClusterOfflineConfigRequest
//
// @return ModifyClusterOfflineConfigResponse
func (client *Client) ModifyClusterOfflineConfig(instanceId *string, request *ModifyClusterOfflineConfigRequest) (_result *ModifyClusterOfflineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterOfflineConfigResponse{}
	_body, _err := client.ModifyClusterOfflineConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the online configurations of a cluster.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/cluster-online-config`
//
// @param request - ModifyClusterOnlineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterOnlineConfigResponse
func (client *Client) ModifyClusterOnlineConfigWithOptions(instanceId *string, request *ModifyClusterOnlineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterOnlineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clusters)) {
		body["clusters"] = request.Clusters
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterOnlineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/cluster-online-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterOnlineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the online configurations of a cluster.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/cluster-online-config`
//
// @param request - ModifyClusterOnlineConfigRequest
//
// @return ModifyClusterOnlineConfigResponse
func (client *Client) ModifyClusterOnlineConfig(instanceId *string, request *ModifyClusterOnlineConfigRequest) (_result *ModifyClusterOnlineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterOnlineConfigResponse{}
	_body, _err := client.ModifyClusterOnlineConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a data source.
//
// Description:
//
// ## Method
//
// `PUT`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param request - ModifyDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSourceWithOptions(instanceId *string, dataSourceName *string, request *ModifyDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataSource"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a data source.
//
// Description:
//
// ## Method
//
// `PUT`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param request - ModifyDataSourceRequest
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSource(instanceId *string, dataSourceName *string, request *ModifyDataSourceRequest) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(instanceId, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改数据源部署信息
//
// @param request - ModifyDataSourceDeployRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceDeployResponse
func (client *Client) ModifyDataSourceDeployWithOptions(instanceId *string, deployName *string, dataSourceName *string, request *ModifyDataSourceDeployRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDataSourceDeployResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.GenerationId)) {
		query["generationId"] = request.GenerationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuildIndex)) {
		body["autoBuildIndex"] = request.AutoBuildIndex
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Processor)) {
		body["processor"] = request.Processor
	}

	if !tea.BoolValue(util.IsUnset(request.Storage)) {
		body["storage"] = request.Storage
	}

	if !tea.BoolValue(util.IsUnset(request.Swift)) {
		body["swift"] = request.Swift
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataSourceDeploy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/deploys/" + tea.StringValue(openapiutil.GetEncodeParam(deployName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataSourceDeployResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改数据源部署信息
//
// @param request - ModifyDataSourceDeployRequest
//
// @return ModifyDataSourceDeployResponse
func (client *Client) ModifyDataSourceDeploy(instanceId *string, deployName *string, dataSourceName *string, request *ModifyDataSourceDeployRequest) (_result *ModifyDataSourceDeployResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDataSourceDeployResponse{}
	_body, _err := client.ModifyDataSourceDeployWithOptions(instanceId, deployName, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a file.
//
// Description:
//
// ## Method
//
//     PUT
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
//
// @param request - ModifyFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileResponse
func (client *Client) ModifyFileWithOptions(instanceId *string, indexName *string, versionName *string, request *ModifyFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFile"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionName)) + "/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a file.
//
// Description:
//
// ## Method
//
//     PUT
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
//
// @param request - ModifyFileRequest
//
// @return ModifyFileResponse
func (client *Client) ModifyFile(instanceId *string, indexName *string, versionName *string, request *ModifyFileRequest) (_result *ModifyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyFileResponse{}
	_body, _err := client.ModifyFileWithOptions(instanceId, indexName, versionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexResponse
func (client *Client) ModifyIndexWithOptions(instanceId *string, indexName *string, request *ModifyIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildParallelNum)) {
		body["buildParallelNum"] = request.BuildParallelNum
	}

	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		body["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterConfigName)) {
		body["clusterConfigName"] = request.ClusterConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInfo)) {
		body["dataSourceInfo"] = request.DataSourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.MergeParallelNum)) {
		body["mergeParallelNum"] = request.MergeParallelNum
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.PushMode)) {
		body["pushMode"] = request.PushMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyIndexRequest
//
// @return ModifyIndexResponse
func (client *Client) ModifyIndex(instanceId *string, indexName *string, request *ModifyIndexRequest) (_result *ModifyIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexResponse{}
	_body, _err := client.ModifyIndexWithOptions(instanceId, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改索引在线策略
//
// @param request - ModifyIndexOnlineStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexOnlineStrategyResponse
func (client *Client) ModifyIndexOnlineStrategyWithOptions(instanceId *string, dataSourceName *string, deployName *string, indexName *string, request *ModifyIndexOnlineStrategyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexOnlineStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeRate)) {
		body["changeRate"] = request.ChangeRate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndexOnlineStrategy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/deploys/" + tea.StringValue(openapiutil.GetEncodeParam(deployName)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/online-strategy"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexOnlineStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改索引在线策略
//
// @param request - ModifyIndexOnlineStrategyRequest
//
// @return ModifyIndexOnlineStrategyResponse
func (client *Client) ModifyIndexOnlineStrategy(instanceId *string, dataSourceName *string, deployName *string, indexName *string, request *ModifyIndexOnlineStrategyRequest) (_result *ModifyIndexOnlineStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexOnlineStrategyResponse{}
	_body, _err := client.ModifyIndexOnlineStrategyWithOptions(instanceId, dataSourceName, deployName, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about index partitions.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/index-partition`
//
// @param request - ModifyIndexPartitionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexPartitionResponse
func (client *Client) ModifyIndexPartitionWithOptions(instanceId *string, request *ModifyIndexPartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexPartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.IndexInfos)) {
		body["indexInfos"] = request.IndexInfos
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndexPartition"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/index-partition"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexPartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about index partitions.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/index-partition`
//
// @param request - ModifyIndexPartitionRequest
//
// @return ModifyIndexPartitionResponse
func (client *Client) ModifyIndexPartition(instanceId *string, request *ModifyIndexPartitionRequest) (_result *ModifyIndexPartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexPartitionResponse{}
	_body, _err := client.ModifyIndexPartitionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the index version of a cluster (an index version rollback).
//
// Description:
//
// ## Method
//
//     PUT
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @param request - ModifyIndexVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexVersionResponse
func (client *Client) ModifyIndexVersionWithOptions(instanceId *string, clusterName *string, request *ModifyIndexVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyIndexVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName)) + "/index-version"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the index version of a cluster (an index version rollback).
//
// Description:
//
// ## Method
//
//     PUT
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @param request - ModifyIndexVersionRequest
//
// @return ModifyIndexVersionResponse
func (client *Client) ModifyIndexVersion(instanceId *string, clusterName *string, request *ModifyIndexVersionRequest) (_result *ModifyIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyIndexVersionResponse{}
	_body, _err := client.ModifyIndexVersionWithOptions(instanceId, clusterName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a node.
//
// Description:
//
// ### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node-config?type=qrs&name=test
//
// ```
//
// @param request - ModifyNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeConfigResponse
func (client *Client) ModifyNodeConfigWithOptions(instanceId *string, request *ModifyNodeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyNodeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["clusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.DataDuplicateNumber)) {
		body["dataDuplicateNumber"] = request.DataDuplicateNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DataFragmentNumber)) {
		body["dataFragmentNumber"] = request.DataFragmentNumber
	}

	if !tea.BoolValue(util.IsUnset(request.FlowRatio)) {
		body["flowRatio"] = request.FlowRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinServicePercent)) {
		body["minServicePercent"] = request.MinServicePercent
	}

	if !tea.BoolValue(util.IsUnset(request.Published)) {
		body["published"] = request.Published
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodeConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a node.
//
// Description:
//
// ### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node-config?type=qrs&name=test
//
// ```
//
// @param request - ModifyNodeConfigRequest
//
// @return ModifyNodeConfigResponse
func (client *Client) ModifyNodeConfig(instanceId *string, request *ModifyNodeConfigRequest) (_result *ModifyNodeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeConfigResponse{}
	_body, _err := client.ModifyNodeConfigWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies online configurations.
//
// Description:
//
// ### Method
//
// ```java
//
// put
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs/{indexName}
//
// ```
//
// @param request - ModifyOnlineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOnlineConfigResponse
func (client *Client) ModifyOnlineConfigWithOptions(instanceId *string, nodeName *string, indexName *string, request *ModifyOnlineConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyOnlineConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOnlineConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/node/" + tea.StringValue(openapiutil.GetEncodeParam(nodeName)) + "/online-configs/" + tea.StringValue(openapiutil.GetEncodeParam(indexName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOnlineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies online configurations.
//
// Description:
//
// ### Method
//
// ```java
//
// put
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs/{indexName}
//
// ```
//
// @param request - ModifyOnlineConfigRequest
//
// @return ModifyOnlineConfigResponse
func (client *Client) ModifyOnlineConfig(instanceId *string, nodeName *string, indexName *string, request *ModifyOnlineConfigRequest) (_result *ModifyOnlineConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyOnlineConfigResponse{}
	_body, _err := client.ModifyOnlineConfigWithOptions(instanceId, nodeName, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例的密码
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/password`
//
// @param request - ModifyPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPasswordResponse
func (client *Client) ModifyPasswordWithOptions(instanceId *string, request *ModifyPasswordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPassword"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/password"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例的密码
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/password`
//
// @param request - ModifyPasswordRequest
//
// @return ModifyPasswordResponse
func (client *Client) ModifyPassword(instanceId *string, request *ModifyPasswordRequest) (_result *ModifyPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyPasswordResponse{}
	_body, _err := client.ModifyPasswordWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyPausePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPausePolicyResponse
func (client *Client) ModifyPausePolicyWithOptions(instanceId *string, request *ModifyPausePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyPausePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPausePolicy"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/pause-policies"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPausePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyPausePolicyRequest
//
// @return ModifyPausePolicyResponse
func (client *Client) ModifyPausePolicy(instanceId *string, request *ModifyPausePolicyRequest) (_result *ModifyPausePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyPausePolicyResponse{}
	_body, _err := client.ModifyPausePolicyWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改公网域名访问白名单
//
// @param request - ModifyPublicUrlIpListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPublicUrlIpListResponse
func (client *Client) ModifyPublicUrlIpListWithOptions(instanceId *string, request *ModifyPublicUrlIpListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyPublicUrlIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPublicUrlIpList"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/public-url-ip-list"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPublicUrlIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改公网域名访问白名单
//
// @param request - ModifyPublicUrlIpListRequest
//
// @return ModifyPublicUrlIpListResponse
func (client *Client) ModifyPublicUrlIpList(instanceId *string, request *ModifyPublicUrlIpListRequest) (_result *ModifyPublicUrlIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyPublicUrlIpListResponse{}
	_body, _err := client.ModifyPublicUrlIpListWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改索引V2
//
// @param request - ModifyTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTableResponse
func (client *Client) ModifyTableWithOptions(instanceId *string, tableName *string, request *ModifyTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataProcessConfig)) {
		body["dataProcessConfig"] = request.DataProcessConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["dataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.FieldSchema)) {
		body["fieldSchema"] = request.FieldSchema
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionCount)) {
		body["partitionCount"] = request.PartitionCount
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryKey)) {
		body["primaryKey"] = request.PrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.RawSchema)) {
		body["rawSchema"] = request.RawSchema
	}

	if !tea.BoolValue(util.IsUnset(request.VectorIndex)) {
		body["vectorIndex"] = request.VectorIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTable"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改索引V2
//
// @param request - ModifyTableRequest
//
// @return ModifyTableResponse
func (client *Client) ModifyTable(instanceId *string, tableName *string, request *ModifyTableRequest) (_result *ModifyTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyTableResponse{}
	_body, _err := client.ModifyTableWithOptions(instanceId, tableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a version of advanced configurations.
//
// Description:
//
// ## Method
//
// ~~~
//
// POST
//
// ~~~
//
// ## URI
//
// ~~~
//
// /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/actions/publish
//
// ~~~
//
// @param request - PublishAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishAdvanceConfigResponse
func (client *Client) PublishAdvanceConfigWithOptions(instanceId *string, configName *string, request *PublishAdvanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishAdvanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishAdvanceConfig"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/advanced-configs/" + tea.StringValue(openapiutil.GetEncodeParam(configName)) + "/actions/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishAdvanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a version of advanced configurations.
//
// Description:
//
// ## Method
//
// ~~~
//
// POST
//
// ~~~
//
// ## URI
//
// ~~~
//
// /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/actions/publish
//
// ~~~
//
// @param request - PublishAdvanceConfigRequest
//
// @return PublishAdvanceConfigResponse
func (client *Client) PublishAdvanceConfig(instanceId *string, configName *string, request *PublishAdvanceConfigRequest) (_result *PublishAdvanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishAdvanceConfigResponse{}
	_body, _err := client.PublishAdvanceConfigWithOptions(instanceId, configName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a specified index version.
//
// Description:
//
// ## Method
//
//     POST
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/actions/publish
//
// @param request - PublishIndexVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishIndexVersionResponse
func (client *Client) PublishIndexVersionWithOptions(instanceId *string, indexName *string, request *PublishIndexVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishIndexVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishIndexVersion"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/actions/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishIndexVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified index version.
//
// Description:
//
// ## Method
//
//     POST
//
// ## URI
//
//     /openapi/ha3/instances/{instanceId}/indexes/{indexName}/actions/publish
//
// @param request - PublishIndexVersionRequest
//
// @return PublishIndexVersionResponse
func (client *Client) PublishIndexVersion(instanceId *string, indexName *string, request *PublishIndexVersionRequest) (_result *PublishIndexVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishIndexVersionResponse{}
	_body, _err := client.PublishIndexVersionWithOptions(instanceId, indexName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PushDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDocumentsResponse
func (client *Client) PushDocumentsWithOptions(instanceId *string, dataSourceName *string, request *PushDocumentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushDocumentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PkField)) {
		query["pkField"] = request.PkField
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("PushDocuments"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceName)) + "/actions/bulk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushDocumentsRequest
//
// @return PushDocumentsResponse
func (client *Client) PushDocuments(instanceId *string, dataSourceName *string, request *PushDocumentsRequest) (_result *PushDocumentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushDocumentsResponse{}
	_body, _err := client.PushDocumentsWithOptions(instanceId, dataSourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores data from an index.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/recover-index`
//
// @param request - RecoverIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverIndexResponse
func (client *Client) RecoverIndexWithOptions(instanceId *string, request *RecoverIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecoverIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildDeployId)) {
		body["buildDeployId"] = request.BuildDeployId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceName)) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Generation)) {
		body["generation"] = request.Generation
	}

	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		body["indexName"] = request.IndexName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/recover-index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores data from an index.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/recover-index`
//
// @param request - RecoverIndexRequest
//
// @return RecoverIndexResponse
func (client *Client) RecoverIndex(instanceId *string, request *RecoverIndexRequest) (_result *RecoverIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecoverIndexResponse{}
	_body, _err := client.RecoverIndexWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 索引重建V2
//
// @param request - ReindexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReindexResponse
func (client *Client) ReindexWithOptions(instanceId *string, tableName *string, request *ReindexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReindexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataTimeSec)) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !tea.BoolValue(util.IsUnset(request.OssDataPath)) {
		body["ossDataPath"] = request.OssDataPath
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Reindex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/reindex"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReindexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 索引重建V2
//
// @param request - ReindexRequest
//
// @return ReindexResponse
func (client *Client) Reindex(instanceId *string, tableName *string, request *ReindexRequest) (_result *ReindexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReindexResponse{}
	_body, _err := client.ReindexWithOptions(instanceId, tableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cluster.
//
// Description:
//
// ### Method
//
// ```java
//
// DELETE
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters/{clusterName}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClusterResponse
func (client *Client) RemoveClusterWithOptions(instanceId *string, clusterName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveClusterResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveCluster"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/clusters/" + tea.StringValue(openapiutil.GetEncodeParam(clusterName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cluster.
//
// Description:
//
// ### Method
//
// ```java
//
// DELETE
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters/{clusterName}
//
// ```
//
// @return RemoveClusterResponse
func (client *Client) RemoveCluster(instanceId *string, clusterName *string) (_result *RemoveClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveClusterResponse{}
	_body, _err := client.RemoveClusterWithOptions(instanceId, clusterName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartIndexResponse
func (client *Client) StartIndexWithOptions(instanceId *string, indexName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartIndexResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/startIndex"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return StartIndexResponse
func (client *Client) StartIndex(instanceId *string, indexName *string) (_result *StartIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartIndexResponse{}
	_body, _err := client.StartIndexWithOptions(instanceId, indexName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopIndexResponse
func (client *Client) StopIndexWithOptions(instanceId *string, indexName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopIndexResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopIndex"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/indexes/" + tea.StringValue(openapiutil.GetEncodeParam(indexName)) + "/stopIndex"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return StopIndexResponse
func (client *Client) StopIndex(instanceId *string, indexName *string) (_result *StopIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopIndexResponse{}
	_body, _err := client.StopIndexWithOptions(instanceId, indexName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an FSM task.
//
// Description:
//
// ### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/stop-task/{fsmId}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskResponse
func (client *Client) StopTaskWithOptions(instanceId *string, fsmId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopTask"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/stop-task/" + tea.StringValue(openapiutil.GetEncodeParam(fsmId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an FSM task.
//
// Description:
//
// ### Method
//
// ```java
//
// PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/stop-task/{fsmId}
//
// ```
//
// @return StopTaskResponse
func (client *Client) StopTask(instanceId *string, fsmId *string) (_result *StopTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTaskResponse{}
	_body, _err := client.StopTaskWithOptions(instanceId, fsmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified instance.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Components)) {
		body["components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["orderType"] = request.OrderType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2021-10-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ha3/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified instance.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(instanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
