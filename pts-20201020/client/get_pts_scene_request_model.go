// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetPtsSceneRequest
	GetSceneId() *string
}

type GetPtsSceneRequest struct {
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// NKJBSH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneRequest) SetSceneId(v string) *GetPtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
