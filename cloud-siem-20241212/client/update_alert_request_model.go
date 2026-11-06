// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertStatus(v string) *UpdateAlertRequest
	GetAlertStatus() *string
	SetAlertUuid(v string) *UpdateAlertRequest
	GetAlertUuid() *string
	SetLang(v string) *UpdateAlertRequest
	GetLang() *string
	SetRegionId(v string) *UpdateAlertRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateAlertRequest
	GetRoleFor() *int64
	SetRoleType(v int64) *UpdateAlertRequest
	GetRoleType() *int64
}

type UpdateAlertRequest struct {
	// The alert status.
	//
	// example:
	//
	// 1
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The ID of the alert associated with the event.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for threat analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
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

func (s UpdateAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRequest) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *UpdateAlertRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *UpdateAlertRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateAlertRequest) GetRoleType() *int64 {
	return s.RoleType
}

func (s *UpdateAlertRequest) SetAlertStatus(v string) *UpdateAlertRequest {
	s.AlertStatus = &v
	return s
}

func (s *UpdateAlertRequest) SetAlertUuid(v string) *UpdateAlertRequest {
	s.AlertUuid = &v
	return s
}

func (s *UpdateAlertRequest) SetLang(v string) *UpdateAlertRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAlertRequest) SetRegionId(v string) *UpdateAlertRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertRequest) SetRoleFor(v int64) *UpdateAlertRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateAlertRequest) SetRoleType(v int64) *UpdateAlertRequest {
	s.RoleType = &v
	return s
}

func (s *UpdateAlertRequest) Validate() error {
	return dara.Validate(s)
}
