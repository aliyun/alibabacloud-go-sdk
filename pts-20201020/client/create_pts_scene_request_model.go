// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v string) *CreatePtsSceneRequest
	GetScene() *string
}

type CreatePtsSceneRequest struct {
	// The scenario details.
	//
	// This parameter is required.
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s CreatePtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneRequest) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneRequest) GetScene() *string {
	return s.Scene
}

func (s *CreatePtsSceneRequest) SetScene(v string) *CreatePtsSceneRequest {
	s.Scene = &v
	return s
}

func (s *CreatePtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
