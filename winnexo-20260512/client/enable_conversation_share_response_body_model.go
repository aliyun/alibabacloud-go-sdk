// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConversationShareResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableConversationShareResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableConversationShareResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableConversationShareResponseBody
  GetRequestId() *string 
  SetSessionId(v string) *EnableConversationShareResponseBody
  GetSessionId() *string 
  SetShareAccessType(v string) *EnableConversationShareResponseBody
  GetShareAccessType() *string 
  SetShareEnabled(v bool) *EnableConversationShareResponseBody
  GetShareEnabled() *bool 
  SetShareEnabledAt(v string) *EnableConversationShareResponseBody
  GetShareEnabledAt() *string 
  SetSharePath(v string) *EnableConversationShareResponseBody
  GetSharePath() *string 
  SetShareToken(v string) *EnableConversationShareResponseBody
  GetShareToken() *string 
  SetShareUrl(v string) *EnableConversationShareResponseBody
  GetShareUrl() *string 
  SetShareUserGroupIds(v []*string) *EnableConversationShareResponseBody
  GetShareUserGroupIds() []*string 
  SetShareUserIds(v []*int64) *EnableConversationShareResponseBody
  GetShareUserIds() []*int64 
}

type EnableConversationShareResponseBody struct {
  // The result code. A value of SUCCESS indicates success. For failures, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The status code description.
  // 
  // example:
  // 
  // Invalid request parameters
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 019FF406-1B10-0065-A97D-2D1920C2A03D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // The ID of the source single-chat session.
  // 
  // example:
  // 
  // exampleSessionId
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // The current sharing scope. Valid values: PUBLIC and SPECIFIED.
  // 
  // example:
  // 
  // PUBLIC
  ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
  // Indicates whether sharing is enabled.
  // 
  // example:
  // 
  // true
  ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
  // The time when sharing was most recently enabled, in ISO 8601 format.
  // 
  // example:
  // 
  // 2026-11-16T13:00:00+00:00
  ShareEnabledAt *string `json:"shareEnabledAt,omitempty" xml:"shareEnabledAt,omitempty"`
  // The relative path for in-site sharing. This value is empty when sharing is disabled.
  // 
  // example:
  // 
  // /10001/share/chat/token
  SharePath *string `json:"sharePath,omitempty" xml:"sharePath,omitempty"`
  // The sharing token. This token is present when sharing is enabled and can be used to access the public sharing preview API operation.
  // 
  // example:
  // 
  // exampleShareToken
  ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
  // The absolute URL of the current sharing link. This value is empty when sharing is disabled.
  // 
  // example:
  // 
  // https://example.com/10001/share/chat/token
  ShareUrl *string `json:"shareUrl,omitempty" xml:"shareUrl,omitempty"`
  // The list of currently specified user group IDs.
  // 
  // example:
  // 
  // []
  ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
  // The list of currently specified user IDs.
  // 
  // example:
  // 
  // []
  ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
}

func (s EnableConversationShareResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableConversationShareResponseBody) GoString() string {
  return s.String()
}

func (s *EnableConversationShareResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableConversationShareResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableConversationShareResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableConversationShareResponseBody) GetSessionId() *string  {
  return s.SessionId
}

func (s *EnableConversationShareResponseBody) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableConversationShareResponseBody) GetShareEnabled() *bool  {
  return s.ShareEnabled
}

func (s *EnableConversationShareResponseBody) GetShareEnabledAt() *string  {
  return s.ShareEnabledAt
}

func (s *EnableConversationShareResponseBody) GetSharePath() *string  {
  return s.SharePath
}

func (s *EnableConversationShareResponseBody) GetShareToken() *string  {
  return s.ShareToken
}

func (s *EnableConversationShareResponseBody) GetShareUrl() *string  {
  return s.ShareUrl
}

func (s *EnableConversationShareResponseBody) GetShareUserGroupIds() []*string  {
  return s.ShareUserGroupIds
}

func (s *EnableConversationShareResponseBody) GetShareUserIds() []*int64  {
  return s.ShareUserIds
}

func (s *EnableConversationShareResponseBody) SetCode(v string) *EnableConversationShareResponseBody {
  s.Code = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetMessage(v string) *EnableConversationShareResponseBody {
  s.Message = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetRequestId(v string) *EnableConversationShareResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetSessionId(v string) *EnableConversationShareResponseBody {
  s.SessionId = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareAccessType(v string) *EnableConversationShareResponseBody {
  s.ShareAccessType = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareEnabled(v bool) *EnableConversationShareResponseBody {
  s.ShareEnabled = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareEnabledAt(v string) *EnableConversationShareResponseBody {
  s.ShareEnabledAt = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetSharePath(v string) *EnableConversationShareResponseBody {
  s.SharePath = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareToken(v string) *EnableConversationShareResponseBody {
  s.ShareToken = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareUrl(v string) *EnableConversationShareResponseBody {
  s.ShareUrl = &v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareUserGroupIds(v []*string) *EnableConversationShareResponseBody {
  s.ShareUserGroupIds = v
  return s
}

func (s *EnableConversationShareResponseBody) SetShareUserIds(v []*int64) *EnableConversationShareResponseBody {
  s.ShareUserIds = v
  return s
}

func (s *EnableConversationShareResponseBody) Validate() error {
  return dara.Validate(s)
}

