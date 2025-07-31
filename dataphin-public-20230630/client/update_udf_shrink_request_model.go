// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateUdfShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateUdfShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateUdfShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateUdfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUdfShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateUdfShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateUdfShrinkRequest) SetOpTenantId(v int64) *UpdateUdfShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUdfShrinkRequest) SetUpdateCommandShrink(v string) *UpdateUdfShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateUdfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
