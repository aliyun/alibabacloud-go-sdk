// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RenameNodeRequest
	GetId() *int64
	SetName(v string) *RenameNodeRequest
	GetName() *string
	SetProjectId(v int64) *RenameNodeRequest
	GetProjectId() *int64
}

type RenameNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameNodeRequest) GoString() string {
	return s.String()
}

func (s *RenameNodeRequest) GetId() *int64 {
	return s.Id
}

func (s *RenameNodeRequest) GetName() *string {
	return s.Name
}

func (s *RenameNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameNodeRequest) SetId(v int64) *RenameNodeRequest {
	s.Id = &v
	return s
}

func (s *RenameNodeRequest) SetName(v string) *RenameNodeRequest {
	s.Name = &v
	return s
}

func (s *RenameNodeRequest) SetProjectId(v int64) *RenameNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameNodeRequest) Validate() error {
	return dara.Validate(s)
}
