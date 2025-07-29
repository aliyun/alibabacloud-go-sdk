// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTagMiningAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunTagMiningAnalysisShrinkRequest
	GetApiKey() *string
	SetBusinessType(v string) *RunTagMiningAnalysisShrinkRequest
	GetBusinessType() *string
	SetContent(v string) *RunTagMiningAnalysisShrinkRequest
	GetContent() *string
	SetExtraInfo(v string) *RunTagMiningAnalysisShrinkRequest
	GetExtraInfo() *string
	SetModelId(v string) *RunTagMiningAnalysisShrinkRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunTagMiningAnalysisShrinkRequest
	GetOutputFormat() *string
	SetTagsShrink(v string) *RunTagMiningAnalysisShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *RunTagMiningAnalysisShrinkRequest
	GetTaskDescription() *string
}

type RunTagMiningAnalysisShrinkRequest struct {
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
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	TagsShrink   *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunTagMiningAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunTagMiningAnalysisShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *RunTagMiningAnalysisShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *RunTagMiningAnalysisShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunTagMiningAnalysisShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunTagMiningAnalysisShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunTagMiningAnalysisShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *RunTagMiningAnalysisShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunTagMiningAnalysisShrinkRequest) SetApiKey(v string) *RunTagMiningAnalysisShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetBusinessType(v string) *RunTagMiningAnalysisShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetContent(v string) *RunTagMiningAnalysisShrinkRequest {
	s.Content = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetExtraInfo(v string) *RunTagMiningAnalysisShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetModelId(v string) *RunTagMiningAnalysisShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetOutputFormat(v string) *RunTagMiningAnalysisShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetTagsShrink(v string) *RunTagMiningAnalysisShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) SetTaskDescription(v string) *RunTagMiningAnalysisShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunTagMiningAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
