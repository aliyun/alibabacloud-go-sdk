// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleTenant interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SingleTenant
	GetDescription() *string
	SetStatus(v string) *SingleTenant
	GetStatus() *string
	SetTenantId(v string) *SingleTenant
	GetTenantId() *string
	SetTenantName(v string) *SingleTenant
	GetTenantName() *string
	SetUUID(v string) *SingleTenant
	GetUUID() *string
}

type SingleTenant struct {
	// Tenant description
	//
	// example:
	//
	// 该租户用来测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Tenant status
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Tenant ID
	//
	// example:
	//
	// GA***134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Tenant name
	//
	// example:
	//
	// demo
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// Tenant UUID
	//
	// example:
	//
	// paiworkspace-0001
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s SingleTenant) String() string {
	return dara.Prettify(s)
}

func (s SingleTenant) GoString() string {
	return s.String()
}

func (s *SingleTenant) GetDescription() *string {
	return s.Description
}

func (s *SingleTenant) GetStatus() *string {
	return s.Status
}

func (s *SingleTenant) GetTenantId() *string {
	return s.TenantId
}

func (s *SingleTenant) GetTenantName() *string {
	return s.TenantName
}

func (s *SingleTenant) GetUUID() *string {
	return s.UUID
}

func (s *SingleTenant) SetDescription(v string) *SingleTenant {
	s.Description = &v
	return s
}

func (s *SingleTenant) SetStatus(v string) *SingleTenant {
	s.Status = &v
	return s
}

func (s *SingleTenant) SetTenantId(v string) *SingleTenant {
	s.TenantId = &v
	return s
}

func (s *SingleTenant) SetTenantName(v string) *SingleTenant {
	s.TenantName = &v
	return s
}

func (s *SingleTenant) SetUUID(v string) *SingleTenant {
	s.UUID = &v
	return s
}

func (s *SingleTenant) Validate() error {
	return dara.Validate(s)
}
