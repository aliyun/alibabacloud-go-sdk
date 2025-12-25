// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateSceneResponseBodyAccessDeniedDetail) *UpdateSceneResponseBody
	GetAccessDeniedDetail() *UpdateSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSceneResponseBody
	GetSuccess() *bool
}

type UpdateSceneResponseBody struct {
	AccessDeniedDetail *UpdateSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s UpdateSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBody) GetAccessDeniedDetail() *UpdateSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSceneResponseBody) SetAccessDeniedDetail(v *UpdateSceneResponseBodyAccessDeniedDetail) *UpdateSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateSceneResponseBody) SetCode(v int64) *UpdateSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSceneResponseBody) SetMessage(v string) *UpdateSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSceneResponseBody) SetRequestId(v string) *UpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSceneResponseBody) SetSuccess(v bool) *UpdateSceneResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
