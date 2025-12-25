// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateSubSceneResponseBodyAccessDeniedDetail) *UpdateSubSceneResponseBody
	GetAccessDeniedDetail() *UpdateSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateSubSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateSubSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSubSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSubSceneResponseBody
	GetSuccess() *bool
}

type UpdateSubSceneResponseBody struct {
	AccessDeniedDetail *UpdateSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s UpdateSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponseBody) GetAccessDeniedDetail() *UpdateSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSubSceneResponseBody) SetAccessDeniedDetail(v *UpdateSubSceneResponseBodyAccessDeniedDetail) *UpdateSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateSubSceneResponseBody) SetCode(v int64) *UpdateSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetMessage(v string) *UpdateSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetRequestId(v string) *UpdateSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetSuccess(v bool) *UpdateSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSubSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
