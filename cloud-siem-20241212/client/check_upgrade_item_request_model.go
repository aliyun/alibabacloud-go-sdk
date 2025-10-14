// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUpgradeItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckUpgradeItemRequest
	GetLang() *string
	SetRegionId(v string) *CheckUpgradeItemRequest
	GetRegionId() *string
	SetRoleFor(v string) *CheckUpgradeItemRequest
	GetRoleFor() *string
	SetUpgradeItemId(v string) *CheckUpgradeItemRequest
	GetUpgradeItemId() *string
}

type CheckUpgradeItemRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// dispose_task_upgrade
	UpgradeItemId *string `json:"UpgradeItemId,omitempty" xml:"UpgradeItemId,omitempty"`
}

func (s CheckUpgradeItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUpgradeItemRequest) GoString() string {
	return s.String()
}

func (s *CheckUpgradeItemRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckUpgradeItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckUpgradeItemRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *CheckUpgradeItemRequest) GetUpgradeItemId() *string {
	return s.UpgradeItemId
}

func (s *CheckUpgradeItemRequest) SetLang(v string) *CheckUpgradeItemRequest {
	s.Lang = &v
	return s
}

func (s *CheckUpgradeItemRequest) SetRegionId(v string) *CheckUpgradeItemRequest {
	s.RegionId = &v
	return s
}

func (s *CheckUpgradeItemRequest) SetRoleFor(v string) *CheckUpgradeItemRequest {
	s.RoleFor = &v
	return s
}

func (s *CheckUpgradeItemRequest) SetUpgradeItemId(v string) *CheckUpgradeItemRequest {
	s.UpgradeItemId = &v
	return s
}

func (s *CheckUpgradeItemRequest) Validate() error {
	return dara.Validate(s)
}
