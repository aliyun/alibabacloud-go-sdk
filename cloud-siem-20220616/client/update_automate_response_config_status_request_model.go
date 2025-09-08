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
	// The IDs of the automatic response rules. The value is a JSON array.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
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
