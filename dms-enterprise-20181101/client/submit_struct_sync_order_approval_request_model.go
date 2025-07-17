// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStructSyncOrderApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *SubmitStructSyncOrderApprovalRequest
	GetOrderId() *int64
	SetTid(v int64) *SubmitStructSyncOrderApprovalRequest
	GetTid() *int64
}

type SubmitStructSyncOrderApprovalRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4324535
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SubmitStructSyncOrderApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalRequest) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SubmitStructSyncOrderApprovalRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SubmitStructSyncOrderApprovalRequest) SetOrderId(v int64) *SubmitStructSyncOrderApprovalRequest {
	s.OrderId = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalRequest) SetTid(v int64) *SubmitStructSyncOrderApprovalRequest {
	s.Tid = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalRequest) Validate() error {
	return dara.Validate(s)
}
