// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportDownloadURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataExportDownloadURLRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *GetDataExportDownloadURLRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *GetDataExportDownloadURLRequest
	GetTid() *int64
}

type GetDataExportDownloadURLRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 546****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the Alibaba Cloud account that is used to call the API operation.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataExportDownloadURLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportDownloadURLRequest) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataExportDownloadURLRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *GetDataExportDownloadURLRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataExportDownloadURLRequest) SetOrderId(v int64) *GetDataExportDownloadURLRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataExportDownloadURLRequest) SetRealLoginUserUid(v string) *GetDataExportDownloadURLRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *GetDataExportDownloadURLRequest) SetTid(v int64) *GetDataExportDownloadURLRequest {
	s.Tid = &v
	return s
}

func (s *GetDataExportDownloadURLRequest) Validate() error {
	return dara.Validate(s)
}
