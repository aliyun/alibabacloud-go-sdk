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
	SetOpUserId(v string) *UpdateDatasetShrinkRequest
	GetOpUserId() *string
	SetProjectId(v string) *UpdateDatasetShrinkRequest
	GetProjectId() *string
	SetUpdateCommandShrink(v string) *UpdateDatasetShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDatasetShrinkRequest struct {
	// The tenant ID.
	//
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
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7273382541481536
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The update request struct.
	//
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

func (s *UpdateDatasetShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
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

func (s *UpdateDatasetShrinkRequest) SetOpUserId(v string) *UpdateDatasetShrinkRequest {
	s.OpUserId = &v
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
