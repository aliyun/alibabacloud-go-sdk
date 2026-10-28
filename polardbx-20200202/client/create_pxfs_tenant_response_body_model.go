// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePxfsTenantResponseBodyAccessDeniedDetail) *CreatePxfsTenantResponseBody
	GetAccessDeniedDetail() *CreatePxfsTenantResponseBodyAccessDeniedDetail
	SetData(v *CreatePxfsTenantResponseBodyData) *CreatePxfsTenantResponseBody
	GetData() *CreatePxfsTenantResponseBodyData
	SetRequestId(v string) *CreatePxfsTenantResponseBody
	GetRequestId() *string
}

type CreatePxfsTenantResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePxfsTenantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *CreatePxfsTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePxfsTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsTenantResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePxfsTenantResponseBody) GetAccessDeniedDetail() *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePxfsTenantResponseBody) GetData() *CreatePxfsTenantResponseBodyData {
	return s.Data
}

func (s *CreatePxfsTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePxfsTenantResponseBody) SetAccessDeniedDetail(v *CreatePxfsTenantResponseBodyAccessDeniedDetail) *CreatePxfsTenantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePxfsTenantResponseBody) SetData(v *CreatePxfsTenantResponseBodyData) *CreatePxfsTenantResponseBody {
	s.Data = v
	return s
}

func (s *CreatePxfsTenantResponseBody) SetRequestId(v string) *CreatePxfsTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePxfsTenantResponseBody) Validate() error {
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

type CreatePxfsTenantResponseBodyAccessDeniedDetail struct {
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

func (s CreatePxfsTenantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsTenantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePxfsTenantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePxfsTenantResponseBodyData struct {
	// The capacity quota.
	//
	// example:
	//
	// 107374182400
	QuotaBytes *int64 `json:"QuotaBytes,omitempty" xml:"QuotaBytes,omitempty"`
	// The file count quota.
	//
	// example:
	//
	// 1000000
	QuotaFiles *int64 `json:"QuotaFiles,omitempty" xml:"QuotaFiles,omitempty"`
	// The request rate limit.
	//
	// example:
	//
	// 1000
	RateLimitRps *int64 `json:"RateLimitRps,omitempty" xml:"RateLimitRps,omitempty"`
	// The write bandwidth limit.
	//
	// example:
	//
	// 104857600
	RateLimitWbps *int64 `json:"RateLimitWbps,omitempty" xml:"RateLimitWbps,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// agent-files
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreatePxfsTenantResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePxfsTenantResponseBodyData) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *CreatePxfsTenantResponseBodyData) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *CreatePxfsTenantResponseBodyData) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *CreatePxfsTenantResponseBodyData) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *CreatePxfsTenantResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreatePxfsTenantResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsTenantResponseBodyData) GetTenantName() *string {
	return s.TenantName
}

func (s *CreatePxfsTenantResponseBodyData) SetQuotaBytes(v int64) *CreatePxfsTenantResponseBodyData {
	s.QuotaBytes = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetQuotaFiles(v int64) *CreatePxfsTenantResponseBodyData {
	s.QuotaFiles = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetRateLimitRps(v int64) *CreatePxfsTenantResponseBodyData {
	s.RateLimitRps = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetRateLimitWbps(v int64) *CreatePxfsTenantResponseBodyData {
	s.RateLimitWbps = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetStatus(v string) *CreatePxfsTenantResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetTenantId(v string) *CreatePxfsTenantResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) SetTenantName(v string) *CreatePxfsTenantResponseBodyData {
	s.TenantName = &v
	return s
}

func (s *CreatePxfsTenantResponseBodyData) Validate() error {
	return dara.Validate(s)
}
