// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIFlightOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyIdList(v []*string) *IFlightOrderListQueryRequest
	GetApplyIdList() []*string
	SetBookTypeList(v []*int32) *IFlightOrderListQueryRequest
	GetBookTypeList() []*int32
	SetBookerId(v []*string) *IFlightOrderListQueryRequest
	GetBookerId() []*string
	SetEndDate(v string) *IFlightOrderListQueryRequest
	GetEndDate() *string
	SetPageSize(v int32) *IFlightOrderListQueryRequest
	GetPageSize() *int32
	SetScrollId(v string) *IFlightOrderListQueryRequest
	GetScrollId() *string
	SetStartDate(v string) *IFlightOrderListQueryRequest
	GetStartDate() *string
	SetThirdPartApplyIdList(v []*string) *IFlightOrderListQueryRequest
	GetThirdPartApplyIdList() []*string
}

type IFlightOrderListQueryRequest struct {
	ApplyIdList  []*string `json:"apply_id_list,omitempty" xml:"apply_id_list,omitempty" type:"Repeated"`
	BookTypeList []*int32  `json:"book_type_list,omitempty" xml:"book_type_list,omitempty" type:"Repeated"`
	BookerId     []*string `json:"booker_id,omitempty" xml:"booker_id,omitempty" type:"Repeated"`
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
	StartDate            *string   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	ThirdPartApplyIdList []*string `json:"third_part_apply_id_list,omitempty" xml:"third_part_apply_id_list,omitempty" type:"Repeated"`
}

func (s IFlightOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s IFlightOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *IFlightOrderListQueryRequest) GetApplyIdList() []*string {
	return s.ApplyIdList
}

func (s *IFlightOrderListQueryRequest) GetBookTypeList() []*int32 {
	return s.BookTypeList
}

func (s *IFlightOrderListQueryRequest) GetBookerId() []*string {
	return s.BookerId
}

func (s *IFlightOrderListQueryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *IFlightOrderListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IFlightOrderListQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *IFlightOrderListQueryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *IFlightOrderListQueryRequest) GetThirdPartApplyIdList() []*string {
	return s.ThirdPartApplyIdList
}

func (s *IFlightOrderListQueryRequest) SetApplyIdList(v []*string) *IFlightOrderListQueryRequest {
	s.ApplyIdList = v
	return s
}

func (s *IFlightOrderListQueryRequest) SetBookTypeList(v []*int32) *IFlightOrderListQueryRequest {
	s.BookTypeList = v
	return s
}

func (s *IFlightOrderListQueryRequest) SetBookerId(v []*string) *IFlightOrderListQueryRequest {
	s.BookerId = v
	return s
}

func (s *IFlightOrderListQueryRequest) SetEndDate(v string) *IFlightOrderListQueryRequest {
	s.EndDate = &v
	return s
}

func (s *IFlightOrderListQueryRequest) SetPageSize(v int32) *IFlightOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *IFlightOrderListQueryRequest) SetScrollId(v string) *IFlightOrderListQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *IFlightOrderListQueryRequest) SetStartDate(v string) *IFlightOrderListQueryRequest {
	s.StartDate = &v
	return s
}

func (s *IFlightOrderListQueryRequest) SetThirdPartApplyIdList(v []*string) *IFlightOrderListQueryRequest {
	s.ThirdPartApplyIdList = v
	return s
}

func (s *IFlightOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
