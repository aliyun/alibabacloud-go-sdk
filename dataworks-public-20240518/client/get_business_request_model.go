// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *GetBusinessRequest
	GetBusinessId() *int64
	SetProjectId(v int64) *GetBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetBusinessRequest
	GetProjectIdentifier() *string
}

type GetBusinessRequest struct {
	// The workflow ID. You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to obtain the folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000111
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

func (s GetBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetBusinessRequest) SetBusinessId(v int64) *GetBusinessRequest {
	s.BusinessId = &v
	return s
}

func (s *GetBusinessRequest) SetProjectId(v int64) *GetBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBusinessRequest) SetProjectIdentifier(v string) *GetBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetBusinessRequest) Validate() error {
	return dara.Validate(s)
}
