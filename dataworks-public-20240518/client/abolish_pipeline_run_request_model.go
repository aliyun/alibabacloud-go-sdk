// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *AbolishPipelineRunRequest
	GetId() *string
	SetProjectId(v int64) *AbolishPipelineRunRequest
	GetProjectId() *int64
}

type AbolishPipelineRunRequest struct {
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606087c-9ac4-43f0-83a8-0b5ced21XXXX
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

func (s AbolishPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s AbolishPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *AbolishPipelineRunRequest) GetId() *string {
	return s.Id
}

func (s *AbolishPipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AbolishPipelineRunRequest) SetId(v string) *AbolishPipelineRunRequest {
	s.Id = &v
	return s
}

func (s *AbolishPipelineRunRequest) SetProjectId(v int64) *AbolishPipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *AbolishPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
