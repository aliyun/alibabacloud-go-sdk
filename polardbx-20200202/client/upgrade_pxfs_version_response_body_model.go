// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePxfsVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpgradePxfsVersionResponseBodyAccessDeniedDetail) *UpgradePxfsVersionResponseBody
	GetAccessDeniedDetail() *UpgradePxfsVersionResponseBodyAccessDeniedDetail
	SetData(v *UpgradePxfsVersionResponseBodyData) *UpgradePxfsVersionResponseBody
	GetData() *UpgradePxfsVersionResponseBodyData
	SetRequestId(v string) *UpgradePxfsVersionResponseBody
	GetRequestId() *string
}

type UpgradePxfsVersionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *UpgradePxfsVersionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *UpgradePxfsVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradePxfsVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePxfsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePxfsVersionResponseBody) GetAccessDeniedDetail() *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpgradePxfsVersionResponseBody) GetData() *UpgradePxfsVersionResponseBodyData {
	return s.Data
}

func (s *UpgradePxfsVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePxfsVersionResponseBody) SetAccessDeniedDetail(v *UpgradePxfsVersionResponseBodyAccessDeniedDetail) *UpgradePxfsVersionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpgradePxfsVersionResponseBody) SetData(v *UpgradePxfsVersionResponseBodyData) *UpgradePxfsVersionResponseBody {
	s.Data = v
	return s
}

func (s *UpgradePxfsVersionResponseBody) SetRequestId(v string) *UpgradePxfsVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePxfsVersionResponseBody) Validate() error {
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

type UpgradePxfsVersionResponseBodyAccessDeniedDetail struct {
	// The authorization action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authorization in the request.
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
	// The authorization principal type.
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
	// The type of missing permission.
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

func (s UpgradePxfsVersionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpgradePxfsVersionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpgradePxfsVersionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type UpgradePxfsVersionResponseBodyData struct {
	// The PolarDB-X instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The target minor version.
	//
	// example:
	//
	// 1.1.0
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
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

func (s UpgradePxfsVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradePxfsVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradePxfsVersionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradePxfsVersionResponseBodyData) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradePxfsVersionResponseBodyData) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *UpgradePxfsVersionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *UpgradePxfsVersionResponseBodyData) SetDBInstanceName(v string) *UpgradePxfsVersionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyData) SetMinorVersion(v string) *UpgradePxfsVersionResponseBodyData {
	s.MinorVersion = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyData) SetPxfsInstanceId(v string) *UpgradePxfsVersionResponseBodyData {
	s.PxfsInstanceId = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyData) SetTaskId(v int32) *UpgradePxfsVersionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpgradePxfsVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
