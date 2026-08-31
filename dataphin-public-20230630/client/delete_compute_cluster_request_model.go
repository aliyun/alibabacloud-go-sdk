// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteComputeClusterRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteComputeClusterRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteComputeClusterRequest
	GetOpUserId() *string
}

type DeleteComputeClusterRequest struct {
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

func (s DeleteComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteComputeClusterRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteComputeClusterRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteComputeClusterRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteComputeClusterRequest) SetId(v int64) *DeleteComputeClusterRequest {
	s.Id = &v
	return s
}

func (s *DeleteComputeClusterRequest) SetOpTenantId(v int64) *DeleteComputeClusterRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteComputeClusterRequest) SetOpUserId(v string) *DeleteComputeClusterRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteComputeClusterRequest) Validate() error {
	return dara.Validate(s)
}
