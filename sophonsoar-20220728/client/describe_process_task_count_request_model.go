// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTaskCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityUuidList(v []*string) *DescribeProcessTaskCountRequest
	GetEntityUuidList() []*string
	SetLang(v string) *DescribeProcessTaskCountRequest
	GetLang() *string
	SetRoleFor(v int64) *DescribeProcessTaskCountRequest
	GetRoleFor() *int64
	SetRoleType(v string) *DescribeProcessTaskCountRequest
	GetRoleType() *string
}

type DescribeProcessTaskCountRequest struct {
	// The UUIDs of the entities.
	//
	// This parameter is required.
	EntityUuidList []*string `json:"EntityUuidList,omitempty" xml:"EntityUuidList,omitempty" type:"Repeated"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member whose permissions you want to use to call the operation.
	//
	// example:
	//
	// 104739******259
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - **0**: The view of the current Alibaba Cloud account.
	//
	// - **1**: The view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProcessTaskCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountRequest) GetEntityUuidList() []*string {
	return s.EntityUuidList
}

func (s *DescribeProcessTaskCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeProcessTaskCountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeProcessTaskCountRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeProcessTaskCountRequest) SetEntityUuidList(v []*string) *DescribeProcessTaskCountRequest {
	s.EntityUuidList = v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetLang(v string) *DescribeProcessTaskCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetRoleFor(v int64) *DescribeProcessTaskCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetRoleType(v string) *DescribeProcessTaskCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeProcessTaskCountRequest) Validate() error {
	return dara.Validate(s)
}
