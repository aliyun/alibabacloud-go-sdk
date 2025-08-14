// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialMaps(v string) *AddEditingProjectMaterialsRequest
	GetMaterialMaps() *string
	SetProjectId(v string) *AddEditingProjectMaterialsRequest
	GetProjectId() *string
}

type AddEditingProjectMaterialsRequest struct {
	// The material ID. Separate multiple material IDs with commas (,). Each type supports up to 10 material IDs. The following material types are supported:
	//
	// 	- video
	//
	// 	- audio
	//
	// 	- image
	//
	// 	- liveStream
	//
	// 	- editingProject
	//
	// This parameter is required.
	//
	// example:
	//
	// {"video":"*****2e057304fcd9b145c5cafc*****", "image":"****8021a8d493da643c8acd98*****,*****cb6307a4edea614d8b3f3c*****", "liveStream": "[{\\"appName\\":\\"testrecord\\",\\"domainName\\":\\"test.alivecdn.com\\",\\"liveUrl\\":\\"rtmp://test.alivecdn.com/testrecord/teststream\\",\\"streamName\\":\\"teststream\\"}]", "editingProject": "*****9b145c5cafc2e057304fcd*****"}
	MaterialMaps *string `json:"MaterialMaps,omitempty" xml:"MaterialMaps,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// *****b2101cb318c*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsRequest) GetMaterialMaps() *string {
	return s.MaterialMaps
}

func (s *AddEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddEditingProjectMaterialsRequest) SetMaterialMaps(v string) *AddEditingProjectMaterialsRequest {
	s.MaterialMaps = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetProjectId(v string) *AddEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
