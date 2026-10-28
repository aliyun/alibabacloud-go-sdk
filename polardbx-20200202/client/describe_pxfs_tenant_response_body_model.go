// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfsTenantResponseBodyAccessDeniedDetail) *DescribePxfsTenantResponseBody
	GetAccessDeniedDetail() *DescribePxfsTenantResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfsTenantResponseBodyData) *DescribePxfsTenantResponseBody
	GetData() *DescribePxfsTenantResponseBodyData
	SetRequestId(v string) *DescribePxfsTenantResponseBody
	GetRequestId() *string
}

type DescribePxfsTenantResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribePxfsTenantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data.
	Data *DescribePxfsTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfsTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsTenantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfsTenantResponseBody) GetAccessDeniedDetail() *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfsTenantResponseBody) GetData() *DescribePxfsTenantResponseBodyData {
	return s.Data
}

func (s *DescribePxfsTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfsTenantResponseBody) SetAccessDeniedDetail(v *DescribePxfsTenantResponseBodyAccessDeniedDetail) *DescribePxfsTenantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfsTenantResponseBody) SetData(v *DescribePxfsTenantResponseBodyData) *DescribePxfsTenantResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfsTenantResponseBody) SetRequestId(v string) *DescribePxfsTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfsTenantResponseBody) Validate() error {
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

type DescribePxfsTenantResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// polardbx:QueryColumnarLog
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

func (s DescribePxfsTenantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsTenantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfsTenantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfsTenantResponseBodyData struct {
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
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// agent-files
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The used capacity.
	//
	// example:
	//
	// 1048576
	UsedBytes *int64 `json:"UsedBytes,omitempty" xml:"UsedBytes,omitempty"`
	// The used file count.
	//
	// example:
	//
	// 128
	UsedFiles *int64 `json:"UsedFiles,omitempty" xml:"UsedFiles,omitempty"`
}

func (s DescribePxfsTenantResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfsTenantResponseBodyData) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *DescribePxfsTenantResponseBodyData) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *DescribePxfsTenantResponseBodyData) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *DescribePxfsTenantResponseBodyData) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *DescribePxfsTenantResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribePxfsTenantResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribePxfsTenantResponseBodyData) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribePxfsTenantResponseBodyData) GetUsedBytes() *int64 {
	return s.UsedBytes
}

func (s *DescribePxfsTenantResponseBodyData) GetUsedFiles() *int64 {
	return s.UsedFiles
}

func (s *DescribePxfsTenantResponseBodyData) SetQuotaBytes(v int64) *DescribePxfsTenantResponseBodyData {
	s.QuotaBytes = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetQuotaFiles(v int64) *DescribePxfsTenantResponseBodyData {
	s.QuotaFiles = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetRateLimitRps(v int64) *DescribePxfsTenantResponseBodyData {
	s.RateLimitRps = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetRateLimitWbps(v int64) *DescribePxfsTenantResponseBodyData {
	s.RateLimitWbps = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetStatus(v string) *DescribePxfsTenantResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetTenantId(v string) *DescribePxfsTenantResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetTenantName(v string) *DescribePxfsTenantResponseBodyData {
	s.TenantName = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetUsedBytes(v int64) *DescribePxfsTenantResponseBodyData {
	s.UsedBytes = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) SetUsedFiles(v int64) *DescribePxfsTenantResponseBodyData {
	s.UsedFiles = &v
	return s
}

func (s *DescribePxfsTenantResponseBodyData) Validate() error {
	return dara.Validate(s)
}
