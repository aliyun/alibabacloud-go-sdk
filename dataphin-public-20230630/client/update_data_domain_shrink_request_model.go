// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataDomainShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateDataDomainShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataDomainShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataDomainShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataDomainShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataDomainShrinkRequest) SetOpTenantId(v int64) *UpdateDataDomainShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataDomainShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataDomainShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
