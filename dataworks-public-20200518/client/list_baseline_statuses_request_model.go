// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineTypes(v string) *ListBaselineStatusesRequest
	GetBaselineTypes() *string
	SetBizdate(v string) *ListBaselineStatusesRequest
	GetBizdate() *string
	SetFinishStatus(v string) *ListBaselineStatusesRequest
	GetFinishStatus() *string
	SetOwner(v string) *ListBaselineStatusesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListBaselineStatusesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBaselineStatusesRequest
	GetPageSize() *int32
	SetPriority(v string) *ListBaselineStatusesRequest
	GetPriority() *string
	SetSearchText(v string) *ListBaselineStatusesRequest
	GetSearchText() *string
	SetStatus(v string) *ListBaselineStatusesRequest
	GetStatus() *string
	SetTopicId(v int64) *ListBaselineStatusesRequest
	GetTopicId() *int64
}

type ListBaselineStatusesRequest struct {
	// The type of the baseline. Valid values: DAILY and HOURLY. Separate multiple values with commas (,).
	//
	// example:
	//
	// DAILY,HOURLY
	BaselineTypes *string `json:"BaselineTypes,omitempty" xml:"BaselineTypes,omitempty"`
	// The business date in UTC format (yyyy-MM-dd\\"T\\"HH:mm:ssZ).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-07T00:00:00+0800
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The completion status of the baseline instance. Valid values: UNFINISH and FINISH. Separate multiple values with commas (,).
	//
	// example:
	//
	// FINISH,UNFINISH
	FinishStatus *string `json:"FinishStatus,omitempty" xml:"FinishStatus,omitempty"`
	// The Alibaba Cloud UID of the baseline owner.
	//
	// example:
	//
	// 9527952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the baseline. Valid values: 1, 3, 5, 7, and 8. Separate multiple priorities with commas (,).
	//
	// example:
	//
	// 1,3,5,7,8
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The search keyword. You can search by baseline name or baseline ID. If you enter an hourly baseline ID, all instances of that hourly baseline are returned.
	//
	// example:
	//
	// keyword
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGROUS, and OVER. Separate multiple values with commas (,).
	//
	// example:
	//
	// SAFE,DANGROUS,OVER
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the related event.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ListBaselineStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineStatusesRequest) GoString() string {
	return s.String()
}

func (s *ListBaselineStatusesRequest) GetBaselineTypes() *string {
	return s.BaselineTypes
}

func (s *ListBaselineStatusesRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *ListBaselineStatusesRequest) GetFinishStatus() *string {
	return s.FinishStatus
}

func (s *ListBaselineStatusesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselineStatusesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBaselineStatusesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineStatusesRequest) GetPriority() *string {
	return s.Priority
}

func (s *ListBaselineStatusesRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListBaselineStatusesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBaselineStatusesRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListBaselineStatusesRequest) SetBaselineTypes(v string) *ListBaselineStatusesRequest {
	s.BaselineTypes = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetBizdate(v string) *ListBaselineStatusesRequest {
	s.Bizdate = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetFinishStatus(v string) *ListBaselineStatusesRequest {
	s.FinishStatus = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetOwner(v string) *ListBaselineStatusesRequest {
	s.Owner = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetPageNumber(v int32) *ListBaselineStatusesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetPageSize(v int32) *ListBaselineStatusesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetPriority(v string) *ListBaselineStatusesRequest {
	s.Priority = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetSearchText(v string) *ListBaselineStatusesRequest {
	s.SearchText = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetStatus(v string) *ListBaselineStatusesRequest {
	s.Status = &v
	return s
}

func (s *ListBaselineStatusesRequest) SetTopicId(v int64) *ListBaselineStatusesRequest {
	s.TopicId = &v
	return s
}

func (s *ListBaselineStatusesRequest) Validate() error {
	return dara.Validate(s)
}
