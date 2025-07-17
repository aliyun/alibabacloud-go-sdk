// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListOrdersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListOrdersResponseBody
	GetErrorMessage() *string
	SetOrders(v *ListOrdersResponseBodyOrders) *ListOrdersResponseBody
	GetOrders() *ListOrdersResponseBodyOrders
	SetRequestId(v string) *ListOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOrdersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListOrdersResponseBody
	GetTotalCount() *int64
}

type ListOrdersResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details about the tickets.
	Orders *ListOrdersResponseBodyOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOrdersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOrdersResponseBody) GetOrders() *ListOrdersResponseBodyOrders {
	return s.Orders
}

func (s *ListOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrdersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrdersResponseBody) SetErrorCode(v string) *ListOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrdersResponseBody) SetErrorMessage(v string) *ListOrdersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrdersResponseBody) SetOrders(v *ListOrdersResponseBodyOrders) *ListOrdersResponseBody {
	s.Orders = v
	return s
}

func (s *ListOrdersResponseBody) SetRequestId(v string) *ListOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrdersResponseBody) SetSuccess(v bool) *ListOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrdersResponseBody) SetTotalCount(v int64) *ListOrdersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOrdersResponseBodyOrders struct {
	Order []*ListOrdersResponseBodyOrdersOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s ListOrdersResponseBodyOrders) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponseBodyOrders) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrders) GetOrder() []*ListOrdersResponseBodyOrdersOrder {
	return s.Order
}

func (s *ListOrdersResponseBodyOrders) SetOrder(v []*ListOrdersResponseBodyOrdersOrder) *ListOrdersResponseBodyOrders {
	s.Order = v
	return s
}

func (s *ListOrdersResponseBodyOrders) Validate() error {
	return dara.Validate(s)
}

type ListOrdersResponseBodyOrdersOrder struct {
	// The remarks of the ticket.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The user who submitted the ticket.
	//
	// example:
	//
	// test
	Committer *string `json:"Committer,omitempty" xml:"Committer,omitempty"`
	// The ID of the user who submitted the ticket.
	//
	// example:
	//
	// 51****
	CommitterId *int64 `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// 2022-04-08 11:15:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the ticket was last modified.
	//
	// example:
	//
	// 2022-04-08 11:27:45
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The ID of the ticket.
	//
	// example:
	//
	// 51****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The type of the ticket.
	//
	// example:
	//
	// DC_COMMON
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The status code of the ticket. Valid values:
	//
	// 	- **fail**: The ticket fails to be executed.
	//
	// 	- **toaudit**: The ticket is waiting for approval.
	//
	// 	- **cancel**: The ticket is cancelled.
	//
	// 	- **processing**: The ticket is being executed.
	//
	// 	- **approved**: The ticket is approved.
	//
	// 	- **reject**: The ticket is rejected.
	//
	// 	- **success**: The ticket is executed.
	//
	// 	- **closed**: The ticket is closed.
	//
	// example:
	//
	// success
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The status description of the ticket.
	//
	// example:
	//
	// changed successfully
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s ListOrdersResponseBodyOrdersOrder) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponseBodyOrdersOrder) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrdersOrder) GetComment() *string {
	return s.Comment
}

func (s *ListOrdersResponseBodyOrdersOrder) GetCommitter() *string {
	return s.Committer
}

func (s *ListOrdersResponseBodyOrdersOrder) GetCommitterId() *int64 {
	return s.CommitterId
}

func (s *ListOrdersResponseBodyOrdersOrder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOrdersResponseBodyOrdersOrder) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *ListOrdersResponseBodyOrdersOrder) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListOrdersResponseBodyOrdersOrder) GetPluginType() *string {
	return s.PluginType
}

func (s *ListOrdersResponseBodyOrdersOrder) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListOrdersResponseBodyOrdersOrder) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListOrdersResponseBodyOrdersOrder) SetComment(v string) *ListOrdersResponseBodyOrdersOrder {
	s.Comment = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCommitter(v string) *ListOrdersResponseBodyOrdersOrder {
	s.Committer = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCommitterId(v int64) *ListOrdersResponseBodyOrdersOrder {
	s.CommitterId = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCreateTime(v string) *ListOrdersResponseBodyOrdersOrder {
	s.CreateTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetLastModifyTime(v string) *ListOrdersResponseBodyOrdersOrder {
	s.LastModifyTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetOrderId(v int64) *ListOrdersResponseBodyOrdersOrder {
	s.OrderId = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetPluginType(v string) *ListOrdersResponseBodyOrdersOrder {
	s.PluginType = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetStatusCode(v string) *ListOrdersResponseBodyOrdersOrder {
	s.StatusCode = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetStatusDesc(v string) *ListOrdersResponseBodyOrdersOrder {
	s.StatusDesc = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) Validate() error {
	return dara.Validate(s)
}
