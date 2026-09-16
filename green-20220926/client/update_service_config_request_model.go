// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceConfigRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateServiceConfigRequest
	GetDryRun() *bool
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
	// The client-generated idempotency token used to prevent duplicate operations caused by network retries. The token must be unique across requests and contain only printable ASCII characters (ASCII 32-126).
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. When set to true, only parameter validation and business logic checks are performed without actually creating or updating resources.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The document configuration.
	//
	// example:
	//
	// {}
	FileConfig *string `json:"FileConfig,omitempty" xml:"FileConfig,omitempty"`
	// The keyword filter libraries.
	//
	// example:
	//
	// []
	KeywordFilterLibs *string `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty"`
	// The keyword hit libraries.
	//
	// example:
	//
	// []
	KeywordHitLibs *string `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty"`
	// The human-machine moderation configuration.
	//
	// example:
	//
	// {}
	ManualMachineConfig *string `json:"ManualMachineConfig,omitempty" xml:"ManualMachineConfig,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The moderation scenario.
	//
	// example:
	//
	// pornographic
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The scenario configuration.
	//
	// example:
	//
	// {}
	SceneConfig *string `json:"SceneConfig,omitempty" xml:"SceneConfig,omitempty"`
	// The service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The service configuration.
	//
	// example:
	//
	// {}
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// The video configuration.
	//
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

func (s *UpdateServiceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceConfigRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *UpdateServiceConfigRequest) SetClientToken(v string) *UpdateServiceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetDryRun(v bool) *UpdateServiceConfigRequest {
	s.DryRun = &v
	return s
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
