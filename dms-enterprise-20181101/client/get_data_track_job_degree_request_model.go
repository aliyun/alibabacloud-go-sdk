// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobDegreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataTrackJobDegreeRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataTrackJobDegreeRequest
	GetTid() *int64
}

type GetDataTrackJobDegreeRequest struct {
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 321****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataTrackJobDegreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobDegreeRequest) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobDegreeRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataTrackJobDegreeRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataTrackJobDegreeRequest) SetOrderId(v int64) *GetDataTrackJobDegreeRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataTrackJobDegreeRequest) SetTid(v int64) *GetDataTrackJobDegreeRequest {
	s.Tid = &v
	return s
}

func (s *GetDataTrackJobDegreeRequest) Validate() error {
	return dara.Validate(s)
}
