// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePhysicalNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *PausePhysicalNodeShrinkRequest
	GetEnv() *string
	SetOpTenantId(v int64) *PausePhysicalNodeShrinkRequest
	GetOpTenantId() *int64
	SetPauseCommandShrink(v string) *PausePhysicalNodeShrinkRequest
	GetPauseCommandShrink() *string
}

type PausePhysicalNodeShrinkRequest struct {
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
	PauseCommandShrink *string `json:"PauseCommand,omitempty" xml:"PauseCommand,omitempty"`
}

func (s PausePhysicalNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PausePhysicalNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *PausePhysicalNodeShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PausePhysicalNodeShrinkRequest) GetPauseCommandShrink() *string {
	return s.PauseCommandShrink
}

func (s *PausePhysicalNodeShrinkRequest) SetEnv(v string) *PausePhysicalNodeShrinkRequest {
	s.Env = &v
	return s
}

func (s *PausePhysicalNodeShrinkRequest) SetOpTenantId(v int64) *PausePhysicalNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *PausePhysicalNodeShrinkRequest) SetPauseCommandShrink(v string) *PausePhysicalNodeShrinkRequest {
	s.PauseCommandShrink = &v
	return s
}

func (s *PausePhysicalNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
