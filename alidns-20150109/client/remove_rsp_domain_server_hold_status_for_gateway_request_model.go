// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest
	GetClientToken() *string
	SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest
	GetDomainName() *string
	SetStatusMsg(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest
	GetStatusMsg() *string
}

type RemoveRspDomainServerHoldStatusForGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// asdf
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// 实名认证通过，解除serverHold状态
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayRequest) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) SetClientToken(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest {
	s.DomainName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) SetStatusMsg(v string) *RemoveRspDomainServerHoldStatusForGatewayRequest {
	s.StatusMsg = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayRequest) Validate() error {
	return dara.Validate(s)
}
