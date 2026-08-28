// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentIMChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentIMChannelResponseBody
	GetCode() *string
	SetData(v *CreateAgentIMChannelResponseBodyData) *CreateAgentIMChannelResponseBody
	GetData() *CreateAgentIMChannelResponseBodyData
	SetHttpStatusCode(v int32) *CreateAgentIMChannelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAgentIMChannelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentIMChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentIMChannelResponseBody
	GetSuccess() *bool
}

type CreateAgentIMChannelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The IM channel information after creation.
	Data *CreateAgentIMChannelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing result message.
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

func (s CreateAgentIMChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentIMChannelResponseBody) GetData() *CreateAgentIMChannelResponseBodyData {
	return s.Data
}

func (s *CreateAgentIMChannelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAgentIMChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentIMChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentIMChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentIMChannelResponseBody) SetCode(v string) *CreateAgentIMChannelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentIMChannelResponseBody) SetData(v *CreateAgentIMChannelResponseBodyData) *CreateAgentIMChannelResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentIMChannelResponseBody) SetHttpStatusCode(v int32) *CreateAgentIMChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentIMChannelResponseBody) SetMessage(v string) *CreateAgentIMChannelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentIMChannelResponseBody) SetRequestId(v string) *CreateAgentIMChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentIMChannelResponseBody) SetSuccess(v bool) *CreateAgentIMChannelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentIMChannelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentIMChannelResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The channel behavior configuration.
	ChannelConfig *CreateAgentIMChannelResponseBodyDataChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
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
	// The channel credential summary. Only non-sensitive fields and configured secret field names are returned. Secret values are not returned.
	CredentialSummary *CreateAgentIMChannelResponseBodyDataCredentialSummary `json:"credentialSummary,omitempty" xml:"credentialSummary,omitempty" type:"Struct"`
	// Specifies whether to enable the IM channel. Default value: true.
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
	// The ID of the ServiceEndpoint to bind. The endpoint must belong to the specified agent and its current version, be in the ready state, and have a public network address.
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
	// The reason for the current IM channel status.
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

func (s CreateAgentIMChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentIMChannelResponseBodyData) GetChannelConfig() *CreateAgentIMChannelResponseBodyDataChannelConfig {
	return s.ChannelConfig
}

func (s *CreateAgentIMChannelResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateAgentIMChannelResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAgentIMChannelResponseBodyData) GetCredentialSummary() *CreateAgentIMChannelResponseBodyDataCredentialSummary {
	return s.CredentialSummary
}

func (s *CreateAgentIMChannelResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAgentIMChannelResponseBodyData) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *CreateAgentIMChannelResponseBodyData) GetImChannelId() *string {
	return s.ImChannelId
}

func (s *CreateAgentIMChannelResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *CreateAgentIMChannelResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateAgentIMChannelResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateAgentIMChannelResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAgentIMChannelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAgentIMChannelResponseBodyData) SetAgentId(v string) *CreateAgentIMChannelResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetChannelConfig(v *CreateAgentIMChannelResponseBodyDataChannelConfig) *CreateAgentIMChannelResponseBodyData {
	s.ChannelConfig = v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetChannelType(v string) *CreateAgentIMChannelResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetCreateTime(v string) *CreateAgentIMChannelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetCredentialSummary(v *CreateAgentIMChannelResponseBodyDataCredentialSummary) *CreateAgentIMChannelResponseBodyData {
	s.CredentialSummary = v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetEnabled(v bool) *CreateAgentIMChannelResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetEndpointUrl(v string) *CreateAgentIMChannelResponseBodyData {
	s.EndpointUrl = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetImChannelId(v string) *CreateAgentIMChannelResponseBodyData {
	s.ImChannelId = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetServiceEndpointId(v string) *CreateAgentIMChannelResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetStatus(v string) *CreateAgentIMChannelResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetStatusReason(v string) *CreateAgentIMChannelResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetUpdateTime(v string) *CreateAgentIMChannelResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) SetWorkspaceId(v string) *CreateAgentIMChannelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyData) Validate() error {
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

type CreateAgentIMChannelResponseBodyDataChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s CreateAgentIMChannelResponseBodyDataChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelResponseBodyDataChannelConfig) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelResponseBodyDataChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *CreateAgentIMChannelResponseBodyDataChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *CreateAgentIMChannelResponseBodyDataChannelConfig) SetShowThinking(v bool) *CreateAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyDataChannelConfig) SetShowToolCalls(v bool) *CreateAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *CreateAgentIMChannelResponseBodyDataChannelConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAgentIMChannelResponseBodyDataCredentialSummary struct {
	// The list of configured secret field names. Secret values are not included.
	ConfiguredSecretFields []*string `json:"configuredSecretFields,omitempty" xml:"configuredSecretFields,omitempty" type:"Repeated"`
	// The non-sensitive credential fields and their values.
	NonSecretFields map[string]*string `json:"nonSecretFields,omitempty" xml:"nonSecretFields,omitempty"`
}

func (s CreateAgentIMChannelResponseBodyDataCredentialSummary) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelResponseBodyDataCredentialSummary) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelResponseBodyDataCredentialSummary) GetConfiguredSecretFields() []*string {
	return s.ConfiguredSecretFields
}

func (s *CreateAgentIMChannelResponseBodyDataCredentialSummary) GetNonSecretFields() map[string]*string {
	return s.NonSecretFields
}

func (s *CreateAgentIMChannelResponseBodyDataCredentialSummary) SetConfiguredSecretFields(v []*string) *CreateAgentIMChannelResponseBodyDataCredentialSummary {
	s.ConfiguredSecretFields = v
	return s
}

func (s *CreateAgentIMChannelResponseBodyDataCredentialSummary) SetNonSecretFields(v map[string]*string) *CreateAgentIMChannelResponseBodyDataCredentialSummary {
	s.NonSecretFields = v
	return s
}

func (s *CreateAgentIMChannelResponseBodyDataCredentialSummary) Validate() error {
	return dara.Validate(s)
}
