// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuid(v string) *GetAlertRequest
	GetAlertUuid() *string
	SetLang(v string) *GetAlertRequest
	GetLang() *string
	SetRegionId(v string) *GetAlertRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetAlertRequest
	GetRoleFor() *int64
	SetRoleType(v int64) *GetAlertRequest
	GetRoleType() *int64
}

type GetAlertRequest struct {
	// The ID of the alert that is associated with the incident.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region in which the threat detection and response data management center resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int64 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GetAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *GetAlertRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetAlertRequest) GetRoleType() *int64 {
	return s.RoleType
}

func (s *GetAlertRequest) SetAlertUuid(v string) *GetAlertRequest {
	s.AlertUuid = &v
	return s
}

func (s *GetAlertRequest) SetLang(v string) *GetAlertRequest {
	s.Lang = &v
	return s
}

func (s *GetAlertRequest) SetRegionId(v string) *GetAlertRequest {
	s.RegionId = &v
	return s
}

func (s *GetAlertRequest) SetRoleFor(v int64) *GetAlertRequest {
	s.RoleFor = &v
	return s
}

func (s *GetAlertRequest) SetRoleType(v int64) *GetAlertRequest {
	s.RoleType = &v
	return s
}

func (s *GetAlertRequest) Validate() error {
	return dara.Validate(s)
}
