// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveModelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SaveModelConfigRequest
	GetData() *string
	SetSceneId(v string) *SaveModelConfigRequest
	GetSceneId() *string
}

type SaveModelConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iyweyteyue****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveModelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveModelConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveModelConfigRequest) GetData() *string {
	return s.Data
}

func (s *SaveModelConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *SaveModelConfigRequest) SetData(v string) *SaveModelConfigRequest {
	s.Data = &v
	return s
}

func (s *SaveModelConfigRequest) SetSceneId(v string) *SaveModelConfigRequest {
	s.SceneId = &v
	return s
}

func (s *SaveModelConfigRequest) Validate() error {
	return dara.Validate(s)
}
