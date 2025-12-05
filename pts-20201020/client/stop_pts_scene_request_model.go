// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StopPtsSceneRequest
	GetSceneId() *string
}

type StopPtsSceneRequest struct {
	// The ID of the scenario that you want to stop, which is the ID that is returned after the scenario is created. You can view scenario IDs on the scenario list page in the PTS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// GV4DEBG
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StopPtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StopPtsSceneRequest) SetSceneId(v string) *StopPtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StopPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
