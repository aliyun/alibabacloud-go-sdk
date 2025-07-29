// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTagMiningAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunTagMiningAnalysisRequest
	GetApiKey() *string
	SetBusinessType(v string) *RunTagMiningAnalysisRequest
	GetBusinessType() *string
	SetContent(v string) *RunTagMiningAnalysisRequest
	GetContent() *string
	SetExtraInfo(v string) *RunTagMiningAnalysisRequest
	GetExtraInfo() *string
	SetModelId(v string) *RunTagMiningAnalysisRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunTagMiningAnalysisRequest
	GetOutputFormat() *string
	SetTags(v []*RunTagMiningAnalysisRequestTags) *RunTagMiningAnalysisRequest
	GetTags() []*RunTagMiningAnalysisRequestTags
	SetTaskDescription(v string) *RunTagMiningAnalysisRequest
	GetTaskDescription() *string
}

type RunTagMiningAnalysisRequest struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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
	OutputFormat *string                            `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*RunTagMiningAnalysisRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunTagMiningAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunTagMiningAnalysisRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *RunTagMiningAnalysisRequest) GetContent() *string {
	return s.Content
}

func (s *RunTagMiningAnalysisRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunTagMiningAnalysisRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunTagMiningAnalysisRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunTagMiningAnalysisRequest) GetTags() []*RunTagMiningAnalysisRequestTags {
	return s.Tags
}

func (s *RunTagMiningAnalysisRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunTagMiningAnalysisRequest) SetApiKey(v string) *RunTagMiningAnalysisRequest {
	s.ApiKey = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetBusinessType(v string) *RunTagMiningAnalysisRequest {
	s.BusinessType = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetContent(v string) *RunTagMiningAnalysisRequest {
	s.Content = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetExtraInfo(v string) *RunTagMiningAnalysisRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetModelId(v string) *RunTagMiningAnalysisRequest {
	s.ModelId = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetOutputFormat(v string) *RunTagMiningAnalysisRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetTags(v []*RunTagMiningAnalysisRequestTags) *RunTagMiningAnalysisRequest {
	s.Tags = v
	return s
}

func (s *RunTagMiningAnalysisRequest) SetTaskDescription(v string) *RunTagMiningAnalysisRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunTagMiningAnalysisRequest) Validate() error {
	return dara.Validate(s)
}

type RunTagMiningAnalysisRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunTagMiningAnalysisRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisRequestTags) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *RunTagMiningAnalysisRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *RunTagMiningAnalysisRequestTags) SetTagDefinePrompt(v string) *RunTagMiningAnalysisRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunTagMiningAnalysisRequestTags) SetTagName(v string) *RunTagMiningAnalysisRequestTags {
	s.TagName = &v
	return s
}

func (s *RunTagMiningAnalysisRequestTags) Validate() error {
	return dara.Validate(s)
}
