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
	// Leave this parameter empty if not specified. The filter condition: within the specified container. Specify the container ID. This parameter is not related to the resource group (ResourceGroupId).
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The node name. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Filter condition: scheduling type. Valid values:
	//
	// - Normal: The node is executed normally.
	//
	// - Pause: The node status is set to paused, and downstream nodes that depend on the current node are blocked from execution.
	//
	// - Skip: The node status is set to dry run. The system directly returns a success result (with an execution duration of 0 seconds), does not block downstream node execution, and does not consume resources.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The rerun property. If not specified, this parameter is left empty. Valid values:
	//
	// - Allowed: The node can be rerun regardless of whether it runs successfully or fails.
	//
	// - FailureAllowed: The node can be rerun only after a failed run, not after a successful run.
	//
	// - Denied: The node cannot be rerun regardless of whether it runs successfully or fails.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The scene in which the node resides. Leave this parameter empty if not specified. This parameter corresponds to the partition of the left-side navigation pane in DataStudio. Valid values:
	//
	// - DataworksProject: project folder.
	//
	// - DataworksManualWorkflow: manual workflow.
	//
	// - DataworksManualTask: manual node.
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
