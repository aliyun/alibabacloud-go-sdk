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
}

type TravelStandardListQueryRequest struct {
	// example:
	//
	// false
	FromGroup *bool `json:"from_group,omitempty" xml:"from_group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name,omitempty"`
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

func (s *TravelStandardListQueryRequest) Validate() error {
	return dara.Validate(s)
}
