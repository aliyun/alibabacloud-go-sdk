// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPxfsTenantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListPxfsTenantsResponseBodyAccessDeniedDetail) *ListPxfsTenantsResponseBody
	GetAccessDeniedDetail() *ListPxfsTenantsResponseBodyAccessDeniedDetail
	SetData(v *ListPxfsTenantsResponseBodyData) *ListPxfsTenantsResponseBody
	GetData() *ListPxfsTenantsResponseBodyData
	SetRequestId(v string) *ListPxfsTenantsResponseBody
	GetRequestId() *string
}

type ListPxfsTenantsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ListPxfsTenantsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data structure.
	Data *ListPxfsTenantsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPxfsTenantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsResponseBody) GetAccessDeniedDetail() *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListPxfsTenantsResponseBody) GetData() *ListPxfsTenantsResponseBodyData {
	return s.Data
}

func (s *ListPxfsTenantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPxfsTenantsResponseBody) SetAccessDeniedDetail(v *ListPxfsTenantsResponseBodyAccessDeniedDetail) *ListPxfsTenantsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListPxfsTenantsResponseBody) SetData(v *ListPxfsTenantsResponseBodyData) *ListPxfsTenantsResponseBody {
	s.Data = v
	return s
}

func (s *ListPxfsTenantsResponseBody) SetRequestId(v string) *ListPxfsTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPxfsTenantsResponseBody) Validate() error {
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

type ListPxfsTenantsResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// polardbx:ListPxfsTenants
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// Sample RAM user
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

func (s ListPxfsTenantsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListPxfsTenantsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListPxfsTenantsResponseBodyData struct {
	// The list of PXFS tenants.
	Tenants []*ListPxfsTenantsResponseBodyDataTenants `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
}

func (s ListPxfsTenantsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsResponseBodyData) GetTenants() []*ListPxfsTenantsResponseBodyDataTenants {
	return s.Tenants
}

func (s *ListPxfsTenantsResponseBodyData) SetTenants(v []*ListPxfsTenantsResponseBodyDataTenants) *ListPxfsTenantsResponseBodyData {
	s.Tenants = v
	return s
}

func (s *ListPxfsTenantsResponseBodyData) Validate() error {
	if s.Tenants != nil {
		for _, item := range s.Tenants {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPxfsTenantsResponseBodyDataTenants struct {
	// The number of bytes saved by data deduplication for the tenant.
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
	// The current number of logical files for the tenant.
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
	// 0
	RateLimitRps *int64 `json:"RateLimitRps,omitempty" xml:"RateLimitRps,omitempty"`
	// The maximum write bandwidth.
	//
	// example:
	//
	// 0
	RateLimitWbps *int64 `json:"RateLimitWbps,omitempty" xml:"RateLimitWbps,omitempty"`
	// The status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// zero-v1zhme7mtddrkfn40hrr9ifk
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// zero-v1zhme7mtddrkfn40hrr9ifk
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s ListPxfsTenantsResponseBodyDataTenants) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsResponseBodyDataTenants) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetDedupBytes() *int64 {
	return s.DedupBytes
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetLogicalBytes() *int64 {
	return s.LogicalBytes
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetLogicalFiles() *int64 {
	return s.LogicalFiles
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetPhysicalBytes() *int64 {
	return s.PhysicalBytes
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetStatus() *string {
	return s.Status
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPxfsTenantsResponseBodyDataTenants) GetTenantName() *string {
	return s.TenantName
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetDedupBytes(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.DedupBytes = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetLogicalBytes(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.LogicalBytes = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetLogicalFiles(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.LogicalFiles = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetPhysicalBytes(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.PhysicalBytes = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetQuotaBytes(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.QuotaBytes = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetQuotaFiles(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.QuotaFiles = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetRateLimitRps(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.RateLimitRps = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetRateLimitWbps(v int64) *ListPxfsTenantsResponseBodyDataTenants {
	s.RateLimitWbps = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetStatus(v string) *ListPxfsTenantsResponseBodyDataTenants {
	s.Status = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetTenantId(v string) *ListPxfsTenantsResponseBodyDataTenants {
	s.TenantId = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) SetTenantName(v string) *ListPxfsTenantsResponseBodyDataTenants {
	s.TenantName = &v
	return s
}

func (s *ListPxfsTenantsResponseBodyDataTenants) Validate() error {
	return dara.Validate(s)
}
