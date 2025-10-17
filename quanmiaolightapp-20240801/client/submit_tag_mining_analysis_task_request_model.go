// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTagMiningAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *SubmitTagMiningAnalysisTaskRequest
	GetApiKey() *string
	SetBusinessType(v string) *SubmitTagMiningAnalysisTaskRequest
	GetBusinessType() *string
	SetContents(v []*string) *SubmitTagMiningAnalysisTaskRequest
	GetContents() []*string
	SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskRequest
	GetExtraInfo() *string
	SetModelId(v string) *SubmitTagMiningAnalysisTaskRequest
	GetModelId() *string
	SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskRequest
	GetOutputFormat() *string
	SetTags(v []*SubmitTagMiningAnalysisTaskRequestTags) *SubmitTagMiningAnalysisTaskRequest
	GetTags() []*SubmitTagMiningAnalysisTaskRequestTags
	SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskRequest
	GetTaskDescription() *string
	SetUrl(v string) *SubmitTagMiningAnalysisTaskRequest
	GetUrl() *string
}

type SubmitTagMiningAnalysisTaskRequest struct {
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// clueMining
	BusinessType *string   `json:"businessType,omitempty" xml:"businessType,omitempty"`
	Contents     []*string `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
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
	OutputFormat *string                                   `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*SubmitTagMiningAnalysisTaskRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetContents() []*string {
	return s.Contents
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetTags() []*SubmitTagMiningAnalysisTaskRequestTags {
	return s.Tags
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *SubmitTagMiningAnalysisTaskRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetApiKey(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetBusinessType(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetContents(v []*string) *SubmitTagMiningAnalysisTaskRequest {
	s.Contents = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetExtraInfo(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.ExtraInfo = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetModelId(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetOutputFormat(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.OutputFormat = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetTags(v []*SubmitTagMiningAnalysisTaskRequestTags) *SubmitTagMiningAnalysisTaskRequest {
	s.Tags = v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetTaskDescription(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) SetUrl(v string) *SubmitTagMiningAnalysisTaskRequest {
	s.Url = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitTagMiningAnalysisTaskRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s SubmitTagMiningAnalysisTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SubmitTagMiningAnalysisTaskRequestTags) GoString() string {
	return s.String()
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) SetTagDefinePrompt(v string) *SubmitTagMiningAnalysisTaskRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) SetTagName(v string) *SubmitTagMiningAnalysisTaskRequestTags {
	s.TagName = &v
	return s
}

func (s *SubmitTagMiningAnalysisTaskRequestTags) Validate() error {
	return dara.Validate(s)
}
