// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardLookupTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardLookupTableShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardLookupTableShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardLookupTableShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardLookupTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardLookupTableShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardLookupTableShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardLookupTableShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardLookupTableShrinkRequest) SetOpTenantId(v int64) *CreateStandardLookupTableShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardLookupTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
