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
	// The ID of the data quality monitoring task that is associated with the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The IDs of the monitoring rules.
	//
	// This parameter is required.
	DataQualityRuleIdsShrink *string `json:"DataQualityRuleIds,omitempty" xml:"DataQualityRuleIds,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
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
