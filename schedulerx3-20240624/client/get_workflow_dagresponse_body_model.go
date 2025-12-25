// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWorkflowDAGResponseBody
	GetCode() *int32
	SetData(v *GetWorkflowDAGResponseBodyData) *GetWorkflowDAGResponseBody
	GetData() *GetWorkflowDAGResponseBodyData
	SetMessage(v string) *GetWorkflowDAGResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkflowDAGResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkflowDAGResponseBody
	GetSuccess() *bool
}

type GetWorkflowDAGResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetWorkflowDAGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowDAGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWorkflowDAGResponseBody) GetData() *GetWorkflowDAGResponseBodyData {
	return s.Data
}

func (s *GetWorkflowDAGResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowDAGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowDAGResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkflowDAGResponseBody) SetCode(v int32) *GetWorkflowDAGResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowDAGResponseBody) SetData(v *GetWorkflowDAGResponseBodyData) *GetWorkflowDAGResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowDAGResponseBody) SetMessage(v string) *GetWorkflowDAGResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowDAGResponseBody) SetRequestId(v string) *GetWorkflowDAGResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDAGResponseBody) SetSuccess(v bool) *GetWorkflowDAGResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkflowDAGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowDAGResponseBodyData struct {
	Edges []*GetWorkflowDAGResponseBodyDataEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	Nodes []*GetWorkflowDAGResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkflowDAGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponseBodyData) GetEdges() []*GetWorkflowDAGResponseBodyDataEdges {
	return s.Edges
}

func (s *GetWorkflowDAGResponseBodyData) GetNodes() []*GetWorkflowDAGResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetWorkflowDAGResponseBodyData) SetEdges(v []*GetWorkflowDAGResponseBodyDataEdges) *GetWorkflowDAGResponseBodyData {
	s.Edges = v
	return s
}

func (s *GetWorkflowDAGResponseBodyData) SetNodes(v []*GetWorkflowDAGResponseBodyDataNodes) *GetWorkflowDAGResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetWorkflowDAGResponseBodyData) Validate() error {
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

type GetWorkflowDAGResponseBodyDataEdges struct {
	// example:
	//
	// 3
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 4
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkflowDAGResponseBodyDataEdges) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponseBodyDataEdges) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponseBodyDataEdges) GetSource() *int64 {
	return s.Source
}

func (s *GetWorkflowDAGResponseBodyDataEdges) GetTarget() *int64 {
	return s.Target
}

func (s *GetWorkflowDAGResponseBodyDataEdges) SetSource(v int64) *GetWorkflowDAGResponseBodyDataEdges {
	s.Source = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataEdges) SetTarget(v int64) *GetWorkflowDAGResponseBodyDataEdges {
	s.Target = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataEdges) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowDAGResponseBodyDataNodes struct {
	// example:
	//
	// oak-payment-async-job
	AppName    *string                                        `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Coordinate *GetWorkflowDAGResponseBodyDataNodesCoordinate `json:"Coordinate,omitempty" xml:"Coordinate,omitempty" type:"Struct"`
	// example:
	//
	// all_success
	DependentStrategy *int32 `json:"DependentStrategy,omitempty" xml:"DependentStrategy,omitempty"`
	// example:
	//
	// 5
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// script_shell
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// job1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkflowDAGResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetCoordinate() *GetWorkflowDAGResponseBodyDataNodesCoordinate {
	return s.Coordinate
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetDependentStrategy() *int32 {
	return s.DependentStrategy
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetJobType() *string {
	return s.JobType
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *GetWorkflowDAGResponseBodyDataNodes) GetStatus() *int32 {
	return s.Status
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetAppName(v string) *GetWorkflowDAGResponseBodyDataNodes {
	s.AppName = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetCoordinate(v *GetWorkflowDAGResponseBodyDataNodesCoordinate) *GetWorkflowDAGResponseBodyDataNodes {
	s.Coordinate = v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetDependentStrategy(v int32) *GetWorkflowDAGResponseBodyDataNodes {
	s.DependentStrategy = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetId(v int64) *GetWorkflowDAGResponseBodyDataNodes {
	s.Id = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetJobType(v string) *GetWorkflowDAGResponseBodyDataNodes {
	s.JobType = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetName(v string) *GetWorkflowDAGResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) SetStatus(v int32) *GetWorkflowDAGResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodes) Validate() error {
	if s.Coordinate != nil {
		if err := s.Coordinate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowDAGResponseBodyDataNodesCoordinate struct {
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

func (s GetWorkflowDAGResponseBodyDataNodesCoordinate) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGResponseBodyDataNodesCoordinate) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) GetHeight() *float32 {
	return s.Height
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) GetWidth() *float32 {
	return s.Width
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) GetX() *float32 {
	return s.X
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) GetY() *float32 {
	return s.Y
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) SetHeight(v float32) *GetWorkflowDAGResponseBodyDataNodesCoordinate {
	s.Height = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) SetWidth(v float32) *GetWorkflowDAGResponseBodyDataNodesCoordinate {
	s.Width = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) SetX(v float32) *GetWorkflowDAGResponseBodyDataNodesCoordinate {
	s.X = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) SetY(v float32) *GetWorkflowDAGResponseBodyDataNodesCoordinate {
	s.Y = &v
	return s
}

func (s *GetWorkflowDAGResponseBodyDataNodesCoordinate) Validate() error {
	return dara.Validate(s)
}
