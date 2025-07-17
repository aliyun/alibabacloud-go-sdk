// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderAttachmentFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetOrderAttachmentFileRequest
	GetOrderId() *int64
	SetTid(v int64) *GetOrderAttachmentFileRequest
	GetTid() *int64
}

type GetOrderAttachmentFileRequest struct {
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

func (s GetOrderAttachmentFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrderAttachmentFileRequest) GoString() string {
	return s.String()
}

func (s *GetOrderAttachmentFileRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetOrderAttachmentFileRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetOrderAttachmentFileRequest) SetOrderId(v int64) *GetOrderAttachmentFileRequest {
	s.OrderId = &v
	return s
}

func (s *GetOrderAttachmentFileRequest) SetTid(v int64) *GetOrderAttachmentFileRequest {
	s.Tid = &v
	return s
}

func (s *GetOrderAttachmentFileRequest) Validate() error {
	return dara.Validate(s)
}
