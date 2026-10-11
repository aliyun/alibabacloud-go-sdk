// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEgressProxyConfigurationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEgressProxyConfigurationRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationId(v string) *DeleteEgressProxyConfigurationRequest
	GetEgressProxyConfigurationId() *string
}

type DeleteEgressProxyConfigurationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on the client side. The value must be unique across different requests. The client token can contain only ASCII characters.
	//
	// > Note If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: sends a check request and does not delete the configuration. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the check fails, an error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a regular request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// epc-hn302f6a9a284c4f****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
}

func (s DeleteEgressProxyConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEgressProxyConfigurationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEgressProxyConfigurationRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *DeleteEgressProxyConfigurationRequest) SetClientToken(v string) *DeleteEgressProxyConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEgressProxyConfigurationRequest) SetDryRun(v bool) *DeleteEgressProxyConfigurationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEgressProxyConfigurationRequest) SetEgressProxyConfigurationId(v string) *DeleteEgressProxyConfigurationRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *DeleteEgressProxyConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
