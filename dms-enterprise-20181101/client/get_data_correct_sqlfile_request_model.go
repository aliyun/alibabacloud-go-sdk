// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectSQLFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetDataCorrectSQLFileRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCorrectSQLFileRequest
	GetTid() *int64
}

type GetDataCorrectSQLFileRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 730000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCorrectSQLFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectSQLFileRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCorrectSQLFileRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCorrectSQLFileRequest) SetOrderId(v int64) *GetDataCorrectSQLFileRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectSQLFileRequest) SetTid(v int64) *GetDataCorrectSQLFileRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectSQLFileRequest) Validate() error {
	return dara.Validate(s)
}
