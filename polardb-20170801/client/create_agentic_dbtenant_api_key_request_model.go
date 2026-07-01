// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAgenticDBTenantApiKeyRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateAgenticDBTenantApiKeyRequest
	GetDescription() *string
	SetExpireTime(v string) *CreateAgenticDBTenantApiKeyRequest
	GetExpireTime() *string
	SetRegionId(v string) *CreateAgenticDBTenantApiKeyRequest
	GetRegionId() *string
	SetTenantName(v string) *CreateAgenticDBTenantApiKeyRequest
	GetTenantName() *string
}

type CreateAgenticDBTenantApiKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// MCP Server专用Key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreateAgenticDBTenantApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantApiKeyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBTenantApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticDBTenantApiKeyRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateAgenticDBTenantApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgenticDBTenantApiKeyRequest) GetTenantName() *string {
	return s.TenantName
}

func (s *CreateAgenticDBTenantApiKeyRequest) SetDBClusterId(v string) *CreateAgenticDBTenantApiKeyRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyRequest) SetDescription(v string) *CreateAgenticDBTenantApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyRequest) SetExpireTime(v string) *CreateAgenticDBTenantApiKeyRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyRequest) SetRegionId(v string) *CreateAgenticDBTenantApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyRequest) SetTenantName(v string) *CreateAgenticDBTenantApiKeyRequest {
	s.TenantName = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
