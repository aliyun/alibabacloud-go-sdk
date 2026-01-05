// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetComputeResourceRequest
	GetId() *int64
	SetProjectId(v int64) *GetComputeResourceRequest
	GetProjectId() *int64
}

type GetComputeResourceRequest struct {
	// The computing resource ID, which is the unique identifier for the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11792
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10003
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *GetComputeResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetComputeResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetComputeResourceRequest) SetId(v int64) *GetComputeResourceRequest {
	s.Id = &v
	return s
}

func (s *GetComputeResourceRequest) SetProjectId(v int64) *GetComputeResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *GetComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
