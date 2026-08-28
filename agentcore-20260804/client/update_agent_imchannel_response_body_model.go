// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentIMChannelResponseBody
	GetCode() *string
	SetData(v *UpdateAgentIMChannelResponseBodyData) *UpdateAgentIMChannelResponseBody
	GetData() *UpdateAgentIMChannelResponseBodyData
	SetHttpStatusCode(v int32) *UpdateAgentIMChannelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAgentIMChannelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentIMChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentIMChannelResponseBody
	GetSuccess() *bool
}

type UpdateAgentIMChannelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated IM channel information.
	Data *UpdateAgentIMChannelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAgentIMChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentIMChannelResponseBody) GetData() *UpdateAgentIMChannelResponseBodyData {
	return s.Data
}

func (s *UpdateAgentIMChannelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAgentIMChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentIMChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentIMChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentIMChannelResponseBody) SetCode(v string) *UpdateAgentIMChannelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) SetData(v *UpdateAgentIMChannelResponseBodyData) *UpdateAgentIMChannelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) SetHttpStatusCode(v int32) *UpdateAgentIMChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) SetMessage(v string) *UpdateAgentIMChannelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) SetRequestId(v string) *UpdateAgentIMChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) SetSuccess(v bool) *UpdateAgentIMChannelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The channel behavior configuration.
	ChannelConfig *UpdateAgentIMChannelResponseBodyDataChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
	// The IM channel type. Valid values:
	//
	// - DINGTALK: DingTalk.
	//
	// - FEISHU: Lark.
	//
	// - WECOM: WeCom.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The creation time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The channel credential summary. Only non-sensitive fields and the names of configured secret fields are returned. Secret values are not returned.
	CredentialSummary *UpdateAgentIMChannelResponseBodyDataCredentialSummary `json:"credentialSummary,omitempty" xml:"credentialSummary,omitempty" type:"Struct"`
	// Specifies whether to enable the IM channel. Default value: true (when created).
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The public network access URL of the attached ServiceEndpoint.
	//
	// example:
	//
	// https://agent.example.com
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The IM channel ID.
	//
	// example:
	//
	// imc-1
	ImChannelId *string `json:"imChannelId,omitempty" xml:"imChannelId,omitempty"`
	// The ID of the bound ServiceEndpoint. The endpoint must belong to the specified agent and its current version, be in the ready state, and have a public endpoint address.
	//
	// example:
	//
	// se-1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// The IM channel status. Valid values:
	//
	// - CREATING: Being created.
	//
	// - READY: Ready.
	//
	// - UPDATING: Being updated.
	//
	// - FAILED: Failed.
	//
	// - DELETING: Being deleted.
	//
	// - DELETE_FAILED: Deletion failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the current status of the IM channel.
	//
	// example:
	//
	// AppFlow creation failed
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateAgentIMChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateAgentIMChannelResponseBodyData) GetChannelConfig() *UpdateAgentIMChannelResponseBodyDataChannelConfig {
	return s.ChannelConfig
}

func (s *UpdateAgentIMChannelResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *UpdateAgentIMChannelResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateAgentIMChannelResponseBodyData) GetCredentialSummary() *UpdateAgentIMChannelResponseBodyDataCredentialSummary {
	return s.CredentialSummary
}

func (s *UpdateAgentIMChannelResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAgentIMChannelResponseBodyData) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *UpdateAgentIMChannelResponseBodyData) GetImChannelId() *string {
	return s.ImChannelId
}

func (s *UpdateAgentIMChannelResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *UpdateAgentIMChannelResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateAgentIMChannelResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateAgentIMChannelResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateAgentIMChannelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAgentIMChannelResponseBodyData) SetAgentId(v string) *UpdateAgentIMChannelResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetChannelConfig(v *UpdateAgentIMChannelResponseBodyDataChannelConfig) *UpdateAgentIMChannelResponseBodyData {
	s.ChannelConfig = v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetChannelType(v string) *UpdateAgentIMChannelResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetCreateTime(v string) *UpdateAgentIMChannelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetCredentialSummary(v *UpdateAgentIMChannelResponseBodyDataCredentialSummary) *UpdateAgentIMChannelResponseBodyData {
	s.CredentialSummary = v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetEnabled(v bool) *UpdateAgentIMChannelResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetEndpointUrl(v string) *UpdateAgentIMChannelResponseBodyData {
	s.EndpointUrl = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetImChannelId(v string) *UpdateAgentIMChannelResponseBodyData {
	s.ImChannelId = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetServiceEndpointId(v string) *UpdateAgentIMChannelResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetStatus(v string) *UpdateAgentIMChannelResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetStatusReason(v string) *UpdateAgentIMChannelResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetUpdateTime(v string) *UpdateAgentIMChannelResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) SetWorkspaceId(v string) *UpdateAgentIMChannelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyData) Validate() error {
	if s.ChannelConfig != nil {
		if err := s.ChannelConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CredentialSummary != nil {
		if err := s.CredentialSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelResponseBodyDataChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s UpdateAgentIMChannelResponseBodyDataChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelResponseBodyDataChannelConfig) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelResponseBodyDataChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *UpdateAgentIMChannelResponseBodyDataChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *UpdateAgentIMChannelResponseBodyDataChannelConfig) SetShowThinking(v bool) *UpdateAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyDataChannelConfig) SetShowToolCalls(v bool) *UpdateAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyDataChannelConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentIMChannelResponseBodyDataCredentialSummary struct {
	// The list of configured secret field names. Secret values are not included.
	ConfiguredSecretFields []*string `json:"configuredSecretFields,omitempty" xml:"configuredSecretFields,omitempty" type:"Repeated"`
	// The non-sensitive credential fields and their values.
	NonSecretFields map[string]*string `json:"nonSecretFields,omitempty" xml:"nonSecretFields,omitempty"`
}

func (s UpdateAgentIMChannelResponseBodyDataCredentialSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelResponseBodyDataCredentialSummary) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelResponseBodyDataCredentialSummary) GetConfiguredSecretFields() []*string {
	return s.ConfiguredSecretFields
}

func (s *UpdateAgentIMChannelResponseBodyDataCredentialSummary) GetNonSecretFields() map[string]*string {
	return s.NonSecretFields
}

func (s *UpdateAgentIMChannelResponseBodyDataCredentialSummary) SetConfiguredSecretFields(v []*string) *UpdateAgentIMChannelResponseBodyDataCredentialSummary {
	s.ConfiguredSecretFields = v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyDataCredentialSummary) SetNonSecretFields(v map[string]*string) *UpdateAgentIMChannelResponseBodyDataCredentialSummary {
	s.NonSecretFields = v
	return s
}

func (s *UpdateAgentIMChannelResponseBodyDataCredentialSummary) Validate() error {
	return dara.Validate(s)
}
