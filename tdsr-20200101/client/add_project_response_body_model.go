// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddProjectResponseBodyAccessDeniedDetail) *AddProjectResponseBody
	GetAccessDeniedDetail() *AddProjectResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddProjectResponseBody
	GetCode() *int64
	SetId(v string) *AddProjectResponseBody
	GetId() *string
	SetMessage(v string) *AddProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddProjectResponseBody
	GetSuccess() *bool
}

type AddProjectResponseBody struct {
	AccessDeniedDetail *AddProjectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb130936****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectResponseBody) GetAccessDeniedDetail() *AddProjectResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddProjectResponseBody) GetId() *string {
	return s.Id
}

func (s *AddProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddProjectResponseBody) SetAccessDeniedDetail(v *AddProjectResponseBodyAccessDeniedDetail) *AddProjectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddProjectResponseBody) SetCode(v int64) *AddProjectResponseBody {
	s.Code = &v
	return s
}

func (s *AddProjectResponseBody) SetId(v string) *AddProjectResponseBody {
	s.Id = &v
	return s
}

func (s *AddProjectResponseBody) SetMessage(v string) *AddProjectResponseBody {
	s.Message = &v
	return s
}

func (s *AddProjectResponseBody) SetRequestId(v string) *AddProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectResponseBody) SetSuccess(v bool) *AddProjectResponseBody {
	s.Success = &v
	return s
}

func (s *AddProjectResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddProjectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddProjectResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddProjectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddProjectResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddProjectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddProjectResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
