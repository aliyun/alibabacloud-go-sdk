// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *PackSourceRequest
	GetSceneId() *string
}

type PackSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// L2omaCMmQZZkEg4p****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PackSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PackSourceRequest) GoString() string {
	return s.String()
}

func (s *PackSourceRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *PackSourceRequest) SetSceneId(v string) *PackSourceRequest {
	s.SceneId = &v
	return s
}

func (s *PackSourceRequest) Validate() error {
	return dara.Validate(s)
}
