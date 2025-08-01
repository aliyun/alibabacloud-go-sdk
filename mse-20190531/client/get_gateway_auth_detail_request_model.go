// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayAuthDetailRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *GetGatewayAuthDetailRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *GetGatewayAuthDetailRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *GetGatewayAuthDetailRequest
	GetId() *int64
}

type GetGatewayAuthDetailRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// 2274
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-6f0dbd108a0249d2b675b3ef50b*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 1100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGatewayAuthDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayAuthDetailRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayAuthDetailRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthDetailRequest) SetAcceptLanguage(v string) *GetGatewayAuthDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayAuthDetailRequest) SetGatewayId(v int64) *GetGatewayAuthDetailRequest {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayAuthDetailRequest) SetGatewayUniqueId(v string) *GetGatewayAuthDetailRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthDetailRequest) SetId(v int64) *GetGatewayAuthDetailRequest {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthDetailRequest) Validate() error {
	return dara.Validate(s)
}
