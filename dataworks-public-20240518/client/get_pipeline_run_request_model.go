// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetPipelineRunRequest
	GetId() *string
	SetProjectId(v int64) *GetPipelineRunRequest
	GetProjectId() *int64
}

type GetPipelineRunRequest struct {
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRunRequest) GetId() *string {
	return s.Id
}

func (s *GetPipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPipelineRunRequest) SetId(v string) *GetPipelineRunRequest {
	s.Id = &v
	return s
}

func (s *GetPipelineRunRequest) SetProjectId(v int64) *GetPipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
