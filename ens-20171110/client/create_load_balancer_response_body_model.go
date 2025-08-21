// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerName() *string
	SetNetworkId(v string) *CreateLoadBalancerResponseBody
	GetNetworkId() *string
	SetRequestId(v string) *CreateLoadBalancerResponseBody
	GetRequestId() *string
	SetVSwitchId(v string) *CreateLoadBalancerResponseBody
	GetVSwitchId() *string
}

type CreateLoadBalancerResponseBody struct {
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5s7crik3yo3bp03gqrbp5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ELB instance.
	//
	// example:
	//
	// gcs-pre-websocket-****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5sax03dh2eyagujgsn7z9****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1ECC937A-AE0E-4626-BE51-DED1D6D1C888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the vSwitch to which the ELB instance belongs.
	//
	// example:
	//
	// vsw-5savh5ngxh8sbj14bu7n****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetNetworkId(v string) *CreateLoadBalancerResponseBody {
	s.NetworkId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetVSwitchId(v string) *CreateLoadBalancerResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
