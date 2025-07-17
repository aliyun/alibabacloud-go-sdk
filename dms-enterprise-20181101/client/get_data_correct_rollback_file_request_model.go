// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectRollbackFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataCorrectRollbackFileRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCorrectRollbackFileRequest
	GetTid() *int64
}

type GetDataCorrectRollbackFileRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51****
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

func (s GetDataCorrectRollbackFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectRollbackFileRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectRollbackFileRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCorrectRollbackFileRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCorrectRollbackFileRequest) SetOrderId(v int64) *GetDataCorrectRollbackFileRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectRollbackFileRequest) SetTid(v int64) *GetDataCorrectRollbackFileRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectRollbackFileRequest) Validate() error {
	return dara.Validate(s)
}
