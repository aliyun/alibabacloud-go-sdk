// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddSubSceneResponseBodyAccessDeniedDetail) *AddSubSceneResponseBody
	GetAccessDeniedDetail() *AddSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddSubSceneResponseBody
	GetCode() *int64
	SetId(v string) *AddSubSceneResponseBody
	GetId() *string
	SetMessage(v string) *AddSubSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSubSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSubSceneResponseBody
	GetSuccess() *bool
}

type AddSubSceneResponseBody struct {
	AccessDeniedDetail *AddSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s AddSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponseBody) GetAccessDeniedDetail() *AddSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddSubSceneResponseBody) GetId() *string {
	return s.Id
}

func (s *AddSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSubSceneResponseBody) SetAccessDeniedDetail(v *AddSubSceneResponseBodyAccessDeniedDetail) *AddSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddSubSceneResponseBody) SetCode(v int64) *AddSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSubSceneResponseBody) SetId(v string) *AddSubSceneResponseBody {
	s.Id = &v
	return s
}

func (s *AddSubSceneResponseBody) SetMessage(v string) *AddSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSubSceneResponseBody) SetRequestId(v string) *AddSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSubSceneResponseBody) SetSuccess(v bool) *AddSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *AddSubSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
