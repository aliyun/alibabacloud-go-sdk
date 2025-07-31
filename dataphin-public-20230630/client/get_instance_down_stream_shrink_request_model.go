// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDownStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetInstanceDownStreamShrinkRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetInstanceDownStreamShrinkRequest
	GetEnv() *string
	SetInstanceGetShrink(v string) *GetInstanceDownStreamShrinkRequest
	GetInstanceGetShrink() *string
	SetOpTenantId(v int64) *GetInstanceDownStreamShrinkRequest
	GetOpTenantId() *int64
	SetRunStatus(v string) *GetInstanceDownStreamShrinkRequest
	GetRunStatus() *string
}

type GetInstanceDownStreamShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceGetShrink *string `json:"InstanceGet,omitempty" xml:"InstanceGet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
}

func (s GetInstanceDownStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamShrinkRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetInstanceDownStreamShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *GetInstanceDownStreamShrinkRequest) GetInstanceGetShrink() *string {
	return s.InstanceGetShrink
}

func (s *GetInstanceDownStreamShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetInstanceDownStreamShrinkRequest) GetRunStatus() *string {
	return s.RunStatus
}

func (s *GetInstanceDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetInstanceDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetEnv(v string) *GetInstanceDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetInstanceGetShrink(v string) *GetInstanceDownStreamShrinkRequest {
	s.InstanceGetShrink = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetOpTenantId(v int64) *GetInstanceDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetRunStatus(v string) *GetInstanceDownStreamShrinkRequest {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
