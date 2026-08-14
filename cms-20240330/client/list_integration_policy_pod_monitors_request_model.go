// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyPodMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonReleaseName(v string) *ListIntegrationPolicyPodMonitorsRequest
	GetAddonReleaseName() *string
	SetCollectorReleaseName(v string) *ListIntegrationPolicyPodMonitorsRequest
	GetCollectorReleaseName() *string
	SetEncryptYaml(v bool) *ListIntegrationPolicyPodMonitorsRequest
	GetEncryptYaml() *bool
	SetNamespace(v string) *ListIntegrationPolicyPodMonitorsRequest
	GetNamespace() *string
}

type ListIntegrationPolicyPodMonitorsRequest struct {
	// The name of the addon release.
	//
	// example:
	//
	// release-123456789
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// The identifier of the collector. If a release exists, pass the release name. If no release exists, pass the component name.
	//
	// example:
	//
	// collector:metric-agent:policy:policy-bfd3d455fd6f4bc8
	CollectorReleaseName *string `json:"collectorReleaseName,omitempty" xml:"collectorReleaseName,omitempty"`
	// Specifies whether to encrypt the YAML.
	//
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyPodMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyPodMonitorsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyPodMonitorsRequest) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyPodMonitorsRequest) GetCollectorReleaseName() *string {
	return s.CollectorReleaseName
}

func (s *ListIntegrationPolicyPodMonitorsRequest) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyPodMonitorsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyPodMonitorsRequest) SetAddonReleaseName(v string) *ListIntegrationPolicyPodMonitorsRequest {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsRequest) SetCollectorReleaseName(v string) *ListIntegrationPolicyPodMonitorsRequest {
	s.CollectorReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsRequest) SetEncryptYaml(v bool) *ListIntegrationPolicyPodMonitorsRequest {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsRequest) SetNamespace(v string) *ListIntegrationPolicyPodMonitorsRequest {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
