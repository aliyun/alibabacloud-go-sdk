// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrompt interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Prompt
	GetAccessibility() *string
	SetCreateTime(v string) *Prompt
	GetCreateTime() *string
	SetDescription(v string) *Prompt
	GetDescription() *string
	SetFrameworkContent(v string) *Prompt
	GetFrameworkContent() *string
	SetFrameworkType(v string) *Prompt
	GetFrameworkType() *string
	SetModifyTime(v string) *Prompt
	GetModifyTime() *string
	SetPromptId(v string) *Prompt
	GetPromptId() *string
	SetPromptName(v string) *Prompt
	GetPromptName() *string
}

type Prompt struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// 2024-11-15T07:06:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 这是一个自动驾驶场景的提示词。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {
	//
	//   "context": "你是一个拥有十年驾龄的老司机，请你针对以下图片场景做出你的分析判断。",
	//
	//   "inputData": "{\\n  \\"反光贴条\\": \\" 通常为黄色，或者黄黑间隔，贴在墙角等永久的突出障碍物上，用于提示车主避让。为条带状，不是锥筒，不是地锁，不是水马！\\",\\n  \\"地锁\\": \\"也叫车位锁，升起时可以阻止车位被占用。存在地锁时，务必输出地锁为升起还是降下状态。有升起架子时为升起状态,否则为降下状态。\\"\\n}"
	//
	// }
	FrameworkContent *string `json:"FrameworkContent,omitempty" xml:"FrameworkContent,omitempty"`
	// example:
	//
	// ICIO
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// 2024-12-26T02:17:18Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// pmt-axbxtc****xxx
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// example:
	//
	// 自动驾驶提示词
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
}

func (s Prompt) String() string {
	return dara.Prettify(s)
}

func (s Prompt) GoString() string {
	return s.String()
}

func (s *Prompt) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Prompt) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Prompt) GetDescription() *string {
	return s.Description
}

func (s *Prompt) GetFrameworkContent() *string {
	return s.FrameworkContent
}

func (s *Prompt) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *Prompt) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *Prompt) GetPromptId() *string {
	return s.PromptId
}

func (s *Prompt) GetPromptName() *string {
	return s.PromptName
}

func (s *Prompt) SetAccessibility(v string) *Prompt {
	s.Accessibility = &v
	return s
}

func (s *Prompt) SetCreateTime(v string) *Prompt {
	s.CreateTime = &v
	return s
}

func (s *Prompt) SetDescription(v string) *Prompt {
	s.Description = &v
	return s
}

func (s *Prompt) SetFrameworkContent(v string) *Prompt {
	s.FrameworkContent = &v
	return s
}

func (s *Prompt) SetFrameworkType(v string) *Prompt {
	s.FrameworkType = &v
	return s
}

func (s *Prompt) SetModifyTime(v string) *Prompt {
	s.ModifyTime = &v
	return s
}

func (s *Prompt) SetPromptId(v string) *Prompt {
	s.PromptId = &v
	return s
}

func (s *Prompt) SetPromptName(v string) *Prompt {
	s.PromptName = &v
	return s
}

func (s *Prompt) Validate() error {
	return dara.Validate(s)
}
