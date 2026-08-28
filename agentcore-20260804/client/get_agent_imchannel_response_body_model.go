// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIMChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentIMChannelResponseBody
	GetCode() *string
	SetData(v *GetAgentIMChannelResponseBodyData) *GetAgentIMChannelResponseBody
	GetData() *GetAgentIMChannelResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentIMChannelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentIMChannelResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentIMChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentIMChannelResponseBody
	GetSuccess() *bool
}

type GetAgentIMChannelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The IM channel details.
	Data *GetAgentIMChannelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAgentIMChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentIMChannelResponseBody) GetData() *GetAgentIMChannelResponseBodyData {
	return s.Data
}

func (s *GetAgentIMChannelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentIMChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentIMChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentIMChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentIMChannelResponseBody) SetCode(v string) *GetAgentIMChannelResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentIMChannelResponseBody) SetData(v *GetAgentIMChannelResponseBodyData) *GetAgentIMChannelResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentIMChannelResponseBody) SetHttpStatusCode(v int32) *GetAgentIMChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentIMChannelResponseBody) SetMessage(v string) *GetAgentIMChannelResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentIMChannelResponseBody) SetRequestId(v string) *GetAgentIMChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentIMChannelResponseBody) SetSuccess(v bool) *GetAgentIMChannelResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentIMChannelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentIMChannelResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The channel behavior configuration.
	ChannelConfig *GetAgentIMChannelResponseBodyDataChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
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
	CredentialSummary *GetAgentIMChannelResponseBodyDataCredentialSummary `json:"credentialSummary,omitempty" xml:"credentialSummary,omitempty" type:"Struct"`
	// Specifies whether the IM channel is enabled. Default value: true.
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
	// The ID of the bound ServiceEndpoint. The endpoint must belong to the specified agent and its current version, be in the ready state, and have a public network address.
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

func (s GetAgentIMChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentIMChannelResponseBodyData) GetChannelConfig() *GetAgentIMChannelResponseBodyDataChannelConfig {
	return s.ChannelConfig
}

func (s *GetAgentIMChannelResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *GetAgentIMChannelResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentIMChannelResponseBodyData) GetCredentialSummary() *GetAgentIMChannelResponseBodyDataCredentialSummary {
	return s.CredentialSummary
}

func (s *GetAgentIMChannelResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAgentIMChannelResponseBodyData) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *GetAgentIMChannelResponseBodyData) GetImChannelId() *string {
	return s.ImChannelId
}

func (s *GetAgentIMChannelResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *GetAgentIMChannelResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAgentIMChannelResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetAgentIMChannelResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAgentIMChannelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAgentIMChannelResponseBodyData) SetAgentId(v string) *GetAgentIMChannelResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetChannelConfig(v *GetAgentIMChannelResponseBodyDataChannelConfig) *GetAgentIMChannelResponseBodyData {
	s.ChannelConfig = v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetChannelType(v string) *GetAgentIMChannelResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetCreateTime(v string) *GetAgentIMChannelResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetCredentialSummary(v *GetAgentIMChannelResponseBodyDataCredentialSummary) *GetAgentIMChannelResponseBodyData {
	s.CredentialSummary = v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetEnabled(v bool) *GetAgentIMChannelResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetEndpointUrl(v string) *GetAgentIMChannelResponseBodyData {
	s.EndpointUrl = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetImChannelId(v string) *GetAgentIMChannelResponseBodyData {
	s.ImChannelId = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetServiceEndpointId(v string) *GetAgentIMChannelResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetStatus(v string) *GetAgentIMChannelResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetStatusReason(v string) *GetAgentIMChannelResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetUpdateTime(v string) *GetAgentIMChannelResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) SetWorkspaceId(v string) *GetAgentIMChannelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyData) Validate() error {
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

type GetAgentIMChannelResponseBodyDataChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s GetAgentIMChannelResponseBodyDataChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelResponseBodyDataChannelConfig) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelResponseBodyDataChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *GetAgentIMChannelResponseBodyDataChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *GetAgentIMChannelResponseBodyDataChannelConfig) SetShowThinking(v bool) *GetAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyDataChannelConfig) SetShowToolCalls(v bool) *GetAgentIMChannelResponseBodyDataChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *GetAgentIMChannelResponseBodyDataChannelConfig) Validate() error {
	return dara.Validate(s)
}

type GetAgentIMChannelResponseBodyDataCredentialSummary struct {
	// The list of configured secret field names. Secret values are not included.
	ConfiguredSecretFields []*string `json:"configuredSecretFields,omitempty" xml:"configuredSecretFields,omitempty" type:"Repeated"`
	// The non-sensitive credential fields and their values.
	NonSecretFields map[string]*string `json:"nonSecretFields,omitempty" xml:"nonSecretFields,omitempty"`
}

func (s GetAgentIMChannelResponseBodyDataCredentialSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelResponseBodyDataCredentialSummary) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelResponseBodyDataCredentialSummary) GetConfiguredSecretFields() []*string {
	return s.ConfiguredSecretFields
}

func (s *GetAgentIMChannelResponseBodyDataCredentialSummary) GetNonSecretFields() map[string]*string {
	return s.NonSecretFields
}

func (s *GetAgentIMChannelResponseBodyDataCredentialSummary) SetConfiguredSecretFields(v []*string) *GetAgentIMChannelResponseBodyDataCredentialSummary {
	s.ConfiguredSecretFields = v
	return s
}

func (s *GetAgentIMChannelResponseBodyDataCredentialSummary) SetNonSecretFields(v map[string]*string) *GetAgentIMChannelResponseBodyDataCredentialSummary {
	s.NonSecretFields = v
	return s
}

func (s *GetAgentIMChannelResponseBodyDataCredentialSummary) Validate() error {
	return dara.Validate(s)
}
