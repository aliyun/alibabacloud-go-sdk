// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOrderApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *SubmitOrderApprovalRequest
	GetOrderId() *int64
	SetRealLoginUserUid(v string) *SubmitOrderApprovalRequest
	GetRealLoginUserUid() *string
	SetTid(v int64) *SubmitOrderApprovalRequest
	GetTid() *int64
}

type SubmitOrderApprovalRequest struct {
	// The ID of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the Alibaba Cloud account that is used to call the API operation.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SubmitOrderApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitOrderApprovalRequest) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *SubmitOrderApprovalRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *SubmitOrderApprovalRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SubmitOrderApprovalRequest) SetOrderId(v int64) *SubmitOrderApprovalRequest {
	s.OrderId = &v
	return s
}

func (s *SubmitOrderApprovalRequest) SetRealLoginUserUid(v string) *SubmitOrderApprovalRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *SubmitOrderApprovalRequest) SetTid(v int64) *SubmitOrderApprovalRequest {
	s.Tid = &v
	return s
}

func (s *SubmitOrderApprovalRequest) Validate() error {
	return dara.Validate(s)
}
