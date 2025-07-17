// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RenameResourceRequest
	GetId() *int64
	SetName(v string) *RenameResourceRequest
	GetName() *string
	SetProjectId(v int64) *RenameResourceRequest
	GetProjectId() *int64
}

type RenameResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
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
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameResourceRequest) GoString() string {
	return s.String()
}

func (s *RenameResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *RenameResourceRequest) GetName() *string {
	return s.Name
}

func (s *RenameResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameResourceRequest) SetId(v int64) *RenameResourceRequest {
	s.Id = &v
	return s
}

func (s *RenameResourceRequest) SetName(v string) *RenameResourceRequest {
	s.Name = &v
	return s
}

func (s *RenameResourceRequest) SetProjectId(v int64) *RenameResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameResourceRequest) Validate() error {
	return dara.Validate(s)
}
