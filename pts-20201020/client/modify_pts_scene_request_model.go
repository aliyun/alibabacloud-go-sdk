// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v string) *ModifyPtsSceneRequest
	GetScene() *string
}

type ModifyPtsSceneRequest struct {
	// null
	//
	// This parameter is required.
	//
	// example:
	//
	// SD6YZCI
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ModifyPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneRequest) GetScene() *string {
	return s.Scene
}

func (s *ModifyPtsSceneRequest) SetScene(v string) *ModifyPtsSceneRequest {
	s.Scene = &v
	return s
}

func (s *ModifyPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
