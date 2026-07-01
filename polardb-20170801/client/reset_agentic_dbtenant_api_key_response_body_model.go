// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticDBTenantApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetApiKey() *string
	SetApiKeyId(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetApiKeyId() *string
	SetCreateTime(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetExpireTime() *string
	SetRequestId(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetRequestId() *string
	SetTenantId(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetTenantId() *string
	SetTenantName(v string) *ResetAgenticDBTenantApiKeyResponseBody
	GetTenantName() *string
}

type ResetAgenticDBTenantApiKeyResponseBody struct {
	// example:
	//
	// pagc_key_cGFnYy1icDFh...newSignature22ch
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// ak-71304e39c7e841a1
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// 2026-06-10T14:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// B2C3D4E5-F6A7-8901-BCDE-F12345678901
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// my-saas-app
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s ResetAgenticDBTenantApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticDBTenantApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) GetTenantName() *string {
	return s.TenantName
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetApiKey(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetApiKeyId(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.ApiKeyId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetCreateTime(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetExpireTime(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetRequestId(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetTenantId(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.TenantId = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) SetTenantName(v string) *ResetAgenticDBTenantApiKeyResponseBody {
	s.TenantName = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
