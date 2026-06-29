// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPipelineByIdResponseBody
	GetCode() *string
	SetData(v *GetPipelineByIdResponseBodyData) *GetPipelineByIdResponseBody
	GetData() *GetPipelineByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetPipelineByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPipelineByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPipelineByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineByIdResponseBody
	GetSuccess() *bool
}

type GetPipelineByIdResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The pipeline task details.
	Data *GetPipelineByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error details returned by the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPipelineByIdResponseBody) GetData() *GetPipelineByIdResponseBodyData {
	return s.Data
}

func (s *GetPipelineByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPipelineByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPipelineByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineByIdResponseBody) SetCode(v string) *GetPipelineByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPipelineByIdResponseBody) SetData(v *GetPipelineByIdResponseBodyData) *GetPipelineByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetPipelineByIdResponseBody) SetHttpStatusCode(v int32) *GetPipelineByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPipelineByIdResponseBody) SetMessage(v string) *GetPipelineByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetPipelineByIdResponseBody) SetRequestId(v string) *GetPipelineByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineByIdResponseBody) SetSuccess(v bool) *GetPipelineByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineByIdResponseBodyData struct {
	// The configuration mode of the integration pipeline.
	//
	// example:
	//
	// PIPELINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The basic information of the pipeline task.
	NodeInfo *GetPipelineByIdResponseBodyDataNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The component configuration of the integration pipeline.
	PipelineConfig *GetPipelineByIdResponseBodyDataPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// The script mode configuration of the integration pipeline.
	//
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// The pipeline task type.
	//
	// example:
	//
	// 123
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// The schedule configuration of the integration pipeline. The value is a JSON string. Deserialize it by using the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.vo.OAScheduleConfigVO.
	//
	// example:
	//
	// {}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The channel configuration of the integration pipeline. The value is a JSON string. Deserialize it by using the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.OAPipelineSetting.
	//
	// example:
	//
	// {}
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s GetPipelineByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBodyData) GetMode() *string {
	return s.Mode
}

func (s *GetPipelineByIdResponseBodyData) GetNodeInfo() *GetPipelineByIdResponseBodyDataNodeInfo {
	return s.NodeInfo
}

func (s *GetPipelineByIdResponseBodyData) GetPipelineConfig() *GetPipelineByIdResponseBodyDataPipelineConfig {
	return s.PipelineConfig
}

func (s *GetPipelineByIdResponseBodyData) GetPipelineJson() *string {
	return s.PipelineJson
}

func (s *GetPipelineByIdResponseBodyData) GetPipelineType() *int32 {
	return s.PipelineType
}

func (s *GetPipelineByIdResponseBodyData) GetScheduleConfig() *string {
	return s.ScheduleConfig
}

func (s *GetPipelineByIdResponseBodyData) GetSettings() *string {
	return s.Settings
}

