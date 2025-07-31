// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeUpDownStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetNodeUpDownStreamShrinkRequest
	GetEnv() *string
	SetNodeIdShrink(v string) *GetNodeUpDownStreamShrinkRequest
	GetNodeIdShrink() *string
	SetOpTenantId(v int64) *GetNodeUpDownStreamShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetNodeUpDownStreamShrinkRequest
	GetProjectId() *int64
	SetUpStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest
	GetUpStreamDepth() *int32
}

type GetNodeUpDownStreamShrinkRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	NodeIdShrink *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 113123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetNodeUpDownStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamShrinkRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetNodeUpDownStreamShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *GetNodeUpDownStreamShrinkRequest) GetNodeIdShrink() *string {
	return s.NodeIdShrink
}

func (s *GetNodeUpDownStreamShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetNodeUpDownStreamShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeUpDownStreamShrinkRequest) GetUpStreamDepth() *int32 {
	return s.UpStreamDepth
}

func (s *GetNodeUpDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetEnv(v string) *GetNodeUpDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetNodeIdShrink(v string) *GetNodeUpDownStreamShrinkRequest {
	s.NodeIdShrink = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetOpTenantId(v int64) *GetNodeUpDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetProjectId(v int64) *GetNodeUpDownStreamShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetUpStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest {
	s.UpStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
