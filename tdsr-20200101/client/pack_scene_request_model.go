// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *PackSceneRequest
	GetSceneId() *string
	SetType(v string) *PackSceneRequest
	GetType() *string
}

type PackSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// L2omaCMmQZZkEg4p****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// download
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PackSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s PackSceneRequest) GoString() string {
	return s.String()
}

func (s *PackSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *PackSceneRequest) GetType() *string {
	return s.Type
}

func (s *PackSceneRequest) SetSceneId(v string) *PackSceneRequest {
	s.SceneId = &v
	return s
}

func (s *PackSceneRequest) SetType(v string) *PackSceneRequest {
	s.Type = &v
	return s
}

func (s *PackSceneRequest) Validate() error {
	return dara.Validate(s)
}
