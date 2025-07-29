// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNetworkContentAuditShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RunNetworkContentAuditShrinkRequest
	GetApiKey() *string
	SetBusinessType(v string) *RunNetworkContentAuditShrinkRequest
	GetBusinessType() *string
	SetContent(v string) *RunNetworkContentAuditShrinkRequest
	GetContent() *string
	SetExtraInfo(v string) *RunNetworkContentAuditShrinkRequest
	GetExtraInfo() *string
	SetModelId(v string) *RunNetworkContentAuditShrinkRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunNetworkContentAuditShrinkRequest
	GetOutputFormat() *string
	SetTagsShrink(v string) *RunNetworkContentAuditShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *RunNetworkContentAuditShrinkRequest
	GetTaskDescription() *string
}

type RunNetworkContentAuditShrinkRequest struct {
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

func (s RunNetworkContentAuditShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunNetworkContentAuditShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *RunNetworkContentAuditShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *RunNetworkContentAuditShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunNetworkContentAuditShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunNetworkContentAuditShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunNetworkContentAuditShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *RunNetworkContentAuditShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunNetworkContentAuditShrinkRequest) SetApiKey(v string) *RunNetworkContentAuditShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetBusinessType(v string) *RunNetworkContentAuditShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetContent(v string) *RunNetworkContentAuditShrinkRequest {
	s.Content = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetExtraInfo(v string) *RunNetworkContentAuditShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetModelId(v string) *RunNetworkContentAuditShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetOutputFormat(v string) *RunNetworkContentAuditShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetTagsShrink(v string) *RunNetworkContentAuditShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) SetTaskDescription(v string) *RunNetworkContentAuditShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunNetworkContentAuditShrinkRequest) Validate() error {
	return dara.Validate(s)
}
