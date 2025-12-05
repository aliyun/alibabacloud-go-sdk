// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTestingJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StartTestingJMeterSceneRequest
	GetSceneId() *string
}

type StartTestingJMeterSceneRequest struct {
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartTestingJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTestingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StartTestingJMeterSceneRequest) SetSceneId(v string) *StartTestingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StartTestingJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
