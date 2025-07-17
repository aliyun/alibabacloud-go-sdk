// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DescribeAddonReleaseRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DescribeAddonReleaseRequest
	GetRegionId() *string
	SetReleaseName(v string) *DescribeAddonReleaseRequest
	GetReleaseName() *string
}

type DescribeAddonReleaseRequest struct {
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
	// The name of the add-on release.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-822567d4-2449
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
}

func (s DescribeAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeAddonReleaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddonReleaseRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *DescribeAddonReleaseRequest) SetEnvironmentId(v string) *DescribeAddonReleaseRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeAddonReleaseRequest) SetRegionId(v string) *DescribeAddonReleaseRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonReleaseRequest) SetReleaseName(v string) *DescribeAddonReleaseRequest {
	s.ReleaseName = &v
	return s
}

func (s *DescribeAddonReleaseRequest) Validate() error {
	return dara.Validate(s)
}
