// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetEditingProjectMaterialsRequest
	GetProjectId() *string
}

type GetEditingProjectMaterialsRequest struct {
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// *****fb2101cb318*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectMaterialsRequest) SetProjectId(v string) *GetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
