// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) *DeleteSchedulerxCalendarResponseBody
	GetAccessDeniedDetail() *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail
	SetCode(v int32) *DeleteSchedulerxCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteSchedulerxCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSchedulerxCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSchedulerxCalendarResponseBody
	GetSuccess() *bool
}

type DeleteSchedulerxCalendarResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information. Returned only when an error occurs.
	//
	// example:
	//
	// calendar \\"2025workday\\" does not exist in year 2025
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSchedulerxCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxCalendarResponseBody) GetAccessDeniedDetail() *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteSchedulerxCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSchedulerxCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchedulerxCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchedulerxCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSchedulerxCalendarResponseBody) SetAccessDeniedDetail(v *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) *DeleteSchedulerxCalendarResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBody) SetCode(v int32) *DeleteSchedulerxCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBody) SetMessage(v string) *DeleteSchedulerxCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBody) SetRequestId(v string) *DeleteSchedulerxCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBody) SetSuccess(v bool) *DeleteSchedulerxCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:DeleteSchedulerxCalendar
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The account of the principal.
	//
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFn/cLPZ/3Cz0YxQkZBMjVGLTY0REUtNTlGNS05NzUwLTgyMUE4M0MwMTFDRQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The permission denial type.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteSchedulerxCalendarResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
