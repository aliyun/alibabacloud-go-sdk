// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteOpenSearchResponseBodyAccessDeniedDetail) *DeleteOpenSearchResponseBody
	GetAccessDeniedDetail() *DeleteOpenSearchResponseBodyAccessDeniedDetail
	SetData(v *DeleteOpenSearchResponseBodyData) *DeleteOpenSearchResponseBody
	GetData() *DeleteOpenSearchResponseBodyData
	SetRequestId(v string) *DeleteOpenSearchResponseBody
	GetRequestId() *string
}

type DeleteOpenSearchResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteOpenSearchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data list.
	Data *DeleteOpenSearchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOpenSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchResponseBody) GetAccessDeniedDetail() *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteOpenSearchResponseBody) GetData() *DeleteOpenSearchResponseBodyData {
	return s.Data
}

func (s *DeleteOpenSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpenSearchResponseBody) SetAccessDeniedDetail(v *DeleteOpenSearchResponseBodyAccessDeniedDetail) *DeleteOpenSearchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteOpenSearchResponseBody) SetData(v *DeleteOpenSearchResponseBodyData) *DeleteOpenSearchResponseBody {
	s.Data = v
	return s
}

func (s *DeleteOpenSearchResponseBody) SetRequestId(v string) *DeleteOpenSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpenSearchResponseBody) Validate() error {
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

type DeleteOpenSearchResponseBodyAccessDeniedDetail struct {
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

func (s DeleteOpenSearchResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteOpenSearchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteOpenSearchResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1111
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteOpenSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteOpenSearchResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteOpenSearchResponseBodyData) SetDBInstanceName(v string) *DeleteOpenSearchResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyData) SetTaskId(v int32) *DeleteOpenSearchResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteOpenSearchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
