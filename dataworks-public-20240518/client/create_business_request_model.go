// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessName(v string) *CreateBusinessRequest
	GetBusinessName() *string
	SetDescription(v string) *CreateBusinessRequest
	GetDescription() *string
	SetOwner(v string) *CreateBusinessRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *CreateBusinessRequest
	GetProjectIdentifier() *string
	SetUseType(v string) *CreateBusinessRequest
	GetUseType() *string
}

type CreateBusinessRequest struct {
	// The name of the workflow. Workflow names must be unique within the same workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// My business process
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud account ID of the person responsible for the workflow. You can view the account ID by hovering over the user avatar in the top-right corner of the [DataWorks console](https://workbench.data.aliyun.com/console). If this parameter is not specified, the account ID of the API caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace. You can view the workspace ID on the workspace management page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, This is the identifier shown in the workspace switch at the top of the Data Studio page. Either this parameter or ProjectId must be specified to determine which DataWorks project this API call operates on.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The module to which the workflow belongs. Valid values:
	//
	// 	- NORMAL: Data Studio
	//
	// 	- MANUAL_BIZ: Manually triggered workflow
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s CreateBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessRequest) GoString() string {
	return s.String()
}

func (s *CreateBusinessRequest) GetBusinessName() *string {
	return s.BusinessName
}

func (s *CreateBusinessRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBusinessRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateBusinessRequest) GetUseType() *string {
	return s.UseType
}

func (s *CreateBusinessRequest) SetBusinessName(v string) *CreateBusinessRequest {
	s.BusinessName = &v
	return s
}

func (s *CreateBusinessRequest) SetDescription(v string) *CreateBusinessRequest {
	s.Description = &v
	return s
}

func (s *CreateBusinessRequest) SetOwner(v string) *CreateBusinessRequest {
	s.Owner = &v
	return s
}

func (s *CreateBusinessRequest) SetProjectId(v int64) *CreateBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateBusinessRequest) SetProjectIdentifier(v string) *CreateBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateBusinessRequest) SetUseType(v string) *CreateBusinessRequest {
	s.UseType = &v
	return s
}

func (s *CreateBusinessRequest) Validate() error {
	return dara.Validate(s)
}
