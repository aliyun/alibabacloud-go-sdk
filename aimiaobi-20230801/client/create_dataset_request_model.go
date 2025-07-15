// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetConfig(v *CreateDatasetRequestDatasetConfig) *CreateDatasetRequest
	GetDatasetConfig() *CreateDatasetRequestDatasetConfig
	SetDatasetDescription(v string) *CreateDatasetRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateDatasetRequest
	GetDatasetName() *string
	SetDatasetType(v string) *CreateDatasetRequest
	GetDatasetType() *string
	SetDocumentHandleConfig(v *CreateDatasetRequestDocumentHandleConfig) *CreateDatasetRequest
	GetDocumentHandleConfig() *CreateDatasetRequestDocumentHandleConfig
	SetInvokeType(v string) *CreateDatasetRequest
	GetInvokeType() *string
	SetSearchDatasetEnable(v int32) *CreateDatasetRequest
	GetSearchDatasetEnable() *int32
	SetWorkspaceId(v string) *CreateDatasetRequest
	GetWorkspaceId() *string
}

type CreateDatasetRequest struct {
	DatasetConfig *CreateDatasetRequestDatasetConfig `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty" type:"Struct"`
	// example:
	//
	// 企业自定义数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// businessDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType          *string                                   `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	DocumentHandleConfig *CreateDatasetRequestDocumentHandleConfig `json:"DocumentHandleConfig,omitempty" xml:"DocumentHandleConfig,omitempty" type:"Struct"`
	// example:
	//
	// portal
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetDatasetConfig() *CreateDatasetRequestDatasetConfig {
	return s.DatasetConfig
}

func (s *CreateDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateDatasetRequest) GetDocumentHandleConfig() *CreateDatasetRequestDocumentHandleConfig {
	return s.DocumentHandleConfig
}

func (s *CreateDatasetRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *CreateDatasetRequest) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *CreateDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetRequest) SetDatasetConfig(v *CreateDatasetRequestDatasetConfig) *CreateDatasetRequest {
	s.DatasetConfig = v
	return s
}

func (s *CreateDatasetRequest) SetDatasetDescription(v string) *CreateDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetType(v string) *CreateDatasetRequest {
	s.DatasetType = &v
	return s
}

func (s *CreateDatasetRequest) SetDocumentHandleConfig(v *CreateDatasetRequestDocumentHandleConfig) *CreateDatasetRequest {
	s.DocumentHandleConfig = v
	return s
}

func (s *CreateDatasetRequest) SetInvokeType(v string) *CreateDatasetRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateDatasetRequest) SetSearchDatasetEnable(v int32) *CreateDatasetRequest {
	s.SearchDatasetEnable = &v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v string) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfig struct {
	SearchSourceConfigs []*CreateDatasetRequestDatasetConfigSearchSourceConfigs `json:"SearchSourceConfigs,omitempty" xml:"SearchSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequestDatasetConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfig) GetSearchSourceConfigs() []*CreateDatasetRequestDatasetConfigSearchSourceConfigs {
	return s.SearchSourceConfigs
}

func (s *CreateDatasetRequestDatasetConfig) SetSearchSourceConfigs(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigs) *CreateDatasetRequestDatasetConfig {
	s.SearchSourceConfigs = v
	return s
}

func (s *CreateDatasetRequestDatasetConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigs struct {
	// example:
	//
	// 可以搜索到的关键词，用来验证是否可用
	DemoQuery                  *string                                                                         `json:"DemoQuery,omitempty" xml:"DemoQuery,omitempty"`
	SearchSourceRequestConfig  *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig  `json:"SearchSourceRequestConfig,omitempty" xml:"SearchSourceRequestConfig,omitempty" type:"Struct"`
	SearchSourceResponseConfig *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig `json:"SearchSourceResponseConfig,omitempty" xml:"SearchSourceResponseConfig,omitempty" type:"Struct"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) GetDemoQuery() *string {
	return s.DemoQuery
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) GetSearchSourceRequestConfig() *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	return s.SearchSourceRequestConfig
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) GetSearchSourceResponseConfig() *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	return s.SearchSourceResponseConfig
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) GetSize() *int32 {
	return s.Size
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) SetDemoQuery(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.DemoQuery = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) SetSearchSourceRequestConfig(v *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) *CreateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.SearchSourceRequestConfig = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) SetSearchSourceResponseConfig(v *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) *CreateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.SearchSourceResponseConfig = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) SetSize(v int32) *CreateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.Size = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig struct {
	// example:
	//
	// {}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// 3000
	ConnectTimeout *int32                                                                                  `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	Headers        []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// example:
	//
	// 请求方式
	Method *string                                                                                `json:"Method,omitempty" xml:"Method,omitempty"`
	Params []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PathParamsEnable *bool `json:"PathParamsEnable,omitempty" xml:"PathParamsEnable,omitempty"`
	// example:
	//
	// 3000
	SocketTimeout *int32 `json:"SocketTimeout,omitempty" xml:"SocketTimeout,omitempty"`
	// example:
	//
	// api地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetBody() *string {
	return s.Body
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetHeaders() []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	return s.Headers
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetMethod() *string {
	return s.Method
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetParams() []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	return s.Params
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetPathParamsEnable() *bool {
	return s.PathParamsEnable
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetSocketTimeout() *int32 {
	return s.SocketTimeout
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetUrl() *string {
	return s.Url
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetBody(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Body = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetConnectTimeout(v int32) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetHeaders(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Headers = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetMethod(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Method = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetParams(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Params = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetPathParamsEnable(v bool) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.PathParamsEnable = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetSocketTimeout(v int32) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.SocketTimeout = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetUrl(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Url = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders struct {
	// example:
	//
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// valueType = time 时有效
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	// example:
	//
	// 参数值数据类型: 默认string
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValue() *string {
	return s.Value
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueType() *string {
	return s.ValueType
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetName(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValue(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Value = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueFormat(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueFormat = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueType(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueType = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams struct {
	// example:
	//
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// valueType = time 时有效
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	// example:
	//
	// 参数值数据类型: 默认string
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetName() *string {
	return s.Name
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValue() *string {
	return s.Value
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueType() *string {
	return s.ValueType
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetName(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValue(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Value = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueFormat(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueFormat = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueType(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueType = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig struct {
	JqNodes []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GetJqNodes() []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	return s.JqNodes
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) SetJqNodes(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	s.JqNodes = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes struct {
	JqNodes []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 节点key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 节点路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 节点数据类型：string number list object base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetJqNodes() []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	return s.JqNodes
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetKey() *string {
	return s.Key
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetPath() *string {
	return s.Path
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetJqNodes(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.JqNodes = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetKey(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Key = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetPath(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Path = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetType(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes struct {
	JqNodes []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
	// example:
	//
	// title
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetJqNodes() []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	return s.JqNodes
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetJqNodes(v []*CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.JqNodes = v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetKey(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetPath(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetType(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes struct {
	// example:
	//
	// title
	Key  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetKey(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetPath(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetType(v string) *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) Validate() error {
	return dara.Validate(s)
}

type CreateDatasetRequestDocumentHandleConfig struct {
	// example:
	//
	// false
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
}

func (s CreateDatasetRequestDocumentHandleConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequestDocumentHandleConfig) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequestDocumentHandleConfig) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *CreateDatasetRequestDocumentHandleConfig) SetDisableHandleMultimodalMedia(v bool) *CreateDatasetRequestDocumentHandleConfig {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *CreateDatasetRequestDocumentHandleConfig) Validate() error {
	return dara.Validate(s)
}
