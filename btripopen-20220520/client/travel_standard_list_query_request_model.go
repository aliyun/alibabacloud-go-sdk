// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardListQueryRequest
	GetFromGroup() *bool
	SetPageNo(v int32) *TravelStandardListQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *TravelStandardListQueryRequest
	GetPageSize() *int32
	SetRuleName(v string) *TravelStandardListQueryRequest
	GetRuleName() *string
	SetUserId(v string) *TravelStandardListQueryRequest
	GetUserId() *string
}

type TravelStandardListQueryRequest struct {
	// Applicable to parent-subsidiary enterprises. Set this parameter to true to query the unified group travel standards. If left empty, the system returns the travel rules that are currently in effect for the enterprise.
	//
	// example:
	//
	// false
	FromGroup *bool `json:"from_group,omitempty" xml:"from_group,omitempty"`
	// The page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The name of the travel standard to search for.
	//
	// example:
	//
	// 普通员工规则
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name,omitempty"`
	// The user ID. Specify this parameter to query the travel standards bound to an employee.
	//
	// example:
	//
	// user_1234
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TravelStandardListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardListQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TravelStandardListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TravelStandardListQueryRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *TravelStandardListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *TravelStandardListQueryRequest) SetFromGroup(v bool) *TravelStandardListQueryRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardListQueryRequest) SetPageNo(v int32) *TravelStandardListQueryRequest {
	s.PageNo = &v
	return s
}

func (s *TravelStandardListQueryRequest) SetPageSize(v int32) *TravelStandardListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TravelStandardListQueryRequest) SetRuleName(v string) *TravelStandardListQueryRequest {
	s.RuleName = &v
	return s
}

func (s *TravelStandardListQueryRequest) SetUserId(v string) *TravelStandardListQueryRequest {
	s.UserId = &v
	return s
}

func (s *TravelStandardListQueryRequest) Validate() error {
	return dara.Validate(s)
}
