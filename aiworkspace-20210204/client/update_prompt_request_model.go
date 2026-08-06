// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdatePromptRequest
	GetDescription() *string
	SetFrameworkContent(v string) *UpdatePromptRequest
	GetFrameworkContent() *string
	SetFrameworkType(v string) *UpdatePromptRequest
	GetFrameworkType() *string
	SetWorkspaceId(v string) *UpdatePromptRequest
	GetWorkspaceId() *string
}

type UpdatePromptRequest struct {
	// The description of the prompt.
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
	//     \\"Reflective strips\\": \\"Usually yellow, or yellow-black alternating, attached to permanent protruding obstacles such as wall corners to remind drivers to avoid them.\\",
	//
	//     \\"Ground lock\\": \\"Also known as a parking space lock. When raised, it prevents the parking space from being occupied. When a ground lock is present, you must indicate whether it is in the raised or lowered state.\\",
	//
	//   }"
	//
	// }
	FrameworkContent *string `json:"FrameworkContent,omitempty" xml:"FrameworkContent,omitempty"`
	// The framework type of the prompt template.
	//
	// example:
	//
	// ICIO
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// The workspace ID. You can obtain the ID by calling the [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 302914
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdatePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromptRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePromptRequest) GetFrameworkContent() *string {
	return s.FrameworkContent
}

func (s *UpdatePromptRequest) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *UpdatePromptRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdatePromptRequest) SetDescription(v string) *UpdatePromptRequest {
	s.Description = &v
	return s
}

func (s *UpdatePromptRequest) SetFrameworkContent(v string) *UpdatePromptRequest {
	s.FrameworkContent = &v
	return s
}

func (s *UpdatePromptRequest) SetFrameworkType(v string) *UpdatePromptRequest {
	s.FrameworkType = &v
	return s
}

func (s *UpdatePromptRequest) SetWorkspaceId(v string) *UpdatePromptRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdatePromptRequest) Validate() error {
	return dara.Validate(s)
}
