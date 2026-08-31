// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeSupplementShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateNodeSupplementShrinkRequest
	GetCreateCommandShrink() *string
	SetEnv(v string) *CreateNodeSupplementShrinkRequest
	GetEnv() *string
	SetOpTenantId(v int64) *CreateNodeSupplementShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateNodeSupplementShrinkRequest
	GetOpUserId() *string
}

type CreateNodeSupplementShrinkRequest struct {
	// The data backfill request.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The environment identifier. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateNodeSupplementShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateNodeSupplementShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *CreateNodeSupplementShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateNodeSupplementShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateNodeSupplementShrinkRequest) SetCreateCommandShrink(v string) *CreateNodeSupplementShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) SetEnv(v string) *CreateNodeSupplementShrinkRequest {
	s.Env = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) SetOpTenantId(v int64) *CreateNodeSupplementShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) SetOpUserId(v string) *CreateNodeSupplementShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) Validate() error {
	return dara.Validate(s)
}
