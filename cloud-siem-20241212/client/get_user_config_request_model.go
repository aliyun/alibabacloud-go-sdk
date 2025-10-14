// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetUserConfigRequest
	GetLang() *string
	SetRegionId(v string) *GetUserConfigRequest
	GetRegionId() *string
	SetRoleFor(v string) *GetUserConfigRequest
	GetRoleFor() *string
}

type GetUserConfigRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigRequest) GoString() string {
	return s.String()
}

func (s *GetUserConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *GetUserConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserConfigRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *GetUserConfigRequest) SetLang(v string) *GetUserConfigRequest {
	s.Lang = &v
	return s
}

func (s *GetUserConfigRequest) SetRegionId(v string) *GetUserConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserConfigRequest) SetRoleFor(v string) *GetUserConfigRequest {
	s.RoleFor = &v
	return s
}

func (s *GetUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
