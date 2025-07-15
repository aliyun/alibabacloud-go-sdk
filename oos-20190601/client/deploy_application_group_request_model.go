// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *DeployApplicationGroupRequest
	GetApplicationName() *string
	SetDeployParameters(v string) *DeployApplicationGroupRequest
	GetDeployParameters() *string
	SetName(v string) *DeployApplicationGroupRequest
	GetName() *string
	SetRegionId(v string) *DeployApplicationGroupRequest
	GetRegionId() *string
}

type DeployApplicationGroupRequest struct {
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
	// The ID of the region in which you want to deploy the application group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeployApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DeployApplicationGroupRequest) GetDeployParameters() *string {
	return s.DeployParameters
}

func (s *DeployApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *DeployApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeployApplicationGroupRequest) SetApplicationName(v string) *DeployApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetDeployParameters(v string) *DeployApplicationGroupRequest {
	s.DeployParameters = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetName(v string) *DeployApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetRegionId(v string) *DeployApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeployApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}
