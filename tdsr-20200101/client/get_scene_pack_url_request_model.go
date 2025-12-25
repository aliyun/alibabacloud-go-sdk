// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePackUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetScenePackUrlRequest
	GetSceneId() *string
}

type GetScenePackUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// L2omaCMmQZZkEg4p****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetScenePackUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenePackUrlRequest) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetScenePackUrlRequest) SetSceneId(v string) *GetScenePackUrlRequest {
	s.SceneId = &v
	return s
}

func (s *GetScenePackUrlRequest) Validate() error {
	return dara.Validate(s)
}
