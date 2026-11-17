// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOutputItemShareResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableOutputItemShareResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableOutputItemShareResponseBody
  GetMessage() *string 
  SetOutputItemId(v string) *EnableOutputItemShareResponseBody
  GetOutputItemId() *string 
  SetRequestId(v string) *EnableOutputItemShareResponseBody
  GetRequestId() *string 
  SetShareAccessType(v string) *EnableOutputItemShareResponseBody
  GetShareAccessType() *string 
  SetShareEnabled(v bool) *EnableOutputItemShareResponseBody
  GetShareEnabled() *bool 
  SetShareEnabledAt(v string) *EnableOutputItemShareResponseBody
  GetShareEnabledAt() *string 
  SetSharePath(v string) *EnableOutputItemShareResponseBody
  GetSharePath() *string 
  SetShareToken(v string) *EnableOutputItemShareResponseBody
  GetShareToken() *string 
  SetShareUrl(v string) *EnableOutputItemShareResponseBody
  GetShareUrl() *string 
  SetShareUserGroupIds(v []*string) *EnableOutputItemShareResponseBody
  GetShareUserGroupIds() []*string 
  SetShareUserIds(v []*int64) *EnableOutputItemShareResponseBody
  GetShareUserIds() []*int64 
}

type EnableOutputItemShareResponseBody struct {
  // The response code.
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The description of the status code.
  // 
  // example:
  // 
  // Invalid request parameters
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The ID of the output item.
  // 
  // example:
  // 
  // exampleOutputItemId
  OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 019FF406-1B10-0065-A97D-2D1920C2A03D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
  // The time when sharing was last enabled.
  // 
  // example:
  // 
  // 2026-11-16T13:00:00+00:00
  ShareEnabledAt *string `json:"shareEnabledAt,omitempty" xml:"shareEnabledAt,omitempty"`
  // The relative path for in-site sharing. This parameter is empty if sharing is disabled.
  // 
  // example:
  // 
  // /10001/share/output/token
  SharePath *string `json:"sharePath,omitempty" xml:"sharePath,omitempty"`
  // The share token. This token is returned when sharing is enabled and can be used to access the public share preview API operation.
  // 
  // example:
  // 
  // exampleShareToken
  ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
  // The absolute URL of the current share. This parameter is empty if sharing is disabled.
  // 
  // example:
  // 
  // https://example.com/10001/share/output/token
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

func (s EnableOutputItemShareResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableOutputItemShareResponseBody) GoString() string {
  return s.String()
}

func (s *EnableOutputItemShareResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableOutputItemShareResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableOutputItemShareResponseBody) GetOutputItemId() *string  {
  return s.OutputItemId
}

func (s *EnableOutputItemShareResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableOutputItemShareResponseBody) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableOutputItemShareResponseBody) GetShareEnabled() *bool  {
  return s.ShareEnabled
}

func (s *EnableOutputItemShareResponseBody) GetShareEnabledAt() *string  {
  return s.ShareEnabledAt
}

func (s *EnableOutputItemShareResponseBody) GetSharePath() *string  {
  return s.SharePath
}

func (s *EnableOutputItemShareResponseBody) GetShareToken() *string  {
  return s.ShareToken
}

func (s *EnableOutputItemShareResponseBody) GetShareUrl() *string  {
  return s.ShareUrl
}

func (s *EnableOutputItemShareResponseBody) GetShareUserGroupIds() []*string  {
  return s.ShareUserGroupIds
}

func (s *EnableOutputItemShareResponseBody) GetShareUserIds() []*int64  {
  return s.ShareUserIds
}

func (s *EnableOutputItemShareResponseBody) SetCode(v string) *EnableOutputItemShareResponseBody {
  s.Code = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetMessage(v string) *EnableOutputItemShareResponseBody {
  s.Message = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetOutputItemId(v string) *EnableOutputItemShareResponseBody {
  s.OutputItemId = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetRequestId(v string) *EnableOutputItemShareResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareAccessType(v string) *EnableOutputItemShareResponseBody {
  s.ShareAccessType = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareEnabled(v bool) *EnableOutputItemShareResponseBody {
  s.ShareEnabled = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareEnabledAt(v string) *EnableOutputItemShareResponseBody {
  s.ShareEnabledAt = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetSharePath(v string) *EnableOutputItemShareResponseBody {
  s.SharePath = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareToken(v string) *EnableOutputItemShareResponseBody {
  s.ShareToken = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareUrl(v string) *EnableOutputItemShareResponseBody {
  s.ShareUrl = &v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareUserGroupIds(v []*string) *EnableOutputItemShareResponseBody {
  s.ShareUserGroupIds = v
  return s
}

func (s *EnableOutputItemShareResponseBody) SetShareUserIds(v []*int64) *EnableOutputItemShareResponseBody {
  s.ShareUserIds = v
  return s
}

func (s *EnableOutputItemShareResponseBody) Validate() error {
  return dara.Validate(s)
}

