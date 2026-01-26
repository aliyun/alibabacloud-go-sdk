// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvironmentResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvironmentResponseBodyData) *DescribeEnvironmentResponseBody
	GetData() *DescribeEnvironmentResponseBodyData
	SetMessage(v string) *DescribeEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvironmentResponseBody
	GetRequestId() *string
}

type DescribeEnvironmentResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *DescribeEnvironmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvironmentResponseBody) GetData() *DescribeEnvironmentResponseBodyData {
	return s.Data
}

func (s *DescribeEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvironmentResponseBody) SetCode(v int32) *DescribeEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvironmentResponseBody) SetData(v *DescribeEnvironmentResponseBodyData) *DescribeEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvironmentResponseBody) SetMessage(v string) *DescribeEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvironmentResponseBody) SetRequestId(v string) *DescribeEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvironmentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvironmentResponseBodyData struct {
	// The ID of the resource associated with the environment, such as the ACK cluster ID or VPC ID.
	//
	// example:
	//
	// vpc-xxxxx
	BindResourceId *string `json:"BindResourceId,omitempty" xml:"BindResourceId,omitempty"`
	// The profile of the resource.
	//
	// example:
	//
	// Default
	BindResourceProfile *string `json:"BindResourceProfile,omitempty" xml:"BindResourceProfile,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// running
	BindResourceStatus *string `json:"BindResourceStatus,omitempty" xml:"BindResourceStatus,omitempty"`
	// The retention period of the resource. Unit: days.
	//
	// example:
	//
	// 15
	BindResourceStoreDuration *string `json:"BindResourceStoreDuration,omitempty" xml:"BindResourceStoreDuration,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ECS
	BindResourceType *string `json:"BindResourceType,omitempty" xml:"BindResourceType,omitempty"`
	// The VPC CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	BindVpcCidr *string `json:"BindVpcCidr,omitempty" xml:"BindVpcCidr,omitempty"`
	// The status of the database that is bound to the Prometheus instance.
	//
	// Valid values:
	//
	// 	- UNINSTALLING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- INSTALLING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- UNINSTALLED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RUNNING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MODIFYING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// RUNNING
	DbInstanceStatus *string `json:"DbInstanceStatus,omitempty" xml:"DbInstanceStatus,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// env1
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// Environment subtypes:
	//
	// - CS: Currently supports ACK.
	//
	// - ECS: ECS is currently supported.
	//
	// - Cloud: Currently supports Cloud.
	//
	// example:
	//
	// ACK
	EnvironmentSubType *string `json:"EnvironmentSubType,omitempty" xml:"EnvironmentSubType,omitempty"`
	// The type of the environment. Valid values:
	//
	// 	- CS: Container Service for Kubernetes (ACK)
	//
	// 	- ECS: Elastic Compute Service
	//
	// 	- Cloud: cloud service
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"EnvironmentType,omitempty" xml:"EnvironmentType,omitempty"`
	// The payable resource plan. Valid values:
	//
	// 	- If the EnvironmentType parameter is set to CS, set the value to CS_Basic or CS_Pro.
	//
	// 	- Otherwise, leave the parameter empty.
	//
	// example:
	//
	// CS_Basic
	FeePackage *string `json:"FeePackage,omitempty" xml:"FeePackage,omitempty"`
	// The name of the Grafana data source.
	//
	// example:
	//
	// datasource1
	GrafaDataSourceName *string `json:"GrafaDataSourceName,omitempty" xml:"GrafaDataSourceName,omitempty"`
	// The unique ID of the Grafana data source.
	//
	// example:
	//
	// zuvw
	GrafanaDatasourceUid *string `json:"GrafanaDatasourceUid,omitempty" xml:"GrafanaDatasourceUid,omitempty"`
	// The name of the Grafana directory.
	//
	// example:
	//
	// folder1
	GrafanaFolderTitle *string `json:"GrafanaFolderTitle,omitempty" xml:"GrafanaFolderTitle,omitempty"`
	// The unique ID of the Grafana directory.
	//
	// example:
	//
	// xyz
	GrafanaFolderUid *string `json:"GrafanaFolderUid,omitempty" xml:"GrafanaFolderUid,omitempty"`
	// The URL of the Grafana directory.
	//
	// example:
	//
	// https://g.console.aliyun.com/dashboards/f/xxx/yyyy
	GrafanaFolderUrl *string `json:"GrafanaFolderUrl,omitempty" xml:"GrafanaFolderUrl,omitempty"`
	// The ID of the Grafana workspace.
	//
	// example:
	//
	// grafana-cn-27a3m8eem0a
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// managed type:
	//
	// - none: unmanaged. The default value for ACK clusters.
	//
	// - agent: managed agent (including KSM). The default values for ASK, ACS, and AckOne clusters.
	//
	// - agent-exporter: managed agent and exporters. The default value for the cloud service type.
	//
	// example:
	//
	// none
	ManagedType *string `json:"ManagedType,omitempty" xml:"ManagedType,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// xxxxyyyyyzzzzz
	PrometheusInstanceId *string `json:"PrometheusInstanceId,omitempty" xml:"PrometheusInstanceId,omitempty"`
	// The name of the Prometheus instance.
	//
	// example:
	//
	// name1
	PrometheusInstanceName *string `json:"PrometheusInstanceName,omitempty" xml:"PrometheusInstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2vezare****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group associated with the environment.
	//
	// example:
	//
	// sg-8vbdgmf4nraiqa9bx0jo
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tags.
	Tags []*DescribeEnvironmentResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 13002222xxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vb02uk57qbcktqcvqqqj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch associated with the environment.
	//
	// example:
	//
	// vsw-2ze7yr3f1x8snryaioo7u
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeEnvironmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentResponseBodyData) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *DescribeEnvironmentResponseBodyData) GetBindResourceProfile() *string {
	return s.BindResourceProfile
}

