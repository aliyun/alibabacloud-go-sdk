// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetApiKey() *string
	SetApiKeyId(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetApiKeyId() *string
	SetCreateTime(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetExpireTime() *string
	SetRequestId(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetRequestId() *string
	SetTenantId(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetTenantId() *string
	SetTenantName(v string) *CreateAgenticDBTenantApiKeyResponseBody
	GetTenantName() *string
}

type CreateAgenticDBTenantApiKeyResponseBody struct {
	// The plaintext of the API key. This value is returned only in this response.
	//
	// example:
	//
	// pagc_key_xxxx.yyyy
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The resource ID of the API key.
	//
	// example:
	//
	// ak-xxxxxxxxxxxx
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The time when the API key was created.
	//
	// example:
	//
	// 2026-06-10T08:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the API key. This value is empty if the API key never expires.
	//
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-1B2C-3D4E-5F6A-7B8C9D0E1F2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t-aaaa111122223333
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// my-tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreateAgenticDBTenantApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) GetTenantName() *string {
	return s.TenantName
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetApiKey(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetApiKeyId(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.ApiKeyId = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetCreateTime(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetExpireTime(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetRequestId(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetTenantId(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) SetTenantName(v string) *CreateAgenticDBTenantApiKeyResponseBody {
	s.TenantName = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
