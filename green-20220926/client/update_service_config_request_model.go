// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileConfig(v string) *UpdateServiceConfigRequest
	GetFileConfig() *string
	SetKeywordFilterLibs(v string) *UpdateServiceConfigRequest
	GetKeywordFilterLibs() *string
	SetKeywordHitLibs(v string) *UpdateServiceConfigRequest
	GetKeywordHitLibs() *string
	SetManualMachineConfig(v string) *UpdateServiceConfigRequest
	GetManualMachineConfig() *string
	SetRegionId(v string) *UpdateServiceConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *UpdateServiceConfigRequest
	GetResourceType() *string
	SetScene(v string) *UpdateServiceConfigRequest
	GetScene() *string
	SetSceneConfig(v string) *UpdateServiceConfigRequest
	GetSceneConfig() *string
	SetServiceCode(v string) *UpdateServiceConfigRequest
	GetServiceCode() *string
	SetServiceConfig(v string) *UpdateServiceConfigRequest
	GetServiceConfig() *string
	SetVideoConfig(v string) *UpdateServiceConfigRequest
	GetVideoConfig() *string
}

type UpdateServiceConfigRequest struct {
	// example:
	//
	// {}
	FileConfig *string `json:"FileConfig,omitempty" xml:"FileConfig,omitempty"`
	// example:
	//
	// []
	KeywordFilterLibs *string `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty"`
	// example:
	//
	// []
	KeywordHitLibs      *string `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty"`
	ManualMachineConfig *string `json:"ManualMachineConfig,omitempty" xml:"ManualMachineConfig,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// pornographic
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// {}
	SceneConfig *string `json:"SceneConfig,omitempty" xml:"SceneConfig,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode   *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// example:
	//
	// {}
	VideoConfig *string `json:"VideoConfig,omitempty" xml:"VideoConfig,omitempty"`
}

func (s UpdateServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigRequest) GetFileConfig() *string {
	return s.FileConfig
}

func (s *UpdateServiceConfigRequest) GetKeywordFilterLibs() *string {
	return s.KeywordFilterLibs
}

func (s *UpdateServiceConfigRequest) GetKeywordHitLibs() *string {
	return s.KeywordHitLibs
}

func (s *UpdateServiceConfigRequest) GetManualMachineConfig() *string {
	return s.ManualMachineConfig
}

func (s *UpdateServiceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateServiceConfigRequest) GetScene() *string {
	return s.Scene
}

func (s *UpdateServiceConfigRequest) GetSceneConfig() *string {
	return s.SceneConfig
}

func (s *UpdateServiceConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *UpdateServiceConfigRequest) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *UpdateServiceConfigRequest) GetVideoConfig() *string {
	return s.VideoConfig
}

func (s *UpdateServiceConfigRequest) SetFileConfig(v string) *UpdateServiceConfigRequest {
	s.FileConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetKeywordFilterLibs(v string) *UpdateServiceConfigRequest {
	s.KeywordFilterLibs = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetKeywordHitLibs(v string) *UpdateServiceConfigRequest {
	s.KeywordHitLibs = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetManualMachineConfig(v string) *UpdateServiceConfigRequest {
	s.ManualMachineConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetRegionId(v string) *UpdateServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetResourceType(v string) *UpdateServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetScene(v string) *UpdateServiceConfigRequest {
	s.Scene = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetSceneConfig(v string) *UpdateServiceConfigRequest {
	s.SceneConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetServiceCode(v string) *UpdateServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetServiceConfig(v string) *UpdateServiceConfigRequest {
	s.ServiceConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetVideoConfig(v string) *UpdateServiceConfigRequest {
	s.VideoConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
