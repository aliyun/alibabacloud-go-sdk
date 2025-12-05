// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebuggingJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *StartDebuggingJMeterSceneRequest
	GetSceneId() *string
}

type StartDebuggingJMeterSceneRequest struct {
	// The ID of the scenario that you want to debug.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartDebuggingJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDebuggingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StartDebuggingJMeterSceneRequest) SetSceneId(v string) *StartDebuggingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StartDebuggingJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
