// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseVocAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetApiKey() *string
	SetContentTags(v []*SubmitEnterpriseVocAnalysisTaskRequestContentTags) *SubmitEnterpriseVocAnalysisTaskRequest
	GetContentTags() []*SubmitEnterpriseVocAnalysisTaskRequestContentTags
	SetContents(v []*SubmitEnterpriseVocAnalysisTaskRequestContents) *SubmitEnterpriseVocAnalysisTaskRequest
	GetContents() []*SubmitEnterpriseVocAnalysisTaskRequestContents
	SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetFileKey() *string
	SetFilterTags(v []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags) *SubmitEnterpriseVocAnalysisTaskRequest
	GetFilterTags() []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags
	SetMaterialType(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetMaterialType() *string
	SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetModelId() *string
	SetPositiveSample(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetPositiveSample() *string
	SetPositiveSampleFileKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetPositiveSampleFileKey() *string
	SetTaskType(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetTaskType() *string
	SetWorkspaceId(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetWorkspaceId() *string
}

type SubmitEnterpriseVocAnalysisTaskRequest struct {
	// The API key for integration access. For more information, see [Get an API key](https://help.aliyun.com/zh/model-studio/get-api-key?spm=a2c4g.11186623.help-menu-2400256.d_2_0_0.1cbdb0a8lsT1n3).
	//
	// example:
	//
	// sk-sdfs2-wewerwe-ere
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The content tags.
	//
	// This parameter is required.
	ContentTags []*SubmitEnterpriseVocAnalysisTaskRequestContentTags `json:"ContentTags,omitempty" xml:"ContentTags,omitempty" type:"Repeated"`
	// The material content to be mined.
	Contents []*SubmitEnterpriseVocAnalysisTaskRequestContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The key of the file.
	//
	// example:
	//
	// oss://default/bucket-name/materialDocument/tenant_agent/fileName
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// The filter tags.
	FilterTags []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags `json:"FilterTags,omitempty" xml:"FilterTags,omitempty" type:"Repeated"`
	// The material type. Valid values: \\`shortContent\\` (long or short comments, or tickets) and \\`dialogue\\` (dialogues).
	//
	// example:
	//
	// shortContent
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// The ID of the model.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The content of the positive sample.
	//
	// example:
	//
	// 正面样本
	PositiveSample *string `json:"PositiveSample,omitempty" xml:"PositiveSample,omitempty"`
	// The key of the positive sample file.
	//
	// example:
	//
	// oss://default/bucket-name/path/xxx.xlsx
	PositiveSampleFileKey *string `json:"PositiveSampleFileKey,omitempty" xml:"PositiveSampleFileKey,omitempty"`
	// The task type. Valid values: \\`lightAppSass\\` (invoked from a Software as a Service (SaaS) page) and \\`sdkBatchTask\\` (SDK batch task).
	//
	// example:
	//
	// lightAppSass
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The ID of the Model Studio workspace. For more information, see [Get a workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetContentTags() []*SubmitEnterpriseVocAnalysisTaskRequestContentTags {
	return s.ContentTags
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetContents() []*SubmitEnterpriseVocAnalysisTaskRequestContents {
	return s.Contents
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetFilterTags() []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	return s.FilterTags
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetPositiveSample() *string {
	return s.PositiveSample
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetPositiveSampleFileKey() *string {
	return s.PositiveSampleFileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetContentTags(v []*SubmitEnterpriseVocAnalysisTaskRequestContentTags) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ContentTags = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetContents(v []*SubmitEnterpriseVocAnalysisTaskRequestContents) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.Contents = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetFilterTags(v []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.FilterTags = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetMaterialType(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.MaterialType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetPositiveSample(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.PositiveSample = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetPositiveSampleFileKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.PositiveSampleFileKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetTaskType(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.TaskType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetWorkspaceId(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) Validate() error {
	if s.ContentTags != nil {
		for _, item := range s.ContentTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterTags != nil {
		for _, item := range s.FilterTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitEnterpriseVocAnalysisTaskRequestContentTags struct {
	// The tag definition.
	//
	// example:
	//
	// 一级标签-二级标签
	TagDefinePrompt *string `json:"TagDefinePrompt,omitempty" xml:"TagDefinePrompt,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 一级标签-二级标签
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The type of the tag mining task. Valid values: \\`singleTagValue\\` (single tag), \\`multiTagValues\\` (multiple tags), \\`summaryAndOverview\\` (summarization and overview), and \\`filter\\` (filtering).
	//
	// example:
	//
	// singleTagValue
	TagTaskType *string `json:"TagTaskType,omitempty" xml:"TagTaskType,omitempty"`
	// The definition of the tag value. For a tag mining task, this is a comma-separated list of enumerated values. For a summarization task, this is the summarization prompt.
	//
	// example:
	//
	// 是,否
	TagValueDefinePrompt *string `json:"TagValueDefinePrompt,omitempty" xml:"TagValueDefinePrompt,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContentTags) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContentTags) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) GetTagName() *string {
	return s.TagName
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) GetTagTaskType() *string {
	return s.TagTaskType
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) GetTagValueDefinePrompt() *string {
	return s.TagValueDefinePrompt
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) SetTagDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestContentTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) SetTagName(v string) *SubmitEnterpriseVocAnalysisTaskRequestContentTags {
	s.TagName = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) SetTagTaskType(v string) *SubmitEnterpriseVocAnalysisTaskRequestContentTags {
	s.TagTaskType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) SetTagValueDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestContentTags {
	s.TagValueDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContentTags) Validate() error {
	return dara.Validate(s)
}

type SubmitEnterpriseVocAnalysisTaskRequestContents struct {
	// Additional information provided directly to the large language model (LLM).
	//
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The material to be mined.
	//
	// example:
	//
	// 内容文本
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContents) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContents) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) GetText() *string {
	return s.Text
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) SetExtraInfo(v string) *SubmitEnterpriseVocAnalysisTaskRequestContents {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) SetText(v string) *SubmitEnterpriseVocAnalysisTaskRequestContents {
	s.Text = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) Validate() error {
	return dara.Validate(s)
}

type SubmitEnterpriseVocAnalysisTaskRequestFilterTags struct {
	// The tag definition.
	//
	// example:
	//
	// 一级标签-二级标签
	TagDefinePrompt *string `json:"TagDefinePrompt,omitempty" xml:"TagDefinePrompt,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 一级标签-二级标签
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The type of the tag mining task. Valid values: \\`singleTagValue\\` (single tag), \\`multiTagValues\\` (multiple tags), \\`summaryAndOverview\\` (summarization and overview), and \\`filter\\` (filtering).
	//
	// example:
	//
	// singleTagValue
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
	// The definition of the tag value. For a tag mining task, this is a comma-separated list of enumerated values. For a summarization task, this is the summarization prompt.
	//
	// example:
	//
	// 是,否
	TagValueDefinePrompt *string `json:"TagValueDefinePrompt,omitempty" xml:"TagValueDefinePrompt,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequestFilterTags) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequestFilterTags) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) GetTagName() *string {
	return s.TagName
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) GetTagType() *string {
	return s.TagType
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) GetTagValueDefinePrompt() *string {
	return s.TagValueDefinePrompt
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagName(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagName = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagType(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagValueDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagValueDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) Validate() error {
	return dara.Validate(s)
}
