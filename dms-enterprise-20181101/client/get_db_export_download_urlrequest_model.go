// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbExportDownloadURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDbExportDownloadURLRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDbExportDownloadURLRequest
	GetTid() *int64
}

type GetDbExportDownloadURLRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 73****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDbExportDownloadURLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDbExportDownloadURLRequest) GoString() string {
	return s.String()
}

func (s *GetDbExportDownloadURLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDbExportDownloadURLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDbExportDownloadURLRequest) SetOrderId(v int64) *GetDbExportDownloadURLRequest {
	s.OrderId = &v
	return s
}

func (s *GetDbExportDownloadURLRequest) SetTid(v int64) *GetDbExportDownloadURLRequest {
	s.Tid = &v
	return s
}

func (s *GetDbExportDownloadURLRequest) Validate() error {
	return dara.Validate(s)
}
