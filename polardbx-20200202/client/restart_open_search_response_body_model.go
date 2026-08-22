// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOpenSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RestartOpenSearchResponseBodyAccessDeniedDetail) *RestartOpenSearchResponseBody
	GetAccessDeniedDetail() *RestartOpenSearchResponseBodyAccessDeniedDetail
	SetData(v *RestartOpenSearchResponseBodyData) *RestartOpenSearchResponseBody
	GetData() *RestartOpenSearchResponseBodyData
	SetRequestId(v string) *RestartOpenSearchResponseBody
	GetRequestId() *string
}

type RestartOpenSearchResponseBody struct {
	AccessDeniedDetail *RestartOpenSearchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *RestartOpenSearchResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AE4F6C34-065F-45AA-B5DC-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartOpenSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartOpenSearchResponseBody) GoString() string {
	return s.String()
}

func (s *RestartOpenSearchResponseBody) GetAccessDeniedDetail() *RestartOpenSearchResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RestartOpenSearchResponseBody) GetData() *RestartOpenSearchResponseBodyData {
	return s.Data
}

func (s *RestartOpenSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartOpenSearchResponseBody) SetAccessDeniedDetail(v *RestartOpenSearchResponseBodyAccessDeniedDetail) *RestartOpenSearchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RestartOpenSearchResponseBody) SetData(v *RestartOpenSearchResponseBodyData) *RestartOpenSearchResponseBody {
	s.Data = v
	return s
}

func (s *RestartOpenSearchResponseBody) SetRequestId(v string) *RestartOpenSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartOpenSearchResponseBody) Validate() error {
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

type RestartOpenSearchResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s RestartOpenSearchResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RestartOpenSearchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RestartOpenSearchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RestartOpenSearchResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RestartOpenSearchResponseBodyData struct {
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartOpenSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartOpenSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartOpenSearchResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *RestartOpenSearchResponseBodyData) SetTaskId(v int32) *RestartOpenSearchResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RestartOpenSearchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
