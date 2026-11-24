// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentDebugTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkType(v string) *CreateCustomAgentDebugTokenRequest
	GetNetworkType() *string
}

type CreateCustomAgentDebugTokenRequest struct {
	// The network type used for local debugging. INTERNET indicates public network access. INTRANET indicates access over a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s CreateCustomAgentDebugTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentDebugTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentDebugTokenRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateCustomAgentDebugTokenRequest) SetNetworkType(v string) *CreateCustomAgentDebugTokenRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateCustomAgentDebugTokenRequest) Validate() error {
	return dara.Validate(s)
}
