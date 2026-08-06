// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreatePromptRequest
	GetAccessibility() *string
	SetDescription(v string) *CreatePromptRequest
	GetDescription() *string
	SetFrameworkContent(v string) *CreatePromptRequest
	GetFrameworkContent() *string
	SetFrameworkType(v string) *CreatePromptRequest
	GetFrameworkType() *string
	SetPromptName(v string) *CreatePromptRequest
	GetPromptName() *string
	SetWorkspaceId(v string) *CreatePromptRequest
	GetWorkspaceId() *string
}

type CreatePromptRequest struct {
	// The workspace visibility. Valid values:
	//
	// - PRIVATE (default): Visible only to you and administrators in this workspace.
	//
	// - PUBLIC: Visible to everyone in this workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The prompt description.
	//
	// example:
	//
	// This is a prompt for information extraction in autonomous driving highway scenarios, focusing on extracting lane and weather information
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The prompt framework content.
	//
	// example:
	//
	// {
	//
	//   "context":"You are an experienced driver with ten years of driving experience. Please analyze and make judgments about the following image scenarios.",
	//
	//   "inputData":"{
	//
	//     \\"Reflective strips\\": \\"Usually yellow or yellow-black alternating, attached to permanent protruding obstacles such as wall corners to remind drivers to avoid them.\\",
	//
	//     \\"Ground lock\\": \\"Also called a parking space lock. When raised, it prevents the parking space from being occupied. When a ground lock is present, you must indicate whether it is in the raised or lowered state.\\",
	//
	//   }"
	//
	// }
	FrameworkContent *string `json:"FrameworkContent,omitempty" xml:"FrameworkContent,omitempty"`
	// The prompt optimization template.
	//
	// example:
	//
	// ICIO
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// The prompt name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Autonomous driving highway information extraction
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
	// The workspace ID. You can obtain the ID by calling the [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 796**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptRequest) GoString() string {
	return s.String()
}

func (s *CreatePromptRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreatePromptRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePromptRequest) GetFrameworkContent() *string {
	return s.FrameworkContent
}

func (s *CreatePromptRequest) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *CreatePromptRequest) GetPromptName() *string {
	return s.PromptName
}

func (s *CreatePromptRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreatePromptRequest) SetAccessibility(v string) *CreatePromptRequest {
	s.Accessibility = &v
	return s
}

func (s *CreatePromptRequest) SetDescription(v string) *CreatePromptRequest {
	s.Description = &v
	return s
}

func (s *CreatePromptRequest) SetFrameworkContent(v string) *CreatePromptRequest {
	s.FrameworkContent = &v
	return s
}

func (s *CreatePromptRequest) SetFrameworkType(v string) *CreatePromptRequest {
	s.FrameworkType = &v
	return s
}

func (s *CreatePromptRequest) SetPromptName(v string) *CreatePromptRequest {
	s.PromptName = &v
	return s
}

func (s *CreatePromptRequest) SetWorkspaceId(v string) *CreatePromptRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePromptRequest) Validate() error {
	return dara.Validate(s)
}
