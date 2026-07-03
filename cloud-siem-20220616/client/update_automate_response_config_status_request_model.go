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
	// A JSON array of automated response rule IDs.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The enabling status of the rule. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The region where the Data Management center for threat analysis is deployed. You must select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member account that the administrator wants to access.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
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
