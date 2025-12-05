// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *DeletePtsSceneRequest
	GetSceneId() *string
}

type DeletePtsSceneRequest struct {
	// The ID of the PTS scenario that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// XANH3H
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeletePtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeletePtsSceneRequest) SetSceneId(v string) *DeletePtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *DeletePtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
