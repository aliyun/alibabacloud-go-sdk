// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkOptimizationResponseBody
	GetRequestId() *string
}

type DeleteNetworkOptimizationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkOptimizationResponseBody) SetRequestId(v string) *DeleteNetworkOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
