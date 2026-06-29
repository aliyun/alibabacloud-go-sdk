// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTenantRequest
	GetDescription() *string
	SetTenantName(v string) *UpdateTenantRequest
	GetTenantName() *string
}

type UpdateTenantRequest struct {
	// Tenant description.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Tenant name.
	//
	// example:
	//
	// 测试任务202208101424
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s UpdateTenantRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTenantRequest) GetTenantName() *string {
	return s.TenantName
}

func (s *UpdateTenantRequest) SetDescription(v string) *UpdateTenantRequest {
	s.Description = &v
	return s
}

func (s *UpdateTenantRequest) SetTenantName(v string) *UpdateTenantRequest {
	s.TenantName = &v
	return s
}

func (s *UpdateTenantRequest) Validate() error {
	return dara.Validate(s)
}
