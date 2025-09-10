// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyStorageRequirementsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListIntegrationPolicyStorageRequirementsRequest
	GetAddonName() *string
	SetAddonReleaseName(v string) *ListIntegrationPolicyStorageRequirementsRequest
	GetAddonReleaseName() *string
	SetStorageType(v string) *ListIntegrationPolicyStorageRequirementsRequest
	GetStorageType() *string
}

type ListIntegrationPolicyStorageRequirementsRequest struct {
	// example:
	//
	// release-1234357
	AddonName        *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	AddonReleaseName *string `json:"addonReleaseName,omitempty" xml:"addonReleaseName,omitempty"`
	// example:
	//
	// LogStore
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
}

func (s ListIntegrationPolicyStorageRequirementsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) SetAddonName(v string) *ListIntegrationPolicyStorageRequirementsRequest {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) SetAddonReleaseName(v string) *ListIntegrationPolicyStorageRequirementsRequest {
	s.AddonReleaseName = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) SetStorageType(v string) *ListIntegrationPolicyStorageRequirementsRequest {
	s.StorageType = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsRequest) Validate() error {
	return dara.Validate(s)
}
