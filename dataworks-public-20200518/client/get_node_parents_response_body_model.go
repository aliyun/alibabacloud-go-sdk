// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeParentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNodeParentsResponseBodyData) *GetNodeParentsResponseBody
	GetData() *GetNodeParentsResponseBodyData
	SetErrorCode(v string) *GetNodeParentsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetNodeParentsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetNodeParentsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeParentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeParentsResponseBody
	GetSuccess() *bool
}

type GetNodeParentsResponseBody struct {
	// The list of node information returned.
	Data *GetNodeParentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The unique ID of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// asfsdfas-adfasdf-asfas-dfasdf-asdf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeParentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeParentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeParentsResponseBody) GetData() *GetNodeParentsResponseBodyData {
	return s.Data
}

func (s *GetNodeParentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeParentsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNodeParentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeParentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeParentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeParentsResponseBody) SetData(v *GetNodeParentsResponseBodyData) *GetNodeParentsResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeParentsResponseBody) SetErrorCode(v string) *GetNodeParentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeParentsResponseBody) SetErrorMessage(v string) *GetNodeParentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNodeParentsResponseBody) SetHttpStatusCode(v int32) *GetNodeParentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeParentsResponseBody) SetRequestId(v string) *GetNodeParentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeParentsResponseBody) SetSuccess(v bool) *GetNodeParentsResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeParentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeParentsResponseBodyData struct {
	// The list of nodes.
	Nodes []*GetNodeParentsResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetNodeParentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNodeParentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeParentsResponseBodyData) GetNodes() []*GetNodeParentsResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetNodeParentsResponseBodyData) SetNodes(v []*GetNodeParentsResponseBodyDataNodes) *GetNodeParentsResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetNodeParentsResponseBodyData) Validate() error {
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

type GetNodeParentsResponseBodyDataNodes struct {
	// The baseline ID.
	//
	// example:
	//
	// 1244564565
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The CRON expression. This expression is used for timed scheduling to execute the node task.
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
	// The node name.
	//
	// example:
	//
	// test_Node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The DataWorks UID of the node owner.
	//
	// example:
	//
	// 123124561341251321
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority. Valid values: 1 to 8. A larger value indicates a higher priority.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The node type.
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
	// Indicates whether the node can be rerun upon failure. Valid values:
	//
	// - true: The node can be rerun.
	//
	// - false: The node cannot be rerun.
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The scheduling type. Valid values:
	//
	// - NORMAL: normal scheduling node.
	//
	// - MANUAL: manual node that is not triggered by daily scheduling.
	//
	// - PAUSE: paused node that is triggered by daily scheduling but is set to failed when scheduling starts.
	//
	// - SKIP: dry-run node that is triggered by daily scheduling but is set to successful when scheduling starts.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// The scheduling dependency type. Valid values:
	//
	// - **0**: same-cycle dependency.
	//
	// - **3**: cross-cycle dependency.
	//
	// example:
	//
	// 0
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
}

func (s GetNodeParentsResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetNodeParentsResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetNodeParentsResponseBodyDataNodes) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetNodeParentsResponseBodyDataNodes) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetNodeParentsResponseBodyDataNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeParentsResponseBodyDataNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeParentsResponseBodyDataNodes) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetNodeParentsResponseBodyDataNodes) GetPriority() *int32 {
	return s.Priority
}

func (s *GetNodeParentsResponseBodyDataNodes) GetProgramType() *string {
	return s.ProgramType
}

func (s *GetNodeParentsResponseBodyDataNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeParentsResponseBodyDataNodes) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *GetNodeParentsResponseBodyDataNodes) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetNodeParentsResponseBodyDataNodes) GetStepType() *string {
	return s.StepType
}

func (s *GetNodeParentsResponseBodyDataNodes) SetBaselineId(v int64) *GetNodeParentsResponseBodyDataNodes {
	s.BaselineId = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetCronExpress(v string) *GetNodeParentsResponseBodyDataNodes {
	s.CronExpress = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetNodeId(v int64) *GetNodeParentsResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetNodeName(v string) *GetNodeParentsResponseBodyDataNodes {
	s.NodeName = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetOwnerId(v string) *GetNodeParentsResponseBodyDataNodes {
	s.OwnerId = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetPriority(v int32) *GetNodeParentsResponseBodyDataNodes {
	s.Priority = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetProgramType(v string) *GetNodeParentsResponseBodyDataNodes {
	s.ProgramType = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetProjectId(v int64) *GetNodeParentsResponseBodyDataNodes {
	s.ProjectId = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetRepeatability(v bool) *GetNodeParentsResponseBodyDataNodes {
	s.Repeatability = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetSchedulerType(v string) *GetNodeParentsResponseBodyDataNodes {
	s.SchedulerType = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) SetStepType(v string) *GetNodeParentsResponseBodyDataNodes {
	s.StepType = &v
	return s
}

func (s *GetNodeParentsResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
