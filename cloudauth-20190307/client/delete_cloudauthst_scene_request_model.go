// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudauthstSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *DeleteCloudauthstSceneRequest
	GetSceneId() *string
}

type DeleteCloudauthstSceneRequest struct {
	// ID of the scene to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4275001
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteCloudauthstSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudauthstSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudauthstSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeleteCloudauthstSceneRequest) SetSceneId(v string) *DeleteCloudauthstSceneRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteCloudauthstSceneRequest) Validate() error {
	return dara.Validate(s)
}
