// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateSceneConfigRequest
	GetConfig() *string
	SetSceneId(v int64) *CreateSceneConfigRequest
	GetSceneId() *int64
	SetType(v string) *CreateSceneConfigRequest
	GetType() *string
}

type CreateSceneConfigRequest struct {
	// Intention authentication configuration, as a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"faceCompareMode\\":\\"AUTHORITY\\",\\"certConfigs\\":[{\\"index\\":0,\\"openVoiceCompare\\":true,\\"openCustomizedContent\\":true,\\"model\\":\\"FOLLOW\\"}],\\"screenEvidence\\":false}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Scene ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000014084
	SceneId *int64 `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// Configuration type.
	//
	// This parameter is required.
	//
	// example:
	//
	// VOLUNTARY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateSceneConfigRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *CreateSceneConfigRequest) GetType() *string {
	return s.Type
}

func (s *CreateSceneConfigRequest) SetConfig(v string) *CreateSceneConfigRequest {
	s.Config = &v
	return s
}

func (s *CreateSceneConfigRequest) SetSceneId(v int64) *CreateSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *CreateSceneConfigRequest) SetType(v string) *CreateSceneConfigRequest {
	s.Type = &v
	return s
}

func (s *CreateSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
