// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkProven(v string) *QueryQuotaRequest
	GetAkProven() *string
	SetMock(v bool) *QueryQuotaRequest
	GetMock() *bool
	SetRegion(v string) *QueryQuotaRequest
	GetRegion() *string
	SetTenantId(v string) *QueryQuotaRequest
	GetTenantId() *string
}

type QueryQuotaRequest struct {
	AkProven *string `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
	Mock     *bool   `json:"mock,omitempty" xml:"mock,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryQuotaRequest) GetAkProven() *string {
	return s.AkProven
}

func (s *QueryQuotaRequest) GetMock() *bool {
	return s.Mock
}

func (s *QueryQuotaRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryQuotaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryQuotaRequest) SetAkProven(v string) *QueryQuotaRequest {
	s.AkProven = &v
	return s
}

func (s *QueryQuotaRequest) SetMock(v bool) *QueryQuotaRequest {
	s.Mock = &v
	return s
}

func (s *QueryQuotaRequest) SetRegion(v string) *QueryQuotaRequest {
	s.Region = &v
	return s
}

func (s *QueryQuotaRequest) SetTenantId(v string) *QueryQuotaRequest {
	s.TenantId = &v
	return s
}

func (s *QueryQuotaRequest) Validate() error {
	return dara.Validate(s)
}
