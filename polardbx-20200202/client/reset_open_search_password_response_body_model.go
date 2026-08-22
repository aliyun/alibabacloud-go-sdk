// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetOpenSearchPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) *ResetOpenSearchPasswordResponseBody
	GetAccessDeniedDetail() *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail
	SetData(v *ResetOpenSearchPasswordResponseBodyData) *ResetOpenSearchPasswordResponseBody
	GetData() *ResetOpenSearchPasswordResponseBodyData
	SetRequestId(v string) *ResetOpenSearchPasswordResponseBody
	GetRequestId() *string
}

type ResetOpenSearchPasswordResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *ResetOpenSearchPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetOpenSearchPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetOpenSearchPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetOpenSearchPasswordResponseBody) GetAccessDeniedDetail() *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ResetOpenSearchPasswordResponseBody) GetData() *ResetOpenSearchPasswordResponseBodyData {
	return s.Data
}

func (s *ResetOpenSearchPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetOpenSearchPasswordResponseBody) SetAccessDeniedDetail(v *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) *ResetOpenSearchPasswordResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ResetOpenSearchPasswordResponseBody) SetData(v *ResetOpenSearchPasswordResponseBodyData) *ResetOpenSearchPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetOpenSearchPasswordResponseBody) SetRequestId(v string) *ResetOpenSearchPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBody) Validate() error {
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

type ResetOpenSearchPasswordResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ResetOpenSearchPasswordResponseBodyData struct {
	// The additional information returned by the operation. "success" is returned if the operation is successful. Otherwise, the corresponding error code is returned.
	//
	// example:
	//
	// 【环境：huanghe】\\nhuanghe 503_UC_OUTBOUND告警超过阈值！\\n\\n详情请查看: https://grafana-cn-lbj34sreu03.grafana.aliyuncs.com/d/_rOiq2lNk/asm-status-code-monitor?var-datasource=DataSource-HUANGHE\\n
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the password was last modified.
	//
	// example:
	//
	// 2026-08-21T12:00:00Z
	PasswordLastModified *string `json:"PasswordLastModified,omitempty" xml:"PasswordLastModified,omitempty"`
}

func (s ResetOpenSearchPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetOpenSearchPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetOpenSearchPasswordResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ResetOpenSearchPasswordResponseBodyData) GetPasswordLastModified() *string {
	return s.PasswordLastModified
}

func (s *ResetOpenSearchPasswordResponseBodyData) SetMessage(v string) *ResetOpenSearchPasswordResponseBodyData {
	s.Message = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyData) SetPasswordLastModified(v string) *ResetOpenSearchPasswordResponseBodyData {
	s.PasswordLastModified = &v
	return s
}

func (s *ResetOpenSearchPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
