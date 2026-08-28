// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentIMChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentIMChannelsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentIMChannelsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListAgentIMChannelsResponseBodyItems) *ListAgentIMChannelsResponseBody
	GetItems() []*ListAgentIMChannelsResponseBodyItems
	SetMaxResults(v int32) *ListAgentIMChannelsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAgentIMChannelsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAgentIMChannelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentIMChannelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentIMChannelsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgentIMChannelsResponseBody
	GetTotalCount() *int64
}

type ListAgentIMChannelsResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The IM channel list.
	Items []*ListAgentIMChannelsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page for this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page. An empty value indicates that the last page has been reached.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAgentIMChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentIMChannelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentIMChannelsResponseBody) GetItems() []*ListAgentIMChannelsResponseBodyItems {
	return s.Items
}

func (s *ListAgentIMChannelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentIMChannelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentIMChannelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentIMChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentIMChannelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentIMChannelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentIMChannelsResponseBody) SetCode(v string) *ListAgentIMChannelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetHttpStatusCode(v int32) *ListAgentIMChannelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetItems(v []*ListAgentIMChannelsResponseBodyItems) *ListAgentIMChannelsResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetMaxResults(v int32) *ListAgentIMChannelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetMessage(v string) *ListAgentIMChannelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetNextToken(v string) *ListAgentIMChannelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetRequestId(v string) *ListAgentIMChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetSuccess(v bool) *ListAgentIMChannelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) SetTotalCount(v int64) *ListAgentIMChannelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentIMChannelsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentIMChannelsResponseBodyItems struct {
	// The agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The channel behavior configuration.
	ChannelConfig *ListAgentIMChannelsResponseBodyItemsChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
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
	CredentialSummary *ListAgentIMChannelsResponseBodyItemsCredentialSummary `json:"credentialSummary,omitempty" xml:"credentialSummary,omitempty" type:"Struct"`
	// Indicates whether the IM channel is enabled. Default value upon creation: true.
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
	// - CREATING: being created.
	//
	// - READY: ready.
	//
	// - UPDATING: being updated.
	//
	// - FAILED: failed.
	//
	// - DELETING: being deleted.
	//
	// - DELETE_FAILED: deletion failed.
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

func (s ListAgentIMChannelsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentIMChannelsResponseBodyItems) GetChannelConfig() *ListAgentIMChannelsResponseBodyItemsChannelConfig {
	return s.ChannelConfig
}

func (s *ListAgentIMChannelsResponseBodyItems) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListAgentIMChannelsResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAgentIMChannelsResponseBodyItems) GetCredentialSummary() *ListAgentIMChannelsResponseBodyItemsCredentialSummary {
	return s.CredentialSummary
}

func (s *ListAgentIMChannelsResponseBodyItems) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAgentIMChannelsResponseBodyItems) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *ListAgentIMChannelsResponseBodyItems) GetImChannelId() *string {
	return s.ImChannelId
}

func (s *ListAgentIMChannelsResponseBodyItems) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *ListAgentIMChannelsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListAgentIMChannelsResponseBodyItems) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListAgentIMChannelsResponseBodyItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAgentIMChannelsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAgentIMChannelsResponseBodyItems) SetAgentId(v string) *ListAgentIMChannelsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetChannelConfig(v *ListAgentIMChannelsResponseBodyItemsChannelConfig) *ListAgentIMChannelsResponseBodyItems {
	s.ChannelConfig = v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetChannelType(v string) *ListAgentIMChannelsResponseBodyItems {
	s.ChannelType = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetCreateTime(v string) *ListAgentIMChannelsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetCredentialSummary(v *ListAgentIMChannelsResponseBodyItemsCredentialSummary) *ListAgentIMChannelsResponseBodyItems {
	s.CredentialSummary = v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetEnabled(v bool) *ListAgentIMChannelsResponseBodyItems {
	s.Enabled = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetEndpointUrl(v string) *ListAgentIMChannelsResponseBodyItems {
	s.EndpointUrl = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetImChannelId(v string) *ListAgentIMChannelsResponseBodyItems {
	s.ImChannelId = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetServiceEndpointId(v string) *ListAgentIMChannelsResponseBodyItems {
	s.ServiceEndpointId = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetStatus(v string) *ListAgentIMChannelsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetStatusReason(v string) *ListAgentIMChannelsResponseBodyItems {
	s.StatusReason = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetUpdateTime(v string) *ListAgentIMChannelsResponseBodyItems {
	s.UpdateTime = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) SetWorkspaceId(v string) *ListAgentIMChannelsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItems) Validate() error {
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

type ListAgentIMChannelsResponseBodyItemsChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s ListAgentIMChannelsResponseBodyItemsChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsResponseBodyItemsChannelConfig) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsResponseBodyItemsChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *ListAgentIMChannelsResponseBodyItemsChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *ListAgentIMChannelsResponseBodyItemsChannelConfig) SetShowThinking(v bool) *ListAgentIMChannelsResponseBodyItemsChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItemsChannelConfig) SetShowToolCalls(v bool) *ListAgentIMChannelsResponseBodyItemsChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItemsChannelConfig) Validate() error {
	return dara.Validate(s)
}

type ListAgentIMChannelsResponseBodyItemsCredentialSummary struct {
	// The list of configured secret field names. Secret values are not included.
	ConfiguredSecretFields []*string `json:"configuredSecretFields,omitempty" xml:"configuredSecretFields,omitempty" type:"Repeated"`
	// The non-sensitive credential fields and their values.
	NonSecretFields map[string]*string `json:"nonSecretFields,omitempty" xml:"nonSecretFields,omitempty"`
}

func (s ListAgentIMChannelsResponseBodyItemsCredentialSummary) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsResponseBodyItemsCredentialSummary) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsResponseBodyItemsCredentialSummary) GetConfiguredSecretFields() []*string {
	return s.ConfiguredSecretFields
}

func (s *ListAgentIMChannelsResponseBodyItemsCredentialSummary) GetNonSecretFields() map[string]*string {
	return s.NonSecretFields
}

func (s *ListAgentIMChannelsResponseBodyItemsCredentialSummary) SetConfiguredSecretFields(v []*string) *ListAgentIMChannelsResponseBodyItemsCredentialSummary {
	s.ConfiguredSecretFields = v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItemsCredentialSummary) SetNonSecretFields(v map[string]*string) *ListAgentIMChannelsResponseBodyItemsCredentialSummary {
	s.NonSecretFields = v
	return s
}

func (s *ListAgentIMChannelsResponseBodyItemsCredentialSummary) Validate() error {
	return dara.Validate(s)
}
