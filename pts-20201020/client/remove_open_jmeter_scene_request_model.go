// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOpenJMeterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *RemoveOpenJMeterSceneRequest
	GetSceneId() *string
}

type RemoveOpenJMeterSceneRequest struct {
	// The ID of the scenario that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s RemoveOpenJMeterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *RemoveOpenJMeterSceneRequest) SetSceneId(v string) *RemoveOpenJMeterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *RemoveOpenJMeterSceneRequest) Validate() error {
	return dara.Validate(s)
}
