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
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPipelineByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// example:
	//
	// PIPELINE
	Mode           *string                                        `json:"Mode,omitempty" xml:"Mode,omitempty"`
	NodeInfo       *GetPipelineByIdResponseBodyDataNodeInfo       `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	PipelineConfig *GetPipelineByIdResponseBodyDataPipelineConfig `json:"PipelineConfig,omitempty" xml:"PipelineConfig,omitempty" type:"Struct"`
	// example:
	//
	// {}
	PipelineJson *string `json:"PipelineJson,omitempty" xml:"PipelineJson,omitempty"`
	// example:
	//
	// 123
	PipelineType *int32 `json:"PipelineType,omitempty" xml:"PipelineType,omitempty"`
	// example:
	//
	// {}
	ScheduleConfig *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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
	// example:
	//
	// comment
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	Hops  []*GetPipelineByIdResponseBodyDataPipelineConfigHops  `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
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
	SendTo *bool `json:"SendTo,omitempty" xml:"SendTo,omitempty"`
	// example:
	//
	// mysql_reader
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
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
	IsDistribute *bool `json:"IsDistribute,omitempty" xml:"IsDistribute,omitempty"`
	// example:
	//
	// mysqlinput
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// example:
	//
	// mysql_reader
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
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
