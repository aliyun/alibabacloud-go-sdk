// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayoutDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateLayoutDataResponseBodyAccessDeniedDetail) *UpdateLayoutDataResponseBody
	GetAccessDeniedDetail() *UpdateLayoutDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateLayoutDataResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateLayoutDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLayoutDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLayoutDataResponseBody
	GetSuccess() *bool
}

type UpdateLayoutDataResponseBody struct {
	AccessDeniedDetail *UpdateLayoutDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
}

func (s UpdateLayoutDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponseBody) GetAccessDeniedDetail() *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateLayoutDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateLayoutDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLayoutDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLayoutDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLayoutDataResponseBody) SetAccessDeniedDetail(v *UpdateLayoutDataResponseBodyAccessDeniedDetail) *UpdateLayoutDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetCode(v int64) *UpdateLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetMessage(v string) *UpdateLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetRequestId(v string) *UpdateLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetSuccess(v bool) *UpdateLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLayoutDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateLayoutDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayoutDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateLayoutDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateLayoutDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
