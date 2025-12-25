// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneSeqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) *UpdateSubSceneSeqResponseBody
	GetAccessDeniedDetail() *UpdateSubSceneSeqResponseBodyAccessDeniedDetail
	SetCode(v int64) *UpdateSubSceneSeqResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateSubSceneSeqResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSubSceneSeqResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSubSceneSeqResponseBody
	GetSuccess() *bool
}

type UpdateSubSceneSeqResponseBody struct {
	AccessDeniedDetail *UpdateSubSceneSeqResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 344794c32937474a9c59eb13093****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubSceneSeqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneSeqResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqResponseBody) GetAccessDeniedDetail() *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateSubSceneSeqResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateSubSceneSeqResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubSceneSeqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubSceneSeqResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSubSceneSeqResponseBody) SetAccessDeniedDetail(v *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) *UpdateSubSceneSeqResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetCode(v int64) *UpdateSubSceneSeqResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetMessage(v string) *UpdateSubSceneSeqResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetRequestId(v string) *UpdateSubSceneSeqResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetSuccess(v bool) *UpdateSubSceneSeqResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSubSceneSeqResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateSubSceneSeqResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateSubSceneSeqResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
