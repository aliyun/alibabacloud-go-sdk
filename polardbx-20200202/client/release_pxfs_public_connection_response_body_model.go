// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePxfsPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) *ReleasePxfsPublicConnectionResponseBody
	GetAccessDeniedDetail() *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *ReleasePxfsPublicConnectionResponseBodyData) *ReleasePxfsPublicConnectionResponseBody
	GetData() *ReleasePxfsPublicConnectionResponseBodyData
	SetRequestId(v string) *ReleasePxfsPublicConnectionResponseBody
	GetRequestId() *string
}

type ReleasePxfsPublicConnectionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *ReleasePxfsPublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePxfsPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePxfsPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePxfsPublicConnectionResponseBody) GetAccessDeniedDetail() *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReleasePxfsPublicConnectionResponseBody) GetData() *ReleasePxfsPublicConnectionResponseBodyData {
	return s.Data
}

func (s *ReleasePxfsPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePxfsPublicConnectionResponseBody) SetAccessDeniedDetail(v *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) *ReleasePxfsPublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBody) SetData(v *ReleasePxfsPublicConnectionResponseBodyData) *ReleasePxfsPublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBody) SetRequestId(v string) *ReleasePxfsPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBody) Validate() error {
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

type ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail struct {
	// The authentication action.
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
	// PolicyType
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReleasePxfsPublicConnectionResponseBodyData struct {
	// The instance ID of the PolarDB-X instance.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance ID of the Agent file service.
	//
	// example:
	//
	// pxfs-xxxxxxxx
	PxfsInstanceId *string `json:"PxfsInstanceId,omitempty" xml:"PxfsInstanceId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 123456789
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReleasePxfsPublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleasePxfsPublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) SetDBInstanceName(v string) *ReleasePxfsPublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) SetPxfsInstanceId(v string) *ReleasePxfsPublicConnectionResponseBodyData {
	s.PxfsInstanceId = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) SetTaskId(v int32) *ReleasePxfsPublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ReleasePxfsPublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
