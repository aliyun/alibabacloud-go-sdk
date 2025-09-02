// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionApplyOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetPermissionApplyOrderDetailRequest
	GetFlowId() *string
}

type GetPermissionApplyOrderDetailRequest struct {
	// The ID of the permission request order. You can call the [ListPermissionApplyOrders](https://help.aliyun.com/document_detail/211008.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 48f36729-05f9-4a40-9286-933fd940f30a
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetPermissionApplyOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetPermissionApplyOrderDetailRequest) SetFlowId(v string) *GetPermissionApplyOrderDetailRequest {
	s.FlowId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
