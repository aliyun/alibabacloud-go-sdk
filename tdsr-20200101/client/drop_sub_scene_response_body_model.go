// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DropSubSceneResponseBodyAccessDeniedDetail) *DropSubSceneResponseBody
	GetAccessDeniedDetail() *DropSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *DropSubSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *DropSubSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DropSubSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DropSubSceneResponseBody
	GetSuccess() *bool
}

type DropSubSceneResponseBody struct {
	AccessDeniedDetail *DropSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s DropSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponseBody) GetAccessDeniedDetail() *DropSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DropSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DropSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DropSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DropSubSceneResponseBody) SetAccessDeniedDetail(v *DropSubSceneResponseBodyAccessDeniedDetail) *DropSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DropSubSceneResponseBody) SetCode(v int64) *DropSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSubSceneResponseBody) SetMessage(v string) *DropSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DropSubSceneResponseBody) SetRequestId(v string) *DropSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSubSceneResponseBody) SetSuccess(v bool) *DropSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DropSubSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DropSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DropSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DropSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DropSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DropSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
