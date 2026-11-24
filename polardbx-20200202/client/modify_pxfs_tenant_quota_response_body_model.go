// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsTenantQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) *ModifyPxfsTenantQuotaResponseBody
	GetAccessDeniedDetail() *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail
	SetData(v *ModifyPxfsTenantQuotaResponseBodyData) *ModifyPxfsTenantQuotaResponseBody
	GetData() *ModifyPxfsTenantQuotaResponseBodyData
	SetRequestId(v string) *ModifyPxfsTenantQuotaResponseBody
	GetRequestId() *string
}

type ModifyPxfsTenantQuotaResponseBody struct {
	// The details of the access denied error.
	AccessDeniedDetail *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The new connection string information.
	Data *ModifyPxfsTenantQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPxfsTenantQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsTenantQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPxfsTenantQuotaResponseBody) GetAccessDeniedDetail() *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyPxfsTenantQuotaResponseBody) GetData() *ModifyPxfsTenantQuotaResponseBodyData {
	return s.Data
}

func (s *ModifyPxfsTenantQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPxfsTenantQuotaResponseBody) SetAccessDeniedDetail(v *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) *ModifyPxfsTenantQuotaResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBody) SetData(v *ModifyPxfsTenantQuotaResponseBodyData) *ModifyPxfsTenantQuotaResponseBody {
	s.Data = v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBody) SetRequestId(v string) *ModifyPxfsTenantQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBody) Validate() error {
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

type ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail struct {
	// The name of the API operation.
	//
	// example:
	//
	// polardbx:ModifyPxfsTenantQuota
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authenticated principal.
	//
	// example:
	//
	// Sample RAM user
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authenticated principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of identity used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user.
	//
	// - AssumedRoleUser: RAM role.
	//
	// - Federated: SSO federated identity.
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyPxfsTenantQuotaResponseBodyData struct {
	// The number of bytes saved for the tenant after data deduplication.
	//
	// example:
	//
	// 268435456
	DedupBytes *int64 `json:"DedupBytes,omitempty" xml:"DedupBytes,omitempty"`
	// The current logical usage of the tenant, in bytes.
	//
	// example:
	//
	// 1073741824
	LogicalBytes *int64 `json:"LogicalBytes,omitempty" xml:"LogicalBytes,omitempty"`
	// The current number of logical files of the tenant.
	//
	// example:
	//
	// 1000
	LogicalFiles *int64 `json:"LogicalFiles,omitempty" xml:"LogicalFiles,omitempty"`
	// The actual physical storage space occupied by the tenant, in bytes.
	//
	// example:
	//
	// 805306368
	PhysicalBytes *int64 `json:"PhysicalBytes,omitempty" xml:"PhysicalBytes,omitempty"`
	// The storage capacity quota.
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
	// The maximum request rate.
	//
	// example:
	//
	// 1000
	RateLimitRps *int64 `json:"RateLimitRps,omitempty" xml:"RateLimitRps,omitempty"`
	// The maximum write bandwidth.
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
}

func (s ModifyPxfsTenantQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsTenantQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetDedupBytes() *int64 {
	return s.DedupBytes
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetLogicalBytes() *int64 {
	return s.LogicalBytes
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetLogicalFiles() *int64 {
	return s.LogicalFiles
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetPhysicalBytes() *int64 {
	return s.PhysicalBytes
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) GetTenantName() *string {
	return s.TenantName
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetDedupBytes(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.DedupBytes = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetLogicalBytes(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.LogicalBytes = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetLogicalFiles(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.LogicalFiles = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetPhysicalBytes(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.PhysicalBytes = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetQuotaBytes(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.QuotaBytes = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetQuotaFiles(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.QuotaFiles = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetRateLimitRps(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.RateLimitRps = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetRateLimitWbps(v int64) *ModifyPxfsTenantQuotaResponseBodyData {
	s.RateLimitWbps = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetStatus(v string) *ModifyPxfsTenantQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetTenantId(v string) *ModifyPxfsTenantQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) SetTenantName(v string) *ModifyPxfsTenantQuotaResponseBodyData {
	s.TenantName = &v
	return s
}

func (s *ModifyPxfsTenantQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
