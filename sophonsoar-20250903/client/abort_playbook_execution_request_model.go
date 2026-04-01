// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortPlaybookExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AbortPlaybookExecutionRequest
	GetLang() *string
	SetPlaybookExecutionUuid(v string) *AbortPlaybookExecutionRequest
	GetPlaybookExecutionUuid() *string
	SetPlaybookUuid(v string) *AbortPlaybookExecutionRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *AbortPlaybookExecutionRequest
	GetRoleFor() *int64
	SetRoleType(v string) *AbortPlaybookExecutionRequest
	GetRoleType() *string
}

type AbortPlaybookExecutionRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e776dab31-499b-4307-9248-xxxxxx
	PlaybookExecutionUuid *string `json:"PlaybookExecutionUuid,omitempty" xml:"PlaybookExecutionUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AbortPlaybookExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortPlaybookExecutionRequest) GoString() string {
	return s.String()
}

func (s *AbortPlaybookExecutionRequest) GetLang() *string {
	return s.Lang
}

func (s *AbortPlaybookExecutionRequest) GetPlaybookExecutionUuid() *string {
	return s.PlaybookExecutionUuid
}

func (s *AbortPlaybookExecutionRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *AbortPlaybookExecutionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *AbortPlaybookExecutionRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *AbortPlaybookExecutionRequest) SetLang(v string) *AbortPlaybookExecutionRequest {
	s.Lang = &v
	return s
}

func (s *AbortPlaybookExecutionRequest) SetPlaybookExecutionUuid(v string) *AbortPlaybookExecutionRequest {
	s.PlaybookExecutionUuid = &v
	return s
}

func (s *AbortPlaybookExecutionRequest) SetPlaybookUuid(v string) *AbortPlaybookExecutionRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *AbortPlaybookExecutionRequest) SetRoleFor(v int64) *AbortPlaybookExecutionRequest {
	s.RoleFor = &v
	return s
}

func (s *AbortPlaybookExecutionRequest) SetRoleType(v string) *AbortPlaybookExecutionRequest {
	s.RoleType = &v
	return s
}

func (s *AbortPlaybookExecutionRequest) Validate() error {
	return dara.Validate(s)
}
