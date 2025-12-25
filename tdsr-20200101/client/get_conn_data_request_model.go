// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetConnDataRequest
	GetSceneId() *string
}

type GetConnDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetConnDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnDataRequest) GoString() string {
	return s.String()
}

func (s *GetConnDataRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetConnDataRequest) SetSceneId(v string) *GetConnDataRequest {
	s.SceneId = &v
	return s
}

func (s *GetConnDataRequest) Validate() error {
	return dara.Validate(s)
}