func (s *DescribeEnvironmentResponseBodyData) GetBindResourceStatus() *string {
	return s.BindResourceStatus
}

func (s *DescribeEnvironmentResponseBodyData) GetBindResourceStoreDuration() *string {
	return s.BindResourceStoreDuration
}

func (s *DescribeEnvironmentResponseBodyData) GetBindResourceType() *string {
	return s.BindResourceType
}

func (s *DescribeEnvironmentResponseBodyData) GetBindVpcCidr() *string {
	return s.BindVpcCidr
}

func (s *DescribeEnvironmentResponseBodyData) GetDbInstanceStatus() *string {
	return s.DbInstanceStatus
}

func (s *DescribeEnvironmentResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvironmentResponseBodyData) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *DescribeEnvironmentResponseBodyData) GetEnvironmentSubType() *string {
	return s.EnvironmentSubType
}

func (s *DescribeEnvironmentResponseBodyData) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *DescribeEnvironmentResponseBodyData) GetFeePackage() *string {
	return s.FeePackage
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafaDataSourceName() *string {
	return s.GrafaDataSourceName
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafanaDatasourceUid() *string {
	return s.GrafanaDatasourceUid
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafanaFolderTitle() *string {
	return s.GrafanaFolderTitle
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafanaFolderUid() *string {
	return s.GrafanaFolderUid
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafanaFolderUrl() *string {
	return s.GrafanaFolderUrl
}

func (s *DescribeEnvironmentResponseBodyData) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *DescribeEnvironmentResponseBodyData) GetManagedType() *string {
	return s.ManagedType
}

func (s *DescribeEnvironmentResponseBodyData) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *DescribeEnvironmentResponseBodyData) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *DescribeEnvironmentResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvironmentResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEnvironmentResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeEnvironmentResponseBodyData) GetTags() []*DescribeEnvironmentResponseBodyDataTags {
	return s.Tags
}

func (s *DescribeEnvironmentResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEnvironmentResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEnvironmentResponseBodyData) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeEnvironmentResponseBodyData) SetBindResourceId(v string) *DescribeEnvironmentResponseBodyData {
	s.BindResourceId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetBindResourceProfile(v string) *DescribeEnvironmentResponseBodyData {
	s.BindResourceProfile = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetBindResourceStatus(v string) *DescribeEnvironmentResponseBodyData {
	s.BindResourceStatus = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetBindResourceStoreDuration(v string) *DescribeEnvironmentResponseBodyData {
	s.BindResourceStoreDuration = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetBindResourceType(v string) *DescribeEnvironmentResponseBodyData {
	s.BindResourceType = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetBindVpcCidr(v string) *DescribeEnvironmentResponseBodyData {
	s.BindVpcCidr = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetDbInstanceStatus(v string) *DescribeEnvironmentResponseBodyData {
	s.DbInstanceStatus = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetEnvironmentId(v string) *DescribeEnvironmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetEnvironmentName(v string) *DescribeEnvironmentResponseBodyData {
	s.EnvironmentName = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetEnvironmentSubType(v string) *DescribeEnvironmentResponseBodyData {
	s.EnvironmentSubType = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetEnvironmentType(v string) *DescribeEnvironmentResponseBodyData {
	s.EnvironmentType = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetFeePackage(v string) *DescribeEnvironmentResponseBodyData {
	s.FeePackage = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafaDataSourceName(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafaDataSourceName = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafanaDatasourceUid(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafanaDatasourceUid = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafanaFolderTitle(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafanaFolderTitle = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafanaFolderUid(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafanaFolderUid = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafanaFolderUrl(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafanaFolderUrl = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetGrafanaWorkspaceId(v string) *DescribeEnvironmentResponseBodyData {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetManagedType(v string) *DescribeEnvironmentResponseBodyData {
	s.ManagedType = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetPrometheusInstanceId(v string) *DescribeEnvironmentResponseBodyData {
	s.PrometheusInstanceId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetPrometheusInstanceName(v string) *DescribeEnvironmentResponseBodyData {
	s.PrometheusInstanceName = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetRegionId(v string) *DescribeEnvironmentResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetResourceGroupId(v string) *DescribeEnvironmentResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetSecurityGroupId(v string) *DescribeEnvironmentResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetTags(v []*DescribeEnvironmentResponseBodyDataTags) *DescribeEnvironmentResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetUserId(v string) *DescribeEnvironmentResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetVpcId(v string) *DescribeEnvironmentResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) SetVswitchId(v string) *DescribeEnvironmentResponseBodyData {
	s.VswitchId = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnvironmentResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// user1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// p_dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnvironmentResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *DescribeEnvironmentResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *DescribeEnvironmentResponseBodyDataTags) SetKey(v string) *DescribeEnvironmentResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyDataTags) SetValue(v string) *DescribeEnvironmentResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *DescribeEnvironmentResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
