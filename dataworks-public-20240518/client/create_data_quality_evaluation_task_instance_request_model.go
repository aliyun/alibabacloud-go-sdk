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
	// The ID of the data quality evaluation task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000011
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The execution parameters of the data quality evaluation in the JSON format. The following keys are available:
	//
	// - triggerTime: the millisecond timestamp of the trigger time. This is the base time for the $[yyyymmdd] expression in the data range of data quality monitoring. This key is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "triggerTime": 1733284062000 }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace in which this API operation is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group information. This parameter is required when you run a data quality evaluation on a non-MaxCompute data source.
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
	// The CU consumption configured for the task to run. This parameter must be specified if a serverless resource group is used.
	//
	// example:
	//
	// 0.25
	Cu *float64 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The identifier of the scheduling resource group configured for the task to run.
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
