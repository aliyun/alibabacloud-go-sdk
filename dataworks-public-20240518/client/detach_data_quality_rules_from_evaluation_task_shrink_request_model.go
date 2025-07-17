// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDataQualityRulesFromEvaluationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest
	GetDataQualityEvaluationTaskId() *int64
	SetDataQualityRuleIdsShrink(v string) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest
	GetDataQualityRuleIdsShrink() *string
	SetProjectId(v int64) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest
	GetProjectId() *int64
}

type DetachDataQualityRulesFromEvaluationTaskShrinkRequest struct {
	// The ID of the data quality monitoring task that is associated with the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The IDs of the monitoring rules.
	//
	// This parameter is required.
	DataQualityRuleIdsShrink *string `json:"DataQualityRuleIds,omitempty" xml:"DataQualityRuleIds,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DetachDataQualityRulesFromEvaluationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDataQualityRulesFromEvaluationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) GetDataQualityRuleIdsShrink() *string {
	return s.DataQualityRuleIdsShrink
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) SetDataQualityEvaluationTaskId(v int64) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) SetDataQualityRuleIdsShrink(v string) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest {
	s.DataQualityRuleIdsShrink = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) SetProjectId(v int64) *DetachDataQualityRulesFromEvaluationTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
