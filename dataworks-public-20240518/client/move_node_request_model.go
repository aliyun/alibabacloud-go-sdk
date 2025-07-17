// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MoveNodeRequest
	GetId() *int64
	SetPath(v string) *MoveNodeRequest
	GetPath() *string
	SetProjectId(v int64) *MoveNodeRequest
	GetProjectId() *int64
}

type MoveNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the node. You do not need to specify a node name in the path.
	//
	// For example, if you want to move the test node to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You can use this parameter to specify the DataWorks workspace on which you want to perform the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveNodeRequest) GoString() string {
	return s.String()
}

func (s *MoveNodeRequest) GetId() *int64 {
	return s.Id
}

func (s *MoveNodeRequest) GetPath() *string {
	return s.Path
}

func (s *MoveNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveNodeRequest) SetId(v int64) *MoveNodeRequest {
	s.Id = &v
	return s
}

func (s *MoveNodeRequest) SetPath(v string) *MoveNodeRequest {
	s.Path = &v
	return s
}

func (s *MoveNodeRequest) SetProjectId(v int64) *MoveNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveNodeRequest) Validate() error {
	return dara.Validate(s)
}
