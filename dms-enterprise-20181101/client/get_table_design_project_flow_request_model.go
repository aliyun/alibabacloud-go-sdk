// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDesignProjectFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetTableDesignProjectFlowRequest
	GetOrderId() *int64
	SetTid(v int64) *GetTableDesignProjectFlowRequest
	GetTid() *int64
}

type GetTableDesignProjectFlowRequest struct {
	// The ID of the schema design ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTableDesignProjectFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectFlowRequest) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectFlowRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetTableDesignProjectFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetTableDesignProjectFlowRequest) SetOrderId(v int64) *GetTableDesignProjectFlowRequest {
	s.OrderId = &v
	return s
}

func (s *GetTableDesignProjectFlowRequest) SetTid(v int64) *GetTableDesignProjectFlowRequest {
	s.Tid = &v
	return s
}

func (s *GetTableDesignProjectFlowRequest) Validate() error {
	return dara.Validate(s)
}
