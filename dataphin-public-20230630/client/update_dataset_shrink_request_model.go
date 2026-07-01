// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDatasetShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v string) *UpdateDatasetShrinkRequest
	GetProjectId() *string
	SetUpdateCommandShrink(v string) *UpdateDatasetShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDatasetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7273382541481536
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDatasetShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateDatasetShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDatasetShrinkRequest) SetOpTenantId(v int64) *UpdateDatasetShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetProjectId(v string) *UpdateDatasetShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDatasetShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
