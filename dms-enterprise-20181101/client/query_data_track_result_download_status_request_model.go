// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataTrackResultDownloadStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadKeyId(v string) *QueryDataTrackResultDownloadStatusRequest
	GetDownloadKeyId() *string
	SetOrderId(v int64) *QueryDataTrackResultDownloadStatusRequest
	GetOrderId() *int64
	SetTid(v int64) *QueryDataTrackResultDownloadStatusRequest
	GetTid() *int64
}

type QueryDataTrackResultDownloadStatusRequest struct {
	// The ID of the download key, which is used to identify the parsing progress of data tracking logs. You can call the DownloadDataTrackResult operation to query the ID of the key.
	//
	// This parameter is required.
	//
	// example:
	//
	// e23dd7ec-a19f-4a69-8eb3-8ffd26e6****
	DownloadKeyId *string `json:"DownloadKeyId,omitempty" xml:"DownloadKeyId,omitempty"`
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the ID of the tenant.
	//
	// example:
	//
	// 1***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s QueryDataTrackResultDownloadStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataTrackResultDownloadStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDataTrackResultDownloadStatusRequest) GetDownloadKeyId() *string {
	return s.DownloadKeyId
}

func (s *QueryDataTrackResultDownloadStatusRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *QueryDataTrackResultDownloadStatusRequest) GetTid() *int64 {
	return s.Tid
}

func (s *QueryDataTrackResultDownloadStatusRequest) SetDownloadKeyId(v string) *QueryDataTrackResultDownloadStatusRequest {
	s.DownloadKeyId = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusRequest) SetOrderId(v int64) *QueryDataTrackResultDownloadStatusRequest {
	s.OrderId = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusRequest) SetTid(v int64) *QueryDataTrackResultDownloadStatusRequest {
	s.Tid = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusRequest) Validate() error {
	return dara.Validate(s)
}
