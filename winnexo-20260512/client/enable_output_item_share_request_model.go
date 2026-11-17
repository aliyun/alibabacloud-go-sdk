// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOutputItemShareRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOutputItemId(v string) *EnableOutputItemShareRequest
  GetOutputItemId() *string 
  SetOutputItemName(v string) *EnableOutputItemShareRequest
  GetOutputItemName() *string 
  SetSessionId(v string) *EnableOutputItemShareRequest
  GetSessionId() *string 
  SetShareAccessType(v string) *EnableOutputItemShareRequest
  GetShareAccessType() *string 
  SetShareUserGroupIds(v []*string) *EnableOutputItemShareRequest
  GetShareUserGroupIds() []*string 
  SetShareUserIds(v []*int64) *EnableOutputItemShareRequest
  GetShareUserIds() []*int64 
  SetTenantId(v string) *EnableOutputItemShareRequest
  GetTenantId() *string 
}

type EnableOutputItemShareRequest struct {
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
  ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
  // The list of platform user IDs within the specified scope.
  // 
  // example:
  // 
  // [10001]
  ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
  // The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
  // 
  // example:
  // 
  // 10000
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s EnableOutputItemShareRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableOutputItemShareRequest) GoString() string {
  return s.String()
}

func (s *EnableOutputItemShareRequest) GetOutputItemId() *string  {
  return s.OutputItemId
}

func (s *EnableOutputItemShareRequest) GetOutputItemName() *string  {
  return s.OutputItemName
}

func (s *EnableOutputItemShareRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EnableOutputItemShareRequest) GetShareAccessType() *string  {
  return s.ShareAccessType
}

func (s *EnableOutputItemShareRequest) GetShareUserGroupIds() []*string  {
  return s.ShareUserGroupIds
}

func (s *EnableOutputItemShareRequest) GetShareUserIds() []*int64  {
  return s.ShareUserIds
}

func (s *EnableOutputItemShareRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableOutputItemShareRequest) SetOutputItemId(v string) *EnableOutputItemShareRequest {
  s.OutputItemId = &v
  return s
}

func (s *EnableOutputItemShareRequest) SetOutputItemName(v string) *EnableOutputItemShareRequest {
  s.OutputItemName = &v
  return s
}

func (s *EnableOutputItemShareRequest) SetSessionId(v string) *EnableOutputItemShareRequest {
  s.SessionId = &v
  return s
}

func (s *EnableOutputItemShareRequest) SetShareAccessType(v string) *EnableOutputItemShareRequest {
  s.ShareAccessType = &v
  return s
}

func (s *EnableOutputItemShareRequest) SetShareUserGroupIds(v []*string) *EnableOutputItemShareRequest {
  s.ShareUserGroupIds = v
  return s
}

func (s *EnableOutputItemShareRequest) SetShareUserIds(v []*int64) *EnableOutputItemShareRequest {
  s.ShareUserIds = v
  return s
}

func (s *EnableOutputItemShareRequest) SetTenantId(v string) *EnableOutputItemShareRequest {
  s.TenantId = &v
  return s
}

func (s *EnableOutputItemShareRequest) Validate() error {
  return dara.Validate(s)
}

