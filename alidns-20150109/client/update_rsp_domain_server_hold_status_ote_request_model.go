// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerHoldStatusOteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetClientToken() *string
	SetDomainName(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetDomainName() *string
	SetOperatorId(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetOperatorId() *string
	SetOperatorType(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetOperatorType() *string
	SetServerHoldStatus(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetServerHoldStatus() *string
	SetStatusMsg(v string) *UpdateRspDomainServerHoldStatusOteRequest
	GetStatusMsg() *string
}

type UpdateRspDomainServerHoldStatusOteRequest struct {
	// example:
	//
	// 443F1A21-XXXX-55C4-93E1-FF020DF93D7B
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gatewayId001
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registryGateway
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// enable
	ServerHoldStatus *string `json:"ServerHoldStatus,omitempty" xml:"ServerHoldStatus,omitempty"`
	// This parameter is required.
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerHoldStatusOteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerHoldStatusOteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetServerHoldStatus() *string {
	return s.ServerHoldStatus
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetClientToken(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetDomainName(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetOperatorId(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetOperatorType(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.OperatorType = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetServerHoldStatus(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.ServerHoldStatus = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) SetStatusMsg(v string) *UpdateRspDomainServerHoldStatusOteRequest {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteRequest) Validate() error {
	return dara.Validate(s)
}
