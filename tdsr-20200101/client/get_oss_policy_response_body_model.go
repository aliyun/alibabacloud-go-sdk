// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetOssPolicyResponseBodyAccessDeniedDetail) *GetOssPolicyResponseBody
	GetAccessDeniedDetail() *GetOssPolicyResponseBodyAccessDeniedDetail
	SetAccessId(v string) *GetOssPolicyResponseBody
	GetAccessId() *string
	SetCallback(v string) *GetOssPolicyResponseBody
	GetCallback() *string
	SetCode(v int64) *GetOssPolicyResponseBody
	GetCode() *int64
	SetDir(v string) *GetOssPolicyResponseBody
	GetDir() *string
	SetExpire(v string) *GetOssPolicyResponseBody
	GetExpire() *string
	SetHost(v string) *GetOssPolicyResponseBody
	GetHost() *string
	SetMessage(v string) *GetOssPolicyResponseBody
	GetMessage() *string
	SetPolicy(v string) *GetOssPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GetOssPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *GetOssPolicyResponseBody
	GetSignature() *string
	SetSuccess(v bool) *GetOssPolicyResponseBody
	GetSuccess() *bool
}

type GetOssPolicyResponseBody struct {
	AccessDeniedDetail *GetOssPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// abc
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// ""
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123/
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 60
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// oss.aliyun.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// def
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ghi
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOssPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBody) GetAccessDeniedDetail() *GetOssPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetOssPolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *GetOssPolicyResponseBody) GetCallback() *string {
	return s.Callback
}

func (s *GetOssPolicyResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetOssPolicyResponseBody) GetDir() *string {
	return s.Dir
}

func (s *GetOssPolicyResponseBody) GetExpire() *string {
	return s.Expire
}

func (s *GetOssPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetOssPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOssPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetOssPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GetOssPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOssPolicyResponseBody) SetAccessDeniedDetail(v *GetOssPolicyResponseBodyAccessDeniedDetail) *GetOssPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetOssPolicyResponseBody) SetAccessId(v string) *GetOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCallback(v string) *GetOssPolicyResponseBody {
	s.Callback = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCode(v int64) *GetOssPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetDir(v string) *GetOssPolicyResponseBody {
	s.Dir = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetExpire(v string) *GetOssPolicyResponseBody {
	s.Expire = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetHost(v string) *GetOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetMessage(v string) *GetOssPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetPolicy(v string) *GetOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetRequestId(v string) *GetOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSignature(v string) *GetOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSuccess(v bool) *GetOssPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetOssPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssPolicyResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetOssPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOssPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetOssPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetOssPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
