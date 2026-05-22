// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkOptimizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateNetworkOptimizationResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateNetworkOptimizationResponseBody
	GetRequestId() *string
}

type CreateNetworkOptimizationResponseBody struct {
	ConfigId  *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkOptimizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkOptimizationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateNetworkOptimizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkOptimizationResponseBody) SetConfigId(v int64) *CreateNetworkOptimizationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateNetworkOptimizationResponseBody) SetRequestId(v string) *CreateNetworkOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkOptimizationResponseBody) Validate() error {
	return dara.Validate(s)
}
