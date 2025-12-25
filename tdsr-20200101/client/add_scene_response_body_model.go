// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddSceneResponseBodyAccessDeniedDetail) *AddSceneResponseBody
	GetAccessDeniedDetail() *AddSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddSceneResponseBody
	GetCode() *int64
	SetId(v string) *AddSceneResponseBody
	GetId() *string
	SetMessage(v string) *AddSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSceneResponseBody
	GetSuccess() *bool
}

type AddSceneResponseBody struct {
	AccessDeniedDetail *AddSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2345****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s AddSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSceneResponseBody) GetAccessDeniedDetail() *AddSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddSceneResponseBody) GetId() *string {
	return s.Id
}

func (s *AddSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSceneResponseBody) SetAccessDeniedDetail(v *AddSceneResponseBodyAccessDeniedDetail) *AddSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddSceneResponseBody) SetCode(v int64) *AddSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSceneResponseBody) SetId(v string) *AddSceneResponseBody {
	s.Id = &v
	return s
}

func (s *AddSceneResponseBody) SetMessage(v string) *AddSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSceneResponseBody) SetRequestId(v string) *AddSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSceneResponseBody) SetSuccess(v bool) *AddSceneResponseBody {
	s.Success = &v
	return s
}

func (s *AddSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
