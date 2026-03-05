// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationAccountDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryInspirationAccountDetailsRequest
	GetEndTime() *string
	SetOrderColumn(v string) *QueryInspirationAccountDetailsRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryInspirationAccountDetailsRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryInspirationAccountDetailsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryInspirationAccountDetailsRequest
	GetPageSize() *int32
	SetSourceType(v string) *QueryInspirationAccountDetailsRequest
	GetSourceType() *string
	SetStartTime(v string) *QueryInspirationAccountDetailsRequest
	GetStartTime() *string
}

type QueryInspirationAccountDetailsRequest struct {
	// example:
	//
	// 2025-07-23T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 2025-06-21T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryInspirationAccountDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryInspirationAccountDetailsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryInspirationAccountDetailsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryInspirationAccountDetailsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInspirationAccountDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInspirationAccountDetailsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryInspirationAccountDetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryInspirationAccountDetailsRequest) SetEndTime(v string) *QueryInspirationAccountDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetOrderColumn(v string) *QueryInspirationAccountDetailsRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetOrderType(v string) *QueryInspirationAccountDetailsRequest {
	s.OrderType = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetPageNum(v int32) *QueryInspirationAccountDetailsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetPageSize(v int32) *QueryInspirationAccountDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetSourceType(v string) *QueryInspirationAccountDetailsRequest {
	s.SourceType = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) SetStartTime(v string) *QueryInspirationAccountDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInspirationAccountDetailsRequest) Validate() error {
	return dara.Validate(s)
}
