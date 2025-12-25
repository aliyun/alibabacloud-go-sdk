// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CheckUserPropertyResponseBodyAccessDeniedDetail) *CheckUserPropertyResponseBody
	GetAccessDeniedDetail() *CheckUserPropertyResponseBodyAccessDeniedDetail
	SetCode(v int64) *CheckUserPropertyResponseBody
	GetCode() *int64
	SetMatch(v bool) *CheckUserPropertyResponseBody
	GetMatch() *bool
	SetMessage(v string) *CheckUserPropertyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckUserPropertyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckUserPropertyResponseBody
	GetSuccess() *bool
}

type CheckUserPropertyResponseBody struct {
	AccessDeniedDetail *CheckUserPropertyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Match *bool `json:"Match,omitempty" xml:"Match,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb1309366493
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyResponseBody) GetAccessDeniedDetail() *CheckUserPropertyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CheckUserPropertyResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CheckUserPropertyResponseBody) GetMatch() *bool {
	return s.Match
}

func (s *CheckUserPropertyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckUserPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserPropertyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckUserPropertyResponseBody) SetAccessDeniedDetail(v *CheckUserPropertyResponseBodyAccessDeniedDetail) *CheckUserPropertyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CheckUserPropertyResponseBody) SetCode(v int64) *CheckUserPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetMatch(v bool) *CheckUserPropertyResponseBody {
	s.Match = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetMessage(v string) *CheckUserPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetRequestId(v string) *CheckUserPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetSuccess(v bool) *CheckUserPropertyResponseBody {
	s.Success = &v
	return s
}

func (s *CheckUserPropertyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUserPropertyResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CheckUserPropertyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CheckUserPropertyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CheckUserPropertyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CheckUserPropertyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
