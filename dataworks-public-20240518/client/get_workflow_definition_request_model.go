// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetWorkflowDefinitionRequest
	GetId() *int64
	SetIncludeScriptContent(v bool) *GetWorkflowDefinitionRequest
	GetIncludeScriptContent() *bool
	SetProjectId(v int64) *GetWorkflowDefinitionRequest
	GetProjectId() *int64
}

type GetWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 查询结果是否包含工作流内部节点的脚本内容（对于内容较多的节点，可能存在较长的网络传输延时）。
	//
	// example:
	//
	// false
	IncludeScriptContent *bool `json:"IncludeScriptContent,omitempty" xml:"IncludeScriptContent,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionRequest) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowDefinitionRequest) GetIncludeScriptContent() *bool {
	return s.IncludeScriptContent
}

func (s *GetWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetWorkflowDefinitionRequest) SetId(v int64) *GetWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) SetIncludeScriptContent(v bool) *GetWorkflowDefinitionRequest {
	s.IncludeScriptContent = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) SetProjectId(v int64) *GetWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
