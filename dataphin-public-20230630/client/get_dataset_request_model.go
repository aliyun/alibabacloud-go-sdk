// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDatasetRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetDatasetRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetDatasetRequest
	GetProjectId() *int64
}

type GetDatasetRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7255018404425088
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
	// 7255013756724992
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDatasetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDatasetRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDatasetRequest) SetId(v int64) *GetDatasetRequest {
	s.Id = &v
	return s
}

func (s *GetDatasetRequest) SetOpTenantId(v int64) *GetDatasetRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDatasetRequest) SetProjectId(v int64) *GetDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDatasetRequest) Validate() error {
	return dara.Validate(s)
}
