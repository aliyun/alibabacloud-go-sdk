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
	// Name of the Business Process.<br>
	//
	// The name must be unique within the same project space.
	//
	// This parameter is required.
	//
	// example:
	//
	// My business process
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// Description of the Business Process.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud account ID of the owner responsible for the Business Process.<br>
	//
	// You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console), move the mouse pointer over the profile picture in the upper-right corner of the menu bar, and view the Account ID. If this parameter is empty, the Alibaba Cloud account ID of the invoker is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace.<br>
	//
	// You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console), go to the Workspace Management page, and view the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, which is the English identifier displayed when you switch workspaces at the top of the Data Development page. You must specify either this parameter or the projectid parameter to identify the DataWorks project for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// Function module to which the Business Process belongs. Valid values:
	//
	// - NORMAL (Data Development)
	//
	// - MANUAL_BIZ (manually triggered workflow)
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
