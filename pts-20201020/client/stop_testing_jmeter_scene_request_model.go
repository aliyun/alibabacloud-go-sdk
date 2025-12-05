// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTestingJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StopTestingJMeterSceneRequest
	GetSceneId() *string
}

type StopTestingJMeterSceneRequest struct {
	// The ID of the JMeter scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopTestingJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTestingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StopTestingJMeterSceneRequest) SetSceneId(v string) *StopTestingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StopTestingJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
