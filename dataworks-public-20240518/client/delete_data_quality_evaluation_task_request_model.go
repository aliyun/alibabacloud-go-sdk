// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDataQualityEvaluationTaskRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteDataQualityEvaluationTaskRequest
	GetProjectId() *int64
}

type DeleteDataQualityEvaluationTaskRequest struct {
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 123123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace management page to obtain the ID.
	//
	// This parameter is used to determine the DataWorks workspaces used for this API call.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDataQualityEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityEvaluationTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataQualityEvaluationTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDataQualityEvaluationTaskRequest) SetId(v int64) *DeleteDataQualityEvaluationTaskRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataQualityEvaluationTaskRequest) SetProjectId(v int64) *DeleteDataQualityEvaluationTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDataQualityEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
