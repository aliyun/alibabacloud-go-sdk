// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneAddOrUpdateSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScenesShrink(v string) *DsgSceneAddOrUpdateSceneShrinkRequest
	GetScenesShrink() *string
}

type DsgSceneAddOrUpdateSceneShrinkRequest struct {
	// The information about the level-2 data masking scenario.
	//
	// This parameter is required.
	ScenesShrink *string `json:"scenes,omitempty" xml:"scenes,omitempty"`
}

func (s DsgSceneAddOrUpdateSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneShrinkRequest) GetScenesShrink() *string {
	return s.ScenesShrink
}

func (s *DsgSceneAddOrUpdateSceneShrinkRequest) SetScenesShrink(v string) *DsgSceneAddOrUpdateSceneShrinkRequest {
	s.ScenesShrink = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
