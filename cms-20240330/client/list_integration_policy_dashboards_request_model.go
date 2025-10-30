// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListIntegrationPolicyDashboardsRequest
	GetAddonName() *string
	SetLanguage(v string) *ListIntegrationPolicyDashboardsRequest
	GetLanguage() *string
	SetScene(v string) *ListIntegrationPolicyDashboardsRequest
	GetScene() *string
}

type ListIntegrationPolicyDashboardsRequest struct {
	// Addon Name.
	//
	// example:
	//
	// cs-default
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// Query Language
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Component Scenario.
	//
	// example:
	//
	// databse
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s ListIntegrationPolicyDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyDashboardsRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPolicyDashboardsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListIntegrationPolicyDashboardsRequest) GetScene() *string {
	return s.Scene
}

func (s *ListIntegrationPolicyDashboardsRequest) SetAddonName(v string) *ListIntegrationPolicyDashboardsRequest {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsRequest) SetLanguage(v string) *ListIntegrationPolicyDashboardsRequest {
	s.Language = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsRequest) SetScene(v string) *ListIntegrationPolicyDashboardsRequest {
	s.Scene = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
