// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeletePlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DeletePlaybookRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *DeletePlaybookRequest
	GetRoleFor() *int64
}

type DeletePlaybookRequest struct {
	// The language type for requests and received messages. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// UUID of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// User ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeletePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *DeletePlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DeletePlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeletePlaybookRequest) SetLang(v string) *DeletePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DeletePlaybookRequest) SetPlaybookUuid(v string) *DeletePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DeletePlaybookRequest) SetRoleFor(v int64) *DeletePlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *DeletePlaybookRequest) Validate() error {
	return dara.Validate(s)
}
