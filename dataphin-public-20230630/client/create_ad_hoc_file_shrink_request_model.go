// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdHocFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateAdHocFileShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateAdHocFileShrinkRequest
	GetOpTenantId() *int64
}

type CreateAdHocFileShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateAdHocFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdHocFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateAdHocFileShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateAdHocFileShrinkRequest) SetCreateCommandShrink(v string) *CreateAdHocFileShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateAdHocFileShrinkRequest) SetOpTenantId(v int64) *CreateAdHocFileShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateAdHocFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
