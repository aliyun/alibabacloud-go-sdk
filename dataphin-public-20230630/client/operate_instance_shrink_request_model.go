// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *OperateInstanceShrinkRequest
	GetEnv() *string
	SetOpTenantId(v int64) *OperateInstanceShrinkRequest
	GetOpTenantId() *int64
	SetOperateCommandShrink(v string) *OperateInstanceShrinkRequest
	GetOperateCommandShrink() *string
}

type OperateInstanceShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	OperateCommandShrink *string `json:"OperateCommand,omitempty" xml:"OperateCommand,omitempty"`
}

func (s OperateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateInstanceShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *OperateInstanceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OperateInstanceShrinkRequest) GetOperateCommandShrink() *string {
	return s.OperateCommandShrink
}

func (s *OperateInstanceShrinkRequest) SetEnv(v string) *OperateInstanceShrinkRequest {
	s.Env = &v
	return s
}

func (s *OperateInstanceShrinkRequest) SetOpTenantId(v int64) *OperateInstanceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OperateInstanceShrinkRequest) SetOperateCommandShrink(v string) *OperateInstanceShrinkRequest {
	s.OperateCommandShrink = &v
	return s
}

func (s *OperateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
