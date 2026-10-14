// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKBSyncLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UpdateKBSyncLinkRequest
	GetClientId() *string
	SetClientSecret(v string) *UpdateKBSyncLinkRequest
	GetClientSecret() *string
	SetFeishuChatMemberStaleTtlSeconds(v int32) *UpdateKBSyncLinkRequest
	GetFeishuChatMemberStaleTtlSeconds() *int32
	SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *UpdateKBSyncLinkRequest
	GetLinkId() *string
	SetMcpEndpoint(v string) *UpdateKBSyncLinkRequest
	GetMcpEndpoint() *string
	SetRegionId(v string) *UpdateKBSyncLinkRequest
	GetRegionId() *string
	SetSheetMcpEndpoint(v string) *UpdateKBSyncLinkRequest
	GetSheetMcpEndpoint() *string
	SetSyncEnabled(v bool) *UpdateKBSyncLinkRequest
	GetSyncEnabled() *bool
	SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest
	GetSyncIntervalMinutes() *int32
	SetUserAccessToken(v string) *UpdateKBSyncLinkRequest
	GetUserAccessToken() *string
	SetUserId(v string) *UpdateKBSyncLinkRequest
	GetUserId() *string
	SetUserRefreshToken(v string) *UpdateKBSyncLinkRequest
	GetUserRefreshToken() *string
}

type UpdateKBSyncLinkRequest struct {
	// The client ID.
	//
	// example:
	//
	// cli_xxxxxxbe8
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret.
	//
	// example:
	//
	// ******
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The validity period of a successful Lark group member snapshot used for fallback reuse, in seconds. This parameter can be explicitly set only for FEISHU synchronization links. A value of -1 indicates that the snapshot never expires. A value of 0 disables snapshot fallback reuse. A positive integer indicates the maximum age of a snapshot allowed for reuse. If this parameter is omitted, the original value is retained. Explicitly passing -1 restores the default behavior.
	//
	// example:
	//
	// -1
	FeishuChatMemberStaleTtlSeconds *int32 `json:"FeishuChatMemberStaleTtlSeconds,omitempty" xml:"FeishuChatMemberStaleTtlSeconds,omitempty"`
	// The unique identifier of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the synchronization link.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkbl-xxxxx
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// The endpoint of the DingTalk Docs MCP server.
	//
	// example:
	//
	// https://mcp-gw.dingtalk.com/server/1a2b3cxxxxxx789?key=d5e6fxxxxxx876
	McpEndpoint *string `json:"McpEndpoint,omitempty" xml:"McpEndpoint,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint of the DingTalk Sheets MCP server.
	//
	// example:
	//
	// https://mcp-gw.dingtalk.com/server/1a2b3cxxxxxx567?key=d5e6fxxxxxx543
	SheetMcpEndpoint *string `json:"SheetMcpEndpoint,omitempty" xml:"SheetMcpEndpoint,omitempty"`
	// Specifies whether to enable automatic synchronization. This parameter must be modified individually and cannot be changed along with other configurations.
	//
	// example:
	//
	// false
	SyncEnabled *bool `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
	// 20
	//
	// example:
	//
	// 20
	SyncIntervalMinutes *int32 `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
	// The new Lark user access token.
	//
	// example:
	//
	// ******
	UserAccessToken *string `json:"UserAccessToken,omitempty" xml:"UserAccessToken,omitempty"`
	// The ID of the DingTalk user that performs the operation.
	//
	// example:
	//
	// 123***56
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The new Lark user refresh token. Pass this parameter together with UserAccessToken. Passing empty strings for both parameters indicates deletion.
	//
	// example:
	//
	// ******
	UserRefreshToken *string `json:"UserRefreshToken,omitempty" xml:"UserRefreshToken,omitempty"`
}

func (s UpdateKBSyncLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKBSyncLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKBSyncLinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateKBSyncLinkRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateKBSyncLinkRequest) GetFeishuChatMemberStaleTtlSeconds() *int32 {
	return s.FeishuChatMemberStaleTtlSeconds
}

func (s *UpdateKBSyncLinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKBSyncLinkRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *UpdateKBSyncLinkRequest) GetMcpEndpoint() *string {
	return s.McpEndpoint
}

func (s *UpdateKBSyncLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKBSyncLinkRequest) GetSheetMcpEndpoint() *string {
	return s.SheetMcpEndpoint
}

func (s *UpdateKBSyncLinkRequest) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateKBSyncLinkRequest) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *UpdateKBSyncLinkRequest) GetUserAccessToken() *string {
	return s.UserAccessToken
}

func (s *UpdateKBSyncLinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateKBSyncLinkRequest) GetUserRefreshToken() *string {
	return s.UserRefreshToken
}

func (s *UpdateKBSyncLinkRequest) SetClientId(v string) *UpdateKBSyncLinkRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetClientSecret(v string) *UpdateKBSyncLinkRequest {
	s.ClientSecret = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetFeishuChatMemberStaleTtlSeconds(v int32) *UpdateKBSyncLinkRequest {
	s.FeishuChatMemberStaleTtlSeconds = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetLinkId(v string) *UpdateKBSyncLinkRequest {
	s.LinkId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetMcpEndpoint(v string) *UpdateKBSyncLinkRequest {
	s.McpEndpoint = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetRegionId(v string) *UpdateKBSyncLinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSheetMcpEndpoint(v string) *UpdateKBSyncLinkRequest {
	s.SheetMcpEndpoint = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSyncEnabled(v bool) *UpdateKBSyncLinkRequest {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetUserAccessToken(v string) *UpdateKBSyncLinkRequest {
	s.UserAccessToken = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetUserId(v string) *UpdateKBSyncLinkRequest {
	s.UserId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetUserRefreshToken(v string) *UpdateKBSyncLinkRequest {
	s.UserRefreshToken = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) Validate() error {
	return dara.Validate(s)
}
