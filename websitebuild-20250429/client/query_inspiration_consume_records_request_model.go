// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationConsumeRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryInspirationConsumeRecordsRequest
	GetEndTime() *string
	SetOrderColumn(v string) *QueryInspirationConsumeRecordsRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryInspirationConsumeRecordsRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryInspirationConsumeRecordsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryInspirationConsumeRecordsRequest
	GetPageSize() *int32
	SetSceneName(v string) *QueryInspirationConsumeRecordsRequest
	GetSceneName() *string
	SetStartTime(v string) *QueryInspirationConsumeRecordsRequest
	GetStartTime() *string
}

type QueryInspirationConsumeRecordsRequest struct {
	// The end time of the query.
	//
	// example:
	//
	// 1762999521
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// 百灵鸟-人员认证
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2025-10-19T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryInspirationConsumeRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryInspirationConsumeRecordsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryInspirationConsumeRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryInspirationConsumeRecordsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInspirationConsumeRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInspirationConsumeRecordsRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryInspirationConsumeRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryInspirationConsumeRecordsRequest) SetEndTime(v string) *QueryInspirationConsumeRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetOrderColumn(v string) *QueryInspirationConsumeRecordsRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetOrderType(v string) *QueryInspirationConsumeRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetPageNum(v int32) *QueryInspirationConsumeRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetPageSize(v int32) *QueryInspirationConsumeRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetSceneName(v string) *QueryInspirationConsumeRecordsRequest {
	s.SceneName = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) SetStartTime(v string) *QueryInspirationConsumeRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInspirationConsumeRecordsRequest) Validate() error {
	return dara.Validate(s)
}
