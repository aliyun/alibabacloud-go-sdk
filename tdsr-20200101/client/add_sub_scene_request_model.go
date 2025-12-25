// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddSubSceneRequest
	GetName() *string
	SetSceneId(v string) *AddSubSceneRequest
	GetSceneId() *string
	SetUploadType(v string) *AddSubSceneRequest
	GetUploadType() *string
}

type AddSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// IMAGE
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s AddSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSubSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSubSceneRequest) GetName() *string {
	return s.Name
}

func (s *AddSubSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *AddSubSceneRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *AddSubSceneRequest) SetName(v string) *AddSubSceneRequest {
	s.Name = &v
	return s
}

func (s *AddSubSceneRequest) SetSceneId(v string) *AddSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *AddSubSceneRequest) SetUploadType(v string) *AddSubSceneRequest {
	s.UploadType = &v
	return s
}

func (s *AddSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
