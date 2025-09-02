// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeChildrenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNodeChildrenResponseBodyData) *GetNodeChildrenResponseBody
	GetData() *GetNodeChildrenResponseBodyData
	SetErrorCode(v string) *GetNodeChildrenResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetNodeChildrenResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetNodeChildrenResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeChildrenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeChildrenResponseBody
	GetSuccess() *bool
}

type GetNodeChildrenResponseBody struct {
	// The information about the descendant nodes.
	Data *GetNodeChildrenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1060010000000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// err
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// adsfasdf-adf-asdf-asdf-asdfadfasdd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeChildrenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeChildrenResponseBody) GetData() *GetNodeChildrenResponseBodyData {
	return s.Data
}

func (s *GetNodeChildrenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeChildrenResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNodeChildrenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeChildrenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeChildrenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeChildrenResponseBody) SetData(v *GetNodeChildrenResponseBodyData) *GetNodeChildrenResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeChildrenResponseBody) SetErrorCode(v string) *GetNodeChildrenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeChildrenResponseBody) SetErrorMessage(v string) *GetNodeChildrenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNodeChildrenResponseBody) SetHttpStatusCode(v int32) *GetNodeChildrenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeChildrenResponseBody) SetRequestId(v string) *GetNodeChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeChildrenResponseBody) SetSuccess(v bool) *GetNodeChildrenResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeChildrenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeChildrenResponseBodyData struct {
	// The descendant nodes.
	Nodes []*GetNodeChildrenResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetNodeChildrenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNodeChildrenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeChildrenResponseBodyData) GetNodes() []*GetNodeChildrenResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetNodeChildrenResponseBodyData) SetNodes(v []*GetNodeChildrenResponseBodyDataNodes) *GetNodeChildrenResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetNodeChildrenResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetNodeChildrenResponseBodyDataNodes struct {
	// The baseline ID.
	//
	// example:
	//
	// 12345656
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The CRON expression. CRON expressions are used to run auto triggered nodes.
	//
	// example:
	//
	// 00 00 	- 	- 	- *
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 1244564565
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test_Node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 123124561341251321
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority. Valid values: 1 to 8. A large value indicates a high priority.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 12315412412
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Indicates whether the node can be rerun if the node fails to run. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node.
	//
	// 	- MANUAL: The node is a manually triggered node. Manually triggered nodes cannot be automatically triggered.
	//
	// 	- PAUSE: The node is a paused node. Paused nodes are started as scheduled but the system sets the status of the nodes to failed when it starts to run them.
	//
	// 	- SKIP: The node is a dry-run node. Dry-run nodes are started as scheduled but the system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The scheduling dependency type. Valid values:
	//
	// 	- **0**: same-cycle scheduling dependency
	//
	// 	- **3**: cross-cycle scheduling dependency
	//
	// example:
	//
	// 0
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s GetNodeChildrenResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetNodeChildrenResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetPriority() *int32 {
	return s.Priority
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetProgramType() *string {
	return s.ProgramType
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetNodeChildrenResponseBodyDataNodes) GetStepType() *string {
	return s.StepType
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetBaselineId(v int64) *GetNodeChildrenResponseBodyDataNodes {
	s.BaselineId = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetCronExpress(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.CronExpress = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetNodeId(v int64) *GetNodeChildrenResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetNodeName(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.NodeName = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetOwnerId(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.OwnerId = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetPriority(v int32) *GetNodeChildrenResponseBodyDataNodes {
	s.Priority = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetProgramType(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.ProgramType = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetProjectId(v int64) *GetNodeChildrenResponseBodyDataNodes {
	s.ProjectId = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetRepeatability(v bool) *GetNodeChildrenResponseBodyDataNodes {
	s.Repeatability = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetSchedulerType(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.SchedulerType = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) SetStepType(v string) *GetNodeChildrenResponseBodyDataNodes {
	s.StepType = &v
	return s
}

func (s *GetNodeChildrenResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
