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
	// Leave this parameter empty if not specified. Filter condition: within a specified container. Specify the container ID. This parameter is independent of the resource group ID (ResourceGroupId).
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the node. Fuzzy search is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the data to retrieve, used for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the data to retrieve, used for pagination.
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
	// Leave this parameter empty if not specified. Filter condition: within a specified container. Specify the container ID. This parameter is independent of the resource group ID (ResourceGroupId).
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
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
