// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskUnderstandDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetCollaborationGroupId() *string
	SetDigitalEmployeeNameShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetOpsUserCode(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetOpsUserId() *string
	SetSegmentsShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetSegmentsShrink() *string
	SetTenantId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetTenantId() *string
	SetUserInput(v string) *GetScheduledTaskUnderstandDetailShrinkRequest
	GetUserInput() *string
}

type GetScheduledTaskUnderstandDetailShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration group to which the task belongs, such as cg_101. If this parameter is specified, a group space task is created and the caller must be a valid group member. If this parameter is left empty, a personal task is created.
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
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetAppToken(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetDigitalEmployeeNameShrink(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetOpsUserCode(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailShrinkRequest) SetOpsUserId(v string) *GetScheduledTaskUnderstandDetailShrinkRequest {
	s.OpsUserId = &v
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
