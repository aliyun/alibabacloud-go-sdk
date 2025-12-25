// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DropProjectResponseBodyAccessDeniedDetail) *DropProjectResponseBody
	GetAccessDeniedDetail() *DropProjectResponseBodyAccessDeniedDetail
	SetCode(v int64) *DropProjectResponseBody
	GetCode() *int64
	SetMessage(v string) *DropProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DropProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DropProjectResponseBody
	GetSuccess() *bool
}

type DropProjectResponseBody struct {
	AccessDeniedDetail *DropProjectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s DropProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DropProjectResponseBody) GetAccessDeniedDetail() *DropProjectResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DropProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DropProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DropProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DropProjectResponseBody) SetAccessDeniedDetail(v *DropProjectResponseBodyAccessDeniedDetail) *DropProjectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DropProjectResponseBody) SetCode(v int64) *DropProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DropProjectResponseBody) SetMessage(v string) *DropProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DropProjectResponseBody) SetRequestId(v string) *DropProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropProjectResponseBody) SetSuccess(v bool) *DropProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DropProjectResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DropProjectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DropProjectResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DropProjectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DropProjectResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DropProjectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DropProjectResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
