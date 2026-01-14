// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigEndpointProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigEndpointProbeResponseBody
	GetRequestId() *string
}

type ConfigEndpointProbeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigEndpointProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigEndpointProbeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigEndpointProbeResponseBody) SetRequestId(v string) *ConfigEndpointProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigEndpointProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
