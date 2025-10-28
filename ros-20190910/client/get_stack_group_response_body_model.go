// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackGroupResponseBody
	GetRequestId() *string
	SetStackGroup(v *GetStackGroupResponseBodyStackGroup) *GetStackGroupResponseBody
	GetStackGroup() *GetStackGroupResponseBodyStackGroup
}

type GetStackGroupResponseBody struct {
	// The details of the stack group.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the stack group.
	StackGroup *GetStackGroupResponseBodyStackGroup `json:"StackGroup,omitempty" xml:"StackGroup,omitempty" type:"Struct"`
}

func (s GetStackGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackGroupResponseBody) GetStackGroup() *GetStackGroupResponseBodyStackGroup {
	return s.StackGroup
}

func (s *GetStackGroupResponseBody) SetRequestId(v string) *GetStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupResponseBody) SetStackGroup(v *GetStackGroupResponseBodyStackGroup) *GetStackGroupResponseBody {
	s.StackGroup = v
	return s
}

func (s *GetStackGroupResponseBody) Validate() error {
	if s.StackGroup != nil {
		if err := s.StackGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackGroupResponseBodyStackGroup struct {
	// The parameters of the stack group.
	//
	// example:
	//
	// AliyunROSStackGroupAdministrationRole
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// 	- true: Automatic deployment is enabled. If a member account is added to the folder to which the stack group belongs after automatic deployment is enabled, the stack group deploys its stack instances in the specified region where the added account is deployed. If the account is deleted from the folder, the stack instances in the specified region are deleted from the stack group.
	//
	// 	- false: Automatic deployment is disabled. After automatic deployment is disabled, the stack instances remain unchanged when the member account in the folder is changed.
	AutoDeployment *GetStackGroupResponseBodyStackGroupAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	CreateTime     *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the stack group.
	//
	// example:
	//
	// StackGroup Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template body.
	//
	// example:
	//
	// AliyunROSStackGroupExecutionRole
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The key of the parameter.
	Parameters []*GetStackGroupResponseBodyStackGroupParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The information about automatic deployment settings.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	//
	// example:
	//
	// SELF_MANAGED
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The folder IDs of the resource directory. This parameter is used to deploy stack instances within all the accounts in the folders.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
	// The permission model.
	//
	// Valid values:
	//
	// 	- SELF_MANAGED: the self-managed permission model
	//
	// 	- SERVICE_MANAGED: the service-managed permission model
	//
	// >  For more information about the permission models of stack groups, see [Overview](https://help.aliyun.com/document_detail/154578.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time when drift detection was performed on the stack group.
	StackGroupDriftDetectionDetail *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
	// The status of the stack group.
	//
	// Valid values:
	//
	// 	- ACTIVE
	//
	// 	- DELETED
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the RAM role that is specified for the execution account when you create the self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role. If this parameter is not specified, the default value AliyunROSStackGroupExecutionRole is returned.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The name of the RAM role that is specified for the administrator account in Resource Orchestration Service (ROS) when you create the self-managed stack group. If this parameter is not specified, the default value AliyunROSStackGroupAdministrationRole is returned.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The structure that contains the template body.
	//
	// > We recommend that you use TemplateContent instead of TemplateBody.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The JSON-formatted structure that contains the template body. For more information, see [Template syntax](https://help.aliyun.com/document_detail/28857.html).
	//
	// example:
	//
	// {
	//
	//       "ROSTemplateFormatVersion": "2015-09-01"
	//
	// }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroup) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroup) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroup) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *GetStackGroupResponseBodyStackGroup) GetAutoDeployment() *GetStackGroupResponseBodyStackGroupAutoDeployment {
	return s.AutoDeployment
}

func (s *GetStackGroupResponseBodyStackGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackGroupResponseBodyStackGroup) GetDescription() *string {
	return s.Description
}

func (s *GetStackGroupResponseBodyStackGroup) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *GetStackGroupResponseBodyStackGroup) GetParameters() []*GetStackGroupResponseBodyStackGroupParameters {
	return s.Parameters
}

func (s *GetStackGroupResponseBodyStackGroup) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *GetStackGroupResponseBodyStackGroup) GetRdFolderIds() []*string {
	return s.RdFolderIds
}

func (s *GetStackGroupResponseBodyStackGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetStackGroupResponseBodyStackGroup) GetStackGroupDriftDetectionDetail() *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	return s.StackGroupDriftDetectionDetail
}

func (s *GetStackGroupResponseBodyStackGroup) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *GetStackGroupResponseBodyStackGroup) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetStackGroupResponseBodyStackGroup) GetStatus() *string {
	return s.Status
}

