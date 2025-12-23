// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListWorkflowsResponseBodyData) *ListWorkflowsResponseBody
	GetData() []*ListWorkflowsResponseBodyData
	SetNextToken(v string) *ListWorkflowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowsResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListWorkflowsResponseBody
	GetTotalSize() *int32
}

type ListWorkflowsResponseBody struct {
	Data []*ListWorkflowsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBody) GetData() []*ListWorkflowsResponseBodyData {
	return s.Data
}

func (s *ListWorkflowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListWorkflowsResponseBody) SetData(v []*ListWorkflowsResponseBodyData) *ListWorkflowsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowsResponseBody) SetNextToken(v string) *ListWorkflowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetRequestId(v string) *ListWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowsResponseBody) SetTotalSize(v int32) *ListWorkflowsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListWorkflowsResponseBody) Validate() error {
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

type ListWorkflowsResponseBodyData struct {
	// example:
	//
	// 2024-01-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// workflow description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workflow_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// wd-3q9jo749ne5****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
}

func (s ListWorkflowsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkflowsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsResponseBodyData) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *ListWorkflowsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListWorkflowsResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowsResponseBodyData) SetCreateTime(v string) *ListWorkflowsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetDescription(v string) *ListWorkflowsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetName(v string) *ListWorkflowsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetParentDirectoryId(v string) *ListWorkflowsResponseBodyData {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetUpdateTime(v string) *ListWorkflowsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) SetWorkflowId(v string) *ListWorkflowsResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
