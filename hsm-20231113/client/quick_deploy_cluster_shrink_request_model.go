// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickDeployClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertManaged(v bool) *QuickDeployClusterShrinkRequest
	GetCertManaged() *bool
	SetClusterName(v string) *QuickDeployClusterShrinkRequest
	GetClusterName() *string
	SetInstanceListShrink(v string) *QuickDeployClusterShrinkRequest
	GetInstanceListShrink() *string
	SetRegionId(v string) *QuickDeployClusterShrinkRequest
	GetRegionId() *string
	SetVSwitchIdListShrink(v string) *QuickDeployClusterShrinkRequest
	GetVSwitchIdListShrink() *string
	SetVpcId(v string) *QuickDeployClusterShrinkRequest
	GetVpcId() *string
	SetWhiteListShrink(v string) *QuickDeployClusterShrinkRequest
	GetWhiteListShrink() *string
}

type QuickDeployClusterShrinkRequest struct {
	CertManaged *bool `json:"CertManaged,omitempty" xml:"CertManaged,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIdListShrink *string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WhiteListShrink *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s QuickDeployClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuickDeployClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuickDeployClusterShrinkRequest) GetCertManaged() *bool {
	return s.CertManaged
}

func (s *QuickDeployClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *QuickDeployClusterShrinkRequest) GetInstanceListShrink() *string {
	return s.InstanceListShrink
}

func (s *QuickDeployClusterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QuickDeployClusterShrinkRequest) GetVSwitchIdListShrink() *string {
	return s.VSwitchIdListShrink
}

func (s *QuickDeployClusterShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *QuickDeployClusterShrinkRequest) GetWhiteListShrink() *string {
	return s.WhiteListShrink
}

func (s *QuickDeployClusterShrinkRequest) SetCertManaged(v bool) *QuickDeployClusterShrinkRequest {
	s.CertManaged = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetClusterName(v string) *QuickDeployClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetInstanceListShrink(v string) *QuickDeployClusterShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetRegionId(v string) *QuickDeployClusterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetVSwitchIdListShrink(v string) *QuickDeployClusterShrinkRequest {
	s.VSwitchIdListShrink = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetVpcId(v string) *QuickDeployClusterShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) SetWhiteListShrink(v string) *QuickDeployClusterShrinkRequest {
	s.WhiteListShrink = &v
	return s
}

func (s *QuickDeployClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
