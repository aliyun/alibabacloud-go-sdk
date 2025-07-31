// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateBizEntityShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateBizEntityShrinkRequest
	GetOpTenantId() *int64
}

type CreateBizEntityShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBizEntityShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateBizEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizEntityShrinkRequest) SetCreateCommandShrink(v string) *CreateBizEntityShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBizEntityShrinkRequest) SetOpTenantId(v int64) *CreateBizEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
