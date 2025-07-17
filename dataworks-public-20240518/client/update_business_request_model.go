// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v int64) *UpdateBusinessRequest
	GetBusinessId() *int64
	SetBusinessName(v string) *UpdateBusinessRequest
	GetBusinessName() *string
	SetDescription(v string) *UpdateBusinessRequest
	GetDescription() *string
	SetOwner(v string) *UpdateBusinessRequest
	GetOwner() *string
	SetProjectId(v int64) *UpdateBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *UpdateBusinessRequest
	GetProjectIdentifier() *string
}

type UpdateBusinessRequest struct {
	// The workflow ID.
	//
	// You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300000
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The name of the workflow.
	//
	// You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the name.
	//
	// example:
	//
	// MyBusiness
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// modified from my first business
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The owner of the workflow.
	//
	// You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the owner.
	//
	// example:
	//
	// 348428****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Workspace page to obtain the workspace ID. You must configure either this parameter or the `ProjectIdentifier` parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Workspace page to obtain the name. You must configure either this parameter or the `ProjectId` parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s UpdateBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBusinessRequest) GoString() string {
	return s.String()
}

func (s *UpdateBusinessRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *UpdateBusinessRequest) GetBusinessName() *string {
	return s.BusinessName
}

func (s *UpdateBusinessRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBusinessRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *UpdateBusinessRequest) SetBusinessId(v int64) *UpdateBusinessRequest {
	s.BusinessId = &v
	return s
}

func (s *UpdateBusinessRequest) SetBusinessName(v string) *UpdateBusinessRequest {
	s.BusinessName = &v
	return s
}

func (s *UpdateBusinessRequest) SetDescription(v string) *UpdateBusinessRequest {
	s.Description = &v
	return s
}

func (s *UpdateBusinessRequest) SetOwner(v string) *UpdateBusinessRequest {
	s.Owner = &v
	return s
}

func (s *UpdateBusinessRequest) SetProjectId(v int64) *UpdateBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBusinessRequest) SetProjectIdentifier(v string) *UpdateBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *UpdateBusinessRequest) Validate() error {
	return dara.Validate(s)
}
