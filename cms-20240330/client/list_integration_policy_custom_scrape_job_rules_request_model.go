// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCustomScrapeJobRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonReleaseName(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetAddonReleaseName() *string
	SetCollectorReleaseName(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetCollectorReleaseName() *string
	SetEncryptYaml(v bool) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetEncryptYaml() *bool
	SetNamespace(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetNamespace() *string
}

type ListIntegrationPolicyCustomScrapeJobRulesRequest struct {
	// The name of the addon release.
	//
	// example:
	//
	// release12345678
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// The probe identifier. If a release exists, pass the release name. If no release exists, pass the component name.
	//
	// example:
	//
	// collector:metric-agent:policy:policy-bfd3d455fd6f4bc8
	CollectorReleaseName *string `json:"collectorReleaseName,omitempty" xml:"collectorReleaseName,omitempty"`
	// Specifies whether to encrypt the YAML content.
	//
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyCustomScrapeJobRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCustomScrapeJobRulesRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) GetCollectorReleaseName() *string {
	return s.CollectorReleaseName
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) SetAddonReleaseName(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) SetCollectorReleaseName(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest {
	s.CollectorReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) SetEncryptYaml(v bool) *ListIntegrationPolicyCustomScrapeJobRulesRequest {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) SetNamespace(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesRequest) Validate() error {
	return dara.Validate(s)
}
