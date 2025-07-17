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
	SetEnvironmentId(v string) *ListAddonReleasesRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListAddonReleasesRequest
	GetRegionId() *string
}

type ListAddonReleasesRequest struct {
	// The name of the add-on.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListAddonReleasesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListAddonReleasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAddonReleasesRequest) SetAddonName(v string) *ListAddonReleasesRequest {
	s.AddonName = &v
	return s
}

func (s *ListAddonReleasesRequest) SetEnvironmentId(v string) *ListAddonReleasesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListAddonReleasesRequest) SetRegionId(v string) *ListAddonReleasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAddonReleasesRequest) Validate() error {
	return dara.Validate(s)
}
