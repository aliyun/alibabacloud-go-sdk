// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentsResponseBody
	GetCode() *int32
	SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody
	GetData() *ListEnvironmentsResponseBodyData
	SetMessage(v string) *ListEnvironmentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentsResponseBody
	GetSuccess() *bool
}

type ListEnvironmentsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *ListEnvironmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentsResponseBody) GetData() *ListEnvironmentsResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentsResponseBody) SetCode(v int32) *ListEnvironmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetMessage(v string) *ListEnvironmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetRequestId(v string) *ListEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetSuccess(v bool) *ListEnvironmentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentsResponseBodyData struct {
	// The queried environments.
	Environments []*ListEnvironmentsResponseBodyDataEnvironments `json:"Environments,omitempty" xml:"Environments,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) GetEnvironments() []*ListEnvironmentsResponseBodyDataEnvironments {
	return s.Environments
}

func (s *ListEnvironmentsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListEnvironmentsResponseBodyData) SetEnvironments(v []*ListEnvironmentsResponseBodyDataEnvironments) *ListEnvironmentsResponseBodyData {
	s.Environments = v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetTotal(v int64) *ListEnvironmentsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentsResponseBodyDataEnvironments struct {
	// The add-ons.
	Addons []*ListEnvironmentsResponseBodyDataEnvironmentsAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The ID of the resource bound to the environment instance. The resource can be a Kubernetes cluster or a VPC.
	//
	// example:
	//
	// vpc-bp1bgo8ronn
	BindResourceId *string `json:"BindResourceId,omitempty" xml:"BindResourceId,omitempty"`
	// The profile that is bound to the resource.
	//
	// example:
	//
	// xxx
	BindResourceProfile *string `json:"BindResourceProfile,omitempty" xml:"BindResourceProfile,omitempty"`
	// The resource type.
	//
	// example:
	//
	// VPC
	BindResourceType *string `json:"BindResourceType,omitempty" xml:"BindResourceType,omitempty"`
	// The CIDR block that is bound to the VPC.
	//
	// example:
	//
	// 172.16.0.0/12
	BindVpcCidr *string `json:"BindVpcCidr,omitempty" xml:"BindVpcCidr,omitempty"`
	// The time when the environment instance was created.
	//
	// example:
	//
	// 2023-03-24 11:58:35 +0800
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12378523784982
	CreatedUserId *string `json:"CreatedUserId,omitempty" xml:"CreatedUserId,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The name of the environment instance.
	//
	// example:
	//
	// feiliks-biz-prod-edas
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The type of the environment instance. Valid values:
	//
	// 	- CS: Container Service
	//
	// 	- ECS: Elastic Compute Service
	//
	// 	- Cloud: cloud service
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"EnvironmentType,omitempty" xml:"EnvironmentType,omitempty"`
	// The parameters of the feature.
	Features []*ListEnvironmentsResponseBodyDataEnvironmentsFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The payable resource plan.
	//
	// 	- If the EnvironmentType parameter is set to CS, set the value to CS_Basic or CS_Pro.
	//
	// 	- Otherwise, leave the parameter empty.
	//
	// example:
	//
	// CS_Pro
	FeePackage *string `json:"FeePackage,omitempty" xml:"FeePackage,omitempty"`
	// The unique ID of the Grafana data source.
	//
	// example:
	//
	// 12374890
	GrafanaDatasourceUid *string `json:"GrafanaDatasourceUid,omitempty" xml:"GrafanaDatasourceUid,omitempty"`
	// The name of the Grafana directory.
	//
	// example:
	//
	// filepath
	GrafanaFolderTitle *string `json:"GrafanaFolderTitle,omitempty" xml:"GrafanaFolderTitle,omitempty"`
	// The unique ID of the Grafana directory.
	//
	// example:
	//
	// 1798319482935
	GrafanaFolderUid *string `json:"GrafanaFolderUid,omitempty" xml:"GrafanaFolderUid,omitempty"`
	// The time when the last add-on was created.
	//
	// example:
	//
	// 2023-09-22T16:56:29+08:00
	LatestReleaseCreateTime *string `json:"LatestReleaseCreateTime,omitempty" xml:"LatestReleaseCreateTime,omitempty"`
	// Indicates whether agents or exporters are managed. Valid values:
	//
	// 	- none: No. By default, no managed agents or exporters are provided for ACK clusters.
	//
	// 	- agent: Agents are managed. By default, managed agents are provided for ASK clusters, ACS clusters, and ACK One clusters.
	//
	// 	- agent-exproter: Agents and exporters are managed. By default, managed agents and exporters are provided for cloud services.
	//
	// example:
	//
	// agent
	ManagedType *string `json:"ManagedType,omitempty" xml:"ManagedType,omitempty"`
	// The Prometheus ID.
	//
	// example:
	//
	// 124769812
	PrometheusId *int64 `json:"PrometheusId,omitempty" xml:"PrometheusId,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// vpc-m5e4alj2i24ndbn
	PrometheusInstanceId *string `json:"PrometheusInstanceId,omitempty" xml:"PrometheusInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of installed add-ons.
	//
	// example:
	//
	// 122
	ReleaseCount *int32 `json:"ReleaseCount,omitempty" xml:"ReleaseCount,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvt3xpr5aema
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the environment resource.
	Tags []*ListEnvironmentsResponseBodyDataEnvironmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 13990957477389
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataEnvironments) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataEnvironments) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetAddons() []*ListEnvironmentsResponseBodyDataEnvironmentsAddons {
	return s.Addons
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetBindResourceProfile() *string {
	return s.BindResourceProfile
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetBindResourceType() *string {
	return s.BindResourceType
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetBindVpcCidr() *string {
	return s.BindVpcCidr
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetCreatedUserId() *string {
	return s.CreatedUserId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetFeatures() []*ListEnvironmentsResponseBodyDataEnvironmentsFeatures {
	return s.Features
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetFeePackage() *string {
	return s.FeePackage
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetGrafanaDatasourceUid() *string {
	return s.GrafanaDatasourceUid
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetGrafanaFolderTitle() *string {
	return s.GrafanaFolderTitle
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetGrafanaFolderUid() *string {
	return s.GrafanaFolderUid
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetLatestReleaseCreateTime() *string {
	return s.LatestReleaseCreateTime
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetManagedType() *string {
	return s.ManagedType
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetPrometheusId() *int64 {
	return s.PrometheusId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetReleaseCount() *int32 {
	return s.ReleaseCount
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetTags() []*ListEnvironmentsResponseBodyDataEnvironmentsTags {
	return s.Tags
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) GetUserId() *string {
	return s.UserId
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetAddons(v []*ListEnvironmentsResponseBodyDataEnvironmentsAddons) *ListEnvironmentsResponseBodyDataEnvironments {
	s.Addons = v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetBindResourceId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.BindResourceId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetBindResourceProfile(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.BindResourceProfile = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetBindResourceType(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.BindResourceType = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetBindVpcCidr(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.BindVpcCidr = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetCreateTime(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.CreateTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetCreatedUserId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.CreatedUserId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetEnvironmentId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetEnvironmentName(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.EnvironmentName = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetEnvironmentType(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.EnvironmentType = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetFeatures(v []*ListEnvironmentsResponseBodyDataEnvironmentsFeatures) *ListEnvironmentsResponseBodyDataEnvironments {
	s.Features = v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetFeePackage(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.FeePackage = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetGrafanaDatasourceUid(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.GrafanaDatasourceUid = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetGrafanaFolderTitle(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.GrafanaFolderTitle = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetGrafanaFolderUid(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.GrafanaFolderUid = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetLatestReleaseCreateTime(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.LatestReleaseCreateTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetManagedType(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.ManagedType = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetPrometheusId(v int64) *ListEnvironmentsResponseBodyDataEnvironments {
	s.PrometheusId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetPrometheusInstanceId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.PrometheusInstanceId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetRegionId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetReleaseCount(v int32) *ListEnvironmentsResponseBodyDataEnvironments {
	s.ReleaseCount = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetResourceGroupId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetTags(v []*ListEnvironmentsResponseBodyDataEnvironmentsTags) *ListEnvironmentsResponseBodyDataEnvironments {
	s.Tags = v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) SetUserId(v string) *ListEnvironmentsResponseBodyDataEnvironments {
	s.UserId = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironments) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentsResponseBodyDataEnvironmentsAddons struct {
	// The alias of the add-on.
	//
	// example:
	//
	// MySQL Exporter
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The description of the add-on.
	//
	// example:
	//
	// Collect mysql indicator information
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// http://xxxx
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The name of the add-on.
	//
	// example:
	//
	// metric-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsAddons) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsAddons) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) GetAlias() *string {
	return s.Alias
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) GetIcon() *string {
	return s.Icon
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) SetAlias(v string) *ListEnvironmentsResponseBodyDataEnvironmentsAddons {
	s.Alias = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) SetDescription(v string) *ListEnvironmentsResponseBodyDataEnvironmentsAddons {
	s.Description = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) SetIcon(v string) *ListEnvironmentsResponseBodyDataEnvironmentsAddons {
	s.Icon = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) SetName(v string) *ListEnvironmentsResponseBodyDataEnvironmentsAddons {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsAddons) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentsResponseBodyDataEnvironmentsFeatures struct {
	// The alias of the feature.
	//
	// example:
	//
	// Prometheus Agent
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The description of the feature.
	//
	// example:
	//
	// Collect Metric data using the Prometheus collection specification
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// http://xxx
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// metirc-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsFeatures) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) GetAlias() *string {
	return s.Alias
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) GetDescription() *string {
	return s.Description
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) GetIcon() *string {
	return s.Icon
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) SetAlias(v string) *ListEnvironmentsResponseBodyDataEnvironmentsFeatures {
	s.Alias = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) SetDescription(v string) *ListEnvironmentsResponseBodyDataEnvironmentsFeatures {
	s.Description = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) SetIcon(v string) *ListEnvironmentsResponseBodyDataEnvironmentsFeatures {
	s.Icon = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) SetName(v string) *ListEnvironmentsResponseBodyDataEnvironmentsFeatures {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsFeatures) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentsResponseBodyDataEnvironmentsTags struct {
	// The tag key.
	//
	// example:
	//
	// fpx-tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// cn-beijing
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataEnvironmentsTags) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsTags) SetKey(v string) *ListEnvironmentsResponseBodyDataEnvironmentsTags {
	s.Key = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsTags) SetValue(v string) *ListEnvironmentsResponseBodyDataEnvironmentsTags {
	s.Value = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataEnvironmentsTags) Validate() error {
	return dara.Validate(s)
}
