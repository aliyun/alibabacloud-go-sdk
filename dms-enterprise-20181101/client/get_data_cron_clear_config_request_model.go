// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataCronClearConfigRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCronClearConfigRequest
	GetTid() *int64
}

type GetDataCronClearConfigRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 420****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCronClearConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataCronClearConfigRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCronClearConfigRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCronClearConfigRequest) SetOrderId(v int64) *GetDataCronClearConfigRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCronClearConfigRequest) SetTid(v int64) *GetDataCronClearConfigRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCronClearConfigRequest) Validate() error {
	return dara.Validate(s)
}
