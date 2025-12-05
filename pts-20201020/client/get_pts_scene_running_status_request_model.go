// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetPtsSceneRunningStatusRequest
	GetSceneId() *string
}

type GetPtsSceneRunningStatusRequest struct {
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// NHBG6V
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRunningStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneRunningStatusRequest) SetSceneId(v string) *GetPtsSceneRunningStatusRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneRunningStatusRequest) Validate() error {
	return dara.Validate(s)
}
