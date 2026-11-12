// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutomateResponseConfigStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *UpdateAutomateResponseConfigStatusRequest
	GetIds() *string
	SetInUse(v bool) *UpdateAutomateResponseConfigStatusRequest
	GetInUse() *bool
	SetRegionId(v string) *UpdateAutomateResponseConfigStatusRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateAutomateResponseConfigStatusRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *UpdateAutomateResponseConfigStatusRequest
	GetRoleType() *int32
}

type UpdateAutomateResponseConfigStatusRequest struct {
	// The IDs of the automated response rules, formatted as a JSON array.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The enabling status of the rule. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The region where the data management center of Threat Analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
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
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusRequest) GetIds() *string {
	return s.Ids
}

func (s *UpdateAutomateResponseConfigStatusRequest) GetInUse() *bool {
	return s.InUse
}

func (s *UpdateAutomateResponseConfigStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAutomateResponseConfigStatusRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateAutomateResponseConfigStatusRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetIds(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.Ids = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetInUse(v bool) *UpdateAutomateResponseConfigStatusRequest {
	s.InUse = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRegionId(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRoleFor(v int64) *UpdateAutomateResponseConfigStatusRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRoleType(v int32) *UpdateAutomateResponseConfigStatusRequest {
	s.RoleType = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) Validate() error {
	return dara.Validate(s)
}
