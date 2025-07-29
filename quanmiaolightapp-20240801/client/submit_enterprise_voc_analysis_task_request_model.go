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
	SetContents(v []*SubmitEnterpriseVocAnalysisTaskRequestContents) *SubmitEnterpriseVocAnalysisTaskRequest
	GetContents() []*SubmitEnterpriseVocAnalysisTaskRequestContents
	SetExtraInfo(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetExtraInfo() *string
	SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetFileKey() *string
	SetFilterTags(v []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags) *SubmitEnterpriseVocAnalysisTaskRequest
	GetFilterTags() []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags
	SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetModelId() *string
	SetOutputFormat(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetOutputFormat() *string
	SetSourceTrace(v bool) *SubmitEnterpriseVocAnalysisTaskRequest
	GetSourceTrace() *bool
	SetTags(v []*SubmitEnterpriseVocAnalysisTaskRequestTags) *SubmitEnterpriseVocAnalysisTaskRequest
	GetTags() []*SubmitEnterpriseVocAnalysisTaskRequestTags
	SetTaskDescription(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetTaskDescription() *string
	SetUrl(v string) *SubmitEnterpriseVocAnalysisTaskRequest
	GetUrl() *string
}

type SubmitEnterpriseVocAnalysisTaskRequest struct {
	ApiKey    *string                                           `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Contents  []*SubmitEnterpriseVocAnalysisTaskRequestContents `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	ExtraInfo *string                                           `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// oss://default/aimiaobi-service-prod/aimiaobi/temp/public/government_service_experience_feedback_summary.txt
	FileKey    *string                                             `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	FilterTags []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags `json:"filterTags,omitempty" xml:"filterTags,omitempty" type:"Repeated"`
	// example:
	//
	// qwen-max
	ModelId         *string                                       `json:"modelId,omitempty" xml:"modelId,omitempty"`
	OutputFormat    *string                                       `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	SourceTrace     *bool                                         `json:"sourceTrace,omitempty" xml:"sourceTrace,omitempty"`
	Tags            []*SubmitEnterpriseVocAnalysisTaskRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TaskDescription *string                                       `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
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

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetContents() []*SubmitEnterpriseVocAnalysisTaskRequestContents {
	return s.Contents
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetFilterTags() []*SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	return s.FilterTags
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetSourceTrace() *bool {
	return s.SourceTrace
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetTags() []*SubmitEnterpriseVocAnalysisTaskRequestTags {
	return s.Tags
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetContents(v []*SubmitEnterpriseVocAnalysisTaskRequestContents) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.Contents = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetExtraInfo(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ExtraInfo = &v
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

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetOutputFormat(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetSourceTrace(v bool) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.SourceTrace = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetTags(v []*SubmitEnterpriseVocAnalysisTaskRequestTags) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.Tags = v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetTaskDescription(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) SetUrl(v string) *SubmitEnterpriseVocAnalysisTaskRequest {
	s.Url = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitEnterpriseVocAnalysisTaskRequestContents struct {
	// example:
	//
	// id-xxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContents) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequestContents) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) GetId() *string {
	return s.Id
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) GetText() *string {
	return s.Text
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestContents) SetId(v string) *SubmitEnterpriseVocAnalysisTaskRequestContents {
	s.Id = &v
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
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	TagName         *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
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

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) SetTagName(v string) *SubmitEnterpriseVocAnalysisTaskRequestFilterTags {
	s.TagName = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestFilterTags) Validate() error {
	return dara.Validate(s)
}

type SubmitEnterpriseVocAnalysisTaskRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskRequestTags) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestTags) SetTagDefinePrompt(v string) *SubmitEnterpriseVocAnalysisTaskRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestTags) SetTagName(v string) *SubmitEnterpriseVocAnalysisTaskRequestTags {
	s.TagName = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskRequestTags) Validate() error {
	return dara.Validate(s)
}
