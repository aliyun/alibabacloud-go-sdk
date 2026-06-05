// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyWorkflowsResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribeComfyWorkflowsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeComfyWorkflowsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyWorkflowsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeComfyWorkflowsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeComfyWorkflowsResponseBody
	GetTotalCount() *int32
	SetWorkflows(v []*DescribeComfyWorkflowsResponseBodyWorkflows) *DescribeComfyWorkflowsResponseBody
	GetWorkflows() []*DescribeComfyWorkflowsResponseBodyWorkflows
}

type DescribeComfyWorkflowsResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workflows  []*DescribeComfyWorkflowsResponseBodyWorkflows `json:"Workflows,omitempty" xml:"Workflows,omitempty" type:"Repeated"`
}

func (s DescribeComfyWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyWorkflowsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyWorkflowsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyWorkflowsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyWorkflowsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeComfyWorkflowsResponseBody) GetWorkflows() []*DescribeComfyWorkflowsResponseBodyWorkflows {
	return s.Workflows
}

func (s *DescribeComfyWorkflowsResponseBody) SetCode(v int64) *DescribeComfyWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetMessage(v string) *DescribeComfyWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetPageNumber(v int32) *DescribeComfyWorkflowsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetPageSize(v int32) *DescribeComfyWorkflowsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetRequestId(v string) *DescribeComfyWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetTotalCount(v int32) *DescribeComfyWorkflowsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) SetWorkflows(v []*DescribeComfyWorkflowsResponseBodyWorkflows) *DescribeComfyWorkflowsResponseBody {
	s.Workflows = v
	return s
}

func (s *DescribeComfyWorkflowsResponseBody) Validate() error {
	if s.Workflows != nil {
		for _, item := range s.Workflows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComfyWorkflowsResponseBodyWorkflows struct {
	// example:
	//
	// 2026-02-03T07:31:45+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1778897586
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DescribeComfyWorkflowsResponseBodyWorkflows) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyWorkflowsResponseBodyWorkflows) GoString() string {
	return s.String()
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) GetDescription() *string {
	return s.Description
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) GetName() *string {
	return s.Name
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) SetCreationTime(v string) *DescribeComfyWorkflowsResponseBodyWorkflows {
	s.CreationTime = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) SetDescription(v string) *DescribeComfyWorkflowsResponseBodyWorkflows {
	s.Description = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) SetName(v string) *DescribeComfyWorkflowsResponseBodyWorkflows {
	s.Name = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) SetUpdatedTime(v string) *DescribeComfyWorkflowsResponseBodyWorkflows {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) SetWorkflowId(v string) *DescribeComfyWorkflowsResponseBodyWorkflows {
	s.WorkflowId = &v
	return s
}

func (s *DescribeComfyWorkflowsResponseBodyWorkflows) Validate() error {
	return dara.Validate(s)
}
