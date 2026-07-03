// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResponseRuleStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetResponseRuleStatisticRequest
	GetLang() *string
	SetRegionId(v string) *GetResponseRuleStatisticRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetResponseRuleStatisticRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *GetResponseRuleStatisticRequest
	GetRoleType() *int32
}

type GetResponseRuleStatisticRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: the Chinese mainland and Hong Kong (China)
	//
	// - ap-southeast-1: regions outside China
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud account ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the current Alibaba Cloud account view.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GetResponseRuleStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResponseRuleStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetResponseRuleStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *GetResponseRuleStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResponseRuleStatisticRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetResponseRuleStatisticRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *GetResponseRuleStatisticRequest) SetLang(v string) *GetResponseRuleStatisticRequest {
	s.Lang = &v
	return s
}

func (s *GetResponseRuleStatisticRequest) SetRegionId(v string) *GetResponseRuleStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *GetResponseRuleStatisticRequest) SetRoleFor(v int64) *GetResponseRuleStatisticRequest {
	s.RoleFor = &v
	return s
}

func (s *GetResponseRuleStatisticRequest) SetRoleType(v int32) *GetResponseRuleStatisticRequest {
	s.RoleType = &v
	return s
}

func (s *GetResponseRuleStatisticRequest) Validate() error {
	return dara.Validate(s)
}
