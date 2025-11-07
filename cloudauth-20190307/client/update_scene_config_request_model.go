// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateSceneConfigRequest
	GetConfig() *string
	SetId(v int64) *UpdateSceneConfigRequest
	GetId() *int64
	SetSceneId(v int64) *UpdateSceneConfigRequest
	GetSceneId() *int64
}

type UpdateSceneConfigRequest struct {
	// Scene configuration information, in JSON format. For the specific structure definition, please refer to more information about the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"faceCompareMode\\":\\"AUTHORITY\\",\\"certConfigs\\":[{\\"index\\":0,\\"model\\":\\"ENROLL\\"}],\\"screenEvidence\\":false}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Willingness configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 607
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Selected authentication scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000012918
	SceneId *int64 `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
}

func (s UpdateSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateSceneConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSceneConfigRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *UpdateSceneConfigRequest) SetConfig(v string) *UpdateSceneConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateSceneConfigRequest) SetId(v int64) *UpdateSceneConfigRequest {
	s.Id = &v
	return s
}

func (s *UpdateSceneConfigRequest) SetSceneId(v int64) *UpdateSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
