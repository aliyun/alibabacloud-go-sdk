// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneBuildTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetSceneBuildTaskStatusRequest
	GetSceneId() *string
}

type GetSceneBuildTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetSceneBuildTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSceneBuildTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSceneBuildTaskStatusRequest) SetSceneId(v string) *GetSceneBuildTaskStatusRequest {
	s.SceneId = &v
	return s
}

func (s *GetSceneBuildTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
