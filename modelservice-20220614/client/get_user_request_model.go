// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneType(v string) *GetUserRequest
	GetSceneType() *string
}

type GetUserRequest struct {
	// example:
	//
	// image_label
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *GetUserRequest) SetSceneType(v string) *GetUserRequest {
	s.SceneType = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
