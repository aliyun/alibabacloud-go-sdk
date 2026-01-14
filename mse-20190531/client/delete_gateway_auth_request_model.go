// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayAuthRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayAuthRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *DeleteGatewayAuthRequest
	GetId() *int64
}

type DeleteGatewayAuthRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGatewayAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayAuthRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayAuthRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayAuthRequest) SetAcceptLanguage(v string) *DeleteGatewayAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayAuthRequest) SetGatewayUniqueId(v string) *DeleteGatewayAuthRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayAuthRequest) SetId(v int64) *DeleteGatewayAuthRequest {
	s.Id = &v
	return s
}

func (s *DeleteGatewayAuthRequest) Validate() error {
	return dara.Validate(s)
}
