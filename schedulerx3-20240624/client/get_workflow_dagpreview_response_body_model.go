// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkflowDAGPreviewResponseBody
	GetCode() *int32
	SetData(v *GetWorkflowDAGPreviewResponseBodyData) *GetWorkflowDAGPreviewResponseBody
	GetData() *GetWorkflowDAGPreviewResponseBodyData
	SetMessage(v string) *GetWorkflowDAGPreviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkflowDAGPreviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkflowDAGPreviewResponseBody
	GetSuccess() *bool
}

type GetWorkflowDAGPreviewResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetWorkflowDAGPreviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B8733786-C045-59F1-8D79-99A52863F62D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowDAGPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkflowDAGPreviewResponseBody) GetData() *GetWorkflowDAGPreviewResponseBodyData {
	return s.Data
}

func (s *GetWorkflowDAGPreviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowDAGPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowDAGPreviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowDAGPreviewResponseBody) SetCode(v int32) *GetWorkflowDAGPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBody) SetData(v *GetWorkflowDAGPreviewResponseBodyData) *GetWorkflowDAGPreviewResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBody) SetMessage(v string) *GetWorkflowDAGPreviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBody) SetRequestId(v string) *GetWorkflowDAGPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBody) SetSuccess(v bool) *GetWorkflowDAGPreviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowDAGPreviewResponseBodyData struct {
	// The list of edges in the workflow.
	Edges []*GetWorkflowDAGPreviewResponseBodyDataEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The list of workflow nodes.
	Nodes []*GetWorkflowDAGPreviewResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkflowDAGPreviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponseBodyData) GetEdges() []*GetWorkflowDAGPreviewResponseBodyDataEdges {
	return s.Edges
}

func (s *GetWorkflowDAGPreviewResponseBodyData) GetNodes() []*GetWorkflowDAGPreviewResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetWorkflowDAGPreviewResponseBodyData) SetEdges(v []*GetWorkflowDAGPreviewResponseBodyDataEdges) *GetWorkflowDAGPreviewResponseBodyData {
	s.Edges = v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyData) SetNodes(v []*GetWorkflowDAGPreviewResponseBodyDataNodes) *GetWorkflowDAGPreviewResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyData) Validate() error {
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

type GetWorkflowDAGPreviewResponseBodyDataEdges struct {
	// The ID of the source node of the edge.
	//
	// example:
	//
	// 3
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the target node of the edge.
	//
	// example:
	//
	// 4
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkflowDAGPreviewResponseBodyDataEdges) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponseBodyDataEdges) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponseBodyDataEdges) GetSource() *int64 {
	return s.Source
}

func (s *GetWorkflowDAGPreviewResponseBodyDataEdges) GetTarget() *int64 {
	return s.Target
}

func (s *GetWorkflowDAGPreviewResponseBodyDataEdges) SetSource(v int64) *GetWorkflowDAGPreviewResponseBodyDataEdges {
	s.Source = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataEdges) SetTarget(v int64) *GetWorkflowDAGPreviewResponseBodyDataEdges {
	s.Target = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataEdges) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowDAGPreviewResponseBodyDataNodes struct {
	// The application name.
	//
	// example:
	//
	// settle-job
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The coordinate information of the node.
	Coordinate *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate `json:"Coordinate,omitempty" xml:"Coordinate,omitempty" type:"Struct"`
	// The dependency policy.
	//
	// example:
	//
	// 1
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node type.
	//
	// example:
	//
	// script_shell
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The node name.
	//
	// example:
	//
	// job1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node status. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkflowDAGPreviewResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetCoordinate() *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate {
	return s.Coordinate
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetDependentStrategy() *int32 {
	return s.DependentStrategy
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetJobType() *string {
	return s.JobType
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetAppName(v string) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.AppName = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetCoordinate(v *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.Coordinate = v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetDependentStrategy(v int32) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.DependentStrategy = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetId(v int64) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.Id = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetJobType(v string) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.JobType = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetName(v string) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) SetStatus(v int32) *GetWorkflowDAGPreviewResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodes) Validate() error {
	if s.Coordinate != nil {
		if err := s.Coordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate struct {
	// The height of the node. This field is optional.
	//
	// example:
	//
	// 20
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The width of the node. This field is optional.
	//
	// example:
	//
	// 100
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-axis coordinate of the node.
	//
	// example:
	//
	// 50
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-axis coordinate of the node.
	//
	// example:
	//
	// 50
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) GetHeight() *float32 {
	return s.Height
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) GetWidth() *float32 {
	return s.Width
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) GetX() *float32 {
	return s.X
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) GetY() *float32 {
	return s.Y
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) SetHeight(v float32) *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate {
	s.Height = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) SetWidth(v float32) *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate {
	s.Width = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) SetX(v float32) *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate {
	s.X = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) SetY(v float32) *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate {
	s.Y = &v
	return s
}

func (s *GetWorkflowDAGPreviewResponseBodyDataNodesCoordinate) Validate() error {
	return dara.Validate(s)
}
