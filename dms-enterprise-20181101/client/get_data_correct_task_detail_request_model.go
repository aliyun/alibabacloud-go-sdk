// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataCorrectTaskDetailRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCorrectTaskDetailRequest
	GetTid() *int64
}

type GetDataCorrectTaskDetailRequest struct {
	// The ID of the ticket. You can call the [CreateDataCorrectOrder](https://help.aliyun.com/document_detail/208388.html), [CreateDataImportOrder](https://help.aliyun.com/document_detail/208387.html), or [CreateFreeLockCorrectOrder](https://help.aliyun.com/document_detail/208386.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12435523
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 14325432
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCorrectTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCorrectTaskDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCorrectTaskDetailRequest) SetOrderId(v int64) *GetDataCorrectTaskDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectTaskDetailRequest) SetTid(v int64) *GetDataCorrectTaskDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
