// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOutputFormatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GenerateOutputFormatRequest
	GetBusinessType() *string
	SetContent(v string) *GenerateOutputFormatRequest
	GetContent() *string
	SetExtraInfo(v string) *GenerateOutputFormatRequest
	GetExtraInfo() *string
	SetTags(v []*GenerateOutputFormatRequestTags) *GenerateOutputFormatRequest
	GetTags() []*GenerateOutputFormatRequestTags
	SetTaskDescription(v string) *GenerateOutputFormatRequest
	GetTaskDescription() *string
}

type GenerateOutputFormatRequest struct {
	// example:
	//
	// clueMining
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// example:
	//
	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 额外信息
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// This parameter is required.
	Tags []*GenerateOutputFormatRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s GenerateOutputFormatRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatRequest) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GenerateOutputFormatRequest) GetContent() *string {
	return s.Content
}

func (s *GenerateOutputFormatRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GenerateOutputFormatRequest) GetTags() []*GenerateOutputFormatRequestTags {
	return s.Tags
}

func (s *GenerateOutputFormatRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *GenerateOutputFormatRequest) SetBusinessType(v string) *GenerateOutputFormatRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetContent(v string) *GenerateOutputFormatRequest {
	s.Content = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetExtraInfo(v string) *GenerateOutputFormatRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateOutputFormatRequest) SetTags(v []*GenerateOutputFormatRequestTags) *GenerateOutputFormatRequest {
	s.Tags = v
	return s
}

func (s *GenerateOutputFormatRequest) SetTaskDescription(v string) *GenerateOutputFormatRequest {
	s.TaskDescription = &v
	return s
}

func (s *GenerateOutputFormatRequest) Validate() error {
	return dara.Validate(s)
}

type GenerateOutputFormatRequestTags struct {
	// example:
	//
	// xxxx
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// example:
	//
	// xxxx
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s GenerateOutputFormatRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatRequestTags) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *GenerateOutputFormatRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *GenerateOutputFormatRequestTags) SetTagDefinePrompt(v string) *GenerateOutputFormatRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *GenerateOutputFormatRequestTags) SetTagName(v string) *GenerateOutputFormatRequestTags {
	s.TagName = &v
	return s
}

func (s *GenerateOutputFormatRequestTags) Validate() error {
	return dara.Validate(s)
}
