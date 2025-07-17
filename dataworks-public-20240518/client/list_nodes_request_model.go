// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v int64) *ListNodesRequest
	GetContainerId() *int64
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
	// The container ID, which is a filter condition. If you do not want to use this condition for filtering, you do not need to configure this parameter. The container ID that you specify is unrelated to the resource group ID indicated by the ResourceGroupId parameter.
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *int64  `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
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
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type, which is a filter condition. Valid values:
	//
	// 	- Normal: The nodes are scheduled as expected.
	//
	// 	- Pause: The nodes are paused, and the running of their descendant nodes is blocked.
	//
	// 	- Skip: The nodes are dry run. The system does not actually run the nodes, but directly returns a success response. The running duration of the nodes is 0 seconds. In addition, the nodes do not occupy resources or block the running of their descendant nodes.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The rerun property, which is a filter condition. If you do not want to use this condition for filtering, you do not need to configure this parameter. Valid values:
	//
	// 	- Allowed: The nodes can be rerun regardless of whether they are successfully run or fail to run.
	//
	// 	- FailureAllowed: The nodes can be rerun only after they fail to run.
	//
	// 	- Denied: The nodes cannot be rerun regardless of whether they are successfully run or fail to run.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The location of the nodes in the left-side navigation pane of the Data Studio page, which is a filter condition. If you do not want to use this condition for filtering, you do not need to configure this parameter. Valid values:
	//
	// 	- DataworksProject
	//
	// 	- DataworksManualWorkflow
	//
	// 	- DataworksManualTask
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetContainerId() *int64 {
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

func (s *ListNodesRequest) SetContainerId(v int64) *ListNodesRequest {
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
