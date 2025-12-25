// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMinimapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SaveMinimapRequest
	GetData() *string
	SetSceneId(v string) *SaveMinimapRequest
	GetSceneId() *string
}

type SaveMinimapRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// This parameter is required.
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveMinimapRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveMinimapRequest) GoString() string {
	return s.String()
}

func (s *SaveMinimapRequest) GetData() *string {
	return s.Data
}

func (s *SaveMinimapRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *SaveMinimapRequest) SetData(v string) *SaveMinimapRequest {
	s.Data = &v
	return s
}

func (s *SaveMinimapRequest) SetSceneId(v string) *SaveMinimapRequest {
	s.SceneId = &v
	return s
}

func (s *SaveMinimapRequest) Validate() error {
	return dara.Validate(s)
}
