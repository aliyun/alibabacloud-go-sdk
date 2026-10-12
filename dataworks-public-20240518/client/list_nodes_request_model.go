// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v string) *ListNodesRequest
	GetContainerId() *string
	SetName(v string) *ListNodesRequest
	GetName() *string
	SetPageNumber(v int32) *ListNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListNodesRequest
	GetProjectId() *int64
	SetRecurrence(v string) *ListNodesRequest
	GetRecurrence() *string
	SetRerunMode(v string) *ListNodesRequest
	GetRerunMode() *string
	SetScene(v string) *ListNodesRequest
	GetScene() *string
}

type ListNodesRequest struct {
	// The ID of the container. This parameter is used as a filter condition to query nodes in a specified container. If you do not specify this parameter, leave it empty. This parameter is independent of the ResourceGroupId parameter.
	//
	// 	Notice: Before SDK version 8.0.0, this field is of the Long type. In SDK version 8.0.0 and later, this field is of the String type. **This change does not affect the normal use of the SDK, and the parameter is still returned based on the type defined in the SDK**. However, if you upgrade the SDK across version 8.0.0, the type change may cause project compilation failures. In this case, you must manually correct the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the node. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the requested data.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type. This parameter is used as a filter condition. Valid values:
	//
	// - Normal: The node runs as expected.
	//
	// - Pause: The node is paused, and the execution of its downstream nodes is blocked.
	//
	//
	//
	// - Skip: The node is set to dry run. The system directly returns a success response with an execution duration of 0 seconds. This does not block the execution of downstream nodes or consume resources.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The rerun mode. If you do not specify this parameter, leave it empty. Valid values:
	//
	// - Allowed: The node can be rerun regardless of whether it is successfully run or fails to run.
	//
	//
	//
	// - FailureAllowed: The node can be rerun only after it fails to run.
	//
	//
	//
	// - Denied: The node cannot be rerun regardless of whether it is successfully run or fails to run.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The scenario to which the node belongs. This parameter is used as a filter condition. If you do not specify this parameter, leave it empty. In Data Studio, this parameter corresponds to the partitions in the left-side directory tree. Valid values:
	//
	// - DataworksProject: project directory.
	//
	// - DataworksManualWorkflow: manual workflow.
	//
	// - DataworksManualTask: manual task.
	//
	// example:
	//
	// DataworksProject
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *ListNodesRequest) GetName() *string {
	return s.Name
}

func (s *ListNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesRequest) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListNodesRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListNodesRequest) GetScene() *string {
	return s.Scene
}

func (s *ListNodesRequest) SetContainerId(v string) *ListNodesRequest {
	s.ContainerId = &v
	return s
}

func (s *ListNodesRequest) SetName(v string) *ListNodesRequest {
	s.Name = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProjectId(v int64) *ListNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequest) SetRecurrence(v string) *ListNodesRequest {
	s.Recurrence = &v
	return s
}

func (s *ListNodesRequest) SetRerunMode(v string) *ListNodesRequest {
	s.RerunMode = &v
	return s
}

func (s *ListNodesRequest) SetScene(v string) *ListNodesRequest {
	s.Scene = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
