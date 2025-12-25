// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CopySceneRequest
	GetProjectId() *string
	SetSceneId(v string) *CopySceneRequest
	GetSceneId() *string
	SetSceneName(v string) *CopySceneRequest
	GetSceneName() *string
}

type CopySceneRequest struct {
	// example:
	//
	// opwuoieywtyqw****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sgyuyewyew****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s CopySceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySceneRequest) GoString() string {
	return s.String()
}

func (s *CopySceneRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CopySceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CopySceneRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *CopySceneRequest) SetProjectId(v string) *CopySceneRequest {
	s.ProjectId = &v
	return s
}

func (s *CopySceneRequest) SetSceneId(v string) *CopySceneRequest {
	s.SceneId = &v
	return s
}

func (s *CopySceneRequest) SetSceneName(v string) *CopySceneRequest {
	s.SceneName = &v
	return s
}

func (s *CopySceneRequest) Validate() error {
	return dara.Validate(s)
}
