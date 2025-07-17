// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearTaskDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataCronClearTaskDetailListRequest
	GetOrderId() *int64
	SetPageNumber(v int64) *GetDataCronClearTaskDetailListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetDataCronClearTaskDetailListRequest
	GetPageSize() *int64
	SetTid(v int64) *GetDataCronClearTaskDetailListRequest
	GetTid() *int64
}

type GetDataCronClearTaskDetailListRequest struct {
	// The ID of the ticket. You can query the ticket ID from the response parameters of the [CreateDataCronClearOrder](https://help.aliyun.com/document_detail/208385.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 432532
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 12345
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCronClearTaskDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearTaskDetailListRequest) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCronClearTaskDetailListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetDataCronClearTaskDetailListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDataCronClearTaskDetailListRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCronClearTaskDetailListRequest) SetOrderId(v int64) *GetDataCronClearTaskDetailListRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetPageNumber(v int64) *GetDataCronClearTaskDetailListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetPageSize(v int64) *GetDataCronClearTaskDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetTid(v int64) *GetDataCronClearTaskDetailListRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) Validate() error {
	return dara.Validate(s)
}
