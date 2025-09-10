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
	SetEncryptYaml(v bool) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetEncryptYaml() *bool
	SetNamespace(v string) *ListIntegrationPolicyCustomScrapeJobRulesRequest
	GetNamespace() *string
}

type ListIntegrationPolicyCustomScrapeJobRulesRequest struct {
	// example:
	//
	// release12345678
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// true
	EncryptYaml *bool   `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	Namespace   *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
