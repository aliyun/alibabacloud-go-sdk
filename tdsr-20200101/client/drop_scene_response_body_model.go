// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DropSceneResponseBodyAccessDeniedDetail) *DropSceneResponseBody
	GetAccessDeniedDetail() *DropSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *DropSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *DropSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DropSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DropSceneResponseBody
	GetSuccess() *bool
}

type DropSceneResponseBody struct {
	AccessDeniedDetail *DropSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s DropSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSceneResponseBody) GetAccessDeniedDetail() *DropSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DropSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DropSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DropSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DropSceneResponseBody) SetAccessDeniedDetail(v *DropSceneResponseBodyAccessDeniedDetail) *DropSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DropSceneResponseBody) SetCode(v int64) *DropSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSceneResponseBody) SetMessage(v string) *DropSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DropSceneResponseBody) SetRequestId(v string) *DropSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSceneResponseBody) SetSuccess(v bool) *DropSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DropSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DropSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DropSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DropSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DropSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DropSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DropSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
