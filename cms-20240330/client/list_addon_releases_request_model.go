// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonReleasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListAddonReleasesRequest
	GetAddonName() *string
	SetParentAddonReleaseId(v string) *ListAddonReleasesRequest
	GetParentAddonReleaseId() *string
}

type ListAddonReleasesRequest struct {
	// Addon component name.
	//
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// Parent AddonReleaseId.
	//
	// example:
	//
	// policy-xxxxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
}

func (s ListAddonReleasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesRequest) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListAddonReleasesRequest) GetParentAddonReleaseId() *string {
	return s.ParentAddonReleaseId
}

func (s *ListAddonReleasesRequest) SetAddonName(v string) *ListAddonReleasesRequest {
	s.AddonName = &v
	return s
}

func (s *ListAddonReleasesRequest) SetParentAddonReleaseId(v string) *ListAddonReleasesRequest {
	s.ParentAddonReleaseId = &v
	return s
}

func (s *ListAddonReleasesRequest) Validate() error {
	return dara.Validate(s)
}
