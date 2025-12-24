// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNotifyTemplateListRequest
	GetLang() *string
	SetRoleFor(v int64) *DescribeNotifyTemplateListRequest
	GetRoleFor() *int64
	SetRoleType(v string) *DescribeNotifyTemplateListRequest
	GetRoleType() *string
}

type DescribeNotifyTemplateListRequest struct {
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 137602425xxx8726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Default value: 0. Valid values:
	//
	// 	- 0: the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeNotifyTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNotifyTemplateListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeNotifyTemplateListRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeNotifyTemplateListRequest) SetLang(v string) *DescribeNotifyTemplateListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNotifyTemplateListRequest) SetRoleFor(v int64) *DescribeNotifyTemplateListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeNotifyTemplateListRequest) SetRoleType(v string) *DescribeNotifyTemplateListRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeNotifyTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
