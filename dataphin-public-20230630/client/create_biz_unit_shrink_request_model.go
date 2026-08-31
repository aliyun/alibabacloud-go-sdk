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
	SetOpUserId(v string) *CreateBizUnitShrinkRequest
	GetOpUserId() *string
}

type CreateBizUnitShrinkRequest struct {
	// The create request.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *CreateBizUnitShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateBizUnitShrinkRequest) SetCreateCommandShrink(v string) *CreateBizUnitShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBizUnitShrinkRequest) SetOpTenantId(v int64) *CreateBizUnitShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizUnitShrinkRequest) SetOpUserId(v string) *CreateBizUnitShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBizUnitShrinkRequest) Validate() error {
	return dara.Validate(s)
}
