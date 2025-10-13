// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonVersion(v string) *UpdateAddonReleaseRequest
	GetAddonVersion() *string
	SetDryRun(v bool) *UpdateAddonReleaseRequest
	GetDryRun() *bool
	SetEntityRules(v *EntityDiscoverRule) *UpdateAddonReleaseRequest
	GetEntityRules() *EntityDiscoverRule
	SetValues(v string) *UpdateAddonReleaseRequest
	GetValues() *string
}

type UpdateAddonReleaseRequest struct {
	// example:
	//
	// 0.0.2
	AddonVersion *string `json:"addonVersion,omitempty" xml:"addonVersion,omitempty"`
	// example:
	//
	// true
	DryRun      *bool               `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	EntityRules *EntityDiscoverRule `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Values *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s UpdateAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateAddonReleaseRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *UpdateAddonReleaseRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateAddonReleaseRequest) GetEntityRules() *EntityDiscoverRule {
	return s.EntityRules
}

func (s *UpdateAddonReleaseRequest) GetValues() *string {
	return s.Values
}

func (s *UpdateAddonReleaseRequest) SetAddonVersion(v string) *UpdateAddonReleaseRequest {
	s.AddonVersion = &v
	return s
}

func (s *UpdateAddonReleaseRequest) SetDryRun(v bool) *UpdateAddonReleaseRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAddonReleaseRequest) SetEntityRules(v *EntityDiscoverRule) *UpdateAddonReleaseRequest {
	s.EntityRules = v
	return s
}

func (s *UpdateAddonReleaseRequest) SetValues(v string) *UpdateAddonReleaseRequest {
	s.Values = &v
	return s
}

func (s *UpdateAddonReleaseRequest) Validate() error {
	if s.EntityRules != nil {
		if err := s.EntityRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}
