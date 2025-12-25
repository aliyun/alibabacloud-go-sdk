// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateWorkflowDAGRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateWorkflowDAGRequest
	GetClusterId() *string
	SetDag(v *UpdateWorkflowDAGRequestDag) *UpdateWorkflowDAGRequest
	GetDag() *UpdateWorkflowDAGRequestDag
	SetDagVersion(v string) *UpdateWorkflowDAGRequest
	GetDagVersion() *string
	SetWorkflowId(v int64) *UpdateWorkflowDAGRequest
	GetWorkflowId() *int64
}

type UpdateWorkflowDAGRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Dag *UpdateWorkflowDAGRequestDag `json:"Dag,omitempty" xml:"Dag,omitempty" type:"Struct"`
	// example:
	//
	// 1137005
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowDAGRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWorkflowDAGRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateWorkflowDAGRequest) GetDag() *UpdateWorkflowDAGRequestDag {
	return s.Dag
}

func (s *UpdateWorkflowDAGRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *UpdateWorkflowDAGRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *UpdateWorkflowDAGRequest) SetAppName(v string) *UpdateWorkflowDAGRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWorkflowDAGRequest) SetClusterId(v string) *UpdateWorkflowDAGRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWorkflowDAGRequest) SetDag(v *UpdateWorkflowDAGRequestDag) *UpdateWorkflowDAGRequest {
	s.Dag = v
	return s
}

func (s *UpdateWorkflowDAGRequest) SetDagVersion(v string) *UpdateWorkflowDAGRequest {
	s.DagVersion = &v
	return s
}

func (s *UpdateWorkflowDAGRequest) SetWorkflowId(v int64) *UpdateWorkflowDAGRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowDAGRequest) Validate() error {
	if s.Dag != nil {
		if err := s.Dag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkflowDAGRequestDag struct {
	Edges []*UpdateWorkflowDAGRequestDagEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	Nodes []*UpdateWorkflowDAGRequestDagNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s UpdateWorkflowDAGRequestDag) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGRequestDag) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGRequestDag) GetEdges() []*UpdateWorkflowDAGRequestDagEdges {
	return s.Edges
}

func (s *UpdateWorkflowDAGRequestDag) GetNodes() []*UpdateWorkflowDAGRequestDagNodes {
	return s.Nodes
}

func (s *UpdateWorkflowDAGRequestDag) SetEdges(v []*UpdateWorkflowDAGRequestDagEdges) *UpdateWorkflowDAGRequestDag {
	s.Edges = v
	return s
}

func (s *UpdateWorkflowDAGRequestDag) SetNodes(v []*UpdateWorkflowDAGRequestDagNodes) *UpdateWorkflowDAGRequestDag {
	s.Nodes = v
	return s
}

func (s *UpdateWorkflowDAGRequestDag) Validate() error {
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

type UpdateWorkflowDAGRequestDagEdges struct {
	// example:
	//
	// 3
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 4
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s UpdateWorkflowDAGRequestDagEdges) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGRequestDagEdges) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGRequestDagEdges) GetSource() *int64 {
	return s.Source
}

func (s *UpdateWorkflowDAGRequestDagEdges) GetTarget() *int64 {
	return s.Target
}

func (s *UpdateWorkflowDAGRequestDagEdges) SetSource(v int64) *UpdateWorkflowDAGRequestDagEdges {
	s.Source = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagEdges) SetTarget(v int64) *UpdateWorkflowDAGRequestDagEdges {
	s.Target = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagEdges) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowDAGRequestDagNodes struct {
	// example:
	//
	// {
	//
	//     "logicType": "status_branch",
	//
	//     "conditionResult": {
	//
	//         "successList": [
	//
	//             {
	//
	//                 "jobName": "status-job4",
	//
	//                 "jobId": "269"
	//
	//             }
	//
	//         ],
	//
	//         "failedList": [
	//
	//             {
	//
	//                 "jobName": "status-job5",
	//
	//                 "jobId": "270"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dependence": {
	//
	//         "relation": "AND",
	//
	//         "dependList": [
	//
	//             {
	//
	//                 "relation": "AND",
	//
	//                 "dependItemList": [
	//
	//                     {
	//
	//                         "jobName": "status-job1",
	//
	//                         "jobId": 265,
	//
	//                         "status": [
	//
	//                             4
	//
	//                         ]
	//
	//                     },
	//
	//                     {
	//
	//                         "jobName": "status-job2",
	//
	//                         "jobId": 266,
	//
	//                         "status": [
	//
	//                             5
	//
	//                         ]
	//
	//                     }
	//
	//                 ]
	//
	//             },
	//
	//             {
	//
	//                 "relation": "AND",
	//
	//                 "dependItemList": [
	//
	//                     {
	//
	//                         "jobName": "status-job3",
	//
	//                         "jobId": 267,
	//
	//                         "status": [
	//
	//                             4
	//
	//                         ]
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Content    *string                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	Coordinate *UpdateWorkflowDAGRequestDagNodesCoordinate `json:"Coordinate,omitempty" xml:"Coordinate,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateWorkflowDAGRequestDagNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGRequestDagNodes) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGRequestDagNodes) GetContent() *string {
	return s.Content
}

func (s *UpdateWorkflowDAGRequestDagNodes) GetCoordinate() *UpdateWorkflowDAGRequestDagNodesCoordinate {
	return s.Coordinate
}

func (s *UpdateWorkflowDAGRequestDagNodes) GetId() *int64 {
	return s.Id
}

func (s *UpdateWorkflowDAGRequestDagNodes) SetContent(v string) *UpdateWorkflowDAGRequestDagNodes {
	s.Content = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodes) SetCoordinate(v *UpdateWorkflowDAGRequestDagNodesCoordinate) *UpdateWorkflowDAGRequestDagNodes {
	s.Coordinate = v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodes) SetId(v int64) *UpdateWorkflowDAGRequestDagNodes {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodes) Validate() error {
	if s.Coordinate != nil {
		if err := s.Coordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkflowDAGRequestDagNodesCoordinate struct {
	// example:
	//
	// 20
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 50
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 50
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s UpdateWorkflowDAGRequestDagNodesCoordinate) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGRequestDagNodesCoordinate) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) GetX() *float32 {
	return s.X
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) GetY() *float32 {
	return s.Y
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) SetHeight(v float32) *UpdateWorkflowDAGRequestDagNodesCoordinate {
	s.Height = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) SetWidth(v float32) *UpdateWorkflowDAGRequestDagNodesCoordinate {
	s.Width = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) SetX(v float32) *UpdateWorkflowDAGRequestDagNodesCoordinate {
	s.X = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) SetY(v float32) *UpdateWorkflowDAGRequestDagNodesCoordinate {
	s.Y = &v
	return s
}

func (s *UpdateWorkflowDAGRequestDagNodesCoordinate) Validate() error {
	return dara.Validate(s)
}
