// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineResponseBody
	GetErrorMessage() *string
	SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody
	GetPipeline() *GetPipelineResponseBodyPipeline
	SetRequestId(v string) *GetPipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineResponseBody
	GetSuccess() *bool
}

type GetPipelineResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Pipeline     *GetPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineResponseBody) GetPipeline() *GetPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *GetPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineResponseBody) SetErrorCode(v string) *GetPipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineResponseBody) SetErrorMessage(v string) *GetPipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineResponseBody) SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetPipelineResponseBody) SetRequestId(v string) *GetPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineResponseBody) SetSuccess(v bool) *GetPipelineResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodyPipeline struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 112222122
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 0
	EnvId *int32 `json:"envId,omitempty" xml:"envId,omitempty"`
	// example:
	//
	// 日常环境
	EnvName *string `json:"envName,omitempty" xml:"envName,omitempty"`
	// example:
	//
	// 1111
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 112222122
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// example:
	//
	// 流水线
	Name           *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	PipelineConfig *GetPipelineResponseBodyPipelinePipelineConfig `json:"pipelineConfig,omitempty" xml:"pipelineConfig,omitempty" type:"Struct"`
	TagList        []*GetPipelineResponseBodyPipelineTagList      `json:"tagList,omitempty" xml:"tagList,omitempty" type:"Repeated"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipeline) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPipelineResponseBodyPipeline) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetPipelineResponseBodyPipeline) GetEnvId() *int32 {
	return s.EnvId
}

func (s *GetPipelineResponseBodyPipeline) GetEnvName() *string {
	return s.EnvName
}

func (s *GetPipelineResponseBodyPipeline) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetPipelineResponseBodyPipeline) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetPipelineResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *GetPipelineResponseBodyPipeline) GetPipelineConfig() *GetPipelineResponseBodyPipelinePipelineConfig {
	return s.PipelineConfig
}

func (s *GetPipelineResponseBodyPipeline) GetTagList() []*GetPipelineResponseBodyPipelineTagList {
	return s.TagList
}

func (s *GetPipelineResponseBodyPipeline) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetPipelineResponseBodyPipeline) SetCreateTime(v int64) *GetPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetCreatorAccountId(v string) *GetPipelineResponseBodyPipeline {
	s.CreatorAccountId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetEnvId(v int32) *GetPipelineResponseBodyPipeline {
	s.EnvId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetEnvName(v string) *GetPipelineResponseBodyPipeline {
	s.EnvName = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetGroupId(v int64) *GetPipelineResponseBodyPipeline {
	s.GroupId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetModifierAccountId(v string) *GetPipelineResponseBodyPipeline {
	s.ModifierAccountId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetName(v string) *GetPipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetPipelineConfig(v *GetPipelineResponseBodyPipelinePipelineConfig) *GetPipelineResponseBodyPipeline {
	s.PipelineConfig = v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetTagList(v []*GetPipelineResponseBodyPipelineTagList) *GetPipelineResponseBodyPipeline {
	s.TagList = v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetUpdateTime(v int64) *GetPipelineResponseBodyPipeline {
	s.UpdateTime = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) Validate() error {
	if s.PipelineConfig != nil {
		if err := s.PipelineConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineResponseBodyPipelinePipelineConfig struct {
	// example:
	//
	// schema: tb pipeline:   - name: 执行命令     stages:       - driven: AUTO         jobs:           - displayName: 执行命令             task: execution-component-production@10             identifier: \"10_1626147407245\"             templateType: task             templateSign: \"\"             templateBatchUpdate: \"N\"             extraInfo: \"\"             params:               version1: pre-jdk1.62               steps:                 - name: 执行命令                   stepType: exec-shell                   stepIdentifier: \"10_1626147407245__11_1626147407249\"                   command: |                     # input your command here                     echo hello,world!                   ARTIFACTS: \"\"                   JSONEncoding: true                   freeInTaskGroupModeFields:                     - ARTIFACTS                   source: 132504-sss_ddd_3mvJ               ENGINE_PIPELINE_NAME: \"${INPUTS.ENGINE_PIPELINE_NAME}\"               ENGINE_PIPELINE_ID: \"${INPUTS.ENGINE_PIPELINE_ID}\"               ENGINE_PIPELINE_INST_ID: \"${INPUTS.ENGINE_PIPELINE_INST_ID}\"               ENGINE_PIPELINE_INST_NUMBER: \"${INPUTS.ENGINE_PIPELINE_INST_NUMBER}\"               buildNodeGroup: K8S-4             plugins: []             output: []             freeInTaskGroupModeFields: []
	Flow *string `json:"flow,omitempty" xml:"flow,omitempty"`
	// example:
	//
	// {}
	Settings *string                                                 `json:"settings,omitempty" xml:"settings,omitempty"`
	Sources  []*GetPipelineResponseBodyPipelinePipelineConfigSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfig) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) GetFlow() *string {
	return s.Flow
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) GetSettings() *string {
	return s.Settings
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) GetSources() []*GetPipelineResponseBodyPipelinePipelineConfigSources {
	return s.Sources
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetFlow(v string) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Flow = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetSettings(v string) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Settings = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetSources(v []*GetPipelineResponseBodyPipelinePipelineConfigSources) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Sources = v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineResponseBodyPipelinePipelineConfigSources struct {
	Data *GetPipelineResponseBodyPipelinePipelineConfigSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// xxsxsxs
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
	// example:
	//
	// Codeup
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSources) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSources) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) GetData() *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	return s.Data
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) GetSign() *string {
	return s.Sign
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetData(v *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Data = v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetSign(v string) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Sign = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetType(v string) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodyPipelinePipelineConfigSourcesData struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 1
	CloneDepth *int64 `json:"cloneDepth,omitempty" xml:"cloneDepth,omitempty"`
	// Credential Id
	//
	// example:
	//
	// 222
	CredentialId *int64 `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Credential Label
	//
	// example:
	//
	// 企业公钥
	CredentialLabel *string `json:"credentialLabel,omitempty" xml:"credentialLabel,omitempty"`
	// Credential Type
	//
	// example:
	//
	// region-ssh
	CredentialType *string   `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	Events         []*string `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsBranchMode *bool `json:"isBranchMode,omitempty" xml:"isBranchMode,omitempty"`
	// example:
	//
	// true
	IsCloneDepth *bool `json:"isCloneDepth,omitempty" xml:"isCloneDepth,omitempty"`
	// example:
	//
	// false
	IsSubmodule *bool `json:"isSubmodule,omitempty" xml:"isSubmodule,omitempty"`
	// example:
	//
	// true
	IsTrigger *bool `json:"isTrigger,omitempty" xml:"isTrigger,omitempty"`
	// example:
	//
	// cdup/ss
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// asasasas
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/test.git
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// example:
	//
	// 12
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// example:
	//
	// .*
	TriggerFilter *string `json:"triggerFilter,omitempty" xml:"triggerFilter,omitempty"`
	// example:
	//
	// https://flow.aliyun/webhook/asassasa
	Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSourcesData) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetBranch() *string {
	return s.Branch
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetCloneDepth() *int64 {
	return s.CloneDepth
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetCredentialId() *int64 {
	return s.CredentialId
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetCredentialLabel() *string {
	return s.CredentialLabel
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetEvents() []*string {
	return s.Events
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetIsBranchMode() *bool {
	return s.IsBranchMode
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetIsCloneDepth() *bool {
	return s.IsCloneDepth
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetIsSubmodule() *bool {
	return s.IsSubmodule
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetIsTrigger() *bool {
	return s.IsTrigger
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetLabel() *string {
	return s.Label
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetRepo() *string {
	return s.Repo
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetTriggerFilter() *string {
	return s.TriggerFilter
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GetWebhook() *string {
	return s.Webhook
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetBranch(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Branch = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCloneDepth(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CloneDepth = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialId(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialId = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialLabel(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialLabel = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialType(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialType = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetEvents(v []*string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Events = v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsBranchMode(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsBranchMode = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsCloneDepth(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsCloneDepth = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsSubmodule(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsSubmodule = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsTrigger(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsTrigger = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetLabel(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Label = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetNamespace(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Namespace = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetRepo(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Repo = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetServiceConnectionId(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.ServiceConnectionId = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetTriggerFilter(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.TriggerFilter = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetWebhook(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Webhook = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodyPipelineTagList struct {
	// example:
	//
	// 22
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 标签1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPipelineResponseBodyPipelineTagList) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipelineTagList) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelineTagList) GetId() *int64 {
	return s.Id
}

func (s *GetPipelineResponseBodyPipelineTagList) GetName() *string {
	return s.Name
}

func (s *GetPipelineResponseBodyPipelineTagList) SetId(v int64) *GetPipelineResponseBodyPipelineTagList {
	s.Id = &v
	return s
}

func (s *GetPipelineResponseBodyPipelineTagList) SetName(v string) *GetPipelineResponseBodyPipelineTagList {
	s.Name = &v
	return s
}

func (s *GetPipelineResponseBodyPipelineTagList) Validate() error {
	return dara.Validate(s)
}
