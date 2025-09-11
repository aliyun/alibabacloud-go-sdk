// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickDeployClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertManaged(v bool) *QuickDeployClusterRequest
	GetCertManaged() *bool
	SetClusterName(v string) *QuickDeployClusterRequest
	GetClusterName() *string
	SetInstanceList(v []*string) *QuickDeployClusterRequest
	GetInstanceList() []*string
	SetRegionId(v string) *QuickDeployClusterRequest
	GetRegionId() *string
	SetVSwitchIdList(v []*string) *QuickDeployClusterRequest
	GetVSwitchIdList() []*string
	SetVpcId(v string) *QuickDeployClusterRequest
	GetVpcId() *string
	SetWhiteList(v []*string) *QuickDeployClusterRequest
	GetWhiteList() []*string
}

type QuickDeployClusterRequest struct {
	CertManaged *bool `json:"CertManaged,omitempty" xml:"CertManaged,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s QuickDeployClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s QuickDeployClusterRequest) GoString() string {
	return s.String()
}

func (s *QuickDeployClusterRequest) GetCertManaged() *bool {
	return s.CertManaged
}

func (s *QuickDeployClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *QuickDeployClusterRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *QuickDeployClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QuickDeployClusterRequest) GetVSwitchIdList() []*string {
	return s.VSwitchIdList
}

func (s *QuickDeployClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *QuickDeployClusterRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *QuickDeployClusterRequest) SetCertManaged(v bool) *QuickDeployClusterRequest {
	s.CertManaged = &v
	return s
}

func (s *QuickDeployClusterRequest) SetClusterName(v string) *QuickDeployClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *QuickDeployClusterRequest) SetInstanceList(v []*string) *QuickDeployClusterRequest {
	s.InstanceList = v
	return s
}

func (s *QuickDeployClusterRequest) SetRegionId(v string) *QuickDeployClusterRequest {
	s.RegionId = &v
	return s
}

func (s *QuickDeployClusterRequest) SetVSwitchIdList(v []*string) *QuickDeployClusterRequest {
	s.VSwitchIdList = v
	return s
}

func (s *QuickDeployClusterRequest) SetVpcId(v string) *QuickDeployClusterRequest {
	s.VpcId = &v
	return s
}

func (s *QuickDeployClusterRequest) SetWhiteList(v []*string) *QuickDeployClusterRequest {
	s.WhiteList = v
	return s
}

func (s *QuickDeployClusterRequest) Validate() error {
	return dara.Validate(s)
}
