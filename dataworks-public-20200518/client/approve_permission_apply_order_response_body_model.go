// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePermissionApplyOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproveSuccess(v bool) *ApprovePermissionApplyOrderResponseBody
	GetApproveSuccess() *bool
	SetRequestId(v string) *ApprovePermissionApplyOrderResponseBody
	GetRequestId() *string
}

type ApprovePermissionApplyOrderResponseBody struct {
	// Indicates whether the permission request order is processed.
	//
	// example:
	//
	// true
	ApproveSuccess *bool `json:"ApproveSuccess,omitempty" xml:"ApproveSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApprovePermissionApplyOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApprovePermissionApplyOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ApprovePermissionApplyOrderResponseBody) GetApproveSuccess() *bool {
	return s.ApproveSuccess
}

func (s *ApprovePermissionApplyOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApprovePermissionApplyOrderResponseBody) SetApproveSuccess(v bool) *ApprovePermissionApplyOrderResponseBody {
	s.ApproveSuccess = &v
	return s
}

func (s *ApprovePermissionApplyOrderResponseBody) SetRequestId(v string) *ApprovePermissionApplyOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApprovePermissionApplyOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
