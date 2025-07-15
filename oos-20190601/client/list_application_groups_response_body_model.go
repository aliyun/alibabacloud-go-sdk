// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGroups(v []*ListApplicationGroupsResponseBodyApplicationGroups) *ListApplicationGroupsResponseBody
	GetApplicationGroups() []*ListApplicationGroupsResponseBodyApplicationGroups
	SetMaxResults(v int32) *ListApplicationGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationGroupsResponseBody
	GetRequestId() *string
}

type ListApplicationGroupsResponseBody struct {
	// The details of the application group.
	ApplicationGroups []*ListApplicationGroupsResponseBodyApplicationGroups `json:"ApplicationGroups,omitempty" xml:"ApplicationGroups,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69D97BF2-5DF2-544C-A650-36A474E17BC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBody) GetApplicationGroups() []*ListApplicationGroupsResponseBodyApplicationGroups {
	return s.ApplicationGroups
}

func (s *ListApplicationGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationGroupsResponseBody) SetApplicationGroups(v []*ListApplicationGroupsResponseBodyApplicationGroups) *ListApplicationGroupsResponseBody {
	s.ApplicationGroups = v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetMaxResults(v int32) *ListApplicationGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetNextToken(v string) *ListApplicationGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetRequestId(v string) *ListApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationGroupsResponseBodyApplicationGroups struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 12345678
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The configuration information of the application group.
	//
	// example:
	//
	// {   "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",   "Parameters": {     "ZoneId": "cn-hangzhou-k",     "ProjectName": "test",     "SystemDiskSize": 40,     "InstanceChargeType": "PostPaid",     "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",     "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",     "SystemDiskCategory": "cloud_essd",     "InstancePassword": "******",     "InternetChargeType": "PayByTraffic",     "InstanceCount": 1,     "InternetMaxBandwidthOut": 0,     "VpcId": "vpc-bp1i99boyas8i8m9t3skp",     "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",     "DataDiskSize": 100,     "EcsInstanceType": "ecs.s6-c1m4.small",     "DataDiskCategory": "cloud_efficiency",     "EnvironmentCommandId": "c-hz028fc3g031gcg"   },   "RegionId": "cn-hangzhou",   "StackName": "stack-1645688523068-3no_AKhOJ",   "DisableRollback": true }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The ID of the region in which the related resources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId      *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	DeployedRevisionIds *string `json:"DeployedRevisionIds,omitempty" xml:"DeployedRevisionIds,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorDetail *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	ErrorType   *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the application group.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The state information of the application group.
	//
	// example:
	//
	// ApplicationGroup is Created.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-08T03:01:53Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetCmsGroupId() *string {
	return s.CmsGroupId
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetDeployParameters() *string {
	return s.DeployParameters
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetDeployedRevisionIds() *string {
	return s.DeployedRevisionIds
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetErrorType() *string {
	return s.ErrorType
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetImportTagKey() *string {
	return s.ImportTagKey
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetImportTagValue() *string {
	return s.ImportTagValue
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetName() *string {
	return s.Name
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetApplicationName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCmsGroupId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CmsGroupId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCreateDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployParameters(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployParameters = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployRegionId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployedRevisionIds(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployedRevisionIds = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDescription(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Description = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetErrorDetail(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ErrorDetail = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetErrorType(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ErrorType = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetImportTagKey(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ImportTagKey = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetImportTagValue(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ImportTagValue = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Name = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetStatus(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Status = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetStatusReason(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.StatusReason = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetUpdateDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) Validate() error {
	return dara.Validate(s)
}
