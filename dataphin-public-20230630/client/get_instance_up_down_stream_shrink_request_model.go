// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUpDownStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetInstanceUpDownStreamShrinkRequest
	GetEnv() *string
	SetInstanceIdShrink(v string) *GetInstanceUpDownStreamShrinkRequest
	GetInstanceIdShrink() *string
	SetOpTenantId(v int64) *GetInstanceUpDownStreamShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetInstanceUpDownStreamShrinkRequest
	GetProjectId() *int64
	SetUpStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest
	GetUpStreamDepth() *int32
}

type GetInstanceUpDownStreamShrinkRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceIdShrink *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetInstanceUpDownStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetInstanceIdShrink() *string {
	return s.InstanceIdShrink
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceUpDownStreamShrinkRequest) GetUpStreamDepth() *int32 {
	return s.UpStreamDepth
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetEnv(v string) *GetInstanceUpDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetInstanceIdShrink(v string) *GetInstanceUpDownStreamShrinkRequest {
	s.InstanceIdShrink = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetOpTenantId(v int64) *GetInstanceUpDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetProjectId(v int64) *GetInstanceUpDownStreamShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetUpStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest {
	s.UpStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
