// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteComputeResourceRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteComputeResourceRequest
	GetProjectId() *int64
}

type DeleteComputeResourceRequest struct {
	// The computing resource ID, the unique identifier of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteComputeResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteComputeResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteComputeResourceRequest) SetId(v int64) *DeleteComputeResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteComputeResourceRequest) SetProjectId(v int64) *DeleteComputeResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
