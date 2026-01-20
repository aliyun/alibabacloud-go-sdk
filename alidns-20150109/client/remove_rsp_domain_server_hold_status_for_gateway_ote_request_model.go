// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayOteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest
	GetClientToken() *string
	SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest
	GetDomainName() *string
	SetStatusMsg(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest
	GetStatusMsg() *string
}

type RemoveRspDomainServerHoldStatusForGatewayOteRequest struct {
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
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StatusMsg  *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteRequest) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) SetClientToken(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest {
	s.DomainName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) SetStatusMsg(v string) *RemoveRspDomainServerHoldStatusForGatewayOteRequest {
	s.StatusMsg = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteRequest) Validate() error {
	return dara.Validate(s)
}
