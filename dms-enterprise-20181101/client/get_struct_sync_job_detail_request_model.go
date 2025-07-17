// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetStructSyncJobDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetStructSyncJobDetailRequest
	GetTid() *int64
}

type GetStructSyncJobDetailRequest struct {
	// The ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4324321
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetStructSyncJobDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetStructSyncJobDetailRequest) SetOrderId(v int64) *GetStructSyncJobDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncJobDetailRequest) SetTid(v int64) *GetStructSyncJobDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetStructSyncJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
