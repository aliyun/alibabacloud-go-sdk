// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAutoRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClusterAutoRenewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterAutoRenewResponseBody
	GetSuccess() *bool
}

type UpdateClusterAutoRenewResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3896A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterAutoRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAutoRenewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterAutoRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterAutoRenewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterAutoRenewResponseBody) SetRequestId(v string) *UpdateClusterAutoRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterAutoRenewResponseBody) SetSuccess(v bool) *UpdateClusterAutoRenewResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterAutoRenewResponseBody) Validate() error {
	return dara.Validate(s)
}
