// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalType(v int32) *UpdateApprovalProcessShrinkRequest
	GetApprovalType() *int32
	SetDescription(v string) *UpdateApprovalProcessShrinkRequest
	GetDescription() *string
	SetEventLabel(v string) *UpdateApprovalProcessShrinkRequest
	GetEventLabel() *string
	SetExternalConfig(v string) *UpdateApprovalProcessShrinkRequest
	GetExternalConfig() *string
	SetMatchSchemaConfigsShrink(v string) *UpdateApprovalProcessShrinkRequest
	GetMatchSchemaConfigsShrink() *string
	SetMatchSchemasShrink(v string) *UpdateApprovalProcessShrinkRequest
	GetMatchSchemasShrink() *string
	SetProcessId(v string) *UpdateApprovalProcessShrinkRequest
	GetProcessId() *string
	SetProcessName(v string) *UpdateApprovalProcessShrinkRequest
	GetProcessName() *string
	SetProcessNodes(v [][]*string) *UpdateApprovalProcessShrinkRequest
	GetProcessNodes() [][]*string
}

type UpdateApprovalProcessShrinkRequest struct {
	// The approval type. Valid values: 0 (built-in approval), 1 (DingTalk approval), 2 (WeCom approval), and 3 (Lark approval). Default value: 0.
	//
	// example:
	//
	// 1
	ApprovalType *int32 `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The description of the approval process. The description must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// Office endpoint software installation and uninstallation approval process
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event callback identifier. When the approval type is 1 (DingTalk approval) or 3 (Lark approval), this parameter is used to generate the event callback URL for the external platform. You can pass a complete callback URL, and the server retains only the part after the callback path prefix.
	//
	// example:
	//
	// https://callback.example.com/sase/api/v1/approval/dingtalk/callback/9m4e2mr0ui3e8a21****
	EventLabel *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	// The integration configuration for the external approval platform. This parameter is required when the approval type is 1 (DingTalk approval) or 3 (Lark approval). It includes integration parameters such as appKey and appSecret, and must be passed as a JSON string.
	//
	// example:
	//
	// {"appKey":"dingn1a2b3c4d5e6f****","appSecret":"f9c8e7d6a5b4c3d2e1f0a9b8c7d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0****","eventVerifyToken":"7a8b9c0d1e2f3a4b5c6d7e8f9a0b****","eventAesKey":"1a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t****"}
	ExternalConfig *string `json:"ExternalConfig,omitempty" xml:"ExternalConfig,omitempty"`
	// The external approval template configurations. Configure the approval template ID, external approval process ID, and field mappings based on the policy type. This parameter is used only when ApprovalType is set to 1 (DingTalk approval), 2 (WeCom approval), or 3 (Lark approval). If this parameter is not passed, the existing template associations of the approval process are cleared. The approval templates associated with business policies cannot be cleared.
	MatchSchemaConfigsShrink *string `json:"MatchSchemaConfigs,omitempty" xml:"MatchSchemaConfigs,omitempty"`
	// The matched approval templates.
	MatchSchemasShrink *string `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty"`
	// The ID of the approval process.
	//
	// This parameter is required.
	//
	// example:
	//
	// approval-process-f16bf74b2b29****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the approval process. The name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Endpoint software installation approval process
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The list of approval nodes. You can define up to 5 approval nodes.
	ProcessNodes [][]*string `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
}

func (s UpdateApprovalProcessShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessShrinkRequest) GetApprovalType() *int32 {
	return s.ApprovalType
}

func (s *UpdateApprovalProcessShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApprovalProcessShrinkRequest) GetEventLabel() *string {
	return s.EventLabel
}

func (s *UpdateApprovalProcessShrinkRequest) GetExternalConfig() *string {
	return s.ExternalConfig
}

func (s *UpdateApprovalProcessShrinkRequest) GetMatchSchemaConfigsShrink() *string {
	return s.MatchSchemaConfigsShrink
}

func (s *UpdateApprovalProcessShrinkRequest) GetMatchSchemasShrink() *string {
	return s.MatchSchemasShrink
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *UpdateApprovalProcessShrinkRequest) GetProcessNodes() [][]*string {
	return s.ProcessNodes
}

func (s *UpdateApprovalProcessShrinkRequest) SetApprovalType(v int32) *UpdateApprovalProcessShrinkRequest {
	s.ApprovalType = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetDescription(v string) *UpdateApprovalProcessShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetEventLabel(v string) *UpdateApprovalProcessShrinkRequest {
	s.EventLabel = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetExternalConfig(v string) *UpdateApprovalProcessShrinkRequest {
	s.ExternalConfig = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetMatchSchemaConfigsShrink(v string) *UpdateApprovalProcessShrinkRequest {
	s.MatchSchemaConfigsShrink = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetMatchSchemasShrink(v string) *UpdateApprovalProcessShrinkRequest {
	s.MatchSchemasShrink = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessId(v string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessId = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessName(v string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessName = &v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) SetProcessNodes(v [][]*string) *UpdateApprovalProcessShrinkRequest {
	s.ProcessNodes = v
	return s
}

func (s *UpdateApprovalProcessShrinkRequest) Validate() error {
	return dara.Validate(s)
}
