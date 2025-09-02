// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetProjectRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetProjectRequest
	GetProjectIdentifier() *string
}

type GetProjectRequest struct {
	// The ID of the DataWorks workspace. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to query the ID.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to query the name.
	//
	// example:
	//
	// test_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetProjectRequest) SetProjectId(v int64) *GetProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRequest) SetProjectIdentifier(v string) *GetProjectRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
