// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeletePxfsResponseBodyAccessDeniedDetail) *DeletePxfsResponseBody
	GetAccessDeniedDetail() *DeletePxfsResponseBodyAccessDeniedDetail
	SetData(v *DeletePxfsResponseBodyData) *DeletePxfsResponseBody
	GetData() *DeletePxfsResponseBodyData
	SetRequestId(v string) *DeletePxfsResponseBody
	GetRequestId() *string
}

type DeletePxfsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeletePxfsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DeletePxfsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePxfsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePxfsResponseBody) GetAccessDeniedDetail() *DeletePxfsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeletePxfsResponseBody) GetData() *DeletePxfsResponseBodyData {
	return s.Data
}

func (s *DeletePxfsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePxfsResponseBody) SetAccessDeniedDetail(v *DeletePxfsResponseBodyAccessDeniedDetail) *DeletePxfsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeletePxfsResponseBody) SetData(v *DeletePxfsResponseBodyData) *DeletePxfsResponseBody {
	s.Data = v
	return s
}

func (s *DeletePxfsResponseBody) SetRequestId(v string) *DeletePxfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePxfsResponseBody) Validate() error {
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

type DeletePxfsResponseBodyAccessDeniedDetail struct {
	// The authorization action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
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
	// The type of the authorization principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
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

func (s DeletePxfsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeletePxfsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeletePxfsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeletePxfsResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 123456789
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeletePxfsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePxfsResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeletePxfsResponseBodyData) SetTaskId(v int64) *DeletePxfsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeletePxfsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
