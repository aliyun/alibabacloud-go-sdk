// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutomateResponseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteAutomateResponseConfigRequest
	GetId() *int64
	SetRegionId(v string) *DeleteAutomateResponseConfigRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteAutomateResponseConfigRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DeleteAutomateResponseConfigRequest
	GetRoleType() *int32
}

type DeleteAutomateResponseConfigRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
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

func (s DeleteAutomateResponseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteAutomateResponseConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAutomateResponseConfigRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteAutomateResponseConfigRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DeleteAutomateResponseConfigRequest) SetId(v int64) *DeleteAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRegionId(v string) *DeleteAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRoleFor(v int64) *DeleteAutomateResponseConfigRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRoleType(v int32) *DeleteAutomateResponseConfigRequest {
	s.RoleType = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) Validate() error {
	return dara.Validate(s)
}
