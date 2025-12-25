// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateProjectResponseBodyAccessDeniedDetail) *UpdateProjectResponseBody
	GetAccessDeniedDetail() *UpdateProjectResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateProjectResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectResponseBody
	GetSuccess() *bool
}

type UpdateProjectResponseBody struct {
	AccessDeniedDetail *UpdateProjectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetAccessDeniedDetail() *UpdateProjectResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectResponseBody) SetAccessDeniedDetail(v *UpdateProjectResponseBodyAccessDeniedDetail) *UpdateProjectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateProjectResponseBody) SetCode(v int64) *UpdateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProjectResponseBody) SetMessage(v string) *UpdateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateProjectResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateProjectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateProjectResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
