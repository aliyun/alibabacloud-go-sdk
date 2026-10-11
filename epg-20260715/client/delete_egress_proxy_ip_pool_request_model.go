// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyIpPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEgressProxyIpPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEgressProxyIpPoolRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *DeleteEgressProxyIpPoolRequest
	GetEgressProxyIpPoolId() *string
}

type DeleteEgressProxyIpPoolRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on the client side. Make sure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: The system sends a check request and does not delete the egress proxy IP address pool. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the check fails, the system returns the corresponding error. If the check succeeds, the system returns the error code DryRunOperation.
	//
	// - **false*	- (default): The system sends a normal request. If the check succeeds, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
}

func (s DeleteEgressProxyIpPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyIpPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyIpPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEgressProxyIpPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEgressProxyIpPoolRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *DeleteEgressProxyIpPoolRequest) SetClientToken(v string) *DeleteEgressProxyIpPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEgressProxyIpPoolRequest) SetDryRun(v bool) *DeleteEgressProxyIpPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEgressProxyIpPoolRequest) SetEgressProxyIpPoolId(v string) *DeleteEgressProxyIpPoolRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *DeleteEgressProxyIpPoolRequest) Validate() error {
	return dara.Validate(s)
}
