// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfigShrink(v string) *CreateComputeClusterShrinkRequest
	GetClusterConfigShrink() *string
	SetOpTenantId(v int64) *CreateComputeClusterShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateComputeClusterShrinkRequest
	GetOpUserId() *string
}

type CreateComputeClusterShrinkRequest struct {
	// The cluster configuration.
	//
	// This parameter is required.
	ClusterConfigShrink *string `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty"`
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

func (s CreateComputeClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterShrinkRequest) GetClusterConfigShrink() *string {
	return s.ClusterConfigShrink
}

func (s *CreateComputeClusterShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateComputeClusterShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateComputeClusterShrinkRequest) SetClusterConfigShrink(v string) *CreateComputeClusterShrinkRequest {
	s.ClusterConfigShrink = &v
	return s
}

func (s *CreateComputeClusterShrinkRequest) SetOpTenantId(v int64) *CreateComputeClusterShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateComputeClusterShrinkRequest) SetOpUserId(v string) *CreateComputeClusterShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateComputeClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
