// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardLookupTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardLookupTableShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateStandardLookupTableShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardLookupTableShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardLookupTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardLookupTableShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardLookupTableShrinkRequest) SetOpTenantId(v int64) *UpdateStandardLookupTableShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardLookupTableShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardLookupTableShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardLookupTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
