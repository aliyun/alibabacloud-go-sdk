// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetRectifyImageResponseBodyAccessDeniedDetail) *GetRectifyImageResponseBody
	GetAccessDeniedDetail() *GetRectifyImageResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetRectifyImageResponseBody
	GetCode() *int64
	SetMessage(v string) *GetRectifyImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRectifyImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRectifyImageResponseBody
	GetSuccess() *bool
	SetUrl(v string) *GetRectifyImageResponseBody
	GetUrl() *string
}

type GetRectifyImageResponseBody struct {
	AccessDeniedDetail *GetRectifyImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// https://image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRectifyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponseBody) GetAccessDeniedDetail() *GetRectifyImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetRectifyImageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetRectifyImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRectifyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRectifyImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRectifyImageResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetRectifyImageResponseBody) SetAccessDeniedDetail(v *GetRectifyImageResponseBodyAccessDeniedDetail) *GetRectifyImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetRectifyImageResponseBody) SetCode(v int64) *GetRectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetMessage(v string) *GetRectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetRequestId(v string) *GetRectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetSuccess(v bool) *GetRectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetUrl(v string) *GetRectifyImageResponseBody {
	s.Url = &v
	return s
}

func (s *GetRectifyImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRectifyImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetRectifyImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetRectifyImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetRectifyImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
