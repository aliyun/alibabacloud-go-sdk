// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNetworkContentAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunNetworkContentAuditRequest
	GetApiKey() *string
	SetBusinessType(v string) *RunNetworkContentAuditRequest
	GetBusinessType() *string
	SetContent(v string) *RunNetworkContentAuditRequest
	GetContent() *string
	SetExtraInfo(v string) *RunNetworkContentAuditRequest
	GetExtraInfo() *string
	SetModelId(v string) *RunNetworkContentAuditRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunNetworkContentAuditRequest
	GetOutputFormat() *string
	SetTags(v []*RunNetworkContentAuditRequestTags) *RunNetworkContentAuditRequest
	GetTags() []*RunNetworkContentAuditRequestTags
	SetTaskDescription(v string) *RunNetworkContentAuditRequest
	GetTaskDescription() *string
}

type RunNetworkContentAuditRequest struct {
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
	OutputFormat *string                              `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	Tags         []*RunNetworkContentAuditRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunNetworkContentAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditRequest) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunNetworkContentAuditRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *RunNetworkContentAuditRequest) GetContent() *string {
	return s.Content
}

func (s *RunNetworkContentAuditRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunNetworkContentAuditRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunNetworkContentAuditRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunNetworkContentAuditRequest) GetTags() []*RunNetworkContentAuditRequestTags {
	return s.Tags
}

func (s *RunNetworkContentAuditRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunNetworkContentAuditRequest) SetApiKey(v string) *RunNetworkContentAuditRequest {
	s.ApiKey = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetBusinessType(v string) *RunNetworkContentAuditRequest {
	s.BusinessType = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetContent(v string) *RunNetworkContentAuditRequest {
	s.Content = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetExtraInfo(v string) *RunNetworkContentAuditRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetModelId(v string) *RunNetworkContentAuditRequest {
	s.ModelId = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetOutputFormat(v string) *RunNetworkContentAuditRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunNetworkContentAuditRequest) SetTags(v []*RunNetworkContentAuditRequestTags) *RunNetworkContentAuditRequest {
	s.Tags = v
	return s
}

func (s *RunNetworkContentAuditRequest) SetTaskDescription(v string) *RunNetworkContentAuditRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunNetworkContentAuditRequest) Validate() error {
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

type RunNetworkContentAuditRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunNetworkContentAuditRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditRequestTags) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *RunNetworkContentAuditRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *RunNetworkContentAuditRequestTags) SetTagDefinePrompt(v string) *RunNetworkContentAuditRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunNetworkContentAuditRequestTags) SetTagName(v string) *RunNetworkContentAuditRequestTags {
	s.TagName = &v
	return s
}

func (s *RunNetworkContentAuditRequestTags) Validate() error {
	return dara.Validate(s)
}
