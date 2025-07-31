// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePhysicalNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ResumePhysicalNodeShrinkRequest
	GetEnv() *string
	SetOpTenantId(v int64) *ResumePhysicalNodeShrinkRequest
	GetOpTenantId() *int64
	SetResumeCommandShrink(v string) *ResumePhysicalNodeShrinkRequest
	GetResumeCommandShrink() *string
}

type ResumePhysicalNodeShrinkRequest struct {
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
	ResumeCommandShrink *string `json:"ResumeCommand,omitempty" xml:"ResumeCommand,omitempty"`
}

func (s ResumePhysicalNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ResumePhysicalNodeShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ResumePhysicalNodeShrinkRequest) GetResumeCommandShrink() *string {
	return s.ResumeCommandShrink
}

func (s *ResumePhysicalNodeShrinkRequest) SetEnv(v string) *ResumePhysicalNodeShrinkRequest {
	s.Env = &v
	return s
}

func (s *ResumePhysicalNodeShrinkRequest) SetOpTenantId(v int64) *ResumePhysicalNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResumePhysicalNodeShrinkRequest) SetResumeCommandShrink(v string) *ResumePhysicalNodeShrinkRequest {
	s.ResumeCommandShrink = &v
	return s
}

func (s *ResumePhysicalNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
