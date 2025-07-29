// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTagMiningAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetApiKey() *string
	SetBusinessType(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetBusinessType() *string
	SetContentsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetContentsShrink() *string
	SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetExtraInfo() *string
	SetModelId(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetModelId() *string
	SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetOutputFormat() *string
	SetTagsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetTaskDescription() *string
	SetUrl(v string) *SubmitTagMiningAnalysisTaskShrinkRequest
	GetUrl() *string
}

type SubmitTagMiningAnalysisTaskShrinkRequest struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// clueMining
	BusinessType   *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	ContentsShrink *string `json:"contents,omitempty" xml:"contents,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 请返回如下JSON格式，{"key1":"","key2":""}
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	TagsShrink   *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetApiKey(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetBusinessType(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetContentsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetTagsShrink(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) SetUrl(v string) *SubmitTagMiningAnalysisTaskShrinkRequest {
	s.Url = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
