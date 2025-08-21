// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCycle(v string) *CreateLoadBalancerRequest
	GetBillingCycle() *string
	SetClientToken(v string) *CreateLoadBalancerRequest
	GetClientToken() *string
	SetEnsRegionId(v string) *CreateLoadBalancerRequest
	GetEnsRegionId() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerRequest
	GetLoadBalancerName() *string
	SetLoadBalancerSpec(v string) *CreateLoadBalancerRequest
	GetLoadBalancerSpec() *string
	SetLoadBalancerType(v string) *CreateLoadBalancerRequest
	GetLoadBalancerType() *string
	SetNetworkId(v string) *CreateLoadBalancerRequest
	GetNetworkId() *string
	SetPayType(v string) *CreateLoadBalancerRequest
	GetPayType() *string
	SetVSwitchId(v string) *CreateLoadBalancerRequest
	GetVSwitchId() *string
}

type CreateLoadBalancerRequest struct {
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The client token that is used to ensure the idempotence of the request. This prevents repeated operations caused by multiple retries.
	//
	// 	- You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// 	- If you retry an API request with the same client token and request parameters after it has completed successfully, the result of the original request is returned. The server status does not change.
	//
	// 	- You can initiate a retry when the operation times out or the error code is PROCESSING. The idempotence is valid. If HTTP status code 200 is returned, the client receives the same result as the last request. However, your server status is not affected. If HTTP status code 4xx is returned and error code is not PROCESSING, the idempotence is invalid.
	//
	// 	- A client token is valid for 10 minutes.
	//
	// example:
	//
	// 26C28756-2586-17AF-B802-0DC50D8FDEBB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the ELB instance. The name must be 1 to 80 characters in length. If you leave this parameter empty, the system randomly allocates a name as the value of this parameter.
	//
	// >  The value cannot start with `http://` or `https://`.
	//
	// example:
	//
	// gcs-pre-websocket-eslb-telecom
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specification of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// elb.s2.medium
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The network ID of the created ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-5sax03dh2eyagujgsn7z9****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The billing method of the cluster. Valid value: PostPaid. PostPaid specifies the pay-as-you-go billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the vSwitch to which the internal-facing ELB instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-5s78haoys9oylle6ln71m****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *CreateLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLoadBalancerRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *CreateLoadBalancerRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateLoadBalancerRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateLoadBalancerRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLoadBalancerRequest) SetBillingCycle(v string) *CreateLoadBalancerRequest {
	s.BillingCycle = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetEnsRegionId(v string) *CreateLoadBalancerRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerSpec(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerType(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetNetworkId(v string) *CreateLoadBalancerRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetPayType(v string) *CreateLoadBalancerRequest {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVSwitchId(v string) *CreateLoadBalancerRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
