// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOutputFormatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GenerateOutputFormatShrinkRequest
	GetBusinessType() *string
	SetContent(v string) *GenerateOutputFormatShrinkRequest
	GetContent() *string
	SetExtraInfo(v string) *GenerateOutputFormatShrinkRequest
	GetExtraInfo() *string
	SetTagsShrink(v string) *GenerateOutputFormatShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *GenerateOutputFormatShrinkRequest
	GetTaskDescription() *string
}

type GenerateOutputFormatShrinkRequest struct {
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
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 给你一条待分析文本数据，请你按照标签体系来对数据进行打标。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s GenerateOutputFormatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOutputFormatShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateOutputFormatShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GenerateOutputFormatShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *GenerateOutputFormatShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GenerateOutputFormatShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *GenerateOutputFormatShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *GenerateOutputFormatShrinkRequest) SetBusinessType(v string) *GenerateOutputFormatShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetContent(v string) *GenerateOutputFormatShrinkRequest {
	s.Content = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetExtraInfo(v string) *GenerateOutputFormatShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetTagsShrink(v string) *GenerateOutputFormatShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) SetTaskDescription(v string) *GenerateOutputFormatShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *GenerateOutputFormatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
