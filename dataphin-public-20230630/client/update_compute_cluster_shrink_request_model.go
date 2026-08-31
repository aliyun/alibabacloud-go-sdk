// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfigShrink(v string) *UpdateComputeClusterShrinkRequest
	GetClusterConfigShrink() *string
	SetId(v int64) *UpdateComputeClusterShrinkRequest
	GetId() *int64
	SetOpTenantId(v int64) *UpdateComputeClusterShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateComputeClusterShrinkRequest
	GetOpUserId() *string
}

type UpdateComputeClusterShrinkRequest struct {
	// This parameter is required.
	ClusterConfigShrink *string `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s UpdateComputeClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterShrinkRequest) GetClusterConfigShrink() *string {
	return s.ClusterConfigShrink
}

func (s *UpdateComputeClusterShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateComputeClusterShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateComputeClusterShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateComputeClusterShrinkRequest) SetClusterConfigShrink(v string) *UpdateComputeClusterShrinkRequest {
	s.ClusterConfigShrink = &v
	return s
}

func (s *UpdateComputeClusterShrinkRequest) SetId(v int64) *UpdateComputeClusterShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateComputeClusterShrinkRequest) SetOpTenantId(v int64) *UpdateComputeClusterShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateComputeClusterShrinkRequest) SetOpUserId(v string) *UpdateComputeClusterShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateComputeClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
