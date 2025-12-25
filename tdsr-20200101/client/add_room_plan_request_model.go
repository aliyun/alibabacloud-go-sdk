// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRoomPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *AddRoomPlanRequest
	GetSceneId() *string
}

type AddRoomPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AddRoomPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRoomPlanRequest) GoString() string {
	return s.String()
}

func (s *AddRoomPlanRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AddRoomPlanRequest) SetSceneId(v string) *AddRoomPlanRequest {
	s.SceneId = &v
	return s
}

func (s *AddRoomPlanRequest) Validate() error {
	return dara.Validate(s)
}
