// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDataQualityRulesToEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *AttachDataQualityRulesToEvaluationTaskRequest
	GetDataQualityEvaluationTaskId() *int64
	SetDataQualityRuleIds(v []*int64) *AttachDataQualityRulesToEvaluationTaskRequest
	GetDataQualityRuleIds() []*int64
	SetProjectId(v int64) *AttachDataQualityRulesToEvaluationTaskRequest
	GetProjectId() *int64
}

type AttachDataQualityRulesToEvaluationTaskRequest struct {
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
	DataQualityRuleIds []*int64 `json:"DataQualityRuleIds,omitempty" xml:"DataQualityRuleIds,omitempty" type:"Repeated"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AttachDataQualityRulesToEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDataQualityRulesToEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) GetDataQualityRuleIds() []*int64 {
	return s.DataQualityRuleIds
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) SetDataQualityEvaluationTaskId(v int64) *AttachDataQualityRulesToEvaluationTaskRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) SetDataQualityRuleIds(v []*int64) *AttachDataQualityRulesToEvaluationTaskRequest {
	s.DataQualityRuleIds = v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) SetProjectId(v int64) *AttachDataQualityRulesToEvaluationTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
