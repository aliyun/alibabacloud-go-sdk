// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) *ManageSchedulerxCalendarResponseBody
	GetAccessDeniedDetail() *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail
	SetCode(v int32) *ManageSchedulerxCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *ManageSchedulerxCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManageSchedulerxCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ManageSchedulerxCalendarResponseBody
	GetSuccess() *bool
}

type ManageSchedulerxCalendarResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information. Returned only if an error occurs.
	//
	// example:
	//
	// calendar \\"2025workday\\" does not exist in year 2025
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

func (s ManageSchedulerxCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxCalendarResponseBody) GetAccessDeniedDetail() *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ManageSchedulerxCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManageSchedulerxCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManageSchedulerxCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageSchedulerxCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ManageSchedulerxCalendarResponseBody) SetAccessDeniedDetail(v *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) *ManageSchedulerxCalendarResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ManageSchedulerxCalendarResponseBody) SetCode(v int32) *ManageSchedulerxCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBody) SetMessage(v string) *ManageSchedulerxCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBody) SetRequestId(v string) *ManageSchedulerxCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBody) SetSuccess(v bool) *ManageSchedulerxCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ManageSchedulerxCalendarResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:ManageSchedulerxCalendar
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The principal name.
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

func (s ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ManageSchedulerxCalendarResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
