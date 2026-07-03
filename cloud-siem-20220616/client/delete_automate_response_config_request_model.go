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
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where the Data Management hub for threat analysis is located. Select the region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this ID to switch to the perspective of the member.
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
