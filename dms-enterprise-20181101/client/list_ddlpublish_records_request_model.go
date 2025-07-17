// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDLPublishRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ListDDLPublishRecordsRequest
	GetOrderId() *int64
	SetTid(v int64) *ListDDLPublishRecordsRequest
	GetTid() *int64
}

type ListDDLPublishRecordsRequest struct {
	// The ID of the ticket.
	//
	// > You can create a schema design ticket in the Data Management (DMS) console. For more information, see [Design schemas](https://help.aliyun.com/document_detail/69711.html). You can also call the [CreateOrder](https://help.aliyun.com/document_detail/144649.html) operation to create a schema design ticket and obtain the ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3214325
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, log on to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDDLPublishRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDDLPublishRecordsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDDLPublishRecordsRequest) SetOrderId(v int64) *ListDDLPublishRecordsRequest {
	s.OrderId = &v
	return s
}

func (s *ListDDLPublishRecordsRequest) SetTid(v int64) *ListDDLPublishRecordsRequest {
	s.Tid = &v
	return s
}

func (s *ListDDLPublishRecordsRequest) Validate() error {
	return dara.Validate(s)
}
