// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneBaseLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *DeletePtsSceneBaseLineRequest
	GetSceneId() *string
}

type DeletePtsSceneBaseLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NHGV4CDG
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeletePtsSceneBaseLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeletePtsSceneBaseLineRequest) SetSceneId(v string) *DeletePtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

func (s *DeletePtsSceneBaseLineRequest) Validate() error {
	return dara.Validate(s)
}
