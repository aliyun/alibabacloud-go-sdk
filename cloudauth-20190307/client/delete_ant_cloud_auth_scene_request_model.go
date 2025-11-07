// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntCloudAuthSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v int64) *DeleteAntCloudAuthSceneRequest
	GetSceneId() *int64
}

type DeleteAntCloudAuthSceneRequest struct {
	// ID of the scene to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000011589
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteAntCloudAuthSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntCloudAuthSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntCloudAuthSceneRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DeleteAntCloudAuthSceneRequest) SetSceneId(v int64) *DeleteAntCloudAuthSceneRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteAntCloudAuthSceneRequest) Validate() error {
	return dara.Validate(s)
}
