// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConversationShareRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSessionId(v string) *EnableConversationShareRequest
  GetSessionId() *string 
  SetShareAccessType(v string) *EnableConversationShareRequest
  GetShareAccessType() *string 
  SetShareUserGroupIds(v []*string) *EnableConversationShareRequest
  GetShareUserGroupIds() []*string 
  SetShareUserIds(v []*int64) *EnableConversationShareRequest
  GetShareUserIds() []*int64 
  SetTenantId(v string) *EnableConversationShareRequest
  GetTenantId() *string 
}

type EnableConversationShareRequest struct {
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
  ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
  // The list of platform user IDs for the specified scope. When shareAccessType is set to SPECIFIED, at least one of this list or the user group list must not be empty.
  // 
  // example:
  // 
  // [10001]
  ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
  // The tenant ID.
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s EnableConversationShareRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableConversationShareRequest) GoString() string {
  return s.String()
}

func (s *EnableConversationShareRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EnableConversationShareRequest) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableConversationShareRequest) GetShareUserGroupIds() []*string  {
  return s.ShareUserGroupIds
}

func (s *EnableConversationShareRequest) GetShareUserIds() []*int64  {
  return s.ShareUserIds
}

func (s *EnableConversationShareRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableConversationShareRequest) SetSessionId(v string) *EnableConversationShareRequest {
  s.SessionId = &v
  return s
}

func (s *EnableConversationShareRequest) SetShareAccessType(v string) *EnableConversationShareRequest {
  s.ShareAccessType = &v
  return s
}

func (s *EnableConversationShareRequest) SetShareUserGroupIds(v []*string) *EnableConversationShareRequest {
  s.ShareUserGroupIds = v
  return s
}

func (s *EnableConversationShareRequest) SetShareUserIds(v []*int64) *EnableConversationShareRequest {
  s.ShareUserIds = v
  return s
}

func (s *EnableConversationShareRequest) SetTenantId(v string) *EnableConversationShareRequest {
  s.TenantId = &v
  return s
}

func (s *EnableConversationShareRequest) Validate() error {
  return dara.Validate(s)
}

