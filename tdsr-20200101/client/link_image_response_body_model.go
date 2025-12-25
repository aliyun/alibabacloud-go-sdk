// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *LinkImageResponseBodyAccessDeniedDetail) *LinkImageResponseBody
	GetAccessDeniedDetail() *LinkImageResponseBodyAccessDeniedDetail
	SetCode(v int64) *LinkImageResponseBody
	GetCode() *int64
	SetMessage(v string) *LinkImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *LinkImageResponseBody
	GetRequestId() *string
	SetResourceId(v string) *LinkImageResponseBody
	GetResourceId() *string
	SetSuccess(v bool) *LinkImageResponseBody
	GetSuccess() *bool
}

type LinkImageResponseBody struct {
	AccessDeniedDetail *LinkImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 234****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LinkImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LinkImageResponseBody) GoString() string {
	return s.String()
}

func (s *LinkImageResponseBody) GetAccessDeniedDetail() *LinkImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *LinkImageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *LinkImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LinkImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LinkImageResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *LinkImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LinkImageResponseBody) SetAccessDeniedDetail(v *LinkImageResponseBodyAccessDeniedDetail) *LinkImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *LinkImageResponseBody) SetCode(v int64) *LinkImageResponseBody {
	s.Code = &v
	return s
}

func (s *LinkImageResponseBody) SetMessage(v string) *LinkImageResponseBody {
	s.Message = &v
	return s
}

func (s *LinkImageResponseBody) SetRequestId(v string) *LinkImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *LinkImageResponseBody) SetResourceId(v string) *LinkImageResponseBody {
	s.ResourceId = &v
	return s
}

func (s *LinkImageResponseBody) SetSuccess(v bool) *LinkImageResponseBody {
	s.Success = &v
	return s
}

func (s *LinkImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LinkImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s LinkImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s LinkImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *LinkImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *LinkImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *LinkImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
