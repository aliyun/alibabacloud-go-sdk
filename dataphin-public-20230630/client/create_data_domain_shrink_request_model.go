// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDataDomainShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDataDomainShrinkRequest
	GetOpTenantId() *int64
}

type CreateDataDomainShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataDomainShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDataDomainShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataDomainShrinkRequest) SetCreateCommandShrink(v string) *CreateDataDomainShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataDomainShrinkRequest) SetOpTenantId(v int64) *CreateDataDomainShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
