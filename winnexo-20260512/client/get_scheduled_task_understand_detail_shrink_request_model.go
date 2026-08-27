// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskUnderstandDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetCollaborationGroupId() *string
	SetDigitalEmployeeNameShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetSegmentsShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetSegmentsShrink() *string
	SetTenantId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetTenantId() *string
	SetUserInput(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetUserInput() *string
}

type GetScheduledTaskUnderstandDetailShrinkRequest struct {
	// The ID of the collaboration group to which the task belongs (such as cg_101). If this parameter is specified, a group space task is created (the caller must be a valid group member). If this parameter is left empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The name of the current effective digital employee. This parameter is empty if not configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The segments.
	SegmentsShrink *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The natural language task description.
	//
	// example:
	//
	// string_value
	UserInput *string `json:"userInput,omitempty" xml:"userInput,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetUserInput() *string {
	return s.UserInput
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetDigitalEmployeeNameShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetSegmentsShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.SegmentsShrink = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetTenantId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetUserInput(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.UserInput = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
