// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkOptimizationResponseBody
	GetRequestId() *string
}

type UpdateNetworkOptimizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkOptimizationResponseBody) SetRequestId(v string) *UpdateNetworkOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
