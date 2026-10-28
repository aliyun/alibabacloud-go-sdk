// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPxfsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RestartPxfsResponseBodyAccessDeniedDetail) *RestartPxfsResponseBody
	GetAccessDeniedDetail() *RestartPxfsResponseBodyAccessDeniedDetail
	SetData(v *RestartPxfsResponseBodyData) *RestartPxfsResponseBody
	GetData() *RestartPxfsResponseBodyData
	SetRequestId(v string) *RestartPxfsResponseBody
	GetRequestId() *string
}

type RestartPxfsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *RestartPxfsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *RestartPxfsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartPxfsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartPxfsResponseBody) GoString() string {
	return s.String()
}

func (s *RestartPxfsResponseBody) GetAccessDeniedDetail() *RestartPxfsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RestartPxfsResponseBody) GetData() *RestartPxfsResponseBodyData {
	return s.Data
}

func (s *RestartPxfsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartPxfsResponseBody) SetAccessDeniedDetail(v *RestartPxfsResponseBodyAccessDeniedDetail) *RestartPxfsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RestartPxfsResponseBody) SetData(v *RestartPxfsResponseBodyData) *RestartPxfsResponseBody {
	s.Data = v
	return s
}

func (s *RestartPxfsResponseBody) SetRequestId(v string) *RestartPxfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartPxfsResponseBody) Validate() error {
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

type RestartPxfsResponseBodyAccessDeniedDetail struct {
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
	// The diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the missing permission.
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

func (s RestartPxfsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RestartPxfsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RestartPxfsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RestartPxfsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RestartPxfsResponseBodyData struct {
	// The PolarDB-X instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The Agent file service instance ID.
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

func (s RestartPxfsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartPxfsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartPxfsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartPxfsResponseBodyData) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *RestartPxfsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *RestartPxfsResponseBodyData) SetDBInstanceName(v string) *RestartPxfsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *RestartPxfsResponseBodyData) SetPxfsInstanceId(v string) *RestartPxfsResponseBodyData {
	s.PxfsInstanceId = &v
	return s
}

func (s *RestartPxfsResponseBodyData) SetTaskId(v int32) *RestartPxfsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RestartPxfsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
