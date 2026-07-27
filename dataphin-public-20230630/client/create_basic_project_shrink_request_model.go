// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateBasicProjectShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateBasicProjectShrinkRequest
	GetOpTenantId() *int64
}

type CreateBasicProjectShrinkRequest struct {
	// The create command.
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
}

func (s CreateBasicProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateBasicProjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBasicProjectShrinkRequest) SetCreateCommandShrink(v string) *CreateBasicProjectShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBasicProjectShrinkRequest) SetOpTenantId(v int64) *CreateBasicProjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBasicProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
