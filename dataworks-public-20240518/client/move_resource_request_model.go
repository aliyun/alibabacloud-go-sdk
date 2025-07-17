// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MoveResourceRequest
	GetId() *int64
	SetPath(v string) *MoveResourceRequest
	GetPath() *string
	SetProjectId(v int64) *MoveResourceRequest
	GetProjectId() *int64
}

type MoveResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the file resource. You do not need to specify a file resource name in the path.
	//
	// For example, if you want to move the test file resource to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
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

func (s MoveResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *MoveResourceRequest) GetPath() *string {
	return s.Path
}

func (s *MoveResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveResourceRequest) SetId(v int64) *MoveResourceRequest {
	s.Id = &v
	return s
}

func (s *MoveResourceRequest) SetPath(v string) *MoveResourceRequest {
	s.Path = &v
	return s
}

func (s *MoveResourceRequest) SetProjectId(v int64) *MoveResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveResourceRequest) Validate() error {
	return dara.Validate(s)
}
