// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayConfigRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayConfigRequest
	GetGatewayUniqueId() *string
}

type GetGatewayConfigRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-61f465fa2dd044069e2208c4912*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s GetGatewayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayConfigRequest) SetAcceptLanguage(v string) *GetGatewayConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayConfigRequest) SetGatewayUniqueId(v string) *GetGatewayConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayConfigRequest) Validate() error {
	return dara.Validate(s)
}