func (s *GetPipelineByIdResponseBodyData) SetMode(v string) *GetPipelineByIdResponseBodyData {
	s.Mode = &v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetNodeInfo(v *GetPipelineByIdResponseBodyDataNodeInfo) *GetPipelineByIdResponseBodyData {
	s.NodeInfo = v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetPipelineConfig(v *GetPipelineByIdResponseBodyDataPipelineConfig) *GetPipelineByIdResponseBodyData {
	s.PipelineConfig = v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetPipelineJson(v string) *GetPipelineByIdResponseBodyData {
	s.PipelineJson = &v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetPipelineType(v int32) *GetPipelineByIdResponseBodyData {
	s.PipelineType = &v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetScheduleConfig(v string) *GetPipelineByIdResponseBodyData {
	s.ScheduleConfig = &v
	return s
}

func (s *GetPipelineByIdResponseBodyData) SetSettings(v string) *GetPipelineByIdResponseBodyData {
	s.Settings = &v
	return s
}

func (s *GetPipelineByIdResponseBodyData) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PipelineConfig != nil {
		if err := s.PipelineConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineByIdResponseBodyDataNodeInfo struct {
	// The task description.
	//
	// example:
	//
	// comment
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The folder of the integration pipeline task node. The default value is the root folder. The folder must exist. If it does not exist, call the relevant API operation to create a folder of the offlinePipeline type.
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The pipeline file ID. This parameter is empty when the task is first created. When updating a pipeline task, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The scheduling node ID of the pipeline task. This parameter is empty when the task is first created. When updating a pipeline task, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the integration pipeline task.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The pipeline task ID. This parameter is empty when the task is first created. When updating a pipeline task, specify at least one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s GetPipelineByIdResponseBodyDataNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBodyDataNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetDesc() *string {
	return s.Desc
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetFileId() *int64 {
	return s.FileId
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetDesc(v string) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.Desc = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetDirectory(v string) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.Directory = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetFileId(v int64) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.FileId = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetNodeId(v string) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.NodeId = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetNodeName(v string) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.NodeName = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) SetPipelineId(v int64) *GetPipelineByIdResponseBodyDataNodeInfo {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataNodeInfo) Validate() error {
	return dara.Validate(s)
}

type GetPipelineByIdResponseBodyDataPipelineConfig struct {
	// The DAG (directed acyclic graph) link configuration that describes the connections between all components.
	Hops []*GetPipelineByIdResponseBodyDataPipelineConfigHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The component configurations, including detailed configurations of all components used.
	Steps []*GetPipelineByIdResponseBodyDataPipelineConfigSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
}

func (s GetPipelineByIdResponseBodyDataPipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBodyDataPipelineConfig) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfig) GetHops() []*GetPipelineByIdResponseBodyDataPipelineConfigHops {
	return s.Hops
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfig) GetSteps() []*GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	return s.Steps
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfig) SetHops(v []*GetPipelineByIdResponseBodyDataPipelineConfigHops) *GetPipelineByIdResponseBodyDataPipelineConfig {
	s.Hops = v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfig) SetSteps(v []*GetPipelineByIdResponseBodyDataPipelineConfigSteps) *GetPipelineByIdResponseBodyDataPipelineConfig {
	s.Steps = v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfig) Validate() error {
	if s.Hops != nil {
		for _, item := range s.Hops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineByIdResponseBodyDataPipelineConfigHops struct {
	// For conditional distribution components, set this parameter to true when the downstream condition is true. Otherwise, set it to false.
	SendTo *bool `json:"SendTo,omitempty" xml:"SendTo,omitempty"`
	// The input step name, which corresponds to Steps[*].StepName.
	//
	// example:
	//
	// mysql_reader
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The output step name, which corresponds to Steps[*].StepName.
	//
	// example:
	//
	// odps_writer
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetPipelineByIdResponseBodyDataPipelineConfigHops) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBodyDataPipelineConfigHops) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) GetSendTo() *bool {
	return s.SendTo
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) GetSource() *string {
	return s.Source
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) GetTarget() *string {
	return s.Target
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) SetSendTo(v bool) *GetPipelineByIdResponseBodyDataPipelineConfigHops {
	s.SendTo = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) SetSource(v string) *GetPipelineByIdResponseBodyDataPipelineConfigHops {
	s.Source = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) SetTarget(v string) *GetPipelineByIdResponseBodyDataPipelineConfigHops {
	s.Target = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigHops) Validate() error {
	return dara.Validate(s)
}

type GetPipelineByIdResponseBodyDataPipelineConfigSteps struct {
	// Specifies the data distribution method when the current component has multiple downstream components. Valid values:
	//
	// - true: The data of the current component is sent to all downstream components in a round-robin manner. For example, if the current component has 100 records and two downstream components, each downstream component receives 50 records. Default value: true.
	//
	// - false: The full data of the current component is sent to all downstream components. For example, if the current component has 100 records and two downstream components, each downstream component receives 100 records.
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// The plugin ID. Each plugin has a unique identifier. Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepKey. Developers should inherit this component configuration class and implement the corresponding component configuration. Each component configuration has the same structure as the pipeline configuration created on the Dataphin console.
	//
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The specific component configuration in JSON string format. Refer to the toJsonString method of the subclasses of the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig. Developers should inherit this component configuration class and implement the corresponding component configuration. Each component configuration has the same structure as the pipeline configuration created on the Dataphin console.
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// The step name. Step names must be unique within the same pipeline task.
	//
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The component type. Valid values:
	//
	// - input: an input component.
	//
	// - output: an output component.
	//
	// - transfrom: a transformation component.
	//
	// - process: a flow control component.
	//
	// Refer to the utility class com.alibaba.dataphin.pipeline.common.facade.openapi.model.plugin.OABasePluginConfig#stepType. Developers should inherit this component configuration class and implement the corresponding component configuration. Each component configuration has the same structure as the pipeline configuration created on the Dataphin console.
	//
	// example:
	//
	// input
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s GetPipelineByIdResponseBodyDataPipelineConfigSteps) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdResponseBodyDataPipelineConfigSteps) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) GetIsDistribute() *bool {
	return s.IsDistribute
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) GetKey() *string {
	return s.Key
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) GetStepName() *string {
	return s.StepName
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) GetStepType() *string {
	return s.StepType
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) SetIsDistribute(v bool) *GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	s.IsDistribute = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) SetKey(v string) *GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	s.Key = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) SetPluginConfig(v string) *GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	s.PluginConfig = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) SetStepName(v string) *GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	s.StepName = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) SetStepType(v string) *GetPipelineByIdResponseBodyDataPipelineConfigSteps {
	s.StepType = &v
	return s
}

func (s *GetPipelineByIdResponseBodyDataPipelineConfigSteps) Validate() error {
	return dara.Validate(s)
}
