// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDataQualityRulesToEvaluationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *AttachDataQualityRulesToEvaluationTaskShrinkRequest
	GetDataQualityEvaluationTaskId() *int64
	SetDataQualityRuleIdsShrink(v string) *AttachDataQualityRulesToEvaluationTaskShrinkRequest
	GetDataQualityRuleIdsShrink() *string
	SetProjectId(v int64) *AttachDataQualityRulesToEvaluationTaskShrinkRequest
	GetProjectId() *int64
}

type AttachDataQualityRulesToEvaluationTaskShrinkRequest struct {
	// The ID of the associated quality check task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The list of data quality rule IDs.
	//
	// This parameter is required.
	DataQualityRuleIdsShrink *string `json:"DataQualityRuleIds,omitempty" xml:"DataQualityRuleIds,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AttachDataQualityRulesToEvaluationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDataQualityRulesToEvaluationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) GetDataQualityRuleIdsShrink() *string {
	return s.DataQualityRuleIdsShrink
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) SetDataQualityEvaluationTaskId(v int64) *AttachDataQualityRulesToEvaluationTaskShrinkRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) SetDataQualityRuleIdsShrink(v string) *AttachDataQualityRulesToEvaluationTaskShrinkRequest {
	s.DataQualityRuleIdsShrink = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) SetProjectId(v int64) *AttachDataQualityRulesToEvaluationTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
