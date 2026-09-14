// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceAuthUserMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResourceId(v int64) *GetComputeResourceAuthUserMappingsRequest
	GetComputeResourceId() *int64
	SetProjectId(v int64) *GetComputeResourceAuthUserMappingsRequest
	GetProjectId() *int64
}

type GetComputeResourceAuthUserMappingsRequest struct {
	// The ID of the compute resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000
	ComputeResourceId *int64 `json:"ComputeResourceId,omitempty" xml:"ComputeResourceId,omitempty"`
	// The DataWorks workspace to which the data source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetComputeResourceAuthUserMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceAuthUserMappingsRequest) GoString() string {
	return s.String()
}

func (s *GetComputeResourceAuthUserMappingsRequest) GetComputeResourceId() *int64 {
	return s.ComputeResourceId
}

func (s *GetComputeResourceAuthUserMappingsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetComputeResourceAuthUserMappingsRequest) SetComputeResourceId(v int64) *GetComputeResourceAuthUserMappingsRequest {
	s.ComputeResourceId = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsRequest) SetProjectId(v int64) *GetComputeResourceAuthUserMappingsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsRequest) Validate() error {
	return dara.Validate(s)
}
