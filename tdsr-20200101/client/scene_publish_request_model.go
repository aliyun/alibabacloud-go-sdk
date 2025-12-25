// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScenePublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *ScenePublishRequest
	GetSceneId() *string
}

type ScenePublishRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ScenePublishRequest) String() string {
	return dara.Prettify(s)
}

func (s ScenePublishRequest) GoString() string {
	return s.String()
}

func (s *ScenePublishRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ScenePublishRequest) SetSceneId(v string) *ScenePublishRequest {
	s.SceneId = &v
	return s
}

func (s *ScenePublishRequest) Validate() error {
	return dara.Validate(s)
}
