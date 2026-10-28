// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) *ModifyPxfsSecurityIpsResponseBody
	GetAccessDeniedDetail() *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *ModifyPxfsSecurityIpsResponseBodyData) *ModifyPxfsSecurityIpsResponseBody
	GetData() *ModifyPxfsSecurityIpsResponseBodyData
	SetRequestId(v string) *ModifyPxfsSecurityIpsResponseBody
	GetRequestId() *string
}

type ModifyPxfsSecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *ModifyPxfsSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPxfsSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPxfsSecurityIpsResponseBody) GetAccessDeniedDetail() *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyPxfsSecurityIpsResponseBody) GetData() *ModifyPxfsSecurityIpsResponseBodyData {
	return s.Data
}

func (s *ModifyPxfsSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPxfsSecurityIpsResponseBody) SetAccessDeniedDetail(v *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) *ModifyPxfsSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBody) SetData(v *ModifyPxfsSecurityIpsResponseBodyData) *ModifyPxfsSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBody) SetRequestId(v string) *ModifyPxfsSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail struct {
	// The authorization action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The description is the same as above.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyPxfsSecurityIpsResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 123456789
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyPxfsSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyPxfsSecurityIpsResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyPxfsSecurityIpsResponseBodyData) SetTaskId(v int64) *ModifyPxfsSecurityIpsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyPxfsSecurityIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
