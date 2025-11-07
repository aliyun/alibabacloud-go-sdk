// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneConfigId(v int64) *DeleteSceneConfigRequest
	GetSceneConfigId() *int64
}

type DeleteSceneConfigRequest struct {
	// ID of the intent authentication configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 559
	SceneConfigId *int64 `json:"sceneConfigId,omitempty" xml:"sceneConfigId,omitempty"`
}

func (s DeleteSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneConfigRequest) GetSceneConfigId() *int64 {
	return s.SceneConfigId
}

func (s *DeleteSceneConfigRequest) SetSceneConfigId(v int64) *DeleteSceneConfigRequest {
	s.SceneConfigId = &v
	return s
}

func (s *DeleteSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
