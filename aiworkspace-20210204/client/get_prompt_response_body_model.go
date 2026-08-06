// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetPromptResponseBody
	GetAccessibility() *string
	SetCreateTime(v string) *GetPromptResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetPromptResponseBody
	GetDescription() *string
	SetFrameworkContent(v string) *GetPromptResponseBody
	GetFrameworkContent() *string
	SetFrameworkType(v string) *GetPromptResponseBody
	GetFrameworkType() *string
	SetModifyTime(v string) *GetPromptResponseBody
	GetModifyTime() *string
	SetPromptName(v string) *GetPromptResponseBody
	GetPromptName() *string
	SetRequestId(v string) *GetPromptResponseBody
	GetRequestId() *string
}

type GetPromptResponseBody struct {
	// The access type. Valid values:
	//
	// - PUBLIC: All members in the current workspace can access the prompt.
	//
	// - PRIVATE: Only the creator can access the prompt.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-16T01:44:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The prompt description.
	//
	// example:
	//
	// This is an information extraction prompt for autonomous driving highway scenarios, focusing on extracting lane and weather information
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The prompt content.
	//
	// example:
	//
	// {
	//
	//   "context":"You are an experienced driver with ten years of driving experience. Please analyze and make judgments about the following image scenarios.",
	//
	//   "inputData":"{
	//
	//     \\"Reflective strips\\": \\"Usually yellow or yellow-black alternating, attached to permanent protruding obstacles such as wall corners to remind drivers to avoid them. They are strip-shaped, not cones, not ground locks, not water barriers!\\",
	//
	//     \\"Ground lock\\": \\"Also called a parking space lock, it can prevent a parking space from being occupied when raised. When a ground lock is present, you must indicate whether it is in the raised or lowered state. It is in the raised state when there is a raised frame, otherwise it is in the lowered state.\\",
	//
	//   }"
	//
	// }
	FrameworkContent *string `json:"FrameworkContent,omitempty" xml:"FrameworkContent,omitempty"`
	// The prompt template framework type.
	//
	// example:
	//
	// ICIO
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-08-27T02:01:10Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The prompt name.
	//
	// example:
	//
	// Autonomous driving prompt
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPromptResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetPromptResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPromptResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetPromptResponseBody) GetFrameworkContent() *string {
	return s.FrameworkContent
}

func (s *GetPromptResponseBody) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *GetPromptResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetPromptResponseBody) GetPromptName() *string {
	return s.PromptName
}

func (s *GetPromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPromptResponseBody) SetAccessibility(v string) *GetPromptResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetPromptResponseBody) SetCreateTime(v string) *GetPromptResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPromptResponseBody) SetDescription(v string) *GetPromptResponseBody {
	s.Description = &v
	return s
}

func (s *GetPromptResponseBody) SetFrameworkContent(v string) *GetPromptResponseBody {
	s.FrameworkContent = &v
	return s
}

func (s *GetPromptResponseBody) SetFrameworkType(v string) *GetPromptResponseBody {
	s.FrameworkType = &v
	return s
}

func (s *GetPromptResponseBody) SetModifyTime(v string) *GetPromptResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetPromptResponseBody) SetPromptName(v string) *GetPromptResponseBody {
	s.PromptName = &v
	return s
}

func (s *GetPromptResponseBody) SetRequestId(v string) *GetPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptResponseBody) Validate() error {
	return dara.Validate(s)
}
