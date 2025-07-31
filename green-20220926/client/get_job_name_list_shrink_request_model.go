// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobNameListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetJobNameListShrinkRequest
	GetEndDate() *string
	SetQuery(v string) *GetJobNameListShrinkRequest
	GetQuery() *string
	SetRegionId(v string) *GetJobNameListShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *GetJobNameListShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *GetJobNameListShrinkRequest
	GetStartDate() *string
}

type GetJobNameListShrinkRequest struct {
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetJobNameListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobNameListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetJobNameListShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetJobNameListShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *GetJobNameListShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobNameListShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *GetJobNameListShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetJobNameListShrinkRequest) SetEndDate(v string) *GetJobNameListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetQuery(v string) *GetJobNameListShrinkRequest {
	s.Query = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetRegionId(v string) *GetJobNameListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetSortShrink(v string) *GetJobNameListShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetStartDate(v string) *GetJobNameListShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetJobNameListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
