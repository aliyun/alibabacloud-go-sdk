// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreatePxfsTenantRequest
	GetDBInstanceName() *string
	SetQuotaBytes(v int64) *CreatePxfsTenantRequest
	GetQuotaBytes() *int64
	SetQuotaFiles(v int64) *CreatePxfsTenantRequest
	GetQuotaFiles() *int64
	SetRateLimitRps(v int64) *CreatePxfsTenantRequest
	GetRateLimitRps() *int64
	SetRateLimitWbps(v int64) *CreatePxfsTenantRequest
	GetRateLimitWbps() *int64
	SetRegionId(v string) *CreatePxfsTenantRequest
	GetRegionId() *string
	SetTenantId(v string) *CreatePxfsTenantRequest
	GetTenantId() *string
	SetTenantName(v string) *CreatePxfsTenantRequest
	GetTenantName() *string
}

type CreatePxfsTenantRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The capacity quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 107374182400
	QuotaBytes *int64 `json:"QuotaBytes,omitempty" xml:"QuotaBytes,omitempty"`
	// The file count quota.
	//
	// This parameter is required.
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-files
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreatePxfsTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsTenantRequest) GoString() string {
	return s.String()
}

func (s *CreatePxfsTenantRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfsTenantRequest) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *CreatePxfsTenantRequest) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *CreatePxfsTenantRequest) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *CreatePxfsTenantRequest) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *CreatePxfsTenantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePxfsTenantRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsTenantRequest) GetTenantName() *string {
	return s.TenantName
}

func (s *CreatePxfsTenantRequest) SetDBInstanceName(v string) *CreatePxfsTenantRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetQuotaBytes(v int64) *CreatePxfsTenantRequest {
	s.QuotaBytes = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetQuotaFiles(v int64) *CreatePxfsTenantRequest {
	s.QuotaFiles = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetRateLimitRps(v int64) *CreatePxfsTenantRequest {
	s.RateLimitRps = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetRateLimitWbps(v int64) *CreatePxfsTenantRequest {
	s.RateLimitWbps = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetRegionId(v string) *CreatePxfsTenantRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetTenantId(v string) *CreatePxfsTenantRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsTenantRequest) SetTenantName(v string) *CreatePxfsTenantRequest {
	s.TenantName = &v
	return s
}

func (s *CreatePxfsTenantRequest) Validate() error {
	return dara.Validate(s)
}
