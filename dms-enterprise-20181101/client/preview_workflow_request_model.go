// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *PreviewWorkflowRequest
	GetOrderId() *int64
	SetTid(v int64) *PreviewWorkflowRequest
	GetTid() *int64
}

type PreviewWorkflowRequest struct {
	// The ticket ID. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1069****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s PreviewWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowRequest) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *PreviewWorkflowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *PreviewWorkflowRequest) SetOrderId(v int64) *PreviewWorkflowRequest {
	s.OrderId = &v
	return s
}

func (s *PreviewWorkflowRequest) SetTid(v int64) *PreviewWorkflowRequest {
	s.Tid = &v
	return s
}

func (s *PreviewWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
