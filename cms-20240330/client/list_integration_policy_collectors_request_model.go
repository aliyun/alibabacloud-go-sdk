// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCollectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonReleaseName(v string) *ListIntegrationPolicyCollectorsRequest
	GetAddonReleaseName() *string
	SetCollectorType(v string) *ListIntegrationPolicyCollectorsRequest
	GetCollectorType() *string
	SetLanguage(v string) *ListIntegrationPolicyCollectorsRequest
	GetLanguage() *string
}

type ListIntegrationPolicyCollectorsRequest struct {
	// example:
	//
	// release-1234567
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Exporter
	CollectorType *string `json:"collectorType,omitempty" xml:"collectorType,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ListIntegrationPolicyCollectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsRequest) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyCollectorsRequest) GetCollectorType() *string {
	return s.CollectorType
}

func (s *ListIntegrationPolicyCollectorsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListIntegrationPolicyCollectorsRequest) SetAddonReleaseName(v string) *ListIntegrationPolicyCollectorsRequest {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsRequest) SetCollectorType(v string) *ListIntegrationPolicyCollectorsRequest {
	s.CollectorType = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsRequest) SetLanguage(v string) *ListIntegrationPolicyCollectorsRequest {
	s.Language = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsRequest) Validate() error {
	return dara.Validate(s)
}
