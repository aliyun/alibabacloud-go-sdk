// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderListQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyIdListShrink(v string) *IFlightOrderListQueryShrinkRequest
	GetApplyIdListShrink() *string
	SetBookTypeListShrink(v string) *IFlightOrderListQueryShrinkRequest
	GetBookTypeListShrink() *string
	SetBookerIdShrink(v string) *IFlightOrderListQueryShrinkRequest
	GetBookerIdShrink() *string
	SetEndDate(v string) *IFlightOrderListQueryShrinkRequest
	GetEndDate() *string
	SetPageSize(v int32) *IFlightOrderListQueryShrinkRequest
	GetPageSize() *int32
	SetScrollId(v string) *IFlightOrderListQueryShrinkRequest
	GetScrollId() *string
	SetStartDate(v string) *IFlightOrderListQueryShrinkRequest
	GetStartDate() *string
	SetThirdPartApplyIdListShrink(v string) *IFlightOrderListQueryShrinkRequest
	GetThirdPartApplyIdListShrink() *string
}

type IFlightOrderListQueryShrinkRequest struct {
	ApplyIdListShrink  *string `json:"apply_id_list,omitempty" xml:"apply_id_list,omitempty"`
	BookTypeListShrink *string `json:"book_type_list,omitempty" xml:"book_type_list,omitempty"`
	BookerIdShrink     *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4ZGViODFlYmM3MYzM4
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	StartDate                  *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	ThirdPartApplyIdListShrink *string `json:"third_part_apply_id_list,omitempty" xml:"third_part_apply_id_list,omitempty"`
}

func (s IFlightOrderListQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryShrinkRequest) GetApplyIdListShrink() *string {
	return s.ApplyIdListShrink
}

func (s *IFlightOrderListQueryShrinkRequest) GetBookTypeListShrink() *string {
	return s.BookTypeListShrink
}

func (s *IFlightOrderListQueryShrinkRequest) GetBookerIdShrink() *string {
	return s.BookerIdShrink
}

func (s *IFlightOrderListQueryShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *IFlightOrderListQueryShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IFlightOrderListQueryShrinkRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *IFlightOrderListQueryShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *IFlightOrderListQueryShrinkRequest) GetThirdPartApplyIdListShrink() *string {
	return s.ThirdPartApplyIdListShrink
}

func (s *IFlightOrderListQueryShrinkRequest) SetApplyIdListShrink(v string) *IFlightOrderListQueryShrinkRequest {
	s.ApplyIdListShrink = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetBookTypeListShrink(v string) *IFlightOrderListQueryShrinkRequest {
	s.BookTypeListShrink = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetBookerIdShrink(v string) *IFlightOrderListQueryShrinkRequest {
	s.BookerIdShrink = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetEndDate(v string) *IFlightOrderListQueryShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetPageSize(v int32) *IFlightOrderListQueryShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetScrollId(v string) *IFlightOrderListQueryShrinkRequest {
	s.ScrollId = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetStartDate(v string) *IFlightOrderListQueryShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) SetThirdPartApplyIdListShrink(v string) *IFlightOrderListQueryShrinkRequest {
	s.ThirdPartApplyIdListShrink = &v
	return s
}

func (s *IFlightOrderListQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
