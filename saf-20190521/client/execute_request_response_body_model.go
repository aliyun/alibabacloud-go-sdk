// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v *ExecuteRequestResponseBodyAccessDeniedDetail) *ExecuteRequestResponseBody
  GetAccessDeniedDetail() *ExecuteRequestResponseBodyAccessDeniedDetail 
  SetCode(v int32) *ExecuteRequestResponseBody
  GetCode() *int32 
  SetData(v map[string]interface{}) *ExecuteRequestResponseBody
  GetData() map[string]interface{} 
  SetMessage(v string) *ExecuteRequestResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteRequestResponseBody
  GetRequestId() *string 
}

type ExecuteRequestResponseBody struct {
  // The diagnostic details for access denial.
  AccessDeniedDetail *ExecuteRequestResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
  // The status code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The response data, which includes fields such as tags, score, and extend.
  // 
  // example:
  // 
  // {"score": 80.4, "tags": "ra0601,rnk0002,rn0312"}
  Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
  // The result message.
  // 
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // A5F41DC6-2070-5CE7-8B2D-7A23D882932D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteRequestResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteRequestResponseBody) GetAccessDeniedDetail() *ExecuteRequestResponseBodyAccessDeniedDetail  {
  return s.AccessDeniedDetail
}

func (s *ExecuteRequestResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteRequestResponseBody) GetData() map[string]interface{}  {
  return s.Data
}

func (s *ExecuteRequestResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteRequestResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteRequestResponseBody) SetAccessDeniedDetail(v *ExecuteRequestResponseBodyAccessDeniedDetail) *ExecuteRequestResponseBody {
  s.AccessDeniedDetail = v
  return s
}

func (s *ExecuteRequestResponseBody) SetCode(v int32) *ExecuteRequestResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteRequestResponseBody) SetData(v map[string]interface{}) *ExecuteRequestResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteRequestResponseBody) SetMessage(v string) *ExecuteRequestResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteRequestResponseBody) SetRequestId(v string) *ExecuteRequestResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteRequestResponseBody) Validate() error {
  if s.AccessDeniedDetail != nil {
    if err := s.AccessDeniedDetail.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteRequestResponseBodyAccessDeniedDetail struct {
  // The specific action that the caller attempted to perform, such as s3:PutObject or ec2:DescribeInstances.
  // 
  // example:
  // 
  // yundun-saf:Execute
  AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
  // The human-readable display name of the request principal, such as the username or role name.
  // 
  // example:
  // 
  // 29918048924096****
  AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
  // The UID of the Alibaba Cloud account to which the request principal belongs.
  // 
  // example:
  // 
  // 178059440630****
  AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
  // The identity type of the request principal. Valid values:
  // 
  // - SubUser: RAM user created by an Alibaba Cloud account.
  // 
  // - AssumedRoleUser: Temporary identity obtained by assuming a role.
  // 
  // example:
  // 
  // SubUser
  AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
  // The Base64-encoded permission diagnostic log used for troubleshooting and automated analysis.
  // 
  // example:
  // 
  // AQIBIAAAA...r+ow==
  EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
  // The specific denial type for the access denied scenario. Valid values:
  // 
  // - ImplicitDeny: Implicit deny due to the lack of explicit authorization.
  // 
  // - ExplicitDeny: Explicit deny where the operation is explicitly denied in the policy.
  // 
  // example:
  // 
  // ImplicitDeny
  NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
  // The scope and type of the policy in the access denied scenario. Valid values:
  // 
  // - AccountLevelIdentityBasedPolicy: Account-level identity-based policy that applies to the entire Alibaba Cloud account.
  // 
  // - ResourceGroupLevelIdentityBasedPolicy: Resource group-level identity-based policy that applies to a specific resource group.
  // 
  // - SessionPolicy: Session policy that applies to a temporary session.
  // 
  // - ControlPolicy: Control policy used for core permission control.
  // 
  // - AssumeRolePolicy: Trust policy that defines which entities can assume the role.
  // 
  // example:
  // 
  // AccountLevelIdentityBasedPolicy
  PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ExecuteRequestResponseBodyAccessDeniedDetail) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestResponseBodyAccessDeniedDetail) GoString() string {
  return s.String()
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetAuthAction() *string  {
  return s.AuthAction
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string  {
  return s.AuthPrincipalDisplayName
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string  {
  return s.AuthPrincipalOwnerId
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string  {
  return s.AuthPrincipalType
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string  {
  return s.EncodedDiagnosticMessage
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetNoPermissionType() *string  {
  return s.NoPermissionType
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) GetPolicyType() *string  {
  return s.PolicyType
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.AuthAction = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.AuthPrincipalDisplayName = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.AuthPrincipalOwnerId = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.AuthPrincipalType = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.EncodedDiagnosticMessage = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.NoPermissionType = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ExecuteRequestResponseBodyAccessDeniedDetail {
  s.PolicyType = &v
  return s
}

func (s *ExecuteRequestResponseBodyAccessDeniedDetail) Validate() error {
  return dara.Validate(s)
}

