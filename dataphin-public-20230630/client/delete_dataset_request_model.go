// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDatasetRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteDatasetRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *DeleteDatasetRequest
	GetProjectId() *int64
}

type DeleteDatasetRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7280840713415040
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7262419688152064
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDatasetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDatasetRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDatasetRequest) SetId(v int64) *DeleteDatasetRequest {
	s.Id = &v
	return s
}

func (s *DeleteDatasetRequest) SetOpTenantId(v int64) *DeleteDatasetRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDatasetRequest) SetProjectId(v int64) *DeleteDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
