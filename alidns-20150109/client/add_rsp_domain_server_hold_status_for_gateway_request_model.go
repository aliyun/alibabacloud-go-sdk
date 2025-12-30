// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddRspDomainServerHoldStatusForGatewayRequest
	GetClientToken() *string
	SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayRequest
	GetDomainName() *string
	SetStatusMsg(v string) *AddRspDomainServerHoldStatusForGatewayRequest
	GetStatusMsg() *string
}

type AddRspDomainServerHoldStatusForGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// token123
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayRequest) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) SetClientToken(v string) *AddRspDomainServerHoldStatusForGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayRequest {
	s.DomainName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) SetStatusMsg(v string) *AddRspDomainServerHoldStatusForGatewayRequest {
	s.StatusMsg = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayRequest) Validate() error {
	return dara.Validate(s)
}
