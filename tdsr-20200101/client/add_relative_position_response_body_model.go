// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRelativePositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddRelativePositionResponseBodyAccessDeniedDetail) *AddRelativePositionResponseBody
	GetAccessDeniedDetail() *AddRelativePositionResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddRelativePositionResponseBody
	GetCode() *int64
	SetMessage(v string) *AddRelativePositionResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRelativePositionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRelativePositionResponseBody
	GetSuccess() *bool
}

type AddRelativePositionResponseBody struct {
	AccessDeniedDetail *AddRelativePositionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s AddRelativePositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRelativePositionResponseBody) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponseBody) GetAccessDeniedDetail() *AddRelativePositionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddRelativePositionResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddRelativePositionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRelativePositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRelativePositionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRelativePositionResponseBody) SetAccessDeniedDetail(v *AddRelativePositionResponseBodyAccessDeniedDetail) *AddRelativePositionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddRelativePositionResponseBody) SetCode(v int64) *AddRelativePositionResponseBody {
	s.Code = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetMessage(v string) *AddRelativePositionResponseBody {
	s.Message = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetRequestId(v string) *AddRelativePositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetSuccess(v bool) *AddRelativePositionResponseBody {
	s.Success = &v
	return s
}

func (s *AddRelativePositionResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRelativePositionResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddRelativePositionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddRelativePositionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddRelativePositionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddRelativePositionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
