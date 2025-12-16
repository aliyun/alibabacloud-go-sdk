// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) *CreateSchedulerxCalendarResponseBody
	GetAccessDeniedDetail() *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail
	SetCode(v int32) *CreateSchedulerxCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateSchedulerxCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSchedulerxCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSchedulerxCalendarResponseBody
	GetSuccess() *bool
}

type CreateSchedulerxCalendarResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// Th status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information. Returned only when an error occurs.
	//
	// example:
	//
	// calendar \\"2025workday\\" already exists in year 2025
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchedulerxCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxCalendarResponseBody) GetAccessDeniedDetail() *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateSchedulerxCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSchedulerxCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSchedulerxCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchedulerxCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSchedulerxCalendarResponseBody) SetAccessDeniedDetail(v *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) *CreateSchedulerxCalendarResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateSchedulerxCalendarResponseBody) SetCode(v int32) *CreateSchedulerxCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBody) SetMessage(v string) *CreateSchedulerxCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBody) SetRequestId(v string) *CreateSchedulerxCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBody) SetSuccess(v bool) *CreateSchedulerxCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSchedulerxCalendarResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas: CreateSchedulerxCalendar
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

func (s CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateSchedulerxCalendarResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
