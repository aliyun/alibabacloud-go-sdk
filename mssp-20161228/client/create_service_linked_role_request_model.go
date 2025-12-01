// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateServiceLinkedRoleRequest
	GetLang() *string
	SetRegionId(v string) *CreateServiceLinkedRoleRequest
	GetRegionId() *string
}

type CreateServiceLinkedRoleRequest struct {
	// Language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceLinkedRoleRequest) SetLang(v string) *CreateServiceLinkedRoleRequest {
	s.Lang = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
