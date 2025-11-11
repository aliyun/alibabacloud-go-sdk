// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDatasetResponseBody
	GetCode() *string
	SetData(v *GetDatasetResponseBodyData) *GetDatasetResponseBody
	GetData() *GetDatasetResponseBodyData
	SetHttpStatusCode(v int32) *GetDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetResponseBody
	GetSuccess() *bool
}

type GetDatasetResponseBody struct {
	// example:
	//
	// NoData
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDatasetResponseBody) GetData() *GetDatasetResponseBodyData {
	return s.Data
}

func (s *GetDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetResponseBody) SetCode(v string) *GetDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatasetResponseBody) SetData(v *GetDatasetResponseBodyData) *GetDatasetResponseBody {
	s.Data = v
	return s
}

func (s *GetDatasetResponseBody) SetHttpStatusCode(v int32) *GetDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDatasetResponseBody) SetMessage(v string) *GetDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSuccess(v bool) *GetDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyData struct {
	// example:
	//
	// 2024-11-12 21:46:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUser    *string                                  `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DatasetConfig *GetDatasetResponseBodyDataDatasetConfig `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// xxx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType          *string                                         `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	DocumentHandleConfig *GetDatasetResponseBodyDataDocumentHandleConfig `json:"DocumentHandleConfig,omitempty" xml:"DocumentHandleConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
}

func (s GetDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDatasetResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDatasetResponseBodyData) GetDatasetConfig() *GetDatasetResponseBodyDataDatasetConfig {
	return s.DatasetConfig
}

func (s *GetDatasetResponseBodyData) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *GetDatasetResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetResponseBodyData) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetResponseBodyData) GetDatasetType() *string {
	return s.DatasetType
}

func (s *GetDatasetResponseBodyData) GetDocumentHandleConfig() *GetDatasetResponseBodyDataDocumentHandleConfig {
	return s.DocumentHandleConfig
}

func (s *GetDatasetResponseBodyData) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *GetDatasetResponseBodyData) SetCreateTime(v string) *GetDatasetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetCreateUser(v string) *GetDatasetResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetDatasetConfig(v *GetDatasetResponseBodyDataDatasetConfig) *GetDatasetResponseBodyData {
	s.DatasetConfig = v
	return s
}

