// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupProductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupProductionsRequest
	GetLang() *string
	SetRoleFor(v int64) *DescribeGroupProductionsRequest
	GetRoleFor() *int64
	SetRoleType(v string) *DescribeGroupProductionsRequest
	GetRoleType() *string
}

type DescribeGroupProductionsRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 118******150980
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- **0*	- (default): the view of the current Alibaba Cloud account.
	//
	// 	- **1**: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeGroupProductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupProductionsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeGroupProductionsRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeGroupProductionsRequest) SetLang(v string) *DescribeGroupProductionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupProductionsRequest) SetRoleFor(v int64) *DescribeGroupProductionsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeGroupProductionsRequest) SetRoleType(v string) *DescribeGroupProductionsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeGroupProductionsRequest) Validate() error {
	return dara.Validate(s)
}
