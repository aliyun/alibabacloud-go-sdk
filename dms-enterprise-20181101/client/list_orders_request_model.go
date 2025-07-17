// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListOrdersRequest
	GetEndTime() *string
	SetOrderResultType(v string) *ListOrdersRequest
	GetOrderResultType() *string
	SetOrderStatus(v string) *ListOrdersRequest
	GetOrderStatus() *string
	SetPageNumber(v int32) *ListOrdersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOrdersRequest
	GetPageSize() *int32
	SetPluginType(v string) *ListOrdersRequest
	GetPluginType() *string
	SetSearchContent(v string) *ListOrdersRequest
	GetSearchContent() *string
	SetSearchDateType(v string) *ListOrdersRequest
	GetSearchDateType() *string
	SetStartTime(v string) *ListOrdersRequest
	GetStartTime() *string
	SetTid(v int64) *ListOrdersRequest
	GetTid() *int64
}

type ListOrdersRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2022-04-09 11:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The scope of the tickets that you want to query. Valid values:
	//
	// 	- **AS_ADMIN**: all tickets.
	//
	// 	- **AS_COMMITTER**: the tickets that are submitted by the current user.
	//
	// 	- **AS_HANDLER**: the tickets to be processed by the current user.
	//
	// 	- **AS_OWNER**: the tickets that are processed by the current user.
	//
	// 	- **AS_Related**: the tickets that are related to the current user.
	//
	// example:
	//
	// AS_ADMIN
	OrderResultType *string `json:"OrderResultType,omitempty" xml:"OrderResultType,omitempty"`
	// The status of the tickets that you want to query. Valid values:
	//
	// 	- **ALL**: queries the tickets of all statuses.
	//
	// 	- **FINISHED**: queries the tickets that are completed.
	//
	// 	- **RUNNING**: queries the tickets that are being processed.
	//
	// example:
	//
	// ALL
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the tickets that you want to query. For more information, see [PluginType parameter](https://help.aliyun.com/document_detail/429109.html).
	//
	// example:
	//
	// DC_COMMON
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The keyword that is used to query tickets.
	//
	// example:
	//
	// test
	SearchContent *string `json:"SearchContent,omitempty" xml:"SearchContent,omitempty"`
	// The time condition based on which you want to query tickets. Valid values:
	//
	// 	- **CREATE_TIME**: the time when a ticket was created.
	//
	// 	- **MODIFY_TIME**: the time when a ticket was last modified.
	//
	// example:
	//
	// CREATE_TIME
	SearchDateType *string `json:"SearchDateType,omitempty" xml:"SearchDateType,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2022-04-08 11:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3000
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListOrdersRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListOrdersRequest) GetOrderResultType() *string {
	return s.OrderResultType
}

func (s *ListOrdersRequest) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListOrdersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrdersRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *ListOrdersRequest) GetSearchContent() *string {
	return s.SearchContent
}

func (s *ListOrdersRequest) GetSearchDateType() *string {
	return s.SearchDateType
}

func (s *ListOrdersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOrdersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListOrdersRequest) SetEndTime(v string) *ListOrdersRequest {
	s.EndTime = &v
	return s
}

func (s *ListOrdersRequest) SetOrderResultType(v string) *ListOrdersRequest {
	s.OrderResultType = &v
	return s
}

func (s *ListOrdersRequest) SetOrderStatus(v string) *ListOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListOrdersRequest) SetPageNumber(v int32) *ListOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrdersRequest) SetPageSize(v int32) *ListOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrdersRequest) SetPluginType(v string) *ListOrdersRequest {
	s.PluginType = &v
	return s
}

func (s *ListOrdersRequest) SetSearchContent(v string) *ListOrdersRequest {
	s.SearchContent = &v
	return s
}

func (s *ListOrdersRequest) SetSearchDateType(v string) *ListOrdersRequest {
	s.SearchDateType = &v
	return s
}

func (s *ListOrdersRequest) SetStartTime(v string) *ListOrdersRequest {
	s.StartTime = &v
	return s
}

func (s *ListOrdersRequest) SetTid(v int64) *ListOrdersRequest {
	s.Tid = &v
	return s
}

func (s *ListOrdersRequest) Validate() error {
	return dara.Validate(s)
}
