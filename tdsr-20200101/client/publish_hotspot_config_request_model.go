// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *PublishHotspotConfigRequest
	GetSceneId() *string
}

type PublishHotspotConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rteyauiiuw****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishHotspotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *PublishHotspotConfigRequest) SetSceneId(v string) *PublishHotspotConfigRequest {
	s.SceneId = &v
	return s
}

func (s *PublishHotspotConfigRequest) Validate() error {
	return dara.Validate(s)
}
