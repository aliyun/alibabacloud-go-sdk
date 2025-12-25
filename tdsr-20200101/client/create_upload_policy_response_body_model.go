// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateUploadPolicyResponseBodyAccessDeniedDetail) *CreateUploadPolicyResponseBody
	GetAccessDeniedDetail() *CreateUploadPolicyResponseBodyAccessDeniedDetail
	SetCode(v int64) *CreateUploadPolicyResponseBody
	GetCode() *int64
	SetData(v *CreateUploadPolicyResponseBodyData) *CreateUploadPolicyResponseBody
	GetData() *CreateUploadPolicyResponseBodyData
	SetMessage(v string) *CreateUploadPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUploadPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUploadPolicyResponseBody
	GetSuccess() *bool
}

type CreateUploadPolicyResponseBody struct {
	AccessDeniedDetail *CreateUploadPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateUploadPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B28A2ECB-AB29-1E01****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponseBody) GetAccessDeniedDetail() *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateUploadPolicyResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateUploadPolicyResponseBody) GetData() *CreateUploadPolicyResponseBodyData {
	return s.Data
}

func (s *CreateUploadPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUploadPolicyResponseBody) SetAccessDeniedDetail(v *CreateUploadPolicyResponseBodyAccessDeniedDetail) *CreateUploadPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetCode(v int64) *CreateUploadPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetData(v *CreateUploadPolicyResponseBodyData) *CreateUploadPolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetMessage(v string) *CreateUploadPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetRequestId(v string) *CreateUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetSuccess(v bool) *CreateUploadPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUploadPolicyResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateUploadPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateUploadPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateUploadPolicyResponseBodyData struct {
	// accessId
	//
	// example:
	//
	// LTAI5t9k9****
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// ""
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// dir/
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 1658812297
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// oss.aliyun.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMj****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// KdnPJFIG25SM****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateUploadPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateUploadPolicyResponseBodyData) GetCallback() *string {
	return s.Callback
}

func (s *CreateUploadPolicyResponseBodyData) GetDir() *string {
	return s.Dir
}

func (s *CreateUploadPolicyResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *CreateUploadPolicyResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *CreateUploadPolicyResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *CreateUploadPolicyResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *CreateUploadPolicyResponseBodyData) SetAccessId(v string) *CreateUploadPolicyResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetCallback(v string) *CreateUploadPolicyResponseBodyData {
	s.Callback = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetDir(v string) *CreateUploadPolicyResponseBodyData {
	s.Dir = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetExpire(v string) *CreateUploadPolicyResponseBodyData {
	s.Expire = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetHost(v string) *CreateUploadPolicyResponseBodyData {
	s.Host = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetPolicy(v string) *CreateUploadPolicyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetSignature(v string) *CreateUploadPolicyResponseBodyData {
	s.Signature = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
