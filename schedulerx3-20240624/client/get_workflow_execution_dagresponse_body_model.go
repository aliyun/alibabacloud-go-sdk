// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowExecutionDAGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkflowExecutionDAGResponseBody
	GetCode() *int32
	SetData(v *GetWorkflowExecutionDAGResponseBodyData) *GetWorkflowExecutionDAGResponseBody
	GetData() *GetWorkflowExecutionDAGResponseBodyData
	SetMessage(v string) *GetWorkflowExecutionDAGResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkflowExecutionDAGResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkflowExecutionDAGResponseBody
	GetSuccess() *bool
}

type GetWorkflowExecutionDAGResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the request.
	Data *GetWorkflowExecutionDAGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter format error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - `true`: The call succeeded.
	//
	// - `false`: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowExecutionDAGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkflowExecutionDAGResponseBody) GetData() *GetWorkflowExecutionDAGResponseBodyData {
	return s.Data
}

func (s *GetWorkflowExecutionDAGResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowExecutionDAGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowExecutionDAGResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowExecutionDAGResponseBody) SetCode(v int32) *GetWorkflowExecutionDAGResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBody) SetData(v *GetWorkflowExecutionDAGResponseBodyData) *GetWorkflowExecutionDAGResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBody) SetMessage(v string) *GetWorkflowExecutionDAGResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBody) SetRequestId(v string) *GetWorkflowExecutionDAGResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBody) SetSuccess(v bool) *GetWorkflowExecutionDAGResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowExecutionDAGResponseBodyData struct {
	// A list of edges in the workflow DAG.
	Edges []*GetWorkflowExecutionDAGResponseBodyDataEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// A list of nodes in the workflow DAG.
	Nodes []*GetWorkflowExecutionDAGResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkflowExecutionDAGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponseBodyData) GetEdges() []*GetWorkflowExecutionDAGResponseBodyDataEdges {
	return s.Edges
}

func (s *GetWorkflowExecutionDAGResponseBodyData) GetNodes() []*GetWorkflowExecutionDAGResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetWorkflowExecutionDAGResponseBodyData) SetEdges(v []*GetWorkflowExecutionDAGResponseBodyDataEdges) *GetWorkflowExecutionDAGResponseBodyData {
	s.Edges = v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyData) SetNodes(v []*GetWorkflowExecutionDAGResponseBodyDataNodes) *GetWorkflowExecutionDAGResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyData) Validate() error {
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

type GetWorkflowExecutionDAGResponseBodyDataEdges struct {
	// The ID of the source node.
	//
	// example:
	//
	// 1000
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the target node.
	//
	// example:
	//
	// 1001
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkflowExecutionDAGResponseBodyDataEdges) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponseBodyDataEdges) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponseBodyDataEdges) GetSource() *string {
	return s.Source
}

func (s *GetWorkflowExecutionDAGResponseBodyDataEdges) GetTarget() *string {
	return s.Target
}

func (s *GetWorkflowExecutionDAGResponseBodyDataEdges) SetSource(v string) *GetWorkflowExecutionDAGResponseBodyDataEdges {
	s.Source = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataEdges) SetTarget(v string) *GetWorkflowExecutionDAGResponseBodyDataEdges {
	s.Target = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataEdges) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowExecutionDAGResponseBodyDataNodes struct {
	// The application name.
	//
	// example:
	//
	// cua-xxl-job-executor
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The coordinates of the node.
	Coordinate *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate `json:"Coordinate,omitempty" xml:"Coordinate,omitempty" type:"Struct"`
	// The unique ID for the job execution (node).
	//
	// example:
	//
	// 7491777526619542799
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 9
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job type.
	//
	// example:
	//
	// script_shell
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The job name.
	//
	// example:
	//
	// job1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution status of the job. Valid values:
	//
	// - 0: Unknown
	//
	// - 1: Waiting
	//
	// - 2: Queued
	//
	// - 3: Running
	//
	// - 4: Succeeded
	//
	// - 5: Failed
	//
	// - 6: Killed
	//
	// - 7: Held
	//
	// - 8: Marked as successful
	//
	// - 9: Skipped
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkflowExecutionDAGResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetCoordinate() *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate {
	return s.Coordinate
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetId() *string {
	return s.Id
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetJobId() *int64 {
	return s.JobId
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetJobType() *string {
	return s.JobType
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetAppName(v string) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.AppName = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetCoordinate(v *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.Coordinate = v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetId(v string) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.Id = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetJobId(v int64) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.JobId = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetJobType(v string) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.JobType = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetName(v string) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) SetStatus(v int32) *GetWorkflowExecutionDAGResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodes) Validate() error {
	if s.Coordinate != nil {
		if err := s.Coordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate struct {
	// The height of the node. This parameter is optional.
	//
	// example:
	//
	// 20
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The width of the node. This parameter is optional.
	//
	// example:
	//
	// 100
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate of the node.
	//
	// example:
	//
	// 50
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the node.
	//
	// example:
	//
	// 50
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) GetHeight() *float32 {
	return s.Height
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) GetWidth() *float32 {
	return s.Width
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) GetX() *float32 {
	return s.X
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) GetY() *float32 {
	return s.Y
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) SetHeight(v float32) *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate {
	s.Height = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) SetWidth(v float32) *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate {
	s.Width = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) SetX(v float32) *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate {
	s.X = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) SetY(v float32) *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate {
	s.Y = &v
	return s
}

func (s *GetWorkflowExecutionDAGResponseBodyDataNodesCoordinate) Validate() error {
	return dara.Validate(s)
}
