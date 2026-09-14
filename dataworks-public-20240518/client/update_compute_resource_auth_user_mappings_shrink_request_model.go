// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceAuthUserMappingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResourceId(v int64) *UpdateComputeResourceAuthUserMappingsShrinkRequest
	GetComputeResourceId() *int64
	SetProjectId(v int64) *UpdateComputeResourceAuthUserMappingsShrinkRequest
	GetProjectId() *int64
	SetRemoveUserIdsShrink(v string) *UpdateComputeResourceAuthUserMappingsShrinkRequest
	GetRemoveUserIdsShrink() *string
	SetUpsertsShrink(v string) *UpdateComputeResourceAuthUserMappingsShrinkRequest
	GetUpsertsShrink() *string
}

type UpdateComputeResourceAuthUserMappingsShrinkRequest struct {
	// The compute resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123455
	ComputeResourceId *int64 `json:"ComputeResourceId,omitempty" xml:"ComputeResourceId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of user mappings to remove.
	RemoveUserIdsShrink *string `json:"RemoveUserIds,omitempty" xml:"RemoveUserIds,omitempty"`
	// The list of objects to update.
	UpsertsShrink *string `json:"Upserts,omitempty" xml:"Upserts,omitempty"`
}

func (s UpdateComputeResourceAuthUserMappingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) GetComputeResourceId() *int64 {
	return s.ComputeResourceId
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) GetRemoveUserIdsShrink() *string {
	return s.RemoveUserIdsShrink
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) GetUpsertsShrink() *string {
	return s.UpsertsShrink
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) SetComputeResourceId(v int64) *UpdateComputeResourceAuthUserMappingsShrinkRequest {
	s.ComputeResourceId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) SetProjectId(v int64) *UpdateComputeResourceAuthUserMappingsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) SetRemoveUserIdsShrink(v string) *UpdateComputeResourceAuthUserMappingsShrinkRequest {
	s.RemoveUserIdsShrink = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) SetUpsertsShrink(v string) *UpdateComputeResourceAuthUserMappingsShrinkRequest {
	s.UpsertsShrink = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
