// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListWorkflowDirectoriesResponseBodyData) *ListWorkflowDirectoriesResponseBody
	GetData() []*ListWorkflowDirectoriesResponseBodyData
	SetMaxResults(v string) *ListWorkflowDirectoriesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListWorkflowDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWorkflowDirectoriesResponseBody
	GetTotalCount() *int32
}

type ListWorkflowDirectoriesResponseBody struct {
	Data []*ListWorkflowDirectoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListWorkflowDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowDirectoriesResponseBody) GetData() []*ListWorkflowDirectoriesResponseBodyData {
	return s.Data
}

func (s *ListWorkflowDirectoriesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListWorkflowDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowDirectoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkflowDirectoriesResponseBody) SetData(v []*ListWorkflowDirectoriesResponseBodyData) *ListWorkflowDirectoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowDirectoriesResponseBody) SetMaxResults(v string) *ListWorkflowDirectoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBody) SetNextToken(v string) *ListWorkflowDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBody) SetRequestId(v string) *ListWorkflowDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBody) SetTotalCount(v int32) *ListWorkflowDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowDirectoriesResponseBodyData struct {
	// example:
	//
	// wd-y98v7non5dx****
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// example:
	//
	// 目录名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-y98v7non5dx****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// p-v7n28j0m4ol****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// w-lxyy60mpgpg****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
}

func (s ListWorkflowDirectoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDirectoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowDirectoriesResponseBodyData) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListWorkflowDirectoriesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListWorkflowDirectoriesResponseBodyData) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *ListWorkflowDirectoriesResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListWorkflowDirectoriesResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowDirectoriesResponseBodyData) SetDirectoryId(v string) *ListWorkflowDirectoriesResponseBodyData {
	s.DirectoryId = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBodyData) SetName(v string) *ListWorkflowDirectoriesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBodyData) SetParentDirectoryId(v string) *ListWorkflowDirectoriesResponseBodyData {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBodyData) SetProjectId(v string) *ListWorkflowDirectoriesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBodyData) SetWorkflowId(v string) *ListWorkflowDirectoriesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowDirectoriesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
