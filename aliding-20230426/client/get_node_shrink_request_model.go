// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *GetNodeShrinkRequest
	GetNodeId() *string
	SetTenantContextShrink(v string) *GetNodeShrinkRequest
	GetTenantContextShrink() *string
	SetWithPermissionRole(v bool) *GetNodeShrinkRequest
	GetWithPermissionRole() *bool
	SetWithStatisticalInfo(v bool) *GetNodeShrinkRequest
	GetWithStatisticalInfo() *bool
}

type GetNodeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a9E05BDRVQ9K600yf1NplNDxV63zgkYA
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// true
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
	// example:
	//
	// true
	WithStatisticalInfo *bool `json:"WithStatisticalInfo,omitempty" xml:"WithStatisticalInfo,omitempty"`
}

func (s GetNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNodeShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetNodeShrinkRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetNodeShrinkRequest) GetWithStatisticalInfo() *bool {
	return s.WithStatisticalInfo
}

func (s *GetNodeShrinkRequest) SetNodeId(v string) *GetNodeShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeShrinkRequest) SetTenantContextShrink(v string) *GetNodeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetNodeShrinkRequest) SetWithPermissionRole(v bool) *GetNodeShrinkRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodeShrinkRequest) SetWithStatisticalInfo(v bool) *GetNodeShrinkRequest {
	s.WithStatisticalInfo = &v
	return s
}

func (s *GetNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
