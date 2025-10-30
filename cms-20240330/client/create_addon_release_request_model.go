// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *CreateAddonReleaseRequest
	GetAddonName() *string
	SetAliyunLang(v string) *CreateAddonReleaseRequest
	GetAliyunLang() *string
	SetDryRun(v bool) *CreateAddonReleaseRequest
	GetDryRun() *bool
	SetEntityRules(v *EntityDiscoverRule) *CreateAddonReleaseRequest
	GetEntityRules() *EntityDiscoverRule
	SetEnvType(v string) *CreateAddonReleaseRequest
	GetEnvType() *string
	SetParentAddonReleaseId(v string) *CreateAddonReleaseRequest
	GetParentAddonReleaseId() *string
	SetReleaseName(v string) *CreateAddonReleaseRequest
	GetReleaseName() *string
	SetValues(v string) *CreateAddonReleaseRequest
	GetValues() *string
	SetVersion(v string) *CreateAddonReleaseRequest
	GetVersion() *string
	SetWorkspace(v string) *CreateAddonReleaseRequest
	GetWorkspace() *string
}

type CreateAddonReleaseRequest struct {
	// The Addon name of the component that needs to be monitored.
	//
	// This parameter is required.
	//
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// The language type of the component.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// Whether it is a dry run, default is false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Field rules
	EntityRules *EntityDiscoverRule `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
	// Environment type. If the Policy type is CS and ECS, use accordingly; otherwise, it is unified as Cloud.
	//
	// example:
	//
	// CS
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// Parent AddonReleaseId.
	//
	// example:
	//
	// policy-xxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
	// The plugin name after access. If not specified, a default rule name will be generated.
	//
	// example:
	//
	// test-gpu-integration-name
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// Input metadata.
	//
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Values *string `json:"values,omitempty" xml:"values,omitempty"`
	// The version of the Addon component that needs to be monitored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace name for installing the component resources.
	//
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *CreateAddonReleaseRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *CreateAddonReleaseRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateAddonReleaseRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAddonReleaseRequest) GetEntityRules() *EntityDiscoverRule {
	return s.EntityRules
}

func (s *CreateAddonReleaseRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateAddonReleaseRequest) GetParentAddonReleaseId() *string {
	return s.ParentAddonReleaseId
}

func (s *CreateAddonReleaseRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *CreateAddonReleaseRequest) GetValues() *string {
	return s.Values
}

func (s *CreateAddonReleaseRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateAddonReleaseRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateAddonReleaseRequest) SetAddonName(v string) *CreateAddonReleaseRequest {
	s.AddonName = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetAliyunLang(v string) *CreateAddonReleaseRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetDryRun(v bool) *CreateAddonReleaseRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetEntityRules(v *EntityDiscoverRule) *CreateAddonReleaseRequest {
	s.EntityRules = v
	return s
}

func (s *CreateAddonReleaseRequest) SetEnvType(v string) *CreateAddonReleaseRequest {
	s.EnvType = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetParentAddonReleaseId(v string) *CreateAddonReleaseRequest {
	s.ParentAddonReleaseId = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetReleaseName(v string) *CreateAddonReleaseRequest {
	s.ReleaseName = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetValues(v string) *CreateAddonReleaseRequest {
	s.Values = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetVersion(v string) *CreateAddonReleaseRequest {
	s.Version = &v
	return s
}

func (s *CreateAddonReleaseRequest) SetWorkspace(v string) *CreateAddonReleaseRequest {
	s.Workspace = &v
	return s
}

func (s *CreateAddonReleaseRequest) Validate() error {
	if s.EntityRules != nil {
		if err := s.EntityRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}
