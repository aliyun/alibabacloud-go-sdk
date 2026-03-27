// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyServiceMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonReleaseName(v string) *ListIntegrationPolicyServiceMonitorsRequest
	GetAddonReleaseName() *string
	SetEncryptYaml(v bool) *ListIntegrationPolicyServiceMonitorsRequest
	GetEncryptYaml() *bool
	SetNamespace(v string) *ListIntegrationPolicyServiceMonitorsRequest
	GetNamespace() *string
}

type ListIntegrationPolicyServiceMonitorsRequest struct {
	// example:
	//
	// release-12345678
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// true
	EncryptYaml *bool `json:"encryptYaml,omitempty" xml:"encryptYaml,omitempty"`
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyServiceMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyServiceMonitorsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) SetAddonReleaseName(v string) *ListIntegrationPolicyServiceMonitorsRequest {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) SetEncryptYaml(v bool) *ListIntegrationPolicyServiceMonitorsRequest {
	s.EncryptYaml = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) SetNamespace(v string) *ListIntegrationPolicyServiceMonitorsRequest {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
