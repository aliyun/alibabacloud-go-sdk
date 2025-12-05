// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StartPtsSceneRequest
	GetSceneId() *string
}

type StartPtsSceneRequest struct {
	// The ID of the scenario that you want to start, which is the ID that is returned after the scenario is created. You can view scenario IDs on the scenario list page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FGSRA3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StartPtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StartPtsSceneRequest) SetSceneId(v string) *StartPtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StartPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
