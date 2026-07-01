// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticDBTenantApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *ResetAgenticDBTenantApiKeyRequest
	GetApiKeyId() *string
	SetDBClusterId(v string) *ResetAgenticDBTenantApiKeyRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ResetAgenticDBTenantApiKeyRequest
	GetRegionId() *string
	SetTenantId(v string) *ResetAgenticDBTenantApiKeyRequest
	GetTenantId() *string
}

type ResetAgenticDBTenantApiKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ak-71304e39c7e841a1
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ResetAgenticDBTenantApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticDBTenantApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ResetAgenticDBTenantApiKeyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ResetAgenticDBTenantApiKeyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetAgenticDBTenantApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetAgenticDBTenantApiKeyRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetAgenticDBTenantApiKeyRequest) SetApiKeyId(v string) *ResetAgenticDBTenantApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyRequest) SetDBClusterId(v string) *ResetAgenticDBTenantApiKeyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyRequest) SetRegionId(v string) *ResetAgenticDBTenantApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyRequest) SetTenantId(v string) *ResetAgenticDBTenantApiKeyRequest {
	s.TenantId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
