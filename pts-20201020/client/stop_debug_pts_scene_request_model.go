// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebugPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *StopDebugPtsSceneRequest
	GetPlanId() *string
	SetSceneId(v string) *StopDebugPtsSceneRequest
	GetSceneId() *string
}

type StopDebugPtsSceneRequest struct {
	// The ID of the stress testing task.
	//
	// This parameter is required.
	//
	// example:
	//
	// FVDC7HB
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDDCF7
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopDebugPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDebugPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *StopDebugPtsSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StopDebugPtsSceneRequest) SetPlanId(v string) *StopDebugPtsSceneRequest {
	s.PlanId = &v
	return s
}

func (s *StopDebugPtsSceneRequest) SetSceneId(v string) *StopDebugPtsSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StopDebugPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
