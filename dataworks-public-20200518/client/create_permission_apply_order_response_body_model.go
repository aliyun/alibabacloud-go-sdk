// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionApplyOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v []*string) *CreatePermissionApplyOrderResponseBody
	GetFlowId() []*string
	SetRequestId(v string) *CreatePermissionApplyOrderResponseBody
	GetRequestId() *string
}

type CreatePermissionApplyOrderResponseBody struct {
	// The request order ID.
	FlowId []*string `json:"FlowId,omitempty" xml:"FlowId,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePermissionApplyOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionApplyOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePermissionApplyOrderResponseBody) GetFlowId() []*string {
	return s.FlowId
}

func (s *CreatePermissionApplyOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePermissionApplyOrderResponseBody) SetFlowId(v []*string) *CreatePermissionApplyOrderResponseBody {
	s.FlowId = v
	return s
}

func (s *CreatePermissionApplyOrderResponseBody) SetRequestId(v string) *CreatePermissionApplyOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePermissionApplyOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
