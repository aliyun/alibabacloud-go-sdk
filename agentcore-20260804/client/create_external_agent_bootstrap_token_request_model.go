// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentBootstrapTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkType(v string) *CreateExternalAgentBootstrapTokenRequest
	GetNetworkType() *string
}

type CreateExternalAgentBootstrapTokenRequest struct {
	// The network type for connection. Valid values:
	//
	// - INTERNET: public network
	//
	// - INTRANET: internal network
	//
	// This parameter is required.
	//
	// example:
	//
	// INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s CreateExternalAgentBootstrapTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentBootstrapTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentBootstrapTokenRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateExternalAgentBootstrapTokenRequest) SetNetworkType(v string) *CreateExternalAgentBootstrapTokenRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateExternalAgentBootstrapTokenRequest) Validate() error {
	return dara.Validate(s)
}
