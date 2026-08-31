// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetComputeClusterRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetComputeClusterRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetComputeClusterRequest
	GetOpUserId() *string
}

type GetComputeClusterRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s GetComputeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeClusterRequest) GoString() string {
	return s.String()
}

func (s *GetComputeClusterRequest) GetId() *int64 {
	return s.Id
}

func (s *GetComputeClusterRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetComputeClusterRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetComputeClusterRequest) SetId(v int64) *GetComputeClusterRequest {
	s.Id = &v
	return s
}

func (s *GetComputeClusterRequest) SetOpTenantId(v int64) *GetComputeClusterRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetComputeClusterRequest) SetOpUserId(v string) *GetComputeClusterRequest {
	s.OpUserId = &v
	return s
}

func (s *GetComputeClusterRequest) Validate() error {
	return dara.Validate(s)
}
