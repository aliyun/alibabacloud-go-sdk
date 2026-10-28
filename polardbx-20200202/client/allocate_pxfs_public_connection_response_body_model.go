// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePxfsPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) *AllocatePxfsPublicConnectionResponseBody
	GetAccessDeniedDetail() *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *AllocatePxfsPublicConnectionResponseBodyData) *AllocatePxfsPublicConnectionResponseBody
	GetData() *AllocatePxfsPublicConnectionResponseBodyData
	SetRequestId(v string) *AllocatePxfsPublicConnectionResponseBody
	GetRequestId() *string
}

type AllocatePxfsPublicConnectionResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *AllocatePxfsPublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePxfsPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocatePxfsPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePxfsPublicConnectionResponseBody) GetAccessDeniedDetail() *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AllocatePxfsPublicConnectionResponseBody) GetData() *AllocatePxfsPublicConnectionResponseBodyData {
	return s.Data
}

func (s *AllocatePxfsPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocatePxfsPublicConnectionResponseBody) SetAccessDeniedDetail(v *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) *AllocatePxfsPublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBody) SetData(v *AllocatePxfsPublicConnectionResponseBodyData) *AllocatePxfsPublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBody) SetRequestId(v string) *AllocatePxfsPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBody) Validate() error {
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

type AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail struct {
	// Same as above.
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
	// Same as above.
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
	// The diagnostic message.
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

func (s AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AllocatePxfsPublicConnectionResponseBodyData struct {
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

func (s AllocatePxfsPublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AllocatePxfsPublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) SetDBInstanceName(v string) *AllocatePxfsPublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) SetPxfsInstanceId(v string) *AllocatePxfsPublicConnectionResponseBodyData {
	s.PxfsInstanceId = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) SetTaskId(v int32) *AllocatePxfsPublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AllocatePxfsPublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
