// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *DeleteAgenticDBTenantApiKeyRequest
	GetApiKeyId() *string
	SetDBClusterId(v string) *DeleteAgenticDBTenantApiKeyRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteAgenticDBTenantApiKeyRequest
	GetRegionId() *string
	SetTenantId(v string) *DeleteAgenticDBTenantApiKeyRequest
	GetTenantId() *string
}

type DeleteAgenticDBTenantApiKeyRequest struct {
	// The ID of the API key to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ak-71304e39c7e841a1
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID, which is used for secondary authentication on the backend.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAgenticDBTenantApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantApiKeyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *DeleteAgenticDBTenantApiKeyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgenticDBTenantApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgenticDBTenantApiKeyRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAgenticDBTenantApiKeyRequest) SetApiKeyId(v string) *DeleteAgenticDBTenantApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyRequest) SetDBClusterId(v string) *DeleteAgenticDBTenantApiKeyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyRequest) SetRegionId(v string) *DeleteAgenticDBTenantApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyRequest) SetTenantId(v string) *DeleteAgenticDBTenantApiKeyRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
