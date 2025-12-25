// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateConnDataResponseBodyAccessDeniedDetail) *UpdateConnDataResponseBody
	GetAccessDeniedDetail() *UpdateConnDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateConnDataResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateConnDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConnDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConnDataResponseBody
	GetSuccess() *bool
}

type UpdateConnDataResponseBody struct {
	AccessDeniedDetail *UpdateConnDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s UpdateConnDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponseBody) GetAccessDeniedDetail() *UpdateConnDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateConnDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateConnDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConnDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConnDataResponseBody) SetAccessDeniedDetail(v *UpdateConnDataResponseBodyAccessDeniedDetail) *UpdateConnDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateConnDataResponseBody) SetCode(v int64) *UpdateConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetMessage(v string) *UpdateConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetRequestId(v string) *UpdateConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetSuccess(v bool) *UpdateConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConnDataResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateConnDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateConnDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateConnDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
