// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConversationShareShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppToken(v string) *EnableConversationShareShrinkRequest
  GetAppToken() *string 
  SetOpsUserCode(v string) *EnableConversationShareShrinkRequest
  GetOpsUserCode() *string 
  SetOpsUserId(v string) *EnableConversationShareShrinkRequest
  GetOpsUserId() *string 
  SetSessionId(v string) *EnableConversationShareShrinkRequest
  GetSessionId() *string 
  SetShareAccessType(v string) *EnableConversationShareShrinkRequest
  GetShareAccessType() *string 
  SetShareUserGroupIdsShrink(v string) *EnableConversationShareShrinkRequest
  GetShareUserGroupIdsShrink() *string 
  SetShareUserIdsShrink(v string) *EnableConversationShareShrinkRequest
  GetShareUserIdsShrink() *string 
  SetTenantId(v string) *EnableConversationShareShrinkRequest
  GetTenantId() *string 
}

type EnableConversationShareShrinkRequest struct {
  // 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
  // 
  // example:
  // 
  // wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
  // The unique identifier of the function session.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // exampleSessionId
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // The sharing scope. Valid values:
  // 
  // - PUBLIC: accessible to all users
  // 
  // - SPECIFIED: accessible to specified users or user groups
  // 
  // example:
  // 
  // PUBLIC
  ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
  // The list of user group IDs for the specified scope. When shareAccessType is set to SPECIFIED, at least one of this list or the user list must not be empty.
  // 
  // example:
  // 
  // ["group-1"]
  ShareUserGroupIdsShrink *string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty"`
  // The list of platform user IDs for the specified scope. When shareAccessType is set to SPECIFIED, at least one of this list or the user group list must not be empty.
  // 
  // example:
  // 
  // [10001]
  ShareUserIdsShrink *string `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty"`
  // The tenant ID.
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s EnableConversationShareShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableConversationShareShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableConversationShareShrinkRequest) GetAppToken() *string  {
  return s.AppToken
}

func (s *EnableConversationShareShrinkRequest) GetOpsUserCode() *string  {
  return s.OpsUserCode
}

func (s *EnableConversationShareShrinkRequest) GetOpsUserId() *string  {
  return s.OpsUserId
}

func (s *EnableConversationShareShrinkRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EnableConversationShareShrinkRequest) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableConversationShareShrinkRequest) GetShareUserGroupIdsShrink() *string  {
  return s.ShareUserGroupIdsShrink
}

func (s *EnableConversationShareShrinkRequest) GetShareUserIdsShrink() *string  {
  return s.ShareUserIdsShrink
}

func (s *EnableConversationShareShrinkRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableConversationShareShrinkRequest) SetAppToken(v string) *EnableConversationShareShrinkRequest {
  s.AppToken = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetOpsUserCode(v string) *EnableConversationShareShrinkRequest {
  s.OpsUserCode = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetOpsUserId(v string) *EnableConversationShareShrinkRequest {
  s.OpsUserId = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetSessionId(v string) *EnableConversationShareShrinkRequest {
  s.SessionId = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetShareAccessType(v string) *EnableConversationShareShrinkRequest {
  s.ShareAccessType = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetShareUserGroupIdsShrink(v string) *EnableConversationShareShrinkRequest {
  s.ShareUserGroupIdsShrink = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetShareUserIdsShrink(v string) *EnableConversationShareShrinkRequest {
  s.ShareUserIdsShrink = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) SetTenantId(v string) *EnableConversationShareShrinkRequest {
  s.TenantId = &v
  return s
}

func (s *EnableConversationShareShrinkRequest) Validate() error {
  return dara.Validate(s)
}

