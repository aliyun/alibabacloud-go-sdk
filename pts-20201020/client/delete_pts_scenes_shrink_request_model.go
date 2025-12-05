// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsScenesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneIdsShrink(v string) *DeletePtsScenesShrinkRequest
	GetSceneIdsShrink() *string
}

type DeletePtsScenesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["XVB4DF","AFG3CV"]
	SceneIdsShrink *string `json:"SceneIds,omitempty" xml:"SceneIds,omitempty"`
}

func (s DeletePtsScenesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsScenesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesShrinkRequest) GetSceneIdsShrink() *string {
	return s.SceneIdsShrink
}

func (s *DeletePtsScenesShrinkRequest) SetSceneIdsShrink(v string) *DeletePtsScenesShrinkRequest {
	s.SceneIdsShrink = &v
	return s
}

func (s *DeletePtsScenesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
