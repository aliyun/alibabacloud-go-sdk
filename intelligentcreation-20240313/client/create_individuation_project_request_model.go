// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectInfo(v string) *CreateIndividuationProjectRequest
	GetProjectInfo() *string
	SetProjectName(v string) *CreateIndividuationProjectRequest
	GetProjectName() *string
	SetPurpose(v string) *CreateIndividuationProjectRequest
	GetPurpose() *string
	SetSceneId(v string) *CreateIndividuationProjectRequest
	GetSceneId() *string
}

type CreateIndividuationProjectRequest struct {
	ProjectInfo *string `json:"projectInfo,omitempty" xml:"projectInfo,omitempty"`
	// example:
	//
	// avatar-1
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Purpose     *string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// example:
	//
	// ail003
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
}

func (s CreateIndividuationProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectRequest) GetProjectInfo() *string {
	return s.ProjectInfo
}

func (s *CreateIndividuationProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateIndividuationProjectRequest) GetPurpose() *string {
	return s.Purpose
}

func (s *CreateIndividuationProjectRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateIndividuationProjectRequest) SetProjectInfo(v string) *CreateIndividuationProjectRequest {
	s.ProjectInfo = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetProjectName(v string) *CreateIndividuationProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetPurpose(v string) *CreateIndividuationProjectRequest {
	s.Purpose = &v
	return s
}

func (s *CreateIndividuationProjectRequest) SetSceneId(v string) *CreateIndividuationProjectRequest {
	s.SceneId = &v
	return s
}

func (s *CreateIndividuationProjectRequest) Validate() error {
	return dara.Validate(s)
}
