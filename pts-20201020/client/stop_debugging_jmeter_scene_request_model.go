// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebuggingJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StopDebuggingJMeterSceneRequest
	GetSceneId() *string
}

type StopDebuggingJMeterSceneRequest struct {
	// The ID of the scenario that you want to stop debugging.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopDebuggingJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDebuggingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StopDebuggingJMeterSceneRequest) SetSceneId(v string) *StopDebuggingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StopDebuggingJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
