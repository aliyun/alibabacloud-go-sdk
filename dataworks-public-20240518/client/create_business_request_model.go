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
	// The name of the business process. The name must be unique within the same workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// My business process
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The description of the business process.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud account ID of the owner of the business process. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and hover over the profile picture in the upper-right corner of the top navigation bar to view the account ID. If this parameter is left empty, the Alibaba Cloud account ID of the caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace management page to view the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, which is the English identifier displayed in the workspace switcher at the top of the DataStudio page. You must specify either this parameter or ProjectId to determine the DataWorks workspace on which the API operation is performed.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The functional module to which the business process belongs. Valid values:
	//
	// - NORMAL: DataStudio.
	//
	// - MANUAL_BIZ: Manual business process.
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
