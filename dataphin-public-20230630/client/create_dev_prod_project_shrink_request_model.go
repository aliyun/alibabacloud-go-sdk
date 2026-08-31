// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDevProdProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDevProdProjectShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDevProdProjectShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateDevProdProjectShrinkRequest
	GetOpUserId() *string
}

type CreateDevProdProjectShrinkRequest struct {
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
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateDevProdProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDevProdProjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDevProdProjectShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateDevProdProjectShrinkRequest) SetCreateCommandShrink(v string) *CreateDevProdProjectShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDevProdProjectShrinkRequest) SetOpTenantId(v int64) *CreateDevProdProjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDevProdProjectShrinkRequest) SetOpUserId(v string) *CreateDevProdProjectShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateDevProdProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
