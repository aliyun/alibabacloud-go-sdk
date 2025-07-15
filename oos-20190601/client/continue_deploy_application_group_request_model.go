// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeployApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ContinueDeployApplicationGroupRequest
	GetApplicationName() *string
	SetDeployParameters(v string) *ContinueDeployApplicationGroupRequest
	GetDeployParameters() *string
	SetName(v string) *ContinueDeployApplicationGroupRequest
	GetName() *string
	SetRegionId(v string) *ContinueDeployApplicationGroupRequest
	GetRegionId() *string
}

type ContinueDeployApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The deployment information about the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// {       "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",       "Parameters": {         "ZoneId": "cn-hangzhou-k",         "ProjectName": "test",         "SystemDiskSize": 40,         "InstanceChargeType": "PostPaid",         "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",         "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",         "SystemDiskCategory": "cloud_essd",         "InstancePassword": "******",         "InternetChargeType": "PayByTraffic",         "InstanceCount": 1,         "InternetMaxBandwidthOut": 0,         "VpcId": "vpc-bp1i99boyas8i8m9t3skp",         "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",         "DataDiskSize": 100,         "EcsInstanceType": "ecs.s6-c1m4.small",         "DataDiskCategory": "cloud_efficiency",         "EnvironmentCommandId": "c-hz028fc3g031gcg"       }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ContinueDeployApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ContinueDeployApplicationGroupRequest) GetDeployParameters() *string {
	return s.DeployParameters
}

func (s *ContinueDeployApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *ContinueDeployApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContinueDeployApplicationGroupRequest) SetApplicationName(v string) *ContinueDeployApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetDeployParameters(v string) *ContinueDeployApplicationGroupRequest {
	s.DeployParameters = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetName(v string) *ContinueDeployApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetRegionId(v string) *ContinueDeployApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}
