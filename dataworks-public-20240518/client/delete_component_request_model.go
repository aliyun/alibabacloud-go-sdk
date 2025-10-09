// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *DeleteComponentRequest
	GetComponentId() *string
	SetProjectId(v int64) *DeleteComponentRequest
	GetProjectId() *int64
}

type DeleteComponentRequest struct {
	// The component ID. It can be used as a request parameter for querying the list of production studio components and modifying production studio components.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123123123123
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter specifies the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentRequest) GoString() string {
	return s.String()
}

func (s *DeleteComponentRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *DeleteComponentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteComponentRequest) SetComponentId(v string) *DeleteComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *DeleteComponentRequest) SetProjectId(v int64) *DeleteComponentRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteComponentRequest) Validate() error {
	return dara.Validate(s)
}
