// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterResponseBody
	GetSuccess() *bool
}

type UpdateClusterResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, possible values:
	//
	// - true: request succeeded
	//
	// - false: request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
