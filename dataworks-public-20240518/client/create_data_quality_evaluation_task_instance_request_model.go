// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *CreateDataQualityEvaluationTaskInstanceRequest
	GetDataQualityEvaluationTaskId() *int64
	SetParameters(v string) *CreateDataQualityEvaluationTaskInstanceRequest
	GetParameters() *string
	SetProjectId(v int64) *CreateDataQualityEvaluationTaskInstanceRequest
	GetProjectId() *int64
	SetRuntimeResource(v *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) *CreateDataQualityEvaluationTaskInstanceRequest
	GetRuntimeResource() *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource
}

type CreateDataQualityEvaluationTaskInstanceRequest struct {
	// The ID of the data quality monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// Data quality verification execution parameters in JSON format. The available keys are as follows:
	//
	// - triggerTime: the millisecond timestamp of the trigger time. The baseline time of the $[yyyymmdd] expression in the data range of data quality monitoring. Required.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "triggerTime": 1733284062000 }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace management page to obtain the ID.
	//
	// This parameter is used to determine the DataWorks workspaces used for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Resource Group information, which must be filled in when running non-MaxCompute data quality verification.
	RuntimeResource *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s CreateDataQualityEvaluationTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) GetRuntimeResource() *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) SetDataQualityEvaluationTaskId(v int64) *CreateDataQualityEvaluationTaskInstanceRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) SetParameters(v string) *CreateDataQualityEvaluationTaskInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) SetProjectId(v int64) *CreateDataQualityEvaluationTaskInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) SetRuntimeResource(v *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) *CreateDataQualityEvaluationTaskInstanceRequest {
	s.RuntimeResource = v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequest) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource struct {
	// The task runs to configure CU consumption. If Serverless resource groups are used, you must specify this parameter.
	//
	// example:
	//
	// 0.25
	Cu *float64 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The identifier of the scheduling resource group configured for running the task.
	//
	// example:
	//
	// 63900680
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) GetCu() *float64 {
	return s.Cu
}

func (s *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) SetCu(v float64) *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) SetResourceGroupId(v string) *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}
