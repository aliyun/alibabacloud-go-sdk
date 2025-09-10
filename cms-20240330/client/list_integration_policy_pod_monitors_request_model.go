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
	SetEncryptYaml(v bool) *ListIntegrationPolicyPodMonitorsRequest
	GetEncryptYaml() *bool
	SetNamespace(v string) *ListIntegrationPolicyPodMonitorsRequest
	GetNamespace() *string
}

type ListIntegrationPolicyPodMonitorsRequest struct {
	// example:
	//
	// release-123456789
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
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
