// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationAvailableVersionShrinkRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationAvailableVersionShrinkRequest
	GetApplicationType() *string
	SetComponentsShrink(v string) *DescribeApplicationAvailableVersionShrinkRequest
	GetComponentsShrink() *string
	SetRegionId(v string) *DescribeApplicationAvailableVersionShrinkRequest
	GetRegionId() *string
}

type DescribeApplicationAvailableVersionShrinkRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The list of application subcomponents.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApplicationAvailableVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) SetApplicationId(v string) *DescribeApplicationAvailableVersionShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) SetApplicationType(v string) *DescribeApplicationAvailableVersionShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) SetComponentsShrink(v string) *DescribeApplicationAvailableVersionShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) SetRegionId(v string) *DescribeApplicationAvailableVersionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAvailableVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
