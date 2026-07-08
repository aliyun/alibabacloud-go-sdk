// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevel(v string) *UpdateDatasetRequest
	GetAccessLevel() *string
	SetDatasetConfig(v *UpdateDatasetRequestDatasetConfig) *UpdateDatasetRequest
	GetDatasetConfig() *UpdateDatasetRequestDatasetConfig
	SetDatasetDescription(v string) *UpdateDatasetRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *UpdateDatasetRequest
	GetDatasetId() *int64
	SetSearchDatasetEnable(v int32) *UpdateDatasetRequest
	GetSearchDatasetEnable() *int32
	SetWorkspaceId(v string) *UpdateDatasetRequest
	GetWorkspaceId() *string
}

type UpdateDatasetRequest struct {
	// example:
	//
	// private
	AccessLevel *string `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	// The configurations for a third-party search dataset.
	DatasetConfig *UpdateDatasetRequestDatasetConfig `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty" type:"Struct"`
	// The description of the dataset.
	//
	// example:
	//
	// 企业自定义数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// Specifies whether to enable search for the dataset.
	//
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// The unique identifier of the Alibaba Cloud Model Studio workspace. For more information, see [Get a workspaceId]().
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetAccessLevel() *string {
	return s.AccessLevel
}

func (s *UpdateDatasetRequest) GetDatasetConfig() *UpdateDatasetRequestDatasetConfig {
	return s.DatasetConfig
}

func (s *UpdateDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *UpdateDatasetRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *UpdateDatasetRequest) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *UpdateDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetRequest) SetAccessLevel(v string) *UpdateDatasetRequest {
	s.AccessLevel = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetConfig(v *UpdateDatasetRequestDatasetConfig) *UpdateDatasetRequest {
	s.DatasetConfig = v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetDescription(v string) *UpdateDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetId(v int64) *UpdateDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetRequest) SetSearchDatasetEnable(v int32) *UpdateDatasetRequest {
	s.SearchDatasetEnable = &v
	return s
}

func (s *UpdateDatasetRequest) SetWorkspaceId(v string) *UpdateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	if s.DatasetConfig != nil {
		if err := s.DatasetConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfig struct {
	// The dataset configuration items.
	SearchSourceConfig *UpdateDatasetRequestDatasetConfigSearchSourceConfig `json:"SearchSourceConfig,omitempty" xml:"SearchSourceConfig,omitempty" type:"Struct"`
	// Third-party search: API definition.
	SearchSourceConfigs []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs `json:"SearchSourceConfigs,omitempty" xml:"SearchSourceConfigs,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestDatasetConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfig) GetSearchSourceConfig() *UpdateDatasetRequestDatasetConfigSearchSourceConfig {
	return s.SearchSourceConfig
}

func (s *UpdateDatasetRequestDatasetConfig) GetSearchSourceConfigs() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	return s.SearchSourceConfigs
}

func (s *UpdateDatasetRequestDatasetConfig) SetSearchSourceConfig(v *UpdateDatasetRequestDatasetConfigSearchSourceConfig) *UpdateDatasetRequestDatasetConfig {
	s.SearchSourceConfig = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfig) SetSearchSourceConfigs(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs) *UpdateDatasetRequestDatasetConfig {
	s.SearchSourceConfigs = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfig) Validate() error {
	if s.SearchSourceConfig != nil {
		if err := s.SearchSourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchSourceConfigs != nil {
		for _, item := range s.SearchSourceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfig struct {
	// Specifies whether the key-value pairs in metadata are used for generation. Default: true.
	//
	// example:
	//
	// true
	MetadataKeyValueGenerateEnable *bool `json:"MetadataKeyValueGenerateEnable,omitempty" xml:"MetadataKeyValueGenerateEnable,omitempty"`
	// Specifies whether the key-value pairs in metadata are included in searches. Default: true.
	//
	// example:
	//
	// true
	MetadataKeyValueSearchEnable *bool `json:"MetadataKeyValueSearchEnable,omitempty" xml:"MetadataKeyValueSearchEnable,omitempty"`
	// Specifies whether tags are used for generation. Default: true.
	//
	// example:
	//
	// true
	TagGenerateEnable *bool `json:"TagGenerateEnable,omitempty" xml:"TagGenerateEnable,omitempty"`
	// Specifies whether tags are included in searches. Default: true.
	//
	// example:
	//
	// true
	TagSearchEnable *bool `json:"TagSearchEnable,omitempty" xml:"TagSearchEnable,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) GetMetadataKeyValueGenerateEnable() *bool {
	return s.MetadataKeyValueGenerateEnable
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) GetMetadataKeyValueSearchEnable() *bool {
	return s.MetadataKeyValueSearchEnable
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) GetTagGenerateEnable() *bool {
	return s.TagGenerateEnable
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) GetTagSearchEnable() *bool {
	return s.TagSearchEnable
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) SetMetadataKeyValueGenerateEnable(v bool) *UpdateDatasetRequestDatasetConfigSearchSourceConfig {
	s.MetadataKeyValueGenerateEnable = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) SetMetadataKeyValueSearchEnable(v bool) *UpdateDatasetRequestDatasetConfigSearchSourceConfig {
	s.MetadataKeyValueSearchEnable = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) SetTagGenerateEnable(v bool) *UpdateDatasetRequestDatasetConfigSearchSourceConfig {
	s.TagGenerateEnable = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) SetTagSearchEnable(v bool) *UpdateDatasetRequestDatasetConfigSearchSourceConfig {
	s.TagSearchEnable = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigs struct {
	// A searchable keyword used to verify availability.
	//
	// example:
	//
	// 可以搜索到的关键词，用来验证是否可用
	DemoQuery *string `json:"DemoQuery,omitempty" xml:"DemoQuery,omitempty"`
	// The API request configuration.
	SearchSourceRequestConfig *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig `json:"SearchSourceRequestConfig,omitempty" xml:"SearchSourceRequestConfig,omitempty" type:"Struct"`
	// The API response configuration.
	SearchSourceResponseConfig *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig `json:"SearchSourceResponseConfig,omitempty" xml:"SearchSourceResponseConfig,omitempty" type:"Struct"`
	// The default number of data entries for requests and responses.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigs) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) GetDemoQuery() *string {
	return s.DemoQuery
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) GetSearchSourceRequestConfig() *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	return s.SearchSourceRequestConfig
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) GetSearchSourceResponseConfig() *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	return s.SearchSourceResponseConfig
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) GetSize() *int32 {
	return s.Size
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) SetDemoQuery(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.DemoQuery = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) SetSearchSourceRequestConfig(v *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) *UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.SearchSourceRequestConfig = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) SetSearchSourceResponseConfig(v *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) *UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.SearchSourceResponseConfig = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) SetSize(v int32) *UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	s.Size = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigs) Validate() error {
	if s.SearchSourceRequestConfig != nil {
		if err := s.SearchSourceRequestConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchSourceResponseConfig != nil {
		if err := s.SearchSourceResponseConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The connection timeout period, in milliseconds.
	//
	// example:
	//
	// 3000
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// The HTTP request headers.
	Headers []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The request method.
	//
	// example:
	//
	// 请求方式
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path parameters.
	Params []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Specifies whether to enable path parameters.
	//
	// example:
	//
	// true
	PathParamsEnable *bool `json:"PathParamsEnable,omitempty" xml:"PathParamsEnable,omitempty"`
	// The read timeout period, in milliseconds.
	//
	// example:
	//
	// 3000
	SocketTimeout *int32 `json:"SocketTimeout,omitempty" xml:"SocketTimeout,omitempty"`
	// The API URL.
	//
	// example:
	//
	// api地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetBody() *string {
	return s.Body
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetHeaders() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	return s.Headers
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetMethod() *string {
	return s.Method
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetParams() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	return s.Params
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetPathParamsEnable() *bool {
	return s.PathParamsEnable
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetSocketTimeout() *int32 {
	return s.SocketTimeout
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetUrl() *string {
	return s.Url
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetBody(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Body = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetConnectTimeout(v int32) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.ConnectTimeout = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetHeaders(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Headers = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetMethod(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Method = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetParams(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Params = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetPathParamsEnable(v bool) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.PathParamsEnable = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetSocketTimeout(v int32) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.SocketTimeout = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetUrl(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Url = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders struct {
	// The parameter name.
	//
	// example:
	//
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is valid only when ValueType is set to time.
	//
	// example:
	//
	// valueType = time 时有效
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	// The data type of the parameter value. Default: string.
	//
	// example:
	//
	// 参数值数据类型：默认string
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValue() *string {
	return s.Value
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetName(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValue(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Value = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueFormat(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueFormat = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueType(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueType = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams struct {
	// The parameter name.
	//
	// example:
	//
	// 参数名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 参数值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is valid only when ValueType is set to time.
	//
	// example:
	//
	// valueType = time 时有效
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	// The data type of the parameter value. Default: string.
	//
	// example:
	//
	// 参数值数据类型：默认string
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValue() *string {
	return s.Value
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetName(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValue(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Value = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueFormat(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueFormat = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueType(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueType = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig struct {
	// The node configuration.
	JqNodes []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GetJqNodes() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	return s.JqNodes
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) SetJqNodes(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	s.JqNodes = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) Validate() error {
	if s.JqNodes != nil {
		for _, item := range s.JqNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes struct {
	// The child node configuration.
	JqNodes []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
	// The node key.
	//
	// example:
	//
	// 节点key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The node path.
	//
	// example:
	//
	// 节点路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The data type of the node. Valid values: string, number, list, object, and base.
	//
	// example:
	//
	// 节点数据类型：string number list object base
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetJqNodes() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	return s.JqNodes
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetKey() *string {
	return s.Key
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetPath() *string {
	return s.Path
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetJqNodes(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.JqNodes = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetKey(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Key = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetPath(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Path = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetType(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) Validate() error {
	if s.JqNodes != nil {
		for _, item := range s.JqNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes struct {
	// The child node configuration.
	JqNodes []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
	// The node key.
	//
	// example:
	//
	// title
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The node path.
	//
	// example:
	//
	// .title
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type.
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetJqNodes() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	return s.JqNodes
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetJqNodes(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.JqNodes = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetKey(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetPath(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetType(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) Validate() error {
	if s.JqNodes != nil {
		for _, item := range s.JqNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes struct {
	// The node key.
	//
	// example:
	//
	// title
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The node path.
	//
	// example:
	//
	// .title
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type.
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetKey(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetPath(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetType(v string) *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) Validate() error {
	return dara.Validate(s)
}