func (s *GetDatasetResponseBodyData) SetDatasetDescription(v string) *GetDatasetResponseBodyData {
	s.DatasetDescription = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetDatasetId(v int64) *GetDatasetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetDatasetName(v string) *GetDatasetResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetDatasetType(v string) *GetDatasetResponseBodyData {
	s.DatasetType = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetDocumentHandleConfig(v *GetDatasetResponseBodyDataDocumentHandleConfig) *GetDatasetResponseBodyData {
	s.DocumentHandleConfig = v
	return s
}

func (s *GetDatasetResponseBodyData) SetSearchDatasetEnable(v int32) *GetDatasetResponseBodyData {
	s.SearchDatasetEnable = &v
	return s
}

func (s *GetDatasetResponseBodyData) Validate() error {
	if s.DatasetConfig != nil {
		if err := s.DatasetConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DocumentHandleConfig != nil {
		if err := s.DocumentHandleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetResponseBodyDataDatasetConfig struct {
	SearchSourceConfigs []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs `json:"SearchSourceConfigs,omitempty" xml:"SearchSourceConfigs,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodyDataDatasetConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfig) GetSearchSourceConfigs() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs {
	return s.SearchSourceConfigs
}

func (s *GetDatasetResponseBodyDataDatasetConfig) SetSearchSourceConfigs(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) *GetDatasetResponseBodyDataDatasetConfig {
	s.SearchSourceConfigs = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfig) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs struct {
	// example:
	//
	// 可以搜索到的关键词，用来验证是否可用
	DemoQuery                  *string                                                                               `json:"DemoQuery,omitempty" xml:"DemoQuery,omitempty"`
	SearchSourceRequestConfig  *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig  `json:"SearchSourceRequestConfig,omitempty" xml:"SearchSourceRequestConfig,omitempty" type:"Struct"`
	SearchSourceResponseConfig *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig `json:"SearchSourceResponseConfig,omitempty" xml:"SearchSourceResponseConfig,omitempty" type:"Struct"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) GetDemoQuery() *string {
	return s.DemoQuery
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) GetSearchSourceRequestConfig() *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	return s.SearchSourceRequestConfig
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) GetSearchSourceResponseConfig() *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	return s.SearchSourceResponseConfig
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) GetSize() *int32 {
	return s.Size
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) SetDemoQuery(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs {
	s.DemoQuery = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) SetSearchSourceRequestConfig(v *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs {
	s.SearchSourceRequestConfig = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) SetSearchSourceResponseConfig(v *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs {
	s.SearchSourceResponseConfig = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) SetSize(v int32) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs {
	s.Size = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigs) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig struct {
	// example:
	//
	// {}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// 30
	ConnectTimeout *int32                                                                                        `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	Headers        []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// example:
	//
	// 请求方式
	Method *string                                                                                      `json:"Method,omitempty" xml:"Method,omitempty"`
	Params []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PathParamsEnable *bool `json:"PathParamsEnable,omitempty" xml:"PathParamsEnable,omitempty"`
	// example:
	//
	// 78
	SocketTimeout *int32 `json:"SocketTimeout,omitempty" xml:"SocketTimeout,omitempty"`
	// example:
	//
	// api地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetBody() *string {
	return s.Body
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetHeaders() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	return s.Headers
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetMethod() *string {
	return s.Method
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetParams() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	return s.Params
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetPathParamsEnable() *bool {
	return s.PathParamsEnable
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetSocketTimeout() *int32 {
	return s.SocketTimeout
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) GetUrl() *string {
	return s.Url
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetBody(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Body = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetConnectTimeout(v int32) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.ConnectTimeout = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetHeaders(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Headers = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetMethod(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Method = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetParams(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Params = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetPathParamsEnable(v bool) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.PathParamsEnable = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetSocketTimeout(v int32) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.SocketTimeout = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) SetUrl(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig {
	s.Url = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfig) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders struct {
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

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValue() *string {
	return s.Value
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) GetValueType() *string {
	return s.ValueType
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetName(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValue(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.Value = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueFormat(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueFormat = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) SetValueType(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders {
	s.ValueType = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigHeaders) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams struct {
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

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetName() *string {
	return s.Name
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValue() *string {
	return s.Value
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) GetValueType() *string {
	return s.ValueType
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetName(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValue(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.Value = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueFormat(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueFormat = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) SetValueType(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams {
	s.ValueType = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceRequestConfigParams) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig struct {
	JqNodes []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) GetJqNodes() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	return s.JqNodes
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) SetJqNodes(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig {
	s.JqNodes = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfig) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes struct {
	JqNodes []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
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

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetJqNodes() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	return s.JqNodes
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetKey() *string {
	return s.Key
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetPath() *string {
	return s.Path
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetJqNodes(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.JqNodes = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetKey(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Key = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetPath(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Path = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) SetType(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodes) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes struct {
	JqNodes []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes `json:"JqNodes,omitempty" xml:"JqNodes,omitempty" type:"Repeated"`
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

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetJqNodes() []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	return s.JqNodes
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetJqNodes(v []*GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.JqNodes = v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetKey(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetPath(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) SetType(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodes) Validate() error {
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

type GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes struct {
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

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetKey() *string {
	return s.Key
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetPath() *string {
	return s.Path
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) GetType() *string {
	return s.Type
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetKey(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Key = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetPath(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Path = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) SetType(v string) *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes {
	s.Type = &v
	return s
}

func (s *GetDatasetResponseBodyDataDatasetConfigSearchSourceConfigsSearchSourceResponseConfigJqNodesJqNodesJqNodes) Validate() error {
	return dara.Validate(s)
}

type GetDatasetResponseBodyDataDocumentHandleConfig struct {
	// example:
	//
	// true
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
}

func (s GetDatasetResponseBodyDataDocumentHandleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBodyDataDocumentHandleConfig) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyDataDocumentHandleConfig) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *GetDatasetResponseBodyDataDocumentHandleConfig) SetDisableHandleMultimodalMedia(v bool) *GetDatasetResponseBodyDataDocumentHandleConfig {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *GetDatasetResponseBodyDataDocumentHandleConfig) Validate() error {
	return dara.Validate(s)
}
