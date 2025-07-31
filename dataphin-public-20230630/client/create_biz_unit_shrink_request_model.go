// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizUnitShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateBizUnitShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateBizUnitShrinkRequest
	GetOpTenantId() *int64
}

type CreateBizUnitShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizUnitShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBizUnitShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateBizUnitShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizUnitShrinkRequest) SetCreateCommandShrink(v string) *CreateBizUnitShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBizUnitShrinkRequest) SetOpTenantId(v int64) *CreateBizUnitShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizUnitShrinkRequest) Validate() error {
	return dara.Validate(s)
}
