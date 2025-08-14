// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialIds(v string) *DeleteEditingProjectMaterialsRequest
	GetMaterialIds() *string
	SetMaterialType(v string) *DeleteEditingProjectMaterialsRequest
	GetMaterialType() *string
	SetProjectId(v string) *DeleteEditingProjectMaterialsRequest
	GetProjectId() *string
}

type DeleteEditingProjectMaterialsRequest struct {
	// The material ID. Separate multiple material IDs with commas (,). You can specify up to 10 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// *****cbd721b418a89a7dafb1dc*****,*****86f5d534c95997c55c96f*****
	MaterialIds *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	// The material type. Valid values:
	//
	// \\- video
	//
	// \\- image
	//
	// \\- audio
	//
	// \\- subtitle
	//
	// \\- text
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// *****fb2101cb318*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsRequest) GetMaterialIds() *string {
	return s.MaterialIds
}

func (s *DeleteEditingProjectMaterialsRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *DeleteEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialIds(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialType(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetProjectId(v string) *DeleteEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
