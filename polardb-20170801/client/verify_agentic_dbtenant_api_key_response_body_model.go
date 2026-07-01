// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgenticDBTenantApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetApiKeyId() *string
	SetDBClusterId(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetDBClusterId() *string
	SetExpireTime(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetExpireTime() *string
	SetReason(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetReason() *string
	SetRegionId(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetRequestId() *string
	SetTenantId(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetTenantId() *string
	SetTenantName(v string) *VerifyAgenticDBTenantApiKeyResponseBody
	GetTenantName() *string
	SetValid(v bool) *VerifyAgenticDBTenantApiKeyResponseBody
	GetValid() *bool
}

type VerifyAgenticDBTenantApiKeyResponseBody struct {
	// example:
	//
	// ak-71304e39c7e841a1
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// D4E5F6A7-B8C9-0123-DEFA-234567890123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// my-saas-app
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s VerifyAgenticDBTenantApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgenticDBTenantApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetReason() *string {
	return s.Reason
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetTenantName() *string {
	return s.TenantName
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) GetValid() *bool {
	return s.Valid
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetApiKeyId(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.ApiKeyId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetDBClusterId(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetExpireTime(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetReason(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.Reason = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetRegionId(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.RegionId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetRequestId(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetTenantId(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.TenantId = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetTenantName(v string) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.TenantName = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) SetValid(v bool) *VerifyAgenticDBTenantApiKeyResponseBody {
	s.Valid = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
