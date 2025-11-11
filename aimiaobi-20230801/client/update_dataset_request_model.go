// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
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
	DatasetConfig *UpdateDatasetRequestDatasetConfig `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty" type:"Struct"`
	// example:
	//
	// 企业自定义数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
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
	SearchSourceConfigs []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs `json:"SearchSourceConfigs,omitempty" xml:"SearchSourceConfigs,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestDatasetConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestDatasetConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestDatasetConfig) GetSearchSourceConfigs() []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs {
	return s.SearchSourceConfigs
}

func (s *UpdateDatasetRequestDatasetConfig) SetSearchSourceConfigs(v []*UpdateDatasetRequestDatasetConfigSearchSourceConfigs) *UpdateDatasetRequestDatasetConfig {
	s.SearchSourceConfigs = v
	return s
}

func (s *UpdateDatasetRequestDatasetConfig) Validate() error {
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

type UpdateDatasetRequestDatasetConfigSearchSourceConfigs struct {
	// example:
	//
	// 可以搜索到的关键词，用来验证是否可用
	DemoQuery                  *string                                                                         `json:"DemoQuery,omitempty" xml:"DemoQuery,omitempty"`
	SearchSourceRequestConfig  *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfig  `json:"SearchSourceRequestConfig,omitempty" xml:"SearchSourceRequestConfig,omitempty" type:"Struct"`
	SearchSourceResponseConfig *UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfig `json:"SearchSourceResponseConfig,omitempty" xml:"SearchSourceResponseConfig,omitempty" type:"Struct"`
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
	// example:
	//
	// {}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// 3000
	ConnectTimeout *int32                                                                                  `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	Headers        []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// example:
	//
	// 请求方式
	Method *string                                                                                `json:"Method,omitempty" xml:"Method,omitempty"`
	Params []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
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
	JqNodes []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
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
	JqNodes []*UpdateDatasetRequestDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
	// example:
	//
	// title
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// .title
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
	// example:
	//
	// title
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// .title
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
