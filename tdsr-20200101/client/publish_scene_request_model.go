// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *PublishSceneRequest
	GetSceneId() *string
}

type PublishSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishSceneRequest) GoString() string {
	return s.String()
}

func (s *PublishSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *PublishSceneRequest) SetSceneId(v string) *PublishSceneRequest {
	s.SceneId = &v
	return s
}

func (s *PublishSceneRequest) Validate() error {
	return dara.Validate(s)
}
