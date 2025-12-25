// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePackUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetScenePackUrlResponseBodyAccessDeniedDetail) *GetScenePackUrlResponseBody
	GetAccessDeniedDetail() *GetScenePackUrlResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetScenePackUrlResponseBody
	GetCode() *int64
	SetData(v *GetScenePackUrlResponseBodyData) *GetScenePackUrlResponseBody
	GetData() *GetScenePackUrlResponseBodyData
	SetMessage(v string) *GetScenePackUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScenePackUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScenePackUrlResponseBody
	GetSuccess() *bool
}

type GetScenePackUrlResponseBody struct {
	AccessDeniedDetail *GetScenePackUrlResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScenePackUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePackUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScenePackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponseBody) GetAccessDeniedDetail() *GetScenePackUrlResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetScenePackUrlResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetScenePackUrlResponseBody) GetData() *GetScenePackUrlResponseBodyData {
	return s.Data
}

func (s *GetScenePackUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScenePackUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScenePackUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScenePackUrlResponseBody) SetAccessDeniedDetail(v *GetScenePackUrlResponseBodyAccessDeniedDetail) *GetScenePackUrlResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetScenePackUrlResponseBody) SetCode(v int64) *GetScenePackUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetData(v *GetScenePackUrlResponseBodyData) *GetScenePackUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePackUrlResponseBody) SetMessage(v string) *GetScenePackUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetRequestId(v string) *GetScenePackUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetSuccess(v bool) *GetScenePackUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetScenePackUrlResponseBody) Validate() error {
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

type GetScenePackUrlResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetScenePackUrlResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetScenePackUrlResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetScenePackUrlResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetScenePackUrlResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetScenePackUrlResponseBodyData struct {
	// example:
	//
	// 2022-05-17 11:00:17
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetScenePackUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScenePackUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *GetScenePackUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetScenePackUrlResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *GetScenePackUrlResponseBodyData) SetExpire(v string) *GetScenePackUrlResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetScenePackUrlResponseBodyData) SetUrl(v string) *GetScenePackUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetScenePackUrlResponseBodyData) SetValid(v bool) *GetScenePackUrlResponseBodyData {
	s.Valid = &v
	return s
}

func (s *GetScenePackUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
