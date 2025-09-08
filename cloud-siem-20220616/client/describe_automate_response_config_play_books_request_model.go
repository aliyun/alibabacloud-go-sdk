// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigPlayBooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoResponseType(v string) *DescribeAutomateResponseConfigPlayBooksRequest
	GetAutoResponseType() *string
	SetEntityType(v string) *DescribeAutomateResponseConfigPlayBooksRequest
	GetEntityType() *string
	SetRegionId(v string) *DescribeAutomateResponseConfigPlayBooksRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAutomateResponseConfigPlayBooksRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAutomateResponseConfigPlayBooksRequest
	GetRoleType() *int32
}

type DescribeAutomateResponseConfigPlayBooksRequest struct {
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The entity type of the playbook. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
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

func (s DescribeAutomateResponseConfigPlayBooksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) GetAutoResponseType() *string {
	return s.AutoResponseType
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetEntityType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRegionId(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) Validate() error {
	return dara.Validate(s)
}
