// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *PublishStatusRequest
	GetSceneId() *string
}

type PublishStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishStatusRequest) GoString() string {
	return s.String()
}

func (s *PublishStatusRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *PublishStatusRequest) SetSceneId(v string) *PublishStatusRequest {
	s.SceneId = &v
	return s
}

func (s *PublishStatusRequest) Validate() error {
	return dara.Validate(s)
}
