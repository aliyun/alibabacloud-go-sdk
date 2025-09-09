// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSlowSqlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDrdsSlowSqlsResponseBodyItems) *DescribeDrdsSlowSqlsResponseBody
	GetItems() *DescribeDrdsSlowSqlsResponseBodyItems
	SetPageNumber(v int32) *DescribeDrdsSlowSqlsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsSlowSqlsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDrdsSlowSqlsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsSlowSqlsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeDrdsSlowSqlsResponseBody
	GetTotal() *int32
}

type DescribeDrdsSlowSqlsResponseBody struct {
	// Indicates the details of the slow SQL query.
	Items *DescribeDrdsSlowSqlsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// Indicates the page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 509BDE17-505A-4B3B-854D-30D3F0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetItems() *DescribeDrdsSlowSqlsResponseBodyItems {
	return s.Items
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsSlowSqlsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetItems(v *DescribeDrdsSlowSqlsResponseBodyItems) *DescribeDrdsSlowSqlsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageNumber(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetPageSize(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetRequestId(v string) *DescribeDrdsSlowSqlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetSuccess(v bool) *DescribeDrdsSlowSqlsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) SetTotal(v int32) *DescribeDrdsSlowSqlsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsSlowSqlsResponseBodyItems struct {
	Item []*DescribeDrdsSlowSqlsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItems) GetItem() []*DescribeDrdsSlowSqlsResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeDrdsSlowSqlsResponseBodyItems) SetItem(v []*DescribeDrdsSlowSqlsResponseBodyItemsItem) *DescribeDrdsSlowSqlsResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsSlowSqlsResponseBodyItemsItem struct {
	// Indicates the IP address of the execution machine.
	//
	// example:
	//
	// 10.0.***.***
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Indicates the response time. Unit: ms.
	//
	// example:
	//
	// 1568267711
	ResponseTime *int64 `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// Indicates the name of the database.
	//
	// example:
	//
	// user
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates the time when the slow SQL query was sent. Unit: ms.
	//
	// example:
	//
	// 1568267711
	SendTime *int64 `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// Indicates the content of the slow SQL query.
	//
	// example:
	//
	// SELECT   count(1) from   payment_order where   order_status = \\"08\\";
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSlowSqlsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) GetHost() *string {
	return s.Host
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) GetSchema() *string {
	return s.Schema
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) GetSendTime() *int64 {
	return s.SendTime
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) GetSql() *string {
	return s.Sql
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetHost(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Host = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetResponseTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.ResponseTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSchema(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSendTime(v int64) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.SendTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) SetSql(v string) *DescribeDrdsSlowSqlsResponseBodyItemsItem {
	s.Sql = &v
	return s
}

func (s *DescribeDrdsSlowSqlsResponseBodyItemsItem) Validate() error {
	return dara.Validate(s)
}
