// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDataQualityRulesFromEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *DetachDataQualityRulesFromEvaluationTaskRequest
	GetDataQualityEvaluationTaskId() *int64
	SetDataQualityRuleIds(v []*int64) *DetachDataQualityRulesFromEvaluationTaskRequest
	GetDataQualityRuleIds() []*int64
	SetProjectId(v int64) *DetachDataQualityRulesFromEvaluationTaskRequest
	GetProjectId() *int64
}

type DetachDataQualityRulesFromEvaluationTaskRequest struct {
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
	DataQualityRuleIds []*int64 `json:"DataQualityRuleIds,omitempty" xml:"DataQualityRuleIds,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DetachDataQualityRulesFromEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDataQualityRulesFromEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) GetDataQualityRuleIds() []*int64 {
	return s.DataQualityRuleIds
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) SetDataQualityEvaluationTaskId(v int64) *DetachDataQualityRulesFromEvaluationTaskRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) SetDataQualityRuleIds(v []*int64) *DetachDataQualityRulesFromEvaluationTaskRequest {
	s.DataQualityRuleIds = v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) SetProjectId(v int64) *DetachDataQualityRulesFromEvaluationTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
