// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseVocAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetApiKey() *string
	SetContentsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetContentsShrink() *string
	SetExtraInfo(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetExtraInfo() *string
	SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetFileKey() *string
	SetFilterTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetFilterTagsShrink() *string
	SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetModelId() *string
	SetOutputFormat(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetOutputFormat() *string
	SetSourceTrace(v bool) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetSourceTrace() *bool
	SetTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetTaskDescription() *string
	SetUrl(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetUrl() *string
}

type SubmitEnterpriseVocAnalysisTaskShrinkRequest struct {
	ApiKey         *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	ContentsShrink *string `json:"contents,omitempty" xml:"contents,omitempty"`
	ExtraInfo      *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// oss://default/aimiaobi-service-prod/aimiaobi/temp/public/government_service_experience_feedback_summary.txt
	FileKey          *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	FilterTagsShrink *string `json:"filterTags,omitempty" xml:"filterTags,omitempty"`
	// example:
	//
	// qwen-max
	ModelId         *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	OutputFormat    *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	SourceTrace     *bool   `json:"sourceTrace,omitempty" xml:"sourceTrace,omitempty"`
	TagsShrink      *string `json:"tags,omitempty" xml:"tags,omitempty"`
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetFilterTagsShrink() *string {
	return s.FilterTagsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetSourceTrace() *bool {
	return s.SourceTrace
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetContentsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetExtraInfo(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetFilterTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.FilterTagsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetOutputFormat(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetSourceTrace(v bool) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.SourceTrace = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetTaskDescription(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetUrl(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.Url = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
