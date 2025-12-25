// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRelativePositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelativePosition(v string) *AddRelativePositionRequest
	GetRelativePosition() *string
	SetSceneId(v string) *AddRelativePositionRequest
	GetSceneId() *string
}

type AddRelativePositionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"key"："value"}
	RelativePosition *string `json:"RelativePosition,omitempty" xml:"RelativePosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AddRelativePositionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRelativePositionRequest) GoString() string {
	return s.String()
}

func (s *AddRelativePositionRequest) GetRelativePosition() *string {
	return s.RelativePosition
}

func (s *AddRelativePositionRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AddRelativePositionRequest) SetRelativePosition(v string) *AddRelativePositionRequest {
	s.RelativePosition = &v
	return s
}

func (s *AddRelativePositionRequest) SetSceneId(v string) *AddRelativePositionRequest {
	s.SceneId = &v
	return s
}

func (s *AddRelativePositionRequest) Validate() error {
	return dara.Validate(s)
}
