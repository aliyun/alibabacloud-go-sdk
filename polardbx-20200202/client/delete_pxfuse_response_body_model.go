// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfuseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeletePxfuseResponseBodyAccessDeniedDetail) *DeletePxfuseResponseBody
	GetAccessDeniedDetail() *DeletePxfuseResponseBodyAccessDeniedDetail
	SetData(v *DeletePxfuseResponseBodyData) *DeletePxfuseResponseBody
	GetData() *DeletePxfuseResponseBodyData
	SetRequestId(v string) *DeletePxfuseResponseBody
	GetRequestId() *string
}

type DeletePxfuseResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeletePxfuseResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *DeletePxfuseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePxfuseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfuseResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePxfuseResponseBody) GetAccessDeniedDetail() *DeletePxfuseResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeletePxfuseResponseBody) GetData() *DeletePxfuseResponseBodyData {
	return s.Data
}

func (s *DeletePxfuseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePxfuseResponseBody) SetAccessDeniedDetail(v *DeletePxfuseResponseBodyAccessDeniedDetail) *DeletePxfuseResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeletePxfuseResponseBody) SetData(v *DeletePxfuseResponseBodyData) *DeletePxfuseResponseBody {
	s.Data = v
	return s
}

func (s *DeletePxfuseResponseBody) SetRequestId(v string) *DeletePxfuseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePxfuseResponseBody) Validate() error {
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

type DeletePxfuseResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
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
	// The type of the permission denial.
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

func (s DeletePxfuseResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfuseResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeletePxfuseResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeletePxfuseResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeletePxfuseResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeletePxfuseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfuseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePxfuseResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeletePxfuseResponseBodyData) SetTaskId(v int32) *DeletePxfuseResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeletePxfuseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
