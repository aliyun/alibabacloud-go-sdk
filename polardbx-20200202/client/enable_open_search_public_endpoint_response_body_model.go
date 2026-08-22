// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOpenSearchPublicEndpointResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) *EnableOpenSearchPublicEndpointResponseBody
  GetAccessDeniedDetail() *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail 
  SetHost(v string) *EnableOpenSearchPublicEndpointResponseBody
  GetHost() *string 
  SetMessage(v string) *EnableOpenSearchPublicEndpointResponseBody
  GetMessage() *string 
  SetPort(v int32) *EnableOpenSearchPublicEndpointResponseBody
  GetPort() *int32 
  SetRequestId(v string) *EnableOpenSearchPublicEndpointResponseBody
  GetRequestId() *string 
}

type EnableOpenSearchPublicEndpointResponseBody struct {
  // The details of the access denial.
  AccessDeniedDetail *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
  // The host required for upload.
  // 
  // example:
  // 
  // https://sddp-datamask.oss-cn-zhangjiakou.aliyuncs.com
  Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
  // The additional information returned by the operation. "success" is returned if the operation is successful. Otherwise, the corresponding error code is returned.
  // 
  // example:
  // 
  // *****
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The public endpoint port.
  // 
  // example:
  // 
  // 3306
  Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 1A586DCB-39A6-4050-81CC-C7BD4CCDB49F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableOpenSearchPublicEndpointResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableOpenSearchPublicEndpointResponseBody) GoString() string {
  return s.String()
}

func (s *EnableOpenSearchPublicEndpointResponseBody) GetAccessDeniedDetail() *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail  {
  return s.AccessDeniedDetail
}

func (s *EnableOpenSearchPublicEndpointResponseBody) GetHost() *string  {
  return s.Host
}

func (s *EnableOpenSearchPublicEndpointResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableOpenSearchPublicEndpointResponseBody) GetPort() *int32  {
  return s.Port
}

func (s *EnableOpenSearchPublicEndpointResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableOpenSearchPublicEndpointResponseBody) SetAccessDeniedDetail(v *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) *EnableOpenSearchPublicEndpointResponseBody {
  s.AccessDeniedDetail = v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBody) SetHost(v string) *EnableOpenSearchPublicEndpointResponseBody {
  s.Host = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBody) SetMessage(v string) *EnableOpenSearchPublicEndpointResponseBody {
  s.Message = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBody) SetPort(v int32) *EnableOpenSearchPublicEndpointResponseBody {
  s.Port = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBody) SetRequestId(v string) *EnableOpenSearchPublicEndpointResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBody) Validate() error {
  if s.AccessDeniedDetail != nil {
    if err := s.AccessDeniedDetail.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail struct {
  // The authentication action.
  // 
  // example:
  // 
  // xxx
  AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
  // The display name of the authentication principal.
  // 
  // example:
  // 
  // xxx
  AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
  // The description is the same as above.
  // 
  // example:
  // 
  // 111
  AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
  // The type of the authentication principal.
  // 
  // example:
  // 
  // 222
  AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
  // The encoded diagnostic message.
  // 
  // example:
  // 
  // AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
  EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
  // The type of the permission denial.
  // 
  // example:
  // 
  // ImplicitDeny
  NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
  // The policy type.
  // 
  // example:
  // 
  // PRIORITY
  PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) String() string {
  return dara.Prettify(s)
}

func (s EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GoString() string {
  return s.String()
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthAction() *string  {
  return s.AuthAction
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string  {
  return s.AuthPrincipalDisplayName
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string  {
  return s.AuthPrincipalOwnerId
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string  {
  return s.AuthPrincipalType
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string  {
  return s.EncodedDiagnosticMessage
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetNoPermissionType() *string  {
  return s.NoPermissionType
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetPolicyType() *string  {
  return s.PolicyType
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthAction(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.AuthAction = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.AuthPrincipalDisplayName = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.AuthPrincipalOwnerId = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.AuthPrincipalType = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.EncodedDiagnosticMessage = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.NoPermissionType = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetPolicyType(v string) *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
  s.PolicyType = &v
  return s
}

func (s *EnableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) Validate() error {
  return dara.Validate(s)
}

