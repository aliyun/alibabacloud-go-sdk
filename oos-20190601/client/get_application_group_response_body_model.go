// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGroup(v *GetApplicationGroupResponseBodyApplicationGroup) *GetApplicationGroupResponseBody
	GetApplicationGroup() *GetApplicationGroupResponseBodyApplicationGroup
	SetRequestId(v string) *GetApplicationGroupResponseBody
	GetRequestId() *string
}

type GetApplicationGroupResponseBody struct {
	// The information about the application group.
	ApplicationGroup *GetApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 92EA60ED-544D-55E9-93D9-237BE9976C0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBody) GetApplicationGroup() *GetApplicationGroupResponseBodyApplicationGroup {
	return s.ApplicationGroup
}

func (s *GetApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationGroupResponseBody) SetApplicationGroup(v *GetApplicationGroupResponseBodyApplicationGroup) *GetApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *GetApplicationGroupResponseBody) SetRequestId(v string) *GetApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationGroupResponseBody) Validate() error {
	if s.ApplicationGroup != nil {
		if err := s.ApplicationGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationGroupResponseBodyApplicationGroup struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The source of application.
	//
	// example:
	//
	// {"Platform":"github","RepoName":"wenle/springboot-ecs-sourcecode-demo","Owner":"wenle","Branch":"main","CommitHash":"8559ff3ac7568fc7951ff63f841883ee3f06c6fe","CommitMessage":"Init computenest project"}
	ApplicationSource *string `json:"ApplicationSource,omitempty" xml:"ApplicationSource,omitempty"`
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
	// The output of the deployment result.
	//
	// example:
	//
	// {       "Outputs": [         {           "Description": "No description given",           "OutputKey": "InstanceIds"         }       ],       "StackId": "6ef4b860-f6e7-4145-8d22-f4a2a32363e1"     }   }
	DeployOutputs *string `json:"DeployOutputs,omitempty" xml:"DeployOutputs,omitempty"`
	// The configuration information of the application group.
	//
	// example:
	//
	// {       "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",       "Parameters": {         "ZoneId": "cn-hangzhou-k",         "ProjectName": "test",         "SystemDiskSize": 40,         "InstanceChargeType": "PostPaid",         "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",         "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",         "SystemDiskCategory": "cloud_essd",         "InstancePassword": "******",         "InternetChargeType": "PayByTraffic",         "InstanceCount": 1,         "InternetMaxBandwidthOut": 0,         "VpcId": "vpc-bp1i99boyas8i8m9t3skp",         "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",         "DataDiskSize": 100,         "EcsInstanceType": "ecs.s6-c1m4.small",         "DataDiskCategory": "cloud_efficiency",         "EnvironmentCommandId": "c-hz028fc3g031gcg"       }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The ID of the region in which you deploy the application group.
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
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
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
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\\\"PrometheusConfigMap\\\\":{\\\\"Template 1\\\\":{\\\\"EnablePrometheus\\\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The creation progress of the application instance.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
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
	// Stack CREATE completed successfully
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the application group was last modified.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationGroupResponseBodyApplicationGroup) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetApplicationSource() *string {
	return s.ApplicationSource
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetCmsGroupId() *string {
	return s.CmsGroupId
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetDeployOutputs() *string {
	return s.DeployOutputs
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetDeployParameters() *string {
	return s.DeployParameters
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetDeployedRevisionIds() *string {
	return s.DeployedRevisionIds
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetErrorType() *string {
	return s.ErrorType
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetImportTagKey() *string {
	return s.ImportTagKey
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetImportTagValue() *string {
	return s.ImportTagValue
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetName() *string {
	return s.Name
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetProgress() *string {
	return s.Progress
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetApplicationSource(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationSource = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCmsGroupId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CmsGroupId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCreateDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployOutputs(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployOutputs = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployParameters(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployParameters = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployedRevisionIds(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployedRevisionIds = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetErrorDetail(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ErrorDetail = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetErrorType(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ErrorType = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetExecutionId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ExecutionId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetOperationMetadata(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.OperationMetadata = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetProgress(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Progress = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetStatus(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Status = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetStatusReason(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.StatusReason = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetUpdateDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.UpdateDate = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) Validate() error {
	return dara.Validate(s)
}
