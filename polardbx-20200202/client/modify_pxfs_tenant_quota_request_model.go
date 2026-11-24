// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsTenantQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyPxfsTenantQuotaRequest
	GetDBInstanceName() *string
	SetQuotaBytes(v int64) *ModifyPxfsTenantQuotaRequest
	GetQuotaBytes() *int64
	SetQuotaFiles(v int64) *ModifyPxfsTenantQuotaRequest
	GetQuotaFiles() *int64
	SetRateLimitRps(v int64) *ModifyPxfsTenantQuotaRequest
	GetRateLimitRps() *int64
	SetRateLimitWbps(v int64) *ModifyPxfsTenantQuotaRequest
	GetRateLimitWbps() *int64
	SetRegionId(v string) *ModifyPxfsTenantQuotaRequest
	GetRegionId() *string
	SetTenantId(v string) *ModifyPxfsTenantQuotaRequest
	GetTenantId() *string
}

type ModifyPxfsTenantQuotaRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyPxfsTenantQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsTenantQuotaRequest) GoString() string {
	return s.String()
}

func (s *ModifyPxfsTenantQuotaRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyPxfsTenantQuotaRequest) GetQuotaBytes() *int64 {
	return s.QuotaBytes
}

func (s *ModifyPxfsTenantQuotaRequest) GetQuotaFiles() *int64 {
	return s.QuotaFiles
}

func (s *ModifyPxfsTenantQuotaRequest) GetRateLimitRps() *int64 {
	return s.RateLimitRps
}

func (s *ModifyPxfsTenantQuotaRequest) GetRateLimitWbps() *int64 {
	return s.RateLimitWbps
}

func (s *ModifyPxfsTenantQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPxfsTenantQuotaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ModifyPxfsTenantQuotaRequest) SetDBInstanceName(v string) *ModifyPxfsTenantQuotaRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetQuotaBytes(v int64) *ModifyPxfsTenantQuotaRequest {
	s.QuotaBytes = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetQuotaFiles(v int64) *ModifyPxfsTenantQuotaRequest {
	s.QuotaFiles = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetRateLimitRps(v int64) *ModifyPxfsTenantQuotaRequest {
	s.RateLimitRps = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetRateLimitWbps(v int64) *ModifyPxfsTenantQuotaRequest {
	s.RateLimitWbps = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetRegionId(v string) *ModifyPxfsTenantQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) SetTenantId(v string) *ModifyPxfsTenantQuotaRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyPxfsTenantQuotaRequest) Validate() error {
	return dara.Validate(s)
}
