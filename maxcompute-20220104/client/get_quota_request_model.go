// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkProven(v string) *GetQuotaRequest
	GetAkProven() *string
	SetMock(v bool) *GetQuotaRequest
	GetMock() *bool
	SetRegion(v string) *GetQuotaRequest
	GetRegion() *string
	SetTenantId(v string) *GetQuotaRequest
	GetTenantId() *string
}

type GetQuotaRequest struct {
	// Deprecated
	//
	// The trusted AccessKey pairs.
	//
	// example:
	//
	// null
	AkProven *string `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
	// Specifies whether to include submodules. Valid values: -true: The request includes submodules. -false: The request does not include submodules. This is the default value.
	//
	// example:
	//
	// false
	Mock *bool `json:"mock,omitempty" xml:"mock,omitempty"`
	// Deprecated
	//
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Deprecated
	//
	// The tenant ID.
	//
	// example:
	//
	// 520539530998273
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) GetAkProven() *string {
	return s.AkProven
}

func (s *GetQuotaRequest) GetMock() *bool {
	return s.Mock
}

func (s *GetQuotaRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQuotaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaRequest) SetAkProven(v string) *GetQuotaRequest {
	s.AkProven = &v
	return s
}

func (s *GetQuotaRequest) SetMock(v bool) *GetQuotaRequest {
	s.Mock = &v
	return s
}

func (s *GetQuotaRequest) SetRegion(v string) *GetQuotaRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaRequest) SetTenantId(v string) *GetQuotaRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaRequest) Validate() error {
	return dara.Validate(s)
}
