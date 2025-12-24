// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CopyPlaybookRequest
	GetDescription() *string
	SetDisplayName(v string) *CopyPlaybookRequest
	GetDisplayName() *string
	SetLang(v string) *CopyPlaybookRequest
	GetLang() *string
	SetReleaseVersion(v string) *CopyPlaybookRequest
	GetReleaseVersion() *string
	SetRoleFor(v int64) *CopyPlaybookRequest
	GetRoleFor() *int64
	SetRoleType(v string) *CopyPlaybookRequest
	GetRoleType() *string
	SetSourcePlaybookUuid(v string) *CopyPlaybookRequest
	GetSourcePlaybookUuid() *string
}

type CopyPlaybookRequest struct {
	// The description of the playbook.
	//
	// example:
	//
	// playbook description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// playbook_xxx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese (default).
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The release version of the playbook that you want to copy. Default value: 0. Valid values:
	//
	// 	- \\-1: The version that is being edited.
	//
	// 	- 0: The latest online version of the current playbook.
	//
	// example:
	//
	// 0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 137602*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- **0*	- (default): the view of the current account.
	//
	// 	- **1**: the view of the global account.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The UUID of the playbook that you want to copy.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 94bc318c-****-4cba-****-801ccb0d739f
	SourcePlaybookUuid *string `json:"SourcePlaybookUuid,omitempty" xml:"SourcePlaybookUuid,omitempty"`
}

func (s CopyPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *CopyPlaybookRequest) GetDescription() *string {
	return s.Description
}

func (s *CopyPlaybookRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CopyPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *CopyPlaybookRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CopyPlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CopyPlaybookRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *CopyPlaybookRequest) GetSourcePlaybookUuid() *string {
	return s.SourcePlaybookUuid
}

func (s *CopyPlaybookRequest) SetDescription(v string) *CopyPlaybookRequest {
	s.Description = &v
	return s
}

func (s *CopyPlaybookRequest) SetDisplayName(v string) *CopyPlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *CopyPlaybookRequest) SetLang(v string) *CopyPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *CopyPlaybookRequest) SetReleaseVersion(v string) *CopyPlaybookRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CopyPlaybookRequest) SetRoleFor(v int64) *CopyPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *CopyPlaybookRequest) SetRoleType(v string) *CopyPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *CopyPlaybookRequest) SetSourcePlaybookUuid(v string) *CopyPlaybookRequest {
	s.SourcePlaybookUuid = &v
	return s
}

func (s *CopyPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
