// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *DeleteBusinessRequest
	GetBusinessId() *int64
	SetProjectId(v int64) *DeleteBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *DeleteBusinessRequest
	GetProjectIdentifier() *string
}

type DeleteBusinessRequest struct {
	// The workflow ID. You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the ID of a workflow by name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must specify either this parameter or the ProjectIdentifier parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace page to query the workspace name. You must specify either this parameter or the ProjectId parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeleteBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessRequest) GoString() string {
	return s.String()
}

func (s *DeleteBusinessRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *DeleteBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *DeleteBusinessRequest) SetBusinessId(v int64) *DeleteBusinessRequest {
	s.BusinessId = &v
	return s
}

func (s *DeleteBusinessRequest) SetProjectId(v int64) *DeleteBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteBusinessRequest) SetProjectIdentifier(v string) *DeleteBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *DeleteBusinessRequest) Validate() error {
	return dara.Validate(s)
}
