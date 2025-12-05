// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StartDebugPtsSceneRequest
	GetSceneId() *string
}

type StartDebugPtsSceneRequest struct {
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// NHBGB8B
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartDebugPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDebugPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StartDebugPtsSceneRequest) SetSceneId(v string) *StartDebugPtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StartDebugPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
