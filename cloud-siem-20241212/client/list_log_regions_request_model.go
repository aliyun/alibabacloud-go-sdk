// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListLogRegionsRequest
	GetLang() *string
	SetRoleFor(v int64) *ListLogRegionsRequest
	GetRoleFor() *int64
}

type ListLogRegionsRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListLogRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListLogRegionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListLogRegionsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListLogRegionsRequest) SetLang(v string) *ListLogRegionsRequest {
	s.Lang = &v
	return s
}

func (s *ListLogRegionsRequest) SetRoleFor(v int64) *ListLogRegionsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListLogRegionsRequest) Validate() error {
	return dara.Validate(s)
}
