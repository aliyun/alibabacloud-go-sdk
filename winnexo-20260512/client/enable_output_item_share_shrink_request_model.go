// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOutputItemShareShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppToken(v string) *EnableOutputItemShareShrinkRequest
  GetAppToken() *string 
  SetOpsUserCode(v string) *EnableOutputItemShareShrinkRequest
  GetOpsUserCode() *string 
  SetOpsUserId(v string) *EnableOutputItemShareShrinkRequest
  GetOpsUserId() *string 
  SetOutputItemId(v string) *EnableOutputItemShareShrinkRequest
  GetOutputItemId() *string 
  SetOutputItemName(v string) *EnableOutputItemShareShrinkRequest
  GetOutputItemName() *string 
  SetSessionId(v string) *EnableOutputItemShareShrinkRequest
  GetSessionId() *string 
  SetShareAccessType(v string) *EnableOutputItemShareShrinkRequest
  GetShareAccessType() *string 
  SetShareUserGroupIdsShrink(v string) *EnableOutputItemShareShrinkRequest
  GetShareUserGroupIdsShrink() *string 
  SetShareUserIdsShrink(v string) *EnableOutputItemShareShrinkRequest
  GetShareUserIdsShrink() *string 
  SetTenantId(v string) *EnableOutputItemShareShrinkRequest
  GetTenantId() *string 
}

type EnableOutputItemShareShrinkRequest struct {
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
  // The ID of the output item.
  // 
  // example:
  // 
  // exampleOutputItemId
  OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
  // The name of the output item. This parameter must be used together with sessionId. If multiple output items have the same name, the request is rejected and you are prompted to use outputItemId instead.
  // 
  // example:
  // 
  // Quarterly Review.pptx
  OutputItemName *string `json:"outputItemName,omitempty" xml:"outputItemName,omitempty"`
  // The session ID used for exact filtering. The length is 1 to 128 characters. This parameter applies to common sessions, skill sessions, and the executionId of task executions.
  // 
  // example:
  // 
  // exampleSessionId
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // The sharing scope. Valid values:
  // 
  // - PUBLIC: Accessible to all users.
  // 
  // - SPECIFIED: Accessible only to specified users or user groups.
  // 
  // example:
  // 
  // PUBLIC
  ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
  // The list of user group IDs within the specified scope.
  // 
  // example:
  // 
  // ["group-1"]
  ShareUserGroupIdsShrink *string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty"`
  // The list of platform user IDs within the specified scope.
  // 
  // example:
  // 
  // [10001]
  ShareUserIdsShrink *string `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty"`
  // The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s EnableOutputItemShareShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableOutputItemShareShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableOutputItemShareShrinkRequest) GetAppToken() *string  {
  return s.AppToken
}

func (s *EnableOutputItemShareShrinkRequest) GetOpsUserCode() *string  {
  return s.OpsUserCode
}

func (s *EnableOutputItemShareShrinkRequest) GetOpsUserId() *string  {
  return s.OpsUserId
}

func (s *EnableOutputItemShareShrinkRequest) GetOutputItemId() *string  {
  return s.OutputItemId
}

func (s *EnableOutputItemShareShrinkRequest) GetOutputItemName() *string  {
  return s.OutputItemName
}

func (s *EnableOutputItemShareShrinkRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EnableOutputItemShareShrinkRequest) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableOutputItemShareShrinkRequest) GetShareUserGroupIdsShrink() *string  {
  return s.ShareUserGroupIdsShrink
}

func (s *EnableOutputItemShareShrinkRequest) GetShareUserIdsShrink() *string  {
  return s.ShareUserIdsShrink
}

func (s *EnableOutputItemShareShrinkRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableOutputItemShareShrinkRequest) SetAppToken(v string) *EnableOutputItemShareShrinkRequest {
  s.AppToken = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetOpsUserCode(v string) *EnableOutputItemShareShrinkRequest {
  s.OpsUserCode = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetOpsUserId(v string) *EnableOutputItemShareShrinkRequest {
  s.OpsUserId = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetOutputItemId(v string) *EnableOutputItemShareShrinkRequest {
  s.OutputItemId = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetOutputItemName(v string) *EnableOutputItemShareShrinkRequest {
  s.OutputItemName = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetSessionId(v string) *EnableOutputItemShareShrinkRequest {
  s.SessionId = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetShareAccessType(v string) *EnableOutputItemShareShrinkRequest {
  s.ShareAccessType = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetShareUserGroupIdsShrink(v string) *EnableOutputItemShareShrinkRequest {
  s.ShareUserGroupIdsShrink = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetShareUserIdsShrink(v string) *EnableOutputItemShareShrinkRequest {
  s.ShareUserIdsShrink = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) SetTenantId(v string) *EnableOutputItemShareShrinkRequest {
  s.TenantId = &v
  return s
}

func (s *EnableOutputItemShareShrinkRequest) Validate() error {
  return dara.Validate(s)
}

