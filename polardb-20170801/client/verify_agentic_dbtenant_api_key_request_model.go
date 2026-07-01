// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgenticDBTenantApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *VerifyAgenticDBTenantApiKeyRequest
	GetApiKey() *string
	SetRegionId(v string) *VerifyAgenticDBTenantApiKeyRequest
	GetRegionId() *string
}

type VerifyAgenticDBTenantApiKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pagc_key_cGFnYy1icDFh...kX9mP2vL7wQ3
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VerifyAgenticDBTenantApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgenticDBTenantApiKeyRequest) GoString() string {
	return s.String()
}

func (s *VerifyAgenticDBTenantApiKeyRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *VerifyAgenticDBTenantApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyAgenticDBTenantApiKeyRequest) SetApiKey(v string) *VerifyAgenticDBTenantApiKeyRequest {
	s.ApiKey = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyRequest) SetRegionId(v string) *VerifyAgenticDBTenantApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
