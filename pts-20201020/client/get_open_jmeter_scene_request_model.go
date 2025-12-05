// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetOpenJMeterSceneRequest
	GetSceneId() *string
}

type GetOpenJMeterSceneRequest struct {
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetOpenJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetOpenJMeterSceneRequest) SetSceneId(v string) *GetOpenJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *GetOpenJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
