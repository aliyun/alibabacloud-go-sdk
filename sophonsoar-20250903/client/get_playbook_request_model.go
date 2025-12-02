// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetPlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *GetPlaybookRequest
	GetPlaybookUuid() *string
	SetPlaybookVersion(v string) *GetPlaybookRequest
	GetPlaybookVersion() *string
	SetPlaybookVersionType(v string) *GetPlaybookRequest
	GetPlaybookVersionType() *string
	SetRoleFor(v int64) *GetPlaybookRequest
	GetRoleFor() *int64
}

type GetPlaybookRequest struct {
	// The language type for requests and received messages.
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The version ID of the playbook, which is only effective when **PlaybookVersionType*	- is **History**.
	//
	// example:
	//
	// 36c9dcd6-****-4262-****-d508464ebd21
	PlaybookVersion *string `json:"PlaybookVersion,omitempty" xml:"PlaybookVersion,omitempty"`
	// The version type of the playbook, with the following values:
	//
	// - **Draft**: Editing state.
	//
	// - **Online**: Online version.
	//
	// - **History**: Historical version.
	//
	// example:
	//
	// History
	PlaybookVersionType *string `json:"PlaybookVersionType,omitempty" xml:"PlaybookVersionType,omitempty"`
	// The user ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlaybookRequest) GoString() string {
	return s.String()
}

func (s *GetPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *GetPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *GetPlaybookRequest) GetPlaybookVersion() *string {
	return s.PlaybookVersion
}

func (s *GetPlaybookRequest) GetPlaybookVersionType() *string {
	return s.PlaybookVersionType
}

func (s *GetPlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetPlaybookRequest) SetLang(v string) *GetPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *GetPlaybookRequest) SetPlaybookUuid(v string) *GetPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *GetPlaybookRequest) SetPlaybookVersion(v string) *GetPlaybookRequest {
	s.PlaybookVersion = &v
	return s
}

func (s *GetPlaybookRequest) SetPlaybookVersionType(v string) *GetPlaybookRequest {
	s.PlaybookVersionType = &v
	return s
}

func (s *GetPlaybookRequest) SetRoleFor(v int64) *GetPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *GetPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
