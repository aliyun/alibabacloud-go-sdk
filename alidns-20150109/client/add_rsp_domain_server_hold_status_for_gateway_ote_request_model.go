// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayOteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest
	GetClientToken() *string
	SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest
	GetDomainName() *string
	SetStatusMsg(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest
	GetStatusMsg() *string
}

type AddRspDomainServerHoldStatusForGatewayOteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qwoefasdf
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

func (s AddRspDomainServerHoldStatusForGatewayOteRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayOteRequest) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) SetClientToken(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest {
	s.ClientToken = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest {
	s.DomainName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) SetStatusMsg(v string) *AddRspDomainServerHoldStatusForGatewayOteRequest {
	s.StatusMsg = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteRequest) Validate() error {
	return dara.Validate(s)
}
