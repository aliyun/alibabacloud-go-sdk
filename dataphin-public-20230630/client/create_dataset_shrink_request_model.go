// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateDatasetShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateDatasetShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v string) *CreateDatasetShrinkRequest
	GetProjectId() *string
}

type CreateDatasetShrinkRequest struct {
	// The creation request.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
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
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateDatasetShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDatasetShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateDatasetShrinkRequest) SetCreateCommandShrink(v string) *CreateDatasetShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetOpTenantId(v int64) *CreateDatasetShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetProjectId(v string) *CreateDatasetShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
