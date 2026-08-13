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
	// 所属协作群组 ID（如 cg_101）；群任务理解时传入（调用者需为有效群成员），候选技能额外并入群绑定技能
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 数字员工名称列表，用于过滤可用技能；必传（传空列表表示仅用租户 global 技能）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	SegmentsShrink            *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 自然语言任务描述
	//
	// This parameter is required.
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
