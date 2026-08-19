// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *CreateDataQualityEvaluationTaskInstanceShrinkRequest
	GetDataQualityEvaluationTaskId() *int64
	SetParameters(v string) *CreateDataQualityEvaluationTaskInstanceShrinkRequest
	GetParameters() *string
	SetProjectId(v int64) *CreateDataQualityEvaluationTaskInstanceShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *CreateDataQualityEvaluationTaskInstanceShrinkRequest
	GetRuntimeResourceShrink() *string
}

type CreateDataQualityEvaluationTaskInstanceShrinkRequest struct {
	// The ID of the data quality check task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000011
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The execution parameters of the data quality check in JSON format. The following keys are available:
	//
	// - triggerTime: the timestamp in milliseconds of the trigger time. This value is used as the base time for the $[yyyymmdd] expression in the data range of the data quality monitoring task. This key is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "triggerTime": 1733284062000 }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks workspace. You can logon to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Storage Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API invoke operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group information. This parameter is required when running data quality checks on non-MaxCompute data.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
}

func (s CreateDataQualityEvaluationTaskInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) SetDataQualityEvaluationTaskId(v int64) *CreateDataQualityEvaluationTaskInstanceShrinkRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) SetParameters(v string) *CreateDataQualityEvaluationTaskInstanceShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) SetProjectId(v int64) *CreateDataQualityEvaluationTaskInstanceShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) SetRuntimeResourceShrink(v string) *CreateDataQualityEvaluationTaskInstanceShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
