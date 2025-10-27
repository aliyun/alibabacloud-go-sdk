// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkFlowResponseBody
	GetCode() *int32
	SetData(v *GetWorkFlowResponseBodyData) *GetWorkFlowResponseBody
	GetData() *GetWorkFlowResponseBodyData
	SetMessage(v string) *GetWorkFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkFlowResponseBody
	GetSuccess() *bool
}

type GetWorkFlowResponseBody struct {
	// Error codes
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the workflow.
	Data *GetWorkFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// workflow is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45678xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the API call.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkFlowResponseBody) GetData() *GetWorkFlowResponseBodyData {
	return s.Data
}

func (s *GetWorkFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkFlowResponseBody) SetCode(v int32) *GetWorkFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetData(v *GetWorkFlowResponseBodyData) *GetWorkFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkFlowResponseBody) SetMessage(v string) *GetWorkFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetRequestId(v string) *GetWorkFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkFlowResponseBody) SetSuccess(v bool) *GetWorkFlowResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkFlowResponseBodyData struct {
	// The basic information of the workflow.
	WorkFlowInfo *GetWorkFlowResponseBodyDataWorkFlowInfo `json:"WorkFlowInfo,omitempty" xml:"WorkFlowInfo,omitempty" type:"Struct"`
	// The node information of the workflow.
	WorkFlowNodeInfo *GetWorkFlowResponseBodyDataWorkFlowNodeInfo `json:"WorkFlowNodeInfo,omitempty" xml:"WorkFlowNodeInfo,omitempty" type:"Struct"`
}

func (s GetWorkFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyData) GetWorkFlowInfo() *GetWorkFlowResponseBodyDataWorkFlowInfo {
	return s.WorkFlowInfo
}

func (s *GetWorkFlowResponseBodyData) GetWorkFlowNodeInfo() *GetWorkFlowResponseBodyDataWorkFlowNodeInfo {
	return s.WorkFlowNodeInfo
}

func (s *GetWorkFlowResponseBodyData) SetWorkFlowInfo(v *GetWorkFlowResponseBodyDataWorkFlowInfo) *GetWorkFlowResponseBodyData {
	s.WorkFlowInfo = v
	return s
}

func (s *GetWorkFlowResponseBodyData) SetWorkFlowNodeInfo(v *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) *GetWorkFlowResponseBodyData {
	s.WorkFlowNodeInfo = v
	return s
}

func (s *GetWorkFlowResponseBodyData) Validate() error {
	if s.WorkFlowInfo != nil {
		if err := s.WorkFlowInfo.Validate(); err != nil {
			return err
		}
	}
	if s.WorkFlowNodeInfo != nil {
		if err := s.WorkFlowNodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkFlowResponseBodyDataWorkFlowInfo struct {
	// The description of the workflow.
	//
	// example:
	//
	// my first workflow
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// workflow_111
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time expression of the workflow.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type of the workflow.
	//
	// example:
	//
	// cron
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 1234xxx
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowInfo) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetDescription() *string {
	return s.Description
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetName() *string {
	return s.Name
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetStatus() *string {
	return s.Status
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetTimeType() *string {
	return s.TimeType
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetDescription(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Description = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetGroupId(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.GroupId = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetMaxConcurrency(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.MaxConcurrency = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetName(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Name = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetNamespace(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Namespace = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetStatus(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.Status = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetTimeExpression(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.TimeExpression = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetTimeType(v string) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.TimeType = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) SetWorkflowId(v int64) *GetWorkFlowResponseBodyDataWorkFlowInfo {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowInfo) Validate() error {
	return dara.Validate(s)
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfo struct {
	// The workflow edges.
	Edges []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The list of workflow nodes.
	Nodes []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfo) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) GetEdges() []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges {
	return s.Edges
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) GetNodes() []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	return s.Nodes
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) SetEdges(v []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) *GetWorkFlowResponseBodyDataWorkFlowNodeInfo {
	s.Edges = v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) SetNodes(v []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) *GetWorkFlowResponseBodyDataWorkFlowNodeInfo {
	s.Nodes = v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfo) Validate() error {
	if s.Edges != nil {
		for _, item := range s.Edges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges struct {
	// The ID of the source job.
	//
	// example:
	//
	// 100
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the object job.
	//
	// example:
	//
	// 200
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) GetSource() *int64 {
	return s.Source
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) GetTarget() *int64 {
	return s.Target
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) SetSource(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges {
	s.Source = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) SetTarget(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges {
	s.Target = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges) Validate() error {
	return dara.Validate(s)
}

type GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes struct {
	// The ID of the job.
	//
	// example:
	//
	// 123456xxx
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// job_111
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) GoString() string {
	return s.String()
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) GetId() *int64 {
	return s.Id
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) GetLabel() *string {
	return s.Label
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetId(v int64) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Id = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetLabel(v string) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Label = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) SetStatus(v int32) *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes {
	s.Status = &v
	return s
}

func (s *GetWorkFlowResponseBodyDataWorkFlowNodeInfoNodes) Validate() error {
	return dara.Validate(s)
}
