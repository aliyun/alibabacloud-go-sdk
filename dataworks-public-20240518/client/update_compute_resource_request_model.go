// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionProperties(v string) *UpdateComputeResourceRequest
	GetConnectionProperties() *string
	SetConnectionPropertiesMode(v string) *UpdateComputeResourceRequest
	GetConnectionPropertiesMode() *string
	SetDescription(v string) *UpdateComputeResourceRequest
	GetDescription() *string
	SetId(v int64) *UpdateComputeResourceRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateComputeResourceRequest
	GetProjectId() *int64
}

type UpdateComputeResourceRequest struct {
	// This parameter is required.
	ConnectionProperties     *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *UpdateComputeResourceRequest) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *UpdateComputeResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateComputeResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateComputeResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateComputeResourceRequest) SetConnectionProperties(v string) *UpdateComputeResourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetConnectionPropertiesMode(v string) *UpdateComputeResourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetDescription(v string) *UpdateComputeResourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetId(v int64) *UpdateComputeResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetProjectId(v int64) *UpdateComputeResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
