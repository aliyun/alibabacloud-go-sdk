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
	// The ID of the container. If you specify this parameter, only nodes in the specified container are returned. This parameter is independent of the resource group (ResourceGroupId).
	//
	// 	Notice:
	//
	// This parameter is of the Long type in SDK versions earlier than 8.0.0 and of the String type in SDK 8.0.0 and later. **This change does not affect SDK usage. The parameter is returned in the type defined for your SDK version.*	- The type change may cause compilation errors only when you upgrade the SDK across version 8.0.0. In this case, you must manually correct the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The node name. Fuzzy search is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the results to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default: 10. Maximum: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. To find this ID, log in to the [DataWorks console](https://workbench.data.aliyun.com/console) and navigate to the workspace configuration page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Filters nodes by their scheduling type. Valid values:
	//
	// - Normal: The node runs as scheduled.
	//
	// - Pause: The node is paused and blocks its dependent downstream nodes.
	//
	// - Skip: The node is skipped, and the system immediately returns a success status with a 0-second execution time. This action does not block downstream nodes or consume resources.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The rerun mode. Valid values:
	//
	// - Allowed: The node can be rerun regardless of whether it succeeded or failed.
	//
	// - FailureAllowed: The node can be rerun only if its previous run failed.
	//
	// - Denied: The node cannot be rerun regardless of whether it succeeded or failed.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The context for filtering nodes. In data development, this corresponds to the sections in the directory tree on the left. If you omit this parameter, no filtering is applied. Valid values:
	//
	// - DataworksProject: Nodes in the project directory.
	//
	// - DataworksManualWorkflow: manual workflow
	//
	// - DataworksManualTask: manual task
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