func (s *GetStackGroupResponseBodyStackGroup) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetStackGroupResponseBodyStackGroup) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetStackGroupResponseBodyStackGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetStackGroupResponseBodyStackGroup) SetAdministrationRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.AdministrationRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetAutoDeployment(v *GetStackGroupResponseBodyStackGroupAutoDeployment) *GetStackGroupResponseBodyStackGroup {
	s.AutoDeployment = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetCreateTime(v string) *GetStackGroupResponseBodyStackGroup {
	s.CreateTime = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetDescription(v string) *GetStackGroupResponseBodyStackGroup {
	s.Description = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetExecutionRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetParameters(v []*GetStackGroupResponseBodyStackGroupParameters) *GetStackGroupResponseBodyStackGroup {
	s.Parameters = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetPermissionModel(v string) *GetStackGroupResponseBodyStackGroup {
	s.PermissionModel = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetRdFolderIds(v []*string) *GetStackGroupResponseBodyStackGroup {
	s.RdFolderIds = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetResourceGroupId(v string) *GetStackGroupResponseBodyStackGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupDriftDetectionDetail(v *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupDriftDetectionDetail = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupId(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupName(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStatus(v string) *GetStackGroupResponseBodyStackGroup {
	s.Status = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetTemplateBody(v string) *GetStackGroupResponseBodyStackGroup {
	s.TemplateBody = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetTemplateContent(v string) *GetStackGroupResponseBodyStackGroup {
	s.TemplateContent = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetUpdateTime(v string) *GetStackGroupResponseBodyStackGroup {
	s.UpdateTime = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) Validate() error {
	if s.AutoDeployment != nil {
		if err := s.AutoDeployment.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StackGroupDriftDetectionDetail != nil {
		if err := s.StackGroupDriftDetectionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackGroupResponseBodyStackGroupAutoDeployment struct {
	// Indicates whether stacks in the member account are retained when the member account is deleted from the folder.
	//
	// Valid values:
	//
	// 	- true: The stacks are retained.
	//
	// 	- false: The stacks are deleted.
	//
	// >  This parameter is returned only when the Enabled parameter is set to true.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The folder IDs of the resource directory. This parameter is used to deploy stack instances within all the accounts in the folders.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	//
	// example:
	//
	// true
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupAutoDeployment) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupAutoDeployment) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) GetRetainStacksOnAccountRemoval() *bool {
	return s.RetainStacksOnAccountRemoval
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) SetEnabled(v bool) *GetStackGroupResponseBodyStackGroupAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *GetStackGroupResponseBodyStackGroupAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupResponseBodyStackGroupParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupParameters) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupParameters) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetStackGroupResponseBodyStackGroupParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterKey(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterValue(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupParameters) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail struct {
	// The number of stack instances that have drifted.
	//
	// example:
	//
	// 0
	CancelledStackInstancesCount *int32 `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	// The drift status of the stack group.
	//
	// Valid values:
	//
	// 	- DRIFTED: At least one stack instance in the stack group has drifted.
	//
	// 	- NOT_CHECKED: No drift detection is completed on the stack group.
	//
	// 	- IN_SYNC: All the stack instances in the stack group are being synchronized.
	//
	// example:
	//
	// COMPLETED
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The number of stack instances.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The ID of the resource group. This parameter is specified when you create the stack group.
	//
	// example:
	//
	// 1
	DriftedStackInstancesCount *int32 `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
	// The status of drift detection on the stack group.
	//
	// Valid values:
	//
	// 	- COMPLETED: Drift detection is performed and completed on all stack instances.
	//
	// 	- FAILED: Drift detection is performed. The number of stack instances that failed the drift detection exceeds the specified threshold.
	//
	// 	- PARTIAL_SUCCESS: Drift detection is performed. The number of stack instances that failed the drift detection does not exceed the specified threshold.
	//
	// 	- IN_PROGRESS: Drift detection is being performed on the stack group.
	//
	// 	- STOPPED: Drift detection is canceled for the stack group.
	//
	// example:
	//
	// 0
	FailedStackInstancesCount *int32 `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	// The number of stack instances that were being synchronized.
	//
	// example:
	//
	// 0
	InProgressStackInstancesCount *int32 `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	// The number of stack instances for which drift detection was canceled.
	//
	// example:
	//
	// 1
	InSyncStackInstancesCount *int32 `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	// The number of stack instances on which drift detection was being performed.
	//
	// example:
	//
	// DRIFTED
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The number of stack instances that failed drift detection.
	//
	// example:
	//
	// 2
	TotalStackInstancesCount *int32 `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetCancelledStackInstancesCount() *int32 {
	return s.CancelledStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetDriftDetectionStatus() *string {
	return s.DriftDetectionStatus
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetDriftedStackInstancesCount() *int32 {
	return s.DriftedStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetFailedStackInstancesCount() *int32 {
	return s.FailedStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetInProgressStackInstancesCount() *int32 {
	return s.InProgressStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetInSyncStackInstancesCount() *int32 {
	return s.InSyncStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetStackGroupDriftStatus() *string {
	return s.StackGroupDriftStatus
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GetTotalStackInstancesCount() *int32 {
	return s.TotalStackInstancesCount
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) Validate() error {
	return dara.Validate(s)
}
